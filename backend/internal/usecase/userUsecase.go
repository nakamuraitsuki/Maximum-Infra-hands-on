package usecase

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"time"

	"example.com/infrahandson/internal/domain/entity"
	"example.com/infrahandson/internal/domain/repository"
	"example.com/infrahandson/internal/domain/service"
	"example.com/infrahandson/internal/interface/adapter"
	"example.com/infrahandson/internal/interface/factory"
)

// UserUseCaseInterface: ユーザーに関するユースケースを管理するインターフェース
type UserUseCaseInterface interface {
	SignUp(ctx context.Context, req SignUpRequest) (SignUpResponse, error)
	AuthenticateUser(ctx context.Context, req AuthenticateUserRequest) (AuthenticateUserResponse, error)
	GetUserByID(ctx context.Context, id entity.UserID) (*entity.User, error)
	SaveUserIcon(ctx context.Context, fh *multipart.FileHeader, id entity.UserID) error
	GetUserIconPath(ctx context.Context, id entity.UserID) (path string, err error)
}

type UserUseCase struct {
	userRepo      repository.UserRepository
	hasher        adapter.HasherAdapter
	tokenSvc      adapter.TokenServiceAdapter
	iconSvc       service.IconStoreService
	userIDFactory factory.UserIDFactory
}

type NewUserUseCaseParams struct {
	UserRepo      repository.UserRepository
	Hasher        adapter.HasherAdapter
	TokenSvc      adapter.TokenServiceAdapter
	IconSvc       service.IconStoreService
	UserIDFactory factory.UserIDFactory
}

func NewUserUseCase(p NewUserUseCaseParams) UserUseCaseInterface {
	return &UserUseCase{
		userRepo:      p.UserRepo,
		hasher:        p.Hasher,
		tokenSvc:      p.TokenSvc,
		iconSvc:       p.IconSvc,
		userIDFactory: p.UserIDFactory,
	}
}

// SignUpRequest構造体: サインアップリクエスト
type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignUpResponse構造体: サインアップレスポンス
type SignUpResponse struct {
	User *entity.User
}

// SignUp ユーザー登録
func (u *UserUseCase) SignUp(ctx context.Context, req SignUpRequest) (SignUpResponse, error) {
	hashedPassword, err := u.hasher.HashPassword(req.Password)
	if err != nil {
		return SignUpResponse{nil}, err
	}

	id, err := u.userIDFactory.NewUserID()
	if err != nil {
		return SignUpResponse{nil}, err
	}

	userParams := entity.UserParams{
		ID:         id,
		Name:       req.Name,
		Email:      req.Email,
		PasswdHash: hashedPassword,
		CreatedAt:  time.Now(),
		UpdatedAt:  nil,
	}

	user := entity.NewUser(userParams)

	res, err := u.userRepo.SaveUser(ctx, user)
	if err != nil {
		return SignUpResponse{nil}, err
	}

	return SignUpResponse{User: res}, nil
}

// AuthenticateUserRequest構造体: 認証リクエスト
type AuthenticateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthenticateUserResponse構造体: 認証レスポンス
type AuthenticateUserResponse struct {
	token *string
	exp   *int
}

// IsTokenNil: トークンがnilかどうかを判定
func (res *AuthenticateUserResponse) IsTokenNil() bool {
	return res.token == nil
}

// GetToken: トークンを取得（nilの場合は空文字を返す）
func (res *AuthenticateUserResponse) GetToken() string {
	if res.token == nil {
		return ""
	}
	return *res.token
}

func (res *AuthenticateUserResponse) GetExp() int {
	if res.exp == nil {
		return 0
	}
	return *res.exp
}

// 外部でのテストのためのセッター
func (res *AuthenticateUserResponse) SetToken(token string) {
	res.token = &token
}

// AuthenticateUser ユーザー認証
func (u *UserUseCase) AuthenticateUser(ctx context.Context, req AuthenticateUserRequest) (AuthenticateUserResponse, error) {
	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return AuthenticateUserResponse{token: nil}, err
	}

	ok, err := u.hasher.ComparePassword(user.GetPasswdHash(), req.Password)
	if err != nil {
		return AuthenticateUserResponse{token: nil}, err
	}

	if !ok {
		return AuthenticateUserResponse{token: nil}, errors.New("password mismatch")
	}

	res, err := u.tokenSvc.GenerateToken(user.GetID())
	if err != nil {
		return AuthenticateUserResponse{token: nil}, err
	}

	exp, err := u.tokenSvc.GetExpireAt(res)
	if err != nil {
		return AuthenticateUserResponse{token: nil}, err
	}

	return AuthenticateUserResponse{
		token: &res,
		exp:   &exp,
	}, nil
}

func (u *UserUseCase) GetUserByID(ctx context.Context, id entity.UserID) (*entity.User, error) {
	user, err := u.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) SaveUserIcon(ctx context.Context, fh *multipart.FileHeader, id entity.UserID) error {
	// ファイルを開く
	file, err := fh.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	// []byte に読み込む
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// IconData を作成
	iconData := service.NewIconData(data, int64(len(data)), fh.Header.Get("Content-Type"))

	// 保存処理を呼ぶ
	return u.iconSvc.SaveIcon(ctx, iconData, id)
}

func (u *UserUseCase) GetUserIconPath(ctx context.Context, id entity.UserID) (path string, err error) {
	path, err = u.iconSvc.GetIconPath(ctx, id)
	if err != nil {
		return "", err
	}
	return path, nil
}
