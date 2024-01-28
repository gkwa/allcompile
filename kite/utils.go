package kite

import (
	"fmt"
	"strings"
)

func (k *Kite) String() string {
	var builder strings.Builder

	builder.WriteString("Brand: ")
	builder.WriteString(k.Brand)

	builder.WriteString("\n")
	builder.WriteString("Color: ")
	builder.WriteString(k.Color)

	builder.WriteString("\n")
	builder.WriteString("Length: ")
	builder.WriteString(fmt.Sprintf("%d", k.Length))

	return builder.String()
}
