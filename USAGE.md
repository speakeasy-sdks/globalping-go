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
        Limit: globalpinggo.Int64(139213),
        Locations: []shared.MeasurementLocationOption{
            shared.MeasurementLocationOption{
                Asn: globalpinggo.Int64(176799),
                City: globalpinggo.String("Maggiobury"),
                Continent: shared.ContinentCodeSa.ToPointer(),
                Country: globalpinggo.String("Mongolia"),
                Limit: globalpinggo.Int64(445560),
                Magic: globalpinggo.String("bluetooth driver Southwest"),
                Network: globalpinggo.String("neutral"),
                Region: shared.RegionNameMicronesia.ToPointer(),
                State: globalpinggo.String("program male Trial"),
                Tags: []string{
                    "Implementation",
                },
            },
        },
        MeasurementOptions: &shared.MeasurementOptions{},
        Target: "Island",
        Type: shared.MeasurementTypeMtr,
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