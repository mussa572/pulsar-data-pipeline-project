package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/google/uuid"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type SensorReading struct {
	SensorId    string    `json:"sensorId"`
	StationId   int       `json:"stationId"`
	Status      string    `json:"status"`
	StartupTime time.Time `json:"startupTime"`
	EventTime   time.Time `json:"eventTime"`
	Reading     string    `json:"reading"`
}

type StationData struct {
	StationId   int    `json:"stationId"`
	StationName string `json:"stationName"`
	StartupTime time.Time
}

func main() {

	iterations := flag.Int("it", 1000, "number of iteration")
	stationCount := flag.Int("st", 11, "number of station")
	flag.Parse()
	// Create the Seed based on the current
	seed := rand.NewSource(time.Now().UnixNano())
	//Passing the Random number seed
	random := rand.New(seed)

	stationData := initStationData(*stationCount, random)
	for i := 1; i <= *iterations; i++ {
		time.Sleep(time.Millisecond * time.Duration(200+random.Int63n(500)))
		//Picking the random station

		station := stationData[rand.Intn(*stationCount)]

		// Generating the random Data

		reading := math.Round(random.NormFloat64() + 40)

		sensorReading := SensorReading{
			SensorId:    uuid.New().String(),
			StationId:   station.StationId,
			Status:      "Running",
			StartupTime: station.StartupTime,
			EventTime:   time.Now(),
			Reading:     strconv.FormatFloat(reading, 'E', -1, 32),
		}
		if random.Float64() < 0.2 {
			sensorReading.Status = "Inactive"
		}

		// Create Sensor Producer
		producer := createSensorProducer()

		_, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
			Value: &SensorReading{
				SensorId:    sensorReading.SensorId,
				StationId:   sensorReading.StationId,
				Status:      sensorReading.Status,
				StartupTime: sensorReading.StartupTime,
				EventTime:   sensorReading.EventTime,
				Reading:     sensorReading.Reading,
			},
		})
		if err != nil {
			log.Fatal(err)
		}

	}

}

func initStationData(stationCount int, random *rand.Rand) map[int]StationData {
	mp := make(map[int]StationData)
	for i := 1; i < stationCount; i++ {
		mp[i] = StationData{
			i,
			fmt.Sprintf("station-%v", i),
			time.Now().Add(-time.Millisecond * time.Duration(random.Int63n(100000))),
		}

	}
	return mp
}

func formatTime(t time.Time) string {
	formatted := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	return formatted
}

func createSensorProducer() pulsar.Producer {

	var (
		exampleSchemaDef = "{\"type\":\"record\",\"name\":\"Sensor\",\"namespace\":\"io.streamnative.models\"," +
			"\"fields\":[{\"name\":\"sensorId\",\"type\":\"string\"},{\"name\":\"stationId\",\"type\":\"int\"},{\"name\":\"status\",\"type\":\"string\"},{\"name\":\"startuptime\",\"type\":\"string\"},{\"name\":\"eventime\",\"type\":\"string\"},{\"name\":\"reading\",\"type\":\"string\"}]}"
	)
	pulsarURL := "pulsar://localhost:6650"

	client, err := connectProducer(pulsarURL)
	if err != nil {
		log.Fatal(err)
	}

	jsonSchemaWithProperties := pulsar.NewJSONSchema(exampleSchemaDef, nil)

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic:  "sensor_readings",
		Schema: jsonSchemaWithProperties,
	})

	if err != nil {
		log.Fatal(err)
	}

	return producer
}

func connectProducer(URL string) (pulsar.Client, error) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	return client, nil

}
