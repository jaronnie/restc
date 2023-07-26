package restc

import (
	"net/http"
	"time"
)

type Interface interface {
	Verb(verb string) *Request
	Post() *Request
	Get() *Request

	GetHeader() http.Header
}

type Opt func(client *RESTClient) error

type RESTClient struct {
	protocol string
	addr     string
	port     string

	gatewayPrefix string

	retryTimes int
	retryDelay time.Duration

	headers http.Header

	// Set specific behavior of the client.  If not set http.DefaultClient will be used.
	client *http.Client
}

func (r *RESTClient) Verb(verb string) *Request {
	return NewRequest(r).Verb(verb)
}

func (r *RESTClient) Post() *Request {
	return r.Verb("POST")
}

func (r *RESTClient) Get() *Request {
	return r.Verb("GET")
}

func (r *RESTClient) GetHeader() http.Header {
	return r.headers
}

func RESTClientFor(config *RESTClient) (*RESTClient, error) {
	rest := &RESTClient{
		protocol:      config.protocol,
		addr:          config.addr,
		port:          config.port,
		gatewayPrefix: config.gatewayPrefix,
		retryTimes:    config.retryTimes,
		retryDelay:    config.retryDelay,
		headers:       config.headers,
		client:        config.client,
	}
	return rest, nil
}
