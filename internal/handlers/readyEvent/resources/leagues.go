package leagues

type LeagueMessage struct {
	LeagueName      string
	ReactionStrings []string // These are strings of form <name:id> and they can be obtained on the server through sending \:emojiname:
}

// Returns a list of messages and their associated reactions to be sent on the team-assign channel.
func GetLeaguesAsList() []LeagueMessage {

	return []LeagueMessage{
		{
			LeagueName: "Liga 1",
			ReactionStrings: []string{
				"FCSB:1165017898666954874",
				"ASC_Otelul_Galati:1165018101729984542",
				"UTA_Arad_logo:1165018059002609664",
			},
		},
		{
			LeagueName:      "Liga 2",
			ReactionStrings: []string{},
		},
	}

}
