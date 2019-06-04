package translator

import (
	"strings"

	"github.com/ikawaha/tinysegmenter"
)

type personType int

const (
	firstPersonSingular  personType = 1
	secondPersonSingular            = 2
	thirdPersonSingular             = 3
	firstSecondBoth                 = 4
)

var nthPersonTable = map[string]personType{}

var firstPerson = []string{
	"僕", "ぼく", "ボク",
	"私", "わたし",
	"あたし",
	"あたい",
	"俺", "オレ",
	"拙者",
	"我輩", "吾輩",
}

var secondPerson = []string{
	"君", "きみ", "キミ",
	"あなた",
	"あんた",
	"ダーリン",
	"ベイベー",
	"YOU",
	"お前", "おまえ",
}

var thirdPerson = []string{
	"彼", "カレ",
	"彼女", "カノジョ",
}

var firstAndSecondPerson = []string{
	"僕ら", "ぼくら",
}

var kouOtsu = map[personType]string{
	firstPersonSingular:  "甲",
	secondPersonSingular: "乙",
	thirdPersonSingular:  "丙",
	firstSecondBoth:      "甲乙双方",
}

func normalize(s string) string {
	return strings.ToLower(s)
}

func init() {
	for _, s := range firstPerson {
		nthPersonTable[normalize(s)] = firstPersonSingular
	}
	for _, s := range secondPerson {
		nthPersonTable[normalize(s)] = secondPersonSingular
	}
	for _, s := range thirdPerson {
		nthPersonTable[normalize(s)] = thirdPersonSingular
	}
	for _, s := range firstAndSecondPerson {
		nthPersonTable[normalize(s)] = firstSecondBoth
	}
}

func nthPerson(s string) (personType, bool) {
	r, ok := nthPersonTable[normalize(s)]
	return r, ok
}

// Translate translates
func Translate(src string) string {
	terms := tinysegmenter.Segment(src)

	dest := make([]string, len(terms))
	for i := range dest {
		term := terms[i]
		if nth, found := nthPerson(term); found {
			term = kouOtsu[nth]
		}
		dest[i] = term
	}
	return strings.Join(dest, "")
}
