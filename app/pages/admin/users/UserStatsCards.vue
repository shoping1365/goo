<template>
  <div class="grid grid-cols-2 md:grid-cols-5 gap-3 mb-4">
    <!-- Total Users -->
    <div
      class="cursor-pointer transition-all duration-200 hover:scale-105"
      @click="clearFilters"
    >
      <TemplateCard 
        title="کل کاربران" 
        :value="stats.total" 
        variant="blue"
      >
        <template #icon>
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M9 20H4v-2a3 3 0 015.356-1.857M15 10a4 4 0 11-8 0 4 4 0 018 0zm6 6v2a2 2 0 01-2 2h-1.5M3 16v2a2 2 0 002 2h1.5" />
          </svg>
        </template>
      </TemplateCard>
    </div>
    <!-- Active -->
    <div
      class="cursor-pointer transition-all duration-200 hover:scale-105"
      @click="filterStatus('active')"
    >
      <TemplateCard 
        title="فعال" 
        :value="stats.active" 
        variant="green"
      >
        <template #icon>
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          </svg>
        </template>
      </TemplateCard>
    </div>
    <!-- Blocked -->
    <div
      class="cursor-pointer transition-all duration-200 hover:scale-105"
      @click="filterStatus('blocked')"
    >
      <TemplateCard 
        title="مسدود شده" 
        :value="stats.blocked" 
        variant="red"
      >
        <template #icon>
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </template>
      </TemplateCard>
    </div>
    <!-- Online -->
    <div
      class="cursor-pointer transition-all duration-200 hover:scale-105"
      @click="filterOnline"
    >
      <TemplateCard 
        title="آنلاین" 
        :value="stats.online" 
        variant="purple"
      >
        <template #icon>
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <circle cx="12" cy="12" r="4" stroke="currentColor" stroke-width="2" fill="none" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 8v1m0 4v1m-8-6v1m0 4v1" />
          </svg>
        </template>
      </TemplateCard>
    </div>
    <div class="cursor-pointer transition-all duration-200 hover:scale-105" @click="exportExcel">
      <TemplateCard 
        title="اکسپورت اکسل" 
        value="" 
        variant="amber"
      >
        <template #icon>
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v2a2 2 0 002 2h12a2 2 0 002-2v-2M7 10V6a5 5 0 0110 0v4M5 20h14a2 2 0 002-2v-2a2 2 0 00-2-2H5a2 2 0 00-2 2v2a2 2 0 002 2z" />
          </svg>
        </template>
      </TemplateCard>
    </div>
  </div>
</template>
<script setup lang="ts">
import TemplateCard from '@/components/common/TemplateCard.vue';
import { computed } from 'vue';

// Props
import type { User } from '~/types/user';

const props = defineProps<{
  users: User[];
  filters: Record<string, unknown>;
}>();

// Emits
const emit = defineEmits<{
  'update:filters': [filters: Record<string, unknown>];
  'export-excel': [];
}>();

// Computed stats
const stats = computed(() => {
  const total = props.users.length;
  const active = props.users.filter(u => u.status === 'active').length;
  const blocked = props.users.filter(u => u.status === 'blocked').length;
  // شمارش آنلاین بر اساس فیلد محاسبه‌شده در صفحه والد (threshold قابل‌تنظیم)
  const online = props.users.filter(u => u.online).length;
  
  return { total, active, blocked, online };
});



function filterStatus(status: string) {
  const newFilters = { ...props.filters, status, onlineOnly: false };
  emit('update:filters', newFilters);
}

function filterOnline() {
  const newFilters = { ...props.filters, status: '', onlineOnly: true };
  emit('update:filters', newFilters);
}

function clearFilters() {
  const newFilters = { ...props.filters, status: '', role: '', onlineOnly: false };
  emit('update:filters', newFilters);
}

function exportExcel() {
  emit('export-excel');
}
</script> 
