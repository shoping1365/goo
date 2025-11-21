<template>
  <div class="automatic-notifications">
    <div class="section-header">
      <h3>اعلان‌های خودکار</h3>
      <p>مدیریت اعلان‌های خودکار برای مشتریان و مدیران</p>
    </div>

    <!-- Notification Status Overview -->
    <div class="notification-status-overview">
      <div class="status-card active">
        <div class="status-icon">
          <i class="fas fa-bell"></i>
        </div>
        <div class="status-content">
          <h4>اعلان‌های فعال</h4>
          <div class="number">{{ activeNotifications }}</div>
          <div class="status-text">در حال ارسال</div>
        </div>
      </div>

      <div class="status-card sent">
        <div class="status-icon">
          <i class="fas fa-paper-plane"></i>
        </div>
        <div class="status-content">
          <h4>ارسال شده امروز</h4>
          <div class="number">{{ todaySent }}</div>
          <div class="status-text">اعلان ارسال شده</div>
        </div>
      </div>

      <div class="status-card failed">
        <div class="status-icon">
          <i class="fas fa-exclamation-triangle"></i>
        </div>
        <div class="status-content">
          <h4>ناموفق</h4>
          <div class="number">{{ failedNotifications }}</div>
          <div class="status-text">نیاز به بررسی</div>
        </div>
      </div>

      <div class="status-card scheduled">
        <div class="status-icon">
          <i class="fas fa-clock"></i>
        </div>
        <div class="status-content">
          <h4>برنامه‌ریزی شده</h4>
          <div class="number">{{ scheduledNotifications }}</div>
          <div class="status-text">در انتظار ارسال</div>
        </div>
      </div>
    </div>

    <!-- Notification Templates -->
    <div class="notification-templates">
      <h4>قالب‌های اعلان</h4>

      <div class="templates-grid">
        <div v-for="template in notificationTemplates" :key="template.id" class="template-card">
          <div class="template-header">
            <div class="template-info">
              <h5>{{ template.name }}</h5>
              <span class="template-type">{{ getTypeText(template.type) }}</span>
            </div>
            <div class="template-status" :class="template.status">
              <i :class="getStatusIcon(template.status)"></i>
              <span>{{ getStatusText(template.status) }}</span>
            </div>
          </div>

          <div class="template-content">
            <div class="template-preview">
              <h6>پیش‌نمایش:</h6>
              <div class="preview-content">
                <div class="preview-subject">
                  <strong>موضوع:</strong> {{ template.subject }}
                </div>
                <div class="preview-body">
                  {{ template.body }}
                </div>
              </div>
            </div>

            <div class="template-settings">
              <div class="setting-item">
                <span class="label">کانال‌های ارسال:</span>
                <div class="channels">
                  <span v-for="channel in template.channels" :key="channel" class="channel-tag">
                    {{ getChannelText(channel) }}
                  </span>
                </div>
              </div>

              <div class="setting-item">
                <span class="label">زمان ارسال:</span>
                <span class="value">{{ getTimingText(template.timing) }}</span>
              </div>

              <div class="setting-item">
                <span class="label">آخرین ارسال:</span>
                <span class="value">{{ formatDate(template.lastSent) }}</span>
              </div>

              <div class="setting-item">
                <span class="label">تعداد ارسال:</span>
                <span class="value">{{ formatNumber(template.totalSent) }}</span>
              </div>
            </div>
          </div>

          <div class="template-actions">
            <button class="btn btn-primary" @click="editTemplate(template.id)">
              <i class="fas fa-edit"></i>
              ویرایش
            </button>
            <button class="btn btn-secondary" @click="testTemplate(template.id)">
              <i class="fas fa-play"></i>
              تست
            </button>
            <button class="btn" :class="template.status === 'active' ? 'btn-warning' : 'btn-success'" @click="toggleTemplate(template.id)">
              <i :class="template.status === 'active' ? 'fas fa-pause' : 'fas fa-play'"></i>
              {{ template.status === 'active' ? 'غیرفعال' : 'فعال' }}
            </button>
            <button class="btn btn-info" @click="viewHistory(template.id)">
              <i class="fas fa-history"></i>
              تاریخچه
            </button>
          </div>
        </div>
      </div>

      <!-- Add New Template -->
      <div class="add-template-section">
        <h5>افزودن قالب جدید</h5>
        <div class="add-template-form">
          <div class="form-row">
            <div class="form-group">
              <label>نام قالب:</label>
              <input v-model="newTemplate.name" type="text" placeholder="مثال: اعلان ارسال سفارش">
            </div>
            <div class="form-group">
              <label>نوع اعلان:</label>
              <select v-model="newTemplate.type">
                <option value="order_shipped">ارسال سفارش</option>
                <option value="order_delivered">تحویل سفارش</option>
                <option value="tracking_update">بروزرسانی ردیابی</option>
                <option value="delivery_delay">تاخیر در تحویل</option>
                <option value="payment_reminder">یادآوری پرداخت</option>
                <option value="custom">سفارشی</option>
              </select>
            </div>
          </div>

          <div class="form-row">
            <div class="form-group">
              <label>موضوع:</label>
              <input v-model="newTemplate.subject" type="text" placeholder="موضوع اعلان">
            </div>
            <div class="form-group">
              <label>کانال‌های ارسال:</label>
              <div class="channels-selection">
                <label v-for="channel in availableChannels" :key="channel.value" class="checkbox-item">
                  <input v-model="newTemplate.channels" type="checkbox" :value="channel.value">
                  <span>{{ channel.label }}</span>
                </label>
              </div>
            </div>
          </div>

          <div class="form-group">
            <label>متن اعلان:</label>
            <textarea v-model="newTemplate.body" rows="6" placeholder="متن اعلان را وارد کنید..."></textarea>
            <div class="template-variables">
              <span class="variables-title">متغیرهای قابل استفاده:</span>
              <div class="variables-list">
                <span v-for="variable in templateVariables" :key="variable.key" class="variable-tag" @click="insertVariable(variable.key)">
                  {{ variable.key }} - {{ variable.description }}
                </span>
              </div>
            </div>
          </div>

          <div class="form-row">
            <div class="form-group">
              <label>زمان ارسال:</label>
              <select v-model="newTemplate.timing">
                <option value="immediate">فوری</option>
                <option value="scheduled">برنامه‌ریزی شده</option>
                <option value="delayed">تاخیری</option>
              </select>
            </div>
            <div class="form-group">
              <label>اولویت:</label>
              <select v-model="newTemplate.priority">
                <option value="low">کم</option>
                <option value="medium">متوسط</option>
                <option value="high">بالا</option>
                <option value="urgent">فوری</option>
              </select>
            </div>
          </div>

          <div class="form-actions">
            <button class="btn btn-success" @click="addTemplate">
              <i class="fas fa-plus"></i>
              افزودن قالب
            </button>
            <button class="btn btn-secondary" @click="resetNewTemplate">
              <i class="fas fa-undo"></i>
              بازنشانی
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Notification Rules -->
    <div class="notification-rules">
      <h4>قوانین اعلان</h4>

      <div class="rules-list">
        <div v-for="rule in notificationRules" :key="rule.id" class="rule-item">
          <div class="rule-header">
            <div class="rule-info">
              <h5>{{ rule.name }}</h5>
              <span class="rule-description">{{ rule.description }}</span>
            </div>
            <div class="rule-status" :class="rule.status">
              <i :class="getStatusIcon(rule.status)"></i>
              <span>{{ getStatusText(rule.status) }}</span>
            </div>
          </div>

          <div class="rule-conditions">
            <h6>شرایط:</h6>
            <div class="conditions-list">
              <div v-for="condition in rule.conditions" :key="condition.id" class="condition-item">
                <span class="condition-field">{{ condition.field }}</span>
                <span class="condition-operator">{{ condition.operator }}</span>
                <span class="condition-value">{{ condition.value }}</span>
              </div>
            </div>
          </div>

          <div class="rule-actions">
            <div class="action-item">
              <span class="label">اقدام:</span>
              <span class="value">{{ rule.action }}</span>
            </div>
            <div class="action-item">
              <span class="label">قالب:</span>
              <span class="value">{{ rule.template }}</span>
            </div>
            <div class="action-item">
              <span class="label">مخاطبان:</span>
              <span class="value">{{ rule.audience }}</span>
            </div>
          </div>

          <div class="rule-controls">
            <button class="btn btn-primary" @click="editRule(rule.id)">
              <i class="fas fa-edit"></i>
              ویرایش
            </button>
            <button class="btn" :class="rule.status === 'active' ? 'btn-warning' : 'btn-success'" @click="toggleRule(rule.id)">
              <i :class="rule.status === 'active' ? 'fas fa-pause' : 'fas fa-play'"></i>
              {{ rule.status === 'active' ? 'غیرفعال' : 'فعال' }}
            </button>
            <button class="btn btn-danger" @click="deleteRule(rule.id)">
              <i class="fas fa-trash"></i>
              حذف
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Notification History -->
    <div class="notification-history">
      <h4>تاریخچه اعلان‌ها</h4>
      <div class="history-controls">
        <div class="form-group">
          <label>فیلتر بر اساس:</label>
          <select v-model="historyFilter">
            <option value="all">همه</option>
            <option value="sent">ارسال شده</option>
            <option value="failed">ناموفق</option>
            <option value="pending">در انتظار</option>
          </select>
        </div>
        <div class="form-group">
          <label>قالب:</label>
          <select v-model="selectedTemplate">
            <option value="all">همه قالب‌ها</option>
            <option v-for="template in notificationTemplates" :key="template.id" :value="template.id">
              {{ template.name }}
            </option>
          </select>
        </div>
        <button class="btn btn-primary" @click="refreshHistory">
          <i class="fas fa-sync"></i>
          بروزرسانی
        </button>
      </div>

      <div class="history-table">
        <table>
          <thead>
          <tr>
            <th>زمان</th>
            <th>قالب</th>
            <th>مخاطب</th>
            <th>کانال</th>
            <th>وضعیت</th>
            <th>جزئیات</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="notification in filteredHistory" :key="notification.id" :class="notification.status">
            <td>{{ formatDateTime(notification.timestamp) }}</td>
            <td>{{ notification.template }}</td>
            <td>{{ notification.recipient }}</td>
            <td>{{ getChannelText(notification.channel) }}</td>
            <td>
                <span class="status-badge" :class="notification.status">
                  {{ getNotificationStatusText(notification.status) }}
                </span>
            </td>
            <td>
              <button class="btn-link" @click="viewNotificationDetails(notification.id)">
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
  templateUpdated: [templateId: string, data: unknown]
  templateAdded: [data: unknown]
  templateTested: [templateId: string, result: unknown]
  ruleUpdated: [ruleId: string, data: unknown]
}>()

