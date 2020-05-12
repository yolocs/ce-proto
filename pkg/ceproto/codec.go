package ceproto

import (
	"fmt"
	"strings"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var jsonDecodeOptions = protojson.UnmarshalOptions{
	DiscardUnknown: true,
	AllowPartial:   true,
}

var protoDecodeOptions = proto.UnmarshalOptions{
	AllowPartial:   true,
	DiscardUnknown: true,
}

func Marshal(msg proto.Message, event *event.Event) error {
	if err := setData(msg, event); err != nil {
		return err
	}
	return setEventAttributes(msg, event)
}

func Unmarshal(event *event.Event, msg proto.Message) error {
	switch event.DataContentType() {
	case "application/json":
		return jsonDecodeOptions.Unmarshal(event.Data(), msg)
	case "application/octet-stream":
		return protoDecodeOptions.Unmarshal(event.Data(), msg)
	}
	return fmt.Errorf("content type %s not supported", event.DataContentType())
}

// func Play(msg proto.Message) {
// 	mref := msg.ProtoReflect()
// 	mref.Descriptor().Options().ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
// 		log.Printf("Message option name: %q\n", fd.Name())
// 		log.Printf("Message option value: %q\n", v.String())
// 		return true
// 	})
// 	mref.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
// 		log.Printf("Field json name: %q\n", fd.JSONName())
// 		log.Printf("Field type: %v\n", fd.Kind())
// 		log.Printf("Field string value: %v\n", v.String())
// 		return true
// 	})
// }

func setData(msg proto.Message, event *event.Event) error {
	encoding := getEncoding(msg)
	switch encoding {
	case "JSON":
		b, err := protojson.Marshal(msg)
		if err != nil {
			return err
		}
		event.SetData("application/json", b)
	case "PROTO":
		b, err := proto.Marshal(msg)
		if err != nil {
			return err
		}
		event.SetData("application/octet-stream", b)
	default:
		return fmt.Errorf("encoding %s not supported", encoding)
	}

	return nil
}

func getEncoding(msg proto.Message) string {
	mref := msg.ProtoReflect()
	mopts := mref.Descriptor().Options().ProtoReflect()
	var encoding string
	mopts.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.Name() == "transport_encoding" {
			encoding = v.String()
		}
		return true
	})

	if encoding == "" {
		encoding = "JSON"
	}
	return encoding
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
