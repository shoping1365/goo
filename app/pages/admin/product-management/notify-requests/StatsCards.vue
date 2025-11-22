<template>
  <div class="w-full px-4 pb-6">
    <!-- Statistics Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
      <TemplateCard
        title="کل درخواست‌ها"
        :value="totalRequests"
        variant="blue"
      >
        <template #icon>
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
          </svg>
        </template>
      </TemplateCard>

      <TemplateCard
        title="در انتظار"
        :value="pendingRequests"
        variant="orange"
      >
        <template #icon>
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
        </template>
      </TemplateCard>

      <TemplateCard
        title="تکمیل شده"
        :value="completedRequests"
        variant="green"
      >
        <template #icon>
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
        </template>
      </TemplateCard>

      <TemplateCard
        title="امروز"
        :value="todayRequests"
        variant="purple"
      >
        <template #icon>
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
          </svg>
        </template>
      </TemplateCard>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import TemplateCard from '@/components/common/TemplateCard.vue';
import { useNotifyRequests } from '~/composables/useNotifyRequests';

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

const { totalRequests, pendingRequests, completedRequests, todayRequests } = useNotifyRequests()
</script>

<style scoped>

</style>
