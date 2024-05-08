package pokeapi

type EvolutionChainResp struct {
	BabyTriggerItem any `json:"baby_trigger_item"`
	Chain           struct {
		EvolutionDetails []any `json:"evolution_details"`
		EvolvesTo        []struct {
			EvolutionDetails []struct {
				Gender                any    `json:"gender"`
				HeldItem              any    `json:"held_item"`
				Item                  any    `json:"item"`
				KnownMove             any    `json:"known_move"`
				KnownMoveType         any    `json:"known_move_type"`
				Location              any    `json:"location"`
				MinAffection          any    `json:"min_affection"`
				MinBeauty             any    `json:"min_beauty"`
				MinHappiness          any    `json:"min_happiness"`
				MinLevel              int    `json:"min_level"`
				NeedsOverworldRain    bool   `json:"needs_overworld_rain"`
				PartySpecies          any    `json:"party_species"`
				PartyType             any    `json:"party_type"`
				RelativePhysicalStats any    `json:"relative_physical_stats"`
				TimeOfDay             string `json:"time_of_day"`
				TradeSpecies          any    `json:"trade_species"`
				Trigger               struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"trigger"`
				TurnUpsideDown bool `json:"turn_upside_down"`
			} `json:"evolution_details"`
			EvolvesTo []struct {
				EvolutionDetails []struct {
					Gender                any    `json:"gender"`
					HeldItem              any    `json:"held_item"`
					Item                  any    `json:"item"`
					KnownMove             any    `json:"known_move"`
					KnownMoveType         any    `json:"known_move_type"`
					Location              any    `json:"location"`
					MinAffection          any    `json:"min_affection"`
					MinBeauty             any    `json:"min_beauty"`
					MinHappiness          any    `json:"min_happiness"`
					MinLevel              int    `json:"min_level"`
					NeedsOverworldRain    bool   `json:"needs_overworld_rain"`
					PartySpecies          any    `json:"party_species"`
					PartyType             any    `json:"party_type"`
					RelativePhysicalStats any    `json:"relative_physical_stats"`
					TimeOfDay             string `json:"time_of_day"`
					TradeSpecies          any    `json:"trade_species"`
					Trigger               struct {
						Name string `json:"name"`
						URL  string `json:"url"`
					} `json:"trigger"`
					TurnUpsideDown bool `json:"turn_upside_down"`
				} `json:"evolution_details"`
				EvolvesTo []any `json:"evolves_to"`
				IsBaby    bool  `json:"is_baby"`
				Species   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"species"`
			} `json:"evolves_to"`
			IsBaby  bool `json:"is_baby"`
			Species struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"species"`
		} `json:"evolves_to"`
		IsBaby  bool `json:"is_baby"`
		Species struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"species"`
	} `json:"chain"`
	ID int `json:"id"`
}