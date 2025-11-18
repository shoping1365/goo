<template>
  <div class="space-y-6">
    <!-- اطلاعات حمل و نقل -->
    <div class="px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-blue-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4"></path>
        </svg>
        اطلاعات حمل و نقل
      </h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- روش ارسال -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">روش ارسال</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">نوع ارسال</label>
              <select 
                v-model="shippingInfo.method"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="post">پست</option>
                <option value="express">پست پیشتاز</option>
                <option value="courier">پیک موتوری</option>
                <option value="pickup">دریافت حضوری</option>
                <option value="truck">حمل با کامیون</option>
                <option value="air">حمل هوایی</option>
                <option value="sea">حمل دریایی</option>
              </select>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">شرکت حمل و نقل</label>
              <select 
                v-model="shippingInfo.company"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">انتخاب شرکت</option>
                <option value="iran_post">پست جمهوری اسلامی ایران</option>
                <option value="tipax">تیپاکس</option>
                <option value="chapar">چاپار</option>
                <option value="delivery">دیلیوری</option>
                <option value="custom">سفارشی</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">هزینه ارسال (تومان)</label>
              <input 
                v-model="shippingInfo.cost"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
              />
            </div>
          </div>
        </div>

        <!-- اطلاعات پیگیری -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">اطلاعات پیگیری</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">شماره پیگیری</label>
              <input 
                v-model="shippingInfo.trackingNumber"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="شماره پیگیری ارسال"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">تاریخ ارسال</label>
              <input 
                v-model="shippingInfo.shippingDate"
                type="date" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">تاریخ تحویل تخمینی</label>
              <input 
                v-model="shippingInfo.estimatedDelivery"
                type="date" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- اطلاعات ارسال سفارش -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-indigo-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
        </svg>
        اطلاعات ارسال سفارش
      </h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">تاریخ‌های مهم</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">تاریخ ارسال</label>
              <input 
                v-model="shippingInfo.shippingDate"
                type="datetime-local" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">تاریخ تحویل تخمینی</label>
              <input 
                v-model="shippingInfo.estimatedDelivery"
                type="datetime-local" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">تاریخ تحویل واقعی</label>
              <input 
                v-model="shippingInfo.actualDelivery"
                type="datetime-local" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
        </div>

        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">وضعیت تحویل</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت تحویل</label>
              <select 
                v-model="shippingInfo.deliveryStatus"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="pending">در انتظار ارسال</option>
                <option value="shipped">ارسال شده</option>
                <option value="in_transit">در حال حمل</option>
                <option value="out_for_delivery">آماده تحویل</option>
                <option value="delivered">تحویل داده شده</option>
                <option value="failed">تحویل ناموفق</option>
                <option value="returned">بازگشت</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">یادداشت تحویل</label>
              <textarea 
                v-model="shippingInfo.deliveryNotes"
                rows="2"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="یادداشت‌های مربوط به تحویل..."
              ></textarea>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- اطلاعات بسته‌بندی -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-green-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
        </svg>
        اطلاعات بسته‌بندی
      </h3>
      
      <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-2">وزن کل (کیلوگرم)</label>
          <input 
            v-model="shippingInfo.weight"
            type="number" 
            step="0.1"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="0.0"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-2">ابعاد (سانتی‌متر)</label>
          <div class="grid grid-cols-3 gap-2">
            <input 
              v-model="shippingInfo.dimensions.length"
              type="number" 
              class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="طول"
            />
            <input 
              v-model="shippingInfo.dimensions.width"
              type="number" 
              class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="عرض"
            />
            <input 
              v-model="shippingInfo.dimensions.height"
              type="number" 
              class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="ارتفاع"
            />
          </div>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-2">تعداد بسته</label>
          <input 
            v-model="shippingInfo.packageCount"
            type="number" 
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="1"
          />
        </div>
      </div>
    </div>

    <!-- وضعیت ارسال -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-purple-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        وضعیت ارسال
      </h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت فعلی ارسال</label>
          <select 
            v-model="shippingInfo.status"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="pending">در انتظار ارسال</option>
            <option value="packaged">بسته‌بندی شده</option>
            <option value="shipped">ارسال شده</option>
            <option value="in_transit">در حال حمل</option>
            <option value="out_for_delivery">آماده تحویل</option>
            <option value="delivered">تحویل داده شده</option>
            <option value="returned">بازگشت</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-2">تاریخ آخرین به‌روزرسانی</label>
          <input 
            v-model="shippingInfo.lastUpdate"
            type="datetime-local" 
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
      </div>

      <!-- یادداشت ارسال -->
      <div class="mt-4">
        <label class="block text-sm font-medium text-gray-600 mb-2">یادداشت ارسال</label>
        <textarea 
          v-model="shippingInfo.notes"
          rows="3"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          placeholder="یادداشت‌های مربوط به ارسال..."
        ></textarea>
      </div>
    </div>

    <!-- تاریخچه ارسال -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-orange-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        تاریخچه ارسال
      </h3>
      
      <div class="space-y-3">
        <div 
          v-for="(history, index) in shippingInfo.history" 
          :key="index"
          class="flex items-center p-3 bg-gray-50 rounded-lg"
        >
          <div class="w-3 h-3 rounded-full mr-3" :class="getShippingStatusColor(history.status)"></div>
          <div class="flex-1">
            <div class="text-sm font-medium text-gray-900">{{ getShippingStatusText(history.status) }}</div>
            <div class="text-xs text-gray-500">{{ formatDate(history.date) }} - {{ history.location }}</div>
          </div>
          <div class="text-xs text-gray-500">{{ history.description }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

// تعریف props
const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({
      method: 'post',
      company: '',
      cost: 0,
      trackingNumber: '',
      shippingDate: '',
      estimatedDelivery: '',
      actualDelivery: '',
      deliveryStatus: 'pending',
      deliveryNotes: '',
      weight: 0,
      dimensions: {
        length: 0,
        width: 0,
        height: 0
      },
      packageCount: 1,
      status: 'pending',
      lastUpdate: '',
      notes: '',
      history: []
    })
  }
})

// تعریف emit
const emit = defineEmits(['update:modelValue'])

// متغیر محلی
const shippingInfo = ref({ ...props.modelValue })

// توابع کمکی
const getShippingStatusText = (status) => {
  const statusMap = {
    'pending': 'در انتظار ارسال',
    'packaged': 'بسته‌بندی شده',
    'shipped': 'ارسال شده',
    'in_transit': 'در حال حمل',
    'out_for_delivery': 'آماده تحویل',
    'delivered': 'تحویل داده شده',
    'returned': 'بازگشت'
  }
  return statusMap[status] || status
}

const getShippingStatusColor = (status) => {
  const colors = {
    'pending': 'bg-yellow-400',
    'packaged': 'bg-blue-400',
    'shipped': 'bg-purple-400',
    'in_transit': 'bg-indigo-400',
    'out_for_delivery': 'bg-orange-400',
    'delivered': 'bg-green-400',
    'returned': 'bg-red-400'
  }
  return colors[status] || 'bg-gray-400'
}

const formatDate = (date) => {
  if (!date) return 'نامشخص'
  return new Date(date).toLocaleDateString('fa-IR')
}

// نظارت بر تغییرات
watch(shippingInfo, (newValue) => {
  emit('update:modelValue', newValue)
}, { deep: true })
</script> 
