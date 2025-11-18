<template>
  <div class="external-apis">
    <div class="section-header">
      <h3>API های خارجی</h3>
      <p>مدیریت اتصالات و تنظیمات API های خارجی</p>
    </div>

    <!-- API Status Overview -->
    <div class="api-status-overview">
      <div class="status-card active">
        <div class="status-icon">
          <i class="fas fa-check-circle"></i>
        </div>
        <div class="status-content">
          <h4>API های فعال</h4>
          <div class="number">{{ activeAPIs }}</div>
          <div class="status-text">همه API ها در حال کار هستند</div>
        </div>
      </div>

      <div class="status-card warning">
        <div class="status-icon">
          <i class="fas fa-exclamation-triangle"></i>
        </div>
        <div class="status-content">
          <h4>API های مشکل‌دار</h4>
          <div class="number">{{ problematicAPIs }}</div>
          <div class="status-text">نیاز به بررسی دارند</div>
        </div>
      </div>

      <div class="status-card inactive">
        <div class="status-icon">
          <i class="fas fa-times-circle"></i>
        </div>
        <div class="status-content">
          <h4>API های غیرفعال</h4>
          <div class="number">{{ inactiveAPIs }}</div>
          <div class="status-text">غیرفعال یا متوقف شده</div>
        </div>
      </div>
    </div>

    <!-- API Configuration -->
    <div class="api-configuration">
      <h4>تنظیمات API</h4>

      <!-- Shipping APIs -->
      <div class="api-section">
        <h5>API های ارسال</h5>
        <div class="api-list">
          <div v-for="api in shippingAPIs" :key="api.id" class="api-item">
            <div class="api-header">
              <div class="api-info">
                <h6>{{ api.name }}</h6>
                <span class="api-provider">{{ api.provider }}</span>
              </div>
              <div class="api-status" :class="api.status">
                <i :class="getStatusIcon(api.status)"></i>
                <span>{{ getStatusText(api.status) }}</span>
              </div>
            </div>
            <div class="api-details">
              <div class="detail-item">
                <span class="label">Endpoint:</span>
                <span class="value">{{ api.endpoint }}</span>
              </div>
              <div class="detail-item">
                <span class="label">کلید API:</span>
                <span class="value">{{ maskAPIKey(api.apiKey) }}</span>
                <button @click="toggleAPIKey(api.id)" class="btn-toggle">
                  <i :class="api.showKey ? 'fas fa-eye-slash' : 'fas fa-eye'"></i>
                </button>
              </div>
              <div class="detail-item">
                <span class="label">آخرین درخواست:</span>
                <span class="value">{{ formatDate(api.lastRequest) }}</span>
              </div>
              <div class="detail-item">
                <span class="label">نرخ موفقیت:</span>
                <span class="value" :class="getSuccessRateClass(api.successRate)">
                  {{ api.successRate }}%
                </span>
              </div>
            </div>
            <div class="api-actions">
              <button @click="testAPI(api.id)" class="btn btn-primary">
                <i class="fas fa-play"></i>
                تست اتصال
              </button>
              <button @click="editAPI(api.id)" class="btn btn-secondary">
                <i class="fas fa-edit"></i>
                ویرایش
              </button>
              <button @click="toggleAPI(api.id)" class="btn" :class="api.status === 'active' ? 'btn-warning' : 'btn-success'">
                <i :class="api.status === 'active' ? 'fas fa-pause' : 'fas fa-play'"></i>
                {{ api.status === 'active' ? 'غیرفعال' : 'فعال' }}
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Tracking APIs -->
      <div class="api-section">
        <h5>API های ردیابی</h5>
        <div class="api-list">
          <div v-for="api in trackingAPIs" :key="api.id" class="api-item">
            <div class="api-header">
              <div class="api-info">
                <h6>{{ api.name }}</h6>
                <span class="api-provider">{{ api.provider }}</span>
              </div>
              <div class="api-status" :class="api.status">
                <i :class="getStatusIcon(api.status)"></i>
                <span>{{ getStatusText(api.status) }}</span>
              </div>
            </div>
            <div class="api-details">
              <div class="detail-item">
                <span class="label">Endpoint:</span>
                <span class="value">{{ api.endpoint }}</span>
              </div>
              <div class="detail-item">
                <span class="label">کلید API:</span>
                <span class="value">{{ maskAPIKey(api.apiKey) }}</span>
                <button @click="toggleAPIKey(api.id)" class="btn-toggle">
                  <i :class="api.showKey ? 'fas fa-eye-slash' : 'fas fa-eye'"></i>
                </button>
              </div>
              <div class="detail-item">
                <span class="label">آخرین درخواست:</span>
                <span class="value">{{ formatDate(api.lastRequest) }}</span>
              </div>
              <div class="detail-item">
                <span class="label">نرخ موفقیت:</span>
                <span class="value" :class="getSuccessRateClass(api.successRate)">
                  {{ api.successRate }}%
                </span>
              </div>
            </div>
            <div class="api-actions">
              <button @click="testAPI(api.id)" class="btn btn-primary">
                <i class="fas fa-play"></i>
                تست اتصال
              </button>
              <button @click="editAPI(api.id)" class="btn btn-secondary">
                <i class="fas fa-edit"></i>
                ویرایش
              </button>
              <button @click="toggleAPI(api.id)" class="btn" :class="api.status === 'active' ? 'btn-warning' : 'btn-success'">
                <i :class="api.status === 'active' ? 'fas fa-pause' : 'fas fa-play'"></i>
                {{ api.status === 'active' ? 'غیرفعال' : 'فعال' }}
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Add New API -->
      <div class="add-api-section">
        <h5>افزودن API جدید</h5>
        <div class="add-api-form">
          <div class="form-row">
            <div class="form-group">
              <label>نام API:</label>
              <input type="text" v-model="newAPI.name" placeholder="مثال: Post API">
            </div>
            <div class="form-group">
              <label>ارائه‌دهنده:</label>
              <input type="text" v-model="newAPI.provider" placeholder="مثال: شرکت پست">
            </div>
          </div>
          <div class="form-row">
            <div class="form-group">
              <label>نوع API:</label>
              <select v-model="newAPI.type">
                <option value="shipping">ارسال</option>
                <option value="tracking">ردیابی</option>
                <option value="pricing">قیمت‌گذاری</option>
              </select>
            </div>
            <div class="form-group">
              <label>Endpoint:</label>
              <input type="url" v-model="newAPI.endpoint" placeholder="https://api.example.com/v1">
            </div>
          </div>
          <div class="form-row">
            <div class="form-group">
              <label>کلید API:</label>
              <input type="password" v-model="newAPI.apiKey" placeholder="کلید API را وارد کنید">
            </div>
            <div class="form-group">
              <label>محدودیت درخواست:</label>
              <input type="number" v-model="newAPI.rateLimit" placeholder="1000">
            </div>
          </div>
          <div class="form-actions">
            <button @click="addAPI" class="btn btn-success">
              <i class="fas fa-plus"></i>
              افزودن API
            </button>
            <button @click="resetNewAPI" class="btn btn-secondary">
              <i class="fas fa-undo"></i>
              بازنشانی
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- API Logs -->
    <div class="api-logs">
      <h4>لاگ‌های API</h4>
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
          <label>API:</label>
          <select v-model="selectedAPI">
            <option value="all">همه API ها</option>
            <option v-for="api in allAPIs" :key="api.id" :value="api.id">
              {{ api.name }}
            </option>
          </select>
        </div>
        <button @click="refreshLogs" class="btn btn-primary">
          <i class="fas fa-sync"></i>
          بروزرسانی
        </button>
      </div>

      <div class="logs-table">
        <table>
          <thead>
          <tr>
            <th>زمان</th>
            <th>API</th>
            <th>نوع درخواست</th>
            <th>وضعیت</th>
            <th>زمان پاسخ</th>
            <th>جزئیات</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="log in filteredLogs" :key="log.id" :class="log.status">
            <td>{{ formatDateTime(log.timestamp) }}</td>
            <td>{{ log.apiName }}</td>
            <td>{{ log.requestType }}</td>
            <td>
                <span class="status-badge" :class="log.status">
                  {{ getLogStatusText(log.status) }}
                </span>
            </td>
            <td>{{ log.responseTime }}ms</td>
            <td>
              <button @click="viewLogDetails(log.id)" class="btn-link">
                مشاهده جزئیات
              </button>
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'

