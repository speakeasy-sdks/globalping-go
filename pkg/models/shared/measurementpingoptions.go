// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MeasurementPingOptions struct {
	Packets *int64 `json:"packets,omitempty"`
}

func (o *MeasurementPingOptions) GetPackets() *int64 {
	if o == nil {
		return nil
	}
	return o.Packets
}
