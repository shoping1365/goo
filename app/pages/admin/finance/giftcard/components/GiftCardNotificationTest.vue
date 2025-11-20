<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center">
        <div>
          <h2 class="text-xl font-bold text-gray-900">تست نوتیفیکیشن</h2>
          <p class="text-gray-600 mt-1">تست و بررسی عملکرد سیستم اطلاع‌رسانی گیفت کارت</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            class="px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
            @click="runAllTests"
          >
            اجرای همه تست‌ها
          </button>
          <button 
            class="px-4 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
            @click="clearResults"
          >
            پاک کردن نتایج
          </button>
        </div>
      </div>
    </div>

    <!-- آمار تست‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">کل تست‌ها</p>
            <p class="text-2xl font-bold text-gray-900">{{ totalTests }}</p>
            <p class="text-xs text-blue-600">انجام شده</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-green-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">موفق</p>
            <p class="text-2xl font-bold text-gray-900">{{ successfulTests }}</p>
            <p class="text-xs text-green-600">{{ successRate }}%</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-red-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">ناموفق</p>
            <p class="text-2xl font-bold text-gray-900">{{ failedTests }}</p>
            <p class="text-xs text-red-600">{{ failureRate }}%</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-yellow-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">زمان متوسط</p>
            <p class="text-2xl font-bold text-gray-900">{{ averageTime }}ms</p>
            <p class="text-xs text-yellow-600">آخرین تست</p>
          </div>
        </div>
      </div>
    </div>

    <!-- تست‌های مختلف -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- تست خرید جدید -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-medium text-gray-900">تست اطلاع‌رسانی خرید جدید</h3>
        </div>
        <div class="p-6 space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ تست (تومان)</label>
            <input
              v-model.number="testData.purchase.amount"
              type="number"
              min="1000"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام خریدار</label>
            <input
              v-model="testData.purchase.buyerName"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">روش پرداخت</label>
            <select
              v-model="testData.purchase.paymentMethod"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="online">آنلاین</option>
              <option value="wallet">کیف پول</option>
              <option value="giftcard">گیفت کارت</option>
            </select>
          </div>
          
          <div class="flex items-center space-x-3 space-x-reverse">
            <button 
              class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
              @click="testPurchaseNotification"
            >
              اجرای تست
            </button>
            <div v-if="testResults.purchase" class="flex items-center">
              <span :class="getTestResultClass(testResults.purchase.success)" class="text-sm">
                {{ testResults.purchase.success ? 'موفق' : 'ناموفق' }}
              </span>
              <span class="text-xs text-gray-500 mr-2">({{ testResults.purchase.time }}ms)</span>
            </div>
          </div>
        </div>
      </div>

      <!-- تست هشدار انقضا -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-medium text-gray-900">تست هشدار انقضا</h3>
        </div>
        <div class="p-6 space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">روزهای باقی‌مانده</label>
            <input
              v-model.number="testData.expiry.daysRemaining"
              type="number"
              min="1"
              max="30"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ باقی‌مانده (تومان)</label>
            <input
              v-model.number="testData.expiry.remainingAmount"
              type="number"
              min="0"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام گیرنده</label>
            <input
              v-model="testData.expiry.recipientName"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
          
          <div class="flex items-center space-x-3 space-x-reverse">
            <button 
              class="px-4 py-2 bg-yellow-600 text-white text-sm font-medium rounded-lg hover:bg-yellow-700 focus:outline-none focus:ring-2 focus:ring-yellow-500 focus:ring-offset-2"
              @click="testExpiryNotification"
            >
              اجرای تست
            </button>
            <div v-if="testResults.expiry" class="flex items-center">
              <span :class="getTestResultClass(testResults.expiry.success)" class="text-sm">
                {{ testResults.expiry.success ? 'موفق' : 'ناموفق' }}
              </span>
              <span class="text-xs text-gray-500 mr-2">({{ testResults.expiry.time }}ms)</span>
            </div>
          </div>
        </div>
      </div>

      <!-- تست تراکنش مشکوک -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-medium text-gray-900">تست تراکنش مشکوک</h3>
        </div>
        <div class="p-6 space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ تراکنش (تومان)</label>
            <input
              v-model.number="testData.suspicious.amount"
              type="number"
              min="1000"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">IP آدرس</label>
            <input
              v-model="testData.suspicious.ipAddress"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دلیل مشکوک بودن</label>
            <select
              v-model="testData.suspicious.reason"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="large_amount">مبلغ بالا</option>
              <option value="multiple_attempts">تلاش‌های متعدد</option>
              <option value="unusual_time">زمان غیرعادی</option>
              <option value="different_ip">IP متفاوت</option>
              <option value="rapid_transactions">تراکنش‌های سریع</option>
            </select>
          </div>
          
          <div class="flex items-center space-x-3 space-x-reverse">
            <button 
              class="px-4 py-2 bg-red-600 text-white text-sm font-medium rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
              @click="testSuspiciousNotification"
            >
              اجرای تست
            </button>
            <div v-if="testResults.suspicious" class="flex items-center">
              <span :class="getTestResultClass(testResults.suspicious.success)" class="text-sm">
                {{ testResults.suspicious.success ? 'موفق' : 'ناموفق' }}
              </span>
              <span class="text-xs text-gray-500 mr-2">({{ testResults.suspicious.time }}ms)</span>
            </div>
          </div>
        </div>
      </div>

      <!-- تست سیستم -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-medium text-gray-900">تست سیستم</h3>
        </div>
        <div class="p-6 space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع پیام</label>
            <select
              v-model="testData.system.messageType"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="info">اطلاعات</option>
              <option value="warning">هشدار</option>
              <option value="error">خطا</option>
              <option value="success">موفقیت</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">متن پیام</label>
            <textarea
              v-model="testData.system.message"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            ></textarea>
          </div>
          
          <div class="flex items-center space-x-3 space-x-reverse">
            <button 
              class="px-4 py-2 bg-purple-600 text-white text-sm font-medium rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2"
              @click="testSystemNotification"
            >
              اجرای تست
            </button>
            <div v-if="testResults.system" class="flex items-center">
              <span :class="getTestResultClass(testResults.system.success)" class="text-sm">
                {{ testResults.system.success ? 'موفق' : 'ناموفق' }}
              </span>
              <span class="text-xs text-gray-500 mr-2">({{ testResults.system.time }}ms)</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- نتایج تست‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">نتایج تست‌ها</h3>
      </div>
      <div class="p-6">
        <div v-if="testHistory.length === 0" class="text-center py-8">
          <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
          <h3 class="mt-2 text-sm font-medium text-gray-900">هیچ تستی انجام نشده</h3>
          <p class="mt-1 text-sm text-gray-500">برای شروع، یکی از تست‌های بالا را اجرا کنید.</p>
        </div>
        
        <div v-else class="space-y-4">
          <div 
            v-for="(test, index) in testHistory" 
            :key="index"
            class="border border-gray-200 rounded-lg p-6"
          >
            <div class="flex items-center justify-between">
              <div class="flex items-center space-x-3 space-x-reverse">
                <div :class="getTestResultIconClass(test.success)" class="w-6 h-6 rounded-full flex items-center justify-center">
                  <component :is="getTestResultIcon(test.success)" class="w-4 h-4" />
                </div>
                <div>
                  <h4 class="text-sm font-medium text-gray-900">{{ test.name }}</h4>
                  <p class="text-xs text-gray-500">{{ formatDate(test.timestamp) }}</p>
                </div>
              </div>
              <div class="flex items-center space-x-4 space-x-reverse">
                <span :class="getTestResultClass(test.success)" class="text-sm font-medium">
                  {{ test.success ? 'موفق' : 'ناموفق' }}
                </span>
                <span class="text-xs text-gray-500">{{ test.time }}ms</span>
              </div>
            </div>
            
            <div v-if="test.message" class="mt-2 p-3 bg-gray-50 rounded-lg">
              <p class="text-sm text-gray-700">{{ test.message }}</p>
            </div>
            
            <div v-if="test.details" class="mt-2">
              <details class="text-sm">
                <summary class="cursor-pointer text-gray-600 hover:text-gray-900">جزئیات بیشتر</summary>
                <div class="mt-2 p-3 bg-gray-50 rounded-lg">
                  <pre class="text-xs text-gray-700 whitespace-pre-wrap">{{ JSON.stringify(test.details, null, 2) }}</pre>
                </div>
              </details>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'

