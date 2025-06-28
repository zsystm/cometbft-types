package events

type EventTyp string

const (
	EventTypeSendVote                      EventTyp = "sendVote"
	EventTypeReceiveVote                   EventTyp = "receiveVote"
	EventTypeReceivedProposal              EventTyp = "receivedProposal"
	EventTypeReceivedCompleteProposalBlock EventTyp = "receivedCompleteProposalBlock"
	EventTypeStepChange                    EventTyp = "stepChange"
	EventTypeEnteringNewRound              EventTyp = "enteringNewRound"
	EventTypeScheduledTimeout              EventTyp = "scheduledTimeout"
)
