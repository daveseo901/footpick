package league

type League struct {
    teams []Team
    freeAgents []FreeAgency
    draftPool []DraftPool
}

func (l League) AddTeam(team Team) {
    l.teams = append(l.teams, team)
}
