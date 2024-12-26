package main

import (
	"beto0607.com/blober/src/data"
)

func main() {
	data.ConnectToDB()
	defer func() {
		data.DisconnectDB()
	}()
}
