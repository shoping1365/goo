<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h3 class="text-lg font-semibold text-gray-800">ارسال خودکار SMS</h3>
        <p class="text-gray-600 text-sm">تنظیم قوانین و زمان‌بندی ارسال خودکار پیامک</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button 
          @click="testSchedule"
          class="px-4 py-2 bg-green-600 text-white rounded-lg text-sm hover:bg-green-700 transition-colors"
        >
          تست زمان‌بندی
        </button>
      </div>
    </div>

    <!-- Main Settings -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gapx-4 py-4">
      <!-- General Settings -->
      <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
        <h4 class="text-lg font-semibold text-gray-800 mb-4">تنظیمات کلی</h4>
        
        <div class="space-y-4">
          <!-- Enable/Disable -->
          <div class="flex items-center justify-between">
            <label class="text-sm font-medium text-gray-700">فعال‌سازی ارسال خودکار</label>
            <button 
              @click="settings.enabled = !settings.enabled"
              :class="[
                'relative inline-flex h-6 w-11 items-center rounded-full transition-colors',
                settings.enabled ? 'bg-green-600' : 'bg-gray-200'
              ]"
            >
              <span 
                :class="[
                  'inline-block h-4 w-4 transform rounded-full bg-white transition-transform',
                  settings.enabled ? 'translate-x-6' : 'translate-x-1'
                ]"
              ></span>
            </button>
          </div>

          <!-- Delay Days -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              تأخیر ارسال (روز)
            </label>
            <input 
              v-model.number="settings.delayDays"
              type="number"
              min="1"
              max="30"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
            <p class="text-xs text-gray-500 mt-1">تعداد روزهای تأخیر بعد از تکمیل سفارش</p>
          </div>

          <!-- Daily Limit -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              محدودیت روزانه
            </label>
            <input 
              v-model.number="settings.dailyLimit"
              type="number"
              min="1"
              max="1000"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
            <p class="text-xs text-gray-500 mt-1">حداکثر تعداد SMS ارسالی در روز</p>
          </div>
        </div>
      </div>

      <!-- Time Settings -->
      <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
        <h4 class="text-lg font-semibold text-gray-800 mb-4">تنظیمات زمانی</h4>
        
        <div class="space-y-4">
          <!-- Allowed Hours -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              ساعات مجاز ارسال
            </label>
            <div class="grid grid-cols-2 gap-3">
              <div>
                <label class="block text-xs text-gray-500 mb-1">از ساعت</label>
                <input 
                  v-model="settings.allowedHours.start"
                  type="time"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                >
              </div>
              <div>
                <label class="block text-xs text-gray-500 mb-1">تا ساعت</label>
                <input 
                  v-model="settings.allowedHours.end"
                  type="time"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                >
              </div>
            </div>
          </div>

          <!-- Days of Week -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              روزهای هفته مجاز
            </label>
            <div class="grid grid-cols-7 gap-2">
              <label 
                v-for="(day, index) in weekDays" 
                :key="index"
                class="flex items-center justify-center p-2 border border-gray-300 rounded-lg cursor-pointer hover:bg-gray-50"
                :class="{ 'bg-blue-100 border-blue-300': settings.allowedDays.includes(index) }"
              >
                <input 
                  v-model="settings.allowedDays"
                  :value="index"
                  type="checkbox"
                  class="sr-only"
                >
                <span class="text-xs font-medium">{{ day }}</span>
              </label>
            </div>
          </div>

          <!-- Time Zone -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              منطقه زمانی
            </label>
            <select 
              v-model="settings.timezone"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="Asia/Tehran">تهران (UTC+3:30)</option>
              <option value="UTC">UTC</option>
              <option value="Europe/London">لندن (UTC+0)</option>
            </select>
          </div>
        </div>
      </div>
    </div>

    <!-- Rules Management -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <div class="flex items-center justify-between mb-4">
        <h4 class="text-lg font-semibold text-gray-800">قوانین ارسال</h4>
        <button 
          @click="addRule"
          class="px-4 py-2 bg-blue-600 text-white rounded-lg text-sm hover:bg-blue-700 transition-colors"
        >
          افزودن قانون جدید
        </button>
      </div>

      <div class="space-y-4">
        <div 
          v-for="rule in settings.rules" 
          :key="rule.id"
          class="border border-gray-200 rounded-lg px-4 py-4"
        >
          <div class="flex items-center justify-between mb-3">
            <div class="flex items-center space-x-3 space-x-reverse">
              <button 
                @click="rule.enabled = !rule.enabled"
                :class="[
                  'relative inline-flex h-5 w-10 items-center rounded-full transition-colors',
                  rule.enabled ? 'bg-green-600' : 'bg-gray-200'
                ]"
              >
                <span 
                  :class="[
                    'inline-block h-3 w-3 transform rounded-full bg-white transition-transform',
                    rule.enabled ? 'translate-x-5' : 'translate-x-1'
                  ]"
                ></span>
              </button>
              <h5 class="font-medium text-gray-800">{{ rule.name }}</h5>
            </div>
            <button 
              @click="deleteRule(rule.id)"
              class="text-red-600 hover:text-red-800 transition-colors"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
              </svg>
            </button>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
            <div>
              <label class="block text-xs text-gray-500 mb-1">شرط</label>
              <select 
                v-model="rule.condition"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="order_completed">تکمیل سفارش</option>
                <option value="high_value">مبلغ بالا</option>
                <option value="first_time">مشتری جدید</option>
                <option value="vip_customer">مشتری VIP</option>
                <option value="specific_product">محصول خاص</option>
              </select>
            </div>

            <div>
              <label class="block text-xs text-gray-500 mb-1">تأخیر (روز)</label>
              <input 
                v-model.number="rule.delay"
                type="number"
                min="0"
                max="30"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
            </div>

            <div>
              <label class="block text-xs text-gray-500 mb-1">اولویت</label>
              <select 
                v-model="rule.priority"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="low">کم</option>
                <option value="medium">متوسط</option>
                <option value="high">زیاد</option>
              </select>
            </div>
          </div>

          <!-- Rule Conditions -->
          <div v-if="rule.condition === 'high_value'" class="mt-3">
            <label class="block text-xs text-gray-500 mb-1">حداقل مبلغ (تومان)</label>
            <input 
              v-model.number="rule.minAmount"
              type="number"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
          </div>

          <div v-if="rule.condition === 'specific_product'" class="mt-3">
            <label class="block text-xs text-gray-500 mb-1">محصول</label>
            <select 
              v-model="rule.productId"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="">انتخاب محصول</option>
              <option value="1">لپ‌تاپ</option>
              <option value="2">موبایل</option>
              <option value="3">کتاب</option>
            </select>
          </div>
        </div>
      </div>
    </div>

    <!-- Statistics -->
    <div class="grid grid-cols-1 md:grid-cols-4 gapx-4 py-4">
      <div class="bg-blue-50 rounded-lg px-4 py-4">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-3">
            <div class="text-2xl font-bold text-blue-800">{{ stats.scheduledToday }}</div>
            <div class="text-sm text-blue-600">برنامه‌ریزی امروز</div>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg px-4 py-4">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
          </div>
          <div class="mr-3">
            <div class="text-2xl font-bold text-green-800">{{ stats.sentToday }}</div>
            <div class="text-sm text-green-600">ارسال شده امروز</div>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg px-4 py-4">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-3">
            <div class="text-2xl font-bold text-yellow-800">{{ stats.pending }}</div>
            <div class="text-sm text-yellow-600">در انتظار</div>
          </div>
        </div>
      </div>

      <div class="bg-red-50 rounded-lg px-4 py-4">
        <div class="flex items-center">
          <div class="p-2 bg-red-100 rounded-lg">
            <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-3">
            <div class="text-2xl font-bold text-red-800">{{ stats.failed }}</div>
            <div class="text-sm text-red-600">ناموفق</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Save Button -->
    <div class="flex justify-end">
      <button 
        @click="saveSettings"
        class="px-6 py-2 bg-green-600 text-white rounded-lg font-medium hover:bg-green-700 transition-colors"
      >
        ذخیره تنظیمات
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'

