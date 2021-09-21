package main

import(
	"encoding/json"
	"os"
)


func main() {

	encoder := json.NewEncoder(os.Stdout)

	encoder.Encode([]int { 1, 2, 3 })
}