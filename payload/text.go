package payload

import (
	"math/rand"
	"time"
)

const (
	throwAwayTrash = iota
	deskWiperA
	deskWiperB
	vacuum
	chairWiper
	tidy
)

func MakeAttachmentText() string {
	var roles []int
	var text string
	// If the number of members you want to assign roles increases, add the SlackID here.
	m := map[int]string{0: "<@UBN2JTJTH>", 1: "<@UKAJQR2R1>", 2: "<@UD464EFPU>", 3: "<@U9YUS6U65>", 4: "<@UCSPXT797>", 5: "<@UGM7FH695>"}
	rand.Seed(time.Now().UnixNano())
	for len(roles) < 6 {
		r := rand.Intn(6)
		b := isUsed(r, roles)
		if b == false {
			roles = append(roles, r)
		}
	}
	for k, v := range roles {
		switch k {
		case throwAwayTrash:
			text += "ゴミ捨て : " + m[v] + "\n"
		case deskWiperA:
			text += "机拭き A : " + m[v] + "\n"
		case deskWiperB:
			text += "机拭き B : " + m[v] + "\n"
		case vacuum:
			text += "掃除機 : " + m[v] + "\n"
		case chairWiper:
			text += "イスコロ : " + m[v] + "\n"
		case tidy:
			text += "整理整頓 : " + m[v] + "\n"
		}
	}
	return text
}

// Check whether a randomly generated number is already used twice
func isUsed(r int, s []int) bool {
	for _, s := range s {
		if s == r {
			return true
		}
	}
	return false
}
