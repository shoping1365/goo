<script setup lang="ts">
import { computed, inject, isRef, onMounted, reactive, ref, watch, watchEffect } from 'vue';
import { useRoute } from 'vue-router';
import { useSlugManagement } from '~/composables/useSlugManagement';

// Props
const props = defineProps<{ 
  slug?: string; 
  defaultTitle?: string; 
  enableOgType?: boolean; 
  metaTitle?: string; 
  initialMetaDescription?: string;
  url?: string;
  ogImage?: string; 
  ogTitle?: string;
  ogDescription?: string;
  ogType?: string;
  ogSiteName?: string;
  indexStatus?: 'index' | 'noindex';
  followStatus?: 'follow' | 'nofollow';
}>()
const emit = defineEmits([
  'update:slug', 
  'select-og-image', 
  'update:metaTitle', 
  'update:metaDescription',
  'update:url',
  'update:ogTitle',
  'update:ogDescription',
  'update:ogType',
  'update:ogSiteName',
  'update:ogImage',
  'update:indexStatus',
  'update:followStatus'
])

// Local reactive slug linked to input


const slug = ref(props.slug || '')
const seoTitle = ref(props.metaTitle ?? props.defaultTitle ?? '')
const storedUrl = ref(props.url || '')

// دریافت slug از productStore اگر موجود باشد
const productStore = inject('productStore', null)
const route = useRoute()

// استفاده از composable
const { checkSlugUnique: checkSlugUniqueAPI, generateUniqueSlug: generateUniqueSlugAPI, generateSlugFromTitle } = useSlugManagement()
// const { buildProductLink } = useProductLink()

// Debug: بررسی inject (فقط در development)
if (process.env.NODE_ENV === 'development') {
  // console.log('ProductStore injected:', !!productStore)
  if (productStore) {
    // console.log('✅ ProductStore successfully injected with data:', {
    //   editingProductId: productStore.editingProductId,
    //   productForm: productStore.productForm
    // })
  } else {
    // console.warn('⚠️ productStore inject نشده است!')
  }
}

// Initialize from productStore if available
if (productStore?.productForm?.slug && !slug.value) {
  slug.value = productStore.productForm.slug
}
if (productStore?.productForm?.url) {
  storedUrl.value = productStore.productForm.url
  // console.log('✅ Initial URL from ProductStore:', productStore.productForm.url)
}

// Watch برای به‌روزرسانی از productStore
watch(() => productStore?.productForm, (newForm) => {
  if (newForm) {
    if (process.env.NODE_ENV === 'development') {
      // console.log('ProductForm updated:', newForm)
    }
    
    // Update slug if not manually changed
    if (newForm.slug && !slugTouched.value) {
      slug.value = newForm.slug
    }
    
    // Update URL
    if (newForm.url) {
      storedUrl.value = newForm.url
      // console.log('✅ URL updated from ProductForm:', newForm.url)
    }
  }
}, { deep: true, immediate: true })

// اگر productStore موجود نیست، از API استفاده کن
const productData = ref(null)
const isLoadingProduct = ref(false)

const loadProductData = async () => {
  // اگر store موجود است و URL هم دارد، نیازی به API نیست
  if (productStore?.productForm?.sku && productStore?.productForm?.url) {
    // console.log('✅ ProductStore has URL, no need for API:', productStore.productForm.url)
    return
  }
  
  // ابتدا از route.params.id استفاده کن، اگر موجود نبود از query.id استفاده کن
  let productId = route.params.id
  if (!productId) {
    productId = route.query.id as string
  }
  
  if (!productId) {
    // console.warn('⚠️ No product ID found in route params or query:', {
    //   params: route.params,
    //   query: route.query
    // })
    return
  }
  
  isLoadingProduct.value = true
  try {
    // console.log('📡 Loading product data from API for ID:', productId)
    const response = await $fetch<Record<string, unknown>>(`/api/admin/products/${productId}`)
    productData.value = response
    // console.log('✅ Product data loaded:', response)
    
    // اگر URL در response موجود است، آن را تنظیم کن
    if (response?.url) {
      storedUrl.value = (response.url as string)
      // console.log('✅ URL updated from API:', response.url)
    } else {
      // console.warn('⚠️ No URL found in API response:', response)
    }
  } catch (error) {
    // console.error('❌ Error loading product data:', error)
  } finally {
    isLoadingProduct.value = false
  }
}

// Load product data on mount
onMounted(() => {
  // اگر productStore موجود نیست یا URL ندارد، از API استفاده کن
  if (!productStore || !productStore.productForm?.url) {
    // console.log('🔄 ProductStore not available or URL missing, loading from API...')
    loadProductData()
  } else {
    // console.log('✅ ProductStore available with URL:', productStore.productForm.url)
  }
})

// Watch برای به‌روزرسانی وقتی productData لود شد
watch(() => productData.value, (newData) => {
  if (newData) {
    // console.log('🔄 ProductData loaded, updating URL:', newData)
    // اگر slug در productData موجود است، آن را تنظیم کن
    const data = newData as { slug?: string; url?: string }
    if (data.slug && !slug.value) {
      slug.value = data.slug
    }
    // اگر URL در productData موجود است، آن را تنظیم کن
    if (data.url) {
      storedUrl.value = data.url
    }
  }
}, { immediate: true })

// Watch برای به‌روزرسانی وقتی productStore تغییر کرد
watch(() => productStore, (newStore) => {
  if (newStore && newStore.productForm) {
    // console.log('🔄 ProductStore updated:', newStore.productForm)
    // اگر URL در store موجود است، آن را تنظیم کن
    if (newStore.productForm.url) {
      storedUrl.value = newStore.productForm.url
      // console.log('✅ URL updated from ProductStore:', newStore.productForm.url)
    }
  }
}, { immediate: true })

// متغیرهای جدید برای مدیریت slug
const slugTouched = ref(false)
const slugError = ref('')
const isCheckingSlug = ref(false)
const isGeneratingSlug = ref(false)

