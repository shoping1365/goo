<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 w-full">
    <!-- هدر بخش -->
    <div class="p-6 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-lg font-semibold text-gray-900">سیستم تخفیف‌های پویا</h2>
          <p class="text-gray-600 mt-1">مدیریت تخفیف‌های هوشمند و خودکار بر اساس رفتار کاربران</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="showRuleForm = true">
            <span class="i-heroicons-plus ml-2"></span>
            قانون جدید
          </button>
        </div>
      </div>
    </div>

    <!-- آمار سیستم -->
    <div class="p-6 border-b border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="bg-blue-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-cpu-chip text-blue-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-blue-600">قوانین فعال</p>
              <p class="text-2xl font-bold text-blue-900">{{ stats.activeRules }}</p>
            </div>
          </div>
        </div>
        <div class="bg-green-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-currency-dollar text-green-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-green-600">تخفیف اعمال شده</p>
              <p class="text-2xl font-bold text-green-900">{{ formatCurrency(stats.totalDiscount) }}</p>
            </div>
          </div>
        </div>
        <div class="bg-purple-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-users text-purple-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-purple-600">کاربران تحت تاثیر</p>
              <p class="text-2xl font-bold text-purple-900">{{ stats.affectedUsers }}</p>
            </div>
          </div>
        </div>
        <div class="bg-orange-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-orange-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-chart-bar text-orange-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-orange-600">نرخ تبدیل</p>
              <p class="text-2xl font-bold text-orange-900">{{ stats.conversionRate }}%</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تب‌های مدیریت -->
    <div class="border-b border-gray-200">
      <div class="flex border-b border-gray-200 overflow-x-auto">
        <button
