package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

// declare global var `c`
var c Config

// create `initConfig` for unmarshalling .yml
func initConfig() {

	yamlFile, err := ioutil.ReadFile("./configs/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatal(err)
	}
}

// create a `serveDynamic` funtcion
func serveDynamic(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the root path")
}

// create a `serveStatic` funtcion
func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}

func main() {
	initConfig()

	// log the message on starting server
	log.Printf("Server starting on port %v\n", 8080)

	http.HandleFunc("/", serveDynamic)
	http.HandleFunc("/static", serveStatic)

	// listen and serve
	http.ListenAndServe(c.Server.Port, nil)

}