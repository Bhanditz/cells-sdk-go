/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

type RestAclCollection struct {

	ACLs []IdmAcl `json:"ACLs,omitempty"`

	Total int32 `json:"Total,omitempty"`
}