## 0.6.0

1. **Rules now live in subfolders** — each rule is its own Go package under `.lintcn/{rule_name}/`, replacing the old flat `.lintcn/*.go` layout. This eliminates the need to rename `options.go` and `schema.json` companions — they stay in the subfolder with their original names, and the Go package name matches the folder. `lintcn add` now fetches the entire rule folder automatically.

   ```
   .lintcn/
       no_floating_promises/
           no_floating_promises.go
           no_floating_promises_test.go
           options.go      ← original name, no renaming
           schema.json
       my_custom_rule/
           my_custom_rule.go
   ```

2. **`lintcn add` fetches whole folders** — both folder URLs (`/tree/`) and file URLs (`/blob/`) now fetch every `.go` and `.json` file in the rule's directory. Passing a file URL auto-detects the parent folder:

   ```bash
   # folder URL
   lintcn add https://github.com/oxc-project/tsgolint/tree/main/internal/rules/no_floating_promises

   # file URL — auto-fetches the whole folder
   lintcn add https://github.com/oxc-project/tsgolint/blob/main/internal/rules/await_thenable/await_thenable.go
   ```

3. **Error for flat `.go` files in `.lintcn/`** — if leftover flat files from older versions are detected, lintcn now prints a clear migration error with instructions instead of silently ignoring them.

4. **Reproducible builds with `-trimpath`** — the Go binary is now built with `-trimpath`, stripping absolute paths from the output. Binaries are identical across machines for the same rule content + tsgolint version + platform.

5. **Faster cache hits** — Go version removed from the content hash. The compiled binary is a standalone executable with no Go runtime dependency, so the Go version used to build it doesn't affect correctness. Also excludes `_test.go` files from the hash since tests don't affect the binary.

6. **Go compilation output is live** — `go build` now inherits stdio, so compilation progress and errors stream directly to the terminal instead of being silently captured.

7. **First-build guidance** — on first compile (cold Go cache), lintcn explains the one-time 30s cost and shows which directories to cache in CI:
   ```
   Compiling custom tsgolint binary (first build — may take 30s+ to compile dependencies)...
   Subsequent builds will be fast (~1s). In CI, cache ~/.cache/lintcn/ and GOCACHE (run `go env GOCACHE`).
   ```

8. **GitHub Actions example** — README now includes a copy-paste workflow that caches the compiled binary. Subsequent CI runs take ~12s (vs ~4min cold):

   ```yaml
   - name: Cache lintcn binary + Go build cache
     uses: actions/cache@v4
     with:
       path: |
         ~/.cache/lintcn
         ~/go/pkg
       key: lintcn-${{ runner.os }}-${{ runner.arch }}-${{ hashFiles('.lintcn/**/*.go') }}
       restore-keys: lintcn-${{ runner.os }}-${{ runner.arch }}-
   ```

## 0.5.0

1. **Security fix — path traversal in `--tsgolint-version`** — the version flag is now validated against a strict pattern. Previously a value like `../../etc` could escape the cache directory.

2. **Fixed intermittent failures with concurrent `lintcn lint` runs** — build workspaces are now per-content-hash instead of shared. Two processes running simultaneously no longer corrupt each other's build.

3. **Cross-platform tar extraction** — replaced shell `tar` command with the npm `tar` package. Works on Windows without needing system tar.

4. **No more `patch` command required** — tsgolint downloads now use a fork with a clean `internal/runner.Run()` entry point. Zero modifications to existing tsgolint files; upstream syncs will never conflict.

5. **Downloads no longer hang** — 120s timeout on all GitHub tarball downloads.

6. **Fixed broken `.tsgolint` symlink** — `lintcn add` now correctly detects and recreates broken symlinks.

## 0.4.0

## 0.3.0

1. **Only custom rules run by default** — previously the binary included all 44 built-in tsgolint rules, producing thousands of noisy errors. Now only your `.lintcn/` rules run. True shadcn model: explicitly add each rule you want.

   Before (0.2.0): `Found 2315 errors (linted 193 files with 45 rules)`
   After (0.3.0): `Found 8 errors (linted 193 files with 1 rule)`

2. **Run `lintcn lint` from any subdirectory** — uses `find-up` to walk parent directories for `.lintcn/`. You no longer need to be at the project root:
   ```bash
   cd packages/my-app
   lintcn lint   # finds .lintcn/ in parent
   ```

3. **No git required** — tsgolint source is now downloaded as a tarball from GitHub instead of cloned. Patches applied with `patch -p1`. Faster first setup, works without git installed.

4. **Fixed stale binary cache** — added `CACHE_SCHEMA_VERSION` to the content hash. Upgrading lintcn now correctly invalidates cached binaries built by older versions.

5. **Fixed partial download corruption** — if the tsgolint download fails midway, the partial directory is cleaned up so the next run starts fresh instead of failing repeatedly.

6. **Fixed GitHub URLs with `/` in branch names** — `lintcn add` now correctly handles branch names like `feature/my-branch` in GitHub blob URLs.

## 0.2.0

1. **Pinned tsgolint version** — each lintcn release bundles a specific tsgolint version (`v0.9.2`). Builds are now reproducible: everyone on the same lintcn version compiles against the same tsgolint API. Previously used `main` branch which was non-deterministic.

2. **`--tsgolint-version` flag** — override the pinned version for testing unreleased tsgolint:
   ```bash
   npx lintcn lint --tsgolint-version v0.10.0
   ```

3. **Version pinning docs** — README now explains why you should pin lintcn in `package.json` (no `^` or `~`) and how to update safely.

## 0.1.0

1. **Initial release** — CLI for adding type-aware TypeScript lint rules as Go files to your project:

   ```bash
   npx lintcn add https://github.com/user/repo/blob/main/rules/no_unhandled_error.go
   npx lintcn lint
   ```

2. **`lintcn add <url>`** — fetch a `.go` rule file by URL into `.lintcn/`. Normalizes GitHub blob URLs to raw URLs automatically. Also fetches the matching `_test.go` if present. Rewrites the package declaration to `package lintcn` and injects a `// lintcn:source` comment.

3. **`lintcn lint`** — builds a custom tsgolint binary (all 50+ built-in rules + your custom rules) and runs it against the project. Binary is cached by SHA-256 content hash — rebuilds only when rules change.

4. **`lintcn build`** — build the custom binary without running it. Prints the binary path.

5. **`lintcn list`** — list installed rules with descriptions parsed from `// lintcn:` metadata comments.

6. **`lintcn remove <name>`** — delete a rule and its test file from `.lintcn/`.

7. **Editor/LSP support** — generates `go.work` and `go.mod` inside `.lintcn/` so gopls provides full autocomplete, go-to-definition, and type checking on tsgolint APIs while writing rules.
