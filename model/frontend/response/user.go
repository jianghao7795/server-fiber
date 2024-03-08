package response

import "server-fiber/model/frontend"

type LoginResponse struct {
	User      frontend.User `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}
