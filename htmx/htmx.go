package htmx

import "net/url"

type Htmx struct {
	Route    url.URL
	Method   string
	Target   string
	Interval string
	Swap     string
}

type HTMX interface {
	AddBeforePath(path string) error
	AppendToPath(path string)

	SetMethod(method string)
	GetMethod() string

	SetTarget(target string)
	GetTarget() string

	SetInterval(interval string)
	GetInterval() string

	SetSwap(swap string)
	GetSwap() string

	GetHtmx() *Htmx
}

func NewEmpty() HTMX {
	var h Htmx
	return &h
}

func (h *Htmx) GetHtmx() *Htmx {
	return h
}

func (h *Htmx) AddBeforePath(path string) error {
	path, err := url.JoinPath(path, h.Route.Path)
	if err != nil {
		return err
	}

	h.Route.Path = path

	return nil
}

func (h *Htmx) AppendToPath(path string) {
	h.Route = *h.Route.JoinPath(path)
}

func (h *Htmx) GetUrl() string {
	return h.Route.String()
}

func (h *Htmx) SetMethod(method string) {
	h.Method = method
}

func (h *Htmx) GetMethod() string {
	return h.Method
}

func (h *Htmx) SetTarget(target string) {
	h.Target = target
}

func (h *Htmx) GetTarget() string {
	return h.Target
}

func (h *Htmx) SetInterval(interval string) {
	h.Interval = interval
}

func (h *Htmx) GetInterval() string {
	return h.Interval
}

func (h *Htmx) SetSwap(swap string) {
	h.Swap = swap
}

func (h *Htmx) GetSwap() string {
	return h.Swap
}
