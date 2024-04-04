package uauth

import (
	"embed"
	"errors"
)

// Internationalizable text. Execute "genl10n_go" in your
// project to generate "l10n.go" with functions to extract
// internationalized texts from this.
type message struct {
	S string
}

// Internationalizable error, Execute "genl10n_go" in your
// project to generate "l10n.go" with functions to extract
// internationalized texts from this.
type userError struct {
	msgUser  message
	msgAdmin error
	field    string
}

//go:embed locale
var locale embed.FS

func (e userError) Error()              string   { return e.msgAdmin.Error() }
func (e userError) GetUserMessage()     string   { return e.msgUser.S  }
func (e userError) GetDomainName()      string   { return "uauth" }
func (e userError) GetDomainLocaleDir() string   { return "./locale" }
func (e userError) GetDomainLocaleFS()  embed.FS { return locale }
func (e userError) GetField()           string   { return e.field }

func l(s string) message { return message{s} }

func newErrorE(uMsg message, err error)                error { return userError{uMsg, err               , ""}    }
func newErrorS(uMsg message, aMsg string)              error { return userError{uMsg, errors.New(aMsg)  , ""}    }
func newErrorF(uMsg message, field string)             error { return userError{uMsg, errors.New(uMsg.S), field} }
func newErrorEF(uMsg message, err error, field string) error { return userError{uMsg, err               , field} }
func newError(uMsg message)                            error { return userError{uMsg, errors.New(uMsg.S), ""}    }

func (m message) GetUserMessage()     string   { return m.S  }
func (m message) String()             string   { return m.S  }
func (m message) GetDomainName()      string   { return "uauth" }
func (m message) GetDomainLocaleDir() string   { return "./locale" }
func (m message) GetDomainLocaleFS()  embed.FS { return locale }

