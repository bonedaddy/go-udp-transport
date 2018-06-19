package udp

import (
	"testing"

	insecure "github.com/libp2p/go-conn-security/insecure"
	tptu "github.com/libp2p/go-libp2p-transport-upgrader"
	utils "github.com/libp2p/go-libp2p-transport/test"
	mplex "github.com/whyrusleeping/go-smux-multiplex"
)

func TestUdpTransport(t *testing.T) {
	for i := 0; i < 2; i++ {
		ta := NewUdpTransport(&tptu.Upgrader{
			Secure: insecure.New("peerA"),
			Muxer:  new(mplex.Transport),
		})
		tb := NewUdpTransport(&tptu.Upgrader{
			Secure: insecure.New("peerB"),
			Muxer:  new(mplex.Transport),
		})

		zero := "/ip4/127.0.0.1/udp/10"
		utils.SubtestTransport(t, ta, tb, zero, "peerA")
	}
}
