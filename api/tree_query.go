/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

type TreeQuery struct {

	PathPrefix []string `json:"PathPrefix,omitempty"`

	MinSize string `json:"MinSize,omitempty"`

	MaxSize string `json:"MaxSize,omitempty"`

	MinDate string `json:"MinDate,omitempty"`

	MaxDate string `json:"MaxDate,omitempty"`

	Type_ *TreeNodeType `json:"Type,omitempty"`

	FileName string `json:"FileName,omitempty"`

	Content string `json:"Content,omitempty"`

	FreeString string `json:"FreeString,omitempty"`

	Extension string `json:"Extension,omitempty"`

	GeoQuery *TreeGeoQuery `json:"GeoQuery,omitempty"`
}