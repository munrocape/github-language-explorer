package main
 
import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"

    "gopkg.in/yaml.v2"
)

type Languages struct {
	Languages []Language
}

type Language struct {

}

func main() {
    
	// Get YAML representing all languages
    response, err := http.Get("https://raw.githubusercontent.com/github/linguist/master/lib/linguist/languages.yml")
    var unmarshalled_yaml []byte
    if err != nil {
        fmt.Printf("%s\n", err)
        os.Exit(1)
    }
    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Printf("%s\n", err)
        os.Exit(1)
    }
    unmarshalled_yaml = []byte(contents)

    // Parse YAML into struct representation
    
    // Write to JSON

}