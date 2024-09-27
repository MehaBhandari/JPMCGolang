package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock implementation of UserService using testify/mock
type MockUserService struct {
	mock.Mock
}

// Implement the mock method for GetUser
func (m *MockUserService) GetUser(id int) string {
	args := m.Called(id)
	return args.String(0)
}

func TestGetUserName(t *testing.T) {
	// Create an instance of the mock
	mockService := new(MockUserService)

	// Define the expected behavior
	mockService.On("GetUser", 1).Return("Alice")
	mockService.On("GetUser", 2).Return("Bob")

	assert.Equal(t, "Alicedfysgfiu", GetUserDetails(mockService, 1))
}

// func TestGetUserNameWithDifferentID(t *testing.T) {
// 	mockService := new(MockUserService)

// 	// Define different behavior for different ID
// 	mockService.On("GetUser", 2).Return("Bob")

// 	result := GetUserDetails(mockService, 2)
// 	assert.Equal(t, "Bob", result)
// 	assert.Equal(t, 1, 1)
// 	assert.Equal(t, 1, 1)

// 	mockService.AssertExpectations(t)
// }

// func TestGetUserNameWithError(t *testing.T) {
// 	mockService := new(MockUserService)

// 	// Mock behavior with an error
// 	mockService.On("GetUser", 3).Return("")

// 	result := GetUserDetails(mockService, 3)

// 	assert.Equal(t, "", result)

// 	mockService.AssertExpectations(t)
// }
