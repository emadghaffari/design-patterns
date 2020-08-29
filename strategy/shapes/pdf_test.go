package shapes

import (
	"os"
	"testing"
)

func TestPDFPrint(t *testing.T) {
	sq := PdfSquare{}
	sq.SetWriter(os.Stdout)
	if err := sq.Print(); err != nil {
		t.Error(err.Error())
	}
}
