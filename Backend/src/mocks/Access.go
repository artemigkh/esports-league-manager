// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import databaseAccess "Server/databaseAccess"
import mock "github.com/stretchr/testify/mock"

// Access is an autogenerated mock type for the Access type
type Access struct {
	mock.Mock
}

// Game provides a mock function with given fields: accessType, leagueId, gameId, userId
func (_m *Access) Game(accessType databaseAccess.AccessType, leagueId int, gameId int, userId int) (bool, error) {
	ret := _m.Called(accessType, leagueId, gameId, userId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(databaseAccess.AccessType, int, int, int) bool); ok {
		r0 = rf(accessType, leagueId, gameId, userId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(databaseAccess.AccessType, int, int, int) error); ok {
		r1 = rf(accessType, leagueId, gameId, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// League provides a mock function with given fields: accessType, leagueId, userId
func (_m *Access) League(accessType databaseAccess.AccessType, leagueId int, userId int) (bool, error) {
	ret := _m.Called(accessType, leagueId, userId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(databaseAccess.AccessType, int, int) bool); ok {
		r0 = rf(accessType, leagueId, userId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(databaseAccess.AccessType, int, int) error); ok {
		r1 = rf(accessType, leagueId, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Player provides a mock function with given fields: accessType, leagueId, teamId, playerId, userId
func (_m *Access) Player(accessType databaseAccess.AccessType, leagueId int, teamId int, playerId int, userId int) (bool, error) {
	ret := _m.Called(accessType, leagueId, teamId, playerId, userId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(databaseAccess.AccessType, int, int, int, int) bool); ok {
		r0 = rf(accessType, leagueId, teamId, playerId, userId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(databaseAccess.AccessType, int, int, int, int) error); ok {
		r1 = rf(accessType, leagueId, teamId, playerId, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Report provides a mock function with given fields: leagueId, gameId, userId
func (_m *Access) Report(leagueId int, gameId int, userId int) (bool, error) {
	ret := _m.Called(leagueId, gameId, userId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int, int, int) bool); ok {
		r0 = rf(leagueId, gameId, userId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, int) error); ok {
		r1 = rf(leagueId, gameId, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Team provides a mock function with given fields: accessType, leagueId, teamId, userId
func (_m *Access) Team(accessType databaseAccess.AccessType, leagueId int, teamId int, userId int) (bool, error) {
	ret := _m.Called(accessType, leagueId, teamId, userId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(databaseAccess.AccessType, int, int, int) bool); ok {
		r0 = rf(accessType, leagueId, teamId, userId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(databaseAccess.AccessType, int, int, int) error); ok {
		r1 = rf(accessType, leagueId, teamId, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
