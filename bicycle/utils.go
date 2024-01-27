package bicycle

import (
	"fmt"
	"strings"
)

func (b *Bicycle) String() string {
	var builder strings.Builder

	builder.WriteString("Brand: ")
	builder.WriteString(b.Brand)

	builder.WriteString("\n")
	builder.WriteString("Color: ")
	builder.WriteString(b.Color)

	builder.WriteString("\n")
	builder.WriteString("Weight: ")
	builder.WriteString(fmt.Sprintf("%d", b.Weight))

	return builder.String()
}
