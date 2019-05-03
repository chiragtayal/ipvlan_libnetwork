package ipvlan

import (
	"github.com/docker/go-plugins-helpers/ipam"
)

type IpamDriver struct{}

// NewIpamDriver() (*Driver)
func NewIpamDriver() (*IpamDriver) {
	return &IpamDriver{}
}

// GetCapabilities() (*CapabilitiesResponse, error)
func (d *IpamDriver) GetCapabilities() (*ipam.CapabilitiesResponse, error) {
	return &ipam.CapabilitiesResponse{}, nil
}

// GetDefaultAddressSpaces() (*AddressSpacesResponse, error)
func (d *IpamDriver) GetDefaultAddressSpaces() (*ipam.AddressSpacesResponse, error) {
	return &ipam.AddressSpacesResponse{}, nil
}

// RequestPool(*RequestPoolRequest) (*RequestPoolResponse, error)
func (d *IpamDriver) RequestPool(*ipam.RequestPoolRequest) (*ipam.RequestPoolResponse, error) {
	return &ipam.RequestPoolResponse{}, nil
}

// ReleasePool(*ReleasePoolRequest) error
func (d *IpamDriver) ReleasePool(*ipam.ReleasePoolRequest) error {
	return nil
}

// RequestAddress(*RequestAddressRequest) (*RequestAddressResponse, error)
func (d *IpamDriver) RequestAddress(*ipam.RequestAddressRequest) (*ipam.RequestAddressResponse, error) {
	return &ipam.RequestAddressResponse{}, nil
}

// ReleaseAddress(*ReleaseAddressRequest) error
func (d *IpamDriver) ReleaseAddress(*ipam.ReleaseAddressRequest) error {
	return nil
}

