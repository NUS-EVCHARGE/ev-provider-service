package authentication

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/sirupsen/logrus"
	_ "golang.org/x/net/context"
)

type AuthenticationController interface {
	RegisterUser(signUpRequest dto.SignUpRequest) error
	LoginUser(loginCredential dto.Credentials) (*dto.LoginResponse, error)
	ConfirmUser(userInfo dto.ConfirmUser) error
	ResendChallengeCode(resendRequest dto.SignUpResendRequest) error
	LogoutUser(accessToken string) error
	GetUserInfo(accessToken string) (*cognitoidentityprovider.GetUserOutput, error)
}

type AuthenticationControllerImpl struct {
}

var (
	awsRegion = "ap-southeast-1" // Your AWS Region
	//userPoolID                  = "ap-southeast-1_wnUcfMgqN"  // Your Cognito User Pool ID
	clientID                    = "og5uq3m2bvhfbghf3jd2q14jm"
	clientSecret                = "16q37emcuik0cbfffo534lsqo2kck4fisjp7gnkpmbil2br6bho"
	AuthenticationControllerObj AuthenticationController
	cognitoClient               = setupCognitoClient()
)

func NewAuthenticationController() {
	AuthenticationControllerObj = &AuthenticationControllerImpl{}
}

func (a AuthenticationControllerImpl) LoginUser(loginCredential dto.Credentials) (*dto.LoginResponse, error) {

	secretHash := generateSecretHash(clientSecret, loginCredential.Email, clientID)

	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME":    aws.String(loginCredential.Email),
			"PASSWORD":    aws.String(loginCredential.Password),
			"SECRET_HASH": aws.String(secretHash),
		},
		ClientId: aws.String(clientID),
	}
	login := &dto.LoginResponse{}

	authResp, err := cognitoClient.InitiateAuth(input)
	if err != nil {
		return nil, err
	}

	// Handle MFA setup and challenges
	if authResp.ChallengeName != nil && *authResp.ChallengeName != "" {
		fmt.Println("Challenge required:", authResp.ChallengeName)
		logrus.WithField("login", login).Info("User authentication failed \n", loginCredential.Email)
		login.Status = *authResp.ChallengeName
		//handleChallenge(cognitoClient, clientID, clientSecret, loginCredential.Email, *authResp.Session, *authResp.ChallengeName)
	} else {
		login.AccessToken = *authResp.AuthenticationResult.AccessToken
		login.RefreshToken = *authResp.AuthenticationResult.RefreshToken
		login.IdToken = *authResp.AuthenticationResult.IdToken
		login.ExpiresIn = int(*authResp.AuthenticationResult.ExpiresIn)
		login.Status = "success"
	}

	logrus.WithField("login", login).Info("User authenticated successfully \n", loginCredential.Email)
	return login, nil
}

func (a AuthenticationControllerImpl) LogoutUser(accessToken string) error {
	input := &cognitoidentityprovider.GlobalSignOutInput{
		AccessToken: aws.String(accessToken),
	}

	_, err := cognitoClient.GlobalSignOut(input)
	if err != nil {
		return err
	}

	fmt.Printf("User %s logged out successfully\n", accessToken)
	return nil
}

// GetUserIno
func (a AuthenticationControllerImpl) GetUserInfo(accessToken string) (*cognitoidentityprovider.GetUserOutput, error) {
	input := &cognitoidentityprovider.GetUserInput{
		AccessToken: aws.String(accessToken),
	}

	userInfo, err := cognitoClient.GetUser(input)
	if err != nil {
		return nil, err
	}

	logrus.WithField("accessToken", accessToken).Info("User authenticated successfully")
	return userInfo, nil
}

func (a AuthenticationControllerImpl) RegisterUser(signUpRequest dto.SignUpRequest) error {

	secretHash := generateSecretHash(clientSecret, signUpRequest.Email, clientID)

	input := &cognitoidentityprovider.SignUpInput{
		ClientId:   aws.String(clientID),
		Username:   aws.String(signUpRequest.Email),
		Password:   aws.String(signUpRequest.Password),
		SecretHash: aws.String(secretHash),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(signUpRequest.Email),
			},
			{
				Name:  aws.String("preferred_username"),
				Value: aws.String(signUpRequest.Email),
			},
			{
				Name:  aws.String("custom:company"),
				Value: aws.String(signUpRequest.CompanyName),
			},
		},
	}

	result, err := cognitoClient.SignUp(input)
	if err != nil {
		return err
	}

	fmt.Printf("User %s registered successfully\n", *result.UserSub)

	return nil
}

func (a AuthenticationControllerImpl) ConfirmUser(userInfo dto.ConfirmUser) error {

	secretHash := generateSecretHash(clientSecret, userInfo.Email, clientID)

	input := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(clientID),
		Username:         aws.String(userInfo.Email),
		ConfirmationCode: aws.String(userInfo.ConfirmationCode),
		SecretHash:       aws.String(secretHash),
	}

	_, err := cognitoClient.ConfirmSignUp(input)
	if err != nil {
		return err
	}

	fmt.Printf("User %s confirmed successfully\n", userInfo.Email)

	return nil
}

func (a AuthenticationControllerImpl) ResendChallengeCode(resendRequest dto.SignUpResendRequest) error {

	secretHash := generateSecretHash(clientSecret, resendRequest.Email, clientID)

	input := &cognitoidentityprovider.ResendConfirmationCodeInput{
		ClientId:   aws.String(clientID),
		Username:   aws.String(resendRequest.Email),
		SecretHash: aws.String(secretHash),
	}

	_, err := cognitoClient.ResendConfirmationCode(input)
	if err != nil {
		return err
	}

	fmt.Printf("Confirmation code resent to %s\n", resendRequest.Email)

	return nil
}

func generateSecretHash(clientSecret, userName, clientID string) string {
	key := []byte(clientSecret)
	message := userName + clientID
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(message))
	hash := mac.Sum(nil)
	secretHash := base64.StdEncoding.EncodeToString(hash)
	return secretHash
}

func setupCognitoClient() *cognitoidentityprovider.CognitoIdentityProvider {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	if err != nil {
		log.Fatalf("failed to create session, %v", err)
	}

	// Create a Cognito Identity Provider client
	cognitoClient := cognitoidentityprovider.New(sess)
	return cognitoClient
}
