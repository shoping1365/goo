<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <!-- هدر -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6 mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">تنظیمات معافیت‌های مالیاتی</h3>
        <p class="text-sm text-gray-600">مدیریت معافیت‌ها، تخفیف‌ها و کدهای تخفیف مالیاتی</p>
      </div>
      
      <!-- دکمه افزودن معافیت جدید -->
      <button 
        class="inline-flex items-center gap-2 px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg transition-colors duration-200"
        @click="showAddExemptionModal = true"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
        </svg>
        افزودن معافیت جدید
      </button>
    </div>

    <!-- کارت‌های آمار معافیت‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-gradient-to-br from-green-50 to-green-100 rounded-xl p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ totalExemptions }}</div>
            <div class="text-sm text-green-700">کل معافیت‌ها</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-xl p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ activeExemptions }}</div>
            <div class="text-sm text-blue-700">معافیت‌های فعال</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-br from-purple-50 to-purple-100 rounded-xl p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-purple-600">{{ totalDiscountCodes }}</div>
            <div class="text-sm text-purple-700">کدهای تخفیف</div>
          </div>
          <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 5v2m0 4v2m0 4v2M5 5a2 2 0 00-2 2v3a2 2 0 110 4v3a2 2 0 002 2h14a2 2 0 002-2v-3a2 2 0 110-4V7a2 2 0 00-2-2H5z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-br from-orange-50 to-orange-100 rounded-xl p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-orange-600">{{ formatCurrency(totalSavings) }}</div>
            <div class="text-sm text-orange-700">کل صرفه‌جویی</div>
          </div>
          <div class="w-10 h-10 bg-orange-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- تب‌های مختلف -->
    <div class="mb-6">
      <div class="border-b border-gray-200">
        <nav class="-mb-px flex space-x-8">
          <button 
            v-for="tab in tabs" 
            :key="tab.id"
            :class="[
              'py-2 px-1 border-b-2 font-medium text-sm transition-colors duration-200',
              activeTab === tab.id 
                ? 'border-blue-500 text-blue-600' 
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
            @click="activeTab = tab.id"
          >
            {{ tab.label }}
          </button>
        </nav>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div v-if="activeTab === 'exemptions'">
      <!-- جدول معافیت‌ها -->
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-200 bg-gray-50">
              <th class="text-right py-3 px-4 font-medium text-gray-600">نوع معافیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">دسته‌بندی</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">درصد معافیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">شرایط</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="exemption in exemptions" :key="exemption.id" class="border-b border-gray-100 hover:bg-gray-50">
              <td class="py-3 px-4">
                <div class="flex items-center gap-2">
                  <div class="w-8 h-8 rounded-lg flex items-center justify-center" :class="getExemptionTypeColor(exemption.type)">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    </svg>
                  </div>
                  <span class="font-medium text-gray-900">{{ getExemptionTypeLabel(exemption.type) }}</span>
                </div>
              </td>
              <td class="py-3 px-4 text-gray-900">{{ exemption.category }}</td>
              <td class="py-3 px-4">
                <span class="font-bold text-green-600">{{ exemption.percentage }}%</span>
              </td>
              <td class="py-3 px-4 text-gray-600">{{ exemption.conditions }}</td>
              <td class="py-3 px-4">
                <span 
                  :class="[
                    'px-2 py-1 rounded-full text-xs font-medium',
                    exemption.status === 'active' ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'
                  ]"
                >
                  {{ exemption.status === 'active' ? 'فعال' : 'غیرفعال' }}
                </span>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center gap-2">
                  <button 
                    class="p-1 text-blue-600 hover:text-blue-800 transition-colors"
                    @click="editExemption(exemption)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                    </svg>
                  </button>
                  <button 
                    :class="[
                      'p-1 transition-colors',
                      exemption.status === 'active' ? 'text-red-600 hover:text-red-800' : 'text-green-600 hover:text-green-800'
                    ]"
                    @click="toggleExemptionStatus(exemption)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="exemption.status === 'active' ? 'M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728' : 'M5 13l4 4L19 7'"></path>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <div v-else-if="activeTab === 'discounts'">
      <!-- جدول کدهای تخفیف -->
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-200 bg-gray-50">
              <th class="text-right py-3 px-4 font-medium text-gray-600">کد تخفیف</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">نوع تخفیف</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">مقدار</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">تاریخ انقضا</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">استفاده شده</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="discount in discountCodes" :key="discount.id" class="border-b border-gray-100 hover:bg-gray-50">
              <td class="py-3 px-4">
                <div class="flex items-center gap-2">
                  <span class="font-mono font-bold text-blue-600">{{ discount.code }}</span>
                  <button 
                    class="p-1 text-gray-400 hover:text-gray-600"
                    @click="copyToClipboard(discount.code)"
                  >
                    <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                    </svg>
                  </button>
                </div>
              </td>
              <td class="py-3 px-4 text-gray-900">{{ getDiscountTypeLabel(discount.type) }}</td>
              <td class="py-3 px-4">
                <span class="font-bold text-green-600">
                  {{ discount.type === 'percentage' ? `${discount.value}%` : formatCurrency(discount.value) }}
                </span>
              </td>
              <td class="py-3 px-4 text-gray-600">{{ formatDate(discount.expiryDate) }}</td>
              <td class="py-3 px-4 text-gray-600">{{ discount.usedCount }} / {{ discount.maxUses || 'نامحدود' }}</td>
              <td class="py-3 px-4">
                <span 
                  :class="[
                    'px-2 py-1 rounded-full text-xs font-medium',
                    discount.status === 'active' ? 'bg-green-100 text-green-700' : 
                    discount.status === 'expired' ? 'bg-red-100 text-red-700' : 
                    'bg-yellow-100 text-yellow-700'
                  ]"
                >
                  {{ getDiscountStatusLabel(discount.status) }}
                </span>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center gap-2">
                  <button 
                    class="p-1 text-blue-600 hover:text-blue-800 transition-colors"
                    @click="editDiscount(discount)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                    </svg>
                  </button>
                  <button 
                    class="p-1 text-red-600 hover:text-red-800 transition-colors"
                    @click="deleteDiscount(discount)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <div v-else-if="activeTab === 'settings'">
      <!-- تنظیمات عمومی معافیت‌ها -->
      <div class="space-y-6">
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="text-md font-semibold text-gray-900 mb-4">تنظیمات عمومی</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر درصد معافیت</label>
              <input 
                v-model.number="settings.maxExemptionPercentage"
                type="number"
                min="0"
                max="100"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ سفارش برای معافیت</label>
              <input 
                v-model.number="settings.minOrderAmount"
                type="number"
                min="0"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
        </div>

        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="text-md font-semibold text-gray-900 mb-4">تنظیمات کدهای تخفیف</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">طول پیش‌فرض کد تخفیف</label>
              <input 
                v-model.number="settings.defaultCodeLength"
                type="number"
                min="4"
                max="20"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">مدت اعتبار پیش‌فرض (روز)</label>
              <input 
                v-model.number="settings.defaultValidityDays"
                type="number"
                min="1"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
        </div>

        <div class="flex justify-end">
          <button 
            class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors duration-200"
            @click="saveSettings"
          >
            ذخیره تنظیمات
          </button>
        </div>
      </div>
    </div>

    <!-- مودال افزودن/ویرایش معافیت -->
    <div v-if="showAddExemptionModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl p-6 w-full max-w-md mx-4">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">
            {{ editingExemption ? 'ویرایش معافیت' : 'افزودن معافیت جدید' }}
          </h3>
          <button class="text-gray-400 hover:text-gray-600" @click="closeModal">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <form class="space-y-4" @submit.prevent="saveExemption">
          <!-- نوع معافیت -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع معافیت</label>
            <select 
              v-model="exemptionForm.type"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="">انتخاب کنید</option>
              <option value="full">معافیت کامل</option>
              <option value="partial">معافیت جزئی</option>
              <option value="conditional">معافیت مشروط</option>
            </select>
          </div>

          <!-- دسته‌بندی -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
            <select 
              v-model="exemptionForm.category"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="">انتخاب کنید</option>
              <option value="products">محصولات</option>
              <option value="services">خدمات</option>
              <option value="digital">محصولات دیجیتال</option>
              <option value="food">مواد غذایی</option>
            </select>
          </div>

          <!-- درصد معافیت -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">درصد معافیت</label>
            <input 
              v-model.number="exemptionForm.percentage"
              type="number"
              step="0.1"
              min="0"
              max="100"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>

          <!-- شرایط -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">شرایط</label>
            <textarea 
              v-model="exemptionForm.conditions"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            ></textarea>
          </div>

          <!-- دکمه‌ها -->
          <div class="flex gap-3 pt-4">
            <button 
              type="submit"
              class="flex-1 px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg transition-colors duration-200"
            >
              {{ editingExemption ? 'ویرایش' : 'افزودن' }}
            </button>
            <button 
              type="button"
              class="flex-1 px-4 py-2 bg-gray-200 hover:bg-gray-300 text-gray-700 rounded-lg transition-colors duration-200"
              @click="closeModal"
            >
              انصراف
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'

