<template>
  <div class="space-y-6">
    <!-- هدر صفحه -->
    <div class="bg-gradient-to-r from-blue-600 to-purple-600 text-white p-6 rounded-lg">
      <div class="flex justify-between items-start">
        <div class="flex-1">
          <h1 class="text-3xl font-bold mb-2">چت بات هوش مصنوعی</h1>
          <p class="text-blue-100">
            با استفاده از هوش مصنوعی، پاسخ‌های خودکار و هوشمند به مشتریان ارائه دهید. این سیستم قابلیت یادگیری و بهبود مداوم دارد.
          </p>
        </div>
        <div class="flex flex-col space-y-3 space-y-reverse">
          <button class="bg-white text-blue-600 px-6 py-3 rounded-lg hover:bg-blue-50 transition-colors font-medium" @click="openCreateBotModal">
            ایجاد بات جدید
          </button>
          <a href="#" class="text-blue-100 hover:text-white text-sm flex items-center space-x-2 space-x-reverse">
            <span>راهنمای استفاده</span>
            <span>❓</span>
          </a>
        </div>
      </div>
    </div>

    <!-- کارت‌های آمار -->
    <StatsCards 
      :total-bots="totalBots"
      :active-bots="activeBots"
      :total-conversations="totalConversations"
      :satisfaction-rate="satisfactionRate"
    />

    <!-- جدول بات‌ها -->
    <BotsTable 
      :chatbots="chatbots"
      @edit="editBot"
      @test="testBot"
      @analytics="viewAnalytics"
      @delete="deleteBot"
    />

    <!-- مودال ایجاد/ویرایش بات -->
    <BotModal 
      :show="showCreateBotModal"
      :is-editing="isEditingBot"
      :form="botForm"
      @close="closeBotModal"
      @save="saveBot"
      @reset="resetBotForm"
    />
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import BotModal from './components/BotModal.vue'
import BotsTable from './components/BotsTable.vue'
import StatsCards from './components/StatsCards.vue'

definePageMeta({ layout: 'admin-main', middleware: 'admin' })

// متغیرهای reactive
interface Bot {
  id: number | string;
  name: string;
  description?: string;
  type: string;
  icon: string;
  status: string;
  todayConversations: number;
  accuracy: number;
  lastUpdate: string;
  is_active?: boolean;
  updated_at?: string;
  created_at?: string;
  [key: string]: unknown;
}

interface BotForm {
  name: string;
  description: string;
  type: string;
  icon: string;
  aiModel: string;
  temperature: number;
  maxTokens: number;
  personality: string;
  baseKnowledge: string;
  autoRespond: boolean;
  escalateToHuman: boolean;
  learnFromConversations: boolean;
  multiLanguage: boolean;
  workStartTime: string;
  workEndTime: string;
  timezone: string;
  isActive: boolean;
}

const showCreateBotModal = ref(false)
const isEditingBot = ref(false)
const currentBot = ref<Partial<Bot>>({})

// فرم بات
const botForm = ref<BotForm>({
  name: '',
  description: '',
  type: '',
  icon: '🤖',
  aiModel: 'gpt-3.5-turbo',
  temperature: 0.7,
  maxTokens: 500,
  personality: '',
  baseKnowledge: '',
  autoRespond: true,
  escalateToHuman: true,
  learnFromConversations: true,
  multiLanguage: false,
  workStartTime: '09:00',
  workEndTime: '18:00',
  timezone: 'Asia/Tehran',
  isActive: true
})

// لیست ربات‌ها از API
const chatbots = ref<Bot[]>([])
async function loadBots() {
  try {
    const res = await $fetch<{ status: string; data: Record<string, unknown>[] }>('/api/admin/chat/ai-bots')
    if (res?.status === 'success') {
      chatbots.value = ((res.data) || []).map((b) => ({
        id: b.id as number | string,
        name: b.name as string,
        description: b.description as string,
        type: 'AI',
        icon: '🤖',
        status: b.is_active ? 'active' : 'inactive',
        todayConversations: 0,
        accuracy: 0,
        lastUpdate: new Date((b.updated_at || b.created_at) as string).toLocaleDateString('fa-IR')
      }))
    }
  } catch (e) { console.error('loadBots failed', e) }
}