interface AutoSettings {
  enabled: boolean
  delayDays: number
  dailyLimit: number
  allowedHours: {
    start: string
    end: string
  }
  allowedDays: number[]
  timezone: string
  rules: Array<{
    id: number
    name: string
    condition: string
    delay: number
    priority: string
    enabled: boolean
    minAmount?: number
    productId?: string
  }>
}

const props = defineProps<{
  settings: AutoSettings
}>()

const emit = defineEmits<{
  'update-settings': [settings: AutoSettings]
  'test-schedule': []
}>()

// Reactive data
const weekDays = ['ش', 'ی', 'د', 'س', 'چ', 'پ', 'ج']

const stats = reactive({
  scheduledToday: 45,
  sentToday: 38,
  pending: 7,
  failed: 2
})

// Methods
const addRule = () => {
  const newRule = {
    id: Date.now(),
    name: 'قانون جدید',
    condition: 'order_completed',
    delay: 3,
    priority: 'medium',
    enabled: true
  }
  props.settings.rules.push(newRule)
}

const deleteRule = (ruleId: number) => {
  const index = props.settings.rules.findIndex(rule => rule.id === ruleId)
  if (index > -1) {
    props.settings.rules.splice(index, 1)
  }
}

const saveSettings = () => {
  emit('update-settings', props.settings)
}

const testSchedule = () => {
  emit('test-schedule')
}

// Lifecycle
onMounted(() => {
  // Initialize default allowed days if not set
  if (!props.settings.allowedDays || props.settings.allowedDays.length === 0) {
    props.settings.allowedDays = [0, 1, 2, 3, 4, 5, 6] // All days
  }
})
</script>

<style scoped>
.space-x-reverse > :not([hidden]) ~ :not([hidden]) {
  --tw-space-x-reverse: 1;
}
</style> 
