package strategy

import "io"

// PrintStrategy interface
type PrintStrategy interface {
	Print() error
	SetLog(io.Writer)
	SetWriter(io.Writer)
}

// PrintOutPut struct
type PrintOutPut struct {
	Writer    io.Writer
	LogWriter io.Writer
}

// SetLog method
func (d *PrintOutPut) SetLog(w io.Writer) {
	d.LogWriter = w
}

// SetWriter method
func (d *PrintOutPut) SetWriter(w io.Writer) {
	d.Writer = w
}
