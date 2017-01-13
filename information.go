package gorocket

import (
	"net/http"
)

type infoResponse struct {
	Info Info `json:"info"`
}

type Commit struct {
	Hash string `json:"hash"`
	Date string `json:"date"`
	Author string `json:"author"`
	Subject string `json:"subject"`
	Tag string `json:"tag"`
	Branch string `json:"branch"`
}

type Info struct {
	Version string `json:"version"`
	Build Build `json:"build"`
	Travis Travis `json:"travis"`
	Commit Commit `json:"commit"`
	GraphicsMagick Library `json:"GraphicsMagick"`
	ImageMagick Library `json:"ImageMagick"`
}

type Travis struct {
	BuildNumber string `json:"buildNumber"`
	Branch string `json:"branch"`
	Tag string `json:"tag"`
}

type Build struct {
	Date string `json:"date"`
	NodeVersion string `json:"nodeVersion"`
	Arch string `json:"arch"`
	Platform string `json:"platform"`
	OsRelease string `json:"osRelease"`
	TotalMemory int64 `json:"totalMemory"`
	FreeMemory int64 `json:"freeMemory"`
	CpuCount int `json:"cpus"`
}

type Library struct {
	Enabled bool `json:"enabled"`
	Version string `json:"version"`
}

// Get information about the server.
// This function does not need a logged in user.
//
// https://rocket.chat/docs/developer-guides/rest-api/miscellaneous/info
func (r *Rocket) GetServerInfo() (*Info, error) {
	request, _ := http.NewRequest("GET", r.getUrl() + "/api/v1/info", nil)

	response := new(infoResponse)

	if err := r.doRequest(request, response); err != nil {
		return nil, err
	}

	return &response.Info, nil
}