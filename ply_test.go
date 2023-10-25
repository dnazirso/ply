package ply

import (
	"html/template"
	"os"
	"testing"
)

func TestPly(t *testing.T) {
	type Item struct {
		Title string
		Is    bool
	}

	type Data struct {
		Title string
		List  []Item
	}

	index := Fold("tests/index", "")

	tmpl, err := template.New("index").Parse(index)

	data := Data{
		Title: "List",
		List: []Item{
			{Title: "One", Is: false},
			{Title: "Two", Is: true},
		},
	}

	err = tmpl.Execute(os.Stdout, data)

	if err != nil {
		t.Errorf(err.Error())
	}
}
