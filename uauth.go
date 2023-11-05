package uauth

import (
	"os"
	"time"
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	gorilla "github.com/gorilla/sessions"
)

// Authority is the main object to be used to authenticate.
type Authority struct {
	Settings   Settings
	CLIChannel chan User
}

// Settings contains the required settings for uauth.
type Settings struct {
	UseHTTPS            bool
	Domain              string
	GoogleClientID      string
	GoogleClientSecret  string
	GoogleCallbackPage  string
	RandomString        string
}

// User contains the user information.
type User goth.User

// LoadSettings loads the required settings for uauth from a json file.
func LoadSettings(file string) (s Settings, err error) {
	var configFile *os.File
	var jsonParser *json.Decoder
	
	configFile, err = os.Open(file)
	defer configFile.Close()
	if err != nil { return }
	jsonParser = json.NewDecoder(configFile)
	jsonParser.Decode(&s)
	
	return s, s.UAuthVerify()
}

// UAuthVerify verifies the settings are correct.
func (s Settings) UAuthVerify() (err error) {
	if s.Domain == "" {
		return newError(l("Missing setting: Domain"))
	}
	if s.GoogleClientID == "" {
		return newError(l("Missing setting: GoogleClientID"))
	}
	if s.GoogleClientSecret == "" {
		return newError(l("Missing setting: GoogleClientSecret"))
	}
	if s.GoogleCallbackPage == "" {
		return newError(l("Missing setting: GoogleCallbackPage"))
	}
	if s.RandomString == "" {
		return newError(l("Missing setting: RandomString"))
	}
	return nil
}


// NewAuthority creates a main object to be used to authenticate
// the users. It also registers a session to be used by gin.
func NewAuthority(s Settings, e *gin.Engine, useCLI bool) (u Authority, err error) {
	var provider *google.Provider
	var gstore   *gorilla.CookieStore
	var mstore    memstore.Store
	var cbURL     string
	
	u.Settings = s
	if useCLI {
		if s.UseHTTPS {
			err = newError(l("Cannot use HTTPS with CLI authentication"))
			return
		}
		if s.Domain != "127.0.0.1:8080" {
			err = newError(l("CLI authentication requires the domain to be 127.0.0.1:8080"))
			return
		}
	}
	
	
	// Initialize memstore session named uauth. 
	mstore = memstore.NewStore([]byte(s.RandomString))	
	e.Use(sessions.Sessions("uauth", mstore))
	
	// Initialize gothic
	gstore = gorilla.NewCookieStore([]byte(s.RandomString))
	gstore.MaxAge(86400 * 30 /* 30 days */)
	gstore.Options.Path = "/"
	gstore.Options.HttpOnly = true   // HttpOnly should always be enabled
	gothic.Store = gstore
	
	// Calculate callback page.
	if s.UseHTTPS {
		cbURL = "https://" + s.Domain + s.GoogleCallbackPage
	} else {
		cbURL = "http://" + s.Domain + s.GoogleCallbackPage
	}
	
	// Configure google provider
	provider = google.New(
		s.GoogleClientID,
		s.GoogleClientSecret,
		cbURL,
		"email",
		"profile",
	)
	provider.SetPrompt("select_account")
	provider.SetAccessType("offline")
	goth.UseProviders(provider)
	
	// Create inbound user channel when required. This is used
	// for CLI authentication.
	if useCLI {
		u.CLIChannel = make(chan User)
	}
	
	// This callback is used only for CLI authentication.
	if useCLI {
		e.GET("authenticate", func(ctx *gin.Context) {
			var session sessions.Session
			var user User
			var found bool
			
			// If the authentication is a success return the user
			// through the channel. Print an HTML page with JavaScript
			// code to close the page.
			user, found = u.VerifyUser(ctx)
			if found {
				u.CLIChannel <- user
				ctx.Writer.WriteString("<html><head><script>window.close();</script></head><body></body></html>")
				return
			}
			
			
			// Set the provider for gothic.
			urlValues := ctx.Request.URL.Query()
			urlValues.Add("provider", "google")
			ctx.Request.URL.RawQuery = urlValues.Encode()
			
			// Set the current URL for recovery.
			session = sessions.Default(ctx)
			session.Set("Redirect", ctx.Request.URL.String())
			session.Save()
			
			// Begin authentication.
			gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
		})
	}
	
	e.GET("oauthcb/google", func(ctx *gin.Context) {
		var user goth.User
		var userB []byte
		var err error
		
		// Get the user
		user, err = gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
		if err != nil {
			ctx.AbortWithStatus(500)
			return
		}
		user.RawData = nil
		
		// Marshal user to json and save to session.
		userB, err = json.Marshal(user)
		if err != nil {
			ctx.AbortWithStatus(500)
			return
		}
		
		// Set the session
		session := sessions.Default(ctx)
		session.Set("user", userB)
		session.Save()
		
		// Redirect to the original page.
		var redirect interface{}
		redirect = session.Get("Redirect")
		if redirect != nil {
			ctx.Redirect(http.StatusSeeOther, redirect.(string))
		} else {
			ctx.Redirect(http.StatusSeeOther, "/")
		}
	})
	return
}

// VerifyUser verifies if the user is authenticated.
func (u Authority) VerifyUser(ctx *gin.Context) (user User, found bool) {
	var session sessions.Session
	
	session = sessions.Default(ctx)
	if session.Get("user") != nil {
		json.Unmarshal(session.Get("user").([]byte), &user)
		found = true
	}
	
	// Check expiration date.
	if user.ExpiresAt.Before(time.Now()) {
		found = false
	}
	
	return
}

// String returns the user information in a string.
func (u User) String() (s string) {
	var userB []byte
	userB, _ = json.Marshal(u)
	return string(userB)
}
