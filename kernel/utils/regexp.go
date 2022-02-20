package utils

import (
	"reflect"
	"regexp"
	"sort"
	"strings"
)

type UtilRegexp struct {
}

func (obj UtilRegexp) ReplaceKeys(body string, sent map[string]string) string {
	//?Se crea la expresion regular para buscar todas las variables que esten dentro de [[]]
	r := regexp.MustCompile(`\[\[([^\[\[\]\]]*)\]\]`)
	matches := r.FindAllStringSubmatch(body, -1)

	var dataKeys []string

	//?v[0] = [[variable]] -> Es el identificador del mapa, aun conserva los [[]]
	//?v[1] = variable -> Es el value del mapa
	for _, v := range matches {
		body = strings.Replace(body, v[1], sent[v[1]], -1)
		dataKeys = append(dataKeys, v[1])
	}

	keys := []string{}
	for key := range sent {
		keys = append(keys, key)
	}

	//?Se organizan los arreglos de forma ascendente para comparar las keys del template y las keys enviadas a travez del endpoint
	keysAsc := keys
	sort.Strings(keysAsc[:])
	dataKeysAsc := dataKeys
	sort.Strings(dataKeysAsc[:])
	result := reflect.DeepEqual(keys, dataKeys)

	if !result {
		return ""
	}

	body = r.ReplaceAllString(body, "$1") //?$1 representa la primera subcomparación, es un valor propio de regex

	return body
}
