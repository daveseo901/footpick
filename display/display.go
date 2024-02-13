package display

import (
    "fmt"
    "errors"
    "strings"

    tsize "github.com/kopoli/go-terminal-size"
)

// TODO: current index?
type Display struct {
    buffer []string
    Dims tsize.Size
}

func NewDisplay() (*Display, error) {
    size, e := tsize.GetSize()
    if e != nil {
        return nil, e
    }
    return &Display{make([]string, 0), size}, nil
}

func (d *Display) Write(line string) error {
    if len(line) > d.Dims.Width {
        return errors.New("Line too long for display")
    }
    d.buffer = append(d.buffer, strings.Clone(line)) 
    return nil
}

func (d Display) ShowBuffer() {
    fmt.Print(d.buffer)
    for _, line := range d.buffer {
        fmt.Println(line)
    }
}

func (d Display) Render(index int) int {
    fmt.Print("\033[H\033[2J")
    for i, line := range d.buffer {
        if i < index {
            continue
        } else if (i - index) == d.Dims.Height {
            return i - 1
        }
        fmt.Println(line) 
        if i == len(d.buffer) - 1 {
            return i
        }
    }
    return -1
}

func (d *Display) Flush() {
    d.buffer = nil
}
