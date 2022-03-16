/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListUserConversationResponse struct for ListUserConversationResponse
type ListUserConversationResponse struct {
	Conversations []ConversationsV1UserConversation    `json:"conversations,omitempty"`
	Meta          ListConfigurationAddressResponseMeta `json:"meta,omitempty"`
}
