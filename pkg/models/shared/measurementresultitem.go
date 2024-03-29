// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// MeasurementResultItemProbe - Information about the probe that performed this test.
type MeasurementResultItemProbe struct {
	// An autonomous system number.
	Asn int64 `json:"asn"`
	// A city name in English.
	City      string        `json:"city"`
	Continent ContinentCode `json:"continent"`
	// An ISO 3166-1 alpha-2 country code.
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	// A network name.
	Network string `json:"network"`
	// A Geographic Region name based on the
	// UN [Standard Country or Area Codes for Statistical Use (M49)](https://unstats.un.org/unsd/methodology/m49/).
	//
	Region RegionName `json:"region"`
	// A list of default DNS resolvers configured on the probe.
	Resolvers []string `json:"resolvers"`
	// Only for US states. A two-letter state code based on ISO 3166-2.
	State *string `json:"state"`
	// An array of additional values that can be used to target the probe.
	// Probes hosted in [AWS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions)
	// and [Google Cloud](https://cloud.google.com/compute/docs/regions-zones#available) are automatically assigned the service region code.
	// Probes are also automatically assigned `datacenter-network` and `eyeball-network` tags to distinguish between datacenter and end-user locations.
	//
	Tags []string `json:"tags"`
}

func (o *MeasurementResultItemProbe) GetAsn() int64 {
	if o == nil {
		return 0
	}
	return o.Asn
}

func (o *MeasurementResultItemProbe) GetCity() string {
	if o == nil {
		return ""
	}
	return o.City
}

func (o *MeasurementResultItemProbe) GetContinent() ContinentCode {
	if o == nil {
		return ContinentCode("")
	}
	return o.Continent
}

func (o *MeasurementResultItemProbe) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}

func (o *MeasurementResultItemProbe) GetLatitude() float64 {
	if o == nil {
		return 0.0
	}
	return o.Latitude
}

func (o *MeasurementResultItemProbe) GetLongitude() float64 {
	if o == nil {
		return 0.0
	}
	return o.Longitude
}

func (o *MeasurementResultItemProbe) GetNetwork() string {
	if o == nil {
		return ""
	}
	return o.Network
}

func (o *MeasurementResultItemProbe) GetRegion() RegionName {
	if o == nil {
		return RegionName("")
	}
	return o.Region
}

func (o *MeasurementResultItemProbe) GetResolvers() []string {
	if o == nil {
		return []string{}
	}
	return o.Resolvers
}

func (o *MeasurementResultItemProbe) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *MeasurementResultItemProbe) GetTags() []string {
	if o == nil {
		return []string{}
	}
	return o.Tags
}

type MeasurementResultItem struct {
	Probe  MeasurementResultItemProbe `json:"probe"`
	Result TestResult                 `json:"result"`
}

func (o *MeasurementResultItem) GetProbe() MeasurementResultItemProbe {
	if o == nil {
		return MeasurementResultItemProbe{}
	}
	return o.Probe
}

func (o *MeasurementResultItem) GetResult() TestResult {
	if o == nil {
		return TestResult{}
	}
	return o.Result
}