// Reactive data
const activeNotifications = ref(0)
const todaySent = ref(0)
const failedNotifications = ref(0)
const scheduledNotifications = ref(0)

const notificationTemplates = ref([
  {
    id: 'template-1',
    name: 'اعلان ارسال سفارش',
    type: 'order_shipped',
    status: 'active',
    subject: 'سفارش شما ارسال شد',
    body: 'سفارش شما با شماره {order_number} در تاریخ {ship_date} ارسال شد. شماره ردیابی: {tracking_number}',
    channels: ['email', 'sms'],
    timing: 'immediate',
    lastSent: new Date('2024-01-15T10:30:00'),
    totalSent: 1250,
    priority: 'medium'
  },
  {
    id: 'template-2',
    name: 'اعلان تحویل سفارش',
    type: 'order_delivered',
    status: 'active',
    subject: 'سفارش شما تحویل داده شد',
    body: 'سفارش شما با شماره {order_number} در تاریخ {delivery_date} تحویل داده شد.',
    channels: ['email', 'sms', 'push'],
    timing: 'immediate',
    lastSent: new Date('2024-01-15T09:15:00'),
    totalSent: 890,
    priority: 'high'
  },
  {
    id: 'template-3',
    name: 'یادآوری پرداخت',
    type: 'payment_reminder',
    status: 'inactive',
    subject: 'یادآوری پرداخت سفارش',
    body: 'سفارش شما با شماره {order_number} در انتظار پرداخت است. مبلغ: {amount}',
    channels: ['email'],
    timing: 'delayed',
    lastSent: new Date('2024-01-14T16:20:00'),
    totalSent: 320,
    priority: 'medium'
  }
])

