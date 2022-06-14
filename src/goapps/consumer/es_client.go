package main

import (
	"context"
	"github.com/olivere/elastic/v7"
	"time"
)

var (
	EsClient *elastic.Client
)

func Init() {

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://localhost:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
	)
	if err != nil {
		panic(err)
	}
	EsClient = client
}

func AddNewsSource(s SensorEReading) error {
	ctx := context.Background()

	_, err := EsClient.Index().Index("sensor_readings_enriched").BodyJson(s).Do(ctx)

	if err != nil {

		return err

	}

	return nil
}
