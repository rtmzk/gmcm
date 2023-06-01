package middleware

import "github.com/gin-gonic/gin"

type AuthStrategy interface {
	AuthFunc() gin.HandlerFunc
}

type AuthOperator struct {
	strategy AuthStrategy
}

func (operator *AuthOperator) SetStrategy(stragegy AuthStrategy) {
	operator.strategy = stragegy
}

func (operator *AuthOperator) AuthFunc() gin.HandlerFunc {
	return operator.strategy.AuthFunc()
}
