package flags

import (
	"fmt"
	"strings"
)

type AdvancedFeatures []string

const (
	Htmx              string = "htmx"
	GoProjectWorkflow string = "githubaction"
	Websocket         string = "websocket"
	Tailwind          string = "tailwind"
	React             string = "react"
	Docker            string = "docker"
	Kafka             string = "kafka"
	Worker            string = "worker"
	Redis             string = "redis"
)

var AllowedAdvancedFeatures = []string{string(React), string(Htmx), string(GoProjectWorkflow), string(Websocket), string(Tailwind), string(Docker), string(Kafka), string(Worker), string(Redis)}

func (f AdvancedFeatures) String() string {
	return strings.Join(f, ",")
}

func (f *AdvancedFeatures) Type() string {
	return "AdvancedFeatures"
}

func (f *AdvancedFeatures) Set(value string) error {
	// Contains isn't available in 1.20 yet
	// if AdvancedFeatures.Contains(value) {
	for _, advancedFeature := range AllowedAdvancedFeatures {
		if advancedFeature == value {
			*f = append(*f, advancedFeature)
			return nil
		}
	}

	return fmt.Errorf("advanced Feature to use. Allowed values: %s", strings.Join(AllowedAdvancedFeatures, ", "))
}
