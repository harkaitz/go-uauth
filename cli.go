package uauth

import (
	"os"
	"time"
	"log"
	"os/exec"
	"encoding/json"
)

// CLIAuthenticateURL returns the URL to authenticate the user.
func (u Authority) CLIAuthenticateURL() (url string) {
	if u.Settings.UseHTTPS {
		url = "https://" + u.Settings.Domain + "/authenticate"
	} else {
		url = "http://" + u.Settings.Domain + "/authenticate"
	}
	return url
}

// CLIAuthenticate opens a browser to authenticate the user.
func (u Authority) CLIAuthenticate() (user User, err error) {
	var cmd *exec.Cmd
	var userB []byte
	
	// Open authenticate page with the browser.
	cmd = exec.Command("xdg-open", u.CLIAuthenticateURL())
	err = cmd.Run()
	if err != nil { return }
	
	// Wait user in channel.
	user = <- u.CLIChannel
	
	// Save JSON file in ~/.uauth
	userB, err = json.Marshal(user)
	if err != nil { return }
	err = os.WriteFile(CLIAuthorizationFile(), userB, 0600)
	if err != nil { return }
	
	return
}

// CLIAuthorizationFile returns the file where the authorization
// information is stored.
func CLIAuthorizationFile() (file string) {
	var home string
	var err  error
	home, err = os.UserHomeDir()
	if err != nil { log.Panic(err) }
	return home + "/.uauth.json"
}

// CLIGetAuthentication gets the user from the saved file.
func CLIGetAuthentication() (user User, found bool, err error) {
	var userB []byte
	
	userB, err = os.ReadFile(CLIAuthorizationFile())
	if err != nil { err = nil; return }
	
	err = json.Unmarshal(userB, &user)
	if err != nil { return }
	
	if user.ExpiresAt.Before(time.Now()) { return }
	
	found = true
	return
}

// CLIPleaseAuthenticateError returns an error to ask the user to
// authenticate.
func CLIPleaseAuthenticateError() (err error) {
	err = newError(l("Please authenticate with \"uauth -f\"."))
	return
}
