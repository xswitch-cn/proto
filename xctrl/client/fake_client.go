package client

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/xswitch-cn/proto/xctrl/codec"
)

const defaultTimeout = 60 * time.Second
const defaultDialTimeout = 5 * time.Second

func newOptions(options ...Option) Options {
	opts := Options{
		Codecs: make(map[string]codec.NewCodec),
		CallOptions: CallOptions{
			Backoff:        DefaultBackoff,
			Retry:          DefaultRetry,
			Retries:        DefaultRetries,
			RequestTimeout: DefaultRequestTimeout,
			DialTimeout:    defaultDialTimeout,
		},
		PoolSize: DefaultPoolSize,
		PoolTTL:  DefaultPoolTTL,
	}

	for _, o := range options {
		o(&opts)
	}

	if opts.Context == nil {
		opts.Context = context.Background()
	}

	return opts
}

type fakeClient struct {
	opts Options
}

type fakeRequest struct {
	service     string
	method      string
	endpoint    string
	contentType string
	codec       codec.Codec
	body        interface{}
	opts        RequestOptions
}

func (r *fakeRequest) ContentType() string {
	return r.contentType
}

func (r *fakeRequest) Service() string {
	return r.service
}

func (r *fakeRequest) Method() string {
	return r.method
}

func (r *fakeRequest) Endpoint() string {
	return r.endpoint
}

func (r *fakeRequest) Body() interface{} {
	return r.body
}

func (r *fakeRequest) Codec() codec.Writer {
	return r.codec
}

func NewFakeClient(opt ...Option) Client {
	opts := newOptions(opt...)
	return &fakeClient{
		opts: opts,
	}
}

func (f *fakeClient) Init(opts ...Option) error {
	for _, o := range opts {
		o(&f.opts)
	}
	return nil
}

func (r *fakeClient) SetAService() error {
	return nil
}

func (r *fakeClient) Options() Options {
	return r.opts
}

func (r *fakeClient) Call(ctx context.Context, request Request, response interface{}, opts ...CallOption) error {
	log.Println("request method:", request.Method())
	switch request.Method() {
	case "XNode.NativeAPI":
		data := []byte(`{
			"code": 200,
			"message": "OK",
			"node_uuid": "test.test",
			"data": "Fake Data",
			"result": "result"
		}`)
		json.Unmarshal(data, response)
	case "getConferenceList":
		data := []byte(`{
			"conferences": [
				{
					"name": "test",
					"domain": "test"
				}
			]
		}`)
		json.Unmarshal(data, response)
	}
	return nil
}

func (r *fakeClient) Publish(ctx context.Context, msg Message, opts ...PublishOption) error {
	return nil
}

func (r *fakeClient) NewMessage(topic string, message interface{}, opts ...MessageOption) Message {
	return nil
}

func (r *fakeClient) NewRequest(service, method string, request interface{}, reqOpts ...RequestOption) Request {
	var opts RequestOptions

	for _, o := range reqOpts {
		o(&opts)
	}

	return &fakeRequest{
		service:     service,
		method:      method,
		endpoint:    "fake",
		body:        request,
		contentType: "application/json",
		opts:        opts,
	}
}

func (r *fakeClient) String() string {
	return "fake"
}
