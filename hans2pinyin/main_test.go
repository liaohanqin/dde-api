/*
 * Copyright (C) 2019 ~ 2021 Uniontech Software Technology Co.,Ltd
 *
 * Author:     dengbo <dengbo@uniontech.com>
 *
 * Maintainer: dengbo <dengbo@uniontech.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/linuxdeepin/go-lib/dbusutil"
)

type UnitTestSuite struct {
	suite.Suite
	m *Manager
}

func (s *UnitTestSuite) SetupSuite() {
	var err error
	s.m = &Manager{}
	s.m.service, err = dbusutil.NewSessionService()
	if err != nil {
		s.T().Skip(fmt.Sprintf("failed to get service: %v", err))
	}
}

func (s *UnitTestSuite) Test_Query() {
	pinyin, err := s.m.Query("Hanz")
	s.Require().Nil(err)
	s.Assert().ElementsMatch(pinyin, []string{"Hanz"})
}

func (s *UnitTestSuite) Test_QueryQueryList() {
	hansList := []string{"Hanz"}
	jsonStr, err := s.m.QueryList(hansList)
	testStr := `{"Hanz":["Hanz"]}`
	s.Require().Nil(err)
	s.Equal(jsonStr, testStr)
}

func (s *UnitTestSuite) Test_GetInterfaceName() {
	s.Equal(s.m.GetInterfaceName(), dbusServiceName)
}

func (s *UnitTestSuite) Test_usage() {
	usage()
}

func (s *UnitTestSuite) Test_queryPinyin() {
	pinyin := queryPinyin("Hanz")
	s.Assert().ElementsMatch(pinyin, []string{"Hanz"})
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}
