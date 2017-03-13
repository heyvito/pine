// Pine is a completely useless (but cute) logging interface
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

// PineWriter is a writter instance already associated to a module.
type PineWriter struct {
	parent *Pine
	name   string
}

// PineExtraWriter is a writer instance associated to a module and with
// a static Extra field.
type PineExtraWriter struct {
	parent *PineWriter
	extra  string
}

// Pine is a completely useless (but cute) logging interface
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

// NewWriter creates a new writer instance with a given module name
func (p *Pine) NewWriter(module string) *PineWriter {
	return &PineWriter{
		parent: p,
		name:   module,
	}
}

// WithExtra returns a new PineExtraWriter with an associated module and
// static extra value
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

// NewWriter creates a new Writer instance using the provided module
// name
func NewWriter(module string) *PineWriter {
	return pine.NewWriter(module)
}

//go:generate go run generators/type-generator.go
