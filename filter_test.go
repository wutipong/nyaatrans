package main

import (
	"github.com/antonmedv/expr"
	"github.com/stretchr/testify/suite"

	"testing"
)

type EvalulateTestSuite struct {
	suite.Suite
}

func TestEvalulateTestSuite(t *testing.T) {
	suite.Run(t, new(EvalulateTestSuite))
}

func (suite *EvalulateTestSuite) TestEmptyExpression() {
	expression := ""
	_, err := expr.Compile(expression, expr.Env(NyaaTorrentItem{}))

	suite.Assert().NotNil(err)
}

func (suite *EvalulateTestSuite) TestSeederExceedExpr() {
	expression := "Seeder > 20"
	program, err := expr.Compile(expression, expr.Env(NyaaTorrentItem{}))
	suite.Assert().Nil(err)

	item := NyaaTorrentItem{
		Seeder: 100,
	}

	res, err := Evaluate(item, program)

	suite.Assert().Nil(err)
	suite.Assert().True(res)
}

func (suite *EvalulateTestSuite) TestSeederNotExceedExpr() {
	expression := "Seeder > 20"
	program, err := expr.Compile(expression, expr.Env(NyaaTorrentItem{}))
	suite.Assert().Nil(err)

	item := NyaaTorrentItem{
		Seeder: 5,
	}

	res, err := Evaluate(item, program)

	suite.Assert().Nil(err)
	suite.Assert().False(res)
}

type FilterNyaaItemsTestSuite struct {
	suite.Suite
}

func TestFilterNyaaItemsTestSuite(t *testing.T) {
	suite.Run(t, new(FilterNyaaItemsTestSuite))
}

func (suite *FilterNyaaItemsTestSuite) TestInvalidExpression() {
	expression := "Hello > 20"
	items := []NyaaTorrentItem{
		{
			Seeder: 100,
		},
		{
			Seeder: 10,
		},
		{
			Seeder: 200,
		},
	}

	_, err := FilterNyaaItems(expression, items)
	suite.Assert().NotNil(err)
}

func (suite *FilterNyaaItemsTestSuite) TestMultipleItems() {
	expression := "Seeder > 20"
	items := []NyaaTorrentItem{
		{
			Seeder: 100,
		},
		{
			Seeder: 10,
		},
		{
			Seeder: 200,
		},
	}

	output, err := FilterNyaaItems(expression, items)

	suite.Assert().Nil(err)
	suite.Assert().ElementsMatch(output, []NyaaTorrentItem{
		{
			Seeder: 100,
		},
		{
			Seeder: 200,
		},
	})
}

func (suite *FilterNyaaItemsTestSuite) TestAllFails() {
	expression := "Seeder > 20"
	items := []NyaaTorrentItem{
		{
			Seeder: 1,
		},
		{
			Seeder: 10,
		},
		{
			Seeder: 2,
		},
	}

	output, err := FilterNyaaItems(expression, items)
	suite.Assert().Nil(err)
	suite.Assert().ElementsMatch(output, []NyaaTorrentItem{})
}

func (suite *FilterNyaaItemsTestSuite) TestEmpty() {
	expression := "Seeder > 20"
	items := []NyaaTorrentItem{}

	output, err := FilterNyaaItems(expression, items)
	suite.Assert().Nil(err)
	suite.Assert().ElementsMatch(output, []NyaaTorrentItem{})
}
