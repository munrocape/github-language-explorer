package main

import (
	//"fmt"
	"strconv"
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

func hexToRGB(str string) (int64, int64, int64) {
	var r, g, b int64
	if(len(str) == 4){
		r, _ = strconv.ParseInt(str[1:3], 16, 0)
		g, _ = strconv.ParseInt(str[1:3], 16, 0)
		b, _ = strconv.ParseInt(str[1:3], 16, 0)
	} else {
		r, _ = strconv.ParseInt(str[1:3], 16, 0)
		g, _ = strconv.ParseInt(str[3:5], 16, 0)
		b, _ = strconv.ParseInt(str[5:7], 16, 0)
	}
	return r, g, b
}

func createLangStruct(name string, hex string) *Language {
	//r, g, b := hexToRGB(hex)
	//fmt.Printf("lang: %s hex: %s rgb: %d %d %d\n", name, hex, r, g, b)
	return new(Language)
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
	// colorMap := make(map[string] *Language)
	colorMap := make(map[string] string)

	for k, v := range yamlMap {
		if val, ok := v["color"]; ok { // color exists
			if str, ok := val.(string); ok { // string type check (required)
				colorMap[k] = str// createLangStruct(k, str)
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
