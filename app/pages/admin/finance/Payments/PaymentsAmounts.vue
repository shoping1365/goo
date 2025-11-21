<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">توزیع مبالغ پرداخت</h3>
        <p class="text-sm text-gray-500 mt-1">Payment Amount Distribution</p>
      </div>
    </div>

    <div class="space-y-4">
      <div 
        v-for="amount in amounts" 
        :key="amount.range"
        class="flex items-center justify-between p-3 bg-gray-50 rounded-lg"
      >
        <div class="flex items-center">
          <div class="w-3 h-3 bg-indigo-500 rounded-full ml-3"></div>
          <div>
            <div class="font-medium text-gray-900">{{ amount.range }}</div>
            <div class="text-sm text-gray-500">{{ amount.percentage }}% از کل</div>
          </div>
        </div>
        <div class="text-left">
          <div class="font-semibold text-gray-900">{{ amount.count }}</div>
          <div class="text-xs text-gray-500">تراکنش</div>
        </div>
      </div>
    </div>

    <div class="mt-6 pt-4 border-t border-gray-200">
      <div class="text-center">
        <div class="text-lg font-semibold text-indigo-600">{{ averageAmount }}</div>
        <div class="text-sm text-gray-500">میانگین مبلغ</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  amounts?: Array<{
    range: string
    count: number
    percentage: number
  }>
}

const props = withDefaults(defineProps<Props>(), {
  amounts: () => []
})

const averageAmount = computed(() => {
  const total = props.amounts.reduce((sum, item) => sum + item.count, 0)
  return total > 0 ? Math.round(total / props.amounts.length) : 0
})
</script> 
