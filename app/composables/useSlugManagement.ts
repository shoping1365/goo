interface SlugCheckResponse {
  exists: boolean
  slug: string
}

interface SlugGenerateResponse {
  slug: string
}

// Composable برای مدیریت URL (Slug)
export const useSlugManagement = () => {
  // تابع تبدیل عنوان به slug
  const slugify = (text: string): string => {
    return text
      .toString()
      .trim()
      .replace(/\s+/g, '-') // فاصله‌ها به -
      .replace(/[\u200C\u200B]/g, '') // حذف نیم‌فاصله و کاراکترهای نامرئی
      .replace(/[^\w\u0600-\u06FF-]+/g, '') // فقط حروف، اعداد و فارسی
      .replace(/-+/g, '-')
      .replace(/^-+|-+$/g, '')
      .toLowerCase()
  }

  // تابع بررسی تکراری بودن slug
  const checkSlugUnique = async (slug: string): Promise<{ exists: boolean; slug: string }> => {
    if (!slug.trim()) {
      throw new Error('URL الزامی است')
    }
    
    try {
      const response = await $fetch<SlugCheckResponse>(`/api/products/check-slug?slug=${encodeURIComponent(slug)}`)
      return response
    } catch (error) {
      console.warn('خطا در بررسی URL:', error)
      throw new Error('خطا در بررسی URL')
    }
  }

  // تابع تولید slug یکتا
  const generateUniqueSlug = async (baseSlug: string): Promise<string> => {
    try {
      const response = await $fetch<SlugGenerateResponse>(`/api/products/generate-slug?base_slug=${encodeURIComponent(baseSlug)}`)
      return response.slug
    } catch (error) {
      console.warn('خطا در تولید URL یکتا:', error)
      throw new Error('خطا در تولید URL یکتا')
    }
  }

  // تابع تولید خودکار slug از عنوان
  const generateSlugFromTitle = async (title: string): Promise<string> => {
    if (!title.trim()) {
      throw new Error('ابتدا عنوان را وارد کنید')
    }
    
    const baseSlug = slugify(title)
    if (!baseSlug) {
      throw new Error('نمی‌توان از عنوان URL تولید کرد')
    }
    
    return await generateUniqueSlug(baseSlug)
  }

  return {
    slugify,
    checkSlugUnique,
    generateUniqueSlug,
    generateSlugFromTitle
  }
} 