// تب‌های مختلف
const tabs = [
  { id: 'exemptions', label: 'معافیت‌ها' },
  { id: 'discounts', label: 'کدهای تخفیف' },
  { id: 'settings', label: 'تنظیمات' }
]

const activeTab = ref('exemptions')

// معافیت‌ها
const exemptions = ref([
  {
    id: 1,
    type: 'full',
    category: 'مواد غذایی',
    percentage: 100,
    conditions: 'معافیت کامل برای مواد غذایی اساسی',
    status: 'active'
  },
  {
    id: 2,
    type: 'partial',
    category: 'محصولات دیجیتال',
    percentage: 50,
    conditions: 'معافیت ۵۰ درصدی برای محصولات دیجیتال آموزشی',
    status: 'active'
  },
  {
    id: 3,
    type: 'conditional',
    category: 'خدمات',
    percentage: 25,
    conditions: 'معافیت مشروط برای خدمات درمانی',
    status: 'inactive'
  }
])

// کدهای تخفیف
const discountCodes = ref([
  {
    id: 1,
    code: 'TAX2024',
    type: 'percentage',
    value: 10,
    expiryDate: '2024-12-31',
    usedCount: 45,
    maxUses: 100,
    status: 'active'
  },
  {
    id: 2,
    code: 'SAVE50K',
    type: 'fixed',
    value: 50000,
    expiryDate: '2024-06-30',
    usedCount: 12,
    maxUses: 50,
    status: 'active'
  },
  {
    id: 3,
    code: 'WELCOME15',
    type: 'percentage',
    value: 15,
    expiryDate: '2024-03-15',
    usedCount: 100,
    maxUses: 100,
    status: 'expired'
  }
])

