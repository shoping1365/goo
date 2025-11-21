<template>
  <div class="edit-mobile-app-header-page">
    <div class="header-bg">
      <div class="page-header-flex">
        <div>
          <h1 class="page-title">ویرایش هدر موبایل و اپلیکیشن</h1>
          <p class="page-subtitle">ویرایش هدر موبایل و اپلیکیشن: {{ mobileAppHeader?.name }}</p>
        </div>
        <button class="btn btn-secondary" @click="goBack">بازگشت</button>
      </div>
    </div>

    <div v-if="loading" class="loading-state">
      <p>در حال بارگذاری...</p>
    </div>

    <div v-else-if="error" class="error-state">
      <p>{{ error }}</p>
      <button class="btn btn-primary" @click="loadMobileAppHeader">تلاش مجدد</button>
    </div>

    <div v-else>
      <!-- نمونه‌های آماده -->
      <div class="templates-section">
        <h2>انتخاب از نمونه‌های آماده</h2>
        <div class="templates-grid">
          <div class="template-card" :class="{ active: selectedTemplate === 'template1' }" @click="selectTemplate('template1')">
            <div class="template-preview">
              <div class="mobile-header-preview">
                <!-- نمونه اول: لوگو + آیکون ماه و حساب کاربری + جستجو + سبد خرید -->
                <div class="preview-header">
                  <div class="preview-top-bar">
                    <div class="preview-logo">لوگو</div>
                    <div class="preview-icons">
                      <div class="preview-icon moon-icon">🌙</div>
                      <div class="preview-icon user-icon">👤</div>
                    </div>
                  </div>
                  <div class="preview-search-section">
                    <div class="preview-search-bar">
                      <div class="preview-search-icon">🔍</div>
                      <span class="preview-search-text">نام سایت</span>
                    </div>
                    <div class="preview-cart-icon">🛒</div>
                  </div>
                </div>
              </div>
            </div>
            <div class="template-info">
              <h4>نمونه کلاسیک</h4>
              <p>لوگو + آیکون ماه و حساب کاربری + جستجو + سبد خرید</p>
            </div>
          </div>

          <div class="template-card" :class="{ active: selectedTemplate === 'template2' }" @click="selectTemplate('template2')">
            <div class="template-preview">
              <div class="mobile-header-preview">
                <!-- نمونه دوم: جستجو با لوگو + نوتیفیکیشن -->
                <div class="preview-header">
                  <div class="preview-search-section-2">
                    <div class="preview-search-bar-2">
                      <span class="preview-search-text-2">نام سایت</span>
                      <div class="preview-logo-left-2">لوگو</div>
                    </div>
                    <div class="preview-notification-icon">🔔</div>
                  </div>
                </div>
              </div>
            </div>
            <div class="template-info">
              <h4>نمونه با بنر</h4>
              <p>جستجو با لوگو + نوتیفیکیشن</p>
            </div>
          </div>

          <div class="template-card" :class="{ active: selectedTemplate === 'template3' }" @click="selectTemplate('template3')">
            <div class="template-preview">
              <div class="mobile-header-preview">
                <!-- نمونه سوم: باکس جستجو با لوگو در سمت چپ -->
                <div class="preview-header">
                  <div class="preview-search-section-3">
                    <div class="preview-search-bar-3">
                      <div class="preview-search-icon-3">🔍</div>
                      <span class="preview-search-text-3">نام سایت</span>
                      <div class="preview-logo-left">
                        <span class="preview-logo-text">لوگو</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="template-info">
              <h4>نمونه مینیمال</h4>
              <p>باکس جستجو با لوگو در سمت راست</p>
            </div>
          </div>
        </div>
      </div>

      <div class="form-container">
        <div class="form-layout">
        <form class="mobile-app-header-form" @submit.prevent="updateMobileAppHeader">
          <!-- اطلاعات اصلی -->
          <div class="form-section">
            <h3>اطلاعات اصلی</h3>
            <div class="form-grid">
              <div class="form-group">
                <label for="name">نام هدر *</label>
                <input 
                  id="name"
                  v-model="formData.name" 
                  type="text" 
                  required
                  placeholder="نام هدر موبایل و اپلیکیشن"
                  class="form-input"
                >
              </div>
              <div class="form-group">
                <label for="platform">پلتفرم *</label>
                <select id="platform" v-model="formData.platform" required class="form-select">
                  <option value="mobile">موبایل</option>
                  <option value="app">اپلیکیشن</option>
                  <option value="both">هر دو</option>
                </select>
              </div>
            </div>
            <div class="form-group">
              <label for="description">توضیحات</label>
              <textarea 
                id="description"
                v-model="formData.description" 
                rows="3"
                placeholder="توضیحات هدر موبایل و اپلیکیشن"
                class="form-textarea"
              ></textarea>
            </div>
            
            <!-- آپلود عکس هدر -->
            <div class="form-group">
              <label>عکس هدر</label>
              <div class="image-upload-section">
                <!-- نمایش عکس‌های موجود -->
                <div v-if="formData.top_image_url || formData.bottom_image_url" class="existing-images">
                  <div v-if="formData.top_image_url" class="image-preview">
                    <span class="image-label">عکس بالای هدر:</span>
                    <img :src="formData.top_image_url" :alt="formData.top_image_alt || 'عکس بالای هدر'" class="preview-image" />
                    <button type="button" class="remove-image-btn" @click="removeTopImage">حذف</button>
                  </div>
                  <div v-if="formData.bottom_image_url" class="image-preview">
                    <span class="image-label">عکس پایین هدر:</span>
                    <img :src="formData.bottom_image_url" :alt="formData.bottom_image_alt || 'عکس پایین هدر'" class="preview-image" />
                    <button type="button" class="remove-image-btn" @click="removeBottomImage">حذف</button>
                  </div>
                </div>
                
                <!-- دکمه آپلود کوچک -->
                <div v-if="!showPositionOptions" class="small-upload-section">
                  <button type="button" class="small-upload-btn" @click="showMediaModal = true">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                    </svg>
                    آپلود عکس جدید
                  </button>
                </div>
                
                <!-- پیش‌نمایش و گزینه‌های موقعیت بعد از آپلود -->
                <div v-if="showPositionOptions && uploadedImage" class="uploaded-image-section">
                  <!-- پیش‌نمایش عکس آپلود شده -->
                  <div class="uploaded-image-preview">
                    <img :src="uploadedImage.url" :alt="uploadedImage.alt || 'عکس آپلود شده'" class="uploaded-preview-image" />
                    <div class="uploaded-image-info">
                      <span class="uploaded-image-name">{{ uploadedImage.name || 'عکس جدید' }}</span>
                      <span class="uploaded-image-size">{{ formatFileSize(uploadedImage.size) }}</span>
                    </div>
                  </div>
                  
                  <!-- گزینه‌های موقعیت -->
                  <div class="position-options">
                    <h4>عکس را در کجا قرار دهید؟</h4>
                    <div class="position-buttons">
                      <button type="button" class="position-btn" @click="setImagePosition('top')">
                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18" />
                        </svg>
                        بالای هدر
                      </button>
                      <button type="button" class="position-btn" @click="setImagePosition('bottom')">
                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3" />
                        </svg>
                        پایین هدر
                      </button>
                    </div>
                    <button type="button" class="cancel-btn" @click="cancelImageUpload">لغو</button>
                    
                    <!-- Debug: نمایش وضعیت -->
                    <div style="background: green; padding: 4px; font-size: 10px; color: white; margin-top: 8px;">
                      Debug: uploadedImage = {{ uploadedImage ? 'موجود' : 'موجود نیست' }}, showPositionOptions = {{ showPositionOptions }}
                    </div>
                  </div>
                </div>
              </div>
              
              <!-- فیلدهای متن جایگزین -->
              <div v-if="formData.top_image_url" class="form-group">
                <label for="top_image_alt">متن جایگزین عکس بالای هدر</label>
                <input 
                  id="top_image_alt"
                  v-model="formData.top_image_alt" 
                  type="text"
                  placeholder="متن جایگزین عکس بالای هدر"
                  class="form-input"
                >
              </div>
              <div v-if="formData.bottom_image_url" class="form-group">
                <label for="bottom_image_alt">متن جایگزین عکس پایین هدر</label>
                <input 
                  id="bottom_image_alt"
                  v-model="formData.bottom_image_alt" 
                  type="text"
                  placeholder="متن جایگزین عکس پایین هدر"
                  class="form-input"
                >
              </div>
            </div>
            
            <!-- تنظیمات رنگ -->
            <div class="form-section">
              <h3>تنظیمات رنگ</h3>
              <div class="form-grid">
                <div class="form-group">
                  <label for="background_color">رنگ پس‌زمینه</label>
                  <div class="color-input-group">
                    <input 
                      id="background_color"
                      v-model="formData.background_color" 
                      type="color"
                      class="color-picker"
                    >
                    <input 
                      v-model="formData.background_color" 
                      type="text"
                      placeholder="#ffffff"
                      class="form-input color-text-input"
                    >
                  </div>
                </div>
                <div class="form-group">
                  <label for="text_color">رنگ متن</label>
                  <div class="color-input-group">
                    <input 
                      id="text_color"
                      v-model="formData.text_color" 
                      type="color"
                      class="color-picker"
                    >
                    <input 
                      v-model="formData.text_color" 
                      type="text"
                      placeholder="#000000"
                      class="form-input color-text-input"
                    >
                  </div>
                </div>
              </div>
            </div>
            
            <!-- تنظیمات نمایش -->
            <div class="form-group">
              <label for="pageSelection">انتخاب صفحات</label>
              <select id="pageSelection" v-model="formData.pageSelection" class="form-select">
                <option value="all">همه صفحات</option>
                <option value="specific">صفحات خاص</option>
                <option value="exclude">مستثنی کردن صفحات</option>
              </select>
            </div>
            <div v-if="formData.pageSelection === 'specific'" class="form-group">
              <label for="specificPages">صفحات خاص</label>
              <textarea 
                id="specificPages"
                v-model="formData.specificPages" 
                rows="2"
                placeholder="آدرس صفحات مورد نظر (هر خط یک آدرس)"
                class="form-textarea"
              ></textarea>
            </div>
            <div v-if="formData.pageSelection === 'exclude'" class="form-group">
              <label for="excludedPages">صفحات مستثنی</label>
              <textarea 
                id="excludedPages"
                v-model="formData.excludedPages" 
                rows="2"
                placeholder="آدرس صفحات مستثنی (هر خط یک آدرس)"
                class="form-textarea"
              ></textarea>
            </div>
            
            <div class="form-group">
              <label class="checkbox-label">
                <input v-model="formData.isActive" type="checkbox">
                <span>فعال</span>
              </label>
            </div>
            
            <!-- دکمه‌های عملیات -->
            <div class="form-actions">
              <button type="button" class="btn btn-secondary" @click="goBack">لغو</button>
              <button type="submit" class="btn btn-primary" :disabled="updating">
                <span v-if="updating">در حال به‌روزرسانی...</span>
                <span v-else>به‌روزرسانی هدر موبایل و اپلیکیشن</span>
              </button>
            </div>
          </div>
        </form>

        <!-- پیش‌نمایش زنده -->
        <div class="live-preview-section">
          <h3>پیش‌نمایش زنده</h3>
          <div class="live-preview-container">
            <div class="live-preview-mobile">
              <div class="live-preview-screen">
                <!-- پیش‌نمایش نمونه کلاسیک -->
                <div v-if="selectedTemplate === 'template1'" class="live-preview-header" :style="{ backgroundColor: formData.background_color, color: formData.text_color }">
                  <!-- عکس بالای هدر -->
                  <div v-if="formData.top_image_url" class="preview-top-image">
                    <img :src="formData.top_image_url" :alt="formData.top_image_alt || 'عکس بالای هدر'" class="preview-image" />
                  </div>
                  
                  <ClassicTemplate
