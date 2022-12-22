package translation

import (
	"encoding/json"
	"fmt"
	"ocean/utils/telemetry"
	"os"
	"reflect"
	"strings"
)

var Languages = []string{"en-us"}
var Ns = []string{"commands"}
var translations = make(map[string]map[string]map[string]any)

func Init() {
	for _, lang := range Languages {
		for _, ns := range Ns {
			var resu = map[string]any{}

			r, err := os.ReadFile(fmt.Sprintf("resources/locales/%s/%s.json", lang, ns))
			if err != nil {
				telemetry.Fatalf("Could not load language file %s/%s", lang, ns)
				return
			}

			err = json.Unmarshal(r, &resu)
			if err != nil {
				telemetry.Fatalf("Could not load language file %s/%s", lang, ns)
				return
			}

			translations[lang] = map[string]map[string]any{}
			translations[lang][ns] = resu
		}
	}

	telemetry.Info("Locales file have been registered")
}

func GetFixedT(language string) func(path string, data ...interface{}) string {
	return func(path string, data ...interface{}) string {
		// `langobj` irá retornar um objeto com todos os ns
		if langobj, ok := translations[language]; ok {
			nspath := strings.SplitN(path, ":", 2)
			ns := nspath[0]
			path = nspath[1]

			// translation um objeto com todas as traduções
			if translation, ok := langobj[ns]; ok {
				splitedPath := strings.Split(path, ".")
				var selector any = translation

                // iremos navegar por todas as keys
                // se colocarmos [a, b, c], irá navegar até chegar ao c
				for _, name := range splitedPath {
					if _selector, ok := selector.(map[string]interface{})[name]; ok {
						selector = _selector.(interface{})
					} else {
						return path
					}
				}

				if reflect.TypeOf(selector).Kind() != reflect.String {
					return path
				}

				return selector.(string)
			} else {
				return path
			}
		} else {
			return path
		}
	}
}
