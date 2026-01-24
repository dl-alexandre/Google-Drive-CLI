package cli

import "testing"

func TestParseSuspendedFlag(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		got, err := parseSuspendedFlag("true")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got == nil || *got != true {
			t.Fatalf("expected true")
		}
	})

	t.Run("false", func(t *testing.T) {
		got, err := parseSuspendedFlag("false")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got == nil || *got != false {
			t.Fatalf("expected false")
		}
	})

	t.Run("invalid", func(t *testing.T) {
		if _, err := parseSuspendedFlag("yes"); err == nil {
			t.Fatalf("expected error")
		}
	})

	t.Run("empty", func(t *testing.T) {
		if _, err := parseSuspendedFlag(""); err == nil {
			t.Fatalf("expected error")
		}
	})

	t.Run("uppercase", func(t *testing.T) {
		if _, err := parseSuspendedFlag("TRUE"); err == nil {
			t.Fatalf("expected error")
		}
	})
}
