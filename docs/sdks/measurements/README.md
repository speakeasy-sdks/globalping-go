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
	"github.com/speakeasy-sdks/globalping-go"
	"github.com/speakeasy-sdks/globalping-go/pkg/models/shared"
)

func main() {
    s := globalping.New()

    ctx := context.Background()
    res, err := s.Measurements.CreateMeasurement(ctx, shared.MeasurementRequest{
        InProgressUpdates: globalping.Bool(false),
        Limit: globalping.Int64(758616),
        Locations: []shared.MeasurementLocationOption{
            shared.MeasurementLocationOption{
                Asn: globalping.Int64(105907),
                City: globalping.String("Jarenboro"),
                Continent: shared.ContinentCodeAn.ToPointer(),
                Country: globalping.String("Samoa"),
                Limit: globalping.Int64(736918),
                Magic: globalping.String("esse"),
                Network: globalping.String("ipsum"),
                Region: shared.RegionNameSouthernAsia.ToPointer(),
                State: globalping.String("aspernatur"),
                Tags: []string{
                    "ad",
                },
            },
            shared.MeasurementLocationOption{
                Asn: globalping.Int64(617636),
                City: globalping.String("New Dameonfurt"),
                Continent: shared.ContinentCodeAs.ToPointer(),
                Country: globalping.String("United States of America"),
                Limit: globalping.Int64(902599),
                Magic: globalping.String("fuga"),
                Network: globalping.String("in"),
                Region: shared.RegionNameSouthAmerica.ToPointer(),
                State: globalping.String("iste"),
                Tags: []string{
                    "saepe",
                    "quidem",
                },
            },
            shared.MeasurementLocationOption{
                Asn: globalping.Int64(99280),
                City: globalping.String("Fort Manuelachester"),
                Continent: shared.ContinentCodeNa.ToPointer(),
                Country: globalping.String("Chad"),
                Limit: globalping.Int64(210382),
                Magic: globalping.String("corporis"),
                Network: globalping.String("explicabo"),
                Region: shared.RegionNameSouthernEurope.ToPointer(),
                State: globalping.String("enim"),
                Tags: []string{
                    "nemo",
                    "minima",
                    "excepturi",
                },
            },
        },
        MeasurementOptions: &shared.MeasurementPingOptions{
            Packets: globalping.Int64(438601),
        },
        Target: "culpa",
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

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.MeasurementRequest](../../models/shared/measurementrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.CreateMeasurementResponse](../../models/operations/createmeasurementresponse.md), error**

