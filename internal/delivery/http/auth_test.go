package http

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/isido5ik/StoryPublishingPlatform/dtos"
	mock_usecase "github.com/isido5ik/StoryPublishingPlatform/internal/usecase/mocks"
	"github.com/magiconair/properties/assert"
)

func TestHttp_signUp(t *testing.T) {
	type mockBehavior func(m *mock_usecase.MockUsecase, user dtos.SignUpInput)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            dtos.SignUpInput
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username":"test", "email":"test", "password":"test"}`,
			inputUser: dtos.SignUpInput{
				Username: "test",
				Email:    "test",
				Password: "test",
			},
			mockBehavior: func(m *mock_usecase.MockUsecase, user dtos.SignUpInput) {
				m.EXPECT().CreateUserAsClient(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"client_id":1}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"email":"", "password":"test"}`,
			inputUser:            dtos.SignUpInput{},
			mockBehavior:         func(m *mock_usecase.MockUsecase, user dtos.SignUpInput) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"username":"test", "email":"test", "password":"test"}`,
			inputUser: dtos.SignUpInput{
				Username: "test",
				Email:    "test",
				Password: "test",
			},
			mockBehavior: func(m *mock_usecase.MockUsecase, user dtos.SignUpInput) {
				m.EXPECT().CreateUserAsClient(user).Return(0, errors.New("Failed to create user as client"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"Failed to create user as client"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			usecase := mock_usecase.NewMockUsecase(c)
			test.mockBehavior(usecase, test.inputUser)

			handler := NewHandler(usecase)

			router := gin.New()
			router.POST("/api/auth/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/auth/sign-up", bytes.NewBufferString(test.inputBody))

			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}

}