:header="{ 
                    show_logo: true, 
                    show_search: true, 
                    show_cart: true, 
                    show_user: true, 
                    show_moon: true,
                    logo_text: formData.name || 'نام هدر',
                    search_placeholder: 'جستجو...'
                  }" />
                  
                  <!-- عکس پایین هدر -->
                  <div v-if="formData.bottom_image_url" class="preview-bottom-image">
                    <img :src="formData.bottom_image_url" :alt="formData.bottom_image_alt || 'عکس پایین هدر'" class="preview-image" />
                  </div>
                </div>

                <!-- پیش‌نمایش نمونه با بنر -->
                <div v-else-if="selectedTemplate === 'template2'" class="live-preview-header" :style="{ backgroundColor: formData.background_color, color: formData.text_color }">
                  <!-- عکس بالای هدر -->
                  <div v-if="formData.top_image_url" class="preview-top-image">
                    <img :src="formData.top_image_url" :alt="formData.top_image_alt || 'عکس بالای هدر'" class="preview-image" />
                  </div>
                  
                  <BannerTemplate
:header="{ 
                    show_logo: true, 
                    show_search: true, 
                    show_cart: true, 
                    show_user: true, 
                    show_moon: true,
                    logo_text: formData.name || 'نام هدر',
                    search_placeholder: 'جستجو...',
                    banner_text: formData.description || 'توضیحات هدر'
                  }" />
                  
                  <!-- عکس پایین هدر -->
                  <div v-if="formData.bottom_image_url" class="preview-bottom-image">
                    <img :src="formData.bottom_image_url" :alt="formData.bottom_image_alt || 'عکس پایین هدر'" class="preview-image" />
                  </div>
                </div>

                <!-- پیش‌نمایش نمونه مینیمال -->
                <div v-else-if="selectedTemplate === 'template3'" class="live-preview-header" :style="{ backgroundColor: formData.background_color, color: formData.text_color }">
                  <!-- عکس بالای هدر -->
                  <div v-if="formData.top_image_url" class="preview-top-image">
                    <img :src="formData.top_image_url" :alt="formData.top_image_alt || 'عکس بالای هدر'" class="preview-image" />
                  </div>
                  
                  <MinimalTemplate
