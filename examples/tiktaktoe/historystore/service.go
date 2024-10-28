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

package historystore

import (
	"context"
	"errors"
	"github.com/crypto-bundle/bc-wallet-common-lib-tinyerrors/pkg/tinyerrors"
	"sync"
	"tiktaktoe/types"

	"tiktaktoe/models"

	"github.com/google/uuid"
)

var (
	ErrMatchInfoAlreadyExist = errors.New("match already exists")
)

type store struct {
	mu sync.Mutex

	matchMap     map[uuid.UUID]models.BattleField
	movementsMap map[uuid.UUID][]models.Movement
}

func (s *store) GetMatchInfo(_ context.Context, matchUUID uuid.UUID) *models.BattleField {
	info, isExists := s.matchMap[matchUUID]
	if !isExists {
		return nil
	}

	return info.Clone()
}

func (s *store) GetAllMatches(_ context.Context) []models.BattleField {
	s.mu.Lock()
	defer s.mu.Unlock()

	list := make([]models.BattleField, 0, len(s.matchMap))
	counter := 0

	for _, matchInfo := range s.matchMap {
		list[counter] = matchInfo
		counter++
	}

	return list
}

func (s *store) AddMatchInfo(_ context.Context, info models.BattleField) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, isExists := s.matchMap[info.UUID]
	if isExists {
		return tinyerrors.ErrorWithCode(ErrMatchInfoAlreadyExist,
			types.TinyErrCodeMatchAlreadyRegistered.Int())
	}

	s.matchMap[info.UUID] = info

	return nil
}

func (s *store) AddMatchMovement(_ context.Context, movement models.Movement) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	matchInfo, isExists := s.matchMap[movement.BattleFieldUUID]
	if !isExists {
		return tinyerrors.ErrorWithCode(ErrMatchInfoAlreadyExist,
			types.TinyErrCodeMatchNotRegistered.Int())
	}

	movementsList, isExists := s.movementsMap[movement.BattleFieldUUID]
	if !isExists {
		s.movementsMap[movement.BattleFieldUUID] = make([]models.Movement, 0, matchInfo.Size^2)
		s.movementsMap[movement.BattleFieldUUID][0] = movement

		return nil
	}

	movementsList = append(movementsList, movement)

	return nil
}

func (s *store) GetAllMatchMovement(_ context.Context, matchUUID uuid.UUID) ([]models.Movement, error) {
	_, isExists := s.matchMap[matchUUID]
	if !isExists {
		return nil, tinyerrors.ErrorWithCode(ErrMatchInfoAlreadyExist,
			types.TinyErrCodeMatchNotRegistered.Int())
	}

	movementsList := s.movementsMap[matchUUID]

	return append(movementsList[:0:0], movementsList...), nil
}
