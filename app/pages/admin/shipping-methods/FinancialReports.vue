<template>
  <div class="financial-reports">
    <div class="section-header">
      <h3>گزارش‌های مالی</h3>
      <p>تحلیل مالی روش‌های ارسال و سودآوری</p>
    </div>

    <!-- Date Range Selector -->
    <div class="date-range-selector">
      <div class="form-group">
        <label>بازه زمانی:</label>
        <select v-model="selectedPeriod" @change="updateFinancialData">
          <option value="7">7 روز گذشته</option>
          <option value="30">30 روز گذشته</option>
          <option value="90">90 روز گذشته</option>
          <option value="365">یک سال گذشته</option>
          <option value="custom">سفارشی</option>
        </select>
      </div>
      
      <div v-if="selectedPeriod === 'custom'" class="custom-date-range">
        <div class="form-group">
          <label>از تاریخ:</label>
          <input v-model="customStartDate" type="date" @change="updateFinancialData">
        </div>
        <div class="form-group">
          <label>تا تاریخ:</label>
          <input v-model="customEndDate" type="date" @change="updateFinancialData">
        </div>
      </div>
    </div>

    <!-- Financial Summary Cards -->
    <div class="financial-summary">
      <div class="card revenue-card">
        <div class="card-icon">
          <i class="fas fa-dollar-sign"></i>
        </div>
        <div class="card-content">
          <h4>درآمد کل</h4>
          <div class="number">{{ formatCurrency(financialData.totalRevenue) }}</div>
          <div class="trend" :class="financialData.revenueGrowth > 0 ? 'positive' : 'negative'">
            {{ financialData.revenueGrowth > 0 ? '+' : '' }}{{ financialData.revenueGrowth }}%
          </div>
        </div>
      </div>

      <div class="card cost-card">
        <div class="card-icon">
          <i class="fas fa-coins"></i>
        </div>
        <div class="card-content">
          <h4>هزینه کل</h4>
          <div class="number">{{ formatCurrency(financialData.totalCost) }}</div>
          <div class="trend" :class="financialData.costChange > 0 ? 'negative' : 'positive'">
            {{ financialData.costChange > 0 ? '+' : '' }}{{ financialData.costChange }}%
          </div>
        </div>
      </div>

      <div class="card profit-card">
        <div class="card-icon">
          <i class="fas fa-chart-line"></i>
        </div>
        <div class="card-content">
          <h4>سود خالص</h4>
          <div class="number">{{ formatCurrency(financialData.netProfit) }}</div>
          <div class="trend" :class="financialData.profitGrowth > 0 ? 'positive' : 'negative'">
            {{ financialData.profitGrowth > 0 ? '+' : '' }}{{ financialData.profitGrowth }}%
          </div>
        </div>
      </div>

      <div class="card margin-card">
        <div class="card-icon">
          <i class="fas fa-percentage"></i>
        </div>
        <div class="card-content">
          <h4>حاشیه سود</h4>
          <div class="number">{{ financialData.profitMargin }}%</div>
          <div class="trend" :class="financialData.marginChange > 0 ? 'positive' : 'negative'">
            {{ financialData.marginChange > 0 ? '+' : '' }}{{ financialData.marginChange }}%
          </div>
        </div>
      </div>
    </div>

    <!-- Revenue Analysis -->
    <div class="revenue-analysis">
      <h4>تحلیل درآمد</h4>
      <div class="analysis-grid">
        <div class="chart-container">
          <h5>درآمد بر اساس روش ارسال</h5>
          <div class="chart-wrapper">
            <canvas ref="revenueChart" width="400" height="200"></canvas>
          </div>
        </div>

        <div class="chart-container">
          <h5>روند درآمد ماهانه</h5>
          <div class="chart-wrapper">
            <canvas ref="trendChart" width="400" height="200"></canvas>
          </div>
        </div>
      </div>
    </div>

    <!-- Cost Breakdown -->
    <div class="cost-breakdown">
      <h4>تجزیه هزینه‌ها</h4>
      <div class="cost-details">
        <div class="cost-item">
          <div class="cost-label">هزینه سوخت</div>
          <div class="cost-bar">
            <div class="bar" :style="{ width: costBreakdown.fuel + '%' }"></div>
            <span>{{ formatCurrency(costBreakdown.fuelAmount) }} ({{ costBreakdown.fuel }}%)</span>
          </div>
        </div>

        <div class="cost-item">
          <div class="cost-label">هزینه نیروی انسانی</div>
          <div class="cost-bar">
            <div class="bar" :style="{ width: costBreakdown.labor + '%' }"></div>
            <span>{{ formatCurrency(costBreakdown.laborAmount) }} ({{ costBreakdown.labor }}%)</span>
          </div>
        </div>

        <div class="cost-item">
          <div class="cost-label">هزینه نگهداری</div>
          <div class="cost-bar">
            <div class="bar" :style="{ width: costBreakdown.maintenance + '%' }"></div>
            <span>{{ formatCurrency(costBreakdown.maintenanceAmount) }} ({{ costBreakdown.maintenance }}%)</span>
          </div>
        </div>

        <div class="cost-item">
          <div class="cost-label">هزینه بیمه</div>
          <div class="cost-bar">
            <div class="bar" :style="{ width: costBreakdown.insurance + '%' }"></div>
            <span>{{ formatCurrency(costBreakdown.insuranceAmount) }} ({{ costBreakdown.insurance }}%)</span>
          </div>
        </div>

        <div class="cost-item">
          <div class="cost-label">سایر هزینه‌ها</div>
          <div class="cost-bar">
            <div class="bar" :style="{ width: costBreakdown.other + '%' }"></div>
            <span>{{ formatCurrency(costBreakdown.otherAmount) }} ({{ costBreakdown.other }}%)</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Profitability Table -->
    <div class="profitability-table">
      <h4>جدول سودآوری روش‌های ارسال</h4>
      <div class="table-container">
        <table>
          <thead>
            <tr>
              <th>روش ارسال</th>
              <th>درآمد</th>
              <th>هزینه</th>
              <th>سود</th>
              <th>حاشیه سود</th>
              <th>نرخ بازگشت</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="method in profitabilityData" :key="method.id">
              <td>{{ method.name }}</td>
              <td>{{ formatCurrency(method.revenue) }}</td>
              <td>{{ formatCurrency(method.cost) }}</td>
              <td :class="method.profit >= 0 ? 'positive' : 'negative'">
                {{ formatCurrency(method.profit) }}
              </td>
              <td :class="method.margin >= 0 ? 'positive' : 'negative'">
                {{ method.margin }}%
              </td>
              <td :class="method.roi >= 0 ? 'positive' : 'negative'">
                {{ method.roi }}%
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Export Options -->
    <div class="export-section">
      <h4>خروجی گزارش مالی</h4>
      <div class="export-buttons">
        <button class="btn btn-secondary" @click="exportToExcel">
          <i class="fas fa-file-excel"></i>
          خروجی اکسل
        </button>
        <button class="btn btn-secondary" @click="exportToPDF">
          <i class="fas fa-file-pdf"></i>
          خروجی PDF
        </button>
        <button class="btn btn-secondary" @click="printReport">
          <i class="fas fa-print"></i>
          چاپ گزارش
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { nextTick, onMounted, reactive, ref } from 'vue';

