package arknights

// Amiya is "アーミヤ".
var Amiya = operator{
	rarity: 5,
	class:  Caster,
	elite1: Materials{
		Device: 4,
		Oriron: 4,
	},
	elite2: Materials{
		OrirockConcentration: 10,
		LoxicKohl:            10,
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
			Aketon: 4,
		},
		{
			IntegratedDevice: 2,
			SugarPack:        3,
		},
	},
	s1: []Materials{
		{
			WhiteHorseKohl: 3,
			Aketon:         5,
		},
		{
			GrindstonePentahydrate: 3,
			WhiteHorseKohl:         3,
		},
		{
			D32Steel:             4,
			OrirockConcentration: 5,
		},
	},
	s2: []Materials{
		{
			ManganeseTrihydrate: 3,
			IntegratedDevice:    2,
		},
		{
			RMA7024:             3,
			ManganeseTrihydrate: 5,
		},
		{
			PolymerizationPreparation: 4,
			WhiteHorseKohl:            5,
		},
	},
	s3: []Materials{
		{
			GrindstonePentahydrate: 3,
			LoxicKohl:              4,
		},
		{
			OrirockConcentration:   3,
			GrindstonePentahydrate: 6,
		},
		{
			BipolarNanoflake: 4,
			PolyesterLump:    4,
		},
	},
}
