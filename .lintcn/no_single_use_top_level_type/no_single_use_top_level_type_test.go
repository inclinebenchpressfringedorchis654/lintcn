// Tests for the no-single-use-top-level-type rule.
package no_single_use_top_level_type

import (
	"testing"

	"github.com/typescript-eslint/tsgolint/internal/rule_tester"
	"github.com/typescript-eslint/tsgolint/internal/rules/fixtures"
)

func TestNoSingleUseTopLevelType(t *testing.T) {
	t.Parallel()
	rule_tester.RunRuleTester(
		fixtures.GetRootDir(),
		"tsconfig.minimal.json",
		t,
		&NoSingleUseTopLevelTypeRule,
		validCases,
		invalidCases,
	)
}

var validCases = []rule_tester.ValidTestCase{
	{Code: `
		type User = { name: string }
		const a: User = { name: 'a' }
		const b: User = { name: 'b' }
	`},
	{Code: `
		type User = { name: string }
	`},
	{Code: `
		interface User { name: string }
		function print(user: User) {
			return user.name
		}
		const user: User = { name: 'a' }
		print(user)
	`},
	{Code: `
		type User = { name: string }
		const pair: [User, User] = [{ name: 'a' }, { name: 'b' }]
	`},
	{Code: `
		function build() {
			type Local = { name: string }
			const user: Local = { name: 'a' }
			return user
		}
	`},
	{Code: `
		declare interface User { name: string }
		const user: User = { name: 'a' }
	`},
	{Code: `
		export type User = { name: string }
		const user: User = { name: 'a' }
	`},
	{Code: `
		export interface User { name: string }
		const user: User = { name: 'a' }
	`},
	{Code: `
		export default interface User { name: string }
		const user: User = { name: 'a' }
	`},
	{Code: `
		type User = { name: string }
		export type { User }
		const user: User = { name: 'a' }
	`},
	{Code: `
		type User = { name: string }
		export { User }
		const user: User = { name: 'a' }
	`},
	{Code: `
		type User = { name: string }
		export type { User as PublicUser }
		const user: User = { name: 'a' }
	`},
	{Code: `
		type User = { name: string }
		export { User as PublicUser }
		const user: User = { name: 'a' }
	`},
	{Code: `
		type Tree = { value: string; children: Tree[] }
		const root: Tree = { value: 'a', children: [] }
	`},
	{Code: `
		type User = {
			id: string
			name: string
			email: string
			avatarUrl: string
			roles: Array<'admin' | 'editor' | 'viewer'>
			preferences: {
				theme: 'light' | 'dark'
				locale: string
				timezone: string
			}
		}

		const user: User = {
			id: '1',
			name: 'a',
			email: 'a@example.com',
			avatarUrl: '/a.png',
			roles: ['viewer'],
			preferences: { theme: 'dark', locale: 'en', timezone: 'UTC' },
		}
	`},
	{Code: `
		interface Tree { value: string; children: Tree[] }
		const root: Tree = { value: 'a', children: [] }
	`},
	{Code: `
		interface Config {
			cwd: string
			entry: string
			outDir: string
			plugins: string[]
			env: {
				mode: 'dev' | 'prod'
				region: string
				debug: boolean
			}
			cache: {
				enabled: boolean
				directory: string
			}
		}

		const config: Config = {
			cwd: process.cwd(),
			entry: 'src/index.ts',
			outDir: 'dist',
			plugins: [],
			env: { mode: 'dev', region: 'local', debug: true },
			cache: { enabled: true, directory: '.cache' },
		}
	`},
	{
		Code: `
			type User = { name: string }
			export type { User }
		`,
		Files: map[string]string{
			"consumer.ts": `
				import type { User } from './file'
				const user: User = { name: 'a' }
			`,
		},
	},
	{
		Code: `
			type User = { name: string }
			export { User }
		`,
		Files: map[string]string{
			"consumer.ts": `
				import type { User } from './file'
				const user: User = { name: 'a' }
			`,
		},
	},
	{
		Code: `
			type User = { name: string }
			export type { User as PublicUser }
		`,
		Files: map[string]string{
			"consumer.ts": `
				import type { PublicUser } from './file'
				const user: PublicUser = { name: 'a' }
			`,
		},
	},
	{
		Code: `
			type User = { name: string }
			export { User as PublicUser }
		`,
		Files: map[string]string{
			"consumer.ts": `
				import type { PublicUser } from './file'
				const user: PublicUser = { name: 'a' }
			`,
		},
	},
}

var invalidCases = []rule_tester.InvalidTestCase{
	{
		Code: `
			type User = { name: string }
			const user: User = { name: 'a' }
		`,
		Errors: []rule_tester.InvalidTestCaseError{
			{MessageId: "singleUseTopLevelType"},
		},
	},
	{
		Code: `
			interface User { name: string }
			const user: User = { name: 'a' }
		`,
		Errors: []rule_tester.InvalidTestCaseError{
			{MessageId: "singleUseTopLevelType"},
		},
	},
	{
		Code: `
			type UserId = string
			type User = { id: UserId }
			const user: User = { id: 'a' }
		`,
		Errors: []rule_tester.InvalidTestCaseError{
			{MessageId: "singleUseTopLevelType"},
			{MessageId: "singleUseTopLevelType"},
		},
	},
	{
		Code: `
			type User = { name: string }
			function print(user: User) {
				return user.name
			}
			print({ name: 'a' })
		`,
		Errors: []rule_tester.InvalidTestCaseError{
			{MessageId: "singleUseTopLevelType"},
		},
	},
	{
		Code: `
			type User = { name: string }
			const user = { name: 'a' } satisfies User
		`,
		Errors: []rule_tester.InvalidTestCaseError{
			{MessageId: "singleUseTopLevelType"},
		},
	},
	{
		Code: `
			interface BaseOptions { cwd: string }
			interface Options extends BaseOptions { mode: 'dev' }
			const options: Options = { cwd: '', mode: 'dev' }
		`,
		Errors: []rule_tester.InvalidTestCaseError{
			{MessageId: "singleUseTopLevelType"},
			{MessageId: "singleUseTopLevelType"},
		},
	},
	{
		Code: `
			type User = { name: string }
			type Box<T extends User> = { value: T }
			const box: Box<{ name: 'a' }> = { value: { name: 'a' } }
		`,
		Errors: []rule_tester.InvalidTestCaseError{
			{MessageId: "singleUseTopLevelType"},
			{MessageId: "singleUseTopLevelType"},
		},
	},
}
