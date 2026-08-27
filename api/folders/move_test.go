package folders

import (
	"context"
	"testing"
)

func TestMoveFolderRequiresFolderAndSpaceIDs(t *testing.T) {
	tests := []struct {
		name     string
		folderID string
		spaceID  string
	}{
		{name: "missing folder", spaceID: "space-1"},
		{name: "missing space", folderID: "folder-1"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service := &Service{}
			if err := service.MoveFolder(context.Background(), test.folderID, test.spaceID); err == nil {
				t.Fatal("MoveFolder() error = nil")
			}
		})
	}
}
