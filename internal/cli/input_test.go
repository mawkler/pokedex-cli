package cli

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase = struct {
	args     string
	wantHead string
	wantTail []string
}

func Test_SplitInput(t *testing.T) {
	tests := []testCase{
		testCase{
			args:     "  foo bar ",
			wantHead: "foo",
			wantTail: []string{"bar"},
		},
		testCase{
			args:     "foo bar baz faz",
			wantHead: "foo",
			wantTail: []string{"bar", "baz", "faz"},
		},
		testCase{
			args:     "",
			wantHead: "",
			wantTail: []string{},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test nr %d", i), func(t *testing.T) {
			gotHead, gotTail := SplitInput(tt.args)
			if gotHead != tt.wantHead {
				t.Errorf("splitInput() gotHead = %v, want %v", gotHead, tt.wantHead)
			}
			if !reflect.DeepEqual(gotTail, tt.wantTail) {
				t.Errorf("splitInput() gotTail = %v, want %v", gotTail, tt.wantTail)
			}
		})
	}
}
