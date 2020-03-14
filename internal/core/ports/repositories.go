package ports

import "github.com/matiasvarela/minesweeper-hex-arch-sample/internal/core/domain"

type GamesRepository interface {
	Get(id string) (domain.Game, error)
	Save(domain.Game) error
}
