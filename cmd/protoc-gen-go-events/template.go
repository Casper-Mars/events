package main

import (
	"bytes"
	"strings"
	"text/template"
)

var publisherTemplate = `

// {{ .Name }}Publisher is an interface to publish {{ .Name }} event
type {{ .Name }}Publisher interface {
	SendEvent(ctx context.Context, in *{{ .Name }}) error
}

// New{{ .Name }}Publisher create a publisher for {{ .Name }} event
func New{{ .Name }}Publisher(metadata events.PublishMetadata, sender events.Sender) {{ .Name }}Publisher {
	return &{{ .ImplName }}Publisher{
		sender:   sender,
		metadata: metadata,
	}
}

// {{ .Name }}Subscriber is an interface to subscribe to {{ .Name }} event
type {{ .Name }}Subscriber interface {
	Receive(ctx context.Context, in *{{ .Name }}) error
}

// Register{{ .Name }}Subscriber register a subscriber 
func Register{{ .Name }}Subscriber(server events.Subscriber, subReq events.SubRequest, srv {{ .Name }}Subscriber) error {
	handler := &{{ .ImplName }}Handler{
		srv: srv,
	}
	return server.Subscribe(subReq, handler)
}

type {{ .ImplName }}Publisher struct {
	sender   events.Sender
	metadata events.PublishMetadata
}

func (b *{{ .ImplName }}Publisher) SendEvent(ctx context.Context, in *{{ .Name }}) error {
	data, err := proto.Marshal(in)
	if err != nil {
		return err
	}
	return b.sender.Send(ctx, events.Message{
		Data:  data,
		Topic: b.metadata.Topic,
	})
}

type {{ .ImplName }}Handler struct {
	srv {{ .Name }}Subscriber
}

func (o *{{ .ImplName }}Handler) Handle(ctx context.Context, msg events.Message) error {
	event := &{{ .Name }}{}
	err := proto.Unmarshal(msg.Data, event)
	if err != nil {
		return err
	}
	return o.srv.Receive(ctx, event)
}

`

type pubSubTemplateData struct {
	Name     string
	ImplName string
}

func newPublisherTemplateData(name string) pubSubTemplateData {
	return pubSubTemplateData{
		Name:     name,
		ImplName: strings.ToLower(name[:1]) + name[1:],
	}
}

type PubSubTemplate struct {
	Name string
}

func (t PubSubTemplate) Exec() string {
	buf := new(bytes.Buffer)
	parse, err := template.New("PubSubTemplate").Parse(publisherTemplate)
	if err != nil {
		panic(err)
	}
	data := newPublisherTemplateData(t.Name)
	if err := parse.Execute(buf, data); err != nil {
		panic(err)
	}
	return buf.String()
}
