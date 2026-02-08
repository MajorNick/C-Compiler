package internal

import (
	"strings"
	"testing"
)

func TestCompiler(t *testing.T) {
	input := `
	int main() {
		return 5 + 3 * 2;
	}`

	c := CompilerFromSource(input)
	output, errors := c.Compile()

	if len(errors) > 0 {
		t.Fatalf("Compilation failed with errors: %v", errors)
	}

	expectedSubstrings := []string{
		"global main",
		"main:",
		"push rbp",
		"mov rax, 5",
		"push rax",
		"mov rax, 3",
		"push rax",
		"mov rax, 2",
		"mov rbx, rax",
		"pop rax",
		"imul rax, rbx",
		"mov rbx, rax",
		"pop rax",
		"add rax, rbx",
		"ret",
	}
	for _, substr := range expectedSubstrings {
		if !strings.Contains(output, substr) {
			t.Errorf("Expected output to contain %q, but it didn't.\nOutput:\n%s", substr, output)
		}
	}
}
