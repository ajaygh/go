package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type GeoTree struct {
	Top   Country `json:"text"`
	Level int     `json:"level_id"`
}

type Country struct {
	Node
	State []States `json:"nodes"`
}

type States struct {
	Node
	City []City `json:"nodes"`
}

type City struct {
	Node
	Zones []Zones `json:"nodes"`
}

type Zones struct {
	Node
	Street []Streets `json:"nodes"`
}

type Streets struct {
	Node
	Lights []Lamps `json:"nodes"`
}

type Lamps struct {
	Node
	Lat    string `json:"lat"`
	Lng    string `json:"lng"`
	Status string `json:"status"`
}

type Node struct {
	ID    int    `json:"id"`
	Name  string `json:"text"`
	Level int    `json:"level_id"`
}

var geomap = make(map[string]Country)

func main() {
	initGeoParams()
}

func initGeoParams() {

	light1 := Lamps{Node: Node{ID: 1, Name: "Light1", Level: 6}, Lat: "12.9159791", Lng: "77.61606649999999", Status: "Green"}
	light2 := Lamps{Node: Node{ID: 2, Name: "Light2", Level: 6}, Lat: "12.9158791", Lng: "77.61616649999999", Status: "Green"}
	light3 := Lamps{Node: Node{ID: 3, Name: "Light3", Level: 6}, Lat: "12.9158791", Lng: "77.61606649999999", Status: "Red"}
	light4 := Lamps{Node: Node{ID: 4, Name: "Light4", Level: 6}, Lat: "12.9158691", Lng: "77.61606649999999", Status: "Green"}

	light5 := Lamps{Node: Node{ID: 5, Name: "Light5", Level: 6}, Lat: "12.8159791", Lng: "77.61606649999999", Status: "Green"}
	light6 := Lamps{Node: Node{ID: 6, Name: "Light6", Level: 6}, Lat: "12.8158791", Lng: "77.61606649999999", Status: "Red"}
	light7 := Lamps{Node: Node{ID: 7, Name: "Light7", Level: 6}, Lat: "12.8159791", Lng: "77.61616649999999", Status: "Red"}
	light8 := Lamps{Node: Node{ID: 8, Name: "Light8", Level: 6}, Lat: "12.8158791", Lng: "77.61616649999999", Status: "Green"}

	street1 := Streets{Node: Node{ID: 1, Name: "1st corss", Level: 5}}
	street2 := Streets{Node: Node{ID: 2, Name: "2st corss", Level: 5}}
	street3 := Streets{Node: Node{ID: 3, Name: "3st corss", Level: 5}}
	street4 := Streets{Node: Node{ID: 4, Name: "4st corss", Level: 5}}
	street1.Lights = append(street1.Lights, light1)
	street1.Lights = append(street1.Lights, light2)
	street2.Lights = append(street2.Lights, light3)
	street2.Lights = append(street2.Lights, light4)
	street3.Lights = append(street3.Lights, light5)
	street3.Lights = append(street3.Lights, light6)
	street4.Lights = append(street4.Lights, light7)
	street4.Lights = append(street4.Lights, light8)
	zone1 := Zones{Node: Node{ID: 1, Name: "Kormangala", Level: 4}}
	zone1.Street = append(zone1.Street, street1)
	zone1.Street = append(zone1.Street, street2)
	zone2 := Zones{Node: Node{ID: 2, Name: "XYZ", Level: 4}}
	zone2.Street = append(zone2.Street, street3)
	zone2.Street = append(zone2.Street, street4)
	city1 := City{Node: Node{ID: 1, Name: "Bangalore", Level: 3}}
	city1.Zones = append(city1.Zones, zone1)
	city2 := City{Node: Node{ID: 2, Name: "City2", Level: 3}}
	city2.Zones = append(city2.Zones, zone2)
	state1 := States{Node: Node{ID: 1, Name: "Karnataka", Level: 2}}
	state2 := States{Node: Node{ID: 2, Name: "State2", Level: 2}}
	state1.City = append(state1.City, city1)
	state2.City = append(state2.City, city2)
	temp1 := Country{Node: Node{ID: 1, Name: "INDIA", Level: 1}}
	temp1.State = append(temp1.State, state1)
	counArr := []Country{temp1}

	res := json.NewEncoder(os.Stdout).Encode(counArr)
	fmt.Println(res)
	geomap["560034"] = temp1
	//geomap["560076"] = temp2
}
