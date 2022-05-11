/*
 * LogCom API
 *
 * LogCom Swagger documentation
 *
 * API version: 1.0.25
 * Contact: laborit@blutspende.de
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package logcomapi

type ModelHealthCheck struct {
	ApiVersion []string `json:"apiVersion,omitempty"`
	Service string `json:"service,omitempty"`
	Status string `json:"status,omitempty"`
}
