package events

type Event string

const (
	EventTypeSendVote                      Event = "sendVote"
	EventTypeReceiveVote                   Event = "receiveVote"
	EventTypeReceivedProposal              Event = "receivedProposal"
	EventTypeReceivedCompleteProposalBlock Event = "receivedCompleteProposalBlock"
	EventTypeStepChange                    Event = "stepChange"
	EventTypeEnteringNewRound              Event = "enteringNewRound"
	EventTypeScheduledTimeout              Event = "scheduledTimeout"
)
