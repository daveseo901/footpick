package display

import (
    "testing"
)

func TestNewDisplay(t *testing.T) {
    if d, _ := NewDisplay(); d == nil {
        t.Errorf("NewDisplay = %v; want &Display{'' <tsize.Size>}", d)
    } else {
        t.Logf("d = %v", d)
    }
}
