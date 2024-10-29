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
	"context"

	"tiktaktoe/models"
	"tiktaktoe/types"

	"github.com/crypto-bundle/bc-wallet-common-lib-tinyerrors/pkg/tinyerrors"

	"github.com/google/uuid"
)

type battleFieldWorker struct {
	matchUUID uuid.UUID

	roles                   matchRolesManager
	fields                  tikTakToeFieldService
	battleFieldStoreDataSvc matchDataStoreService

	nextPlayer int
}

func (w *battleFieldWorker) WhoIsNext(_ context.Context) uuid.UUID {
	return w.roles.GePlayerUUIDBySymbol(w.nextPlayer)
}

func (w *battleFieldWorker) SetMovement(ctx context.Context,
	playerUUID uuid.UUID,
	position [2]uint8,
) (*types.MatchResult, error) {
	playerSymbol := w.roles.GetSymbolByPlayerUUID(playerUUID)
	if playerSymbol != w.nextPlayer {
		return nil, tinyerrors.ErrWithCode(ErrSetMove, types.TinyErrNotYourMovementOrder)
	}

	x, y := position[0], position[1]

	winner, err := w.fields.SetMove(x, y, playerSymbol)
	if err != nil {
		return nil, tinyerrors.ErrorNoWrap(err)
	}

	err = w.battleFieldStoreDataSvc.AddMatchMovement(ctx, &models.Movement{
		PlayerUUID:      playerUUID,
		Position:        position,
		BattleFieldUUID: w.matchUUID,
	})
	if err != nil {
		return nil, tinyerrors.ErrorNoWrap(err)
	}

	movementsCount, err := w.battleFieldStoreDataSvc.GetMatchMovementsCount(ctx, w.matchUUID)
	if err != nil {
		return nil, tinyerrors.ErrorNoWrap(err)
	}

	if winner > 0 { // match is over. return match result
		return &types.MatchResult{
			MatchUUID:     w.matchUUID,
			WinnerUUID:    w.roles.GePlayerUUIDBySymbol(winner),
			WinnerSign:    winner,
			MovementCount: uint(movementsCount),
		}, nil
	}

	// match still in progress, set who's next
	if movementsCount%2 == 0 { // first move of "X". "X" equal to 1 symbol
		w.nextPlayer = 1

		return nil, nil
	}

	// second move of "O". "O" equal to 0 symbol
	w.nextPlayer = 1

	return nil, nil
}
