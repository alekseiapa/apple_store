package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 32

// JWTMaker implements the token maker interface
type JWTMaker struct {
	// use symmetric key algorithm to sign the tokens so this struct will have a field to store the secret key
	secretKey string
}

// create a new JWTMaker
// By returning the interface, we will make sure that
// our JWTMaker must implement the token maker interface.
func NewJWTMaker(secretKey string) (Maker, error) {
	// ensure that the key should not be too short
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	// create a new jwtToken, First param is the signing algorithm
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	// function that receives the parsed but unverified token. You should verify its header to make sure that the signing algorithm matches with what you normally use to sign the tokens. if it matches then return the key
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// get its signing algorithm and convert it to a specific SigningMethodHMAC because we’re using HS256 which is an instance of the SigningMethodHMAC struct
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			// the algorithm of the token doesn’t match with our signing algorithm
			return nil, ErrInvalidToken
		}
		// secretKey using to sign the token
		return []byte(maker.secretKey), nil
	}
	// parse the token
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		// in order to figure out the real error type we have to convert the returned error of the ParseWithClaims function to jwt.ValidationError
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}
	// get token payload data
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
