// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2020-2021 Intel Corporation

package module

import (
	"context"
	"encoding/json"
	"strings"
	"sync"
	"time"

	pkgerrors "github.com/pkg/errors"
	"google.golang.org/grpc"

	"gitlab.com/project-emco/core/emco-base/src/orchestrator/common"
	"gitlab.com/project-emco/core/emco-base/src/orchestrator/pkg/appcontext"
	"gitlab.com/project-emco/core/emco-base/src/orchestrator/pkg/infra/db"
	log "gitlab.com/project-emco/core/emco-base/src/orchestrator/pkg/infra/logutils"
	"gitlab.com/project-emco/core/emco-base/src/orchestrator/pkg/infra/rpc"
	"gitlab.com/project-emco/core/emco-base/src/orchestrator/pkg/state"
	readynotifypb "gitlab.com/project-emco/core/emco-base/src/rsync/pkg/grpc/readynotify"
)

type RsyncInfo struct {
	RsyncName  string
	hostName   string
	portNumber int
}

var rsyncInfo RsyncInfo
var mutex = &sync.Mutex{}

type _testvars struct {
	UseGrpcMock       bool
	ReadyNotifyClient readynotifypb.ReadyNotifyClient
}

var Testvars _testvars

// InitRsyncClient initializes connections to the Resource Synchronizer service
func initRsyncClient() bool {
	if (RsyncInfo{}) == rsyncInfo {
		mutex.Lock()
		defer mutex.Unlock()
		log.Error("[ReadyNotify gRPC] RsyncInfo not set - InitRsyncClient failed", log.Fields{
			"Rsyncname":  rsyncInfo.RsyncName,
			"Hostname":   rsyncInfo.hostName,
			"PortNumber": rsyncInfo.portNumber,
		})
		return false
	}
	rpc.UpdateRpcConn(rsyncInfo.RsyncName, rsyncInfo.hostName, rsyncInfo.portNumber)
	return true
}

// NewRsyncInfo shall return a newly created RsyncInfo object
func NewRsyncInfo(rName, h string, pN int) RsyncInfo {
	mutex.Lock()
	defer mutex.Unlock()
	rsyncInfo = RsyncInfo{RsyncName: rName, hostName: h, portNumber: pN}
	return rsyncInfo

}

// InvokeReadyNotify will make a gRPC call to the resource synchronizer and
// will subscribe DCM to alerts from the rsync gRPC server ("ready-notify")
func InvokeReadyNotify(appContextID string) error {
	var rpcClient readynotifypb.ReadyNotifyClient
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Unit test helper code
	if Testvars.UseGrpcMock {
		rpcClient = Testvars.ReadyNotifyClient
		readynotifyReq := new(readynotifypb.Topic)
		// readynotifyReq.AppContext = appContextID
		rpcClient.Alert(ctx, readynotifyReq)
		return nil
	}

	conn := rpc.GetRpcConn(rsyncName)
	if conn == nil {
		initRsyncClient()
		conn = rpc.GetRpcConn(rsyncName)
		if conn == nil {
			log.Error("[ReadyNotify gRPC] connection error", log.Fields{"grpc-server": rsyncName})
			return pkgerrors.Errorf("[ReadyNotify gRPC] connection error. grpc-server[%v]", rsyncName)
		}
	}

	client := readynotifypb.NewReadyNotifyClient(conn)
	if client == nil {
		log.Error("[ReadyNotify gRPC] Couldn't create a gRPC client", log.Fields{})
		return pkgerrors.Errorf("[ReadyNotify gRPC] Couldn't create a gRPC client")
	}

	subscribe(client, appContextID)
	return nil
}

func processAlert(client readynotifypb.ReadyNotifyClient, stream readynotifypb.ReadyNotify_AlertClient) {
	var ac appcontext.AppContext
	var appContextID string
	var dcc *ClusterClient
	var lcmeta appcontext.CompositeAppMeta
	var project string
	var logicalCloud string

	allCertsReady := false
	for !allCertsReady {
		resp, err := stream.Recv()
		if err != nil {
			log.Error("[ReadyNotify gRPC] Failed to receive notification", log.Fields{"err": err})
			time.Sleep(5 * time.Second) // protect against potential deluge of errors in the for loop
			continue
		}

		appContextID = resp.AppContext
		log.Info("[ReadyNotify gRPC] Received alert from rsync", log.Fields{"appContextID": appContextID, "err": err})

		allCertsReady = checkAppContext(appContextID) // check whether all certificates have been issued

		// Abort path?
	}
	log.Info("[ReadyNotify gRPC] All Logical Cloud certificates have been issued", log.Fields{"appContextID": appContextID})

	// if this point is reached, it means all clusters' certificates have been issued,
	// so it's time for DCM to build all the L1 kubeconfigs and store them in CloudConfig

	// get logical cloud using context logicalcloud meta:
	_, err := ac.LoadAppContext(appContextID)
	if err != nil {
		log.Error("[ReadyNotify gRPC] Error getting Logical Cloud using AppContext ID", log.Fields{"err": err})
		return
	}
	lcmeta, err = ac.GetCompositeAppMeta()
	if err != nil {
		log.Error("[ReadyNotify gRPC] Couldn't get Logical Cloud using AppContext ID", log.Fields{"err": err})
		return
	}
	project = lcmeta.Project
	logicalCloud = lcmeta.LogicalCloud
	log.Info("[ReadyNotify gRPC] Project and Logical Cloud obtained", log.Fields{"project": project, "logicalCloud": logicalCloud})

	// Get all clusters of the Logical Cloud
	dcc = NewClusterClient() // in cluster.go
	clusterList, err := dcc.GetAllClusters(project, logicalCloud)
	if err != nil {
		log.Error("[ReadyNotify gRPC] Failed getting all clusters of Logical Cloud", log.Fields{"logicalCloud": logicalCloud, "project": project})
		return
	}
	for _, cluster := range clusterList {
		_, err = dcc.GetClusterConfig(project, logicalCloud, cluster.MetaData.Name)
		// discard kubeconfig returned because it's not needed here
		if err != nil {
			log.Error("[ReadyNotify gRPC] Generating kubeconfig or storing CloudConfig failed", log.Fields{"logicalCloud": logicalCloud, "project": project, "cluster": cluster.MetaData.Name, "error": err.Error()})
			return
		}
		log.Info("[ReadyNotify gRPC] Generated kubeconfig and created CloudConfig for cluster", log.Fields{"project": project, "logicalCloud": logicalCloud, "cluster": cluster.MetaData.Name})
		// if this point is reached, the kubeconfig is already stored in CloudConfig
	}
	log.Info("[ReadyNotify gRPC] All CloudConfigs for Logical Cloud have been created", log.Fields{"project": project, "logicalCloud": logicalCloud})

	if err != nil {
		return // error already logged
	}

	_ = unsubscribe(client, appContextID)
}

