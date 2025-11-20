<template>
  <div class="rounded-lg shadow-sm border border-gray-200 bg-white">
    <!-- Header -->
    <div class="px-16 py-10 border-b border-gray-200 bg-white shadow-sm">
      <div class="flex items-center justify-between">
        <div>
          <h3 class="text-lg font-medium text-gray-900">درگاه‌های پرداخت</h3>
          <p class="text-sm text-gray-500">مدیریت و تنظیم درگاه‌های پرداخت</p>
        </div>
        <TemplateButton
          bg-gradient="bg-gradient-to-r from-emerald-500 to-teal-600"
          text-color="text-white"
          border-color="border border-emerald-500"
          hover-class="hover:from-emerald-600 hover:to-teal-700"
          focus-class="focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2"
          size="medium"
          @click="openAddModal"
        >
          <Icon name="heroicons:plus" class="w-4 h-4 mr-2" />
          افزودن درگاه جدید
        </TemplateButton>
      </div>
    </div>

    <!-- Filters -->
    <div class="px-6 py-4 border-b border-gray-200 bg-white">
      <div class="flex items-center space-x-4 space-x-reverse">
        <div class="flex-1">
          <input
            v-model="searchQuery"
            type="text"
            placeholder="جستجو در درگاه‌ها..."
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
        <select
          v-model="statusFilter"
          class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
            <option value="">همه وضعیت‌ها</option>
            <option value="active">فعال</option>
            <option value="inactive">غیرفعال</option>
          <option value="maintenance">تعمیر و نگهداری</option>
        </select>
        <select
          v-model="typeFilter"
          class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="">همه انواع</option>
          <option value="iranian">ایرانی</option>
          <option value="international">بین‌المللی</option>
          <option value="digital">دیجیتال</option>
          </select>
        </div>
        </div>

    <!-- Loading State -->
    <div v-if="loading" class="px-6 py-8 text-center">
      <div class="inline-flex items-center">
        <Icon name="heroicons:arrow-path" class="w-5 h-5 mr-2 animate-spin" />
        در حال بارگذاری...
      </div>
    </div>

    <!-- Gateway List -->
    <div v-else-if="filteredGateways.length > 0" class="grid grid-cols-1 lg:grid-cols-2 gap-6 p-6">
      <div
        v-for="gateway in filteredGateways"
        :key="gateway.id"
        class="bg-white rounded-xl shadow-sm border border-gray-200 hover:shadow-md transition-all duration-200 overflow-hidden"
      >
        <!-- Header Card -->
        <div class="p-6 border-b border-gray-100">
          <div class="flex items-start justify-between">
            <!-- Gateway Info -->
            <div class="flex items-start space-x-6 space-x-reverse">
              <!-- Gateway Icon with Name Overlay -->
              <div class="relative w-20 h-20 rounded-xl flex items-center justify-center bg-blue-50 border-2 border-blue-200">
                <span class="text-2xl text-blue-600">{{ gateway.icon || gateway.name.charAt(0) }}</span>
                
                <!-- Name Overlay -->
                <div class="absolute inset-0 flex items-center justify-center">
                  <h3 class="text-sm font-bold text-white text-center px-1 drop-shadow-lg">{{ gateway.name }}</h3>
                </div>
              </div>

              <!-- Gateway Details -->
              <div class="flex-1 mr-12">
                <h3 class="text-xl font-bold text-gray-900 mb-1">{{ gateway.name }}</h3>
                <p class="text-sm text-gray-500 mb-2">{{ gateway.english_name }}</p>
                
                <!-- Status Badges -->
                <div class="flex items-center space-x-2 space-x-reverse mt-4">
                  <span
                    :class="[
                      'inline-flex items-center px-3 py-1 rounded-full text-xs font-semibold',
                      getStatusClass(gateway.status)
                    ]"
                  >
                    <div
