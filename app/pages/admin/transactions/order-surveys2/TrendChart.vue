<template>
  <div class="relative">
    <canvas ref="chartCanvas"></canvas>
    <div v-if="loading" class="absolute inset-0 bg-white/50 flex items-center justify-center">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch, nextTick } from 'vue'
import Chart from 'chart.js/auto'

interface ChartData {
  labels: string[]
  datasets: {
    label: string
    data: number[]
    borderColor: string
    backgroundColor: string
    tension?: number
  }[]
}

const props = defineProps<{
  data: ChartData
}>()

const chartCanvas = ref<HTMLCanvasElement>()
const chart = ref<Chart>()
const loading = ref(false)

const createChart = () => {
  if (!chartCanvas.value) return

  const ctx = chartCanvas.value.getContext('2d')
  if (!ctx) return

  chart.value = new Chart(ctx, {
    type: 'line',
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
        },
        tooltip: {
          mode: 'index',
          intersect: false,
          backgroundColor: 'rgba(0, 0, 0, 0.8)',
          titleColor: '#fff',
          bodyColor: '#fff',
          borderColor: '#6366f1',
          borderWidth: 1
        }
      },
      scales: {
        x: {
          display: true,
          title: {
            display: true,
            text: 'زمان',
            font: {
              family: 'IRANSans',
              size: 14
            }
          },
          ticks: {
            font: {
              family: 'IRANSans',
              size: 12
            }
          }
        },
        y: {
          display: true,
          title: {
            display: true,
            text: 'تعداد',
            font: {
              family: 'IRANSans',
              size: 14
            }
          },
          ticks: {
            font: {
              family: 'IRANSans',
              size: 12
            }
          }
        }
      },
      interaction: {
        mode: 'nearest',
        axis: 'x',
        intersect: false
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

watch(() => props.data, updateChart, { deep: true })

onMounted(async () => {
  await nextTick()
  createChart()
})
</script>
