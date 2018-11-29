// Code generated by protoc-gen-defaults. DO NOT EDIT.

package scopesval

import "fmt"
import "context"
import "google.golang.org/grpc/codes"
import "google.golang.org/grpc/status"
import "go.uber.org/fx"
import "go.appointy.com/google/pb/rights"
import "go.appointy.com/google/userinfo"
import "github.com/ayushbpl10/protoc-gen-rights/example/rights"

import "github.com/golang/protobuf/ptypes/timestamp"

import "github.com/golang/protobuf/ptypes/empty"

import "github.com/golang/protobuf/ptypes/empty"

import "github.com/ayushbpl10/protoc-gen-scopes/example/example/pb"

type ScopesUsersServer struct {
	pb.UsersServer
	rightsCli rights.ScopeValidatorsClient
	user      right.UserIDer
}

func init() {
	options = append(options, fx.Provide(NewScopesUsersClient))
}

type ScopesUsersClientResult struct {
	fx.Out
	UsersClient pb.AcceptancesClient `name:"r"`
}

func NewScopesUsersClient(c rights.ScopeValidatorsClient, s pb.UsersServer) ScopesUsersClientResult {
	return ScopesUsersClientResult{UsersClient: pb.NewLocalUsersClient(NewScopesUsersServer(c, s))}
}
func NewScopesUsersServer(c rights.ScopeValidatorsClient, s pb.UsersServer, u right.UserIDer) pb.UsersServer {
	return &ScopesUsersServer{
		s,
		c,
		u,
	}
}

func (s *ScopesUsersServer) AddUser(ctx context.Context, rightsvar *pb.User) (*empty.Empty, error) {

	ResourcePathOR := make([]string, 0)

	ResourcePathOR = append(ResourcePath,

		"/users/{id}/cards.read/{blocked}",

		"/users/{id}/cards/user.write",
	)

	res, err := s.rightsCli.IsValid(ctx, &rights.IsValidReq{
		ResourcePathOR: ResourcePath,
		UserId:         s.user.UserID(ctx),
		ModuleName:     "Users",
	})
	if err != nil {
		return nil, err
	}

	if !res.IsValid {
		return nil, status.Errorf(codes.PermissionDenied, res.Reason)
	}
	return s.UsersServer.AddUser(ctx, rightsvar)
}

func (s *ScopesUsersServer) GetUser(ctx context.Context, rightsvar *pb.GetUserReq) (*pb.User, error) {

	ResourcePathOR := make([]string, 0)

	ResourcePathOR = append(ResourcePath,

		"/{user_email.email}/users/{user_id}/cards/{tent_id.tent}/email/{user_email.email.checks.check.check_id.val_id}",

		"/users/{user_id}/cards/{tent_id.tent}/ex.write",
	)

	res, err := s.rightsCli.IsValid(ctx, &rights.IsValidReq{
		ResourcePathOR: ResourcePath,
		UserId:         s.user.UserID(ctx),
		ModuleName:     "Users",
	})
	if err != nil {
		return nil, err
	}

	if !res.IsValid {
		return nil, status.Errorf(codes.PermissionDenied, res.Reason)
	}
	return s.UsersServer.GetUser(ctx, rightsvar)
}

func (s *ScopesUsersServer) UpdateUser(ctx context.Context, rightsvar *pb.UpdateUserReq) (*empty.Empty, error) {

	ResourcePathOR := make([]string, 0)

	ResourcePathOR = append(ResourcePath,

		"/users/{email_ids.emails}/cards.read/",

		"/users/{id}/cards/user.write",
	)

	res, err := s.rightsCli.IsValid(ctx, &rights.IsValidReq{
		ResourcePathOR: ResourcePath,
		UserId:         s.user.UserID(ctx),
		ModuleName:     "Users",
	})
	if err != nil {
		return nil, err
	}

	if !res.IsValid {
		return nil, status.Errorf(codes.PermissionDenied, res.Reason)
	}
	return s.UsersServer.UpdateUser(ctx, rightsvar)
}