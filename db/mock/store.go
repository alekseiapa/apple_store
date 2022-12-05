// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/alekseiapa/apple_store/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/alekseiapa/apple_store/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// BuyProductTx mocks base method.
func (m *MockStore) BuyProductTx(arg0 context.Context, arg1 db.BuyProductTxParams) (db.BuyProductTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyProductTx", arg0, arg1)
	ret0, _ := ret[0].(db.BuyProductTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuyProductTx indicates an expected call of BuyProductTx.
func (mr *MockStoreMockRecorder) BuyProductTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyProductTx", reflect.TypeOf((*MockStore)(nil).BuyProductTx), arg0, arg1)
}

// CreateOrder mocks base method.
func (m *MockStore) CreateOrder(arg0 context.Context, arg1 db.CreateOrderParams) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockStoreMockRecorder) CreateOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockStore)(nil).CreateOrder), arg0, arg1)
}

// CreateOrderProduct mocks base method.
func (m *MockStore) CreateOrderProduct(arg0 context.Context, arg1 db.CreateOrderProductParams) (db.OrderProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderProduct", arg0, arg1)
	ret0, _ := ret[0].(db.OrderProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrderProduct indicates an expected call of CreateOrderProduct.
func (mr *MockStoreMockRecorder) CreateOrderProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderProduct", reflect.TypeOf((*MockStore)(nil).CreateOrderProduct), arg0, arg1)
}

// CreateProduct mocks base method.
func (m *MockStore) CreateProduct(arg0 context.Context, arg1 db.CreateProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockStoreMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockStore)(nil).CreateProduct), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// CreateUserToUser mocks base method.
func (m *MockStore) CreateUserToUser(arg0 context.Context, arg1 db.CreateUserToUserParams) (db.UserToUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserToUser", arg0, arg1)
	ret0, _ := ret[0].(db.UserToUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserToUser indicates an expected call of CreateUserToUser.
func (mr *MockStoreMockRecorder) CreateUserToUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserToUser", reflect.TypeOf((*MockStore)(nil).CreateUserToUser), arg0, arg1)
}

// DeleteOrder mocks base method.
func (m *MockStore) DeleteOrder(arg0 context.Context, arg1 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockStoreMockRecorder) DeleteOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockStore)(nil).DeleteOrder), arg0, arg1)
}

// DeleteProduct mocks base method.
func (m *MockStore) DeleteProduct(arg0 context.Context, arg1 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockStoreMockRecorder) DeleteProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockStore)(nil).DeleteProduct), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// GetOrder mocks base method.
func (m *MockStore) GetOrder(arg0 context.Context, arg1 int64) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", arg0, arg1)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockStoreMockRecorder) GetOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockStore)(nil).GetOrder), arg0, arg1)
}

// GetOrderProduct mocks base method.
func (m *MockStore) GetOrderProduct(arg0 context.Context, arg1 db.GetOrderProductParams) (db.OrderProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderProduct", arg0, arg1)
	ret0, _ := ret[0].(db.OrderProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderProduct indicates an expected call of GetOrderProduct.
func (mr *MockStoreMockRecorder) GetOrderProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderProduct", reflect.TypeOf((*MockStore)(nil).GetOrderProduct), arg0, arg1)
}

// GetProduct mocks base method.
func (m *MockStore) GetProduct(arg0 context.Context, arg1 int64) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockStoreMockRecorder) GetProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockStore)(nil).GetProduct), arg0, arg1)
}

// GetProductForUpdate mocks base method.
func (m *MockStore) GetProductForUpdate(arg0 context.Context, arg1 int64) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductForUpdate indicates an expected call of GetProductForUpdate.
func (mr *MockStoreMockRecorder) GetProductForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductForUpdate", reflect.TypeOf((*MockStore)(nil).GetProductForUpdate), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserForUpdate mocks base method.
func (m *MockStore) GetUserForUpdate(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserForUpdate indicates an expected call of GetUserForUpdate.
func (mr *MockStoreMockRecorder) GetUserForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserForUpdate", reflect.TypeOf((*MockStore)(nil).GetUserForUpdate), arg0, arg1)
}

