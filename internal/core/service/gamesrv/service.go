package gamesrv

import (
	"github.com/matiasvarela/errors"
	"github.com/matiasvarela/minesweeper-hex-arch-sample/internal/core/domain"
	"github.com/matiasvarela/minesweeper-hex-arch-sample/internal/core/ports"
	"github.com/matiasvarela/minesweeper-hex-arch-sample/pkg/apperrors"
	"github.com/matiasvarela/minesweeper-hex-arch-sample/pkg/uidgen"
)

type service struct {
	gamesRepository ports.GamesRepository
	uidGen          uidgen.UIDGen
}

func New(gamesRepository ports.GamesRepository, uidGen uidgen.UIDGen) ports.GamesService {
	return &service{
		gamesRepository: gamesRepository,
		uidGen:          uidGen,
	}
}

func (srv *service) Get(id string) (domain.Game, error) {
	game, err := srv.gamesRepository.Get(id)
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return domain.Game{}, errors.New(apperrors.NotFound, err, "game not found")
		}

		return domain.Game{}, errors.New(apperrors.Internal, err, "get game from repository has failed")
	}

	game.Board = game.Board.HideBombs()

	return game, nil
}

func (srv *service) Create(name string, size uint, bombs uint) (domain.Game, error) {
	if bombs >= size*size {
		return domain.Game{}, errors.New(apperrors.InvalidInput, nil, "the number of bombs is too high")
	}

	game := domain.NewGame(srv.uidGen.New(), name, size, bombs)

	if err := srv.gamesRepository.Save(game); err != nil {
		return domain.Game{}, errors.New(apperrors.Internal, err, "create game into repository has failed")
	}

	game.Board = game.Board.HideBombs()

	return game, nil
}

func (srv *service) Reveal(id string, row uint, col uint) (domain.Game, error) {
	game, err := srv.gamesRepository.Get(id)
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return domain.Game{}, errors.New(apperrors.NotFound, err, "game not found")
		}

		return domain.Game{}, errors.New(apperrors.Internal, err, "get game from repository has failed")
	}

	if !game.Board.IsValidPosition(row, col) {
		return domain.Game{}, errors.New(apperrors.InvalidInput, nil, "invalid position")
	}

	if game.IsOver() {
		return domain.Game{}, errors.New(apperrors.IllegalOperation, nil, "game is over")
	}

	if game.Board.Contains(row, col, domain.CELL_BOMB) {
		game.State = domain.GAME_STATE_LOST
	} else {
		game.Board.Set(row, col, domain.CELL_REVEALED)

		if !game.Board.HasEmptyCells() {
			game.State = domain.GAME_STATE_WON
		}
	}

	if err := srv.gamesRepository.Save(game); err != nil {
		return domain.Game{}, errors.New(apperrors.Internal, err, "update game into repository has failed")
	}

	game.Board = game.Board.HideBombs()

	return game, nil
}
