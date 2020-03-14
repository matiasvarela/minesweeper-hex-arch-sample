package domain

const (
	GAME_STATE_WON  = "won"
	GAME_STATE_LOST = "lost"
	GAME_STATE_NEW  = "new"
)

type Game struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	State         string        `json:"state"`
	BoardSettings BoardSettings `json:"board_settings"`
	Board         Board         `json:"board"`
}

func NewGame(id string, name string, size uint, bombs uint) Game {
	return Game{
		ID:    id,
		Name:  name,
		State: GAME_STATE_NEW,
		BoardSettings: BoardSettings{
			Size:  size,
			Bombs: bombs,
		},
		Board: NewBoard(size, bombs),
	}
}
func (game *Game) IsOver() bool {
	return game.State == GAME_STATE_LOST || game.State == GAME_STATE_WON
}
