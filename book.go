// Copyright (c) 2025, The CCN Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"

	"cogentcore.org/core/content"
	"cogentcore.org/core/core"
	_ "cogentcore.org/core/text/tex" // include this to get math
	"cogentcore.org/core/tree"
)

// NOTE: you must make a symbolic link to the zotero CCNLab CSL file as ccnlab.json
// in this directory, to generate references and have the generated reference links
// use the official APA style. https://www.zotero.org/groups/340666/ccnlab
// Must configure using BetterBibTeX for zotero: https://retorque.re/zotero-better-bibtex/
// todo: include link for configuring here

//go:generate mdcite -vv -refs ./ccnlab.json -d ./content

//go:embed content
var econtent embed.FS

//go:embed icon.svg
var icon string

func main() {
	core.AppIcon = icon
	content.Settings.SiteTitle = "Computational Cognitive Neuroscience"
	content.OfflineURL = "https://book.compcogneuro.org"
	b := core.NewBody(content.Settings.SiteTitle)
	ct := content.NewContent(b).SetContent(econtent)
	// refs, err := csl.OpenFS(econtent, "citedrefs.json")
	// if err == nil {
	// 	ct.References = csl.NewKeyList(refs)
	// }
	b.AddTopBar(func(bar *core.Frame) {
		core.NewToolbar(bar).Maker(func(p *tree.Plan) {
			ct.MakeToolbar(p)
			ct.MakeToolbarPDF(p)
		})
	})
	b.RunMainWindow()
}
