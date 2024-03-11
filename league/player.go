package league

import (
	"errors"
	"fmt"
)

// TODO: change Position to an enum type?
// For now every team runs 11 personnel offense and 4-3 defense
var positions = map[string]struct{}{
	"QB":  {}, // quarterback
	"RB":  {}, // running back
	"WR":  {}, // wide receiver
	"TE":  {}, // tight end
	"OT":  {}, // offensive tackle
	"OG":  {}, // offensive guard
	"C":   {}, // center
	"CB":  {}, // cornerback
	"S":   {}, // safety
	"ILB": {}, // inside linebacker
	"OLB": {}, // outside linebacker
	"DT":  {}, // defensive tackle
	"DE":  {}, // defensive end
}

var attributes = map[string]struct{}{
	"SPD": {}, // speed
	"STR": {}, // strength
	"ACC": {}, // acceleration
	"THP": {}, // throw power
	"THA": {}, // throw accuracy
	"CAR": {}, // carry
	"CAT": {}, // catch
	"BLK": {}, // block
	"TAC": {}, // tackle
	"COV": {}, // coverage
	"BLS": {}, // block shed
	"PSR": {}, // pass rush
}

// TODO: add number and validation
type Player struct {
	FirstName string
	LastName  string
	Pos       string
	Attr      map[string]int
}

type Attribute struct {
	name string
	val  int
}

func NewPlayer(firstName string, lastName string, pos string) (*Player, error) {
	if _, ok := positions[pos]; !ok {
		return nil, errors.New(fmt.Sprintf("Invalid position %s", pos))
	}
	return &Player{firstName, lastName, pos, make(map[string]int)}, nil
}

func (p Player) UpdateAttr(attrName string, attrVal int) error {
	if attrVal < 0 || attrVal > 99 {
		return errors.New("Invalid attribute value")
	}
	if _, ok := attributes[attrName]; !ok {
		return errors.New("Invalid attribute name")
	}
	p.Attr[attrName] = attrVal
	return nil
}

func (p Player) UpdateAttrs(attrs []Attribute) error {
	var e error
	for _, a := range attrs {
		e = p.UpdateAttr(a.name, a.val)
		if e != nil {
			return e
		}
	}
	return nil
}

// TODO: Add IDs to make this easier
func (p Player) Compare(player Player) bool {
	return p.FirstName == player.FirstName && p.LastName == player.LastName
}
