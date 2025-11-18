/**
 * ابزارهای مربوط به مدیریت تصاویر
 */

/**
 * تبدیل URL تصویر به سایز مشخص
 * @param url - URL اصلی تصویر
 * @param size - سایز مورد نظر (thumbnail, small, medium, large)
 * @returns URL تصویر با سایز جدید
 */
export function getImageSize(url: string, size: 'thumbnail' | 'small' | 'medium' | 'large'): string {
  if (!url) return '/default-product.jpg'
  
  // اگر URL از قبل شامل سایز است، آن را حذف کن
  const urlWithoutSize = url.replace(/_(thumbnail|small|medium|large)(\.[^.]+)?$/, '')
  
  // اگر URL شامل پسوند فایل است، سایز جدید را اضافه کن
  const dotIndex = urlWithoutSize.lastIndexOf('.')
  if (dotIndex !== -1) {
    const baseName = urlWithoutSize.substring(0, dotIndex)
    const extension = urlWithoutSize.substring(dotIndex)
    return `${baseName}_${size}${extension}`
  }
  
  // اگر پسوند فایل ندارد، سایز جدید را اضافه کن
  return `${urlWithoutSize}_${size}`
}

/**
 * تبدیل تصویر به سایز بزرگ (large)
 * @param url - URL اصلی تصویر
 * @returns URL تصویر با سایز large
 */
export function toLargeImage(url: string): string {
  return getImageSize(url, 'large')
}

/**
 * تبدیل تصویر به سایز متوسط (medium)
 * @param url - URL اصلی تصویر
 * @returns URL تصویر با سایز medium
 */
export function toMediumImage(url: string): string {
  return getImageSize(url, 'medium')
}

/**
 * تبدیل تصویر به سایز کوچک (small)
 * @param url - URL اصلی تصویر
 * @returns URL تصویر با سایز small
 */
export function toSmallImage(url: string): string {
  return getImageSize(url, 'small')
}

/**
 * تبدیل تصویر به سایز thumbnail
 * @param url - URL اصلی تصویر
 * @returns URL تصویر با سایز thumbnail
 */
export function toThumbnailImage(url: string): string {
  return getImageSize(url, 'thumbnail')
}
