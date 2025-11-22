<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="border-b border-gray-200 pb-4">
      <h3 class="text-lg font-medium text-gray-900">تشخیص تقلب و امنیت</h3>
      <p class="mt-1 text-sm text-gray-600">تشخیص و جلوگیری از فعالیت‌های مشکوک و تقلبی در سیستم</p>
    </div>

    <!-- تشخیص بوت -->
    <div class="bg-gray-50 rounded-lg p-6">
      <div class="flex items-center justify-between mb-4">
        <h4 class="text-md font-medium text-gray-900">تشخیص بوت</h4>
        <div class="flex items-center space-x-2 space-x-reverse">
          <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
            فعال
          </span>
          <button class="text-blue-600 hover:text-blue-800 text-sm">تنظیمات</button>
        </div>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="bg-white rounded-lg p-6 border border-gray-200 text-center">
          <div class="text-2xl font-bold text-red-600 mb-1">۲۳۴</div>
          <div class="text-sm text-gray-600">بوت تشخیص داده شده</div>
        </div>
        
        <div class="bg-white rounded-lg p-6 border border-gray-200 text-center">
          <div class="text-2xl font-bold text-green-600 mb-1">۹۶.۸%</div>
          <div class="text-sm text-gray-600">دقت تشخیص</div>
        </div>
        
        <div class="bg-white rounded-lg p-6 border border-gray-200 text-center">
          <div class="text-2xl font-bold text-blue-600 mb-1">۱.۲ ثانیه</div>
          <div class="text-sm text-gray-600">متوسط زمان تشخیص</div>
        </div>
      </div>
      
      <div class="mt-4 bg-white rounded-lg p-6 border border-gray-200">
        <h5 class="text-sm font-medium text-gray-900 mb-3">آخرین بوت‌های تشخیص داده شده</h5>
        <div class="space-y-2">
          <div class="flex items-center justify-between text-sm">
            <span class="text-gray-700">IP: 192.168.1.100 - فعالیت مشکوک</span>
            <span class="text-red-600 font-medium">مسدود شده</span>
          </div>
          <div class="flex items-center justify-between text-sm">
            <span class="text-gray-700">IP: 10.0.0.45 - کلیک‌های خودکار</span>
            <span class="text-red-600 font-medium">مسدود شده</span>
          </div>
          <div class="flex items-center justify-between text-sm">
            <span class="text-gray-700">IP: 172.16.0.23 - رفتار غیرعادی</span>
            <span class="text-red-600 font-medium">مسدود شده</span>
          </div>
        </div>
      </div>
    </div>

    <!-- تشخیص خرید جعلی -->
    <div class="bg-gray-50 rounded-lg p-6">
      <div class="flex items-center justify-between mb-4">
        <h4 class="text-md font-medium text-gray-900">تشخیص خرید جعلی</h4>
        <button class="text-blue-600 hover:text-blue-800 text-sm">تنظیمات تشخیص</button>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="bg-white rounded-lg p-6 border border-gray-200">
          <h5 class="text-sm font-medium text-gray-900 mb-3">آمار خریدهای مشکوک</h5>
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">کل خریدهای بررسی شده</span>
              <span class="text-sm font-medium">۱۵,۶۷۸</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">خریدهای مشکوک</span>
              <span class="text-sm font-medium text-red-600">۱۲۳</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">خریدهای تایید شده جعلی</span>
              <span class="text-sm font-medium text-red-600">۸۹</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">دقت تشخیص</span>
              <span class="text-sm font-medium text-green-600">۹۲.۳%</span>
            </div>
          </div>
        </div>
        
        <div class="bg-white rounded-lg p-6 border border-gray-200">
          <h5 class="text-sm font-medium text-gray-900 mb-3">معیارهای تشخیص</h5>
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">سرعت خرید غیرعادی</span>
              <div class="flex items-center">
                <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                  <div class="bg-red-600 h-2 rounded-full" style="width: 85%"></div>
                </div>
                <span class="text-sm font-medium">۸۵%</span>
              </div>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">الگوی خرید غیرمعمول</span>
              <div class="flex items-center">
                <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                  <div class="bg-orange-600 h-2 rounded-full" style="width: 72%"></div>
                </div>
                <span class="text-sm font-medium">۷۲%</span>
              </div>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">اطلاعات پرداخت مشکوک</span>
              <div class="flex items-center">
                <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                  <div class="bg-yellow-600 h-2 rounded-full" style="width: 68%"></div>
                </div>
                <span class="text-sm font-medium">۶۸%</span>
              </div>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">رفتار غیرعادی کاربر</span>
              <div class="flex items-center">
                <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                  <div class="bg-green-600 h-2 rounded-full" style="width: 91%"></div>
                </div>
                <span class="text-sm font-medium">۹۱%</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تشخیص نظرات جعلی -->
    <div class="bg-gray-50 rounded-lg p-6">
      <div class="flex items-center justify-between mb-4">
        <h4 class="text-md font-medium text-gray-900">تشخیص نظرات جعلی</h4>
        <button class="text-blue-600 hover:text-blue-800 text-sm">تنظیمات فیلتر</button>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="bg-white rounded-lg p-6 border border-gray-200 text-center">
          <div class="text-2xl font-bold text-red-600 mb-1">۴۵۶</div>
          <div class="text-sm text-gray-600">نظر جعلی مسدود شده</div>
        </div>
        
        <div class="bg-white rounded-lg p-6 border border-gray-200 text-center">
          <div class="text-2xl font-bold text-green-600 mb-1">۹۸.۷%</div>
          <div class="text-sm text-gray-600">دقت تشخیص</div>
        </div>
        
        <div class="bg-white rounded-lg p-6 border border-gray-200 text-center">
          <div class="text-2xl font-bold text-blue-600 mb-1">۰.۸ ثانیه</div>
          <div class="text-sm text-gray-600">متوسط زمان بررسی</div>
        </div>
      </div>
      
      <div class="mt-4 bg-white rounded-lg p-6 border border-gray-200">
        <h5 class="text-sm font-medium text-gray-900 mb-3">آخرین نظرات جعلی تشخیص داده شده</h5>
        <div class="space-y-2">
          <div class="flex items-center justify-between text-sm">
            <span class="text-gray-700">"محصول فوق‌العاده، قیمت باورنکردنی..."</span>
            <span class="text-red-600 font-medium">جعلی</span>
          </div>
          <div class="flex items-center justify-between text-sm">
            <span class="text-gray-700">"کیفیت عالی، حتماً بخرید..."</span>
            <span class="text-red-600 font-medium">جعلی</span>
          </div>
          <div class="flex items-center justify-between text-sm">
            <span class="text-gray-700">"بهترین محصول دنیا..."</span>
            <span class="text-red-600 font-medium">جعلی</span>
          </div>
        </div>
      </div>
    </div>

    <!-- تشخیص حملات -->
    <div class="bg-gray-50 rounded-lg p-6">
      <div class="flex items-center justify-between mb-4">
        <h4 class="text-md font-medium text-gray-900">تشخیص حملات</h4>
        <button class="text-blue-600 hover:text-blue-800 text-sm">تنظیمات امنیت</button>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="bg-white rounded-lg p-6 border border-gray-200">
          <h5 class="text-sm font-medium text-gray-900 mb-3">آمار حملات</h5>
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">حملات DDoS</span>
              <span class="text-sm font-medium text-red-600">۲۳</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">حملات Brute Force</span>
              <span class="text-sm font-medium text-red-600">۱۵۶</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">حملات SQL Injection</span>
              <span class="text-sm font-medium text-red-600">۴۵</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">حملات XSS</span>
              <span class="text-sm font-medium text-red-600">۶۷</span>
            </div>
          </div>
        </div>
        
        <div class="bg-white rounded-lg p-6 border border-gray-200">
          <h5 class="text-sm font-medium text-gray-900 mb-3">وضعیت امنیت</h5>
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">سطح امنیت کلی</span>
              <span class="text-sm font-medium text-green-600">عالی</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">آخرین به‌روزرسانی</span>
              <span class="text-sm font-medium">۲ ساعت پیش</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">تعداد قوانین فعال</span>
              <span class="text-sm font-medium">۲۳۴</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">نرخ مسدودسازی</span>
              <span class="text-sm font-medium text-green-600">۹۹.۸%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تشخیص ناهنجاری -->
    <div class="bg-gray-50 rounded-lg p-6">
      <div class="flex items-center justify-between mb-4">
        <h4 class="text-md font-medium text-gray-900">تشخیص ناهنجاری</h4>
        <button class="text-blue-600 hover:text-blue-800 text-sm">تنظیمات تشخیص</button>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="bg-white rounded-lg p-6 border border-gray-200">
          <h5 class="text-sm font-medium text-gray-900 mb-3">ناهنجاری‌های تشخیص داده شده</h5>
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">رفتار غیرعادی کاربر</span>
              <span class="text-sm font-medium text-red-600">۸۹</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">الگوی خرید غیرمعمول</span>
              <span class="text-sm font-medium text-red-600">۱۲۳</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">فعالیت مشکوک</span>
              <span class="text-sm font-medium text-red-600">۶۷</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">دسترسی غیرمجاز</span>
              <span class="text-sm font-medium text-red-600">۳۴</span>
            </div>
          </div>
        </div>
        
        <div class="bg-white rounded-lg p-6 border border-gray-200">
          <h5 class="text-sm font-medium text-gray-900 mb-3">عملکرد تشخیص</h5>
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">دقت تشخیص</span>
              <span class="text-sm font-medium text-green-600">۹۴.۲%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">نرخ خطای مثبت</span>
              <span class="text-sm font-medium text-red-600">۵.۸%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">زمان تشخیص</span>
              <span class="text-sm font-medium">۰.۵ ثانیه</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">مدل‌های فعال</span>
              <span class="text-sm font-medium">۸</span>
            </div>
          </div>
        </div>
      </div>
      
      <div class="mt-4 bg-white rounded-lg p-6 border border-gray-200">
        <h5 class="text-sm font-medium text-gray-900 mb-3">آخرین ناهنجاری‌ها</h5>
        <div class="space-y-2">
          <div class="flex items-center justify-between text-sm">
            <span class="text-gray-700">کاربر ID: 12345 - رفتار غیرعادی در خرید</span>
            <span class="text-red-600 font-medium">تحت بررسی</span>
          </div>
          <div class="flex items-center justify-between text-sm">
            <span class="text-gray-700">IP: 192.168.1.50 - فعالیت مشکوک</span>
            <span class="text-red-600 font-medium">مسدود شده</span>
          </div>
          <div class="flex items-center justify-between text-sm">
            <span class="text-gray-700">Session: ABC123 - دسترسی غیرمجاز</span>
            <span class="text-red-600 font-medium">مسدود شده</span>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات پیشرفته امنیت -->
    <div class="bg-gray-50 rounded-lg p-6">
      <h4 class="text-md font-medium text-gray-900 mb-4">تنظیمات پیشرفته امنیت</h4>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <span class="text-sm font-medium text-gray-900">تشخیص بوت</span>
            <button class="relative inline-flex h-6 w-11 items-center rounded-full bg-blue-600">
              <span class="inline-block h-4 w-4 transform rounded-full bg-white transition translate-x-6"></span>
            </button>
          </div>
          
          <div class="flex items-center justify-between">
            <span class="text-sm font-medium text-gray-900">تشخیص خرید جعلی</span>
            <button class="relative inline-flex h-6 w-11 items-center rounded-full bg-blue-600">
              <span class="inline-block h-4 w-4 transform rounded-full bg-white transition translate-x-6"></span>
            </button>
          </div>
          
          <div class="flex items-center justify-between">
            <span class="text-sm font-medium text-gray-900">تشخیص نظرات جعلی</span>
            <button class="relative inline-flex h-6 w-11 items-center rounded-full bg-blue-600">
              <span class="inline-block h-4 w-4 transform rounded-full bg-white transition translate-x-6"></span>
            </button>
          </div>
        </div>
        
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <span class="text-sm font-medium text-gray-900">تشخیص حملات</span>
            <button class="relative inline-flex h-6 w-11 items-center rounded-full bg-blue-600">
              <span class="inline-block h-4 w-4 transform rounded-full bg-white transition translate-x-6"></span>
            </button>
          </div>
          
          <div class="flex items-center justify-between">
            <span class="text-sm font-medium text-gray-900">تشخیص ناهنجاری</span>
            <button class="relative inline-flex h-6 w-11 items-center rounded-full bg-blue-600">
              <span class="inline-block h-4 w-4 transform rounded-full bg-white transition translate-x-6"></span>
            </button>
          </div>
          
          <div class="flex items-center justify-between">
            <span class="text-sm font-medium text-gray-900">یادگیری مداوم</span>
            <button class="relative inline-flex h-6 w-11 items-center rounded-full bg-blue-600">
              <span class="inline-block h-4 w-4 transform rounded-full bg-white transition translate-x-6"></span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})
