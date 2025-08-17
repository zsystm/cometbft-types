package core

import (
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Header     Header       `bson:"header" json:"header"`
	Data       Data         `bson:"data" json:"data"`
	Evidence   EvidenceData `bson:"evidence" json:"evidence"`
	LastCommit *Commit      `bson:"lastCommit" json:"lastCommit"`
}

type Header struct {
	Version            Consensus `bson:"version" json:"version"`
	ChainID            string    `bson:"chainId" json:"chainId"`
	Height             int64     `bson:"height" json:"height"`
	Time               time.Time `bson:"time" json:"time"`
	LastBlockID        BlockID   `bson:"lastBlockId" json:"lastBlockId"`
	LastCommitHash     string    `bson:"lastCommitHash" json:"lastCommitHash"`
	DataHash           string    `bson:"dataHash" json:"dataHash"`
	ValidatorsHash     string    `bson:"validatorsHash" json:"validatorsHash"`
	NextValidatorsHash string    `bson:"nextValidatorsHash" json:"nextValidatorsHash"`
	ConsensusHash      string    `bson:"consensusHash" json:"consensusHash"`
	AppHash            string    `bson:"appHash" json:"appHash"`
	LastResultsHash    string    `bson:"lastResultsHash" json:"lastResultsHash"`
	EvidenceHash       string    `bson:"evidenceHash" json:"evidenceHash"`
	ProposerAddress    string    `bson:"proposerAddress" json:"proposerAddress"`
}

type Consensus struct {
	Block uint64 `bson:"block" json:"block"`
	App   uint64 `bson:"app" json:"app"`
}

type Data struct {
	Txs []Tx `bson:"txs" json:"txs"`
}

type Tx []byte

type EvidenceData struct {
	Evidence []Evidence `bson:"evidence" json:"evidence"`
}

type Evidence interface{}

type Commit struct {
	Height     int64       `bson:"height" json:"height"`
	Round      int32       `bson:"round" json:"round"`
	BlockID    BlockID     `bson:"blockId" json:"blockId"`
	Signatures []CommitSig `bson:"signatures" json:"signatures"`
}

type CommitSig struct {
	BlockIDFlag      BlockIDFlag `bson:"blockIdFlag" json:"blockIdFlag"`
	ValidatorAddress string      `bson:"validatorAddress" json:"validatorAddress"`
	Timestamp        time.Time   `bson:"timestamp" json:"timestamp"`
	Signature        string      `bson:"signature" json:"signature"`
}

type BlockIDFlag int

const (
	BlockIDFlagAbsent BlockIDFlag = 1
	BlockIDFlagCommit BlockIDFlag = 2
	BlockIDFlagNil    BlockIDFlag = 3
)

func NewCommitSigAbsent() CommitSig {
	return CommitSig{
		BlockIDFlag: BlockIDFlagAbsent,
	}
}

func (b *Block) String() string {
	if b == nil {
		return "nil-Block"
	}
	return "Block{" +
		"\n  " + b.Header.String() +
		"\n  " + b.Data.String() +
		"\n  " + b.Evidence.String() +
		"\n  " + b.LastCommit.String() +
		"\n}#" + b.Hash()
}

func (h Header) String() string {
	return "Header{" +
		"\n    Version: " + h.Version.String() +
		"\n    ChainID: " + h.ChainID +
		"\n    Height: " + fmt.Sprintf("%d", h.Height) +
		"\n    Time: " + h.Time.Format(time.RFC3339Nano) +
		"\n    LastBlockID: " + h.LastBlockID.String() +
		"\n    LastCommit: " + h.LastCommitHash +
		"\n    Data: " + h.DataHash +
		"\n    Validators: " + h.ValidatorsHash +
		"\n    NextValidators: " + h.NextValidatorsHash +
		"\n    App: " + h.AppHash +
		"\n    Consensus: " + h.ConsensusHash +
		"\n    Results: " + h.LastResultsHash +
		"\n    Evidence: " + h.EvidenceHash +
		"\n    Proposer: " + h.ProposerAddress +
		"\n  }#" + h.Hash()
}

func (c Consensus) String() string {
	return fmt.Sprintf("{%d %d}", c.Block, c.App)
}

func (d Data) String() string {
	result := "Data{"
	for _, tx := range d.Txs {
		result += fmt.Sprintf("\n    %s (%d bytes)", hex.EncodeToString(tx), len(tx))
	}
	return result + "\n  }#" + d.Hash()
}

func (e EvidenceData) String() string {
	return "EvidenceData{" +
		"\n  }#" + e.Hash()
}

func (c *Commit) String() string {
	if c == nil {
		return "nil-Commit"
	}
	result := "Commit{" +
		fmt.Sprintf("\n    Height: %d", c.Height) +
		fmt.Sprintf("\n    Round: %d", c.Round) +
		"\n    BlockID: " + c.BlockID.String() +
		"\n    Signatures:"
	for _, sig := range c.Signatures {
		result += "\n      " + sig.String()
	}
	return result + "\n  }#" + c.Hash()
}

func (cs CommitSig) String() string {
	return fmt.Sprintf("CommitSig{%s by %s on %d @ %s}", cs.Signature, cs.ValidatorAddress, cs.BlockIDFlag, cs.Timestamp.Format(time.RFC3339))
}

func (bid BlockID) String() string {
	return fmt.Sprintf("%s:%d:%s", bid.Hash, bid.PartSetHeader.Total, bid.PartSetHeader.Hash)
}

func (b *Block) Hash() string {
	return ""
}

func (h Header) Hash() string {
	return ""
}

func (d Data) Hash() string {
	return ""
}

func (e EvidenceData) Hash() string {
	return ""
}

func (c *Commit) Hash() string {
	return ""
}

func (b *Block) ValidateBasic() error {
	return nil
}
