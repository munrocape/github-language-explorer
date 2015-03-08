package main

import (
	//"fmt"
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"sort"
	"strconv"

	"gopkg.in/yaml.v2"
)

type Language struct {
	Name  string
	Color Color
}

func NewLanguage(name string, hex string) Language {
	color := NewColor(hex)
	l := Language{Name: name, Color: color}
	return l
}

type Color struct {
	Hex        string
	R          float64
	G          float64
	B          float64
	MaxRGB     float64
	MinRGB     float64
	Hue        float64
	Saturation float64
	Lightness  float64
}

func NewColor(hex string) Color {
	c := new(Color)
	c.Hex = hex
	c.R, c.G, c.B = hexToRGB(hex)
	c.MaxRGB = ternaryMax(c.R, c.G, c.B)
	c.MinRGB = ternaryMin(c.R, c.G, c.B)
	c.Hue = hueFromRGB(c.MaxRGB, c.MinRGB, c.R, c.G, c.B)
	return *c
}

// Sorting implementation for language list
type Languages []Language

func (slice Languages) Len() int {
	return len(slice)
}

func (slice Languages) Less(i, j int) bool {
	return slice[i].Color.Hue < slice[j].Color.Hue
	//if slice[i].Color.Hue == slice[j].Color.Hue {
	//	return slice[i].Color.Lightness < slice[j].Color.Lightness
	//} else {
	//	return slice[i].Color.Hue < slice[j].Color.Hue
	//}
}

func (slice Languages) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ternaryMax(a float64, b float64, c float64) float64 {
	if (a > b) && (a > c) {
		return a
	} else if (b > a) && (b > c) {
		return b
	} else {
		return c
	}
}

func ternaryMin(a float64, b float64, c float64) float64 {
	if (a < b) && (a < c) {
		return a
	} else if (b < a) && (b < c) {
		return b
	} else {
		return c
	}
}

func hexToRGB(str string) (float64, float64, float64) {
	var r, g, b int64
	if len(str) == 4 {
		r, _ = strconv.ParseInt(str[1:3], 16, 0)
		g, _ = strconv.ParseInt(str[1:3], 16, 0)
		b, _ = strconv.ParseInt(str[1:3], 16, 0)
	} else {
		r, _ = strconv.ParseInt(str[1:3], 16, 0)
		g, _ = strconv.ParseInt(str[3:5], 16, 0)
		b, _ = strconv.ParseInt(str[5:7], 16, 0)
	}
	return float64(r), float64(g), float64(b)
}

func hueFromRGB(max float64, min float64, r float64, g float64, b float64) float64 {
	delta := max - min
	if min == 0 {
		return float64(0)
	}
	if (max - min) == 0 {
		return float64(0)
	}
	var hue float64
	if max == r {
		hue = (math.Mod((g-b)/delta, 6))
	} else if max == g {
		hue = (2 + (b-r)/delta)
	} else {
		hue = (4 + (r-g)/delta)
	}
	return hue * 60
}

func lightnessFromRGB(max float64, min float64, r float64, g float64, b float64) float64 {
	return 0.30*r + 0.59*g + 0.11*b
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
	colorMap := make(map[string]Language)

	for k, v := range yamlMap {
		if val, ok := v["color"]; ok { // color exists
			if str, ok := val.(string); ok { // string type check (required)
				colorMap[k] = NewLanguage(k, str)
			}
		}
	}

	languages := make(Languages, 0, len(colorMap))
	for k := range colorMap {
		languages = append(languages, colorMap[k])
	}
	sort.Sort(languages)

	colorJSON, err := json.MarshalIndent(languages, "", "    ")
	check(err)
	allJSON, err := json.MarshalIndent(yamlMap, "", "    ")
	check(err)

	// Write to file
	err = ioutil.WriteFile("color-info.json", []byte("color_data = "+string(colorJSON)), 0644)
	check(err)
	err = ioutil.WriteFile("all-info.json", []byte("all_data = "+string(allJSON)), 0644)
	check(err)

}
