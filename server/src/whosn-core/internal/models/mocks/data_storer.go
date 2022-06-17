// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	testing "testing"

	models "github.com/Ahmad-Ibra/whosn-core/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// DataStorer is an autogenerated mock type for the DataStorer type
type DataStorer struct {
	mock.Mock
}

// DeleteEventByID provides a mock function with given fields: eventID
func (_m *DataStorer) DeleteEventByID(eventID string) error {
	ret := _m.Called(eventID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(eventID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUserByID provides a mock function with given fields: userID
func (_m *DataStorer) DeleteUserByID(userID string) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetEventByID provides a mock function with given fields: eventID
func (_m *DataStorer) GetEventByID(eventID string) (*models.Event, error) {
	ret := _m.Called(eventID)

	var r0 *models.Event
	if rf, ok := ret.Get(0).(func(string) *models.Event); ok {
		r0 = rf(eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(eventID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: userID
func (_m *DataStorer) GetUserByID(userID string) (*models.User, error) {
	ret := _m.Called(userID)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUsername provides a mock function with given fields: username
func (_m *DataStorer) GetUserByUsername(username string) (*models.User, error) {
	ret := _m.Called(username)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertEvent provides a mock function with given fields: event
func (_m *DataStorer) InsertEvent(event models.Event) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Event) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertUser provides a mock function with given fields: user
func (_m *DataStorer) InsertUser(user models.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListAllEvents provides a mock function with given fields:
func (_m *DataStorer) ListAllEvents() (*[]models.Event, error) {
	ret := _m.Called()

	var r0 *[]models.Event
	if rf, ok := ret.Get(0).(func() *[]models.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllUsers provides a mock function with given fields:
func (_m *DataStorer) ListAllUsers() (*[]models.User, error) {
	ret := _m.Called()

	var r0 *[]models.User
	if rf, ok := ret.Get(0).(func() *[]models.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEventByID provides a mock function with given fields: eventUpdate, eventID
func (_m *DataStorer) UpdateEventByID(eventUpdate models.Event, eventID string) (*models.Event, error) {
	ret := _m.Called(eventUpdate, eventID)

	var r0 *models.Event
	if rf, ok := ret.Get(0).(func(models.Event, string) *models.Event); ok {
		r0 = rf(eventUpdate, eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Event, string) error); ok {
		r1 = rf(eventUpdate, eventID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserByID provides a mock function with given fields: userUpdate, userID
func (_m *DataStorer) UpdateUserByID(userUpdate models.User, userID string) (*models.User, error) {
	ret := _m.Called(userUpdate, userID)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(models.User, string) *models.User); ok {
		r0 = rf(userUpdate, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.User, string) error); ok {
		r1 = rf(userUpdate, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDataStorer creates a new instance of DataStorer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewDataStorer(t testing.TB) *DataStorer {
	mock := &DataStorer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
