import { readBody } from 'h3'

export default defineEventHandler(async (event) => {
  const base = getGoApiBaseUrl()
  const { proxy } = await import('../_utils/fetchProxy')
  const body = await readBody(event)
  return proxy(event, `${base}/api/brands`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body)
  })
})