<template>
  <div v-if="hasAccess" class="min-h-screen">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-4">
          <div>
            <h1 class="text-2xl font-bold text-gray-900">مدیریت منوها</h1>
            <p class="text-sm text-gray-600">ایجاد و مدیریت منوهای سایت</p>
          </div>
          <NuxtLink
            to="/admin/menu-management"
            class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg text-sm font-medium"
          >
            ایجاد منوی جدید
          </NuxtLink>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
      <!-- Loading State -->
      <div v-if="isLoading" class="text-center py-12">
        <div class="inline-flex items-center px-4 py-2 font-semibold leading-6 text-sm shadow rounded-md text-white bg-blue-500 hover:bg-blue-400 transition ease-in-out duration-150 cursor-not-allowed">
          <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          در حال بارگذاری...
        </div>
      </div>

      <!-- Empty State -->
      <div v-else-if="menus.length === 0" class="text-center py-12">
        <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"/>
        </svg>
        <h3 class="mt-2 text-sm font-medium text-gray-900">هیچ منویی وجود ندارد</h3>
        <p class="mt-1 text-sm text-gray-500">برای شروع، منوی جدیدی ایجاد کنید.</p>
        <div class="mt-6">
          <NuxtLink
            to="/admin/menu-management"
            class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
          >
            ایجاد منوی جدید
          </NuxtLink>
        </div>
      </div>

      <!-- Menus List -->
      <div v-else class="bg-white shadow overflow-hidden sm:rounded-md">
        <ul class="divide-y divide-gray-200">
          <li v-for="menu in menus" :key="menu.id" class="px-6 py-4">
            <div class="flex items-center justify-between">
              <div class="flex items-center">
                <div class="flex-shrink-0">
                  <div class="h-10 w-10 rounded-full bg-blue-100 flex items-center justify-center">
                    <svg class="h-6 w-6 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"/>
                    </svg>
                  </div>
                </div>
                <div class="mr-4">
                  <div class="text-sm font-medium text-gray-900">{{ menu.name }}</div>
                  <div class="text-sm text-gray-500">
                    {{ menu.items?.length || 0 }} آیتم
                    <span v-if="menu.slug" class="mr-2">• شناسه: {{ menu.slug }}</span>
                  </div>
                </div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <span
                  :class="[
                    'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                    menu.enabled
                      ? 'bg-green-100 text-green-800'
                      : 'bg-red-100 text-red-800'
                  ]"
                >
                  {{ menu.enabled ? 'فعال' : 'غیرفعال' }}
                </span>
                <div class="flex space-x-2 space-x-reverse">
                  <NuxtLink
                    :to="`/admin/menu-management?id=${menu.id}`"
                    class="text-blue-600 hover:text-blue-900 text-sm font-medium"
                  >
                    ویرایش
                  </NuxtLink>
                  <button
                    class="text-red-600 hover:text-red-900 text-sm font-medium"
                    @click="deleteMenu(menu.id)"
                  >
                    حذف
                  </button>
                </div>
              </div>
            </div>
            
            <!-- Menu Items Preview -->
            <div v-if="menu.items && menu.items.length > 0" class="mt-3">
              <div class="flex flex-wrap gap-2">
                <span
                  v-for="item in menu.items.slice(0, 5)"
                  :key="item.id"
                  class="inline-flex items-center px-2 py-1 rounded-full text-xs bg-gray-100 text-gray-800"
                >
                  {{ item.title }}
                </span>
                <span
                  v-if="menu.items.length > 5"
                  class="inline-flex items-center px-2 py-1 rounded-full text-xs bg-gray-100 text-gray-600"
                >
                  +{{ menu.items.length - 5 }} آیتم دیگر
                </span>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';
import { useMenuManagement } from '~/composables/useMenuManagement'

// احراز هویت
const { user, isAuthenticated } = useAuth();

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false;
  }

  const userRole = user.value?.role?.toLowerCase() || '';
  const adminRoles = ['admin', 'developer'];
  return adminRoles.includes(userRole);
});

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async (): Promise<void> => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false });
  }
};

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth();
  fetchMenus()
});

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth();
  }
});

// Use the composable
const {
  menus,
  isLoading,
  fetchMenus,
  deleteMenu: deleteMenuFromAPI
} = useMenuManagement()

// Methods
const deleteMenu = async (id: string | number): Promise<void> => {
  try {
    await deleteMenuFromAPI(Number(id))
    // TODO: نمایش پیام موفقیت با toast

  } catch (error) {
    // TODO: نمایش پیام خطا با toast
    console.error('خطا در حذف منو')
  }
}

</script>
