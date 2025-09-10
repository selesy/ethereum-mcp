package proxy

// ParamsSpec contains the "specification" needed to service each JSON-RPC
// method.
type ParamsSpec struct {
	order    []string
	required map[string]struct{}
}

// Order returns an ordered list of parameter names which can be used to
// convert a JSON-RPC "params" value from named parameters to positional
// parameters.
func (p *ParamsSpec) Order() []string {
	return p.order
}

// Required returns a set of which parameters are required.  For positional
// parameters, the parameters in this set should be those that are first
// in the list returned by Order.
func (p *ParamsSpec) Required() map[string]struct{} {
	return p.required
}
