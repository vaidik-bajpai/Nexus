package mock

import (
	"github.com/stretchr/testify/mock"
)

type MockMailer struct {
	mock.Mock
}

func (m *MockMailer) SendEmailVerificationEmail(to []string, subject string, verificationURL string) error {
	args := m.Called(to, subject, verificationURL)
	return args.Error(0)
}

func (m *MockMailer) SendPasswordResetEmail(to []string, subject string, passwordResetURL string) error {
	args := m.Called(to, subject, passwordResetURL)
	return args.Error(0)
}

func (m *MockMailer) SendBoardInvitationEmail(to []string, subject, inviterName, boardName, invitationURL string) error {
	args := m.Called(to, subject, inviterName, boardName, invitationURL)
	return args.Error(0)
}
