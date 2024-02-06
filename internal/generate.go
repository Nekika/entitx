package internal

import (
	"bytes"
	_ "embed"
	"entitx/lib"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
)

var (
	//go:embed templates/entity.go.tmpl
	entityTemplate string
)

func Generate(entity, dest string) error {
	t, err := template.New("entity").Parse(entityTemplate)
	if err != nil {
		return err
	}

	data := make(map[string]any)
	data["Entity"] = lib.Capitalize(entity)
	data["Record"] = fmt.Sprintf("%sRecord", data["Entity"])
	data["Repository"] = fmt.Sprintf("%sRepository", data["Entity"])
	data["DefaultRepository"] = fmt.Sprintf("Default%s", data["Repository"])
	data["Service"] = fmt.Sprintf("%sService", data["Entity"])
	data["Package"] = path.Base(dest)

	buff := bytes.NewBuffer([]byte{})
	if err = t.Execute(buff, data); err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.go", strings.ToLower(entity))
	filepath := path.Join(dest, filename)
	return os.WriteFile(filepath, buff.Bytes(), 0o644)
}

func MustGenerate(entity, path string) {
	if err := Generate(entity, path); err != nil {
		panic(err)
	}
}
