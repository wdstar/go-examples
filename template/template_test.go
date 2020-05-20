package template

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func init() {
}

func TestMain(m *testing.M) {
	// before all
	code := m.Run()

	// after all
	os.Exit(code)
}

// test helpers for each test case.
func setUp() {
}

func tearDown() {
}

type Album struct {
	Title  string
	Labels []string
	Join   func(strs []string, sep string) string
}

func TestTemplate(t *testing.T) {
	assert := assert.New(t)

	album := Album{
		Title:  "Mozart: Requiem K. 626",
		Labels: []string{"Classical", "SACD"},
		Join:   strings.Join,
	}

	expected := "Title: Mozart: Requiem K. 626, labels: Classical,SACD"

	// 1. Calling function variables version
	tpl, err := template.New("test").
		Parse(`Title: {{.Title}}, labels: {{(call .Join .Labels ",")}}`)
	if err != nil {
		panic(err)
	}
	var output bytes.Buffer
	err = tpl.Execute(&output, album)
	// err = tpl.Execute(os.Stdout, album)
	if err != nil {
		panic(err)
	}
	assert.Equal(expected, output.String())

	// 2. Creating custom functions for template version
	tpl, err = template.New("test").Funcs(template.FuncMap{
		"join": strings.Join,
	}).Parse(`Title: {{.Title}}, labels: {{join .Labels ","}}`)
	if err != nil {
		panic(err)
	}
	output = bytes.Buffer{}
	err = tpl.Execute(&output, album)
	if err != nil {
		panic(err)
	}
	assert.Equal(expected, output.String())
}