:header="{ 
                    show_logo: true, 
                    show_search: true, 
                    show_cart: true, 
                    show_user: true, 
                    show_moon: true,
                    logo_text: formData.name || 'نام هدر',
                    search_placeholder: 'جستجو...'
                  }" />
                  
                  <!-- عکس پایین هدر -->
                  <div v-if="formData.bottom_image_url" class="preview-bottom-image">
                    <img :src="formData.bottom_image_url" :alt="formData.bottom_image_alt || 'عکس پایین هدر'" class="preview-image" />
                  </div>
                </div>

                <!-- پیش‌نمایش پیش‌فرض -->
                <div v-else class="live-preview-placeholder">
                  <p>یک نمونه را انتخاب کنید تا پیش‌نمایش آن را ببینید</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    </div>
  </div>
  
  <!-- مودال آپلود رسانه -->
  <MediaLibraryModal 
    v-model="showMediaModal" 
    file-type="image"
    default-category="banners"
    @select="handleImageSelect"
  />
</template>

<script setup lang="ts">
import { useRoute, useRouter } from 'nuxt/app'
import { onMounted, ref } from 'vue'
import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue'
import BannerTemplate from '../templates/banner-template.vue'
import ClassicTemplate from '../templates/classic-template.vue'
import MinimalTemplate from '../templates/minimal-template.vue'

