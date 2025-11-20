<template>
  <div class="webhooks">
    <div class="section-header">
      <h3>وب‌هوک‌ها</h3>
      <p>مدیریت وب‌هوک‌ها برای اعلان‌های خودکار</p>
    </div>

    <!-- Webhook Status Overview -->
    <div class="webhook-status-overview">
      <div class="status-card active">
        <div class="status-icon">
          <i class="fas fa-check-circle"></i>
        </div>
        <div class="status-content">
          <h4>وب‌هوک‌های فعال</h4>
          <div class="number">{{ activeWebhooks }}</div>
          <div class="status-text">در حال ارسال اعلان</div>
        </div>
      </div>

      <div class="status-card failed">
        <div class="status-icon">
          <i class="fas fa-exclamation-triangle"></i>
        </div>
        <div class="status-content">
          <h4>وب‌هوک‌های ناموفق</h4>
          <div class="number">{{ failedWebhooks }}</div>
          <div class="status-text">نیاز به بررسی</div>
        </div>
      </div>

      <div class="status-card pending">
        <div class="status-icon">
          <i class="fas fa-clock"></i>
        </div>
        <div class="status-content">
          <h4>در انتظار ارسال</h4>
          <div class="number">{{ pendingWebhooks }}</div>
          <div class="status-text">در صف ارسال</div>
        </div>
      </div>
    </div>

    <!-- Webhook Configuration -->
    <div class="webhook-configuration">
      <h4>تنظیمات وب‌هوک</h4>

      <!-- Webhook List -->
      <div class="webhook-list">
        <div v-for="webhook in webhooks" :key="webhook.id" class="webhook-item">
          <div class="webhook-header">
            <div class="webhook-info">
              <h5>{{ webhook.name }}</h5>
              <span class="webhook-url">{{ webhook.url }}</span>
            </div>
            <div class="webhook-status" :class="webhook.status">
              <i :class="getStatusIcon(webhook.status)"></i>
              <span>{{ getStatusText(webhook.status) }}</span>
            </div>
          </div>

          <div class="webhook-details">
            <div class="detail-row">
              <div class="detail-item">
                <span class="label">رویدادها:</span>
                <div class="events-list">
                  <span v-for="event in webhook.events" :key="event" class="event-tag">
                    {{ getEventText(event) }}
                  </span>
                </div>
              </div>
            </div>

            <div class="detail-row">
              <div class="detail-item">
                <span class="label">آخرین ارسال:</span>
                <span class="value">{{ formatDate(webhook.lastSent) }}</span>
              </div>
              <div class="detail-item">
                <span class="label">نرخ موفقیت:</span>
                <span class="value" :class="getSuccessRateClass(webhook.successRate)">
                  {{ webhook.successRate }}%
                </span>
              </div>
              <div class="detail-item">
                <span class="label">تعداد ارسال:</span>
                <span class="value">{{ formatNumber(webhook.totalSent) }}</span>
              </div>
            </div>

            <div class="detail-row">
              <div class="detail-item">
                <span class="label">کلید امنیتی:</span>
                <span class="value">{{ maskSecret(webhook.secret) }}</span>
                <button class="btn-toggle" @click="toggleSecret(webhook.id)">
                  <i :class="webhook.showSecret ? 'fas fa-eye-slash' : 'fas fa-eye'"></i>
                </button>
              </div>
              <div class="detail-item">
                <span class="label">تایم‌اوت:</span>
                <span class="value">{{ webhook.timeout }} ثانیه</span>
              </div>
            </div>
          </div>

          <div class="webhook-actions">
            <button class="btn btn-primary" @click="testWebhook(webhook.id)">
              <i class="fas fa-play"></i>
              تست وب‌هوک
            </button>
            <button class="btn btn-secondary" @click="editWebhook(webhook.id)">
              <i class="fas fa-edit"></i>
              ویرایش
            </button>
            <button class="btn" :class="webhook.status === 'active' ? 'btn-warning' : 'btn-success'" @click="toggleWebhook(webhook.id)">
              <i :class="webhook.status === 'active' ? 'fas fa-pause' : 'fas fa-play'"></i>
              {{ webhook.status === 'active' ? 'غیرفعال' : 'فعال' }}
            </button>
            <button class="btn btn-info" @click="viewLogs(webhook.id)">
              <i class="fas fa-list"></i>
              لاگ‌ها
            </button>
          </div>
        </div>
      </div>

      <!-- Add New Webhook -->
      <div class="add-webhook-section">
        <h5>افزودن وب‌هوک جدید</h5>
        <div class="add-webhook-form">
          <div class="form-row">
            <div class="form-group">
              <label>نام وب‌هوک:</label>
              <input v-model="newWebhook.name" type="text" placeholder="مثال: اعلان سفارش">
            </div>
            <div class="form-group">
              <label>URL:</label>
              <input v-model="newWebhook.url" type="url" placeholder="https://your-domain.com/webhook">
            </div>
          </div>

          <div class="form-row">
            <div class="form-group">
              <label>رویدادها:</label>
              <div class="events-selection">
                <label v-for="event in availableEvents" :key="event.value" class="checkbox-item">
                  <input v-model="newWebhook.events" type="checkbox" :value="event.value">
                  <span>{{ event.label }}</span>
                </label>
              </div>
            </div>
            <div class="form-group">
              <label>کلید امنیتی:</label>
              <input v-model="newWebhook.secret" type="password" placeholder="کلید امنیتی را وارد کنید">
            </div>
          </div>

          <div class="form-row">
            <div class="form-group">
              <label>تایم‌اوت (ثانیه):</label>
              <input v-model="newWebhook.timeout" type="number" min="5" max="60" placeholder="30">
            </div>
            <div class="form-group">
              <label>حداکثر تلاش:</label>
              <input v-model="newWebhook.maxRetries" type="number" min="1" max="10" placeholder="3">
            </div>
          </div>

          <div class="form-actions">
            <button class="btn btn-success" @click="addWebhook">
              <i class="fas fa-plus"></i>
              افزودن وب‌هوک
            </button>
            <button class="btn btn-secondary" @click="resetNewWebhook">
              <i class="fas fa-undo"></i>
              بازنشانی
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Webhook Logs -->
    <div class="webhook-logs">
      <h4>لاگ‌های وب‌هوک</h4>
      <div class="logs-controls">
        <div class="form-group">
          <label>فیلتر بر اساس:</label>
          <select v-model="logFilter">
            <option value="all">همه</option>
            <option value="success">موفق</option>
            <option value="failed">ناموفق</option>
            <option value="pending">در انتظار</option>
          </select>
        </div>
        <div class="form-group">
          <label>وب‌هوک:</label>
          <select v-model="selectedWebhook">
            <option value="all">همه وب‌هوک‌ها</option>
            <option v-for="webhook in webhooks" :key="webhook.id" :value="webhook.id">
              {{ webhook.name }}
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
            <th>وب‌هوک</th>
            <th>رویداد</th>
            <th>وضعیت</th>
            <th>زمان پاسخ</th>
            <th>کد پاسخ</th>
            <th>جزئیات</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="log in filteredLogs" :key="log.id" :class="log.status">
            <td>{{ formatDateTime(log.timestamp) }}</td>
            <td>{{ log.webhookName }}</td>
            <td>{{ getEventText(log.event) }}</td>
            <td>
                <span class="status-badge" :class="log.status">
                  {{ getLogStatusText(log.status) }}
                </span>
            </td>
            <td>{{ log.responseTime }}ms</td>
            <td>{{ log.statusCode }}</td>
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
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue';

