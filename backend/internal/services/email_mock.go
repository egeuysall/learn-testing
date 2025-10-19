package services

import "errors"

// MockEmailSender - for testing
type MockEmailSender struct {
	SendWelcomeEmailCalled bool
	SentTo                 string
	SentName               string
	ShouldReturnError      bool
}

func (m *MockEmailSender) SendWelcomeEmail(to, name string) error {
	m.SendWelcomeEmailCalled = true
	m.SentTo = to
	m.SentName = name

	if m.ShouldReturnError {
		return errors.New("email failed")
	}

	return nil
}