// Reactive data
const testHistory = ref([])

// داده‌های تست
const testData = reactive({
  purchase: {
    amount: 500000,
    buyerName: 'علی احمدی',
    paymentMethod: 'online'
  },
  expiry: {
    daysRemaining: 3,
    remainingAmount: 750000,
    recipientName: 'فاطمه محمدی'
  },
  suspicious: {
    amount: 2000000,
    ipAddress: '192.168.1.100',
    reason: 'large_amount'
  },
  system: {
    messageType: 'info',
    message: 'این یک پیام تست سیستم است'
  }
})

// نتایج تست‌ها
const testResults = reactive({
  purchase: null,
  expiry: null,
  suspicious: null,
  system: null
})

// Computed properties
const totalTests = computed(() => testHistory.value.length)
const successfulTests = computed(() => testHistory.value.filter(t => t.success).length)
const failedTests = computed(() => testHistory.value.filter(t => !t.success).length)
const successRate = computed(() => totalTests.value > 0 ? Math.round((successfulTests.value / totalTests.value) * 100) : 0)
const failureRate = computed(() => totalTests.value > 0 ? Math.round((failedTests.value / totalTests.value) * 100) : 0)
const averageTime = computed(() => {
  if (testHistory.value.length === 0) return 0
  const totalTime = testHistory.value.reduce((sum, test) => sum + test.time, 0)
  return Math.round(totalTime / testHistory.value.length)
})

