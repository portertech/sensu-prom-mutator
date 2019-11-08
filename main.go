package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/sensu/sensu-go/types"
)

func main() {
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
		tags := make(map[string]string)
		if event.Entity != nil {
			tags["sensu_entity_name"] = event.Entity.Name
		}
		if event.Check != nil {
			tags["sensu_check_name"] = event.Check.Name
		}
		for i, tag := range point.Tags {
			tags[tag.Name] = tag.Value
		}
		tagString := ""
		for i, tag := range tags {
			if i == 0 {
				tagString = tagString + fmt.Sprintf("%s=\"%v\"", tag.Name, tag.Value)
			} else {
				tagString = tagString + fmt.Sprintf(" , %s=\"%v\"", tag.Name, tag.Value)
			}
		}

		timestamp := point.Timestamp

		if timestamp < 1000000000000 {
			timestamp = time.Unix(timestamp, 0).UnixNano() / int64(time.Millisecond)
		}

		fmt.Printf("%s{%s} %v %v\n", point.Name, tagString, point.Value, timestamp)
	}
}
