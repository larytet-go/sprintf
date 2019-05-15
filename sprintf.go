package sprintf

import (
	"bytes"
	"fmt"
	"reflect"
)
// Reflection based loop which prints all fields in a structure
// This function is slow.
// Reflection based loop which prints all fields in a structure
// This function is slow.
func SprintfStructure(data interface{}, columns int, format string, skip []string) string {
	skipSet := make(map[string]bool)
	for _, v := range skip {
		skipSet[v] = true
	}

	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	if format == "" {
		format = "%-24s %9v "
	}
	var buffer bytes.Buffer
	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		if _, ok := skipSet[fieldName]; ok {
			continue
		}
		buffer.WriteString(fmt.Sprintf(format, fieldName, v.Field(i)))
		if (i % columns) == (columns - 1) {
			buffer.WriteString("\n")
		}
	}

	return buffer.String()
}
