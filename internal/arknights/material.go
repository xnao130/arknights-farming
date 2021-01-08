package arknights

type material string

const (
	// Polyester is "初級エステル".
	Polyester material = "Polyester"
	// Sugar is "初級糖原".
	Sugar material = "Sugar"
	// OrironBlock is "上級異鉄".
	OrironBlock material = "Oriron Block"
	// Device is "初級装置".
	Device material = "Device"
	// OrirockCube is "初級源岩".
	OrirockCube material = "Orirock Cube"
	// OrirockConcentration is "上級源岩".
	OrirockConcentration material = "Orirock Concentration"
	// KetonColloid is "上級アケトン".
	KetonColloid material = "Keton Colloid"
	// BipolarNanoflake is "ナノフレーク".
	BipolarNanoflake material = "Bipolar Nanoflake"
	// D32Steel is "D32鋼".
	D32Steel material = "D32 Steel"
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
