<template>
  <div dir="rtl">
    <div class="bg-white rounded-lg shadow p-6 mb-8 flex items-center justify-between">
      <div class="flex items-center gap-6 w-full justify-between">
        <h1 class="text-2xl font-bold">ثبت درگاه پیامک</h1>
        <button class="inline-flex items-center justify-center gap-2 w-40 h-12 text-base px-0 py-0 rounded-lg bg-gradient-to-r from-purple-400 to-purple-600 text-white font-bold shadow hover:from-purple-500 hover:to-purple-700 transition" @click="$router.back()">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" /></svg>
          بازگشت
        </button>
      </div>
    </div>
    <div class="bg-white rounded-lg shadow p-6 py-2 mt-8 max-w-4xl mx-auto">
      <!-- انتخاب نوع درگاه پیامک -->
      <div class="mb-6">
        <label class="block mb-6 font-semibold">نوع درگاه پیامک</label>
        <select v-model="selectedGateway" class="form-select w-full rounded-lg border border-purple-400 focus:border-purple-500 focus:ring focus:ring-purple-200">
          <option disabled value="">یک درگاه را انتخاب کنید...</option>
          <option v-for="gateway in gateways" :key="gateway.value" :value="gateway.value">
            {{ gateway.label }}
          </option>
        </select>
      </div>
    </div>
    <!-- تنظیمات اختصاصی هر درگاه در کانتینر جدا -->
    <div v-if="selectedGateway" class="bg-white rounded-2xl shadow-lg p-8 mt-8 max-w-4xl mx-auto">
      <h2 class="text-xl font-bold mb-8 flex items-center gap-2">
        <svg class="w-6 h-6 text-purple-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 01.88 7.903A4.001 4.001 0 0112 21a4 4 0 01-4.88-6.097A4 4 0 018 7V5a4 4 0 118 0v2z" /></svg>
        تنظیمات {{ getGatewayLabel(selectedGateway) }}
      </h2>

      <!-- آدرس API -->
      <div class="mb-8">
        <div class="flex flex-row-reverse items-center gap-3 bg-purple-50 border border-purple-200 rounded-lg p-6">
          <label class="font-bold whitespace-nowrap">:آدرس API</label>
          <input v-model="apiUrl" type="text" class="form-input w-full rounded-lg border-gray-300 focus:border-purple-500 focus:ring focus:ring-purple-200 font-mono text-xs select-all text-left" dir="ltr" placeholder="مثال: https://api.example.com/send" />
        </div>
      </div>

      <!-- کلید API -->
      <div class="mb-8">
        <div class="flex flex-row-reverse items-center gap-3 bg-purple-50 border border-purple-200 rounded-lg p-6">
          <label class="font-bold whitespace-nowrap">:کلید API</label>
          <input v-model="apiKey" type="password" maxlength="255" class="form-input w-full rounded-lg border-gray-300 focus:border-purple-500 focus:ring focus:ring-purple-200 font-mono text-xs select-all text-left" dir="ltr" placeholder="کلید API درگاه" />
        </div>
      </div>

      <!-- فیلدهای اضافی هر درگاه -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
        <template v-for="field in getGatewayFields(selectedGateway)" :key="field.name">
          <div class="bg-white border-2 border-purple-300 rounded-xl p-6 flex flex-col gap-2">
            <label :for="field.name" class="block mb-2 font-semibold">{{ field.label }}</label>
            <input
              :id="field.name"
              v-model="form[field.name]"
              :type="field.name === 'password' ? 'password' : 'text'"
              class="form-input w-full rounded-lg border-gray-300 focus:border-purple-500 focus:ring focus:ring-purple-200"
              :placeholder="field.placeholder || ''"
            />
          </div>
        </template>
        <!-- فیلد ارسال بر اساس پترن -->
        <div class="bg-white border-2 border-purple-300 rounded-xl p-6 flex flex-col gap-2">
          <label class="block mb-2 font-semibold">ارسال بر اساس پترن</label>
          <div class="flex gap-6">
            <label class="inline-flex items-center gap-2 cursor-pointer">
              <input v-model="form.pattern_based" type="radio" value="yes" class="form-radio text-purple-500" />
              <span>بله</span>
            </label>
            <label class="inline-flex items-center gap-2 cursor-pointer">
              <input v-model="form.pattern_based" type="radio" value="no" class="form-radio text-purple-500" />
              <span>خیر</span>
            </label>
          </div>
        </div>
      </div>
      <!-- باکس تست ارسال پیامک (فقط بعد از ثبت درگاه) -->
      <div class="bg-purple-50 border-2 border-purple-300 rounded-2xl p-8 mt-8 mb-8 shadow flex flex-col gap-6">
        <div class="font-bold text-purple-700 text-lg mb-2 flex items-center gap-2">
          <svg class="w-5 h-5 text-purple-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" /></svg>
          تست ارسال پیامک
        </div>
        <div class="bg-blue-50 border border-blue-200 rounded-lg p-6 text-blue-700">
          <p class="text-sm">برای تست ارسال پیامک، ابتدا درگاه را ثبت کنید و سپس از صفحه ویرایش استفاده کنید.</p>
        </div>
      </div>
      <!-- دکمه تست اتصال و ثبت درگاه (در انتهای فرم) -->
      <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-6 mt-8">
        <div class="flex items-center gap-6 order-2 md:order-1">
          <button :disabled="!isFormValid" class="inline-flex items-center justify-center gap-2 w-40 h-12 text-base px-0 py-0 rounded-lg bg-gradient-to-r from-blue-400 to-blue-600 text-white font-bold shadow hover:from-blue-500 hover:to-blue-700 transition disabled:opacity-50 disabled:cursor-not-allowed" @click="testConnection">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 11V7m0 8h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
            تست اتصال
          </button>
        </div>
        <div class="flex justify-end order-1 md:order-2 w-full md:w-auto">
          <button :disabled="!selectedGateway || !isFormValid" class="inline-flex items-center justify-center gap-2 w-40 h-12 text-base px-0 py-0 rounded-lg bg-gradient-to-r from-purple-400 to-purple-600 text-white font-bold shadow hover:from-purple-500 hover:to-purple-700 transition disabled:opacity-50 disabled:cursor-not-allowed" @click="submitForm">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" /></svg>
            ثبت درگاه
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// @ts-ignore: Nuxt macro
definePageMeta({ layout: 'admin-main' })

