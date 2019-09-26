# Sensu Go Prometheus Mutator
TravisCI: [![TravisCI Build Status](https://travis-ci.org/portertech/sensu-prom-mutator.svg?branch=master)](https://travis-ci.org/portertech/sensu-prom-mutator)

The Sensu Go Prometheus Mutator is a [Sensu Event Mutator][1] which
mutates Sensu Go Event metrics into to the [Prometheus metric
format][2]. This mutator is intended to be used in combination with a
TCP Event Handler to send metrics to a [Sumo Logic][3] collector.

[1]: https://docs.sensu.io/sensu-go/5.13/reference/mutators/#how-do-mutators-work
[2]: https://github.com/prometheus/docs/blob/master/content/docs/instrumenting/exposition_formats.md
[3]: https://www.sumologic.com/