// آمار کلی
const totalBots = computed(() => chatbots.value.length)
const activeBots = computed(() => chatbots.value.filter(bot => bot.status === 'active').length)
const totalConversations = computed(() => chatbots.value.reduce((sum, bot) => sum + (bot.todayConversations || 0), 0))
const satisfactionRate = computed(() => {
  if (chatbots.value.length === 0) return 0
  const avgAccuracy = chatbots.value.reduce((sum, bot) => sum + (bot.accuracy || 0), 0) / chatbots.value.length
  return Math.round(avgAccuracy)
})

// توابع
function openCreateBotModal() {
  isEditingBot.value = false
  currentBot.value = {}
  resetBotForm()
  showCreateBotModal.value = true
}

function editBot(bot: Bot) {
  isEditingBot.value = true
  currentBot.value = { ...bot }
  botForm.value = {
    name: bot.name,
    description: bot.description || '',
    type: bot.type,
    icon: bot.icon,
    aiModel: 'gpt-3.5-turbo',
    temperature: 0.7,
    maxTokens: 500,
    personality: '',
    baseKnowledge: '',
    autoRespond: true,
    escalateToHuman: true,
    learnFromConversations: true,
    multiLanguage: false,
    workStartTime: '09:00',
    workEndTime: '18:00',
    timezone: 'Asia/Tehran',
    isActive: bot.status === 'active'
  }
  showCreateBotModal.value = true
}

function testBot(bot: Bot) {
  alert(`تست بات "${bot.name}" در حال آماده‌سازی است...`)
}

function viewAnalytics(bot: Bot) {
  alert(`آمار بات "${bot.name}" در حال آماده‌سازی است...`)
}

function deleteBot(id) {
  if (!confirm('آیا از حذف این چت بات اطمینان دارید؟')) return
  $fetch(`/api/admin/chat/ai-bots/${id}`, { method: 'DELETE' })
    .then(async () => { await loadBots(); alert('حذف شد') })
    .catch((e) => { console.error('deleteBot failed', e); alert('خطا در حذف') })
}

function closeBotModal() {
  showCreateBotModal.value = false
  isEditingBot.value = false
  currentBot.value = {}
  resetBotForm()
}

function resetBotForm() {
  botForm.value = {
    name: '',
    description: '',
    type: '',
    icon: '🤖',
    aiModel: 'gpt-3.5-turbo',
    temperature: 0.7,
    maxTokens: 500,
    personality: '',
    baseKnowledge: '',
    autoRespond: true,
    escalateToHuman: true,
    learnFromConversations: true,
    multiLanguage: false,
    workStartTime: '09:00',
    workEndTime: '18:00',
    timezone: 'Asia/Tehran',
    isActive: true
  }
}

function saveBot(formData: BotForm) {
  if (!formData.name.trim()) { alert('لطفاً نام بات را وارد کنید'); return }
  if (!formData.type) { alert('لطفاً نوع بات را انتخاب کنید'); return }

  const payload = {
    name: formData.name,
    description: formData.description,
    model: formData.aiModel,
    temperature: formData.temperature,
    max_tokens: formData.maxTokens,
    system_prompt: formData.personality,
    is_active: !!formData.isActive
  }

  const doRequest = async () => {
    if (isEditingBot.value) {
      await $fetch(`/api/admin/chat/ai-bots/${currentBot.value.id}`, { method: 'PUT', body: payload })
    } else {
      await $fetch('/api/admin/chat/ai-bots', { method: 'POST', body: payload })
    }
  }

  doRequest()
    .then(async () => { await loadBots(); alert('ذخیره شد'); closeBotModal() })
    .catch((e) => { console.error('saveBot failed', e); alert('خطا در ذخیره') })
}

onMounted(() => { loadBots() })
</script>
