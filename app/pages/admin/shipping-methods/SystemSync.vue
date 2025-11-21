<template>
  <div class="external-system-sync">
    <div class="section-header">
      <h3>همگام‌سازی با سیستم‌های خارجی</h3>
      <p>مدیریت اتصال و همگام‌سازی با سیستم‌های خارجی</p>
    </div>

    <!-- Sync Status Overview -->
    <div class="sync-status-overview">
      <div class="status-card connected">
        <div class="status-icon">
          <i class="fas fa-link"></i>
        </div>
        <div class="status-content">
          <h4>سیستم‌های متصل</h4>
          <div class="number">{{ connectedSystems }}</div>
          <div class="status-text">در حال همگام‌سازی</div>
        </div>
      </div>

      <div class="status-card syncing">
        <div class="status-icon">
          <i class="fas fa-sync-alt"></i>
        </div>
        <div class="status-content">
          <h4>در حال همگام‌سازی</h4>
          <div class="number">{{ syncingSystems }}</div>
          <div class="status-text">در حال انتقال داده</div>
        </div>
      </div>

      <div class="status-card error">
        <div class="status-icon">
          <i class="fas fa-exclamation-triangle"></i>
        </div>
        <div class="status-content">
          <h4>خطا در اتصال</h4>
          <div class="number">{{ errorSystems }}</div>
          <div class="status-text">نیاز به بررسی</div>
        </div>
      </div>

      <div class="status-card pending">
        <div class="status-icon">
          <i class="fas fa-clock"></i>
        </div>
        <div class="status-content">
          <h4>در انتظار</h4>
          <div class="number">{{ pendingSystems }}</div>
          <div class="status-text">در صف همگام‌سازی</div>
        </div>
      </div>
    </div>

    <!-- External Systems Configuration -->
    <div class="external-systems">
      <h4>تنظیمات سیستم‌های خارجی</h4>

      <div class="systems-grid">
        <div v-for="system in externalSystems" :key="system.id" class="system-card">
          <div class="system-header">
            <div class="system-info">
              <div class="system-logo">
                <img v-if="system.logo" :src="system.logo" :alt="system.name">
                <i v-else :class="system.icon"></i>
              </div>
              <div class="system-details">
                <h5>{{ system.name }}</h5>
                <span class="system-type">{{ system.type }}</span>
              </div>
            </div>
            <div class="system-status" :class="system.status">
              <i :class="getStatusIcon(system.status)"></i>
              <span>{{ getStatusText(system.status) }}</span>
            </div>
          </div>

          <div class="system-connection">
            <div class="connection-info">
              <div class="info-item">
                <span class="label">آخرین همگام‌سازی:</span>
                <span class="value">{{ formatDateTime(system.lastSync) }}</span>
              </div>
              <div class="info-item">
                <span class="label">نرخ موفقیت:</span>
                <span class="value" :class="getSuccessRateClass(system.successRate)">
                  {{ system.successRate }}%
                </span>
              </div>
              <div class="info-item">
                <span class="label">تعداد رکوردها:</span>
                <span class="value">{{ formatNumber(system.recordCount) }}</span>
              </div>
              <div class="info-item">
                <span class="label">اندازه داده:</span>
                <span class="value">{{ formatDataSize(system.dataSize) }}</span>
              </div>
            </div>

            <div v-if="system.status === 'syncing'" class="sync-progress">
              <div class="progress-bar">
                <div class="progress-fill" :style="{ width: system.syncProgress + '%' }"></div>
              </div>
              <span class="progress-text">{{ system.syncProgress }}%</span>
            </div>
          </div>

          <div class="system-settings">
            <div class="settings-row">
              <div class="setting-item">
                <span class="label">نوع همگام‌سازی:</span>
                <span class="value">{{ getSyncTypeText(system.syncType) }}</span>
              </div>
              <div class="setting-item">
                <span class="label">فرکانس:</span>
                <span class="value">{{ getFrequencyText(system.frequency) }}</span>
              </div>
            </div>

            <div class="settings-row">
              <div class="setting-item">
                <span class="label">فیلدهای همگام:</span>
                <div class="sync-fields">
                  <span v-for="field in system.syncFields" :key="field" class="field-tag">
                    {{ field }}
                  </span>
                </div>
              </div>
            </div>
          </div>

          <div class="system-actions">
            <button class="btn btn-primary" @click="testConnection(system.id)">
              <i class="fas fa-plug"></i>
              تست اتصال
            </button>
            <button class="btn btn-success" :disabled="system.status === 'syncing'" @click="startSync(system.id)">
              <i class="fas fa-sync-alt"></i>
              شروع همگام‌سازی
            </button>
            <button class="btn btn-secondary" @click="editSystem(system.id)">
              <i class="fas fa-edit"></i>
              ویرایش
            </button>
            <button class="btn btn-info" @click="viewLogs(system.id)">
              <i class="fas fa-list"></i>
              لاگ‌ها
            </button>
            <button class="btn" :class="system.status === 'connected' ? 'btn-warning' : 'btn-success'" @click="toggleSystem(system.id)">
              <i :class="system.status === 'connected' ? 'fas fa-pause' : 'fas fa-play'"></i>
              {{ system.status === 'connected' ? 'قطع اتصال' : 'اتصال' }}
            </button>
          </div>
        </div>
      </div>

      <!-- Add New System -->
      <div class="add-system-section">
        <h5>افزودن سیستم جدید</h5>
        <div class="add-system-form">
          <div class="form-row">
            <div class="form-group">
              <label>نام سیستم:</label>
              <input v-model="newSystem.name" type="text" placeholder="مثال: سیستم انبارداری">
            </div>
            <div class="form-group">
              <label>نوع سیستم:</label>
              <select v-model="newSystem.type">
                <option value="inventory">انبارداری</option>
                <option value="accounting">حسابداری</option>
                <option value="crm">مدیریت مشتری</option>
                <option value="erp">برنامه‌ریزی منابع</option>
                <option value="custom">سفارشی</option>
              </select>
            </div>
          </div>

          <div class="form-row">
            <div class="form-group">
              <label>URL اتصال:</label>
              <input v-model="newSystem.connectionUrl" type="url" placeholder="https://api.example.com">
            </div>
            <div class="form-group">
              <label>کلید API:</label>
              <input v-model="newSystem.apiKey" type="password" placeholder="کلید API را وارد کنید">
            </div>
          </div>

          <div class="form-row">
            <div class="form-group">
              <label>نوع همگام‌سازی:</label>
              <select v-model="newSystem.syncType">
                <option value="full">کامل</option>
                <option value="incremental">افزایشی</option>
                <option value="real_time">زمان واقعی</option>
              </select>
            </div>
            <div class="form-group">
              <label>فرکانس همگام‌سازی:</label>
              <select v-model="newSystem.frequency">
                <option value="manual">دستی</option>
                <option value="hourly">ساعتی</option>
                <option value="daily">روزانه</option>
                <option value="weekly">هفتگی</option>
              </select>
            </div>
          </div>

          <div class="form-group">
            <label>فیلدهای همگام:</label>
            <div class="fields-selection">
              <label v-for="field in availableFields" :key="field.value" class="checkbox-item">
                <input v-model="newSystem.syncFields" type="checkbox" :value="field.value">
                <span>{{ field.label }}</span>
              </label>
            </div>
          </div>

          <div class="form-actions">
            <button class="btn btn-success" @click="addSystem">
              <i class="fas fa-plus"></i>
              افزودن سیستم
            </button>
            <button class="btn btn-secondary" @click="resetNewSystem">
              <i class="fas fa-undo"></i>
              بازنشانی
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Sync Logs -->
    <div class="sync-logs">
      <h4>لاگ‌های همگام‌سازی</h4>
      <div class="logs-controls">
        <div class="form-group">
          <label>فیلتر بر اساس:</label>
          <select v-model="logFilter">
            <option value="all">همه</option>
            <option value="success">موفق</option>
            <option value="error">خطا</option>
            <option value="warning">هشدار</option>
          </select>
        </div>
        <div class="form-group">
          <label>سیستم:</label>
          <select v-model="selectedSystem">
            <option value="all">همه سیستم‌ها</option>
            <option v-for="system in externalSystems" :key="system.id" :value="system.id">
              {{ system.name }}
            </option>
          </select>
        </div>
        <button class="btn btn-primary" @click="refreshLogs">
          <i class="fas fa-sync"></i>
          بروزرسانی
        </button>
      </div>

      <div class="logs-table">
        <table>
          <thead>
          <tr>
            <th>زمان</th>
            <th>سیستم</th>
            <th>عملیات</th>
            <th>وضعیت</th>
            <th>رکوردها</th>
            <th>زمان اجرا</th>
            <th>جزئیات</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="log in filteredLogs" :key="log.id" :class="log.status">
            <td>{{ formatDateTime(log.timestamp) }}</td>
            <td>{{ log.systemName }}</td>
            <td>{{ getOperationText(log.operation) }}</td>
            <td>
                <span class="status-badge" :class="log.status">
                  {{ getLogStatusText(log.status) }}
                </span>
            </td>
            <td>{{ formatNumber(log.records) }}</td>
            <td>{{ log.duration }}ms</td>
            <td>
              <button class="btn-link" @click="viewLogDetails(log.id)">
                مشاهده جزئیات
              </button>
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Data Mapping -->
    <div class="data-mapping">
      <h4>نقشه‌برداری داده</h4>

      <div class="mapping-tables">
        <div v-for="mapping in dataMappings" :key="mapping.id" class="mapping-table">
          <h5>{{ mapping.name }}</h5>
          <div class="mapping-content">
            <table>
              <thead>
              <tr>
                <th>فیلد منبع</th>
                <th>نوع داده</th>
                <th>فیلد مقصد</th>
                <th>تبدیل</th>
                <th>وضعیت</th>
              </tr>
              </thead>
              <tbody>
              <tr v-for="field in mapping.fields" :key="field.id">
                <td>{{ field.sourceField }}</td>
                <td>{{ field.dataType }}</td>
                <td>{{ field.targetField }}</td>
                <td>{{ field.transformation || '-' }}</td>
                <td>
                    <span class="status-badge" :class="field.status">
                      {{ getMappingStatusText(field.status) }}
                    </span>
                </td>
              </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, reactive, ref } from 'vue'

