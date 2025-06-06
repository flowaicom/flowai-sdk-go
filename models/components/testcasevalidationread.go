// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/flowaicom/flowai-sdk-go/internal/utils"
	"time"
)

// TestCaseValidationRead - Schema for reading/returning TestCaseValidation data
type TestCaseValidationRead struct {
	IsAccepted       bool                     `json:"is_accepted"`
	Feedback         *string                  `json:"feedback,omitempty"`
	ID               string                   `json:"id"`
	CreatedAt        time.Time                `json:"created_at"`
	UpdatedAt        time.Time                `json:"updated_at"`
	TestCaseID       string                   `json:"test_case_id"`
	ValidationTaskID string                   `json:"validation_task_id"`
	ValidatorTaskID  string                   `json:"validator_task_id"`
	ValidatorUserID  string                   `json:"validator_user_id"`
	ItemFeedbacks    []ValidationItemFeedback `json:"item_feedbacks,omitempty"`
}

func (t TestCaseValidationRead) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TestCaseValidationRead) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TestCaseValidationRead) GetIsAccepted() bool {
	if o == nil {
		return false
	}
	return o.IsAccepted
}

func (o *TestCaseValidationRead) GetFeedback() *string {
	if o == nil {
		return nil
	}
	return o.Feedback
}

func (o *TestCaseValidationRead) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TestCaseValidationRead) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *TestCaseValidationRead) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *TestCaseValidationRead) GetTestCaseID() string {
	if o == nil {
		return ""
	}
	return o.TestCaseID
}

func (o *TestCaseValidationRead) GetValidationTaskID() string {
	if o == nil {
		return ""
	}
	return o.ValidationTaskID
}

func (o *TestCaseValidationRead) GetValidatorTaskID() string {
	if o == nil {
		return ""
	}
	return o.ValidatorTaskID
}

func (o *TestCaseValidationRead) GetValidatorUserID() string {
	if o == nil {
		return ""
	}
	return o.ValidatorUserID
}

func (o *TestCaseValidationRead) GetItemFeedbacks() []ValidationItemFeedback {
	if o == nil {
		return nil
	}
	return o.ItemFeedbacks
}
