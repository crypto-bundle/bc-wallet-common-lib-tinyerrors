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
	"github.com/crypto-bundle/bc-wallet-common-lib-tinyerrors/pkg/tinyerrors"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type pingWorker struct {
	ticker *time.Ticker
	client *http.Client
	logger *log.Logger
}

func (w *pingWorker) Run(ctx context.Context) {
	w.ticker = time.NewTicker(time.Second * 3)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-w.ticker.C:
				err := w.onTick(ctx)
				if err != nil {
					w.logger.Println("ping flow failed", err)
				}

				w.logger.Println("ping-pong success")
			}
		}
	}()
}

func (w *pingWorker) onTick(_ context.Context) error {
	resp, err := w.client.Get("http://localhost:8082/ping")
	if err != nil {
		return tinyerrors.ErrorNoWrap(err)
	}
	defer func() {
		closeErr := tinyerrors.ErrorNoWrap(resp.Body.Close())
		if closeErr != nil {
			w.logger.Printf("unable to close resp body %e", closeErr)
		}
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return tinyerrors.NewErrorf("unable to read response: %s, status-code: %s",
			err.Error(), resp.Status)
	}

	return tinyerrors.NewErrorf("request ended with wrong status-code: %s", resp.Status)
}

func NewPingWorker() *pingWorker {
	return &pingWorker{
		ticker: nil,
		client: &http.Client{
			Transport:     http.DefaultTransport,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       time.Second,
		},
		logger: log.New(os.Stdout, "ping_worker ", log.LstdFlags),
	}
}
