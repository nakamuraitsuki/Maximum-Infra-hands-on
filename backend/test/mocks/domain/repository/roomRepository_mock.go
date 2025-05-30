// Code generated by MockGen. DO NOT EDIT.
// Source: /home/nakamura/program/Maximum/infra/Maximum-Infra-hands-on/backend/internal/domain/repository/roomRepository.go
//
// Generated by this command:
//
//	mockgen -source=/home/nakamura/program/Maximum/infra/Maximum-Infra-hands-on/backend/internal/domain/repository/roomRepository.go -destination=test/mocks/domain/repository/roomRepository_mock.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	entity "example.com/infrahandson/internal/domain/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockRoomRepository is a mock of RoomRepository interface.
type MockRoomRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRoomRepositoryMockRecorder
	isgomock struct{}
}

// MockRoomRepositoryMockRecorder is the mock recorder for MockRoomRepository.
type MockRoomRepositoryMockRecorder struct {
	mock *MockRoomRepository
}

// NewMockRoomRepository creates a new mock instance.
func NewMockRoomRepository(ctrl *gomock.Controller) *MockRoomRepository {
	mock := &MockRoomRepository{ctrl: ctrl}
	mock.recorder = &MockRoomRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomRepository) EXPECT() *MockRoomRepositoryMockRecorder {
	return m.recorder
}

// AddMemberToRoom mocks base method.
func (m *MockRoomRepository) AddMemberToRoom(arg0 context.Context, arg1 entity.RoomID, arg2 entity.UserID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMemberToRoom", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMemberToRoom indicates an expected call of AddMemberToRoom.
func (mr *MockRoomRepositoryMockRecorder) AddMemberToRoom(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMemberToRoom", reflect.TypeOf((*MockRoomRepository)(nil).AddMemberToRoom), arg0, arg1, arg2)
}

// DeleteRoom mocks base method.
func (m *MockRoomRepository) DeleteRoom(arg0 context.Context, arg1 entity.RoomID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoom", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRoom indicates an expected call of DeleteRoom.
func (mr *MockRoomRepositoryMockRecorder) DeleteRoom(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoom", reflect.TypeOf((*MockRoomRepository)(nil).DeleteRoom), arg0, arg1)
}

// GetAllRooms mocks base method.
func (m *MockRoomRepository) GetAllRooms(arg0 context.Context) ([]*entity.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRooms", arg0)
	ret0, _ := ret[0].([]*entity.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRooms indicates an expected call of GetAllRooms.
func (mr *MockRoomRepositoryMockRecorder) GetAllRooms(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRooms", reflect.TypeOf((*MockRoomRepository)(nil).GetAllRooms), arg0)
}

// GetRoomByID mocks base method.
func (m *MockRoomRepository) GetRoomByID(ctx context.Context, id entity.RoomID) (*entity.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomByID", ctx, id)
	ret0, _ := ret[0].(*entity.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomByID indicates an expected call of GetRoomByID.
func (mr *MockRoomRepositoryMockRecorder) GetRoomByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomByID", reflect.TypeOf((*MockRoomRepository)(nil).GetRoomByID), ctx, id)
}

// GetRoomByNameLike mocks base method.
func (m *MockRoomRepository) GetRoomByNameLike(ctx context.Context, name string) ([]*entity.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomByNameLike", ctx, name)
	ret0, _ := ret[0].([]*entity.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomByNameLike indicates an expected call of GetRoomByNameLike.
func (mr *MockRoomRepositoryMockRecorder) GetRoomByNameLike(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomByNameLike", reflect.TypeOf((*MockRoomRepository)(nil).GetRoomByNameLike), ctx, name)
}

// GetUsersInRoom mocks base method.
func (m *MockRoomRepository) GetUsersInRoom(arg0 context.Context, arg1 entity.RoomID) ([]*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersInRoom", arg0, arg1)
	ret0, _ := ret[0].([]*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersInRoom indicates an expected call of GetUsersInRoom.
func (mr *MockRoomRepositoryMockRecorder) GetUsersInRoom(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersInRoom", reflect.TypeOf((*MockRoomRepository)(nil).GetUsersInRoom), arg0, arg1)
}

// RemoveMemberFromRoom mocks base method.
func (m *MockRoomRepository) RemoveMemberFromRoom(arg0 context.Context, arg1 entity.RoomID, arg2 entity.UserID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMemberFromRoom", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMemberFromRoom indicates an expected call of RemoveMemberFromRoom.
func (mr *MockRoomRepositoryMockRecorder) RemoveMemberFromRoom(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMemberFromRoom", reflect.TypeOf((*MockRoomRepository)(nil).RemoveMemberFromRoom), arg0, arg1, arg2)
}

// SaveRoom mocks base method.
func (m *MockRoomRepository) SaveRoom(arg0 context.Context, arg1 *entity.Room) (entity.RoomID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveRoom", arg0, arg1)
	ret0, _ := ret[0].(entity.RoomID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveRoom indicates an expected call of SaveRoom.
func (mr *MockRoomRepositoryMockRecorder) SaveRoom(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveRoom", reflect.TypeOf((*MockRoomRepository)(nil).SaveRoom), arg0, arg1)
}

// UpdateRoomName mocks base method.
func (m *MockRoomRepository) UpdateRoomName(arg0 context.Context, arg1 entity.RoomID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoomName", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRoomName indicates an expected call of UpdateRoomName.
func (mr *MockRoomRepositoryMockRecorder) UpdateRoomName(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoomName", reflect.TypeOf((*MockRoomRepository)(nil).UpdateRoomName), arg0, arg1, arg2)
}
