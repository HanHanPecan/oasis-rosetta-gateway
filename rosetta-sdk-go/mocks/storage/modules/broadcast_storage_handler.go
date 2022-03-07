// Code generated by mockery v1.0.0. DO NOT EDIT.

package modules

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	database "github.com/oasisprotocol/oasis-rosetta-gateway/rosetta-sdk-go/storage/database"
	types "github.com/oasisprotocol/oasis-rosetta-gateway/rosetta-sdk-go/types"
)

// BroadcastStorageHandler is an autogenerated mock type for the BroadcastStorageHandler type
type BroadcastStorageHandler struct {
	mock.Mock
}

// BroadcastFailed provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *BroadcastStorageHandler) BroadcastFailed(_a0 context.Context, _a1 database.Transaction, _a2 string, _a3 *types.TransactionIdentifier, _a4 []*types.Operation) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, string, *types.TransactionIdentifier, []*types.Operation) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TransactionConfirmed provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *BroadcastStorageHandler) TransactionConfirmed(_a0 context.Context, _a1 database.Transaction, _a2 string, _a3 *types.BlockIdentifier, _a4 *types.Transaction, _a5 []*types.Operation) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, string, *types.BlockIdentifier, *types.Transaction, []*types.Operation) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TransactionStale provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *BroadcastStorageHandler) TransactionStale(_a0 context.Context, _a1 database.Transaction, _a2 string, _a3 *types.TransactionIdentifier) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, string, *types.TransactionIdentifier) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
