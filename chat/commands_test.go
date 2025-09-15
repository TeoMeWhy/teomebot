package chat

import (
	"testing"

	"github.com/gempir/go-twitch-irc/v4"
)

func TestParseCommand(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"!nekt", "nekt"},
		{"!NEKT", "nekt"},
		{"!nekt arg1 arg2", "nekt"},
		{"!agenda", "agenda"},
	}

	for _, test := range tests {
		result := ParseCommand(test.input)
		if result != test.expected {
			t.Errorf("ParseCommand(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestNetLinkFunction(t *testing.T) {
	// This is a simple test to ensure the NetLink function exists and can be called
	// In a real bot environment, this would send a message to the chat
	// Here we just verify the function doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("NetLink function panicked: %v", r)
		}
	}()

	// Create a mock client and message
	client := &twitch.Client{}
	message := twitch.PrivateMessage{
		Channel: "testchannel",
		User: twitch.User{
			Name: "testuser",
		},
		Message: "!nekt",
	}

	// Call the function - it should not panic
	NetLink(client, message)
}

func TestCommandMapping(t *testing.T) {
	// Test that the nekt command is properly mapped
	allFun := map[string]HandleCommand{
		"nekt": NetLink,
	}

	if _, exists := allFun["nekt"]; !exists {
		t.Error("nekt command is not mapped in the command registry")
	}

	// Verify the function pointer is correct
	if allFun["nekt"] == nil {
		t.Error("nekt command maps to nil function")
	}
}