// Methods
const testPurchaseNotification = async () => {
  const startTime = Date.now()
  
  try {
    console.log('تست اطلاع‌رسانی خرید جدید در حال اجرا...', testData.purchase)
    
    // شبیه‌سازی API call
    await new Promise(resolve => setTimeout(resolve, 500 + Math.random() * 1000))
    
    const time = Date.now() - startTime
    const success = Math.random() > 0.1 // 90% موفقیت
    
    testResults.purchase = { success, time }
    
    testHistory.value.unshift({
      name: 'تست اطلاع‌رسانی خرید جدید',
      success,
      time,
      timestamp: new Date(),
      message: success ? 'نوتیفیکیشن خرید جدید با موفقیت ارسال شد' : 'خطا در ارسال نوتیفیکیشن خرید جدید',
      details: {
        amount: testData.purchase.amount,
        buyerName: testData.purchase.buyerName,
        paymentMethod: testData.purchase.paymentMethod
      }
    })
    
    console.log('تست خرید جدید تکمیل شد:', { success, time })
  } catch (error) {
    console.error('خطا در تست خرید جدید:', error)
    testResults.purchase = { success: false, time: Date.now() - startTime }
  }
}

const testExpiryNotification = async () => {
  const startTime = Date.now()
  
  try {
    console.log('تست هشدار انقضا در حال اجرا...', testData.expiry)
    
    // شبیه‌سازی API call
    await new Promise(resolve => setTimeout(resolve, 300 + Math.random() * 800))
    
    const time = Date.now() - startTime
    const success = Math.random() > 0.05 // 95% موفقیت
    
    testResults.expiry = { success, time }
    
    testHistory.value.unshift({
      name: 'تست هشدار انقضا',
      success,
      time,
      timestamp: new Date(),
      message: success ? 'هشدار انقضا با موفقیت ارسال شد' : 'خطا در ارسال هشدار انقضا',
      details: {
        daysRemaining: testData.expiry.daysRemaining,
        remainingAmount: testData.expiry.remainingAmount,
        recipientName: testData.expiry.recipientName
      }
    })
    
    console.log('تست هشدار انقضا تکمیل شد:', { success, time })
  } catch (error) {
    console.error('خطا در تست هشدار انقضا:', error)
    testResults.expiry = { success: false, time: Date.now() - startTime }
  }
}

