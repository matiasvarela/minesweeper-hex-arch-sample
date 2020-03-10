package gamesrv_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"

	"github.com/matiasvarela/errors"
	"github.com/matiasvarela/minesweeper-hex-arch-sample/internal/core/domain"
	"github.com/matiasvarela/minesweeper-hex-arch-sample/internal/core/service/gamesrv"
	"github.com/matiasvarela/minesweeper-hex-arch-sample/mocks/mockups"
	"github.com/matiasvarela/minesweeper-hex-arch-sample/pkg/apperrors"
)

type mocks struct {
	gameRepository *mockups.MockGamesRepository
	uidGen         *mockups.MockUIDGen
}

func TestGet(t *testing.T) {
	// · Mocks · //

	game := domain.NewGame("1001-1001-1001-1001", "my game", 10, 50)

	gameWithBombsHided := domain.NewGame("1001-1001-1001-1001", "my game", 10, 50)
	gameWithBombsHided.Board = game.Board.HideBombs()

	// · Tests · //
	type args struct {
		id string
	}

	type want struct {
		result domain.Game
		err    error
	}

	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mocks)
	}{
		{
			name: "Should get game successfully",
			args: args{id: "1001-1001-1001-1001"},
			want: want{result: gameWithBombsHided},
			mocks: func(m mocks) {
				m.gameRepository.EXPECT().Get("1001-1001-1001-1001").Return(game, nil)
			},
		},
		{
			name: "Should return error - game not found",
			args: args{id: "1001-1001-1001-1001"},
			want: want{err: errors.New(apperrors.NotFound, nil, "game not found")},
			mocks: func(m mocks) {
				m.gameRepository.EXPECT().Get("1001-1001-1001-1001").Return(domain.Game{}, errors.New(apperrors.NotFound, nil, ""))
			},
		},
		{
			name: "Should return error - get from repository fails",
			args: args{id: "1001-1001-1001-1001"},
			want: want{err: errors.New(apperrors.Internal, nil, "get game from repository has failed")},
			mocks: func(m mocks) {
				m.gameRepository.EXPECT().Get("1001-1001-1001-1001").Return(domain.Game{}, errors.New(apperrors.Internal, nil, "game not found"))
			},
		},
	}

	// · Runner · //

	for _, tt := range tests {
		tt := tt

		// Prepare
		m := mocks{
			gameRepository: mockups.NewMockGamesRepository(gomock.NewController(t)),
			uidGen:         mockups.NewMockUIDGen(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := gamesrv.New(m.gameRepository, m.uidGen)

		// Execute
		result, err := service.Get(tt.args.id)

		// Verify
		if tt.want.err != nil && err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)
	}
}