// Props
interface Props {
  shippingMethods?: unknown[]
}

defineProps<Props>()

// Emits
const emit = defineEmits<{
  exportData: [data: unknown, format: string]
}>()

// Reactive data
const selectedPeriod = ref('30')
const customStartDate = ref('')
const customEndDate = ref('')

const financialData = reactive({
  totalRevenue: 0,
  revenueGrowth: 0,
  totalCost: 0,
  costChange: 0,
  netProfit: 0,
  profitGrowth: 0,
  profitMargin: 0,
  marginChange: 0
})

const costBreakdown = reactive({
  fuel: 0,
  fuelAmount: 0,
  labor: 0,
  laborAmount: 0,
  maintenance: 0,
  maintenanceAmount: 0,
  insurance: 0,
  insuranceAmount: 0,
  other: 0,
  otherAmount: 0
})

const profitabilityData = ref([
  {
    id: 1,
    name: 'ارسال استاندارد',
    revenue: 450000000,
    cost: 320000000,
    profit: 130000000,
    margin: 28.9,
    roi: 40.6
  },
  {
    id: 2,
    name: 'ارسال سریع',
    revenue: 680000000,
    cost: 520000000,
    profit: 160000000,
    margin: 23.5,
    roi: 30.8
  },
  {
    id: 3,
    name: 'ارسال رایگان',
    revenue: 280000000,
    cost: 250000000,
    profit: 30000000,
    margin: 10.7,
    roi: 12.0
  }
])

