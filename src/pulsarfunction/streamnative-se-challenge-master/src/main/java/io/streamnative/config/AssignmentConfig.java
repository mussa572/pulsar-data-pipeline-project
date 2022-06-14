
package io.streamnative.config;

public class AssignmentConfig {
    public static final String SERVICE_URL  = "pulsar://localhost:6650";
    public static final String INPUT_TOPIC  = "sensor_readings";
    public static final String OUTPUT_TOPIC = "sensor_readings_enriched";
    public static final String SCHEMA_TYPE  = "JSON"; // change this to avro if you wanna use avro schema
}

