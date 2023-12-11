// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/globalping-go/pkg/utils"
)

type MeasurementPingOptions struct {
	Packets *int64 `default:"3" json:"packets"`
}

func (m MeasurementPingOptions) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeasurementPingOptions) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *MeasurementPingOptions) GetPackets() *int64 {
	if o == nil {
		return nil
	}
	return o.Packets
}
