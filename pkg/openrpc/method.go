package openrpc

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/selesy/jsonschema"
)

// Method is an idempotent representation of an OpenRPC method.  Only the
// fields needed to create a "raw schema" for an MCP tool are defined and
// processed while unmarshaling the JSON.  See [the specification] for
// more details on the available fields when updating or expanding this
// type.
//
// [the specification]: https://spec.open-rpc.org/#method-object
type Method struct {
	method method
}

func (m *Method) WithDefs(defs Definitions) *Method {

	return &Method{
		method: method{
			summarizedAndDescribed: summarizedAndDescribed{
				Summary:     m.method.Summary,
				Description: m.method.Description,
			},
			Name:   m.method.Name,
			Params: m.method.Params,
			Refs:   m.method.Refs,
			Defs:   defs,
		},
	}
}

// Name returns the name of the method.
func (m *Method) Name() string {
	return m.method.Name
}

// Description returns the description of the method.
func (m *Method) Description() string {
	return m.method.description()
}

// Params returns the parameters of the method.
func (m *Method) Params() []Param {
	return m.method.Params
}

// Refs returns the references found within the method.
func (m *Method) Refs() []string {
	return m.method.Refs
}

// Defs returns the definitions used within the method.
func (m *Method) Defs() Definitions {
	return m.method.Defs
}

// MarshalJSON implements the json.Marshaler interface.
func (m Method) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.method)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (m *Method) UnmarshalJSON(data []byte) error {
	data = rewriteReferences(data)

	if err := json.Unmarshal(data, &m.method); err != nil {
		return err
	}

	refs, err := findParamReferences(data)
	if err != nil {
		return err
	}

	m.method.Refs = refs

	return err
}

type method struct {
	summarizedAndDescribed
	Name   string      `json:"name"`
	Params []Param     `json:"params"`
	Refs   []string    `json:"-"`
	Defs   Definitions `json:"defs"`
}

// Param is an idempotent representation of an OpenRPC method parameter.
// For the Ethereum RPC, this is always a ContentDescriptorObject and all
// fields are decoded.  See [the specification] for more details.
//
// [the specification]: https://spec.open-rpc.org/#content-descriptor-object
type Param struct {
	param param
}

// Name returns the name of the parameter.
func (p *Param) Name() string {
	return p.param.Name
}

// Description returns the description of the parameter.
func (p *Param) Description() string {
	return p.param.description()
}

// Schema returns the schema of the parameter.
func (p *Param) Schema() jsonschema.Schema {
	return p.param.Schema
}

// Required returns whether the parameter is required.
func (p *Param) Required() bool {
	return p.param.Required
}

// Deprecated returns whether the parameter is deprecated.
func (p *Param) Deprecated() bool {
	return p.param.Deprecated
}

// MarshalJSON implements the json.Marshaler interface.
func (p Param) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.param)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (p *Param) UnmarshalJSON(data []byte) error {
	var err error

	if err = json.Unmarshal(data, &p.param); err != nil {
		err = errors.Join(ErrUnmarshalingParams, err)
	}

	if p.param.description() == "" {
		p.param.Description = p.Name()
	}

	p.param.Name = strings.ToLower(strings.ReplaceAll(p.param.Name, " ", "_"))

	return err
}

type param struct {
	summarizedAndDescribed

	Name       string            `json:"name"`
	Schema     jsonschema.Schema `json:"schema"`
	Required   bool              `json:"required"`
	Deprecated bool              `json:"deprecated"`
}

type summarizedAndDescribed struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

func (d *summarizedAndDescribed) description() string {
	desc := d.Description
	if desc == "" {
		desc = d.Summary
	}

	return desc
}
