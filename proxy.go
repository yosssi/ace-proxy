package proxy

import (
	"html/template"

	"github.com/yosssi/ace"
)

// Proxy represents a proxy for the Ace template engine.
type Proxy struct {
	Opts *ace.Options
}

// Load calls the `Load` function of the Ace template engine.
func (p *Proxy) Load(basePath, innerPath string, opts *ace.Options) (<-chan *template.Template, <-chan error) {
	var o *ace.Options

	if opts == nil {
		o = p.Opts
	} else {
		o = opts
	}

	tplc := make(chan *template.Template)
	errc := make(chan error)

	go func() {
		tpl, err := ace.Load(basePath, innerPath, o)
		if err != nil {
			errc <- err
			return
		}
		tplc <- tpl
	}()

	return tplc, errc
}

// New creates and returns a proxy.
func New(opts *ace.Options) *Proxy {
	return &Proxy{
		Opts: opts,
	}
}
