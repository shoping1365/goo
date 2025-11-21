<template>
  <div class="font-iranyekan">
    <!-- عنوان صفحه -->
    <div class="header-bg">
      <div class="page-header-flex">
        <h1 class="page-title">مدیریت فوتر</h1>
        <button class="btn btn-primary new-item-btn" @click="addNewFooter">افزودن فوتر</button>
      </div>
    </div>

    <!-- محتوای اصلی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div v-if="loading" class="text-center text-gray-500 py-16">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600 mx-auto mb-4"></div>
        <p class="font-iranyekan text-lg">در حال بارگذاری...</p>
      </div>
      
      <div v-else-if="footerSettings.length === 0" class="text-center text-gray-500 py-16">
        <div class="text-6xl mb-4">📄</div>
        <p class="font-iranyekan text-lg mb-2">هیچ فوتری ایجاد نشده است</p>
        <p class="text-sm text-gray-400 font-iranyekan">برای شروع، روی دکمه "افزودن فوتر" کلیک کنید</p>
      </div>
      
      <div v-else>
        <div class="grid gap-6">
          <div 
            v-for="footer in footerSettings" 
            :key="footer.id"
            class="border border-gray-200 rounded-xl p-6 hover:shadow-lg transition-all duration-300 bg-gradient-to-r from-gray-50 to-white"
          >
            <div class="flex items-start justify-between">
              <div class="flex-1">
                <div class="flex items-center gap-3 mb-3">
                  <h3 class="text-xl font-bold text-gray-900 font-iranyekan">{{ footer.title }}</h3>
                  <span
class="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium font-iranyekan"
                        :class="footer.is_active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'">
                    <span
class="w-2 h-2 rounded-full mr-2"
                          :class="footer.is_active ? 'bg-green-500' : 'bg-red-500'"></span>
                    {{ footer.is_active ? 'فعال' : 'غیرفعال' }}
                  </span>
                </div>
                
                <div class="mb-3">
                  <span class="text-sm text-gray-500 font-iranyekan">نوع نمایش: </span>
                  <span class="text-sm font-medium text-gray-700 font-iranyekan">{{ getPageSelectionLabel(footer.page_selection) }}</span>
                </div>
                
                <!-- 
                  ⚠️ امنیت XSS: استفاده از v-html خطرناک است!
                  
                  این کد محتوای HTML را بدون sanitization نمایش می‌دهد که می‌تواند منجر به حملات XSS شود.
                  
                  ✅ راه حل صحیح:
                  1. قبل از استفاده از v-html، محتوا را با کتابخانه sanitization (مثل DOMPurify) پاکسازی کنید
                  2. یا از {{ }} به جای v-html استفاده کنید اگر HTML نیاز نیست
                  3. محتوای کاربر را هرگز بدون sanitization در v-html قرار ندهید
                  
                  مثال صحیح:
                  import DOMPurify from 'dompurify'
                  const sanitizedContent = computed(() => DOMPurify.sanitize(truncateContent(footer.content, 150)))
                  <div v-html="sanitizedContent"></div>
                -->
                <div class="text-gray-700 text-sm leading-relaxed font-iranyekan mb-4">{{ truncateContent(footer.content, 150) }}</div>
                
                <div class="text-xs text-gray-400 font-iranyekan">
                  ایجاد شده در: {{ formatDate(footer.created_at) }}
                </div>
              </div>
              
              <div class="flex flex-col gap-2 ml-4">
                <NuxtLink
                  :to="`/admin/content/footer-management/edit/${footer.id}`"
                  class="bg-blue-500 text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-blue-600 transition-all duration-200 shadow-sm hover:shadow-md font-iranyekan flex items-center justify-center min-w-[80px]"
                >
                  <span class="ml-1">✏️</span>
                  ویرایش
                </NuxtLink>
                
                <button
                  class="bg-red-500 text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-red-600 transition-all duration-200 shadow-sm hover:shadow-md font-iranyekan flex items-center justify-center min-w-[80px]"
                  @click="openDeleteConfirm(footer.id)"
                >
                  <span class="ml-1">🗑️</span>
                  حذف
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Toast Notification -->
    <div
