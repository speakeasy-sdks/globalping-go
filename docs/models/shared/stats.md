# Stats

Summary `rtt` and packet loss statistics.
All times are in milliseconds.



## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `Avg`                                            | *float64*                                        | :heavy_check_mark:                               | N/A                                              |
| `Drop`                                           | *int64*                                          | :heavy_check_mark:                               | The number of dropped packets (`total` - `rcv`). |
| `Loss`                                           | *float64*                                        | :heavy_check_mark:                               | The percentage of dropped packets.               |
| `Max`                                            | *float64*                                        | :heavy_check_mark:                               | N/A                                              |
| `Min`                                            | *float64*                                        | :heavy_check_mark:                               | N/A                                              |
| `Rcv`                                            | *int64*                                          | :heavy_check_mark:                               | The number of received packets.                  |
| `Total`                                          | *int64*                                          | :heavy_check_mark:                               | The number of sent packets.                      |