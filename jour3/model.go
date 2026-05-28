package main

type Station struct {
	Id           string
	Country      string
	Altitude     int
	DeviceModel  string
	Location     Coordinate
	Observations []Observation
}

type Coordinate struct {
	Lat  float64
	Long float64
}

type Observation struct {
	TimeStamp   string
	Temperature float64
	Condition   string
	Wind        Wind
	Notes       *string
}

type Wind struct {
	Speed     float64
	Direction int
}