// کامپوننت تشخیص تقلب و امنیت
// قابلیت‌های ۸۱-۸۵: تشخیص بوت، تشخیص خرید جعلی، تشخیص نظرات جعلی، تشخیص حملات، تشخیص ناهنجاری

// متغیرهای کامپوننت
const _botDetectionStats = ref({
  detectedBots: 234,
  accuracy: 96.8,
  avgDetectionTime: 1.2
})

const _fakePurchaseStats = ref({
  totalChecked: 15678,
  suspiciousPurchases: 123,
  confirmedFake: 89,
  accuracy: 92.3
})

const _fakeReviewStats = ref({
  blockedFakeReviews: 456,
  accuracy: 98.7,
  avgCheckTime: 0.8
})

const _attackStats = ref({
  ddosAttacks: 23,
  bruteForceAttacks: 156,
  sqlInjectionAttacks: 45,
  xssAttacks: 67
})

const _securityStatus = ref({
  overallSecurityLevel: 'عالی',
  lastUpdate: '۲ ساعت پیش',
  activeRules: 234,
  blockRate: 99.8
})

const _anomalyStats = ref({
  abnormalUserBehavior: 89,
  unusualPurchasePattern: 123,
  suspiciousActivity: 67,
  unauthorizedAccess: 34
})

const _anomalyPerformance = ref({
  accuracy: 94.2,
  falsePositiveRate: 5.8,
  detectionTime: 0.5,
  activeModels: 8
})

// توابع کامپوننت
const _updateBotDetection = () => {
  // به‌روزرسانی تشخیص بوت
}

const _updateFakePurchaseDetection = () => {
  // به‌روزرسانی تشخیص خرید جعلی
}

const _updateFakeReviewDetection = () => {
  // به‌روزرسانی تشخیص نظرات جعلی
}

const _updateAttackDetection = () => {
  // به‌روزرسانی تشخیص حملات
}

const _updateAnomalyDetection = () => {
  // به‌روزرسانی تشخیص ناهنجاری
}

onMounted(() => {
  // مقداردهی اولیه کامپوننت
})
</script> 