// Chart references
const revenueChart = ref<HTMLCanvasElement>()
const trendChart = ref<HTMLCanvasElement>()

// Methods
const updateFinancialData = async () => {
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // Update financial data
    financialData.totalRevenue = Math.floor(Math.random() * 2000000000) + 1000000000
    financialData.revenueGrowth = Math.floor(Math.random() * 30) - 10
    financialData.totalCost = Math.floor(Math.random() * 1500000000) + 800000000
    financialData.costChange = Math.floor(Math.random() * 20) - 10
    financialData.netProfit = financialData.totalRevenue - financialData.totalCost
    financialData.profitGrowth = Math.floor(Math.random() * 40) - 20
    financialData.profitMargin = Math.round((financialData.netProfit / financialData.totalRevenue) * 100)
    financialData.marginChange = Math.floor(Math.random() * 15) - 7
    
    // Update cost breakdown
    const totalCost = financialData.totalCost
    costBreakdown.fuel = 35
    costBreakdown.fuelAmount = Math.round(totalCost * 0.35)
    costBreakdown.labor = 25
    costBreakdown.laborAmount = Math.round(totalCost * 0.25)
    costBreakdown.maintenance = 20
    costBreakdown.maintenanceAmount = Math.round(totalCost * 0.20)
    costBreakdown.insurance = 15
    costBreakdown.insuranceAmount = Math.round(totalCost * 0.15)
    costBreakdown.other = 5
    costBreakdown.otherAmount = Math.round(totalCost * 0.05)
    
    updateCharts()
  } catch (error) {
    console.error('Error updating financial data:', error)
  }
}

const updateCharts = () => {
  nextTick(() => {
    createRevenueChart()
    createTrendChart()
  })
}

const createRevenueChart = () => {
  if (!revenueChart.value) return
  
  const ctx = revenueChart.value.getContext('2d')
  if (!ctx) return
  
  // Simple chart implementation without Chart.js
  const canvas = revenueChart.value
  const width = canvas.width
  const height = canvas.height
  
  ctx.clearRect(0, 0, width, height)
  
  // Draw simple bar chart
  const data = profitabilityData.value.map(item => item.revenue)
  const maxValue = Math.max(...data)
  const barWidth = width / data.length - 10
  
  data.forEach((value, index) => {
    const barHeight = (value / maxValue) * (height - 40)
    const x = index * (barWidth + 10) + 5
    const y = height - barHeight - 20
    
    ctx.fillStyle = '#3498db'
    ctx.fillRect(x, y, barWidth, barHeight)
    
    // Draw value label
    ctx.fillStyle = '#2c3e50'
    ctx.font = '12px Arial'
    ctx.textAlign = 'center'
    ctx.fillText(formatCurrency(value), x + barWidth / 2, y - 5)
  })
}

const createTrendChart = () => {
  if (!trendChart.value) return
  
  const ctx = trendChart.value.getContext('2d')
  if (!ctx) return
  
  // Simple line chart implementation
  const canvas = trendChart.value
  const width = canvas.width
  const height = canvas.height
  
  ctx.clearRect(0, 0, width, height)
  
  // Generate sample trend data
  const data = Array.from({ length: 12 }, () => Math.floor(Math.random() * 100000000) + 50000000)
  const maxValue = Math.max(...data)
  
  ctx.strokeStyle = '#e74c3c'
  ctx.lineWidth = 2
  ctx.beginPath()
  
  data.forEach((value, index) => {
    const x = (index / (data.length - 1)) * (width - 40) + 20
    const y = height - ((value / maxValue) * (height - 40) + 20)
    
    if (index === 0) {
      ctx.moveTo(x, y)
    } else {
      ctx.lineTo(x, y)
    }
  })
  
  ctx.stroke()
}

