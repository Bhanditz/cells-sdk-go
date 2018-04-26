/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

type ServiceQuery struct {

	SubQueries []ProtobufAny `json:"SubQueries,omitempty"`

	Operation *ServiceOperationType `json:"Operation,omitempty"`

	ResourcePolicyQuery *ServiceResourcePolicyQuery `json:"ResourcePolicyQuery,omitempty"`

	Offset string `json:"Offset,omitempty"`

	Limit string `json:"Limit,omitempty"`

	GroupBy int32 `json:"groupBy,omitempty"`
}