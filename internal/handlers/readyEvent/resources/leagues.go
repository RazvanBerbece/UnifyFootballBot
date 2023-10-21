package leagues

type LeagueMessage struct {
	LeagueName      string
	ReactionStrings []string
}

// Returns a list of messages and their associated reactions to be sent on the team-assign channel.
func GetLeaguesAsList() []LeagueMessage {

	return []LeagueMessage{
		{
			LeagueName:      "Liga 1",
			ReactionStrings: []string{":FCSB:", ":ASC_Otelul_Galati:", ":UTA_Arad_logo:"},
		},
		{
			LeagueName:      "Liga 2",
			ReactionStrings: []string{},
		},
	}

}
