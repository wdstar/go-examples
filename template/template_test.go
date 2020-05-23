package template

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// before all
	code := m.Run()

	// after all
	os.Exit(code)
}

type Album struct {
	Title  string
	Labels []string
	Join   func(strs []string, sep string) string
}

func TestTemplate(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		tpl *template.Template
	}{
		// 1. Calling function variables version
		{tpl: template.Must(
			template.New("test").
				Parse(`Title: {{.Title}}, labels: {{(call .Join .Labels ",")}}`))},
		// 2. Creating custom functions for template version
		{tpl: template.Must(
			template.New("test").Funcs(template.FuncMap{
				"join": strings.Join,
			}).Parse(`Title: {{.Title}}, labels: {{join .Labels ","}}`))},
	}

	album := Album{
		Title:  "Mozart: Requiem K. 626",
		Labels: []string{"Classical", "SACD"},
		Join:   strings.Join,
	}

	expected := "Title: Mozart: Requiem K. 626, labels: Classical,SACD"

	for _, c := range cases {
		var output bytes.Buffer
		err := c.tpl.Execute(&output, album)
		// err := c.tpl.Execute(os.Stdout, album)
		if err != nil {
			panic(err)
		}
		assert.Equal(expected, output.String())
	}
}
