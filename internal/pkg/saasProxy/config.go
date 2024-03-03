package saasProxy

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type Configuration struct {
	Webhooks    []Webhook `json:"webhooks" yaml:"webhooks" toml:"webhooks"`
	Port        int       `json:"port" yaml:"port" toml:"port"`
	Destination string    `json:"destination" yaml:"destination" toml:"destination"`
}

type Webhook struct {
	IncomingSlug     string            `json:"incoming_slug" yaml:"incoming_slug" toml:"incoming_slug"`
	HttpResponseCode int               `json:"http_response_code" yaml:"http_response_code" toml:"http_response_code"`
	ResponseBody     string            `json:"response_body" yaml:"response_body" toml:"response_body"`
	RequestVerb      string            `json:"request_verb" yaml:"request_verb" toml:"request_verb"`
	Headers          map[string]string `json:"headers" yaml:"headers" toml:"headers"`
}

func (c *Configuration) ToServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	for i := 0; i < len(c.Webhooks); i++ {
		c.Webhooks[i].ToHttpHandlerMux(mux)
	}
	return mux
}

func (w *Webhook) ToHttpHandlerMux(mux *http.ServeMux) *http.ServeMux {
	incomingPath := fmt.Sprintf("%s %s", w.RequestVerb, w.IncomingSlug)
	mux.HandleFunc(fmt.Sprintf(incomingPath), w.GetResponseBody())
	return mux
}

func (w *Webhook) GetResponseBody() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.WithFields(log.Fields{
			"request.URL.Path":   &request.URL.Path,
			"w.HttpResponseCode": w.HttpResponseCode,
			"request.Method":     request.Method,
		}).Info("saasProxy incoming request received.")
		w.addHeadersToHandler(writer)
		if w.ResponseBody == "pass-through" {
			var b []byte
			var err error
			b, err, done := readResponseBodyBytes(request, b, err)
			if done {
				return
			}

			_, err = writer.Write(b)
			if err != nil {
				handleGetResponseBodyWriterError(err, w)
				return
			}
		} else {
			_, err := writer.Write([]byte(w.ResponseBody))
			if err != nil {
				handleGetResponseBodyWriterError(err, w)
				return
			}
		}
	}
}

func readResponseBodyBytes(request *http.Request, b []byte, err error) ([]byte, error, bool) {
	if request.Body != nil {
		b, err = io.ReadAll(request.Body)
		if err != nil {
			log.Error(fmt.Sprintf("Body reading error: %v", err))
			return nil, nil, true
		}
		defer request.Body.Close()
	}
	return b, err, false
}

func (w *Webhook) addHeadersToHandler(writer http.ResponseWriter) {
	for k, v := range w.Headers {
		log.WithFields(log.Fields{
			"k": k,
			"v": v,
		}).Info("Adding header key and value...")
		writer.Header().Set(k, v)
	}
	log.WithFields(log.Fields{
		"writer.Header()": writer.Header(),
	}).Info("HTTP response writer headers added")
}

func handleGetResponseBodyWriterError(err error, webhook *Webhook) {
	log.WithFields(log.Fields{
		"webhook.IncomingSlug":     webhook.IncomingSlug,
		"webhook.HttpResponseCode": webhook.HttpResponseCode,
		"webhook.ResponseBody":     webhook.ResponseBody,
		"webhook.RequestVerb":      webhook.RequestVerb,
	}).Error("Unable to write webhook.ResponseBody as []byte")
	return
}
