package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/moguriso/thp"
)

func main() {
	macAddr := os.Getenv("DOCODEMO_MAC")
	token := os.Getenv("DOCODEMO_TOKEN")
	endpoint := thp.BuildEndpoint(macAddr, token)
	jsonByte := thp.GetRequest(endpoint)
	sd := thp.ParseJson(jsonByte)
	js, _ := json.Marshal(sd)
	log.Println(string(js))
}
