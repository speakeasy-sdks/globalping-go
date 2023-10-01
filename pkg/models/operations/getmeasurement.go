// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/globalping-go/pkg/models/shared"
	"github.com/speakeasy-sdks/globalping-go/pkg/utils"
	"net/http"
	"time"
)

type GetMeasurementRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetMeasurementRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetMeasurement404ApplicationJSONError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func (o *GetMeasurement404ApplicationJSONError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *GetMeasurement404ApplicationJSONError) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

// GetMeasurement404ApplicationJSON - Not Found
type GetMeasurement404ApplicationJSON struct {
	Error GetMeasurement404ApplicationJSONError `json:"error"`
}

func (o *GetMeasurement404ApplicationJSON) GetError() GetMeasurement404ApplicationJSONError {
	if o == nil {
		return GetMeasurement404ApplicationJSONError{}
	}
	return o.Error
}

// GetMeasurement200ApplicationJSON - Success
type GetMeasurement200ApplicationJSON struct {
	// Time when the measurement was created.
	CreatedAt          time.Time                          `json:"createdAt"`
	ID                 string                             `json:"id"`
	Limit              *int64                             `default:"1" json:"limit"`
	Locations          []shared.MeasurementLocationOption `json:"locations,omitempty"`
	MeasurementOptions *shared.MeasurementOptions         `json:"measurementOptions,omitempty"`
	// The number of probes that performed the measurement. Smaller or equal to `limit`.
	ProbesCount int64 `json:"probesCount"`
	// The measurement results.
	Results []shared.MeasurementResultItem `json:"results"`
	// The measurement status. Any value other than `in-progress` is final.
	Status shared.MeasurementStatus `json:"status"`
	// A public endpoint on which the measurement is executed.
	// Typically a hostname or an IPv4 address. The exact format depends on the measurement type.
	//
	Target string                 `json:"target"`
	Type   shared.MeasurementType `json:"type"`
	// Time when the measurement was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
}

func (g GetMeasurement200ApplicationJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetMeasurement200ApplicationJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetMeasurement200ApplicationJSON) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *GetMeasurement200ApplicationJSON) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetMeasurement200ApplicationJSON) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetMeasurement200ApplicationJSON) GetLocations() []shared.MeasurementLocationOption {
	if o == nil {
		return nil
	}
	return o.Locations
}

func (o *GetMeasurement200ApplicationJSON) GetMeasurementOptions() *shared.MeasurementOptions {
	if o == nil {
		return nil
	}
	return o.MeasurementOptions
}

func (o *GetMeasurement200ApplicationJSON) GetProbesCount() int64 {
	if o == nil {
		return 0
	}
	return o.ProbesCount
}

func (o *GetMeasurement200ApplicationJSON) GetResults() []shared.MeasurementResultItem {
	if o == nil {
		return []shared.MeasurementResultItem{}
	}
	return o.Results
}

func (o *GetMeasurement200ApplicationJSON) GetStatus() shared.MeasurementStatus {
	if o == nil {
		return shared.MeasurementStatus("")
	}
	return o.Status
}

func (o *GetMeasurement200ApplicationJSON) GetTarget() string {
	if o == nil {
		return ""
	}
	return o.Target
}

func (o *GetMeasurement200ApplicationJSON) GetType() shared.MeasurementType {
	if o == nil {
		return shared.MeasurementType("")
	}
	return o.Type
}

func (o *GetMeasurement200ApplicationJSON) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

type GetMeasurementResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	GetMeasurement200ApplicationJSONObject *GetMeasurement200ApplicationJSON
	// Not Found
	GetMeasurement404ApplicationJSONObject *GetMeasurement404ApplicationJSON
}

func (o *GetMeasurementResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeasurementResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeasurementResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeasurementResponse) GetGetMeasurement200ApplicationJSONObject() *GetMeasurement200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetMeasurement200ApplicationJSONObject
}

func (o *GetMeasurementResponse) GetGetMeasurement404ApplicationJSONObject() *GetMeasurement404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetMeasurement404ApplicationJSONObject
}