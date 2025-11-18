<template>
  <div class="relative">
    <!-- انتخاب نوع نمودار -->
    <div class="mb-4 flex justify-center space-x-4 space-x-reverse">
      <button 
        v-for="type in chartTypes" 
        :key="type.value"
        @click="changeChartType(type.value)"
        :class="[
          'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
          selectedType === type.value 
            ? 'bg-blue-500 text-white shadow-md' 
            : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
        ]"
      >
        {{ type.label }}
      </button>
    </div>
    
    <canvas ref="chartCanvas" class="w-full h-full"></canvas>

    <!-- صفحه‌بندی -->
    <div v-if="pagination && pagination.total_pages > 1" class="mt-4">
      <Pagination
        :current-page="pagination.current_page"
        :total-pages="pagination.total_pages"
        :total="pagination.total_items"
        :items-per-page="pagination.items_per_page"
        @page-changed="handlePageChange"
        @items-per-page-changed="handleItemsPerPageChange"
      />
    </div>
  </div>
</template>

<script setup>
import { Chart, registerables } from 'chart.js'

// ثبت تمام کامپوننت‌های Chart.js
Chart.register(...registerables)

const props = defineProps({
  data: {
    type: Object,
    required: true,
    default: () => ({
      daily_traffic: [],
      weekly_traffic: [],
      monthly_traffic: [],
      pagination: {
        current_page: 1,
        total_pages: 1,
        total_items: 0,
        items_per_page: 10,
        has_next: false,
        has_prev: false
      }
    })
  }
})

const emit = defineEmits(['chart-type-changed', 'page-changed', 'items-per-page-changed'])

const chartCanvas = ref(null)
let chartInstance = null
const selectedType = ref('daily')

// انواع نمودار
const chartTypes = [
  { value: 'daily', label: 'روزانه' },
  { value: 'weekly', label: 'هفتگی' },
  { value: 'monthly', label: 'ماهانه' }
]

// به‌روزرسانی نمودار
const updateChart = () => {
  if (!chartCanvas.value) return

  // حذف نمودار قبلی
  if (chartInstance) {
    chartInstance.destroy()
  }

  const ctx = chartCanvas.value.getContext('2d')
  
  let labels = []
  let chartData = []
  let title = ''

  // انتخاب داده‌ها بر اساس نوع نمودار
  switch (selectedType.value) {
    case 'daily':
      if (props.data.daily_traffic) {
        labels = props.data.daily_traffic.map(item => {
          const date = new Date(item.date)
          return date.toLocaleDateString('fa-IR', { month: 'short', day: 'numeric' })
        })
        chartData = props.data.daily_traffic.map(item => item.count)
        title = 'ترافیک روزانه (7 روز گذشته)'
      }
      break
    case 'weekly':
      if (props.data.weekly_traffic) {
        labels = props.data.weekly_traffic.map(item => `هفته ${item.week}`)
        chartData = props.data.weekly_traffic.map(item => item.count)
        title = 'ترافیک هفتگی (12 هفته گذشته)'
      }
      break
    case 'monthly':
      if (props.data.monthly_traffic) {
        labels = props.data.monthly_traffic.map(item => {
          const [year, month] = item.month.split('-')
          const date = new Date(parseInt(year), parseInt(month) - 1)
          return date.toLocaleDateString('fa-IR', { year: 'numeric', month: 'short' })
        })
        chartData = props.data.monthly_traffic.map(item => item.count)
        title = 'ترافیک ماهانه (12 ماه گذشته)'
      }
      break
  }

  chartInstance = new Chart(ctx, {
    type: 'line',
    data: {
      labels,
      datasets: [{
        label: 'تعداد درخواست‌ها',
        data: chartData,
        borderColor: 'rgb(59, 130, 246)',
        backgroundColor: 'rgba(59, 130, 246, 0.1)',
        borderWidth: 2,
        fill: true,
        tension: 0.4,
        pointBackgroundColor: 'rgb(59, 130, 246)',
        pointBorderColor: '#fff',
        pointBorderWidth: 2,
        pointRadius: 4,
        pointHoverRadius: 6
      }]
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        legend: {
          display: false
        },
        tooltip: {
          mode: 'index',
          intersect: false,
          backgroundColor: 'rgba(0, 0, 0, 0.8)',
          titleColor: '#fff',
          bodyColor: '#fff',
          borderColor: 'rgb(59, 130, 246)',
          borderWidth: 1,
          callbacks: {
            title: function(context) {
              return title
            }
          }
        }
      },
      scales: {
        x: {
          grid: {
            color: 'rgba(0, 0, 0, 0.1)'
          },
          ticks: {
            color: '#6b7280',
            maxRotation: 45
          }
        },
        y: {
          beginAtZero: true,
          grid: {
            color: 'rgba(0, 0, 0, 0.1)'
          },
          ticks: {
            stepSize: 1,
            color: '#6b7280'
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

// تغییر نوع نمودار
const changeChartType = (type) => {
  selectedType.value = type
  emit('chart-type-changed', type)
}

// تغییر صفحه
const handlePageChange = (page) => {
  emit('page-changed', page)
}

// تغییر تعداد آیتم در هر صفحه
const handleItemsPerPageChange = (itemsPerPage) => {
  emit('items-per-page-changed', itemsPerPage)
}

// محاسبه اطلاعات صفحه‌بندی
const pagination = computed(() => props.data.pagination)

// نظارت بر تغییرات props و نوع نمودار
watch([() => props.data, selectedType], () => {
  updateChart()
}, { deep: true })

onMounted(() => {
  updateChart()
})

onUnmounted(() => {
  if (chartInstance) {
    chartInstance.destroy()
  }
})
</script> 