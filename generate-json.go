package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Get YAML representing all languages
	var unmarshalled_yaml []byte
	response, err := http.Get("https://raw.githubusercontent.com/github/linguist/master/lib/linguist/languages.yml")
	check(err)
	
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	check(err)
	
	unmarshalled_yaml = []byte(contents)

	// Parse YAML into map representation

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(unmarshalled_yaml, &m)
	check(err)

	fmt.Printf("%#v\n", m)

	// Convert map to JSON

	// Write to file


}
