package models

type GameType int

const (
    Game1 GameType = iota + 1
    Game2
    Game3
)

func (g GameType) String() string {
    switch g {
    case Game1:
        return "Game1"
    case Game2:
        return "Game2"
    case Game3:
        return "Game3"
    default:
        return "Unknown"
    }
}