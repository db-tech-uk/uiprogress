package uiprogress

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/db-tech-uk/csync/waitgroup"
)

func TestStoppingPrintout(t *testing.T) {
	p := New()
	p.SetRefreshInterval(time.Millisecond * 10)

	var buffer = &bytes.Buffer{}
	p.SetOut(buffer)

	bar := p.AddBar(100)
	p.Start()

	wg := waitgroup.New()

	wg.Add(1)

	go func() {
		for i := 0; i <= 80; i = i + 10 {
			if err := bar.Set(i); err != nil {
				t.Error(err)
			}
			time.Sleep(time.Millisecond * 5)
		}

		wg.Done()
	}()

	wg.Wait()

	p.Stop()
	fmt.Fprintf(buffer, "foo")

	var wantSuffix = "[==============================>-------]\nfoo"

	if !strings.HasSuffix(buffer.String(), wantSuffix) {
		t.Errorf("Content that should be printed after stop not appearing on buffer. [%v]", buffer.String())
	}
}
