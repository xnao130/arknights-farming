package arknights

// Lappland is "ラップランド".
var Lappland = operator{
	rarity: 5,
	class:  Guard,
	elite1: Materials{
		Device:      3,
		OrirockCube: 4,
	},
	elite2: Materials{
		OptimizedDevice: 6,
		OrironCluster:   10,
	},
	skill: []Materials{
		{},
		{
			DamagedDevice: 4,
		},
		{
			OrirockCube: 4,
		},
		{
			Sugar: 5,
		},
		{
			Grindstone: 3,
		},
		{
			RMA7012:       2,
			PolyesterPack: 3,
		},
	},
	s1: []Materials{
		{
			OrirockConcentration: 3,
			Grindstone:           4,
		},
		{
			PolyesterLump:        3,
			OrirockConcentration: 6,
		},
		{
			D32Steel:     4,
			KetonColloid: 4,
		},
	},
	s2: []Materials{
		{
			SugarLump: 3,
			RMA7012:   3,
		},
		{
			OrironBlock: 3,
			SugarLump:   5,
		},
		{
			PolymerizationPreparation: 4,
			GrindstonePentahydrate:    4,
		},
	},
}
