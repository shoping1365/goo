<template>
  <div class="p-8">
    <div class="bg-white rounded-lg shadow p-6">
      <h2 class="text-2xl font-bold text-gray-900 mb-6">مدیریت دسترسی‌ها</h2>

      <!-- Permissions Matrix -->
      <div v-if="permService.loading.value || roleService.loading.value" class="text-center py-8">
        <p class="text-gray-600">درحال بارگذاری...</p>
      </div>

      <div v-else class="overflow-x-auto">
        <table class="w-full border-collapse">
          <thead class="bg-gray-50">
            <tr>
              <th class="border border-gray-200 px-4 py-2 text-right text-sm font-semibold">دسترسی</th>
              <th class="border border-gray-200 px-4 py-2 text-center text-sm font-semibold">توضیح</th>
              <th class="border border-gray-200 px-4 py-2 text-center text-sm font-semibold">
                <span
                  v-for="role in roleService.roles.value"
                  :key="role.id"
                  class="block text-xs"
                >
                  {{ role.name }}
                </span>
              </th>
            </tr>
          </thead>
          <tbody class="divide-y">
            <tr v-for="perm in permService.permissions.value" :key="perm.id" class="hover:bg-gray-50">
              <td class="border border-gray-200 px-4 py-3 text-sm font-medium text-gray-900">
                {{ perm.name }}
              </td>
              <td class="border border-gray-200 px-4 py-3 text-sm text-gray-600">
                {{ perm.description }}
              </td>
              <td class="border border-gray-200 px-4 py-3 text-center">
                <div class="space-y-2">
                  <div v-for="role in roleService.roles.value" :key="role.id" class="flex justify-center">
                    <input
                      type="checkbox"
                      :checked="roleHasPermission(role.id, perm.id)"
                      class="w-4 h-4 rounded border-gray-300"
                      @change="togglePermission(role.id, perm.id, $event)"
                    />
                  </div>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Action Buttons -->
      <div class="mt-6 flex gap-3">
        <button
          class="px-6 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors"
          @click="savePermissions"
        >
          💾 ذخیره
        </button>
        <button
          class="px-6 py-2 bg-gray-400 text-white rounded-lg hover:bg-gray-500 transition-colors"
          @click="resetPermissions"
        >
          ↺ بازنشانی
        </button>
      </div>
    </div>

    <!-- Permission Categories -->
    <div class="mt-8 grid grid-cols-1 md:grid-cols-2 gap-6">
      <div class="bg-blue-50 rounded-lg p-6 border border-blue-200">
        <h3 class="text-lg font-bold text-blue-900 mb-4">🔐 دسترسی‌های امنیتی</h3>
        <ul class="space-y-2 text-sm text-blue-800">
          <li>✓ view_audit_logs - مشاهده لاگ‌های تدقیق</li>
          <li>✓ manage_permissions - مدیریت دسترسی‌ها</li>
          <li>✓ manage_security - مدیریت تنظیمات امنیتی</li>
        </ul>
      </div>

      <div class="bg-green-50 rounded-lg p-6 border border-green-200">
        <h3 class="text-lg font-bold text-green-900 mb-4">👥 دسترسی‌های کاربر</h3>
        <ul class="space-y-2 text-sm text-green-800">
          <li>✓ manage_users - مدیریت کاربران</li>
          <li>✓ view_profile - مشاهده پروفایل</li>
          <li>✓ edit_profile - ویرایش پروفایل</li>
        </ul>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6 border border-yellow-200">
        <h3 class="text-lg font-bold text-yellow-900 mb-4">🎭 دسترسی‌های نقش</h3>
        <ul class="space-y-2 text-sm text-yellow-800">
          <li>✓ manage_roles - مدیریت نقش‌ها</li>
          <li>✓ assign_roles - انتساب نقش‌ها</li>
        </ul>
      </div>

      <div class="bg-purple-50 rounded-lg p-6 border border-purple-200">
        <h3 class="text-lg font-bold text-purple-900 mb-4">⚙️ دسترسی‌های سیستم</h3>
        <ul class="space-y-2 text-sm text-purple-800">
          <li>✓ manage_settings - مدیریت تنظیمات</li>
          <li>✓ view_logs - مشاهده لاگ‌ها</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const navigateTo: (to: string) => Promise<void>
</script>

<script setup lang="ts">
import { usePermission } from '@/composables/usePermission'
import { usePermissionsAPI } from '@/composables/usePermissionsAPI'
import { useRoles } from '@/composables/useRoles'
import { onMounted, ref } from 'vue'

definePageMeta({
  middleware: 'admin',
  layout: 'admin'
})

const permission = usePermission()
const permService = usePermissionsAPI()
const roleService = useRoles()

// بررسی دسترسی
if (!permission.hasPermission('manage_permissions')) {
  navigateTo('/403')
}

const rolePermissions = ref(new Map())

// بارگذاری داده‌ها هنگام mount
onMounted(async () => {
  await permService.fetchPermissions()
  await roleService.fetchRoles()
  
  // مقداردهی rolePermissions
  for (const role of roleService.roles.value) {
    const perms = await roleService.getRolePermissions(role.id)
    rolePermissions.value.set(role.id, new Set(perms))
  }
})

const roleHasPermission = (roleId: string, permId: string): boolean => {
  return rolePermissions.value.get(roleId)?.has(permId) ?? false
}

const togglePermission = (roleId: string, permId: string, event: Event) => {
  const perms = rolePermissions.value.get(roleId)
  if (perms) {
    const target = event.target as HTMLInputElement
    if (target.checked) {
      perms.add(permId)
    } else {
      perms.delete(permId)
    }
  }
}

const savePermissions = async () => {
  try {
    for (const role of roleService.roles.value) {
      const perms = Array.from(rolePermissions.value.get(role.id) || new Set()) as string[]
      await roleService.assignPermissionsToRole(role.id, perms)
    }
    alert('تنظیمات دسترسی‌ها ذخیره شد')
  } catch (err) {
    alert('خطا در ذخیره تنظیمات')
  }
}

const resetPermissions = async () => {
  await permService.fetchPermissions()
  await roleService.fetchRoles()
}
</script>
