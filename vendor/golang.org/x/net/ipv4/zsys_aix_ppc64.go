// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs defs_aix.go

// Added for go1.11 compatibility
//go:build aix
// +build aix

package ipv4

const (
	sizeofIPMreq = 0x8
)

type ipMreq struct {
	Multiaddr [4]byte /* in_addr */
	Interface [4]byte /* in_addr */
}
