import { defineEventHandler, readBody } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const payload = await readBody(event)

  try {
    const res = await $fetch(`${base}/api/product-categories/reorder`, {
      method: 'POST',
      body: payload,
    })
    return res
  } catch (error) {
    console.error('Error reordering product categories:', error)
    throw error
  }
})