import { computed, ref, watch } from 'vue'
import { useRouter } from 'vue-router'

// لیست انواع درگاه و فیلدهای هرکدام
const gateways = [
  {
    value: 'meli_payamak',
    label: 'ملی پیامک',
    fields: [
      { name: 'username', label: 'نام کاربری', placeholder: 'نام کاربری ملی پیامک' },
      { name: 'password', label: 'رمز عبور', placeholder: 'رمز عبور ملی پیامک' },
      { name: 'from', label: 'شماره ارسال‌کننده', placeholder: 'مثال: 1000xxxxxx' }
    ]
  },
  {
    value: 'farazsms',
    label: 'فراز اس‌ام‌اس',
    fields: [
      { name: 'from', label: 'شماره ارسال‌کننده', placeholder: 'مثال: 1000xxxxxx' }
    ]
  },
  {
    value: 'ippanel',
    label: 'آی‌پی‌پنل',
    fields: [
      { name: 'from', label: 'شماره ارسال‌کننده', placeholder: 'مثال: 1000xxxxxx' }
    ]
  },
  {
    value: 'kavenegar',
    label: 'کاوه‌نگار',
    fields: [
      { name: 'username', label: 'نام کاربری', placeholder: 'نام کاربری کاوه‌نگار' },
      { name: 'password', label: 'رمز عبور', placeholder: 'رمز عبور کاوه‌نگار' }
    ]
  }
]

const selectedGateway = ref('')
const form = ref<Record<string, string>>({})

const apiUrl = ref('')
const apiKey = ref('')

const gatewayApiInfo = {
  meli_payamak: 'https://rest.payamak-panel.com/api/SendSMS/',
  farazsms: 'https://edge.ippanel.com/v1/',
        ippanel: 'https://rest.ippanel.com/v1/messages/patterns/send',
  kavenegar: 'https://api.kavenegar.com/v1/{API-KEY}/sms/send.json'
}
function getGatewayApiInfo(val: string) {
  return gatewayApiInfo[val] || ''
}

watch(selectedGateway, (val) => {
  apiUrl.value = getGatewayApiInfo(val)
  // مقداردهی اولیه فیلدهای فرم بر اساس فیلدهای درگاه انتخابی
  const fields = getGatewayFields(val)
  const newForm: Record<string, string> = {}
  fields.forEach(f => { newForm[f.name] = '' })
  // مقداردهی pattern_based به 'no' به صورت پیش‌فرض
  newForm.pattern_based = 'no'
  form.value = newForm
})

// گرفتن لیبل درگاه بر اساس مقدار
function getGatewayLabel(val: string) {
  return gateways.find(g => g.value === val)?.label || ''
}
// گرفتن فیلدهای هر درگاه
function getGatewayFields(val: string) {
  return gateways.find(g => g.value === val)?.fields || []
}

// اعتبارسنجی ساده: فیلدهای اضافی باید پر باشند
const isFormValid = computed(() => {
  if (!selectedGateway.value) return false
  const fields = getGatewayFields(selectedGateway.value)
  return fields.every(f => form.value[f.name] && form.value[f.name].trim() !== '')
})

// ثبت فرم (نمونه)
const router = useRouter()

async function submitForm() {
  if (!isFormValid.value) return

  const payload = {
    type: selectedGateway.value,
    name: getGatewayLabel(selectedGateway.value),
    api_url: apiUrl.value, // آدرس API
    api_key: apiKey.value, // کلید API
    sender_number: form.value.from, // نگاشت from به sender_number
    username: form.value.username || '', // نام کاربری
    password: form.value.password || '', // رمز عبور
    pattern_based: form.value.pattern_based === 'yes',
    is_active: true,
    priority: 1
  }

  try {
    await $fetch('/api/sms-gateways', {
      method: 'POST',
      body: payload
    })
    alert('درگاه با موفقیت ثبت شد!')
    form.value = {}
    selectedGateway.value = ''
    router.push('/admin/notifications/system-settings')
  } catch {
    alert('خطا در ثبت درگاه!')
  }
}



// تست اتصال (فقط بعد از ثبت درگاه)
const _testConnectionStatus = ref('') // '', 'success', 'error'
function testConnection() {
  alert('برای تست اتصال، ابتدا درگاه را ثبت کنید و سپس از صفحه ویرایش استفاده کنید.')
}

defineExpose({
  getGatewayApiInfo,
  apiUrl,
  apiKey
})
</script> 
