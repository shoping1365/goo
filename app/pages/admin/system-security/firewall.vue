<template>
  <div class="p-6" dir="rtl">
    <div class="mb-6 bg-white rounded-lg shadow-md border border-gray-200 p-6">
      <h1 class="text-2xl font-bold text-gray-900 mb-2">فایروال</h1>
      <p class="text-gray-600">مدیریت قوانین فایروال و محافظت از سیستم</p>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <TemplateCard
        title="وضعیت فایروال"
        :value="firewallStatus"
        variant="blue"
      />
      <TemplateCard
        title="قوانین فعال"
        :value="activeRules"
        variant="green"
      />
      <TemplateCard
        title="درخواست‌های مسدود شده"
        :value="blockedRequests"
        variant="red"
      />
      <TemplateCard
        title="تهدیدات شناسایی شده"
        :value="detectedThreats"
        variant="yellow"
      />
    </div>

    <!-- کنترل‌های فایروال -->
    <div class="bg-white rounded-2xl shadow-lg border border-gray-200 mb-8 p-6">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 items-center">
        <!-- وضعیت فایروال -->
        <div>
          <div class="text-xs text-gray-500 mb-2 font-medium">وضعیت فایروال</div>
          <div class="flex gap-2">
            <button @click="toggleFirewall(true)" :class="firewallEnabled ? 'bg-green-500 text-white' : 'bg-gray-200 text-gray-700'" class="px-4 py-1 rounded-full text-sm font-bold transition-colors">فعال</button>
            <button @click="toggleFirewall(false)" :class="!firewallEnabled ? 'bg-gray-400 text-white' : 'bg-gray-200 text-gray-700'" class="px-4 py-1 rounded-full text-sm font-bold transition-colors">غیرفعال</button>
          </div>
        </div>
        <!-- حالت امنیتی -->
        <div>
          <div class="text-xs text-gray-500 mb-2 font-medium">حالت امنیتی</div>
          <select v-model="securityMode" @change="changeSecurityMode" class="w-full px-3 py-2 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-400 text-sm bg-gray-50">
            <option value="low">کم</option>
            <option value="medium">متوسط</option>
            <option value="high">زیاد</option>
            <option value="strict">سختگیرانه</option>
          </select>
        </div>
        <!-- عملیات سریع -->
        <div>
          <div class="flex gap-6 justify-end">
            <button @click="clearLogs" class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-gray-400 to-gray-600 hover:from-gray-500 hover:to-gray-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105">
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
              </svg>
              پاک کردن لاگ‌ها
            </button>
            <button @click="backupRules" class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105">
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4"></path>
              </svg>
              پشتیبان‌گیری از قوانین
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودار ترافیک -->
    <div class="bg-white p-6 rounded-lg shadow-md border border-gray-200 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">نمودار ترافیک فایروال</h3>
      <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
        <p class="text-gray-500">نمودار ترافیک بر اساس لاگ‌ها (placeholder)</p>
      </div>
    </div>

    <!-- قوانین فایروال -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 mb-8">
      <div class="px-6 py-4 border-b border-gray-200 flex justify-between items-center">
        <h3 class="text-lg font-semibold text-gray-900">قوانین فایروال</h3>
        <button @click="showAddRuleModal = true" class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-green-400 to-green-600 hover:from-green-500 hover:to-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105">
          <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
          </svg>
          افزودن قانون جدید
        </button>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام قانون</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">منبع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مقصد</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">پورت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="rule in firewallRules" :key="rule.name">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ rule.name }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ rule.type }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ rule.source }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ rule.destination }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ rule.port }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ rule.action }}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="rule.enabled ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ rule.enabled ? 'فعال' : 'غیرفعال' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button @click="toggleRule(rule.name)" class="text-blue-600 hover:text-blue-900 ml-2">
                  {{ rule.enabled ? 'غیرفعال' : 'فعال' }}
                </button>
                <button @click="openEdit(rule)" class="text-green-600 hover:text-green-900 ml-2">ویرایش</button>
                <button v-if="hasPermission('security_firewall.write')" @click="deleteRule(rule.name)" class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- لاگ‌های فایروال -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">لاگ‌های فایروال</h3>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">زمان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">IP منبع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">IP مقصد</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">پورت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع ترافیک</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="log in firewallLogs" :key="log.id" :class="getLogRowClass(log.action)">
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ log.timestamp }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ log.sourceIP }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ log.destinationIP }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ log.port }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ log.trafficType }}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getActionClass(log.action)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ getActionText(log.action) }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- مودال افزودن قانون -->
    <div v-if="showAddRuleModal" class="fixed inset-0 bg-black bg-opacity-50 backdrop-blur-sm flex items-center justify-center z-50 p-6">
      <div class="bg-white rounded-2xl shadow-2xl w-full max-w-lg transform transition-all">
        <!-- Header -->
        <div class="px-6 py-4 border-b border-gray-100 flex items-center justify-between">
          <h3 class="text-lg font-bold text-gray-900">{{ isEdit ? 'ویرایش قانون' : 'افزودن قانون جدید' }}</h3>
          <button @click="showAddRuleModal = false" class="p-2 hover:bg-gray-100 rounded-lg transition-colors">
            <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>
        
        <!-- Content -->
        <div class="px-6 py-6 space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام قانون</label>
            <input v-model="newRule.name" :disabled="isEdit" type="text" placeholder="نام قانون" class="w-full px-4 py-3 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors">
          </div>
          
          <div class="grid grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع</label>
              <select v-model="newRule.type" class="w-full px-4 py-3 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors bg-white">
                <option value="inbound">ورودی</option>
                <option value="outbound">خروجی</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">عملیات</label>
              <select v-model="newRule.action" class="w-full px-4 py-3 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors bg-white">
                <option value="allow">اجازه</option>
                <option value="deny">رد</option>
                <option value="drop">حذف</option>
              </select>
            </div>
          </div>
          
          <div class="grid grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">IP منبع</label>
              <input v-model="newRule.source" type="text" placeholder="0.0.0.0/0" class="w-full px-4 py-3 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">IP مقصد</label>
              <input v-model="newRule.destination" type="text" placeholder="0.0.0.0/0" class="w-full px-4 py-3 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors">
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">پورت</label>
            <input v-model="newRule.port" type="text" placeholder="80,443" class="w-full px-4 py-3 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors">
          </div>
        </div>
        
        <!-- Footer -->
        <div class="px-6 py-4 border-t border-gray-100 flex justify-end gap-3">
          <button @click="showAddRuleModal = false" class="px-6 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors font-medium">
            انصراف
          </button>
          <button @click="isEdit ? updateRule() : addRule()" class="inline-flex items-center px-6 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-green-400 to-green-600 hover:from-green-500 hover:to-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 shadow-md transition-all duration-200 hover:shadow-lg">
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
            {{ isEdit ? 'ذخیره تغییرات' : 'افزودن قانون' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
declare const useAuth: () => { user: { id?: number; name?: string; email?: string } | null; hasPermission: (perm: string) => boolean }
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; body?: unknown }) => Promise<T>
</script>

