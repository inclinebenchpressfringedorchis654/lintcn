# lintcn

## tsgolint fork

lintcn uses `remorses/tsgolint` (forked from `oxc-project/tsgolint`).

The fork adds 1 commit: `internal/runner/runner.go` with `Run(rules, args)`.
Zero modifications to existing files. Upstream merges should never conflict.

User rules import from `internal/rule`, `internal/utils` etc. — same paths
as tsgolint's own code. The Go workspace allows this because the user module
name is a child path: `github.com/typescript-eslint/tsgolint/lintcn-rules`.

## updating tsgolint

Two constants in `src/cache.ts`:

- `DEFAULT_TSGOLINT_VERSION` — commit hash from `remorses/tsgolint`
- `TYPESCRIPT_GO_COMMIT` — base commit from `microsoft/typescript-go`
  (before patches). Changes only when upstream updates its submodule.

To sync: merge upstream into fork, push, update both constants, clear
`~/.cache/lintcn`, rebuild, test.

typescript-go is managed by tsgolint — never fork it independently.
