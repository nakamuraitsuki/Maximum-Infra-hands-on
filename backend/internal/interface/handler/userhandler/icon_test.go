package userhandler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/infrahandson/internal/domain/entity"
	"example.com/infrahandson/internal/interface/handler/userhandler"
	"example.com/infrahandson/internal/usecase/usercase"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSaveUserIcon(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler, mockDeps, e := userhandler.NewTestUserHandler(ctrl)

	userID := entity.UserID("user-123")
	iconContent := []byte("dummy image data")
	contentType := "image/png"
	fh := userhandler.NewMultipartFileHeader(iconContent, contentType)

	t.Run("正常系: アイコン保存成功", func(t *testing.T) {
		mockDeps.UserUseCase.EXPECT().
			SaveUserIcon(gomock.Any(), gomock.Any(), userID).
			Return(nil)
		req := httptest.NewRequest(http.MethodPost, "/user/icon", fh)
		req.SetParamNames("user_id")
		req.SetParamValues(string(userID))

		rec := e.NewResponseRecorder()
		c := e.NewContext(req, rec)
		c.Set("user_id", string(userID))

		err := handler.SaveUserIcon(c)
		assert.NoError(t, err)
		assert.Equal(t, 200, rec.Code)
		assert.JSONEq(t, `{"message": "Icon saved successfully"}`, rec.Body.String())
	})
}

func TestGetUserIconPath(t *testing.T) {
	userID := entity.UserID("user-123")
	expectedPath := "/path/to/icon.png"

	t.Run("正常系: パス取得成功", func(t *testing.T) {
		mockSvc := &mockIconService{
			getIconPathFunc: func(ctx context.Context, id entity.UserID) (string, error) {
				assert.Equal(t, userID, id)
				return expectedPath, nil
			},
		}
		uc := &usercase.UserUseCase{IconSvc: mockSvc}
		path, err := uc.GetUserIconPath(context.Background(), userID)
		assert.NoError(t, err)
		assert.Equal(t, expectedPath, path)
	})

	t.Run("異常系: サービスでエラー", func(t *testing.T) {
		mockSvc := &mockIconService{
			getIconPathFunc: func(ctx context.Context, id entity.UserID) (string, error) {
				return "", errors.New("not found")
			},
		}
		uc := &usercase.UserUseCase{IconSvc: mockSvc}
		path, err := uc.GetUserIconPath(context.Background(), userID)
		assert.Error(t, err)
		assert.Empty(t, path)
	})
}
