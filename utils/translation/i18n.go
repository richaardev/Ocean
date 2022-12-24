package translation

import (
	"encoding/json"
	"fmt"
	"github.com/disgoorg/disgo/discord"
	"github.com/richaardev/Ocean/utils/telemetry"
	"os"
	"reflect"
	"strings"
)

var Languages = []string{"en-us"} // soon portuguese
var DefaultLanguage = Languages[0]
var Ns = []string{"commands"}
var translations = make(map[string]map[string]map[string]any)

func init() {
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

func Translate(query string, data ...interface{}) string {
	return GetFixedT(DefaultLanguage)(query, data)
}

type TFunction func(path string, data ...interface{}) string

func GetFixedT(language string) TFunction {
	return func(query string, data ...interface{}) string {
		// `langobj` irá retornar um objeto com todos os ns
		if langobj, ok := translations[language]; ok {
			nspath := strings.SplitN(query, ":", 2)
			ns := nspath[0]
			path := nspath[1]

			// translation um objeto com todas as traduções
			if translation, ok := langobj[ns]; ok {
				splitedPath := strings.Split(path, ".")
				var selector any = translation

				// iremos navegar por todas as keys
				// se colocarmos [a, b, c], irá navegar até chegar ao c, caso não consiga chegar ao c, irá retornar a path
				for _, name := range splitedPath {
					if _selector, ok := selector.(map[string]interface{})[name]; ok {
						selector = _selector
					} else {
						telemetry.Debugf("Could not reach %s, key %s does not exist", query, name)
						return query
					}
				}

				if reflect.TypeOf(selector).Kind() != reflect.String {
					telemetry.Debugf("%s is not a string", query)
					return query
				}

				return selector.(string)
			} else {
				telemetry.Debugf("NS %s does not exists", ns)
				return query
			}
		} else {
			telemetry.Debugf("Language %s does not exists", language)
			return query
		}
	}
}

func GetLocalizationsValues(path string) map[discord.Locale]string {
	result := map[discord.Locale]string{}

	for _, language := range Languages {
		dlang := language
		if strings.Contains(language, "-") {
			s := strings.Split(language, "-")
			language = s[0] + "-" + strings.ToUpper(s[1])
		}

		result[discord.Locale(language)] = GetFixedT(dlang)(path)
	}

	return result
}
