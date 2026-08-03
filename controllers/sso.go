package controllers

import (
	"context"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/astaxie/beego"
	beecontext "github.com/astaxie/beego/context"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorilla/securecookie"
	"golang.org/x/oauth2"
)

const (
	oidcStateCookieName   = "pa_oidc_state"
	oidcSessionCookieName = "pa_oidc_session"
	oidcStateTTL          = 10 * time.Minute
	defaultOIDCSessionTTL = 8 * time.Hour
)

type oidcState struct {
	State     string
	Nonce     string
	ExpiresAt int64
}

type oidcSession struct {
	Subject             string
	Username            string
	AuthorizationValues []string
	ExpiresAt           int64
}

func IsSSOEnabled() bool {
	return beego.AppConfig.DefaultString("sso_enabled", "0") == "1"
}

func ssoCookie(maxAge time.Duration) (*securecookie.SecureCookie, error) {
	secret := strings.TrimSpace(beego.AppConfig.String("sso_cookie_secret"))
	if len(secret) < 32 {
		return nil, fmt.Errorf("sso_cookie_secret must contain at least 32 characters")
	}

	keyMaterial := sha512.Sum512([]byte(secret))
	codec := securecookie.New(keyMaterial[:32], keyMaterial[32:])
	codec.MaxAge(int(maxAge.Seconds()))
	return codec, nil
}

func ssoCookieSecure() bool {
	return beego.AppConfig.DefaultString("sso_cookie_secure", "1") == "1"
}

func setSSOCookie(w http.ResponseWriter, name, value string, maxAge time.Duration) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		MaxAge:   int(maxAge.Seconds()),
		HttpOnly: true,
		Secure:   ssoCookieSecure(),
		SameSite: http.SameSiteLaxMode,
	})
}

func clearSSOCookies(w http.ResponseWriter) {
	for _, name := range []string{oidcStateCookieName, oidcSessionCookieName} {
		http.SetCookie(w, &http.Cookie{
			Name:     name,
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
			Secure:   ssoCookieSecure(),
			SameSite: http.SameSiteLaxMode,
		})
	}
}

func randomString() (string, error) {
	value := make([]byte, 32)
	if _, err := rand.Read(value); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(value), nil
}

func oidcConfig(ctxContext context.Context) (*oauth2.Config, *oidc.IDTokenVerifier, error) {
	issuerURL := strings.TrimSpace(beego.AppConfig.String("sso_issuer_url"))
	clientID := strings.TrimSpace(beego.AppConfig.String("sso_client_id"))
	clientSecret := beego.AppConfig.String("sso_client_secret")
	redirectURL := strings.TrimSpace(beego.AppConfig.String("sso_redirect_url"))
	if issuerURL == "" || clientID == "" || clientSecret == "" || redirectURL == "" {
		return nil, nil, fmt.Errorf("SSO is enabled but the issuer URL, client ID, client secret, or redirect URL is missing")
	}

	provider, err := oidc.NewProvider(ctxContext, issuerURL)
	if err != nil {
		return nil, nil, fmt.Errorf("discover OIDC provider: %w", err)
	}

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  redirectURL,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}, provider.Verifier(&oidc.Config{ClientID: clientID}), nil
}

func ssoSessionTTL() time.Duration {
	seconds, err := beego.AppConfig.Int("sso_session_ttl_seconds")
	if err != nil || seconds <= 0 {
		return defaultOIDCSessionTTL
	}
	return time.Duration(seconds) * time.Second
}

func allowedAuthorizationValues() map[string]struct{} {
	allowedValues := make(map[string]struct{})
	for _, value := range strings.Split(beego.AppConfig.String("sso_authorization_values"), ",") {
		if value = strings.TrimSpace(value); value != "" {
			allowedValues[value] = struct{}{}
		}
	}
	return allowedValues
}

func hasAllowedAuthorizationValue(values []string) bool {
	allowed := allowedAuthorizationValues()
	if len(allowed) == 0 {
		return false
	}
	for _, value := range values {
		if _, ok := allowed[value]; ok {
			return true
		}
	}
	return false
}

func authorizationValuesFromClaims(claims map[string]interface{}) ([]string, error) {
	claimPath := strings.TrimSpace(beego.AppConfig.DefaultString("sso_authorization_claim", "groups"))
	var claimValue interface{} = claims
	for _, part := range strings.Split(claimPath, ".") {
		claimMap, ok := claimValue.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("OIDC authorization claim %q is not an object path", claimPath)
		}
		claimValue, ok = claimMap[part]
		if !ok {
			return nil, fmt.Errorf("OIDC authorization claim %q is missing", claimPath)
		}
	}

	switch values := claimValue.(type) {
	case []interface{}:
		result := make([]string, 0, len(values))
		for _, value := range values {
			valueName, ok := value.(string)
			if !ok {
				return nil, fmt.Errorf("OIDC authorization claim %q contains a non-string value", claimPath)
			}
			result = append(result, valueName)
		}
		return result, nil
	case []string:
		return values, nil
	case string:
		return strings.Split(values, ","), nil
	default:
		return nil, fmt.Errorf("OIDC authorization claim %q must be a string or string array", claimPath)
	}
}

