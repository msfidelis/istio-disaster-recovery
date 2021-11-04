package httpclient

import (
	"orders-api/pkg/logger"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Request(method string, host string, path string, headers map[string][]string, body string) (*http.Response, string) {

	log := logger.Instance()
	log.Info().
		Str("action", "request").
		Str("method", strings.ToUpper(method)).
		Str("host", host).
		Str("path", path).
		Str("body", body).
		Msg("Execution an HTTP request as proxy")

	reqURL, _ := url.Parse(fmt.Sprintf("%s%s", host, path))
	reqBody := ioutil.NopCloser(strings.NewReader(body))

	req := &http.Request{
		Method: strings.ToUpper(method),
		URL:    reqURL,
		Header: headers,
		Body:   reqBody,
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Error().
			Str("action", "request").
			Str("method", strings.ToUpper(method)).
			Str("host", host).
			Str("path", path).
			Str("body", body).
			Str("error", err.Error()).
			Msg("Error during request execution")
	}

	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	log.Info().
		Str("action", "proxy").
		Str("method", strings.ToUpper(method)).
		Str("host", host).
		Str("path", path).
		Str("body", string(data)).
		Int("status_code", res.StatusCode).
		Msg("Success on request execution")

	return res, string(data)

}