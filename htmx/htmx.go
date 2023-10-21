package htmx

type Htmx struct {
	Route    string
	Method   string
	Target   string
	Interval string
	Swap     string
}

type HTMX interface {
	SetRoute(route string)
	GetRoute() string

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

func (h *Htmx) SetRoute(route string) {
	h.Route = route
}

func (h *Htmx) GetRoute() string {
	return h.Route
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
