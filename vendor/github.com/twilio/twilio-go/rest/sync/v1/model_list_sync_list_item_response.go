/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSyncListItemResponse struct for ListSyncListItemResponse
type ListSyncListItemResponse struct {
	Items []SyncV1SyncListItem    `json:"items,omitempty"`
	Meta  ListServiceResponseMeta `json:"meta,omitempty"`
}
