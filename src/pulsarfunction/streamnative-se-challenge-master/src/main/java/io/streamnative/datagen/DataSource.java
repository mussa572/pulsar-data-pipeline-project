package io.streamnative.datagen;
import io.streamnative.models.StationData;
import java.sql.Timestamp;
import java.util.Map;
import java.util.Random;
import java.util.function.Function;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class DataSource {
    private static final Random random = new Random();

    public static final Map<Integer, StationData> stationData = initStationData();

    private static Map<Integer, StationData> initStationData() {
        return IntStream.range(1, 11).mapToObj(i -> {
            return new StationData(
                    i,
                    "station-" + i,
                    new Timestamp(System.currentTimeMillis() - random.nextInt(100000))
            );
        }).collect(Collectors.toMap(StationData::getStationId, Function.identity()));
    }
}

