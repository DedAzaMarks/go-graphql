package main

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	"log"
	"os"
)

func main() {
	client := graphql.NewClient("http://localhost:8080")
	for {
		var s string
		_, err := fmt.Fscanf(os.Stdin, "%s", s)
		var respData graphql.Request
		err = client.Run(context.Background(), graphql.NewRequest(s), respData)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("RESPOND: %s", respData)
	}
}
