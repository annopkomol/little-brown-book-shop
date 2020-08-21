// Code generated by mockery v2.0.3. DO NOT EDIT.

package mocks

import (
	domain "lbbs-service/domain"

	mock "github.com/stretchr/testify/mock"
)

// CartRepository is an autogenerated mock type for the CartRepository type
type CartRepository struct {
	mock.Mock
}

// CountOrdersInCart provides a mock function with given fields: cartID
func (_m *CartRepository) CountOrdersInCart(cartID int) (int, error) {
	ret := _m.Called(cartID)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(cartID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(cartID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNewCart provides a mock function with given fields: posID
func (_m *CartRepository) CreateNewCart(posID int) (domain.Cart, error) {
	ret := _m.Called(posID)

	var r0 domain.Cart
	if rf, ok := ret.Get(0).(func(int) domain.Cart); ok {
		r0 = rf(posID)
	} else {
		r0 = ret.Get(0).(domain.Cart)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(posID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNewOrder provides a mock function with given fields: cartID, bookID
func (_m *CartRepository) CreateNewOrder(cartID int, bookID int) (domain.Order, error) {
	ret := _m.Called(cartID, bookID)

	var r0 domain.Order
	if rf, ok := ret.Get(0).(func(int, int) domain.Order); ok {
		r0 = rf(cartID, bookID)
	} else {
		r0 = ret.Get(0).(domain.Order)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(cartID, bookID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOrder provides a mock function with given fields: orderID
func (_m *CartRepository) DeleteOrder(orderID int) error {
	ret := _m.Called(orderID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(orderID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindCartByID provides a mock function with given fields: cartID
func (_m *CartRepository) FindCartByID(cartID int) (domain.Cart, error) {
	ret := _m.Called(cartID)

	var r0 domain.Cart
	if rf, ok := ret.Get(0).(func(int) domain.Cart); ok {
		r0 = rf(cartID)
	} else {
		r0 = ret.Get(0).(domain.Cart)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(cartID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCartByPosID provides a mock function with given fields: posID
func (_m *CartRepository) FindCartByPosID(posID int) (domain.Cart, error) {
	ret := _m.Called(posID)

	var r0 domain.Cart
	if rf, ok := ret.Get(0).(func(int) domain.Cart); ok {
		r0 = rf(posID)
	} else {
		r0 = ret.Get(0).(domain.Cart)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(posID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOrder provides a mock function with given fields: cartID, BookID
func (_m *CartRepository) FindOrder(cartID int, BookID int) (domain.Order, error) {
	ret := _m.Called(cartID, BookID)

	var r0 domain.Order
	if rf, ok := ret.Get(0).(func(int, int) domain.Order); ok {
		r0 = rf(cartID, BookID)
	} else {
		r0 = ret.Get(0).(domain.Order)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(cartID, BookID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FlushCart provides a mock function with given fields: cart
func (_m *CartRepository) FlushCart(cart domain.Cart) error {
	ret := _m.Called(cart)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Cart) error); ok {
		r0 = rf(cart)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOrderQty provides a mock function with given fields: orderID, qty
func (_m *CartRepository) UpdateOrderQty(orderID int, qty int) error {
	ret := _m.Called(orderID, qty)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(orderID, qty)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}