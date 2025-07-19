package lib

import (
	"encoding/hex"
	cryptov1 "github.com/cometbft/cometbft/api/cometbft/crypto/v1"
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
			Hash:          hex.EncodeToString(blockId.Hash),
			PartSetHeader: CometPartSetHeaderToPartSetHeader(&blockId.PartSetHeader),
		},
		Height:           uint64(vote.Height),
		Round:            uint64(vote.Round),
		Timestamp:        vote.Timestamp,
		ValidatorIndex:   uint64(vote.ValidatorIndex),
		ValidatorAddress: hex.EncodeToString(vote.ValidatorAddress),
		Signature:        hex.EncodeToString(vote.Signature),
	}
}

func StepIntToString(step uint32) string {
	switch step {
	case 1:
		return "newHeight"
	case 2:
		return "newRound"
	case 3:
		return "propose"
	case 4:
		return "prevote"
	case 5:
		return "prevoteWait"
	case 6:
		return "precommit"
	case 7:
		return "precommitWait"
	case 8:
		return "commit"
	default:
		return "unknown"
	}
}

func CometPartSetHeaderToPartSetHeader(psh *v1.PartSetHeader) core.PartSetHeader {
	if psh == nil {
		return core.PartSetHeader{}
	}
	return core.PartSetHeader{
		Total: uint64(psh.Total),
		Hash:  hex.EncodeToString(psh.Hash),
	}
}

func CometSignedMsgTypeToString(typ int32) string {
	switch typ {
	case 1:
		return "prevote"
	case 2:
		return "precommit"
	case 32:
		return "proposal"
	default:
		return "unknown"
	}
}

func CometBlockIDToBlockID(blockID *v1.BlockID) core.BlockID {
	if blockID == nil {
		return core.BlockID{}
	}
	return core.BlockID{
		Hash: hex.EncodeToString(blockID.Hash),
		PartSetHeader: core.PartSetHeader{
			Total: uint64(blockID.PartSetHeader.Total),
			Hash:  hex.EncodeToString(blockID.PartSetHeader.Hash),
		},
	}
}

func CometPartToPart(part *v1.Part) core.Part {
	if part == nil {
		return core.Part{}
	}
	return core.Part{
		Index: part.Index,
		Bytes: hex.EncodeToString(part.Bytes),
	}
}

func CometProofToProof(proof *cryptov1.Proof) core.Proof {
	if proof == nil {
		return core.Proof{}
	}

	return core.Proof{
		Total:    proof.Total,
		Index:    proof.Index,
		LeafHash: hex.EncodeToString(proof.LeafHash),
		Aunts:    BytesArrayToHexStringArr(proof.Aunts),
	}
}

func BytesArrayToHexStringArr(arr [][]byte) []string {
	if arr == nil {
		return nil
	}
	hexStrings := make([]string, len(arr))
	for i, b := range arr {
		hexStrings[i] = hex.EncodeToString(b)
	}
	return hexStrings
}
