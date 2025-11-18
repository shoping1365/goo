<template>
  <form @submit.prevent="saveCompressionSettings" class="space-y-6 bg-white rounded-2xl shadow-lg border border-gray-200 p-6">
    <!-- Quality -->
    <div>
      <label class="block font-bold mb-1">کیفیت</label>
      <select v-model="compressionSettings.quality" class="w-full border border-gray-300 rounded-lg px-3 py-2">
        <option value="smart">هوشمند</option>
        <option value="high">بالا</option>
        <option value="medium">متوسط</option>
        <option value="low">پایین</option>
        <option value="custom">سفارشی</option>
      </select>
      <div v-if="compressionSettings.quality==='custom'" class="mt-2">
        <input type="number" v-model.number="compressionSettings.customQuality" min="1" max="100" class="w-24 border border-gray-300 rounded-lg px-2 py-1" />
        <span class="ml-2 text-sm">%</span>
      </div>
    </div>

    <!-- Format -->
    <div>
      <label class="block font-bold mb-1">فرمت خروجی</label>
      <select v-model="compressionSettings.format" class="w-full border border-gray-300 rounded-lg px-3 py-2">
        <option value="original">همان فرمت اصلی</option>
        <option value="jpeg">JPEG</option>
        <option value="png">PNG</option>
        <option value="webp">WEBP</option>
      </select>
    </div>

    <!-- Checkboxes -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <label class="flex items-center gap-2">
        <input type="checkbox" v-model="compressionSettings.progressive" class="w-4 h-4 text-blue-600 border-gray-300 rounded" />
        <span class="text-sm">JPEG Progressive</span>
      </label>
      <label class="flex items-center gap-2">
        <input type="checkbox" v-model="compressionSettings.removeMetadata" class="w-4 h-4 text-blue-600 border-gray-300 rounded" />
        <span class="text-sm">حذف متادیتا</span>
      </label>
      <label class="flex items-center gap-2">
        <input type="checkbox" v-model="compressionSettings.optimizeColors" class="w-4 h-4 text-blue-600 border-gray-300 rounded" />
        <span class="text-sm">بهینه‌سازی رنگ</span>
      </label>
      <label class="flex items-center gap-2">
        <input type="checkbox" v-model="compressionSettings.createBackup" class="w-4 h-4 text-blue-600 border-gray-300 rounded" />
        <span class="text-sm">ساخت نسخه پشتیبان</span>
      </label>
      <label class="flex items-center gap-2">
        <input type="checkbox" v-model="compressionSettings.keepOriginalDate" class="w-4 h-4 text-blue-600 border-gray-300 rounded" />
        <span class="text-sm">حفظ تاریخ اصلی فایل</span>
      </label>
      <label class="flex items-center gap-2">
        <input type="checkbox" v-model="compressionSettings.autoOrient" class="w-4 h-4 text-blue-600 border-gray-300 rounded" />
        <span class="text-sm">Auto-Orientation</span>
      </label>
      <label class="flex items-center gap-2">
        <input type="checkbox" v-model="compressionSettings.enabled" class="w-4 h-4 text-blue-600 border-gray-300 rounded" />
        <span class="text-sm">فعال</span>
      </label>
    </div>

    <!-- Save Button -->
    <div class="flex justify-end">
      <button type="submit" :disabled="!unsavedChanges" class="bg-gradient-to-l from-pink-400 via-purple-500 to-blue-500 text-white px-8 py-2 rounded-lg font-bold disabled:opacity-50">ذخیره</button>
    </div>
  </form>
</template>

<script setup lang="ts">
import { useImageCompression } from '~/composables/useImageCompression';
const { compressionSettings, unsavedChanges, saveCompressionSettings } = useImageCompression()
</script> 
