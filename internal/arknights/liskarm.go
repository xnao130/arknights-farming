package arknights

// Liskarm is "リスカム".
var Liskarm = operator{
	rarity: 5,
	class:  Defender,
	elite1: Materials{
		Sugar:       5,
		OrirockCube: 3,
	},
	elite2: Materials{
		GrindstonePentahydrate: 7,
		Aketon:                 15,
	},
	skill: []Materials{
		{},
		{
			SugarSubstitute: 7,
		},
		{
			Polyester: 3,
		},
		{
			Oriron: 4,
		},
		{
			PolyesterPack: 5,
		},
		{
			OrirockCluster: 2,
			Grindstone:     3,
		},
	},
	s1: []Materials{
		{
			KetonColloid:  3,
			PolyesterPack: 3,
		},
		{
			WhiteHorseKohl: 3,
			KetonColloid:   5,
		},
		{
			BipolarNanoflake: 4,
			KetonColloid:     3,
		},
	},
	s2: []Materials{
		{
			OptimizedDevice: 2,
			OrirockCluster:  3,
		},
		{
			ManganeseTrihydrate: 3,
			OptimizedDevice:     4,
		},
		{
			D32Steel:  4,
			SugarLump: 4,
		},
	},
}
