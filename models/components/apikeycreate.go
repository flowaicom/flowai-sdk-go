// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type APIKeyCreate struct {
	// A user-defined name to identify the key
	KeyName string `json:"keyName"`
}

func (o *APIKeyCreate) GetKeyName() string {
	if o == nil {
		return ""
	}
	return o.KeyName
}
