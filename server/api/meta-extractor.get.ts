import { defineEventHandler, getQuery, createError } from 'h3'

/**
 * API endpoint برای استخراج متا تگ‌ها از URL های مختلف
 */
export default defineEventHandler(async (event) => {
  const query = getQuery(event)
  const url = query.url as string

  if (!url) {
    throw createError({
      statusCode: 400,
      message: 'URL الزامی است'
    })
  }

  try {
    // بررسی معتبر بودن URL
    new URL(url)
  } catch {
    throw createError({
      statusCode: 400,
      message: 'URL نامعتبر است'
    })
  }

  try {
    // دریافت محتوای صفحه
    const response = await fetch(url, {
      headers: {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36'
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
    const metaData = extractMetaTags(html)
    
    return {
      success: true,
      url: url,
      ...metaData
    }

  } catch (error) {
    console.error('خطا در استخراج متا:', error)
    
    throw createError({
      statusCode: 500,
      message: error instanceof Error ? error.message : 'خطا در استخراج متا تگ‌ها'
    })
  }
})

/**
 * استخراج متا تگ‌ها از HTML
 */
function extractMetaTags(html: string) {
  // استخراج title
  const titleMatch = html.match(/<title[^>]*>([^<]*)<\/title>/i)
  const title = titleMatch ? titleMatch[1].trim() : ''

  // استخراج meta description
  const descMatch = html.match(/<meta[^>]*name=["']description["'][^>]*content=["']([^"']*)["'][^>]*>/i)
  const description = descMatch ? descMatch[1].trim() : ''

  // استخراج meta keywords
  const keywordsMatch = html.match(/<meta[^>]*name=["']keywords["'][^>]*content=["']([^"']*)["'][^>]*>/i)
  const keywords = keywordsMatch ? keywordsMatch[1].trim() : ''

  // استخراج Open Graph title
  const ogTitleMatch = html.match(/<meta[^>]*property=["']og:title["'][^>]*content=["']([^"']*)["'][^>]*>/i)
  const ogTitle = ogTitleMatch ? ogTitleMatch[1].trim() : ''

  // استخراج Open Graph description
  const ogDescMatch = html.match(/<meta[^>]*property=["']og:description["'][^>]*content=["']([^"']*)["'][^>]*>/i)
  const ogDescription = ogDescMatch ? ogDescMatch[1].trim() : ''

  // استخراج Open Graph image
  const ogImageMatch = html.match(/<meta[^>]*property=["']og:image["'][^>]*content=["']([^"']*)["'][^>]*>/i)
  const ogImage = ogImageMatch ? ogImageMatch[1].trim() : ''

  // استخراج canonical URL
  const canonicalMatch = html.match(/<link[^>]*rel=["']canonical["'][^>]*href=["']([^"']*)["'][^>]*>/i)
  const canonical = canonicalMatch ? canonicalMatch[1].trim() : ''

  // استخراج robots meta
  const robotsMatch = html.match(/<meta[^>]*name=["']robots["'][^>]*content=["']([^"']*)["'][^>]*>/i)
  const robots = robotsMatch ? robotsMatch[1].trim() : ''

  return {
    title: cleanText(title),
    description: cleanText(description),
    keywords: cleanText(keywords),
    ogTitle: cleanText(ogTitle),
    ogDescription: cleanText(ogDescription),
    ogImage: cleanText(ogImage),
    canonical: cleanText(canonical),
    robots: cleanText(robots)
  }
}

/**
 * پاکسازی متن
 */
function cleanText(text: string): string {
  if (!text) return ''
  
  return text
    .replace(/\s+/g, ' ') // حذف فاصله‌های اضافی
    .replace(/&amp;/g, '&')
    .replace(/&lt;/g, '<')
    .replace(/&gt;/g, '>')
    .replace(/&quot;/g, '"')
    .replace(/&#39;/g, "'")
    .trim()
}



































