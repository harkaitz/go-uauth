package uauth

import (
	"embed"
	"errors"
)

// Internationalizable text, import "github.com/harkaitz/go-gettext-l10n"
// to get localized text with "i10n.GetMessage(m, languages...)". For
// text in english simply "m.S".
type Message struct {
	S string
}

// Internationalizable error, import "github.com/harkaitz/go-gettext-l10n"
// to get localized text with "i10n.GetError(m, languages...)".
type userError struct {
	msgUser  Message
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

func l(s string) Message { return Message{s} }

func newErrorE(uMsg Message, err error)                error { return userError{uMsg, err               , ""}    }
func newErrorS(uMsg Message, aMsg string)              error { return userError{uMsg, errors.New(aMsg)  , ""}    }
func newErrorF(uMsg Message, field string)             error { return userError{uMsg, errors.New(uMsg.S), field} }
func newErrorEF(uMsg Message, err error, field string) error { return userError{uMsg, err               , field} }
func newError(uMsg Message)                            error { return userError{uMsg, errors.New(uMsg.S), ""}    }

func (m Message) GetUserMessage()     string   { return m.S  }
func (m Message) String()             string   { return m.S  }
func (m Message) GetDomainName()      string   { return "uauth" }
func (m Message) GetDomainLocaleDir() string   { return "./locale" }
func (m Message) GetDomainLocaleFS()  embed.FS { return locale }

