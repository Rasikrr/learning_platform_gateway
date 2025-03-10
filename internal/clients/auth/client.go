package auth

import (
	"context"
	"github.com/Rasikrr/learning_platform_core/grpc"
	"github.com/Rasikrr/learning_platform_core/http/session"
	"github.com/Rasikrr/learning_platform_gateway/internal/domain/entity"
	pb "github.com/Rasikrr/learning_platform_gateway/pkg/deps/api/proto/auth"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	Register(ctx context.Context, email, password, passwordConfirm string) error
	ConfirmRegister(ctx context.Context, email, code string) (*entity.Auth, error)
	ConfirmAdminRegister(ctx context.Context, email, code string) (*entity.Auth, error)
	Login(ctx context.Context, email, password string) (*entity.Auth, error)
	ResetPassword(ctx context.Context, email, password, passwordConfirm string) error
	ConfirmResetPassword(ctx context.Context, email, code string) error
	CheckToken(ctx context.Context, token string) (*session.Session, error)
	RefreshToken(ctx context.Context, token string) (*entity.Auth, error)
}

type client struct {
	client pb.AuthClient
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
		client: pb.NewAuthClient(conn),
	}, nil
}

func (c *client) Register(ctx context.Context, email, password, passwordConfirm string) error {
	_, err := c.client.Register(ctx, &pb.RegisterRequest{
		Email:                email,
		Password:             password,
		PasswordConfirmation: passwordConfirm,
	})
	return err
}

func (c *client) ConfirmRegister(ctx context.Context, email, code string) (*entity.Auth, error) {
	reply, err := c.client.ConfirmRegister(ctx, &pb.ConfirmRegisterRequest{
		Email: email,
		Code:  code,
	})
	if err != nil {
		return nil, err
	}
	return convertAuth(reply)
}

func (c *client) ConfirmAdminRegister(ctx context.Context, email, code string) (*entity.Auth, error) {
	reply, err := c.client.ConfirmAdminRegister(ctx, &pb.ConfirmRegisterRequest{
		Email: email,
		Code:  code,
	})
	if err != nil {
		return nil, err
	}
	return convertAuth(reply)
}

func (c *client) Login(ctx context.Context, email, password string) (*entity.Auth, error) {
	reply, err := c.client.Login(ctx, &pb.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return convertAuth(reply)
}

func (c *client) ResetPassword(ctx context.Context, email, password, passwordConfirm string) error {
	_, err := c.client.ResetPassword(ctx, &pb.ResetPasswordRequest{
		Email:                email,
		Password:             password,
		PasswordConfirmation: passwordConfirm,
	})
	return err
}

func (c *client) ConfirmResetPassword(ctx context.Context, email, code string) error {
	_, err := c.client.ConfirmResetPassword(ctx, &pb.ConfirmResetPasswordRequest{
		Email: email,
		Code:  code,
	})
	return err
}

func (c *client) CheckToken(ctx context.Context, token string) (*session.Session, error) {
	reply, err := c.client.CheckToken(ctx, &pb.CheckTokenRequest{
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	return convertSession(reply.Session)
}

func (c *client) RefreshToken(ctx context.Context, refreshToken string) (*entity.Auth, error) {
	reply, err := c.client.RefreshToken(ctx, &pb.RefreshTokenRequest{
		RefreshToken: refreshToken,
	})
	if err != nil {
		return nil, err
	}
	return convertAuth(reply)
}
