// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/globalping-go/pkg/utils"
)

// FailedTestResult - Represents a `failed` test where most fields are not available.
type FailedTestResult struct {
	// The raw output can be presented to users but is not meant to be parsed clients.
	// Please use the individual values provided in other fields for automated processing.
	//
	RawOutput string `json:"rawOutput"`
	status    string `const:"failed" json:"status"`
}

func (f FailedTestResult) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FailedTestResult) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *FailedTestResult) GetRawOutput() string {
	if o == nil {
		return ""
	}
	return o.RawOutput
}

func (o *FailedTestResult) GetStatus() string {
	return "failed"
}
