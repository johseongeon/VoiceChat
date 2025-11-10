package track

import (
	"sync"

	"github.com/johseongeon/WHIPTube/webrtc/peer"
	"github.com/pion/webrtc/v4"
)

// Add to list of tracks and fire renegotation for all PeerConnections.
func AddTrack(t *webrtc.TrackRemote, listLock *sync.RWMutex, trackLocals map[string]*webrtc.TrackLocalStaticRTP, peerConnections []peer.PeerConnectionState, trackNames map[string]string, streamNames map[string]string) *webrtc.TrackLocalStaticRTP { // nolint
	listLock.Lock()
	defer func() {
		listLock.Unlock()
		peer.SignalPeerConnections(listLock, trackLocals, peerConnections, trackNames, streamNames)
	}()

	// Create a new TrackLocal with the same codec as our incoming
	trackLocal, err := webrtc.NewTrackLocalStaticRTP(t.Codec().RTPCodecCapability, t.ID(), t.StreamID())
	if err != nil {
		panic(err)
	}

	trackLocals[t.ID()] = trackLocal

	return trackLocal
}

// Remove from list of tracks and fire renegotation for all PeerConnections.
func RemoveTrack(t *webrtc.TrackLocalStaticRTP, listLock *sync.RWMutex, trackLocals map[string]*webrtc.TrackLocalStaticRTP, peerConnections []peer.PeerConnectionState, trackNames map[string]string, streamNames map[string]string) {
	listLock.Lock()
	defer func() {
		listLock.Unlock()
		peer.SignalPeerConnections(listLock, trackLocals, peerConnections, trackNames, streamNames)
	}()

	delete(trackLocals, t.ID())
}
