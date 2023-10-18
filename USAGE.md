<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	globalpinggo "github.com/speakeasy-sdks/globalping-go"
	"github.com/speakeasy-sdks/globalping-go/pkg/models/shared"
	"log"
)

func main() {
	s := globalpinggo.New()

	ctx := context.Background()
	res, err := s.Measurements.CreateMeasurement(ctx, &shared.MeasurementRequest{
		Locations: []shared.MeasurementLocationOption{
			shared.MeasurementLocationOption{
				Tags: []string{
					"Florida",
				},
			},
		},
		MeasurementOptions: shared.CreateMeasurementOptionsMeasurementMtrOptions(
			shared.MeasurementMtrOptions{},
		),
		Target: "nisi",
		Type:   shared.MeasurementTypeHTTP,
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