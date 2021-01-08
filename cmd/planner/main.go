package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/xnao130/arknights-farming/internal/arknights"
)

func main() {
	var input struct {
		Plans    []arknights.Plan
		Workshop arknights.Materials
	}

	bytes, err := ioutil.ReadFile("plan.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(bytes, &input)

	materials := arknights.CountPromotionMaterials(input.Plans...).Subtract(input.Workshop).Add(arknights.CountWorkshop(input.Workshop))

	bytes, err = json.MarshalIndent(&materials, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
