package caas_segmentation

import (
	"net/http"

	"calebsideras.com/temporary/src/components/server"
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

func Page(w http.ResponseWriter, r *http.Request) templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/projects/_markdown/caas_segmentation.md")

	if err != nil {
		panic(err)
	}

	return server.Project(server.ProjectType{
		Title:      "CAAS",
		ProjectURL: "https://github.com/caleb-sideras/CAAS/blob/main/AudioSegmentation.py",
		Body:       newTempl,
	})
}
