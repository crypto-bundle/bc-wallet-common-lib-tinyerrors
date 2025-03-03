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
	"fmt"
	"github.com/crypto-bundle/bc-wallet-common-lib-tinyerrors/pkg/tinyerrors"
)

type CoinSide bool

const (
	CoinSideHead = true
	CoinSideTail = false

	CoinSideHeadText = "head"
	CoinSideTailText = "tail"
)

func (s CoinSide) String() string {
	if s {
		return CoinSideHeadText
	}

	return CoinSideTailText
}

type responseModel struct {
	ActionName string   `json:"action_name"`
	Message    string   `json:"message"`
	ActionID   uint     `json:"action_id"`
	CoinSide   CoinSide `json:"coin_side"`
}

type responseModelError struct {
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
	ActionID     uint   `json:"action_id"`
}

func newResponseModelPresenter(side CoinSide,
	reqUrl string,
	actionID uint,
) ([]byte, error) {
	respData := &responseModel{
		ActionName: reqUrl,
		Message:    fmt.Sprintf("you have got a %s coin side", side),
		ActionID:   actionID,
		CoinSide:   side,
	}

	rawData, err := json.Marshal(respData)
	if err != nil {
		// example of pseudo-wrap error here. We don't need wrap error because trying to prevent re-wrap error in
		// serveCoinFlip function. In case if we wrap error here via tinyerrors.ErrorOnly, caller function will
		// wrap error again
		return nil, tinyerrors.ErrorNoWrap(err)
	}

	return rawData, nil
}
