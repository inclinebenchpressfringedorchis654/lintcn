// lintcn clean — remove cached tsgolint source and compiled binaries.
// Frees disk space from old versions that accumulate over time.

import fs from 'node:fs'
import { getCacheDir } from '../cache.ts'

export function clean(): void {
  const cacheDir = getCacheDir()

  if (!fs.existsSync(cacheDir)) {
    console.log('No cache to clean')
    return
  }

  const stats = getCacheStats(cacheDir)
  fs.rmSync(cacheDir, { recursive: true })
  console.log(`Removed ${cacheDir} (${formatBytes(stats.totalBytes)})`)
}

function getCacheStats(dir: string): { totalBytes: number } {
  let totalBytes = 0
  const walk = (d: string): void => {
    for (const entry of fs.readdirSync(d, { withFileTypes: true })) {
      const fullPath = `${d}/${entry.name}`
      if (entry.isDirectory()) {
        walk(fullPath)
      } else {
        totalBytes += fs.statSync(fullPath).size
      }
    }
  }
  try {
    walk(dir)
  } catch {
    // ignore errors during stat
  }
  return { totalBytes }
}

function formatBytes(bytes: number): string {
  if (bytes < 1024) {
    return `${bytes}B`
  }
  if (bytes < 1024 * 1024) {
    return `${(bytes / 1024).toFixed(0)}KB`
  }
  return `${(bytes / (1024 * 1024)).toFixed(0)}MB`
}
