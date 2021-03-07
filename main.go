package main

import (
	"context"
	"fmt"

	"court.com/src/pkg/mongodb"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	fmt.Println("Hello")

	cli, err := mongodb.InitClient()
	if err != nil {
		fmt.Println(err)
	}
	cli.Ping(context.Background(), readpref.Primary())
}
