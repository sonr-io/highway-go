/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListModelBuildResponse struct for ListModelBuildResponse
type ListModelBuildResponse struct {
	Meta        ListAssistantResponseMeta `json:"meta,omitempty"`
	ModelBuilds []AutopilotV1ModelBuild   `json:"model_builds,omitempty"`
}
