// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs defs_freebsd.go

package ipv4

const (
	sizeofSockaddrStorage = 0x80
	sizeofSockaddrInet    = 0x10

	sizeofIPMreq         = 0x8
	sizeofIPMreqSource   = 0xc
	sizeofGroupReq       = 0x88
	sizeofGroupSourceReq = 0x108
)

type sockaddrStorage struct {
	Len         uint8
	Family      uint8
	X__ss_pad1  [6]int8
	X__ss_align int64
	X__ss_pad2  [112]int8
}

type sockaddrInet struct {
	Len    uint8
	Family uint8
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]int8
}

type ipMreq struct {
	Multiaddr [4]byte /* in_addr */
	Interface [4]byte /* in_addr */
}

type ipMreqSource struct {
	Multiaddr  [4]byte /* in_addr */
	Sourceaddr [4]byte /* in_addr */
	Interface  [4]byte /* in_addr */
}

type groupReq struct {
	Interface uint32
	Pad_cgo_0 [4]byte
	Group     sockaddrStorage
}

type groupSourceReq struct {
	Interface uint32
	Pad_cgo_0 [4]byte
	Group     sockaddrStorage
	Source    sockaddrStorage
}
