package league

type PlayerGroup struct {
    players []Player
}

type FreeAgency struct {
    PlayerGroup
}

type DraftPool struct {
    year int
    players []Player
}
