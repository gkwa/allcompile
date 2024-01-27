package sailboat

import (
	"fmt"
	"strings"
)

func (b *Sailboat) String() string {
	var builder strings.Builder

	builder.WriteString("Brand: ")
	builder.WriteString(b.Brand)

	builder.WriteString("\n")
	builder.WriteString("Color: ")
	builder.WriteString(b.Color)

	builder.WriteString("\n")
	builder.WriteString("Length: ")
	builder.WriteString(fmt.Sprintf("%d", b.Length))

	return builder.String()
}