:class="[
                      'w-2 h-2 rounded-full mr-2',
                      gateway.status === 'active' ? 'bg-green-400' : 
                      gateway.status === 'inactive' ? 'bg-gray-400' : 'bg-yellow-400'
                    ]"></div>
                    {{ getStatusText(gateway.status) }}
                  </span>
                  
                  <span class="text-xs text-gray-400">•</span>
                  
                  <span class="text-xs text-gray-600 bg-gray-100 px-2 py-1 rounded">
                    {{ getTypeText(gateway.type) }}
                  </span>
                  
                  <span v-if="gateway.is_test_mode" class="text-xs text-orange-600 bg-orange-100 px-2 py-1 rounded">
                    حالت تست
                  </span>
                </div>
              </div>
            </div>

            <!-- Actions -->
            <div class="flex flex-col space-y-2">
              <!-- Test Connection -->
              <TemplateButton
                :disabled="testingConnection === gateway.id"
                bg-gradient="bg-gradient-to-r from-blue-100 to-blue-200"
                text-color="text-blue-700"
                border-color="border border-blue-200"
                hover-class="hover:from-blue-200 hover:to-blue-300"
                focus-class="focus:ring-2 focus:ring-blue-200 focus:ring-offset-2"
                size="medium"
                :custom-class="testingConnection === gateway.id ? 'opacity-50 cursor-not-allowed' : ''"
                @click="testConnection(gateway.id)"
              >
                <Icon
                  :name="testingConnection === gateway.id ? 'heroicons:arrow-path' : 'heroicons:signal'"
                  :class="[
                    'w-4 h-4 mr-2',
                    testingConnection === gateway.id ? 'animate-spin' : ''
                  ]"
                />
                {{ testingConnection === gateway.id ? 'در حال تست...' : 'تست اتصال' }}
              </TemplateButton>

              <!-- Edit -->
              <TemplateButton
                bg-gradient="bg-gradient-to-r from-purple-100 to-purple-200"
                text-color="text-purple-700"
                border-color="border border-purple-200"
                hover-class="hover:from-purple-200 hover:to-purple-300"
                focus-class="focus:ring-2 focus:ring-purple-200 focus:ring-offset-2"
                size="medium"
                @click="editGateway(gateway)"
              >
                <Icon name="heroicons:pencil-square" class="w-4 h-4 mr-2" />
                ویرایش
              </TemplateButton>

              <!-- Toggle Status -->
              <TemplateButton
                :bg-gradient="gateway.status === 'active' ? 'bg-gradient-to-r from-yellow-100 to-yellow-200' : 'bg-gradient-to-r from-green-100 to-green-200'"
                :text-color="gateway.status === 'active' ? 'text-yellow-700' : 'text-green-700'"
                :border-color="gateway.status === 'active' ? 'border border-yellow-200' : 'border border-green-200'"
                :hover-class="gateway.status === 'active' ? 'hover:from-yellow-200 hover:to-yellow-300' : 'hover:from-green-200 hover:to-green-300'"
                :focus-class="gateway.status === 'active' ? 'focus:ring-2 focus:ring-yellow-200 focus:ring-offset-2' : 'focus:ring-2 focus:ring-green-200 focus:ring-offset-2'"
                size="medium"
                @click="toggleStatus(gateway)"
              >
                <Icon
                  :name="gateway.status === 'active' ? 'heroicons:pause' : 'heroicons:play'"
                  class="w-4 h-4 mr-2"
                />
                {{ gateway.status === 'active' ? 'غیرفعال' : 'فعال' }}
              </TemplateButton>

              <!-- Delete -->
              <TemplateButton
                bg-gradient="bg-gradient-to-r from-red-100 to-red-200"
                text-color="text-red-700"
                border-color="border border-red-200"
                hover-class="hover:from-red-200 hover:to-red-300"
                focus-class="focus:ring-2 focus:ring-red-200 focus:ring-offset-2"
                size="medium"
                @click="confirmDelete(gateway)"
              >
                <Icon name="heroicons:trash" class="w-4 h-4 mr-2" />
                حذف
              </TemplateButton>
            </div>
          </div>
        </div>

        <!-- Gateway Configuration -->
        <div class="p-6 border-b border-gray-100">
          <h4 class="text-sm font-semibold text-gray-700 mb-3">تنظیمات درگاه</h4>
                      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="bg-gray-50 rounded-lg p-3">
              <div class="text-xs text-gray-500 mb-1">کارمزد</div>
              <div class="text-lg font-bold text-gray-900">{{ gateway.fee }}%</div>
            </div>
            <div class="bg-gray-50 rounded-lg p-3">
              <div class="text-xs text-gray-500 mb-1">حداقل مبلغ</div>
              <div class="text-sm font-semibold text-gray-900">{{ formatCurrency(gateway.min_amount) }}</div>
            </div>
            <div class="bg-gray-50 rounded-lg p-3">
              <div class="text-xs text-gray-500 mb-1">حداکثر مبلغ</div>
              <div class="text-sm font-semibold text-gray-900">{{ formatCurrency(gateway.max_amount) }}</div>
            </div>
          </div>
        </div>

        <!-- Statistics -->
        <div class="p-6">
          <h4 class="text-sm font-semibold text-gray-700 mb-3">آمار امروز</h4>
          <div class="grid grid-cols-2 gap-6">
            <div class="bg-blue-50 rounded-lg p-3">
              <div class="text-xs text-blue-600 mb-1">تراکنش‌های امروز</div>
              <div class="text-lg font-bold text-blue-900">{{ gateway.today_transactions }}</div>
            </div>
            <div class="bg-green-50 rounded-lg p-3">
              <div class="text-xs text-green-600 mb-1">درآمد امروز</div>
              <div class="text-lg font-bold text-green-900">{{ formatCurrency(gateway.today_revenue) }}</div>
            </div>
            <div class="bg-purple-50 rounded-lg p-3">
              <div class="text-xs text-purple-600 mb-1">کل تراکنش‌ها</div>
              <div class="text-lg font-bold text-purple-900">{{ gateway.total_transactions }}</div>
            </div>
            <div class="bg-orange-50 rounded-lg p-3">
              <div class="text-xs text-orange-600 mb-1">کل درآمد</div>
              <div class="text-lg font-bold text-orange-900">{{ formatCurrency(gateway.total_revenue) }}</div>
            </div>
          </div>
        </div>

        <!-- تراکنش‌های درگاه -->
        <div class="p-6 border-t border-gray-100">
          <GatewayTransactionSummary 
            :gateway-id="gateway.id"
          />
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else class="px-6 py-8 text-center">
      <Icon name="heroicons:credit-card" class="w-12 h-12 mx-auto text-gray-400" />
      <h3 class="mt-2 text-sm font-medium text-gray-900">هیچ درگاهی یافت نشد</h3>
      <p class="mt-1 text-sm text-gray-500">درگاه پرداخت جدیدی اضافه کنید.</p>
            </div>

    <!-- Test Connection Result Modal -->
    <div
      v-if="showTestResult"
      class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50"
      @click="closeTestResult"
    >
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white" @click.stop>
        <div class="mt-3 text-center">
          <div
            :class="[
              'mx-auto flex items-center justify-center h-12 w-12 rounded-full',
              testResult.success ? 'bg-green-100' : 'bg-red-100'
            ]"
          >
            <Icon
              :name="testResult.success ? 'heroicons:check' : 'heroicons:x-mark'"
              :class="[
                'h-6 w-6',
                testResult.success ? 'text-green-600' : 'text-red-600'
              ]"
            />
                </div>
          <h3 class="text-lg leading-6 font-medium text-gray-900 mt-4">
            {{ testResult.success ? 'اتصال موفق' : 'اتصال ناموفق' }}
          </h3>
          <div class="mt-2 px-7 py-3">
            <p class="text-sm text-gray-500">{{ testResult.message }}</p>
                </div>
          <div class="items-center px-4 py-3">
            <TemplateButton
              bg-gradient="bg-gradient-to-r from-gray-500 to-gray-600"
              text-color="text-white"
              border-color="border border-gray-500"
              hover-class="hover:from-gray-600 hover:to-gray-700"
              focus-class="focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
              size="medium"
              custom-class="w-full"
              @click="closeTestResult"
            >
              بستن
            </TemplateButton>
          </div>
        </div>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div
      v-if="showDeleteModal"
      class="fixed inset-0 overflow-y-auto h-full w-full z-50"
      @click="closeDeleteModal"
    >
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white" @click.stop>
        <div class="mt-3 text-center">
          <div class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-red-100">
            <Icon name="heroicons:exclamation-triangle" class="h-6 w-6 text-red-600" />
          </div>
          <h3 class="text-lg leading-6 font-medium text-gray-900 mt-4">
            حذف درگاه پرداخت
          </h3>
          <div class="mt-2 px-7 py-3">
            <p class="text-sm text-gray-500">
              آیا از حذف درگاه <strong>{{ gatewayToDelete?.name }}</strong> اطمینان دارید؟
            </p>
            <p class="text-xs text-red-500 mt-2">
              این عملیات غیرقابل بازگشت است و تمام اطلاعات مربوط به این درگاه حذف خواهد شد.
            </p>
          </div>
          <div class="flex items-center space-x-3 space-x-reverse px-4 py-3">
            <TemplateButton
              bg-gradient="bg-gradient-to-r from-gray-500 to-gray-600"
              text-color="text-white"
              border-color="border border-gray-500"
              hover-class="hover:from-gray-600 hover:to-gray-700"
              focus-class="focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
              size="medium"
              custom-class="flex-1"
              @click="closeDeleteModal"
            >
              انصراف
            </TemplateButton>
            <TemplateButton
              :disabled="deleting"
              bg-gradient="bg-gradient-to-r from-red-300 to-pink-400"
              text-color="text-white"
              border-color="border border-red-300"
              hover-class="hover:from-red-400 hover:to-pink-500"
              focus-class="focus:ring-2 focus:ring-red-300 focus:ring-offset-2"
              size="medium"
              custom-class="flex-1"
              @click="deleteGateway"
            >
              <Icon v-if="deleting" name="heroicons:arrow-path" class="w-4 h-4 mr-2 animate-spin" />
              {{ deleting ? 'در حال حذف...' : 'حذف' }}
            </TemplateButton>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import TemplateButton from '~/components/common/TemplateButton.vue'
