package draft

// TODO: implement everything
// create a draft pool of players
func createDraftPool() map[string]string {
    draftPool := make(map[string]string)
    draftPool["Arian Foster"] = "RB"
    draftPool["Andre Johnson"] = "WR"
    return draftPool
}

// run a draft
func Draft() int {
    draftPool := createDraftPool()
    return len(draftPool)
}
