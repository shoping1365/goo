<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <!-- هدر -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6 mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">تست محاسبات مالیاتی</h3>
        <p class="text-sm text-gray-600">تست و اعتبارسنجی محاسبات مالیاتی با سناریوهای مختلف</p>
      </div>
      
      <!-- دکمه اجرای تست‌ها -->
      <button 
        class="inline-flex items-center gap-2 px-4 py-2 bg-red-600 hover:bg-red-700 text-white rounded-lg transition-colors duration-200"
        @click="runAllTests"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        اجرای تمام تست‌ها
      </button>
    </div>

    <!-- آمار تست‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ testStats.passed }}</div>
            <div class="text-sm text-green-700">تست‌های موفق</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-red-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-red-600">{{ testStats.failed }}</div>
            <div class="text-sm text-red-700">تست‌های ناموفق</div>
          </div>
          <div class="w-10 h-10 bg-red-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-yellow-600">{{ testStats.warning }}</div>
            <div class="text-sm text-yellow-700">هشدارها</div>
          </div>
          <div class="w-10 h-10 bg-yellow-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ testStats.total }}</div>
            <div class="text-sm text-blue-700">کل تست‌ها</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- تب‌های تست -->
    <div class="mb-6">
      <div class="border-b border-gray-200">
        <nav class="-mb-px flex space-x-8">
          <button 
            v-for="tab in testTabs" 
            :key="tab.id"
            :class="[
              'py-2 px-1 border-b-2 font-medium text-sm transition-colors duration-200',
              activeTestTab === tab.id 
                ? 'border-red-500 text-red-600' 
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
            @click="activeTestTab = tab.id"
          >
            {{ tab.label }}
          </button>
        </nav>
      </div>
    </div>

    <!-- تست‌های محاسباتی -->
    <div v-if="activeTestTab === 'calculations'">
      <div class="space-y-4">
        <div v-for="test in calculationTests" :key="test.id" class="border border-gray-200 rounded-lg p-6">
          <div class="flex items-center justify-between mb-3">
            <h4 class="font-medium text-gray-900">{{ test.name }}</h4>
            <div class="flex items-center gap-2">
              <span 
                :class="[
                  'px-2 py-1 rounded-full text-xs font-medium',
                  test.status === 'passed' ? 'bg-green-100 text-green-700' : 
                  test.status === 'failed' ? 'bg-red-100 text-red-700' : 
                  'bg-yellow-100 text-yellow-700'
                ]"
              >
                {{ getTestStatusLabel(test.status) }}
              </span>
              <button 
                class="px-3 py-1 bg-blue-100 text-blue-600 rounded-lg text-sm hover:bg-blue-200"
                @click="runTest(test)"
              >
                اجرا
              </button>
            </div>
          </div>
          
          <div class="text-sm text-gray-600 mb-3">{{ test.description }}</div>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-sm">
            <div>
              <div class="font-medium text-gray-700 mb-1">ورودی:</div>
              <div class="bg-gray-50 rounded p-2 font-mono text-xs">
                {{ JSON.stringify(test.input, null, 2) }}
              </div>
            </div>
            <div>
              <div class="font-medium text-gray-700 mb-1">خروجی مورد انتظار:</div>
              <div class="bg-gray-50 rounded p-2 font-mono text-xs">
                {{ JSON.stringify(test.expectedOutput, null, 2) }}
              </div>
            </div>
          </div>
          
          <div v-if="test.actualOutput" class="mt-3">
            <div class="font-medium text-gray-700 mb-1">خروجی واقعی:</div>
            <div class="bg-gray-50 rounded p-2 font-mono text-xs">
              {{ JSON.stringify(test.actualOutput, null, 2) }}
            </div>
          </div>
          
          <div v-if="test.error" class="mt-3 p-3 bg-red-50 rounded-lg">
            <div class="text-sm text-red-700">{{ test.error }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- تست‌های معافیت -->
    <div v-else-if="activeTestTab === 'exemptions'">
      <div class="space-y-4">
        <div v-for="test in exemptionTests" :key="test.id" class="border border-gray-200 rounded-lg p-6">
          <div class="flex items-center justify-between mb-3">
            <h4 class="font-medium text-gray-900">{{ test.name }}</h4>
            <div class="flex items-center gap-2">
              <span 
                :class="[
                  'px-2 py-1 rounded-full text-xs font-medium',
                  test.status === 'passed' ? 'bg-green-100 text-green-700' : 
                  test.status === 'failed' ? 'bg-red-100 text-red-700' : 
                  'bg-yellow-100 text-yellow-700'
                ]"
              >
                {{ getTestStatusLabel(test.status) }}
              </span>
              <button 
                class="px-3 py-1 bg-blue-100 text-blue-600 rounded-lg text-sm hover:bg-blue-200"
                @click="runExemptionTest(test)"
              >
                اجرا
              </button>
            </div>
          </div>
          
          <div class="text-sm text-gray-600 mb-3">{{ test.description }}</div>
          
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6 text-sm">
            <div>
              <div class="font-medium text-gray-700 mb-1">نوع محصول:</div>
              <div class="bg-gray-50 rounded p-2">{{ test.productType }}</div>
            </div>
            <div>
              <div class="font-medium text-gray-700 mb-1">مبلغ:</div>
              <div class="bg-gray-50 rounded p-2">{{ formatCurrency(test.amount) }}</div>
            </div>
            <div>
              <div class="font-medium text-gray-700 mb-1">معافیت مورد انتظار:</div>
              <div class="bg-gray-50 rounded p-2">{{ test.expectedExemption }}%</div>
            </div>
          </div>
          
          <div v-if="test.actualExemption !== undefined" class="mt-3">
            <div class="font-medium text-gray-700 mb-1">معافیت واقعی:</div>
            <div class="bg-gray-50 rounded p-2">{{ test.actualExemption }}%</div>
          </div>
        </div>
      </div>
    </div>

    <!-- تست‌های تخفیف -->
    <div v-else-if="activeTestTab === 'discounts'">
      <div class="space-y-4">
        <div v-for="test in discountTests" :key="test.id" class="border border-gray-200 rounded-lg p-6">
          <div class="flex items-center justify-between mb-3">
            <h4 class="font-medium text-gray-900">{{ test.name }}</h4>
            <div class="flex items-center gap-2">
              <span 
                :class="[
                  'px-2 py-1 rounded-full text-xs font-medium',
                  test.status === 'passed' ? 'bg-green-100 text-green-700' : 
                  test.status === 'failed' ? 'bg-red-100 text-red-700' : 
                  'bg-yellow-100 text-yellow-700'
                ]"
              >
                {{ getTestStatusLabel(test.status) }}
              </span>
              <button 
                class="px-3 py-1 bg-blue-100 text-blue-600 rounded-lg text-sm hover:bg-blue-200"
                @click="runDiscountTest(test)"
              >
                اجرا
              </button>
            </div>
          </div>
          
          <div class="text-sm text-gray-600 mb-3">{{ test.description }}</div>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-sm">
            <div>
              <div class="font-medium text-gray-700 mb-1">کد تخفیف:</div>
              <div class="bg-gray-50 rounded p-2 font-mono">{{ test.discountCode }}</div>
            </div>
            <div>
              <div class="font-medium text-gray-700 mb-1">مبلغ:</div>
              <div class="bg-gray-50 rounded p-2">{{ formatCurrency(test.amount) }}</div>
            </div>
          </div>
          
          <div v-if="test.actualDiscount !== undefined" class="mt-3">
            <div class="font-medium text-gray-700 mb-1">تخفیف واقعی:</div>
            <div class="bg-gray-50 rounded p-2">{{ formatCurrency(test.actualDiscount) }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- گزارش تست‌ها -->
    <div v-else-if="activeTestTab === 'reports'">
      <div class="space-y-6">
        <!-- نمودار نتایج -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">نمودار نتایج تست‌ها</h4>
          <div class="h-64 flex items-end justify-center gap-6">
            <div class="flex flex-col items-center">
              <div class="w-16 bg-green-500 rounded-t-lg" :style="{ height: `${(testStats.passed / testStats.total) * 200}px` }"></div>
              <div class="mt-2 text-xs text-gray-600">موفق</div>
            </div>
            <div class="flex flex-col items-center">
              <div class="w-16 bg-red-500 rounded-t-lg" :style="{ height: `${(testStats.failed / testStats.total) * 200}px` }"></div>
              <div class="mt-2 text-xs text-gray-600">ناموفق</div>
            </div>
            <div class="flex flex-col items-center">
              <div class="w-16 bg-yellow-500 rounded-t-lg" :style="{ height: `${(testStats.warning / testStats.total) * 200}px` }"></div>
              <div class="mt-2 text-xs text-gray-600">هشدار</div>
            </div>
          </div>
        </div>

        <!-- جدول نتایج -->
        <div class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead>
              <tr class="border-b border-gray-200 bg-gray-50">
                <th class="text-right py-3 px-4 font-medium text-gray-600">نوع تست</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">تعداد</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">موفق</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">ناموفق</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">درصد موفقیت</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="report in testReports" :key="report.type" class="border-b border-gray-100">
                <td class="py-3 px-4 text-gray-900">{{ report.type }}</td>
                <td class="py-3 px-4 text-gray-900">{{ report.total }}</td>
                <td class="py-3 px-4 text-green-600">{{ report.passed }}</td>
                <td class="py-3 px-4 text-red-600">{{ report.failed }}</td>
                <td class="py-3 px-4">
                  <span :class="report.successRate >= 80 ? 'text-green-600' : report.successRate >= 60 ? 'text-yellow-600' : 'text-red-600'">
                    {{ report.successRate }}%
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- دکمه‌های عملیاتی -->
        <div class="flex gap-3">
          <button 
            class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors duration-200"
            @click="exportTestReport"
          >
            خروجی گزارش
          </button>
          <button 
            class="px-4 py-2 bg-gray-200 hover:bg-gray-300 text-gray-700 rounded-lg transition-colors duration-200"
            @click="clearTestResults"
          >
            پاک کردن نتایج
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'

// تب‌های تست
const testTabs = [
  { id: 'calculations', label: 'تست محاسبات' },
  { id: 'exemptions', label: 'تست معافیت‌ها' },
  { id: 'discounts', label: 'تست تخفیف‌ها' },
  { id: 'reports', label: 'گزارش تست‌ها' }
]

const activeTestTab = ref('calculations')

// آمار تست‌ها
const testStats = ref({
  passed: 0,
  failed: 0,
  warning: 0,
  total: 0
})

// تست‌های محاسباتی
const calculationTests = ref([
  {
    id: 1,
    name: 'محاسبه مالیات بر ارزش افزوده ساده',
    description: 'تست محاسبه مالیات بر ارزش افزوده با نرخ ۹ درصد',
    status: 'pending',
    input: { amount: 1000000, vatRate: 9 },
    expectedOutput: { vatAmount: 90000, totalAmount: 1090000 },
    actualOutput: null,
    error: null
  },
  {
    id: 2,
    name: 'محاسبه مالیات بر درآمد',
    description: 'تست محاسبه مالیات بر درآمد با نرخ ۱۵ درصد',
    status: 'pending',
    input: { amount: 5000000, incomeTaxRate: 15 },
    expectedOutput: { incomeTaxAmount: 750000, totalAmount: 5750000 },
    actualOutput: null,
    error: null
  },
  {
    id: 3,
    name: 'محاسبه مالیات گمرکی',
    description: 'تست محاسبه مالیات گمرکی با حداقل مبلغ',
    status: 'pending',
    input: { amount: 100000, customsRate: 5, minAmount: 50000 },
    expectedOutput: { customsAmount: 50000, totalAmount: 150000 },
    actualOutput: null,
    error: null
  }
])

// تست‌های معافیت
const exemptionTests = ref([
  {
    id: 1,
    name: 'معافیت مواد غذایی',
    description: 'تست معافیت کامل برای مواد غذایی',
    status: 'pending',
    productType: 'food',
    amount: 500000,
    expectedExemption: 100,
    actualExemption: undefined
  },
  {
    id: 2,
    name: 'معافیت محصولات پزشکی',
    description: 'تست معافیت ۲۵ درصدی برای محصولات پزشکی',
    status: 'pending',
    productType: 'medical',
    amount: 1000000,
    expectedExemption: 25,
    actualExemption: undefined
  }
])

// تست‌های تخفیف
const discountTests = ref([
  {
    id: 1,
    name: 'تخفیف درصدی',
    description: 'تست تخفیف ۱۰ درصدی',
    status: 'pending',
    discountCode: 'TAX2024',
    amount: 1000000,
    expectedDiscount: 100000,
    actualDiscount: undefined
  },
  {
    id: 2,
    name: 'تخفیف مبلغ ثابت',
    description: 'تست تخفیف ۵۰ هزار تومانی',
    status: 'pending',
    discountCode: 'SAVE50K',
    amount: 500000,
    expectedDiscount: 50000,
    actualDiscount: undefined
  }
])

// گزارش تست‌ها
const testReports = computed(() => [
  {
    type: 'محاسبات',
    total: calculationTests.value.length,
    passed: calculationTests.value.filter(t => t.status === 'passed').length,
    failed: calculationTests.value.filter(t => t.status === 'failed').length,
    successRate: Math.round((calculationTests.value.filter(t => t.status === 'passed').length / calculationTests.value.length) * 100)
  },
  {
    type: 'معافیت‌ها',
    total: exemptionTests.value.length,
    passed: exemptionTests.value.filter(t => t.status === 'passed').length,
    failed: exemptionTests.value.filter(t => t.status === 'failed').length,
    successRate: Math.round((exemptionTests.value.filter(t => t.status === 'passed').length / exemptionTests.value.length) * 100)
  },
  {
    type: 'تخفیف‌ها',
    total: discountTests.value.length,
    passed: discountTests.value.filter(t => t.status === 'passed').length,
    failed: discountTests.value.filter(t => t.status === 'failed').length,
    successRate: Math.round((discountTests.value.filter(t => t.status === 'passed').length / discountTests.value.length) * 100)
  }
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

// برچسب وضعیت تست
const getTestStatusLabel = (status: string) => {
  const labels = {
    passed: 'موفق',
    failed: 'ناموفق',
    warning: 'هشدار',
    pending: 'در انتظار'
  }
  return labels[status] || status
}

// اجرای تست محاسباتی
const runTest = async (test: any) => {
  try {
    // شبیه‌سازی محاسبه
    const { amount, vatRate, incomeTaxRate, customsRate, minAmount } = test.input
    
    const actualOutput: any = {}
    
    if (vatRate) {
      actualOutput.vatAmount = (amount * vatRate) / 100
      actualOutput.totalAmount = amount + actualOutput.vatAmount
    }
    
    if (incomeTaxRate) {
      actualOutput.incomeTaxAmount = (amount * incomeTaxRate) / 100
      actualOutput.totalAmount = amount + actualOutput.incomeTaxAmount
    }
    
    if (customsRate) {
      actualOutput.customsAmount = Math.max((amount * customsRate) / 100, minAmount || 0)
      actualOutput.totalAmount = amount + actualOutput.customsAmount
    }
    
    test.actualOutput = actualOutput
    
    // بررسی نتیجه
    const isPassed = JSON.stringify(actualOutput) === JSON.stringify(test.expectedOutput)
    test.status = isPassed ? 'passed' : 'failed'
    
    if (!isPassed) {
      test.error = 'نتایج محاسبه با انتظار مطابقت ندارد'
    }
    
    updateTestStats()
  } catch (error) {
    test.status = 'failed'
    test.error = error.message
    updateTestStats()
  }
}

// اجرای تست معافیت
const runExemptionTest = async (test: any) => {
  try {
    // شبیه‌سازی محاسبه معافیت
    const exemptions = {
      food: 100,
      medical: 25,
      education: 50,
      digital: 0
    }
    
    test.actualExemption = exemptions[test.productType] || 0
    test.status = test.actualExemption === test.expectedExemption ? 'passed' : 'failed'
    
    updateTestStats()
  } catch (error) {
    test.status = 'failed'
    updateTestStats()
  }
}

// اجرای تست تخفیف
const runDiscountTest = async (test: any) => {
  try {
    // شبیه‌سازی محاسبه تخفیف
    const discounts = {
      'TAX2024': { type: 'percentage', value: 10 },
      'SAVE50K': { type: 'fixed', value: 50000 }
    }
    
    const discount = discounts[test.discountCode]
    if (discount) {
      if (discount.type === 'percentage') {
        test.actualDiscount = (test.amount * discount.value) / 100
      } else {
        test.actualDiscount = discount.value
      }
    }
    
    test.status = test.actualDiscount === test.expectedDiscount ? 'passed' : 'failed'
    
    updateTestStats()
  } catch (error) {
    test.status = 'failed'
    updateTestStats()
  }
}

// بروزرسانی آمار تست‌ها
const updateTestStats = () => {
  const allTests = [...calculationTests.value, ...exemptionTests.value, ...discountTests.value]
  
  testStats.value = {
    passed: allTests.filter(t => t.status === 'passed').length,
    failed: allTests.filter(t => t.status === 'failed').length,
    warning: allTests.filter(t => t.status === 'warning').length,
    total: allTests.length
  }
}

// اجرای تمام تست‌ها
const runAllTests = async () => {
  // اجرای تست‌های محاسباتی
  for (const test of calculationTests.value) {
    await runTest(test)
  }
  
  // اجرای تست‌های معافیت
  for (const test of exemptionTests.value) {
    await runExemptionTest(test)
  }
  
  // اجرای تست‌های تخفیف
  for (const test of discountTests.value) {
    await runDiscountTest(test)
  }
}

// خروجی گزارش تست
const exportTestReport = () => {
  // TODO: تولید گزارش
  console.log('گزارش تست تولید شد')
}

// پاک کردن نتایج تست
const clearTestResults = () => {
  calculationTests.value.forEach(test => {
    test.status = 'pending'
    test.actualOutput = null
    test.error = null
  })
  
  exemptionTests.value.forEach(test => {
    test.status = 'pending'
    test.actualExemption = undefined
  })
  
  discountTests.value.forEach(test => {
    test.status = 'pending'
    test.actualDiscount = undefined
  })
  
  updateTestStats()
}

// بارگذاری اولیه
onMounted(() => {
  updateTestStats()
})
</script>

<!--
  کامپوننت تست محاسبات مالیاتی
  شامل:
  1. تست‌های محاسباتی مختلف
  2. تست معافیت‌ها و تخفیف‌ها
  3. گزارش‌گیری از نتایج تست‌ها
  4. نمودار و آمار تست‌ها
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
