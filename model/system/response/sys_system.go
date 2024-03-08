package response

import "server-fiber/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
