package arknights

// Elysium is "エリジウム".
var Elysium = operator{
	rarity: 5,
	class:  Vanguard,
	elite1: Materials{
		Polyester: 4,
		Sugar:     3,
	},
	elite2: Materials{
		IncandescentAlloyBlock: 7,
		Aketon:                 16,
	},
	skill: []Materials{
		{},
		{
			Ester: 7,
		},
		{
			Oriron: 3,
		},
		{
			Polyketon: 4,
		},
		{
			IntegratedDevice: 3,
		},
		{
			Grindstone:     2,
			OrirockCluster: 4,
		},
	},
	s1: []Materials{
		{
			PolymerizedGel: 3,
			OrirockCluster: 6,
		},
		{
			IncandescentAlloyBlock: 3,
			RMA7024:                5,
		},
		{
			D32Steel:     4,
			KetonColloid: 4,
		},
	},
	s2: []Materials{
		{
			OptimizedDevice: 2,
			OrironCluster:   3,
		},
		{
			KetonColloid:   3,
			PolymerizedGel: 6,
		},
		{
			PolymerizationPreparation: 4,
			GrindstonePentahydrate:    4,
		},
	},
}
