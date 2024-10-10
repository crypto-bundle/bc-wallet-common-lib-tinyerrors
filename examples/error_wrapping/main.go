/*
 *
 *
 * MIT NON-AI License
 *
 * Copyright (c) 2024-2024 Aleksei Kotelnikov(gudron2s@gmail.com)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of the software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions.
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * In addition, the following restrictions apply:
 *
 * 1. The Software and any modifications made to it may not be used for the purpose of training or improving machine learning algorithms,
 * including but not limited to artificial intelligence, natural language processing, or data mining. This condition applies to any derivatives,
 * modifications, or updates based on the Software code. Any usage of the Software in an AI-training dataset is considered a breach of this License.
 *
 * 2. The Software may not be included in any dataset used for training or improving machine learning algorithms,
 * including but not limited to artificial intelligence, natural language processing, or data mining.
 *
 * 3. Any person or organization found to be in violation of these restrictions will be subject to legal action and may be held liable
 * for any damages resulting from such use.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
 * DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
 * OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 *
 */

package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"bc-wallet-common-lib-tinyerrors/pkg/tinyerrors"
)

type ErrorCode int

const (
	RandLimit = 6000

	CodeUnableToMarshalData ErrorCode = iota + 5000
	CodeUnableToProcessRequestUnluckyNumber
)

var (
	CodeUnableToMarshalDataName                 = "unable to marshal response data"
	CodeUnableToProcessRequestUnluckyNumberName = "just unlucky request number value. " +
		"Please try again later"
)

var (
	ErrUnableProcessRequest = errors.New("unable to process request")
)

func (c ErrorCode) String() string {
	switch c {
	case CodeUnableToMarshalData:
		return CodeUnableToMarshalDataName
	case CodeUnableToProcessRequestUnluckyNumber:
		return CodeUnableToProcessRequestUnluckyNumberName
	default:
		return "<nil>"
	}
}

type responseModel struct {
	ActionName string `json:"action_name"`
	Message    string `json:"message"`
	ActionID   uint   `json:"action_id"`
	Weight     uint   `json:"weight"`
}

type responseModelError struct {
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
	ActionID     uint   `json:"action_id"`
}

type handler struct {
	logger  *log.Logger
	counter atomic.Uint64
}

func (h *handler) ServeHTTP(respWriter http.ResponseWriter, _ *http.Request) {
	newCounterValue := h.counter.Add(1)
	if newCounterValue%2 == 0 {
		respData := responseModel{
			ActionName: "serve_http",
			Message: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. " +
				"Donec sed nunc at mauris efficitur iaculis sit amet ac dui. " +
				"Sed lacinia suscipit risus sit amet sodales.",
			ActionID: uint(newCounterValue),
			Weight:   rand.UintN(RandLimit),
		}

		rawData, err := json.Marshal(respData)
		if err != nil {
			h.serveError(respWriter, tinyerrors.ErrorOnly(ErrUnableProcessRequest,
				CodeUnableToMarshalData.String()))
		}

		respWriter.WriteHeader(http.StatusOK)
		_, err = respWriter.Write(rawData)
		if err != nil {
			h.logger.Println("unable to write response")

			return
		}

		return
	}

	h.serveError(respWriter, tinyerrors.ErrorOnly(ErrUnableProcessRequest,
		CodeUnableToProcessRequestUnluckyNumber.String()))
}

func (h *handler) serveError(respWriter http.ResponseWriter, err error) {
	respData := &responseModelError{
		Message:      "unable to process request",
		ActionID:     uint(h.counter.Load()),
		ErrorDetails: err.Error(),
	}

	rawData, err := json.Marshal(respData)
	if err != nil {
		respWriter.WriteHeader(http.StatusInternalServerError)
		respWriter.Header().Set("Content-Type", "text/plain")

		_, err = respWriter.Write([]byte("internal error"))
		if err != nil {
			h.logger.Println("unable to write response")

			return
		}
	}

	respWriter.WriteHeader(http.StatusOK)
	respWriter.Header().Set("Content-Type", "application/json")

	_, err = respWriter.Write(rawData)
	if err != nil {
		h.logger.Println("unable to write response")

		return
	}
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	logger := log.Default()

	mux := http.NewServeMux()
	mux.Handle("/handle", &handler{
		logger:  logger,
		counter: atomic.Uint64{},
	})

	//nolint:exhaustruct // it's ok here. we don't need to fully fill up http.Server struct
	server := &http.Server{
		Addr:         "localhost:8083",
		Handler:      mux,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 3,
		ErrorLog:     logger,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Println("unable to start http server", err)

			return
		}
	}()

	go func() {
		<-ctx.Done()

		err := server.Shutdown(ctx)
		if err != nil {
			logger.Println("unable to shutdown http server", err)

			return
		}

		logger.Println("http server successfully shutdown")
	}()

	logger.Println("application started successfully")

	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	cancelFunc()

	time.Sleep(time.Second * 5)

	logger.Println("application closed")
}
