package keys

import (
	"crypto/sha256"
	"github.com/zsystm/cometbft-types/lib"
)

type ProposalPOLKey struct {
	Height           int64
	ProposalPolRound int32
	Sender           string
	Receiver         string
}

func (ppk *ProposalPOLKey) Hash() string {
	hash := sha256.New()
	hash.Write(lib.Int64ToBytes(ppk.Height))
	hash.Write(lib.Int32ToBytes(ppk.ProposalPolRound))
	hash.Write([]byte(ppk.Sender))
	hash.Write([]byte(ppk.Receiver))
	return string(hash.Sum(nil))
}
