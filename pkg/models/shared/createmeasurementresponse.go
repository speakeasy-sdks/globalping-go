// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateMeasurementResponse struct {
	ID string `json:"id"`
	// The number of probes that performed the measurement. Smaller or equal to `limit`.
	ProbesCount int64 `json:"probesCount"`
}

func (o *CreateMeasurementResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateMeasurementResponse) GetProbesCount() int64 {
	if o == nil {
		return 0
	}
	return o.ProbesCount
}
