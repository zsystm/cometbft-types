package keys

import (
	"crypto/sha256"
	"github.com/zsystm/cometbft-types/lib"
)

type BlockPartKey struct {
	Height   int64
	Round    int32
	Index    uint32
	Sender   string
	Receiver string
}

func (bpk *BlockPartKey) Hash() string {
	hash := sha256.New()
	hash.Write(lib.Int64ToBytes(bpk.Height))
	hash.Write(lib.Int32ToBytes(bpk.Round))
	hash.Write(lib.Uint32ToBytes(bpk.Index))
	hash.Write([]byte(bpk.Sender))
	hash.Write([]byte(bpk.Receiver))
	return string(hash.Sum(nil))
}
