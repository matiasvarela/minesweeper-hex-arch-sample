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

	game := easymockGame("1001", "mygame", 4, "", false, []pos{{1, 1}}, []pos{})
	gameWithBombsHided := easymockGame("1001", "mygame", 4, "", true, []pos{{1, 1}}, []pos{})

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
				m.gameRepository.EXPECT().Get("1001-1001-1001-1001").Return(domain.Game{}, errors.New(apperrors.Internal, nil, ""))
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

func TestCreate(t *testing.T) {
	// · Mocks · //

	gameWithBombsHided := easymockGame("1001", "mygame", 4, "", true, []pos{{0, 0}, {1, 1}}, []pos{})

	// · Tests · //

	type args struct {
		name  string
		size  uint
		bombs uint
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
			name: "Should create a new game successfully",
			args: args{name: "mygame", size: 4, bombs: 2},
			want: want{result: gameWithBombsHided},
			mocks: func(m mocks) {
				m.uidGen.EXPECT().New().Return("1001")
				m.gameRepository.EXPECT().Save(gomock.Any()).Return(nil)
			},
		},
		{
			name: "Should return an error - create game into repository fails",
			args: args{name: "mygame", size: 4, bombs: 2},
			want: want{err: errors.New(apperrors.Internal, nil, "create game into repository has failed")},
			mocks: func(m mocks) {
				m.uidGen.EXPECT().New().Return("1001")
				m.gameRepository.EXPECT().Save(gomock.Any()).Return(errors.New(apperrors.Internal, nil, ""))
			},
		},
		{
			name:  "Should return an error - invalid bombs number",
			args:  args{name: "mygame", size: 4, bombs: 40},
			want:  want{err: errors.New(apperrors.InvalidInput, nil, "the number of bombs is too high")},
			mocks: func(m mocks) {},
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
		result, err := service.Create(tt.args.name, tt.args.size, tt.args.bombs)

		// Verify
		if tt.want.err != nil && err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result.ID, result.ID)
		assert.Equal(t, tt.want.result.Name, result.Name)
		assert.Equal(t, tt.want.result.State, result.State)
		assert.Equal(t, tt.want.result.BoardSettings, result.BoardSettings)
		assert.Equal(t, len(tt.want.result.Board), len(result.Board))
	}
}

