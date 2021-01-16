package arknights

// Ethan is "イーサン".
var Ethan = operator{
	rarity: 4,
	class:  Specialist,
	elite1: Materials{
		Sugar:  1,
		Oriron: 1,
	},
	elite2: Materials{
		SugarPack:      17,
		OrirockCluster: 14,
	},
	skill: []Materials{
		{},
		{
			SugarSubstitute: 4,
		},
		{
			Polyester: 2,
		},
		{
			Oriron: 2,
		},
		{
			Aketon: 2,
		},
		{
			IntegratedDevice: 2,
		},
	},
	s1: []Materials{
		{
			WhiteHorseKohl: 1,
			Aketon:         4,
		},
		{
			GrindstonePentahydrate: 2,
			WhiteHorseKohl:         3,
		},
		{
			PolymerizationPreparation: 2,
			RMA7024:                   2,
		},
	},
	s2: []Materials{
		{
			ManganeseTrihydrate: 1,
			IntegratedDevice:    3,
		},
		{
			RMA7024:             2,
			ManganeseTrihydrate: 2,
		},
		{
			BipolarNanoflake: 2,
			SugarLump:        2,
		},
	},
}
