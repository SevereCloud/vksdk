package foaf_test

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/SevereCloud/vksdk/foaf"
)

func ExampleGetGroup() {
	ctx := context.Background()

	person, err := foaf.GetGroup(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(person)
}

func ExampleGetPerson() {
	ctx := context.Background()

	person, err := foaf.GetPerson(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(person)
}

func ExampleGetPerson_customHTTP() {
	ctx := context.Background()

	// Use the custom HTTP client
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, foaf.HTTPClient, httpClient)

	person, err := foaf.GetPerson(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(person)
}
