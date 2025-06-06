// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ProjectValidatorListResponse - Schema for listing project validators.
type ProjectValidatorListResponse struct {
	Validators []ProjectValidatorResponse `json:"validators"`
}

func (o *ProjectValidatorListResponse) GetValidators() []ProjectValidatorResponse {
	if o == nil {
		return []ProjectValidatorResponse{}
	}
	return o.Validators
}
