<template>
  <div v-if="hasAccess">
    <div v-if="isOpen" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen px-4 pt-4 pb-20 text-center sm:block sm:p-0">
        <!-- Overlay -->
        <div class="fixed inset-0 transition-opacity bg-gray-500 bg-opacity-75" @click="$emit('close')"></div>

        <!-- Modal -->
        <div class="inline-block w-full max-w-4xl px-6 py-6 my-8 overflow-hidden text-right transition-all transform bg-white rounded-lg shadow-xl align-middle">
          <!-- Header -->
          <div class="flex items-center justify-between mb-6">
            <h3 class="text-lg font-medium text-gray-900">تنظیمات پیشرفته A/B تست</h3>
            <button class="text-gray-400 hover:text-gray-600" @click="$emit('close')">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>

          <!-- Content -->
          <div class="space-y-6">
            <!-- تنظیمات عمومی -->
            <div class="bg-gray-50 rounded-lg p-6">
              <h4 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات عمومی</h4>
              
              <div class="space-y-4">
                <!-- فعال/غیرفعال کردن A/B تست -->
                <div class="flex items-center justify-between">
                  <div>
                    <div class="text-sm font-medium text-gray-700">فعال کردن A/B تست</div>
                    <div class="text-sm text-gray-500">اجازه اجرای تست‌های A/B در سایت</div>
                  </div>
                  <label class="relative inline-flex items-center cursor-pointer">
                    <input v-model="settings.enableABTesting" type="checkbox" class="sr-only peer">
                    <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                  </label>
                </div>

                <!-- حداکثر تعداد تست‌های همزمان -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تعداد تست‌های همزمان</label>
                  <input
                    v-model.number="settings.maxConcurrentTests"
                    type="number"
                    min="1"
                    max="50"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  />
                  <p class="text-xs text-gray-500 mt-1">تعداد تست‌هایی که می‌توانند همزمان اجرا شوند</p>
                </div>

                <!-- مدت زمان پیش‌فرض تست -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">مدت زمان پیش‌فرض تست (روز)</label>
                  <input
                    v-model.number="settings.defaultTestDuration"
                    type="number"
                    min="1"
                    max="365"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  />
                </div>

                <!-- درصد تقسیم ترافیک پیش‌فرض -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">درصد تقسیم ترافیک پیش‌فرض</label>
                  <div class="flex items-center space-x-4 space-x-reverse">
                    <div class="flex-1">
                      <label class="block text-xs text-gray-500 mb-1">نسخه A</label>
                      <input
                        v-model.number="settings.defaultTrafficSplitA"
                        type="number"
                        min="0"
                        max="100"
                        class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                      />
                    </div>
                    <div class="flex-1">
                      <label class="block text-xs text-gray-500 mb-1">نسخه B</label>
                      <input
                        v-model.number="settings.defaultTrafficSplitB"
                        type="number"
                        min="0"
                        max="100"
                        class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                      />
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- تنظیمات آماری -->
            <div class="bg-gray-50 rounded-lg p-6">
              <h4 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات آماری</h4>
              
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <!-- سطح اطمینان پیش‌فرض -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">سطح اطمینان پیش‌فرض</label>
                  <select
                  v-model="settings.defaultConfidenceLevel"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                >
                  <option value="90">90%</option>
                  <option value="95">95%</option>
                  <option value="99">99%</option>
                </select>
              </div>

              <!-- حداقل نمونه‌گیری پیش‌فرض -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">حداقل نمونه‌گیری پیش‌فرض</label>
                <input
                  v-model.number="settings.defaultMinSampleSize"
                  type="number"
                  min="100"
                  step="100"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                />
              </div>

              <!-- قدرت تست پیش‌فرض -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">قدرت تست پیش‌فرض</label>
                <select
                  v-model="settings.defaultTestPower"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                >
                  <option value="80">80%</option>
                  <option value="85">85%</option>
                  <option value="90">90%</option>
                </select>
              </div>

              <!-- روش محاسبه آماری -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">روش محاسبه آماری</label>
                <select
                  v-model="settings.statisticalMethod"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                >
                  <option value="z_test">Z-Test</option>
                  <option value="t_test">T-Test</option>
                  <option value="chi_square">Chi-Square Test</option>
                  <option value="bayesian">Bayesian Analysis</option>
                </select>
              </div>
            </div>
          </div>

          <!-- تنظیمات اعلان‌ها -->
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات اعلان‌ها</h4>
            
            <div class="space-y-4">
              <!-- اعلان‌های ایمیل -->
              <div class="flex items-center justify-between">
                <div>
                  <div class="text-sm font-medium text-gray-700">اعلان‌های ایمیل</div>
                  <div class="text-sm text-gray-500">ارسال اعلان‌ها از طریق ایمیل</div>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="settings.emailNotifications" type="checkbox" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>

              <!-- اعلان‌های درون‌برنامه‌ای -->
              <div class="flex items-center justify-between">
                <div>
                  <div class="text-sm font-medium text-gray-700">اعلان‌های درون‌برنامه‌ای</div>
                  <div class="text-sm text-gray-500">نمایش اعلان‌ها در پنل مدیریت</div>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="settings.inAppNotifications" type="checkbox" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>

              <!-- هشدار تست‌های نزدیک به پایان -->
              <div class="flex items-center justify-between">
                <div>
                  <div class="text-sm font-medium text-gray-700">هشدار تست‌های نزدیک به پایان</div>
                  <div class="text-sm text-gray-500">اعلان برای تست‌هایی که به زودی تمام می‌شوند</div>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="settings.endDateWarnings" type="checkbox" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>

              <!-- اعلان نتایج معنادار -->
              <div class="flex items-center justify-between">
                <div>
                  <div class="text-sm font-medium text-gray-700">اعلان نتایج معنادار</div>
                  <div class="text-sm text-gray-500">اعلان وقتی نتایج آماری معنادار می‌شوند</div>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="settings.significantResultsNotifications" type="checkbox" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>

              <!-- تعداد روزهای هشدار -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">تعداد روزهای هشدار قبل از پایان</label>
                <input
                  v-model.number="settings.warningDaysBeforeEnd"
                  type="number"
                  min="1"
                  max="30"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
            </div>
          </div>

          <!-- تنظیمات یکپارچه‌سازی -->
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات یکپارچه‌سازی</h4>
            
            <div class="space-y-4">
              <!-- Google Analytics -->
              <div class="flex items-center justify-between">
                <div>
                  <div class="text-sm font-medium text-gray-700">Google Analytics</div>
                  <div class="text-sm text-gray-500">ارسال داده‌های تست به Google Analytics</div>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="settings.googleAnalyticsIntegration" type="checkbox" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>

              <!-- Facebook Pixel -->
              <div class="flex items-center justify-between">
                <div>
                  <div class="text-sm font-medium text-gray-700">Facebook Pixel</div>
                  <div class="text-sm text-gray-500">ارسال داده‌های تست به Facebook Pixel</div>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="settings.facebookPixelIntegration" type="checkbox" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>

              <!-- CRM Integration -->
              <div class="flex items-center justify-between">
                <div>
                  <div class="text-sm font-medium text-gray-700">یکپارچه‌سازی CRM</div>
                  <div class="text-sm text-gray-500">ارسال نتایج تست به سیستم CRM</div>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="settings.crmIntegration" type="checkbox" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>
            </div>
          </div>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex items-center justify-end space-x-4 space-x-reverse pt-6 border-t mt-6">
          <button
            class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50"
            @click="$emit('close')"
          >
            انصراف
          </button>
          <button
            class="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-100 border border-gray-300 rounded-lg hover:bg-gray-200"
            @click="resetToDefaults"
          >
            بازنشانی به پیش‌فرض
          </button>
          <button
            class="px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-lg hover:bg-blue-700"
            @click="saveSettings"
          >
            ذخیره تنظیمات
          </button>
        </div>
      </div>
    </div>
  </div>
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '@/composables/useAuth'
import { computed, onMounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'

const { user } = useAuth()
const router = useRouter()

const hasAccess = computed(() => {
  return user.value?.role === 'admin' || user.value?.role === 'developer'
})

watch(() => user.value, (newUser) => {
  if (!newUser || (newUser.role !== 'admin' && newUser.role !== 'developer')) {
    router.push('/404')
  }
})

onMounted(() => {
  if (!user.value || (user.value.role !== 'admin' && user.value.role !== 'developer')) {
    router.push('/404')
  }
})

// Props
const _props = defineProps<{
  isOpen: boolean
}>()

// Events
const emit = defineEmits(['close', 'save'])

// تنظیمات پیش‌فرض
const defaultSettings = {
  enableABTesting: true,
  maxConcurrentTests: 10,
  defaultTestDuration: 30,
  defaultTrafficSplitA: 50,
  defaultTrafficSplitB: 50,
  defaultConfidenceLevel: 95,
  defaultMinSampleSize: 1000,
  defaultTestPower: 80,
  statisticalMethod: 'z_test',
  emailNotifications: true,
  inAppNotifications: true,
  endDateWarnings: true,
  significantResultsNotifications: true,
  warningDaysBeforeEnd: 3,
  googleAnalyticsIntegration: false,
  facebookPixelIntegration: false,
  crmIntegration: false
}

// تنظیمات فعلی
const settings = ref({ ...defaultSettings })

// تنظیم خودکار درصد تقسیم ترافیک
watch(() => settings.value.defaultTrafficSplitA, (newVal) => {
  settings.value.defaultTrafficSplitB = 100 - newVal
})

watch(() => settings.value.defaultTrafficSplitB, (newVal) => {
  settings.value.defaultTrafficSplitA = 100 - newVal
})

// بازنشانی به تنظیمات پیش‌فرض
const resetToDefaults = () => {
  if (confirm('آیا از بازنشانی تنظیمات به حالت پیش‌فرض اطمینان دارید؟')) {
    settings.value = { ...defaultSettings }
  }
}

// ذخیره تنظیمات
const saveSettings = () => {
  emit('save', settings.value)
  emit('close')
}
</script>
