<template>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- هدر صفحه -->
      <div class="mb-6">
        <h1 class="text-2xl font-bold text-gray-900 text-right">مدیریت احراز هویت کاربران</h1>
        <p class="text-sm text-gray-600 text-right mt-1">بررسی و تایید مدارک ارسالی کاربران</p>
      </div>

      <!-- فیلترها و آمار -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <div class="flex items-center justify-between mb-4">
          <div class="flex gap-2">
            <button
              :class="[
                'px-4 py-2 rounded text-sm font-medium transition-colors',
                currentFilter === 'all' 
                  ? 'bg-blue-600 text-white' 
                  : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
              ]"
              @click="currentFilter = 'all'"
            >
              همه ({{ stats.total }})
            </button>
            <button
              :class="[
                'px-4 py-2 rounded text-sm font-medium transition-colors',
                currentFilter === 'pending' 
                  ? 'bg-yellow-600 text-white' 
                  : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
              ]"
              @click="currentFilter = 'pending'"
            >
              در انتظار ({{ stats.pending }})
            </button>
            <button
              :class="[
                'px-4 py-2 rounded text-sm font-medium transition-colors',
                currentFilter === 'verified' 
                  ? 'bg-green-600 text-white' 
                  : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
              ]"
              @click="currentFilter = 'verified'"
            >
              تایید شده ({{ stats.verified }})
            </button>
            <button
              :class="[
                'px-4 py-2 rounded text-sm font-medium transition-colors',
                currentFilter === 'rejected' 
                  ? 'bg-red-600 text-white' 
                  : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
              ]"
              @click="currentFilter = 'rejected'"
            >
              رد شده ({{ stats.rejected }})
            </button>
          </div>
          <button
            class="px-4 py-2 bg-gray-100 rounded hover:bg-gray-200 transition-colors"
            @click="loadVerifications"
          >
            🔄 بارگیری مجدد
          </button>
        </div>
      </div>

      <!-- لیست درخواست‌ها -->
      <div v-if="loading" class="text-center py-12">
        <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
        <p class="text-gray-600 mt-2">در حال بارگیری...</p>
      </div>

      <div v-else-if="filteredVerifications.length === 0" class="bg-white rounded-lg shadow p-12 text-center">
        <p class="text-gray-500">درخواستی برای نمایش وجود ندارد</p>
      </div>

      <div v-else class="space-y-4">
        <div
          v-for="verification in filteredVerifications"
          :key="verification.id"
          class="bg-white rounded-lg shadow overflow-hidden"
        >
          <div class="p-6">
            <div class="flex items-start justify-between">
              <!-- اطلاعات کاربر و درخواست -->
              <div class="flex-1">
                <div class="flex items-center gap-3 mb-2">
                  <h3 class="text-lg font-bold text-gray-900">
                    شناسه کاربر: {{ verification.user_id }}
                  </h3>
                  <span
                    :class="[
                      'text-xs px-3 py-1 rounded-full font-medium',
                      verification.status === 'pending' ? 'bg-yellow-100 text-yellow-800' :
                      verification.status === 'verified' ? 'bg-green-100 text-green-800' :
                      'bg-red-100 text-red-800'
                    ]"
                  >
                    {{ getStatusLabel(verification.status) }}
                  </span>
                </div>
                
                <div class="text-sm text-gray-600 space-y-1">
                  <p><strong>نوع درخواست:</strong> {{ getFieldLabel(verification.field_name) }}</p>
                  <p><strong>مقدار:</strong> {{ verification.field_value }}</p>
                  <p><strong>تاریخ ثبت:</strong> {{ formatDate(verification.created_at) }}</p>
                  <p v-if="verification.verification_method">
                    <strong>روش احراز:</strong> {{ verification.verification_method }}
                  </p>
                </div>

                <!-- نمایش مدارک (اگر وجود دارد) -->
                <div v-if="verification.documents && verification.documents.length > 0" class="mt-4">
                  <p class="text-sm font-medium text-gray-700 mb-2">مدارک ارسالی:</p>
                  <div class="flex gap-3 flex-wrap">
                    <div
                      v-for="(doc, idx) in verification.documents"
                      :key="idx"
                      class="relative group"
                    >
                      <img
                        :src="doc"
                        class="w-32 h-32 object-cover rounded border cursor-pointer hover:opacity-75 transition-opacity"
                        alt="مدرک"
                        @click="openImageModal(doc)"
                      />
                      <button
                        class="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-20 transition-opacity flex items-center justify-center"
                        @click="openImageModal(doc)"
                      >
                        <span class="text-white opacity-0 group-hover:opacity-100">🔍 مشاهده</span>
                      </button>
                    </div>
                  </div>
                </div>

                <!-- دلیل رد (اگر رد شده) -->
                <div v-if="verification.rejection_reason" class="mt-3 p-3 bg-red-50 rounded">
                  <p class="text-sm text-red-800">
                    <strong>دلیل رد:</strong> {{ verification.rejection_reason }}
                  </p>
                </div>
              </div>

              <!-- دکمه‌های عملیات -->
              <div v-if="verification.status === 'pending'" class="mr-4 flex flex-col gap-2">
                <button
                  :disabled="actionLoading === verification.id"
                  class="px-4 py-2 bg-green-600 text-white rounded hover:bg-green-700 disabled:bg-gray-400 transition-colors text-sm font-medium whitespace-nowrap"
                  @click="verifyRequest(verification.id)"
                >
                  ✓ تایید
                </button>
                <button
                  :disabled="actionLoading === verification.id"
                  class="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700 disabled:bg-gray-400 transition-colors text-sm font-medium whitespace-nowrap"
                  @click="openRejectModal(verification)"
                >
                  ✗ رد
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Modal رد درخواست -->
  <div
    v-if="rejectModal.show"
    class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
    @click.self="closeRejectModal"
  >
    <div class="bg-white rounded-lg p-6 max-w-md w-full mx-4">
      <h3 class="text-lg font-bold text-gray-900 mb-4 text-right">رد درخواست احراز هویت</h3>
      
      <div class="mb-4">
        <label class="block text-sm font-medium text-gray-700 mb-2 text-right">
          دلیل رد (اختیاری)
        </label>
        <textarea
          v-model="rejectModal.reason"
          rows="4"
          class="w-full px-3 py-2 border border-gray-300 rounded text-sm text-right focus:ring-2 focus:ring-red-500 focus:border-red-500"
          placeholder="توضیحات مربوط به رد درخواست را وارد کنید..."
        ></textarea>
      </div>

      <div class="flex gap-2 justify-end">
        <button
          class="px-4 py-2 bg-gray-200 text-gray-700 rounded hover:bg-gray-300 transition-colors"
          @click="closeRejectModal"
        >
          انصراف
        </button>
        <button
          :disabled="actionLoading"
          class="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700 disabled:bg-gray-400 transition-colors"
          @click="confirmReject"
        >
          {{ actionLoading ? 'در حال رد...' : 'رد درخواست' }}
        </button>
      </div>
    </div>
  </div>

  <!-- Modal نمایش تصویر -->
  <div
    v-if="imageModal.show"
    class="fixed inset-0 bg-black bg-opacity-75 flex items-center justify-center z-50"
    @click="closeImageModal"
  >
    <div class="relative max-w-4xl max-h-screen p-6">
      <button
        class="absolute top-6 left-6 text-white bg-black bg-opacity-50 rounded-full w-10 h-10 flex items-center justify-center hover:bg-opacity-75 transition-opacity"
        @click="closeImageModal"
      >
        ✕
      </button>
      <img
        :src="imageModal.url"
        class="max-w-full max-h-screen rounded"
        alt="نمایش تصویر"
      />
    </div>
  </div>
