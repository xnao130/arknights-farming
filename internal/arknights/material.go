package arknights

type material string

const (
	// Polyester is "初級エステル".
	Polyester material = "Polyester"
	// PolyesterLump is "上級エステル".
	PolyesterLump material = "Polyester Lump"
	// SugarSubstitute is "ブドウ糖".
	SugarSubstitute material = "Sugar Substitute"
	// Sugar is "初級糖原".
	Sugar material = "Sugar"
	// SugarPack is "中級糖原".
	SugarPack material = "Sugar Pack"
	// SugarLump is "上級糖原".
	SugarLump material = "Sugar Lump"
	// OrironShard is "異鉄の欠片".
	OrironShard material = "Oriron Shard"
	// Oriron is "初級異鉄".
	Oriron material = "Oriron"
	// OrironCluster is "中級異鉄".
	OrironCluster material = "Oriron Cluster"
	// OrironBlock is "上級異鉄".
	OrironBlock material = "Oriron Block"
	// Device is "初級装置".
	Device material = "Device"
	// OptimizedDevice is "上級装置".
	OptimizedDevice material = "Optimized Device"
	// OrirockCube is "初級源岩".
	OrirockCube material = "Orirock Cube"
	// OrirockCluster is "中級源岩".
	OrirockCluster material = "Orirock Cluster"
	// OrirockConcentration is "上級源岩".
	OrirockConcentration material = "Orirock Concentration"
	// Diketon is "アケトン試剤".
	Diketon material = "Diketon"
	// Polyketon is "初級アケトン".
	Polyketon material = "Polyketon"
	// Aketon is "中級アケトン".
	Aketon material = "Aketon"
	// KetonColloid is "上級アケトン".
	KetonColloid material = "Keton Colloid"
	// ManganeseOre is "マンガン".
	ManganeseOre material = "Manganese Ore"
	// ManganeseTrihydrate is "上級マンガン".
	ManganeseTrihydrate material = "Manganese Trihydrate"
	// Grindstone is "砥石".
	Grindstone material = "Grindstone"
	// GrindstonePentahydrate is "上級砥石".
	GrindstonePentahydrate material = "Grindstone Pentahydrate"
	// LoxicKohl is "合成コール".
	LoxicKohl material = "Loxic Kohl"
	// WhiteHorseKohl is "上級合成コール".
	WhiteHorseKohl material = "White Horse Kohl"
	// RMA7012 is "RMA70-12".
	RMA7012 material = "RMA70-12"
	// RMA7024 is "RMA70-24".
	RMA7024 material = "RMA70-24"
	// IncandescentAlloyBlock is "上級熾合金".
	IncandescentAlloyBlock material = "Incandescent Alloy Block"
	// BipolarNanoflake is "ナノフレーク".
	BipolarNanoflake material = "Bipolar Nanoflake"
	// PolymerizationPreparation is "融合剤".
	PolymerizationPreparation material = "Polymerization Preparation"
	// D32Steel is "D32鋼".
	D32Steel material = "D32 Steel"
	// ChipCatalyst is "SoC強化剤".
	ChipCatalyst material = "Chip Catalyst"
	// SkillSummary1 is "アーツ学I".
	SkillSummary1 material = "Skill Summary - 1"
	// SkillSummary2 is "アーツ学II".
	SkillSummary2 material = "Skill Summary - 2"
	// SkillSummary3 is "アーツ学III".
	SkillSummary3 material = "Skill Summary - 3"
)

// Materials holds name-amount pairs.
type Materials map[material]int

func (m Materials) normalize() Materials {
	res := make(Materials)

	for k, v := range m {
		if v != 0 {
			res[k] = v
		}
	}

	return res
}

func (m Materials) multiply(x int) Materials {
	res := make(Materials)

	for k, v := range m {
		res[k] += v * x
	}

	return res.normalize()
}

// Add adds x to m.
func (m Materials) Add(x Materials) Materials {
	res := make(Materials)

	for k, v := range m {
		res[k] += v
	}

	for k, v := range x {
		res[k] += v
	}

	return res.normalize()
}

// Subtract subtracts x from m.
func (m Materials) Subtract(x Materials) Materials {
	res := make(Materials)

	for k, v := range m {
		res[k] += v
	}

	for k, v := range x {
		res[k] -= v
	}

	return res.normalize()
}
