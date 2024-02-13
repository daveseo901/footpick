package league

import (
    "testing"
)

// NewPlayer should return a valid *Player when given a valid position
func TestValidPosition(t *testing.T) {
    if p, _ := NewPlayer("John", "Smith", "RB"); p == nil {
        t.Errorf("NewPlayer('John', 'Smith', 'RB') = %v; " +
            "want &{'John' 'Smith' 'RB'}", p)
    }
}

// NewPlayer should return `nil` when given an invalid position
func TestInvalidPosition(t *testing.T) {
    if p, _ := NewPlayer("John", "Smith", "RC"); p != nil {
        t.Errorf("NewPlayer('John', 'Smith', 'RC') = %v; " +
            "want <nil>", p)
    }
}
