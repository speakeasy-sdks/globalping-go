# Probes
(*Probes*)

### Available Operations

* [ListProbes](#listprobes) - List currently connected probes

## ListProbes

Returns a list of all currently connected probes and their metadata.


### Example Usage

```go
package main

import(
	globalpinggo "github.com/speakeasy-sdks/globalping-go"
	"context"
	"log"
)

func main() {
    s := globalpinggo.New()

    ctx := context.Background()
    res, err := s.Probes.ListProbes(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Probes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListProbesResponse](../../pkg/models/operations/listprobesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
