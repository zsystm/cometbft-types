package events

type EventTyp string

const (
	EventTypeSendVote                          EventTyp = "sendVote"
	EventTypeSendNewRoundStep                  EventTyp = "sendNewRoundStep"
	EventTypeSendNewValidBlock                 EventTyp = "sendNewValidBlock"
	EventTypeSendProposal                      EventTyp = "sendProposal"
	EventTypeSendProposalPOL                   EventTyp = "sendProposalPOL"
	EventTypeSendBlockPart                     EventTyp = "sendBlockPart"
	EventTypeSendHasVote                       EventTyp = "sendHasVote"
	EventTypeSendVoteSetMaj23                  EventTyp = "sendVoteSetMaj23"
	EventTypeSendVoteSetBits                   EventTyp = "sendVoteSetBits"
	EventTypeSendHasProposalBlockPart          EventTyp = "sendHasProposalBlockPart"
	EventTypeReceivePacketNewRoundStep         EventTyp = "receivePacketNewRoundStep"
	EventTypeReceivePacketNewValidBlock        EventTyp = "receivePacketNewValidBlock"
	EventTypeReceivePacketProposal             EventTyp = "receivePacketProposal"
	EventTypeReceivePacketProposalPOL          EventTyp = "receivePacketProposalPOL"
	EventTypeReceivePacketBlockPart            EventTyp = "receivePacketBlockPart"
	EventTypeReceivePacketVote                 EventTyp = "receivePacketVote"
	EventTypeReceivePacketHasVote              EventTyp = "receivePacketHasVote"
	EventTypeReceivePacketVoteSetMaj23         EventTyp = "receivePacketVoteSetMaj23"
	EventTypeReceivePacketVoteSetBits          EventTyp = "receivePacketVoteSetBits"
	EventTypeReceivePacketHasProposalBlockPart EventTyp = "receivePacketHasProposalBlockPart"
	EventTypeReceiveVote                       EventTyp = "receiveVote"
	EventTypeReceivedProposal                  EventTyp = "receivedProposal"
	EventTypeReceivedCompleteProposalBlock     EventTyp = "receivedCompleteProposalBlock"
	EventTypeStepChange                        EventTyp = "stepChange"
	EventTypeEnteringNewRound                  EventTyp = "enteringNewRound"
	EventTypeScheduledTimeout                  EventTyp = "scheduledTimeout"
)
