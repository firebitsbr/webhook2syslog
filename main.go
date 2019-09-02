package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"log/syslog"
	"net/http"
	"strconv"
)

var hostPtr string
var portPtr int
var messagePtr string

func isJson(s string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(s), &js) == nil
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("[Error] Message body not parsed. " + err.Error())
		return
	}

	bs := string(body)

	if isJson(bs) {
		logMessage(bs)
	} else {
		log.Println("[Error] Invalid JSON: " + bs)
	}
}

func logMessage(b string) {
	logwriter, e := syslog.New(syslog.LOG_NOTICE, messagePtr)

	if e == nil {
		log.SetOutput(logwriter)
	}

	log.Print(b)
}

func main() {
	flag.StringVar(&hostPtr,"host", "localhost", "HTTP server address (IPv4 or hostname).")
	flag.IntVar(&portPtr, "port", 5001, "HTTP listener port.")
	flag.StringVar(&messagePtr, "message", "[TheHive Webhook]", "Syslog message prefix.")
	flag.Parse()

	http.HandleFunc("/", handleMessage)
	log.Fatal(http.ListenAndServe(hostPtr + ":" + strconv.Itoa(portPtr), nil))
}
