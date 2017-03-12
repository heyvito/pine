package pine

import (
	"fmt"
	"time"

	"github.com/logrusorgru/aurora"
)

type msgType int

type pineMsg struct {
	t      msgType
	at     time.Time
	module string
	extra  *string
	msg    string
	params []interface{}
}

type PineWriter struct {
	parent *Pine
	name   string
}

type PineExtraWriter struct {
	parent *PineWriter
	extra  string
}

type Pine struct {
	timeProvider   func() time.Time
	outputProvider func(msg string)
}

func (p *Pine) write(t msgType, module string, extra *string, msg string, params ...interface{}) {
	at := p.timeProvider()

	prefix := fmt.Sprintf("%s %s  %s ", aurora.Gray(at.Format("15:04:05")), typeEmoji[t], aurora.Magenta(module))
	ex := ""
	if extra != nil {
		ex = fmt.Sprintf("%s ", aurora.Cyan(*extra))
	}
	suffix := fmt.Sprintf(msg, params...)
	p.outputProvider(fmt.Sprintf("%s%s%s\n", prefix, ex, suffix))

}

func (p *Pine) NewWriter(module string) *PineWriter {
	return &PineWriter{
		parent: p,
		name:   module,
	}
}

func (p *PineWriter) WithExtra(extra string) *PineExtraWriter {
	return &PineExtraWriter{
		parent: p,
		extra:  extra,
	}
}

var pine = &Pine{
	timeProvider:   func() time.Time { return time.Now() },
	outputProvider: func(msg string) { fmt.Print(msg) },
}

func NewWriter(module string) *PineWriter {
	return pine.NewWriter(module)
}

//go:generate go run generators/type-generator.go
