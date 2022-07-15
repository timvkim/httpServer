package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

// create a `HttpHanfler` struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// create response data
	data := []byte("Hello, World!\n")

	// write `data` to response
	res.Write(data)
}

func initConfig() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./configs")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

}

func main() {

	// create a new handler
	handler := HttpHandler{}

	// log the message on starting server
	log.Printf("Server starting on port %v\n", 8080)

	initConfig()
	port := viper.GetString(":port")
	// listen and serve
	http.ListenAndServe(port, handler)

}
