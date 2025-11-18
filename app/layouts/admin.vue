<template>
  <div class="flex h-screen bg-gray-100">
    <!-- Sidebar -->
    <nav class="w-64 bg-gray-900 text-white">
      <div class="p-6 border-b border-gray-800">
        <h1 class="text-2xl font-bold">Admin</h1>
      </div>

      <ul class="mt-6 space-y-1">
        <li>
          <NuxtLink
            to="/admin"
            class="block px-6 py-3 hover:bg-gray-800 transition-colors"
            active-class="bg-blue-600"
          >
            داشبورد
          </NuxtLink>
        </li>

        <li v-if="permission.hasPermission('manage_users')">
          <NuxtLink
            to="/admin/users"
            class="block px-6 py-3 hover:bg-gray-800 transition-colors"
            active-class="bg-blue-600"
          >
            مدیریت کاربران
          </NuxtLink>
        </li>

        <li v-if="permission.hasPermission('manage_roles')">
          <NuxtLink
            to="/admin/roles"
            class="block px-6 py-3 hover:bg-gray-800 transition-colors"
            active-class="bg-blue-600"
          >
            مدیریت نقش‌ها
          </NuxtLink>
        </li>

        <li v-if="permission.hasPermission('manage_permissions')">
          <NuxtLink
            to="/admin/permissions"
            class="block px-6 py-3 hover:bg-gray-800 transition-colors"
            active-class="bg-blue-600"
          >
            مدیریت دسترسی‌ها
          </NuxtLink>
        </li>

        <li v-if="permission.hasPermission('view_audit_logs')">
          <NuxtLink
            to="/admin/audit-logs"
            class="block px-6 py-3 hover:bg-gray-800 transition-colors"
            active-class="bg-blue-600"
          >
            لاگ‌های تدقیق
          </NuxtLink>
        </li>

        <li v-if="permission.hasPermission('manage_settings')">
          <NuxtLink
            to="/admin/settings"
            class="block px-6 py-3 hover:bg-gray-800 transition-colors"
            active-class="bg-blue-600"
          >
            تنظیمات سیستم
          </NuxtLink>
        </li>

        <li class="border-t border-gray-800 mt-6 pt-6">
          <NuxtLink
            to="/"
            class="block px-6 py-3 hover:bg-gray-800 transition-colors text-yellow-400"
          >
            ← بازگشت به سایت
          </NuxtLink>
        </li>
      </ul>
    </nav>

    <!-- Main Content -->
    <div class="flex-1 flex flex-col overflow-hidden">
      <!-- Top Bar -->
      <header class="bg-white shadow">
        <div class="px-6 py-4 flex justify-between items-center">
          <h2 class="text-xl font-semibold text-gray-900">{{ pageTitle }}</h2>
          <div class="flex items-center space-x-4">
            <span class="text-sm text-gray-600">{{ authStore.user?.mobile }}</span>
            <button
              @click="handleLogout"
              class="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700 transition-colors text-sm"
            >
              خروج
            </button>
          </div>
        </div>
      </header>

      <!-- Page Content -->
      <main class="flex-1 overflow-auto">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '@/composables/useAuth'
import { usePermission } from '@/composables/usePermission'
import { useAuthStore } from '@/stores/auth'
import { ref } from 'vue'

const authStore = useAuthStore()
const permission = usePermission()
const auth = useAuth()

const pageTitle = ref('صفحه ادمین')

const handleLogout = async () => {
  await auth.logout()
}
</script>
