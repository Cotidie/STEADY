package main

import (
	"encoding/json"
	"os"
)

func main() {
	jsonBytes, _ := os.ReadFile("./trace.txt")

	json.Unmarshal()
}
