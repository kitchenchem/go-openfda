package fda

type DeviceService struct {
	Client *Client
}

type Device struct{}

func (d Device) String() string {
	// TODO Implement nice formatting for fda types
	return ""
}
