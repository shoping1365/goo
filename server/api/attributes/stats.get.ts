import { defineEventHandler } from 'h3'
import { getGoApiBaseUrl } from '~/server/utils/api-config'

export default defineEventHandler(async () => {
  const base = getGoApiBaseUrl()
  try {
    const res = await fetch(`${base}/api/attributes/stats`)
    const json = await res.json()
    if (process.env.NODE_ENV === 'development') {
      console.debug('[dev] /api/attributes/stats →', json)
    }
    return json
  } catch (error) {
    console.error('Error fetching attributes stats', error)
    throw error
  }
});