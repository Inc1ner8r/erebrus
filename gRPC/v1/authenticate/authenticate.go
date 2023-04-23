package authenticate

import (
	"context"
	"fmt"
	"os"

	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/authenticate/challengeid"
	"github.com/TheLazarusNetwork/erebrus/util/pkg/auth"
	"github.com/TheLazarusNetwork/erebrus/util/pkg/claims"
	"github.com/TheLazarusNetwork/erebrus/util/pkg/cryptosign"
	log "github.com/sirupsen/logrus"
)

type AuthenticateService struct {
	UnimplementedAuthenticateServiceServer
}

func (as *AuthenticateService) Authenticate(ctx context.Context, request *AuthenticateRequest) (*AuthenticatePayload, error) {
	userAuthEULA := os.Getenv("AUTH_EULA")
	message := userAuthEULA + request.ChallengeId
	walletAddress, isCorrect, err := cryptosign.CheckSign(request.Signature, request.ChallengeId, message)
	if err == cryptosign.ErrFlowIdNotFound {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("FlowId Not Found")
		errResponse := ErrAuthenticate(err.Error())
		return errResponse, err
	}
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to CheckSignature")
		errResponse := ErrAuthenticate(err.Error())
		return errResponse, err
	}
	if !isCorrect {
		fmt.Println(ErrAuthenticate("Forbidden"))
		errResponse := ErrAuthenticate(err.Error())
		return errResponse, err
	}
	customClaims := claims.New(walletAddress)
	pasetoToken, err := auth.GenerateTokenPaseto(customClaims)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to generate token")
		errResponse := ErrAuthenticate(err.Error())
		return errResponse, err
	}

	delete(challengeid.Data, request.ChallengeId)
	payload := &AuthenticatePayload{
		Status:  200,
		Success: true,
		Message: "Successfully Authenticated",
		Token:   pasetoToken,
	}
	return payload, nil

}

func ErrAuthenticate(errvalue string) *AuthenticatePayload {
	payload := &AuthenticatePayload{
		Status:  401,
		Success: false,
		Message: errvalue,
	}
	return payload
}
