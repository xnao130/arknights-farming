package arknights

// Siege is "シージ".
var Siege = operator{
	rarity: 6,
	class:  Vanguard,
	elite1: Materials{
		Sugar:  9,
		Device: 3,
	},
	elite2: Materials{
		BipolarNanoflake:     4,
		OrirockConcentration: 6,
	},
	skill: []Materials{
		{},
		{
			SugarSubstitute: 5,
			Diketon:         4,
		},
		{
			Polyester: 5,
		},
		{
			Oriron: 4,
			Sugar:  3,
		},
		{
			Aketon: 6,
		},
		{
			IntegratedDevice: 3,
			SugarPack:        5,
		},
	},
	s1: []Materials{
		{
			WhiteHorseKohl: 4,
			Aketon:         8,
		},
		{
			RMA7024:             4,
			ManganeseTrihydrate: 7,
		},
		{
			BipolarNanoflake:     6,
			OrirockConcentration: 6,
		},
	},
	s2: []Materials{
		{
			ManganeseTrihydrate: 4,
			IntegratedDevice:    4,
		},
		{
			OrirockConcentration:   4,
			GrindstonePentahydrate: 9,
		},

		{
			BipolarNanoflake: 6,
			PolyesterLump:    5,
		},
	},
	s3: []Materials{
		{
			GrindstonePentahydrate: 4,
			LoxicKohl:              7,
		},
		{
			SugarLump: 4,
			RMA7024:   8,
		},
		{
			PolymerizationPreparation: 6,
			OptimizedDevice:           4,
		},
	},
}
