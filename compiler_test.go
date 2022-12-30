package main

import "testing"

func TestCompiler(t *testing.T) {
	t.Run("can compile an input string", func(t *testing.T) {
		compiler := Compiler{}
		result, err := compiler.compile("(add 2 (subtract 4 2))")

		want := "add(2, subtract(4, 2));"
		if err != nil {
			t.Errorf("want no error but got error: %s", err)
		}
		if result != want {
			t.Errorf("want %v but got %v", want, result)
		}
	})

	t.Run("returns error for invalid input", func(t *testing.T) {
		compiler := Compiler{}
		result, err := compiler.compile("(add 2 )subtract 4 2))")

		wantError := "syntax error: 'subtract' at position: 4"
		if err.Error() != wantError {
			t.Errorf("want error %s but got %v", wantError, err)
		}
		if result != "" {
			t.Errorf("want empty string but got %v", result)
		}
	})
}
