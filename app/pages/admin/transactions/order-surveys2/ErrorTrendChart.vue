<template>
  <div class="relative">
    <canvas ref="chartCanvas"></canvas>
  </div>
</template>

<script setup lang="ts">
import Chart, { type ChartData } from 'chart.js/auto';
import { nextTick, onMounted, ref, watch } from 'vue';

const props = defineProps<{ data: ChartData<'line'> }>()
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
      },
      scales: {
        y: {
          beginAtZero: true
        }
      }
    }
  })
}

watch(() => props.data, createChart, { deep: true })
onMounted(() => nextTick(createChart))
</script>

<style scoped>
/* Custom styles if needed */
</style>
