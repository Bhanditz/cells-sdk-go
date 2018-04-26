/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

type JobsActionMessage struct {

	Event *ProtobufAny `json:"Event,omitempty"`

	Nodes []TreeNode `json:"Nodes,omitempty"`

	Users []IdmUser `json:"Users,omitempty"`

	Activities []ActivityObject `json:"Activities,omitempty"`

	OutputChain []JobsActionOutput `json:"OutputChain,omitempty"`
}