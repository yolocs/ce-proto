package main

import (
	"log"
	"os"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/yolocs/ce-proto/pkg/ceproto"
	"github.com/yolocs/ce-proto/pkg/eventtypes/samples/hello/v1alpha1"
	"github.com/yolocs/ce-proto/pkg/eventtypes/samples/hello/v1beta1"
)

func main() {
	helloSent := &v1beta1.Hello{
		GreetingWord:   "hello",
		GreetingTarget: "Chen",
		GreetingTimes:  3,
	}
	log.Printf("Hello (v1beta1) sent: %s\n", helloSent.String())

	e := event.New()
	if err := ceproto.Marshal(helloSent, &e); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Printf("CloudEvent sent: %+v\n", e)

	var helloReceived1 v1alpha1.Hello
	if err := ceproto.Unmarshal(&e, &helloReceived1); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Printf("Hello (v1alpha1) received: %s\n", helloReceived1.String())

	var helloReceived2 v1beta1.Hello
	if err := ceproto.Unmarshal(&e, &helloReceived2); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Printf("Hello (v1beta1) received: %s\n", helloReceived2.String())

	// Output

	// 2020/05/12 16:57:40 Hello (v1beta1) sent: greeting_word:"hello"  greeting_target:"Chen"  greeting_times:3
	// 2020/05/12 16:57:40 CloudEvent sent: Validation: valid
	// Context Attributes,
	// 	specversion: 1.0
	// 	type: sample.hello.v1beta1.Hello
	// 	source: sample.hello.v1beta1
	// 	subject: sample.hello.v1beta1.Hello
	// 	id: 9d5b58a9-0a53-40b6-baf3-ad63bdca53c5
	// 	datacontenttype: application/json
	// Data (binary),
	// 	{
	// 		"greetingWord": "hello",
	// 		"greetingTarget": "Chen",
	// 		"greetingTimes": 3
	// 	}

	// 2020/05/12 16:57:40 Hello (v1alpha1) received: greeting_word:"hello"  greeting_target:"Chen"
	// 2020/05/12 16:57:40 Hello (v1beta1) received: greeting_word:"hello"  greeting_target:"Chen"  greeting_times:3
}
