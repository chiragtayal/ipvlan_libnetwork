package ipvlan

import (
	"github.com/docker/go-plugins-helpers/network"
)

type Driver struct{}

// NewDriver() (*Driver)
func NewDriver() (*Driver) {
	return &Driver{}
}

// GetCapabilities() (*CapabilitiesResponse, error)
func (d *Driver) GetCapabilities() (*network.CapabilitiesResponse, error) {
	return &network.CapabilitiesResponse{Scope: network.LocalScope}, nil
}

// CreateNetwork(*CreateNetworkRequest) error
func (d *Driver) CreateNetwork(r *network.CreateNetworkRequest) error {
	return nil
}

// AllocateNetwork(*AllocateNetworkRequest) (*AllocateNetworkResponse, error)
func (d *Driver) AllocateNetwork(*network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {
	return  &network.AllocateNetworkResponse{}, nil
}

// DeleteNetwork(*DeleteNetworkRequest) error
func (d *Driver) DeleteNetwork(*network.DeleteNetworkRequest) error {
	return nil
}

// FreeNetwork(*FreeNetworkRequest) error
func (d *Driver) FreeNetwork(*network.FreeNetworkRequest) error {
	return nil
}

// CreateEndpoint(*CreateEndpointRequest) (*CreateEndpointResponse, error)
func (d *Driver) CreateEndpoint(*network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	return &network.CreateEndpointResponse{}, nil
}

// DeleteEndpoint(*DeleteEndpointRequest) error
func (d *Driver) DeleteEndpoint(*network.DeleteEndpointRequest) error {
	return nil
}

// EndpointInfo(*InfoRequest) (*InfoResponse, error)
func (d *Driver) EndpointInfo(*network.InfoRequest) (*network.InfoResponse, error) {
	return &network.InfoResponse{}, nil
}

// Join(*JoinRequest) (*JoinResponse, error)
func (d *Driver) Join(*network.JoinRequest) (*network.JoinResponse, error) {
	return &network.JoinResponse{}, nil
}

// Leave(*LeaveRequest) error
func (d *Driver) Leave(*network.LeaveRequest) error {
	return nil
}

// DiscoverNew(*DiscoveryNotification) error
func (d *Driver) DiscoverNew(*network.DiscoveryNotification) error {
	return nil
}

// DiscoverDelete(*DiscoveryNotification) error
func (d *Driver) DiscoverDelete(*network.DiscoveryNotification) error {
	return nil
}

// ProgramExternalConnectivity(*ProgramExternalConnectivityRequest) error
func (d *Driver) ProgramExternalConnectivity(*network.ProgramExternalConnectivityRequest) error {
	return nil
}

// RevokeExternalConnectivity(*RevokeExternalConnectivityRequest) error
func (d *Driver) RevokeExternalConnectivity(*network.RevokeExternalConnectivityRequest) error {
	return nil
}
