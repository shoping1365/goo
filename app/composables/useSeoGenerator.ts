// Composable برای تولید پویای فیلدهای SEO
export interface SeoData {
  title: string
  excerpt: string
  content: string
  slug: string
  meta_title?: string
  meta_description?: string
  meta_keywords?: string
  featured_image?: string
  robots_meta?: string
  author_name?: string
  category_name?: string
  tags?: string[]
  language?: string
  site_url?: string
  
  // فیلدهای اضافی برای اسکیمای خاص
  location?: string
  event_date?: string
  estimated_time?: number
  product_price?: number
  rating?: number
}

// محاسبه تعداد کلمات
function calculateWordCount(content: string): number {
  // حذف HTML tags
  const cleanContent = content.replace(/<[^>]*>/g, '');
  // تقسیم بر فاصله و شمارش
  const words = cleanContent.trim().split(/\s+/);
  return words.length;
}

// محاسبه زمان مطالعه (به دقیقه)
function calculateReadingTime(wordCount: number): number {
  // میانگین سرعت مطالعه: 200 کلمه در دقیقه
  const readingTime = Math.ceil(wordCount / 200);
  return readingTime < 1 ? 1 : readingTime;
}

export interface SeoTags {
  html_meta_tags: string
  open_graph_tags: string
  twitter_card_tags: string
  json_ld_schema: string
  og_title: string
  og_description: string
  og_site_name: string
}

export interface SeoPreview {
  seo_tags: SeoTags
  preview: {
    og_title: string
    og_description: string
    og_site_name: string
    meta_title: string
    meta_description: string
    canonical_url: string
  }
}

// تعریف interface برای response های API
interface ApiResponse<T = unknown> {
  data?: T
  success?: boolean
  message?: string
}

