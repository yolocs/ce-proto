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

	// 2020/05/06 21:10:10 File upload event: Validation: valid
	// Context Attributes,
	// 	specversion: 1.0
	// 	type: com.example.samplesvc.FileUploadedEvent
	// 	source: com.example.samplesvc
	// 	subject: com.example.samplesvc.FileUploadedEvent
	// 	id: f5c3f979-c822-47f3-aec4-724180a45230
	// 	datacontenttype: application/json
	// Extensions,
	// 	ext: txt
	// Data (binary),
	// 	{
	// 		"name": "sample.txt",
	// 		"path": "/folder/foo",
	// 		"ext": "txt"
	// 	}

	// 2020/05/06 21:10:10 Paymend received event: Validation: valid
	// Context Attributes,
	// 	specversion: 1.0
	// 	type: com.example.samplesvc.PaymentReceivedEvent
	// 	source: com.example.samplesvc
	// 	subject: com.example.samplesvc.PaymentReceivedEvent
	// 	id: b5dfc798-7021-44ca-aa10-1d687bcdc50e
	// 	datacontenttype: application/json
	// Extensions,
	// 	user: user1
	// 	usergroup: ceprototesters
	// Data (binary),
	// 	{
	// 		"user": "user1",
	// 		"userGroup": "ceprototesters",
	// 		"amount": "$100"
	// 	}

}
