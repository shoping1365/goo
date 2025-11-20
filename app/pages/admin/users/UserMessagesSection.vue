<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-blue-700 mb-2 flex justify-between items-center">
      <span>پیام‌های کاربر</span>
      <div class="flex gap-2">
        <button class="bg-blue-100 hover:bg-blue-200 text-blue-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="sendNewMessage">
          <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/></svg>
          ارسال پیام جدید
        </button>
        <button v-if="messages.length > 5" class="bg-blue-100 hover:bg-blue-200 text-blue-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="showAll = !showAll">
          <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
          مشاهده همه
        </button>
      </div>
    </div>
    <div class="space-y-4">
      <div v-for="message in displayedMessages" :key="message.id" class="border rounded-lg p-6 hover:bg-gray-50">
        <div class="flex justify-between items-start mb-2">
          <div class="flex items-center gap-2">
            <span
class="px-2 py-1 rounded text-xs font-bold" :class="{
              'bg-green-100 text-green-700': message.type === 'success',
              'bg-red-100 text-red-700': message.type === 'error',
              'bg-blue-100 text-blue-700': message.type === 'info'
            }">{{ message.type === 'success' ? 'موفق' : message.type === 'error' ? 'خطا' : 'اطلاعات' }}</span>
            <span class="text-gray-500 text-sm">{{ message.date }}</span>
          </div>
          <span class="text-gray-500 text-sm">{{ message.sender }}</span>
        </div>
        <p class="text-gray-700">{{ message.content }}</p>
      </div>
    </div>

    <!-- Modal -->
     <ViewAllModal v-model="showAll" title="پیام‌های کاربر">
      <div class="space-y-4">
        <div v-for="message in messages" :key="message.id" class="border rounded-lg p-6 hover:bg-gray-50">
          <div class="flex justify-between items-start mb-2">
            <div class="flex items-center gap-2">
              <span
class="px-2 py-1 rounded text-xs font-bold" :class="{
                'bg-green-100 text-green-700': message.type === 'success',
                'bg-red-100 text-red-700': message.type === 'error',
                'bg-blue-100 text-blue-700': message.type === 'info'
              }">{{ message.type === 'success' ? 'موفق' : message.type === 'error' ? 'خطا' : 'اطلاعات' }}</span>
              <span class="text-gray-500 text-sm">{{ message.date }}</span>
            </div>
            <span class="text-gray-500 text-sm">{{ message.sender }}</span>
          </div>
          <p class="text-gray-700">{{ message.content }}</p>
        </div>
      </div>
    </ViewAllModal>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import ViewAllModal from '~/components/admin/modals/ViewAllModal.vue';
import type { User } from '~/types/user';

interface Message {
  id: number | string;
  type: string;
  content: string;
  date: string;
  sender: string;
}

interface NotificationApiResponse {
  id?: number | string;
  status?: string;
  title?: string;
  body?: string;
  sent_at?: string;
  created_at?: string;
  source?: string;
  [key: string]: unknown;
}

const props = defineProps<{ user: User }>()
const showAll = ref(false)
const messages = ref<Message[]>([])

function toFaDate(d?: string | Date | null) {
  if (!d) return '-'
  try { return new Date(d as string).toLocaleDateString('fa-IR') } catch { return '-' }
}

async function loadMessages(userId?: number | string) {
  if (!userId) return
  const res = await $fetch<NotificationApiResponse[]>(`/api/users/${userId}/notifications?limit=8`, { credentials: 'include' })
  messages.value = (res || []).map(m => ({
    id: m.id || Math.random(), // Ensure ID exists
    type: m.status === 'failed' ? 'error' : (m.status === 'sent' ? 'success' : 'info'),
    content: m.title || m.body || '-',
    date: toFaDate(m.sent_at || m.created_at),
    sender: m.source === 'admin' ? 'ادمین' : 'سیستم'
  }))
}

function sendNewMessage() {
  // TODO: بعداً فرم ارسال پیام اضافه می‌شود
}

const displayedMessages = computed(() => showAll.value ? messages.value : messages.value.slice(0, 5))

watch(() => props.user?.id, (id) => loadMessages(id), { immediate: true })
</script>
