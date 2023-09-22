<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	globalpinggo "github.com/speakeasy-sdks/globalping-go"
	"github.com/speakeasy-sdks/globalping-go/pkg/models/shared"
)

func main() {
    s := globalpinggo.New()

    ctx := context.Background()
    res, err := s.Measurements.CreateMeasurement(ctx, shared.MeasurementRequest{
        InProgressUpdates: globalpinggo.Bool(false),
        Limit: globalpinggo.Int64(548814),
        Locations: []shared.MeasurementLocationOption{
            shared.MeasurementLocationOption{
                Asn: globalpinggo.Int64(592845),
                City: globalpinggo.String("Sporerstead"),
                Continent: shared.ContinentCodeSa.ToPointer(),
                Country: globalpinggo.String("Mali"),
                Limit: globalpinggo.Int64(847252),
                Magic: globalpinggo.String("vel"),
                Network: globalpinggo.String("error"),
                Region: shared.RegionNameEasternEurope.ToPointer(),
                State: globalpinggo.String("suscipit"),
                Tags: []string{
                    "iure",
                },
            },
        },
        MeasurementOptions: &shared.MeasurementOptions{},
        Target: "magnam",
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