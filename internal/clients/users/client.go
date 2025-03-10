package users

import (
	"context"
	"github.com/Rasikrr/learning_platform_core/grpc"
	"github.com/Rasikrr/learning_platform_gateway/internal/domain/entity"
	pb "github.com/Rasikrr/learning_platform_gateway/pkg/deps/api/proto/users"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Client interface {
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}

type client struct {
	client pb.UsersClient
}

func NewClient(ctx context.Context, addr string) (Client, error) {
	conn, err := grpc.NewClient(
		ctx,
		addr,
		grpc2.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}
	return &client{
		client: pb.NewUsersClient(conn),
	}, nil
}

func (c *client) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	reply, err := c.client.GetByEmail(ctx, &pb.GetByEmailRequest{Email: email})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return convert(reply)
}
