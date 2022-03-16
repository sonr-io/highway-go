/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListCommandResponse struct for ListCommandResponse
type ListCommandResponse struct {
	Commands []WirelessV1Command     `json:"commands,omitempty"`
	Meta     ListCommandResponseMeta `json:"meta,omitempty"`
}
