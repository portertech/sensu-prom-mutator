# Sensu Go Prometheus Mutator
TravisCI: [![TravisCI Build Status](https://travis-ci.org/portertech/sensu-prom-mutator.svg?branch=master)](https://travis-ci.org/portertech/sensu-prom-mutator)

The Sensu Go Prometheus Mutator is a [Sensu Event Mutator][1] which
mutates Sensu Go Event metrics into to the [Prometheus metric
format][2]. This mutator is intended to be used in combination with a
TCP Event Handler to send metrics to a [Sumo Logic][3] collector.

This mutator turns this:

``` json
{
  "entity": "...",
  "metrics": {
    "handlers": [
      "sumologic"
    ],
    "points": [
      {
        "name": "sensu-go-sandbox.curl_timings.time_total",
        "tags": [
          {
            "name": "foo",
            "value": "42"
          }
        ],
        "timestamp": 1552506033,
        "value": 0.005
      },
      {
        "name": "sensu-go-sandbox.curl_timings.time_namelookup",
        "tags": [],
        "timestamp": 1552506033,
        "value": 0.004
      }
    ]
  },
  "timestamp": 1552506033
}
```

Into this:

```
sensu-go-sandbox.curl_timings.time_total{foo="42"} 0.005 1552506033000
sensu-go-sandbox.curl_timings.time_namelookup{} 0.004 1552506033000
```

[1]: https://docs.sensu.io/sensu-go/5.13/reference/mutators/#how-do-mutators-work
[2]: https://github.com/prometheus/docs/blob/master/content/docs/instrumenting/exposition_formats.md
[3]: https://www.sumologic.com/
