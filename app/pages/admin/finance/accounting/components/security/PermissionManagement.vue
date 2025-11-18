<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">مدیریت مجوزهای دسترسی</h4>
        <p class="text-sm text-gray-600 mt-1">مدیریت نقش‌ها و مجوزهای کاربران</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          @click="createRole"
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
          </svg>
          ایجاد نقش جدید
        </button>
      </div>
    </div>

    <!-- آمار مجوزها -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">آمار مجوزها</h5>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">کل نقش‌ها</h6>
            <div class="w-3 h-3 bg-blue-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-blue-600">{{ permissionStats.totalRoles }}</div>
          <div class="text-xs text-gray-500 mt-1">نقش</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">کاربران فعال</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ permissionStats.activeUsers }}</div>
          <div class="text-xs text-gray-500 mt-1">کاربر</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">مجوزهای تعریف شده</h6>
            <div class="w-3 h-3 bg-yellow-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-yellow-600">{{ permissionStats.totalPermissions }}</div>
          <div class="text-xs text-gray-500 mt-1">مجوز</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">آخرین تغییر</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ permissionStats.lastChange }}</div>
          <div class="text-xs text-gray-500 mt-1">ساعت پیش</div>
        </div>
      </div>
    </div>

    <!-- نقش‌های کاربری -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">نقش‌های کاربری</h5>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div
            v-for="role in userRoles"
            :key="role.id"
            class="flex items-center justify-between p-6 border border-gray-200 rounded-lg hover:bg-gray-50"
          >
            <div class="flex items-center">
              <div
                class="w-3 h-3 rounded-full ml-3"
                :class="getRoleStatusColor(role.status)"
              ></div>
              <div>
                <h6 class="text-sm font-medium text-gray-900">{{ role.name }}</h6>
                <p class="text-xs text-gray-500">{{ role.description }}</p>
              </div>
            </div>
            <div class="flex items-center space-x-4 space-x-reverse">
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ role.users }}</div>
                <div class="text-xs text-gray-500">کاربر</div>
              </div>
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ role.permissions }}</div>
                <div class="text-xs text-gray-500">مجوز</div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button
                  @click="viewRole(role)"
                  class="text-blue-600 hover:text-blue-800"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                </button>
                <button
                  @click="editRole(role)"
                  class="text-green-600 hover:text-green-800"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                </button>
                <button
                  @click="deleteRole(role)"
                  class="text-red-600 hover:text-red-800"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- مجوزهای سیستم -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">مجوزهای سیستم</h5>
      </div>
      <div class="p-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <h6 class="text-sm font-medium text-gray-900 mb-3">مدیریت اتصالات</h6>
            <div class="space-y-2">
              <label class="flex items-center">
                <input
                  v-model="permissions.connections.view"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <span class="mr-3 text-sm text-gray-700">مشاهده اتصالات</span>
              </label>
              <label class="flex items-center">
                <input
                  v-model="permissions.connections.create"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <span class="mr-3 text-sm text-gray-700">ایجاد اتصال</span>
              </label>
              <label class="flex items-center">
                <input
                  v-model="permissions.connections.edit"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <span class="mr-3 text-sm text-gray-700">ویرایش اتصال</span>
              </label>
              <label class="flex items-center">
                <input
                  v-model="permissions.connections.delete"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <span class="mr-3 text-sm text-gray-700">حذف اتصال</span>
              </label>
            </div>
          </div>

          <div>
            <h6 class="text-sm font-medium text-gray-900 mb-3">مدیریت داده‌ها</h6>
            <div class="space-y-2">
              <label class="flex items-center">
                <input
                  v-model="permissions.data.sync"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <span class="mr-3 text-sm text-gray-700">همگام‌سازی داده‌ها</span>
              </label>
              <label class="flex items-center">
                <input
                  v-model="permissions.data.backup"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <span class="mr-3 text-sm text-gray-700">پشتیبان‌گیری</span>
              </label>
              <label class="flex items-center">
                <input
                  v-model="permissions.data.restore"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <span class="mr-3 text-sm text-gray-700">بازیابی</span>
              </label>
              <label class="flex items-center">
                <input
                  v-model="permissions.data.cleanup"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <span class="mr-3 text-sm text-gray-700">پاکسازی</span>
              </label>
            </div>
          </div>

          <div>
            <h6 class="text-sm font-medium text-gray-900 mb-3">گزارش‌گیری</h6>
            <div class="space-y-2">
              <label class="flex items-center">
                <input
                  v-model="permissions.reports.view"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <span class="mr-3 text-sm text-gray-700">مشاهده گزارش‌ها</span>
              </label>
              <label class="flex items-center">
                <input
                  v-model="permissions.reports.create"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <span class="mr-3 text-sm text-gray-700">ایجاد گزارش</span>
              </label>
              <label class="flex items-center">
                <input
                  v-model="permissions.reports.export"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <span class="mr-3 text-sm text-gray-700">خروجی گزارش</span>
              </label>
              <label class="flex items-center">
                <input
                  v-model="permissions.reports.schedule"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <span class="mr-3 text-sm text-gray-700">زمان‌بندی گزارش</span>
              </label>
            </div>
          </div>

          <div>
            <h6 class="text-sm font-medium text-gray-900 mb-3">تنظیمات سیستم</h6>
            <div class="space-y-2">
              <label class="flex items-center">
                <input
                  v-model="permissions.settings.view"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <span class="mr-3 text-sm text-gray-700">مشاهده تنظیمات</span>
              </label>
              <label class="flex items-center">
                <input
                  v-model="permissions.settings.edit"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <span class="mr-3 text-sm text-gray-700">ویرایش تنظیمات</span>
              </label>
              <label class="flex items-center">
                <input
                  v-model="permissions.settings.security"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <span class="mr-3 text-sm text-gray-700">تنظیمات امنیتی</span>
              </label>
              <label class="flex items-center">
                <input
                  v-model="permissions.settings.users"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <span class="mr-3 text-sm text-gray-700">مدیریت کاربران</span>
              </label>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات مجوزها -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات مجوزها</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">سیاست مجوزها</label>
          <select
            v-model="permissionSettings.policy"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="deny">انکار پیش‌فرض</option>
            <option value="allow">اجازه پیش‌فرض</option>
            <option value="prompt">درخواست تأیید</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">مدت اعتبار جلسه</label>
          <select
            v-model="permissionSettings.sessionTimeout"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="15">۱۵ دقیقه</option>
            <option value="30">۳۰ دقیقه</option>
            <option value="60">۱ ساعت</option>
            <option value="480">۸ ساعت</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تلاش ورود</label>
          <input
            v-model="permissionSettings.maxLoginAttempts"
            type="number"
            min="3"
            max="10"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">مدت قفل حساب</label>
          <select
            v-model="permissionSettings.lockoutDuration"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="15">۱۵ دقیقه</option>
            <option value="30">۳۰ دقیقه</option>
            <option value="60">۱ ساعت</option>
            <option value="1440">۲۴ ساعت</option>
          </select>
        </div>
      </div>

      <div class="mt-4 space-y-3">
        <label class="flex items-center">
          <input
            v-model="permissionSettings.auditLog"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">ثبت فعالیت‌های کاربران</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="permissionSettings.ipRestriction"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">محدودیت IP</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="permissionSettings.mfaRequired"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">اجبار احراز هویت دو مرحله‌ای</span>
        </label>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// آمار مجوزها