const availableChannels = [
  { value: 'email', label: 'ایمیل' },
  { value: 'sms', label: 'پیامک' },
  { value: 'push', label: 'اعلان مرورگر' },
  { value: 'in_app', label: 'اعلان درون برنامه' }
]

const templateVariables = [
  { key: '{customer_name}', description: 'نام مشتری' },
  { key: '{order_number}', description: 'شماره سفارش' },
  { key: '{tracking_number}', description: 'شماره ردیابی' },
  { key: '{ship_date}', description: 'تاریخ ارسال' },
  { key: '{delivery_date}', description: 'تاریخ تحویل' },
  { key: '{amount}', description: 'مبلغ' },
  { key: '{shipping_method}', description: 'روش ارسال' }
]

const newTemplate = reactive({
  name: '',
  type: 'order_shipped',
  subject: '',
  body: '',
  channels: [],
  timing: 'immediate',
  priority: 'medium'
})

const notificationRules = ref([
  {
    id: 'rule-1',
    name: 'اعلان ارسال خودکار',
    description: 'ارسال اعلان هنگام ارسال سفارش',
    status: 'active',
    conditions: [
      { id: 1, field: 'order_status', operator: 'equals', value: 'shipped' }
    ],
    action: 'send_notification',
    template: 'اعلان ارسال سفارش',
    audience: 'customer'
  },
  {
    id: 'rule-2',
    name: 'یادآوری پرداخت',
    description: 'ارسال یادآوری برای سفارشات پرداخت نشده',
    status: 'active',
    conditions: [
      { id: 1, field: 'payment_status', operator: 'equals', value: 'pending' },
      { id: 2, field: 'order_age', operator: 'greater_than', value: '24 hours' }
    ],
    action: 'send_reminder',
    template: 'یادآوری پرداخت',
    audience: 'customer'
  }
])