// @ts-ignore
definePageMeta({
  layout: 'admin-main'
})

const router = useRouter()
const route = useRoute()

// State
const loading = ref(false)
const updating = ref(false)
const error = ref('')
const showMediaModal = ref(false)
const uploadedImage = ref(null) // عکس آپلود شده
const showPositionOptions = ref(false) // نمایش گزینه‌های موقعیت
const mobileAppHeader = ref(null)
const selectedTemplate = ref('')
const formData = ref({
  name: '',
  description: '',
  platform: 'mobile',
  background_color: '#ffffff',
  text_color: '#000000',
  pageSelection: 'all',
  specificPages: '',
  excludedPages: '',
  isActive: true,
  top_image_url: '',
  top_image_alt: '',
  bottom_image_url: '',
  bottom_image_alt: ''
})

// Methods
const goBack = () => {
  router.push('/admin/content/mobile-app-header-management')
}

const loadMobileAppHeader = async () => {
  loading.value = true
  error.value = ''
  
  try {
    const response = await $fetch<Record<string, unknown>>(`/api/admin/mobile-app-header-settings/${route.params.id}`)
    
    const data = (response.data || response) as Record<string, unknown> // Handle both response formats
    
    mobileAppHeader.value = data
    formData.value = {
      name: data.name || '',
      description: data.description || '',
      platform: data.platform || 'mobile',
      background_color: data.background_color || '#ffffff',
      text_color: data.text_color || '#000000',
      pageSelection: data.pageSelection || 'all',
      specificPages: data.specificPages || '',
      excludedPages: data.excludedPages || '',
      isActive: data.isActive !== false,
      top_image_url: data.top_image_url || '',
      top_image_alt: data.top_image_alt || '',
      bottom_image_url: data.bottom_image_url || '',
      bottom_image_alt: data.bottom_image_alt || ''
    }
    
    // اگر فیلدهای عکس وجود ندارند، آن‌ها را با مقادیر پیش‌فرض اضافه کن
    if (!data.hasOwnProperty('top_image_url')) {
      formData.value.top_image_url = ''
      formData.value.top_image_alt = ''
    }
    if (!data.hasOwnProperty('bottom_image_url')) {
      formData.value.bottom_image_url = ''
      formData.value.bottom_image_alt = ''
    }
    
    // تشخیص نمونه بر اساس نام و توضیحات
    detectTemplate(data)
  } catch (err: unknown) {
    const message = (err as { data?: { message?: string } })?.data?.message || 'خطا در بارگذاری هدر موبایل و اپلیکیشن'
    error.value = message
    console.error('Error loading mobile app header:', err)
  } finally {
    loading.value = false
  }
}

const detectTemplate = (data: Record<string, unknown>) => {
  const name = (data.name as string)?.toLowerCase() || ''
  const description = (data.description as string)?.toLowerCase() || ''
  
  if (name.includes('کلاسیک') || description.includes('کلاسیک')) {
    selectedTemplate.value = 'template1'
  } else if (name.includes('بنر') || description.includes('بنر')) {
    selectedTemplate.value = 'template2'
  } else if (name.includes('مینیمال') || description.includes('مینیمال')) {
    selectedTemplate.value = 'template3'
  } else {
    // پیش‌فرض: نمونه کلاسیک
    selectedTemplate.value = 'template1'
  }
}

