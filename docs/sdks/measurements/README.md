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
        Limit: globalping.Int64(383441),
        Locations: []shared.MeasurementLocationOption{
            shared.MeasurementLocationOption{
                Asn: globalping.Int64(477665),
                City: globalping.String("Schuliststad"),
                Continent: shared.ContinentCodeEu.ToPointer(),
                Country: globalping.String("Mayotte"),
                Limit: globalping.Int64(392785),
                Magic: globalping.String("recusandae"),
                Network: globalping.String("temporibus"),
                Region: shared.RegionNameEasternAfrica.ToPointer(),
                State: globalping.String("quis"),
                Tags: []string{
                    "veritatis",
                },
            },
        },
        MeasurementOptions: &shared.MeasurementMtrOptions{
            Packets: globalping.Int64(20218),
            Port: globalping.Int64(368241),
            Protocol: shared.MeasurementMtrOptionsProtocolUDP.ToPointer(),
        },
        Target: "sapiente",
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

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.MeasurementRequest](../../models/shared/measurementrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.CreateMeasurementResponse](../../models/operations/createmeasurementresponse.md), error**