// Props
interface Props {
  shippingMethods?: unknown[]
}

defineProps<Props>()

// Emits
const emit = defineEmits<{
  webhookUpdated: [webhookId: string, data: unknown]
  webhookAdded: [data: unknown]
  webhookTested: [webhookId: string, result: unknown]
}>()

// Reactive data
const activeWebhooks = ref(0)
const failedWebhooks = ref(0)
const pendingWebhooks = ref(0)

const webhooks = ref([
  {
    id: 'webhook-1',
    name: 'اعلان سفارش',
    url: 'https://api.example.com/webhooks/orders',
    status: 'active',
    events: ['order.created', 'order.shipped', 'order.delivered'],
    lastSent: new Date('2024-01-15T10:30:00'),
    successRate: 98.5,
    totalSent: 1250,
    secret: 'whsec_1234567890abcdef',
    showSecret: false,
    timeout: 30,
    maxRetries: 3
  },
  {
    id: 'webhook-2',
    name: 'اعلان ردیابی',
    url: 'https://tracking.example.com/webhook',
    status: 'active',
    events: ['tracking.updated'],
    lastSent: new Date('2024-01-15T09:15:00'),
    successRate: 95.2,
    totalSent: 890,
    secret: 'whsec_abcdef123456',
    showSecret: false,
    timeout: 25,
    maxRetries: 3
  },
  {
    id: 'webhook-3',
    name: 'اعلان خطا',
    url: 'https://error.example.com/webhook',
    status: 'failed',
    events: ['error.occurred'],
    lastSent: new Date('2024-01-15T08:45:00'),
    successRate: 45.8,
    totalSent: 120,
    secret: 'whsec_789xyz',
    showSecret: false,
    timeout: 20,
    maxRetries: 5
  }
])

