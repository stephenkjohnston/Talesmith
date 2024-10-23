package parser

import "testing"

func TestDuplicateScene(t *testing.T) {
	input := `
	SCENE "Start"
	SCENE "Start"
	`

	parser := NewParser(input)
	err := parser.Parse()

	if err == nil {
		t.Errorf("Expected duplicate scene error, but got none")
	} else if err.Error() != "duplicate scene name found \"Start\"" {
		t.Errorf("Unexpected error message: %v", err)
	}
}

func TestAddScene(t *testing.T) {
	input := `SCENE "Start"`

	parser := NewParser(input)
	if err := parser.Parse(); err != nil {
		t.Errorf("Unexpected error message: %v", err)
	}

	if len(parser.Scenes) != 1 {
		t.Errorf("Exepected one scene, but got %d", len(parser.Scenes))
	}
}
