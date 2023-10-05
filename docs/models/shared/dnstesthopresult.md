# DNSTestHopResult


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Answers`                                                                 | [][DNSTestAnswer](../../models/shared/dnstestanswer.md)                   | :heavy_check_mark:                                                        | The list of received resource records.                                    |
| `Resolver`                                                                | *string*                                                                  | :heavy_check_mark:                                                        | The hostname or IP of the resolver that answered the query.               |
| `Timings`                                                                 | [DNSTestHopResultTimings](../../models/shared/dnstesthopresulttimings.md) | :heavy_check_mark:                                                        | N/A                                                                       |