</template>

<script setup>
// تعریف layout (middleware در layout خودش هست)
definePageMeta({
  layout: 'admin-main'
})

// State
const verifications = ref([])
const loading = ref(false)
const actionLoading = ref(null)
const currentFilter = ref('pending')

const stats = ref({
  total: 0,
  pending: 0,
  verified: 0,
  rejected: 0
})

const rejectModal = ref({
  show: false,
  verification: null,
  reason: ''
})

const imageModal = ref({
  show: false,
  url: ''
})

// Computed
const filteredVerifications = computed(() => {
  if (currentFilter.value === 'all') {
    return verifications.value
  }
  return verifications.value.filter(v => v.status === currentFilter.value)
})

// Methods
const loadVerifications = async () => {
  loading.value = true
  try {
    const response = await $fetch('/api/admin/verifications', {
      credentials: 'include'
    })

    verifications.value = response.data || []
    
    // محاسبه آمار
    stats.value.total = verifications.value.length
    stats.value.pending = verifications.value.filter(v => v.status === 'pending').length
    stats.value.verified = verifications.value.filter(v => v.status === 'verified').length
    stats.value.rejected = verifications.value.filter(v => v.status === 'rejected').length

  } catch (error) {
    console.error('خطا در بارگیری درخواست‌ها:', error)
    alert('خطا در بارگیری درخواست‌ها')
  } finally {
    loading.value = false
  }
}

const verifyRequest = async (id) => {
  if (!confirm('آیا از تایید این درخواست اطمینان دارید؟')) return

  actionLoading.value = id
  try {
    await $fetch(`/api/admin/verifications/${id}/verify`, {
      method: 'POST',
      credentials: 'include'
    })

    alert('درخواست با موفقیت تایید شد')
    await loadVerifications()
  } catch (error) {
    console.error('خطا در تایید درخواست:', error)
    alert('خطا در تایید درخواست')
  } finally {
    actionLoading.value = null
  }
}

const openRejectModal = (verification) => {
  rejectModal.value = {
    show: true,
    verification: verification,
    reason: ''
  }
}

const closeRejectModal = () => {
  rejectModal.value = {
    show: false,
    verification: null,
    reason: ''
  }
}

const confirmReject = async () => {
  if (!rejectModal.value.verification) return

  actionLoading.value = rejectModal.value.verification.id
  try {
    await $fetch(`/api/admin/verifications/${rejectModal.value.verification.id}/reject`, {
      method: 'POST',
      body: {
        reason: rejectModal.value.reason || 'مدارک نامعتبر یا ناقص'
      },
      credentials: 'include'
    })

    alert('درخواست رد شد')
    closeRejectModal()
    await loadVerifications()
  } catch (error) {
    console.error('خطا در رد درخواست:', error)
    alert('خطا در رد درخواست')
  } finally {
    actionLoading.value = null
  }
}

const openImageModal = (url) => {
  imageModal.value = {
    show: true,
    url: url
  }
}

const closeImageModal = () => {
  imageModal.value = {
    show: false,
    url: ''
  }
}

const getStatusLabel = (status) => {
  const labels = {
    pending: 'در انتظار بررسی',
    verified: 'تایید شده',
    rejected: 'رد شده'
  }
  return labels[status] || status
}

const getFieldLabel = (fieldName) => {
  const labels = {
    national_code: 'کد ملی',
    sheba_number: 'شماره شبا',
    birth_date: 'تاریخ تولد',
    identity_documents: 'مدارک هویتی',
    legal_info: 'اطلاعات حقوقی'
  }
  return labels[fieldName] || fieldName
}

const formatDate = (dateString) => {
  if (!dateString) return '-'
  const date = new Date(dateString)
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date)
}

// Load on mount
onMounted(() => {
  loadVerifications()
})
</script>