// Props
interface Props {
  shippingMethods?: any[]
}

const props = defineProps<Props>()

// Emits
const emit = defineEmits<{
  apiUpdated: [apiId: string, data: any]
  apiAdded: [data: any]
  apiTested: [apiId: string, result: any]
}>()

// Reactive data
const activeAPIs = ref(0)
const problematicAPIs = ref(0)
const inactiveAPIs = ref(0)

const shippingAPIs = ref([
  {
    id: 'shipping-1',
    name: 'Post API',
    provider: 'شرکت پست ایران',
    status: 'active',
    endpoint: 'https://api.post.ir/v1/shipping',
    apiKey: 'sk_test_1234567890abcdef',
    showKey: false,
    lastRequest: new Date('2024-01-15T10:30:00'),
    successRate: 98.5
  },
  {
    id: 'shipping-2',
    name: 'TNT API',
    provider: 'TNT Express',
    status: 'active',
    endpoint: 'https://api.tnt.com/v2/shipping',
    apiKey: 'tnt_key_abcdef123456',
    showKey: false,
    lastRequest: new Date('2024-01-15T09:15:00'),
    successRate: 95.2
  },
  {
    id: 'shipping-3',
    name: 'DHL API',
    provider: 'DHL Express',
    status: 'warning',
    endpoint: 'https://api.dhl.com/v1/shipping',
    apiKey: 'dhl_key_789xyz',
    showKey: false,
    lastRequest: new Date('2024-01-15T08:45:00'),
    successRate: 87.3
  }
])

