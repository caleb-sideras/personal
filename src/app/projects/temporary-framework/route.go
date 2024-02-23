package temporaryframework

import (
	"calebsideras.com/temporary/src/components/server"
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

func Readme() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/projects/_markdown/temporary_framework.md")
	if err != nil {
		panic(err)
	}

	return newTempl
}

func Videos() templ.Component {
	tLinkList := []utils.TLinkContent{
		{
			Title:       "Temporary Framework Deep Dive",
			Description: "Temporary is a Go framework for building full-stack web applications using Templ and HTMX. It tightly couples these technologies together and introduces other primitives to help you build web-applications easier and faster. Under the hood Temporary abstracts routing, building, serving and more.",
			Src:         "https://img.youtube.com/vi/sQeqaKnk8oc/hqdefault.jpg",
			Alt:         "Temporary Thumbnail",
			Href:        "https://youtu.be/sQeqaKnk8oc?si=_ilvEk3oKjNymEek",
			Boost:       "false",
		},
	}
	return server.Grid(tLinkList)
}
