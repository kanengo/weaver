package generate

import "github.com/ServiceWeaver/weaver/internal/tool/generate"

type Options = generate.Options

var Usage = generate.Usage

func Generate(dir string, pkgs []string, opt Options) error {
	return generate.Generate(dir, pkgs, opt)
}
