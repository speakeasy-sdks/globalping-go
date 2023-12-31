// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/globalping-go/pkg/utils"
)

type MeasurementLocationOption struct {
	// An autonomous system number.
	Asn *int64 `json:"asn,omitempty"`
	// A city name in English.
	City      *string        `json:"city,omitempty"`
	Continent *ContinentCode `json:"continent,omitempty"`
	// An ISO 3166-1 alpha-2 country code.
	Country *string `json:"country,omitempty"`
	// Mutually exclusive with the global `limit`.
	// Specifies the maximum number of probes that run the measurement in this location.
	//
	Limit *int64 `default:"1" json:"limit"`
	// Fuzzy matching based on the `country`, `city`, `state`, `continent`, `region`, `asn` (using `AS` prefix, e.g., `AS123`), `tags`, and `network` values.
	// Includes the full names, ISO codes (where applicable), and common aliases.
	// Multiple conditions can be combined using the `+` character.
	//
	Magic *string `json:"magic,omitempty"`
	// A network name.
	Network *string `json:"network,omitempty"`
	// A Geographic Region name based on the
	// UN [Standard Country or Area Codes for Statistical Use (M49)](https://unstats.un.org/unsd/methodology/m49/).
	//
	Region *RegionName `json:"region,omitempty"`
	// Only for US states. A two-letter state code based on ISO 3166-2.
	State *string `json:"state,omitempty"`
	// An array of additional values that can be used to target the probe.
	// Probes hosted in [AWS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions)
	// and [Google Cloud](https://cloud.google.com/compute/docs/regions-zones#available) are automatically assigned the service region code.
	// Probes are also automatically assigned `datacenter-network` and `eyeball-network` tags to distinguish between datacenter and end-user locations.
	//
	Tags []string `json:"tags,omitempty"`
}

func (m MeasurementLocationOption) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeasurementLocationOption) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeasurementLocationOption) GetAsn() *int64 {
	if o == nil {
		return nil
	}
	return o.Asn
}

func (o *MeasurementLocationOption) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *MeasurementLocationOption) GetContinent() *ContinentCode {
	if o == nil {
		return nil
	}
	return o.Continent
}

func (o *MeasurementLocationOption) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *MeasurementLocationOption) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *MeasurementLocationOption) GetMagic() *string {
	if o == nil {
		return nil
	}
	return o.Magic
}

func (o *MeasurementLocationOption) GetNetwork() *string {
	if o == nil {
		return nil
	}
	return o.Network
}

func (o *MeasurementLocationOption) GetRegion() *RegionName {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *MeasurementLocationOption) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *MeasurementLocationOption) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
