package events

import "time"

type BaseEvent struct {
	Type      Event     `bson:"type"`
	Timestamp time.Time `bson:"timestamp"`
}
