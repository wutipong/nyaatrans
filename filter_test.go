package main

import (
	"testing"

	"github.com/apaxa-go/eval"
	"github.com/stretchr/testify/suite"
)

type EvalulateTestSuite struct {
	suite.Suite
}

func TestEvalulateTestSuite(t *testing.T) {
	suite.Run(t, new(EvalulateTestSuite))
}

func (suite *EvalulateTestSuite) TestEmptyExpression() {
	expr := ""
	_, err := eval.ParseString(expr, "")
	suite.Assert().NotNil(err)
}

func (suite *EvalulateTestSuite) TestSeederExceedExpr() {
	expr := "item.Seeder > 20"
	exprObj, _ := eval.ParseString(expr, "")

	item := NyaaTorrentItem{
		Seeder: 100,
	}

	res, err := Evaluate(item, exprObj)

	suite.Assert().Nil(err)
	suite.Assert().True(res)
}

func (suite *EvalulateTestSuite) TestSeederNotExceedExpr() {
	expr := "item.Seeder > 20"
	exprObj, _ := eval.ParseString(expr, "")

	item := NyaaTorrentItem{
		Seeder: 5,
	}

	res, err := Evaluate(item, exprObj)

	suite.Assert().Nil(err)
	suite.Assert().False(res)
}
