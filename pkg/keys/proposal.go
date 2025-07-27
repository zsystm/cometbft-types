package keys

import (
	"crypto/sha256"
	"github.com/zsystm/cometbft-types/lib"
)

type ProposalKey struct {
	Height   int64
	Round    int32
	Sender   string
	Receiver string
}

func (pk *ProposalKey) Hash() string {
	hash := sha256.New()
	hash.Write(lib.Int64ToBytes(pk.Height))
	hash.Write(lib.Int32ToBytes(pk.Round))
	hash.Write([]byte(pk.Sender))
	hash.Write([]byte(pk.Receiver))
	return string(hash.Sum(nil))
}
