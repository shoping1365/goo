export type ProductLike = {
  id?: string | number
  sku?: string | number
  slug?: string | null
  seo_title?: string | null
}

/**
 * ساخت مسیر استاندارد نمایش محصول بر اساس الگوی پوشه `pages/product/[sku]/[[seo]].vue`.
 * فقط در یک نقطه «sku-» را نگه می‌داریم تا از تکرار در کل پروژه جلوگیری شود.
 */
export function useProductLink() {
  function buildProductLink(product: ProductLike | (string | number), seo?: string | null, includeSlug?: boolean): string {
    const rawSku = typeof product === 'object' ? (product.sku ?? product.id) : product
    const sku = String(rawSku ?? '').trim()
    
    // تعیین اینکه آیا slug باید در URL باشد یا نه
    let shouldIncludeSlug = includeSlug
    if (shouldIncludeSlug === undefined) {
      // اگر includeSlug مشخص نشده، از تنظیمات عمومی استفاده کن
      if (import.meta.client) {
        try {
          const settings = JSON.parse(localStorage.getItem('app-settings') || '{}')
          shouldIncludeSlug = settings.includeSlugInURL !== false // پیش‌فرض true
        } catch {
          shouldIncludeSlug = true // پیش‌فرض true
        }
      } else {
        shouldIncludeSlug = true // پیش‌فرض true
      }
    }
    
    const seoPart = typeof product === 'object'
      ? (product.slug || product.seo_title || null)
      : (seo ?? null)
    
    // دقیقاً همان منطق ریدایرکت‌ها در WordPress migration
    if (shouldIncludeSlug && seoPart) {
      return `/product/sku-${sku}/${seoPart}`
    } else {
      return `/product/sku-${sku}`
    }
  }
  return { buildProductLink }
}


