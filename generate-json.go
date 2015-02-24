package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Language struct {
	Name string
	Hex string
	R float64 // To find the hue, 1 / {R, G, B} is required
	G float64 // Typical uint8 assignment is therefore not ideal
	B float64
	Hue float64

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Get YAML representing all languages
	response, err := http.Get("https://raw.githubusercontent.com/github/linguist/master/lib/linguist/languages.yml")
	check(err)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	check(err)

	unmarshalled_yaml := []byte(contents)

	// Parse YAML into map representation
	yamlMap := make(map[string]map[string]interface{})
	err = yaml.Unmarshal(unmarshalled_yaml, &yamlMap)
	check(err)

	// Convert map to JSON
	colorMap := make(map[string]string)

	for k, v := range yamlMap {
		if val, ok := v["color"]; ok { // color exists
			if str, ok := val.(string); ok { // string type check (required)
				colorMap[k] = str
			}
		}
	}

	colorJSON, err := json.MarshalIndent(colorMap, "", "    ")
	check(err)
	allJSON, err := json.MarshalIndent(yamlMap, "", "    ")
	check(err)

	// Write to file
	err = ioutil.WriteFile("color-info.json", []byte("color_data = " + string(colorJSON)), 0644)
	check(err)
	err = ioutil.WriteFile("all-info.json", []byte("all_data = " + string(allJSON)), 0644)
	check(err)

}
