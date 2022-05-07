// This file is generated by "./lib/proto/generate"

package proto

/*

DeviceOrientation

*/

// DeviceOrientationClearDeviceOrientationOverride Clears the overridden Device Orientation.
type DeviceOrientationClearDeviceOrientationOverride struct {
}

// ProtoReq name
func (m DeviceOrientationClearDeviceOrientationOverride) ProtoReq() string {
	return "DeviceOrientation.clearDeviceOrientationOverride"
}

// Call sends the request
func (m DeviceOrientationClearDeviceOrientationOverride) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// DeviceOrientationSetDeviceOrientationOverride Overrides the Device Orientation.
type DeviceOrientationSetDeviceOrientationOverride struct {

	// Alpha Mock alpha
	Alpha float64 `json:"alpha"`

	// Beta Mock beta
	Beta float64 `json:"beta"`

	// Gamma Mock gamma
	Gamma float64 `json:"gamma"`
}

// ProtoReq name
func (m DeviceOrientationSetDeviceOrientationOverride) ProtoReq() string {
	return "DeviceOrientation.setDeviceOrientationOverride"
}

// Call sends the request
func (m DeviceOrientationSetDeviceOrientationOverride) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}
