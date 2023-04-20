package test

import (
	"context"
	"fmt"

	"github.com/TheLazarusNetwork/erebrus/util/pkg/auth"
)

type TestService struct {
	UnimplementedTestServiceServer
}

func (s *TestService) TestFn(ctx context.Context, request *Empty) error {
	// secretKey := gopaseto.NewV4AsymmetricSecretKey()
	// PublicKey := secretKey.Public()
	fmt.Println(auth.Getpublickey().ExportHex())

	return nil
}
