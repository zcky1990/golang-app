// Package mailers mirrors Rails' app/mailers directory. Concrete mailers
// embed ApplicationMailer, just like Rails' ApplicationMailer base class.
package mailers

import "errors"

// Message represents an outbound email.
type Message struct {
	To      string
	Subject string
	Body    string
}

// DeliverFunc sends a message using the configured backend (SMTP, API, ...).
type DeliverFunc func(from string, msg Message) error

// ApplicationMailer is the base struct every mailer should embed.
type ApplicationMailer struct {
	From        string
	DeliverFunc DeliverFunc
}

// New returns a base mailer with the given sender address.
func New(from string) *ApplicationMailer {
	return &ApplicationMailer{From: from}
}

// Deliver sends the message through the configured backend.
func (m *ApplicationMailer) Deliver(msg Message) error {
	if m.From == "" {
		return errors.New("mailer: sender address is empty")
	}
	if msg.To == "" {
		return errors.New("mailer: recipient is empty")
	}
	if m.DeliverFunc == nil {
		return errors.New("mailer: delivery backend not configured")
	}
	return m.DeliverFunc(m.From, msg)
}
