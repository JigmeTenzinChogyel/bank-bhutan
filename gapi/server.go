package gapi

import (
	"fmt"

	db "github.com/JigmeTenzinChogyel/bank-bhutan/db/sqlc"
	"github.com/JigmeTenzinChogyel/bank-bhutan/pb"
	"github.com/JigmeTenzinChogyel/bank-bhutan/token"
	"github.com/JigmeTenzinChogyel/bank-bhutan/util"
)

// Server servers HTTP request for our backend service
type Server struct {
	pb.UnimplementedBankBhutanServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// newServer creates a new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %W", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
