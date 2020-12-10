package response

import "go-deck/app/model/entity"

type LoginResponse struct {
	User      entity.User `json:"user"`
	Token     string      `json:"token"`
	ExpiresAt int64       `json:"expiresAt"`
}