const testSuspiciousNotification = async () => {
  const startTime = Date.now()
  
  try {
    console.log('تست تراکنش مشکوک در حال اجرا...', testData.suspicious)
    
    // شبیه‌سازی API call
    await new Promise(resolve => setTimeout(resolve, 400 + Math.random() * 600))
    
    const time = Date.now() - startTime
    const success = Math.random() > 0.15 // 85% موفقیت
    
    testResults.suspicious = { success, time }
    
    testHistory.value.unshift({
      name: 'تست تراکنش مشکوک',
      success,
      time,
      timestamp: new Date(),
      message: success ? 'نوتیفیکیشن تراکنش مشکوک با موفقیت ارسال شد' : 'خطا در ارسال نوتیفیکیشن تراکنش مشکوک',
      details: {
        amount: testData.suspicious.amount,
        ipAddress: testData.suspicious.ipAddress,
        reason: testData.suspicious.reason
      }
    })
    
    console.log('تست تراکنش مشکوک تکمیل شد:', { success, time })
  } catch (error) {
    console.error('خطا در تست تراکنش مشکوک:', error)
    testResults.suspicious = { success: false, time: Date.now() - startTime }
  }
}

const testSystemNotification = async () => {
  const startTime = Date.now()
  
  try {
    console.log('تست سیستم در حال اجرا...', testData.system)
    
    // شبیه‌سازی API call
    await new Promise(resolve => setTimeout(resolve, 200 + Math.random() * 400))
    
    const time = Date.now() - startTime
    const success = Math.random() > 0.02 // 98% موفقیت
    
    testResults.system = { success, time }
    
    testHistory.value.unshift({
      name: 'تست سیستم',
      success,
      time,
      timestamp: new Date(),
      message: success ? 'پیام سیستم با موفقیت ارسال شد' : 'خطا در ارسال پیام سیستم',
      details: {
        messageType: testData.system.messageType,
        message: testData.system.message
      }
    })
    
    console.log('تست سیستم تکمیل شد:', { success, time })
  } catch (error) {
    console.error('خطا در تست سیستم:', error)
    testResults.system = { success: false, time: Date.now() - startTime }
  }
}

const runAllTests = async () => {
  console.log('اجرای همه تست‌ها شروع شد')
  
  await testPurchaseNotification()
  await new Promise(resolve => setTimeout(resolve, 1000))
  
  await testExpiryNotification()
  await new Promise(resolve => setTimeout(resolve, 1000))
  
  await testSuspiciousNotification()
  await new Promise(resolve => setTimeout(resolve, 1000))
  
  await testSystemNotification()
  
  console.log('همه تست‌ها تکمیل شد')
}

const clearResults = () => {
  if (confirm('آیا مطمئن هستید که می‌خواهید همه نتایج تست‌ها را پاک کنید؟')) {
    testHistory.value = []
    Object.keys(testResults).forEach(key => {
      testResults[key] = null
    })
    console.log('نتایج تست‌ها پاک شد')
  }
}

const formatDate = (date: Date) => {
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit'
  }).format(date)
}

const getTestResultClass = (success: boolean) => {
  return success ? 'text-green-600' : 'text-red-600'
}

const getTestResultIconClass = (success: boolean) => {
  return success ? 'bg-green-100 text-green-600' : 'bg-red-100 text-red-600'
}

const getTestResultIcon = (success: boolean) => {
  return success ? 'CheckIcon' : 'XIcon'
}

// Lifecycle
onMounted(() => {
  console.log('Gift card notification test component mounted')
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
