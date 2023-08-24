<!-- Start SDK Example Usage -->


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
        Limit: globalping.Int64(548814),
        Locations: []shared.MeasurementLocationOption{
            shared.MeasurementLocationOption{
                Asn: globalping.Int64(715190),
                City: globalping.String("New Orleans"),
                Continent: shared.ContinentCodeSa.ToPointer(),
                Country: globalping.String("Mali"),
                Limit: globalping.Int64(847252),
                Magic: globalping.String("vel"),
                Network: globalping.String("error"),
                Region: shared.RegionNameEasternEurope.ToPointer(),
                State: globalping.String("suscipit"),
                Tags: []string{
                    "magnam",
                    "debitis",
                },
            },
            shared.MeasurementLocationOption{
                Asn: globalping.Int64(56713),
                City: globalping.String("Edinburg"),
                Continent: shared.ContinentCodeAs.ToPointer(),
                Country: globalping.String("Kyrgyz Republic"),
                Limit: globalping.Int64(791725),
                Magic: globalping.String("placeat"),
                Network: globalping.String("voluptatum"),
                Region: shared.RegionNameEasternAsia.ToPointer(),
                State: globalping.String("excepturi"),
                Tags: []string{
                    "recusandae",
                    "temporibus",
                },
            },
            shared.MeasurementLocationOption{
                Asn: globalping.Int64(71036),
                City: globalping.String("North Lydia"),
                Continent: shared.ContinentCodeAf.ToPointer(),
                Country: globalping.String("Guinea-Bissau"),
                Limit: globalping.Int64(832620),
                Magic: globalping.String("sapiente"),
                Network: globalping.String("quo"),
                Region: shared.RegionNameSouthernAfrica.ToPointer(),
                State: globalping.String("at"),
                Tags: []string{
                    "maiores",
                    "molestiae",
                    "quod",
                    "quod",
                },
            },
        },
        MeasurementOptions: &shared.MeasurementDNSOptions{
            Port: globalping.Int64(520478),
            Protocol: shared.MeasurementDNSOptionsProtocolUDP.ToPointer(),
            Query: &shared.MeasurementDNSOptionsQuery{
                Type: shared.MeasurementDNSOptionsQueryTypePtr.ToPointer(),
            },
            Resolver: globalping.String("184.163.148.36"),
            Trace: globalping.Bool(false),
        },
        Target: "deleniti",
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