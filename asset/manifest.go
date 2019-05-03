package asset

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

// Manifest ...
type Manifest struct {
	JSONPath string
	ManifestFile
}

// ManifestFile ...
type ManifestFile struct {
	AppJs string `json:"app.js"`
}

// NewManifest return Manifest instance
func NewManifest(jsonPath string) *Manifest {
	m := new(Manifest)
	m.JSONPath = jsonPath

	bytes, err := ioutil.ReadFile(m.JSONPath)

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(bytes, &m.ManifestFile); err != nil {
		log.Fatal(err)
	}

	return m
}

// FileName アセットのHash値がついた値を返す
func (m *Manifest) FileName(assetName string) string {
	splitted := strings.Split(assetName, ".") // app.js => ["app", "js"]

	var camelizedKey = ""
	for _, str := range splitted {
		rune := []rune(str)
		camelizedKey += strings.ToUpper(string(rune[0])) + string(rune[1:])
	}

	// 動的に Manifest構造体から値を取り出す
	r := reflect.ValueOf(m.ManifestFile)
	f := reflect.Indirect(r).FieldByName(camelizedKey)

	return f.String()
}
