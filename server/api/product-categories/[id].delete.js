export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const { id } = event.context.params

  try {
    const res = await $fetch(`${base}/api/product-categories/${id}`, {
      method: 'DELETE',
    })
    return res
  } catch (error) {
    console.error('Error deleting product category:', error)
    throw error
  }
})