// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/flowaicom/flowai-sdk-go/models/components"
)

type RevokeAPIKeyAPIKeysKeyIDDeleteRequest struct {
	// The ID of the API key to revoke
	KeyID string `pathParam:"style=simple,explode=false,name=key_id"`
}

func (o *RevokeAPIKeyAPIKeysKeyIDDeleteRequest) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

type RevokeAPIKeyAPIKeysKeyIDDeleteResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *RevokeAPIKeyAPIKeysKeyIDDeleteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
