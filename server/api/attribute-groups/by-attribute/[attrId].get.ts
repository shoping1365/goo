import { defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const attrId = event.context.params?.attrId
  try {
    const res = await fetch(`${base}/api/attribute-groups/by-attribute/${attrId}`)
    const json = await res.json()
    if (!res.ok) throw json
    return json
  } catch (err) {
    console.error('Error fetching attribute groups', err)
    throw err
  }
})
