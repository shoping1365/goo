<template>
  <div class="relative">
    <canvas ref="chartCanvas"></canvas>
    <div class="mt-4 grid grid-cols-2 gap-6">
      <div class="text-center">
        <div class="text-2xl font-bold text-blue-600">{{ totalSent }}</div>
        <div class="text-sm text-gray-600">کل ارسال‌ها</div>
      </div>
      <div class="text-center">
        <div class="text-2xl font-bold text-green-600">{{ totalResponses }}</div>
        <div class="text-sm text-gray-600">کل پاسخ‌ها</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch, nextTick } from 'vue'
import Chart from 'chart.js/auto'

interface ChartData {
  labels: string[]
  datasets: {
    label: string
    data: number[]
    borderColor: string
    backgroundColor: string
  }[]
}

const props = defineProps<{
  data: ChartData
}>()

const chartCanvas = ref<HTMLCanvasElement>()
const chart = ref<Chart>()

const totalSent = computed(() => {
  return props.data.datasets[0]?.data.reduce((sum, val) => sum + val, 0) || 0
})

const totalResponses = computed(() => {
  return props.data.datasets[1]?.data.reduce((sum, val) => sum + val, 0) || 0
})

const createChart = () => {
  if (!chartCanvas.value) return

  const ctx = chartCanvas.value.getContext('2d')
  if (!ctx) return

  chart.value = new Chart(ctx, {
    type: 'bar',
    data: props.data,
    options: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        legend: {
          position: 'top',
          labels: {
            font: {
              family: 'IRANSans',
              size: 12
            }
          }
        }
      },
      scales: {
        x: {
          grid: {
            display: false
          },
          ticks: {
            font: {
              family: 'IRANSans',
              size: 11
            }
          }
        },
        y: {
          beginAtZero: true,
          ticks: {
            font: {
              family: 'IRANSans',
              size: 11
            }
          }
        }
      }
    }
  })
}

const updateChart = () => {
  if (chart.value) {
    chart.value.data = props.data
    chart.value.update()
  }
}

onMounted(() => {
  createChart()
})

watch(() => props.data, () => {
  nextTick(() => {
    updateChart()
  })
}, { deep: true })
</script>

<style scoped>
/* Custom styles if needed */
</style>