// Debug: بررسی وضعیت productStore
if (process.env.NODE_ENV === 'development') {
  // console.log('Product URL Debug:', {
  //   productSku: productStore?.productForm?.sku,
  //   productId: productStore?.editingProductId,
  //   routeId: route.params.id,
  //   queryId: route.query.id,
  //   englishName: productStore?.productForm?.englishName,
  //   currentSlug: slug.value,
  //   storedUrl: storedUrl.value,
  //   productStore: !!productStore,
  //   productForm: productStore?.productForm,
  //   routeParams: route.params,
  //   routeQuery: route.query
  // })
  
  if (!productStore) {
    // console.warn('⚠️ هیچ SKU یا ID محصولی یافت نشد!', { 
    //   productStore: false, 
    //   routeId: route.params.id, 
    //   queryId: route.query.id,
    //   routeParams: route.params,
    //   routeQuery: route.query,
    //   productForm: undefined 
    // })
  } else {
    // console.log('✅ ProductStore available with data:', {
    //   editingProductId: productStore.editingProductId,
    //   sku: productStore.productForm?.sku,
    //   url: productStore.productForm?.url
    // })
  }
}

// Fallback ID برای زمانی که productStore موجود نیست
// const fallbackId = computed(() => {
//   if (productStore?.editingProductId) return productStore.editingProductId
//   if (productStore?.productForm?.sku) return productStore.productForm.sku
//   if (route.params.id) return route.params.id
//   if (route.query.id) return route.query.id as string
//   return '[id]'
// })

if (process.env.NODE_ENV === 'development') {
  // console.log('🔄 Using fallback ID:', fallbackId)
}
// const { showSuccess, showError, showWarning, showInfo } = useToast()

// Try to get title provided by parent via provide('pageTitle')
const providedTitle = inject<unknown>('pageTitle', null)
const seoTitleTouched = ref(false)

function getProvided(){ return providedTitle ? (isRef(providedTitle) ? providedTitle.value : providedTitle) : '' }

if (!props.defaultTitle && !seoTitle.value && providedTitle) {
  seoTitle.value = (getProvided() as string) || ''
}

if (providedTitle) {
  watch(() => getProvided(), v => {
    if (!seoTitleTouched.value) seoTitle.value = (v as string) || ''
  })
}

// همگام‌سازی prop با مقدار داخلی
watch(() => props.metaTitle, (val) => {
  if (val !== undefined && val !== seoTitle.value) {
    seoTitle.value = val
  }
})

// اگر مقدار اولیه defaultTitle داده شده بود و metaTitle نبود، مقداردهی کن
if (props.defaultTitle && !props.metaTitle) {
  seoTitle.value = props.defaultTitle
}

// هر بار که seoTitle تغییر کرد، مقدار را به والد emit کن
watch(seoTitle, (val) => {
  emit('update:metaTitle', val)
  seoTitleTouched.value = true
})

// Meta Description live length feedback ------------------------------------
const metaDescription = ref(props.initialMetaDescription || '')
const metaDescriptionTouched = ref(false)
watch(() => props.initialMetaDescription, (val) => {
  if (val !== undefined && val !== metaDescription.value) {
    metaDescription.value = val || ''
  }
})
const metaLength = computed(()=> metaDescription.value.length)
const metaQuality = computed(()=>{
  const len = metaLength.value
  if(len < 100) return 'کم'
  if(len < 150) return 'متوسط'
  if(len <= 160) return 'خوب'
  return 'زیاد'
})
const metaColor = computed(()=>{
  switch(metaQuality.value){
    case 'کم': return 'bg-red-500'
    case 'متوسط': return 'bg-yellow-500'
    case 'خوب': return 'bg-green-500'
    default: return 'bg-red-500'
  }
})

// همگام‌سازی metaDescription با والد
watch(metaDescription, (val) => {
  emit('update:metaDescription', val)
  metaDescriptionTouched.value = true
})

// compute for title
const titleLength = computed(()=> seoTitle.value.length)
const titleQuality = computed(()=>{
  const len = titleLength.value
  if(len < 30) return 'کم'
  if(len < 50) return 'متوسط'
  if(len <= 60) return 'خوب'
  return 'زیاد'
})
const titleColor = computed(()=>{
  switch(titleQuality.value){
    case 'کم': return 'bg-red-500'
    case 'متوسط': return 'bg-yellow-500'
    case 'خوب': return 'bg-green-500'
    default: return 'bg-red-500'
  }
})

// slug length
watch(() => props.slug, val => slug.value = val || '')
watch(() => props.defaultTitle, val => { if(!seoTitle.value) seoTitle.value = val || '' })
watch(slug, val => emit('update:slug', val))

// تابع بررسی تکراری بودن slug
const checkSlugUnique = async (slugToCheck: string) => {
  if (!slugToCheck.trim()) {
    slugError.value = 'URL الزامی است'
    return false
  }
  
  isCheckingSlug.value = true
  slugError.value = ''
  
  try {
    const response = await checkSlugUniqueAPI(slugToCheck)
    if (response.exists) {
      slugError.value = 'این URL قبلاً استفاده شده است'
      return false
    }
    return true
  } catch (error) {
    console.error('خطا در بررسی URL:', error)
    slugError.value = error instanceof Error ? error.message : 'خطا در بررسی URL'
    return false
  } finally {
    isCheckingSlug.value = false
  }
}

// تابع تولید slug یکتا
const generateUniqueSlug = async (baseSlug: string) => {
  isGeneratingSlug.value = true
  
  try {
    const uniqueSlug = await generateUniqueSlugAPI(baseSlug)
    slug.value = uniqueSlug
    emit('update:slug', uniqueSlug)
    slugError.value = ''
  } catch (error) {
    console.error('خطا در تولید URL یکتا:', error)
    slugError.value = error instanceof Error ? error.message : 'خطا در تولید URL یکتا'
  } finally {
    isGeneratingSlug.value = false
  }
}

// تنظیم touched برای slug
watch(slug, () => {
  slugTouched.value = true
})

// همگام‌سازی خودکار slug از عنوان مقاله
watch(seoTitle, async (newTitle) => {
  // اگر کاربر هنوز slug را دستی تغییر نداده باشد
  if (!slugTouched.value && newTitle.trim()) {
    try {
      const uniqueSlug = await generateSlugFromTitle(newTitle)
      slug.value = uniqueSlug
      emit('update:slug', uniqueSlug)
      slugError.value = ''
    } catch (error) {
      console.error('خطا در تولید خودکار URL از عنوان:', error)
    }
  }
})

// همگام‌سازی خودکار seoTitle با defaultTitle
watch(() => props.defaultTitle, (newTitle) => {
  if (newTitle && !seoTitleTouched.value) {
    seoTitle.value = String(newTitle)
  }
}, { immediate: true })

