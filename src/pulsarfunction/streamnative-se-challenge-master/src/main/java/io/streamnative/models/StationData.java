package io.streamnative.models;

import java.sql.Timestamp;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class StationData {
    private int stationId;
    private String stationName;
    private Timestamp statupTime;
}
