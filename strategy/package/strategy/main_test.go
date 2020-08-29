package strategy

import (
	"fmt"
	"os"
	"testing"
)

// TestSetLog test func
func TestSetWriter(t *testing.T) {
	printer := PrintOutPut{
		Writer:    os.Stdout,
		LogWriter: os.Stdout,
	}
	count, err := printer.Writer.Write([]byte("circle"))
	if err != nil {
		t.Error(err.Error())
	}
	if count != 6 {
		t.Error("Error in output and count!")
	}
}

func ExampleWriter() {
	printer := PrintOutPut{
		Writer:    os.Stdout,
		LogWriter: os.Stdout,
	}
	count, _ := printer.Writer.Write([]byte("circle"))
	fmt.Print(count)
	// Output:
	// circle6

}
