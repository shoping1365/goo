<template>
  <div class="bg-white rounded-xl shadow-lg border-4 border-blue-200 hover:shadow-xl transition-all duration-300 overflow-hidden max-w-md">
    <!-- هدر کارت -->
    <div class="p-3 border-b border-gray-100">
      <div class="flex items-start justify-between mb-2">
        <div class="flex-1">
          <h3 class="text-lg font-bold text-gray-900 mb-1">{{ pattern.name }}</h3>
          <p class="text-sm text-gray-600 line-clamp-2">{{ pattern.description }}</p>
        </div>
        <div class="flex-shrink-0 mr-3">
          <span 
            :class="[
              'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
              pattern.status === 'active' 
                ? 'bg-green-100 text-green-800' 
                : pattern.status === 'inactive'
                ? 'bg-red-100 text-red-800'
                : 'bg-yellow-100 text-yellow-800'
            ]"
          >
            {{ pattern.status === 'active' ? 'فعال' : pattern.status === 'inactive' ? 'غیرفعال' : 'پیش‌نویس' }}
          </span>
        </div>
      </div>
      
      <!-- اطلاعات درگاه -->
      <div class="flex items-center text-sm text-gray-500 mb-2">
        <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
        </svg>
        {{ pattern.gatewayName }}
      </div>

      <!-- آمار -->
      <div class="flex items-center space-x-4 space-x-reverse text-sm">
        <div class="flex items-center text-blue-600">
          <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
          </svg>
          {{ pattern.usageCount ?? 0 }} استفاده
        </div>
        <div class="flex items-center text-gray-500">
          <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          {{ pattern.patternCode }}
        </div>
      </div>
    </div>

    <!-- محتوای پترن -->
    <div class="p-3">
      <div class="mb-3">
        <h4 class="text-sm font-medium text-gray-700 mb-1">محتوا:</h4>
        <div class="bg-white border border-gray-200 rounded-lg p-2 text-sm text-gray-800 font-mono">
          {{ pattern.content }}
        </div>
      </div>

      <!-- متغیرها -->
      <div v-if="pattern.variables" class="mb-3">
        <h4 class="text-sm font-medium text-gray-700 mb-1">متغیرها:</h4>
        <div class="bg-white border border-gray-200 rounded-lg p-2 text-sm text-gray-800">
          {{ pattern.variables }}
        </div>
      </div>

      <!-- دکمه‌های عملیات -->
      <div class="flex items-center justify-between pt-3 border-t border-gray-100">
        <div class="flex space-x-2 space-x-reverse">
          <button
            class="inline-flex items-center px-3 py-1.5 border border-gray-300 shadow-sm text-xs font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors"
            @click="$emit('edit', pattern)"
          >
            <svg class="w-3 h-3 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
            </svg>
            ویرایش
          </button>
          
          <button
            class="inline-flex items-center px-3 py-1.5 border border-gray-300 shadow-sm text-xs font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors"
            @click="$emit('duplicate', pattern)"
          >
            <svg class="w-3 h-3 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
            </svg>
            کپی
          </button>
        </div>

        <div class="flex space-x-2 space-x-reverse">
          <button
            :disabled="isTesting"
            class="inline-flex items-center px-3 py-1.5 border border-transparent text-xs font-medium rounded-md text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transition-colors disabled:opacity-60"
            @click="testPatternDirect"
          >
            <svg class="w-3 h-3 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
            </svg>
            <span v-if="isTesting">در حال ارسال...</span>
            <span v-else>تست</span>
          </button>
          
          <button
            v-if="!isReserved"
            class="inline-flex items-center px-3 py-1.5 border border-transparent text-xs font-medium rounded-md text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 transition-colors"
            @click="$emit('delete', pattern)"
          >
            <svg class="w-3 h-3 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
            </svg>
            حذف
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRuntimeConfig } from 'nuxt/app'

const config = useRuntimeConfig()
const apiBase = config.public.goApiBase
if (!apiBase) {
  throw new Error('Go API base URL is not configured. Please set NUXT_PUBLIC_GO_API_BASE.')
}
const isTesting = ref(false)
const testResult = ref<string|null>(null)

async function testPatternDirect() {
  isTesting.value = true
  testResult.value = null
  try {
    // مقداردهی اولیه شماره موبایل تست (اینجا به صورت ثابت، قابل سفارشی‌سازی است)
    const phone = prompt('شماره موبایل تست را وارد کنید:')
    if (!phone) {
      isTesting.value = false
      return
    }
    // استخراج متغیرها از pattern.variables
    const variables: Record<string, string> = {}
    if (props.pattern.variables) {
      const vars = props.pattern.variables.split(',').map(v => v.trim()).filter(v => v.startsWith('%') && v.endsWith('%')).map(v => v.slice(1, -1))
      vars.forEach(v => {
        variables[v] = prompt(`مقدار متغیر ${v}:`, '') || `[${v}]`
      })
    }
  const response = await $fetch(`${apiBase}/api/admin/sms-patterns/test`, {
      method: 'POST',
      body: {
        pattern_id: props.pattern.id,
        phone,
        variables
      }
    })
    testResult.value = 'ارسال تست موفقیت‌آمیز بود.'
    alert('ارسال تست موفقیت‌آمیز بود.')
  } catch (e: any) {
    testResult.value = 'خطا در ارسال تست: ' + (e?.data?.message || e?.message || 'خطای ناشناخته')
    alert(testResult.value)
  } finally {
    isTesting.value = false
  }
}
interface Pattern {
  id: number
  name: string
  description: string
  type: 'sms' | 'email' | 'push'
  category: 'order' | 'marketing' | 'notification' | 'verification'
  content: string
  status: 'active' | 'inactive' | 'draft'
  usageCount: number
  maxLength?: number
  gatewayId: number
  gatewayName: string
  patternCode: string
  variables: string
}

interface Props {
  pattern: Pattern
}

const props = defineProps<Props>()

const reservedCodes = ["admin_failover", "admin_test", "otp_login", "admin_order", "low_stock", "security_issue"]
const reservedNames = ["اعلان خطای سیستم", "تست پیامک", "کد تأیید ورود", "اعلان سفارشات به مدیر", "کمبود موجودی محصول", "مشکلات امنیتی"]
const isReserved = computed(() => 
  reservedCodes.includes(props.pattern.patternCode) || 
  reservedNames.includes(props.pattern.name) ||
  props.pattern.patternCode === "13333" // برای OTP که pattern_code آن تغییر کرده
)

defineEmits<{
  edit: [pattern: Pattern]
  duplicate: [pattern: Pattern]
  test: [pattern: Pattern]
  delete: [pattern: Pattern]
}>()
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  line-clamp: 2;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style> 