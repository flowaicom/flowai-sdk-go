// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/flowaicom/flowai-sdk-go/internal/utils"
	"time"
)

type JobCancelResponse struct {
	ID        string         `json:"id"`
	ProjectID string         `json:"project_id"`
	Status    string         `json:"status"`
	StartTime *time.Time     `json:"start_time,omitempty"`
	EndTime   *time.Time     `json:"end_time,omitempty"`
	Error     *string        `json:"error,omitempty"`
	Config    map[string]any `json:"config,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func (j JobCancelResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(j, "", false)
}

func (j *JobCancelResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &j, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *JobCancelResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *JobCancelResponse) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *JobCancelResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *JobCancelResponse) GetStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *JobCancelResponse) GetEndTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *JobCancelResponse) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *JobCancelResponse) GetConfig() map[string]any {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *JobCancelResponse) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *JobCancelResponse) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
