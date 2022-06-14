
package io.streamnative.function;
import io.streamnative.datagen.DataSource;
import io.streamnative.models.ReadingEnriched;
import io.streamnative.models.SensorReading;
import io.streamnative.models.StationData;
import org.apache.pulsar.functions.api.Context;
import org.apache.pulsar.functions.api.Function;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import java.util.Map;

public class EnrichmentFunc implements Function<SensorReading, ReadingEnriched> {
    private static final Logger logger = LoggerFactory.getLogger(EnrichmentFunc.class);

    public static void main(String[] args) throws Exception{

    }

    private final Map<Integer, StationData> inMemoryStationCache = DataSource.stationData;

    @Override
    public ReadingEnriched process(SensorReading input, Context context) throws Exception {

        logger.info("Received reading: " + input);

        // Check that the sensor is in 'RUNNING' status

        if (input.getStatus().equals("Running") )  {
        ReadingEnriched readingEnriched = new ReadingEnriched();
        readingEnriched.setStationId(input.getStationId());
        readingEnriched.setSensorId(input.getSensorId());
        readingEnriched.setEventTime(input.getEventTime());
        readingEnriched.setStatus(input.getStatus());
        readingEnriched.setReading(input.getReading());
        readingEnriched.setStartupTime(input.getStartupTime());

        // Create a new record and send it downstream

        readingEnriched.setStationName(inMemoryStationCache.get(input.getStationId()).getStationName());
            return readingEnriched;

        }
return null;


    }



}

