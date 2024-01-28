package hockey

import (
	"fmt"
	"strings"
)

func (h *HockeyStick) String() string {
	var builder strings.Builder

	builder.WriteString("Brand: ")
	builder.WriteString(h.Brand)

	builder.WriteString("\n")
	builder.WriteString("Length: ")
	builder.WriteString(fmt.Sprintf("%d", h.Length))

	builder.WriteString("\n")
	builder.WriteString("Flex: ")
	builder.WriteString(h.Color)

	return builder.String()
}
