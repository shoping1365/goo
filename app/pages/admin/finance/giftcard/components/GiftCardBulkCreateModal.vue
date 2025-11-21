<template>
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
    <div class="bg-white rounded-2xl shadow-xl w-full max-w-3xl mx-4 max-h-[90vh] overflow-y-auto">
      <!-- Header -->
      <div class="flex items-center justify-between p-6 border-b border-gray-200">
        <h2 class="text-xl font-bold text-gray-900">ایجاد گروهی گیفت کارت‌ها</h2>
        <button 
          class="text-gray-400 hover:text-gray-600 transition-colors"
          @click="$emit('close')"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>
      
      <!-- Form -->
      <form class="p-6 space-y-6" @submit.prevent="handleSubmit">
        <!-- تنظیمات اصلی -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <!-- تعداد کارت‌ها -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد کارت‌ها</label>
            <input
              v-model="form.quantity"
              type="number"
              min="1"
              max="100"
              required
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: 10"
            />
          </div>
          
          <!-- مبلغ هر کارت -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ هر کارت (تومان)</label>
            <input
              v-model="form.amount"
              type="number"
              min="1000"
              step="1000"
              required
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: 500000"
            />
          </div>
          
          <!-- نوع کارت -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع کارت</label>
            <select
              v-model="form.type"
              required
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="">انتخاب کنید</option>
              <option value="digital">دیجیتال</option>
              <option value="physical">فیزیکی</option>
            </select>
          </div>
        </div>
        
        <!-- تاریخ انقضا -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ انقضا</label>
          <input
            v-model="form.expiresAt"
            type="date"
            :min="minDate"
            required
            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          />
        </div>
        
        <!-- پیشوند کد -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">پیشوند کد (اختیاری)</label>
          <input
            v-model="form.codePrefix"
            type="text"
            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: BULK-2024"
          />
        </div>
        
        <!-- تنظیمات اضافی -->
        <div class="border-t border-gray-200 pt-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات اضافی</h3>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="space-y-4">
              <!-- ارسال ایمیل -->
              <div class="flex items-center justify-between">
                <div>
                  <label class="text-sm font-medium text-gray-700">ارسال ایمیل</label>
                  <p class="text-xs text-gray-500">اطلاعات کارت‌ها به ایمیل ارسال شود</p>
                </div>
                <div class="flex items-center">
                  <input
                    v-model="form.sendEmail"
                    type="checkbox"
                    class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                  />
                </div>
              </div>
              
              <!-- فعال‌سازی فوری -->
              <div class="flex items-center justify-between">
                <div>
                  <label class="text-sm font-medium text-gray-700">فعال‌سازی فوری</label>
                  <p class="text-xs text-gray-500">کارت‌ها بلافاصله قابل استفاده باشند</p>
                </div>
                <div class="flex items-center">
                  <input
                    v-model="form.activateImmediately"
                    type="checkbox"
                    class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                  />
                </div>
              </div>
            </div>
            
            <div class="space-y-4">
              <!-- دانلود فایل -->
              <div class="flex items-center justify-between">
                <div>
                  <label class="text-sm font-medium text-gray-700">دانلود فایل</label>
                  <p class="text-xs text-gray-500">فایل CSV با کدهای کارت دانلود شود</p>
                </div>
                <div class="flex items-center">
                  <input
                    v-model="form.downloadFile"
                    type="checkbox"
                    class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                  />
                </div>
              </div>
              
              <!-- چاپ کارت‌ها -->
              <div class="flex items-center justify-between">
                <div>
                  <label class="text-sm font-medium text-gray-700">چاپ کارت‌ها</label>
                  <p class="text-xs text-gray-500">کارت‌های فیزیکی چاپ شوند</p>
                </div>
                <div class="flex items-center">
                  <input
                    v-model="form.printCards"
                    type="checkbox"
                    :disabled="form.type !== 'physical'"
                    class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500 disabled:opacity-50"
                  />
                </div>
              </div>
            </div>
          </div>
        </div>
        
        <!-- خلاصه -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="text-sm font-medium text-gray-900 mb-3">خلاصه عملیات</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-sm">
            <div class="space-y-2">
              <div class="flex justify-between">
                <span class="text-gray-600">تعداد کارت‌ها:</span>
                <span class="font-medium">{{ form.quantity || 0 }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">مبلغ هر کارت:</span>
                <span class="font-medium">{{ formatPrice(form.amount) }}</span>
              </div>
            </div>
            <div class="space-y-2">
              <div class="flex justify-between">
                <span class="text-gray-600">کل مبلغ:</span>
                <span class="font-medium text-green-600">{{ formatPrice(totalAmount) }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">نوع:</span>
                <span class="font-medium">{{ getTypeText(form.type) }}</span>
              </div>
            </div>
          </div>
        </div>
        
        <!-- Actions -->
        <div class="flex items-center justify-end space-x-3 space-x-reverse pt-6 border-t border-gray-200">
          <button
            type="button"
            class="px-6 py-3 text-gray-700 bg-gray-100 hover:bg-gray-200 rounded-lg transition-colors"
            @click="$emit('close')"
          >
            انصراف
          </button>
          <button
            type="submit"
            :disabled="isSubmitting"
            class="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors flex items-center space-x-2 space-x-reverse"
          >
            <svg v-if="isSubmitting" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <span>{{ isSubmitting ? 'در حال ایجاد...' : 'ایجاد کارت‌ها' }}</span>
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

interface GiftCard {
  id?: number | string
  [key: string]: unknown
}

// تعریف events
const emit = defineEmits<{
  close: []
  created: [giftCards: GiftCard[]]
}>()

// Reactive data
const isSubmitting = ref(false)
const form = ref({
  quantity: '',
  amount: '',
  type: '',
  expiresAt: '',
  codePrefix: '',
  sendEmail: false,
  activateImmediately: true,
  downloadFile: true,
  printCards: false
})

// Computed properties
const minDate = computed(() => {
  const today = new Date()
  return today.toISOString().split('T')[0]
})

const totalAmount = computed(() => {
  return String(Number(form.value.quantity || 0) * Number(form.value.amount || 0))
})

// Methods
const formatPrice = (price: string): string => {
  if (!price) return '0 تومان'
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(Number(price))
}

const getTypeText = (type: string): string => {
  const typeMap = {
    digital: 'دیجیتال',
    physical: 'فیزیکی'
  }
  return typeMap[type] || 'انتخاب نشده'
}

const generateGiftCardCode = (index: number): string => {
  const prefix = form.value.codePrefix || 'BULK'
  const year = new Date().getFullYear()
  const random = Math.random().toString(36).substr(2, 4).toUpperCase()
  return `${prefix}-${year}-${String(index + 1).padStart(3, '0')}-${random}`
}

const handleSubmit = async () => {
  if (isSubmitting.value) return
  
  isSubmitting.value = true
  
  try {
    // شبیه‌سازی API call
    await new Promise(resolve => setTimeout(resolve, 3000))
    
    const giftCards = []
    const quantity = Number(form.value.quantity)
    
    for (let i = 0; i < quantity; i++) {
      const giftCard = {
        id: Date.now() + i,
        code: generateGiftCardCode(i),
        amount: Number(form.value.amount),
        remainingAmount: Number(form.value.amount),
        status: form.value.activateImmediately ? 'active' : 'inactive',
        type: form.value.type,
        createdBy: 'مدیر سیستم',
        recipient: '',
        recipientEmail: '',
        createdAt: new Date(),
        expiresAt: new Date(form.value.expiresAt),
        isUsed: false,
        usageHistory: []
      }
      giftCards.push(giftCard)
    }
    
    emit('created', giftCards)
  } catch (error) {
    console.error('خطا در ایجاد گروهی گیفت کارت‌ها:', error)
  } finally {
    isSubmitting.value = false
  }
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
