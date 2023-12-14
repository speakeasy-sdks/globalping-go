# DNSTestHopResult


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Answers`                                                             | [][shared.DNSTestAnswer](../../../pkg/models/shared/dnstestanswer.md) | :heavy_check_mark:                                                    | The list of received resource records.                                |
| `Resolver`                                                            | *string*                                                              | :heavy_check_mark:                                                    | The hostname or IP of the resolver that answered the query.           |
| `Timings`                                                             | [shared.Timings](../../../pkg/models/shared/timings.md)               | :heavy_check_mark:                                                    | N/A                                                                   |