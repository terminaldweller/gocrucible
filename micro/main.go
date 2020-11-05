package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"strings"
	"time"
)

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

var ErrEmpty = errors.New("Empty string")

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"s"`
}

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}
}

func main() {
	// logger := log.NewLogfmtLogger(os.Stderr)
	// var svc StringService
	// svc = stringService{}
	// svc = loggingMiddleware{logger, svc}

	// uppercaseHandler := httptransport.NewServer(makeUppercaseEndpoint(svc), decodeUppercaseRequest, encodeResponse)
	// countHandler := httptransport.NewServer(makeCountEndpoint(svc), decodeCountRequest, encodeResponse)

	// http.Handle("/uppercase", uppercaseHandler)
	// http.Handle("/count", countHandler)

	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)

	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "total duration of requests in microseconds",
	}, fieldKeys)

	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "the result of each count method.",
	}, []string{})

	var svc StringService
	svc = stringService{}
	svc = loggingMiddleware{logger, svc}
	svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	uppercaseHandler := httptransport.NewServer(makeUppercaseEndpoint(svc), decodeUppercaseRequest, encodeResponse)
	countHandler := httptransport.NewServer(makeCountEndpoint(svc), decodeCountRequest, encodeResponse)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", "8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// var uppercase endpoint.Endpoint
// uppercase = makeUppercaseEndpoint()
// uppercase = loggingMiddleware(log.With(logger, "method", "uppercase"))(uppercase)

// var count endpoint.Endpoint
// count = makeCountEndpoint(svc)
// count = loggingMiddleware(log.With(logger, "method", "count"))(count)

// uppercaseHandler := httptransport.NewServer(uppercase, )
// countHandler := httptransport.NewServer(count, )

type loggingMiddleware struct {
	logger log.Logger
	next   StringService
}

func (mw loggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "uppercase", "input", s, "output", output, "err", err, "took", time.Since(begin))
	}(time.Now())

	output, err = mw.next.Uppercase(s)
	return
}

func (mw loggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "count", "inputl", s, "n", n, "took", time.Since(begin))
		return
	}(time.Now())

	n = mw.next.Count(s)
	return
}

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           StringService
}

func (mw instrumentingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "uppercase", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Uppercase(s)
	return
}

func (mw instrumentingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		lvs := []string{"method", "count", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	n = mw.next.Count(s)
	return
}

// type proxymw struct {
// 	next      StringService
// 	uppercase endpoint.Endpoint
// }

// func (mw proxymw) Uppercase(s string) (string, error) {
// 	response, err := mw.uppercase(uppercaseRequest{S: s})
// 	if err != nil {
// 		return "", err
// 	}
// 	resp := response.(uppercaseResponse)
// 	if resp.Err != "" {
// 		return resp.V, errors.New(resp.Err)
// 	}
// 	return resp.V, nil
// }

// func proxyingMiddleware(proxyURL string) ServiceMiddleware {
// 	return func(next StringService) StringService {
// 		return proxymw{next, makeUppercaseProxy(proxyURL)}
// 	}
// }

// func makeUppercaseProxy(proxyURL string) endpoint.Endpoint {
// 	return httptransport.NewClient(
// 		"GET",
// 		mustParseURL(proxyURL),
// 		encodeUppercaseRequest,
// 		decodeUppercaseResponse,
// 	).Endpoint()
// }
