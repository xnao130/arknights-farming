package arknights

// Weedy is "ウィーディ".
var Weedy = operator{
	rarity: 6,
	class:  Specialist,
	elite1: Materials{
		Device: 6,
		Sugar:  4,
	},
	elite2: Materials{
		D32Steel:            4,
		ManganeseTrihydrate: 6,
	},
	skill: []Materials{
		{},
		{
			DamagedDevice: 4,
			Ester:         4,
		},
		{
			OrirockCube: 7,
		},
		{
			Sugar:     4,
			Polyketon: 4,
		},
		{
			LoxicKohl: 7,
		},
		{
			Orirock:          5,
			IntegratedDevice: 4,
		},
	},
	s1: []Materials{
		{
			ManganeseTrihydrate: 4,
			IntegratedDevice:    4,
		},
		{
			OrirockConcentration:   4,
			GrindstonePentahydrate: 9,
		},
		{
			BipolarNanoflake:       6,
			IncandescentAlloyBlock: 5,
		},
	},
	s2: []Materials{
		{
			GrindstonePentahydrate: 4,
			LoxicKohl:              7,
		},
		{
			OptimizedDevice: 3,
			OrironBlock:     6,
		},
		{
			BipolarNanoflake: 6,
			PolymerizedGel:   6,
		},
	},
	s3: []Materials{
		{
			RMA7024:      4,
			ManganeseOre: 5,
		},
		{
			WhiteHorseKohl: 4,
			KetonColloid:   8,
		},
		{
			PolymerizationPreparation: 6,
			PolymerizedGel:            7,
		},
	},
}
