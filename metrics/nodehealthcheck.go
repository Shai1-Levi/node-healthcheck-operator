/*
Copyright 2020 The Machine API Operator authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"

	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	// NodeHealthCheckOldRemediationCR is a Prometheus metric, which reports the number of old Remediation CRs.
	// It is an indication for remediation that is pending for a long while, which might indicate a problem with the external remediation mechanism.
	NodeHealthCheckOldRemediationCR = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "nodehealthcheck_old_remediation_cr",
			Help: "Number of old remediation CRs detected by NodeHealthChecks",
		}, []string{"name", "namespace"},
	)
)

var (
	// NodeHealthCheckOngoingRemediation is a Prometheus metric, which reports an ongoing remediation of a node
	NodeHealthCheckOngoingRemediation = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "nodehealthcheck_ongoing_remediation",
			Help: "Indication of an ongoing remediation of an unhealthy node",
		}, []string{"name", "namespace", "remediation"},
	)
)

func InitializeNodeHealthCheckMetrics() {
	metrics.Registry.MustRegister(
		NodeHealthCheckOldRemediationCR,
		NodeHealthCheckOngoingRemediation,
	)
}

func ObserveNodeHealthCheckOldRemediationCR(name, namespace string) {
	NodeHealthCheckOldRemediationCR.With(prometheus.Labels{
		"name":      name,
		"namespace": namespace,
	}).Inc()
}

func ObserveNodeHealthCheckRemediationCreated(name, namespace, remediation string) {
	NodeHealthCheckOngoingRemediation.With(prometheus.Labels{
		"name":        name,
		"namespace":   namespace,
		"remediation": remediation,
	}).Set(1)
}

func ObserveNodeHealthCheckRemediationDeleted(name, namespace, remediation string) {
	NodeHealthCheckOngoingRemediation.With(prometheus.Labels{
		"name":        name,
		"namespace":   namespace,
		"remediation": remediation,
	}).Set(0)
}
