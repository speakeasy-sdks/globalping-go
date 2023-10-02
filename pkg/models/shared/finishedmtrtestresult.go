// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/globalping-go/pkg/utils"
)

// FinishedMtrTestResultHopsStats - Summary `rtt` and packet loss statistics.
// All times are in milliseconds.
type FinishedMtrTestResultHopsStats struct {
	// The average `rtt` value.
	Avg float64 `json:"avg"`
	// The number of dropped packets (`total` - `rcv`).
	Drop int64 `json:"drop"`
	// The average jitter value.
	JAvg float64 `json:"jAvg"`
	// The highest jitter value.
	JMax float64 `json:"jMax"`
	// The lowest jitter value.
	JMin float64 `json:"jMin"`
	// The percentage of dropped packets.
	Loss float64 `json:"loss"`
	// The highest `rtt` value.
	Max float64 `json:"max"`
	// The lowest `rtt` value.
	Min float64 `json:"min"`
	// The number of received packets.
	Rcv int64 `json:"rcv"`
	// The standard deviation of the `rtt` values.
	StDev float64 `json:"stDev"`
	// The number of sent packets.
	Total int64 `json:"total"`
}

func (o *FinishedMtrTestResultHopsStats) GetAvg() float64 {
	if o == nil {
		return 0.0
	}
	return o.Avg
}

func (o *FinishedMtrTestResultHopsStats) GetDrop() int64 {
	if o == nil {
		return 0
	}
	return o.Drop
}

func (o *FinishedMtrTestResultHopsStats) GetJAvg() float64 {
	if o == nil {
		return 0.0
	}
	return o.JAvg
}

func (o *FinishedMtrTestResultHopsStats) GetJMax() float64 {
	if o == nil {
		return 0.0
	}
	return o.JMax
}

func (o *FinishedMtrTestResultHopsStats) GetJMin() float64 {
	if o == nil {
		return 0.0
	}
	return o.JMin
}

func (o *FinishedMtrTestResultHopsStats) GetLoss() float64 {
	if o == nil {
		return 0.0
	}
	return o.Loss
}

func (o *FinishedMtrTestResultHopsStats) GetMax() float64 {
	if o == nil {
		return 0.0
	}
	return o.Max
}

func (o *FinishedMtrTestResultHopsStats) GetMin() float64 {
	if o == nil {
		return 0.0
	}
	return o.Min
}

func (o *FinishedMtrTestResultHopsStats) GetRcv() int64 {
	if o == nil {
		return 0
	}
	return o.Rcv
}

func (o *FinishedMtrTestResultHopsStats) GetStDev() float64 {
	if o == nil {
		return 0.0
	}
	return o.StDev
}

func (o *FinishedMtrTestResultHopsStats) GetTotal() int64 {
	if o == nil {
		return 0
	}
	return o.Total
}

type FinishedMtrTestResultHopsTimings struct {
	// The round-trip time for this packet.
	Rtt float64 `json:"rtt"`
}

func (o *FinishedMtrTestResultHopsTimings) GetRtt() float64 {
	if o == nil {
		return 0.0
	}
	return o.Rtt
}

type FinishedMtrTestResultHops struct {
	// The list of ASNs assigned to this hop.
	Asn              []int64 `json:"asn"`
	ResolvedAddress  string  `json:"resolvedAddress"`
	ResolvedHostname *string `json:"resolvedHostname"`
	// Summary `rtt` and packet loss statistics.
	// All times are in milliseconds.
	//
	Stats FinishedMtrTestResultHopsStats `json:"stats"`
	// Details for each sent packet.
	// All times are in milliseconds.
	//
	Timings []FinishedMtrTestResultHopsTimings `json:"timings"`
}

func (o *FinishedMtrTestResultHops) GetAsn() []int64 {
	if o == nil {
		return []int64{}
	}
	return o.Asn
}

func (o *FinishedMtrTestResultHops) GetResolvedAddress() string {
	if o == nil {
		return ""
	}
	return o.ResolvedAddress
}

func (o *FinishedMtrTestResultHops) GetResolvedHostname() *string {
	if o == nil {
		return nil
	}
	return o.ResolvedHostname
}

func (o *FinishedMtrTestResultHops) GetStats() FinishedMtrTestResultHopsStats {
	if o == nil {
		return FinishedMtrTestResultHopsStats{}
	}
	return o.Stats
}

func (o *FinishedMtrTestResultHops) GetTimings() []FinishedMtrTestResultHopsTimings {
	if o == nil {
		return []FinishedMtrTestResultHopsTimings{}
	}
	return o.Timings
}

type FinishedMtrTestResult struct {
	Hops []FinishedMtrTestResultHops `json:"hops"`
	// The raw output can be presented to users but is not meant to be parsed clients.
	// Please use the individual values provided in other fields for automated processing.
	//
	RawOutput string `json:"rawOutput"`
	// The resolved IP address of the `target`.
	ResolvedAddress string `json:"resolvedAddress"`
	// The resolved hostname of the `target`.
	ResolvedHostname *string `json:"resolvedHostname"`
	status           string  `const:"finished" json:"status"`
}

func (f FinishedMtrTestResult) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FinishedMtrTestResult) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *FinishedMtrTestResult) GetHops() []FinishedMtrTestResultHops {
	if o == nil {
		return []FinishedMtrTestResultHops{}
	}
	return o.Hops
}

func (o *FinishedMtrTestResult) GetRawOutput() string {
	if o == nil {
		return ""
	}
	return o.RawOutput
}

func (o *FinishedMtrTestResult) GetResolvedAddress() string {
	if o == nil {
		return ""
	}
	return o.ResolvedAddress
}

func (o *FinishedMtrTestResult) GetResolvedHostname() *string {
	if o == nil {
		return nil
	}
	return o.ResolvedHostname
}

func (o *FinishedMtrTestResult) GetStatus() string {
	return "finished"
}
