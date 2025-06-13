package userhandler

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/textproto"

	"example.com/infrahandson/internal/infrastructure/validator"
	mock_adapter "example.com/infrahandson/test/mocks/interface/adapter"
	mock_factory "example.com/infrahandson/test/mocks/interface/factory"
	mock_usercase "example.com/infrahandson/test/mocks/usecase/usercase"
	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

// mockDeps は UserHandler のテストで使用する依存関係モックをまとめた構造体です
type mockDeps struct {
	UserUseCase   mock_usercase.MockUserUseCaseInterface
	UserIDFactory mock_factory.MockUserIDFactory
	Logger        mock_adapter.MockLoggerAdapter
}

// NewTestUserHandler ( ハンドラ, モック依存関係, Echoインスタンス ) を生成する
func NewTestUserHandler(
	ctrl *gomock.Controller,
) (*UserHandler, mockDeps, *echo.Echo) {
	mockUserUseCase := mock_usercase.NewMockUserUseCaseInterface(ctrl)
	mockUserIDFactory := mock_factory.NewMockUserIDFactory(ctrl)
	mockLogger := mock_adapter.NewMockLoggerAdapter(ctrl)
	params := NewUserHandlerParams{
		UserUseCase:   mockUserUseCase,
		UserIDFactory: mockUserIDFactory,
		Logger:        mockLogger,
	}
	handler := NewUserHandler(params)

	mockDeps := mockDeps{
		UserUseCase:   *mockUserUseCase,
		UserIDFactory: *mockUserIDFactory,
		Logger:        *mockLogger,
	}

	e := echo.New()
	e.Validator = validator.NewEchoValidator()

	return handler, mockDeps, e
}

// NewMultipartFileHeader は， icno_test.go で使用するダミーリクエストの multipart.FileHeader を生成するヘルパー関数です
func NewMultipartFileHeader(content []byte, contentType string) *multipart.FileHeader {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", `form-data; name="file"; filename="icon.png"`)
	h.Set("Content-Type", contentType)
	part, _ := writer.CreatePart(h)
	part.Write(content)
	writer.Close()

	// Parse the multipart body to get the FileHeader
	r := multipart.NewReader(bytes.NewReader(body.Bytes()), writer.Boundary())
	form, _ := r.ReadForm(1024 * 1024)
	files := form.File["file"]
	if len(files) > 0 {
		return files[0]
	}
	return nil
}

// BadFile は、icon_test.go 用の不正なファイルを模倣するための構造体です
type BadFile struct{}

func (b *BadFile) Read(p []byte) (n int, err error)             { return 0, io.ErrUnexpectedEOF }
func (b *BadFile) Close() error                                 { return nil }
func (b *BadFile) Seek(offset int64, whence int) (int64, error) { return 0, nil }
