package projects

import (
	"calebsideras.com/temporary/src/components/server"
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

func Page_() templ.Component {

	tLinkList := []utils.TLinkContent{
		{
			Title:       "Temporary Framework",
			Description: "The HTMX and Templ Framework you are looking for.",
			Src:         "https://temporary-framework.org/static/assets/temporary.png",
			Alt:         "Temporary Image",
			Href:        "/projects/temporary-framework",
			Boost:       "true",
		},
		{
			Title:       "MusicGPT",
			Description: "Multimodal AI that generates custom visualizations and facilitates natural language dialogue about technical aspects of a song - spanning from its musical structure to its production features.",
			Src:         "https://music-gpt.xyz/musicgpt.png",
			Alt:         "MusicGPT Image",
			Href:        "/projects/musicgpt",
			Boost:       "true",
		},
		{
			Title:       "GoX Framework",
			Description: "A framework designed to make working with Go and HTMX easier. Achieves this by employing certain primitives that tightly couple them, aiding scaling.",
			Src:         "https://gox-framework.org/static/assets/gox-mascot-hor.png",
			Href:        "/projects/gox-framework",
			Alt:         "GoX Framework Moscot",
			Boost:       "true",
		},
		{
			Title:       "CAAS",
			Description: "CAAS is an audio segmentation algorithm that iteratively finds the best intervals and patterns in a given audio signal, clusters them based on similarity and then extracts relevant audio features (Lyrics, MIDI etc).",
			Src:         "https://opengraph.githubassets.com/1a43d7c7f54d054a1094b6f53a76668f5d192eafd3f041c978cce36355cf3006/caleb-sideras/CAAS",
			Href:        "/projects/caas",
			Alt:         "CAAS Github",
			Boost:       "true",
		},
		{
			Title:       "People Pedia",
			Description: "AI powered search that summarizes and visualizes information for anyone who has data on the internet.",
			Src:         "https://opengraph.githubassets.com/be1d1b6657697642586856c35f9916242a1108d34d8cee509b61c2ca5b5e2e56/caleb-sideras/PeoplePedia-Backend",
			Href:        "/projects/peoplepedia",
			Alt:         "People Pedia Github",
			Boost:       "true",
		},
		{
			Title:       "Tweetailyze",
			Description: "A web app that performs Twitter account summarization through tweet embedding, clustering, sentiment analysis and topic modeling.",
			Src:         "https://opengraph.githubassets.com/bc52bfa035af32496e45f674130251555d2549ca38d8b333461df836979cebbb/caleb-sideras/tweetailyze-frontend",
			Href:        "/projects/tweetailyze",
			Alt:         "Tweetailyze Github",
			Boost:       "true",
		},
		{
			Title:       "Temporary UI",
			Description: "A UI library built on top of Material 3.",
			Src:         "/static/assets/coming_soon.jpg",
			Alt:         "Coming Soon",
			Disabled:    true,
		},
	}

	return server.Grid(tLinkList)

}
