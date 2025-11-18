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
    data: number[]
    backgroundColor: string[]
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
    type: 'doughnut',
    data: props.data,
    options: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        legend: {
          position: 'bottom',
          labels: {
            font: {
              family: 'IRANSans',
              size: 12
            },
            usePointStyle: true,
            padding: 20
          }
        },
        tooltip: {
          backgroundColor: 'rgba(0, 0, 0, 0.8)',
          titleColor: '#fff',
          bodyColor: '#fff',
          borderColor: '#6366f1',
          borderWidth: 1,
          callbacks: {
            label: function(context) {
              const label = context.label || ''
              const value = context.parsed
              const total = context.dataset.data.reduce((a: number, b: number) => a + b, 0)
              const percentage = ((value / total) * 100).toFixed(1)
              return `${label}: ${value} (${percentage}%)`
            }
          }
        }
      },
      cutout: '60%'
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
