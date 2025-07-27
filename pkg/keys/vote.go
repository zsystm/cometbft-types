package keys

import (
	"crypto/sha256"
	"strconv"
)

type VoteKey struct {
	Height   uint64
	Round    uint64
	ValIdx   uint64
	Sender   string
	Receiver string
}

func (vk *VoteKey) Hash() string {
	hash := sha256.New()
	hash.Write([]byte(
		strconv.FormatUint(vk.Height, 10),
	))
	hash.Write([]byte(
		strconv.FormatUint(vk.Round, 10),
	))
	hash.Write([]byte(
		strconv.FormatUint(vk.ValIdx, 10),
	))
	hash.Write([]byte(vk.Sender))
	hash.Write([]byte(vk.Receiver))
	return string(hash.Sum(nil))
}
