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

package grpcserver

import (
	"context"
	"github.com/google/uuid"
	"tiktaktoe/models"
	pb "tiktaktoe/pkg"
)

type marshallerJoinToLobbyService interface {
	marshallJoinToLobbyResponse(dataModel *models.BattleField) *pb.JoinToLobbyResponse
	marshallJoinToLobbyError(err error) error
}

type marshallerNewGameService interface {
	marshallNewGameResponse(dataModel *models.BattleField) *pb.StartMatchResponse
	marshallNewGameError(err error) error
}

type gameEngineService interface {
	StartNewGame(ctx context.Context,
		playerOneUUID,
		playerTwoUUID uuid.UUID,
		fieldSize uint,
	) (*models.BattleField, error)
	StopGame(ctx context.Context,
		matchUUID uuid.UUID,
	) error
}

type joinToLobbyHandlerService interface {
	Handle(ctx context.Context, req *pb.JoinToLobbyRequest) (*pb.JoinToLobbyResponse, error)
}

type newGameHandlerService interface {
	Handle(ctx context.Context, req *pb.StartMatchRequest) (*pb.StartMatchResponse, error)
}

type stopGameHandlerService interface {
	Handle(ctx context.Context, request *pb.StopMatchRequest) (*pb.StopMatchResponse, error)
}

type playerMoveHandlerService interface {
	Handle(ctx context.Context, request *pb.PlayerMoveRequest) (*pb.PlayerMoveResponse, error)
}

type getMatchStatusHandlerService interface {
	Handle(ctx context.Context, request *pb.MatchStatusRequest) (*pb.MatchStatusResponse, error)
}

type getMatchListHandlerService interface {
	Handle(ctx context.Context, request *pb.MathListRequest) (*pb.MathListResponse, error)
}
