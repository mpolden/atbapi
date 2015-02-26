package api

import (
	"github.com/martinp/atbapi/atb"
	"strconv"
)

type BusStops struct {
	Stops []BusStop `json:"stops"`
}

type BusStop struct {
	StopId      int     `json:"stopId"`
	NodeId      int     `json:"nodeId"`
	Description string  `json:"description"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	MobileCode  string  `json:"mobileCode"`
	MobileName  string  `json:"mobileName"`
}

type Departures struct {
	TowardsCentrum bool        `json:"isGoingTowardsCentrum"`
	Departures     []Departure `json:"departures"`
}

type Departure struct {
	LineId                  string `json:"line"`
	RegisteredDepartureTime string `json:"registeredDepartureTime"`
	ScheduledDepartureTime  string `json:"scheduledDepartureTime"`
	Destination             string `json:"destination"`
}

func convertBusStop(s atb.BusStop) (BusStop, error) {
	nodeId, err := strconv.Atoi(s.NodeId)
	if err != nil {
		return BusStop{}, err
	}
	longitude, err := strconv.ParseFloat(s.Longitude, 64)
	if err != nil {
		return BusStop{}, err
	}
	latitude := float64(s.Latitude)
	return BusStop{
		StopId:      s.StopId,
		NodeId:      nodeId,
		Description: s.Description,
		Longitude:   longitude,
		Latitude:    latitude,
		MobileCode:  s.MobileCode,
		MobileName:  s.MobileName,
	}, nil
}

func convertBusStops(s atb.BusStops) (BusStops, error) {
	stops := make([]BusStop, 0, len(s.Stops))
	for _, stop := range s.Stops {
		converted, err := convertBusStop(stop)
		if err != nil {
			return BusStops{}, err
		}
		stops = append(stops, converted)
	}
	return BusStops{Stops: stops}, nil
}

func convertForecast(f atb.Forecast) (Departure, error) {
	return Departure{
		LineId:                  f.LineId,
		Destination:             f.Destination,
		RegisteredDepartureTime: f.RegisteredDepartureTime,
		ScheduledDepartureTime:  f.ScheduledDepartureTime,
	}, nil
}

func convertForecasts(f atb.Forecasts) (Departures, error) {
	towardsCentrum := false
	if len(f.Nodes) > 0 {
		nodeId, err := strconv.Atoi(f.Nodes[0].NodeId)
		if err != nil {
			return Departures{}, err
		}
		towardsCentrum = (nodeId/1000)%2 == 1
	}
	departures := make([]Departure, 0, len(f.Forecasts))
	for _, forecast := range f.Forecasts {
		departure, err := convertForecast(forecast)
		if err != nil {
			return Departures{}, err
		}
		departures = append(departures, departure)
	}
	return Departures{
		TowardsCentrum: towardsCentrum,
		Departures:     departures,
	}, nil
}