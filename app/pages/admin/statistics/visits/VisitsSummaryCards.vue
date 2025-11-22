<template>
  <div v-if="hasAccess" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
    <!-- Total Visits -->
    <div class="bg-gradient-to-br from-blue-500 to-blue-600 text-white rounded-xl shadow-lg p-6 relative overflow-hidden">
      <div class="absolute top-0 right-0 w-24 h-24 bg-white bg-opacity-10 rounded-full -translate-y-6 translate-x-6"></div>
      <div class="relative">
        <div class="flex items-center justify-between mb-2">
          <div class="w-10 h-10 bg-white bg-opacity-20 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M3 3a1 1 0 000 2v8a2 2 0 002 2h2.586l-1.293 1.293a1 1 0 101.414 1.414L10 15.414l2.293 2.293a1 1 0 001.414-1.414L12.414 15H15a2 2 0 002-2V5a1 1 0 100-2H3zm11.707 4.707a1 1 0 00-1.414-1.414L10 9.586 8.707 8.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
            </svg>
          </div>
          <div class="text-sm opacity-90">کل بازدیدها</div>
        </div>
        <div class="text-3xl font-bold mb-1">{{ stats.total.toLocaleString('fa-IR') }}</div>
        <div class="flex items-center gap-1 text-sm opacity-90">
          <svg class="w-4 h-4 text-green-200" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M5.293 7.707a1 1 0 010-1.414l4-4a1 1 0 011.414 0l4 4a1 1 0 01-1.414 1.414L11 5.414V17a1 1 0 11-2 0V5.414L6.707 7.707a1 1 0 01-1.414 0z" clip-rule="evenodd" />
          </svg>
          {{ stats.growth }}% نسبت به ماه قبل
        </div>
      </div>
    </div>

    <!-- Today's Visits -->
    <div class="bg-gradient-to-br from-green-500 to-green-600 text-white rounded-xl shadow-lg p-6 relative overflow-hidden">
      <div class="absolute top-0 right-0 w-24 h-24 bg-white bg-opacity-10 rounded-full -translate-y-6 translate-x-6"></div>
      <div class="relative">
        <div class="flex items-center justify-between mb-2">
          <div class="w-10 h-10 bg-white bg-opacity-20 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd" />
            </svg>
          </div>
          <div class="text-sm opacity-90">بازدید امروز</div>
        </div>
        <div class="text-3xl font-bold mb-1">{{ stats.today.toLocaleString('fa-IR') }}</div>
        <div class="text-sm opacity-90">از ساعت 00:00 تاکنون</div>
      </div>
    </div>

    <!-- Unique Visitors -->
    <div class="bg-gradient-to-br from-purple-500 to-purple-600 text-white rounded-xl shadow-lg p-6 relative overflow-hidden">
      <div class="absolute top-0 right-0 w-24 h-24 bg-white bg-opacity-10 rounded-full -translate-y-6 translate-x-6"></div>
      <div class="relative">
        <div class="flex items-center justify-between mb-2">
          <div class="w-10 h-10 bg-white bg-opacity-20 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
              <path d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3z" />
            </svg>
          </div>
          <div class="text-sm opacity-90">بازدیدکنندگان یکتا</div>
        </div>
        <div class="text-3xl font-bold mb-1">{{ stats.unique.toLocaleString('fa-IR') }}</div>
        <div class="text-sm opacity-90">در این ماه</div>
      </div>
    </div>

    <!-- Bounce Rate -->
    <div class="bg-gradient-to-br from-orange-500 to-red-500 text-white rounded-xl shadow-lg p-6 relative overflow-hidden">
      <div class="absolute top-0 right-0 w-24 h-24 bg-white bg-opacity-10 rounded-full -translate-y-6 translate-x-6"></div>
      <div class="relative">
        <div class="flex items-center justify-between mb-2">
          <div class="w-10 h-10 bg-white bg-opacity-20 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z" clip-rule="evenodd" />
            </svg>
          </div>
          <div class="text-sm opacity-90">نرخ پرش</div>
        </div>
        <div class="text-3xl font-bold mb-1">{{ stats.bounce }}%</div>
        <div class="text-sm opacity-90">میانگین مدت: {{ stats.avgDuration }}</div>
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

// احراز هویت
const { user, isAuthenticated } = useAuth()

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false
  }

  const userRole = user.value?.role?.toLowerCase() || ''
  const adminRoles = ['admin', 'developer']
  return adminRoles.includes(userRole)
})

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async () => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false })
  }
}

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth()
})

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth()
  }
})

interface Stats {
  total: number
  today: number
  week: number
  month: number
  unique: number
  bounce: number
  avgDuration: string
  growth: number
}

defineProps<{
  stats: Stats
}>()
</script> 
