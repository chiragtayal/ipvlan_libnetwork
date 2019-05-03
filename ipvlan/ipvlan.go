package ipvlan

import (
	"github.com/docker/go-plugins-helpers/network"
)

type IpvlanDriver struct{}

// NewIpvlanDriver() (*Driver)
func NewIpvlanDriver() (*IpvlanDriver) {
	return &IpvlanDriver{}
}

// GetCapabilities() (*CapabilitiesResponse, error)
func (d *IpvlanDriver) GetCapabilities() (*network.CapabilitiesResponse, error) {
	return &network.CapabilitiesResponse{Scope: network.LocalScope}, nil
}

// CreateNetwork(*CreateNetworkRequest) error
func (d *IpvlanDriver) CreateNetwork(r *network.CreateNetworkRequest) error {
	return nil
}

// AllocateNetwork(*AllocateNetworkRequest) (*AllocateNetworkResponse, error)
func (d *IpvlanDriver) AllocateNetwork(*network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {
	return  &network.AllocateNetworkResponse{}, nil
}

// DeleteNetwork(*DeleteNetworkRequest) error
func (d *IpvlanDriver) DeleteNetwork(*network.DeleteNetworkRequest) error {
	return nil
}

// FreeNetwork(*FreeNetworkRequest) error
func (d *IpvlanDriver) FreeNetwork(*network.FreeNetworkRequest) error {
	return nil
}

// CreateEndpoint(*CreateEndpointRequest) (*CreateEndpointResponse, error)
func (d *IpvlanDriver) CreateEndpoint(*network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	return &network.CreateEndpointResponse{}, nil
}

// DeleteEndpoint(*DeleteEndpointRequest) error
func (d *IpvlanDriver) DeleteEndpoint(*network.DeleteEndpointRequest) error {
	return nil
}

// EndpointInfo(*InfoRequest) (*InfoResponse, error)
func (d *IpvlanDriver) EndpointInfo(*network.InfoRequest) (*network.InfoResponse, error) {
	return &network.InfoResponse{}, nil
}

// Join(*JoinRequest) (*JoinResponse, error)
func (d *IpvlanDriver) Join(*network.JoinRequest) (*network.JoinResponse, error) {
	return &network.JoinResponse{}, nil
}

// Leave(*LeaveRequest) error
func (d *IpvlanDriver) Leave(*network.LeaveRequest) error {
	return nil
}

// DiscoverNew(*DiscoveryNotification) error
func (d *IpvlanDriver) DiscoverNew(*network.DiscoveryNotification) error {
	return nil
}

// DiscoverDelete(*DiscoveryNotification) error
func (d *IpvlanDriver) DiscoverDelete(*network.DiscoveryNotification) error {
	return nil
}

// ProgramExternalConnectivity(*ProgramExternalConnectivityRequest) error
func (d *IpvlanDriver) ProgramExternalConnectivity(*network.ProgramExternalConnectivityRequest) error {
	return nil
}

// RevokeExternalConnectivity(*RevokeExternalConnectivityRequest) error
func (d *IpvlanDriver) RevokeExternalConnectivity(*network.RevokeExternalConnectivityRequest) error {
	return nil
}
