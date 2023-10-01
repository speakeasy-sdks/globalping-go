// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/globalping-go/pkg/utils"
)

type FinishedTraceDNSTestResult struct {
	Hops []DNSTestHopResult `json:"hops"`
	// The raw output can be presented to users but is not meant to be parsed clients.
	// Please use the individual values provided in other fields for automated processing.
	//
	RawOutput string `json:"rawOutput"`
	status    string `const:"finished" json:"status"`
}

func (f FinishedTraceDNSTestResult) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FinishedTraceDNSTestResult) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *FinishedTraceDNSTestResult) GetHops() []DNSTestHopResult {
	if o == nil {
		return []DNSTestHopResult{}
	}
	return o.Hops
}

func (o *FinishedTraceDNSTestResult) GetRawOutput() string {
	if o == nil {
		return ""
	}
	return o.RawOutput
}

func (o *FinishedTraceDNSTestResult) GetStatus() string {
	return "finished"
}
