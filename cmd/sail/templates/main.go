package templates

import (
	"bytes"
	"text/template"
)

type ComponentData struct {
	Name      string // e.g. "App"
	LowerName string // e.g. "app" (you can compute this)
}

func RenderTemplate(tmplStr string, data ComponentData) (string, error) {
	tmpl, err := template.New("component").Parse(tmplStr)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
