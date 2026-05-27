package main

type Station struct {
	Country      string
	Altitude     int
	DeviceModel  string
	Location     Coordinate
	Observations []Observation
	Notes        *string
}

type Coordinate struct {
	Lat  float64
	Long float64
}

type Observation struct {
	Temperature float64
	Condition   string
	Wind        Wind
}

type Wind struct {
	Speed     float64
	Direction int
}