// Section collapse/expand states for SEO tab
const sections = reactive({
  seoBasic: false,
  seoSocial: false,
  seoSchema: false,
  seoAnalysis: false
})

const toggleSection = (section) => {
  sections[section] = !sections[section]
}

// تابع تبدیل عنوان انگلیسی به slug مناسب
const convertEnglishToSlug = (englishTitle: string): string => {
  if (!englishTitle) return ''
  
  return englishTitle
    .toLowerCase()
    .replace(/[^a-z0-9\s-]/g, '') // حذف کاراکترهای غیرمجاز
    .replace(/\s+/g, '-') // جایگزینی فاصله با خط تیره
    .replace(/-+/g, '-') // حذف خط تیره‌های تکراری
    .replace(/^-|-$/g, '') // حذف خط تیره از ابتدا و انتها
}

// ------ Product URL (complete product URL) - دقیقاً همان منطق ریدایرکت‌ها -----------------
const productUrl = computed(() => {
  if (import.meta.client) {
    try {
      const url = new URL(window.location.href)
      const baseUrl = url.origin
      
      // استفاده از productStore که در بالا inject شده
      const productSku = productStore?.productForm?.sku || productData.value?.sku || ''
      const productId = productStore?.editingProductId || route.params.id || ''
      const englishName = productStore?.productForm?.englishName || productData.value?.name_en || ''
      const currentSlug = (slug.value || '').toString().trim()
      
      // اگر slug خالی است، از productData استفاده کن
      const productSlug = productData.value?.slug || ''
      
      // تعیین slug مناسب: اول slug فعلی، سپس slug از API، سپس عنوان انگلیسی، در نهایت خالی
      let finalSlug = currentSlug || productSlug
      if (!finalSlug && englishName) {
        // تبدیل عنوان انگلیسی به slug مناسب
        finalSlug = convertEnglishToSlug(englishName)
      }
      
      // Debug: نمایش اطلاعات دریافتی
      if (process.env.NODE_ENV === 'development') {
        // console.log('Product URL Debug:', {
        //   productSku: productStore?.productForm?.sku,
        //   productId: productStore?.editingProductId,
        //   routeId: route.params.id,
        //   queryId: route.query.id,
        //   englishName: productStore?.productForm?.englishName,
        //   currentSlug: slug.value,
        //   storedUrl: storedUrl.value,
        //   productStore: !!productStore,
        //   productForm: productStore?.productForm,
        //   routeParams: route.params,
        //   routeQuery: route.query
        // })
        
        if (!productStore) {
          // console.warn('⚠️ هیچ SKU یا ID محصولی یافت نشد!', { 
          //   productStore: false, 
          //   routeId: route.params.id, 
          //   queryId: route.query.id,
          //   routeParams: route.params,
          //   routeQuery: route.query,
          //   productForm: undefined 
          // })
        } else {
          // console.log('✅ ProductStore available with data:', {
          //   editingProductId: productStore.editingProductId,
          //   sku: productStore.productForm?.sku,
          //   url: productStore.productForm?.url
          // })
        }
      }
      
      // اگر هیچ اطلاعاتی موجود نیست، هشدار بده
      if (!productSku && !productId) {
        console.warn('⚠️ هیچ SKU یا ID محصولی یافت نشد!', {
          productStore: !!productStore,
          routeId: route.params.id,
          queryId: route.query.id,
          routeParams: route.params,
          routeQuery: route.query,
          productForm: productStore?.productForm,
          productData: productData.value
        })
      }
      
      // دقیقاً همان منطق ریدایرکت‌ها در WordPress migration
      if (productSku) {
        if (finalSlug) {
          return baseUrl + `/product/sku-${productSku}/${finalSlug}`
        } else {
          return baseUrl + `/product/sku-${productSku}`
        }
      }
      // اگر SKU وجود ندارد ولی ID محصول موجود است، از ID استفاده کن
      else if (productId) {
        if (finalSlug) {
          return baseUrl + `/product/sku-${productId}/${finalSlug}`
        } else {
          return baseUrl + `/product/sku-${productId}`
        }
      }
      // اگر هیچ‌کدام موجود نیست، از route ID استفاده کن
      const fallbackId = route.params.id || '[id]'
      console.log('🔄 Using fallback ID:', fallbackId)
      if (finalSlug) {
        return baseUrl + `/product/sku-${fallbackId}/${finalSlug}`
      } else {
        return baseUrl + `/product/sku-${fallbackId}`
      }
    } catch {
      return '/product/sku-[sku]/'
    }
  }
  return '/product/sku-[sku]/'
})




// ------ Canonical URL (same as product URL) -----------------
const canonicalUrl = ref('')
const canonicalFromSlug = computed(() => {
  // اگر URL در دیتابیس موجود است، از آن استفاده کن
  if (storedUrl.value && storedUrl.value.trim()) {
    return storedUrl.value
  }
  // در غیر این صورت، از URL محاسبه شده استفاده کن
  const computedUrl = productUrl.value
  if (process.env.NODE_ENV === 'development') {
    // console.log('🔄 Canonical URL computed:', { storedUrl: storedUrl.value, computedUrl })
  }
  return computedUrl
})

// Watch برای به‌روزرسانی canonical URL وقتی slug تغییر کرد
watch([slug, storedUrl], ([newSlug, newStoredUrl]) => {
  if (process.env.NODE_ENV === 'development') {
    // اگر slug یا storedUrl تغییر کرده باشد، canonicalUrl را آپدیت کن
    if (slug.value !== newSlug || storedUrl.value !== newStoredUrl) {
      // console.log('🔄 Slug or storedUrl changed:', { newSlug, newStoredUrl })
    }
    
    // console.log('✅ Canonical URL updated:', canonicalUrl.value)
  }
}, { immediate: true })

// Keep exposed canonicalUrl in sync with product URL
watchEffect(() => { canonicalUrl.value = canonicalFromSlug.value })

// ---------------- Open Graph fields -------------------
const ogTitle = ref(props.ogTitle ?? seoTitle.value)
const ogTitleTouched = ref(false)
const ogDescriptionTouched = ref(false)

