package latency

import "time"

// LatencyHistogram represents latency distribution for a normalized node pair and message type.
type LatencyHistogram struct {
	NodePairKey    string `json:"nodePairKey" bson:"nodePairKey"` // Canonical "nodeA:nodeB" (alphabetically ordered)
	Node1ID        string `json:"node1Id" bson:"node1Id"`         // First node (alphabetically)
	Node2ID        string `json:"node2Id" bson:"node2Id"`         // Second node (alphabetically)
	MessageType    string `json:"messageType" bson:"messageType"`
	Node1Validator string `json:"node1Validator" bson:"node1Validator"` // Node1's validator address
	Node2Validator string `json:"node2Validator" bson:"node2Validator"` // Node2's validator address

	// Latency measurements (milliseconds)
	Latencies []int64 `json:"latenciesMs" bson:"latenciesMs"`
	Count     int     `json:"count" bson:"count"`

	// Statistical metrics
	MinLatency    int64 `json:"minLatencyMs" bson:"minLatencyMs"`
	MaxLatency    int64 `json:"maxLatencyMs" bson:"maxLatencyMs"`
	MeanLatency   int64 `json:"meanLatencyMs" bson:"meanLatencyMs"`
	MedianLatency int64 `json:"medianLatencyMs" bson:"medianLatencyMs"`
	P95Latency    int64 `json:"p95LatencyMs" bson:"p95LatencyMs"`
	P99Latency    int64 `json:"p99LatencyMs" bson:"p99LatencyMs"`

	// Time range
	FirstSeen time.Time `json:"firstSeen" bson:"firstSeen"`
	LastSeen  time.Time `json:"lastSeen" bson:"lastSeen"`

	// Percentile-based categorization
	BelowP50Count int `json:"belowP50Count" bson:"belowP50Count"` // Messages below median
	P50ToP95Count int `json:"p50ToP95Count" bson:"p50ToP95Count"` // Messages between P50-P95
	P95ToP99Count int `json:"p95ToP99Count" bson:"p95ToP99Count"` // Messages between P95-P99
	AboveP99Count int `json:"aboveP99Count" bson:"aboveP99Count"` // Messages above P99
}

// UnmatchedMessageStats provides statistics about unmatched messages across all nodes.
type UnmatchedMessageStats struct {
	// Unmatched sends (sends without receives)
	TotalUnmatchedSends int `json:"totalUnmatchedSends" bson:"totalUnmatchedSends"`

	// Unmatched receives (receives without sends)
	TotalUnmatchedReceives int `json:"totalUnmatchedReceives" bson:"totalUnmatchedReceives"`

	// Overall health metrics across all nodes
	TotalMessages int `json:"totalMessages" bson:"totalMessages"`
}

// NodeNetworkStats provides network statistics for a specific node.
type NodeNetworkStats struct {
	NodeID           string `json:"nodeId" bson:"nodeId"`
	ValidatorAddress string `json:"validatorAddress" bson:"validatorAddress"`

	// Message counts
	TotalSends    int `json:"totalSends" bson:"totalSends"`
	TotalReceives int `json:"totalReceives" bson:"totalReceives"`

	// Unmatched messages for this node
	UnmatchedSends    int `json:"unmatchedSends" bson:"unmatchedSends"`
	UnmatchedReceives int `json:"unmatchedReceives" bson:"unmatchedReceives"`

	// Peer connections for this node
	ConnectedPeers []string `json:"connectedPeers" bson:"connectedPeers"`
	PeerCount      int      `json:"peerCount" bson:"peerCount"`
}

// NodePairLatencyStats aggregates latency data across all message types for a node pair.
type NodePairLatencyStats struct {
	NodePairKey    string                       `json:"nodePairKey" bson:"nodePairKey"`       // Canonical "nodeA:nodeB"
	Node1ID        string                       `json:"node1Id" bson:"node1Id"`               // First node (alphabetically)
	Node2ID        string                       `json:"node2Id" bson:"node2Id"`               // Second node (alphabetically)
	Node1Validator string                       `json:"node1Validator" bson:"node1Validator"` // Node1's validator address
	Node2Validator string                       `json:"node2Validator" bson:"node2Validator"` // Node2's validator address
	MessageTypes   map[string]*LatencyHistogram `json:"messageTypes" bson:"messageTypes"`
	OverallStats   *LatencyHistogram            `json:"overallStats" bson:"overallStats"`
}
