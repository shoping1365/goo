<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center">
        <div>
          <h2 class="text-xl font-bold text-gray-900">جزئیات گیفت کارت</h2>
          <p class="text-gray-600 mt-1">مشاهده اطلاعات کامل کارت {{ card.code }}</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            @click="printCard"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z"></path>
            </svg>
            چاپ کارت
          </button>
          <button 
            class="px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
            @click="editCard"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
            </svg>
            ویرایش
          </button>
        </div>
      </div>
    </div>

    <!-- اطلاعات اصلی کارت -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- کارت اصلی -->
      <div class="lg:col-span-2">
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">نمای کارت</h3>
          
          <div 
            class="w-full h-48 rounded-lg flex flex-col items-center justify-center text-white font-bold text-2xl relative overflow-hidden"
            :style="{
              background: `linear-gradient(135deg, ${card.primaryColor || '#3B82F6'} 0%, ${card.secondaryColor || '#8B5CF6'} 100%)`,
              fontFamily: getFontFamily(card.font)
            }"
          >
            <!-- پس‌زمینه تزئینی -->
            <div class="absolute inset-0 opacity-10">
              <div class="absolute top-6 right-4 text-4xl">🎁</div>
              <div class="absolute bottom-4 left-4 text-4xl">💝</div>
            </div>
            
            <!-- محتوای کارت -->
            <div class="text-center z-10">
              <div class="text-3xl font-bold mb-2">{{ formatCurrency(card.amount) }}</div>
              <div class="text-lg opacity-90">{{ card.recipientName || 'گیرنده' }}</div>
              <div class="text-sm opacity-75 mt-2">{{ card.code }}</div>
            </div>
            
            <!-- وضعیت کارت -->
            <div class="absolute top-6 left-4">
              <span :class="getStatusClass(card.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                {{ getStatusLabel(card.status) }}
              </span>
            </div>
          </div>
          
          <!-- اطلاعات اضافی -->
          <div class="mt-4 grid grid-cols-2 gap-6">
            <div class="text-center">
              <div class="text-sm text-gray-500">مبلغ باقی‌مانده</div>
              <div class="text-lg font-bold text-gray-900">{{ formatCurrency(card.remainingAmount) }}</div>
            </div>
            <div class="text-center">
              <div class="text-sm text-gray-500">تعداد استفاده</div>
              <div class="text-lg font-bold text-gray-900">{{ card.usageCount || 0 }} / {{ getUsageLimitText(card.usageLimit) }}</div>
            </div>
          </div>
        </div>
      </div>

      <!-- اطلاعات کلی -->
      <div class="space-y-6">
        <!-- وضعیت کلی -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">وضعیت کلی</h3>
          
          <div class="space-y-3">
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-600">وضعیت:</span>
              <span :class="getStatusClass(card.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                {{ getStatusLabel(card.status) }}
              </span>
            </div>
            
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-600">نوع کارت:</span>
              <span class="text-sm font-medium text-gray-900">{{ getTypeLabel(card.type) }}</span>
            </div>
            
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-600">دسته‌بندی:</span>
              <span class="text-sm font-medium text-gray-900">{{ getCategoryLabel(card.category) }}</span>
            </div>
            
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-600">روش تحویل:</span>
              <span class="text-sm font-medium text-gray-900">{{ getDeliveryMethodLabel(card.deliveryMethod) }}</span>
            </div>
          </div>
        </div>

        <!-- اطلاعات مالی -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">اطلاعات مالی</h3>
          
          <div class="space-y-3">
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-600">مبلغ اصلی:</span>
              <span class="text-sm font-medium text-gray-900">{{ formatCurrency(card.amount) }}</span>
            </div>
            
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-600">مبلغ باقی‌مانده:</span>
              <span class="text-sm font-medium text-gray-900">{{ formatCurrency(card.remainingAmount) }}</span>
            </div>
            
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-600">مبلغ استفاده شده:</span>
              <span class="text-sm font-medium text-gray-900">{{ formatCurrency(card.amount - card.remainingAmount) }}</span>
            </div>
            
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-600">درصد استفاده:</span>
              <span class="text-sm font-medium text-gray-900">{{ getUsagePercentage() }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- اطلاعات گیرنده -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-medium text-gray-900 mb-4">اطلاعات گیرنده</h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نام گیرنده</label>
          <p class="text-sm text-gray-900">{{ card.recipientName || 'تعیین نشده' }}</p>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">ایمیل گیرنده</label>
          <p class="text-sm text-gray-900">{{ card.recipientEmail || 'تعیین نشده' }}</p>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">شماره موبایل</label>
          <p class="text-sm text-gray-900">{{ card.recipientPhone || 'تعیین نشده' }}</p>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ تحویل</label>
          <p class="text-sm text-gray-900">{{ formatDate(card.deliveryDate) || 'تعیین نشده' }}</p>
        </div>
        
        <div class="md:col-span-2">
          <label class="block text-sm font-medium text-gray-700 mb-2">پیام شخصی</label>
          <p class="text-sm text-gray-900 bg-gray-50 p-3 rounded-md">{{ card.personalMessage || 'پیام شخصی وجود ندارد' }}</p>
        </div>
      </div>
    </div>

    <!-- تنظیمات استفاده -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات استفاده</h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">محدودیت استفاده</label>
          <p class="text-sm text-gray-900">{{ getUsageLimitText(card.usageLimit) }}</p>
        </div>
        
        <div v-if="card.usageLimit === 'multiple'">
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تعداد استفاده</label>
          <p class="text-sm text-gray-900">{{ card.maxUsageCount || 'نامحدود' }}</p>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ سفارش</label>
          <p class="text-sm text-gray-900">{{ card.minOrderAmount ? formatCurrency(card.minOrderAmount) : 'بدون محدودیت' }}</p>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر مبلغ سفارش</label>
          <p class="text-sm text-gray-900">{{ card.maxOrderAmount ? formatCurrency(card.maxOrderAmount) : 'بدون محدودیت' }}</p>
        </div>
        
        <div class="md:col-span-2">
          <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی‌های مجاز</label>
          <div class="flex flex-wrap gap-2">
            <span 
              v-for="category in card.allowedCategories" 
              :key="category"
              class="inline-flex px-2 py-1 text-xs font-medium bg-blue-100 text-blue-800 rounded-full"
            >
              {{ getCategoryLabel(category) }}
            </span>
            <span v-if="!card.allowedCategories || card.allowedCategories.length === 0" class="text-sm text-gray-500">
              همه دسته‌بندی‌ها مجاز
            </span>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات امنیتی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات امنیتی</h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="flex items-center">
          <input
            :checked="card.requiresApproval"
            type="checkbox"
            disabled
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <label class="mr-2 text-sm text-gray-700">نیاز به تأیید قبل از استفاده</label>
        </div>
        
        <div class="flex items-center">
          <input
            :checked="card.isActive"
            type="checkbox"
            disabled
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <label class="mr-2 text-sm text-gray-700">کارت فعال</label>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کد امنیتی</label>
          <p class="text-sm text-gray-900 font-mono">{{ card.securityCode || 'تعیین نشده' }}</p>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">آخرین بررسی امنیتی</label>
          <p class="text-sm text-gray-900">{{ formatDate(card.lastSecurityCheck) || 'انجام نشده' }}</p>
        </div>
      </div>
    </div>

    <!-- تاریخچه استفاده -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-medium text-gray-900 mb-4">تاریخچه استفاده</h3>
      
      <div v-if="card.usageHistory && card.usageHistory.length > 0" class="space-y-4">
        <div 
          v-for="usage in card.usageHistory" 
          :key="usage.id"
          class="border border-gray-200 rounded-lg p-6"
        >
          <div class="flex justify-between items-start">
            <div>
              <div class="text-sm font-medium text-gray-900">{{ usage.orderNumber }}</div>
              <div class="text-sm text-gray-500">{{ usage.userName }} ({{ usage.userEmail }})</div>
              <div class="text-sm text-gray-500">{{ formatDate(usage.usedAt) }}</div>
            </div>
            <div class="text-right">
              <div class="text-sm font-medium text-gray-900">{{ formatCurrency(usage.amount) }}</div>
              <span :class="getUsageStatusClass(usage.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                {{ getUsageStatusLabel(usage.status) }}
              </span>
            </div>
          </div>
        </div>
      </div>
      
      <div v-else class="text-center py-8">
        <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
        </svg>
        <h3 class="mt-2 text-sm font-medium text-gray-900">تاریخچه‌ای وجود ندارد</h3>
        <p class="mt-1 text-sm text-gray-500">این کارت هنوز استفاده نشده است.</p>
      </div>
    </div>

    <!-- اطلاعات زمانی -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">تاریخ‌های مهم</h3>
        
        <div class="space-y-3">
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">تاریخ ایجاد:</span>
            <span class="text-sm font-medium text-gray-900">{{ formatDate(card.createdAt) }}</span>
          </div>
          
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">تاریخ انقضا:</span>
            <span class="text-sm font-medium text-gray-900">{{ formatDate(card.expiresAt) }}</span>
          </div>
          
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">تاریخ تحویل:</span>
            <span class="text-sm font-medium text-gray-900">{{ formatDate(card.deliveryDate) }}</span>
          </div>
          
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">آخرین استفاده:</span>
            <span class="text-sm font-medium text-gray-900">{{ formatDate(card.lastUsedAt) || 'استفاده نشده' }}</span>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">آمار استفاده</h3>
        
        <div class="space-y-3">
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">تعداد استفاده:</span>
            <span class="text-sm font-medium text-gray-900">{{ card.usageCount || 0 }}</span>
          </div>
          
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">میانگین مبلغ استفاده:</span>
            <span class="text-sm font-medium text-gray-900">{{ formatCurrency(card.avgUsageAmount || 0) }}</span>
          </div>
          
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">بیشترین مبلغ استفاده:</span>
            <span class="text-sm font-medium text-gray-900">{{ formatCurrency(card.maxUsageAmount || 0) }}</span>
          </div>
          
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">آخرین مبلغ استفاده:</span>
            <span class="text-sm font-medium text-gray-900">{{ formatCurrency(card.lastUsageAmount || 0) }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

// Props
interface GiftCard {
  id?: number | string
  [key: string]: unknown
}

const props = defineProps<{
  card: GiftCard
}>()

// Emits
const emit = defineEmits<{
  'edit-card': [card: GiftCard]
}>()

// Computed properties
const card = computed(() => props.card)

// Methods
const printCard = () => {

  // اینجا می‌توانید منطق چاپ اضافه کنید
}

const editCard = () => {
  emit('edit-card', card.value)
}

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const formatDate = (dateString: string) => {
  if (!dateString) return ''
  return new Date(dateString).toLocaleDateString('fa-IR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'used':
      return 'bg-blue-100 text-blue-800'
    case 'expired':
      return 'bg-red-100 text-red-800'
    case 'suspended':
      return 'bg-yellow-100 text-yellow-800'
    case 'cancelled':
      return 'bg-gray-100 text-gray-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusLabel = (status: string) => {
  switch (status) {
    case 'active':
      return 'فعال'
    case 'used':
      return 'استفاده شده'
    case 'expired':
      return 'منقضی شده'
    case 'suspended':
      return 'معلق'
    case 'cancelled':
      return 'لغو شده'
    default:
      return 'نامشخص'
  }
}

const getTypeLabel = (type: string) => {
  switch (type) {
    case 'digital':
      return 'دیجیتال'
    case 'physical':
      return 'فیزیکی'
    case 'virtual':
      return 'مجازی'
    default:
      return 'نامشخص'
  }
}

const getCategoryLabel = (category: string) => {
  const categories = {
    'birthday': 'تولد',
    'wedding': 'عروسی',
    'newyear': 'سال نو',
    'christmas': 'کریسمس',
    'graduation': 'فارغ‌التحصیلی',
    'anniversary': 'سالگرد',
    'general': 'عمومی',
    'electronics': 'الکترونیک',
    'clothing': 'پوشاک',
    'books': 'کتاب',
    'home': 'خانه و آشپزخانه',
    'sports': 'ورزشی',
    'beauty': 'زیبایی',
    'toys': 'اسباب بازی',
    'food': 'مواد غذایی'
  }
  return categories[category] || category
}

const getDeliveryMethodLabel = (method: string) => {
  switch (method) {
    case 'email':
      return 'ایمیل'
    case 'sms':
      return 'پیامک'
    case 'both':
      return 'هر دو'
    case 'manual':
      return 'دستی'
    default:
      return 'نامشخص'
  }
}

const getUsageLimitText = (limit: string) => {
  switch (limit) {
    case 'unlimited':
      return 'نامحدود'
    case 'once':
      return 'یک بار'
    case 'multiple':
      return 'چند بار'
    default:
      return 'نامشخص'
  }
}

const getUsagePercentage = () => {
  if (!card.value.amount || !card.value.remainingAmount) return 0
  const used = card.value.amount - card.value.remainingAmount
  return Math.round((used / card.value.amount) * 100)
}

const getFontFamily = (font: string) => {
  const fontMap = {
    'IRANSans': 'IRANSans, sans-serif',
    'Vazir': 'Vazir, sans-serif',
    'Shabnam': 'Shabnam, sans-serif',
    'Yekan': 'Yekan, sans-serif'
  }
  return fontMap[font] || 'IRANSans, sans-serif'
}

const getUsageStatusClass = (status: string) => {
  switch (status) {
    case 'success':
      return 'bg-green-100 text-green-800'
    case 'pending':
      return 'bg-yellow-100 text-yellow-800'
    case 'failed':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getUsageStatusLabel = (status: string) => {
  switch (status) {
    case 'success':
      return 'موفق'
    case 'pending':
      return 'در انتظار'
    case 'failed':
      return 'ناموفق'
    default:
      return 'نامشخص'
  }
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
