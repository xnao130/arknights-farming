package arknights

// Bibeak is "バイビーク".
var Bibeak = operator{
	rarity: 5,
	class:  Guard,
	elite1: Materials{
		Oriron:      4,
		OrirockCube: 3,
	},
	elite2: Materials{
		ManganeseTrihydrate: 8,
		RMA7012:             8,
	},
	skill: []Materials{
		{},
		{
			OrironShard: 5,
		},
		{
			Polyketon: 3,
		},
		{
			Device: 3,
		},
		{
			ManganeseOre: 4,
		},
		{
			LoxicKohl: 3,
			Aketon:    3,
		},
	},
	s1: []Materials{
		{
			PolymerizedGel: 3,
			OrirockCluster: 6,
		},
		{
			ManganeseTrihydrate: 3,
			OptimizedDevice:     4,
		},
		{
			BipolarNanoflake: 4,
			PolymerizedGel:   4,
		},
	},
	s2: []Materials{
		{
			OrironBlock:       3,
			IncandescentAlloy: 1,
		},
		{
			OrirockConcentration:   3,
			GrindstonePentahydrate: 6,
		},
		{
			D32Steel:     4,
			KetonColloid: 4,
		},
	},
}
