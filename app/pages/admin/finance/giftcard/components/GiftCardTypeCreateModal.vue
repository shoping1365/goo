<template>
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
    <div class="bg-white rounded-2xl shadow-xl w-full max-w-2xl mx-4 max-h-[90vh] overflow-y-auto">
      <!-- Header -->
      <div class="flex items-center justify-between p-6 border-b border-gray-200">
        <h2 class="text-xl font-bold text-gray-900">ایجاد نوع جدید گیفت کارت</h2>
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
        <!-- اطلاعات اصلی -->
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام نوع کارت</label>
            <input
              v-model="form.name"
              type="text"
              required
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: کارت‌های تولد"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
            <textarea
              v-model="form.description"
              rows="3"
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent resize-none"
              placeholder="توضیحات مربوط به این نوع کارت..."
            ></textarea>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
            <select
              v-model="form.category"
              required
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="">انتخاب کنید</option>
              <option value="fixed">کارت‌های با مبلغ ثابت</option>
              <option value="variable">کارت‌های با مبلغ متغیر</option>
              <option value="thematic">کارت‌های موضوعی</option>
              <option value="corporate">کارت‌های شرکتی</option>
              <option value="discount">کارت‌های تخفیف ویژه</option>
            </select>
          </div>
        </div>
        
        <!-- تنظیمات مبلغ -->
        <div class="border-t border-gray-200 pt-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات مبلغ</h3>
          
          <div v-if="form.category === 'fixed'" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ ثابت (تومان)</label>
              <input
                v-model="form.fixedAmount"
                type="number"
                min="1000"
                step="1000"
                required
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="مثال: 500000"
              />
            </div>
          </div>
          
          <div v-else-if="form.category === 'variable'" class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ (تومان)</label>
              <input
                v-model="form.minAmount"
                type="number"
                min="1000"
                step="1000"
                required
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="مثال: 100000"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر مبلغ (تومان)</label>
              <input
                v-model="form.maxAmount"
                type="number"
                min="1000"
                step="1000"
                required
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="مثال: 5000000"
              />
            </div>
          </div>
          
          <div v-else-if="form.category === 'discount'" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">درصد تخفیف</label>
              <input
                v-model="form.discountPercentage"
                type="number"
                min="1"
                max="100"
                required
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="مثال: 20"
              />
            </div>
          </div>
          
          <div v-else-if="form.category === 'corporate'" class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ (تومان)</label>
              <input
                v-model="form.minAmount"
                type="number"
                min="100000"
                step="100000"
                required
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="مثال: 1000000"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر مبلغ (تومان)</label>
              <input
                v-model="form.maxAmount"
                type="number"
                min="100000"
                step="100000"
                required
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="مثال: 10000000"
              />
            </div>
          </div>
        </div>
        
        <!-- تنظیمات موضوعی -->
        <div v-if="form.category === 'thematic'" class="border-t border-gray-200 pt-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات موضوعی</h3>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">موضوعات</label>
            <div class="space-y-2">
              <div class="flex items-center">
                <input
                  v-model="form.themes"
                  type="checkbox"
                  value="تولد"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
                <label class="mr-2 text-sm text-gray-700">تولد</label>
              </div>
              <div class="flex items-center">
                <input
                  v-model="form.themes"
                  type="checkbox"
                  value="عروسی"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
                <label class="mr-2 text-sm text-gray-700">عروسی</label>
              </div>
              <div class="flex items-center">
                <input
                  v-model="form.themes"
                  type="checkbox"
                  value="سال نو"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
                <label class="mr-2 text-sm text-gray-700">سال نو</label>
              </div>
              <div class="flex items-center">
                <input
                  v-model="form.themes"
                  type="checkbox"
                  value="ولنتاین"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
                <label class="mr-2 text-sm text-gray-700">ولنتاین</label>
              </div>
              <div class="flex items-center">
                <input
                  v-model="form.themes"
                  type="checkbox"
                  value="مادر"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
                <label class="mr-2 text-sm text-gray-700">مادر</label>
              </div>
              <div class="flex items-center">
                <input
                  v-model="form.themes"
                  type="checkbox"
                  value="پدر"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
                <label class="mr-2 text-sm text-gray-700">پدر</label>
              </div>
            </div>
          </div>
        </div>
        
        <!-- تنظیمات اضافی -->
        <div class="border-t border-gray-200 pt-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات اضافی</h3>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <div>
                <label class="text-sm font-medium text-gray-700">فعال</label>
                <p class="text-xs text-gray-500">این نوع کارت قابل استفاده باشد</p>
              </div>
              <div class="flex items-center">
                <input
                  v-model="form.isActive"
                  type="checkbox"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
              </div>
            </div>
            
            <div class="flex items-center justify-between">
              <div>
                <label class="text-sm font-medium text-gray-700">نمایش در فروشگاه</label>
                <p class="text-xs text-gray-500">در صفحه فروشگاه نمایش داده شود</p>
              </div>
              <div class="flex items-center">
                <input
                  v-model="form.showInStore"
                  type="checkbox"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
              </div>
            </div>
            
            <div class="flex items-center justify-between">
              <div>
                <label class="text-sm font-medium text-gray-700">نیاز به تأیید</label>
                <p class="text-xs text-gray-500">برای استفاده نیاز به تأیید مدیر باشد</p>
              </div>
              <div class="flex items-center">
                <input
                  v-model="form.requiresApproval"
                  type="checkbox"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
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
            <span>{{ isSubmitting ? 'در حال ایجاد...' : 'ایجاد نوع جدید' }}</span>
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// تعریف interface برای نوع گیفت کارت
interface GiftCardType {
  id: string
  name: string
  description: string
  category: 'fixed' | 'variable' | 'thematic' | 'corporate' | 'discount'
  minAmount?: number
  maxAmount?: number
  fixedAmount?: number
  discountPercentage?: number
  themes?: string[]
  isActive: boolean
  showInStore: boolean
  requiresApproval: boolean
  createdAt: Date
}

const emit = defineEmits<{
  close: []
  created: [giftCardType: GiftCardType]
}>()

// Reactive data
const isSubmitting = ref(false)
const form = ref({
  name: '',
  description: '',
  category: '',
  minAmount: '',
  maxAmount: '',
  fixedAmount: '',
  discountPercentage: '',
  themes: [] as string[],
  isActive: true,
  showInStore: true,
  requiresApproval: false
})

// Methods
const handleSubmit = async () => {
  if (isSubmitting.value) return
  
  isSubmitting.value = true
  
  try {
    // شبیه‌سازی API call
    await new Promise(resolve => setTimeout(resolve, 2000))
    
    const giftCardType: GiftCardType = {
      id: `type-${Date.now()}`,
      name: form.value.name,
      description: form.value.description,
      category: form.value.category as 'fixed' | 'variable' | 'thematic' | 'corporate' | 'discount',
      minAmount: form.value.minAmount ? Number(form.value.minAmount) : undefined,
      maxAmount: form.value.maxAmount ? Number(form.value.maxAmount) : undefined,
      fixedAmount: form.value.fixedAmount ? Number(form.value.fixedAmount) : undefined,
      discountPercentage: form.value.discountPercentage ? Number(form.value.discountPercentage) : undefined,
      themes: form.value.themes.length > 0 ? form.value.themes : undefined,
      isActive: form.value.isActive,
      showInStore: form.value.showInStore,
      requiresApproval: form.value.requiresApproval,
      createdAt: new Date()
    }
    
    emit('created', giftCardType)
  } catch (error) {
    console.error('خطا در ایجاد نوع گیفت کارت:', error)
  } finally {
    isSubmitting.value = false
  }
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
