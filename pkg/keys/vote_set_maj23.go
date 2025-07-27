package keys

import (
	"crypto/sha256"
	"github.com/zsystm/cometbft-types/lib"
)

type VoteSetMaj23Key struct {
	Height   int64
	Round    int32
	Type     string
	Sender   string
	Receiver string
}

func (vsmk *VoteSetMaj23Key) Hash() string {
	hash := sha256.New()
	hash.Write(lib.Int64ToBytes(vsmk.Height))
	hash.Write(lib.Int32ToBytes(vsmk.Round))
	hash.Write([]byte(vsmk.Type))
	hash.Write([]byte(vsmk.Sender))
	hash.Write([]byte(vsmk.Receiver))
	return string(hash.Sum(nil))
}
