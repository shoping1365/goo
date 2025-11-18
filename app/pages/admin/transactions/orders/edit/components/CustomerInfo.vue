<template>
  <div class="px-4 py-4">
    <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
      <svg class="w-5 h-5 text-blue-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
      </svg>
      اطلاعات مشتری
    </h3>
    
    <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
      <!-- اطلاعات شخصی -->
      <div>
        <h4 class="text-md font-medium text-gray-700 mb-3">اطلاعات شخصی</h4>
        <div class="space-y-3">
          <div>
            <label class="block text-sm font-medium text-gray-600 mb-1">نام و نام خانوادگی</label>
            <input 
              v-model="customerInfo.name"
              type="text" 
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="نام مشتری"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-600 mb-1">شماره موبایل</label>
            <input 
              v-model="customerInfo.phone"
              type="tel" 
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="09123456789"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-600 mb-1">ایمیل</label>
            <input 
              v-model="customerInfo.email"
              type="email" 
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="customer@example.com"
            />
          </div>
        </div>
      </div>

      <!-- آدرس ارسال -->
      <div>
        <h4 class="text-md font-medium text-gray-700 mb-3">آدرس ارسال</h4>
        <div class="space-y-3">
          <div>
            <label class="block text-sm font-medium text-gray-600 mb-1">استان</label>
            <select 
              v-model="customerInfo.address.province"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="">انتخاب استان</option>
              <option value="tehran">تهران</option>
              <option value="isfahan">اصفهان</option>
              <option value="shiraz">شیراز</option>
              <option value="mashhad">مشهد</option>
              <option value="tabriz">تبریز</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-600 mb-1">شهر</label>
            <input 
              v-model="customerInfo.address.city"
              type="text" 
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="نام شهر"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-600 mb-1">آدرس کامل</label>
            <textarea 
              v-model="customerInfo.address.fullAddress"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="آدرس کامل پستی"
            ></textarea>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-600 mb-1">کد پستی</label>
            <input 
              v-model="customerInfo.address.postalCode"
              type="text" 
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="1234567890"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- اطلاعات اضافی -->
    <div class="mt-6 pt-6 border-t border-gray-200">
      <h4 class="text-md font-medium text-gray-700 mb-3">اطلاعات اضافی</h4>
      <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-1">تاریخ عضویت</label>
          <div class="text-sm text-gray-900 bg-gray-50 px-3 py-2 rounded-lg">
            {{ formatDate(customerInfo.joinDate) }}
          </div>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-1">تعداد سفارشات</label>
          <div class="text-sm text-gray-900 bg-gray-50 px-3 py-2 rounded-lg">
            {{ customerInfo.totalOrders }} سفارش
          </div>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-1">امتیاز مشتری</label>
          <div class="text-sm text-gray-900 bg-gray-50 px-3 py-2 rounded-lg">
            {{ customerInfo.rating }}/5
          </div>
        </div>
      </div>
    </div>

    <!-- سفارشات قبلی -->
    <div class="mt-6 pt-6 border-t border-gray-200">
      <h4 class="text-md font-medium text-gray-700 mb-4 flex items-center">
        <svg class="w-5 h-5 text-green-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
        </svg>
        سفارشات قبلی
      </h4>
      
      <div class="space-y-3">
        <!-- سفارش 1 -->
        <div class="bg-gray-100 rounded-lg px-4 py-4 hover:bg-gray-200 transition-colors">
            <div class="flex justify-between items-start">
              <div class="flex-1">
                <div class="flex items-center space-x-2 space-x-reverse mb-2">
                  <span class="text-sm font-medium text-gray-900">سفارش #ORD-2023-045</span>
                  <span class="px-2 py-1 text-xs bg-green-100 text-green-800 rounded-full">تکمیل شده</span>
                </div>
                <div class="text-sm text-gray-600 mb-2">
                  <span class="font-medium">تاریخ:</span> 15 دی 1402
                </div>
                <div class="text-sm text-gray-600 mb-2">
                  <span class="font-medium">مبلغ:</span> 2,450,000 تومان
                </div>
                <div class="text-sm text-gray-600">
                  <span class="font-medium">محصولات:</span> لپ‌تاپ اپل مک‌بوک پرو، گوشی سامسونگ گلکسی S23
                </div>
              </div>
              <button class="text-blue-600 hover:text-blue-800 text-sm font-medium">
                مشاهده جزئیات
              </button>
            </div>
          </div>

          <!-- سفارش 2 -->
          <div class="bg-gray-100 rounded-lg px-4 py-4 hover:bg-gray-200 transition-colors">
            <div class="flex justify-between items-start">
              <div class="flex-1">
                <div class="flex items-center space-x-2 space-x-reverse mb-2">
                  <span class="text-sm font-medium text-gray-900">سفارش #ORD-2023-032</span>
                  <span class="px-2 py-1 text-xs bg-blue-100 text-blue-800 rounded-full">در حال ارسال</span>
                </div>
                <div class="text-sm text-gray-600 mb-2">
                  <span class="font-medium">تاریخ:</span> 8 دی 1402
                </div>
                <div class="text-sm text-gray-600 mb-2">
                  <span class="font-medium">مبلغ:</span> 1,850,000 تومان
                </div>
                <div class="text-sm text-gray-600">
                  <span class="font-medium">محصولات:</span> تبلت اپل iPad Pro، قلم Apple Pencil
                </div>
              </div>
              <button class="text-blue-600 hover:text-blue-800 text-sm font-medium">
                مشاهده جزئیات
              </button>
            </div>
          </div>

          <!-- سفارش 3 -->
          <div class="bg-gray-100 rounded-lg px-4 py-4 hover:bg-gray-200 transition-colors">
            <div class="flex justify-between items-start">
              <div class="flex-1">
                <div class="flex items-center space-x-2 space-x-reverse mb-2">
                  <span class="text-sm font-medium text-gray-900">سفارش #ORD-2023-018</span>
                  <span class="px-2 py-1 text-xs bg-yellow-100 text-yellow-800 rounded-full">در حال پردازش</span>
                </div>
                <div class="text-sm text-gray-600 mb-2">
                  <span class="font-medium">تاریخ:</span> 1 دی 1402
                </div>
                <div class="text-sm text-gray-600 mb-2">
                  <span class="font-medium">مبلغ:</span> 3,200,000 تومان
                </div>
                <div class="text-sm text-gray-600">
                  <span class="font-medium">محصولات:</span> ساعت هوشمند اپل Apple Watch Series 9
                </div>
              </div>
              <button class="text-blue-600 hover:text-blue-800 text-sm font-medium">
                مشاهده جزئیات
              </button>
            </div>
          </div>
        </div>

        <!-- دکمه مشاهده همه -->
        <div class="mt-4 text-center">
          <button class="inline-flex items-center px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 transition-colors">
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
            </svg>
            مشاهده همه سفارشات
          </button>
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
      name: '',
      phone: '',
      email: '',
      address: {
        province: '',
        city: '',
        fullAddress: '',
        postalCode: ''
      },
      joinDate: '',
      totalOrders: 0,
      rating: 0
    })
  }
})

// تعریف emit
const emit = defineEmits(['update:modelValue'])

// متغیر محلی
const customerInfo = ref({ ...props.modelValue })

// نظارت بر تغییرات
watch(customerInfo, (newValue) => {
  emit('update:modelValue', newValue)
}, { deep: true })

// تابع فرمت تاریخ
const formatDate = (date) => {
  if (!date) return 'نامشخص'
  return new Date(date).toLocaleDateString('fa-IR')
}
</script> 
