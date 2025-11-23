<template>
  <div class="p-8">
    <div class="bg-white rounded-lg shadow p-6">
      <h2 class="text-2xl font-bold text-gray-900 mb-6">لاگ‌های تدقیق</h2>

      <!-- Filters -->
      <div class="mb-6 grid grid-cols-1 md:grid-cols-3 gap-4">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="جستجو..."
          class="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
        />
        <select
          v-model="selectedAction"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
        >
          <option value="">تمام فعالیت‌ها</option>
          <option value="LOGIN">ورود</option>
          <option value="LOGOUT">خروج</option>
          <option value="CREATE">ایجاد</option>
          <option value="UPDATE">ویرایش</option>
          <option value="DELETE">حذف</option>
        </select>
        <select
          v-model="selectedUser"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
        >
          <option value="">تمام کاربران</option>
          <option v-for="u in uniqueUsers" :key="u" :value="u">
            {{ u }}
          </option>
        </select>
      </div>

      <!-- Logs Table -->
      <div class="overflow-x-auto">
        <div v-if="loading" class="text-center py-8">
          <p class="text-gray-600">درحال بارگذاری...</p>
        </div>

        <div v-else-if="error" class="p-4 bg-red-100 text-red-700 rounded mb-4">
          {{ error }}
        </div>

        <table v-else class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-right text-sm font-semibold text-gray-900">کاربر</th>
              <th class="px-6 py-3 text-right text-sm font-semibold text-gray-900">فعالیت</th>
              <th class="px-6 py-3 text-right text-sm font-semibold text-gray-900">تفاصیل</th>
              <th class="px-6 py-3 text-right text-sm font-semibold text-gray-900">زمان</th>
              <th class="px-6 py-3 text-right text-sm font-semibold text-gray-900">IP</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="log in filteredLogs" :key="log.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 text-sm font-medium text-gray-900">
                {{ log.user }}
              </td>
              <td class="px-6 py-4 text-sm">
                <span
                  :class="[
                    'px-3 py-1 rounded-full text-xs font-medium',
                    log.action === 'LOGIN'
                      ? 'bg-green-100 text-green-800'
                      : log.action === 'LOGOUT'
                      ? 'bg-yellow-100 text-yellow-800'
                      : log.action === 'DELETE'
                      ? 'bg-red-100 text-red-800'
                      : 'bg-blue-100 text-blue-800',
                  ]"
                >
                  {{ actionLabel(log.action) }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-600">
                {{ log.details }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-600">
                {{ formatTime(log.timestamp) }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-600 font-mono text-xs">
                {{ log.ip }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="mt-6 flex justify-between items-center">
        <div class="text-sm text-gray-600">
          نمایش {{ filteredLogs.length }} از {{ logs.length }} لاگ
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const navigateTo: (to: string) => Promise<void>
</script>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { usePermission } from '@/composables/usePermission'
import { useApiClient } from '@/utils/api'

definePageMeta({ layout: 'admin-main', middleware: 'admin' })

const permission = usePermission()
const { api } = useApiClient()

// بررسی دسترسی
if (!permission.hasPermission('view_audit_logs')) {
  navigateTo('/403')
}

interface AuditLog {
  id: string
  user: string
  action: 'LOGIN' | 'LOGOUT' | 'CREATE' | 'UPDATE' | 'DELETE'
  details: string
  timestamp: string
  ip: string
}

const logs = ref<AuditLog[]>([])
const loading = ref(false)
const error = ref('')
const searchQuery = ref('')
const selectedAction = ref('')
const selectedUser = ref('')

// بارگذاری لاگ‌ها
const fetchLogs = async () => {
  loading.value = true
  error.value = ''

  try {
    const response = await api.get<AuditLog[]>('/api/admin/audit-logs')
    logs.value = response
  } catch (err: unknown) {
    const errorMessage = (err && typeof err === 'object' && 'error' in err)
      ? (err as { error?: string }).error || 'خطا در دریافت لاگ‌ها'
      : 'خطا در دریافت لاگ‌ها'
    error.value = errorMessage
  } finally {
    loading.value = false
  }
}

// بارگذاری هنگام mount
onMounted(async () => {
  await fetchLogs()
})

const uniqueUsers = computed(() => {
  return [...new Set(logs.value.map(log => log.user))]
})

const filteredLogs = computed(() => {
  return logs.value.filter(log => {
    const matchesSearch =
      log.user.includes(searchQuery.value) ||
      log.details.includes(searchQuery.value) ||
      log.ip.includes(searchQuery.value)
    const matchesAction = selectedAction.value === '' || log.action === selectedAction.value
    const matchesUser = selectedUser.value === '' || log.user === selectedUser.value

    return matchesSearch && matchesAction && matchesUser
  })
})

const actionLabel = (action: string): string => {
  const labels: Record<string, string> = {
    LOGIN: 'ورود',
    LOGOUT: 'خروج',
    CREATE: 'ایجاد',
    UPDATE: 'ویرایش',
    DELETE: 'حذف',
  }
  return labels[action] || action
}

const formatTime = (timestamp: string): string => {
  const date = new Date(timestamp)
  return date.toLocaleString('fa-IR')
}
</script>