// Reactive data
const connectedSystems = ref(3)
const syncingSystems = ref(1)
const errorSystems = ref(0)
const pendingSystems = ref(2)

const logFilter = ref('all')
const selectedSystem = ref('all')

const newSystem = reactive({
  name: '',
  type: 'inventory',
  connectionUrl: '',
  apiKey: '',
  syncType: 'full',
  frequency: 'daily',
  syncFields: []
})

const externalSystems = ref([
  {
    id: 1,
    name: 'سیستم انبارداری',
    type: 'انبارداری',
    status: 'connected',
    lastSync: new Date(),
    successRate: 98,
    recordCount: 15420,
    dataSize: '2.5 MB',
    syncProgress: 0,
    syncType: 'incremental',
    frequency: 'hourly',
    syncFields: ['products', 'inventory', 'orders'],
    logo: null,
    icon: 'fas fa-boxes'
  },
  {
    id: 2,
    name: 'سیستم حسابداری',
    type: 'حسابداری',
    status: 'syncing',
    lastSync: new Date(Date.now() - 3600000),
    successRate: 95,
    recordCount: 8920,
    dataSize: '1.8 MB',
    syncProgress: 65,
    syncType: 'full',
    frequency: 'daily',
    syncFields: ['invoices', 'payments', 'customers'],
    logo: null,
    icon: 'fas fa-calculator'
  }
])