// تنظیمات
const settings = ref({
  maxExemptionPercentage: 100,
  minOrderAmount: 100000,
  defaultCodeLength: 8,
  defaultValidityDays: 30
})

// مودال
const showAddExemptionModal = ref(false)
const editingExemption = ref(null)

// فرم معافیت
const exemptionForm = ref({
  type: '',
  category: '',
  percentage: 0,
  conditions: ''
})

// آمار
const totalExemptions = computed(() => exemptions.value.length)
const activeExemptions = computed(() => exemptions.value.filter(e => e.status === 'active').length)
const totalDiscountCodes = computed(() => discountCodes.value.length)
const totalSavings = computed(() => 125000000) // 125 میلیون تومان

// رنگ نوع معافیت
const getExemptionTypeColor = (type: string) => {
  const colors = {
    full: 'bg-green-500',
    partial: 'bg-blue-500',
    conditional: 'bg-orange-500'
  }
  return colors[type] || 'bg-gray-500'
}

// برچسب نوع معافیت
const getExemptionTypeLabel = (type: string) => {
  const labels = {
    full: 'معافیت کامل',
    partial: 'معافیت جزئی',
    conditional: 'معافیت مشروط'
  }
  return labels[type] || type
}

// برچسب نوع تخفیف
const getDiscountTypeLabel = (type: string) => {
  const labels = {
    percentage: 'درصدی',
    fixed: 'مبلغ ثابت'
  }
  return labels[type] || type
}

