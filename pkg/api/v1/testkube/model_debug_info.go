/*
 * Testkube API
 *
 * Testkube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

// Testkube debug info
type DebugInfo struct {
	ClusterVersion string   `json:"clusterVersion,omitempty"`
	ApiLogs        []string `json:"apiLogs,omitempty"`
	OperatorLogs   []string `json:"operatorLogs,omitempty"`
	ExecutionLogs  []string `json:"executionLogs,omitempty"`
}