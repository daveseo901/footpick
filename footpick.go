package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"

    "github.com/daveseo901/footpick/draft"
    "github.com/daveseo901/footpick/league" 
    "github.com/daveseo901/footpick/display" 
)

func main() {
    // TODO: hook up database (do we even need one?)
    numPlayers := draft.Draft()
    fmt.Printf("%d players\n", numPlayers)
    fmt.Println(league.AllTeams)
    bos, e := league.NewTeam("Boston", "Lobsters", "BOS")
    if e != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", e)
    }
    league.AllTeams = append(league.AllTeams, *bos)
    for _, t := range league.AllTeams {
        fmt.Println(t)
    }
    var players []*league.Player
    var player *league.Player
    player, e = league.NewPlayer("David", "Seo", "QB")
    if e != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", e)
        os.Exit(1)
    }
    players = append(players, player)
    e = player.UpdateAttr("SPD", 85)
    if e != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", e)
        os.Exit(1)
    }
    for _, p := range players {
        fmt.Println(*p)
    }

    display, e := display.NewDisplay()
    e = display.Write("hello")
    if e != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", e)
        os.Exit(1)
    }
    e = display.Write("welcome to footpick")
    if e != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", e)
        os.Exit(1)
    }
    e = display.Write("what is your name?")
    if e != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", e)
        os.Exit(1)
    }
    index := display.Render(0)
    if (index < 0) {
        fmt.Fprintf(os.Stderr, "error: we fucked up\n")
        os.Exit(1)
    }
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input = strings.TrimSuffix(input, "\n")
    display.Flush()
    display.Write(fmt.Sprintf("hello %s", input))
    display.Render(0)
    input, _ = reader.ReadString('\n')
}
