package main

import (
	"edulx/qbot/api"
	"edulx/qbot/db"
)

func main() {
	api.Run()
}

func init() {
	err := db.InitClirnt()
	if err != nil {
		return
	}
}