export const useSeoGenerator = () => {
  // تولید تگ‌های SEO برای یک نوشته موجود
  const generateSeoTagsForPost = async (postId: number): Promise<SeoTags | null> => {
    try {
      const res = await $fetch<ApiResponse<SeoTags>>(`/api/seo-generator/post/${postId}`)
      return res.data || null
    } catch (error) {
      console.warn('خطا در تولید تگ‌های SEO:', error)
      return null
    }
  }

  // تولید تگ‌های SEO از داده‌های ارسالی
  const generateSeoTagsFromData = async (seoData: SeoData): Promise<SeoTags | null> => {
    try {
      const res = await $fetch<ApiResponse<SeoTags>>('/api/seo-generator/generate', {
        method: 'POST',
        body: seoData
      })
      return res.data || null
    } catch (error) {
      console.warn('خطا در تولید تگ‌های SEO:', error)
      return null
    }
  }

  // پیش‌نمایش تگ‌های SEO
  const previewSeoTags = async (seoData: SeoData): Promise<SeoPreview | null> => {
    try {
      // تبدیل تگ‌ها از آرایه به string
      const previewData = {
        ...seoData,
        tags: seoData.tags ? seoData.tags.join(', ') : ''
      }

      const res = await $fetch<ApiResponse<SeoPreview>>('/api/seo-generator/preview', {
        method: 'POST',
        body: previewData
      })
      return res.data || null
    } catch (error) {
      console.warn('خطا در پیش‌نمایش تگ‌های SEO:', error)
      return null
    }
  }

  // تولید تگ‌های HTML Meta به صورت محلی
  const generateHTMLMetaTags = (seoData: SeoData): string => {
    const tags: string[] = []

    // Meta Title
    if (seoData.meta_title) {
      tags.push(`<meta name="title" content="${escapeHtml(seoData.meta_title)}">`)
    }

    // Meta Description
    if (seoData.meta_description) {
      tags.push(`<meta name="description" content="${escapeHtml(seoData.meta_description)}">`)
    }

    // Meta Keywords
    if (seoData.meta_keywords) {
      tags.push(`<meta name="keywords" content="${escapeHtml(seoData.meta_keywords)}">`)
    }

    // Robots
    const robotsMeta = seoData.robots_meta || 'index,follow'
    tags.push(`<meta name="robots" content="${robotsMeta}">`)

    // Canonical URL
    if (seoData.slug && seoData.site_url) {
      const canonicalURL = `${seoData.site_url}/blog/${seoData.slug}`
      tags.push(`<link rel="canonical" href="${canonicalURL}">`)
    }

    // Language
    if (seoData.language) {
      tags.push(`<meta http-equiv="content-language" content="${seoData.language}">`)
    }

    return tags.join('\n')
  }

  // تولید تگ‌های Open Graph به صورت محلی
  const generateOpenGraphTags = (seoData: SeoData): string => {
    const tags: string[] = []

    // OG Title
    const ogTitle = seoData.meta_title || seoData.title
    if (ogTitle) {
      tags.push(`<meta property="og:title" content="${escapeHtml(ogTitle)}">`)
    }

    // OG Description
    const ogDescription = seoData.meta_description || seoData.excerpt || seoData.content
    if (ogDescription) {
      const cleanDescription = stripHTML(ogDescription)
      const truncatedDescription = cleanDescription.length > 160 
        ? cleanDescription.substring(0, 157) + '...' 
        : cleanDescription
      tags.push(`<meta property="og:description" content="${escapeHtml(truncatedDescription)}">`)
    }

    // OG Type - همیشه article برای مقالات
    tags.push(`<meta property="og:type" content="article">`)

    // OG Image - از تصویر شاخص استفاده می‌کنیم
    const ogImage = seoData.featured_image || `${seoData.site_url}/default-og-image.jpg`
    tags.push(`<meta property="og:image" content="${ogImage}">`)

    // OG Site Name
    const ogSiteName = extractDomain(seoData.site_url) || 'وب‌سایت من'
    tags.push(`<meta property="og:site_name" content="${escapeHtml(ogSiteName)}">`)

    // OG URL
    if (seoData.slug && seoData.site_url) {
      const ogURL = `${seoData.site_url}/blog/${seoData.slug}`
      tags.push(`<meta property="og:url" content="${ogURL}">`)
    }

    // OG Locale
    if (seoData.language === 'fa') {
      tags.push('<meta property="og:locale" content="fa_IR">')
    }

    return tags.join('\n')
  }

  // تولید تگ‌های Twitter Card به صورت محلی
  const generateTwitterCardTags = (seoData: SeoData): string => {
    const tags: string[] = []

    // Twitter Card Type
    tags.push('<meta name="twitter:card" content="summary_large_image">')

    // Twitter Title
    const ogTitle = seoData.meta_title || seoData.title
    if (ogTitle) {
      tags.push(`<meta name="twitter:title" content="${escapeHtml(ogTitle)}">`)
    }

    // Twitter Description
    const ogDescription = seoData.meta_description || seoData.excerpt || seoData.content
    if (ogDescription) {
      const cleanDescription = stripHTML(ogDescription)
      const truncatedDescription = cleanDescription.length > 160 
        ? cleanDescription.substring(0, 157) + '...' 
        : cleanDescription
      tags.push(`<meta name="twitter:description" content="${escapeHtml(truncatedDescription)}">`)
    }

    // Twitter Image - از تصویر شاخص استفاده می‌کنیم
    const ogImage = seoData.featured_image || `${seoData.site_url}/default-og-image.jpg`
    tags.push(`<meta name="twitter:image" content="${ogImage}">`)

    return tags.join('\n')
  }

  // تولید JSON-LD Schema به صورت کاملاً پویا
  const generateJSONLDSchema = (seoData: SeoData): string => {
    // تعیین نوع اسکیما بر اساس محتوا
    const schemaType = determineSchemaType(seoData)
    
    const schema: Record<string, unknown> = {
      '@context': 'https://schema.org',
      '@type': schemaType
    }

    // عنوان
    if (seoData.title) {
      schema.headline = seoData.title
    }

    // توضیحات
    if (seoData.meta_description) {
      schema.description = seoData.meta_description
    }

    // نویسنده
    if (seoData.author_name) {
      schema.author = {
        '@type': 'Person',
        name: seoData.author_name
      }
    }

    // ناشر
    const ogSiteName = extractDomain(seoData.site_url) || 'وب‌سایت من'
    schema.publisher = {
      '@type': 'Organization',
      name: ogSiteName
    }

    // URL اصلی
    if (seoData.slug && seoData.site_url) {
      const mainEntityURL = `${seoData.site_url}/blog/${seoData.slug}`
      schema.mainEntityOfPage = {
        '@type': 'WebPage',
        '@id': mainEntityURL
      }
    }

    // تصویر - از تصویر شاخص استفاده می‌کنیم
    const ogImage = seoData.featured_image || `${seoData.site_url}/default-og-image.jpg`
    schema.image = ogImage

    // دسته‌بندی
    if (seoData.category_name) {
      schema.articleSection = seoData.category_name
    }

    // کلمات کلیدی
    if (seoData.meta_keywords) {
      schema.keywords = seoData.meta_keywords
    }

    // تعداد کلمات - محاسبه پویا
    const wordCount = calculateWordCount(seoData.content)
    if (wordCount > 0) {
      schema.wordCount = wordCount
    }

    // زمان مطالعه - محاسبه پویا
    const readingTime = calculateReadingTime(wordCount)
    if (readingTime > 0) {
      schema.timeRequired = `PT${readingTime}M`
    }

    // زبان
    if (seoData.language) {
      schema.inLanguage = seoData.language
    }

    // تگ‌ها
    if (seoData.tags && seoData.tags.length > 0) {
      schema.keywords = seoData.tags.join(', ')
    }

    // فیلدهای خاص بر اساس نوع اسکیما
    addSchemaSpecificFields(schema, seoData, schemaType)

    return JSON.stringify(schema, null, 2)
  }

  // تعیین نوع اسکیما بر اساس محتوا
  const determineSchemaType = (seoData: SeoData): string => {
    // بر اساس محتوا و فیلدهای موجود، نوع اسکیما را تعیین کن
    if (seoData.category_name && seoData.category_name.toLowerCase().includes('خبر')) {
      return 'NewsArticle'
    }
    
    if (seoData.category_name && seoData.category_name.toLowerCase().includes('آموزش')) {
      return 'TechArticle'
    }
    
    if (seoData.category_name && seoData.category_name.toLowerCase().includes('بررسی')) {
      return 'Review'
    }
    
    return 'Article'
  }

  // اضافه کردن فیلدهای خاص بر اساس نوع اسکیما
  const addSchemaSpecificFields = (schema: Record<string, unknown>, seoData: SeoData, schemaType: string) => {
    switch (schemaType) {
      case 'NewsArticle':
        // فیلدهای خاص مقالات خبری
        if (seoData.location) {
          schema.dateline = seoData.location
        }
        if (seoData.event_date) {
          schema.dateCreated = seoData.event_date
        }
        break
        
      case 'TechArticle':
        // فیلدهای خاص مقالات آموزشی
        if (seoData.estimated_time && seoData.estimated_time > 0) {
          schema.timeRequired = `PT${seoData.estimated_time}M`
        }
        break
        
      case 'Review':
        // فیلدهای خاص بررسی محصول
        if (seoData.product_price && seoData.product_price > 0) {
          schema.offers = {
            '@type': 'Offer',
            price: seoData.product_price,
            priceCurrency: 'IRR'
          }
        }
        if (seoData.rating && seoData.rating > 0) {
          schema.reviewRating = {
            '@type': 'Rating',
            ratingValue: seoData.rating,
            bestRating: 5
          }
        }
        break
    }
  }

  // تولید تمام تگ‌های SEO به صورت محلی
  const generateAllSeoTags = (seoData: SeoData): SeoTags => {
    const ogTitle = seoData.meta_title || seoData.title
    const ogDescription = seoData.meta_description || seoData.excerpt || seoData.content
    const cleanDescription = stripHTML(ogDescription || '')
    const truncatedDescription = cleanDescription.length > 160 
      ? cleanDescription.substring(0, 157) + '...' 
      : cleanDescription
    const ogSiteName = extractDomain(seoData.site_url) || 'وب‌سایت من'

    return {
      html_meta_tags: generateHTMLMetaTags(seoData),
      open_graph_tags: generateOpenGraphTags(seoData),
      twitter_card_tags: generateTwitterCardTags(seoData),
      json_ld_schema: generateJSONLDSchema(seoData),
      og_title: ogTitle || '',
      og_description: truncatedDescription,
      og_site_name: ogSiteName
    }
  }

  return {
    generateSeoTagsForPost,
    generateSeoTagsFromData,
    previewSeoTags,
    generateHTMLMetaTags,
    generateOpenGraphTags,
    generateTwitterCardTags,
    generateJSONLDSchema,
    generateAllSeoTags
  }
}

// Helper functions
function escapeHtml(text: string): string {
  const div = document.createElement('div')
  div.textContent = text
  return div.innerHTML
}

function stripHTML(html: string): string {
  if (typeof document === 'undefined') {
    // Server-side fallback
    return html
      .replace(/<br\s*\/?>/gi, ' ')
      .replace(/<p[^>]*>/gi, ' ')
      .replace(/<\/p>/gi, ' ')
      .replace(/<div[^>]*>/gi, ' ')
      .replace(/<\/div>/gi, ' ')
      .replace(/<span[^>]*>/gi, '')
      .replace(/<\/span>/gi, '')
      .replace(/<[^>]*>/g, '')
      .replace(/\s+/g, ' ')
      .trim()
  }

  const div = document.createElement('div')
  div.innerHTML = html
  return div.textContent || div.innerText || ''
}

function extractDomain(url?: string): string {
  if (!url) return ''
  
  try {
    const urlObj = new URL(url)
    return urlObj.hostname
  } catch {
    // اگر URL نامعتبر باشد، سعی کن از string استخراج کن
    const match = url.match(/https?:\/\/([^\/]+)/)
    return match ? match[1] : ''
  }
} 