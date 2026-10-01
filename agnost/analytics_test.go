package agnost

import (
	"testing"

	"github.com/mark3labs/mcp-go/server"
)

func TestTrackMCPInitializationErrors(t *testing.T) {
	tests := []struct {
		name   string
		server *server.MCPServer
		orgID  string
	}{
		{"nil server", nil, "test-org"},
		{"missing organization", server.NewMCPServer("test", "1.0"), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewAgnostAnalytics()
			if err := client.TrackMCP(tt.server, tt.orgID, nil); err == nil {
				t.Fatal("expected initialization error")
			}
			if client.IsInitialized() {
				t.Fatal("client must remain uninitialized")
			}
			// The lock must also remain usable after the failed call.
			if err := client.TrackMCP(tt.server, tt.orgID, nil); err == nil {
				t.Fatal("expected repeated initialization error")
			}
			client.Shutdown()
		})
	}
}
