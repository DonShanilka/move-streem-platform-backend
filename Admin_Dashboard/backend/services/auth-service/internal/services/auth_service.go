package services

import (
    "github.com/DonShanilka/auth-service/internal/models"
    "github.com/DonShanilka/auth-service/internal/repository"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
    "time"
    "errors"
)

type AuthService interface {
    Register(user *models.User) error
    Login(email, password string) (*models.TokenResponse, error)
}

type authService struct {
    repo repository.UserRepository
    JWTSecret string
}

func NewAuthService(repo repository.UserRepository, jwtSecret string) AuthService {
    return &authService{
        repo: repo,
        JWTSecret: jwtSecret,
    }
}

func (s *authService) Register(user *models.User) error {
    hash, _:= bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    user.Password = string(hash)
    return s.repo.CreateUser(user)
}

func (s *authService) Login(email, password string) (*models.TokenResponse, error) {
    user, err := s.repo.FindUserByEmail(email)
    if err != nil {
        return nil, errors.New("Email not found")
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return nil, errors.New("Invalid password")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userID": user.ID,
        "email":  user.Email,
        "exp":    time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, _ := token.SignedString([]byte(s.JWTSecret))
    
    return &models.TokenResponse{Token: tokenString}, nil
}