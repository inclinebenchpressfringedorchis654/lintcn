// Content hash for binary caching.
// Combines tsgolint version, rule file contents, Go version, and platform
// into a single SHA-256 hash used as the cached binary filename.

import crypto from 'node:crypto'
import fs from 'node:fs'
import path from 'node:path'
import { execAsync } from './exec.ts'

export async function computeContentHash({
  lintcnDir,
  tsgolintVersion,
}: {
  lintcnDir: string
  tsgolintVersion: string
}): Promise<string> {
  const hash = crypto.createHash('sha256')

  hash.update(`tsgolint:${tsgolintVersion}\n`)
  hash.update(`platform:${process.platform}-${process.arch}\n`)

  // add Go version
  try {
    const { stdout } = await execAsync('go', ['version'])
    hash.update(`go:${stdout.trim()}\n`)
  } catch {
    hash.update('go:unknown\n')
  }

  // add all rule file contents in sorted order
  const files = fs
    .readdirSync(lintcnDir)
    .filter((f) => {
      return f.endsWith('.go')
    })
    .sort()

  for (const file of files) {
    const content = fs.readFileSync(path.join(lintcnDir, file), 'utf-8')
    hash.update(`file:${file}\n`)
    hash.update(content)
  }

  return hash.digest('hex').slice(0, 16)
}
