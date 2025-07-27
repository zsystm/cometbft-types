package keys

import (
	"crypto/sha256"
	"github.com/zsystm/cometbft-types/lib"
)

type HasVoteKey struct {
	Height   int64
	Round    int32
	Type     string
	Index    int32
	Sender   string
	Receiver string
}

func (hvk *HasVoteKey) Hash() string {
	hash := sha256.New()
	hash.Write(lib.Int64ToBytes(hvk.Height))
	hash.Write(lib.Int32ToBytes(hvk.Round))
	hash.Write([]byte(hvk.Type))
	hash.Write(lib.Int32ToBytes(hvk.Index))
	hash.Write([]byte(hvk.Sender))
	hash.Write([]byte(hvk.Receiver))
	return string(hash.Sum(nil))
}