func (c *LoginController) SSOLogin() {
	if !IsSSOEnabled() {
		c.Abort("404")
		return
	}

	state, err := randomString()
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "generate SSO state: "+err.Error(), http.StatusInternalServerError)
		return
	}
	nonce, err := randomString()
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "generate SSO nonce: "+err.Error(), http.StatusInternalServerError)
		return
	}

	codec, err := ssoCookie(oidcStateTTL)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	encodedState, err := codec.Encode(oidcStateCookieName, oidcState{
		State:     state,
		Nonce:     nonce,
		ExpiresAt: time.Now().Add(oidcStateTTL).Unix(),
	})
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "encode SSO state: "+err.Error(), http.StatusInternalServerError)
		return
	}

	oauthConfig, _, err := oidcConfig(c.Ctx.Request.Context())
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusServiceUnavailable)
		return
	}
	setSSOCookie(c.Ctx.ResponseWriter, oidcStateCookieName, encodedState, oidcStateTTL)
	c.Redirect(oauthConfig.AuthCodeURL(state, oidc.Nonce(nonce)), http.StatusFound)
}

func (c *LoginController) SSOCallback() {
	if !IsSSOEnabled() {
		c.Abort("404")
		return
	}
	if providerError := c.GetString("error"); providerError != "" {
		http.Error(c.Ctx.ResponseWriter, "OIDC login failed: "+providerError, http.StatusUnauthorized)
		return
	}

	stateCookie, err := c.Ctx.Request.Cookie(oidcStateCookieName)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "missing SSO state", http.StatusBadRequest)
		return
	}
	codec, err := ssoCookie(oidcStateTTL)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	var savedState oidcState
	if err := codec.Decode(oidcStateCookieName, stateCookie.Value, &savedState); err != nil ||
		savedState.ExpiresAt < time.Now().Unix() ||
		savedState.State != c.GetString("state") {
		http.Error(c.Ctx.ResponseWriter, "invalid SSO state", http.StatusBadRequest)
		return
	}

	oauthConfig, verifier, err := oidcConfig(c.Ctx.Request.Context())
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusServiceUnavailable)
		return
	}
	token, err := oauthConfig.Exchange(c.Ctx.Request.Context(), c.GetString("code"))
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "exchange authorization code: "+err.Error(), http.StatusUnauthorized)
		return
	}
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(c.Ctx.ResponseWriter, "OIDC provider response did not contain an ID token", http.StatusUnauthorized)
		return
	}
	idToken, err := verifier.Verify(c.Ctx.Request.Context(), rawIDToken)
	if err != nil || idToken.Nonce != savedState.Nonce {
		http.Error(c.Ctx.ResponseWriter, "invalid ID token", http.StatusUnauthorized)
		return
	}

	var claims map[string]interface{}
	if err := idToken.Claims(&claims); err != nil {
		http.Error(c.Ctx.ResponseWriter, "decode ID token claims: "+err.Error(), http.StatusUnauthorized)
		return
	}
	authorizationValues, err := authorizationValuesFromClaims(claims)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusForbidden)
		return
	}
	if !hasAllowedAuthorizationValue(authorizationValues) {
		http.Error(c.Ctx.ResponseWriter, "your OIDC account is not authorized for PrometheusAlert", http.StatusForbidden)
		return
	}

	subject, _ := claims["sub"].(string)
	if subject == "" {
		http.Error(c.Ctx.ResponseWriter, "OIDC ID token did not contain a subject", http.StatusUnauthorized)
		return
	}
	username, _ := claims["preferred_username"].(string)
	if username == "" {
		username, _ = claims["email"].(string)
	}
	if username == "" {
		username = subject
	}
	sessionTTL := ssoSessionTTL()
	sessionCodec, err := ssoCookie(sessionTTL)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	encodedSession, err := sessionCodec.Encode(oidcSessionCookieName, oidcSession{
		Subject:             subject,
		Username:            username,
		AuthorizationValues: authorizationValues,
		ExpiresAt:           time.Now().Add(sessionTTL).Unix(),
	})
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "encode SSO session: "+err.Error(), http.StatusInternalServerError)
		return
	}

	clearSSOCookies(c.Ctx.ResponseWriter)
	setSSOCookie(c.Ctx.ResponseWriter, oidcSessionCookieName, encodedSession, sessionTTL)
	c.Redirect("/", http.StatusFound)
}

func CheckSSOAccount(ctx *beecontext.Context) bool {
	_, ok := currentSSOSession(ctx)
	return ok
}

func SSOAccountName(ctx *beecontext.Context) string {
	session, ok := currentSSOSession(ctx)
	if !ok {
		return ""
	}
	return session.Username
}

func currentSSOSession(ctx *beecontext.Context) (*oidcSession, bool) {
	if !IsSSOEnabled() {
		return nil, false
	}
	cookie, err := ctx.Request.Cookie(oidcSessionCookieName)
	if err != nil {
		return nil, false
	}
	codec, err := ssoCookie(ssoSessionTTL())
	if err != nil {
		return nil, false
	}
	var session oidcSession
	if err := codec.Decode(oidcSessionCookieName, cookie.Value, &session); err != nil {
		return nil, false
	}
	if session.Subject == "" || session.ExpiresAt < time.Now().Unix() || !hasAllowedAuthorizationValue(session.AuthorizationValues) {
		return nil, false
	}
	return &session, true
}

func SSOLogoutRedirect() string {
	logoutURL := strings.TrimSpace(beego.AppConfig.String("sso_logout_url"))
	if logoutURL == "" {
		return ""
	}
	parsedURL, err := url.Parse(logoutURL)
	if err != nil {
		return ""
	}
	query := parsedURL.Query()
	if redirectURL := strings.TrimSpace(beego.AppConfig.String("sso_post_logout_redirect_url")); redirectURL != "" {
		query.Set("post_logout_redirect_uri", redirectURL)
	}
	if clientID := strings.TrimSpace(beego.AppConfig.String("sso_client_id")); clientID != "" {
		query.Set("client_id", clientID)
	}
	parsedURL.RawQuery = query.Encode()
	return parsedURL.String()
}
