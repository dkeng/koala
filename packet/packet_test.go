package packet

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	p, err := New([]byte("da keng is good."))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(p.Bytes())
}

func TestNewQrs(t *testing.T) {
	p := NewQrs()
	fmt.Println(p.Bytes())
}

func TestNewClose(t *testing.T) {
	p := NewClose()
	fmt.Println(p.Bytes())
}
