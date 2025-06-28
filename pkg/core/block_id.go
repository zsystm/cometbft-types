package core

type BlockID struct {
	Hash          string        `bson:"hash"`
	PartSetHeader PartSetHeader `bson:"partSetHeader"`
}

type PartSetHeader struct {
	Total uint64 `bson:"total"`
	Hash  string `bson:"hash"`
}
