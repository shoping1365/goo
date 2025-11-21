<script setup lang="ts">
// Interface for PostForm
interface PostForm {
  tags?: string[]
  tagsInput?: string
}

// Props
const props = defineProps<{ 
  slug?: string; 
  defaultTitle?: string; 
  enableOgType?: boolean; 
  metaTitle?: string; 
  ogTitle?: string;
  ogDescription?: string;
  ogType?: string;
  ogSiteName?: string;
  ogImage?: string; 
  robotsMeta?: string;
  postForm: PostForm;
}>()

const emit = defineEmits([
  'update:slug', 
  'select-og-image', 
  'update:metaTitle', 
  'update:ogTitle',
  'update:ogDescription',
  'add-tag',
  'remove-tag',
  'update:ogType',
  'update:ogSiteName',
  'update:ogImage',
  'update:robotsMeta',
  'update:postForm'
])

// Local reactive slug linked to input
import { computed, inject, isRef, onMounted, reactive, ref, watch } from 'vue';
import { useSlugManagement } from '~/composables/useSlugManagement';

const slug = ref(props.slug || '')
const seoTitle = ref(props.metaTitle ?? props.defaultTitle ?? '')

// متغیرهای جدید برای مدیریت slug
const slugTouched = ref(false)
const slugError = ref('')
const isCheckingSlug = ref(false)
const isGeneratingSlug = ref(false)

// استفاده از composable
const { checkSlugUnique: checkSlugUniqueAPI, generateUniqueSlug: generateUniqueSlugAPI, generateSlugFromTitle } = useSlugManagement()

// Try to get title provided by parent via provide('pageTitle')
const providedTitle = inject<string | Ref<string> | null>('pageTitle', null)
const seoTitleTouched = ref(false)

function getProvided(){ return providedTitle ? (isRef(providedTitle) ? providedTitle.value : providedTitle) : '' }

if (!props.defaultTitle && !seoTitle.value && providedTitle) {
  seoTitle.value = getProvided() || ''
}

