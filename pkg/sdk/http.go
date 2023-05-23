package library

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/sirupsen/logrus"
)

func RawGetRequest(logger *logrus.Entry, uri string, timeOut int64, queryParams map[string]string) (response []byte, e error) {

	request, e := http.NewRequest("GET", uri, nil)
	q := request.URL.Query()

	for k, v := range queryParams {
		q.Add(k, v)
	}

	request.URL.RawQuery = q.Encode()

	if e != nil {
		logger.WithField("error", e).Error("Catch error unable to marshal request to json format")
		return
	}

	var cfg *http.Transport
	cfg = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).DialContext,
		MaxIdleConns:        10,
		IdleConnTimeout:     10 * time.Second,
		TLSHandshakeTimeout: 5 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Timeout:   time.Millisecond * time.Duration(timeOut),
		Transport: cfg,
	}

	resp, e := client.Do(request)
	if e != nil {
		if isTimeout(e) {
			logger.WithField("error", e).Error("Catch timeout detected")
		} else {
			logger.WithField("error", e).Error("Catch error unable to send GET")
		}
		return
	} else {
		defer resp.Body.Close()

		if resp == nil {
			e = errors.New("emtpy response")
			logger.WithField("error", e).Error("Catch error client do error")
			return nil, e
		}

		responseDump, e := httputil.DumpResponse(resp, true)
		if e != nil {
			return nil, e
		}
		logger.WithField("response", string(responseDump)).Info("Receiving response from " + request.URL.String())
	}

	response, e = ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, e
	}

	return
}

func RawGetRequestV2(logger *logrus.Entry, uri string, timeOut int64, id string) (response []byte, e error) {

	request, e := http.NewRequest("GET", uri+id, nil)
	q := request.URL.Query()

	request.URL.RawQuery = q.Encode()

	if e != nil {
		logger.WithField("error", e).Error("Catch error unable to marshal request to json format")
		return
	}

	var cfg *http.Transport
	cfg = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).DialContext,
		MaxIdleConns:        10,
		IdleConnTimeout:     10 * time.Second,
		TLSHandshakeTimeout: 5 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Timeout:   time.Millisecond * time.Duration(timeOut),
		Transport: cfg,
	}

	resp, e := client.Do(request)
	if e != nil {
		if isTimeout(e) {
			logger.WithField("error", e).Error("Catch timeout detected")
		} else {
			logger.WithField("error", e).Error("Catch error unable to send GET")
		}
		return
	} else {
		defer resp.Body.Close()

		if resp == nil {
			e = errors.New("emtpy response")
			logger.WithField("error", e).Error("Catch error client do error")
			return nil, e
		}

		responseDump, e := httputil.DumpResponse(resp, true)
		if e != nil {
			return nil, e
		}
		logger.WithField("response", string(responseDump)).Info("Receiving response from " + request.URL.String())
	}

	response, e = ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, e
	}

	return
}

func isTimeout(err error) bool {
	if err, ok := err.(net.Error); ok && err.Timeout() {
		return true
	}
	return false
}
