import { defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
  const base = getGoApiBaseUrl()
  // @ts-ignore
  const attrId = event.context.params?.attrId
  try {
    const res = await fetch(`${base}/api/attribute-values/by-attribute/${attrId}`)
    const json = await res.json()
    if (!res.ok) {
      throw json
    }
    return json
  } catch (err) {
    console.error('Error fetching attribute values', err)
    throw err
  }
})