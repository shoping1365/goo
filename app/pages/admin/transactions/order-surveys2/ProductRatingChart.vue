<template>
  <div class="relative">
    <canvas ref="chartCanvas"></canvas>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch, nextTick } from 'vue'
import Chart from 'chart.js/auto'

const props = defineProps<{ data: any }>()
const chartCanvas = ref<HTMLCanvasElement>()
let chart: Chart | null = null

const createChart = () => {
  if (!chartCanvas.value) return
  const ctx = chartCanvas.value.getContext('2d')
  if (!ctx) return
  if (chart) chart.destroy()
  chart = new Chart(ctx, {
    type: 'line',
    data: props.data,
    options: {
      responsive: true,
      plugins: {
        legend: { position: 'top' }
      }
    }
  })
}

watch(() => props.data, createChart, { deep: true })
onMounted(() => nextTick(createChart))
</script>
