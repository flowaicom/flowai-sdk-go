// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/flowaicom/flowai-sdk-go/internal/utils"
	"time"
)

type APIKeyInfo struct {
	// Optional user-friendly name for the key
	KeyName    *string    `json:"key_name,omitempty"`
	ID         string     `json:"id"`
	UserID     string     `json:"user_id"`
	KeyPrefix  string     `json:"key_prefix"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
	IsActive   bool       `json:"is_active"`
}

func (a APIKeyInfo) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *APIKeyInfo) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *APIKeyInfo) GetKeyName() *string {
	if o == nil {
		return nil
	}
	return o.KeyName
}

func (o *APIKeyInfo) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *APIKeyInfo) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *APIKeyInfo) GetKeyPrefix() string {
	if o == nil {
		return ""
	}
	return o.KeyPrefix
}

func (o *APIKeyInfo) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *APIKeyInfo) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *APIKeyInfo) GetLastUsedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastUsedAt
}

func (o *APIKeyInfo) GetIsActive() bool {
	if o == nil {
		return false
	}
	return o.IsActive
}
