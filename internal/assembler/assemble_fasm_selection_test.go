package assembler

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type fasmRunCapture struct {
	fasmCalled bool
}

func TestAsm_ForceFASMRoutesToFasm(t *testing.T) {
	oldForceFASM := ForceFASM
	oldUseNasm := UseNasm
	defer func() {
		ForceFASM = oldForceFASM
		UseNasm = oldUseNasm
	}()

	ForceFASM = true
	UseNasm = false

	cap := &fasmRunCapture{}
	oldRun := runCommand
	defer func() { runCommand = oldRun }()
	SetRunCommand(func(ctx context.Context, verbose bool, name string, args ...string) (string, error) {
		if strings.Contains(name, "fasm") {
			cap.fasmCalled = true
		}
		return "", nil
	})

	dir := t.TempDir()
	src := filepath.Join(dir, "t.asm")
	if err := os.WriteFile(src, []byte("format ELF64\nsection '.text'\nstart:\n  nop\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	obj := filepath.Join(dir, "t.o")

	_ = Assemble(context.Background(), src, obj, false, false, "raw")

	if !cap.fasmCalled {
		t.Fatal("fasm was not invoked, expected FASM routing when ForceFASM=true")
	}
}
