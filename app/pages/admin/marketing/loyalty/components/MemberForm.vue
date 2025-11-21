<template>
  <div v-if="member !== undefined" class="fixed inset-0 bg-gray-600 bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg shadow-lg p-8 w-full max-w-md relative">
      <button class="absolute left-4 top-6 text-gray-400 hover:text-gray-600 text-2xl" @click="$emit('cancel')">×</button>
      <h3 class="text-xl font-bold text-gray-900 mb-6">{{ member ? 'ویرایش عضو' : 'افزودن عضو جدید' }}</h3>
      <form @submit.prevent="saveMember">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700">نام</label>
            <input v-model="form.name" type="text" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500" required>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">ایمیل</label>
            <input v-model="form.email" type="email" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500" required>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">سطح</label>
            <select v-model="form.level" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500">
              <option value="برنزی">برنزی</option>
              <option value="نقره‌ای">نقره‌ای</option>
              <option value="طلایی">طلایی</option>
              <option value="الماس">الماس</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">امتیاز کل</label>
            <input v-model.number="form.totalPoints" type="number" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">امتیاز استفاده شده</label>
            <input v-model.number="form.usedPoints" type="number" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">وضعیت</label>
            <select v-model="form.status" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500">
              <option value="active">فعال</option>
              <option value="inactive">غیرفعال</option>
            </select>
          </div>
        </div>
        <div class="flex justify-end space-x-3 space-x-reverse mt-6">
          <button type="button" class="bg-gray-300 text-gray-700 px-4 py-2 rounded-md hover:bg-gray-400 transition-colors" @click="$emit('cancel')">انصراف</button>
          <button type="submit" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors">ذخیره</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, watch } from 'vue'

const props = defineProps({
  member: { type: Object, default: null }
})
const emit = defineEmits(['save', 'cancel'])

const form = reactive({
  name: '',
  email: '',
  level: 'برنزی',
  totalPoints: 0,
  usedPoints: 0,
  status: 'active'
})

watch(() => props.member, (val) => {
  if (val) {
    form.name = val.name
    form.email = val.email
    form.level = val.level
    form.totalPoints = val.totalPoints
    form.usedPoints = val.usedPoints
    form.status = val.status
  } else {
    form.name = ''
    form.email = ''
    form.level = 'برنزی'
    form.totalPoints = 0
    form.usedPoints = 0
    form.status = 'active'
  }
}, { immediate: true })

function saveMember() {
  emit('save', { ...form })
}
</script> 