// تعریف متغیرهای واکنشی برای Open Graph
const ogImage = ref(props.ogImage ?? '')
const ogDescription = ref(props.ogDescription ?? metaDescription.value)
const ogType = ref(props.ogType ?? 'article')
const ogSiteName = ref(props.ogSiteName ?? '')

// Determine if OG type should be shown
const showOgType = computed(() => props.enableOgType !== false)

// همگام‌سازی ogImage داخلی با prop ogImage
watch(() => props.ogImage, (val) => {
  if (val !== undefined && val !== ogImage.value) {
    ogImage.value = val
  }
})

// همگام‌سازی ogDescription داخلی با prop ogDescription
watch(() => props.ogDescription, (val) => {
  if (val !== undefined && val !== ogDescription.value) {
    ogDescription.value = val
  }
})

// همگام‌سازی ogType داخلی با prop ogType
watch(() => props.ogType, (val) => {
  if (val !== undefined && val !== ogType.value) {
    ogType.value = val
  }
})

// همگام‌سازی ogSiteName داخلی با prop ogSiteName
watch(() => props.ogSiteName, (val) => {
  if (val !== undefined && val !== ogSiteName.value) {
    ogSiteName.value = val
  }
})

// همگام‌سازی ogTitle با seoTitle تا زمانی که کاربر دستی ogTitle را تغییر نداده
watch(seoTitle, (val) => {
  if (!ogTitleTouched.value) {
    ogTitle.value = val
  }
})

// اگر مقدار اولیه ogTitle نبود، مقداردهی کن
if (!props.ogTitle) {
  ogTitle.value = seoTitle.value
}

// هر بار که ogTitle تغییر کرد، مقدار را به والد emit کن
watch(ogTitle, (val) => {
  emit('update:ogTitle', val)
  ogTitleTouched.value = true
})

// هر بار که ogImage تغییر کرد، مقدار را به والد emit کن
watch(ogImage, (_val) => {
  emit('update:ogImage', ogImage.value)
})

// هر بار که ogDescription تغییر کرد، مقدار را به والد emit کن
watch(ogDescription, (val) => {
  emit('update:ogDescription', val)
  ogDescriptionTouched.value = true
})

// هر بار که ogType تغییر کرد، مقدار را به والد emit کن
watch(ogType, (val) => {
  emit('update:ogType', val)
})

// هر بار که ogSiteName تغییر کرد، مقدار را به والد emit کن
watch(ogSiteName, (val) => {
  emit('update:ogSiteName', val)
})

watchEffect(() => {
  if (!ogTitle.value) ogTitle.value = seoTitle.value
  if (!ogDescription.value) ogDescription.value = metaDescription.value
})

// ---------------- Index/Follow -------------------
const indexStatus = ref<'index'|'noindex'>(props.indexStatus || 'index')
const followStatus = ref<'follow'|'nofollow'>(props.followStatus || 'follow')

// Watch for prop changes
watch(() => props.indexStatus, (newVal) => {
  if (newVal) indexStatus.value = newVal
})

watch(() => props.followStatus, (newVal) => {
  if (newVal) followStatus.value = newVal
})

// Emit changes
watch(indexStatus, (newVal) => {
  emit('update:indexStatus', newVal)
})

watch(followStatus, (newVal) => {
  emit('update:followStatus', newVal)
})

defineExpose({
  seoTitle,
  metaDescription,
  indexStatus,
  followStatus,
  ogTitle,
  ogDescription,
  ogImage,
  ogType,
  ogSiteName,
  canonicalUrl
})
</script>

