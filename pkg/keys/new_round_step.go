package keys

import (
	"crypto/sha256"
	"github.com/zsystm/cometbft-types/lib"
)

type NewRoundStepKey struct {
	Height   int64
	Round    int32
	Step     int32
	Sender   string
	Receiver string
}

func (nrsk *NewRoundStepKey) Hash() string {
	hash := sha256.New()
	hash.Write(lib.Int64ToBytes(nrsk.Height))
	hash.Write(lib.Int32ToBytes(nrsk.Round))
	hash.Write(lib.Int32ToBytes(nrsk.Step))
	hash.Write([]byte(nrsk.Sender))
	hash.Write([]byte(nrsk.Receiver))
	return string(hash.Sum(nil))
}
