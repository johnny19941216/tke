/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package monitor

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +genclient:onlyVerbs=create
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Metric defines the structure for querying monitoring data requests and results.
type Metric struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta

	// +optional
	Query MetricQuery
	// +optional
	JSONResult string
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MetricList is the whole list of all metrics.
type MetricList struct {
	metav1.TypeMeta
	// +optional
	metav1.ListMeta

	// List of metrics
	Items []Metric
}

type MetricQuery struct {
	Table string
	// +optional
	StartTime *int64
	// +optional
	EndTime    *int64
	Fields     []string
	Conditions []MetricQueryCondition
	// +optional
	OrderBy string
	// +optional
	Order   string
	GroupBy []string
	Limit   int32
	Offset  int32
}

type MetricQueryCondition struct {
	Key   string
	Expr  string
	Value string
}

// +genclient
// +genclient:nonNamespaced
// +genclient:skipVerbs=deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigMap holds configuration data for tke to consume.
type ConfigMap struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta

	// Data contains the configuration data.
	// Each key must consist of alphanumeric characters, '-', '_' or '.'.
	// Values with non-UTF-8 byte sequences must use the BinaryData field.
	// The keys stored in Data must not overlap with the keys in
	// the BinaryData field, this is enforced during validation process.
	// +optional
	Data map[string]string

	// BinaryData contains the binary data.
	// Each key must consist of alphanumeric characters, '-', '_' or '.'.
	// BinaryData can contain byte sequences that are not in the UTF-8 range.
	// The keys stored in BinaryData must not overlap with the ones in
	// the Data field, this is enforced during validation process.
	// +optional
	BinaryData map[string][]byte
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigMapList is a resource containing a list of ConfigMap objects.
type ConfigMapList struct {
	metav1.TypeMeta

	// +optional
	metav1.ListMeta

	// Items is the list of ConfigMaps.
	Items []ConfigMap
}

// +genclient
// +genclient:nonNamespaced
// +genclient:skipVerbs=deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Collector is a monitor component.
type Collector struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta

	// Spec defines the desired identities of clusters in this set.
	// +optional
	Spec CollectorSpec
	// +optional
	Status CollectorStatus
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CollectorList is the whole list of all collectors which owned by a tenant.
type CollectorList struct {
	metav1.TypeMeta
	// +optional
	metav1.ListMeta

	// List of Collector
	Items []Collector
}

// CollectorSpec describes the attributes on a Collector.
type CollectorSpec struct {
	TenantID    string
	ClusterName string
	// +optional
	Type CollectorType
	// Version is the components version.
	// +optional
	Version string
	// Storage is the remote address for collector when writing/reading outside of cluster.
	// +optional
	Storage CollectorStorage
	// NotifyWebhook is the address that alert messages send to, optional. If not set, a default webhook address "https://[notify-api-address]/webhook" will be used.
	// +optional
	NotifyWebhook string
}

// CollectorStatus is information about the current status of a Collector.
type CollectorStatus struct {
	// Version is the version of collector.
	// +optional
	Version string
	// Components is the components version such as node-exporter and alert manager.
	// +optional
	Components map[string]string
	// Phase is the current lifecycle phase of the helm of cluster.
	// +optional
	Phase CollectorPhase
	// Reason is a brief CamelCase string that describes any failure.
	// +optional
	Reason string
	// RetryCount is a int between 0 and 5 that describes the time of retrying initializing.
	// +optional
	RetryCount int32
	// LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.
	// +optional
	LastReInitializingTimestamp metav1.Time
}

// CollectorStorage is the remote write/read address for collector.
type CollectorStorage struct {
	WriteAddr []string
	ReadAddr  []string
}

// CollectorPhase defines the phase of collector constructor.
type CollectorPhase string

const (
	// CollectorPhaseInitializing means is wait initializing.
	CollectorPhaseInitializing CollectorPhase = "Initializing"
	// CollectorPhaseReinitializing means is reinitializing.
	CollectorPhaseReinitializing CollectorPhase = "Reinitializing"
	// CollectorPhaseChecking means is wait checking.
	CollectorPhaseChecking CollectorPhase = "Checking"
	// CollectorPhaseRunning means is running.
	CollectorPhaseRunning CollectorPhase = "Running"
	// CollectorPhaseUpgrading means is upgrading.
	CollectorPhaseUpgrading CollectorPhase = "Upgrading"
	// CollectorPhaseFailed means has been failed.
	CollectorPhaseFailed CollectorPhase = "Failed"
)

