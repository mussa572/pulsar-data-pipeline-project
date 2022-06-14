package io.streamnative.models;

import java.sql.Timestamp;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class ReadingEnriched {
    private String sensorId;
    private int stationId;
    private String stationName;
    private String status;
    private String startupTime;
    private String eventTime;
    private String reading;
}
