package arknights

type material string

const (
	// Polyester is "初級エステル".
	Polyester material = "Polyester"
	// PolyesterLump is "上級エステル".
	PolyesterLump material = "Polyester Lump"
	// Sugar is "初級糖原".
	Sugar material = "Sugar"
	// SugarPack is "中級糖原".
	SugarPack material = "Sugar Pack"
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
	// Polyketon is "初級アケトン".
	Polyketon material = "Polyketon"
	// Aketon is "中級アケトン".
	Aketon material = "Aketon"
	// KetonColloid is "上級アケトン".
	KetonColloid material = "Keton Colloid"
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
	// IncandescentAlloyBlock is "上級熾合金".
	IncandescentAlloyBlock material = "Incandescent Alloy Block"
	// BipolarNanoflake is "ナノフレーク".
	BipolarNanoflake material = "Bipolar Nanoflake"
	// D32Steel is "D32鋼".
	D32Steel material = "D32 Steel"
	// ChipCatalyst is "SoC強化剤".
	ChipCatalyst material = "Chip Catalyst"
)

// Materials holds name-amount pairs.
type Materials map[material]int

func (m Materials) merge(x Materials) Materials {
	res := make(Materials)

	for k, v := range m {
		res[k] += v
	}

	for k, v := range x {
		res[k] += v
	}

	return res
}

func (m Materials) multiply(x int) Materials {
	res := make(Materials)

	for k, v := range m {
		res[k] += v * x
	}

	return res
}
