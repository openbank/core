package idp

import "net/http"

func (s *Server) getLogoutHandler(w http.ResponseWriter, r *http.Request) {
	challenge := r.URL.Query().Get("logout_challenge")
	ctx := r.Context()
	_, _, err := s.hydra.AdminApi.GetLogoutRequest(ctx).LogoutChallenge(challenge).Execute()
	if err != nil {
		// TODO: Handle the case where the user passed an invalid challenge.
		s.logger.Err(err).Msg("GetLogoutRequest")
		http.Error(w, genericErrMsg, http.StatusInternalServerError)
		return
	}
	// Always accept the logout request.
	acceptResult, _, err := s.hydra.AdminApi.AcceptLogoutRequest(ctx).LogoutChallenge(challenge).Execute()
	if err != nil {
		s.logger.Err(err).Msg("AcceptLogoutRequest")
		http.Error(w, genericErrMsg, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, acceptResult.RedirectTo, http.StatusFound)
}
