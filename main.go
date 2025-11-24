package main

import (
	"github.com/SoeldnerConsult/tofulint-plugin-sdk/plugin"
	"github.com/SoeldnerConsult/tofulint-plugin-sdk/tflint"
	"github.com/arsiba/tofulint-ruleset-aws/aws"
	"github.com/arsiba/tofulint-ruleset-aws/project"
	"github.com/arsiba/tofulint-ruleset-aws/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &aws.RuleSet{
			BuiltinRuleSet: tflint.BuiltinRuleSet{
				Name:    "aws",
				Version: project.Version,
				Rules:   rules.Rules,
			},
		},
	})
}
