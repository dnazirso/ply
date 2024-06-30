package ply

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"text/template"
)

func componentBuilder(tmplStr string) (string, error) {
	lastPlyIndex := strings.LastIndex(tmplStr, "<ply")
	followingEndPlyIndex := strings.Index(tmplStr[lastPlyIndex:], "</ply>")

	isolate := tmplStr[lastPlyIndex:(lastPlyIndex + followingEndPlyIndex + 6)]
	ply_part, contain_component_path, has_component := strings.Cut(isolate, "\"")

	if !strings.Contains(ply_part, "ply") {
		return isolate, nil
	}

	if !strings.Contains(ply_part, "src") {
		return "", fmt.Errorf(`doesn't contain path to component ("src" property not found)`)
	}

	if !has_component {
		return "", fmt.Errorf(`has no component`)
	}

	component_path, contain_children, has_component := strings.Cut(contain_component_path, `"`)

	children := contain_children[1:][:len(contain_children[1:])-6]

	component := fold(component_path, children)

	tmplStr = tmplStr[:lastPlyIndex] + component + tmplStr[lastPlyIndex+followingEndPlyIndex+6:]

	return tmplStr, nil
}

func fold(componentPath string, children string) string {
	tmpl, _ := template.ParseFiles(componentPath + ".html")
	tmplStr := tmpl.Tree.Root.String()

	if strings.Contains(tmplStr, "{{.Children}}") {
		tmplStr = strings.ReplaceAll(tmplStr, "{{.Children}}", children)
	}
	var err error
	if strings.Contains(tmplStr, "<ply") {
		tmplStr, err = componentBuilder(tmplStr)
	}

	if err != nil {
		log.Panic(err)
	}

	return tmplStr
}

// func isolateScripts(tmplStr string) (string, []string) {
// 	isolateArr := []string{}
//
// 	for strings.Contains(tmplStr, "<script") {
//
// 		lastScriptIndex := strings.LastIndex(tmplStr, "<script")
// 		endScriptIndex := strings.Index(tmplStr[lastScriptIndex:], "</script>")
//
// 		isolate := tmplStr[lastScriptIndex:(lastScriptIndex + endScriptIndex + 9)]
// 		isolateArr = append(isolateArr, isolate)
//
// 		tmplStr = tmplStr[:lastScriptIndex] + tmplStr[lastScriptIndex+endScriptIndex+9:]
// 	}
//
// 	return tmplStr, isolateArr
// }

func replaceBlanks(tmplStr string) string {
	regex := regexp.MustCompile(`\n\s*`)

	tmplStr = regex.ReplaceAllString(tmplStr, " ")

	return tmplStr
}

func Fold(componentPath string, children string) string {
	tmplStr := fold(componentPath, children)

	var err error
	for strings.Contains(tmplStr, "<ply") {
		tmplStr, err = componentBuilder(tmplStr)
	}

	if err != nil {
		log.Panic(err)
	}

	// tmplStr, scripts := isolateScripts(tmplStr)

	tmplStr = replaceBlanks(tmplStr)

	// endBodyIndex := strings.Index(tmplStr, "</body>")

	// scriptPack := strings.Join(scripts, "")

	// tmplStr = tmplStr[:endBodyIndex] + "\n" + scriptPack + "\n" + tmplStr[endBodyIndex:]

	return tmplStr
}
