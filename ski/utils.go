package ski

import (
	"fmt"
	"strings"
)

func (s *Ski) String() string {
	var builder strings.Builder

	builder.WriteString("Brand: ")
	builder.WriteString(s.Brand)

	builder.WriteString("\n")
	builder.WriteString("Color: ")
	builder.WriteString(s.Color)

	builder.WriteString("\n")
	builder.WriteString("Weight: ")
	builder.WriteString(fmt.Sprintf("%d", s.Weight))

	return builder.String()
}
