package uauth

import (
	"os"
	"time"
)

// CLIGetAuthentication gets the user from the saved file.
func CLIGetAuthentication() (user User, err error) {
	user.Email = os.Getenv("UAUTH_EMAIL")
	user.ExpiresAt = time.Now().Add(time.Hour)
	if user.Email == "" {
		err = newError(l("Please set UAUTH_EMAIL environment variable"))
		return
	}
	return
}

// Root should be allowed to everything.
func Root() (user User) {
	user.Email = "root"
	user.ExpiresAt = time.Now().Add(time.Hour*24*365*5)
	return
}
