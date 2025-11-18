import { defineEventHandler, readBody, createError } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    const { url } = body

    if (!url) {
      throw createError({
        statusCode: 400,
        message: 'آدرس URL الزامی است'
      })
    }

    // دریافت صفحه
    const response = await fetch(url, {
      headers: {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36'
      }
    })

    if (!response.ok) {
      throw createError({
        statusCode: response.status,
        message: `خطا در دریافت صفحه: ${response.statusText}`
      })
    }

    const html = await response.text()

    // استخراج متا تگ‌ها
    const extractMeta = (name: string): string => {
      const patterns = [
        new RegExp(`<meta\\s+name=["']${name}["']\\s+content=["']([^"']+)["']`, 'i'),
        new RegExp(`<meta\\s+content=["']([^"']+)["']\\s+name=["']${name}["']`, 'i'),
        new RegExp(`<meta\\s+property=["']${name}["']\\s+content=["']([^"']+)["']`, 'i'),
        new RegExp(`<meta\\s+content=["']([^"']+)["']\\s+property=["']${name}["']`, 'i')
      ]

      for (const pattern of patterns) {
        const match = html.match(pattern)
        if (match && match[1]) {
          return match[1]
        }
      }

      return ''
    }

    // استخراج تایتل
    const extractTitle = (): string => {
      const titleMatch = html.match(/<title[^>]*>([^<]+)<\/title>/i)
      return titleMatch ? titleMatch[1].trim() : ''
    }

    // استخراج تصویر اصلی
    const extractImage = (): string => {
      const ogImage = extractMeta('og:image')
      if (ogImage) return ogImage

      const twitterImage = extractMeta('twitter:image')
      if (twitterImage) return twitterImage

      // جستجوی اولین تصویر در صفحه
      const imgMatch = html.match(/<img[^>]+src=["']([^"']+)["']/i)
      return imgMatch ? imgMatch[1] : ''
    }

    const metaData = {
      title: extractTitle(),
      description: extractMeta('description') || extractMeta('og:description'),
      keywords: extractMeta('keywords'),
      image: extractImage(),
      og_title: extractMeta('og:title'),
      og_description: extractMeta('og:description'),
      og_image: extractMeta('og:image'),
      twitter_title: extractMeta('twitter:title'),
      twitter_description: extractMeta('twitter:description'),
      twitter_image: extractMeta('twitter:image'),
      canonical: extractMeta('canonical')
    }

    return {
      success: true,
      message: 'استخراج متادیتا با موفقیت انجام شد',
      data: metaData
    }
  } catch (error: any) {
    console.error('خطا در استخراج متادیتا:', error)
    
    if (error.statusCode) {
      throw error
    }

    throw createError({
      statusCode: 500,
      message: error.message || 'خطا در استخراج متادیتا'
    })
  }
})
