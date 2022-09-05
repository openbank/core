package idp

import (
	"net/http"

	hydra "github.com/ory/hydra-client-go"
)

// getConsentHandler handles the 'GET /consent' endpoint.
func (s *Server) getConsentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	challenge := r.URL.Query().Get("consent_challenge")
	consentReq, _, err := s.hydra.AdminApi.GetConsentRequest(ctx).ConsentChallenge(challenge).Execute()
	if err != nil {
		// TODO: Handle the case where the user passed an invalid challenge.
		s.logger.Err(err).Msg("GetConsentRequest")
		http.Error(w, genericErrMsg, http.StatusInternalServerError)
		return
	}
	// Always accept consent requests.
	body := hydra.NewAcceptConsentRequest()
	body.SetGrantScope(consentReq.RequestedScope)
	acceptResult, _, err := s.hydra.AdminApi.AcceptConsentRequest(ctx).
		AcceptConsentRequest(*body).
		ConsentChallenge(challenge).
		Execute()
	if err != nil {
		s.logger.Err(err).Msg("AcceptConsentRequest")
		http.Error(w, genericErrMsg, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, acceptResult.RedirectTo, http.StatusFound)
}