const historyFilter = ref('all')
const selectedTemplate = ref('all')

const notificationHistory = ref([
  {
    id: 'notif-1',
    timestamp: new Date('2024-01-15T10:30:00'),
    template: 'اعلان ارسال سفارش',
    recipient: 'علی احمدی',
    channel: 'email',
    status: 'sent'
  },
  {
    id: 'notif-2',
    timestamp: new Date('2024-01-15T10:28:00'),
    template: 'اعلان ارسال سفارش',
    recipient: 'فاطمه محمدی',
    channel: 'sms',
    status: 'failed'
  },
  {
    id: 'notif-3',
    timestamp: new Date('2024-01-15T10:25:00'),
    template: 'اعلان تحویل سفارش',
    recipient: 'محمد رضایی',
    channel: 'push',
    status: 'sent'
  }
])

// Computed
const filteredHistory = computed(() => {
  let history = notificationHistory.value

  if (historyFilter.value !== 'all') {
    history = history.filter(notif => notif.status === historyFilter.value)
  }

  if (selectedTemplate.value !== 'all') {
    history = history.filter(notif => {
      const template = notificationTemplates.value.find(t => t.id === selectedTemplate.value)
      return template && notif.template === template.name
    })
  }

  return history
})

// Methods
const getStatusIcon = (status: string): string => {
  switch (status) {
    case 'active': return 'fas fa-check-circle'
    case 'inactive': return 'fas fa-times-circle'
    case 'pending': return 'fas fa-clock'
    default: return 'fas fa-question-circle'
  }
}

const getStatusText = (status: string): string => {
  switch (status) {
    case 'active': return 'فعال'
    case 'inactive': return 'غیرفعال'
    case 'pending': return 'در انتظار'
    default: return 'نامشخص'
  }
}

const getTypeText = (type: string): string => {
  switch (type) {
    case 'order_shipped': return 'ارسال سفارش'
    case 'order_delivered': return 'تحویل سفارش'
    case 'tracking_update': return 'بروزرسانی ردیابی'
    case 'delivery_delay': return 'تاخیر در تحویل'
    case 'payment_reminder': return 'یادآوری پرداخت'
    case 'custom': return 'سفارشی'
    default: return 'نامشخص'
  }
}

const getChannelText = (channel: string): string => {
  const channelObj = availableChannels.find(c => c.value === channel)
  return channelObj ? channelObj.label : channel
}

