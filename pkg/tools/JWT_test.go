package tools

import (
	"fmt"
	"testing"
)

func TestGenToken(t *testing.T) {
	Token, _ := GenToken("hello", 5)

	fmt.Println(Token)

	fmt.Println(ParseToken(Token))
	claims, _ := ParseToken(Token)
	fmt.Println(claims.Username)
	fmt.Println(claims.User_id)
}