v-for="tab in tabs" :key="tab.value" :class="['px-6 py-3 -mb-px font-medium text-sm focus:outline-none whitespace-nowrap', activeTab === tab.value ? 'border-b-2 border-blue-600 text-blue-700' : 'text-gray-500 hover:text-blue-600']"
          @click="activeTab = tab.value">
          {{ tab.label }}
        </button>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="p-6">
      <!-- قوانین تخفیف -->
      <div v-if="activeTab === 'rules'" class="space-y-6">
        <div class="flex justify-between items-center">
          <div class="flex items-center space-x-4 space-x-reverse">
            <select v-model="filterPriority" class="px-3 py-2 border border-gray-300 rounded-lg text-sm" @change="filterRules">
              <option value="">همه اولویت‌ها</option>
              <option value="high">بالا</option>
              <option value="medium">متوسط</option>
              <option value="low">پایین</option>
            </select>
            <select v-model="filterStatus" class="px-3 py-2 border border-gray-300 rounded-lg text-sm" @change="filterRules">
              <option value="">همه وضعیت‌ها</option>
              <option value="active">فعال</option>
              <option value="inactive">غیرفعال</option>
              <option value="draft">پیش‌نویس</option>
            </select>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <button class="px-3 py-1 bg-green-100 text-green-700 rounded text-sm hover:bg-green-200" @click="bulkAction('activate')">
              <span class="i-heroicons-play ml-1"></span>
              فعال‌سازی
            </button>
            <button class="px-3 py-1 bg-red-100 text-red-700 rounded text-sm hover:bg-red-200" @click="bulkAction('deactivate')">
              <span class="i-heroicons-pause ml-1"></span>
              غیرفعال‌سازی
            </button>
          </div>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div v-for="rule in filteredRules" :key="rule.id" class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow">
            <div class="flex items-center justify-between mb-4">
              <div class="flex items-center">
                <div class="w-10 h-10 rounded-lg flex items-center justify-center ml-3" :style="{ backgroundColor: rule.color + '20', color: rule.color }">
                  <span class="i-heroicons-cpu-chip text-lg"></span>
                </div>
                <div>
                  <h4 class="font-medium text-gray-900">{{ rule.name }}</h4>
                  <p class="text-sm text-gray-500">{{ rule.type }} • {{ rule.priority }}</p>
                </div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <span :class="['px-2 py-1 rounded-full text-xs', getStatusClass(rule.status)]">
                  {{ getStatusText(rule.status) }}
                </span>
                <button class="text-blue-600 hover:text-blue-900" @click="editRule(rule)">
                  <span class="i-heroicons-pencil-square"></span>
                </button>
                <button class="text-red-600 hover:text-red-900" @click="deleteRule(rule)">
                  <span class="i-heroicons-trash"></span>
                </button>
              </div>
            </div>
            
            <div class="space-y-3">
              <div class="text-sm text-gray-600">{{ rule.description }}</div>
              <div class="flex flex-wrap gap-2">
                <span v-for="condition in rule.conditions" :key="condition" class="px-2 py-1 bg-gray-100 text-gray-700 rounded text-xs">
                  {{ condition }}
                </span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-500">تخفیف:</span>
                <span class="font-medium">{{ rule.discount }}%</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-500">استفاده شده:</span>
                <span class="font-medium">{{ rule.usageCount }}</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-500">آخرین اجرا:</span>
                <span class="font-medium">{{ formatDate(rule.lastExecuted) }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- موتور هوشمند -->
      <div v-if="activeTab === 'engine'" class="space-y-6">
        <div class="bg-blue-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-blue-900 mb-2">موتور هوشمند تخفیف</h4>
          <p class="text-sm text-blue-700">تنظیمات پیشرفته موتور تصمیم‌گیری هوشمند</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- تنظیمات الگوریتم -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">تنظیمات الگوریتم</h5>
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نوع الگوریتم</label>
                <select v-model="engineSettings.algorithm" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
                  <option value="rule-based">قاعده‌محور</option>
                  <option value="machine-learning">یادگیری ماشین</option>
                  <option value="hybrid">ترکیبی</option>
                </select>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">دقت تصمیم‌گیری (%)</label>
                <input v-model="engineSettings.accuracy" type="range" min="70" max="100" class="w-full">
                <div class="flex justify-between text-xs text-gray-500 mt-1">
                  <span>70%</span>
                  <span>{{ engineSettings.accuracy }}%</span>
                  <span>100%</span>
                </div>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تخفیف (%)</label>
                <input v-model="engineSettings.maxDiscount" type="number" min="1" max="50" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
              </div>
              
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">یادگیری خودکار</span>
                  <p class="text-xs text-gray-500">بهبود مداوم الگوریتم</p>
                </div>
                <input v-model="engineSettings.autoLearning" type="checkbox" class="rounded border-gray-300">
              </div>
            </div>
          </div>

          <!-- فاکتورهای تصمیم‌گیری -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">فاکتورهای تصمیم‌گیری</h5>
            <div class="space-y-4">
              <div v-for="factor in decisionFactors" :key="factor.name" class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">{{ factor.name }}</span>
                  <p class="text-xs text-gray-500">{{ factor.description }}</p>
                </div>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <input v-model="factor.weight" type="range" min="0" max="100" class="w-20">
                  <span class="text-sm text-gray-600 w-8">{{ factor.weight }}%</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- نمودار عملکرد -->
        <div class="border border-gray-200 rounded-lg p-6">
          <h5 class="font-medium text-gray-900 mb-4">عملکرد موتور</h5>
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="text-center">
              <div class="text-2xl font-bold text-blue-600">{{ enginePerformance.accuracy }}%</div>
              <div class="text-sm text-gray-600">دقت تصمیم‌گیری</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-green-600">{{ enginePerformance.responseTime }}ms</div>
              <div class="text-sm text-gray-600">زمان پاسخ</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-purple-600">{{ enginePerformance.throughput }}/sec</div>
              <div class="text-sm text-gray-600">توان عملیاتی</div>
            </div>
          </div>
        </div>
      </div>

      <!-- تحلیل رفتار -->
      <div v-if="activeTab === 'behavior'" class="space-y-6">
        <div class="bg-green-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-green-900 mb-2">تحلیل رفتار کاربران</h4>
          <p class="text-sm text-green-700">تحلیل الگوهای رفتاری و پیش‌بینی نیازهای کاربران</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- الگوهای رفتاری -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">الگوهای رفتاری</h5>
            <div class="space-y-4">
              <div v-for="pattern in behaviorPatterns" :key="pattern.name" class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center justify-between mb-2">
                  <h6 class="font-medium text-gray-900">{{ pattern.name }}</h6>
                  <span :class="['px-2 py-1 rounded-full text-xs', getConfidenceClass(pattern.confidence)]">
                    {{ pattern.confidence }}% اطمینان
                  </span>
                </div>
                <p class="text-sm text-gray-600 mb-3">{{ pattern.description }}</p>
                <div class="flex justify-between text-sm">
                  <span class="text-gray-500">تعداد کاربران:</span>
                  <span class="font-medium">{{ pattern.userCount }}</span>
                </div>
                <div class="flex justify-between text-sm">
                  <span class="text-gray-500">نرخ تکرار:</span>
                  <span class="font-medium">{{ pattern.repeatRate }}%</span>
                </div>
              </div>
            </div>
          </div>

          <!-- پیش‌بینی‌ها -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">پیش‌بینی‌های آینده</h5>
            <div class="space-y-4">
              <div v-for="prediction in predictions" :key="prediction.id" class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center justify-between mb-2">
                  <h6 class="font-medium text-gray-900">{{ prediction.title }}</h6>
                  <span :class="['px-2 py-1 rounded-full text-xs', getPredictionClass(prediction.probability)]">
                    {{ prediction.probability }}% احتمال
                  </span>
                </div>
                <p class="text-sm text-gray-600 mb-3">{{ prediction.description }}</p>
                <div class="flex justify-between text-sm">
                  <span class="text-gray-500">تاریخ پیش‌بینی:</span>
                  <span class="font-medium">{{ formatDate(prediction.date) }}</span>
                </div>
                <div class="flex justify-between text-sm">
                  <span class="text-gray-500">تاثیر تخفیف:</span>
                  <span class="font-medium">{{ prediction.discountImpact }}%</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- نمودار رفتار -->
        <div class="border border-gray-200 rounded-lg p-6">
          <h5 class="font-medium text-gray-900 mb-4">نمودار رفتار کاربران</h5>
          <div class="h-64 flex items-center justify-center text-gray-400">[نمودار تحلیل رفتار کاربران]</div>
        </div>
      </div>

      <!-- تنظیمات پیشرفته -->
      <div v-if="activeTab === 'settings'" class="space-y-6">
        <div class="bg-purple-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-purple-900 mb-2">تنظیمات پیشرفته</h4>
          <p class="text-sm text-purple-700">تنظیمات دقیق و پیشرفته سیستم تخفیف‌های پویا</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- تنظیمات عملکرد -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">تنظیمات عملکرد</h5>
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">فاصله زمانی بررسی (دقیقه)</label>
                <input v-model="advancedSettings.checkInterval" type="number" min="1" max="60" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر قوانین همزمان</label>
                <input v-model="advancedSettings.maxConcurrentRules" type="number" min="1" max="20" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
              </div>
              
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">کش کردن نتایج</span>
                  <p class="text-xs text-gray-500">ذخیره نتایج برای بهبود سرعت</p>
                </div>
                <input v-model="advancedSettings.cacheResults" type="checkbox" class="rounded border-gray-300">
              </div>
              
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">لاگ کامل</span>
                  <p class="text-xs text-gray-500">ثبت تمام تصمیم‌گیری‌ها</p>
                </div>
                <input v-model="advancedSettings.fullLogging" type="checkbox" class="rounded border-gray-300">
              </div>
            </div>
          </div>

          <!-- تنظیمات امنیتی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">تنظیمات امنیتی</h5>
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">تشخیص تقلب</span>
                  <p class="text-xs text-gray-500">تشخیص رفتارهای مشکوک</p>
                </div>
                <input v-model="advancedSettings.fraudDetection" type="checkbox" class="rounded border-gray-300">
              </div>
              
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">محدودیت نرخ</span>
                  <p class="text-xs text-gray-500">محدود کردن درخواست‌ها</p>
                </div>
                <input v-model="advancedSettings.rateLimiting" type="checkbox" class="rounded border-gray-300">
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تخفیف روزانه</label>
                <input v-model="advancedSettings.maxDailyDiscount" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">آستانه هشدار</label>
                <input v-model="advancedSettings.alertThreshold" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
              </div>
            </div>
          </div>
        </div>

        <!-- تنظیمات اعلان‌ها -->
        <div class="border border-gray-200 rounded-lg p-6">
          <h5 class="font-medium text-gray-900 mb-4">تنظیمات اعلان‌ها</h5>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">هشدار عملکرد</span>
                  <p class="text-xs text-gray-500">هشدار در صورت کاهش عملکرد</p>
                </div>
                <input v-model="advancedSettings.performanceAlert" type="checkbox" class="rounded border-gray-300">
              </div>
              
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">هشدار تقلب</span>
                  <p class="text-xs text-gray-500">هشدار در صورت تشخیص تقلب</p>
                </div>
                <input v-model="advancedSettings.fraudAlert" type="checkbox" class="rounded border-gray-300">
              </div>
            </div>
            
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">گزارش روزانه</span>
                  <p class="text-xs text-gray-500">ارسال گزارش عملکرد روزانه</p>
                </div>
                <input v-model="advancedSettings.dailyReport" type="checkbox" class="rounded border-gray-300">
              </div>
              
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">هشدار بودجه</span>
                  <p class="text-xs text-gray-500">هشدار نزدیک به سقف بودجه</p>
                </div>
                <input v-model="advancedSettings.budgetAlert" type="checkbox" class="rounded border-gray-300">
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'

const activeTab = ref('rules')
const showRuleForm = ref(false)
const selectedRules = ref<number[]>([])
const filterPriority = ref('')
const filterStatus = ref('')

const tabs = [
  { value: 'rules', label: 'قوانین تخفیف' },
  { value: 'engine', label: 'موتور هوشمند' },
  { value: 'behavior', label: 'تحلیل رفتار' },
  { value: 'settings', label: 'تنظیمات پیشرفته' }
]

const stats = ref({
  activeRules: 15,
  totalDiscount: 45000000,
  affectedUsers: 2340,
  conversionRate: 23.5
})

const rules = ref([
  {
    id: 1,
    name: 'تخفیف کاربران جدید',
    type: 'behavioral',
    priority: 'high',
    status: 'active',
    color: '#3b82f6',
    description: 'تخفیف ۱۰٪ برای کاربرانی که برای اولین بار خرید می‌کنند',
    conditions: ['کاربر جدید', 'اولین خرید'],
    discount: 10,
    usageCount: 156,
    lastExecuted: '2024-01-15T10:30:00'
  },
  {
    id: 2,
    name: 'تخفیف سبد خرید بالا',
    type: 'cart',
    priority: 'medium',
    status: 'active',
    color: '#10b981',
    description: 'تخفیف ۱۵٪ برای سبد خرید بالای ۵۰۰ هزار تومان',
    conditions: ['سبد خرید > 500K', 'کاربر ثبت‌نام شده'],
    discount: 15,
    usageCount: 89,
    lastExecuted: '2024-01-14T15:45:00'
  }
])

const engineSettings = reactive({
  algorithm: 'hybrid',
  accuracy: 85,
  maxDiscount: 25,
  autoLearning: true
})

const decisionFactors = ref([
  { name: 'تاریخچه خرید', description: 'الگوی خریدهای قبلی', weight: 30 },
  { name: 'رفتار مرور', description: 'صفحات مشاهده شده', weight: 25 },
  { name: 'زمان روز', description: 'ساعت و روز هفته', weight: 15 },
  { name: 'محل جغرافیایی', description: 'موقعیت کاربر', weight: 10 },
  { name: 'دستگاه', description: 'نوع دستگاه و مرورگر', weight: 10 },
  { name: 'منبع ترافیک', description: 'نحوه ورود به سایت', weight: 10 }
])

const enginePerformance = ref({
  accuracy: 87,
  responseTime: 45,
  throughput: 1250
})

const behaviorPatterns = ref([
  {
    name: 'خرید در آخر هفته',
    description: 'کاربران تمایل به خرید در روزهای آخر هفته دارند',
    confidence: 92,
    userCount: 1250,
    repeatRate: 78
  },
  {
    name: 'مقایسه محصولات',
    description: 'کاربران قبل از خرید چندین محصول را مقایسه می‌کنند',
    confidence: 88,
    userCount: 890,
    repeatRate: 65
  }
])

const predictions = ref([
  {
    id: 1,
    title: 'افزایش تقاضا در تعطیلات',
    description: 'پیش‌بینی افزایش ۳۰٪ تقاضا در هفته آینده',
    probability: 85,
    date: '2024-01-20T00:00:00',
    discountImpact: 12
  },
  {
    id: 2,
    title: 'محبوبیت محصولات جدید',
    description: 'محصولات جدید احتمالاً ۲۰٪ بیشتر فروش خواهند داشت',
    probability: 72,
    date: '2024-01-25T00:00:00',
    discountImpact: 8
  }
])

const advancedSettings = reactive({
  checkInterval: 5,
  maxConcurrentRules: 10,
  cacheResults: true,
  fullLogging: true,
  fraudDetection: true,
  rateLimiting: true,
  maxDailyDiscount: 1000000,
  alertThreshold: 80,
  performanceAlert: true,
  fraudAlert: true,
  dailyReport: true,
  budgetAlert: true
})

const filteredRules = computed(() => {
  let filtered = rules.value
  
  if (filterPriority.value) {
    filtered = filtered.filter(r => r.priority === filterPriority.value)
  }
  
  if (filterStatus.value) {
    filtered = filtered.filter(r => r.status === filterStatus.value)
  }
  
  return filtered
})

const formatDate = (date: string): string => {
  return new Intl.DateTimeFormat('fa-IR').format(new Date(date))
}

const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const getStatusClass = (status: string): string => {
  const classes = {
    active: 'bg-green-100 text-green-800',
    inactive: 'bg-red-100 text-red-800',
    draft: 'bg-yellow-100 text-yellow-800'
  }
  return classes[status as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string): string => {
  const texts = {
    active: 'فعال',
    inactive: 'غیرفعال',
    draft: 'پیش‌نویس'
  }
  return texts[status as keyof typeof texts] || status
}

const getConfidenceClass = (confidence: number): string => {
  if (confidence >= 90) return 'bg-green-100 text-green-800'
  if (confidence >= 70) return 'bg-yellow-100 text-yellow-800'
  return 'bg-red-100 text-red-800'
}

const getPredictionClass = (probability: number): string => {
  if (probability >= 80) return 'bg-green-100 text-green-800'
  if (probability >= 60) return 'bg-yellow-100 text-yellow-800'
  return 'bg-red-100 text-red-800'
}

const filterRules = () => {
  // پیاده‌سازی فیلتر کردن قوانین

}

const bulkAction = (action: string) => {
  if (selectedRules.value.length === 0) {
    alert('لطفاً قوانینی را انتخاب کنید')
    return
  }
  
  switch (action) {
    case 'activate':
      selectedRules.value.forEach(id => {
        const rule = rules.value.find(r => r.id === id)
        if (rule) {
          rule.status = 'active'
        }
      })
      break
    case 'deactivate':
      selectedRules.value.forEach(id => {
        const rule = rules.value.find(r => r.id === id)
        if (rule) {
          rule.status = 'inactive'
        }
      })
      break
  }
  selectedRules.value = []
}

interface Rule {
  id?: number | string
  name?: string
  [key: string]: unknown
}

const editRule = (_rule: Rule) => {

  showRuleForm.value = true
}

const deleteRule = (rule: Rule) => {
  if (confirm(`آیا از حذف قانون "${rule.name}" اطمینان دارید؟`)) {
    const index = rules.value.findIndex(r => r.id === rule.id)
    if (index !== -1) {
      rules.value.splice(index, 1)
    }
  }
}
</script>

<!--
  کامپوننت سیستم تخفیف‌های پویا
  شامل:
  1. مدیریت قوانین تخفیف
  2. موتور هوشمند تصمیم‌گیری
  3. تحلیل رفتار کاربران
  4. تنظیمات پیشرفته
  توضیحات کامل به فارسی در کد
--> 
