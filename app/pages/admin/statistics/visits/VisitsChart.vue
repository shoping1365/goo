<template>
  <div class="bg-white rounded-xl shadow-lg p-6 h-96">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-bold text-gray-900">نمودار بازدیدهای هفتگی</h3>
      <div class="flex items-center gap-6">
        <div class="flex items-center gap-2">
          <div class="w-3 h-3 bg-blue-500 rounded-full"></div>
          <span class="text-sm text-gray-600">کل بازدیدها</span>
        </div>
        <div class="flex items-center gap-2">
          <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          <span class="text-sm text-gray-600">بازدیدهای یکتا</span>
        </div>
      </div>
    </div>
    
    <div class="h-64 flex items-end justify-between gap-2 px-4">
      <div 
        v-for="(item, index) in data" 
        :key="item.date"
        class="flex flex-col items-center gap-2 flex-1"
      >
        <div class="relative w-full h-48 flex items-end gap-1">
          <!-- Total Visits Bar -->
          <div 
            class="bg-gradient-to-t from-blue-400 to-blue-600 rounded-t-lg transition-all duration-500 hover:from-blue-500 hover:to-blue-700 cursor-pointer group relative"
            :style="{ height: (item.visits / maxVisits) * 100 + '%', width: '60%' }"
            @mouseover="showTooltip(index, 'visits', $event)"
            @mouseleave="hideTooltip"
          >
            <div v-if="tooltipVisible && tooltipIndex === index && tooltipType === 'visits'" 
                 class="absolute bottom-full left-1/2 transform -translate-x-1/2 mb-2 px-3 py-1 bg-gray-900 text-white text-sm rounded-lg z-10">
              {{ item.visits.toLocaleString('fa-IR') }} بازدید
              <div class="absolute top-full left-1/2 transform -translate-x-1/2 w-0 h-0 border-l-4 border-r-4 border-t-4 border-transparent border-t-gray-900"></div>
            </div>
          </div>
          
          <!-- Unique Visits Bar -->
          <div 
            class="bg-gradient-to-t from-green-400 to-green-600 rounded-t-lg transition-all duration-500 hover:from-green-500 hover:to-green-700 cursor-pointer group relative"
            :style="{ height: (item.unique / maxVisits) * 100 + '%', width: '40%' }"
            @mouseover="showTooltip(index, 'unique', $event)"
            @mouseleave="hideTooltip"
          >
            <div v-if="tooltipVisible && tooltipIndex === index && tooltipType === 'unique'" 
                 class="absolute bottom-full left-1/2 transform -translate-x-1/2 mb-2 px-3 py-1 bg-gray-900 text-white text-sm rounded-lg z-10">
              {{ item.unique.toLocaleString('fa-IR') }} یکتا
              <div class="absolute top-full left-1/2 transform -translate-x-1/2 w-0 h-0 border-l-4 border-r-4 border-t-4 border-transparent border-t-gray-900"></div>
            </div>
          </div>
        </div>
        
        <div class="text-xs text-gray-600 font-medium">
          {{ formatDate(item.date) }}
        </div>
      </div>
    </div>
    
    <!-- Chart Summary -->
    <div class="mt-6 grid grid-cols-3 gap-6">
      <div class="text-center">
        <div class="text-sm text-gray-500">میانگین روزانه</div>
        <div class="text-lg font-bold text-gray-900">{{ avgDaily.toLocaleString('fa-IR') }}</div>
      </div>
      <div class="text-center">
        <div class="text-sm text-gray-500">بیشترین بازدید</div>
        <div class="text-lg font-bold text-blue-600">{{ maxVisits.toLocaleString('fa-IR') }}</div>
      </div>
      <div class="text-center">
        <div class="text-sm text-gray-500">کمترین بازدید</div>
        <div class="text-lg font-bold text-gray-600">{{ minVisits.toLocaleString('fa-IR') }}</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

interface ChartData {
  date: string
  visits: number
  unique: number
}

const props = defineProps<{
  data: ChartData[]
}>()

const tooltipVisible = ref(false)
const tooltipIndex = ref(-1)
const tooltipType = ref('')

const maxVisits = computed(() => 
  Math.max(...props.data.map(item => Math.max(item.visits, item.unique)))
)

const minVisits = computed(() => 
  Math.min(...props.data.map(item => item.visits))
)

const avgDaily = computed(() => 
  Math.round(props.data.reduce((sum, item) => sum + item.visits, 0) / props.data.length)
)

const formatDate = (dateStr: string) => {
  const parts = dateStr.split('/')
  return `${parts[2]}/${parts[1]}`
}

const showTooltip = (index: number, type: string, event: MouseEvent) => {
  tooltipIndex.value = index
  tooltipType.value = type
  tooltipVisible.value = true
}

const hideTooltip = () => {
  tooltipVisible.value = false
}
</script> 
