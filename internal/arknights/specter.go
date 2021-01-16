package arknights

// Specter is "スペクター".
var Specter = operator{
	rarity: 5,
	class:  Guard,
	elite1: Materials{
		OrirockCube: 6,
		Polyketon:   3,
	},
	elite2: Materials{
		WhiteHorseKohl: 8,
		Aketon:         15,
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
			RMA7012: 3,
		},
		{
			OrirockCluster:   4,
			IntegratedDevice: 2,
		},
	},
	s1: []Materials{
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
	s2: []Materials{
		{
			PolyesterLump:  3,
			OrirockCluster: 4,
		},
		{
			KetonColloid:  3,
			PolyesterLump: 3,
		},
		{
			BipolarNanoflake:    4,
			ManganeseTrihydrate: 3,
		},
	},
}
