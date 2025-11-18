<template>
  <div class="space-y-6" dir="rtl">
    <div class="bg-white p-6 rounded-lg shadow border border-gray-200">
      <h2 class="text-lg font-bold text-gray-900 mb-4">تنظیمات نظرات</h2>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فرمت‌های مجاز تصویر</label>
          <input v-model="local.allowedImageTypes" type="text" class="w-full px-3 py-2 border rounded" placeholder="image/jpeg,image/png,image/webp">
          <p class="text-xs text-gray-500 mt-1">با کاما جدا کنید</p>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فرمت‌های مجاز ویدیو</label>
          <input v-model="local.allowedVideoTypes" type="text" class="w-full px-3 py-2 border rounded" placeholder="video/mp4,video/webm">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تعداد فایل برای هر نظر</label>
          <input v-model.number="local.maxFilesPerReview" type="number" min="0" class="w-full px-3 py-2 border rounded">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر حجم هر فایل (MB)</label>
          <input v-model.number="local.maxFileSizeMb" type="number" min="1" class="w-full px-3 py-2 border rounded">
        </div>
      </div>

      <div class="mt-6 flex gap-3">
        <button @click="onSave" class="px-4 py-2 bg-blue-600 text-white rounded">ذخیره</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'

const props = defineProps<{ config: any }>()
const emit = defineEmits(['save'])

const local = reactive({
  allowedImageTypes: props.config?.allowedImageTypes || 'image/jpeg,image/png,image/webp',
  allowedVideoTypes: props.config?.allowedVideoTypes || 'video/mp4,video/webm',
  maxFilesPerReview: props.config?.maxFilesPerReview || 7,
  maxFileSizeMb: props.config?.maxFileSizeMb || 50,
})

const onSave = () => {
  emit('save', { ...local })
  // ذخیره به بک‌اند Go از طریق نیترو
  $fetch('/api/admin/reviews/settings', {
    method: 'PUT',
    body: { ...local }
  }).catch(() => {})
}
</script>


