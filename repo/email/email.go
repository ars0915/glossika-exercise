package email

type EmailClient struct {
}

func NewEmailClient() EmailClient {
	return EmailClient{}
}

func (e EmailClient) SendVerificationEmail(email, verificationCode string) error {
	return nil
}
