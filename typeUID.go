package uauth

import (
	"strings"
)

type UID string

func (u UID) CheckFormat() (UID, error) {
	if len(u) < 5 {
		return u, newError(l("Missing user mail address (E1001)"))
	}
	if strings.Index(string(u), "@") == -1 {
		return u, newError(l("Invalid user mail address (E1002)"))
	}
	return u, nil
}

func (u User) GetUID() UID {
	return UID(u.Email)
}

func Str2UID(s string) UID {
	return UID(strings.TrimSpace(strings.ToLower(s)))
}
