package main

import (
	"fmt"
	"github.com/SevereCloud/vksdk/adapter"
	"github.com/go-joe/joe"
	"github.com/go-joe/redis-memory"
	"os"
	"strings"
)

type ExampleBot struct {
	*joe.Bot
}

func main() {
	b := &ExampleBot{
		Bot: joe.New(
			"example",
			redis.Memory("localhost:6379"),
			adapter.Adapter(os.Getenv("BOT_TOKEN")),
		),
	}

	b.Respond("(.+)", b.Taco)
	b.Respond("remember (.+) is (.+)", b.Remember)
	b.Respond("what is (.+)", b.WhatIs)

	if err := b.Run(); err != nil {
		b.Logger.Fatal(err.Error())
	}
}

func (b *ExampleBot) Taco(msg joe.Message) error {
	text := msg.Matches[0]

	if strings.Contains(text, "тако") {
		msg.Respond("🌮")
	}

	return nil
}

func (b *ExampleBot) Remember(msg joe.Message) error {
	key, value := msg.Matches[0], msg.Matches[1]
	msg.Respond("OK, I'll remember %s is %s", key, value)

	return b.Store.Set(key, value)
}

func (b *ExampleBot) WhatIs(msg joe.Message) error {
	key := msg.Matches[0]
	value := ""

	ok, err := b.Store.Get(key, &value)
	if err != nil {
		return fmt.Errorf("failed to retrieve key %q from brain: %w", key, err)
	}

	if ok {
		msg.Respond("%s is %s", key, value)
	} else {
		msg.Respond("I do not remember %q", key)
	}

	return nil
}
