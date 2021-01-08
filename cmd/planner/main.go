package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/xnao130/arknights-farming/internal/arknights"
)

func main() {
	var plans []arknights.Plan

	bytes, err := ioutil.ReadFile("operators.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(bytes, &plans)

	socs := arknights.CountSoC(plans...)

	bytes, err = json.MarshalIndent(&socs, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	materials := arknights.CountPromotionMaterials(plans...)

	bytes, err = json.MarshalIndent(&materials, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
