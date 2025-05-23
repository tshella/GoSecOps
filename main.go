package main

import (
	"gosecops/api"
)

func main() {
	r := api.SetupRouter()
	r.Run(":8080")
}
