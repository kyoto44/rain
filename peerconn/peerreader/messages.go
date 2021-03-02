package peerreader

import (
	"github.com/kyoto44/rain/bufferpool"
	"github.com/kyoto44/rain/peerprotocol"
)

// Piece message that is read from peers.
// Data of the piece is wrapped with a bufferpool.Buffer object.
type Piece struct {
	peerprotocol.PieceMessage
	Buffer bufferpool.Buffer
}