import GatewayTransactionSummary from './GatewayTransactionSummary.vue'

// Types
interface PaymentGateway {
  id: number
  name: string
  english_name: string
  type: string
  status: string
  icon: string
  color: string
  fee: number
  min_amount: number
  max_amount: number
  today_transactions: number
  today_revenue: number
  total_transactions: number
  total_revenue: number
  is_test_mode: boolean
}

// Reactive state
const gateways = ref<PaymentGateway[]>([])
const loading = ref(true)
const searchQuery = ref('')
const statusFilter = ref('')
const typeFilter = ref('')
const testingConnection = ref<number | null>(null)
const showTestResult = ref(false)
const testResult = ref({
  success: false,
  message: ''
})
const showDeleteModal = ref(false)
const gatewayToDelete = ref<PaymentGateway | null>(null)
const deleting = ref(false)

const maxRetries = 10
const retryCount = ref(0)
const lastError = ref(null)

// Computed properties
const filteredGateways = computed(() => {
  return gateways.value.filter(gateway => {
    const matchesSearch = gateway.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
                         gateway.english_name.toLowerCase().includes(searchQuery.value.toLowerCase())
    const matchesStatus = !statusFilter.value || gateway.status === statusFilter.value
    const matchesType = !typeFilter.value || gateway.type === typeFilter.value
    
    return matchesSearch && matchesStatus && matchesType
  })
})

