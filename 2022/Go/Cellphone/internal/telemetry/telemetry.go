package telemetry

import (
	"fmt"
	"io"
)

type Telemetry struct {
	out io.Writer
}

func NewTelemetry(out io.Writer) Telemetry {
	return Telemetry{out}
}

func (self *Telemetry) Info(str string) {
	self.out.Write([]byte(fmt.Sprintf("INFO %s", str)))
}
