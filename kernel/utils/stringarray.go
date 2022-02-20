package utils

import (
	"fmt"
	"strings"
)

type UtilStringArray struct {
	Value []string
}

func (obj *UtilStringArray) AppendQuoted(item interface{}) {
	obj.Value = append(obj.Value, fmt.Sprintf("'%v'", item))
}

func (obj *UtilStringArray) Append(item interface{}) {
	obj.Value = append(obj.Value, fmt.Sprintf("%v", item))
}

func (obj *UtilStringArray) ToString(separator string) string {
	return strings.Join(obj.Value, separator)
}

func (obj UtilStringArray) ToInPattern(key string) string {
	filters := []map[string]interface{}{
		{
			"key": key,
			"val": strings.Join(obj.Value, ","),
		},
	}

	return UtilsMap{}.MapToString(filters)
}
