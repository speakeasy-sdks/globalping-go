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
                Asn: globalping.Int64(592845),
                City: globalping.String("Sporerstead"),
                Continent: shared.ContinentCodeSa.ToPointer(),
                Country: globalping.String("Mali"),
                Limit: globalping.Int64(847252),
                Magic: globalping.String("vel"),
                Network: globalping.String("error"),
                Region: shared.RegionNameEasternEurope.ToPointer(),
                State: globalping.String("suscipit"),
                Tags: []string{
                    "iure",
                },
            },
        },
        MeasurementOptions: &shared.MeasurementTracerouteOptions{
            Port: globalping.Int64(891773),
            Protocol: shared.MeasurementTracerouteOptionsProtocolIcmp.ToPointer(),
        },
        Target: "delectus",
        Type: shared.MeasurementTypeTraceroute,
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
