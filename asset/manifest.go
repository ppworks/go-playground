package asset

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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
	AppJs  string `json:"app.js"`
	AppCss string `json:"app.css"`
}

// NewManifest return Manifest instance
func NewManifest(jsonPath string) *Manifest {
	m := new(Manifest)
	if os.Getenv("APP_ENV") != "production" {
		return m
	}

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

// Path アセットのHash値がついた値を返す
func (m *Manifest) Path(assetName string) string {
	if os.Getenv("APP_ENV") != "production" {
		return "https://lvh.me:8080/assets/" + assetName
	}

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
