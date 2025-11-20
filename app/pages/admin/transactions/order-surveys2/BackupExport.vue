<template>
  <div class="space-y-8">
    <!-- Header -->
    <div>
      <h3 class="text-lg font-semibold text-gray-800">پشتیبان‌گیری و خروجی</h3>
      <p class="text-gray-600 text-sm">مدیریت پشتیبان‌گیری و خروجی گزارشات</p>
    </div>
    <!-- Excel Export -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">خروجی اکسل گزارشات</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4 py-4">
        <div class="space-y-2">
          <button :disabled="exporting" class="w-full px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg text-sm flex items-center justify-center space-x-2 space-x-reverse" @click="exportSurveyResponses">
            <svg v-if="exporting" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <span>{{ exporting ? 'در حال خروجی...' : 'خروجی پاسخ‌های نظرسنجی' }}</span>
          </button>
          <button :disabled="exporting" class="w-full px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg text-sm flex items-center justify-center space-x-2 space-x-reverse" @click="exportSMSLogs">
            <svg v-if="exporting" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <span>{{ exporting ? 'در حال خروجی...' : 'خروجی لاگ پیامک‌ها' }}</span>
          </button>
        </div>
        <div class="space-y-2">
          <button :disabled="exporting" class="w-full px-4 py-2 bg-red-600 hover:bg-red-700 text-white rounded-lg text-sm flex items-center justify-center space-x-2 space-x-reverse" @click="exportErrorLogs">
            <svg v-if="exporting" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <span>{{ exporting ? 'در حال خروجی...' : 'خروجی خطاهای ارسال' }}</span>
          </button>
          <button :disabled="exporting" class="w-full px-4 py-2 bg-purple-600 hover:bg-purple-700 text-white rounded-lg text-sm flex items-center justify-center space-x-2 space-x-reverse" @click="exportAllData">
            <svg v-if="exporting" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <span>{{ exporting ? 'در حال خروجی...' : 'خروجی همه داده‌ها' }}</span>
          </button>
        </div>
      </div>
    </div>
    <!-- Backup Management -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">مدیریت پشتیبان‌گیری</h4>
      <div class="flex flex-col md:flex-row gap-4 py-4 items-center">
        <button :disabled="backingUp" class="px-6 py-2 bg-orange-600 hover:bg-orange-700 text-white rounded-lg text-sm flex items-center space-x-2 space-x-reverse" @click="createBackup">
          <svg v-if="backingUp" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <span>{{ backingUp ? 'در حال پشتیبان‌گیری...' : 'ایجاد پشتیبان' }}</span>
        </button>
        <input ref="fileInput" type="file" accept=".json,.zip" class="hidden" @change="restoreBackup">
        <button class="px-6 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg text-sm" @click="fileInput?.click()">بازیابی پشتیبان</button>
      </div>
    </div>
    <!-- Backup History -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">تاریخچه پشتیبان‌گیری</h4>
      <div v-if="backups.length" class="space-y-3">
        <div v-for="backup in backups" :key="backup.id" class="flex items-center justify-between p-3 border border-gray-200 rounded-lg">
          <div>
            <div class="text-sm font-medium text-gray-900">{{ backup.name }}</div>
            <div class="text-xs text-gray-500">{{ formatDate(backup.date) }} - {{ backup.size }}</div>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <button class="text-blue-600 hover:text-blue-800 text-sm" @click="downloadBackup(backup.id)">دانلود</button>
            <button class="text-red-600 hover:text-red-800 text-sm" @click="deleteBackup(backup.id)">حذف</button>
          </div>
        </div>
      </div>
      <div v-else class="text-sm text-gray-500">هیچ پشتیبانی وجود ندارد.</div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref } from 'vue'

// تعریف interface برای backup
interface Backup {
  id: number
  name: string
  date: string
  size: string
}

const exporting = ref(false)
const backingUp = ref(false)
const fileInput = ref<HTMLInputElement>()
const backups = ref<Backup[]>([
  { id: 1, name: 'backup_2024_06_07.json', date: '2024-06-07T10:00:00Z', size: '2.5 MB' },
  { id: 2, name: 'backup_2024_06_06.json', date: '2024-06-06T10:00:00Z', size: '2.3 MB' },
  { id: 3, name: 'backup_2024_06_05.json', date: '2024-06-05T10:00:00Z', size: '2.4 MB' }
])
const exportSurveyResponses = async () => {
  exporting.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 2000))
    alert('فایل اکسل پاسخ‌های نظرسنجی آماده شد!')
  } finally {
    exporting.value = false
  }
}
const exportSMSLogs = async () => {
  exporting.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 2000))
    alert('فایل اکسل لاگ پیامک‌ها آماده شد!')
  } finally {
    exporting.value = false
  }
}
const exportErrorLogs = async () => {
  exporting.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 2000))
    alert('فایل اکسل خطاهای ارسال آماده شد!')
  } finally {
    exporting.value = false
  }
}
const exportAllData = async () => {
  exporting.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 3000))
    alert('فایل اکسل همه داده‌ها آماده شد!')
  } finally {
    exporting.value = false
  }
}
const createBackup = async () => {
  backingUp.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 3000))
    const newBackup: Backup = {
      id: Date.now(),
      name: `backup_${new Date().toISOString().split('T')[0]}.json`,
      date: new Date().toISOString(),
      size: '2.5 MB'
    }
    backups.value.unshift(newBackup)
    alert('پشتیبان با موفقیت ایجاد شد!')
  } finally {
    backingUp.value = false
  }
}
const restoreBackup = (event: Event) => {
  const file = (event.target as HTMLInputElement).files?.[0]
  if (file) {
    if (confirm('آیا مطمئن هستید که می‌خواهید این پشتیبان را بازیابی کنید؟')) {
      alert('پشتیبان با موفقیت بازیابی شد!')
    }
  }
}
const downloadBackup = (id: number) => {
  alert(`دانلود پشتیبان ${id} شروع شد!`)
}
const deleteBackup = (id: number) => {
  if (confirm('آیا مطمئن هستید که می‌خواهید این پشتیبان را حذف کنید؟')) {
    backups.value = backups.value.filter(b => b.id !== id)
    alert('پشتیبان حذف شد!')
  }
}
const formatDate = (date: string) => new Date(date).toLocaleString('fa-IR')
</script> 