<template>
  <!-- تنظیمات اصلی SEO -->
  <div class="bg-gradient-to-br from-emerald-50 to-blue-50 rounded-2xl border border-emerald-200 shadow-lg overflow-hidden mb-8">
    <div class="bg-gradient-to-r from-emerald-600 to-blue-600 px-6 py-4 cursor-pointer" @click="toggleSection('seoBasic')">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 rounded-xl">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold text-white">تنظیمات اصلی SEO</h3>
        </div>
        <div class="p-2 bg-white/20 rounded-lg">
          <span class="text-white font-bold text-lg">{{ sections.seoBasic ? '−' : '+' }}</span>
        </div>
      </div>
    </div>

    <div v-show="sections.seoBasic" class="p-8">
      <div class="grid grid-cols-1 gap-6">
        <!-- عنوان SEO -->
        <div class="bg-white rounded-xl shadow-sm border border-emerald-100 p-6">
          <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
            <div class="p-2 bg-emerald-100 rounded-lg">
              <svg class="w-4 h-4 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
              </svg>
            </div>
            عنوان SEO (Title Tag)
          </label>
          <input v-model="seoTitle" type="text" class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 text-gray-900 transition-all duration-200" dir="rtl" placeholder="عنوان برای موتورهای جستجو" />
          <!-- Progress bar below field -->
          <div class="flex items-center gap-3 mt-3">
            <div class="relative h-2 flex-1 bg-gray-200 rounded overflow-hidden">
              <div :style="{ width: Math.min(titleLength, 60)/60*100 + '%' }" :class="[titleColor,'absolute left-0 top-0 h-full transition-all duration-300'].join(' ')" ></div>
            </div>
            <span class="text-sm font-medium" :class="titleColor.replace('bg','text')">{{ titleLength }} / 60</span>
            <span class="text-xs px-2 py-1 rounded-full" :class="titleColor + ' text-white'">{{ titleQuality }}</span>
          </div>
        </div>

        <!-- توضیحات متا -->
        <div class="bg-white rounded-xl shadow-sm border border-blue-100 p-6">
          <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h7"></path>
              </svg>
            </div>
            توضیحات متا (Meta Description)
          </label>
          
          
          <textarea v-model="metaDescription" class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-gray-900 transition-all duration-200 min-h-[80px] resize-none" dir="rtl" placeholder="توضیح کوتاه محصول برای موتورهای جستجو"></textarea>
          <!-- Progress bar below field -->
          <div class="flex items-center gap-3 mt-3">
            <div class="relative h-2 flex-1 bg-gray-200 rounded overflow-hidden">
              <div :style="{ width: Math.min(metaLength, 160)/160*100 + '%' }" :class="[metaColor,'absolute left-0 top-0 h-full transition-all duration-300'].join(' ')" ></div>
            </div>
            <span class="text-sm font-medium" :class="metaColor.replace('bg','text')">{{ metaLength }} / 160</span>
            <span class="text-xs px-2 py-1 rounded-full" :class="metaColor + ' text-white'">{{ metaQuality }}</span>
          </div>
        </div>



        <!-- URL سفارشی -->
        <div class="bg-white rounded-xl shadow-sm border border-orange-100 p-6">
          <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
            <div class="p-2 bg-orange-100 rounded-lg">
              <svg class="w-4 h-4 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"></path>
              </svg>
            </div>
            URL سفارشی (Slug)
          </label>
          
          <div class="space-y-3">
            <!-- فیلد URL -->
            <div class="relative">
              <input 
                v-model="slug" 
                type="text" 
                :class="[
                  'w-full px-4 py-3 bg-white border rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-orange-500 text-gray-900 transition-all duration-200',
                  slugError ? 'border-red-300 focus:ring-red-500 focus:border-red-500' : 'border-gray-300'
                ]"
                dir="ltr" 
                placeholder="custom-article-url"
                @blur="checkSlugUnique(slug)"
              />
              
              <!-- نمایش خطا -->
              <div v-if="slugError" class="mt-2 text-sm text-red-600 flex items-center gap-2">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
                {{ slugError }}
              </div>
              
              <!-- نمایش موفقیت -->
              <div v-else-if="slug && !isCheckingSlug" class="mt-2 text-sm text-green-600 flex items-center gap-2">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                </svg>
                URL در دسترس است
              </div>
              
              <!-- نمایش loading -->
              <div v-if="isCheckingSlug" class="mt-2 text-sm text-blue-600 flex items-center gap-2">
                <svg class="w-4 h-4 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
                </svg>
                در حال بررسی...
              </div>
            </div>
            
            <!-- دکمه‌های عملیات -->
            <div class="flex items-center gap-3">
              <button 
                v-if="slugError && slug"
                :disabled="isGeneratingSlug"
                class="inline-flex items-center px-3 py-2 text-sm font-medium text-white bg-gradient-to-r from-blue-500 to-indigo-600 hover:from-blue-600 hover:to-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-all duration-200 rounded-lg disabled:opacity-50 disabled:cursor-not-allowed"
                @click="generateUniqueSlug(slug)"
              >
                <svg v-if="isGeneratingSlug" class="w-4 h-4 ml-2 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
                </svg>
                <svg v-else class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
                </svg>
                {{ isGeneratingSlug ? 'در حال تولید...' : 'تولید URL یکتا' }}
              </button>
              
            </div>
          </div>
          
          <p class="text-xs text-gray-500 mt-3">
            <span v-if="!slug">URL الزامی است. URL به طور خودکار از عنوان مقاله تولید می‌شود.</span>
            <span v-else>این URL در آدرس صفحه مقاله نمایش داده می‌شود. پس از ویرایش دستی، URL دیگر تغییر نمی‌کند.</span>
          </p>
        </div>

        <!-- URL کامل محصول -->
        <div class="bg-white rounded-xl shadow-sm border border-blue-100 p-6">
          <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"></path>
              </svg>
            </div>
            URL کامل محصول
          </label>
          
          <div class="space-y-3">
            <!-- فیلد URL کامل -->
            <div class="relative">
              <input 
                v-model="storedUrl" 
                type="text" 
                class="w-full px-4 py-3 bg-gray-50 border border-gray-300 rounded-lg text-gray-600 transition-all duration-200"
                dir="ltr" 
                placeholder="/product/sku-ABC123/product-name"
                readonly
              />
              <!-- نمایش loading اگر URL خالی است -->
              <div v-if="!storedUrl && productStore" class="absolute right-3 top-1/2 transform -translate-y-1/2">
                <svg class="w-4 h-4 animate-spin text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
                </svg>
              </div>
            </div>
          </div>
          
          <p class="text-xs text-gray-500 mt-3">
            <span v-if="!storedUrl">در حال بارگذاری URL محصول...</span>
            <span v-else>این URL کامل محصول است که در دیتابیس ذخیره شده و برای ریدایرکت استفاده می‌شود.</span>
          </p>
        </div>

        <!-- وضعیت نمایه‌سازی -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="bg-white rounded-xl shadow-sm border border-indigo-100 p-6">
            <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
              <div class="p-2 bg-indigo-100 rounded-lg">
                <svg class="w-4 h-4 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                </svg>
              </div>
              وضعیت نمایه‌سازی
            </label>
            <select v-model="indexStatus" class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-900 transition-all duration-200 appearance-none">
              <option value="index">نمایه‌سازی شود (Index)</option>
              <option value="noindex">نمایه‌سازی نشود (No Index)</option>
            </select>
          </div>
          <div class="bg-white rounded-xl shadow-sm border border-teal-100 p-6">
            <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
              <div class="p-2 bg-teal-100 rounded-lg">
                <svg class="w-4 h-4 text-teal-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"></path>
                </svg>
              </div>
              پیگیری لینک‌ها
            </label>
            <select v-model="followStatus" class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-teal-500 text-gray-900 transition-all duration-200 appearance-none">
              <option value="follow">پیگیری شود (Follow)</option>
              <option value="nofollow">پیگیری نشود (No Follow)</option>
            </select>
          </div>
        </div>


        <!-- آدرس کانونیکال -->
        <div class="bg-white rounded-xl shadow-sm border border-rose-100 p-6">
          <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
            <div class="p-2 bg-rose-100 rounded-lg">
              <svg class="w-4 h-4 text-rose-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
            </div>
            آدرس کانونیکال
          </label>
          <div>
            <input :value="canonicalFromSlug" readonly type="text" class="w-full px-4 py-3 bg-gray-50 border border-gray-300 rounded-lg text-gray-900 transition-all duration-200" dir="ltr" placeholder="https://example.com/product/sku-123/slug" />
          </div>
          <p class="text-xs text-gray-500 mt-2">
            <span v-if="canonicalFromSlug.includes('[sku]')">
              ابتدا محصول را ذخیره کنید تا URL نمایش داده شود. از دکمه "ذخیره و ادامه ویرایش" استفاده کنید.
            </span>
            <span v-else>
              URL اصلی این صفحه برای جلوگیری از محتوای تکراری. این آدرس همان URL محصول است و با تغییر آن به‌روزرسانی می‌شود.
            </span>
          </p>
        </div>
      </div>
    </div>
  </div>



  <!-- Open Graph و شبکه‌های اجتماعی -->
  <div class="bg-gradient-to-br from-orange-50 to-pink-50 rounded-2xl border border-orange-200 shadow-lg overflow-hidden mb-8">
    <div class="bg-gradient-to-r from-orange-600 to-pink-600 px-6 py-4 cursor-pointer" @click="toggleSection('seoSocial')">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 rounded-xl">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.367 2.684 3 3 0 00-5.367-2.684z"></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold text-white">Open Graph و شبکه‌های اجتماعی</h3>
        </div>
        <div class="p-2 bg-white/20 rounded-lg">
          <span class="text-white font-bold text-lg">{{ sections.seoSocial ? '−' : '+' }}</span>
        </div>
      </div>
    </div>

    <div v-show="sections.seoSocial" class="p-8">
      <div class="grid grid-cols-1 gap-6">
        <!-- تصویر OG -->
        <div class="md:flex md:flex-row-reverse md:gap-6">
          <!-- ستون تصویر (راست) -->
          <div class="bg-white rounded-xl shadow-sm border border-orange-100 p-6 md:w-[400px]">
            <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
              <div class="p-2 bg-orange-100 rounded-lg">
                <svg class="w-4 h-4 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                </svg>
              </div>
              تصویر Open Graph
            </label>
            <div class="border-2 border-dashed border-gray-300 rounded relative flex items-center justify-center w-full h-[400px]">
              <template v-if="ogImage"><img :src="ogImage" class="object-cover w-full h-full rounded" /></template>
              <template v-else>
                <div class="flex flex-col items-center text-gray-400">
                  <svg xmlns="http://www.w3.org/2000/svg" class="w-8 h-8" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M3 5a2 2 0 012-2h3.172a2 2 0 011.414.586l.828.828A2 2 0 0011.828 5H19a2 2 0 012 2v10a2 2 0 01-2 2H5a2 2 0 01-2-2V5z" /><path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" /></svg>
                  <span class="text-xs mt-1">تصویری انتخاب نشده</span>
                </div>
              </template>
            </div>
            <p class="text-xs text-gray-500 mt-2">تصویری که در شبکه‌های اجتماعی نمایش داده می‌شود (گرفته شده از تصویر شاخص)</p>
          </div>

          <!-- ستون چپ: فیلدها -->
          <div class="flex flex-col gap-6 md:flex-1 mt-6 md:mt-0">
            <!-- عنوان OG -->
            <div class="bg-white rounded-xl shadow-sm border border-pink-100 p-6">
              <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
                <div class="p-2 bg-pink-100 rounded-lg">
                  <svg class="w-4 h-4 text-pink-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
              </svg>
                </div>
                عنوان Open Graph
              </label>
              <input v-model="ogTitle" type="text" class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-pink-500 focus:border-pink-500 text-gray-900 transition-all duration-200" dir="rtl" placeholder="عنوان برای اشتراک در شبکه‌های اجتماعی" />
            </div>

            <!-- توضیحات OG -->
            <div class="bg-white rounded-xl shadow-sm border border-indigo-100 p-6">
              <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
                <div class="p-2 bg-indigo-100 rounded-lg">
                  <svg class="w-4 h-4 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h7"></path>
                  </svg>
                </div>
                توضیحات Open Graph
              </label>
              <textarea v-model="ogDescription" class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-900 transition-all duration-200 min-h-[120px] resize-none" dir="rtl" placeholder="توضیح برای اشتراک در شبکه‌های اجتماعی"></textarea>
            </div>

            <!-- پیش‌نمایش شبکه‌های اجتماعی -->
            <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
              <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-6">
                <div class="p-2 bg-blue-100 rounded-lg">
                  <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                  </svg>
                </div>
                پیش‌نمایش شبکه‌های اجتماعی
              </label>
              
              <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
                <!-- پیش‌نمایش Facebook -->
                <div class="border border-gray-200 rounded-lg overflow-hidden">
                  <div class="bg-blue-600 text-white px-4 py-2 text-sm font-medium">
                    <div class="flex items-center gap-2">
                      <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/>
                      </svg>
                      Facebook
                    </div>
                  </div>
                  <div class="p-6 bg-white">
                    <div class="border border-gray-200 rounded-lg overflow-hidden">
                      <!-- تصویر -->
                      <div class="w-full h-48 bg-gray-100 flex items-center justify-center">
                        <div v-if="ogImage" class="w-full h-full bg-cover bg-center" :style="{ backgroundImage: `url(${ogImage})` }"></div>
                        <div v-else class="text-gray-400 text-center">
                          <svg class="w-12 h-12 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                          </svg>
                          <p class="text-sm">تصویر Open Graph</p>
                        </div>
                      </div>
                      <!-- محتوا -->
                      <div class="p-3">
                        <div class="text-xs text-blue-600 mb-1">{{ ogSiteName || 'نام سایت' }}</div>
                        <h3 class="text-sm font-semibold text-gray-900 mb-2 line-clamp-2">
                          {{ ogTitle || seoTitle || 'عنوان محصول' }}
                        </h3>
                        <p class="text-xs text-gray-600 line-clamp-3">
                          {{ ogDescription || metaDescription || 'توضیحات محصول برای نمایش در شبکه‌های اجتماعی' }}
                        </p>
                      </div>
                    </div>
                  </div>
                </div>

                <!-- پیش‌نمایش Twitter -->
                <div class="border border-gray-200 rounded-lg overflow-hidden">
                  <div class="bg-sky-500 text-white px-4 py-2 text-sm font-medium">
                    <div class="flex items-center gap-2">
                      <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M23.953 4.57a10 10 0 01-2.825.775 4.958 4.958 0 002.163-2.723c-.951.555-2.005.959-3.127 1.184a4.92 4.92 0 00-8.384 4.482C7.69 8.095 4.067 6.13 1.64 3.162a4.822 4.822 0 00-.666 2.475c0 1.71.87 3.213 2.188 4.096a4.904 4.904 0 01-2.228-.616v.06a4.923 4.923 0 003.946 4.827 4.996 4.996 0 01-2.212.085 4.936 4.936 0 004.604 3.417 9.867 9.867 0 01-6.102 2.105c-.39 0-.779-.023-1.17-.067a13.995 13.995 0 007.557 2.209c9.053 0 13.998-7.496 13.998-13.985 0-.21 0-.42-.015-.63A9.935 9.935 0 0024 4.59z"/>
                      </svg>
                      Twitter
                    </div>
                  </div>
                  <div class="p-6 bg-white">
                    <div class="border border-gray-200 rounded-lg overflow-hidden">
                      <!-- تصویر -->
                      <div class="w-full h-48 bg-gray-100 flex items-center justify-center">
                        <div v-if="ogImage" class="w-full h-full bg-cover bg-center" :style="{ backgroundImage: `url(${ogImage})` }"></div>
                        <div v-else class="text-gray-400 text-center">
                          <svg class="w-12 h-12 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                          </svg>
                          <p class="text-sm">تصویر Open Graph</p>
                        </div>
                      </div>
                      <!-- محتوا -->
                      <div class="p-3">
                        <div class="text-xs text-sky-500 mb-1">{{ ogSiteName || 'نام سایت' }}</div>
                        <h3 class="text-sm font-semibold text-gray-900 mb-2 line-clamp-2">
                          {{ ogTitle || seoTitle || 'عنوان محصول' }}
                        </h3>
                        <p class="text-xs text-gray-600 line-clamp-3">
                          {{ ogDescription || metaDescription || 'توضیحات محصول برای نمایش در شبکه‌های اجتماعی' }}
                        </p>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              
              <div class="mt-4 p-3 bg-blue-50 rounded-lg">
                <div class="flex items-start gap-2">
                  <svg class="w-4 h-4 text-blue-600 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                  </svg>
                  <div class="text-xs text-blue-800">
                    <p class="font-medium mb-1">نکات مهم:</p>
                    <ul class="space-y-1 text-blue-700">
                      <li>• تصویر باید حداقل 1200×630 پیکسل باشد</li>
                      <li>• عنوان بهتر است کمتر از 60 کاراکتر باشد</li>
                      <li>• توضیحات بهتر است بین 150-160 کاراکتر باشد</li>
                      <li>• پیش‌نمایش بر اساس Open Graph tags نمایش داده می‌شود</li>
                    </ul>
                  </div>
                </div>
              </div>
            </div>

            <!-- برند/سایت -->
            <div class="flex flex-col gap-2 bg-blue-50 border border-blue-200 rounded-md p-3">
              <label class="text-xs text-gray-700 font-semibold">برند/سایت</label>
              <input v-model="ogSiteName" type="text" class="input input-bordered bg-blue-50 border-gray-300 focus:bg-white focus:border-blue-500 w-full text-right" dir="rtl" placeholder="نام برند یا سایت" />
            </div>

            <!-- نوع محتوا -->
            <template v-if="showOgType">
              <div class="flex flex-col gap-2 bg-blue-50 border border-blue-200 rounded-md p-3">
                <label class="text-xs text-gray-700 font-semibold">نوع محتوا (OG Type)</label>
                <select v-model="ogType" class="input input-bordered bg-blue-50 border-gray-300 focus:bg-white focus:border-blue-500 w-full">
                  <option value="product">محصول</option>
                  <option value="article">مقاله</option>
                  <option value="website">وب‌سایت</option>
                </select>
              </div>
            </template>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Schema و داده‌های ساختاریافته -->
  <div class="bg-gradient-to-br from-purple-50 to-indigo-50 rounded-2xl border border-purple-200 shadow-lg overflow-hidden mb-8">
    <div class="bg-gradient-to-r from-purple-600 to-indigo-600 px-6 py-4 cursor-pointer" @click="toggleSection('seoSchema')">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 rounded-xl">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293ل5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold text-white">Schema و داده‌های ساختاریافته</h3>
        </div>
        <div class="p-2 bg-white/20 rounded-lg">
          <span class="text-white font-bold text-lg">{{ sections.seoSchema ? '−' : '+' }}</span>
        </div>
      </div>
    </div>

    <div v-show="sections.seoSchema" class="p-8">
      <div class="grid grid-cols-1 gap-6">
        <!-- نوع Schema -->
        <div class="bg-white rounded-xl shadow-sm border border-purple-100 p-6">
          <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zM7 3H5a2 2 0 00-2 2v12a4 4 0 004 4h2a2 2 0 002-2V5a2 2 0 00-2-2H7z"></path>
              </svg>
            </div>
            نوع Schema
          </label>
          <select class="کامل w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 text-gray-900 transition-all duration-200 appearance-none">
            <option value="Product">محصول (Product)</option>
            <option value="Book">کتاب (Book)</option>
            <option value="SoftwareApplication">نرم‌افزار (Software)</option>
            <option value="CreativeWork">اثر خلاقانه (Creative Work)</option>
          </select>
        </div>

        <!-- اطلاعات محصول -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="bg-white rounded-xl shadow-sm border border-blue-100 p-6">
            <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
              <div class="p-2 bg-blue-100 rounded-lg">
                <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                </svg>
              </div>
              GTIN/UPC/EAN
            </label>
            <input type="text" class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-gray-900 transition-all duration-200" placeholder="کد بارکود جهانی" />
          </div>
          <div class="bg-white rounded-xl shadow-sm border border-green-100 p-6">
            <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
              <div class="p-2 bg-green-100 rounded-lg">
                <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 20l4-16m2 16l4-16M6 9h14M4 15h14"></path>
                </svg>
              </div>
              MPN (شماره قطعه سازنده)
            </label>
            <input type="text" class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500 text-gray-900 transition-all duration-200" placeholder="شماره قطعه" />
          </div>
        </div>
        <!-- rest omitted for brevity -->
      </div>
    </div>
  </div>

  <!-- آنالیز و گزارش SEO -->
  <div class="bg-gradient-to-br from-slate-50 to-emerald-50 rounded-2xl border border-slate-200 shadow-lg overflow-hidden">
    <div class="bg-gradient-to-r from-slate-700 to-emerald-600 px-6 py-4 cursor-pointer" @click="toggleSection('seoAnalysis')">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 rounded-xl">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2v-5a1.5 1.5 0 013 0v6a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold text-white">آنالیز و گزارش SEO</h3>
        </div>
        <div class="p-2 bg-white/20 rounded-lg">
          <span class="text-white font-bold text-lg">{{ sections.seoAnalysis ? '−' : '+' }}</span>
        </div>
      </div>
    </div>

    <div v-show="sections.seoAnalysis" class="p-8">
      <div class="grid grid-cols-1 gap-6">
        <!-- نمره SEO -->
        <div class="bg-white rounded-2xl shadow-sm border border-emerald-100 p-8 text-center">
          <div class="relative inline-flex items-center justify-center">
            <div class="w-24 h-24 rounded-full bg-gradient-to-br from-emerald-400 to-green-500 flex items-center justify-center shadow-lg">
              <div class="text-3xl font-bold text-white">85</div>
            </div>
            <div class="absolute -top-2 -right-2 w-8 h-8 bg-yellow-400 rounded-full flex items-center justify-center text-white font-bold text-sm shadow-md">
              /100
            </div>
          </div>
          <div class="text-lg font-semibold text-gray-700 mt-4">نمره SEO</div>
          <div class="text-sm text-gray-500">عملکرد عالی</div>
        </div>

        <!-- چک‌لیست SEO -->
        <div class="border-t pt-4 border-gray-300">
          <h4 class="text-xs font-semibold text-gray-700 mb-3">چک‌لیست SEO</h4>
          <div class="space-y-2">
            <div class="flex items-center gap-2">
              <span class="text-green-500">✓</span>
              <span class="text-xs text-gray-700">عنوان SEO تعریف شده</span>
            </div>
            <div class="flex items-center gap-2">
              <span class="text-green-500">✓</span>
              <span class="text-xs text-gray-700">Meta Description تعریف شده</span>
            </div>
            <div class="flex items-center gap-2">
              <span class="text-yellow-500">⚠</span>
              <span class="text-xs text-gray-700">کلمات کلیدی کم است</span>
            </div>
            <div class="flex items-center gap-2">
              <span class="text-red-500">✗</span>
              <span class="text-xs text-gray-700">تصویر Alt Text ندارد</span>
            </div>
            <div class="flex items-center gap-2">
              <span class="text-green-500">✓</span>
              <span class="text-xs text-gray-700">URL سفارشی تعریف شده</span>
            </div>
          </div>
        </div>

        <!-- پیشنهادات بهبود -->
        <div class="border-t pt-4 border-gray-300">
          <h4 class="text-xs font-semibold text-gray-700 mb-3">پیشنهادات بهبود</h4>
          <div class="space-y-2">
            <div class="bg-yellow-50 p-2 rounded text-xs">
              <span class="font-semibold">کلمات کلیدی:</span> حداقل 5 کلمه کلیدی مرتبط اضافه کنید
            </div>
            <div class="bg-red-50 p-2 rounded text-xs">
              <span class="font-semibold">تصاویر:</span> متن جایگزین (Alt Text) برای تصاویر اضافه کنید
            </div>
            <div class="bg-gray-50 p-2 rounded text-xs">
              <span class="font-semibold">Schema:</span> اطلاعات برند و سازنده را تکمیل کنید
           
            </div>
          </div>
        </div>

        <!-- آمار عملکرد -->
        <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-6">
          <h4 class="text-lg font-semibold text-gray-800 mb-6 text-center">آمار عملکرد (شبیه‌سازی)</h4>
          <div class="grid grid-cols-2 md:grid-cols-4 gap-6">
            <div class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-xl p-6 text-center border border-blue-200">
              <div class="w-8 h-8 bg-blue-500 rounded-lg mx-auto mb-2 flex items-center justify-center">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                </svg>
              </div>
              <div class="text-2xl font-bold text-blue-700">1,250</div>
              <div class="text-xs text-blue-600 font-medium">نمایش در جستجو</div>
            </div>
            <div class="bg-gradient-to-br from-green-50 to-green-100 rounded-xl p-6 text-center border border-green-200">
              <div class="w-8 h-8 bg-green-500 rounded-lg mx-auto mb-2 flex items-center justify-center">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 15l-2 5L9 9l11 4-5 2zm0 0l5 5M7.188 2.239l.777 2.897M5.136 7.965l-2.898-.777M13.95 4.05l-2.122 2.122m-5.657 5.656l-2.12 2.122"></path>
                </svg>
              </div>
              <div class="text-2xl font-bold text-green-700">85</div>
              <div class="text-xs text-green-600 font-medium">کلیک از جستجو</div>
            </div>
            <div class="bg-gradient-to-br from-purple-50 to-purple-100 rounded-xl p-6 text-center border border-purple-200">
              <div class="w-8 h-8 bg-purple-500 rounded-lg mx-auto mb-2 flex items-center justify-center">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
                </svg>
              </div>
              <div class="text-2xl font-bold text-purple-700">6.8%</div>
              <div class="text-xs text-purple-600 font-medium">نرخ کلیک</div>
            </div>
            <div class="bg-gradient-to-br from-orange-50 to-orange-100 rounded-xl p-6 text-center border border-orange-200">
              <div class="w-8 h-8 bg-orange-500 rounded-lg mx-auto mb-2 flex items-center justify-center">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11.5V14m0-2.5v-6a1.5 1.5 0 113 0m-3 6a1.5 1.5 0 00-3 0v2a7.5 7.5 0 0015 0v-5a1.5 1.5 0 00-3 0m-6-3V11m0-5.5v-1a1.5 1.5 0 013 0v1m0 0V11m0-5.5T16.5 4l-3-3-3 3"></path>
                </svg>
              </div>
              <div class="text-2xl font-bold text-orange-700">15</div>
              <div class="text-xs text-orange-600 font-medium">رتبه متوسط</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="border-t border-gray-300 my-6"></div>
</template>

<style scoped>
.input,
textarea,
select {
  background-color: #ffffff;
}

/* Custom Select Arrow */
select {
  background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 20 20'%3e%3cpath stroke='%236b7280' stroke-linecap='round' stroke-linejoin='round' stroke-width='1.5' d='m6 8 4 4 4-4'/%3e%3c/svg%3e");
  background-position: right 0.5rem center;
  background-repeat: no-repeat;
  background-size: 1.5em 1.5em;
  padding-right: 2.5rem;
}

/* Hover transitions */

/* Line clamp utilities */
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
.transition-all {
  transition: all 0.3s ease;
}

/* Custom focus ring */
.focus\:ring-2:focus {
  --tw-ring-offset-shadow: var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color);
  --tw-ring-shadow: var(--tw-ring-inset) 0 0 0 calc(2px + var(--tw-ring-offset-width)) var(--tw-ring-color);
  box-shadow: var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000);
}
</style>
