package ply

import (
	"html/template"
	"os"
	"strings"
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

	println(index)
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

func TestCut(t *testing.T) {
	entire := `<ply as="path/to/component">this is the child you want to isolate from the component</ply>`
	ply_part, contain_component_path, has_component := strings.Cut(entire, "\"")
	t.Log(ply_part, contain_component_path, has_component)

	if !strings.Contains(ply_part, "as") {
		t.Error(`doesn't containe path to component ("as" property not found)`)
	}

	if !has_component {
		t.Error(`has no component`)
	}

	component_path, contain_children, has_component := strings.Cut(contain_component_path, `"`)
	t.Log("component path:", component_path)

	children := contain_children[1:][:len(contain_children[1:])-6]
	t.Log("children:", children)
}

func TestCutEmptyPly(t *testing.T) {
	entire := `<ply as="path/to/component"></ply>`
	ply_part, contain_component_path, has_component := strings.Cut(entire, "\"")
	t.Log(ply_part, contain_component_path, has_component)

	if !strings.Contains(ply_part, "as") {
		t.Error(`doesn't containe path to component ("as" property not found)`)
	}

	if !has_component {
		t.Error(`has no component`)
	}

	component_path, contain_children, has_component := strings.Cut(contain_component_path, `"`)
	t.Log("component path:", component_path)

	children := contain_children[1:][:len(contain_children[1:])-6]
	t.Log("children:", children)
}

func TestCutNotPly(t *testing.T) {
	entire := `<h1>Foo page !</h1>
	<button type="input">Not a BAZ button !</button>
	<a href="/foo">Foo to Bar</a>`

	ply_part, contain_component_path, has_component := strings.Cut(entire, "\"")
	t.Log(ply_part, contain_component_path, has_component)

	if !strings.Contains(ply_part, "ply") {
		t.Log("Has no ply")
		return
	}

	if !strings.Contains(ply_part, "as") {
		t.Error(`doesn't containe path to component ("as" property not found)`)
	}

	if !has_component {
		t.Error(`has no component`)
	}

	component_path, contain_children, has_component := strings.Cut(contain_component_path, `"`)
	t.Log("component path:", component_path)

	children := contain_children[1:][:len(contain_children[1:])-6]
	t.Log("children:", children)
}
