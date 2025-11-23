<template>
  <div v-if="hasAccess" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg shadow-xl w-full max-w-2xl mx-4 max-h-[90vh] overflow-y-auto">
      <div class="p-6 border-b border-gray-200">
        <h2 class="text-lg font-semibold text-gray-900">{{ isEditing ? 'ویرایش کمپین' : 'افزودن کمپین جدید' }}</h2>
        <p class="text-gray-600 mt-1">اطلاعات کمپین تخفیف را وارد کنید</p>
      </div>
      <form class="p-6 space-y-6" @submit.prevent="saveCampaign">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام کمپین</label>
            <input v-model="form.name" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" required>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع کمپین</label>
            <select v-model="form.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
              <option value="seasonal">فصلی</option>
              <option value="flash">فلش</option>
              <option value="loyalty">وفاداری</option>
              <option value="referral">ارجاع</option>
              <option value="custom">سفارشی</option>
            </select>
          </div>
          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
            <textarea v-model="form.description" rows="3" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm"></textarea>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ شروع</label>
            <input v-model="form.startsAt" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" required>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ پایان</label>
            <input v-model="form.endsAt" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" required>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">بودجه (تومان)</label>
            <input v-model.number="form.budget" type="number" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
            <select v-model="form.status" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
              <option value="draft">پیش‌نویس</option>
              <option value="active">فعال</option>
              <option value="paused">متوقف</option>
              <option value="ended">پایان یافته</option>
            </select>
          </div>
        </div>
        <div class="flex gap-3 justify-end">
          <button type="button" class="px-4 py-2 text-gray-700 bg-gray-100 rounded-lg hover:bg-gray-200 transition-colors" @click="$emit('cancel')">لغو</button>
          <button type="submit" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">{{ isEditing ? 'ویرایش' : 'افزودن' }}</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
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

interface Campaign {
  id?: number
  name: string
  type: string
  description: string
  startsAt: string
  endsAt: string
  budget: number
  status: string
}

const props = defineProps<{
  campaign?: Campaign | null
}>()

const emit = defineEmits<{
  cancel: []
  save: [campaign: Campaign]
}>()

const form = ref<Campaign>({
  name: '',
  type: 'seasonal',
  description: '',
  startsAt: '',
  endsAt: '',
  budget: 0,
  status: 'draft'
})

const isEditing = computed(() => !!props.campaign)

onMounted(() => {
  if (!hasAccess.value) {
    router.push('/404')
    return
  }
  if (props.campaign) {
    form.value = { ...props.campaign }
  }
})

function saveCampaign() {
  emit('save', { ...form.value })
}
</script> 