<script setup lang="ts">
import TemplateCard from '@/components/common/TemplateCard.vue';
import { onMounted, ref } from 'vue';

definePageMeta({ layout: 'admin-main' })

// استفاده از کامپوزابل احراز هویت
// Auth disabled
const { hasPermission } = useAuth()

type FirewallRule = {
  name: string
  type: string
  source: string
  destination: string
  port: string
  action: string
  enabled: boolean
}

type FirewallLog = {
  id?: number
  timestamp?: string
  source?: string
  sourceIP?: string
  destination?: string
  destinationIP?: string
  port?: string
  trafficType?: string
  action: string
  reason?: string
}

type FirewallStatus = {
  enabled?: boolean
  mode?: string
}

type FirewallResponse<T> = {
  data?: T
  success?: boolean
  message?: string
}

type FirewallRuleAPI = {
  name: string
  direction: 'inbound' | 'outbound'
  source: string
  destination: string
  port: string
  action: 'allow' | 'deny'
  enabled: boolean
}

// داده‌های نمونه
const firewallStatus = ref('—')
const activeRules = ref(0)
const blockedRequests = ref(0)
const detectedThreats = ref(0)

const firewallEnabled = ref(false)
const securityMode = ref('medium')
const showAddRuleModal = ref(false)
const isEdit = ref(false)

const newRule = ref({
  name: '',
  type: 'inbound',
  action: 'allow',
  source: '',
  destination: '',
  port: ''
})

const firewallRules = ref<FirewallRule[]>([])

const firewallLogs = ref<FirewallLog[]>([])

async function toggleFirewall(enabled: boolean) {
  await $fetch('/api/admin/system-security/firewall/toggle', { method: 'POST', body: { enabled } })
  await fetchStatus()
}

async function changeSecurityMode() {
  await $fetch('/api/admin/system-security/firewall/mode', { method: 'POST', body: { mode: securityMode.value } })
  await fetchStatus()
}

