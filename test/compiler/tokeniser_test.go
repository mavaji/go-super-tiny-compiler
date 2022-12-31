package compiler

import (
	"go-super-tiny-compiler/src/compiler"
	"reflect"
	"testing"
)

func TestTokenizer(t *testing.T) {
	tokenizer := compiler.NewTokenizer()
	t.Run("can tokenize an input string", func(t *testing.T) {
		input := "(add 2 (subtract 4 2))"
		want := []compiler.Token{
			{
				Type:  "paren",
				Value: "(",
			},
			{
				Type:  "name",
				Value: "add",
			},
			{
				Type:  "number",
				Value: "2",
			},
			{
				Type:  "paren",
				Value: "(",
			},
			{
				Type:  "name",
				Value: "subtract",
			},
			{
				Type:  "number",
				Value: "4",
			},
			{
				Type:  "number",
				Value: "2",
			},
			{
				Type:  "paren",
				Value: ")",
			},
			{
				Type:  "paren",
				Value: ")",
			},
		}
		tokens, err := tokenizer.Tokenize(input)
		if err != nil {
			t.Errorf("want no error but got error: %s", err)
		}
		if !reflect.DeepEqual(tokens, want) {
			t.Errorf("want %v but got %v", want, tokens)
		}
	})

	t.Run("returns error for lexical errors", func(t *testing.T) {
		input := "(add 2, (subtract 4 2))"
		tokens, err := tokenizer.Tokenize(input)
		wantError := "lexical error: ',' at position: 6"
		if err.Error() != wantError {
			t.Errorf("want error %s but got %v", wantError, err)
		}
		if tokens != nil {
			t.Errorf("want nil but got %v", tokens)
		}
	})
}
