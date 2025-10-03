package user

import (
	"context"
	"errors"
	"os"
	"regexp"
	"server/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrJWTSecretNotSet = errors.New("JWT_SECRET environment variable not set")
	ErrInvalidEmail = errors.New("invalid email format")
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
)

type service struct {
	Repository
	timeout time.Duration
	jwtSecret string
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
		os.Getenv("JWT_SECRET"),
	}
}

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq)(*CreateUserRes, error){
	if !isValidEmail(req.Email) {
		return nil, ErrInvalidEmail
	}
	
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)

	if err != nil {
		return nil, err
	}
	
	u := &User{
		Username : req.Username,
		Email: req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID: strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email: r.Email,
	}

	return res, nil 
}
type Claims struct {
	ID string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims

}

func (s *service) Login(c context.Context, req *LoginUserReq)(*LoginUserRes, error){
	
	if len(s.jwtSecret) == 0 {
		return nil, ErrJWTSecretNotSet
	}
	
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return &LoginUserRes{}, ErrUserNotFound
	}

	err = util.CheckPassword(req.Password, u.Password)

	if err != nil {
		return &LoginUserRes{}, ErrInvalidPassword
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,Claims{
		ID: strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24*time.Hour)),
		},
	} )

	ss, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return &LoginUserRes{}, err
	}

	return &LoginUserRes{accessToken: ss, Username: u.Username, ID: strconv.Itoa(int(u.ID))}, nil
}