async function clearLogs() {
  if (!confirm('آیا از پاک کردن لاگ‌ها اطمینان دارید؟')) return
  await $fetch('/api/admin/system-security/firewall/logs', { method: 'DELETE' })
  await fetchLogs()
}

async function backupRules() {
  await $fetch('/api/admin/system-security/firewall/backup', { method: 'POST' })
}

async function toggleRule(name: string) {
  const rule = firewallRules.value.find(r => r.name === name)
  if (!rule) return
  await $fetch(`/api/admin/system-security/firewall/rules/${encodeURIComponent(name)}/toggle?enabled=${rule.enabled ? '0' : '1'}`, { method: 'PUT' })
  await fetchRules()
}

function openCreate() {
  isEdit.value = false
  showAddRuleModal.value = true
}

function openEdit(rule: FirewallRule) {
  isEdit.value = true
  showAddRuleModal.value = true
  newRule.value = {
    name: rule.name,
    type: rule.type === 'ورودی' ? 'inbound' : 'outbound',
    action: rule.action === 'اجازه' ? 'allow' : rule.action === 'رد' ? 'deny' : 'drop',
    source: rule.source,
    destination: rule.destination,
    port: rule.port,
  }
}
async function updateRule() {
  await $fetch(`/api/admin/system-security/firewall/rules/${encodeURIComponent(newRule.value.name)}`, { method: 'PUT', body: {
    direction: newRule.value.type,
    action: newRule.value.action,
    source: newRule.value.source,
    destination: newRule.value.destination,
    port: newRule.value.port,
  }})
  showAddRuleModal.value = false
  await fetchRules()
}

async function deleteRule(name: string) {
  if (!confirm('آیا از حذف این قانون اطمینان دارید؟')) return
  await $fetch(`/api/admin/system-security/firewall/rules/${encodeURIComponent(name)}`, { method: 'DELETE' })
  await fetchRules()
}

function getLogRowClass(action: string) {
  if (action === 'blocked') {
    return 'bg-red-50'
  }
  return ''
}

function getActionClass(action: string) {
  switch (action) {
    case 'allowed':
      return 'bg-green-100 text-green-800'
    case 'blocked':
      return 'bg-red-100 text-red-800'
    case 'dropped':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

function getActionText(action: string) {
  switch (action) {
    case 'allowed':
      return 'اجازه داده شد'
    case 'blocked':
      return 'مسدود شد'
    case 'dropped':
      return 'حذف شد'
    default:
      return 'نامشخص'
  }
}

async function addRule() {
  if (!newRule.value.name) {
    alert('لطفاً نام قانون را وارد کنید')
    return
  }
  await $fetch('/api/admin/system-security/firewall/rules', { method: 'POST', body: {
    name: newRule.value.name,
    direction: newRule.value.type,
    action: newRule.value.action,
    source: newRule.value.source,
    destination: newRule.value.destination,
    port: newRule.value.port,
  }})
  showAddRuleModal.value = false
  await fetchRules()

  // ریست کردن فرم
  newRule.value = {
    name: '',
    type: 'inbound',
    action: 'allow',
    source: '',
    destination: '',
    port: ''
  }
}

async function fetchStatus() {
  const res = await $fetch<FirewallResponse<FirewallStatus>>('/api/admin/system-security/firewall/status')
  firewallEnabled.value = !!res?.data?.enabled
  firewallStatus.value = firewallEnabled.value ? 'فعال' : 'غیرفعال'
  securityMode.value = res?.data?.mode || 'medium'
}

async function fetchRules() {
  const res = await $fetch<FirewallResponse<FirewallRuleAPI[]>>('/api/admin/system-security/firewall/rules')
  firewallRules.value = (res?.data || []).map((r: FirewallRuleAPI) => ({
    name: r.name,
    type: r.direction === 'inbound' ? 'ورودی' : 'خروجی',
    source: r.source,
    destination: r.destination,
    port: r.port,
    action: r.action === 'allow' ? 'اجازه' : 'رد',
    enabled: r.enabled,
  }))
  activeRules.value = firewallRules.value.filter(r => r.enabled).length
}

async function fetchLogs() {
  const res = await $fetch<FirewallResponse<FirewallLog[]>>('/api/admin/system-security/firewall/logs')
  firewallLogs.value = res?.data || []
  // آمار بالا بر اساس لاگ‌ها
  blockedRequests.value = firewallLogs.value.filter((l: FirewallLog) => l.action === 'blocked').length
  detectedThreats.value = Math.max(0, Math.round(blockedRequests.value * 0.1))
}

onMounted(async () => {
  await Promise.all([fetchStatus(), fetchRules(), fetchLogs()])
})
</script> 