// برچسب وضعیت تخفیف
const getDiscountStatusLabel = (status: string) => {
  const labels = {
    active: 'فعال',
    expired: 'منقضی شده',
    used: 'استفاده شده'
  }
  return labels[status] || status
}

// فرمت مبلغ
const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

// فرمت تاریخ
const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('fa-IR')
}

// کپی کردن کد
const copyToClipboard = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text)
    // TODO: نمایش پیام موفقیت
  } catch (error) {
    console.error('خطا در کپی کردن:', error)
  }
}

interface TaxExemption {
  id?: number | string
  name?: string
  [key: string]: unknown
}

// ویرایش معافیت
const editExemption = (exemption: TaxExemption) => {
  editingExemption.value = exemption
  exemptionForm.value = { ...exemption }
  showAddExemptionModal.value = true
}

// تغییر وضعیت معافیت
const toggleExemptionStatus = async (exemption: TaxExemption) => {
  try {
    // TODO: ارسال درخواست به API
    exemption.status = exemption.status === 'active' ? 'inactive' : 'active'
  } catch (error) {
    console.error('خطا در تغییر وضعیت معافیت:', error)
  }
}

interface Discount {
  id?: number | string
  [key: string]: unknown
}

// ویرایش تخفیف
const editDiscount = (_discount: Discount) => {
  // TODO: پیاده‌سازی ویرایش تخفیف

}

// حذف تخفیف
const deleteDiscount = async (discount: Discount) => {
  if (confirm('آیا از حذف این کد تخفیف اطمینان دارید؟')) {
    try {
      // TODO: ارسال درخواست به API
      const index = discountCodes.value.findIndex(d => d.id === discount.id)
      if (index > -1) {
        discountCodes.value.splice(index, 1)
      }
    } catch (error) {
      console.error('خطا در حذف تخفیف:', error)
    }
  }
}

// ذخیره معافیت
const saveExemption = async () => {
  try {
    if (editingExemption.value) {
      // ویرایش معافیت موجود
      const index = exemptions.value.findIndex(e => e.id === editingExemption.value.id)
      if (index > -1) {
        exemptions.value[index] = { ...editingExemption.value, ...exemptionForm.value }
      }
    } else {
      // افزودن معافیت جدید
      const newExemption = {
        id: Date.now(),
        ...exemptionForm.value,
        status: 'active'
      }
      exemptions.value.push(newExemption)
    }
    
    closeModal()
  } catch (error) {
    console.error('خطا در ذخیره معافیت:', error)
  }
}

// ذخیره تنظیمات
const saveSettings = async () => {
  try {
    // TODO: ارسال درخواست به API

  } catch (error) {
    console.error('خطا در ذخیره تنظیمات:', error)
  }
}

// بستن مودال
const closeModal = () => {
  showAddExemptionModal.value = false
  editingExemption.value = null
  exemptionForm.value = {
    type: '',
    category: '',
    percentage: 0,
    conditions: ''
  }
}

// بارگذاری اولیه
onMounted(() => {
  // TODO: بارگذاری داده‌ها از API
})
</script>

<!--
  کامپوننت تنظیمات معافیت‌های مالیاتی
  شامل:
  1. مدیریت معافیت‌های مالیاتی
  2. مدیریت کدهای تخفیف
  3. تنظیمات عمومی
  4. آمار و گزارش‌گیری
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
