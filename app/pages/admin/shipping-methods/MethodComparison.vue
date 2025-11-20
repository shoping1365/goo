<template>
  <div class="method-comparison">
    <div class="section-header">
      <h3>مقایسه روش‌ها</h3>
      <p>مقایسه ویژگی‌های مختلف روش‌های ارسال</p>
    </div>

    <!-- Comparison Table -->
    <div class="comparison-table">
      <h4>جدول مقایسه</h4>
      <table>
        <thead>
        <tr>
          <th>ویژگی</th>
          <th v-for="method in shippingMethods" :key="method.id">{{ method.name }}</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="feature in comparisonFeatures" :key="feature.id">
          <td>{{ feature.name }}</td>
          <td v-for="method in shippingMethods" :key="method.id">
            {{ getFeatureValue(method.id, feature.id) }}
          </td>
        </tr>
        </tbody>
      </table>
    </div>

    <!-- Comparison Chart -->
    <div class="comparison-chart">
      <h4>نمودار مقایسه</h4>
      <div class="chart-container">
        <canvas ref="comparisonChart"></canvas>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'

const shippingMethods = ref([
  { id: 1, name: 'ارسال سریع' },
  { id: 2, name: 'ارسال استاندارد' },
  { id: 3, name: 'ارسال اقتصادی' }
])

const comparisonFeatures = ref([
  { id: 1, name: 'سرعت' },
  { id: 2, name: 'هزینه' },
  { id: 3, name: 'امنیت' },
  { id: 4, name: 'پوشش' }
])

const getFeatureValue = (_methodId, _featureId) => {
  // Get feature value logic
  return 'متوسط'
}

onMounted(() => {
  // Initialize comparison chart
})
</script>

<style scoped>
.method-comparison {
  padding: 20px;
}

.section-header {
  margin-bottom: 30px;
}

.section-header h3 {
  color: #2c3e50;
  margin-bottom: 10px;
}

.section-header p {
  color: #7f8c8d;
}

.comparison-table,
.comparison-chart {
  background: white;
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 20px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 15px;
}

th, td {
  padding: 12px;
  text-align: center;
  border: 1px solid #ecf0f1;
}

th {
  background-color: #f8f9fa;
  font-weight: 600;
  color: #2c3e50;
}

.chart-container {
  height: 400px;
  margin-top: 15px;
}
</style>
