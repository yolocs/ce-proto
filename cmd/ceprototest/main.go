package main

import (
	"log"
	"os"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/yolocs/ce-proto/pkg/ceproto"
	"github.com/yolocs/ce-proto/pkg/eventtypes/examples"
)

func main() {
	et1 := &examples.FileUploadedEvent{
		Name: "sample.txt",
		Path: "/folder/foo",
		Ext:  "txt",
	}
	et2 := &examples.PaymentReceivedEvent{
		Amount:    "$100",
		User:      "user1",
		UserGroup: "ceprototesters",
	}

	e1 := event.New()
	if err := ceproto.Marshal(et1, &e1); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Printf("File upload event: %+v\n", e1)

	e2 := event.New()
	if err := ceproto.Marshal(et2, &e2); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Printf("Paymend received event: %+v\n", e2)

	// Output

	// 2020/05/01 21:57:56 File upload event: Validation: valid
	// Context Attributes,
	// 	specversion: 1.0
	// 	type: com.example.samplesvc.FileUploadedEvent
	// 	source: com.example.samplesvc
	// 	subject: com.example.samplesvc.FileUploadedEvent
	// 	id: 01ffbef3-89e8-40a9-9212-f10ce24e4971
	// 	datacontenttype: application/octet-stream
	// Extensions,
	// 	ext: txt
	// Data (binary),

	// sample.txt
	// 					/folder/footxt

	// 2020/05/01 21:57:56 Paymend received event: Validation: valid
	// Context Attributes,
	// 	specversion: 1.0
	// 	type: com.example.samplesvc.PaymentReceivedEvent
	// 	source: com.example.samplesvc
	// 	subject: com.example.samplesvc.PaymentReceivedEvent
	// 	id: c5d14874-f33b-4069-8898-a0b3470987a7
	// 	datacontenttype: application/octet-stream
	// Extensions,
	// 	user: user1
	// 	usergroup: ceprototesters
	// Data (binary),

	// user1ceprototesters$100

}
