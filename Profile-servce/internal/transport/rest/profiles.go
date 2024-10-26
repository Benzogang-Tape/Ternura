package rest

import (
	"context"
	"encoding/json"
	"github.com/Benzogang-Tape/Ternura/Profile-servce/internal/domain"
	"go.uber.org/zap"
	"net/http"
)

type ProfileApi interface {
	GetAllProfiles(ctx context.Context) ([]*domain.UserProfile, error)
	GetSuitableProfiles(ctx context.Context, profile domain.UserProfile) ([]*domain.UserProfile, error)
	GetProfilesByGender(ctx context.Context, gender string) ([]*domain.UserProfile, error)
	GetProfileByID(ctx context.Context, uuid string) (*domain.UserProfile, error)
}

type ProfileHandler struct {
	logger  *zap.SugaredLogger
	service ProfileApi
}

func NewProfileAPI(profileAPi ProfileApi, logger *zap.SugaredLogger) *ProfileHandler {
	return &ProfileHandler{
		logger:  logger,
		service: profileAPi,
	}
}

func (p *ProfileHandler) GetSuggestedProfiles(w http.ResponseWriter, r *http.Request) {

}

func (p *ProfileHandler) GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	profiles, err := p.service.GetAllProfiles(r.Context())
	if err != nil {
		jsonSimpleErr(w, http.StatusInternalServerError, domain.NewSimpleErr(domain.ErrUnknownError.Error()))
		return
	}

	resp, err := json.Marshal(profiles)
	if _, err = w.Write(resp); err != nil {
		jsonSimpleErr(w, http.StatusInternalServerError, domain.NewSimpleErr(domain.ErrResponseError.Error()))
	}
	if _, err = w.Write(resp); err != nil {
		jsonSimpleErr(w, http.StatusInternalServerError, domain.NewSimpleErr(domain.ErrResponseError.Error()))
	}
}

func (p *ProfileHandler) GetProfileByID(w http.ResponseWriter, r *http.Request) {

}

func (p *ProfileHandler) GetProfilesByGender(w http.ResponseWriter, r *http.Request) {

}