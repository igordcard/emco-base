// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Intel Corporation

package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/project-emco/core/emco-base/src/ca-certs/pkg/client/logicalcloud"
	"gitlab.com/project-emco/core/emco-base/src/ca-certs/pkg/module"
	"gitlab.com/project-emco/core/emco-base/src/orchestrator/pkg/infra/apierror"
	"gitlab.com/project-emco/core/emco-base/src/orchestrator/pkg/infra/logutils"
)

// lcCertHandler implements the logicalCloud caCert handler functions
type lcCertHandler struct {
	manager logicalcloud.CaCertManager
}

// handleCertificateCreate handles the route for creating a new caCert
func (h *lcCertHandler) handleCertificateCreate(w http.ResponseWriter, r *http.Request) {
	h.createOrUpdateCertificate(w, r)
}

// handleCertificateDelete handles the route for deleting a caCert
func (h *lcCertHandler) handleCertificateDelete(w http.ResponseWriter, r *http.Request) {
	// get the route variables
	vars := _lcVars(mux.Vars(r))
	if err := h.manager.DeleteCert(vars.cert, vars.project); err != nil {
		apiErr := apierror.HandleErrors(mux.Vars(r), err, nil, apiErrors)
		http.Error(w, apiErr.Message, apiErr.Status)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// handleCertificateGet handles the route for retrieving a caCert
func (h *lcCertHandler) handleCertificateGet(w http.ResponseWriter, r *http.Request) {
	var (
		certs interface{}
		err   error
	)

	// get the route variables
	vars := _lcVars(mux.Vars(r))
	if len(vars.cert) == 0 {
		certs, err = h.manager.GetAllCert(vars.project)
	} else {
		certs, err = h.manager.GetCert(vars.cert, vars.project)
	}

	if err != nil {
		apiErr := apierror.HandleErrors(mux.Vars(r), err, nil, apiErrors)
		http.Error(w, apiErr.Message, apiErr.Status)
		return
	}

	sendResponse(w, certs, http.StatusOK)
}

// handleCertificateUpdate handles the route for updating a caCert
func (h *lcCertHandler) handleCertificateUpdate(w http.ResponseWriter, r *http.Request) {
	h.createOrUpdateCertificate(w, r)
}

// createOrUpdateCertificate create/update the caCert based on the request method
func (h *lcCertHandler) createOrUpdateCertificate(w http.ResponseWriter, r *http.Request) {
	var cert module.CaCert
	if code, err := validateRequestBody(r.Body, &cert, CertificateSchemaJson); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	// get the route variables
	vars := _lcVars(mux.Vars(r))

	methodPost := false
	if r.Method == http.MethodPost {
		methodPost = true
	}

	if !methodPost {
		// name in the URL should match the name in the body
		if cert.MetaData.Name != vars.cert {
			err := "caCert name is not matching with the name in the request"
			logutils.Error(err,
				logutils.Fields{
					"Cert":     cert,
					"CertName": vars.cert})
			http.Error(w, err, http.StatusBadRequest)
			return
		}
	}

	crt, certExists, err := h.manager.CreateCert(cert, vars.project, methodPost)
	if err != nil {
		apiErr := apierror.HandleErrors(mux.Vars(r), err, cert, apiErrors)
		http.Error(w, apiErr.Message, apiErr.Status)
		return
	}

	code := http.StatusCreated
	if certExists {
		// caCert does have a current representation and that representation is successfully modified
		code = http.StatusOK
	}

	sendResponse(w, crt, code)
}
