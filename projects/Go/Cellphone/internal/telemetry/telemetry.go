package telemetry

import (
	"fmt"
	"io"
)

type Telemetry struct {
	out      io.Writer
	counters map[string]int
	state    map[string]interface{}
}

func NewTelemetry(out io.Writer) *Telemetry {
	return &Telemetry{out: out}
}

func (self *Telemetry) Info(str string) {
	self.out.Write([]byte(fmt.Sprintf("[ INFO ] %s\n", str)))
}
