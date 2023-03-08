package httpie

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

type Server struct {
	instance *httptest.Server
	URL      string
}

func (s *Server) Stop() {
	s.instance.Close()
}

type JSONKV map[string]interface{}

type Response struct {
	url     string
	handler func(w http.ResponseWriter, r *http.Request)
	headers map[string]string
}

func (r *Response) AddHeader(key, value string) *Response {
	if r.headers == nil {
		r.headers = map[string]string{}
	}
	r.headers[key] = value
	return r
}

func (r *Response) AppendHeaders(w http.ResponseWriter) {
	for k, v := range r.headers {
		w.Header().Add(k, v)
	}
}

func WithJSON(url string, data interface{}) *Response {
	res := &Response{
		url: url,
	}
	res.handler = func(w http.ResponseWriter, r *http.Request) {
		res.AppendHeaders(w)
		j, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		_, _ = w.Write(j)
	}
	return res
}

func WithBytes(url, mime string, data []byte) *Response {
	res := &Response{
		url: url,
	}
	res.handler = func(w http.ResponseWriter, r *http.Request) {
		res.AppendHeaders(w)
		w.Header().Add("Content-Type", mime)
		_, _ = w.Write(data)
	}
	return res
}

func WithCustom(url string, fn func(w http.ResponseWriter, r *http.Request)) *Response {
	res := &Response{
		url: url,
	}
	res.handler = func(w http.ResponseWriter, r *http.Request) {
		res.AppendHeaders(w)
		fn(w, r)
	}
	return res
}

func (c *config) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if hand, ok := c.urls[r.URL.Path]; ok {
		hand.handler(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("Not found"))
}

type config struct {
	urls map[string]*Response
}

func New(responses ...*Response) *Server {
	c := &config{urls: map[string]*Response{}}
	for _, r := range responses {
		c.urls[r.url] = r
	}
	s := &Server{
		instance: httptest.NewUnstartedServer(c),
	}
	s.instance.Start()
	s.URL = s.instance.URL
	return s
}
