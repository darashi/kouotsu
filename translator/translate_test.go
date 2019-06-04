package translator

import (
	"testing"
)

func TestExampleSuccess(t *testing.T) {
	tt := []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "私「わたしは君のこと絶対に忘れないから！」",
			Expected: "甲「甲は乙のこと絶対に忘れないから！」",
		},
		{
			Input:    "わたしたいものがあります",
			Expected: "わたしたいものがあります",
		},
		{
			Input:    "YoU",
			Expected: "乙",
		},
		{
			Input:    "いつだって僕らは",
			Expected: "いつだって甲乙双方は",
		},
	}
	for _, tc := range tt {
		actual := Translate(tc.Input)
		if actual != tc.Expected {
			t.Fatalf("expected=%v, actual=%v; input=%v", tc.Expected, actual, tc.Input)
		}
	}
}
