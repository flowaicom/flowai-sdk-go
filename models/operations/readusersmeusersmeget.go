// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/flowaicom/flowai-sdk-go/models/components"
)

type ReadUsersMeUsersMeGetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	UserRead *components.UserRead
}

func (o *ReadUsersMeUsersMeGetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ReadUsersMeUsersMeGetResponse) GetUserRead() *components.UserRead {
	if o == nil {
		return nil
	}
	return o.UserRead
}
