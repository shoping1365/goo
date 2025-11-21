import { defineStore } from 'pinia'
import { computed, reactive, ref, watch } from 'vue'
// حذف وابستگی به dayjs؛ از Date داخلی استفاده می‌کنیم
import { useNotifier } from '~/composables/useNotifier'
import { resolveErrorMessage } from '~/utils/errorMessages'

export const useProductCreateStore = defineStore('productCreate', () => {
  const notifier = useNotifier()
  // --- Saving State
  const isSaving = ref(false)
  const isEditMode = ref(false)
  const editingProductId = ref<string | number | null>(null)
  const isLoadingProduct = ref(false)
  const isCategoriesLoading = ref(false)
  // کنترل فعال‌سازی ذخیره خودکار پس از تثبیت شناسه صحیح محصول
  const autoSaveReady = ref(false)
  const pricingLoaded = ref(false)

  // --- Main Product Form ---
  const productForm = reactive({
    name: '',
    englishName: '',
    description: '',
    fullDescription: '',
    status: 'active',
    sku: '',
    slug: '',
    seo_title: '',
    meta_description: '',
    url: '',
    brand_id: '',
    category_id: '',
    sub_category_id: ''
  })

  // --- Pricing Form ---
  const pricingForm = reactive({
    price: 0,
    old_price: 0,
    cost: 0,
    discount_percent: 0,
    discount_amount: 0,
    profit: 0,
    disableBuyButton: false,
    callForPrice: false,
    // فروش ویژه تک‌قیمتی (برای سازگاری قدیمی؛ برای طرح‌های چندمرحله‌ای از specialOffers استفاده می‌شود)
    sale_price: 0 as number,
    // --- فیلدهای زمان‌بندی قیمت ویژه ---
    sale_start_at: null as string | null,
    sale_end_at: null as string | null,
    // رشته‌های نمایشی شمسی برای ورودی‌های UI
    sale_start_jalali: '' as string,
    sale_end_jalali: '' as string,
  })

  // --- طرح‌های فروش ویژه چندمرحله‌ای (قیمت/تعداد) ---
  type SpecialOffer = { base_price: number; price: number; quantity: number }
  const specialOffers = ref<SpecialOffer[]>([])

  // --- Inventory Form ---
  const inventoryForm = reactive({
    stock_quantity: 0,
    min_stock_quantity: 0,
    max_stock_quantity: 0,
    stock_status: 'in_stock' as string | null,
    track_inventory: true,
    show_stock_to_customer: false,
    allow_reservation: false,
    warehouse_id: null as number | null,
    shipping_enabled: true
  })

  // Computed profit calculations
  const computedProfit = computed(() => {
    return pricingForm.price - pricingForm.cost
  })

  const computedProfitPercent = computed(() => {
    if (pricingForm.cost === 0) return 0
    return ((pricingForm.price - pricingForm.cost) / pricingForm.cost) * 100
  })

  const computedDiscountAmount = computed(() => {
    return pricingForm.discount_amount || 0
  })

  const computedDiscountPercent = computed(() => {
    return pricingForm.discount_percent || 0
  })

  // Watch for automatic profit calculation only
  watch(() => [pricingForm.price, pricingForm.cost], ([newPrice, newCost]) => {
    pricingForm.profit = newPrice - newCost
  })

  // ---------------- تبدیل تاریخ: ISO ↔ جلالی ----------------
  /**
   * تبدیل رشتهٔ جلالی به ISO8601 (UTC)
   * ورودی قابل قبول: yyyy/MM/dd یا yyyy/MM/dd HH:mm
   */
  // مبدل ارقام فارسی/عربی به لاتین
  function toEnglishDigits(input: string): string {
    const map: Record<string, string> = {
      '۰': '0', '۱': '1', '۲': '2', '۳': '3', '۴': '4', '۵': '5', '۶': '6', '۷': '7', '۸': '8', '۹': '9',
      '٠': '0', '١': '1', '٢': '2', '٣': '3', '٤': '4', '٥': '5', '٦': '6', '٧': '7', '٨': '8', '٩': '9'
    }
    return input.replace(/[۰-۹٠-٩]/g, ch => map[ch] || ch)
  }

  // تبدیل تاریخ شمسی به میلادی (الگوریتم استاندارد)
  function jalaliToGregorian(jy: number, jm: number, jd: number): { gy: number; gm: number; gd: number } {
    let gy: number
    if (jy > 979) {
      gy = 1600
      jy -= 979
    } else {
      gy = 621
    }
    let days = 365 * jy + Math.floor(jy / 33) * 8 + Math.floor(((jy % 33) + 3) / 4) + jd + (jm < 7 ? (jm - 1) * 31 : ((jm - 7) * 30) + 186)
    gy += 400 * Math.floor(days / 146097)
    days %= 146097
    if (days > 36524) {
      gy += 100 * Math.floor(--days / 36524)
      days %= 36524
      if (days >= 365) days++
    }
    gy += 4 * Math.floor(days / 1461)
    days %= 1461
    if (days > 365) {
      gy += Math.floor((days - 1) / 365)
      days = (days - 1) % 365
    }
    let gd = days + 1
    const sal_a = [0, 31, ((gy % 4 === 0 && gy % 100 !== 0) || (gy % 400 === 0)) ? 29 : 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
    let gm = 0
    for (gm = 1; gm <= 12; gm++) {
      const v = sal_a[gm]
      if (gd <= v) break
      gd -= v
    }
    return { gy, gm, gd }
  }

  // پارس رشتهٔ شمسی «yyyy/mm/dd hh:mm» → ISO8601
  function parseJalaliToISO(jalali: string): string | null {
    if (!jalali || typeof jalali !== 'string') return null
    const normalized = toEnglishDigits(jalali).replace(/-/g, '/').trim()
    const m = normalized.match(/^(\d{4})\/(\d{1,2})\/(\d{1,2})(?:\s+(\d{1,2}):(\d{2}))?$/)
    if (!m) {
      // اگر الگوی جلالی منطبق نشد، تلاش کن به عنوان تاریخ میلادی yyyy/MM/dd HH:mm تجزیه کنی
      const dt = new Date(normalized.replace(/\//g, '-') + (normalized.includes(':') ? '' : ' 00:00'))
      if (!isNaN(dt.getTime())) return dt.toISOString()
      return null
    }
    const jy = Number(m[1])
    const jm = Number(m[2])
    const jd = Number(m[3])
    const hh = m[4] ? Number(m[4]) : 0
    const mm = m[5] ? Number(m[5]) : 0
    const g = jalaliToGregorian(jy, jm, jd)
    const dt = new Date(g.gy, g.gm - 1, g.gd, hh, mm, 0, 0)
    if (isNaN(dt.getTime())) return null
    return dt.toISOString()
  }

  /**
   * تبدیل ISO8601 به رشتهٔ جلالی (yyyy/MM/dd HH:mm)
   */
  // فرمت تاریخ میلادی به نمایش شمسی «yyyy/mm/dd hh:mm» با ارقام فارسی
  function formatJalaliFromISO(iso: string | null | undefined): string {
    if (!iso) return ''
    const d = new Date(iso)
    if (isNaN(d.getTime())) return ''
    const dateFmt = new Intl.DateTimeFormat('fa-IR-u-ca-persian', { year: 'numeric', month: '2-digit', day: '2-digit' })
    const timeFmt = new Intl.DateTimeFormat('fa-IR-u-ca-persian', { hour: '2-digit', minute: '2-digit' })
    return `${dateFmt.format(d)} ${timeFmt.format(d)}`
  }

  // Auto-save pricing when values change (edit mode only)
  watch(() => [
    pricingForm.price,
    pricingForm.old_price,
    pricingForm.cost,
    pricingForm.discount_percent,
    pricingForm.discount_amount,
    pricingForm.disableBuyButton,
    pricingForm.callForPrice,
    pricingForm.sale_price,
    pricingForm.sale_start_at,
    pricingForm.sale_end_at,
  ], async () => {
    if (isEditMode.value && editingProductId.value && pricingLoaded.value) {
      try {
        await savePricingData(editingProductId.value)
      } catch {
        // خطا در ذخیره خودکار قیمت
      }
    }
  })

  // همگام‌سازی خودکار: تغییر رشتهٔ شمسی → به‌روزرسانی ISO
  watch(() => [pricingForm.sale_start_jalali, pricingForm.sale_end_jalali], ([sj, ej]) => {
    pricingForm.sale_start_at = parseJalaliToISO(sj)
    pricingForm.sale_end_at = parseJalaliToISO(ej)
  })

  // TinyMCE API Key (keep the same key used previously)
  const tinyApiKey = 'qwa4j6x5mh2e3241igpyi345b4uhe2d5qeq6f8hy9qfkw2ro'

  // --- Collapsible Sections (visibility toggles) ---
  const sections = reactive({
    // Info tab
    mainInfo: true,
    technicalInfo: true,
    displaySettings: true,
    scheduling: true,
    management: true,
    // Images tab
    images: true,
    mainImages: true,
    galleryList: true,
    gallerySettings: true,
    // Pricing tab
    mainPrices: true,
    priceSettings: true,
    taxSettings: true,
    tierPricing: true,
    discountCodes: true,
    // تحلیل قیمت رقبا
    priceAnalysis: true,
    // Shipping tab
    dimensions: true,
    shippingSettings: true,
    shippingCosts: true,
    deliverySchedule: true,
    specialSettings: true,
    packaging: true,
    // Inventory tab
    inventoryMain: true,
    purchaseLimits: true,
    alerts: true,
    advancedSettings: true,
    multiWarehouse: true,
    // SEO tab
    seoBasic: true,
    seoSchema: true,
    seoSocial: true,
    seoIndexing: true,
    seoAdvanced: true,
    seoPreview: true,
    // Variants tab
    variantAttributes: true,
    variantCombinations: true,
    variantGeneration: true,
    variantBulkEdit: true,
    variantManagement: true,
    variantPricing: true,
    variantImages: true,
    // Misc
    strengthsWeaknesses: true,
    faq: true,
    videoForm: true,
    videoList: true,
    aparatVideos: true
  })

  // --- Images ---
  interface MediaImage { id: number; url: string; thumbnail: string; name?: string; size?: number;[key: string]: unknown }
  const images = ref<MediaImage[]>([])
  // --- Product Specifications (attribute values) ---
  interface SpecPayload { attribute_id: number; option_id?: number | null; option_ids?: number[]; value_text?: string | null }
  const productSpecs = ref<SpecPayload[]>([])
  function addImages(newImages: MediaImage[]) {
    const existingIds = new Set(images.value.map(i => i.id))
    newImages.forEach(img => {
      if (!existingIds.has(img.id)) images.value.push(img)
    })
  }
  function removeImage(id: number) {
    images.value = images.value.filter(i => i.id !== id)
  }

  /**
   * Toggles visibility of a section.
   */
  function toggleSection(key: keyof typeof sections) {
    sections[key] = !sections[key]
  }

  // ---------------------------
  // Reset functions
  // ---------------------------
  function resetForm() {
    productForm.name = ''
    productForm.englishName = ''
    productForm.description = ''
    productForm.fullDescription = ''
    productForm.status = 'active'
    productForm.sku = ''
    productForm.slug = ''
    productForm.seo_title = ''
    productForm.meta_description = ''
    productForm.url = ''
    productForm.brand_id = ''
    productForm.category_id = ''
    productForm.sub_category_id = ''

    // Reset pricing
    pricingForm.price = 0
    pricingForm.old_price = 0
    pricingForm.cost = 0
    pricingForm.discount_percent = 0
    pricingForm.discount_amount = 0
    pricingForm.profit = 0
    pricingForm.disableBuyButton = false
    pricingForm.callForPrice = false
    specialOffers.value = []

    // Reset inventory
    inventoryForm.stock_quantity = 0
    inventoryForm.min_stock_quantity = 0
    inventoryForm.max_stock_quantity = 0
    inventoryForm.stock_status = 'in_stock'
    inventoryForm.track_inventory = true
    inventoryForm.show_stock_to_customer = false

    images.value = []
    productSpecs.value = []
    isEditMode.value = false
    editingProductId.value = null
    autoSaveReady.value = false
  }

  // The new $reset function to be called from components
  function $reset() {
    // Reset saving state
    isSaving.value = false
    isEditMode.value = false
    editingProductId.value = null
    autoSaveReady.value = false

    // Reset the main form data
    productForm.name = ''
    productForm.englishName = ''
    productForm.description = ''
    productForm.fullDescription = ''
    productForm.status = 'active'
    productForm.sku = ''
    productForm.slug = ''
    productForm.seo_title = ''
    productForm.meta_description = ''
    productForm.url = ''
    productForm.brand_id = ''
    productForm.category_id = ''
    productForm.sub_category_id = ''

    // Reset pricing
    pricingForm.price = 0
    pricingForm.old_price = 0
    pricingForm.cost = 0
    pricingForm.discount_percent = 0
    pricingForm.discount_amount = 0
    pricingForm.profit = 0
    pricingForm.disableBuyButton = false
    pricingForm.callForPrice = false
    specialOffers.value = []

    // Reset inventory
    inventoryForm.stock_quantity = 0
    inventoryForm.min_stock_quantity = 0
    inventoryForm.max_stock_quantity = 0
    inventoryForm.stock_status = 'in_stock'
    inventoryForm.track_inventory = true
    inventoryForm.show_stock_to_customer = false

    // Reset images
    images.value = []
    productSpecs.value = []

    // Reset categories list
    categories.value = []

    // Reset all collapsible sections to their default state (true)
    for (const key in sections) {
      sections[key as keyof typeof sections] = true
    }
  }

  // ---------------------------
  // Load product for editing
  // ---------------------------
  async function loadProductForEdit(productId: string | number) {
    isLoadingProduct.value = true
    try {
      isSaving.value = true
      isEditMode.value = true
      // تا قبل از لود موفق محصول، شناسه ویرایش را خالی نگه می‌داریم تا کال‌های خودکار با ID اشتباه انجام نشوند
      editingProductId.value = null

      // استفاده از any برای جلوگیری از خطای TypeScript «excessively deep»
      const response = await $fetch<Record<string, unknown>>(`/api/admin/products/${productId}`)

      const product = response

      // اطمینان از ست شدن شناسه صحیح برای کال‌های بعدی (قیمت/موجودی)
      if (product && product.id) {
        editingProductId.value = Number(product.id)
        autoSaveReady.value = true
      }

      // Fill form with existing data
      productForm.name = (product.name as string) || ''
      productForm.englishName = (product.name_en as string) || ''
      productForm.description = (product.description as string) || ''
      productForm.fullDescription = (product.full_description as string) || ''
      productForm.status = (product.status as string) || 'active'
      productForm.sku = (product.sku as string) || ''
      productForm.slug = (product.slug as string) || ''
      productForm.seo_title = (product.seo_title as string) || (product.name as string) || ''
      productForm.meta_description = (product.meta_description as string) || ''
      productForm.url = (product.url as string) || ''
      productForm.brand_id = (product.brand_id as string | number) ? String(product.brand_id) : ''
      // اگر دستهبندی والد وجود داشته باشد، آن را بعنوان دسته اصلی و خود دسته محصول را بعنوان فرعی تنظیم میکنیم
      const category = product.category as { parent_id?: number; id?: number } | undefined
      if (category && category.parent_id) {
        productForm.category_id = String(category.parent_id)
        productForm.sub_category_id = (product.category_id as number) ? String(product.category_id) : (category.id ? String(category.id) : '')
      } else {
        productForm.category_id = (product.category_id as number) ? String(product.category_id) : ''
        productForm.sub_category_id = ''
      }

      // Load pricing data
      pricingForm.price = (product.price as number) || 0
      pricingForm.old_price = (product.old_price as number) || 0
      pricingForm.cost = (product.cost as number) || 0
      pricingForm.discount_percent = (product.discount_percent as number) || 0
      pricingForm.discount_amount = (product.discount_amount as number) || 0
      pricingForm.profit = (product.profit as number) || 0
      pricingForm.disableBuyButton = (product.disable_buy_button as boolean) || false
      pricingForm.callForPrice = (product.call_for_price as boolean) || false

      // Load images if exists (normalize image_url -> url)
      if (product.images && Array.isArray(product.images)) {
        images.value = product.images.map((img: Record<string, unknown>) => ({
          ...img,
          url: img.url || img.image_url,
          thumbnail: img.thumbnail || img.url || img.image_url
        })) as MediaImage[]
      }

      // ----- Load existing specifications -----
      try {
        const specsRes = await $fetch<Record<string, unknown>>(`/api/admin/products/${productId}/specs`)
        const arr = (specsRes?.data as Record<string, unknown>[]) || (specsRes as unknown as Record<string, unknown>[]) || []
        if (Array.isArray(arr)) {
          const map: Record<number, { attribute_id: number; option_ids: number[]; option_id?: number; value_text?: string }> = {}
          for (const s of arr) {
            const aid = Number(s.attribute_id)
            const optId = s.attribute_value_id !== undefined && s.attribute_value_id !== null && s.attribute_value_id !== 0
              ? Number(s.attribute_value_id)
              : null
            const valText = (s.value_text as string) ?? null

            if (!map[aid]) {
              map[aid] = { attribute_id: aid, option_ids: [] }
            }

            // multi-select: accumulate option_ids
            if (optId) {
              map[aid].option_ids = map[aid].option_ids || []
              map[aid].option_ids.push(optId)
              // Also keep single option_id for convenience (will be used by single-select fields)
              if (!map[aid].option_id) map[aid].option_id = optId
            } else if (valText !== null) {
              map[aid].value_text = valText
            }
          }
          productSpecs.value = Object.values(map) as SpecPayload[]
          if (process.env.NODE_ENV === 'development') {
            // console.debug('💡 Loaded productSpecs', JSON.stringify(productSpecs.value))
          }
        } else {
          productSpecs.value = []
        }
      } catch (err) {
        // در صورت نبودن مشخصات (404) این وضعیت را عادی در نظر می‌گیریم
        const e = err as { status?: number; response?: { status?: number; statusCode?: number } }
        const status = e?.status || e?.response?.status || e?.response?.statusCode
        if (status === 404) {
          productSpecs.value = []
        } else {
          // Failed loading specs
          productSpecs.value = []
        }
      }

      // Load inventory data
      await loadInventoryData(productId)

      // Load categories to ensure they are available in the form
      await loadCategories()

      return product
    } finally {
      isSaving.value = false
      isLoadingProduct.value = false
    }
  }

  // ---------------------------
  // Update product
  // ---------------------------
  async function updateProduct(productId: string | number) {
    try {
      isSaving.value = true

      const payload = {
        name: productForm.name,
        name_en: productForm.englishName,
        slug: typeof productForm.slug === 'string' ? productForm.slug.trim() : '',
        description: productForm.description,
        full_description: productForm.fullDescription,
        status: productForm.status,
        brand_id: productForm.brand_id ? Number(productForm.brand_id) : null,
        seo_title: productForm.seo_title || productForm.name,
        meta_description: productForm.meta_description || '',
        url: productForm.url || '',
        images: images.value,
        // اولویت با دستهبندی فرعی است؛ در غیر این صورت، دستهبندی اصلی ذخیره میشود
        category_id: productForm.sub_category_id ? Number(productForm.sub_category_id) : (productForm.category_id ? Number(productForm.category_id) : null)
      }



      const response = await $fetch(`/api/admin/products/${productId}`, {
        method: 'PUT',
        body: payload
      })

      // send specifications only if user provided values
      if (response && (response as Record<string, unknown>).id && productSpecs.value.length > 0) {
        await saveSpecsToBackend((response as Record<string, unknown>).id as number)
      }

      // Save pricing data after core product updated
      try {
        await savePricingData(productId)
      } catch {
        // Failed to save pricing data after update
      }

      // Save inventory data after core product updated
      try {
        await saveInventoryData(productId)
      } catch {
        // Failed to save inventory data after update
      }

      return response
    } finally {
      isSaving.value = false
    }
  }

  // ---------------------------
  // Saving handlers (create new)
  // ---------------------------
  async function saveProduct() {
    try {
      isSaving.value = true
      const payload = {
        name: productForm.name,
        name_en: productForm.englishName,
        slug: typeof productForm.slug === 'string' ? productForm.slug.trim() : '',
        description: productForm.description,
        full_description: productForm.fullDescription,
        status: productForm.status,
        brand_id: productForm.brand_id ? Number(productForm.brand_id) : null,
        seo_title: productForm.seo_title || productForm.name,
        meta_description: productForm.meta_description || '',
        url: productForm.url || '',
        images: images.value,
        // قیمت و موجودی اولیه در همان create ارسال می‌شود تا به روت‌های مجزا وابسته نباشیم
        price: Number.isFinite(Number(pricingForm.price)) ? Number(pricingForm.price) : 0,
        stock_quantity: Number.isFinite(Number(inventoryForm.stock_quantity)) ? Number(inventoryForm.stock_quantity) : 0,
        // اولویت با دستهبندی فرعی است؛ در غیر این صورت، دستهبندی اصلی ذخیره میشود
        category_id: productForm.sub_category_id ? Number(productForm.sub_category_id) : (productForm.category_id ? Number(productForm.category_id) : null),
        sub_category_id: productForm.sub_category_id ? Number(productForm.sub_category_id) : null
      }



      let response;
      try {
        if (!isEditMode.value) {
          if (!productForm.name || !productForm.name.trim()) {
            notifier.error('نام محصول الزامی است', 'خطای اعتبارسنجی')
            const err = new Error('validation:name');
            (err as unknown as { skipToast: boolean }).skipToast = true
            throw err
          }
          if (!productForm.sub_category_id && !productForm.category_id) {
            notifier.error('دسته‌بندی الزامی است', 'خطای اعتبارسنجی')
            const err = new Error('validation:category');
            (err as unknown as { skipToast: boolean }).skipToast = true
            throw err
          }
        }
        if (isEditMode.value && editingProductId.value) {
          // Update existing product
          response = await $fetch(`/api/admin/products/${editingProductId.value}`, {
            method: 'PUT',
            body: payload
          })
        } else {
          // Create new product
          response = await $fetch('/api/admin/products', {
            method: 'POST',
            body: payload
          })
          // بعد از ساخت موفق، حالت ویرایش فعال شود و موجودی ذخیره شود
          if (response && response.id) {
            isEditMode.value = true;
            editingProductId.value = response.id;
          }
        }

        // همگام‌سازی فوری SKU و slug با پاسخ سرور تا لینک پیش‌نمایش درست ساخته شود
        if (response && (response as Record<string, unknown>).sku) {
          productForm.sku = String((response as Record<string, unknown>).sku)
        }
        if (response && (response as Record<string, unknown>).slug) {
          productForm.slug = String((response as Record<string, unknown>).slug)
        }

        // Save specifications only if user provided any values
        if (response && (response as Record<string, unknown>).id && productSpecs.value.length > 0) {
          await saveSpecsToBackend((response as Record<string, unknown>).id as number).catch(() => { })
        }

        return response // product object
      } catch (fetchError) {
        // خطا در ذخیره محصول
        throw fetchError;
      }
    } finally {
      isSaving.value = false
    }
  }

  // A simpler function just for creating a new product without state side-effects
  async function createProduct() {
    isSaving.value = true
    try {
      // پیش‌اعتبارسنجی سمت کلاینت
      if (!productForm.name || !productForm.name.trim()) {
        notifier.error('نام محصول الزامی است', 'خطای اعتبارسنجی')
        const err = new Error('validation:name');
        (err as unknown as { skipToast: boolean }).skipToast = true
        throw err
      }
      if (!productForm.sub_category_id && !productForm.category_id) {
        notifier.error('دسته‌بندی الزامی است', 'خطای اعتبارسنجی')
        const err = new Error('validation:category');
        (err as unknown as { skipToast: boolean }).skipToast = true
        throw err
      }
      // Debug: بررسی مقدار قیمت

      const payload = {
        name: productForm.name?.trim() || '',
        name_en: productForm.englishName?.trim() || '',
        slug: typeof productForm.slug === 'string' ? productForm.slug.trim() : '',
        description: productForm.description || '',
        description_en: productForm.fullDescription || '',
        price: Number.isFinite(Number(pricingForm.price)) ? Number(pricingForm.price) : 0,
        sale_price: null as number | null,
        stock_quantity: Number.isFinite(Number(inventoryForm.stock_quantity)) ? Number(inventoryForm.stock_quantity) : 0,
        weight: null as number | null,
        dimensions: '',
        status: productForm.status || 'active',
        brand_id: productForm.brand_id ? Number(productForm.brand_id) : null,
        seo_title: productForm.name || '',
        meta_description: productForm.meta_description || '',
        url: productForm.url || '',
        images: images.value,
        // اولویت با دستهبندی فرعی است؛ در غیر این صورت، دستهبندی اصلی ذخیره می‌شود
        category_id: productForm.sub_category_id ? Number(productForm.sub_category_id) : (productForm.category_id ? Number(productForm.category_id) : null),
        sub_category_id: productForm.sub_category_id ? Number(productForm.sub_category_id) : null
      }

      // Debug - Final payload price


      try {
        const response = await $fetch<Record<string, unknown>>('/api/admin/products', {
          method: 'POST',
          body: payload
        })
        // بعد از ایجاد موفق محصول، اطلاعات تکمیلی را ذخیره کنیم
        if (response && response.id) {
          // فقط در صورت داشتن مشخصات، ذخیره کن (به صورت موازی)
          const followUps: Promise<void>[] = []
          if (productSpecs.value && productSpecs.value.length > 0) {
            followUps.push(saveSpecsToBackend(response.id as number).catch(() => { }))
          }
          // قیمت و موجودی در create ارسال شده‌اند، نیازی به فراخوانی روت‌های جدا نیست
          if (followUps.length > 0) {
            await Promise.allSettled(followUps)
          }
        }

        return response;
      } catch (fetchError) {
        // خطا در ایجاد محصول
        const e = fetchError as Record<string, unknown>
        const msg = (e?.response as { _data?: { user_message?: string } })?._data?.user_message || (e?.message as string) || 'خطا در ایجاد محصول'
        notifier.error(msg, 'خطا')
        throw fetchError
      }
    } finally {
      isSaving.value = false
    }
  }

  async function saveAndContinueEditing() {
    // This function will create or update the product
    const product = await saveProduct()

    if (product && product.id) {
      // فرانت حق ذخیره یا تولید شناسه ندارد؛ فقط از Backend می‌خوانیم

      // After a successful save, we re-load the product data to ensure the form 
      // is populated with the very latest data from the server (like SKU, etc.).
      await loadProductForEdit(product.id)
      autoSaveReady.value = true

      // Also, ensure the category list is still loaded.
      await loadCategories()

      return product // Return the full product object
    }

    return null // Return null on failure
  }

  // ---------------------------
  // AI content generation (stub)
  // ---------------------------
  function generateAIContent(_type: 'short' | 'full') {
    const n = useNotifier(); n.info('قابلیت تولید محتوا با AI به زودی اضافه خواهد شد')
  }

  const categories = ref<Record<string, unknown>[]>([])
  const brands = ref<Record<string, unknown>[]>([])

  /**
   * Load product brands.
   */
  async function loadBrands() {
    try {
      const response = await $fetch<Record<string, unknown>[]>('/api/admin/brands')
      if (Array.isArray(response)) {
        brands.value = response.map(brand => ({
          ...brand,
          id: Number(brand.id),
          label: brand.name,
          value: Number(brand.id)
        }))
      } else {
        brands.value = []
      }
    } catch {
      brands.value = []
    }
  }

  /**
   * Load product categories.
   * @param all  when true, fetches *all* categories (published + unpublished).
   *             Defaults to true when in edit mode so that the product's current
   *             category (even if unpublished) is present in the dropdown.
   */
  async function loadCategories(all = true as boolean) {

    isCategoriesLoading.value = true
    try {
      const endpoint = all ? '/api/admin/product-categories?all=1' : '/api/admin/product-categories'
      // جلوگیری از خطای TypeScript 'Excessive stack depth' روی امضای $fetch
      const response = await $fetch<Record<string, unknown>[]>(endpoint)

      // The API now reliably returns an array of category objects.
      if (Array.isArray(response)) {
        categories.value = response.map(cat => ({
          ...cat,
          id: Number(cat.id),
          parent_id: cat.parent_id !== undefined && cat.parent_id !== null ? Number(cat.parent_id) : null,
          label: cat.name,
          value: Number(cat.id)
        }))
      } else {
        // Handle cases where the response is not an array, maybe log an error.
        categories.value = []
      }
    } catch {
      // Failed to load categories
      categories.value = [] // Ensure categories are empty on error
    } finally {
      isCategoriesLoading.value = false
    }

  }

  async function saveSpecsToBackend(prodId: number) {
    const vals = productSpecs.value ?? []
    try {
      // همواره از شناسهٔ تثبیت‌شده در Store استفاده می‌کنیم تا اشتباهاً با ID قدیمی ارسال نشود
      const id = Number(editingProductId.value ?? prodId)
      const res = await $fetch.raw(`/api/admin/products/${id}/specs`, {
        method: 'POST',
        body: { values: vals },
        retry: 0
      })

      if (!res.ok) {
        const d = res._data as Record<string, unknown>
        const msg = (d?.user_message as string) || (d?.error as string) || (d?.statusMessage as string) || (d?.message as string) || `خطا در ذخیره مشخصات (کد ${res.status})`
        notifier.error(msg, 'خطا در ذخیره مشخصات')
        throw new Error(msg)
      }
    } catch (err) {
      // Failed saving specs
      const e = err as Record<string, unknown>
      let msg = (e?.message as string) || 'خطا در ذخیره مشخصات فنی'
      if ((e?.response as { _data?: unknown })?._data) {
        msg = resolveErrorMessage((e.response as { _data: unknown })._data)
      }
      notifier.error(msg, 'خطا')
      throw err
    }
  }

  // ---------------------------
  // Pricing operations
  // ---------------------------
  async function loadPricingData(productId: string | number) {
    try {
      const response = await $fetch<Record<string, unknown>>(`/api/product-prices/${productId}`)
      if (response) {
        pricingForm.price = (response.price as number) || 0
        pricingForm.old_price = (response.old_price as number) || 0
        pricingForm.cost = (response.cost as number) || 0
        pricingForm.discount_percent = (response.discount_percent as number) || 0
        pricingForm.discount_amount = (response.discount_amount as number) || 0
        pricingForm.profit = (response.profit as number) || 0
        pricingForm.disableBuyButton = (response.disable_buy_button as boolean) || false
        pricingForm.callForPrice = (response.call_for_price as boolean) || false
        pricingForm.sale_price = (response.sale_price as number) || 0
        // زمان‌بندی قیمت ویژه
        pricingForm.sale_start_at = (response.sale_start_at as string) || null
        pricingForm.sale_end_at = (response.sale_end_at as string) || null
        pricingForm.sale_start_jalali = formatJalaliFromISO(pricingForm.sale_start_at)
        pricingForm.sale_end_jalali = formatJalaliFromISO(pricingForm.sale_end_at)
        // پله‌های فروش ویژه
        specialOffers.value = Array.isArray(response.special_offers)
          ? (response.special_offers as Record<string, unknown>[]).map((o: Record<string, unknown>) => ({
            base_price: Number(o.base_price ?? pricingForm.price ?? 0),
            price: Number(o.price) || 0,
            quantity: Number(o.quantity) || 0
          }))
          : []
        pricingLoaded.value = true
      }
    } catch {
      // خطا در بارگیری قیمت‌ها
    }
  }

  async function savePricingData(productId: string | number) {
    try {
      // همواره ترجیح با شناسه فعلی در Store است تا از ارسال با ID قدیمی جلوگیری شود
      let numericId: number | null = Number(editingProductId.value)
      if (!Number.isFinite(numericId)) {
        numericId = Number(productId)
      }
      if (!Number.isFinite(numericId)) {
        try {
          const prod = await $fetch<Record<string, unknown>>(`/api/admin/products/${productId}`)
          numericId = Number(prod?.id)
        } catch { }
      }
      if (!Number.isFinite(numericId)) throw new Error('invalid_product_id')

      // اگر sale_price مشخص نشده ولی حداقل یک پلهٔ فروش ویژه وجود دارد،
      // اولین پله را به عنوان sale_price برای سازگاری صفحات فعلی ارسال می‌کنیم.
      let effectiveSalePrice: number | null | undefined = pricingForm.sale_price
      if ((!effectiveSalePrice || effectiveSalePrice === 0) && Array.isArray(specialOffers.value) && specialOffers.value.length > 0) {
        effectiveSalePrice = specialOffers.value[0].price || 0
      }

      const payload = {
        price: pricingForm.price,
        old_price: pricingForm.old_price,
        cost: pricingForm.cost,
        discount_percent: pricingForm.discount_percent,
        discount_amount: pricingForm.discount_amount,
        disable_buy_button: pricingForm.disableBuyButton,
        call_for_price: pricingForm.callForPrice,
        sale_price: effectiveSalePrice,
        sale_start_at: pricingForm.sale_start_at,
        sale_end_at: pricingForm.sale_end_at,
        special_offers: (specialOffers.value || []).map((o, idx) => ({ base_price: Number(o.base_price || pricingForm.price || 0), price: Number(o.price) || 0, quantity: Number(o.quantity) || 0, sort_order: idx + 1 }))
      }

      const response = await $fetch<Record<string, unknown>>(`/api/product-prices/${numericId}`, {
        method: 'PUT',
        body: payload
      })

      // Update computed fields from response
      if (response && (response as { product?: { profit?: number } }).product) {
        pricingForm.profit = (response as { product: { profit: number } }).product.profit || 0
      }

      return response
    } catch (error) {
      // خطا در ذخیره قیمت‌ها
      throw error
    }
  }

  // ---------------------------
  // Inventory operations
  // ---------------------------
  async function loadInventoryData(productId: string | number) {
    try {
      const response = await $fetch<Record<string, unknown>>(`/api/product-inventories/${productId}`)
      if (response) {
        inventoryForm.stock_quantity = (response.stock_quantity as number) || 0
        inventoryForm.min_stock_quantity = (response.min_stock_quantity as number) || 0
        inventoryForm.max_stock_quantity = (response.max_stock_quantity as number) || 0
        inventoryForm.stock_status = (response.stock_status as string) || 'in_stock'
        inventoryForm.show_stock_to_customer = !!response.show_stock_to_customer
        inventoryForm.track_inventory = response.track_inventory !== undefined ? !!response.track_inventory : true
        inventoryForm.allow_reservation = !!response.allow_reservation
        // اگر بک‌اند انبار پیش‌فرض را برگرداند، همان را ست کن؛ وگرنه 1st warehouse
        if (response.default_warehouse_id) {
          inventoryForm.warehouse_id = Number(response.default_warehouse_id)
          inventoryForm.shipping_enabled = true
        }
      }
    } catch (error) {
      // اگر رکورد موجودی هنوز ایجاد نشده باشد، 404 حالت طبیعی است
      const e = error as Record<string, unknown>
      const status = (e?.status as number) || (e?.response as { status?: number; statusCode?: number })?.status || (e?.response as { status?: number; statusCode?: number })?.statusCode
      if (status === 404) {
        // مقداردهی پیش‌فرض برای فرم موجودی
        inventoryForm.stock_quantity = 0
        inventoryForm.min_stock_quantity = 0
        inventoryForm.max_stock_quantity = 0
        inventoryForm.stock_status = 'in_stock'
        inventoryForm.show_stock_to_customer = false
        inventoryForm.track_inventory = true
      } else {
        // خطا در بارگیری موجودی
      }
    }
  }

  async function saveInventoryData(productId: string | number) {
    try {
      // همواره ترجیح با شناسه فعلی در Store است تا از ارسال با ID قدیمی جلوگیری شود
      let numericId: number | null = Number(editingProductId.value)
      if (!Number.isFinite(numericId)) {
        numericId = Number(productId)
      }
      if (!Number.isFinite(numericId)) {
        try {
          const prod = await $fetch<Record<string, unknown>>(`/api/admin/products/${productId}`)
          numericId = Number(prod?.id)
        } catch { }
      }
      if (!Number.isFinite(numericId)) throw new Error('invalid_product_id')

      const payload = {
        stock_quantity: inventoryForm.stock_quantity,
        min_stock_quantity: inventoryForm.min_stock_quantity,
        max_stock_quantity: inventoryForm.max_stock_quantity,
        stock_status: inventoryForm.stock_status,
        show_stock_to_customer: inventoryForm.show_stock_to_customer,
        track_inventory: inventoryForm.track_inventory,
        allow_reservation: inventoryForm.allow_reservation
      }

      await $fetch(`/api/product-inventories/${numericId}`, {
        method: 'PUT',
        body: payload,
      })

      // قبلاً اینجا موجودی انبار پیش‌فرض را هم‌زمان ست می‌کردیم که باعث بازنویسی و صفر شدن می‌شد.
      // از این پس موجودی هر انبار فقط از طریق جدول سطری (PUT /api/product-warehouse-stocks/:id) بروزرسانی می‌شود.
    } catch (error) {
      // خطا در ذخیره موجودی
      throw error
    }
  }

  // ---------------------------
  // Shipping operations
  // ---------------------------
  async function loadShippingData(productId: string | number) {
    try {
      const response = await $fetch<Record<string, unknown>>(`/api/product-shipping/${productId}`)
      if (response) {
        // map to local UI state if needed (shipping.vue currently uses local inputs)
        // This enables prefill when editing existing product
      }
    } catch {
      // خطا در بارگیری اطلاعات حمل‌ونقل
    }
  }

  async function saveShippingData(productId: string | number, payload: { weight?: number; length?: number; width?: number; height?: number; shipping_cost?: number; shipping_time?: number }) {
    try {
      await $fetch(`/api/product-shipping/${productId}`, {
        method: 'PUT',
        body: payload
      })
    } catch (error) {
      // خطا در ذخیره اطلاعات حمل‌ونقل
      throw error
    }
  }

  // Auto-save inventory when values change (edit mode only)
  // @ts-ignore
  watch(() => [
    inventoryForm.stock_quantity,
    inventoryForm.min_stock_quantity,
    inventoryForm.max_stock_quantity,
    inventoryForm.stock_status,
    inventoryForm.show_stock_to_customer,
    inventoryForm.track_inventory,
    inventoryForm.allow_reservation,
  ], async () => {
    if (isEditMode.value && editingProductId.value) {
      await saveInventoryData(editingProductId.value)
    }
  })

  return {
    // state
    isSaving,
    isEditMode,
    editingProductId,
    isLoadingProduct,
    isCategoriesLoading,
    productForm,
    tinyApiKey,
    sections,
    images,
    categories,
    brands,
    productSpecs,
    pricingForm,
    specialOffers,
    computedProfit,
    computedProfitPercent,
    computedDiscountAmount,
    computedDiscountPercent,
    inventoryForm,
    // actions
    toggleSection,
    addImages,
    removeImage,
    resetForm,
    $reset,
    loadProductForEdit,
    updateProduct,
    saveProduct,
    createProduct,
    saveAndContinueEditing,
    generateAIContent,
    loadCategories,
    loadBrands,
    saveSpecsToBackend,
    loadPricingData,
    savePricingData,
    loadInventoryData,
    saveInventoryData,
    loadShippingData,
    saveShippingData
  }
})