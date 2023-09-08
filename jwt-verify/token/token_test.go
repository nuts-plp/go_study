package token

import (
	"reflect"
	"testing"
	"time"
)

const (
	id     = 121923108723812741
	email  = "nuts_plp@126.com"
	issuer = "pss"
	secret = "密文"
	expire = time.Hour * 3
)

var tokenS string

func TestGenerateToken(t *testing.T) {
	_, _ = GenerateToken(int64(id), email, issuer, secret, expire)

}
func TestParseToken(t *testing.T) {
	ok, claim := ParseToken(tokenS, secret)
	if !ok {

	}
	if !reflect.DeepEqual(claim.UserId, id) {

	}
	if !reflect.DeepEqual(claim.UserEmail, email) {

	}
}
