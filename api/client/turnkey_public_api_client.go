// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/tkhq/tkcli/api/client/activities"
	"github.com/tkhq/tkcli/api/client/organizations"
	"github.com/tkhq/tkcli/api/client/policies"
	"github.com/tkhq/tkcli/api/client/private_keys"
	"github.com/tkhq/tkcli/api/client/users"
)

// Default turnkey public API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "coordinator-beta.turnkey.io"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new turnkey public API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *TurnkeyPublicAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new turnkey public API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *TurnkeyPublicAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new turnkey public API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *TurnkeyPublicAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(TurnkeyPublicAPI)
	cli.Transport = transport
	cli.Activities = activities.New(transport, formats)
	cli.Organizations = organizations.New(transport, formats)
	cli.Policies = policies.New(transport, formats)
	cli.PrivateKeys = private_keys.New(transport, formats)
	cli.Users = users.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// TurnkeyPublicAPI is a client for turnkey public API
type TurnkeyPublicAPI struct {
	Activities activities.ClientService

	Organizations organizations.ClientService

	Policies policies.ClientService

	PrivateKeys private_keys.ClientService

	Users users.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *TurnkeyPublicAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Activities.SetTransport(transport)
	c.Organizations.SetTransport(transport)
	c.Policies.SetTransport(transport)
	c.PrivateKeys.SetTransport(transport)
	c.Users.SetTransport(transport)
}
