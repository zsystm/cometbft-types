package lib

import (
	"encoding/hex"
	v1 "github.com/cometbft/cometbft/api/cometbft/types/v1"
	"github.com/zsystm/cometbft-types/pkg/core"
	"strings"
)

func CometVoteToVote(vote *v1.Vote) *core.Vote {
	if vote == nil {
		return nil
	}
	voteType := vote.GetType().String()
	typeStrings := strings.Split(voteType, "_")
	endIdx := len(typeStrings) - 1
	blockId := vote.GetBlockID()
	return &core.Vote{
		Type: strings.ToLower(typeStrings[endIdx]),
		BlockId: core.BlockID{
			Hash: hex.EncodeToString(blockId.Hash),
			PartSetHeader: core.PartSetHeader{
				Total: uint64(blockId.PartSetHeader.Total),
				Hash:  hex.EncodeToString(blockId.PartSetHeader.Hash),
			},
		},
		Height:           uint64(vote.Height),
		Round:            uint64(vote.Round),
		Timestamp:        vote.Timestamp,
		ValidatorIndex:   uint64(vote.ValidatorIndex),
		ValidatorAddress: hex.EncodeToString(vote.ValidatorAddress),
		Signature:        hex.EncodeToString(vote.Signature),
	}
}
