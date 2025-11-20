<template>
  <div class="international-currencies">
    <div class="section-header">
      <h3>ارزهای مختلف</h3>
      <p>مدیریت ارزهای مختلف برای روش‌های ارسال بین‌المللی</p>
    </div>

    <!-- Currency List -->
    <div class="currency-list">
      <h4>لیست ارزهای فعال</h4>
      <table>
        <thead>
        <tr>
          <th>نام ارز</th>
          <th>نماد</th>
          <th>نرخ تبدیل</th>
          <th>آخرین بروزرسانی</th>
          <th>وضعیت</th>
          <th>عملیات</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="currency in currencies" :key="currency.id">
          <td>{{ currency.name }}</td>
          <td>{{ currency.symbol }}</td>
          <td>{{ currency.exchangeRate }}</td>
          <td>{{ currency.lastUpdate }}</td>
          <td>
            <span :class="['status', currency.status]">{{ currency.statusText }}</span>
          </td>
          <td>
            <button class="btn-edit" @click="editCurrency(currency)">ویرایش</button>
            <button class="btn-delete" @click="deleteCurrency(currency.id)">حذف</button>
          </td>
        </tr>
        </tbody>
      </table>
    </div>

    <!-- Add New Currency -->
    <div class="add-currency-section">
      <h4>افزودن ارز جدید</h4>
      <form @submit.prevent="addNewCurrency">
        <div class="form-row">
          <div class="form-group">
            <label>نام ارز:</label>
            <input v-model="newCurrency.name" type="text" required>
          </div>
          <div class="form-group">
            <label>نماد:</label>
            <input v-model="newCurrency.symbol" type="text" required>
          </div>
        </div>
        <div class="form-row">
          <div class="form-group">
            <label>نرخ تبدیل:</label>
            <input v-model="newCurrency.exchangeRate" type="number" step="0.01" required>
          </div>
          <div class="form-group">
            <label>وضعیت:</label>
            <select v-model="newCurrency.status" required>
              <option value="active">فعال</option>
              <option value="inactive">غیرفعال</option>
            </select>
          </div>
        </div>
        <button type="submit" class="btn-primary">افزودن ارز</button>
      </form>
    </div>

    <!-- Exchange Rate History -->
    <div class="exchange-rate-history">
      <h4>تاریخچه نرخ تبدیل</h4>
      <div class="chart-container">
        <canvas ref="exchangeRateChart"></canvas>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, reactive, ref } from 'vue'

const currencies = ref([
  {
    id: 1,
    name: 'دلار آمریکا',
    symbol: 'USD',
    exchangeRate: 1.0,
    lastUpdate: '1402/10/15',
    status: 'active',
    statusText: 'فعال'
  },
  {
    id: 2,
    name: 'یورو',
    symbol: 'EUR',
    exchangeRate: 0.85,
    lastUpdate: '1402/10/15',
    status: 'active',
    statusText: 'فعال'
  },
  {
    id: 3,
    name: 'پوند انگلیس',
    symbol: 'GBP',
    exchangeRate: 0.73,
    lastUpdate: '1402/10/15',
    status: 'active',
    statusText: 'فعال'
  }
])

const newCurrency = reactive({
  name: '',
  symbol: '',
  exchangeRate: '',
  status: 'active'
})

const editCurrency = (_currency) => {
  // Edit currency logic
}

const deleteCurrency = (_id) => {
  // Delete currency logic
}

const addNewCurrency = () => {
  // Add new currency logic
}

onMounted(() => {
  // Initialize exchange rate chart
})
</script>

<style scoped>
.international-currencies {
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

.currency-list,
.add-currency-section,
.exchange-rate-history {
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
  text-align: right;
  border-bottom: 1px solid #ecf0f1;
}

th {
  background-color: #f8f9fa;
  font-weight: 600;
  color: #2c3e50;
}

.status {
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 500;
}

.status.active {
  background-color: #d4edda;
  color: #155724;
}

.status.inactive {
  background-color: #f8d7da;
  color: #721c24;
}

.btn-edit,
.btn-delete {
  padding: 6px 12px;
  border: none;
  border-radius: 4px;
  margin-left: 5px;
  cursor: pointer;
  font-size: 12px;
}

.btn-edit {
  background-color: #007bff;
  color: white;
}

.btn-delete {
  background-color: #dc3545;
  color: white;
}

.form-row {
  display: flex;
  gap: 20px;
  margin-bottom: 15px;
}

.form-group {
  flex: 1;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
  font-weight: 500;
  color: #2c3e50;
}

.form-group input,
.form-group select {
  width: 100%;
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.btn-primary {
  background-color: #28a745;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
}

.btn-primary:hover {
  background-color: #218838;
}

.chart-container {
  height: 300px;
  margin-top: 15px;
}
</style>