const selectTemplate = (templateId: string) => {
  selectedTemplate.value = templateId
  
  if (templateId === 'template1') {
    formData.value.name = 'هدر کلاسیک موبایل'
    formData.value.description = 'هدر کلاسیک با لوگو، آیکون ماه و حساب کاربری، جستجو و سبد خرید'
    formData.value.platform = 'mobile'
  } else if (templateId === 'template2') {
    formData.value.name = 'هدر با بنر تبلیغاتی'
    formData.value.description = 'هدر با بنر تبلیغاتی، جستجو با لوگو و آیکون نوتیفیکیشن'
    formData.value.platform = 'mobile'
  } else if (templateId === 'template3') {
    formData.value.name = 'هدر مینیمال موبایل'
    formData.value.description = 'هدر مینیمال با باکس جستجو و لوگو در سمت چپ'
    formData.value.platform = 'mobile'
  }
}

// مدیریت انتخاب تصویر از مودال
const handleImageSelect = (image: { url: string; alt?: string }) => {
  uploadedImage.value = image
  showMediaModal.value = false
  showPositionOptions.value = true
}

// تنظیم موقعیت عکس
const setImagePosition = (position: string) => {
  if (!uploadedImage.value) return
  
  if (position === 'top') {
    formData.value.top_image_url = uploadedImage.value.url
    formData.value.top_image_alt = uploadedImage.value.alt || 'عکس بالای هدر'
  } else if (position === 'bottom') {
    formData.value.bottom_image_url = uploadedImage.value.url
    formData.value.bottom_image_alt = uploadedImage.value.alt || 'عکس پایین هدر'
  }
  
  // ریست کردن حالت‌ها
  uploadedImage.value = null
  showPositionOptions.value = false
}

// لغو آپلود
const cancelImageUpload = () => {
  uploadedImage.value = null
  showPositionOptions.value = false
}

// حذف عکس بالای هدر
const removeTopImage = () => {
  formData.value.top_image_url = ''
  formData.value.top_image_alt = ''
}

// حذف عکس پایین هدر
const removeBottomImage = () => {
  formData.value.bottom_image_url = ''
  formData.value.bottom_image_alt = ''
}

// فرمت کردن اندازه فایل
const formatFileSize = (bytes: number) => {
  if (!bytes) return ''
  if (bytes < 1024) return bytes + ' B'
  if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(1) + ' KB'
  return (bytes / (1024 * 1024)).toFixed(1) + ' MB'
}

const updateMobileAppHeader = async () => {
  updating.value = true
  
  try {
    await $fetch(`/api/admin/mobile-app-header-settings/${route.params.id}`, {
      method: 'PUT',
      body: formData.value
    })
    
    alert('هدر موبایل و اپلیکیشن با موفقیت به‌روزرسانی شد')
    router.push('/admin/content/mobile-app-header-management')
  } catch (err: unknown) {
    const message = (err as { data?: { message?: string } })?.data?.message || 'خطا در به‌روزرسانی هدر موبایل و اپلیکیشن'
    alert(message)
    console.error('Error updating mobile app header:', err)
  } finally {
    updating.value = false
  }
}

// Lifecycle
onMounted(() => {
  loadMobileAppHeader()
})
</script>

<style scoped>
.edit-mobile-app-header-page {
  padding: 20px;
}

.header-bg {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 30px;
  border-radius: 12px;
  margin-bottom: 30px;
}

.page-header-flex {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.page-title {
  font-size: 2rem;
  font-weight: bold;
  margin: 0 0 8px 0;
}

.page-subtitle {
  margin: 0;
  opacity: 0.9;
  font-size: 1rem;
}

.loading-state, .error-state {
  text-align: center;
  padding: 40px;
  color: #6b7280;
}

.error-state {
  color: #dc2626;
}

.form-container {
  padding: 30px;
}

.mobile-app-header-form {
  max-width: none;
  background: white;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  height: 100%;
  min-height: 700px;
}

.form-section {
  margin-bottom: 40px;
}

.form-section h3 {
  font-size: 1.25rem;
  color: #374151;
  margin-bottom: 20px;
  padding-bottom: 10px;
  border-bottom: 2px solid #e5e7eb;
}

.form-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 6px;
  font-weight: 600;
  color: #374151;
}

.form-input, .form-select, .form-textarea {
  width: 100%;
  padding: 12px 16px;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  font-size: 1rem;
  transition: border-color 0.2s ease;
}

.form-input:focus, .form-select:focus, .form-textarea:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.form-textarea {
  resize: vertical;
  min-height: 80px;
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
}

.checkbox-label input[type="checkbox"] {
  width: auto;
  margin: 0;
}

.color-input-group {
  display: flex;
  gap: 8px;
  align-items: center;
}

.color-picker {
  width: 50px;
  height: 50px;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  cursor: pointer;
  padding: 0;
}

.color-text-input {
  flex: 1;
}

