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
	"encoding/json"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/crypto-bundle/bc-wallet-common-lib-tinyerrors/pkg/tinyerrors"
)

type handler struct {
	logger          *log.Logger
	counter         atomic.Uint64
	coinFlipCounter atomic.Uint64
}

func (h *handler) ServeHTTP(respWriter http.ResponseWriter, req *http.Request) {
	newCounterValue := h.counter.Add(1)
	if newCounterValue%2 == 0 {
		h.serveCoinFlip(respWriter, req)

		return
	}

	// Example of error wrap. Error message will contain one additional detail.
	// Result of wrap will something like this:
	// unable to process request -> just unlucky request number value. Please try again later
	h.serveError(respWriter, tinyerrors.ErrorOnly(ErrUnableProcessRequest,
		CodeUnableToProcessRequestUnluckyNumber.String()))
}

func (h *handler) serveCoinFlip(respWriter http.ResponseWriter, req *http.Request) {
	var coinSide CoinSide = false

	newCounterValue := h.coinFlipCounter.Add(1)
	if newCounterValue%2 == 0 {
		coinSide = true
	}

	rawData, err := newResponseModelPresenter(coinSide, req.RequestURI, uint(h.counter.Load()))
	if err != nil {
		// Example of error wrap. Error message will contain two additional details.
		// Result of wrap will something like this:
		// <error_message> -> unable marshal coin flip result message, tail
		h.serveError(respWriter, tinyerrors.ErrorOnly(err,
			"unable marshal coin flip result message", coinSide.String()))

		return
	}

	respWriter.Header().Set("Content-Type", "application/json")
	respWriter.WriteHeader(http.StatusOK)
	_, err = respWriter.Write(rawData)
	if err != nil {
		h.logger.Println("unable to write response", err)

		return
	}
}

func (h *handler) serveError(respWriter http.ResponseWriter, err error) {
	respData := &responseModelError{
		Message:      "unable to process request",
		ActionID:     uint(h.counter.Load()),
		ErrorDetails: err.Error(),
	}

	rawData, err := json.Marshal(respData)
	if err != nil {
		respWriter.Header().Set("Content-Type", "text/plain")
		respWriter.WriteHeader(http.StatusInternalServerError)

		_, err = respWriter.Write([]byte("internal error"))
		if err != nil {
			h.logger.Println("unable to write response")

			return
		}
	}

	respWriter.Header().Set("Content-Type", "application/json")
	respWriter.WriteHeader(http.StatusOK)
	_, err = respWriter.Write(rawData)
	if err != nil {
		h.logger.Println("unable to write response")

		return
	}
}
