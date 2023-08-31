// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/globalping-go/pkg/models/shared"
	"net/http"
)

type ListProbesResponse struct {
	ContentType string
	// Success
	Probes      []shared.Probe
	StatusCode  int
	RawResponse *http.Response
}

func (o *ListProbesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListProbesResponse) GetProbes() []shared.Probe {
	if o == nil {
		return nil
	}
	return o.Probes
}

func (o *ListProbesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListProbesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}