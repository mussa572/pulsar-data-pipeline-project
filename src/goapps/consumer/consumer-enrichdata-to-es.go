package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

type SensorEReading struct {
	SensorId    string    `json:"sensorId"`
	StationId   int       `json:"stationId"`
	Status      string    `json:"status"`
	StartupTime time.Time `json:"startupTime"`
	EventTime   time.Time `json:"eventTime"`
	Reading     string    `json:"reading"`
	StationName string    `json:"stationName"`
}

func main() {

	Init()

	var (
		exampleSchemaDef = "{\"type\":\"record\",\"name\":\"ReadingEnriched\",\"namespace\":\"io.streamnative.models\"," +
			"\"fields\":[{\"name\":\"eventTime\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"reading\",\"type\":[\"null\",\"string\"],\"default\": null},{\"name\":\"sensorId\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"startupTime\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"stationId\",\"type\":\"int\"}," +
			"{\"name\":\"stationName\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"status\",\"type\":[\"null\",\"string\"],\"default\":\"null\"}]}"
	)

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})
	if err != nil {
		log.Fatal(err)
	}

	var s SensorEReading

	consumerJS := pulsar.NewJSONSchema(exampleSchemaDef, nil)
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:                       "sensor_readings_enriched",
		SubscriptionName:            "sub-1",
		Schema:                      consumerJS,
		SubscriptionInitialPosition: pulsar.SubscriptionPositionLatest,
	})

	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		err = msg.GetSchemaValue(&s)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
			msg.ID(), string(msg.Payload()))

		fmt.Printf("Received message StationID: %#v", s.SensorId)

		AddNewsSource(s)
	}

}
