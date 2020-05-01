package ceproto

import (
	"log"
	"strings"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Marshal(msg proto.Message, event *event.Event) error {
	b, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	event.SetData("application/octet-stream", b)
	return setEventAttributes(msg, event)
}

func Unmarshal(event *event.Event, msg proto.Message) error {
	if err := proto.Unmarshal(event.Data(), msg); err != nil {
		return err
	}
	return nil
}

func Play(msg proto.Message) {
	mref := msg.ProtoReflect()
	mref.Descriptor().Options().ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		log.Printf("Message option name: %q\n", fd.Name())
		log.Printf("Message option value: %q\n", v.String())
		return true
	})
	mref.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		log.Printf("Field json name: %q\n", fd.JSONName())
		log.Printf("Field type: %v\n", fd.Kind())
		log.Printf("Field string value: %v\n", v.String())
		return true
	})
}

func setEventAttributes(msg proto.Message, event *event.Event) error {
	event.SetID(uuid.New().String())

	mref := msg.ProtoReflect()
	mopts := mref.Descriptor().Options().ProtoReflect()
	mopts.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		switch fd.Name() {
		case "event_subject":
			event.SetSubject(v.String())
		case "event_source":
			event.SetSource(v.String())
		case "event_type":
			event.SetType(v.String())
		}
		return true
	})
	// Fill default values.
	fullName := mref.Descriptor().FullName()
	if event.Source() == "" {
		event.SetSource(string(fullName.Parent()))
	}
	if event.Type() == "" {
		event.SetType(string(fullName))
	}
	if event.Subject() == "" {
		event.SetSubject(string(fullName))
	}

	mref.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		fopts := fd.Options().ProtoReflect()
		fopts.Range(func(ofd protoreflect.FieldDescriptor, ov protoreflect.Value) bool {
			if ofd.Name() == "ce_extension" {
				if ov.Bool() {
					// We should only allow certain types to be extensions.
					event.SetExtension(strings.ToLower(fd.JSONName()), v.String())
					return false
				}
			}
			return true
		})
		return true
	})

	return nil
}
