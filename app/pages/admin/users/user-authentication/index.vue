<template>
  <div class="px-4 py-4">
    <!-- Header -->
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900 mb-2">احراز هویت کاربران</h1>
      <p class="text-gray-600">مدیریت تنظیمات احراز هویت، کاربران و تاریخچه ورود</p>
    </div>

    <!-- Navigation Tabs -->
    <div class="mb-6">
      <nav class="flex space-x-8 border-b border-gray-200">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          @click="activeTab = tab.id"
          :class="[
            'py-2 px-1 border-b-2 font-medium text-sm',
            activeTab === tab.id
              ? 'border-blue-500 text-blue-600'
              : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
          ]"
        >
          {{ tab.name }}
        </button>
      </nav>
    </div>

    <!-- Tab Content -->
    <div class="bg-white rounded-lg shadow">
      <Suspense>
        <template #default>
          <div class="px-4 py-4">
            <AuthSettings v-if="activeTab === 'settings'" />
            <UserManagement v-else-if="activeTab === 'users'" />
            <LoginHistory v-else />
          </div>
        </template>

        <template #fallback>
          <div class="px-4 py-4 flex justify-center items-center">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500"></div>
            <span class="mr-3 text-gray-600">در حال بارگذاری...</span>
          </div>
        </template>
      </Suspense>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

import AuthSettings from './components/AuthSettings.vue'
import UserManagement from './components/UserManagement.vue'
import LoginHistory from './components/LoginHistory.vue'

// Tab management
const activeTab = ref('settings')

const tabs = [
  { id: 'settings', name: 'تنظیمات احراز هویت' },
  { id: 'users', name: 'مدیریت کاربران' },
  { id: 'history', name: 'تاریخچه ورود' }
]

// Page meta
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})
</script> 
