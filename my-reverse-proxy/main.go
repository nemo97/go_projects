package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

/*
	Structs
*/

type AppConfig struct {
	//Context string
	Target string
	//Source  string
}

type Config struct {
	Rules map[string]AppConfig
}

/*
	Utilities
*/

// Get env var or default
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

/*
	Getters
*/

// Get the port to listen on
func getListenAddress() string {
	port := getEnv("MY_REVERSE_PROXY_PORT", "8083")
	return ":" + port
}

/*
	Logging
*/

// Log the env variables required for a reverse proxy
func logSetup() {
	log.Printf("Server will run on: %s\n", getListenAddress())
}

/*
	Reverse Proxy Logic
*/

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	//req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Header.Set("X-Forwarded-Proto", "http")
	//req.Host = url.Host
	//req.Host = "localhost"

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}
func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}
func before(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request, config Config) {
	contextPathWithoutSlace := before(html.EscapeString(req.URL.Path)[1:], "/")
	fmt.Println("PATH map , ", html.EscapeString(req.URL.Path), " -> ", contextPathWithoutSlace, " -> ", config.Rules[contextPathWithoutSlace])
	var appConfig AppConfig
	appConfig = config.Rules[contextPathWithoutSlace]
	var appDefaultConfig AppConfig
	appDefaultConfig = config.Rules["default-host"]

	url := ""
	if &appConfig != nil && appConfig.Target != "" {
		url = appConfig.Target
	} else if &appDefaultConfig != nil {
		fmt.Println("Using default config ", appDefaultConfig)
		url = appDefaultConfig.Target
	}
	fmt.Println("Route , ", html.EscapeString(req.URL.Path), " -> ", url)
	serveReverseProxy(url, res, req)
}

/*
	Entry
*/

func main() {
	var config Config

	filename, _ := filepath.Abs("./config.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Problem reading config file. error: %v", err)
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Problem parsing config file %s. error: %v", yamlFile, err)
		panic(err)
	}
	fmt.Println("--- config rules\n", config)
	// Log setup values
	logSetup()

	// start server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleRequestAndRedirect(w, r, config)
	})
	// http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
		panic(err)
	}
}
