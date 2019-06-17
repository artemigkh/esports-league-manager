// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import databaseAccess "Server/databaseAccess"
import lolApi "Server/lolApi"
import mock "github.com/stretchr/testify/mock"

// LeagueOfLegendsDAO is an autogenerated mock type for the LeagueOfLegendsDAO type
type LeagueOfLegendsDAO struct {
	mock.Mock
}

// GetChampionStats provides a mock function with given fields: leagueId
func (_m *LeagueOfLegendsDAO) GetChampionStats(leagueId int) ([]*databaseAccess.ChampionStats, error) {
	ret := _m.Called(leagueId)

	var r0 []*databaseAccess.ChampionStats
	if rf, ok := ret.Get(0).(func(int) []*databaseAccess.ChampionStats); ok {
		r0 = rf(leagueId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*databaseAccess.ChampionStats)
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

// GetPlayerStats provides a mock function with given fields: leagueId
func (_m *LeagueOfLegendsDAO) GetPlayerStats(leagueId int) ([]*databaseAccess.PlayerStats, error) {
	ret := _m.Called(leagueId)

	var r0 []*databaseAccess.PlayerStats
	if rf, ok := ret.Get(0).(func(int) []*databaseAccess.PlayerStats); ok {
		r0 = rf(leagueId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*databaseAccess.PlayerStats)
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

// GetTeamStats provides a mock function with given fields: leagueId
func (_m *LeagueOfLegendsDAO) GetTeamStats(leagueId int) ([]*databaseAccess.TeamStats, error) {
	ret := _m.Called(leagueId)

	var r0 []*databaseAccess.TeamStats
	if rf, ok := ret.Get(0).(func(int) []*databaseAccess.TeamStats); ok {
		r0 = rf(leagueId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*databaseAccess.TeamStats)
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

// ReportEndGameStats provides a mock function with given fields: leagueId, gameId, winTeamId, loseTeamId, match
func (_m *LeagueOfLegendsDAO) ReportEndGameStats(leagueId int, gameId int, winTeamId int, loseTeamId int, match *lolApi.MatchInformation) error {
	ret := _m.Called(leagueId, gameId, winTeamId, loseTeamId, match)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, int, int, *lolApi.MatchInformation) error); ok {
		r0 = rf(leagueId, gameId, winTeamId, loseTeamId, match)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}