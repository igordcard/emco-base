// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2020-22 Intel Corporation

package subresources

import k8scorev1 "k8s.io/api/core/v1"

// The ApprovalSubresource type defines the 4 necessary parameters
// that the "approval" subresource of a CertificateSigningRequest in K8s
// requires, in a forma tto be exchanged over AppContext
type ApprovalSubresource struct {
	LastUpdateTime string                    `json:"lastUpdateTime,omitempty"`
	Message        string                    `json:"message,omitempty"`
	Reason         string                    `json:"reason,omitempty"`
	Type           string                    `json:"type,omitempty"`
	Status         k8scorev1.ConditionStatus `json:"status" protobuf:"bytes,6,opt,name=status,casttype=k8s.io/api/core/v1.ConditionStatus"`
}
