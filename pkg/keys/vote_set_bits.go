package keys

import (
	"crypto/sha256"
	"github.com/zsystm/cometbft-types/lib"
)

type VoteSetBitsKey struct {
	Height   int64
	Round    int32
	Type     string
	Sender   string
	Receiver string
}

func (vsbk *VoteSetBitsKey) Hash() string {
	hash := sha256.New()
	hash.Write(lib.Int64ToBytes(vsbk.Height))
	hash.Write(lib.Int32ToBytes(vsbk.Round))
	hash.Write([]byte(vsbk.Type))
	hash.Write([]byte(vsbk.Sender))
	hash.Write([]byte(vsbk.Receiver))
	return string(hash.Sum(nil))
}