// GetUserToUser mocks base method.
func (m *MockStore) GetUserToUser(arg0 context.Context, arg1 db.GetUserToUserParams) (db.UserToUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserToUser", arg0, arg1)
	ret0, _ := ret[0].(db.UserToUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserToUser indicates an expected call of GetUserToUser.
func (mr *MockStoreMockRecorder) GetUserToUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserToUser", reflect.TypeOf((*MockStore)(nil).GetUserToUser), arg0, arg1)
}

// ListOrderProducts mocks base method.
func (m *MockStore) ListOrderProducts(arg0 context.Context, arg1 db.ListOrderProductsParams) ([]db.OrderProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrderProducts", arg0, arg1)
	ret0, _ := ret[0].([]db.OrderProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrderProducts indicates an expected call of ListOrderProducts.
func (mr *MockStoreMockRecorder) ListOrderProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrderProducts", reflect.TypeOf((*MockStore)(nil).ListOrderProducts), arg0, arg1)
}

// ListOrders mocks base method.
func (m *MockStore) ListOrders(arg0 context.Context, arg1 db.ListOrdersParams) ([]db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrders", arg0, arg1)
	ret0, _ := ret[0].([]db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrders indicates an expected call of ListOrders.
func (mr *MockStoreMockRecorder) ListOrders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrders", reflect.TypeOf((*MockStore)(nil).ListOrders), arg0, arg1)
}

// ListProducts mocks base method.
func (m *MockStore) ListProducts(arg0 context.Context, arg1 db.ListProductsParams) ([]db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", arg0, arg1)
	ret0, _ := ret[0].([]db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockStoreMockRecorder) ListProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockStore)(nil).ListProducts), arg0, arg1)
}

// ListUserToUser mocks base method.
func (m *MockStore) ListUserToUser(arg0 context.Context, arg1 db.ListUserToUserParams) ([]db.UserToUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserToUser", arg0, arg1)
	ret0, _ := ret[0].([]db.UserToUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserToUser indicates an expected call of ListUserToUser.
func (mr *MockStoreMockRecorder) ListUserToUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserToUser", reflect.TypeOf((*MockStore)(nil).ListUserToUser), arg0, arg1)
}

// ListUsers mocks base method.
func (m *MockStore) ListUsers(arg0 context.Context, arg1 db.ListUsersParams) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0, arg1)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockStoreMockRecorder) ListUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockStore)(nil).ListUsers), arg0, arg1)
}

// ReduceProductInStock mocks base method.
func (m *MockStore) ReduceProductInStock(arg0 context.Context, arg1 db.ReduceProductInStockParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReduceProductInStock", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReduceProductInStock indicates an expected call of ReduceProductInStock.
func (mr *MockStoreMockRecorder) ReduceProductInStock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReduceProductInStock", reflect.TypeOf((*MockStore)(nil).ReduceProductInStock), arg0, arg1)
}

// ReduceUserBalance mocks base method.
func (m *MockStore) ReduceUserBalance(arg0 context.Context, arg1 db.ReduceUserBalanceParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReduceUserBalance", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReduceUserBalance indicates an expected call of ReduceUserBalance.
func (mr *MockStoreMockRecorder) ReduceUserBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReduceUserBalance", reflect.TypeOf((*MockStore)(nil).ReduceUserBalance), arg0, arg1)
}

// UpdateOrder mocks base method.
func (m *MockStore) UpdateOrder(arg0 context.Context, arg1 db.UpdateOrderParams) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", arg0, arg1)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockStoreMockRecorder) UpdateOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockStore)(nil).UpdateOrder), arg0, arg1)
}

// UpdateOrderProduct mocks base method.
func (m *MockStore) UpdateOrderProduct(arg0 context.Context, arg1 db.UpdateOrderProductParams) (db.OrderProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderProduct", arg0, arg1)
	ret0, _ := ret[0].(db.OrderProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrderProduct indicates an expected call of UpdateOrderProduct.
func (mr *MockStoreMockRecorder) UpdateOrderProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderProduct", reflect.TypeOf((*MockStore)(nil).UpdateOrderProduct), arg0, arg1)
}

// UpdateProduct mocks base method.
func (m *MockStore) UpdateProduct(arg0 context.Context, arg1 db.UpdateProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockStoreMockRecorder) UpdateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockStore)(nil).UpdateProduct), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}

// UpdateUserToUser mocks base method.
func (m *MockStore) UpdateUserToUser(arg0 context.Context, arg1 db.UpdateUserToUserParams) (db.UserToUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserToUser", arg0, arg1)
	ret0, _ := ret[0].(db.UserToUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserToUser indicates an expected call of UpdateUserToUser.
func (mr *MockStoreMockRecorder) UpdateUserToUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserToUser", reflect.TypeOf((*MockStore)(nil).UpdateUserToUser), arg0, arg1)
}