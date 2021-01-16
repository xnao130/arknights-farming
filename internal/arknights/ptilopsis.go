package arknights

// Ptilopsis is "フィリオプシス".
var Ptilopsis = operator{
	rarity: 5,
	class:  Medic,
	elite1: Materials{
		OrirockCube: 8,
		Sugar:       2,
	},
	elite2: Materials{
		OrirockConcentration: 9,
		Grindstone:           10,
	},
	skill: []Materials{
		{},
		{
			Orirock: 10,
		},
		{
			Sugar: 3,
		},
		{
			Polyester: 5,
		},
		{
			OrirockCluster: 4,
		},
		{
			Aketon:    2,
			LoxicKohl: 4,
		},
	},
	s1: []Materials{
		{
			OptimizedDevice: 2,
			OrironCluster:   3,
		},
		{
			ManganeseTrihydrate: 3,
			OptimizedDevice:     4,
		},
		{
			PolymerizationPreparation: 4,
			OrirockConcentration:      4,
		},
	},
	s2: []Materials{
		{
			WhiteHorseKohl: 3,
			Aketon:         5,
		},
		{
			GrindstonePentahydrate: 3,
			WhiteHorseKohl:         6,
		},
		{
			BipolarNanoflake:     4,
			OrirockConcentration: 4,
		},
	},
}
