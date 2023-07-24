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
