export default defineEventHandler(async (event) => {
  const { readBody } = await import('h3')
  const { proxy } = await import('../_utils/fetchProxy')
  
  const body = await readBody(event)
  const base = getGoApiBaseUrl()
  
  return proxy(event, `${base}/api/dev/restore-soft-delete`, {
    method: 'POST',
    body,
  })
})