const availableEvents = [
  { value: 'order.created', label: 'ایجاد سفارش' },
  { value: 'order.shipped', label: 'ارسال سفارش' },
  { value: 'order.delivered', label: 'تحویل سفارش' },
  { value: 'order.cancelled', label: 'لغو سفارش' },
  { value: 'tracking.updated', label: 'بروزرسانی ردیابی' },
  { value: 'error.occurred', label: 'رخداد خطا' },
  { value: 'payment.completed', label: 'تکمیل پرداخت' },
  { value: 'refund.processed', label: 'پردازش بازپرداخت' }
]

const newWebhook = reactive({
  name: '',
  url: '',
  events: [],
  secret: '',
  timeout: 30,
  maxRetries: 3
})

const logFilter = ref('all')
const selectedWebhook = ref('all')

const webhookLogs = ref([
  {
    id: 'log-1',
    timestamp: new Date('2024-01-15T10:30:00'),
    webhookName: 'اعلان سفارش',
    event: 'order.created',
    status: 'success',
    responseTime: 245,
    statusCode: 200
  },
  {
    id: 'log-2',
    timestamp: new Date('2024-01-15T10:28:00'),
    webhookName: 'اعلان خطا',
    event: 'error.occurred',
    status: 'failed',
    responseTime: 5000,
    statusCode: 500
  },
  {
    id: 'log-3',
    timestamp: new Date('2024-01-15T10:25:00'),
    webhookName: 'اعلان ردیابی',
    event: 'tracking.updated',
    status: 'success',
    responseTime: 189,
    statusCode: 200
  }
])

// Computed
const filteredLogs = computed(() => {
  let logs = webhookLogs.value

  if (logFilter.value !== 'all') {
    logs = logs.filter(log => log.status === logFilter.value)
  }

  if (selectedWebhook.value !== 'all') {
    logs = logs.filter(log => {
      const webhook = webhooks.value.find(w => w.id === selectedWebhook.value)
      return webhook && log.webhookName === webhook.name
    })
  }

  return logs
})

// Methods
const getStatusIcon = (status: string): string => {
  switch (status) {
    case 'active': return 'fas fa-check-circle'
    case 'failed': return 'fas fa-exclamation-triangle'
    case 'pending': return 'fas fa-clock'
    case 'inactive': return 'fas fa-times-circle'
    default: return 'fas fa-question-circle'
  }
}

const getStatusText = (status: string): string => {
  switch (status) {
    case 'active': return 'فعال'
    case 'failed': return 'ناموفق'
    case 'pending': return 'در انتظار'
    case 'inactive': return 'غیرفعال'
    default: return 'نامشخص'
  }
}

const getEventText = (event: string): string => {
  const eventObj = availableEvents.find(e => e.value === event)
  return eventObj ? eventObj.label : event
}

const maskSecret = (secret: string): string => {
  if (secret.length <= 8) return '*'.repeat(secret.length)
  return secret.substring(0, 4) + '*'.repeat(secret.length - 8) + secret.substring(secret.length - 4)
}

const toggleSecret = (webhookId: string) => {
  const webhook = webhooks.value.find(w => w.id === webhookId)
  if (webhook) {
    webhook.showSecret = !webhook.showSecret
  }
}

const testWebhook = async (webhookId: string) => {
  try {
    // Simulate webhook test
    await new Promise(resolve => setTimeout(resolve, 2000))

    const result = {
      success: Math.random() > 0.2,
      responseTime: Math.floor(Math.random() * 1000) + 100,
      statusCode: Math.random() > 0.2 ? 200 : 500,
      message: Math.random() > 0.2 ? 'ارسال موفق' : 'خطا در ارسال'
    }

    emit('webhookTested', webhookId, result)
  } catch (error) {
    console.error('Error testing webhook:', error)
  }
}

const editWebhook = (_webhookId: string) => {
  // Implementation for editing webhook
  // console.log('Edit webhook:', webhookId)
}

const toggleWebhook = (webhookId: string) => {
  const webhook = webhooks.value.find(w => w.id === webhookId)
  if (webhook) {
    webhook.status = webhook.status === 'active' ? 'inactive' : 'active'
    emit('webhookUpdated', webhookId, webhook)
  }
}

