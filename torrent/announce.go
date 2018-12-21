package torrent

import (
	"math"

	"github.com/cenkalti/rain/torrent/internal/tracker"
)

func (t *Torrent) announcerFields() tracker.Torrent {
	tr := tracker.Torrent{
		InfoHash:        t.infoHash,
		PeerID:          t.peerID,
		Port:            t.port,
		BytesDownloaded: t.byteStats.BytesDownloaded,
		BytesUploaded:   t.byteStats.BytesUploaded,
	}
	if t.bitfield == nil {
		tr.BytesLeft = math.MaxUint32
	} else {
		tr.BytesLeft = t.info.TotalLength - t.bytesComplete()
	}
	return tr
}
