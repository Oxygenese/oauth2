package models

import (
	"time"

	"github.com/go-oauth2/oauth2/v4"
)

// NewToken create to token model instance
func NewToken() *Token {
	return &Token{}
}

// Token token model
type Token struct {
	ClientID            string        `bson:"client_id" json:"client_id" json:"client_id,omitempty"`
	UserID              string        `bson:"user_id" json:"user_id" json:"user_id,omitempty"`
	RedirectURI         string        `bson:"redirect_uri" json:"redirect_uri" json:"redirect_uri,omitempty"`
	Scope               string        `bson:"scope" json:"scope,omitempty"`
	Code                string        `bson:"code" json:"code,omitempty"`
	CodeChallenge       string        `bson:"code_challenge" json:"code_challenge,omitempty"`
	CodeChallengeMethod string        `bson:"code_challenge_method" json:"code_challenge_method,omitempty"`
	CodeCreateAt        time.Time     `bson:"code_create_at" json:"code_create_at"`
	CodeExpiresIn       time.Duration `bson:"code_expires_in" json:"code_expires_in,omitempty"`
	AccessToken         string        `bson:"access_token" json:"access_token,omitempty"`
	AccessCreateAt      time.Time     `bson:"access_create_at" json:"access_create_at"`
	AccessExpiresIn     time.Duration `bson:"access_expires_in" json:"access_expires_in,omitempty"`
	RefreshToken        string        `bson:"refresh_token" json:"refresh_token,omitempty"`
	RefreshCreateAt     time.Time     `bson:"refresh_create_at" json:"refresh_create_at"`
	RefreshExpiresIn    time.Duration `bson:"refresh_expires_in" json:"refresh_expires_in,omitempty"`
	ExtensionClaims     interface{}   `bson:"extension_claims" json:"extension_claims"`
}

func (t *Token) SetExtensionClaims(claim interface{}) {
	t.ExtensionClaims = claim
}

// New create to token model instance
func (t *Token) New() oauth2.TokenInfo {
	return NewToken()
}

// GetClientID the client id
func (t *Token) GetClientID() string {
	return t.ClientID
}

// SetClientID the client id
func (t *Token) SetClientID(clientID string) {
	t.ClientID = clientID
}

// GetUserID the user id
func (t *Token) GetUserID() string {
	return t.UserID
}

// SetUserID the user id
func (t *Token) SetUserID(userID string) {
	t.UserID = userID
}

// GetRedirectURI redirect URI
func (t *Token) GetRedirectURI() string {
	return t.RedirectURI
}

// SetRedirectURI redirect URI
func (t *Token) SetRedirectURI(redirectURI string) {
	t.RedirectURI = redirectURI
}

// GetScope get scope of authorization
func (t *Token) GetScope() string {
	return t.Scope
}

// SetScope get scope of authorization
func (t *Token) SetScope(scope string) {
	t.Scope = scope
}

// GetCode authorization code
func (t *Token) GetCode() string {
	return t.Code
}

// SetCode authorization code
func (t *Token) SetCode(code string) {
	t.Code = code
}

// GetCodeCreateAt create Time
func (t *Token) GetCodeCreateAt() time.Time {
	return t.CodeCreateAt
}

// SetCodeCreateAt create Time
func (t *Token) SetCodeCreateAt(createAt time.Time) {
	t.CodeCreateAt = createAt
}

// GetCodeExpiresIn the lifetime in seconds of the authorization code
func (t *Token) GetCodeExpiresIn() time.Duration {
	return t.CodeExpiresIn
}

// SetCodeExpiresIn the lifetime in seconds of the authorization code
func (t *Token) SetCodeExpiresIn(exp time.Duration) {
	t.CodeExpiresIn = exp
}

// GetCodeChallenge challenge code
func (t *Token) GetCodeChallenge() string {
	return t.CodeChallenge
}

// SetCodeChallenge challenge code
func (t *Token) SetCodeChallenge(code string) {
	t.CodeChallenge = code
}

// GetCodeChallengeMethod challenge method
func (t *Token) GetCodeChallengeMethod() oauth2.CodeChallengeMethod {
	return oauth2.CodeChallengeMethod(t.CodeChallengeMethod)
}

// SetCodeChallengeMethod challenge method
func (t *Token) SetCodeChallengeMethod(method oauth2.CodeChallengeMethod) {
	t.CodeChallengeMethod = string(method)
}

// GetAccess access Token
func (t *Token) GetAccess() string {
	return t.AccessToken
}

// SetAccess access Token
func (t *Token) SetAccess(access string) {
	t.AccessToken = access
}

// GetAccessCreateAt create Time
func (t *Token) GetAccessCreateAt() time.Time {
	return t.AccessCreateAt
}

// SetAccessCreateAt create Time
func (t *Token) SetAccessCreateAt(createAt time.Time) {
	t.AccessCreateAt = createAt
}

// GetAccessExpiresIn the lifetime in seconds of the access token
func (t *Token) GetAccessExpiresIn() time.Duration {
	return t.AccessExpiresIn
}

// SetAccessExpiresIn the lifetime in seconds of the access token
func (t *Token) SetAccessExpiresIn(exp time.Duration) {
	t.AccessExpiresIn = exp
}

// GetRefresh refresh Token
func (t *Token) GetRefresh() string {
	return t.RefreshToken
}

// SetRefresh refresh Token
func (t *Token) SetRefresh(refresh string) {
	t.RefreshToken = refresh
}

// GetRefreshCreateAt create Time
func (t *Token) GetRefreshCreateAt() time.Time {
	return t.RefreshCreateAt
}

// SetRefreshCreateAt create Time
func (t *Token) SetRefreshCreateAt(createAt time.Time) {
	t.RefreshCreateAt = createAt
}

// GetRefreshExpiresIn the lifetime in seconds of the refresh token
func (t *Token) GetRefreshExpiresIn() time.Duration {
	return t.RefreshExpiresIn
}

// SetRefreshExpiresIn the lifetime in seconds of the refresh token
func (t *Token) SetRefreshExpiresIn(exp time.Duration) {
	t.RefreshExpiresIn = exp
}
