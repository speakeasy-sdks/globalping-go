<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	globalpinggo "github.com/speakeasy-sdks/globalping-go"
	"github.com/speakeasy-sdks/globalping-go/pkg/models/operations"
	"github.com/speakeasy-sdks/globalping-go/pkg/models/shared"
	"log"
)

func main() {
	s := globalpinggo.New()

	operationSecurity := operations.CreateMeasurementSecurity{
		BearerAuth: globalpinggo.String("<YOUR_BEARER_TOKEN_HERE>"),
	}

	ctx := context.Background()
	res, err := s.Measurements.CreateMeasurement(ctx, &shared.MeasurementRequest{
		Locations: shared.CreateMeasurementLocationsArrayOfMeasurementLocationOption(
			[]shared.MeasurementLocationOption{
				shared.MeasurementLocationOption{
					Tags: []string{
						"string",
					},
				},
			},
		),
		MeasurementOptions: shared.CreateMeasurementOptionsMeasurementPingOptions(
			shared.MeasurementPingOptions{},
		),
		Target: "string",
		Type:   shared.MeasurementTypeMtr,
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateMeasurementResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->