func TestReveal(t *testing.T) {
	// · Mocks · //

	// · Tests · //

	type args struct {
		id  string
		row uint
		col uint
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
			name: "Should reveal cell successfully - result in game not over",
			args: args{id: "1001", row: 2, col: 2},
			want: want{result: easymockGame("1001", "mygame", 4, "", true, []pos{{1, 1}}, []pos{{2, 2}})},
			mocks: func(m mocks) {
				game := easymockGame("1001", "mygame", 4, "", false, []pos{{1, 1}}, []pos{})
				gameToSave := easymockGame("1001", "mygame", 4, "", false, []pos{{1, 1}}, []pos{{2, 2}})

				m.gameRepository.EXPECT().Get("1001").Return(game, nil)
				m.gameRepository.EXPECT().Save(gameToSave).Return(nil)
			},
		},
		{
			name: "Should reveal cell successfully - result in game over - lost",
			args: args{id: "1001", row: 1, col: 1},
			want: want{result: easymockGame("1001", "mygame", 4, domain.GAME_STATE_LOST, true, []pos{{1, 1}}, []pos{})},
			mocks: func(m mocks) {
				game := easymockGame("1001", "mygame", 4, "", false, []pos{{1, 1}}, []pos{})
				gameToSave := easymockGame("1001", "mygame", 4, domain.GAME_STATE_LOST, false, []pos{{1, 1}}, []pos{})

				m.gameRepository.EXPECT().Get("1001").Return(game, nil)
				m.gameRepository.EXPECT().Save(gameToSave).Return(nil)
			},
		},
		{
			name: "Should reveal cell successfully - result in game over - won",
			args: args{id: "1001", row: 0, col: 0},
			want: want{result: easymockGame("1001", "mygame", 2, domain.GAME_STATE_WON, true, []pos{{1, 1}}, []pos{{0,0},{0,1},{1,0}})},
			mocks: func(m mocks) {
				game := easymockGame("1001", "mygame", 2, "", false, []pos{{1, 1}}, []pos{{0,1},{1,0}})
				gameToSave := easymockGame("1001", "mygame", 2, domain.GAME_STATE_WON, false, []pos{{1, 1}}, []pos{{0,0},{0,1},{1,0}})

				m.gameRepository.EXPECT().Get("1001").Return(game, nil)
				m.gameRepository.EXPECT().Save(gameToSave).Return(nil)
			},
		},
		{
			name: "Should return an error - game not found",
			args: args{id: "1001", row: 2, col: 2},
			want: want{err: errors.New(apperrors.NotFound, nil, "game not found")},
			mocks: func(m mocks) {
				m.gameRepository.EXPECT().Get("1001").Return(domain.Game{}, errors.New(apperrors.NotFound, nil, ""))
			},
		},
		{
			name: "Should return an error - fail to get the game from repository",
			args: args{id: "1001", row: 2, col: 2},
			want: want{err: errors.New(apperrors.Internal, nil, "get game from repository has failed")},
			mocks: func(m mocks) {
				m.gameRepository.EXPECT().Get("1001").Return(domain.Game{}, errors.New(apperrors.Internal, nil, ""))
			},
		},
		{
			name: "Should return an error - invalid position",
			args: args{id: "1001", row: 20, col: 2},
			want: want{err: errors.New(apperrors.InvalidInput, nil, "invalid position")},
			mocks: func(m mocks) {
				game := easymockGame("1001", "mygame", 4, "", false, []pos{{1, 1}}, []pos{})

				m.gameRepository.EXPECT().Get("1001").Return(game, nil)
			},
		},
		{
			name: "Should return an error - gameis over",
			args: args{id: "1001", row: 2, col: 2},
			want: want{err: errors.New(apperrors.IllegalOperation, nil, "game is over")},
			mocks: func(m mocks) {
				game := easymockGame("1001", "mygame", 4, domain.GAME_STATE_LOST, false, []pos{{1, 1}}, []pos{})

				m.gameRepository.EXPECT().Get("1001").Return(game, nil)
			},
		},
		{
			name: "Should return an error - save game has fail",
			args: args{id: "1001", row: 2, col: 2},
			want: want{err: errors.New(apperrors.Internal, nil, "update game into repository has failed")},
			mocks: func(m mocks) {
				game := easymockGame("1001", "mygame", 4, "", false, []pos{{1, 1}}, []pos{})
				gameToSave := easymockGame("1001", "mygame", 4, "", false, []pos{{1, 1}}, []pos{{2, 2}})

				m.gameRepository.EXPECT().Get("1001").Return(game, nil)
				m.gameRepository.EXPECT().Save(gameToSave).Return(errors.New(apperrors.Internal, nil, ""))
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
		result, err := service.Reveal(tt.args.id, tt.args.row, tt.args.col)

		// Verify
		if tt.want.err != nil && err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)
	}
}

type pos struct {
	row uint
	col uint
}

func easymockGame(id string, name string, size uint, state string, hideBombs bool, bombs []pos, revealed []pos) domain.Game {
	game := domain.NewGame(id, name, size, 0)
	game.BoardSettings.Bombs = uint(len(bombs))

	for _, pos := range bombs {
		game.Board[pos.row][pos.col] = domain.CELL_BOMB
	}

	for _, pos := range revealed {
		game.Board[pos.row][pos.col] = domain.CELL_REVEALED
	}

	if hideBombs {
		game.Board = game.Board.HideBombs()
	}

	if state != "" {
		game.State = state
	}

	return game
}
