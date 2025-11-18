<template>
  <div class="max-w-xl mx-auto bg-white rounded-lg shadow p-8 mt-8" dir="rtl">
    <h1 class="text-2xl font-bold mb-6">تنظیمات پرسش و پاسخ</h1>
    <form @submit.prevent="saveSettings">
      <div class="mb-4">
        <label class="block mb-1 font-semibold">فعال بودن سیستم پرسش و پاسخ</label>
        <input type="checkbox" v-model="settings.enabled" class="form-checkbox ml-2" />
        <span>سیستم فعال باشد</span>
      </div>
      <div class="mb-4">
        <label class="block mb-1 font-semibold">نیاز به تایید مدیر برای نمایش پرسش</label>
        <input type="checkbox" v-model="settings.requireApproval" class="form-checkbox ml-2" />
        <span>پرسش‌ها پس از تایید مدیر نمایش داده شوند</span>
      </div>
      <div class="mb-4">
        <label class="block mb-1 font-semibold">نمایش عمومی پاسخ‌ها</label>
        <input type="checkbox" v-model="settings.publicAnswers" class="form-checkbox ml-2" />
        <span>پاسخ‌ها برای همه کاربران قابل مشاهده باشد</span>
      </div>
      <div class="flex justify-end">
        <button type="submit" class="px-6 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">ذخیره تنظیمات</button>
      </div>
    </form>
  </div>
</template>

<script setup>
definePageMeta({ layout: 'admin-main' })
import { ref } from 'vue'

const settings = ref({
  enabled: true,
  requireApproval: true,
  publicAnswers: true
})

function saveSettings() {
  // اینجا می‌توانید ذخیره‌سازی سمت سرور را اضافه کنید
  alert('تنظیمات با موفقیت ذخیره شد!')
}
</script> 
