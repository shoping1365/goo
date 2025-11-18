import { defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
  const base = getGoApiBaseUrl()
  // @ts-ignore - params defined by Nitro runtime
  const id = event.context.params?.id
  try {
    const res = await fetch(`${base}/api/attributes/${id}`)
    const json = await res.json()
    if (!res.ok) {
      throw json
    }
    return json
  } catch (err) {
    console.error('Error fetching attribute detail', err)
    throw err
  }
})