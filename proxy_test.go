package proxy

import (
	"testing"

	"github.com/yosssi/ace"
)

func TestNew(t *testing.T) {
	p := New(nil)

	if p.opts != nil {
		t.Errorf("the options should be nil. [actual: %+v]", p.opts)
	}
}

func TestProxy_Load_opts_not_specified(t *testing.T) {
	p := New(&ace.Options{
		BaseDir: "test",
	})

	if _, err := p.Load("base", "inner", nil); err != nil {
		t.Errorf("an error occurred [err: %+v]", err)
	}
}

func TestProxy_Load_opts_specified(t *testing.T) {
	p := New(&ace.Options{
		BaseDir: "test",
	})

	expectedErrMsg := "open base.ace: no such file or directory"

	if _, err := p.Load("base", "inner", &ace.Options{DynamicReload: true}); err == nil || err.Error() != expectedErrMsg {
		t.Errorf("the error (%s) should be occurred [actual: %+v]", expectedErrMsg, err)
	}
}
