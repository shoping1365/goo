<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <h3 class="text-lg font-bold text-gray-900 mb-4">روند بازدیدها (24 ساعت گذشته)</h3>
    
    <div class="h-48 flex items-end justify-between gap-1 px-2">
      <div 
        v-for="trend in trends" 
        :key="trend.hour"
        class="flex flex-col items-center gap-2 flex-1"
      >
        <div class="relative w-full h-32 flex items-end">
          <div 
            class="w-full bg-gradient-to-t from-blue-400 to-blue-600 rounded-t hover:from-blue-500 hover:to-blue-700 transition-all duration-300 cursor-pointer"
            :style="{ height: (trend.visits / maxTrendVisits) * 100 + '%' }"
            :title="`${trend.hour}:00 - ${trend.visits} بازدید`"
          ></div>
        </div>
        
        <div class="text-xs text-gray-600 font-medium">
          {{ trend.hour }}
        </div>
      </div>
    </div>
    
    <!-- Peak Hours Summary -->
    <div class="mt-6 grid grid-cols-3 gap-6">
      <div class="text-center">
        <div class="text-sm text-gray-500">ساعت پیک</div>
        <div class="text-lg font-bold text-blue-600">{{ peakHour }}:00</div>
      </div>
      <div class="text-center">
        <div class="text-sm text-gray-500">بیشترین بازدید</div>
        <div class="text-lg font-bold text-gray-900">{{ maxTrendVisits.toLocaleString('fa-IR') }}</div>
      </div>
      <div class="text-center">
        <div class="text-sm text-gray-500">کمترین بازدید</div>
        <div class="text-lg font-bold text-gray-600">{{ minTrendVisits.toLocaleString('fa-IR') }}</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { computed } from 'vue'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

interface Trend {
  hour: string
  visits: number
}

const props = defineProps<{
  trends: Trend[]
}>()

const maxTrendVisits = computed(() => 
  Math.max(...props.trends.map(trend => trend.visits))
)

const minTrendVisits = computed(() => 
  Math.min(...props.trends.map(trend => trend.visits))
)

const peakHour = computed(() => {
  const peak = props.trends.reduce((max, trend) => 
    trend.visits > max.visits ? trend : max
  )
  return peak.hour
})
</script> 