const getTimingText = (timing: string): string => {
  switch (timing) {
    case 'immediate': return 'فوری'
    case 'scheduled': return 'برنامه‌ریزی شده'
    case 'delayed': return 'تاخیری'
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

const insertVariable = (variable: string) => {
  newTemplate.body += variable
}

const editTemplate = (_templateId: string) => {
  // Implementation for editing template
}

const testTemplate = async (templateId: string) => {
  try {
    // Simulate template test
    await new Promise(resolve => setTimeout(resolve, 2000))

    const result = {
      success: Math.random() > 0.2,
      message: Math.random() > 0.2 ? 'ارسال موفق' : 'خطا در ارسال'
    }

    emit('templateTested', templateId, result)
  } catch (error) {
    console.error('Error testing template:', error)
  }
}

const toggleTemplate = (templateId: string) => {
  const template = notificationTemplates.value.find(t => t.id === templateId)
  if (template) {
    template.status = template.status === 'active' ? 'inactive' : 'active'
    emit('templateUpdated', templateId, template)
  }
}

const viewHistory = (_templateId: string) => {
  // Implementation for viewing template history
}

const addTemplate = () => {
  if (!newTemplate.name || !newTemplate.subject || !newTemplate.body || newTemplate.channels.length === 0) {
    alert('لطفاً تمام فیلدهای ضروری را پر کنید')
    return
  }

  const templateData = {
    id: `template-${Date.now()}`,
    ...newTemplate,
    status: 'inactive',
    lastSent: new Date(),
    totalSent: 0
  }

  notificationTemplates.value.push(templateData)
  emit('templateAdded', templateData)
  resetNewTemplate()
}

const resetNewTemplate = () => {
  newTemplate.name = ''
  newTemplate.type = 'order_shipped'
  newTemplate.subject = ''
  newTemplate.body = ''
  newTemplate.channels = []
  newTemplate.timing = 'immediate'
  newTemplate.priority = 'medium'
}

const editRule = (_ruleId: string) => {
  // Implementation for editing rule
}

const toggleRule = (ruleId: string) => {
  const rule = notificationRules.value.find(r => r.id === ruleId)
  if (rule) {
    rule.status = rule.status === 'active' ? 'inactive' : 'active'
    emit('ruleUpdated', ruleId, rule)
  }
}

const deleteRule = (ruleId: string) => {
  if (confirm('آیا از حذف این قانون اطمینان دارید؟')) {
    notificationRules.value = notificationRules.value.filter(r => r.id !== ruleId)
  }
}

const getNotificationStatusText = (status: string): string => {
  switch (status) {
    case 'sent': return 'ارسال شده'
    case 'failed': return 'ناموفق'
    case 'pending': return 'در انتظار'
    default: return 'نامشخص'
  }
}

const viewNotificationDetails = (_notificationId: string) => {
  // Implementation for viewing notification details
}

const refreshHistory = () => {
  // Implementation for refreshing history
}

// Lifecycle
onMounted(() => {
  // Calculate notification counts
  activeNotifications.value = notificationTemplates.value.filter(t => t.status === 'active').length
  todaySent.value = Math.floor(Math.random() * 500) + 200
  failedNotifications.value = Math.floor(Math.random() * 50) + 10
  scheduledNotifications.value = Math.floor(Math.random() * 100) + 50
})
</script>

<style scoped>
.automatic-notifications {
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

.notification-status-overview {
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

.status-card.sent {
  background: linear-gradient(135deg, #3498db 0%, #2980b9 100%);
}

.status-card.failed {
  background: linear-gradient(135deg, #e74c3c 0%, #c0392b 100%);
}

.status-card.scheduled {
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

.notification-templates {
  margin-bottom: 30px;
}

.notification-templates h4 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.templates-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  gap: 20px;
  margin-bottom: 30px;
}

.template-card {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  border: 1px solid #e9ecef;
}

.template-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

.template-info h5 {
  margin: 0 0 5px 0;
  color: #2c3e50;
  font-size: 16px;
}

.template-type {
  font-size: 12px;
  color: #6c757d;
  background: #e9ecef;
  padding: 2px 8px;
  border-radius: 12px;
}

.template-status {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
}

.template-status.active {
  background: rgba(39, 174, 96, 0.1);
  color: #27ae60;
}

.template-status.inactive {
  background: rgba(108, 117, 125, 0.1);
  color: #6c757d;
}

.template-content {
  margin-bottom: 15px;
}

.template-preview {
  margin-bottom: 15px;
}

.template-preview h6 {
  margin: 0 0 10px 0;
  color: #2c3e50;
  font-size: 14px;
}

.preview-content {
  background: white;
  padding: 15px;
  border-radius: 6px;
  border: 1px solid #dee2e6;
}

.preview-subject {
  margin-bottom: 10px;
  font-size: 14px;
}

.preview-body {
  font-size: 14px;
  line-height: 1.6;
  color: #495057;
}

.template-settings {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.setting-item {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 14px;
}

.setting-item .label {
  color: #6c757d;
  min-width: 100px;
}

.setting-item .value {
  font-weight: 600;
  color: #2c3e50;
  flex: 1;
}

.channels {
  display: flex;
  flex-wrap: wrap;
  gap: 5px;
}

.channel-tag {
  background: #e9ecef;
  color: #495057;
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 12px;
}

.template-actions {
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

.btn-danger {
  background: #dc3545;
  color: white;
}

.btn-danger:hover {
  background: #c82333;
}

.add-template-section {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  border: 2px dashed #dee2e6;
}

.add-template-section h5 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.add-template-form {
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
.form-group select,
.form-group textarea {
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.form-group textarea {
  resize: vertical;
  min-height: 100px;
}

.channels-selection {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
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

.template-variables {
  margin-top: 10px;
  padding: 10px;
  background: #e9ecef;
  border-radius: 4px;
}

.variables-title {
  font-weight: 600;
  color: #2c3e50;
  font-size: 12px;
  margin-bottom: 8px;
  display: block;
}

.variables-list {
  display: flex;
  flex-wrap: wrap;
  gap: 5px;
}

.variable-tag {
  background: #fff;
  color: #495057;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 11px;
  cursor: pointer;
  border: 1px solid #dee2e6;
  transition: all 0.2s ease;
}

.variable-tag:hover {
  background: #f8f9fa;
  border-color: #adb5bd;
}

.form-actions {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
}

.notification-rules {
  margin-bottom: 30px;
}

.notification-rules h4 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.rules-list {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.rule-item {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  border: 1px solid #e9ecef;
}

.rule-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

.rule-info h5 {
  margin: 0 0 5px 0;
  color: #2c3e50;
  font-size: 16px;
}

.rule-description {
  font-size: 12px;
  color: #6c757d;
}

.rule-status {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
}

.rule-status.active {
  background: rgba(39, 174, 96, 0.1);
  color: #27ae60;
}

.rule-status.inactive {
  background: rgba(108, 117, 125, 0.1);
  color: #6c757d;
}

.rule-conditions {
  margin-bottom: 15px;
}

.rule-conditions h6 {
  margin: 0 0 10px 0;
  color: #2c3e50;
  font-size: 14px;
}

.conditions-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.condition-item {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 14px;
  padding: 8px;
  background: white;
  border-radius: 4px;
  border: 1px solid #dee2e6;
}

.condition-field {
  font-weight: 600;
  color: #2c3e50;
}

.condition-operator {
  color: #6c757d;
  font-style: italic;
}

.condition-value {
  color: #495057;
  background: #e9ecef;
  padding: 2px 6px;
  border-radius: 4px;
}

.rule-actions {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 10px;
  margin-bottom: 15px;
}

.action-item {
  display: flex;
  justify-content: space-between;
  font-size: 14px;
}

.action-item .label {
  color: #6c757d;
}

.action-item .value {
  font-weight: 600;
  color: #2c3e50;
}

.rule-controls {
  display: flex;
  gap: 10px;
  flex-wrap: wrap;
}

.notification-history {
  margin-bottom: 30px;
}

.notification-history h4 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.history-controls {
  display: flex;
  gap: 20px;
  margin-bottom: 20px;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 8px;
  flex-wrap: wrap;
}

.history-table {
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

.status-badge.sent {
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

@media (max-width: 768px) {
  .notification-status-overview {
    grid-template-columns: 1fr;
  }

  .templates-grid {
    grid-template-columns: 1fr;
  }

  .template-actions {
    flex-direction: column;
  }

  .form-row {
    grid-template-columns: 1fr;
  }

  .history-controls {
    flex-direction: column;
  }

  .form-actions {
    justify-content: stretch;
  }

  .channels-selection {
    grid-template-columns: 1fr;
  }

  .rule-controls {
    flex-direction: column;
  }

  .rule-actions {
    grid-template-columns: 1fr;
  }
}
</style>
