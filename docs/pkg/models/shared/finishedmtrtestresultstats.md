# FinishedMtrTestResultStats

Summary `rtt` and packet loss statistics.
All times are in milliseconds.



## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `Avg`                                            | *float64*                                        | :heavy_check_mark:                               | The average `rtt` value.                         |
| `Drop`                                           | *int64*                                          | :heavy_check_mark:                               | The number of dropped packets (`total` - `rcv`). |
| `JAvg`                                           | *float64*                                        | :heavy_check_mark:                               | The average jitter value.                        |
| `JMax`                                           | *float64*                                        | :heavy_check_mark:                               | The highest jitter value.                        |
| `JMin`                                           | *float64*                                        | :heavy_check_mark:                               | The lowest jitter value.                         |
| `Loss`                                           | *float64*                                        | :heavy_check_mark:                               | The percentage of dropped packets.               |
| `Max`                                            | *float64*                                        | :heavy_check_mark:                               | The highest `rtt` value.                         |
| `Min`                                            | *float64*                                        | :heavy_check_mark:                               | The lowest `rtt` value.                          |
| `Rcv`                                            | *int64*                                          | :heavy_check_mark:                               | The number of received packets.                  |
| `StDev`                                          | *float64*                                        | :heavy_check_mark:                               | The standard deviation of the `rtt` values.      |
| `Total`                                          | *int64*                                          | :heavy_check_mark:                               | The number of sent packets.                      |