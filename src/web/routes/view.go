package routes

import (
	"net/http"
	"root/src/lib/libblockchain"
	"root/src/lib/libresponse"
	"root/src/lib/libtemplate"

	"github.com/go-chi/chi/v5"
)

func Initialize(engine libtemplate.TemplateEngine) chi.Router {
	testchain := libblockchain.BlockChainLite{}.New()
	for i := 1; i < 101; i++ {
		curBlock, _ := libblockchain.BlockLite{}.New(80+i, testchain.Tail)
		testchain.Push(&curBlock)
	}

	renderedBlockchainWC := engine.RenderWithLogs("blockchain.wc", map[string]string{
		"blockchain": testchain.Render(),
	})
	blockChainPage := map[string]string{
		"title": "Blockchain Visual Playground",
		"head":  engine.RenderWithLogs("blockchain.header.wc", nil),
		"body":  renderedBlockchainWC,
	}

	router := chi.NewRouter()
	router.Get("/blockchain", func(w http.ResponseWriter, r *http.Request) {
		rendered := engine.RenderWithLogs("index.layout.html", blockChainPage)
		libresponse.WithHTML(w, http.StatusOK, rendered)
	})

	return router
}
