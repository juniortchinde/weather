package json

type JsonStation struct {
	country  string
	altitude int
	obs      Observation
	location Coordinate
}

type Coordinate struct {
	lat  float64
	long float64
}

type Observation struct {
	Temperature float64
	Condition   string
	Wind        Wind
}

type Wind struct {
	Speed     float64
	direction float64
}
