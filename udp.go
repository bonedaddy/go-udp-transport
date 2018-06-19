package udp

import (
	"context"

	logging "github.com/ipfs/go-log"
	peer "github.com/libp2p/go-libp2p-peer"
	tpt "github.com/libp2p/go-libp2p-transport"
	tptu "github.com/libp2p/go-libp2p-transport-upgrader"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr-net"
	mafmt "github.com/whyrusleeping/mafmt"
)

var log = logging.Logger("tcp-udp")

// UdpTransport is the TCP transport.
type UdpTransport struct {
	// Connection upgrader for upgrading insecure stream connections to
	// secure multiplex connections.
	Upgrader *tptu.Upgrader
}

var _ tpt.PacketTransport = &UdpTransport{}

// NewUdpTransport creates a tcp transport object that tracks dialers and listeners
// created. It represents an entire tcp stack (though it might not necessarily be)
func NewUdpTransport(upgrader *tptu.Upgrader) *UdpTransport {
	return &UdpTransport{Upgrader: upgrader}
}

// CanDial returns true if this transport believes it can dial the given
// multiaddr.
func (t *UdpTransport) CanDial(addr ma.Multiaddr) bool {
	return mafmt.UDP.Matches(addr)
}

func (t *UdpTransport) maDial(ctx context.Context, raddr ma.Multiaddr) (manet.Conn, error) {
	var d manet.Dialer
	return d.DialContext(ctx, raddr)
}

// Dial dials the peer at the remote address.
func (t *UdpTransport) Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (tpt.Conn, error) {
	conn, err := t.maDial(ctx, raddr)
	if err != nil {
		return nil, err
	}
	return t.Upgrader.UpgradeOutbound(ctx, t, conn, p)
}

func (t *UdpTransport) maListen(laddr ma.Multiaddr) (manet.Listener, error) {
	return manet.ListenPacket(laddr)
}

// Listen listens on the given multiaddr.
func (t *UdpTransport) Listen(laddr ma.Multiaddr) (tpt.Listener, error) {
	list, err := t.maListen(laddr)
	if err != nil {
		return nil, err
	}
	return t.Upgrader.UpgradeListener(t, list), nil
}

// Protocols returns the list of terminal protocols this transport can dial.
func (t *UdpTransport) Protocols() []int {
	return []int{ma.P_UDP}
}

// Proxy always returns false for the TCP transport.
func (t *UdpTransport) Proxy() bool {
	return false
}

func (t *UdpTransport) String() string {
	return "UDP"
}
