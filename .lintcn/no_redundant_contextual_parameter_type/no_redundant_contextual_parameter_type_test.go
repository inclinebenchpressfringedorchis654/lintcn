package no_redundant_contextual_parameter_type

import (
	"testing"

	"github.com/typescript-eslint/tsgolint/internal/rule_tester"
	"github.com/typescript-eslint/tsgolint/internal/rules/fixtures"
)

func TestNoRedundantContextualParameterType(t *testing.T) {
	t.Parallel()
	rule_tester.RunRuleTester(
		fixtures.GetRootDir(),
		"tsconfig.minimal.json",
		t,
		&NoRedundantContextualParameterTypeRule,
		validCases,
		invalidCases,
	)
}

var validCases = []rule_tester.ValidTestCase{
	{
		Code: `
			type Context = { command: string }
			declare function tool(config: { execute: (ctx: Context) => Promise<void> }): void
			tool({
				execute: async (ctx) => {
					ctx.command
				},
			})
		`,
	},
	{
		Code: `
			const execute = async (ctx: { command: string }) => {
				ctx.command
			}
		`,
	},
	{
		Code: `
			declare function useRunner(fn: (ctx: { command: string; cwd: string }) => void): void
			useRunner((ctx: { command: string }) => {
				ctx.command
			})
		`,
	},
	{
		Code: `
			declare function withValue<T>(fn: (value: T) => void): void
			withValue((value: string) => {
				value.toUpperCase()
			})
		`,
	},
	{
		Code: `
			const obj = {
				execute(ctx: { command: string }) {
					return ctx.command
				},
			}
		`,
	},
}

var invalidCases = []rule_tester.InvalidTestCase{
	{
		Code: `
			type Context = { command: string }
			declare const bash: { exec(command: string): Promise<{ stdout: string; stderr: string; exitCode: number }> }
			declare function tool(config: { execute: (ctx: Context) => Promise<{ stdout: string; stderr: string; exitCode: number }> }): void

			tool({
				execute: async (ctx: { command: string }): Promise<{
					stdout: string
					stderr: string
					exitCode: number
				}> => {
					const result = await bash.exec(ctx.command)
					return {
						stdout: result.stdout,
						stderr: result.stderr,
						exitCode: result.exitCode,
					}
				},
			})
		`,
		Errors: []rule_tester.InvalidTestCaseError{{MessageId: "redundantContextualParameterType"}},
	},
	{
		Code: `
			declare function useRunner(fn: (ctx: { command: string }) => void): void
			useRunner((ctx: { command: string }) => {
				ctx.command
			})
		`,
		Errors: []rule_tester.InvalidTestCaseError{{MessageId: "redundantContextualParameterType"}},
	},
	{
		Code: `
			type Context = { command: string }
			declare function useRunner(fn: (ctx: Context) => void): void
			useRunner((ctx: Context) => {
				ctx.command
			})
		`,
		Errors: []rule_tester.InvalidTestCaseError{{MessageId: "redundantContextualParameterType"}},
	},
	{
		Code: `
			declare const wrapped: (ctx: { command: string }) => void
			const run: typeof wrapped = (ctx: { command: string }) => {
				ctx.command
			}
		`,
		Errors: []rule_tester.InvalidTestCaseError{{MessageId: "redundantContextualParameterType"}},
	},
}
