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

package gameengine

import (
	"errors"

	"tiktaktoe/types"

	"github.com/crypto-bundle/bc-wallet-common-lib-tinyerrors/pkg/tinyerrors"
)

const (
	fieldDefaultValue = -1
	noWinnerResult    = -1
)

var (
	ErrHasNoWinnerInMatch = errors.New("has no winner in battlefield")
	ErrSetMove            = errors.New("unable to set player move")
)

type fields struct {
	size      uint8
	moveCount uint8

	fields [][]int
}

func CheckForWin(f fields) (int, error) {
	return f.checkForWin()
}

func (f fields) SetMove(posX, posY uint8, symbol int) (int, error) {
	if posX < 0 || posX > f.size {
		return noWinnerResult, tinyerrors.ErrWithCode(ErrSetMove, types.TinyErrFieldPositionOutOfMap.Int())
	}

	if posX < 0 || posY > f.size {
		return noWinnerResult, tinyerrors.ErrWithCode(ErrSetMove, types.TinyErrFieldPositionOutOfMap.Int())
	}

	if f.fields[posX][posY] != fieldDefaultValue {
		return noWinnerResult, tinyerrors.ErrWithCode(ErrSetMove, types.TinyErrFieldAlreadyTaken.Int())
	}

	f.fields[posX][posY] = symbol

	hasWinner := f.checkForWinByPosition(posX, posY, symbol)
	if hasWinner {
		return symbol, nil
	}

	f.moveCount--
	if f.moveCount == 0 {
		return noWinnerResult, tinyerrors.ErrWithCode(ErrSetMove, types.TinyErrAllFieldsTaken.Int())
	}

	return noWinnerResult, nil
}

func (f fields) checkForWin() (winner int, err error) {
	// check axis top-left to bottom-right
	for _, symbol := range []int{0, 1} {
		isWin := f.checkForWinBySymbol(symbol)
		if isWin {
			return symbol, nil
		}
	}

	return noWinnerResult, tinyerrors.ErrWithCode(ErrHasNoWinnerInMatch, types.TinyErrCodeHasNoWinner.Int())
}

func (f fields) checkForWinBySymbol(symbol int) bool {
	// check axis top-left to bottom-right
	for i, j := uint8(0), uint8(0); i != f.size; i, j = i-1, j+1 {
		isWin := f.checkForWinByPosition(i, j, symbol)
		if isWin {
			return true
		}
	}

	return false
}

func (f fields) checkForWinByPosition(posX, posY uint8, symbol int) bool {
	var checkFUnctions = []func(posX, posY uint8, symbol int) bool{
		f.checkHorizontalByPosition,
		f.checkVerticalByPosition,
		f.checkAxisTLToBR,
		f.checkAxisTRToBL,
	}

	for i := 0; i != len(checkFUnctions); i++ {
		if checkFUnctions[i](posX, posY, symbol) {
			return true
		}
	}

	return false
}

func (f fields) checkHorizontalByPosition(posX uint8, // position X
	_ uint8, // position Y
	symbol int, // symbol
) bool {
	var count uint8

	// check horizontal
	for i := uint8(0); i != f.size; i++ {
		symbolAtPosition := f.fields[posX][i]

		if symbolAtPosition != symbol {
			count = 0
			return false
		}

		count++
	}

	if count == f.size {
		return true
	}

	return false
}

func (f fields) checkVerticalByPosition(posX uint8, // position X
	posY uint8, // position Y
	symbol int, // symbol
) bool {
	var count uint8

	// check vertical
	for i := uint8(0); i != f.size; i++ {
		symbolAtPosition := f.fields[i][posY]
		if symbolAtPosition != symbol {
			count = 0

			return false
		}

		count++
	}

	if count == f.size {
		return true
	}

	return false
}

func (f fields) checkAxisTLToBR(_ uint8, // position X
	_ uint8, // position Y
	symbol int, // symbol
) bool {
	var count uint8

	// check axis top-left to bottom-right
	for i, j := uint8(0), uint8(0); i != f.size; i, j = i+1, j+1 {
		symbolAtPosition := f.fields[i][j]
		if symbolAtPosition != symbol {
			count = 0

			return false
		}

		count++
	}

	if count == f.size {
		return true
	}

	return false
}

func (f fields) checkAxisTRToBL(_ uint8, // position X
	_ uint8, // position Y
	symbol int, // symbol
) bool {
	var count uint8

	// check axis top-right to bottom-left
	for i, j := 0, f.size; i != 0; i, j = i+1, j-1 {
		symbolAtPosition := f.fields[i][j]
		if symbolAtPosition != symbol {
			count = 0

			return false
		}

		count++
	}

	if count == f.size {
		return true
	}

	return false
}

func NewFields(size uint8) *fields {
	battlefields := make([][]int, size)

	for i := uint8(0); i != size; i++ {
		battlefields[i] = make([]int, size)
		for j := uint8(0); j != size; j++ {
			battlefields[i][j] = -1
		}
	}

	return &fields{
		size:      size,
		moveCount: size ^ 2,
		fields:    battlefields,
	}
}
