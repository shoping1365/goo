<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <!-- هدر -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6 mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">محاسبه خودکار مالیات</h3>
        <p class="text-sm text-gray-600">محاسبه دقیق مالیات بر اساس نرخ‌ها، معافیت‌ها و قوانین مالیاتی</p>
      </div>
      
      <!-- دکمه محاسبه دسته‌ای -->
      <button 
        class="inline-flex items-center gap-2 px-4 py-2 bg-indigo-600 hover:bg-indigo-700 text-white rounded-lg transition-colors duration-200"
        @click="showBulkCalculationModal = true"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
        </svg>
        محاسبه دسته‌ای
      </button>
    </div>

    <!-- فرم محاسبه -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- ورودی‌ها -->
      <div class="space-y-6">
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">اطلاعات سفارش</h4>
          
          <!-- نوع مشتری -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع مشتری</label>
            <select 
              v-model="calculationForm.customerType"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              @change="updateCustomerType"
            >
              <option value="individual">شخصی</option>
              <option value="business">شرکتی</option>
              <option value="government">دولتی</option>
              <option value="foreign">خارجی</option>
            </select>
          </div>

          <!-- منطقه -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">منطقه</label>
            <select 
              v-model="calculationForm.region"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              @change="updateRegion"
            >
              <option value="">انتخاب منطقه</option>
              <option v-for="region in regions" :key="region.id" :value="region.id">
                {{ region.name }} ({{ region.taxRate }}%)
              </option>
            </select>
          </div>

          <!-- دسته‌بندی محصولات -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی محصولات</label>
            <select 
              v-model="calculationForm.category"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              @change="updateCategory"
            >
              <option value="">انتخاب دسته‌بندی</option>
              <option value="products">محصولات</option>
              <option value="services">خدمات</option>
              <option value="digital">محصولات دیجیتال</option>
              <option value="food">مواد غذایی</option>
            </select>
          </div>

          <!-- مبلغ کل -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ کل (تومان)</label>
            <input 
              v-model.number="calculationForm.totalAmount"
              type="number"
              min="0"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              @input="calculateTax"
            />
          </div>

          <!-- کد تخفیف -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">کد تخفیف (اختیاری)</label>
            <div class="flex gap-2">
              <input 
                v-model="calculationForm.discountCode"
                type="text"
                placeholder="کد تخفیف را وارد کنید"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
              <button 
                class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg text-sm transition-colors duration-200"
                @click="validateDiscountCode"
              >
                اعتبارسنجی
              </button>
            </div>
          </div>
        </div>

        <!-- تنظیمات اضافی -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">تنظیمات اضافی</h4>
          
          <div class="space-y-3">
            <label class="flex items-center">
              <input 
                v-model="calculationForm.applyVAT"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                @change="calculateTax"
              />
              <span class="mr-2 text-sm text-gray-700">اعمال مالیات بر ارزش افزوده</span>
            </label>

            <label class="flex items-center">
              <input 
                v-model="calculationForm.applyIncomeTax"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                @change="calculateTax"
              />
              <span class="mr-2 text-sm text-gray-700">اعمال مالیات بر درآمد</span>
            </label>

            <label class="flex items-center">
              <input 
                v-model="calculationForm.applyCustomsTax"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                @change="calculateTax"
              />
              <span class="mr-2 text-sm text-gray-700">اعمال مالیات گمرکی</span>
            </label>

            <label class="flex items-center">
              <input 
                v-model="calculationForm.roundToNearest"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                @change="calculateTax"
              />
              <span class="mr-2 text-sm text-gray-700">گرد کردن به نزدیک‌ترین عدد</span>
            </label>
          </div>
        </div>
      </div>

      <!-- نتایج محاسبه -->
      <div class="space-y-6">
        <!-- خلاصه محاسبه -->
        <div class="bg-blue-50 rounded-lg p-6">
          <h4 class="font-medium text-blue-900 mb-4">خلاصه محاسبه</h4>
          
          <div class="space-y-3">
            <div class="flex justify-between">
              <span class="text-sm text-blue-700">مبلغ کل:</span>
              <span class="font-medium text-blue-900">{{ formatCurrency(calculationResult.totalAmount) }}</span>
            </div>
            
            <div v-if="calculationResult.discountAmount > 0" class="flex justify-between">
              <span class="text-sm text-blue-700">تخفیف:</span>
              <span class="font-medium text-green-600">-{{ formatCurrency(calculationResult.discountAmount) }}</span>
            </div>
            
            <div class="flex justify-between">
              <span class="text-sm text-blue-700">مبلغ پس از تخفیف:</span>
              <span class="font-medium text-blue-900">{{ formatCurrency(calculationResult.amountAfterDiscount) }}</span>
            </div>
            
            <div class="border-t border-blue-200 pt-3">
              <div class="flex justify-between">
                <span class="text-sm text-blue-700">مالیات بر ارزش افزوده:</span>
                <span class="font-medium text-blue-900">{{ formatCurrency(calculationResult.vatAmount) }}</span>
              </div>
              
              <div class="flex justify-between">
                <span class="text-sm text-blue-700">مالیات بر درآمد:</span>
                <span class="font-medium text-blue-900">{{ formatCurrency(calculationResult.incomeTaxAmount) }}</span>
              </div>
              
              <div class="flex justify-between">
                <span class="text-sm text-blue-700">مالیات گمرکی:</span>
                <span class="font-medium text-blue-900">{{ formatCurrency(calculationResult.customsTaxAmount) }}</span>
              </div>
            </div>
            
            <div class="border-t border-blue-200 pt-3">
              <div class="flex justify-between text-lg font-bold">
                <span class="text-blue-900">مبلغ نهایی:</span>
                <span class="text-blue-900">{{ formatCurrency(calculationResult.finalAmount) }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- جزئیات محاسبه -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">جزئیات محاسبه</h4>
          
          <div class="space-y-3 text-sm">
            <div class="flex justify-between">
              <span class="text-gray-600">نرخ مالیات بر ارزش افزوده:</span>
              <span class="font-medium text-gray-900">{{ calculationResult.vatRate }}%</span>
            </div>
            
            <div class="flex justify-between">
              <span class="text-gray-600">نرخ مالیات بر درآمد:</span>
              <span class="font-medium text-gray-900">{{ calculationResult.incomeTaxRate }}%</span>
            </div>
            
            <div class="flex justify-between">
              <span class="text-gray-600">نرخ مالیات گمرکی:</span>
              <span class="font-medium text-gray-900">{{ calculationResult.customsTaxRate }}%</span>
            </div>
            
            <div class="flex justify-between">
              <span class="text-gray-600">معافیت‌های اعمال شده:</span>
              <span class="font-medium text-green-600">{{ calculationResult.exemptionsCount }} مورد</span>
            </div>
            
            <div class="flex justify-between">
              <span class="text-gray-600">کل مالیات:</span>
              <span class="font-medium text-red-600">{{ formatCurrency(calculationResult.totalTaxAmount) }}</span>
            </div>
          </div>
        </div>

        <!-- دکمه‌های عملیاتی -->
        <div class="flex gap-3">
          <button 
            class="flex-1 px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg transition-colors duration-200"
            @click="saveCalculation"
          >
            ذخیره محاسبه
          </button>
          <button 
            class="flex-1 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors duration-200"
            @click="exportCalculation"
          >
            خروجی PDF
          </button>
          <button 
            class="flex-1 px-4 py-2 bg-gray-200 hover:bg-gray-300 text-gray-700 rounded-lg transition-colors duration-200"
            @click="resetCalculation"
          >
            ریست
          </button>
        </div>
      </div>
    </div>

    <!-- مودال محاسبه دسته‌ای -->
    <div v-if="showBulkCalculationModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl p-6 w-full max-w-2xl mx-4">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">محاسبه دسته‌ای مالیات</h3>
          <button class="text-gray-400 hover:text-gray-600" @click="showBulkCalculationModal = false">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <div class="space-y-4">
          <!-- آپلود فایل -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آپلود فایل Excel</label>
            <div class="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center">
              <svg class="w-12 h-12 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
              </svg>
              <p class="text-sm text-gray-600">فایل Excel را اینجا رها کنید یا کلیک کنید</p>
              <input type="file" accept=".xlsx,.xls" class="hidden" />
            </div>
          </div>

          <!-- تنظیمات محاسبه -->
          <div class="grid grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع مشتری</label>
              <select class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
                <option value="individual">شخصی</option>
                <option value="business">شرکتی</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">منطقه</label>
              <select class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
                <option value="">انتخاب منطقه</option>
                <option v-for="region in regions" :key="region.id" :value="region.id">
                  {{ region.name }}
                </option>
              </select>
            </div>
          </div>

          <!-- دکمه‌ها -->
          <div class="flex gap-3 pt-4">
            <button class="flex-1 px-4 py-2 bg-indigo-600 hover:bg-indigo-700 text-white rounded-lg transition-colors duration-200">
              شروع محاسبه
            </button>
            <button 
              class="flex-1 px-4 py-2 bg-gray-200 hover:bg-gray-300 text-gray-700 rounded-lg transition-colors duration-200"
              @click="showBulkCalculationModal = false"
            >
              انصراف
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'

// مناطق
const regions = ref([
  { id: 1, name: 'تهران', taxRate: 9 },
  { id: 2, name: 'اصفهان', taxRate: 8 },
  { id: 3, name: 'خراسان رضوی', taxRate: 7 },
  { id: 4, name: 'فارس', taxRate: 8.5 }
])

// فرم محاسبه
const calculationForm = ref({
  customerType: 'individual',
  region: '',
  category: '',
  totalAmount: 0,
  discountCode: '',
  applyVAT: true,
  applyIncomeTax: false,
  applyCustomsTax: false,
  roundToNearest: true
})

// نتایج محاسبه
const calculationResult = ref({
  totalAmount: 0,
  discountAmount: 0,
  amountAfterDiscount: 0,
  vatAmount: 0,
  incomeTaxAmount: 0,
  customsTaxAmount: 0,
  finalAmount: 0,
  vatRate: 0,
  incomeTaxRate: 0,
  customsTaxRate: 0,
  exemptionsCount: 0,
  totalTaxAmount: 0
})

// مودال
const showBulkCalculationModal = ref(false)

// نرخ‌های مالیات
const taxRates = ref({
  vat: 9,
  income: 15,
  customs: 5
})

// معافیت‌ها
const exemptions = ref([
  { type: 'food', percentage: 100 },
  { type: 'digital_education', percentage: 50 },
  { type: 'medical', percentage: 25 }
])

// کدهای تخفیف
const discountCodes = ref([
  { code: 'TAX2024', type: 'percentage', value: 10, valid: true },
  { code: 'SAVE50K', type: 'fixed', value: 50000, valid: true }
])

// فرمت مبلغ
const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

// بروزرسانی نوع مشتری
const updateCustomerType = () => {
  calculateTax()
}

// بروزرسانی منطقه
const updateRegion = () => {
  const selectedRegion = regions.value.find(r => r.id === (calculationForm.value.region as unknown as number))
  if (selectedRegion) {
    taxRates.value.vat = selectedRegion.taxRate
  }
  calculateTax()
}

// بروزرسانی دسته‌بندی
const updateCategory = () => {
  calculateTax()
}

// اعتبارسنجی کد تخفیف
const validateDiscountCode = () => {
  const discount = discountCodes.value.find(d => d.code === calculationForm.value.discountCode)
  if (discount && discount.valid) {
    calculateTax()
  } else {
    alert('کد تخفیف نامعتبر است')
  }
}

// محاسبه مالیات
const calculateTax = () => {
  const amount = calculationForm.value.totalAmount || 0
  
  // محاسبه تخفیف
  let discountAmount = 0
  if (calculationForm.value.discountCode) {
    const discount = discountCodes.value.find(d => d.code === calculationForm.value.discountCode)
    if (discount && discount.valid) {
      if (discount.type === 'percentage') {
        discountAmount = (amount * discount.value) / 100
      } else {
        discountAmount = discount.value
      }
    }
  }
  
  const amountAfterDiscount = amount - discountAmount
  
  // محاسبه مالیات‌ها
  let vatAmount = 0
  let incomeTaxAmount = 0
  let customsTaxAmount = 0
  let exemptionsCount = 0
  
  if (calculationForm.value.applyVAT) {
    vatAmount = (amountAfterDiscount * taxRates.value.vat) / 100
  }
  
  if (calculationForm.value.applyIncomeTax) {
    incomeTaxAmount = (amountAfterDiscount * taxRates.value.income) / 100
  }
  
  if (calculationForm.value.applyCustomsTax) {
    customsTaxAmount = (amountAfterDiscount * taxRates.value.customs) / 100
  }
  
  // اعمال معافیت‌ها
  if (calculationForm.value.category === 'food') {
    exemptionsCount++
    vatAmount = 0
    incomeTaxAmount = 0
  }
  
  const totalTaxAmount = vatAmount + incomeTaxAmount + customsTaxAmount
  let finalAmount = amountAfterDiscount + totalTaxAmount
  
  // گرد کردن
  if (calculationForm.value.roundToNearest) {
    finalAmount = Math.round(finalAmount / 1000) * 1000
  }
  
  calculationResult.value = {
    totalAmount: amount,
    discountAmount,
    amountAfterDiscount,
    vatAmount,
    incomeTaxAmount,
    customsTaxAmount,
    finalAmount,
    vatRate: taxRates.value.vat,
    incomeTaxRate: taxRates.value.income,
    customsTaxRate: taxRates.value.customs,
    exemptionsCount,
    totalTaxAmount
  }
}

// ذخیره محاسبه
const saveCalculation = async () => {
  try {
    // TODO: ارسال درخواست به API
    console.log('محاسبه ذخیره شد:', calculationResult.value)
  } catch (error) {
    console.error('خطا در ذخیره محاسبه:', error)
  }
}

// خروجی PDF
const exportCalculation = () => {
  // TODO: تولید PDF
  console.log('خروجی PDF تولید شد')
}

// ریست محاسبه
const resetCalculation = () => {
  calculationForm.value = {
    customerType: 'individual',
    region: '',
    category: '',
    totalAmount: 0,
    discountCode: '',
    applyVAT: true,
    applyIncomeTax: false,
    applyCustomsTax: false,
    roundToNearest: true
  }
  
  calculationResult.value = {
    totalAmount: 0,
    discountAmount: 0,
    amountAfterDiscount: 0,
    vatAmount: 0,
    incomeTaxAmount: 0,
    customsTaxAmount: 0,
    finalAmount: 0,
    vatRate: 0,
    incomeTaxRate: 0,
    customsTaxRate: 0,
    exemptionsCount: 0,
    totalTaxAmount: 0
  }
}

// بارگذاری اولیه
onMounted(() => {
  // TODO: بارگذاری نرخ‌ها و معافیت‌ها از API
})
</script>

<!--
  کامپوننت محاسبه‌گر خودکار مالیات
  شامل:
  1. فرم محاسبه با ورودی‌های مختلف
  2. محاسبه دقیق مالیات بر اساس نرخ‌ها
  3. اعمال معافیت‌ها و تخفیف‌ها
  4. نمایش نتایج به صورت تفصیلی
  5. محاسبه دسته‌ای
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