if (providedTitle) {
  watch(() => getProvided(), v => {
    if (!seoTitleTouched.value) seoTitle.value = v || ''
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

// بررسی تکراری بودن slug هنگام تغییر
watch(slug, async (newSlug) => {
  if (newSlug && slugTouched.value) {
    await checkSlugUnique(newSlug)
  }
  emit('update:slug', newSlug)
})

// همگام‌سازی slug با prop
watch(() => props.slug, (val) => {
  if (val !== undefined && val !== slug.value) {
    slug.value = val || ''
  }
})

// Meta Description live length feedback ------------------------------------
const metaDescription = ref('')
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
// const slugLength = computed(()=> slug.value.length)



watch(() => props.slug, val => slug.value = val || '')
watch(() => props.defaultTitle, val => { if(!seoTitle.value) seoTitle.value = val || '' })
watch(slug, val => emit('update:slug', val))

// Section collapse/expand states for SEO tab
const sections = reactive({
  seoBasic: false,
  seoSocial: false,
  seoSchema: false,
  seoAnalysis: false,
  settings: false
})

const toggleSection = (section) => {
  sections[section] = !sections[section]
}

// ------ Canonical URL -----------------
const canonicalUrl = ref('')
const canonicalTouched = ref(false)

// تابع ساخت آدرس کانونیکال صحیح (آدرس مقاله، نه صفحه ویرایش)
function buildDefaultCanonical(){
  if (import.meta.client) {
    try {
      const url = new URL(window.location.href)
      // اگر در صفحه ویرایش هستیم، آدرس کانونیکال باید آدرس مقاله باشد
      if (url.pathname.includes('/edit-post/')) {
        // آدرس کانونیکال باید آدرس مقاله باشد، نه صفحه ویرایش
        return slug.value ? url.origin + '/blog/' + slug.value : url.origin
      }
      // اگر در صفحه افزودن هستیم، آدرس کانونیکال بر اساس slug باشد
      return slug.value ? url.origin + '/blog/' + slug.value : url.origin
    } catch {
      return slug.value ? '/blog/' + slug.value : ''
    }
  }
  // Fallback to relative slug path during SSR/build
  return slug.value ? '/blog/' + slug.value : ''
}

canonicalUrl.value = buildDefaultCanonical()

watch(canonicalUrl, () => { canonicalTouched.value = true })
// Update canonical when slug changes and user hasn't changed it manually
watch(slug, (val) => {
  if (!canonicalTouched.value) {
    if (import.meta.client) {
      try {
        const url = new URL(window.location.href)
        // آدرس کانونیکال همیشه باید آدرس مقاله باشد
        canonicalUrl.value = val ? url.origin + '/blog/' + val : url.origin
      } catch {
        canonicalUrl.value = val ? '/blog/' + val : canonicalUrl.value
      }
    } else {
      canonicalUrl.value = val ? '/blog/' + val : canonicalUrl.value
    }
  }
})

// ---------------- Open Graph fields -------------------
// فیلدهای Open Graph قابل ویرایش
const ogTitle = ref(props.ogTitle ?? '')
const ogDescription = ref(props.ogDescription ?? '')
const ogType = ref(props.ogType ?? 'article')
const ogSiteName = ref(props.ogSiteName ?? '')
// تصویر Open Graph (فقط این فیلد قابل ویرایش است)
const ogImage = ref(props.ogImage ?? '')

// متغیر برای نمایش در شبکه‌های اجتماعی (پیش‌فرض فعال)
// const showSocialMedia = ref(true)

// Interface for ShopSettings
interface ShopSettings {
  shopNameFa?: string
  shopNameEn?: string
}

// تنظیمات فروشگاه برای نام سایت
const shopSettings = ref<ShopSettings>({
  shopNameFa: 'فروشگاه من',
  shopNameEn: 'My Shop'
})

// دریافت تنظیمات فروشگاه
const fetchShopSettings = async () => {
  try {
    interface ApiResponse {
      data?: ShopSettings
    }
    const response = await $fetch<ApiResponse>('/api/admin/shop-settings')
    if (response.data) {
      shopSettings.value = {
        shopNameFa: response.data.shopNameFa || 'فروشگاه من',
        shopNameEn: response.data.shopNameEn || 'My Shop'
      }
    }
  } catch (error) {
    console.error('خطا در دریافت تنظیمات فروشگاه:', error)
  }
}

// Determine if OG type should be shown
// const showOgType = computed(() => props.enableOgType !== false)

// همگام‌سازی ogTitle داخلی با prop ogTitle
watch(() => props.ogTitle, (val) => {
  if (val !== undefined && val !== ogTitle.value) {
    ogTitle.value = val
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

// همگام‌سازی ogImage داخلی با prop ogImage
watch(() => props.ogImage, (val) => {
  if (val !== undefined && val !== ogImage.value) {
    ogImage.value = val
  }
})

// هر بار که ogTitle تغییر کرد، مقدار را به والد emit کن
watch(ogTitle, (val) => {
  emit('update:ogTitle', val)
})

// هر بار که ogDescription تغییر کرد، مقدار را به والد emit کن
watch(ogDescription, (val) => {
  emit('update:ogDescription', val)
})

// هر بار که ogType تغییر کرد، مقدار را به والد emit کن
watch(ogType, (val) => {
  emit('update:ogType', val)
})

// هر بار که ogSiteName تغییر کرد، مقدار را به والد emit کن
watch(ogSiteName, (val) => {
  emit('update:ogSiteName', val)
})

// هر بار که ogImage تغییر کرد، مقدار را به والد emit کن
watch(ogImage, (val) => {
  emit('update:ogImage', val)
})

// ---------------- Robots Meta -------------------
type RobotsMetaType = 'index,follow' | 'noindex,follow' | 'index,nofollow' | 'noindex,nofollow'
const robotsMeta = ref<RobotsMetaType>((props.robotsMeta as RobotsMetaType) || 'index,follow')

// Watch for prop changes
watch(() => props.robotsMeta, (newVal) => {
  if (newVal) robotsMeta.value = newVal as RobotsMetaType
})

// Emit changes
watch(robotsMeta, (newVal) => {
  emit('update:robotsMeta', newVal)
})

// اجرای توابع اولیه
onMounted(() => {
  fetchShopSettings()
})

defineExpose({
  seoTitle,
  metaDescription,
  robotsMeta,
  ogTitle,
  ogDescription,
  ogType,
  ogSiteName,
  ogImage
})

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

const schemaType = ref('Article')

// همگام‌سازی ogType با schemaType
watch(schemaType, (val) => {
  ogType.value = val
  emit('update:ogType', val)
})

// Computed for tags input to avoid prop mutation
const tagsInput = computed({
  get: () => props.postForm.tagsInput,
  set: (val) => emit('update:postForm', { ...props.postForm, tagsInput: val })
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
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- عنوان SEO (کوچک‌تر) -->
        <div class="bg-white rounded-xl shadow-sm border border-emerald-100 p-6">
          <label class="flex items-center gap-2 text-xs font-semibold text-gray-900 mb-3">
            <div class="p-2 bg-emerald-100 rounded-lg">
              <svg class="w-4 h-4 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
              </svg>
            </div>
            عنوان SEO (Title Tag)
          </label>
          <input v-model="seoTitle" type="text" class="w-full px-3 py-2 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500 text-gray-900 transition-all duration-200 text-sm" dir="rtl" placeholder="عنوان برای موتورهای جستجو" />
          <p class="text-xs text-gray-500 mt-1">می‌توانید عنوان SEO را به صورت دستی تغییر دهید یا اجازه دهید از عنوان مقاله گرفته شود</p>
          <!-- Progress bar below field -->
          <div class="flex items-center gap-2 mt-2">
            <div class="relative h-2 flex-1 bg-gray-200 rounded overflow-hidden">
              <div :style="{ width: Math.min(titleLength, 60)/60*100 + '%' }" :class="[titleColor,'absolute left-0 top-0 h-full transition-all duration-300'].join(' ')" ></div>
            </div>
            <span class="text-xs font-medium" :class="titleColor.replace('bg','text')">{{ titleLength }} / 60</span>
            <span class="text-xs px-2 py-1 rounded-full" :class="titleColor + ' text-white'">{{ titleQuality }}</span>
          </div>
        </div>

        <!-- توضیحات متا کنار عنوان -->
        <div class="bg-white rounded-xl shadow-sm border border-blue-100 p-6">
          <label class="flex items-center gap-2 text-xs font-semibold text-gray-900 mb-3">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h7"></path>
              </svg>
            </div>
            توضیحات متا (Meta Description)
          </label>
          <textarea v-model="metaDescription" class="w-full px-3 py-2 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-gray-900 transition-all duration-200 min-h-[80px] resize-none text-sm" dir="rtl" placeholder="توضیح کوتاه محصول برای موتورهای جستجو"></textarea>
          <!-- Progress bar below field -->
          <div class="flex items-center gap-2 mt-2">
            <div class="relative h-2 flex-1 bg-gray-200 rounded overflow-hidden">
              <div :style="{ width: Math.min(metaLength, 160)/160*100 + '%' }" :class="[metaColor,'absolute left-0 top-0 h-full transition-all duration-300'].join(' ')" ></div>
            </div>
            <span class="text-xs font-medium" :class="metaColor.replace('bg','text')">{{ metaLength }} / 160</span>
            <span class="text-xs px-2 py-1 rounded-full" :class="metaColor + ' text-white'">{{ metaQuality }}</span>
          </div>
        </div>

        <!-- پیش‌نمایش Google -->
        <div class="bg-white rounded-xl shadow-sm border border-green-100 p-6">
          <label class="flex items-center gap-2 text-xs font-semibold text-gray-900 mb-3">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-4 h-4 text-green-600" viewBox="0 0 24 24" fill="currentColor">
                <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/>
                <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/>
                <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"/>
                <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/>
              </svg>
            </div>
            پیش‌نمایش Google
          </label>
          <div class="border border-gray-300 rounded-lg p-3 bg-gray-50 min-h-[80px] flex flex-col justify-center">
            <div class="space-y-1">
              <div class="text-blue-600 text-xs truncate">{{ canonicalUrl }}</div>
              <div class="text-sm text-blue-800 font-medium line-clamp-2">{{ seoTitle || 'عنوان صفحه' }}</div>
              <div class="text-xs text-gray-600 line-clamp-2">{{ metaDescription || 'توضیحات صفحه' }}</div>
            </div>
          </div>
          <p class="text-xs text-gray-500 mt-2">نمایش چگونگی نمایش صفحه در نتایج جستجوی Google</p>
        </div>
      </div>

      <!-- فاصله بین بخش‌های مختلف -->
      <div class="mt-8"></div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6" style="column-gap: 40px;">
          <!-- URL سفارشی (Slug) -->
          <div class="bg-white rounded-xl shadow-sm border border-orange-100 p-6">
            <label class="flex items-center gap-2 text-xs font-semibold text-gray-900 mb-3">
              <div class="p-2 bg-orange-100 rounded-lg">
                <svg class="w-4 h-4 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"></path>
                </svg>
              </div>
              URL سفارشی (Slug)
            </label>
            <div class="space-y-2">
              <div class="relative">
                <input 
                  v-model="slug" 
                  type="text" 
                  :class="[
                    'w-full px-3 py-2 bg-white border rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-orange-500 text-gray-900 transition-all duration-200 text-sm',
                    slugError ? 'border-red-300 focus:ring-red-500 focus:border-red-500' : 'border-gray-300'
                  ]"
                  dir="ltr" 
                  placeholder="custom-article-url"
                  @blur="checkSlugUnique(slug)"
                />
                <!-- نمایش خطا -->
                <div v-if="slugError" class="mt-2 text-xs text-red-600 flex items-center gap-2">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                  </svg>
                  {{ slugError }}
                </div>
                <!-- نمایش موفقیت -->
                <div v-else-if="slug && !isCheckingSlug" class="mt-2 text-xs text-green-600 flex items-center gap-2">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                  </svg>
                  URL در دسترس است
                </div>
                <!-- نمایش loading -->
                <div v-if="isCheckingSlug" class="mt-2 text-xs text-blue-600 flex items-center gap-2">
                  <svg class="w-4 h-4 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
                  </svg>
                  در حال بررسی...
                </div>
              </div>
              <!-- دکمه‌های عملیات -->
              <div class="flex items-center gap-2">
                <button 
                  v-if="slugError && slug"
                  :disabled="isGeneratingSlug"
                  class="inline-flex items-center px-3 py-2 text-xs font-medium text-white bg-gradient-to-r from-blue-500 to-indigo-600 hover:from-blue-600 hover:to-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-all duration-200 rounded-lg disabled:opacity-50 disabled:cursor-not-allowed"
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
            <p class="text-xs text-gray-500 mt-2">
              <span v-if="!slug">URL الزامی است. URL به طور خودکار از عنوان مقاله تولید می‌شود.</span>
              <span v-else>این URL در آدرس صفحه مقاله نمایش داده می‌شود. پس از ویرایش دستی، URL دیگر تغییر نمی‌کند.</span>
            </p>
          </div>

          <!-- تنظیمات روبات‌ها -->
          <div class="bg-white rounded-xl shadow-sm border border-indigo-100 p-6">
            <label class="flex items-center gap-2 text-xs font-semibold text-gray-900 mb-3">
              <div class="p-2 bg-indigo-100 rounded-lg">
                <svg class="w-4 h-4 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                </svg>
              </div>
              تنظیمات روبات‌ها
            </label>
            <select v-model="robotsMeta" class="w-full px-3 py-2 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-900 transition-all duration-200 appearance-none text-sm">
              <option value="index,follow">نمایه‌سازی و پیگیری لینک‌ها</option>
              <option value="noindex,follow">عدم نمایه‌سازی، پیگیری لینک‌ها</option>
              <option value="index,nofollow">نمایه‌سازی، عدم پیگیری لینک‌ها</option>
              <option value="noindex,nofollow">عدم نمایه‌سازی و عدم پیگیری لینک‌ها</option>
            </select>
            <p class="text-xs text-gray-500 mt-2">تعیین می‌کند که موتورهای جستجو چگونه با این صفحه رفتار کنند</p>
          </div>
        </div>

        <!-- فاصله بین بخش‌های مختلف -->
        <div class="mt-8"></div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6" style="column-gap: 40px;">
          <!-- آدرس کانونیکال -->
          <div class="bg-white rounded-xl shadow-sm border border-rose-100 p-6">
            <label class="flex items-center gap-2 text-xs font-semibold text-gray-900 mb-3">
              <div class="p-2 bg-rose-100 rounded-lg">
                <svg class="w-4 h-4 text-rose-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                </svg>
              </div>
              آدرس کانونیکال
            </label>
            <input v-model="canonicalUrl" type="text" class="w-full px-3 py-2 bg-gray-100 border border-gray-300 rounded-lg text-gray-600 transition-all duration-200 text-sm" dir="ltr" placeholder="https://example.com/blog/article-slug" readonly />
            <p class="text-xs text-gray-500 mt-2">آدرس کانونیکال به طور خودکار بر اساس slug مقاله تولید می‌شود</p>
          </div>

          <!-- باکس تگ‌ها -->
          <div class="bg-white rounded-xl shadow-sm border border-blue-100 p-6">
            <div class="flex items-center gap-2 mb-3">
              <h4 class="text-xs font-bold text-gray-900">تگ‌ها</h4>
            </div>
            <div class="flex items-center gap-2 mb-2">
              <input 
                v-model="tagsInput"
                type="text"
                placeholder="تگ جدید..."
                class="flex-1 px-3 py-2 border border-gray-300 rounded-lg text-xs focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                dir="rtl"
                @keyup.enter="$emit('add-tag')"
              >
              <button 
                class="px-3 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors text-xs"
                @click="$emit('add-tag')"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
                </svg>
              </button>
            </div>
            <div v-if="props.postForm.tags && props.postForm.tags.length > 0" class="flex flex-wrap gap-2">
              <div 
                v-for="(tag, index) in props.postForm.tags" 
                :key="index"
                class="flex items-center gap-1 bg-blue-100 text-blue-800 px-2 py-1 rounded-lg text-xs"
              >
                <span>{{ tag }}</span>
                <button 
                  class="text-blue-600 hover:text-blue-800 transition-colors"
                  @click="$emit('remove-tag', index)"
                >
                  <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                  </svg>
                </button>
              </div>
            </div>
            <div v-else class="text-xs text-gray-500 text-center py-2">
              هیچ تگی اضافه نشده است
            </div>
          </div>
        </div>
      </div>
    </div>
  

  <!-- تنظیمات شبکه‌های اجتماعی -->
  <div class="bg-gradient-to-br from-pink-50 to-purple-50 rounded-2xl border border-pink-200 shadow-lg overflow-hidden mb-8">
    <div class="bg-gradient-to-r from-pink-600 to-purple-600 px-6 py-4 cursor-pointer" @click="toggleSection('seoSocial')">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 rounded-xl">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 4V2a1 1 0 011-1h8a1 1 0 011 1v2m-9 0h10m-10 0a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V6a2 2 0 00-2-2"></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold text-white">تنظیمات شبکه‌های اجتماعی</h3>
        </div>
        <div class="p-2 bg-white/20 rounded-lg">
          <span class="text-white font-bold text-lg">{{ sections.seoSocial ? '−' : '+' }}</span>
        </div>
      </div>
    </div>

    <div v-show="sections.seoSocial" class="p-8">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mt-6">
        <!-- عنوان Open Graph -->
        <div class="bg-white rounded-xl shadow-sm border border-pink-100 p-6">
          <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
            <div class="p-2 bg-pink-100 rounded-lg">
              <svg class="w-4 h-4 text-pink-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
              </svg>
            </div>
            عنوان Open Graph
          </label>
          <input v-model="ogTitle" type="text" class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-pink-500 focus:border-pink-500 text-gray-900 transition-all duration-200" dir="rtl" placeholder="عنوان برای شبکه‌های اجتماعی" />
          <p class="text-xs text-gray-500 mt-2">عنوانی که در شبکه‌های اجتماعی نمایش داده می‌شود</p>
        </div>
        <!-- توضیحات Open Graph -->
        <div class="bg-white rounded-xl shadow-sm border border-purple-100 p-6">
          <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h7"></path>
              </svg>
            </div>
            توضیحات Open Graph
          </label>
          <textarea v-model="ogDescription" class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 text-gray-900 transition-all duration-200 min-h-[80px] resize-none" dir="rtl" placeholder="توضیحات برای شبکه‌های اجتماعی"></textarea>
          <p class="text-xs text-gray-500 mt-2">توضیحی که در شبکه‌های اجتماعی نمایش داده می‌شود</p>
        </div>
      </div>

        <!-- پیش‌نمایش شبکه‌های اجتماعی -->
        <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 mt-6">
          
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
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 012 2z"></path>
                      </svg>
                      <p class="text-sm">تصویر Open Graph</p>
                    </div>
                  </div>
                  <!-- محتوا -->
                  <div class="p-3">
                    <h3 class="text-sm font-semibold text-gray-900 mb-2 line-clamp-2">
                      {{ ogTitle || seoTitle || 'عنوان مقاله' }}
                    </h3>
                    <p class="text-xs text-gray-600 line-clamp-3">
                      {{ ogDescription || metaDescription || 'توضیحات مقاله برای نمایش در شبکه‌های اجتماعی' }}
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
                    <h3 class="text-sm font-semibold text-gray-900 mb-2 line-clamp-2">
                      {{ ogTitle || seoTitle || 'عنوان مقاله' }}
                    </h3>
                    <p class="text-xs text-gray-600 line-clamp-3">
                      {{ ogDescription || metaDescription || 'توضیحات مقاله برای نمایش در شبکه‌های اجتماعی' }}
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
          
        </div>
      </div>
    </div>
  

  <!-- تنظیمات Schema -->
  <div class="bg-gradient-to-br from-yellow-50 to-orange-50 rounded-2xl border border-yellow-200 shadow-lg overflow-hidden mb-8">
    <div class="bg-gradient-to-r from-yellow-600 to-orange-600 px-6 py-4 cursor-pointer" @click="toggleSection('seoSchema')">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 rounded-xl">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold text-white">تنظیمات Schema</h3>
        </div>
        <div class="p-2 bg-white/20 rounded-lg">
          <span class="text-white font-bold text-lg">{{ sections.seoSchema ? '−' : '+' }}</span>
        </div>
      </div>
    </div>

    <div v-show="sections.seoSchema" class="p-8">
      <!-- انتخاب نوع اسکیما -->
      <div class="bg-white rounded-xl shadow-sm border border-purple-100 p-6 mb-6">
        <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zM7 3H5a2 2 0 00-2 2v12a4 4 0 004 4h2a2 2 0 002-2V5a2 2 0 00-2-2H7z"></path>
            </svg>
          </div>
          نوع Schema
        </label>
        <select v-model="schemaType" class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 text-gray-900 transition-all duration-200 appearance-none">
          <option value="Article">مقاله (Article)</option>
          <option value="Product">محصول (Product)</option>
          <option value="Book">کتاب (Book)</option>
          <option value="SoftwareApplication">نرم‌افزار (Software)</option>
          <option value="CreativeWork">اثر خلاقانه (Creative Work)</option>
          <option value="Organization">سازمان (Organization)</option>
          <option value="WebPage">صفحه وب (WebPage)</option>
          <option value="BreadcrumbList">مسیر ناوبری (BreadcrumbList)</option>
          <option value="FAQPage">سوالات متداول (FAQPage)</option>
          <option value="LocalBusiness">کسب و کار محلی (LocalBusiness)</option>
          <option value="Review">نظر (Review)</option>
        </select>
      </div>
      
      <!-- تقسیم به دو قسمت: کدهای اسکیما و خود اسکیما -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- سمت راست: کدهای اسکیما -->
        <div class="bg-white rounded-xl shadow-sm border border-blue-100 p-6">
          <div class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"></path>
              </svg>
            </div>
            کدهای Schema
          </div>
          <div class="bg-gray-900 rounded-lg p-6 overflow-x-auto">
            <pre class="text-green-400 text-xs leading-relaxed"><code>{
  "@context": "https://schema.org",
  "@type": "{{ schemaType }}",
  "headline": "{{ seoTitle || 'عنوان مقاله' }}",
  "description": "{{ metaDescription || 'توضیحات مقاله' }}",
  "author": {
    "@type": "Person",
    "name": "نویسنده"
  },
  "publisher": {
    "@type": "Organization",
    "name": "{{ shopSettings.shopNameFa }}"
  },
  "datePublished": "{{ new Date().toISOString() }}",
  "dateModified": "{{ new Date().toISOString() }}",
  "mainEntityOfPage": {
    "@type": "WebPage",
    "@id": "{{ canonicalUrl }}"
  }
}</code></pre>
          </div>
          <p class="text-xs text-gray-500 mt-3">کد Schema Markup که به طور خودکار تولید می‌شود</p>
        </div>

        <!-- سمت چپ: خود اسکیما -->
        <div class="bg-white rounded-xl shadow-sm border border-green-100 p-6">
          <div class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            Schema Markup
          </div>
          <div class="bg-gray-50 rounded-lg p-6">
            <div class="flex items-center gap-2 text-sm text-gray-700 mb-3">
              <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
              </svg>
              Schema Markup فعال است
            </div>
            <div class="space-y-2 text-sm text-gray-600">
              <div class="flex items-center gap-2">
                <span class="font-medium">نوع:</span>
                <span class="bg-blue-100 text-blue-800 px-2 py-1 rounded text-xs">{{ schemaType }}</span>
              </div>
              <div class="flex items-center gap-2">
                <span class="font-medium">عنوان:</span>
                <span class="text-gray-800">{{ seoTitle || 'عنوان مقاله' }}</span>
              </div>
              <div class="flex items-center gap-2">
                <span class="font-medium">توضیحات:</span>
                <span class="text-gray-800">{{ metaDescription || 'توضیحات مقاله' }}</span>
              </div>
              <div class="flex items-center gap-2">
                <span class="font-medium">نویسنده:</span>
                <span class="text-gray-800">نویسنده</span>
              </div>
              <div class="flex items-center gap-2">
                <span class="font-medium">ناشر:</span>
                <span class="text-gray-800">{{ shopSettings.shopNameFa }}</span>
              </div>
            </div>
          </div>
          <p class="text-xs text-gray-500 mt-3">اطلاعات Schema که در موتورهای جستجو نمایش داده می‌شود</p>
        </div>
      </div>
    </div>
  </div>

  <!-- تحلیل SEO -->
  <div class="bg-gradient-to-br from-red-50 to-pink-50 rounded-2xl border border-red-200 shadow-lg overflow-hidden mb-8">
    <div class="bg-gradient-to-r from-red-600 to-pink-600 px-6 py-4 cursor-pointer" @click="toggleSection('seoAnalysis')">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 rounded-xl">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold text-white">تحلیل SEO</h3>
        </div>
        <div class="p-2 bg-white/20 rounded-lg">
          <span class="text-white font-bold text-lg">{{ sections.seoAnalysis ? '−' : '+' }}</span>
        </div>
      </div>
    </div>

    <div v-show="sections.seoAnalysis" class="p-8">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <!-- امتیاز کلی -->
        <div class="bg-white rounded-xl shadow-sm border border-red-100 p-6">
          <div class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
            <div class="p-2 bg-red-100 rounded-lg">
              <svg class="w-4 h-4 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"></path>
              </svg>
            </div>
            امتیاز کلی SEO
          </div>
          <div class="text-center">
            <div class="text-4xl font-bold text-green-600 mb-2">85</div>
            <div class="text-sm text-gray-600">از 100</div>
          </div>
        </div>

        <!-- وضعیت فیلدها -->
        <div class="bg-white rounded-xl shadow-sm border border-orange-100 p-6">
          <div class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
            <div class="p-2 bg-orange-100 rounded-lg">
              <svg class="w-4 h-4 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
              </svg>
            </div>
            وضعیت فیلدها
          </div>
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">عنوان SEO</span>
              <span class="text-sm font-medium text-green-600">تکمیل شده</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">توضیحات متا</span>
              <span class="text-sm font-medium text-green-600">تکمیل شده</span>
            </div>

            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">URL سفارشی</span>
              <span class="text-sm font-medium text-green-600">تکمیل شده</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
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
.transition-all {
  transition: all 0.3s ease;
}

/* Custom focus ring */
.focus\:ring-2:focus {
  --tw-ring-offset-shadow: var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color);
  --tw-ring-shadow: var(--tw-ring-inset) 0 0 0 calc(2px + var(--tw-ring-offset-width)) var(--tw-ring-color);
  box-shadow: var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000);
}

/* Line clamp utilities */
.line-clamp-2 {
  display: -webkit-box;
  line-clamp: 2;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.line-clamp-3 {
  display: -webkit-box;
  line-clamp: 3;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
