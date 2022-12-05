package token

import "time"

// THis is the interface for the managing tokens
type Maker interface {
	// CreateToken creates a new token for a specific username and valid duration
	CreateToken(username string, duration time.Duration) (string, error)

	// Check if the input token is valid or not. If it is valid, the method will return the payload data stored inside the body of the token.
	VerifyToken(token string) (*Payload, error)
}
