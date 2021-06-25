/*******************************************************************************
* IBM Cloud Kubernetes Service, 5737-D43
* (C) Copyright IBM Corp. 2017, 2021 All Rights Reserved.
*
* SPDX-License-Identifier: Apache2.0
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*******************************************************************************/

package ibm

import (
	"errors"
	"fmt"

	"k8s.io/klog/v2"

	apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
)

// CloudEventRecorder is the cloud event recorder data
type CloudEventRecorder struct {
	Name     string
	Recorder record.EventRecorder
}

// CloudEventReason describes the reason for the cloud event
type CloudEventReason string

const (
	// CloudLoadBalancerNormalEvent cloud event reason
	CloudLoadBalancerNormalEvent CloudEventReason = "CloudLoadBalancerNormalEvent"
	// CreatingCloudLoadBalancerFailed cloud event reason
	CreatingCloudLoadBalancerFailed CloudEventReason = "CreatingCloudLoadBalancerFailed"
	// UpdatingCloudLoadBalancerFailed cloud event reason
	UpdatingCloudLoadBalancerFailed CloudEventReason = "UpdatingCloudLoadBalancerFailed"
	// DeletingCloudLoadBalancerFailed cloud event reason
	DeletingCloudLoadBalancerFailed CloudEventReason = "DeletingCloudLoadBalancerFailed"
	// DeletingLoadBalancerPodFailed cloud event reason
	DeletingLoadBalancerPodFailed CloudEventReason = "DeletingLoadBalancerPodFailed"
	// GettingCloudLoadBalancerFailed cloud event reason
	GettingCloudLoadBalancerFailed CloudEventReason = "GettingCloudLoadBalancerFailed"
	// VerifyingCloudLoadBalancerFailed cloud event reason
	VerifyingCloudLoadBalancerFailed CloudEventReason = "VerifyingCloudLoadBalancerFailed"
	// MovingCloudLoadBalancerFailedLocalOnlyTraffic cloud event reason
	MovingCloudLoadBalancerFailedLocalOnlyTraffic CloudEventReason = "MovingCloudLoadBalancerFailedLocalOnlyTraffic"
	// CloudVPCLoadBalancerNormalEvent cloud event reason
	CloudVPCLoadBalancerNormalEvent CloudEventReason = "CloudVPCLoadBalancerNormalEvent"
	// CloudVPCLoadBalancerMaintenance cloud event reason
	CloudVPCLoadBalancerMaintenance CloudEventReason = "CloudVPCLoadBalancerMaintenance"
	// CloudVPCLoadBalancerFailed cloud event reason
	CloudVPCLoadBalancerFailed CloudEventReason = "CloudVPCLoadBalancerFailed"
	// CloudVPCLoadBalancerNotFound cloud event reason
	CloudVPCLoadBalancerNotFound CloudEventReason = "CloudVPCLoadBalancerNotFound"
)

// NewCloudEventRecorder returns a cloud event recorder.
func NewCloudEventRecorder(providerName string, kubeClient clientset.Interface) *CloudEventRecorder {
	return NewCloudEventRecorderV1(providerName, v1core.New(kubeClient.CoreV1().RESTClient()).Events(""))
}

// NewCloudEventRecorderV1 returns a cloud event recorder for v1 client
func NewCloudEventRecorderV1(providerName string, eventInterface v1core.EventInterface) *CloudEventRecorder {
	name := providerName + "-cloud-provider"
	broadcaster := record.NewBroadcaster()
	broadcaster.StartLogging(klog.Infof)
	broadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: eventInterface})
	eventRecorder := CloudEventRecorder{
		Name:     name,
		Recorder: broadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: name}),
	}
	return &eventRecorder
}

// LoadBalancerNormalEvent logs a load balancer service event
func (c *CloudEventRecorder) LoadBalancerNormalEvent(lbDeployment *apps.Deployment, lbService *v1.Service, reason CloudEventReason, eventMessage string) {
	message := fmt.Sprintf(
		"Event on cloud load balancer %v with associated deployment %v for service %v with UID %v: %v",
		GetCloudProviderLoadBalancerName(lbService),
		types.NamespacedName{Namespace: lbDeployment.ObjectMeta.Namespace, Name: lbDeployment.ObjectMeta.Name},
		types.NamespacedName{Namespace: lbService.ObjectMeta.Namespace, Name: lbService.ObjectMeta.Name},
		lbService.ObjectMeta.UID,
		eventMessage,
	)
	c.Recorder.Event(lbDeployment, v1.EventTypeNormal, fmt.Sprintf("%v", reason), message)
	c.Recorder.Event(lbService, v1.EventTypeNormal, fmt.Sprintf("%v", reason), message)
}

