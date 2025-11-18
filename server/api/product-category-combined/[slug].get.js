import { defineEventHandler, getQuery, createError } from 'h3'

export default defineEventHandler(async (event) => {
  const { slug } = event.context.params // slug will be "parent-slug/child-slug"
  const q = getQuery(event)
  const isPreview = q.preview === '1' || q.preview === 'true'
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const previewParam = isPreview ? '?preview=1' : ''

  console.log('🔍 Fetching combined category slug:', slug)

  try {
    // تلاش مستقیم برای دریافت از بک‌اند با پیش‌نمایش
    const bySlug = await $fetch(`${base}/api/product-categories/slug/${slug}${previewParam}`)
    console.log('✅ Found category by combined slug:', bySlug)
    return bySlug
  } catch (e) {
    console.log('⚠️ Direct slug fetch failed, trying fallback method')
    // بازگشت به لیست کامل در صورت نبود مسیر بالا
    const categories = await $fetch(`${base}/api/product-categories?all=1`)
    const list = Array.isArray(categories) ? categories : (categories.data || [])

    // جستجو برای دسته‌بندی با slug ترکیبی
    const cat = list.find((c) => {
      // اگر دسته‌بندی والد دارد، slug کامل را بساز
      if (c.parent_slug && c.parent_slug !== '') {
        const fullSlug = `${c.parent_slug}/${c.slug}`
        return fullSlug === slug
      }
      // اگر دسته‌بندی اصلی است، فقط slug خودش را بررسی کن
      return c.slug === slug
    })

    if (!cat) {
      console.log('❌ Category not found with slug:', slug)
      throw createError({ statusCode: 404, statusMessage: 'Category not found' })
    }

    console.log('✅ Found category by fallback method:', cat)
    return cat
  }
})