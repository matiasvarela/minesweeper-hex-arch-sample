package ports

import "github.com/matiasvarela/minesweeper-hex-arch-sample/internal/core/domain"

type GamesRepository interface {
	Get(id string) (domain.Game, error)
	Create(domain.Game) error
	Update(domain.Game) error
}