const trackingAPIs = ref([
  {
    id: 'tracking-1',
    name: 'Post Tracking',
    provider: 'شرکت پست ایران',
    status: 'active',
    endpoint: 'https://api.post.ir/v1/tracking',
    apiKey: 'sk_test_1234567890abcdef',
    showKey: false,
    lastRequest: new Date('2024-01-15T10:25:00'),
    successRate: 99.1
  },
  {
    id: 'tracking-2',
    name: 'Global Tracking',
    provider: 'Global Express',
    status: 'inactive',
    endpoint: 'https://api.global.com/v1/tracking',
    apiKey: 'global_key_456def',
    showKey: false,
    lastRequest: new Date('2024-01-14T16:20:00'),
    successRate: 0
  }
])

const newAPI = reactive({
  name: '',
  provider: '',
  type: 'shipping',
  endpoint: '',
  apiKey: '',
  rateLimit: 1000
})

const logFilter = ref('all')
const selectedAPI = ref('all')

const apiLogs = ref([
  {
    id: 'log-1',
    timestamp: new Date('2024-01-15T10:30:00'),
    apiName: 'Post API',
    requestType: 'POST /shipping',
    status: 'success',
    responseTime: 245
  },
  {
    id: 'log-2',
    timestamp: new Date('2024-01-15T10:28:00'),
    apiName: 'DHL API',
    requestType: 'GET /tracking',
    status: 'error',
    responseTime: 5000
  },
  {
    id: 'log-3',
    timestamp: new Date('2024-01-15T10:25:00'),
    apiName: 'TNT API',
    requestType: 'POST /shipping',
    status: 'success',
    responseTime: 189
  }
])

// Computed
const allAPIs = computed(() => [...shippingAPIs.value, ...trackingAPIs.value])

const filteredLogs = computed(() => {
  let logs = apiLogs.value

  if (logFilter.value !== 'all') {
    logs = logs.filter(log => log.status === logFilter.value)
  }

  if (selectedAPI.value !== 'all') {
    logs = logs.filter(log => {
      const api = allAPIs.value.find(a => a.id === selectedAPI.value)
      return api && log.apiName === api.name
    })
  }

  return logs
})

