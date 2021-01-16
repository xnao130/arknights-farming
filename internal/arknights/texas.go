package arknights

// Texas is "テキサス".
var Texas = operator{
	rarity: 5,
	class:  Vanguard,
	elite1: Materials{
		Polyester: 5,
		Oriron:    3,
	},
	elite2: Materials{
		PolyesterLump:  8,
		OrirockCluster: 16,
	},
	skill: []Materials{
		{},
		{
			Ester: 7,
		},
		{
			Oriron: 3,
		},
		{
			Polyketon: 4,
		},
		{
			IntegratedDevice: 3,
		},
		{
			LoxicKohl: 3,
			Aketon:    3,
		},
	},
	s1: []Materials{
		{
			ManganeseTrihydrate: 3,
			IntegratedDevice:    2,
		},
		{
			RMA7024:             3,
			ManganeseTrihydrate: 5,
		},
		{
			D32Steel:      4,
			PolyesterLump: 4,
		},
	},
	s2: []Materials{
		{
			GrindstonePentahydrate: 3,
			LoxicKohl:              4,
		},
		{
			OrirockConcentration:   3,
			GrindstonePentahydrate: 6,
		},
		{
			PolymerizationPreparation: 4,
			OptimizedDevice:           3,
		},
	},
}
