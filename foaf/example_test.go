package foaf_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SevereCloud/vksdk/v2/foaf"
)

func ExampleGetGroup() {
	ctx := context.Background()

	group, err := foaf.GetGroup(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(group.Name)
	// Output: ВКонтакте API
}

func ExampleGetPerson() {
	ctx := context.Background()

	person, err := foaf.GetPerson(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(person.FirstName)
	// Output: Павел
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

	fmt.Println(person.FirstName)
	// Output: Павел
}