// Methods
const getStatusIcon = (status: string): string => {
  switch (status) {
    case 'active': return 'fas fa-check-circle'
    case 'warning': return 'fas fa-exclamation-triangle'
    case 'inactive': return 'fas fa-times-circle'
    default: return 'fas fa-question-circle'
  }
}

const getStatusText = (status: string): string => {
  switch (status) {
    case 'active': return 'فعال'
    case 'warning': return 'هشدار'
    case 'inactive': return 'غیرفعال'
    default: return 'نامشخص'
  }
}

const maskAPIKey = (key: string): string => {
  if (key.length <= 8) return '*'.repeat(key.length)
  return key.substring(0, 4) + '*'.repeat(key.length - 8) + key.substring(key.length - 4)
}

const toggleAPIKey = (apiId: string) => {
  const api = [...shippingAPIs.value, ...trackingAPIs.value].find(a => a.id === apiId)
  if (api) {
    api.showKey = !api.showKey
  }
}

const testAPI = async (apiId: string) => {
  try {
    // Simulate API test
    await new Promise(resolve => setTimeout(resolve, 2000))

    const result = {
      success: Math.random() > 0.2,
      responseTime: Math.floor(Math.random() * 1000) + 100,
      message: Math.random() > 0.2 ? 'اتصال موفق' : 'خطا در اتصال'
    }

    emit('apiTested', apiId, result)
  } catch (error) {
    console.error('Error testing API:', error)
  }
}

const editAPI = (apiId: string) => {
  // Implementation for editing API
  console.log('Edit API:', apiId)
}

const toggleAPI = (apiId: string) => {
  const api = [...shippingAPIs.value, ...trackingAPIs.value].find(a => a.id === apiId)
  if (api) {
    api.status = api.status === 'active' ? 'inactive' : 'active'
    emit('apiUpdated', apiId, api)
  }
}

const addAPI = () => {
  if (!newAPI.name || !newAPI.provider || !newAPI.endpoint || !newAPI.apiKey) {
    alert('لطفاً تمام فیلدها را پر کنید')
    return
  }

  const apiData = {
    id: `api-${Date.now()}`,
    ...newAPI,
    status: 'inactive',
    showKey: false,
    lastRequest: new Date(),
    successRate: 0
  }

  if (newAPI.type === 'shipping') {
    shippingAPIs.value.push(apiData)
  } else if (newAPI.type === 'tracking') {
    trackingAPIs.value.push(apiData)
  }

  emit('apiAdded', apiData)
  resetNewAPI()
}

const resetNewAPI = () => {
  newAPI.name = ''
  newAPI.provider = ''
  newAPI.type = 'shipping'
  newAPI.endpoint = ''
  newAPI.apiKey = ''
  newAPI.rateLimit = 1000
}

const getSuccessRateClass = (rate: number): string => {
  if (rate >= 95) return 'success'
  if (rate >= 80) return 'warning'
  return 'error'
}

const getLogStatusText = (status: string): string => {
  switch (status) {
    case 'success': return 'موفق'
    case 'error': return 'خطا'
    case 'warning': return 'هشدار'
    default: return 'نامشخص'
  }
}

const formatDate = (date: Date): string => {
  return new Intl.DateTimeFormat('fa-IR').format(date)
}

const formatDateTime = (date: Date): string => {
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date)
}

const viewLogDetails = (logId: string) => {
  // Implementation for viewing log details
  console.log('View log details:', logId)
}

const refreshLogs = () => {
  // Implementation for refreshing logs
  console.log('Refreshing logs...')
}

// Lifecycle
onMounted(() => {
  // Calculate API counts
  activeAPIs.value = allAPIs.value.filter(api => api.status === 'active').length
  problematicAPIs.value = allAPIs.value.filter(api => api.status === 'warning').length
  inactiveAPIs.value = allAPIs.value.filter(api => api.status === 'inactive').length
})
</script>

