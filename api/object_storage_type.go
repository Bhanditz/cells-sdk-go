/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

type ObjectStorageType string

// List of objectStorageType
const (
	LOCAL ObjectStorageType = "LOCAL"
	S3    ObjectStorageType = "S3"
	SMB   ObjectStorageType = "SMB"
)