// Methods
const fetchGateways = async (isManual = false) => {
  try {
    loading.value = true
    lastError.value = null
    interface GatewayData {
      [key: string]: unknown
    }
    interface ApiResponse {
      data?: GatewayData[]
    }
    const response = await $fetch<ApiResponse>('/api/payment-gateways/filters?status=active')
    gateways.value = (response.data || []) as unknown as PaymentGateway[]
    retryCount.value = 0 // موفقیت، شمارنده ریست شود
  } catch (error) {
    lastError.value = error
    if (!isManual && retryCount.value < maxRetries) {
      retryCount.value++
      setTimeout(() => fetchGateways(false), 1000)
    }
    // اگر دستی بود یا به سقف رسید، دیگر تلاش نکن
  } finally {
    loading.value = false
  }
}

const manualRefresh = () => {
  // فقط یک بار تلاش مجدد
  fetchGateways(true)
}

const testConnection = async (gatewayId: number) => {
  try {
    testingConnection.value = gatewayId
    const response: any = await $fetch(`/api/payments/admin/gateway/${gatewayId}/test`)
    
    testResult.value = {
      success: response.success,
      message: response.message
    }
    showTestResult.value = true
  } catch (error) {
    console.error('خطا در تست اتصال:', error)
    testResult.value = {
      success: false,
      message: 'خطا در تست اتصال'
    }
    showTestResult.value = true
  } finally {
    testingConnection.value = null
  }
}