const availableFields = ref([
  { value: 'products', label: 'محصولات' },
  { value: 'inventory', label: 'موجودی' },
  { value: 'orders', label: 'سفارشات' },
  { value: 'customers', label: 'مشتریان' },
  { value: 'invoices', label: 'فاکتورها' },
  { value: 'payments', label: 'پرداخت‌ها' }
])

const syncLogs = ref([
  {
    id: 1,
    timestamp: new Date(),
    systemName: 'سیستم انبارداری',
    operation: 'sync_products',
    status: 'success',
    records: 1250,
    duration: 4500
  }
])

const dataMappings = ref([
  {
    id: 1,
    name: 'نقشه‌برداری محصولات',
    fields: [
      {
        id: 1,
        sourceField: 'product_id',
        dataType: 'string',
        targetField: 'sku',
        transformation: null,
        status: 'mapped'
      }
    ]
  }
])

// Computed properties
const filteredLogs = computed(() => {
  let logs = syncLogs.value
  
  if (logFilter.value !== 'all') {
    logs = logs.filter(log => log.status === logFilter.value)
  }
  
  if (selectedSystem.value !== 'all') {
    logs = logs.filter(log => log.systemName === selectedSystem.value)
  }
  
  return logs
})

// Methods
function getStatusIcon(status) {
  const icons = {
    connected: 'fas fa-check-circle',
    syncing: 'fas fa-sync-alt',
    error: 'fas fa-exclamation-triangle',
    pending: 'fas fa-clock'
  }
  return icons[status] || 'fas fa-question-circle'
}

function getStatusText(status) {
  const texts = {
    connected: 'متصل',
    syncing: 'در حال همگام‌سازی',
    error: 'خطا',
    pending: 'در انتظار'
  }
  return texts[status] || 'نامشخص'
}

