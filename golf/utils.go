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
	builder.WriteString("Type: ")
	builder.WriteString(c.Type)

	builder.WriteString("\n")
	builder.WriteString("Length: ")
	builder.WriteString(fmt.Sprintf("%d", c.Length))

	return builder.String()
}
