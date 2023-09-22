# Measurements

### Available Operations

* [CreateMeasurement](#createmeasurement) - Create measurement

## CreateMeasurement

Creates a new measurement with the configured parameters.
The measurement runs asynchronously and its current state can be retrieved
at the URL returned in the `Location` header.

### Client guidelines

- Set the `inProgressUpdates` option to `true` if the application is running in interactive mode so that the user sees the results right away.
  - If the application is interactive by default but also implements a "CI" mode to be used in scripts, do not set the flag in the CI mode.


### Example Usage

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
        Limit: globalpinggo.Int64(87129),
        Locations: []shared.MeasurementLocationOption{
            shared.MeasurementLocationOption{
                Asn: globalpinggo.Int64(648172),
                City: globalpinggo.String("West Ritaworth"),
                Continent: shared.ContinentCodeOc.ToPointer(),
                Country: globalpinggo.String("Burundi"),
                Limit: globalpinggo.Int64(870013),
                Magic: globalpinggo.String("at"),
                Network: globalpinggo.String("maiores"),
                Region: shared.RegionNameEasternAsia.ToPointer(),
                State: globalpinggo.String("quod"),
                Tags: []string{
                    "quod",
                },
            },
        },
        MeasurementOptions: &shared.MeasurementOptions{},
        Target: "esse",
        Type: shared.MeasurementTypeDNS,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateMeasurementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.MeasurementRequest](../../models/shared/measurementrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.CreateMeasurementResponse](../../models/operations/createmeasurementresponse.md), error**

