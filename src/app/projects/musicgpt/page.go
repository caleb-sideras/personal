package musicgpt

import (
	"net/http"

	"calebsideras.com/temporary/src/components/server"
	"github.com/a-h/templ"
)

func Page(w http.ResponseWriter, r *http.Request) templ.Component {

	return server.ProjectTab(server.ProjectTabType{
		Title:       "MusicGPT",
		ProjectURL:  "https://music-gpt.xyz",
		ReadMeURL:   "/projects/musicgpt/readme",
		VideoURL:    "/projects/musicgpt/videos",
		InitialBody: Readme(),
	})
}
