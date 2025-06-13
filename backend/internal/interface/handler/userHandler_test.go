package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"example.com/infrahandson/internal/domain/entity"
	"example.com/infrahandson/internal/interface/handler"
	mock_adapter "example.com/infrahandson/test/mocks/interface/adapter"
	mock_factory "example.com/infrahandson/test/mocks/interface/factory"
	mock_user "example.com/infrahandson/test/mocks/usecase/user"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetMe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userParams := entity.UserParams{
		ID:         "mockUserID",
		Name:       "Test User",
		Email:      "test",
		PasswdHash: "hashedPassword",
		CreatedAt:  time.Now(),
		UpdatedAt:  nil,
	}

	user := entity.NewUser(userParams)

	mockUserUseCase := mock_user.NewMockUserUseCaseInterface(ctrl)
	mockUserIDFactory := mock_factory.NewMockUserIDFactory(ctrl)
	mockLogger := mock_adapter.NewMockLoggerAdapter(ctrl)

	handler := handler.NewUserHandler(handler.NewUserHandlerParams{
		UserUseCase:   mockUserUseCase,
		UserIDFactory: mockUserIDFactory,
		Logger:        mockLogger,
	})

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/me", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user_id", "mockUserID")

	mockLogger.EXPECT().
		Info(gomock.Any(), gomock.Any()).
		AnyTimes() // ロガーは何回呼ばれてもいい（呼ばれなくても怒らない）設定

	mockUserUseCase.EXPECT().GetUserByID(context.Background(), entity.UserID("mockUserID")).Return(user, nil)

	if assert.NoError(t, handler.GetMe(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "{\"id\":\"mockUserID\",\"name\":\"Test User\",\"email\":\"test\"}")
	}
}
