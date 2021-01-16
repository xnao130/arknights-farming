package arknights

// DurNar is "ジュナー".
var DurNar = operator{
	rarity: 4,
	class:  Defender,
	elite1: Materials{
		OrirockCube: 1,
		Polyester:   1,
	},
	elite2: Materials{
		OrirockCluster: 19,
		RMA7012:        8,
	},
	skill: []Materials{
		{},
		{
			Orirock: 5,
		},
		{
			Sugar: 2,
		},
		{
			Polyester: 3,
		},
		{
			Oriron: 2,
		},
		{
			Aketon: 3,
		},
	},
	s1: []Materials{
		{
			KetonColloid:  1,
			PolyesterPack: 4,
		},
		{
			WhiteHorseKohl: 2,
			KetonColloid:   2,
		},
		{
			BipolarNanoflake:    2,
			ManganeseTrihydrate: 2,
		},
	},
	s2: []Materials{
		{
			OptimizedDevice: 1,
			OrironCluster:   2,
		},
		{
			ManganeseTrihydrate: 2,
			OptimizedDevice:     2,
		},
		{
			D32Steel:       2,
			WhiteHorseKohl: 2,
		},
	},
}
