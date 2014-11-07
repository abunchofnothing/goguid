package guid

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	str1 := Generate()

	if len(str1) < 36 {
		t.Error("Error, GUID too short")
	}

	str2 := Generate()

	if str1 == str2 {
		t.Error("Error, duplicate generated")
	}
}
