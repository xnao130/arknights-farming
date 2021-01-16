package arknights

// Schwarz is "シュヴァルツ".
var Schwarz = operator{
	rarity: 6,
	class:  Sniper,
	elite1: Materials{
		Polyester: 8,
		Sugar:     6,
	},
	elite2: Materials{
		D32Steel:    4,
		OrironBlock: 5,
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
			OrironShard: 4,
			Sugar:       3,
		},
		{
			LoxicKohl: 7,
		},
		{
			ManganeseOre: 3,
			RMA7012:      4,
		},
	},
	s1: []Materials{
		{
			GrindstonePentahydrate: 4,
			LoxicKohl:              7,
		},
		{
			SugarLump: 4,
			RMA7024:   8,
		},
		{
			BipolarNanoflake: 6,
			PolyesterLump:    5,
		},
	},
	s2: []Materials{
		{
			RMA7024:      4,
			ManganeseOre: 5,
		},
		{
			PolyesterLump:        4,
			OrirockConcentration: 10,
		},
		{
			PolymerizationPreparation: 6,
			OptimizedDevice:           4,
		},
	},
	s3: []Materials{
		{
			OrirockConcentration: 4,
			Grindstone:           7,
		},
		{
			OrironBlock: 4,
			SugarLump:   7,
		},
		{
			PolymerizationPreparation: 6,
			ManganeseTrihydrate:       6,
		},
	},
}
