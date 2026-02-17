package utils

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

func formatString(params []any) string {
	var buff bytes.Buffer

	for i := range params {
		buff.WriteString("%v")
		if i+1 != len(params) {
			buff.WriteString("\t")
		}
	}

	buff.WriteString("\n")

	return buff.String()
}

func AddItem(writter *tabwriter.Writer, params []any) {
	fmt.Fprintf(writter, formatString(params), params...)
}
