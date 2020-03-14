package ports

import "github.com/matiasvarela/minesweeper-hex-arch-sample/internal/core/domain"

type GamesService interface {
	Get(id string) (domain.Game, error)
	Create(name string, size uint, bombs uint) (domain.Game, error)
	Reveal(id string, row uint, col uint) (domain.Game, error)
}