.form-actions {
  display: flex;
  gap: 16px;
  justify-content: flex-end;
  padding-top: 20px;
  border-top: 1px solid #e5e7eb;
}

.btn {
  padding: 12px 24px;
  border-radius: 8px;
  font-weight: 600;
  transition: all 0.2s ease;
  border: none;
  cursor: pointer;
  font-size: 1rem;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn-primary {
  background: #3b82f6;
  color: white;
}

.btn-primary:hover:not(:disabled) {
  background: #2563eb;
}

.btn-secondary {
  background: #6b7280;
  color: white;
}

.btn-secondary:hover {
  background: #4b5563;
}

/* نمونه‌های آماده */
.templates-section {
  background: white;
  border-radius: 12px;
  padding: 30px;
  margin-bottom: 30px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.templates-section h2 {
  font-size: 1.5rem;
  color: #374151;
  margin-bottom: 20px;
  text-align: center;
}

.templates-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 20px;
}

.template-card {
  border: 2px solid #e5e7eb;
  border-radius: 12px;
  padding: 20px;
  cursor: pointer;
  transition: all 0.3s ease;
  background: #f9fafb;
}

.template-card:hover {
  border-color: #3b82f6;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.15);
  transform: translateY(-2px);
}

.template-card.active {
  border-color: #3b82f6;
  background: #eff6ff;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.15);
}

.template-preview {
  margin-bottom: 15px;
}

.mobile-header-preview {
  background: white;
  border-radius: 8px;
  padding: 15px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  max-width: 280px;
  margin: 0 auto;
}

.preview-header {
  background: #f8fafc;
  border-radius: 6px;
  padding: 12px;
}

.preview-top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.preview-logo {
  font-weight: bold;
  color: #f97316;
  font-size: 1.1rem;
}

.preview-icons {
  display: flex;
  gap: 8px;
}

.preview-icon {
  width: 32px;
  height: 32px;
  border-radius: 6px;
  background: #e5e7eb;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
}

.preview-search-section {
  display: flex;
  align-items: center;
  gap: 8px;
}

.preview-search-bar {
  flex: 1;
  background: white;
  border-radius: 20px;
  padding: 8px 12px;
  display: flex;
  align-items: center;
  gap: 8px;
  border: 1px solid #d1d5db;
}

.preview-search-icon {
  font-size: 12px;
  color: #6b7280;
}

.preview-search-text {
  font-size: 12px;
  color: #9ca3af;
}

.preview-cart-icon {
  width: 36px;
  height: 36px;
  border-radius: 18px;
  background: #f3f4f6;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  position: relative;
}

.preview-cart-icon::after {
  content: '1';
  position: absolute;
  top: -2px;
  right: -2px;
  background: #10b981;
  color: white;
  border-radius: 50%;
  width: 16px;
  height: 16px;
  font-size: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.template-info h4 {
  font-size: 1.1rem;
  color: #374151;
  margin: 0 0 8px 0;
  text-align: center;
}

.template-info p {
  font-size: 0.9rem;
  color: #6b7280;
  margin: 0;
  text-align: center;
  line-height: 1.4;
}

/* استایل‌های نمونه دوم */
.preview-banner-section {
  margin-bottom: 10px;
}

.preview-banner {
  background: #fef3c7;
  border: 1px solid #f59e0b;
  border-radius: 6px;
  padding: 8px 12px;
  position: relative;
  min-height: 40px;
  display: flex;
  align-items: center;
}

.preview-banner-content {
  flex: 1;
}

.preview-banner-text {
  font-size: 11px;
  color: #374151;
  line-height: 1.3;
}

.preview-paperclip {
  position: absolute;
  top: -2px;
  right: 8px;
  font-size: 12px;
  color: #6b7280;
}

.preview-search-section-2 {
  display: flex;
  align-items: center;
  gap: 8px;
}

.preview-logo-left-2 {
  font-size: 0.8rem;
  color: #374151;
  font-weight: 600;
  flex-shrink: 0;
}

.preview-search-bar-2 {
  flex: 1;
  background: white;
  border-radius: 20px;
  padding: 8px 12px;
  display: flex;
  align-items: center;
  gap: 8px;
  border: 1px solid #d1d5db;
}

.preview-search-text-2 {
  font-size: 11px;
  color: #6b7280;
  flex: 1;
}

.preview-notification-icon {
  width: 32px;
  height: 32px;
  border-radius: 6px;
  background: #f3f4f6;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  color: #6b7280;
}

/* استایل‌های نمونه سوم */
.preview-search-section-3 {
  display: flex;
  align-items: center;
}

.preview-search-bar-3 {
  width: 100%;
  background: white;
  border-radius: 8px;
  padding: 12px 16px;
  display: flex;
  align-items: center;
  gap: 12px;
  border: 1px solid #d1d5db;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.preview-logo-left {
  display: flex;
  align-items: center;
  gap: 6px;
  flex-shrink: 0;
}

.preview-logo-text {
  font-size: 11px;
  color: #1f2937;
  font-weight: 600;
}

.preview-search-text-3 {
  font-size: 12px;
  color: #9ca3af;
  flex: 1;
  text-align: right;
}

.preview-search-icon-3 {
  font-size: 14px;
  color: #6b7280;
  flex-shrink: 0;
}

/* پیش‌نمایش زنده */
.live-preview-section {
  background: white;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  position: sticky;
  top: 20px;
  height: 100%;
  min-height: 700px;
  display: flex;
  flex-direction: column;
}

.live-preview-section h3 {
  font-size: 1.25rem;
  color: #374151;
  margin-bottom: 15px;
  text-align: center;
}

.live-preview-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 700px;
}

.live-preview-mobile {
  background: linear-gradient(145deg, #2d3748, #1a202c);
  border-radius: 30px;
  padding: 8px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
  max-width: 350px;
  width: 100%;
  position: relative;
  border: 2px solid #4a5568;
  min-height: 700px;
}

.live-preview-mobile::before {
  content: '';
  position: absolute;
  top: 20px;
  left: 50%;
  transform: translateX(-50%);
  width: 60px;
  height: 4px;
  background: #4a5568;
  border-radius: 2px;
}

.live-preview-screen {
  background: white;
  border-radius: 25px;
  padding: 20px;
  min-height: 650px;
  position: relative;
  overflow: hidden;
  padding-top: 50px;
}

.live-preview-screen::before {
  content: '2:39 PM';
  position: absolute;
  top: 10px;
  left: 20px;
  color: #374151;
  font-size: 14px;
  font-weight: 600;
  z-index: 10;
}

.live-preview-screen::after {
  content: '●●●●●';
  position: absolute;
  top: 10px;
  right: 20px;
  color: #374151;
  font-size: 12px;
  z-index: 10;
}

.live-preview-header {
  background: #f8fafc;
  border-radius: 12px;
  padding: 16px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  margin-bottom: 20px;
  border: 1px solid #e5e7eb;
}

.live-preview-placeholder {
  text-align: center;
  padding: 60px 20px;
  color: #6b7280;
  background: white;
  border-radius: 12px;
  margin-bottom: 20px;
}

.live-preview-placeholder p {
  font-size: 1.1rem;
  margin: 0;
}

/* استایل‌های پیش‌نمایش زنده - نمونه کلاسیک */
.live-preview-top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.live-preview-logo {
  font-weight: bold;
  color: #f97316;
  font-size: 1.4rem;
}

.live-preview-icons {
  display: flex;
  gap: 6px;
}

.live-preview-icon {
  width: 36px;
  height: 36px;
  border-radius: 8px;
  background: #e5e7eb;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
}

.live-preview-search-section {
  display: flex;
  align-items: center;
  gap: 6px;
}

.live-preview-search-bar {
  flex: 1;
  background: white;
  border-radius: 24px;
  padding: 12px 16px;
  display: flex;
  align-items: center;
  gap: 8px;
  border: 1px solid #d1d5db;
}

.live-preview-search-icon {
  font-size: 14px;
  color: #6b7280;
}

.live-preview-search-text {
  font-size: 14px;
  color: #9ca3af;
}

.live-preview-cart-icon {
  width: 44px;
  height: 44px;
  border-radius: 22px;
  background: #f3f4f6;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  position: relative;
}

.live-preview-cart-icon::after {
  content: '1';
  position: absolute;
  top: -2px;
  right: -2px;
  background: #10b981;
  color: white;
  border-radius: 50%;
  width: 18px;
  height: 18px;
  font-size: 11px;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* استایل‌های پیش‌نمایش زنده - نمونه با بنر */
.live-preview-banner-section {
  margin-bottom: 8px;
}

.live-preview-banner {
  background: #fef3c7;
  border: 1px solid #f59e0b;
  border-radius: 8px;
  padding: 12px 16px;
  position: relative;
  min-height: 50px;
  display: flex;
  align-items: center;
}

.live-preview-banner-content {
  flex: 1;
}

.live-preview-banner-text {
  font-size: 13px;
  color: #374151;
  line-height: 1.4;
}

.live-preview-paperclip {
  position: absolute;
  top: -2px;
  right: 12px;
  font-size: 16px;
  color: #6b7280;
}

.live-preview-search-section-2 {
  display: flex;
  align-items: center;
  gap: 8px;
}

.live-preview-logo-left-2 {
  font-size: 1rem;
  color: #374151;
  font-weight: 600;
  flex-shrink: 0;
}

.live-preview-search-bar-2 {
  flex: 1;
  background: white;
  border-radius: 24px;
  padding: 12px 16px;
  display: flex;
  align-items: center;
  gap: 8px;
  border: 1px solid #d1d5db;
}

.live-preview-search-text-2 {
  font-size: 13px;
  color: #6b7280;
  flex: 1;
}

.live-preview-notification-icon {
  width: 44px;
  height: 44px;
  border-radius: 8px;
  background: #f3f4f6;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  color: #6b7280;
}

/* استایل‌های پیش‌نمایش زنده - نمونه مینیمال */
.live-preview-search-section-3 {
  display: flex;
  align-items: center;
}

.live-preview-search-bar-3 {
  width: 100%;
  background: white;
  border-radius: 12px;
  padding: 16px 20px;
  display: flex;
  align-items: center;
  gap: 12px;
  border: 1px solid #d1d5db;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.live-preview-logo-left {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-shrink: 0;
}

.live-preview-logo-text {
  font-size: 13px;
  color: #1f2937;
  font-weight: 600;
}

.live-preview-search-text-3 {
  font-size: 14px;
  color: #9ca3af;
  flex: 1;
  text-align: right;
}

.live-preview-search-icon-3 {
  font-size: 18px;
  color: #6b7280;
  flex-shrink: 0;
}

/* استایل‌های آپلود عکس */
.image-upload-section {
  margin-top: 8px;
}

.existing-images {
  margin-bottom: 16px;
}

.image-preview {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  border: 2px dashed #d1d5db;
  border-radius: 8px;
  background: #f9fafb;
  margin-bottom: 8px;
}

.image-label {
  font-size: 12px;
  font-weight: 500;
  color: #6b7280;
  min-width: 100px;
}

.preview-image {
  width: 80px;
  height: 60px;
  object-fit: cover;
  border-radius: 4px;
  border: 1px solid #e5e7eb;
}

.remove-image-btn {
  padding: 6px 12px;
  background: #ef4444;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 12px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.remove-image-btn:hover {
  background: #dc2626;
}

/* دکمه آپلود کوچک */
.small-upload-section {
  margin-top: 12px;
}

.small-upload-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 8px 12px;
  background: #f3f4f6;
  color: #374151;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.small-upload-btn:hover {
  background: #e5e7eb;
  border-color: #9ca3af;
}

/* بخش پیش‌نمایش عکس آپلود شده */
.uploaded-image-section {
  margin-top: 16px;
}

.uploaded-image-preview {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  border: 2px solid #10b981;
  border-radius: 8px;
  background: #f0fdf4;
  margin-bottom: 12px;
}

.uploaded-preview-image {
  width: 60px;
  height: 45px;
  object-fit: cover;
  border-radius: 4px;
  border: 1px solid #d1d5db;
}

.uploaded-image-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.uploaded-image-name {
  font-size: 13px;
  font-weight: 500;
  color: #065f46;
}

.uploaded-image-size {
  font-size: 11px;
  color: #6b7280;
}

/* گزینه‌های موقعیت */
.position-options {
  padding: 20px;
  border: 2px solid #3b82f6;
  border-radius: 8px;
  background: #eff6ff;
  text-align: center;
}

.position-options h4 {
  margin: 0 0 16px 0;
  color: #1e40af;
  font-size: 16px;
}

.position-buttons {
  display: flex;
  gap: 12px;
  justify-content: center;
  margin-bottom: 16px;
}

.position-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 20px;
  background: #3b82f6;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.2s;
}

.position-btn:hover {
  background: #2563eb;
}

.cancel-btn {
  padding: 8px 16px;
  background: #6b7280;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 12px;
  cursor: pointer;
  transition: background-color 0.2s;
}

/* استایل‌های پیش‌نمایش عکس‌ها */
.preview-top-image,
.preview-bottom-image {
  width: 100%;
  margin: 4px 0;
  border-radius: 4px;
  overflow: hidden;
}

.preview-top-image {
  margin-bottom: 8px;
}

.preview-bottom-image {
  margin-top: 8px;
}

.preview-image {
  width: 100%;
  height: auto;
  max-height: 60px;
  object-fit: cover;
  border-radius: 4px;
  border: 1px solid #e5e7eb;
}

/* چیدمان کناری */
.form-layout {
  display: grid;
  grid-template-columns: 1fr 450px;
  gap: 40px;
  align-items: stretch;
}

@media (max-width: 768px) {
  .form-grid {
    grid-template-columns: 1fr;
  }
  
  .page-header-flex {
    flex-direction: column;
    gap: 20px;
    text-align: center;
  }
  
  .form-actions {
    flex-direction: column;
  }
  
  .form-layout {
    grid-template-columns: 1fr;
  }
  
  .live-preview-section {
    position: static;
  }
}
</style>
