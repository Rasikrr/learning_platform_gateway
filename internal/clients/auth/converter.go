package auth

import (
	coreEnum "github.com/Rasikrr/learning_platform_core/enum"
	"github.com/Rasikrr/learning_platform_core/http/session"
	"github.com/Rasikrr/learning_platform_gateway/internal/domain/entity"
	pb "github.com/Rasikrr/learning_platform_gateway/pkg/deps/api/proto/auth"
)

func convertAuth(auth *pb.AuthResponse) (*entity.Auth, error) {
	return &entity.Auth{
		AccessToken:  auth.AccessToken,
		RefreshToken: auth.RefreshToken,
	}, nil
}

func convertSession(sesPb *pb.Session) (*session.Session, error) {
	ses := session.Session{}
	ses.SetEmail(sesPb.Email)
	role, err := coreEnum.AccountRoleString(sesPb.Role)
	if err != nil {
		return nil, err
	}
	ses.SetRole(role)
	for k, v := range sesPb.Claims {
		ses.SetClaim(k, v)
	}
	return &ses, nil
}
