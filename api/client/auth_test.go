package client_test

import (
	"context"
	"log"

	"github.com/reddec/api-notes/api/client"
)

func ExampleNewClient() {
	notes, err := client.NewClient("https://example.com", client.HeaderToken("deadbeaf"))
	if err != nil {
		// panic is used for illustration only
		panic("create notes client: " + err.Error())
	}
	note, err := notes.CreateNote(context.Background(), &client.DraftMultipart{
		Title:  "Hello",
		Text:   "## hello world\nThis is sample text",
		Author: client.NewOptString("demo"),
	})
	if err != nil {
		panic("create note: " + err.Error())
	}

	log.Println("Note ID:", note.ID)
	log.Println("Note URL:", note.PublicURL)
}
