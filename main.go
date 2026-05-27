package main

import (
	"fmt"
	"weather/json"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	jsonStations, _ := json.ExtractJson("./data/weather_data.json")

	for _, jsonStation := range jsonStations.Stations {
		fmt.Println(jsonStation.Country)
	}
	fmt.Println(len(jsonStations.Stations))
}
