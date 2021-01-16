package arknights

// Utage is "ウタゲ".
var Utage = operator{
	rarity: 4,
	class:  Guard,
	elite1: Materials{
		Device: 1,
		Sugar:  1,
	},
	elite2: Materials{
		Aketon:         14,
		OrirockCluster: 14,
	},
	skill: []Materials{
		{},
		{
			DamagedDevice: 2,
		},
		{
			OrirockCube: 2,
		},
		{
			Sugar: 3,
		},
		{
			LoxicKohl: 2,
		},
		{
			ManganeseOre: 3,
		},
	},
	s1: []Materials{
		{
			RMA7024:      1,
			ManganeseOre: 3,
		},
		{
			OrironBlock:            2,
			IncandescentAlloyBlock: 2,
		},
		{
			PolymerizationPreparation: 2,
			ManganeseTrihydrate:       2,
		},
	},
	s2: []Materials{
		{
			IncandescentAlloyBlock: 1,
			RMA7012:                3,
		},
		{
			KetonColloid:   2,
			PolymerizedGel: 2,
		},
		{
			BipolarNanoflake:       2,
			GrindstonePentahydrate: 2,
		},
	},
}
