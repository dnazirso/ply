package ply

import (
	"strings"
	"text/template"
)

func componentBuilder(tmplStr string) string {
	lastPlyIndex := strings.LastIndex(tmplStr, "<ply")
	followingEndPlyIndex := strings.Index(tmplStr[lastPlyIndex:], "</ply>")

	isolate := tmplStr[lastPlyIndex:(lastPlyIndex + followingEndPlyIndex)]
	isolateArr := strings.Split(isolate, "\"")

	componentPath := isolateArr[1]
	children := isolateArr[2][1:]
	component := Fold(componentPath, children)

	tmplStr = tmplStr[:lastPlyIndex] + component + tmplStr[lastPlyIndex+followingEndPlyIndex+6:]

	return tmplStr
}

func Fold(componentPath string, children string) string {
	tmpl, _ := template.ParseFiles(componentPath + ".html")
	tmplStr := tmpl.Tree.Root.String()

	if strings.Contains(tmplStr, "{{.Children}}") {
		tmplStr = strings.ReplaceAll(tmplStr, "{{.Children}}", children)
	}

	if strings.Contains(tmplStr, "<ply") {
		tmplStr = componentBuilder(tmplStr)
	}

	return tmplStr
}
