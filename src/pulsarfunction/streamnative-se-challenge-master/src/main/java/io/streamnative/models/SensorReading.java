package io.streamnative.models;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class SensorReading {
    private String sensorId;
    private int stationId;
    private String status;
    private String startupTime;
    private String eventTime;
    private String reading;

    public String getSensorId() {
        return sensorId;
    }

    @Override
    public String toString() {
        return "SensorReading{" +
                "sensorId='" + sensorId + '\'' +
                ", stationId=" + stationId +
                ", status='" + status + '\'' +
                ", startupTime='" + startupTime + '\'' +
                ", eventTime='" + eventTime + '\'' +
                ", reading='" + reading + '\'' +
                '}';
    }

    public void setSensorId(String sensorId) {
        this.sensorId = sensorId;
    }

    public int getStationId() {
        return stationId;
    }

    public void setStationId(int stationId) {
        this.stationId = stationId;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public String getStartupTime() {
        return startupTime;
    }

    public void setStartupTime(String startupTime) {
        this.startupTime = startupTime;
    }

    public String getEventTime() {
        return eventTime;
    }

    public void setEventTime(String eventTime) {
        this.eventTime = eventTime;
    }

    public String getReading() {
        return reading;
    }

    public void setReading(String reading) {
        this.reading = reading;
    }
}

