package main

import (
	pwl "github.com/justjanne/powerline-go/powerline"
	"os"
)

func segmentDotEnv(p *powerline) {
	files := []string{".env", ".envrc"}
	dotEnv := false
	for _, file := range files {
		stat, err := os.Stat(file)
		if err == nil && !stat.IsDir() {
			dotEnv = true
			break
		}
	}
	if dotEnv {
		p.appendSegment("dotenv", pwl.Segment{
			Content:    "\u2235",
			Foreground: p.theme.DotEnvFg,
			Background: p.theme.DotEnvBg,
		})
	}
}
