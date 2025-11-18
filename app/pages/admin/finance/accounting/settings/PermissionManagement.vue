<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">مدیریت مجوزهای دسترسی</h4>
        <p class="text-sm text-gray-600 mt-1">مدیریت مجوزهای دسترسی کاربران به اتصال حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          @click="showCreateRoleModal = true"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
          </svg>
          ایجاد نقش جدید
        </button>
        <button
          @click="savePermissions"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          </svg>
          ذخیره مجوزها
        </button>
      </div>
    </div>

    <!-- آمار مجوزها -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ permissionStats.totalUsers }}</div>
            <div class="text-sm text-blue-700">کل کاربران</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ permissionStats.activeRoles }}</div>
            <div class="text-sm text-green-700">نقش‌های فعال</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-yellow-600">{{ permissionStats.pendingApprovals }}</div>
            <div class="text-sm text-yellow-700">در انتظار تأیید</div>
          </div>
          <div class="w-10 h-10 bg-yellow-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-purple-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-purple-600">{{ permissionStats.permissionGroups }}</div>
            <div class="text-sm text-purple-700">گروه‌های مجوز</div>
          </div>
          <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- تب‌های مدیریت -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="border-b border-gray-200">
        <nav class="-mb-px flex space-x-8 space-x-reverse px-6">
          <button
            v-for="tab in permissionTabs"
            :key="tab.id"
            @click="activePermissionTab = tab.id"
            :class="[
              activePermissionTab === tab.id
                ? 'border-blue-500 text-blue-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
              'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm'
            ]"
          >
            {{ tab.name }}
          </button>
        </nav>
      </div>

      <div class="p-6">
        <!-- تب نقش‌ها -->
        <div v-if="activePermissionTab === 'roles'" class="space-y-6">
          <div class="flex items-center justify-between">
            <h5 class="text-md font-medium text-gray-900">نقش‌های دسترسی</h5>
            <button
              @click="showCreateRoleModal = true"
              class="inline-flex items-center px-3 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
            >
              <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
              </svg>
              ایجاد نقش
            </button>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div
              v-for="role in roles"
              :key="role.id"
              class="bg-gray-50 rounded-lg p-6 border border-gray-200"
            >
              <div class="flex items-center justify-between mb-4">
                <div>
                  <h6 class="font-medium text-gray-900">{{ role.name }}</h6>
                  <p class="text-sm text-gray-600">{{ role.description }}</p>
                </div>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button
                    @click="editRole(role)"
                    class="p-1 text-blue-600 hover:text-blue-800"
                    title="ویرایش"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                    </svg>
                  </button>
                  <button
                    @click="deleteRole(role)"
                    class="p-1 text-red-600 hover:text-red-800"
                    title="حذف"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                    </svg>
                  </button>
                </div>
              </div>

              <div class="space-y-2">
                <div class="flex items-center justify-between text-sm">
                  <span class="text-gray-600">تعداد کاربران:</span>
                  <span class="font-medium">{{ role.userCount }}</span>
                </div>
                <div class="flex items-center justify-between text-sm">
                  <span class="text-gray-600">تعداد مجوزها:</span>
                  <span class="font-medium">{{ role.permissionCount }}</span>
                </div>
                <div class="flex items-center justify-between text-sm">
                  <span class="text-gray-600">وضعیت:</span>
                  <span
                    class="px-2 py-1 rounded-full text-xs font-medium"
                    :class="getRoleStatusClass(role.status)"
                  >
                    {{ getRoleStatusLabel(role.status) }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- تب کاربران -->
        <div v-if="activePermissionTab === 'users'" class="space-y-6">
          <h5 class="text-md font-medium text-gray-900">مدیریت کاربران</h5>
          
          <div class="bg-white rounded-lg border border-gray-200">
            <div class="overflow-x-auto">
              <table class="w-full text-sm">
                <thead>
                  <tr class="border-b border-gray-200 bg-gray-50">
                    <th class="text-right py-3 px-4 font-medium text-gray-600">کاربر</th>
                    <th class="text-right py-3 px-4 font-medium text-gray-600">نقش</th>
                    <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
                    <th class="text-right py-3 px-4 font-medium text-gray-600">آخرین ورود</th>
                    <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="user in users" :key="user.id" class="border-b border-gray-100 hover:bg-gray-50">
                    <td class="py-3 px-4">
                      <div class="flex items-center space-x-3 space-x-reverse">
                        <div class="w-8 h-8 bg-gray-300 rounded-full flex items-center justify-center">
                          <span class="text-sm font-medium text-gray-700">{{ user.initials }}</span>
                        </div>
                        <div>
                          <div class="font-medium text-gray-900">{{ user.name }}</div>
                          <div class="text-xs text-gray-500">{{ user.email }}</div>
                        </div>
                      </div>
                    </td>
                    <td class="py-3 px-4">
                      <span class="px-2 py-1 rounded-full text-xs font-medium bg-blue-100 text-blue-700">
                        {{ user.role }}
                      </span>
                    </td>
                    <td class="py-3 px-4">
                      <span
                        class="px-2 py-1 rounded-full text-xs font-medium"
                        :class="getUserStatusClass(user.status)"
                      >
                        {{ getUserStatusLabel(user.status) }}
                      </span>
                    </td>
                    <td class="py-3 px-4">
                      <div class="text-sm text-gray-900">{{ user.lastLogin }}</div>
                      <div class="text-xs text-gray-500">{{ user.lastLoginTime }}</div>
                    </td>
                    <td class="py-3 px-4">
                      <div class="flex items-center space-x-2 space-x-reverse">
                        <button
                          @click="editUserPermissions(user)"
                          class="p-1 text-blue-600 hover:text-blue-800"
                          title="ویرایش مجوزها"
                        >
                          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                          </svg>
                        </button>
                        <button
                          @click="revokeUserAccess(user)"
                          class="p-1 text-red-600 hover:text-red-800"
                          title="لغو دسترسی"
                        >
                          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728" />
                          </svg>
                        </button>
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>

        <!-- تب مجوزها -->
        <div v-if="activePermissionTab === 'permissions'" class="space-y-6">
          <h5 class="text-md font-medium text-gray-900">تنظیمات مجوزها</h5>
          
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div class="bg-gray-50 rounded-lg p-6">
              <h6 class="font-medium text-gray-900 mb-4">مجوزهای اتصال</h6>
              <div class="space-y-3">
                <label class="flex items-center">
                  <input
                    v-model="permissions.connectionView"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">مشاهده اتصالات</span>
                </label>
                <label class="flex items-center">
                  <input
                    v-model="permissions.connectionCreate"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">ایجاد اتصال</span>
                </label>
                <label class="flex items-center">
                  <input
                    v-model="permissions.connectionEdit"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">ویرایش اتصال</span>
                </label>
                <label class="flex items-center">
                  <input
                    v-model="permissions.connectionDelete"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">حذف اتصال</span>
                </label>
              </div>
            </div>

            <div class="bg-gray-50 rounded-lg p-6">
              <h6 class="font-medium text-gray-900 mb-4">مجوزهای همگام‌سازی</h6>
              <div class="space-y-3">
                <label class="flex items-center">
                  <input
                    v-model="permissions.syncView"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">مشاهده همگام‌سازی</span>
                </label>
                <label class="flex items-center">
                  <input
                    v-model="permissions.syncStart"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">شروع همگام‌سازی</span>
                </label>
                <label class="flex items-center">
                  <input
                    v-model="permissions.syncStop"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">توقف همگام‌سازی</span>
                </label>
                <label class="flex items-center">
                  <input
                    v-model="permissions.syncSettings"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">تنظیمات همگام‌سازی</span>
                </label>
              </div>
            </div>

            <div class="bg-gray-50 rounded-lg p-6">
              <h6 class="font-medium text-gray-900 mb-4">مجوزهای گزارش‌گیری</h6>
              <div class="space-y-3">
                <label class="flex items-center">
                  <input
                    v-model="permissions.reportView"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">مشاهده گزارش‌ها</span>
                </label>
                <label class="flex items-center">
                  <input
                    v-model="permissions.reportExport"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">صادر کردن گزارش</span>
                </label>
                <label class="flex items-center">
                  <input
                    v-model="permissions.reportCreate"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">ایجاد گزارش سفارشی</span>
                </label>
              </div>
            </div>

            <div class="bg-gray-50 rounded-lg p-6">
              <h6 class="font-medium text-gray-900 mb-4">مجوزهای مدیریتی</h6>
              <div class="space-y-3">
                <label class="flex items-center">
                  <input
                    v-model="permissions.userManagement"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">مدیریت کاربران</span>
                </label>
                <label class="flex items-center">
                  <input
                    v-model="permissions.roleManagement"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">مدیریت نقش‌ها</span>
                </label>
                <label class="flex items-center">
                  <input
                    v-model="permissions.systemSettings"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">تنظیمات سیستم</span>
                </label>
                <label class="flex items-center">
                  <input
                    v-model="permissions.securitySettings"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">تنظیمات امنیتی</span>
                </label>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال ایجاد نقش -->
    <div v-if="showCreateRoleModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">ایجاد نقش جدید</h3>
          <form @submit.prevent="createRole" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نام نقش</label>
              <input
                v-model="newRole.name"
                type="text"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="نام نقش"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
              <textarea
                v-model="newRole.description"
                rows="3"
                class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="توضیحات نقش"
              ></textarea>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">مجوزها</label>
              <div class="space-y-2 max-h-40 overflow-y-auto">
                <label v-for="permission in availablePermissions" :key="permission.id" class="flex items-center">
                  <input
                    v-model="newRole.permissions"
                    :value="permission.id"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">{{ permission.name }}</span>
                </label>
              </div>
            </div>

            <div class="flex items-center justify-end space-x-3 space-x-reverse">
              <button
                type="button"
                @click="showCreateRoleModal = false"
                class="px-4 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50"
              >
                انصراف
              </button>
              <button
                type="submit"
                class="px-4 py-2 bg-blue-600 text-white rounded-md text-sm font-medium hover:bg-blue-700"
              >
                ایجاد نقش
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// وضعیت مودال
const showCreateRoleModal = ref(false)

// تب فعال
const activePermissionTab = ref('roles')

// تب‌های مدیریت
const permissionTabs = ref([
  { id: 'roles', name: 'نقش‌ها' },
  { id: 'users', name: 'کاربران' },
  { id: 'permissions', name: 'مجوزها' }
])

// آمار مجوزها
const permissionStats = ref({
  totalUsers: 25,
  activeRoles: 8,
  pendingApprovals: 3,
  permissionGroups: 5
})

// نقش‌ها
const roles = ref([
  {
    id: 1,
    name: 'مدیر سیستم',
    description: 'دسترسی کامل به تمام بخش‌ها',
    userCount: 2,
    permissionCount: 15,
    status: 'active'
  },
  {
    id: 2,
    name: 'مدیر حسابداری',
    description: 'مدیریت اتصالات و همگام‌سازی',
    userCount: 5,
    permissionCount: 12,
    status: 'active'
  },
  {
    id: 3,
    name: 'کاربر عادی',
    description: 'مشاهده و گزارش‌گیری',
    userCount: 18,
    permissionCount: 6,
    status: 'active'
  }
])

// کاربران
const users = ref([
  {
    id: 1,
    name: 'علی احمدی',
    email: 'ali@example.com',
    initials: 'عا',
    role: 'مدیر سیستم',
    status: 'active',
    lastLogin: 'امروز',
    lastLoginTime: '۱۴:۳۰'
  },
  {
    id: 2,
    name: 'فاطمه محمدی',
    email: 'fateme@example.com',
    initials: 'فم',
    role: 'مدیر حسابداری',
    status: 'active',
    lastLogin: 'دیروز',
    lastLoginTime: '۱۸:۴۵'
  },
  {
    id: 3,
    name: 'محمد رضایی',
    email: 'mohammad@example.com',
    initials: 'مر',
    role: 'کاربر عادی',
    status: 'inactive',
    lastLogin: 'هفته گذشته',
    lastLoginTime: '۱۰:۱۵'
  }
])

// مجوزها
const permissions = ref({
  connectionView: true,
  connectionCreate: false,
  connectionEdit: false,
  connectionDelete: false,
  syncView: true,
  syncStart: false,
  syncStop: false,
  syncSettings: false,
  reportView: true,
  reportExport: false,
  reportCreate: false,
  userManagement: false,
  roleManagement: false,
  systemSettings: false,
  securitySettings: false
})

// نقش جدید
const newRole = ref({
  name: '',
  description: '',
  permissions: []
})

// مجوزهای موجود
const availablePermissions = ref([
  { id: 'connectionView', name: 'مشاهده اتصالات' },
  { id: 'connectionCreate', name: 'ایجاد اتصال' },
  { id: 'connectionEdit', name: 'ویرایش اتصال' },
  { id: 'connectionDelete', name: 'حذف اتصال' },
  { id: 'syncView', name: 'مشاهده همگام‌سازی' },
  { id: 'syncStart', name: 'شروع همگام‌سازی' },
  { id: 'syncStop', name: 'توقف همگام‌سازی' },
  { id: 'syncSettings', name: 'تنظیمات همگام‌سازی' },
  { id: 'reportView', name: 'مشاهده گزارش‌ها' },
  { id: 'reportExport', name: 'صادر کردن گزارش' },
  { id: 'reportCreate', name: 'ایجاد گزارش سفارشی' },
  { id: 'userManagement', name: 'مدیریت کاربران' },
  { id: 'roleManagement', name: 'مدیریت نقش‌ها' },
  { id: 'systemSettings', name: 'تنظیمات سیستم' },
  { id: 'securitySettings', name: 'تنظیمات امنیتی' }
])

// کلاس وضعیت نقش
const getRoleStatusClass = (status: string) => {
  const classes = {
    active: 'bg-green-100 text-green-700',
    inactive: 'bg-gray-100 text-gray-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// برچسب وضعیت نقش
const getRoleStatusLabel = (status: string) => {
  const labels = {
    active: 'فعال',
    inactive: 'غیرفعال'
  }
  return labels[status] || status
}

// کلاس وضعیت کاربر
const getUserStatusClass = (status: string) => {
  const classes = {
    active: 'bg-green-100 text-green-700',
    inactive: 'bg-red-100 text-red-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// برچسب وضعیت کاربر
const getUserStatusLabel = (status: string) => {
  const labels = {
    active: 'فعال',
    inactive: 'غیرفعال'
  }
  return labels[status] || status
}

// ذخیره مجوزها
const savePermissions = () => {
  // TODO: ذخیره مجوزها
  console.log('ذخیره مجوزها:', permissions.value)
}

// ایجاد نقش جدید
const createRole = () => {
  // TODO: ایجاد نقش جدید
  console.log('ایجاد نقش جدید:', newRole.value)
  showCreateRoleModal.value = false
  
  // پاک کردن فرم
  newRole.value = {
    name: '',
    description: '',
    permissions: []
  }
}

// ویرایش نقش
const editRole = (role: any) => {
  // TODO: ویرایش نقش
  console.log('ویرایش نقش:', role)
}

// حذف نقش
const deleteRole = (role: any) => {
  // TODO: حذف نقش
  console.log('حذف نقش:', role)
}

// ویرایش مجوزهای کاربر
const editUserPermissions = (user: any) => {
  // TODO: ویرایش مجوزهای کاربر
  console.log('ویرایش مجوزهای کاربر:', user)
}

// لغو دسترسی کاربر
const revokeUserAccess = (user: any) => {
  // TODO: لغو دسترسی کاربر
  console.log('لغو دسترسی کاربر:', user)
}
</script>

<!--
  کامپوننت مدیریت مجوزهای دسترسی
  شامل:
  1. آمار مجوزها
  2. مدیریت نقش‌ها
  3. مدیریت کاربران
  4. تنظیمات مجوزها
  5. مودال ایجاد نقش
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