const permissionStats = ref({
  totalRoles: 6,
  activeUsers: 24,
  totalPermissions: 32,
  lastChange: 3
})

// نقش‌های کاربری
const userRoles = ref([
  {
    id: 1,
    name: 'مدیر سیستم',
    description: 'دسترسی کامل به تمام بخش‌های سیستم',
    status: 'active',
    users: 2,
    permissions: 32
  },
  {
    id: 2,
    name: 'کاربر حسابداری',
    description: 'دسترسی به بخش‌های حسابداری و گزارش‌گیری',
    status: 'active',
    users: 8,
    permissions: 24
  },
  {
    id: 3,
    name: 'کاربر مشاهده‌کننده',
    description: 'فقط مشاهده گزارش‌ها و آمار',
    status: 'active',
    users: 14,
    permissions: 8
  }
])

// مجوزهای سیستم
const permissions = ref({
  connections: {
    view: true,
    create: true,
    edit: true,
    delete: false
  },
  data: {
    sync: true,
    backup: true,
    restore: false,
    cleanup: false
  },
  reports: {
    view: true,
    create: true,
    export: true,
    schedule: false
  },
  settings: {
    view: true,
    edit: false,
    security: false,
    users: false
  }
})

// تنظیمات مجوزها
const permissionSettings = ref({
  policy: 'deny',
  sessionTimeout: 60,
  maxLoginAttempts: 5,
  lockoutDuration: 30,
  auditLog: true,
  ipRestriction: false,
  mfaRequired: true
})

// رنگ‌های وضعیت
const getRoleStatusColor = (status: string) => {
  const colors = {
    active: 'bg-green-500',
    inactive: 'bg-gray-500',
    suspended: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

// عملیات
const createRole = () => {
  console.log('ایجاد نقش جدید')
}

const viewRole = (role: any) => {
  console.log('مشاهده نقش:', role)
}

const editRole = (role: any) => {
  console.log('ویرایش نقش:', role)
}

const deleteRole = (role: any) => {
  console.log('حذف نقش:', role)
}
</script>

<!--
  کامپوننت مدیریت مجوزهای دسترسی
  شامل:
  1. آمار مجوزها
  2. نقش‌های کاربری
  3. مجوزهای سیستم
  4. تنظیمات مجوزها
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
