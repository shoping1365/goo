<template>
  <div class="min-h-screen">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 sticky top-0 z-30">
      <div class="px-6 py-4">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-2xl font-bold text-gray-900">مجوزها</h1>
            <p class="text-sm text-gray-500 mt-1">مدیریت مجوزهای مختلف در سیستم</p>
          </div>
          
          <div class="flex items-center gap-3">
            
            <NuxtLink 
              to="/admin/access-management"
              class="inline-flex items-center px-4 py-2 border border-gray-200 rounded-lg bg-white hover:bg-gray-50 transition-all shadow-md"
            >
              <svg class="w-5 h-5 ml-2 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
              </svg>
              <span class="text-gray-700">بازگشت</span>
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="p-6">
      <div class="max-w-7xl mx-auto">
        <!-- لیست نقش‌ها و مجوزهایشان -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200">
          <div class="px-6 py-4 border-b border-gray-200">
            <h2 class="text-lg font-semibold text-gray-900">نقش‌ها و مجوزهایشان</h2>
          </div>
          <div class="p-6">
            <div v-if="roles.length === 0" class="text-center py-8">
              <svg class="w-12 h-12 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
              </svg>
              <h3 class="text-lg font-medium text-gray-900 mb-2">هیچ نقشی وجود ندارد</h3>
              <p class="text-gray-500 mb-4">ابتدا در صفحه نقش‌های کاربری، نقش جدید ایجاد کنید</p>
              <NuxtLink 
                to="/admin/access-management"
                class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-blue-600 hover:bg-blue-700"
              >
                ایجاد نقش جدید
              </NuxtLink>
            </div>
            
            <div v-else class="space-y-6">
              <div 
                v-for="role in roles" 
                :key="role.id"
                class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
              >
                <div class="flex items-center justify-between mb-4">
                  <div class="flex items-center space-x-3">
                    <div class="w-12 h-12 bg-gradient-to-r from-blue-400 to-blue-600 rounded-lg flex items-center justify-center">
                      <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                      </svg>
                    </div>
                    <div>
                      <h3 class="text-xl font-semibold text-gray-900">{{ role.display_name }}</h3>
                      <p class="text-sm text-gray-500">{{ role.description }}</p>
                      <div class="flex items-center mt-1 space-x-4 space-x-reverse">
                        <span class="text-xs text-gray-400">{{ role.users_count }} کاربر</span>
                        <span class="text-xs text-gray-400">{{ role.permissions_count }} مجوز</span>
                      </div>
                    </div>
                  </div>
                  
                  <NuxtLink 
                    :to="`/admin/access-management/role-permissions/edit/${role.id}`"
                    class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-blue-600 hover:bg-blue-700 transition-colors"
                  >
                    <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                    </svg>
                    تنظیم مجوزها
                  </NuxtLink>
                </div>
                
                <!-- نمایش مجوزهای فعلی نقش -->
                <!-- این بخش حذف شد: دیگر مجوزها نمایش داده نمی‌شوند -->
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>


  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';

// تعریف definePageMeta و useHead برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void

// تعریف interface برای نقش

interface Role {
  id: number
  name: string
  display_name: string
  description: string
  users_count: number
  permissions_count: number
  permissions?: string[]
  created_at: string
  updated_at: string
}

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

useHead({
  title: 'مجوزها - پنل مدیریت'
})

// Reactive data
const roles = ref<Role[]>([])

// Methods
const fetchRoles = async () => {
  try {
    const response = await $fetch('/api/roles') as { data: Role[] }
    // فیلتر کردن نقش customer از لیست نمایشی
    roles.value = response.data.filter(role => role.name !== 'customer')
  } catch (error) {
    console.error('خطا در دریافت نقش‌ها:', error)
  }
}



// Lifecycle
onMounted(async () => {
  await fetchRoles()
})
</script> 

