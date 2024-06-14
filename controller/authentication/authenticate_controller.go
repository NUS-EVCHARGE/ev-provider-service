package authentication

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"log"
)

type AuthenticationController interface {
	RegisterUser(loginCredential dto.Credentials) error
	LoginUser(loginCredential dto.Credentials) error
	ConfirmUser(userInfo dto.ConfirmUser) error
	ResendChallengeCode(email string) error
}

type AuthenticationControllerImpl struct {
}

var (
	awsRegion                   = "ap-southeast-1"            // Your AWS Region
	userPoolID                  = "ap-southeast-1_wnUcfMgqN"  // Your Cognito User Pool ID
	clientID                    = "og5uq3m2bvhfbghf3jd2q14jm" // Your Cognito App Client ID
	clientSecret                = "16q37emcuik0cbfffo534lsqo2kck4fisjp7gnkpmbil2br6bho"
	AuthenticationControllerObj AuthenticationController
)

func NewAuthenticationController() {
	AuthenticationControllerObj = &AuthenticationControllerImpl{}
}

func (a AuthenticationControllerImpl) LoginUser(loginCredential dto.Credentials) error {

	cognitoClient := setupCognitoClient()

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

	authResp, err := cognitoClient.InitiateAuth(input)
	if err != nil {
		return err
	}

	fmt.Printf("User %s authenticated successfully %s\n", loginCredential.Email, authResp)
	return nil
}

func (a AuthenticationControllerImpl) RegisterUser(loginCredential dto.Credentials) error {

	cognitoClient := setupCognitoClient()

	secretHash := generateSecretHash(clientSecret, loginCredential.Email, clientID)

	input := &cognitoidentityprovider.SignUpInput{
		ClientId:   aws.String(clientID),
		Username:   aws.String(loginCredential.Email),
		Password:   aws.String(loginCredential.Password),
		SecretHash: aws.String(secretHash),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(loginCredential.Email),
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

	cognitoClient := setupCognitoClient()
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

func (a AuthenticationControllerImpl) ResendChallengeCode(email string) error {

	cognitoClient := setupCognitoClient()
	secretHash := generateSecretHash(clientSecret, email, clientID)

	input := &cognitoidentityprovider.ResendConfirmationCodeInput{
		ClientId:   aws.String(clientID),
		Username:   aws.String(email),
		SecretHash: aws.String(secretHash),
	}

	_, err := cognitoClient.ResendConfirmationCode(input)
	if err != nil {
		return err
	}

	fmt.Printf("Confirmation code resent to %s\n", email)

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
