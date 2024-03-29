// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/globalping-go/pkg/utils"
)

// OfflineTestResult - Represents an `offline` test where the requested probe is currently offline and most fields are not available.
// Only possible when passing an `id` of a previous measurement to the `locations` field.
type OfflineTestResult struct {
	// The raw output can be presented to users but is not meant to be parsed clients.
	// Please use the individual values provided in other fields for automated processing.
	//
	RawOutput string `json:"rawOutput"`
	status    string `const:"offline" json:"status"`
}

func (o OfflineTestResult) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OfflineTestResult) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *OfflineTestResult) GetRawOutput() string {
	if o == nil {
		return ""
	}
	return o.RawOutput
}

func (o *OfflineTestResult) GetStatus() string {
	return "offline"
}
