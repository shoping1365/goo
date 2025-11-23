<template>
  <div v-if="hasAccess && member" class="fixed inset-0 bg-gray-600 bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg shadow-lg p-8 w-full max-w-md relative">
      <button class="absolute left-4 top-6 text-gray-400 hover:text-gray-600 text-2xl" @click="$emit('close')">×</button>
      <div class="flex flex-col items-center mb-6">
        <img :src="member.avatar" :alt="member.name" class="w-20 h-20 rounded-full mb-2">
        <h3 class="text-xl font-bold text-gray-900">{{ member.name }}</h3>
        <p class="text-gray-500">{{ member.email }}</p>
        <span class="mt-2 px-3 py-1 rounded-full text-sm font-semibold" :class="getLevelColor(member.level)">
          {{ member.level }}
        </span>
      </div>
      <div class="space-y-3 mb-6">
        <div class="flex justify-between">
          <span class="text-gray-600">امتیاز کل:</span>
          <span class="font-bold text-gray-900">{{ member.totalPoints.toLocaleString() }}</span>
        </div>
        <div class="flex justify-between">
          <span class="text-gray-600">امتیاز استفاده شده:</span>
          <span class="font-bold text-gray-900">{{ member.usedPoints.toLocaleString() }}</span>
        </div>
        <div class="flex justify-between">
          <span class="text-gray-600">تاریخ عضویت:</span>
          <span class="font-bold text-gray-900">{{ formatDate(member.joinDate) }}</span>
        </div>
        <div class="flex justify-between">
          <span class="text-gray-600">آخرین فعالیت:</span>
          <span class="font-bold text-gray-900">{{ formatDate(member.lastActivity) }}</span>
        </div>
        <div class="flex justify-between">
          <span class="text-gray-600">وضعیت:</span>
          <span :class="member.status === 'active' ? 'text-green-600' : 'text-red-600'" class="font-bold">{{ member.status === 'active' ? 'فعال' : 'غیرفعال' }}</span>
        </div>
      </div>
      <div class="flex justify-end space-x-3 space-x-reverse">
        <button class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors" @click="$emit('edit', member)">ویرایش</button>
        <button class="bg-gray-300 text-gray-700 px-4 py-2 rounded-md hover:bg-gray-400 transition-colors" @click="$emit('close')">بستن</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'

const { user } = useAuth()
const router = useRouter()

const hasAccess = computed(() => {
  return ['admin', 'developer'].includes(user.value?.role || '')
})

watch(hasAccess, (newValue) => {
  if (!newValue) {
    router.push('/404')
  }
})

onMounted(() => {
  if (!hasAccess.value) {
    router.push('/404')
  }
})

defineProps({
  member: { type: Object, required: true }
})

defineEmits(['close', 'edit'])

function getLevelColor(level) {
  const colors = {
    'برنزی': 'bg-yellow-600 text-white',
    'نقره‌ای': 'bg-gray-500 text-white',
    'طلایی': 'bg-yellow-500 text-white',
    'الماس': 'bg-blue-600 text-white'
  }
  return colors[level] || 'bg-gray-400 text-white'
}

function formatDate(dateString) {
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR')
}
</script> 
