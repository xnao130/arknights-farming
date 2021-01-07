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

	res := arknights.CountSoC(plans...)

	bytes, err = json.MarshalIndent(&res, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
