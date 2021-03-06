// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Intel Corporation

package module

import (
	"reflect"

	"github.com/pkg/errors"
	"gitlab.com/project-emco/core/emco-base/src/orchestrator/pkg/appcontext"
	"gitlab.com/project-emco/core/emco-base/src/orchestrator/pkg/infra/db"
	"gitlab.com/project-emco/core/emco-base/src/orchestrator/pkg/infra/logutils"
	"gitlab.com/project-emco/core/emco-base/src/orchestrator/pkg/state"
)

// CaCertManager exposes all the caCert functionalities
type CaCertManager interface {
	CreateCert(cert CaCert, failIfExists bool) (CaCert, bool, error)
	DeleteCert() error
	GetAllCert() ([]CaCert, error)
	GetCert() (CaCert, error)
}

// CaCertClient holds the client properties
type CaCertClient struct {
	dbInfo db.DbInfo
	dbKey  interface{}
}

// NewCertClients returns an instance of the CaCertClient which implements the Manager
func NewCaCertClient(dbKey interface{}) *CaCertClient {
	return &CaCertClient{
		dbInfo: db.DbInfo{
			StoreName: "resources",
			TagMeta:   "data",
			TagState:  "stateInfo"},
		dbKey: dbKey}
}

// CreateCert creates a caCert
func (c *CaCertClient) CreateCert(cert CaCert, failIfExists bool) (CaCert, bool, error) {
	certExists := false

	if cer, err := c.GetCert(); err == nil &&
		!reflect.DeepEqual(cer, CaCert{}) {
		certExists = true
	}

	if certExists &&
		failIfExists {
		return CaCert{}, certExists, errors.New("Certificate already exists")
	}

	if err := db.DBconn.Insert(c.dbInfo.StoreName, c.dbKey, nil, c.dbInfo.TagMeta, cert); err != nil {
		return CaCert{}, certExists, err
	}

	return cert, certExists, nil
}

// DeleteCert deletes a caCert
func (c *CaCertClient) DeleteCert() error {
	return db.DBconn.Remove(c.dbInfo.StoreName, c.dbKey)
}

// GetAllCert
func (c *CaCertClient) GetAllCert() ([]CaCert, error) {
	values, err := db.DBconn.Find(c.dbInfo.StoreName, c.dbKey, c.dbInfo.TagMeta)
	if err != nil {
		return []CaCert{}, err
	}

	var certs []CaCert
	for _, value := range values {
		cert := CaCert{}
		if err = db.DBconn.Unmarshal(value, &cert); err != nil {
			return []CaCert{}, err
		}
		certs = append(certs, cert)
	}

	return certs, nil
}

// GetCert returns the caCert
func (c *CaCertClient) GetCert() (CaCert, error) {
	value, err := db.DBconn.Find(c.dbInfo.StoreName, c.dbKey, c.dbInfo.TagMeta)
	if err != nil {
		return CaCert{}, err
	}

	if len(value) == 0 {
		return CaCert{}, errors.New("Certificate not found")
	}

	if value != nil {
		cert := CaCert{}
		if err = db.DBconn.Unmarshal(value[0], &cert); err != nil {
			return CaCert{}, err
		}
		return cert, nil
	}

	return CaCert{}, errors.New("Unknown Error")
}

// UpdateCert update the caCert
func (c *CaCertClient) UpdateCert(cert CaCert) error {
	return db.DBconn.Insert(c.dbInfo.StoreName, c.dbKey, nil, c.dbInfo.TagMeta, cert)
}

// VerifyStateBeforeDelete verifies a caCert can be deleted or not
func (c *CaCertClient) VerifyStateBeforeDelete(cert, lifecycle string) error {
	sc := NewStateClient(c.dbKey)
	stateInfo, err := sc.Get()
	if err != nil {
		return err
	}

	cState, err := state.GetCurrentStateFromStateInfo(stateInfo)
	if err != nil {
		return err
	}

	if cState == state.StateEnum.Instantiated ||
		cState == state.StateEnum.InstantiateStopped {
		err := errors.Errorf("%s must be terminated for CaCert %s before it can be deleted", lifecycle, cert)
		logutils.Error("",
			logutils.Fields{
				"Error": err.Error()})
		return err
	}

	if cState == state.StateEnum.Terminated ||
		cState == state.StateEnum.TerminateStopped {
		// verify that the appcontext has completed terminating
		ctxID := state.GetLastContextIdFromStateInfo(stateInfo)
		acStatus, err := state.GetAppContextStatus(ctxID)
		if err == nil &&
			!(acStatus.Status == appcontext.AppContextStatusEnum.Terminated ||
				acStatus.Status == appcontext.AppContextStatusEnum.TerminateFailed) {
			err := errors.Errorf("%s termination has not completed for CaCert %s", lifecycle, cert)
			logutils.Error("",
				logutils.Fields{
					"Error": err.Error()})
			return err
		}

		for _, id := range state.GetContextIdsFromStateInfo(stateInfo) {
			context, err := state.GetAppContextFromId(id)
			if err != nil {
				logutils.Error("Failed to get appContext from id",
					logutils.Fields{
						"ID":    id,
						"Error": err.Error()})
				return err
			}
			err = context.DeleteCompositeApp()
			if err != nil {
				logutils.Error("Failed to delete the appContext",
					logutils.Fields{
						"Error": err.Error()})
				return err
			}
		}
	}

	return nil
}

// VerifyStateBeforeUpdate verifies a caCert can be updated or not
func (c *CaCertClient) VerifyStateBeforeUpdate(cert, lifecycle string) error {
	sc := NewStateClient(c.dbKey)
	stateInfo, err := sc.Get()
	if err != nil {
		return err
	}

	cState, err := state.GetCurrentStateFromStateInfo(stateInfo)
	if err != nil {
		return err
	}

	if cState != state.StateEnum.Created {
		err := errors.Errorf("Failed to update the CaCert. %s for the CaCert %s is in %s state", lifecycle, cert, cState)
		logutils.Error("",
			logutils.Fields{
				"Error": err.Error()})
		return err
	}

	return nil
}