function getSuccessRateClass(rate) {
  if (rate >= 95) return 'success'
  if (rate >= 80) return 'warning'
  return 'error'
}

function getSyncTypeText(type) {
  const texts = {
    full: 'کامل',
    incremental: 'افزایشی',
    real_time: 'زمان واقعی'
  }
  return texts[type] || type
}

function getFrequencyText(frequency) {
  const texts = {
    manual: 'دستی',
    hourly: 'ساعتی',
    daily: 'روزانه',
    weekly: 'هفتگی'
  }
  return texts[frequency] || frequency
}

function getOperationText(operation) {
  const texts = {
    sync_products: 'همگام‌سازی محصولات',
    sync_orders: 'همگام‌سازی سفارشات',
    sync_customers: 'همگام‌سازی مشتریان'
  }
  return texts[operation] || operation
}

function getLogStatusText(status) {
  const texts = {
    success: 'موفق',
    error: 'خطا',
    warning: 'هشدار'
  }
  return texts[status] || status
}

function getMappingStatusText(status) {
  const texts = {
    mapped: 'نقشه‌برداری شده',
    pending: 'در انتظار',
    error: 'خطا'
  }
  return texts[status] || status
}

function formatDateTime(date) {
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date)
}

function formatNumber(num) {
  return new Intl.NumberFormat('fa-IR').format(num)
}

function formatDataSize(bytes) {
  const sizes = ['B', 'KB', 'MB', 'GB']
  if (bytes === 0) return '0 B'
  const i = Math.floor(Math.log(bytes) / Math.log(1024))
  return Math.round(bytes / Math.pow(1024, i) * 100) / 100 + ' ' + sizes[i]
}

function testConnection(_systemId) {
}

function startSync(_systemId) {
}

function editSystem(_systemId) {
}

function viewLogs(_systemId) {
}

function toggleSystem(_systemId) {
}

function addSystem() {
}

function resetNewSystem() {
  Object.assign(newSystem, {
    name: '',
    type: 'inventory',
    connectionUrl: '',
    apiKey: '',
    syncType: 'full',
    frequency: 'daily',
    syncFields: []
  })
}

function refreshLogs() {
}

function viewLogDetails(_logId) {
}

onMounted(() => {
})
</script>

<style scoped>
.external-system-sync {
  direction: rtl;
  padding: 1rem;
}

.section-header {
  margin-bottom: 2rem;
}

.section-header h3 {
  margin: 0 0 0.5rem 0;
  font-size: 1.5rem;
  color: #333;
}

.section-header p {
  margin: 0;
  color: #666;
}

.sync-status-overview {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
  margin-bottom: 2rem;
}

.status-card {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  display: flex;
  align-items: center;
  gap: 1rem;
}

.status-card.connected {
  border-left: 4px solid #28a745;
}

.status-card.syncing {
  border-left: 4px solid #007bff;
}

.status-card.error {
  border-left: 4px solid #dc3545;
}

.status-card.pending {
  border-left: 4px solid #ffc107;
}

.status-icon {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.5rem;
}

.status-card.connected .status-icon {
  background: #d4edda;
  color: #155724;
}

.status-card.syncing .status-icon {
  background: #cce7ff;
  color: #004085;
}

.status-card.error .status-icon {
  background: #f8d7da;
  color: #721c24;
}

.status-card.pending .status-icon {
  background: #fff3cd;
  color: #856404;
}

.status-content h4 {
  margin: 0 0 0.5rem 0;
  font-size: 1rem;
  color: #333;
}

.status-content .number {
  font-size: 2rem;
  font-weight: bold;
  color: #333;
  margin-bottom: 0.25rem;
}

.status-content .status-text {
  font-size: 0.9rem;
  color: #666;
}

.external-systems {
  margin-bottom: 2rem;
}

.external-systems h4 {
  margin: 0 0 1rem 0;
  font-size: 1.2rem;
  color: #333;
}

.systems-grid {
  display: grid;
  gap: 1rem;
}

