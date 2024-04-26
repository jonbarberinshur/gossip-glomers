package main

import (
	"encoding/json"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"log"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("generate", func(msg maelstrom.Message) error {
		// Unmarshal the message body as a loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update the message type to return back.
		body["type"] = "generate_ok"
		body["id"] = "1" // this needs to be a globally unique value each time

		// Echo the original message back with the updated message type and content.
		return n.Reply(msg, body)
	})

	n.Handle("init", func(msg maelstrom.Message) error {
		var b maelstrom.InitMessageBody
		if err := json.Unmarshal(msg.Body, &b); err != nil {
			return err
		}

		log.Printf("init body %+v", b)

		return nil
	})

	n.Handle("topology", func(msg maelstrom.Message) error {
		type topology struct {
			Topology map[string][]string `json:"topology"`
		}

		var t topology
		if err := json.Unmarshal(msg.Body, &t); err != nil {
			return err
		}

		log.Printf("topology %+v", t)

		return nil
	})

	if err := n.Run(); err != nil {
		log.Fatalf("Unable to run the handler : %v", err)
	}

}