const viewLogs = (_webhookId: string) => {
  // Implementation for viewing webhook logs
  // console.log('View webhook logs:', webhookId)
}

const addWebhook = () => {
  if (!newWebhook.name || !newWebhook.url || newWebhook.events.length === 0 || !newWebhook.secret) {
    alert('لطفاً تمام فیلدهای ضروری را پر کنید')
    return
  }

  const webhookData = {
    id: `webhook-${Date.now()}`,
    ...newWebhook,
    status: 'inactive',
    showSecret: false,
    lastSent: new Date(),
    successRate: 0,
    totalSent: 0
  }

  webhooks.value.push(webhookData)
  emit('webhookAdded', webhookData)
  resetNewWebhook()
}

const resetNewWebhook = () => {
  newWebhook.name = ''
  newWebhook.url = ''
  newWebhook.events = []
  newWebhook.secret = ''
  newWebhook.timeout = 30
  newWebhook.maxRetries = 3
}

const getSuccessRateClass = (rate: number): string => {
  if (rate >= 95) return 'success'
  if (rate >= 80) return 'warning'
  return 'error'
}

const getLogStatusText = (status: string): string => {
  switch (status) {
    case 'success': return 'موفق'
    case 'failed': return 'ناموفق'
    case 'pending': return 'در انتظار'
    default: return 'نامشخص'
  }
}

const formatNumber = (num: number): string => {
  return new Intl.NumberFormat('fa-IR').format(num)
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

const viewLogDetails = (_logId: string) => {
  // Implementation for viewing log details
  // console.log('View log details:', logId)
}

const refreshLogs = () => {
  // Implementation for refreshing logs
  // console.log('Refreshing logs...')
}

// Lifecycle
onMounted(() => {
  // Calculate webhook counts
  activeWebhooks.value = webhooks.value.filter(w => w.status === 'active').length
  failedWebhooks.value = webhooks.value.filter(w => w.status === 'failed').length
  pendingWebhooks.value = webhooks.value.filter(w => w.status === 'pending').length
})
</script>

<style scoped>
.webhooks {
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

.webhook-status-overview {
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

.status-card.failed {
  background: linear-gradient(135deg, #e74c3c 0%, #c0392b 100%);
}

.status-card.pending {
  background: linear-gradient(135deg, #f39c12 0%, #e67e22 100%);
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

.webhook-configuration {
  margin-bottom: 30px;
}

.webhook-configuration h4 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.webhook-list {
  display: flex;
  flex-direction: column;
  gap: 15px;
  margin-bottom: 30px;
}

.webhook-item {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  border: 1px solid #e9ecef;
}

.webhook-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

.webhook-info h5 {
  margin: 0 0 5px 0;
  color: #2c3e50;
  font-size: 16px;
}

.webhook-url {
  font-size: 12px;
  color: #6c757d;
  font-family: monospace;
}

.webhook-status {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
}

.webhook-status.active {
  background: rgba(39, 174, 96, 0.1);
  color: #27ae60;
}

.webhook-status.failed {
  background: rgba(231, 76, 60, 0.1);
  color: #e74c3c;
}

.webhook-status.pending {
  background: rgba(243, 156, 18, 0.1);
  color: #f39c12;
}

.webhook-status.inactive {
  background: rgba(108, 117, 125, 0.1);
  color: #6c757d;
}

.webhook-details {
  margin-bottom: 15px;
}

.detail-row {
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

.events-list {
  display: flex;
  flex-wrap: wrap;
  gap: 5px;
}

.event-tag {
  background: #e9ecef;
  color: #495057;
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 12px;
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

.webhook-actions {
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

.btn-info {
  background: #17a2b8;
  color: white;
}

.btn-info:hover {
  background: #138496;
}

.add-webhook-section {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  border: 2px dashed #dee2e6;
}

.add-webhook-section h5 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.add-webhook-form {
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

.events-selection {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 10px;
}

.checkbox-item {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  cursor: pointer;
}

.checkbox-item input[type="checkbox"] {
  margin: 0;
}

.form-actions {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
}

.webhook-logs {
  margin-bottom: 30px;
}

.webhook-logs h4 {
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

.status-badge.failed {
  background: rgba(231, 76, 60, 0.1);
  color: #e74c3c;
}

.status-badge.pending {
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
  .webhook-status-overview {
    grid-template-columns: 1fr;
  }

  .detail-row {
    grid-template-columns: 1fr;
  }

  .webhook-actions {
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

  .events-selection {
    grid-template-columns: 1fr;
  }
}
</style>
