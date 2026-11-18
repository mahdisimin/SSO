package auth

import (
	"SSO/config"
	auth2 "SSO/services/auth"
	"SSO/session"
	"SSO/user"
	"html/template"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v5"
)

func TestLoginHandler(t *testing.T) {

	err := config.LoadDotEnv(".env")
	if err != nil {
		log.Fatal("error on loading env variables : ", err)
	}
	cfg, confErr := config.Load()
	if confErr != nil {
		t.Fatal("error on loading config : ", confErr)
	}

	e := echo.New()
	form := strings.NewReader("username=mahdi&password=1234")
	req := httptest.NewRequest(http.MethodPost, "/login", form)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userStore := user.NewMemoryStore()
	sessionStore := session.NewMemoryStore()
	auth := auth2.NewAuthService(userStore, sessionStore, time.Hour)

	userIns := user.NewUser("mahdi", "test@gmail.com", "1234")
	if err := userStore.Create(userIns); err != nil {
		t.Fatalf("error on create user : %v", err)
	}

	tmpl := template.Must(template.New("login.html").Parse(`{{.Error}}`))

	handler := LoginHandler(auth, cfg, tmpl)
	errHandler := handler(c)

	if errHandler != nil {
		t.Error("error on login : ", errHandler)
	}

	if rec.Code != http.StatusSeeOther {
		t.Fatalf("expecting %d, got %d", http.StatusSeeOther, rec.Code)
	}

	if rec.Header().Get("Location") != "/me" {
		t.Errorf("expecting %s, got %s", "/me", rec.Header().Get("Location"))
	}
	coockie := rec.Header().Get("Set-Cookie")

	if !(strings.Contains(coockie, "HttpOnly") || strings.Contains(coockie, "SameSite=Lax")) {
		t.Errorf("expecting Set-Cookie header got %s", coockie)
	}
}
