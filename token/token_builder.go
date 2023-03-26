package token

type TokenBuilder interface {
	CreateToken(UserID int, Username string, UserType string) (string, error)
	VerifyToken(token string) (*Payload, error)
}