// Updates the State of an existing logical cloud, adding to the timestamped history of States
func addState(lcc *LogicalCloudClient, project, logicalCloud, cid, newState string) error {
	s, err := lcc.GetState(project, logicalCloud)
	if err != nil {
		return err
	}
	lckey := common.LogicalCloudKey{
		Project:          project,
		LogicalCloudName: logicalCloud,
	}
	lastRevision, err := state.GetLatestRevisionFromStateInfo(s)
	if err != nil {
		log.Error("Latest revision not found", log.Fields{})
		return err
	}
	// TODO: make atomic
	newRevision := lastRevision + 1

	a := state.ActionEntry{
		State:     newState,
		ContextId: cid,
		TimeStamp: time.Now(),
		Revision:  newRevision,
	}
	s.StatusContextId = cid
	s.Actions = append(s.Actions, a)

	err = db.DBconn.Insert(lcc.storeName, lckey, nil, lcc.tagState, s)
	if err != nil {
		log.Error("Error updating the state info of the LogicalCloud: ", log.Fields{"logicalCloud": logicalCloud})
		return err
	}
	return nil
}

func subscribe(client readynotifypb.ReadyNotifyClient, appContextID string) {
	stream, err := client.Alert(context.Background(), &readynotifypb.Topic{ClientName: "dcm", AppContext: appContextID}, grpc.WaitForReady(true))
	if err != nil {
		log.Error("[ReadyNotify gRPC] Failed to subscribe to alerts", log.Fields{"err": err, "appContextID": appContextID})
	}

	log.Info("[ReadyNotify gRPC] Subscribing to alerts about appcontext ID", log.Fields{"appContextID": appContextID})
	go processAlert(client, stream)
	stream.CloseSend()
}

func unsubscribe(client readynotifypb.ReadyNotifyClient, appContextID string) error {
	_, err := client.Unsubscribe(context.Background(), &readynotifypb.Topic{ClientName: "dcm", AppContext: appContextID})
	if err != nil {
		log.Error("[ReadyNotify gRPC] Failed to unsubscribe to alerts", log.Fields{"err": err, "appContextID": appContextID})
	}
	return err
}

// checkAppContext checks whether the LC from the provided appcontext has had all cluster certificates issued
func checkAppContext(appContextID string) bool {
	// Get the contextId from the label (id)
	var ac appcontext.AppContext
	_, err := ac.LoadAppContext(appContextID)
	if err != nil {
		log.Error("AppContext not found", log.Fields{"appContextID": appContextID})
		return false
	}

	appsOrder, err := ac.GetAppInstruction("order")
	if err != nil {
		return false
	}
	var appList map[string][]string
	json.Unmarshal([]byte(appsOrder.(string)), &appList)

	for _, app := range appList["apporder"] {
		clusterNames, err := ac.GetClusterNames(app)
		if err != nil {
			return false
		}
		// iterate over all clusters of appcontext
		for k := 0; k < len(clusterNames); k++ {
			gitOps, err := IsGitOpsCluster(clusterNames[k])
			if err != nil {
				return false
			}
			if gitOps {
				continue
			}
			chandle, err := ac.GetClusterHandle(app, clusterNames[k])
			if err != nil {
				log.Info("Error getting cluster handle", log.Fields{"cluster": clusterNames[k]})
				return false
			}
			// Get the handle for the cluster status object
			handle, err := ac.GetLevelHandle(chandle, "status")
			if err != nil {
				log.Error("Couldn't fetch the handle for the cluster status object", log.Fields{"chandle": chandle})
				return false
			}

			clusterStatus, err := ac.GetValue(handle)
			if err != nil {
				log.Error("Couldn't fetch cluster status from its handle", log.Fields{"handle": handle})
				return false
			}
			// detect if certificate has been issued - assumes K8s base64-encoded PEM certificate
			if strings.Contains(clusterStatus.(string), "certificate\":\"LS0t") {
				log.Info("Cluster status contains the certificate", log.Fields{"cluster": clusterNames[k]})
			} else {
				log.Info("Cluster status doesn't contain the certificate yet", log.Fields{"cluster": clusterNames[k]})
				return false
			}
		}
	}
	return true
}
