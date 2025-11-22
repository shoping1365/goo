<template>
  <div class="p-6" dir="rtl">
    <!-- تب‌ها -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 mb-8">
      <div class="border-b border-gray-200">
        <nav class="flex space-x-8 space-x-reverse px-6" aria-label="Tabs">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            :class="[
              activeTab === tab.id
                ? 'border-blue-500 text-blue-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
              'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm transition-colors'
            ]"
            @click="activeTab = tab.id"
          >
            {{ tab.name }}
          </button>
        </nav>
      </div>

      <!-- محتوای تب‌ها -->
      <div class="p-6">
        <!-- تب لیست اپراتورها -->
        <div v-if="activeTab === 'operators-list'">
          <!-- هدر با دکمه دعوت -->
          <div class="bg-purple-600 text-white p-6 rounded-lg mb-8">
            <div class="flex justify-between items-start">
              <div>
                <h1 class="text-2xl font-bold mb-2">لیست اپراتورها</h1>
                <p class="text-purple-100">نمایش ساعات کاری و مکالمات هر اپراتور، اعمال تنظیمات، افزودن و حذف اپراتورها در این بخش انجام می‌شود.</p>
              </div>
              <button 
                class="bg-white text-purple-600 px-6 py-3 rounded-lg font-medium hover:bg-purple-50 transition-colors"
                @click="showInviteModal = true"
              >
                دعوت اپراتور جدید
              </button>
            </div>
          </div>

          <!-- اپراتور منتخب هفته قبل -->
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6 mb-8">
            <h2 class="text-xl font-bold text-gray-900 mb-6">اپراتور منتخب هفته قبل</h2>
            
            <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
              <!-- نمودار -->
              <div class="lg:col-span-2">
                <div class="bg-gray-50 rounded-lg p-6 h-64 flex items-center justify-center">
                  <div class="text-center text-gray-500">
                    <svg class="w-16 h-16 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                    </svg>
                    <p>نمودار عملکرد هفتگی</p>
                    <p class="text-sm">امروز، سه‌شنبه، دوشنبه، یکشنبه، شنبه، جمعه، پنج‌شنبه</p>
                  </div>
                </div>
                
                <!-- کارت‌های آمار -->
                <div class="grid grid-cols-2 gap-6 mt-4">
                  <div class="bg-white border border-gray-200 rounded-lg p-6 text-center">
                    <div class="text-2xl font-bold text-green-600">۳۶</div>
                    <div class="text-sm text-gray-600">مکالمه‌های قبول شده</div>
                  </div>
                  <div class="bg-white border border-gray-200 rounded-lg p-6 text-center">
                    <div class="text-2xl font-bold text-blue-600">۱۲۸:۱۷</div>
                    <div class="text-sm text-gray-600">کل ساعات کاری</div>
                  </div>
                </div>
              </div>

              <!-- کارت اپراتور منتخب -->
              <div class="bg-white border border-gray-200 rounded-lg p-6">
                <div class="text-center">
                  <div class="w-20 h-20 bg-purple-100 rounded-full flex items-center justify-center mx-auto mb-4">
                    <span class="text-2xl font-bold text-purple-600">آ</span>
                  </div>
                  <h3 class="text-lg font-bold text-gray-900 mb-2">آ . آذری</h3>
                  <p class="text-sm text-green-600 mb-4">اگر سوالی هست در خدمتم</p>
                  
                  <div class="space-y-3">
                    <div class="flex justify-between items-center">
                      <span class="text-sm text-gray-600">کل ساعات کاری:</span>
                      <span class="text-sm font-medium">۱۲۸:۱۷</span>
                    </div>
                    <div class="flex justify-between items-center">
                      <span class="text-sm text-gray-600">مکالمه‌های قبول شده:</span>
                      <span class="text-sm font-medium">۳۶</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- لیست اپراتورها و عملکردشان -->
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <div class="flex justify-between items-center mb-6">
              <h2 class="text-xl font-bold text-gray-900">لیست اپراتورها و عملکردشان - مرداد ۱۴۰۴</h2>
              <div class="flex space-x-4 space-x-reverse">
                <select class="px-3 py-2 border border-gray-300 rounded-md text-sm">
                  <option>۱۴۰۴</option>
                </select>
                <select class="px-3 py-2 border border-gray-300 rounded-md text-sm">
                  <option>مرداد</option>
                </select>
              </div>
            </div>

            <!-- کارت‌های اپراتورها -->
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              <div v-for="operator in operators" :key="operator.id" class="bg-white border border-gray-200 rounded-lg p-6">
                <div class="flex items-center mb-4">
                  <div class="w-12 h-12 bg-purple-100 rounded-full flex items-center justify-center mr-3">
                    <span class="text-lg font-bold text-purple-600">{{ operator.name.charAt(0) }}</span>
                  </div>
                  <div>
                    <h3 class="font-medium text-gray-900">{{ operator.name }}</h3>
                    <p class="text-sm text-gray-500">{{ operator.email }}</p>
                  </div>
                </div>
                
                <div class="space-y-2">
                  <div class="flex justify-between items-center">
                    <span class="text-sm text-gray-600">قبول شده:</span>
                    <span class="text-sm font-medium text-green-600">{{ operator.acceptedCount }}</span>
                  </div>
                  <div class="flex justify-between items-center">
                    <span class="text-sm text-gray-600">کل ساعات کاری:</span>
                    <span class="text-sm font-medium">{{ operator.totalHours }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- تب تخصیص اپراتور -->
        <div v-if="activeTab === 'operator-assignment'">
          <div class="max-w-2xl mx-auto">
            <!-- هدر -->
            <div class="text-center mb-8">
              <h1 class="text-2xl font-bold text-gray-900 mb-2">تخصیص / عدم تخصیص اپراتور</h1>
              <p class="text-gray-600">iranxia.net</p>
              <p class="text-sm text-gray-500 mt-2">{{ selectedOperators.length }} اپراتور</p>
            </div>

            <!-- جستجو -->
            <div class="bg-gray-100 rounded-lg p-6 mb-6">
              <div class="flex items-center">
                <svg class="w-5 h-5 text-gray-400 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                </svg>
                <input 
                  v-model="operatorSearch" 
                  type="text" 
                  placeholder="جست و جوی اپراتور" 
                  class="flex-1 bg-transparent border-none outline-none text-gray-700 placeholder-gray-500"
                >
              </div>
            </div>

            <!-- لیست اپراتورها -->
            <div class="bg-white rounded-lg border border-gray-200 p-6">
              <h2 class="text-lg font-bold text-gray-900 mb-2">اپراتورها</h2>
              <p class="text-sm text-gray-600 mb-6">برای تخصیص اپراتور به سایت یا عدم تخصیص آن بر روی آواتار کلیک کنید.</p>
              
              <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
                <div 
                  v-for="operator in filteredOperators" 
                  :key="operator.id"
                  class="text-center cursor-pointer group"
                  @click="toggleOperatorAssignment(operator.id)"
                >
                  <div class="relative inline-block">
                    <div class="w-16 h-16 rounded-full border-2 border-orange-400 flex items-center justify-center mx-auto mb-2 group-hover:border-orange-500 transition-colors">
                      <div class="w-12 h-12 bg-gradient-to-br from-green-400 to-green-600 rounded-full flex items-center justify-center">
                        <span class="text-white font-bold text-lg">{{ operator.name.charAt(0) }}</span>
                      </div>
                    </div>
                    <!-- نشانگر انتخاب -->
                    <div 
                      v-if="isOperatorAssigned(operator.id)"
                      class="absolute -bottom-1 -right-1 w-6 h-6 bg-purple-500 rounded-full flex items-center justify-center"
                    >
                      <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                      </svg>
                    </div>
                  </div>
                  <p class="text-sm font-medium text-gray-900">{{ operator.name }}</p>
                </div>
              </div>
            </div>

            <!-- دکمه‌های عملیات -->
            <div class="flex justify-between items-center mt-8">
              <button 
                class="bg-purple-600 text-white px-8 py-3 rounded-lg font-medium hover:bg-purple-700 transition-colors"
                @click="saveOperatorAssignment"
              >
                ذخیره
              </button>
              <button 
                class="text-gray-600 hover:text-gray-800 font-medium"
                @click="cancelOperatorAssignment"
              >
                انصراف
              </button>
            </div>
          </div>
        </div>

        <!-- تب تخصیص لینک به اپراتور -->
        <div v-if="activeTab === 'link-assignment'" class="min-h-screen bg-gradient-to-b from-purple-900 via-purple-700 to-gray-300 p-6">
          <div class="max-w-4xl mx-auto">
            <!-- کارت اصلی -->
            <div class="bg-white rounded-2xl shadow-xl p-8">
              <!-- هدر -->
              <div class="mb-8">
                <h1 class="text-2xl font-bold text-gray-900 mb-4">تخصیص لینک به اپراتور</h1>
                <p class="text-gray-600 text-sm leading-relaxed">
                  با این قابلیت به جای ایجاد تیم‌ها می‌توانید مستقیماً یک اپراتور را به صفحه خاصی اختصاص دهید. این اپراتور مسئولیت پاسخگویی به تمام مکالمات صورت گرفته در آن صفحه را بر عهده خواهد داشت.
                </p>
              </div>

              <!-- فرم تخصیص -->
              <div class="space-y-6">
                <!-- فیلد لینک -->
                <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
                  <div class="lg:col-span-1">
                    <label class="block text-sm font-medium text-gray-700 mb-2">لینک</label>
                    <input 
                      v-model="linkAssignment.link" 
                      type="text" 
                      placeholder="https://example.com/page"
                      class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                    >
                  </div>

                  <!-- فیلدهای انتخاب اپراتور -->
                  <div class="lg:col-span-2">
                    <div class="space-y-4">
                      <div v-for="(operatorField, index) in linkAssignment.operators" :key="index" class="flex items-start space-x-3 space-x-reverse">
                        <div class="flex-1">
                          <label class="block text-sm font-medium text-gray-700 mb-2">انتخاب اپراتور</label>
                          <select 
                            v-model="operatorField.operatorId"
                            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                          >
                            <option value="">انتخاب کنید</option>
                            <option v-for="operator in operators" :key="operator.id" :value="operator.id">
                              {{ operator.name }}
                            </option>
                          </select>
                        </div>
                        
                        <!-- دکمه‌های عملیات -->
                        <div class="flex space-x-2 space-x-reverse mt-8">
                          <!-- دکمه اضافه کردن -->
                          <button 
                            v-if="index === linkAssignment.operators.length - 1"
                            class="w-10 h-10 bg-green-500 text-white rounded-lg flex items-center justify-center hover:bg-green-600 transition-colors"
                            title="افزودن اپراتور"
                            @click="addOperatorField"
                          >
                            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
                            </svg>
                          </button>
                          
                          <!-- دکمه حذف -->
                          <button 
                            v-if="linkAssignment.operators.length > 1"
                            class="w-10 h-10 bg-gray-500 text-white rounded-lg flex items-center justify-center hover:bg-gray-600 transition-colors"
                            title="حذف اپراتور"
                            @click="removeOperatorField(index)"
                          >
                            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                            </svg>
                          </button>
                          
                          <!-- دکمه تنظیمات -->
                          <button 
                            class="w-10 h-10 bg-white border border-gray-300 text-gray-600 rounded-lg flex items-center justify-center hover:bg-gray-50 transition-colors"
                            title="تنظیمات"
                          >
                            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                            </svg>
                          </button>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <!-- دکمه ذخیره -->
              <div class="flex justify-end mt-8">
                <button 
                  class="bg-purple-800 text-white px-8 py-3 rounded-lg font-medium hover:bg-purple-900 transition-colors"
                  @click="saveLinkAssignment"
                >
                  ذخیره
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال دعوت اپراتور جدید -->
    <div v-if="showInviteModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-xl max-w-2xl w-full mx-4 max-h-[90vh] overflow-y-auto">
        <!-- هدر مودال -->
        <div class="flex justify-between items-center p-6 border-b border-gray-200">
          <h2 class="text-xl font-bold text-gray-900">دعوت اپراتور جدید</h2>
          <button class="text-gray-400 hover:text-gray-600" @click="showInviteModal = false">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <!-- محتوای مودال -->
        <div class="p-6">
          <form @submit.prevent="inviteOperator">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <!-- ستون راست -->
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">عنوان</label>
                  <input v-model="newOperator.title" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">ایمیل</label>
                  <input v-model="newOperator.email" type="email" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">رمز عبور</label>
                  <input v-model="newOperator.password" type="password" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
                </div>
              </div>

              <!-- ستون چپ -->
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">نام کامل</label>
                  <input v-model="newOperator.fullName" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">شماره تلفن</label>
                  <div class="flex">
                    <div class="flex items-center px-3 py-2 border border-r-0 border-gray-300 rounded-r-md bg-gray-50">
                      <span class="text-sm text-gray-500">🇮🇷</span>
                      <span class="text-sm text-gray-500 mr-2">+۹۸</span>
                      <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                      </svg>
                    </div>
                    <input v-model="newOperator.phone" type="tel" class="flex-1 px-3 py-2 border border-gray-300 rounded-l-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500" placeholder="شماره تلفن">
                  </div>
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">تکرار رمز عبور</label>
                  <input v-model="newOperator.confirmPassword" type="password" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
                </div>
              </div>
            </div>

            <!-- دسترسی به وب‌سایت -->
            <div class="mt-6">
              <label class="block text-sm font-medium text-gray-700 mb-2">اپراتور به این وب‌سایتها دسترسی داشته باشد:</label>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="newOperator.websiteAccess" type="checkbox" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="text-sm text-gray-700">iranxia.net</span>
              </div>
            </div>

            <!-- سطح دسترسی -->
            <div class="mt-6">
              <label class="block text-sm font-medium text-gray-700 mb-2">سطح دسترسی:</label>
              <select v-model="newOperator.accessLevel" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
                <option value="operator">اپراتور</option>
                <option value="manager">مدیر</option>
              </select>
            </div>

            <!-- دکمه‌های عملیات -->
            <div class="flex justify-between items-center mt-8">
              <button type="submit" class="bg-purple-600 text-white px-6 py-2 rounded-md hover:bg-purple-700 transition-colors">
                دعوت اپراتور
              </button>
              <button type="button" class="text-gray-600 hover:text-gray-800" @click="showInviteModal = false">
                انصراف
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';

definePageMeta({ layout: 'admin-main', middleware: 'admin' })

interface Operator {
  id: number | string;
  name: string;
  email: string;
  acceptedCount: number;
  totalHours: string;
  [key: string]: unknown;
}

interface NewOperator {
  title: string;
  email: string;
  password: string;
  fullName: string;
  phone: string;
  confirmPassword: string;
  websiteAccess: boolean;
  accessLevel: string;
}

// تب‌های موجود
const tabs = [
  { id: 'operators-list', name: 'لیست اپراتورها' },
  { id: 'operator-assignment', name: 'تخصیص اپراتور' },
  { id: 'link-assignment', name: 'تخصیص لینک به اپراتور' }
]

// تب فعال
const activeTab = ref('operators-list')

// نمایش مودال دعوت
const showInviteModal = ref(false)

// داده‌های اپراتور جدید
const newOperator = ref<NewOperator>({
  title: '',
  email: '',
  password: '',
  fullName: '',
  phone: '',
  confirmPassword: '',
  websiteAccess: true,
  accessLevel: 'operator'
})

// لیست اپراتورها از API
const operators = ref<Operator[]>([])
async function loadAvailableOperators() {
  try {
    const res = await $fetch<{ status: string; data: Record<string, unknown>[] }>('/api/admin/chat/operators/available')
    if (res?.status === 'success' && Array.isArray(res.data)) {
      operators.value = (res.data).map((op) => {
        const user = op.user as Record<string, unknown> | undefined
        return {
          id: op.id as number | string,
          name: (user?.name as string) || (user?.username as string) || `اپراتور ${op.id}`,
          email: (user?.email as string) || '-',
          acceptedCount: (op.current_chats as number) || 0,
          totalHours: '۰:۰۰',
        }
      })
    }
  } catch (e) {
    // console.error('loadAvailableOperators failed', e)
  }
}

// متغیرهای تخصیص اپراتور
const operatorSearch = ref('')
const selectedOperators = ref<(number | string)[]>([])

// فیلتر کردن اپراتورها بر اساس جستجو
const filteredOperators = computed(() => {
  if (!operatorSearch.value) return operators.value
  return operators.value.filter(operator => 
    operator.name.toLowerCase().includes(operatorSearch.value.toLowerCase())
  )
})

// بررسی اینکه آیا اپراتور انتخاب شده است
function isOperatorAssigned(operatorId: number | string) {
  return selectedOperators.value.includes(operatorId)
}

// تغییر وضعیت تخصیص اپراتور
function toggleOperatorAssignment(operatorId: number | string) {
  const index = selectedOperators.value.indexOf(operatorId)
  if (index > -1) {
    selectedOperators.value.splice(index, 1)
  } else {
    selectedOperators.value.push(operatorId)
  }
}

// ذخیره تخصیص اپراتور
function saveOperatorAssignment() {
  // اینجا منطق ذخیره تخصیص پیاده‌سازی می‌شود
  alert('تخصیص اپراتور با موفقیت ذخیره شد!')
}

// انصراف از تخصیص
function cancelOperatorAssignment() {
  selectedOperators.value = []
  operatorSearch.value = ''
}

// متغیرهای تخصیص لینک
const linkAssignment = ref({
  link: '',
  operators: [
    { operatorId: '' }
  ]
})

// اضافه کردن فیلد اپراتور جدید
function addOperatorField() {
  linkAssignment.value.operators.push({ operatorId: '' })
}

// حذف فیلد اپراتور
function removeOperatorField(index: number) {
  if (linkAssignment.value.operators.length > 1) {
    linkAssignment.value.operators.splice(index, 1)
  }
}

// ذخیره تخصیص لینک
function saveLinkAssignment() {
  // اینجا منطق ذخیره تخصیص لینک پیاده‌سازی می‌شود
  alert('تخصیص لینک با موفقیت ذخیره شد!')
}

// تابع دعوت اپراتور
async function inviteOperator() {
  try {
    const payload = {
      full_name: newOperator.value.fullName,
      username: newOperator.value.email?.split('@')[0] || newOperator.value.phone,
      email: newOperator.value.email,
      mobile: newOperator.value.phone,
      password: newOperator.value.password,
      max_concurrent_chats: 5,
      auto_accept: true,
    }
    const res = await $fetch<{ success: boolean; message?: string }>('/api/admin/chat/operators/invite', { method: 'POST', body: payload })
    if (res?.success) {
      await loadAvailableOperators()
      alert('اپراتور با موفقیت دعوت شد')
      showInviteModal.value = false
      newOperator.value = {
        title: '',
        email: '',
        password: '',
        fullName: '',
        phone: '',
        confirmPassword: '',
        websiteAccess: true,
        accessLevel: 'operator'
      }
    } else {
      alert(res?.message || 'خطا در دعوت اپراتور')
    }
  } catch (e) {
    console.error('inviteOperator failed', e)
    alert('خطا در دعوت اپراتور')
  }
}

onMounted(() => {
  loadAvailableOperators()
})
</script>
