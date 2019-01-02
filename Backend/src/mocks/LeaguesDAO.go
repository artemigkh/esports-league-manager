// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import databaseAccess "Server/databaseAccess"
import mock "github.com/stretchr/testify/mock"

// LeaguesDAO is an autogenerated mock type for the LeaguesDAO type
type LeaguesDAO struct {
	mock.Mock
}

// CanJoinLeague provides a mock function with given fields: leagueId, userId
func (_m *LeaguesDAO) CanJoinLeague(leagueId int, userId int) (bool, error) {
	ret := _m.Called(leagueId, userId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int, int) bool); ok {
		r0 = rf(leagueId, userId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(leagueId, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLeague provides a mock function with given fields: userId, name, description, game, publicView, publicJoin, signupStart, signupEnd, leagueStart, leagueEnd
func (_m *LeaguesDAO) CreateLeague(userId int, name string, description string, game string, publicView bool, publicJoin bool, signupStart int, signupEnd int, leagueStart int, leagueEnd int) (int, error) {
	ret := _m.Called(userId, name, description, game, publicView, publicJoin, signupStart, signupEnd, leagueStart, leagueEnd)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, string, string, string, bool, bool, int, int, int, int) int); ok {
		r0 = rf(userId, name, description, game, publicView, publicJoin, signupStart, signupEnd, leagueStart, leagueEnd)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string, string, string, bool, bool, int, int, int, int) error); ok {
		r1 = rf(userId, name, description, game, publicView, publicJoin, signupStart, signupEnd, leagueStart, leagueEnd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGameSummary provides a mock function with given fields: leagueId
func (_m *LeaguesDAO) GetGameSummary(leagueId int) ([]databaseAccess.GameSummaryInformation, error) {
	ret := _m.Called(leagueId)

	var r0 []databaseAccess.GameSummaryInformation
	if rf, ok := ret.Get(0).(func(int) []databaseAccess.GameSummaryInformation); ok {
		r0 = rf(leagueId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]databaseAccess.GameSummaryInformation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(leagueId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLeagueInformation provides a mock function with given fields: leagueId
func (_m *LeaguesDAO) GetLeagueInformation(leagueId int) (*databaseAccess.LeagueInformation, error) {
	ret := _m.Called(leagueId)

	var r0 *databaseAccess.LeagueInformation
	if rf, ok := ret.Get(0).(func(int) *databaseAccess.LeagueInformation); ok {
		r0 = rf(leagueId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*databaseAccess.LeagueInformation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(leagueId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLeaguePermissions provides a mock function with given fields: leagueId, userId
func (_m *LeaguesDAO) GetLeaguePermissions(leagueId int, userId int) (*databaseAccess.LeaguePermissions, error) {
	ret := _m.Called(leagueId, userId)

	var r0 *databaseAccess.LeaguePermissions
	if rf, ok := ret.Get(0).(func(int, int) *databaseAccess.LeaguePermissions); ok {
		r0 = rf(leagueId, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*databaseAccess.LeaguePermissions)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(leagueId, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPublicLeagueList provides a mock function with given fields:
func (_m *LeaguesDAO) GetPublicLeagueList() ([]databaseAccess.PublicLeagueInformation, error) {
	ret := _m.Called()

	var r0 []databaseAccess.PublicLeagueInformation
	if rf, ok := ret.Get(0).(func() []databaseAccess.PublicLeagueInformation); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]databaseAccess.PublicLeagueInformation)
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

// GetTeamManagerInformation provides a mock function with given fields: leagueId
func (_m *LeaguesDAO) GetTeamManagerInformation(leagueId int) ([]databaseAccess.TeamManagerInformation, error) {
	ret := _m.Called(leagueId)

	var r0 []databaseAccess.TeamManagerInformation
	if rf, ok := ret.Get(0).(func(int) []databaseAccess.TeamManagerInformation); ok {
		r0 = rf(leagueId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]databaseAccess.TeamManagerInformation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(leagueId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeamSummary provides a mock function with given fields: leagueId
func (_m *LeaguesDAO) GetTeamSummary(leagueId int) ([]databaseAccess.TeamSummaryInformation, error) {
	ret := _m.Called(leagueId)

	var r0 []databaseAccess.TeamSummaryInformation
	if rf, ok := ret.Get(0).(func(int) []databaseAccess.TeamSummaryInformation); ok {
		r0 = rf(leagueId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]databaseAccess.TeamSummaryInformation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(leagueId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsLeagueViewable provides a mock function with given fields: leagueId, userId
func (_m *LeaguesDAO) IsLeagueViewable(leagueId int, userId int) (bool, error) {
	ret := _m.Called(leagueId, userId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int, int) bool); ok {
		r0 = rf(leagueId, userId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(leagueId, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsNameInUse provides a mock function with given fields: leagueId, name
func (_m *LeaguesDAO) IsNameInUse(leagueId int, name string) (bool, error) {
	ret := _m.Called(leagueId, name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int, string) bool); ok {
		r0 = rf(leagueId, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(leagueId, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// JoinLeague provides a mock function with given fields: leagueId, userId
func (_m *LeaguesDAO) JoinLeague(leagueId int, userId int) error {
	ret := _m.Called(leagueId, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(leagueId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetLeaguePermissions provides a mock function with given fields: leagueId, userId, administrator, createTeams, editTeams, editGames
func (_m *LeaguesDAO) SetLeaguePermissions(leagueId int, userId int, administrator bool, createTeams bool, editTeams bool, editGames bool) error {
	ret := _m.Called(leagueId, userId, administrator, createTeams, editTeams, editGames)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, bool, bool, bool, bool) error); ok {
		r0 = rf(leagueId, userId, administrator, createTeams, editTeams, editGames)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateLeague provides a mock function with given fields: leagueId, name, description, game, publicView, publicJoin, signupStart, signupEnd, leagueStart, leagueEnd
func (_m *LeaguesDAO) UpdateLeague(leagueId int, name string, description string, game string, publicView bool, publicJoin bool, signupStart int, signupEnd int, leagueStart int, leagueEnd int) error {
	ret := _m.Called(leagueId, name, description, game, publicView, publicJoin, signupStart, signupEnd, leagueStart, leagueEnd)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string, string, string, bool, bool, int, int, int, int) error); ok {
		r0 = rf(leagueId, name, description, game, publicView, publicJoin, signupStart, signupEnd, leagueStart, leagueEnd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}