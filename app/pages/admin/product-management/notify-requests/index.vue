<template>
  <ClientOnly>
    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
          <div>
            <h1 class="text-lg font-bold text-gray-900">مدیریت اطلاع‌رسانی‌ها</h1>
            <p class="text-xs text-gray-600 mt-1">داشبورد جامع درخواست‌های خبرم کن</p>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <TemplateButton
              bg-gradient="bg-gradient-to-r from-green-400 to-green-600"
              text-color="text-white"
              hover-class="hover:from-green-500 hover:to-green-700 hover:shadow-lg hover:scale-105"
              focus-class="focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
              size="medium"
              @click="exportToExcel"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              خروجی اکسل
            </TemplateButton>
            <TemplateButton
              bg-gradient="bg-white"
              text-color="text-gray-500"
              border-color="border border-gray-200"
              hover-class="hover:bg-gray-50"
              size="medium"
              @click="refreshData"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
              </svg>
            </TemplateButton>
          </div>
        </div>
      </div>
      
      <!-- Navigation Menu -->
      <div class="border-t border-gray-200 bg-gray-50 px-4 py-3">
        <nav class="flex space-x-6 space-x-reverse overflow-x-auto">
          <button
            v-for="navItem in navigationItems"
            :key="navItem.id"
            :class="[
              'flex items-center space-x-2 space-x-reverse px-4 py-2 rounded-lg transition-all duration-200 whitespace-nowrap text-sm font-medium',
              activeNavItem === navItem.id
                ? 'bg-white text-blue-600 shadow-md border border-blue-200'
                : 'text-gray-600 hover:text-gray-900 hover:bg-white/50'
            ]"
            @click="activeNavItem = navItem.id"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="navItem.icon"></path>
            </svg>
            <span>{{ navItem.name }}</span>
            <span v-if="getNavItemCount(navItem.id)" class="bg-blue-100 text-blue-600 text-xs px-2 py-1 rounded-full">{{ getNavItemCount(navItem.id) }}</span>
          </button>
        </nav>
      </div>
    </div>

    <div class="w-full px-4 pb-6">
      <!-- Statistics Cards -->
      <StatsCards />

      <!-- Dashboard Tab -->
      <DashboardTab v-if="activeNavItem === 'dashboard'" :active-nav-item="activeNavItem" />

      <!-- Stock Tab -->
      <StockTab v-if="activeNavItem === 'stock'" :active-nav-item="activeNavItem" />

      <!-- Discount Tab -->
      <DiscountTab v-if="activeNavItem === 'discount'" :active-nav-item="activeNavItem" />

      <!-- Stock Completed Tab -->
      <StockCompletedTab v-if="activeNavItem === 'stock-completed'" :active-nav-item="activeNavItem" />

      <!-- Discount Completed Tab -->
      <DiscountCompletedTab v-if="activeNavItem === 'discount-completed'" :active-nav-item="activeNavItem" />

      <!-- Settings Tab -->
      <SettingsTab v-if="activeNavItem === 'settings'" :active-nav-item="activeNavItem" />

      <!-- Advanced Settings Modal -->
      <AdvancedSettingsModal v-if="showSettings" />
    </div>
  </ClientOnly>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useHead: (head: { title?: string }) => void
</script>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useNotifyRequests } from '~/composables/useNotifyRequests'
// Import components
import TemplateButton from '@/components/common/TemplateButton.vue'
import AdvancedSettingsModal from './AdvancedSettingsModal.vue'
import DashboardTab from './DashboardTab.vue'
import DiscountCompletedTab from './DiscountCompletedTab.vue'
import DiscountTab from './DiscountTab.vue'
import SettingsTab from './SettingsTab.vue'
import StatsCards from './StatsCards.vue'
import StockCompletedTab from './StockCompletedTab.vue'
import StockTab from './StockTab.vue'

const { 
  // Data
  activeNavItem,
  navigationItems,
  showSettings,
  // Methods
  exportToExcel,
  refreshData,
  getNavItemCount
} = useNotifyRequests()

definePageMeta({
  layout: 'admin-main'
})

useHead({
  title: 'مدیریت درخواست‌های اطلاع‌رسانی - داشبورد جامع'
})

// بارگذاری اولیه داده‌ها از API
onMounted(() => {
  refreshData()
})
</script>
