package application

type MessageEmailService struct {
}

func NewMessageEmailService() *MessageEmailService {
	return &MessageEmailService{}
}

func (s *MessageEmailService) Send(emailMessage string) error {
	return nil
}

func (s *MessageEmailService) SendConfirmEmail(email, callbackUrl string) error {
	return nil
}