<style scoped>
.external-apis {
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

.api-status-overview {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
  margin-bottom: 30px;
}

.status-card {
  padding: 20px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  gap: 15px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
  color: white;
}

.status-card.active {
  background: linear-gradient(135deg, #27ae60 0%, #2ecc71 100%);
}

.status-card.warning {
  background: linear-gradient(135deg, #f39c12 0%, #e67e22 100%);
}

.status-card.inactive {
  background: linear-gradient(135deg, #e74c3c 0%, #c0392b 100%);
}

.status-icon {
  font-size: 24px;
  opacity: 0.8;
}

.status-content h4 {
  margin: 0 0 5px 0;
  font-size: 14px;
  opacity: 0.9;
}

.number {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 5px;
}

.status-text {
  font-size: 12px;
  opacity: 0.8;
}

.api-configuration {
  margin-bottom: 30px;
}

.api-configuration h4 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.api-section {
  margin-bottom: 30px;
}

.api-section h5 {
  margin-bottom: 15px;
  color: #2c3e50;
  font-size: 18px;
}

.api-list {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.api-item {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  border: 1px solid #e9ecef;
}

.api-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

.api-info h6 {
  margin: 0 0 5px 0;
  color: #2c3e50;
  font-size: 16px;
}

.api-provider {
  font-size: 12px;
  color: #6c757d;
}

.api-status {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
}

.api-status.active {
  background: rgba(39, 174, 96, 0.1);
  color: #27ae60;
}

.api-status.warning {
  background: rgba(243, 156, 18, 0.1);
  color: #f39c12;
}

.api-status.inactive {
  background: rgba(231, 76, 60, 0.1);
  color: #e74c3c;
}

.api-details {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 15px;
  margin-bottom: 15px;
}

.detail-item {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 14px;
}

.detail-item .label {
  color: #6c757d;
  min-width: 80px;
}

.detail-item .value {
  font-weight: 600;
  color: #2c3e50;
  flex: 1;
}

.btn-toggle {
  background: none;
  border: none;
  color: #3498db;
  cursor: pointer;
  padding: 2px;
}

.btn-toggle:hover {
  color: #2980b9;
}

.api-actions {
  display: flex;
  gap: 10px;
  flex-wrap: wrap;
}

.btn {
  padding: 8px 16px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  display: flex;
  align-items: center;
  gap: 6px;
  transition: all 0.3s ease;
}

.btn-primary {
  background: #3498db;
  color: white;
}

.btn-primary:hover {
  background: #2980b9;
}

.btn-secondary {
  background: #6c757d;
  color: white;
}

.btn-secondary:hover {
  background: #5a6268;
}

.btn-success {
  background: #27ae60;
  color: white;
}

.btn-success:hover {
  background: #229954;
}

.btn-warning {
  background: #f39c12;
  color: white;
}

.btn-warning:hover {
  background: #e67e22;
}

.add-api-section {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  border: 2px dashed #dee2e6;
}

.add-api-section h5 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.add-api-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-row {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 15px;
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

.form-group input,
.form-group select {
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.form-actions {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
}

.api-logs {
  margin-bottom: 30px;
}

.api-logs h4 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.logs-controls {
  display: flex;
  gap: 20px;
  margin-bottom: 20px;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 8px;
  flex-wrap: wrap;
}

.logs-table {
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

.status-badge {
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 600;
}

.status-badge.success {
  background: rgba(39, 174, 96, 0.1);
  color: #27ae60;
}

.status-badge.error {
  background: rgba(231, 76, 60, 0.1);
  color: #e74c3c;
}

.status-badge.warning {
  background: rgba(243, 156, 18, 0.1);
  color: #f39c12;
}

.btn-link {
  background: none;
  border: none;
  color: #3498db;
  cursor: pointer;
  text-decoration: underline;
}

.btn-link:hover {
  color: #2980b9;
}

.success {
  color: #27ae60;
}

.warning {
  color: #f39c12;
}

.error {
  color: #e74c3c;
}

@media (max-width: 768px) {
  .api-status-overview {
    grid-template-columns: 1fr;
  }

  .api-details {
    grid-template-columns: 1fr;
  }

  .api-actions {
    flex-direction: column;
  }

  .form-row {
    grid-template-columns: 1fr;
  }

  .logs-controls {
    flex-direction: column;
  }

  .form-actions {
    justify-content: stretch;
  }
}
</style>