// CollectorType defines the type of collector.
type CollectorType string

const (
	// CollectorManaged means the collector managed by TKE.
	CollectorManaged CollectorType = "Managed"
	// CollectorImported means the prometheus installed by other.
	CollectorImportedPrometheus CollectorType = "ImportedPrometheus"
)

// +genclient
// +genclient:nonNamespaced
// +genclient:skipVerbs=deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlarmPolicy is a policy of alarm system.
type AlarmPolicy struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta

	// Spec defines the desired identities of alarm policies in this set.
	// +optional
	Spec AlarmPolicySpec
	// +optional
	Status AlarmPolicyStatus
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlarmPolicyList is the whole list of all alarm policies which owned by a tenant.
type AlarmPolicyList struct {
	metav1.TypeMeta
	// +optional
	metav1.ListMeta

	// List of alarm policies.
	Items []AlarmPolicy
}

// AlarmPolicySpec describes the attributes on an alarm policy.
type AlarmPolicySpec struct {
	TenantID    string
	ClusterName string
	Type        AlarmPolicyType
	// +patchMergeKey=metricName
	// +patchStrategy=merge
	Metrics          []AlarmMetric
	Objects          string
	ObjectsType      AlarmObjectsType
	StatisticsPeriod int64
	// +optional
	Namespace *string
	// +optional
	WorkloadType *WorkloadType
	// +optional
	// +patchStrategy=merge
	ReceiverGroups []string
	// +optional
	// +patchStrategy=merge
	Receivers []string
	// +optional
	// +patchMergeKey=templateName
	// +patchStrategy=merge
	NotifyWays []AlarmNotifyWay
}

// AlarmPolicyStatus is information about the current status of a AlarmPolicy.
type AlarmPolicyStatus struct {
	// +optional
	Phase AlarmPolicyPhase
	// The last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time
	// The reason for the condition's last transition.
	// +optional
	Reason string
	// A human readable message indicating details about the transition.
	// +optional
	Message string
}

// AlarmPolicyType defines the type of alarm policy.
type AlarmPolicyType string

const (
	// AlarmPolicyCluster indicates a cluster-wide alarm policy.
	AlarmPolicyCluster AlarmPolicyType = "Cluster"
	// AlarmPolicyNode indicates a node-wide alarm policy.
	AlarmPolicyNode AlarmPolicyType = "Node"
	// AlarmPolicyPod indicates a pod-wide alarm policy.
	AlarmPolicyPod AlarmPolicyType = "Pod"
)

type AlarmObjectsType string

const (
	AlarmObjectsAll  AlarmObjectsType = "All"
	AlarmObjectsPart AlarmObjectsType = "Part"
)

type AlarmMetric struct {
	Measurement    string
	MetricName     string
	ContinuePeriod int64
	// +optional
	DisplayName string
	// +optional
	Evaluator *AlarmEvaluator
	// +optional
	Unit string
}

type AlarmEvaluator struct {
	Type  string
	Value string
}

type WorkloadType string

const (
	WorkloadDeployment  WorkloadType = "Deployment"
	WorkloadDaemonset   WorkloadType = "Daemonset"
	WorkloadStatefulset WorkloadType = "Statefulset"
)

type AlarmNotifyWay struct {
	ChannelName  string
	TemplateName string
}

// AlarmPolicyPhase indicates the status of policy alarm in cluster.
type AlarmPolicyPhase string

// These are valid alarm policy status.
const (
	// AlarmPolicyPending indicates that the alarm policy has been declared, when
	// the alarm policy has not actually been created in the cluster.
	AlarmPolicyPending AlarmPolicyPhase = "Pending"
	// AlarmPolicyAvailable indicates the alarm policy is available.
	AlarmPolicyAvailable AlarmPolicyPhase = "Available"
	// AlarmPolicyFailed indicates that the alarm policy failed to be created in the
	// cluster or deleted in the cluster after it has been created.
	AlarmPolicyFailed AlarmPolicyPhase = "Failed"
	// AlarmPolicyTerminating means the alarm policy is undergoing graceful
	// termination.
	AlarmPolicyTerminating AlarmPolicyPhase = "Terminating"
)
