package services

import (
	"fmt"
	"go-auth/dtos"
	"go-auth/models"
	"go-auth/repositories"
	"math/rand"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService interface {
	Signup(model dtos.AuthSignupRequest) (dtos.AuthSignupResponse, error)
	SignupConfirmOtp(model dtos.AuthSignupConfirmOtpRequest) (dtos.AuthTokenResponse, error)
}

type authService struct {
	userRepo        repositories.UserRepository
	authRegisterReo repositories.AuthRegisterRepository
	jwtService      JwtService
}

func NewAuthService(userRepo repositories.UserRepository, authRegisterRepo repositories.AuthRegisterRepository, jwtService JwtService) AuthService {
	return &authService{
		userRepo:        userRepo,
		authRegisterReo: authRegisterRepo,
		jwtService:      jwtService,
	}
}

func generateOtpCode(authRegisterReo repositories.AuthRegisterRepository) (otpCode string, err error) {
	for {
		var otpCodeNum = rand.Intn(9000) + 1000

		exists, err := authRegisterReo.ExistsByOtpCode(otpCodeNum)
		if err != nil {
			break
		}

		if !exists {
			otpCode = strconv.Itoa(otpCodeNum)
			break
		}
	}

	red := "\033[31m"
	reset := "\033[0m"

	fmt.Println(red + "OtpCode: " + otpCode + reset)

	return
}

func (s *authService) Signup(model dtos.AuthSignupRequest) (dtos.AuthSignupResponse, error) {
	existsBefore, err := s.authRegisterReo.ExistsByPhoneNumber(model.PhoneNumber)

	if err != nil {
		return dtos.AuthSignupResponse{Succeed: false, Message: "Something went wrong"}, err
	}

	// If the user has no record of signup
	if existsBefore == nil {
		randomOtpCode, err := generateOtpCode(s.authRegisterReo)

		if err != nil {
			return dtos.AuthSignupResponse{Succeed: false, Message: "Something went wrong"}, err
		}

		s.authRegisterReo.Create(models.AuthRegister{
			ID:          primitive.NewObjectID(),
			PhoneNumber: model.PhoneNumber,
			OtpCode:     randomOtpCode,
			Trys:        0,
			CreatedDate: primitive.NewDateTimeFromTime(time.Now()),
			UpdatedDate: primitive.NewDateTimeFromTime(time.Now()),
		})

		return dtos.AuthSignupResponse{Succeed: true, Message: "Successfully Registered and OtpCode Sent(as fake)"}, nil
	}

	if time.Since(existsBefore.CreatedDate.Time()) >= 10*time.Minute {
		existsBefore.Trys = 0
		existsBefore.CreatedDate = primitive.NewDateTimeFromTime(time.Now())
	}

	// Check if it less than 3 trys and less than 10 minutes ago
	if existsBefore.Trys < 3 && time.Since(existsBefore.CreatedDate.Time()) < 10*time.Minute {

		if time.Since(existsBefore.UpdatedDate.Time()) < 2*time.Minute {
			return dtos.AuthSignupResponse{Succeed: false, Message: "The Code is not expired, therefore could not send it again"}, nil
		}

		randomOtpCode, err := generateOtpCode(s.authRegisterReo)

		if err != nil {
			return dtos.AuthSignupResponse{Succeed: false, Message: "Something went wrong"}, err
		}

		existsBefore.OtpCode = randomOtpCode
		existsBefore.Trys += 1
		existsBefore.UpdatedDate = primitive.NewDateTimeFromTime(time.Now())

		s.authRegisterReo.Update(*existsBefore)

		return dtos.AuthSignupResponse{Succeed: true, Message: "Code Sent Again"}, nil
	}

	return dtos.AuthSignupResponse{Succeed: false, Message: "Max Code Try Reached, Please Try Again Later"}, nil
}

func (s *authService) SignupConfirmOtp(model dtos.AuthSignupConfirmOtpRequest) (dtos.AuthTokenResponse, error) {
	exists, err := s.authRegisterReo.Exists(model.PhoneNumber, model.OtpCode)

	if err != nil {
		return dtos.AuthTokenResponse{Succeed: false, Message: "Something went wrong"}, err
	}

	if exists {
		s.authRegisterReo.DeleteByPhoneNumber(model.PhoneNumber)

		var user *models.User
		user, err = s.userRepo.ExistsByPhoneNumber(model.PhoneNumber)

		if err != nil {
			return dtos.AuthTokenResponse{Succeed: false, Message: "Something went wrong: In Creating User"}, err
		}

		if user == nil {
			user, err = s.userRepo.Create(models.User{
				ID:          primitive.NewObjectID(),
				PhoneNumber: model.PhoneNumber,
				CreatedDate: time.Now(),
			})

			if err != nil {
				return dtos.AuthTokenResponse{Succeed: false, Message: "Something went wrong: In Creating User"}, err
			}
		}

		token, err := s.jwtService.GenerateToken(user.ID.Hex())

		if err != nil {
			fmt.Println("generate token error: " + err.Error())
			return dtos.AuthTokenResponse{Succeed: false, Message: "Something went wrong: In Generating the Jwt Token"}, err
		}

		return dtos.AuthTokenResponse{Succeed: true, Message: "Succeed", Token: token}, err
	}

	return dtos.AuthTokenResponse{Succeed: false, Message: "Invalid Otp Code"}, err
}
