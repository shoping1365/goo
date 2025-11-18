export default defineEventHandler(async (event) => {
  const { readBody } = await import('h3')
  const { proxy } = await import('../_utils/fetchProxy')
  
  const base = getGoApiBaseUrl()
  try {
    const body = await readBody(event)
    return proxy(event, `${base}/api/dev/purge-soft-deletes`, {
      method: 'POST',
      body,
    })
  } catch (error) {
    console.error('Purging soft deletes error:', error)
    throw error
  }
})