v-if="showToast" 
         class="fixed top-6 right-4 z-50 transition-all duration-300"
         :class="toastType === 'success' ? 'bg-green-500' : 'bg-red-500'">
      <div class="text-white px-6 py-3 rounded-lg shadow-lg font-iranyekan">
        {{ toastMessage }}
      </div>
    </div>

    <!-- Delete Confirm Modal -->
    <DeleteConfirmModal 
      ref="deleteModalRef"
      :single-delete-title="'تأیید حذف فوتر'"
      :single-delete-message="'آیا از حذف این فوتر مطمئن هستید؟ این عملیات غیرقابل بازگشت است.'"
      @single-delete="handleSingleDelete"
    />
  </div>
</template>

<script setup>
// تعریف متا صفحه
definePageMeta({
  layout: 'admin-main'
})

// متغیرهای reactive
const loading = ref(true)
const footerSettings = ref([])
const showToast = ref(false)
const toastMessage = ref('')
const toastType = ref('success')
const deleteModalRef = ref(null)

// دریافت تنظیمات فوتر
const fetchFooterSettings = async () => {
  try {
    loading.value = true
  const response = await $fetch('/api/admin/footer-settings')
    if (response.success) {
      footerSettings.value = response.data || []
    }
  } catch (error) {
    console.error('خطا در دریافت تنظیمات فوتر:', error)
    showToastMessage('خطا در دریافت تنظیمات فوتر', 'error')
  } finally {
    loading.value = false
  }
}

// باز کردن dialog تأیید حذف
const openDeleteConfirm = (footerId) => {
  if (deleteModalRef.value) {
    deleteModalRef.value.openDeleteConfirm(footerId)
  }
}

// مدیریت حذف تک آیتم
const handleSingleDelete = async (footerId) => {
  await deleteFooter(footerId)
}

// حذف فوتر
const deleteFooter = async (id) => {
  try {
  const response = await $fetch(`/api/admin/footer-settings/${id}`, {
      method: 'DELETE'
    })
    
    if (response.success) {
      showToastMessage('فوتر با موفقیت حذف شد', 'success')
      await fetchFooterSettings() // بارگذاری مجدد
    } else {
      showToastMessage('خطا در حذف فوتر', 'error')
    }
  } catch (error) {
    console.error('خطا در حذف فوتر:', error)
    showToastMessage('خطا در حذف فوتر', 'error')
  }
}

// نمایش پیام toast
const showToastMessage = (message, type = 'success') => {
  toastMessage.value = message
  toastType.value = type
  showToast.value = true
  
  setTimeout(() => {
    showToast.value = false
  }, 3000)
}

// برچسب نوع نمایش صفحه
const getPageSelectionLabel = (pageSelection) => {
  const labels = {
    'all': 'همه صفحات',
    'specific': 'صفحات خاص',
    'exclude': 'به جز صفحات خاص'
  }
  return labels[pageSelection] || pageSelection
}

// کوتاه کردن محتوا
const truncateContent = (content, maxLength) => {
  if (!content) return ''
  const strippedContent = content.replace(/<[^>]*>/g, '')
  if (strippedContent.length <= maxLength) return content
  return strippedContent.substring(0, maxLength) + '...'
}

// فرمت تاریخ
const formatDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

// افزودن فوتر جدید
const addNewFooter = () => {
  navigateTo('/admin/content/footer-management/create')
}

// اجرا در mount
onMounted(() => {
  fetchFooterSettings()
})
</script>

<style scoped>
.font-iranyekan {
  font-family: 'Yekan', 'Tahoma', sans-serif;
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
  margin: 0;
}

.new-item-btn {
  background: rgba(255, 255, 255, 0.2);
  border: 2px solid rgba(255, 255, 255, 0.3);
  color: white;
  padding: 12px 24px;
  border-radius: 8px;
  font-weight: 600;
  transition: all 0.3s ease;
}

.new-item-btn:hover {
  background: rgba(255, 255, 255, 0.3);
  border-color: rgba(255, 255, 255, 0.5);
}

/* بهبود انیمیشن‌ها */
.transition-all {
  transition-property: all;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}

/* بهبود hover effects */
.hover\:shadow-lg:hover {
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

.hover\:shadow-md:hover {
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
}
</style>
