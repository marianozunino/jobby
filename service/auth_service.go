package service

import (
	"context"
	"crypto"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/ent"
	"github.com/marianozunino/cc-backend-go/store"
)

var _ AuthService = &authService{}

type authService struct {
	Store store.Store
}

func (a *authService) CreateToken(ctx context.Context, input *ent.User) (*dtos.Auth, error) {
	// set sub to user id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": input.ID,
	})

	// sign token with secret
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &dtos.Auth{
		AccessToken: tokenString,
	}, nil
}

// DecodeToken implements AuthService.
func (a *authService) DecodeToken(ctx context.Context, token string) (*ent.User, error) {
	// get user id from token claims
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, err
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return nil, err
	}

	subAsUUID, err := uuid.Parse(sub)
	if err != nil {
		return nil, err
	}

	// get user from Store
	user, err := a.Store.User(ctx, subAsUUID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *authService) Login(ctx context.Context, input dtos.AuthInput) (*dtos.Auth, error) {
	// get user from Store
	user, err := a.Store.UserByEmail(ctx, input.Email)

	if err != nil {
		return nil, err
	}

	isCorrect := mustComparePassword(ctx, input.Password, user)

	if !isCorrect {
		return nil, fmt.Errorf("invalid credentials")
	}

	// create token
	auth, err := a.CreateToken(ctx, user)

	if err != nil {
		return nil, err
	}

	return auth, nil
}

func mustHashPassword(ctx context.Context, plainPassword string, salt string) string {
	hash := crypto.SHA1.New()
	hash.Write([]byte(salt + plainPassword))
	hashedPassword := hash.Sum(nil)
	return hex.EncodeToString(hashedPassword)
}

func mustGenenerateSalt(ctx context.Context) string {
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		panic(err) // out of randomness, should never happen
	}
	hexString := hex.EncodeToString(randomBytes)
	return hexString
}

func mustComparePassword(ctx context.Context, plainPassword string, user *ent.User) bool {
	hashedPassword := mustHashPassword(ctx, plainPassword, user.Salt)
	return hashedPassword == user.Password
}
