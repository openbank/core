package idp

import (
	"errors"
	"net/http"

	"bnk.to/core/tools/idp/auth"
	hydra "github.com/ory/hydra-client-go"
)

// genericErrMsg is the error message to return to prevent leaking implementation details.
const genericErrMsg = "An unknown error has occurred while processing the request."

// getLoginHandler handles the 'GET /login' endpoint.
func (s *Server) getLoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	challenge := r.URL.Query().Get("login_challenge")
	loginReq, _, err := s.hydra.AdminApi.GetLoginRequest(ctx).LoginChallenge(challenge).Execute()
	if err != nil {
		// TODO: Handle the case where the user passed an invalid challenge.
		s.logger.Err(err).Msg("GetLoginRequest")
		http.Error(w, genericErrMsg, http.StatusInternalServerError)
		return
	}
	if loginReq.Skip {
		id, err := s.provider.Lookup(ctx, loginReq.GetSubject())
		if err == nil {
			// The Skip is valid. Skip authentication.
			s.acceptLoginRequest(w, r, challenge, id)
			return
		}
	}
	s.showLoginForm(w, r, LoginFormParam{
		Challenge: challenge,
	})
}

// postLoginHandler handles the 'POST /login' endpoint.
func (s *Server) postLoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := r.ParseForm()
	if err != nil {
		s.logger.Err(err).Msg("ParseForm")
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	challenge := r.Form.Get("login_challenge")
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	loginErrors := make([]string, 0, 2)
	if username == "" {
		loginErrors = append(loginErrors, "Username must be provided")
	}
	if password == "" {
		loginErrors = append(loginErrors, "Password must be provided")
	}
	if len(loginErrors) > 0 {
		s.logger.Debug().Strs("login_errors", loginErrors).Msg("POST login fail")
		s.showLoginForm(w, r, LoginFormParam{
			Challenge: challenge,
			Errors:    loginErrors,
		})
		return
	}
	id, err := s.provider.Authenticate(ctx, auth.Challenge{
		Username: username,
		Password: password,
	})
	switch {
	case errors.Is(err, auth.ErrInvalidCredentials):
		loginErrors = append(loginErrors, "Invalid username and/or password")
		s.logger.Debug().Strs("login_errors", loginErrors).Msg("POST login fail")
		s.showLoginForm(w, r, LoginFormParam{
			Challenge: challenge,
			Errors:    loginErrors,
		})
	case err != nil:
		s.logger.Err(err).Msg("provider.Authenticate")
		http.Error(w, genericErrMsg, http.StatusInternalServerError)
	default:
		s.logger.Debug().Str("challenge", challenge).Msg("AcceptLoginRequest")
		s.acceptLoginRequest(w, r, challenge, id)
	}
}

// acceptLoginRequest approves the login request with the specified challenge
// code and redirects the user.
func (s *Server) acceptLoginRequest(w http.ResponseWriter, r *http.Request, challenge string, id *auth.Identity) {
	ctx := r.Context()
	body := hydra.NewAcceptLoginRequest(id.UserID)
	acceptResult, _, err := s.hydra.AdminApi.AcceptLoginRequest(ctx).
		AcceptLoginRequest(*body).
		LoginChallenge(challenge).
		Execute()
	if err != nil {
		s.logger.Err(err).Msg("acceptLoginRequest")
		http.Error(w, genericErrMsg, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, acceptResult.RedirectTo, http.StatusFound)
}

// showLoginForm shows the login form to the user with the provided parameters.
func (s *Server) showLoginForm(w http.ResponseWriter, r *http.Request, param LoginFormParam) {
	w.Header().Set("Content-Type", "text/html")
	err := tpl.ExecuteTemplate(w, "login.html", param)
	if err != nil {
		s.logger.Err(err).Msg("ExecuteTemplate login.html")
		http.Error(w, genericErrMsg, http.StatusInternalServerError)
		return
	}
}

// LoginFormParam is the parameters of the login.html template.
type LoginFormParam struct {
	Challenge string
	Errors    []string
}
