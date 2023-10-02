# Measurements
(*Measurements*)

### Available Operations

* [CreateMeasurement](#createmeasurement) - Create measurement
* [GetMeasurement](#getmeasurement) - Get measurement

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

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.MeasurementRequest](../../models/shared/measurementrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.CreateMeasurementResponse](../../models/operations/createmeasurementresponse.md), error**


## GetMeasurement

Retrieves data of an existing measurement.
A link to this endpoint is returned in the `Location` response header when creating the measurement.
The measurement is typically available for up to 7 days after creation.

### Client guidelines

- Use the following algorithm for measurement result pooling:
  1. Request the measurement status.
  2. If the status is `in-progress`, wait 500 ms and repeat from step 1. Note that it is important to wait 500 ms *after* receiving the response, instead of simply using an "every 500 ms" interval. For large measurements, the request itself may take a few hundred milliseconds to complete.
  3. If the status is anything else, stop. The measurement is no longer running. Any value other than `in-progress` is final.


### Example Usage

```go
package main

import(
	"context"
	"log"
	globalpinggo "github.com/speakeasy-sdks/globalping-go"
	"github.com/speakeasy-sdks/globalping-go/pkg/models/operations"
)

func main() {
    s := globalpinggo.New()

    ctx := context.Background()
    res, err := s.Measurements.GetMeasurement(ctx, operations.GetMeasurementRequest{
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetMeasurement200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetMeasurementRequest](../../models/operations/getmeasurementrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetMeasurementResponse](../../models/operations/getmeasurementresponse.md), error**

