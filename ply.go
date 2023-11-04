package ply

import (
	"regexp"
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

func fold(componentPath string, children string) string {
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

func minify(tmplStr string) string {
	regex := regexp.MustCompile(`\n\s*`)

	tmplStr = regex.ReplaceAllString(tmplStr, "")

	// (<script.*)(\s*\S*)*(</script>)

	return tmplStr
}

func Fold(componentPath string, children string) string {
	tmplStr := fold(componentPath, children)

	for strings.Contains(tmplStr, "<ply") {
		tmplStr = componentBuilder(tmplStr)
	}

	tmplStr = minify(tmplStr)
	return tmplStr
}
