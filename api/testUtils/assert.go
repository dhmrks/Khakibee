package testUtils

import (
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt"
	"gitlab.com/khakibee/khakibee/api/user"
)

func AssertAuthJWT(t testing.TB, tokenString string, jwtSecret string, au user.AuthUserJWTClaims) {
	t.Helper()

	token, err := jwt.ParseWithClaims(tokenString, &user.AuthUserJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	AssertError(t, err, false)

	if token == nil || !token.Valid {
		t.Errorf("couldn't handle this token")
		return
	}

	claims := token.Claims.(*user.AuthUserJWTClaims)

	if au.UserID != claims.UserID {
		t.Errorf("wanted user id %d but got %d", au.UserID, claims.UserID)
	}
}

func AssertHTTPStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted http status %d but got http status %d", want, got)
	}
}

func AssertEqual(t testing.TB, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %q \nbut got %q", want, got)
	}
}

func AssertFunctionCalled(t testing.TB, got, want int, funcName string) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %d call(s) to '%s' but got %d", want, funcName, got)
	}
}

func AssertFunctionCalleWith(t testing.TB, got, want interface{}, funcName string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted to call  %[1]s( %[2]q )  and not %[1]s( %[3]q )", funcName, want, got)
	}
}

func AssertError(t testing.TB, err error, expected bool) {
	t.Helper()
	if expected && err == nil {
		t.Errorf("expected to get an error")
	} else if !expected && err != nil {
		t.Errorf("not expected to get an error but got '%s'", err.Error())
	}
}

func AssertNil(t testing.TB, got interface{}) {
	t.Helper()
	if !reflect.ValueOf(got).IsNil() {
		t.Errorf("Wanted nil but got %q", got)
	}
}
