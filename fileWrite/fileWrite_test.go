package fileWrite

import "testing"

func TestFileWrite(t *testing.T) {
    result := Hello("Ryunosuke")
    want := "Hi, Ryunosuke. Welcome!"
    if result != want {
        t.Errorf("fileWrite.Hello() = %q want %q", result, want)
    }
}
