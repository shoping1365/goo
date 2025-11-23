import { promises as fs } from 'fs'
import { defineEventHandler, readBody } from 'h3'
import { join } from 'path'

export default defineEventHandler(async (event) => {
  const file = join(process.cwd(), 'server', 'data', 'compression-results.json')
  await fs.mkdir(join(process.cwd(), 'server', 'data'), { recursive: true })
  let existing: unknown[] = []
  try {
    const prev = await fs.readFile(file, 'utf-8')
    existing = JSON.parse(prev || '[]')
  } catch {}

  const body = await readBody(event)
  if (!body) return { success: false }

  const incoming: unknown[] = Array.isArray(body) ? body : [body]
  // Merge by id (replace existing)
  const merged = [...incoming]
  for (const item of existing) {
    if (!incoming.find((i) => i.id === item.id)) {
      merged.push(item)
    }
  }
  await fs.writeFile(file, JSON.stringify(merged, null, 2), 'utf-8')
  return { success: true }
}) 