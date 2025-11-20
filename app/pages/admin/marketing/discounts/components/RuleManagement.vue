<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 w-full">
    <!-- هدر بخش -->
    <div class="p-6 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-lg font-semibold text-gray-900">مدیریت قوانین و شرط‌ها</h2>
          <p class="text-gray-600 mt-1">تعریف قوانین پیچیده، شرط‌های ترکیبی و منطق کسب‌وکار</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="showRuleForm = true">
            <span class="i-heroicons-plus ml-2"></span>
            افزودن قانون جدید
          </button>
        </div>
      </div>
    </div>

    <!-- آمار قوانین -->
    <div class="p-6 border-b border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="bg-blue-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-cog-6-tooth text-blue-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-blue-600">قوانین فعال</p>
              <p class="text-2xl font-bold text-blue-900">{{ stats.activeRules }}</p>
            </div>
          </div>
        </div>
        <div class="bg-green-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-check-circle text-green-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-green-600">قوانین اجرا شده</p>
              <p class="text-2xl font-bold text-green-900">{{ stats.executedRules }}</p>
            </div>
          </div>
        </div>
        <div class="bg-purple-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-clock text-purple-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-purple-600">زمان اجرای متوسط</p>
              <p class="text-2xl font-bold text-purple-900">{{ stats.avgExecutionTime }}ms</p>
            </div>
          </div>
        </div>
        <div class="bg-orange-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-orange-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-exclamation-triangle text-orange-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-orange-600">خطاهای اجرا</p>
              <p class="text-2xl font-bold text-orange-900">{{ stats.executionErrors }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تب‌های قوانین -->
    <div class="border-b border-gray-200">
      <div class="flex border-b border-gray-200 overflow-x-auto">
        <button
v-for="tab in tabs" :key="tab.value" :class="['px-6 py-3 -mb-px font-medium text-sm focus:outline-none whitespace-nowrap', activeTab === tab.value ? 'border-b-2 border-blue-600 text-blue-700' : 'text-gray-500 hover:text-blue-600']"
          @click="activeTab = tab.value">
          {{ tab.label }}
        </button>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="p-6">
      <!-- قوانین فعال -->
      <div v-if="activeTab === 'rules'" class="space-y-6">
        <div class="flex justify-between items-center">
          <div class="flex items-center space-x-4 space-x-reverse">
            <select v-model="filterCategory" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
              <option value="">همه دسته‌ها</option>
              <option value="coupon">کوپن</option>
              <option value="campaign">کمپین</option>
              <option value="user">کاربر</option>
              <option value="order">سفارش</option>
            </select>
            <select v-model="filterStatus" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
              <option value="">همه وضعیت‌ها</option>
              <option value="active">فعال</option>
              <option value="inactive">غیرفعال</option>
              <option value="draft">پیش‌نویس</option>
            </select>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <button class="px-3 py-1 bg-green-100 text-green-700 rounded text-sm hover:bg-green-200" @click="bulkAction('activate')">
              فعال کردن
            </button>
            <button class="px-3 py-1 bg-yellow-100 text-yellow-700 rounded text-sm hover:bg-yellow-200" @click="bulkAction('deactivate')">
              غیرفعال کردن
            </button>
            <button class="px-3 py-1 bg-red-100 text-red-700 rounded text-sm hover:bg-red-200" @click="bulkAction('delete')">
              حذف
            </button>
          </div>
        </div>

        <div class="space-y-4">
          <div v-for="rule in filteredRules" :key="rule.id" class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow">
            <div class="flex items-center justify-between">
              <div class="flex items-center space-x-4 space-x-reverse">
                <input v-model="selectedRules" type="checkbox" :value="rule.id" class="rounded border-gray-300">
                <div class="w-12 h-12 rounded-lg flex items-center justify-center" :style="{ backgroundColor: rule.color + '20', color: rule.color }">
                  <span class="i-heroicons-cog-6-tooth text-lg"></span>
                </div>
                <div>
                  <h4 class="font-medium text-gray-900">{{ rule.name }}</h4>
                  <p class="text-sm text-gray-500">{{ rule.category }} • {{ rule.description }}</p>
                </div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <span :class="['px-2 py-1 rounded-full text-xs', getStatusClass(rule.status)]">
                  {{ getStatusText(rule.status) }}
                </span>
                <button class="text-blue-600 hover:text-blue-900" @click="editRule(rule)">
                  <span class="i-heroicons-pencil-square"></span>
                </button>
                <button class="text-green-600 hover:text-green-900" @click="duplicateRule(rule)">
                  <span class="i-heroicons-document-duplicate"></span>
                </button>
                <button class="text-red-600 hover:text-red-900" @click="deleteRule(rule)">
                  <span class="i-heroicons-trash"></span>
                </button>
              </div>
            </div>
            
            <div class="mt-4">
              <div class="bg-gray-50 p-6 rounded-lg">
                <h5 class="text-sm font-medium text-gray-900 mb-2">شرط‌های قانون:</h5>
                <div class="space-y-2">
                  <div v-for="(condition, index) in rule.conditions" :key="index" class="flex items-center text-sm">
                    <span v-if="index > 0" class="text-gray-400 mx-2">{{ rule.logicOperator }}</span>
                    <span class="bg-blue-100 text-blue-800 px-2 py-1 rounded text-xs">{{ condition.field }} {{ condition.operator }} {{ condition.value }}</span>
                  </div>
                </div>
              </div>
            </div>
            
            <div class="mt-4 grid grid-cols-1 md:grid-cols-3 gap-6 text-sm">
              <div>
                <span class="text-gray-500">آخرین اجرا:</span>
                <span class="font-medium mr-2">{{ formatDate(rule.lastExecuted) }}</span>
              </div>
              <div>
                <span class="text-gray-500">تعداد اجرا:</span>
                <span class="font-medium mr-2">{{ rule.executionCount }}</span>
              </div>
              <div>
                <span class="text-gray-500">نرخ موفقیت:</span>
                <span class="font-medium mr-2">{{ rule.successRate }}%</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- شرط‌های پیشرفته -->
      <div v-if="activeTab === 'conditions'" class="space-y-6">
        <div class="bg-blue-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-blue-900 mb-2">شرط‌های پیشرفته</h4>
          <p class="text-sm text-blue-700">ایجاد شرط‌های پیچیده با عملگرهای منطقی و توابع پیشرفته</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- عملگرهای موجود -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">عملگرهای موجود</h5>
            <div class="space-y-4">
              <div v-for="operator in operators" :key="operator.category" class="space-y-2">
                <h6 class="text-sm font-medium text-gray-700">{{ operator.category }}</h6>
                <div class="grid grid-cols-2 gap-2">
                  <div v-for="op in operator.operators" :key="op.symbol" class="flex items-center justify-between p-2 bg-gray-50 rounded text-sm">
                    <span class="font-mono">{{ op.symbol }}</span>
                    <span class="text-gray-600">{{ op.description }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- توابع پیشرفته -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">توابع پیشرفته</h5>
            <div class="space-y-4">
              <div v-for="func in advancedFunctions" :key="func.name" class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center justify-between mb-2">
                  <h6 class="text-sm font-medium text-gray-900">{{ func.name }}</h6>
                  <span class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded">{{ func.category }}</span>
                </div>
                <p class="text-sm text-gray-600 mb-2">{{ func.description }}</p>
                <div class="bg-gray-50 p-2 rounded text-xs font-mono text-gray-700">
                  {{ func.syntax }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- تست قوانین -->
      <div v-if="activeTab === 'testing'" class="space-y-6">
        <div class="bg-green-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-green-900 mb-2">تست قوانین</h4>
          <p class="text-sm text-green-700">تست قوانین با داده‌های نمونه و مشاهده نتایج</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- انتخاب قانون -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">انتخاب قانون</h5>
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">قانون</label>
                <select v-model="selectedRuleForTest" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" @change="loadRuleDetails">
                  <option value="">انتخاب کنید</option>
                  <option v-for="rule in rules" :key="rule.id" :value="rule.id">{{ rule.name }}</option>
                </select>
              </div>
              
              <div v-if="selectedRuleForTest">
                <h6 class="text-sm font-medium text-gray-700 mb-2">شرح قانون:</h6>
                <div class="bg-gray-50 p-3 rounded text-sm">
                  {{ getSelectedRuleDetails() }}
                </div>
              </div>
            </div>
          </div>

          <!-- داده‌های تست -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">داده‌های تست</h5>
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نوع داده</label>
                <select v-model="testDataType" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" @change="loadTestData">
                  <option value="user">کاربر</option>
                  <option value="order">سفارش</option>
                  <option value="coupon">کوپن</option>
                  <option value="campaign">کمپین</option>
                </select>
              </div>
              
              <div v-if="testData">
                <h6 class="text-sm font-medium text-gray-700 mb-2">داده‌های نمونه:</h6>
                <div class="bg-gray-50 p-3 rounded text-sm max-h-40 overflow-y-auto">
                  <pre class="text-xs">{{ JSON.stringify(testData, null, 2) }}</pre>
                </div>
              </div>
              
              <button :disabled="!selectedRuleForTest || !testData" class="w-full px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors disabled:bg-gray-300 disabled:cursor-not-allowed" @click="runTest">
                اجرای تست
              </button>
            </div>
          </div>
        </div>

        <!-- نتایج تست -->
        <div v-if="testResults" class="border border-gray-200 rounded-lg p-6">
          <h5 class="font-medium text-gray-900 mb-4">نتایج تست</h5>
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm font-medium text-gray-700">نتیجه:</span>
              <span :class="['px-3 py-1 rounded-full text-sm', testResults.success ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800']">
                {{ testResults.success ? 'موفق' : 'ناموفق' }}
              </span>
            </div>
            
            <div v-if="testResults.details" class="bg-gray-50 p-6 rounded">
              <h6 class="text-sm font-medium text-gray-700 mb-2">جزئیات اجرا:</h6>
              <div class="space-y-2 text-sm">
                <div v-for="(detail, index) in testResults.details" :key="index" class="flex items-center justify-between">
                  <span class="text-gray-600">{{ detail.condition }}</span>
                  <span :class="['px-2 py-1 rounded text-xs', detail.result ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800']">
                    {{ detail.result ? 'صحیح' : 'غلط' }}
                  </span>
                </div>
              </div>
            </div>
            
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">زمان اجرا:</span>
              <span class="font-medium">{{ testResults.executionTime }}ms</span>
            </div>
          </div>
        </div>
      </div>

      <!-- آمار و گزارش‌ها -->
      <div v-if="activeTab === 'analytics'" class="space-y-6">
        <div class="bg-purple-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-purple-900 mb-2">آمار و گزارش‌های قوانین</h4>
          <p class="text-sm text-purple-700">مشاهده عملکرد قوانین و تحلیل کارایی</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- نمودار عملکرد -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">عملکرد قوانین</h5>
            <div class="h-64 flex items-center justify-center text-gray-400">[نمودار عملکرد قوانین]</div>
          </div>

          <!-- جدول آمار -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">آمار قوانین</h5>
            <div class="space-y-3">
              <div v-for="stat in ruleStats" :key="stat.name" class="flex items-center justify-between p-3 bg-gray-50 rounded">
                <span class="text-sm text-gray-700">{{ stat.name }}</span>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <div class="w-32 bg-gray-200 rounded-full h-2">
                    <div class="h-2 rounded-full bg-blue-600" :style="{ width: stat.percentage + '%' }"></div>
                  </div>
                  <span class="text-sm text-gray-600">{{ stat.percentage }}%</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال فرم قانون -->
    <div v-if="showRuleForm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-xl max-w-4xl w-full mx-4 max-h-[90vh] overflow-y-auto">
        <div class="p-6 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold text-gray-900">
              {{ editingRule ? 'ویرایش قانون' : 'افزودن قانون جدید' }}
            </h3>
            <button class="text-gray-400 hover:text-gray-600" @click="closeForm">
              <span class="i-heroicons-x-mark text-xl"></span>
            </button>
          </div>
        </div>
        
        <form class="p-6 space-y-6" @submit.prevent="handleSubmit">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نام قانون *</label>
              <input v-model="form.name" type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="نام قانون">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی *</label>
              <select v-model="form.category" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="coupon">کوپن</option>
                <option value="campaign">کمپین</option>
                <option value="user">کاربر</option>
                <option value="order">سفارش</option>
              </select>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
            <textarea v-model="form.description" rows="3" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="توضیحات قانون"></textarea>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">شرط‌های قانون *</label>
            <div class="space-y-4">
              <div v-for="(condition, index) in form.conditions" :key="index" class="flex items-center space-x-3 space-x-reverse p-6 border border-gray-200 rounded-lg">
                <select v-model="condition.field" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
                  <option value="">انتخاب فیلد</option>
                  <option value="user_age">سن کاربر</option>
                  <option value="order_amount">مبلغ سفارش</option>
                  <option value="user_level">سطح کاربر</option>
                  <option value="coupon_count">تعداد کوپن‌ها</option>
                </select>
                <select v-model="condition.operator" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
                  <option value="">عملگر</option>
                  <option value="equals">برابر است</option>
                  <option value="not_equals">برابر نیست</option>
                  <option value="greater_than">بزرگتر از</option>
                  <option value="less_than">کوچکتر از</option>
                  <option value="contains">شامل</option>
                </select>
                <input v-model="condition.value" type="text" class="px-3 py-2 border border-gray-300 rounded-lg text-sm" placeholder="مقدار">
                <button type="button" class="text-red-600 hover:text-red-900" @click="removeCondition(index)">
                  <span class="i-heroicons-trash"></span>
                </button>
              </div>
              <button type="button" class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors" @click="addCondition">
                <span class="i-heroicons-plus ml-2"></span>
                افزودن شرط
              </button>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">عملگر منطقی</label>
              <select v-model="form.logicOperator" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="AND">AND (و)</option>
                <option value="OR">OR (یا)</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
              <select v-model="form.status" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="active">فعال</option>
                <option value="inactive">غیرفعال</option>
                <option value="draft">پیش‌نویس</option>
              </select>
            </div>
          </div>
        </form>
        
        <div class="p-6 border-t border-gray-200 flex justify-end space-x-3 space-x-reverse">
          <button class="px-6 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors" @click="closeForm">
            انصراف
          </button>
          <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="handleSubmit">
            ذخیره
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, reactive, ref } from 'vue'
import { useAuth } from '~/composables/useAuth'
// احراز هویت auto-import از sidebase استفاده می‌شود

const { user } = useAuth()
const canDeleteRule = computed(() => {
  const roles = user.value?.role ? [user.value.role] : []
  return roles.includes('developer')
})

const activeTab = ref('rules')
const showRuleForm = ref(false)
const editingRule = ref<any>(null)
const selectedRules = ref<number[]>([])
const filterCategory = ref('')
const filterStatus = ref('')
const selectedRuleForTest = ref('')
const testDataType = ref('user')
const testData = ref<any>(null)
const testResults = ref<any>(null)

const tabs = [
  { value: 'rules', label: 'قوانین فعال' },
  { value: 'conditions', label: 'شرط‌های پیشرفته' },
  { value: 'testing', label: 'تست قوانین' },
  { value: 'analytics', label: 'آمار و گزارش‌ها' }
]

const stats = ref({
  activeRules: 15,
  executedRules: 1247,
  avgExecutionTime: 45,
  executionErrors: 3
})

const rules = ref([
  {
    id: 1,
    name: 'قانون تخفیف کاربران جدید',
    category: 'user',
    description: 'اعطای تخفیف ۲۰٪ به کاربران جدید',
    color: '#3b82f6',
    status: 'active',
    conditions: [
      { field: 'user_registration_date', operator: 'greater_than', value: '30 days ago' }
    ],
    logicOperator: 'AND',
    lastExecuted: '2024-01-15T10:30:00',
    executionCount: 156,
    successRate: 95.2
  },
  {
    id: 2,
    name: 'قانون حداقل مبلغ سفارش',
    category: 'order',
    description: 'اعمال کوپن فقط برای سفارش‌های بالای ۵۰۰ هزار تومان',
    color: '#10b981',
    status: 'active',
    conditions: [
      { field: 'order_amount', operator: 'greater_than', value: '500000' }
    ],
    logicOperator: 'AND',
    lastExecuted: '2024-01-14T09:15:00',
    executionCount: 89,
    successRate: 87.6
  }
])

const operators = ref([
  {
    category: 'مقایسه',
    operators: [
      { symbol: '==', description: 'برابر است' },
      { symbol: '!=', description: 'برابر نیست' },
      { symbol: '>', description: 'بزرگتر از' },
      { symbol: '<', description: 'کوچکتر از' },
      { symbol: '>=', description: 'بزرگتر یا برابر' },
      { symbol: '<=', description: 'کوچکتر یا برابر' }
    ]
  },
  {
    category: 'رشته',
    operators: [
      { symbol: 'contains', description: 'شامل' },
      { symbol: 'starts_with', description: 'شروع با' },
      { symbol: 'ends_with', description: 'پایان با' },
      { symbol: 'regex', description: 'عبارت منظم' }
    ]
  },
  {
    category: 'منطقی',
    operators: [
      { symbol: 'AND', description: 'و' },
      { symbol: 'OR', description: 'یا' },
      { symbol: 'NOT', description: 'نه' }
    ]
  }
])

const advancedFunctions = ref([
  {
    name: 'COUNT',
    category: 'آمار',
    description: 'شمارش تعداد موارد',
    syntax: 'COUNT(field_name)'
  },
  {
    name: 'SUM',
    category: 'ریاضی',
    description: 'مجموع مقادیر',
    syntax: 'SUM(field_name)'
  },
  {
    name: 'AVG',
    category: 'ریاضی',
    description: 'میانگین مقادیر',
    syntax: 'AVG(field_name)'
  },
  {
    name: 'DATE_DIFF',
    category: 'تاریخ',
    description: 'تفاوت بین دو تاریخ',
    syntax: 'DATE_DIFF(date1, date2, unit)'
  }
])

const ruleStats = ref([
  { name: 'قوانین موفق', percentage: 85 },
  { name: 'قوانین با خطا', percentage: 10 },
  { name: 'قوانین غیرفعال', percentage: 5 }
])

const form = reactive({
  name: '',
  category: 'coupon',
  description: '',
  conditions: [{ field: '', operator: '', value: '' }],
  logicOperator: 'AND',
  status: 'active',
  color: '#3b82f6'
})

const filteredRules = computed(() => {
  let filtered = rules.value
  
  if (filterCategory.value) {
    filtered = filtered.filter(r => r.category === filterCategory.value)
  }
  
  if (filterStatus.value) {
    filtered = filtered.filter(r => r.status === filterStatus.value)
  }
  
  return filtered
})

const formatDate = (date: string): string => {
  return new Intl.DateTimeFormat('fa-IR').format(new Date(date))
}

const getStatusClass = (status: string): string => {
  const classes = {
    active: 'bg-green-100 text-green-800',
    inactive: 'bg-red-100 text-red-800',
    draft: 'bg-gray-100 text-gray-800'
  }
  return classes[status as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string): string => {
  const texts = {
    active: 'فعال',
    inactive: 'غیرفعال',
    draft: 'پیش‌نویس'
  }
  return texts[status as keyof typeof texts] || status
}

const editRule = (rule: any) => {
  editingRule.value = rule
  Object.assign(form, rule)
  showRuleForm.value = true
}

const duplicateRule = (rule: any) => {
  const duplicate = { ...rule, id: Date.now(), name: `${rule.name} (کپی)` }
  rules.value.unshift(duplicate)
}

const deleteRule = async (rule: any) => {
  if (confirm(`آیا از حذف قانون "${rule.name}" اطمینان دارید؟`)) {
    try {
      const index = rules.value.findIndex(r => r.id === rule.id)
      if (index !== -1) {
        rules.value.splice(index, 1)
      }
    } catch (error) {
      console.error('خطا در حذف قانون:', error)
    }
  }
}

const bulkAction = (action: string) => {
  if (action === 'delete' && !canDeleteRule.value) {
    alert('شما مجوز حذف قوانین را ندارید.');
    return;
  }
  if (selectedRules.value.length === 0) {
    alert('لطفاً قوانینی را انتخاب کنید')
    return
  }
  
  switch (action) {
    case 'activate':
      rules.value.forEach(r => {
        if (selectedRules.value.includes(r.id)) {
          r.status = 'active'
        }
      })
      break
    case 'deactivate':
      rules.value.forEach(r => {
        if (selectedRules.value.includes(r.id)) {
          r.status = 'inactive'
        }
      })
      break
    case 'delete':
      if (confirm(`آیا از حذف ${selectedRules.value.length} قانون اطمینان دارید؟`)) {
        rules.value = rules.value.filter(r => !selectedRules.value.includes(r.id))
        selectedRules.value = []
      }
      break
  }
}

const addCondition = () => {
  form.conditions.push({ field: '', operator: '', value: '' })
}

const removeCondition = (index: number) => {
  if (form.conditions.length > 1) {
    form.conditions.splice(index, 1)
  }
}

const loadRuleDetails = () => {
  // پیاده‌سازی بارگذاری جزئیات قانون
  console.log('بارگذاری جزئیات قانون:', selectedRuleForTest.value)
}

const getSelectedRuleDetails = () => {
  const rule = rules.value.find(r => r.id === parseInt(selectedRuleForTest.value))
  return rule ? rule.description : ''
}

const loadTestData = () => {
  // پیاده‌سازی بارگذاری داده‌های تست
  const testDataMap = {
    user: { name: 'علی احمدی', age: 25, level: 'silver', registration_date: '2024-01-01' },
    order: { id: 'ORD001', amount: 750000, items: 3, date: '2024-01-15' },
    coupon: { code: 'WELCOME2024', discount: 20, type: 'percentage', valid_until: '2024-02-15' },
    campaign: { name: 'تخفیف زمستانه', status: 'active', start_date: '2024-01-01', end_date: '2024-03-01' }
  }
  testData.value = testDataMap[testDataType.value as keyof typeof testDataMap]
}

const runTest = () => {
  // پیاده‌سازی اجرای تست
  const startTime = Date.now()
  
  // شبیه‌سازی اجرای قانون
  setTimeout(() => {
    const executionTime = Date.now() - startTime
    testResults.value = {
      success: Math.random() > 0.3,
      executionTime,
      details: [
        { condition: 'سن کاربر > 18', result: true },
        { condition: 'مبلغ سفارش > 500000', result: true },
        { condition: 'سطح کاربر = silver', result: false }
      ]
    }
  }, 500)
}

const handleSubmit = async () => {
  if (!form.name || !form.category || form.conditions.length === 0) {
    alert('لطفاً فیلدهای اجباری را پر کنید')
    return
  }
  
  if (editingRule.value) {
    Object.assign(editingRule.value, form)
  } else {
    rules.value.unshift({
      id: Date.now(),
      ...form,
      lastExecuted: null,
      executionCount: 0,
      successRate: 0
    })
  }
  closeForm()
}

const closeForm = () => {
  showRuleForm.value = false
  editingRule.value = null
  Object.assign(form, { name: '', category: 'coupon', description: '', conditions: [{ field: '', operator: '', value: '' }], logicOperator: 'AND', status: 'active', color: '#3b82f6' })
}
</script>

<!--
  کامپوننت مدیریت قوانین و شرط‌ها
  شامل:
  1. مدیریت قوانین فعال
  2. شرط‌های پیشرفته و عملگرها
  3. تست قوانین با داده‌های نمونه
  4. آمار و گزارش‌های قوانین
  توضیحات کامل به فارسی در کد
--> 
