package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

const (
	operatoerTemplate = `package arknights

// {{.En}} is "{{.Ja}}".
var {{.En}} = operator{
	rarity: {{.Rarity}},
	class:  Specialist,
	elite1: Materials{},
	elite2: Materials{},
	skill: []Materials{
		{},
		{},
		{},
		{},
		{},
		{},
	},
	s1: []Materials{
		{},
		{},
		{},
	},
	s2: []Materials{
		{},
		{},
		{},
	},
	s3: []Materials{
		{},
		{},
		{},
	},
}	
`
)

func main() {
	en := flag.String("en", "", "operator name in english")
	ja := flag.String("ja", "", "operator name in japanse")
	rarity := flag.Int("rarity", 0, "operator rarity")

	flag.Parse()

	values := struct {
		En     string
		Ja     string
		Rarity int
	}{
		En:     strings.Title(*en),
		Ja:     *ja,
		Rarity: *rarity,
	}

	tpl, err := template.New(values.En).Parse(operatoerTemplate)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, values)
	if err != nil {
		panic(err)
	}

	file := fmt.Sprintf("internal/arknights/%s.go", strings.ToLower(*en))

	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Printf("write %s\n", file)

		err = ioutil.WriteFile(file, buf.Bytes(), 0644)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("%s already exists", file)
	}
}
