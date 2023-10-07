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
    res, err := s.Measurements.CreateMeasurement(ctx, &shared.MeasurementRequest{
        AdditionalProperties: map[string]interface{}{
            "Florida": "salmon",
        },
        Locations: []shared.MeasurementLocationOption{
            shared.MeasurementLocationOption{
                AdditionalProperties: map[string]interface{}{
                    "superstructure": "Funk",
                },
                Tags: []string{
                    "Analyst",
                },
            },
        },
        MeasurementOptions: shared.CreateMeasurementOptionsMeasurementMtrOptions(
                shared.MeasurementMtrOptions{
                    AdditionalProperties: map[string]interface{}{
                        "driver": "Southwest",
                    },
                },
        ),
        Target: "neutral",
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