// LoadBalancerWarningEvent logs load balancer deployment and service warning
// events and returns an error representing the events.
func (c *CloudEventRecorder) LoadBalancerWarningEvent(lbDeployment *apps.Deployment, lbService *v1.Service, reason CloudEventReason, errorMessage string) error {
	message := fmt.Sprintf(
		"Error on cloud load balancer %v with associated deployment %v for service %v with UID %v: %v",
		GetCloudProviderLoadBalancerName(lbService),
		types.NamespacedName{Namespace: lbDeployment.ObjectMeta.Namespace, Name: lbDeployment.ObjectMeta.Name},
		types.NamespacedName{Namespace: lbService.ObjectMeta.Namespace, Name: lbService.ObjectMeta.Name},
		lbService.ObjectMeta.UID,
		errorMessage,
	)
	c.Recorder.Event(lbDeployment, v1.EventTypeWarning, fmt.Sprintf("%v", reason), message)
	c.Recorder.Event(lbService, v1.EventTypeWarning, fmt.Sprintf("%v", reason), message)
	return errors.New(message)
}

// getLoadBalancerPortableSubnetPossibleErrors parse the errors for the portable subnet configmap
// and create a message that can be appended to any failed loadbalancers event
func getLoadBalancerPortableSubnetPossibleErrors(portableSubnetVlanErrors map[string][]subnetConfigErrorField) string {
	type subnetConfigErrors struct {
		subnetConfigErrorField subnetConfigErrorField
		occurrences            int
	}
	errors := make(map[string]*subnetConfigErrors)
	errMsg := ""

	// Loop through each vlan
	for _, portableSubnetVlanError := range portableSubnetVlanErrors {
		// Loop through each subnet error in the vlan
		for _, portableSubnetError := range portableSubnetVlanError {
			if _, ok := errors[portableSubnetError.ErrorReasonCode]; !ok {
				errors[portableSubnetError.ErrorReasonCode] = &subnetConfigErrors{portableSubnetError, 1}
			} else {
				errors[portableSubnetError.ErrorReasonCode].occurrences++
			}
		}

		// Loop through the subnet errors and append the occurrences of the error
		for _, tempError := range errors {
			if errMsg != "" {
				errMsg += ", "
			}
			errMsg += fmt.Sprintf("[%s: %s - Number of Occurrences: %d.]", tempError.subnetConfigErrorField.ErrorReasonCode, tempError.subnetConfigErrorField.ErrorMessage, tempError.occurrences)
		}
	}

	if errMsg != "" {
		return lbPortableSubnetMessage + " " + errMsg + " " + lbDocTroubleshootMessage
	}
	return lbNoIPsMessage + " " + lbDocReferenceMessage
}

// LoadBalancerServiceWarningEvent logs a load balancer service warning
// event and returns an error representing the event.
func (c *CloudEventRecorder) LoadBalancerServiceWarningEvent(lbService *v1.Service, reason CloudEventReason, errorMessage string) error {
	message := fmt.Sprintf(
		"Error on cloud load balancer %v for service %v with UID %v: %v",
		GetCloudProviderLoadBalancerName(lbService),
		types.NamespacedName{Namespace: lbService.ObjectMeta.Namespace, Name: lbService.ObjectMeta.Name},
		lbService.ObjectMeta.UID,
		errorMessage,
	)
	c.Recorder.Event(lbService, v1.EventTypeWarning, fmt.Sprintf("%v", reason), message)
	return errors.New(message)
}

// VpcLoadBalancerServiceWarningEvent logs a VPC load balancer service warning
// event and returns an error representing the event.
func (c *CloudEventRecorder) VpcLoadBalancerServiceWarningEvent(lbService *v1.Service, reason CloudEventReason, lbName string, errorMessage string) error {
	message := fmt.Sprintf(
		"Error on cloud load balancer %v for service %v with UID %v: %v",
		lbName,
		types.NamespacedName{Namespace: lbService.ObjectMeta.Namespace, Name: lbService.ObjectMeta.Name},
		lbService.ObjectMeta.UID,
		errorMessage,
	)
	c.Recorder.Event(lbService, v1.EventTypeWarning, fmt.Sprintf("%v", reason), message)
	return errors.New(message)
}

// VpcLoadBalancerServiceNormalEvent logs a VPC load balancer service event
func (c *CloudEventRecorder) VpcLoadBalancerServiceNormalEvent(lbService *v1.Service, reason CloudEventReason, lbName string, eventMessage string) {
	message := fmt.Sprintf(
		"Event on cloud load balancer %v for service %v with UID %v: %v",
		lbName,
		types.NamespacedName{Namespace: lbService.ObjectMeta.Namespace, Name: lbService.ObjectMeta.Name},
		lbService.ObjectMeta.UID,
		eventMessage,
	)
	c.Recorder.Event(lbService, v1.EventTypeNormal, fmt.Sprintf("%v", reason), message)
}
