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
	"github.com/sirupsen/logrus"
	_ "golang.org/x/net/context"
	"log"
)

type AuthenticationController interface {
	RegisterUser(loginCredential dto.Credentials) error
	LoginUser(loginCredential dto.Credentials) (*dto.LoginResponse, error)
	ConfirmUser(userInfo dto.ConfirmUser) error
	ResendChallengeCode(resendRequest dto.SignUpResendRequest) error
}

type AuthenticationControllerImpl struct {
}

var (
	awsRegion = "ap-southeast-1" // Your AWS Region
	//userPoolID                  = "ap-southeast-1_wnUcfMgqN"  // Your Cognito User Pool ID
	clientID                    = "og5uq3m2bvhfbghf3jd2q14jm" // Your Cognito App Client ID
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
		logrus.WithField("login", login).Info("User %s authentication failed %s\n", loginCredential.Email)
		login.Status = *authResp.ChallengeName
		//handleChallenge(cognitoClient, clientID, clientSecret, loginCredential.Email, *authResp.Session, *authResp.ChallengeName)
	} else {
		login.AccessToken = *authResp.AuthenticationResult.AccessToken
		login.RefreshToken = *authResp.AuthenticationResult.RefreshToken
		login.IdToken = *authResp.AuthenticationResult.IdToken
		login.ExpiresIn = int(*authResp.AuthenticationResult.ExpiresIn)
		login.Status = "success"
	}

	logrus.WithField("login", login).Info("User %s authenticated successfully %s\n", loginCredential.Email)
	return login, nil
}

//func handleChallenge(client *cognitoidentityprovider, clientID, clientSecret, userName, session, challengeName string) {
//	switch challengeName {
//	case "SMS_MFA":
//		var mfaCode string
//		fmt.Print("Enter SMS MFA code: ")
//		//fmt.Scan(&mfaCode)
//
//		respondToChallenge(client, clientID, clientSecret, userName, session, challengeName, "SMS_MFA_CODE", mfaCode)
//	case "SOFTWARE_TOKEN_MFA":
//		var mfaCode string
//		fmt.Print("Enter TOTP MFA code: ")
//		//fmt.Scan(&mfaCode)
//
//		respondToChallenge(client, clientID, clientSecret, userName, session, challengeName, "SOFTWARE_TOKEN_MFA_CODE", mfaCode)
//	case "MFA_SETUP":
//		setupMFA(client, clientID, clientSecret, userName, session)
//	default:
//		fmt.Println("Unknown challenge:", challengeName)
//	}
//}
//
//func respondToChallenge(client *cognitoidentityprovider.Client, clientID, clientSecret, userName, session, challengeName, mfaType, mfaCode string) {
//	input := &cognitoidentityprovider.RespondToAuthChallengeInput{
//		ChallengeName: types.ChallengeNameType(challengeName),
//		ClientId:      &clientID,
//		Session:       &session,
//		ChallengeResponses: map[string]string{
//			"USERNAME":    userName,
//			"SECRET_HASH": generateSecretHash(clientSecret, userName, clientID),
//			mfaType:       mfaCode,
//		},
//	}
//
//	resp, err := client.RespondToAuthChallenge(context.TODO(), input)
//	if err != nil {
//		log.Fatalf("Error responding to challenge: %v", err)
//	}
//
//	fmt.Println("Authentication successful:", resp.AuthenticationResult)
//}
//
//func setupMFA(client *cognitoidentityprovider.Client, clientID, clientSecret, userName, session string) {
//	input := &cognitoidentityprovider.AssociateSoftwareTokenInput{
//		Session: &session,
//	}
//
//	resp, err := client.AssociateSoftwareToken(context.TODO(), input)
//	if err != nil {
//		log.Fatalf("Error setting up MFA: %v", err)
//	}
//
//	fmt.Printf("Set up MFA with TOTP: %v\n", resp)
//
//	var mfaCode string
//	fmt.Print("Enter the TOTP code from your authenticator app: ")
//	fmt.Scan(&mfaCode)
//
//	verifyMFA(client, clientID, clientSecret, userName, session, resp.SecretCode, mfaCode)
//}
//
//func verifyMFA(client *cognitoidentityprovider.Client, clientID, clientSecret, userName, session, secretCode, mfaCode string) {
//	input := &cognitoidentityprovider.VerifySoftwareTokenInput{
//		AccessToken:        &session,
//		FriendlyDeviceName: aws.String("MyDevice"),
//		UserCode:           &mfaCode,
//	}
//
//	resp, err := client.VerifySoftwareToken(context.TODO(), input)
//	if err != nil {
//		log.Fatalf("Error verifying MFA: %v", err)
//	}
//
//	fmt.Printf("MFA verified: %v\n", resp)
//}

func (a AuthenticationControllerImpl) RegisterUser(loginCredential dto.Credentials) error {

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
			{
				Name:  aws.String("preferred_username"),
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
