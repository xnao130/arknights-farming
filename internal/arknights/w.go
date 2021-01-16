package arknights

// W is "W".
var W = operator{
	rarity: 6,
	class:  Sniper,
	elite1: Materials{
		OrirockCube: 12,
		Sugar:       5,
	},
	elite2: Materials{
		BipolarNanoflake: 4,
		KetonColloid:     7,
	},
	skill: []Materials{
		{},
		{
			Diketon: 6,
			Orirock: 4,
		},
		{
			Device: 3,
		},
		{
			OrirockCube: 5,
			Device:      3,
		},
		{
			IntegratedDevice: 4,
		},
		{
			ManganeseOre: 3,
			RMA7012:      4,
		},
	},
	s1: []Materials{
		{
			IncandescentAlloyBlock: 4,
			RMA7012:                5,
		},
		{
			RMA7024:             4,
			ManganeseTrihydrate: 7,
		},
		{
			D32Steel:       6,
			PolymerizedGel: 6,
		},
	},
	s2: []Materials{
		{
			OptimizedDevice: 3,
			OrironCluster:   4,
		},
		{
			PolymerizedGel:       4,
			OrirockConcentration: 10,
		},
		{
			BipolarNanoflake:       6,
			IncandescentAlloyBlock: 5,
		},
	},
	s3: []Materials{
		{
			WhiteHorseKohl: 4,
			Aketon:         8,
		},
		{
			OrironBlock:            4,
			IncandescentAlloyBlock: 7,
		},
		{
			BipolarNanoflake: 6,
			PolymerizedGel:   6,
		},
	},
}
