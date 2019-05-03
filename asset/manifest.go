package asset

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

// Manifest アセットファイルをHash値がついた値にマッピングする
type Manifest struct {
	AppJs string `json:"app.js"`
}

var manifest Manifest

func init() {
	bytes, err := ioutil.ReadFile("../public/js/manifest.json")

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(bytes, &manifest); err != nil {
		log.Fatal(err)
	}
}

// Path アセットのHash値がついた値を返す
func Path(assetName string) string {
	splitted := strings.Split(assetName, ".") // app.js => ["app", "js"]

	var camelizedKey = ""
	for _, str := range splitted {
		rune := []rune(str)
		camelizedKey += strings.ToUpper(string(rune[0])) + string(rune[1:])
	}

	// 動的に Manifest構造体から値を取り出す
	r := reflect.ValueOf(manifest)
	f := reflect.Indirect(r).FieldByName(camelizedKey)

	return f.String()
}
