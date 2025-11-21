<template>
  <div class="space-y-6">
    <!-- هدر -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h2 class="text-xl font-bold text-gray-900">تنظیمات امنیتی گیفت کارت</h2>
      <p class="text-gray-600 mt-1">مدیریت امنیت داده‌ها و رمزگذاری اطلاعات حساس</p>
    </div>

    <!-- بخش رمزگذاری اطلاعات حساس -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-4">
        <div>
          <h3 class="text-lg font-medium text-gray-900">رمزگذاری اطلاعات حساس</h3>
          <p class="text-gray-500 text-sm mt-1">
            برای افزایش امنیت، اطلاعات حساس گیفت کارت (کد، مبلغ، گیرنده و ...) به صورت رمزگذاری‌شده ذخیره می‌شوند.
          </p>
        </div>
        <div>
          <button
            :class="encryptionEnabled ? 'bg-green-600 hover:bg-green-700' : 'bg-gray-400 hover:bg-gray-500'"
            class="px-4 py-2 text-white text-sm font-medium rounded-lg focus:outline-none focus:ring-2 focus:ring-offset-2"
            @click="toggleEncryption"
          >
            {{ encryptionEnabled ? 'غیرفعال‌سازی رمزگذاری' : 'فعال‌سازی رمزگذاری' }}
          </button>
        </div>
      </div>
      <div class="flex items-center space-x-2 space-x-reverse mt-2">
        <span :class="encryptionEnabled ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
          {{ encryptionEnabled ? 'رمزگذاری فعال است' : 'رمزگذاری غیرفعال است' }}
        </span>
        <svg v-if="encryptionEnabled" class="w-5 h-5 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
        </svg>
        <svg v-else class="w-5 h-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </div>
      <div class="mt-4 text-sm text-gray-500">
        <ul class="list-disc pr-5 space-y-1">
          <li>در صورت فعال بودن، تمام داده‌های حساس به صورت رمزگذاری‌شده در پایگاه داده ذخیره می‌شوند.</li>
          <li>در صورت غیرفعال‌سازی، اطلاعات به صورت متنی ذخیره می‌شوند (توصیه نمی‌شود).</li>
          <li>برای امنیت بیشتر، کلید رمزگذاری را در محیط امن نگهداری کنید و هرگز در کد قرار ندهید.</li>
        </ul>
      </div>
    </div>

    <!-- بخش پشتیبان‌گیری خودکار -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-4">
        <div>
          <h3 class="text-lg font-medium text-gray-900">پشتیبان‌گیری خودکار</h3>
          <p class="text-gray-500 text-sm mt-1">
            برای جلوگیری از از دست رفتن داده‌ها، سیستم به صورت خودکار و دستی از اطلاعات گیفت کارت‌ها بکاپ تهیه می‌کند.
          </p>
        </div>
        <div>
          <button
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            @click="createBackup"
          >
            ایجاد بکاپ دستی
          </button>
        </div>
      </div>
      <div class="flex items-center space-x-2 space-x-reverse mt-2">
        <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-gray-100 text-gray-800">
          آخرین بکاپ: {{ lastBackupTime }}
        </span>
        <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800">
          زمان‌بندی: {{ backupScheduleLabel }}
        </span>
      </div>
      <div class="mt-4">
        <label class="block text-sm font-medium text-gray-700 mb-2">زمان‌بندی بکاپ خودکار</label>
        <select v-model="backupSchedule" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500">
          <option value="daily">روزانه</option>
          <option value="weekly">هفتگی</option>
          <option value="monthly">ماهانه</option>
        </select>
      </div>
      <div class="mt-6">
        <h4 class="text-sm font-bold text-gray-700 mb-2">لیست آخرین بکاپ‌ها</h4>
        <div v-if="backups.length > 0" class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">تاریخ</th>
                <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">ساعت</th>
                <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">وضعیت</th>
                <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="backup in backups" :key="backup.id">
                <td class="px-4 py-2 whitespace-nowrap">{{ formatDate(backup.date) }}</td>
                <td class="px-4 py-2 whitespace-nowrap">{{ formatTime(backup.date) }}</td>
                <td class="px-4 py-2 whitespace-nowrap">
                  <span :class="backup.status === 'success' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                    {{ backup.status === 'success' ? 'موفق' : 'ناموفق' }}
                  </span>
                </td>
                <td class="px-4 py-2 whitespace-nowrap flex gap-2">
                  <button class="text-blue-600 hover:underline text-xs" @click="downloadBackup(backup)">دانلود</button>
                  <button class="text-green-600 hover:underline text-xs" @click="restoreBackup(backup)">بازیابی</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <div v-else class="text-gray-500 text-sm mt-2">هیچ بکاپی وجود ندارد.</div>
      </div>
      <div class="mt-4 text-xs text-gray-500">
        <ul class="list-disc pr-5 space-y-1">
          <li>بکاپ‌ها به صورت رمزگذاری‌شده ذخیره می‌شوند.</li>
          <li>پیشنهاد می‌شود فایل بکاپ را در محل امن نگهداری کنید.</li>
          <li>در صورت بازیابی، اطلاعات فعلی جایگزین می‌شوند.</li>
        </ul>
      </div>
    </div>

    <!-- بخش کنترل دسترسی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="mb-4">
        <h3 class="text-lg font-medium text-gray-900">کنترل دسترسی</h3>
        <p class="text-gray-500 text-sm mt-1">
          تعیین نقش کاربران و سطوح دسترسی (CRUD) برای بخش‌های مختلف سیستم گیفت کارت.
        </p>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">نقش</th>
              <th class="px-4 py-2 text-center text-xs font-medium text-gray-500">مشاهده</th>
              <th class="px-4 py-2 text-center text-xs font-medium text-gray-500">ایجاد</th>
              <th class="px-4 py-2 text-center text-xs font-medium text-gray-500">ویرایش</th>
              <th class="px-4 py-2 text-center text-xs font-medium text-gray-500">حذف</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="role in roles" :key="role.name">
              <td class="px-4 py-2 whitespace-nowrap font-bold">{{ role.label }}</td>
              <td class="px-4 py-2 text-center">
                <input v-model="role.permissions.read" type="checkbox" :disabled="role.name === 'admin'" />
              </td>
              <td class="px-4 py-2 text-center">
                <input v-model="role.permissions.create" type="checkbox" :disabled="role.name === 'admin'" />
              </td>
              <td class="px-4 py-2 text-center">
                <input v-model="role.permissions.update" type="checkbox" :disabled="role.name === 'admin'" />
              </td>
              <td class="px-4 py-2 text-center">
                <input v-model="role.permissions.delete" type="checkbox" :disabled="role.name === 'admin'" />
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="mt-6">
        <h4 class="text-sm font-bold text-gray-700 mb-2">تخصیص نقش به کاربران</h4>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">کاربر</label>
            <select v-model="selectedUser" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500">
              <option v-for="user in users" :key="user.id" :value="user.id">{{ user.name }}</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نقش</label>
            <select v-model="selectedRole" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500">
              <option v-for="role in roles" :key="role.name" :value="role.name">{{ role.label }}</option>
            </select>
          </div>
        </div>
        <div class="mt-4 flex justify-end">
          <button class="px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2" @click="assignRole">
            تخصیص نقش
          </button>
        </div>
      </div>
      <div class="mt-4 text-xs text-gray-500">
        <ul class="list-disc pr-5 space-y-1">
          <li>ادمین همیشه به همه بخش‌ها دسترسی کامل دارد و قابل تغییر نیست.</li>
          <li>تغییر نقش کاربران بلافاصله اعمال می‌شود.</li>
          <li>برای امنیت بیشتر، نقش‌ها را فقط به افراد مورد اعتماد اختصاص دهید.</li>
        </ul>
      </div>
    </div>

    <!-- بخش لاگ امنیتی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-4">
        <div>
          <h3 class="text-lg font-medium text-gray-900">لاگ امنیتی</h3>
          <p class="text-gray-500 text-sm mt-1">
            ثبت و نمایش تمام فعالیت‌های امنیتی شامل ورود/خروج، تغییرات مهم و تلاش‌های ناموفق.
          </p>
        </div>
        <div class="flex gap-2">
          <button class="px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2" @click="exportLogs">
            خروجی لاگ‌ها
          </button>
          <button class="px-4 py-2 bg-red-600 text-white text-sm font-medium rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2" @click="clearLogs">
            پاک کردن لاگ‌ها
          </button>
        </div>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع رویداد</label>
          <select v-model="logFilter.eventType" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500">
            <option value="">همه رویدادها</option>
            <option value="login">ورود</option>
            <option value="logout">خروج</option>
            <option value="create">ایجاد</option>
            <option value="update">ویرایش</option>
            <option value="delete">حذف</option>
            <option value="failed_attempt">تلاش ناموفق</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کاربر</label>
          <select v-model="logFilter.user" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500">
            <option value="">همه کاربران</option>
            <option v-for="user in users" :key="user.id" :value="user.name">{{ user.name }}</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ</label>
          <input v-model="logFilter.date" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500" />
        </div>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">تاریخ و ساعت</th>
              <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">کاربر</th>
              <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">نوع رویداد</th>
              <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">جزئیات</th>
              <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">IP</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="log in filteredLogs" :key="log.id">
              <td class="px-4 py-2 whitespace-nowrap text-sm">{{ formatDateTime(log.timestamp) }}</td>
              <td class="px-4 py-2 whitespace-nowrap text-sm">{{ log.user }}</td>
              <td class="px-4 py-2 whitespace-nowrap">
                <span :class="getEventTypeClass(log.eventType)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getEventTypeLabel(log.eventType) }}
                </span>
              </td>
              <td class="px-4 py-2 text-sm">{{ log.details }}</td>
              <td class="px-4 py-2 whitespace-nowrap text-sm">{{ log.ip }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="mt-4 text-xs text-gray-500">
        <ul class="list-disc pr-5 space-y-1">
          <li>لاگ‌ها به مدت ۳۰ روز نگهداری می‌شوند.</li>
          <li>برای امنیت بیشتر، لاگ‌ها را به صورت منظم بررسی کنید.</li>
          <li>در صورت مشاهده فعالیت مشکوک، فوراً به مدیر سیستم اطلاع دهید.</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// وضعیت رمزگذاری (در حالت واقعی باید از API خوانده شود)
const encryptionEnabled = ref(true)

// تغییر وضعیت رمزگذاری
const toggleEncryption = () => {
  if (encryptionEnabled.value) {
    if (confirm('آیا مطمئن هستید که می‌خواهید رمزگذاری را غیرفعال کنید؟ این کار امنیت داده‌ها را کاهش می‌دهد.')) {
      encryptionEnabled.value = false
    }
  } else {
    if (confirm('آیا مطمئن هستید که می‌خواهید رمزگذاری را فعال کنید؟')) {
      encryptionEnabled.value = true
    }
  }
}

// پشتیبان‌گیری خودکار
const lastBackupTime = ref('۱۴۰۳/۰۳/۲۵ - ۱۲:۳۰')
const backupSchedule = ref('daily')
const backups = ref([
  { id: 1, date: new Date(), status: 'success' },
  { id: 2, date: new Date(Date.now() - 86400000), status: 'success' },
  { id: 3, date: new Date(Date.now() - 2 * 86400000), status: 'failed' }
])

const backupScheduleLabel = computed(() => {
  switch (backupSchedule.value) {
    case 'daily': return 'روزانه'
    case 'weekly': return 'هفتگی'
    case 'monthly': return 'ماهانه'
    default: return 'نامشخص'
  }
})

const createBackup = () => {
  if (confirm('آیا از ایجاد بکاپ جدید اطمینان دارید؟')) {
    // اینجا می‌توانید منطق ایجاد بکاپ واقعی را اضافه کنید
    alert('بکاپ جدید با موفقیت ایجاد شد.')
    lastBackupTime.value = new Date().toLocaleString('fa-IR')
    backups.value.unshift({ id: Date.now(), date: new Date(), status: 'success' })
  }
}

interface Backup {
  date?: Date
  [key: string]: unknown
}

const downloadBackup = (backup: Backup) => {
  if (backup.date) {
    alert('دانلود بکاپ: ' + formatDate(backup.date) + ' - ' + formatTime(backup.date))
  }
  // اینجا می‌توانید منطق دانلود را اضافه کنید
}

const restoreBackup = (_backup: Backup) => {
  if (confirm('آیا از بازیابی این بکاپ اطمینان دارید؟ اطلاعات فعلی جایگزین خواهند شد.')) {
    alert('بازیابی انجام شد!')
    // اینجا می‌توانید منطق بازیابی را اضافه کنید
  }
}

const formatDate = (date: Date) => {
  return date.toLocaleDateString('fa-IR')
}
const formatTime = (date: Date) => {
  return date.toLocaleTimeString('fa-IR', { hour: '2-digit', minute: '2-digit' })
}

// کنترل دسترسی
const roles = ref([
  { name: 'admin', label: 'مدیر (ادمین)', permissions: { read: true, create: true, update: true, delete: true } },
  { name: 'operator', label: 'اپراتور', permissions: { read: true, create: true, update: true, delete: false } },
  { name: 'viewer', label: 'مشاهده‌گر', permissions: { read: true, create: false, update: false, delete: false } }
])

const users = ref([
  { id: 1, name: 'علی احمدی', role: 'admin' },
  { id: 2, name: 'فاطمه محمدی', role: 'operator' },
  { id: 3, name: 'حسن رضایی', role: 'viewer' }
])

const selectedUser = ref(users.value[0].id)
const selectedRole = ref(users.value[0].role)

const assignRole = () => {
  const user = users.value.find(u => u.id === selectedUser.value)
  if (user) {
    user.role = selectedRole.value
    alert('نقش کاربر با موفقیت تغییر کرد.')
  }
}

// لاگ امنیتی
const logFilter = ref({
  eventType: '',
  user: '',
  date: ''
})

const securityLogs = ref([
  { id: 1, timestamp: new Date(), user: 'علی احمدی', eventType: 'login', details: 'ورود موفق به سیستم', ip: '192.168.1.100' },
  { id: 2, timestamp: new Date(Date.now() - 3600000), user: 'فاطمه محمدی', eventType: 'create', details: 'ایجاد گیفت کارت جدید', ip: '192.168.1.101' },
  { id: 3, timestamp: new Date(Date.now() - 7200000), user: 'حسن رضایی', eventType: 'failed_attempt', details: 'تلاش ناموفق ورود', ip: '192.168.1.102' },
  { id: 4, timestamp: new Date(Date.now() - 10800000), user: 'علی احمدی', eventType: 'update', details: 'ویرایش تنظیمات امنیتی', ip: '192.168.1.100' }
])

const filteredLogs = computed(() => {
  return securityLogs.value.filter(log => {
    if (logFilter.value.eventType && log.eventType !== logFilter.value.eventType) return false
    if (logFilter.value.user && log.user !== logFilter.value.user) return false
    if (logFilter.value.date) {
      const logDate = new Date(log.timestamp).toISOString().split('T')[0]
      if (logDate !== logFilter.value.date) return false
    }
    return true
  })
})

const getEventTypeClass = (eventType: string) => {
  switch (eventType) {
    case 'login': return 'bg-green-100 text-green-800'
    case 'logout': return 'bg-blue-100 text-blue-800'
    case 'create': return 'bg-purple-100 text-purple-800'
    case 'update': return 'bg-yellow-100 text-yellow-800'
    case 'delete': return 'bg-red-100 text-red-800'
    case 'failed_attempt': return 'bg-red-100 text-red-800'
    default: return 'bg-gray-100 text-gray-800'
  }
}

const getEventTypeLabel = (eventType: string) => {
  switch (eventType) {
    case 'login': return 'ورود'
    case 'logout': return 'خروج'
    case 'create': return 'ایجاد'
    case 'update': return 'ویرایش'
    case 'delete': return 'حذف'
    case 'failed_attempt': return 'تلاش ناموفق'
    default: return 'نامشخص'
  }
}

const formatDateTime = (timestamp: Date) => {
  return timestamp.toLocaleString('fa-IR')
}

const exportLogs = () => {
  alert('خروجی لاگ‌ها در حال دانلود است...')
  // اینجا می‌توانید منطق خروجی لاگ‌ها را اضافه کنید
}

const clearLogs = () => {
  if (confirm('آیا از پاک کردن تمام لاگ‌ها اطمینان دارید؟ این کار قابل بازگشت نیست.')) {
    securityLogs.value = []
    alert('لاگ‌ها با موفقیت پاک شدند.')
  }
}

// مستندسازی فارسی:
// این کامپوننت برای مدیریت رمزگذاری اطلاعات حساس گیفت کارت استفاده می‌شود.
// وضعیت رمزگذاری به کاربر نمایش داده می‌شود و امکان فعال/غیرفعال‌سازی وجود دارد.
// توصیه می‌شود رمزگذاری همیشه فعال باشد و کلید رمزگذاری در محیط امن نگهداری شود.
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
