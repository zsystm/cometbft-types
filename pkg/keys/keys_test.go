package keys

import "testing"

func TestVoteKey(t *testing.T) {
	vk1 := &VoteKey{
		Height:   1,
		Round:    0,
		ValIdx:   2,
		Sender:   "sender1",
		Receiver: "receiver1",
	}

	vk2 := &VoteKey{
		Height:   1,
		Round:    0,
		ValIdx:   2,
		Sender:   "sender1",
		Receiver: "receiver1",
	}

	// Test map behavior
	m := make(map[string]string)
	m[vk1.Hash()] = "first"
	m[vk2.Hash()] = "second"

	if len(m) != 1 {
		t.Errorf("Expected map length to be 1, got %d", len(m))
	}
	if m[vk1.Hash()] != "second" {
		t.Errorf("Expected value for vk1 to be 'second', got '%s'", m[vk1.Hash()])
	}
	if m[vk2.Hash()] != "second" {
		t.Errorf("Expected value for vk2 to be 'second', got '%s'", m[vk2.Hash()])
	}
	if vk1.Hash() != vk2.Hash() {
		t.Errorf("Expected vk1 and vk2 hashes to be equal, got '%s' and '%s'", vk1.Hash(), vk2.Hash())
	}
}