const closeTestResult = () => {
  showTestResult.value = false
}

const toggleStatus = async (gateway: any) => {
  try {
    const newStatus = gateway.status === 'active' ? 'inactive' : 'active'
    await $fetch(`/api/payment-gateways/${gateway.id}`, {
      method: 'PUT' as any,
      body: { status: newStatus }
    })
    
    // Update local state
    gateway.status = newStatus
  } catch (error) {
    console.error('خطا در تغییر وضعیت:', error)
  }
}

const router = useRouter()

const openAddModal = () => {
  // هدایت به صفحه افزودن درگاه جدید
  router.push('/admin/finance/payment-gateways/new')
}

const editGateway = (gateway: any) => {
  // هدایت به صفحه ویرایش درگاه
  window.location.href = `/admin/finance/payment-gateways/${gateway.id}`
}

const confirmDelete = (gateway: any) => {
  gatewayToDelete.value = gateway
  showDeleteModal.value = true
}

const deleteGateway = async () => {
  if (!gatewayToDelete.value) return

  try {
    deleting.value = true
    await $fetch(`/api/payment-gateways/${gatewayToDelete.value.id}`, {
      method: 'DELETE'
    })
    gateways.value = gateways.value.filter(gateway => gateway.id !== gatewayToDelete.value.id)
    showDeleteModal.value = false
    gatewayToDelete.value = null
    // Optionally, show a success message
  } catch (error) {
    console.error('خطا در حذف درگاه:', error)
    // Optionally, show an error message
  } finally {
    deleting.value = false
  }
}

const closeDeleteModal = () => {
  showDeleteModal.value = false
  gatewayToDelete.value = null
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'inactive':
      return 'bg-gray-100 text-gray-800'
    case 'maintenance':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'active':
      return 'فعال'
    case 'inactive':
      return 'غیرفعال'
    case 'maintenance':
      return 'تعمیر و نگهداری'
    default:
      return 'نامشخص'
  }
}

const getTypeText = (type: string) => {
  switch (type) {
    case 'iranian':
      return ''
    case 'international':
      return 'بین‌المللی'
    case 'digital':
      return 'دیجیتال'
    default:
      return ''
  }
}

const formatCurrency = (amount: number) => {
  if (!amount) return '0 تومان'
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const getGatewayIcon = (type: string, name: string = '') => {
  // استفاده از آیکون SVG به جای تصویر برای سرعت بیشتر
  return 'i-heroicons-credit-card'
}

// Lifecycle
onMounted(() => {
  fetchGateways()
})
</script>
