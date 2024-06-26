// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dtos "github.com/isido5ik/StoryPublishingPlatform/dtos"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// AddComment mocks base method.
func (m *MockUsecase) AddComment(userId, postId int, comment dtos.AddCommentInput) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddComment", userId, postId, comment)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddComment indicates an expected call of AddComment.
func (mr *MockUsecaseMockRecorder) AddComment(userId, postId, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddComment", reflect.TypeOf((*MockUsecase)(nil).AddComment), userId, postId, comment)
}

// CreateStory mocks base method.
func (m *MockUsecase) CreateStory(story dtos.AddPostInput, userId int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStory", story, userId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStory indicates an expected call of CreateStory.
func (mr *MockUsecaseMockRecorder) CreateStory(story, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStory", reflect.TypeOf((*MockUsecase)(nil).CreateStory), story, userId)
}

// CreateUserAsClient mocks base method.
func (m *MockUsecase) CreateUserAsClient(input dtos.SignUpInput) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserAsClient", input)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserAsClient indicates an expected call of CreateUserAsClient.
func (mr *MockUsecaseMockRecorder) CreateUserAsClient(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserAsClient", reflect.TypeOf((*MockUsecase)(nil).CreateUserAsClient), input)
}

// DeleteComment mocks base method.
func (m *MockUsecase) DeleteComment(userId, postId, commentId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", userId, postId, commentId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockUsecaseMockRecorder) DeleteComment(userId, postId, commentId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockUsecase)(nil).DeleteComment), userId, postId, commentId)
}

// DeleteStory mocks base method.
func (m *MockUsecase) DeleteStory(postId, userId int, role string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStory", postId, userId, role)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStory indicates an expected call of DeleteStory.
func (mr *MockUsecaseMockRecorder) DeleteStory(postId, userId, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStory", reflect.TypeOf((*MockUsecase)(nil).DeleteStory), postId, userId, role)
}

// DeleteUser mocks base method.
func (m *MockUsecase) DeleteUser(userId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUsecaseMockRecorder) DeleteUser(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUsecase)(nil).DeleteUser), userId)
}

// GenerateToken mocks base method.
func (m *MockUsecase) GenerateToken(username, password string) (string, []dtos.Roles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", username, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]dtos.Roles)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockUsecaseMockRecorder) GenerateToken(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockUsecase)(nil).GenerateToken), username, password)
}

// GetStories mocks base method.
func (m *MockUsecase) GetStories(pagination dtos.PaginationParams) ([]dtos.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStories", pagination)
	ret0, _ := ret[0].([]dtos.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStories indicates an expected call of GetStories.
func (mr *MockUsecaseMockRecorder) GetStories(pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStories", reflect.TypeOf((*MockUsecase)(nil).GetStories), pagination)
}

// GetStoriesByCategory mocks base method.
func (m *MockUsecase) GetStoriesByCategory(pagination dtos.PaginationParams, categoryId int) ([]dtos.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoriesByCategory", pagination, categoryId)
	ret0, _ := ret[0].([]dtos.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStoriesByCategory indicates an expected call of GetStoriesByCategory.
func (mr *MockUsecaseMockRecorder) GetStoriesByCategory(pagination, categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoriesByCategory", reflect.TypeOf((*MockUsecase)(nil).GetStoriesByCategory), pagination, categoryId)
}

// GetStory mocks base method.
func (m *MockUsecase) GetStory(postId int) (dtos.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStory", postId)
	ret0, _ := ret[0].(dtos.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStory indicates an expected call of GetStory.
func (mr *MockUsecaseMockRecorder) GetStory(postId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStory", reflect.TypeOf((*MockUsecase)(nil).GetStory), postId)
}

// GetUserById mocks base method.
func (m *MockUsecase) GetUserById(userId int) (dtos.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", userId)
	ret0, _ := ret[0].(dtos.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUsecaseMockRecorder) GetUserById(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUsecase)(nil).GetUserById), userId)
}

// GetUsers mocks base method.
func (m *MockUsecase) GetUsers() ([]dtos.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]dtos.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUsecaseMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUsecase)(nil).GetUsers))
}

// GetUsersStories mocks base method.
func (m *MockUsecase) GetUsersStories(userId int) (string, []dtos.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersStories", userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]dtos.Post)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUsersStories indicates an expected call of GetUsersStories.
func (mr *MockUsecaseMockRecorder) GetUsersStories(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersStories", reflect.TypeOf((*MockUsecase)(nil).GetUsersStories), userId)
}

// Like mocks base method.
func (m *MockUsecase) Like(userId, postId int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Like", userId, postId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Like indicates an expected call of Like.
func (mr *MockUsecaseMockRecorder) Like(userId, postId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Like", reflect.TypeOf((*MockUsecase)(nil).Like), userId, postId)
}

// ParseToken mocks base method.
func (m *MockUsecase) ParseToken(token string) (int, []dtos.Roles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]dtos.Roles)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockUsecaseMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockUsecase)(nil).ParseToken), token)
}

// RemoveLike mocks base method.
func (m *MockUsecase) RemoveLike(userId, postId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveLike", userId, postId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveLike indicates an expected call of RemoveLike.
func (mr *MockUsecaseMockRecorder) RemoveLike(userId, postId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLike", reflect.TypeOf((*MockUsecase)(nil).RemoveLike), userId, postId)
}

// ReplyToComment mocks base method.
func (m *MockUsecase) ReplyToComment(userId, postId, parentId int, comment dtos.AddCommentInput) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplyToComment", userId, postId, parentId, comment)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReplyToComment indicates an expected call of ReplyToComment.
func (mr *MockUsecaseMockRecorder) ReplyToComment(userId, postId, parentId, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplyToComment", reflect.TypeOf((*MockUsecase)(nil).ReplyToComment), userId, postId, parentId, comment)
}

// UpdateComment mocks base method.
func (m *MockUsecase) UpdateComment(userId, postId, commentId int, newComment dtos.UpdateCommentInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", userId, postId, commentId, newComment)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateComment indicates an expected call of UpdateComment.
func (mr *MockUsecaseMockRecorder) UpdateComment(userId, postId, commentId, newComment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockUsecase)(nil).UpdateComment), userId, postId, commentId, newComment)
}

// UpdateStory mocks base method.
func (m *MockUsecase) UpdateStory(postId, userId int, role string, input dtos.UpdateStoryInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStory", postId, userId, role, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStory indicates an expected call of UpdateStory.
func (mr *MockUsecaseMockRecorder) UpdateStory(postId, userId, role, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStory", reflect.TypeOf((*MockUsecase)(nil).UpdateStory), postId, userId, role, input)
}
