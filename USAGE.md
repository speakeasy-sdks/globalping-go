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