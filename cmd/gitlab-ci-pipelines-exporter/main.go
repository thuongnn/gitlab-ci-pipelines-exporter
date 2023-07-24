package main

import (
	"os"

	"github.com/thuongnn/gitlab-ci-pipelines-exporter/internal/cli"
)

var version = "devel"

func main() {
	cli.Run(version, os.Args)
}
