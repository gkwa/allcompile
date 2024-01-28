package golf

import (
	"fmt"
	"strings"
)

func (c *GolfClub) String() string {
	var builder strings.Builder

	builder.WriteString("Brand: ")
	builder.WriteString(c.Brand)

	builder.WriteString("\n")
	builder.WriteString("Color: ")
	builder.WriteString(c.Color)

	builder.WriteString("\n")
	builder.WriteString("Weight: ")
	builder.WriteString(fmt.Sprintf("%d", c.Weight))

	return builder.String()
}