const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR'
  }).format(amount)
}

const exportToExcel = () => {
  const data = {
    financialData,
    costBreakdown,
    profitabilityData: profitabilityData.value,
    period: selectedPeriod.value
  }
  emit('exportData', data, 'excel')
}

const exportToPDF = () => {
  const data = {
    financialData,
    costBreakdown,
    profitabilityData: profitabilityData.value,
    period: selectedPeriod.value
  }
  emit('exportData', data, 'pdf')
}

const printReport = () => {
  window.print()
}

// Lifecycle
onMounted(() => {
  updateFinancialData()
})
</script>

<style scoped>
.financial-reports {
  padding: 20px;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.section-header {
  margin-bottom: 30px;
  text-align: center;
}

.section-header h3 {
  color: #2c3e50;
  margin-bottom: 10px;
  font-size: 24px;
}

.section-header p {
  color: #7f8c8d;
  font-size: 14px;
}

.date-range-selector {
  display: flex;
  gap: 20px;
  margin-bottom: 30px;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 8px;
  flex-wrap: wrap;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.form-group label {
  font-weight: 600;
  color: #2c3e50;
  font-size: 14px;
}

.form-group select,
.form-group input {
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.custom-date-range {
  display: flex;
  gap: 15px;
}

.financial-summary {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
  margin-bottom: 30px;
}

.card {
  padding: 20px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  gap: 15px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
  color: white;
}

.revenue-card {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.cost-card {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
}

.profit-card {
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
}

.margin-card {
  background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%);
}

.card-icon {
  font-size: 24px;
  opacity: 0.8;
}

.card-content h4 {
  margin: 0 0 5px 0;
  font-size: 14px;
  opacity: 0.9;
}

.number {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 5px;
}

.trend {
  font-size: 12px;
  opacity: 0.8;
}

.trend.positive {
  color: #2ecc71;
}

.trend.negative {
  color: #e74c3c;
}

.revenue-analysis {
  margin-bottom: 30px;
}

.revenue-analysis h4 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.analysis-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  gap: 30px;
}

.chart-container {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
}

.chart-container h5 {
  margin-bottom: 20px;
  color: #2c3e50;
  text-align: center;
}

.chart-wrapper {
  position: relative;
  height: 300px;
}

.cost-breakdown {
  margin-bottom: 30px;
}

.cost-breakdown h4 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.cost-details {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
}

.cost-item {
  margin-bottom: 15px;
}

.cost-label {
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 5px;
}

.cost-bar {
  display: flex;
  align-items: center;
  gap: 10px;
}

.bar {
  height: 12px;
  background: #3498db;
  border-radius: 6px;
  min-width: 100px;
}

.profitability-table {
  margin-bottom: 30px;
}

.profitability-table h4 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.table-container {
  overflow-x: auto;
}

table {
  width: 100%;
  border-collapse: collapse;
  background: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

th, td {
  padding: 12px;
  text-align: right;
  border-bottom: 1px solid #eee;
}

th {
  background: #f8f9fa;
  font-weight: 600;
  color: #2c3e50;
}

.positive {
  color: #27ae60;
  font-weight: 600;
}

.negative {
  color: #e74c3c;
  font-weight: 600;
}

.export-section {
  text-align: center;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 8px;
}

.export-section h4 {
  margin-bottom: 15px;
  color: #2c3e50;
}

.export-buttons {
  display: flex;
  gap: 15px;
  justify-content: center;
  flex-wrap: wrap;
}

.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  display: flex;
  align-items: center;
  gap: 8px;
  transition: all 0.3s ease;
}

.btn-secondary {
  background: #6c757d;
  color: white;
}

.btn-secondary:hover {
  background: #5a6268;
  transform: translateY(-2px);
}

@media (max-width: 768px) {
  .financial-summary {
    grid-template-columns: 1fr;
  }
  
  .analysis-grid {
    grid-template-columns: 1fr;
  }
  
  .custom-date-range {
    flex-direction: column;
  }
  
  .export-buttons {
    flex-direction: column;
  }
}
</style> 
