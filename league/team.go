package league

import (
	"errors"
	"fmt"
	"strings"
)

var league = League{}

// TODO: initialize this list in the driver function
// TODO: change to local list to support multiple leagues
// global list of teams in the league
var AllTeams = []Team{
	Team{"Houston", "Longhorns", "HOU", league, PlayerGroup{[]Player{}}},
	Team{"Las Vegas", "Rattlers", "LVR", league, PlayerGroup{[]Player{}}},
	Team{"Seattle", "Needlers", "SEA", league, PlayerGroup{[]Player{}}},
	Team{"Washington", "Washers", "WAS", league, PlayerGroup{[]Player{}}},
}

// IDs
type Team struct {
	city   string
	name   string
	abbr   string
	league League
	PlayerGroup
}

// string formatting for a team
func (t Team) String() string {
	return fmt.Sprintf("%v %v", t.city, t.name)
}

func NewTeam(city string, name string, abbr string) (*Team, error) {
	if len(abbr) != 3 {
		return nil, errors.New("Team abbreviation must be exactly 3 characters")
	}
	if abbr != strings.ToUpper(abbr) {
		return nil, errors.New("Team abbreviation must be all uppercase")
	}
	return &Team{city, name, abbr, league, PlayerGroup{}}, nil
}

func (pg PlayerGroup) MovePlayer(player Player, recGroup PlayerGroup) error {
	if e := pg.DropPlayer(player); e != nil {
		return e
	}
	recGroup.AddPlayer(player)
	return nil
}

func (pg PlayerGroup) AddPlayer(player Player) {
	pg.players = append(pg.players, player)
}

// TODO: Find and drop player from group. If not found, return an error
func (pg PlayerGroup) DropPlayer(player Player) error {
	for i, p := range pg.players {
		if p.Compare(player) {
			pg.players[i] = pg.players[len(pg.players)-1]
			pg.players = pg.players[:len(pg.players)-1]
			return nil
		}
	}
	return fmt.Errorf("Player not found in %v", pg)
}
