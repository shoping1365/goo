<template>
  <div class="p-8">
    <div class="bg-white rounded-lg shadow p-6">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-2xl font-bold text-gray-900">مدیریت نقش‌ها</h2>
        <button
          class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors"
          @click="showNewRoleModal = true"
        >
          + نقش جدید
        </button>
      </div>

      <!-- Roles List -->
      <div v-if="roleService.loading.value" class="text-center py-8">
        <p class="text-gray-600">درحال بارگذاری...</p>
      </div>

      <div v-else-if="roleService.error.value" class="p-4 bg-red-100 text-red-700 rounded mb-4">
        {{ roleService.error.value }}
      </div>

      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="role in roleService.roles.value" :key="role.id" class="border border-gray-200 rounded-lg p-4">
          <div class="flex justify-between items-start mb-4">
            <h3 class="text-lg font-semibold text-gray-900">{{ role.name }}</h3>
            <span class="px-2 py-1 bg-blue-100 text-blue-800 text-xs rounded">
              {{ role.permissions.length }} دسترسی
            </span>
          </div>

          <p class="text-sm text-gray-600 mb-4">{{ role.description }}</p>

          <div class="mb-4">
            <p class="text-xs font-medium text-gray-700 mb-2">دسترسی‌ها:</p>
            <div class="flex flex-wrap gap-2">
              <span
                v-for="perm in role.permissions.slice(0, 3)"
                :key="perm"
                class="text-xs bg-gray-100 text-gray-700 px-2 py-1 rounded"
              >
                {{ perm }}
              </span>
              <span v-if="role.permissions.length > 3" class="text-xs text-gray-600">
                +{{ role.permissions.length - 3 }}
              </span>
            </div>
          </div>

          <div class="flex gap-2">
            <button
              class="flex-1 px-3 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 transition-colors text-sm"
              @click="editRole(role)"
            >
              ویرایش
            </button>
            <button
              class="flex-1 px-3 py-2 bg-red-600 text-white rounded hover:bg-red-700 transition-colors text-sm"
              @click="deleteRole(role.id)"
            >
              حذف
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- New Role Modal -->
    <div
      v-if="showNewRoleModal"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
    >
      <div class="bg-white rounded-lg p-6 max-w-md w-full mx-4">
        <h3 class="text-lg font-bold mb-4">نقش جدید</h3>
        <form class="space-y-4" @submit.prevent="createRole">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">نام نقش</label>
            <input
              v-model="newRole.name"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
              required
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">توضیح</label>
            <textarea
              v-model="newRole.description"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
              rows="3"
            />
          </div>

          <div class="flex gap-3">
            <button
              type="submit"
              class="flex-1 px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors"
            >
              ایجاد
            </button>
            <button
              type="button"
              class="flex-1 px-4 py-2 bg-gray-300 text-gray-700 rounded-lg hover:bg-gray-400 transition-colors"
              @click="showNewRoleModal = false"
            >
              لغو
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const navigateTo: (to: string) => Promise<void>
</script>

<script setup lang="ts">
import { usePermission } from '@/composables/usePermission';
import { useRoles } from '@/composables/useRoles';
import { onMounted, ref } from 'vue';

definePageMeta({
  layout: 'admin',
  middleware: 'auth',
})

const permission = usePermission()
const roleService = useRoles()

// بررسی دسترسی
if (!permission.hasPermission('manage_roles')) {
  navigateTo('/403')
}

const showNewRoleModal = ref(false)
const newRole = ref({
  name: '',
  description: '',
})

// بارگذاری نقش‌ها هنگام mount
onMounted(async () => {
  await roleService.fetchRoles()
})

const createRole = async () => {
  if (!newRole.value.name) return

  const created = await roleService.createRole({
    name: newRole.value.name,
    description: newRole.value.description,
    permissions: [],
  })

  if (created) {
    newRole.value = { name: '', description: '' }
    showNewRoleModal.value = false
  }
}

interface Role {
  id?: number | string
  name?: string
  [key: string]: unknown
}
const editRole = (role: Role) => {
  // TODO: Open edit modal
}

const deleteRole = async (roleId: string) => {
  if (roleId === '1') {
    alert('نمی‌توان نقش admin را حذف کرد')
    return
  }

  if (confirm('آیا مطمئن هستید؟')) {
    const success = await roleService.deleteRole(roleId)
    if (success) {
      // Role deleted
    }
  }
}
</script>
