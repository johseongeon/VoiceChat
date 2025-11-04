package peer

import (
	"github.com/johseongeon/WHIPTube/ws"
	"github.com/pion/webrtc/v4"
)

type PeerConnectionState struct {
	PeerConnection *webrtc.PeerConnection
	Websocket      *ws.ThreadSafeWriter
}
