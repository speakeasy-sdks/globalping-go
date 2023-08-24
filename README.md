# github.com/speakeasy-sdks/globalping-go

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/globalping-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/globalping-go"
	"github.com/speakeasy-sdks/globalping-go/pkg/models/shared"
)

func main() {
    s := globalping.New()

    ctx := context.Background()
    res, err := s.Measurements.CreateMeasurement(ctx, shared.MeasurementRequest{
        InProgressUpdates: globalping.Bool(false),
        Limit: globalping.Int64(548814),
        Locations: []shared.MeasurementLocationOption{
            shared.MeasurementLocationOption{
                Asn: globalping.Int64(715190),
                City: globalping.String("New Orleans"),
                Continent: shared.ContinentCodeSa.ToPointer(),
                Country: globalping.String("Mali"),
                Limit: globalping.Int64(847252),
                Magic: globalping.String("vel"),
                Network: globalping.String("error"),
                Region: shared.RegionNameEasternEurope.ToPointer(),
                State: globalping.String("suscipit"),
                Tags: []string{
                    "magnam",
                    "debitis",
                },
            },
            shared.MeasurementLocationOption{
                Asn: globalping.Int64(56713),
                City: globalping.String("Edinburg"),
                Continent: shared.ContinentCodeAs.ToPointer(),
                Country: globalping.String("Kyrgyz Republic"),
                Limit: globalping.Int64(791725),
                Magic: globalping.String("placeat"),
                Network: globalping.String("voluptatum"),
                Region: shared.RegionNameEasternAsia.ToPointer(),
                State: globalping.String("excepturi"),
                Tags: []string{
                    "recusandae",
                    "temporibus",
                },
            },
            shared.MeasurementLocationOption{
                Asn: globalping.Int64(71036),
                City: globalping.String("North Lydia"),
                Continent: shared.ContinentCodeAf.ToPointer(),
                Country: globalping.String("Guinea-Bissau"),
                Limit: globalping.Int64(832620),
                Magic: globalping.String("sapiente"),
                Network: globalping.String("quo"),
                Region: shared.RegionNameSouthernAfrica.ToPointer(),
                State: globalping.String("at"),
                Tags: []string{
                    "maiores",
                    "molestiae",
                    "quod",
                    "quod",
                },
            },
        },
        MeasurementOptions: &shared.MeasurementDNSOptions{
            Port: globalping.Int64(520478),
            Protocol: shared.MeasurementDNSOptionsProtocolUDP.ToPointer(),
            Query: &shared.MeasurementDNSOptionsQuery{
                Type: shared.MeasurementDNSOptionsQueryTypePtr.ToPointer(),
            },
            Resolver: globalping.String("184.163.148.36"),
            Trace: globalping.Bool(false),
        },
        Target: "deleniti",
        Type: shared.MeasurementTypeHTTP,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateMeasurementResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Measurements](docs/sdks/measurements/README.md)

* [CreateMeasurement](docs/sdks/measurements/README.md#createmeasurement) - Create measurement

### [Probes](docs/sdks/probes/README.md)

* [ListProbes](docs/sdks/probes/README.md#listprobes) - List currently connected probes
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
