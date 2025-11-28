package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/spf13/cobra"
)

// captureOutput captura o stdout durante a execução da função f.
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	_ = w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

func TestParseDuration(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Duration
		isErr    bool
	}{
		{"1h30m", time.Hour + 30*time.Minute, false},
		{"45m", 45 * time.Minute, false},
		{"10s", 10 * time.Second, false},
		{"2", 2 * time.Minute, false}, // raw number → minutes
		{"invalid", 0, true},          // expect error
	}

	for _, tt := range tests {
		got, err := parseDuration(tt.input)

		if tt.isErr && err == nil {
			t.Errorf("expected error for %s", tt.input)
		}
		if !tt.isErr && err != nil {
			t.Errorf("unexpected error for %s: %v", tt.input, err)
		}
		if !tt.isErr && got != tt.expected {
			t.Errorf("for %s expected %v but got %v", tt.input, tt.expected, got)
		}
	}
}

func TestVerboseOutput(t *testing.T) {
	verbose = true
	defer func() { verbose = false }()

	out := captureOutput(func() {
		printf("hello %s", "world")
		println("line test")
	})

	if !strings.Contains(out, "hello world") {
		t.Errorf("printf did not output expected text; got: %q", out)
	}

	if !strings.Contains(out, "line test") {
		t.Errorf("println did not output expected text; got: %q", out)
	}
}

func TestRunAutoPress(t *testing.T) {
	// garantir que verbose não polui a saída
	verbose = false

	// salvar o valor original e restaurar depois
	origTypeStr := typeStr
	defer func() { typeStr = origTypeStr }()

	var typed []string

	// mock com a assinatura correta: func(str string, args ...int)
	typeStr = func(str string, args ...int) {
		typed = append(typed, str)
	}

	// usar intervalos pequenos para o teste
	interval = "10ms"
	duration = "35ms"

	out := captureOutput(func() {
		runAutoPress(&cobra.Command{}, []string{})
	})

	if len(typed) == 0 {
		t.Fatalf("expected at least one key press, got none")
	}

	if !strings.Contains(out, "Duration completed") {
		t.Fatalf("missing final message; got output: %q", out)
	}
}
