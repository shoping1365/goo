<template>
  <div class="max-w-xl mx-auto bg-white rounded-lg shadow p-8 mt-8" dir="rtl">
    <h1 class="text-2xl font-bold mb-6">تنظیمات نظرات</h1>
    <form @submit.prevent="saveSettings">
      <div class="mb-4">
        <label class="block mb-1 font-semibold">حداکثر تعداد تصویر مجاز برای هر نظر</label>
        <input v-model.number="settings.maxImages" type="number" min="1" max="10" class="form-input w-full" />
      </div>
      <div class="mb-4">
        <label class="block mb-1 font-semibold">اجازه ثبت نظر توسط مهمان</label>
        <input v-model="settings.allowGuest" type="checkbox" class="form-checkbox ml-2" />
        <span>مهمان‌ها بتوانند نظر ثبت کنند</span>
      </div>
      <div class="mb-4">
        <label class="block mb-1 font-semibold">اجازه آپلود ویدیو</label>
        <input v-model="settings.allowVideo" type="checkbox" class="form-checkbox ml-2" />
        <span>کاربران بتوانند ویدیو ارسال کنند</span>
      </div>
      <div class="mb-4">
        <label class="block mb-1 font-semibold">آپلود فایل فعال باشد</label>
        <input v-model="settings.enableFileUpload" type="checkbox" class="form-checkbox ml-2" />
        <span>کاربران بتوانند فایل ضمیمه کنند</span>
      </div>
      <div class="mb-4">
        <label class="block mb-1 font-semibold">ایمیل الزامی باشد</label>
        <input v-model="settings.requireEmail" type="checkbox" class="form-checkbox ml-2" />
        <span>ایمیل برای ثبت نظر الزامی باشد</span>
      </div>
      <div class="mb-4">
        <label class="block mb-1 font-semibold">تلفن الزامی باشد</label>
        <input v-model="settings.requirePhone" type="checkbox" class="form-checkbox ml-2" />
        <span>تلفن برای ثبت نظر الزامی باشد</span>
      </div>
      <div class="flex justify-end">
        <button type="submit" :disabled="loading" class="px-6 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">{{ loading ? 'در حال ذخیره...' : 'ذخیره تنظیمات' }}</button>
      </div>
    </form>
  </div>
</template>

<script setup>
definePageMeta({ layout: 'admin-main' })
import { onMounted, ref } from 'vue'

const settings = ref({
  maxImages: 5,
  allowGuest: false,
  allowVideo: true,
  enableFileUpload: true,
  requireEmail: false,
  requirePhone: false
})
const loading = ref(false)

async function loadSettings() {
  try {
    const res = await fetch('/api/settings/review', { method: 'GET' })
    if (res.ok) {
      const data = await res.json()
      if (data.value) Object.assign(settings.value, data.value)
    }
  } catch { /* silent */ }
}
onMounted(loadSettings)

async function saveSettings() {
  loading.value = true
  try {
    const res = await fetch('/api/settings/review', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ value: settings.value })
    })
    if (res.ok) {
      alert('تنظیمات با موفقیت ذخیره شد!')
    } else {
      alert('خطا در ذخیره تنظیمات')
    }
  } catch {
    alert('خطا در ارتباط با سرور')
  }
  loading.value = false
}
</script> 
