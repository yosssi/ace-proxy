package proxy

import (
	"testing"

	"github.com/yosssi/ace"
)

func TestNew(t *testing.T) {
	p := New(nil)

	if p.Opts != nil {
		t.Errorf("the options should be nil. [actual: %+v]", p.Opts)
	}
}

func TestProxy_Load_opts_not_specified(t *testing.T) {
	p := New(&ace.Options{
		BaseDir: "test",
	})

	_, err := p.Load("base", "inner", nil)

	if err != nil {
		t.Errorf("an error occurred [err: %+v]", err)
	}
}

func TestProxy_Load_opts_specified(t *testing.T) {
	p := New(&ace.Options{
		BaseDir: "test",
	})

	expectedErrMsg := "open base.ace: no such file or directory"

	_, err := p.Load("base", "inner", &ace.Options{DynamicReload: true})

	if err == nil {
		t.Error("an error should be occurred")
	}

	if err.Error() != expectedErrMsg {
		t.Errorf("the error (%s) should be occurred [actual: %+v]", expectedErrMsg, err)
	}
}
