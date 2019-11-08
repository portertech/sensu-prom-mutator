package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/sensu/sensu-go/types"
)

func main() {
	checkPtr := flag.Bool("check-prefix", false, "enable check name prefix")
	entityPtr := flag.Bool("entity-prefix", false, "enable entity name prefix")
	flag.Parse()

	eventJSON, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read STDIN: %s", err)
		os.Exit(1)
	}

	event := &types.Event{}
	err = json.Unmarshal(eventJSON, event)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to unmarshal STDIN data: %s", err)
		os.Exit(1)
	}
	if event.Metrics == nil {
		fmt.Fprintf(os.Stderr, "No Metrics in event\n")
		os.Exit(1)
	}
	for _, point := range event.Metrics.Points {
		tags := ""
		pointName := ""
		for _, tag := range point.Tags {
			tags = tags + fmt.Sprintf("%s=\"%v\"", tag.Name, tag.Value)
		}

		timestamp := point.Timestamp
		if *entityPtr {
			if event.Entity != nil {
				pointName += event.Entity.Name + "."
			}
		}
		if *checkPtr {
			if event.Check != nil {
				pointName += event.Check.Name + "."
			}
		}
		pointName += point.Name
		if timestamp < 1000000000000 {
			timestamp = time.Unix(timestamp, 0).UnixNano() / int64(time.Millisecond)
		}

		fmt.Printf("%s{%s} %v %v\n", pointName, tags, point.Value, timestamp)
	}
}
