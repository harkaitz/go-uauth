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

// CLIAuthorizationFile returns the file where the authorization
// information is stored.
func (u Authority) CLIAuthorizationFile() (file string) {
	var home string
	var err  error
	home, err = os.UserHomeDir()
	if err != nil { log.Panic(err) }
	return home + "/.uauth.json"
}

// CLIConfigurationFile returns the file where the configuration
// information is stored.
func CLIConfigurationFile() (file string) {
	var home string
	var err  error
	home, err = os.UserHomeDir()
	if err != nil { log.Panic(err) }
	return home + "/.config.testing.json"
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
	err = os.WriteFile(u.CLIAuthorizationFile(), userB, 0600)
	if err != nil { return }
	
	return
}

// CLIGetAuthentication gets the user from the saved file.
func (u Authority) CLIGetAuthentication() (user User, found bool, err error) {
	var userB []byte
	
	userB, err = os.ReadFile(u.CLIAuthorizationFile())
	if err != nil { err = nil; return }
	
	err = json.Unmarshal(userB, &user)
	if err != nil { return }
	
	if user.ExpiresAt.Before(time.Now()) { return }
	
	found = true
	return
}
