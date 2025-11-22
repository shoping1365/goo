import { defineEventHandler, getQuery } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const { slug } = event.context.params
  const q = getQuery(event)
  const isPreview = q.preview === '1' || q.preview === 'true'
  const previewParam = isPreview ? '?preview=1' : ''
  try {
    // تلاش مستقیم برای دریافت از بک‌اند با پیش‌نمایش
    const bySlug = await $fetch(`${base}/api/product-categories/slug/${slug}${previewParam}`)
    return bySlug
  } catch (_e) {
    // بازگشت به لیست کامل در صورت نبود مسیر بالا
    const categories = await $fetch(`${base}/api/product-categories?all=1`)
    const list = Array.isArray(categories) ? categories : (categories.data || [])
    const cat = list.find((c) => c.slug === slug)
    if (!cat) throw createError({ statusCode: 404, statusMessage: 'Category not found' })
    return cat
  }
})