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

import (
	"time"
)

// ConversationsV1ServiceParticipantConversation struct for ConversationsV1ServiceParticipantConversation
type ConversationsV1ServiceParticipantConversation struct {
	// The unique ID of the Account responsible for this conversation.
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique ID of the Conversation Service this conversation belongs to.
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
	// An optional string metadata field you can use to store any data you wish.
	ConversationAttributes *string `json:"conversation_attributes,omitempty"`
	// Creator of this conversation.
	ConversationCreatedBy *string `json:"conversation_created_by,omitempty"`
	// The date that this conversation was created.
	ConversationDateCreated *time.Time `json:"conversation_date_created,omitempty"`
	// The date that this conversation was last updated.
	ConversationDateUpdated *time.Time `json:"conversation_date_updated,omitempty"`
	// The human-readable name of this conversation.
	ConversationFriendlyName *string `json:"conversation_friendly_name,omitempty"`
	// The unique ID of the Conversation this Participant belongs to.
	ConversationSid *string `json:"conversation_sid,omitempty"`
	// The current state of this User Conversation
	ConversationState *string `json:"conversation_state,omitempty"`
	// Timer date values for this conversation.
	ConversationTimers *map[string]interface{} `json:"conversation_timers,omitempty"`
	// An application-defined string that uniquely identifies the Conversation resource.
	ConversationUniqueName *string `json:"conversation_unique_name,omitempty"`
	// Absolute URLs to access the participant and conversation of this Participant Conversation.
	Links *map[string]interface{} `json:"links,omitempty"`
	// A unique string identifier for the conversation participant as Conversation User.
	ParticipantIdentity *string `json:"participant_identity,omitempty"`
	// Information about how this participant exchanges messages with the conversation.
	ParticipantMessagingBinding *map[string]interface{} `json:"participant_messaging_binding,omitempty"`
	// The unique ID of the Participant.
	ParticipantSid *string `json:"participant_sid,omitempty"`
	// The unique ID for the conversation participant as Conversation User.
	ParticipantUserSid *string `json:"participant_user_sid,omitempty"`
}