.system-card {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.system-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.system-info {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.system-logo {
  width: 40px;
  height: 40px;
  border-radius: 8px;
  background: #f8f9fa;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.2rem;
}

.system-details h5 {
  margin: 0 0 0.25rem 0;
  font-size: 1.1rem;
  color: #333;
}

.system-type {
  font-size: 0.9rem;
  color: #666;
}

.system-status {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  border-radius: 20px;
  font-size: 0.9rem;
}

.system-status.connected {
  background: #d4edda;
  color: #155724;
}

.system-status.syncing {
  background: #cce7ff;
  color: #004085;
}

.system-status.error {
  background: #f8d7da;
  color: #721c24;
}

.system-status.pending {
  background: #fff3cd;
  color: #856404;
}

.system-connection {
  margin-bottom: 1rem;
}

.connection-info {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
  margin-bottom: 1rem;
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.info-item .label {
  font-size: 0.9rem;
  color: #666;
}

.info-item .value {
  font-weight: 500;
  color: #333;
}

.info-item .value.success {
  color: #28a745;
}

.info-item .value.warning {
  color: #ffc107;
}

.info-item .value.error {
  color: #dc3545;
}

.sync-progress {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.progress-bar {
  flex: 1;
  height: 8px;
  background: #e9ecef;
  border-radius: 4px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: #007bff;
  transition: width 0.3s ease;
}

.progress-text {
  font-size: 0.9rem;
  color: #666;
  min-width: 40px;
}

.system-settings {
  margin-bottom: 1rem;
}

.settings-row {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
  margin-bottom: 1rem;
}

.setting-item {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.setting-item .label {
  font-size: 0.9rem;
  color: #666;
}

.setting-item .value {
  font-weight: 500;
  color: #333;
}

.sync-fields {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.field-tag {
  background: #e9ecef;
  color: #495057;
  padding: 0.25rem 0.5rem;
  border-radius: 12px;
  font-size: 0.8rem;
}

.system-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.btn {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.9rem;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  transition: all 0.2s;
}

.btn-primary {
  background: #007bff;
  color: white;
}

.btn-success {
  background: #28a745;
  color: white;
}

.btn-secondary {
  background: #6c757d;
  color: white;
}

.btn-info {
  background: #17a2b8;
  color: white;
}

.btn-warning {
  background: #ffc107;
  color: #212529;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.add-system-section {
  background: #f8f9fa;
  border-radius: 12px;
  padding: 1.5rem;
  margin-top: 2rem;
}

.add-system-section h5 {
  margin: 0 0 1rem 0;
  font-size: 1.1rem;
  color: #333;
}

.form-row {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
  margin-bottom: 1rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.form-group label {
  font-size: 0.9rem;
  color: #555;
  font-weight: 500;
}

.form-group input,
.form-group select {
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 8px;
  font-size: 0.9rem;
}

.fields-selection {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 0.5rem;
}

.checkbox-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  cursor: pointer;
}

.form-actions {
  display: flex;
  gap: 1rem;
  margin-top: 1rem;
}

.sync-logs {
  margin-bottom: 2rem;
}

.sync-logs h4 {
  margin: 0 0 1rem 0;
  font-size: 1.2rem;
  color: #333;
}

.logs-controls {
  display: flex;
  gap: 1rem;
  margin-bottom: 1rem;
  flex-wrap: wrap;
}

.logs-table {
  background: white;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.logs-table table {
  width: 100%;
  border-collapse: collapse;
}

.logs-table th,
.logs-table td {
  padding: 1rem;
  text-align: right;
  border-bottom: 1px solid #e9ecef;
}

.logs-table th {
  background: #f8f9fa;
  font-weight: 600;
  color: #555;
}

.status-badge {
  padding: 0.25rem 0.5rem;
  border-radius: 12px;
  font-size: 0.8rem;
}

.status-badge.success {
  background: #d4edda;
  color: #155724;
}

.status-badge.error {
  background: #f8d7da;
  color: #721c24;
}

.status-badge.warning {
  background: #fff3cd;
  color: #856404;
}

.btn-link {
  background: none;
  border: none;
  color: #007bff;
  cursor: pointer;
  text-decoration: underline;
}

.data-mapping {
  margin-bottom: 2rem;
}

.data-mapping h4 {
  margin: 0 0 1rem 0;
  font-size: 1.2rem;
  color: #333;
}

.mapping-tables {
  display: grid;
  gap: 1rem;
}

.mapping-table {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.mapping-table h5 {
  margin: 0 0 1rem 0;
  font-size: 1.1rem;
  color: #333;
}

.mapping-content table {
  width: 100%;
  border-collapse: collapse;
}

.mapping-content th,
.mapping-content td {
  padding: 0.75rem;
  text-align: right;
  border-bottom: 1px solid #e9ecef;
}

.mapping-content th {
  background: #f8f9fa;
  font-weight: 600;
  color: #555;
}
</style>
