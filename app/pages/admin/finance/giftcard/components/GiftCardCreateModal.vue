<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg p-6 w-full max-w-4xl max-h-[90vh] overflow-y-auto">
      <div class="flex justify-between items-center mb-6">
        <h3 class="text-lg font-bold text-gray-900">ایجاد گیفت کارت جدید</h3>
        <button class="text-gray-400 hover:text-gray-600" @click="$emit('close')">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <form class="space-y-6" @submit.prevent="handleSubmit">
        <!-- تب‌های ایجاد -->
        <div class="border-b border-gray-200">
          <nav class="-mb-px flex space-x-8 space-x-reverse">
            <button
              type="button"
              :class="{
                'border-blue-500 text-blue-600': activeTab === 'basic',
                'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'basic'
              }"
              class="whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm"
              @click="activeTab = 'basic'"
            >
              اطلاعات پایه
            </button>
            <button
              type="button"
              :class="{
                'border-blue-500 text-blue-600': activeTab === 'design',
                'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'design'
              }"
              class="whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm"
              @click="activeTab = 'design'"
            >
              طراحی و ظاهر
            </button>
            <button
              type="button"
              :class="{
                'border-blue-500 text-blue-600': activeTab === 'recipient',
                'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'recipient'
              }"
              class="whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm"
              @click="activeTab = 'recipient'"
            >
              گیرنده و ارسال
            </button>
            <button
              type="button"
              :class="{
                'border-blue-500 text-blue-600': activeTab === 'advanced',
                'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'advanced'
              }"
              class="whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm"
              @click="activeTab = 'advanced'"
            >
              تنظیمات پیشرفته
            </button>
          </nav>
        </div>

        <!-- تب اطلاعات پایه -->
        <div v-if="activeTab === 'basic'" class="space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- نوع کارت -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع کارت *</label>
              <select v-model="formData.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
                <option value="digital">دیجیتال</option>
                <option value="physical">فیزیکی</option>
                <option value="hybrid">ترکیبی</option>
              </select>
            </div>

            <!-- دسته‌بندی کارت -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی *</label>
              <select v-model="formData.category" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
                <option value="birthday">تولد</option>
                <option value="wedding">عروسی</option>
                <option value="newyear">سال نو</option>
                <option value="corporate">شرکتی</option>
                <option value="discount">تخفیف ویژه</option>
                <option value="general">عمومی</option>
              </select>
            </div>

            <!-- مبلغ -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ (تومان) *</label>
              <div class="relative">
                <input 
                  v-model.number="formData.amount" 
                  type="number" 
                  min="1000" 
                  step="1000"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="مثال: 500000"
                />
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <span class="text-gray-500 text-sm">تومان</span>
                </div>
              </div>
            </div>

            <!-- نوع مبلغ -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع مبلغ</label>
              <select v-model="formData.amountType" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
                <option value="fixed">ثابت</option>
                <option value="variable">متغیر</option>
                <option value="range">محدوده</option>
              </select>
            </div>

            <!-- محدوده مبلغ (اگر متغیر باشد) -->
            <div v-if="formData.amountType === 'range'" class="md:col-span-2">
              <label class="block text-sm font-medium text-gray-700 mb-2">محدوده مبلغ</label>
              <div class="grid grid-cols-2 gap-6">
                <div>
                  <label class="block text-xs text-gray-600 mb-1">حداقل</label>
                  <input 
                    v-model.number="formData.minAmount" 
                    type="number" 
                    min="1000" 
                    step="1000"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  />
                </div>
                <div>
                  <label class="block text-xs text-gray-600 mb-1">حداکثر</label>
                  <input 
                    v-model.number="formData.maxAmount" 
                    type="number" 
                    min="1000" 
                    step="1000"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  />
                </div>
              </div>
            </div>

            <!-- تاریخ انقضا -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ انقضا *</label>
              <input 
                v-model="formData.expiryDate" 
                type="date" 
                :min="minExpiryDate"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>

            <!-- تعداد کارت -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تعداد کارت</label>
              <input 
                v-model.number="formData.quantity" 
                type="number" 
                min="1" 
                max="1000"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>
          </div>

          <!-- تولید خودکار کد -->
          <div class="flex items-center space-x-4 space-x-reverse">
            <input 
              id="autoGenerateCode" 
              v-model="formData.autoGenerateCode" 
              type="checkbox" 
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <label for="autoGenerateCode" class="block text-sm text-gray-900">تولید خودکار کد منحصر به فرد</label>
          </div>

          <!-- کد سفارشی -->
          <div v-if="!formData.autoGenerateCode">
            <label class="block text-sm font-medium text-gray-700 mb-2">کد سفارشی</label>
            <input 
              v-model="formData.customCode" 
              type="text" 
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: GC-BIRTHDAY-2024"
            />
          </div>
        </div>

        <!-- تب طراحی و ظاهر -->
        <div v-if="activeTab === 'design'" class="space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- انتخاب قالب -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">قالب کارت</label>
              <select v-model="formData.templateId" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
                <option value="">انتخاب قالب...</option>
                <option v-for="template in availableTemplates" :key="template.id" :value="template.id">
                  {{ template.name }}
                </option>
              </select>
            </div>

            <!-- رنگ اصلی -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">رنگ اصلی</label>
              <input 
                v-model="formData.primaryColor" 
                type="color" 
                class="w-full h-10 border border-gray-300 rounded-lg"
              />
            </div>

            <!-- رنگ ثانویه -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">رنگ ثانویه</label>
              <input 
                v-model="formData.secondaryColor" 
                type="color" 
                class="w-full h-10 border border-gray-300 rounded-lg"
              />
            </div>

            <!-- فونت -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">فونت</label>
              <select v-model="formData.fontFamily" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
                <option value="IRANSans">ایران سنس</option>
                <option value="Vazir">وزیر</option>
                <option value="Samim">صمیم</option>
                <option value="Shabnam">شبنم</option>
                <option value="Arial">Arial</option>
              </select>
            </div>
          </div>

          <!-- آپلود لوگو -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">لوگو (اختیاری)</label>
            <div class="mt-1 flex justify-center px-6 pt-5 pb-6 border-2 border-gray-300 border-dashed rounded-lg">
              <div class="space-y-1 text-center">
                <svg class="mx-auto h-12 w-12 text-gray-400" stroke="currentColor" fill="none" viewBox="0 0 48 48">
                  <path d="M28 8H12a4 4 0 00-4 4v20m32-12v8m0 0v8a4 4 0 01-4 4H12a4 4 0 01-4-4v-4m32-4l-3.172-3.172a4 4 0 00-5.656 0L28 28M8 32l9.172-9.172a4 4 0 015.656 0L28 28m0 0l4 4m4-24h8m-4-4v8m-12 4h.02" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
                </svg>
                <div class="flex text-sm text-gray-600">
                  <label for="logo-upload" class="relative cursor-pointer bg-white rounded-md font-medium text-blue-600 hover:text-blue-500 focus-within:outline-none focus-within:ring-2 focus-within:ring-offset-2 focus-within:ring-blue-500">
                    <span>آپلود فایل</span>
                    <input id="logo-upload" name="logo-upload" type="file" class="sr-only" accept="image/*" @change="handleLogoUpload" />
                  </label>
                  <p class="pr-1">یا فایل را اینجا بکشید</p>
                </div>
                <p class="text-xs text-gray-500">PNG, JPG تا 10MB</p>
              </div>
            </div>
          </div>

          <!-- پیش‌نمایش -->
          <div v-if="formData.templateId || formData.primaryColor" class="border rounded-lg p-6">
            <h4 class="text-sm font-medium text-gray-700 mb-3">پیش‌نمایش کارت</h4>
            <div class="w-64 h-40 mx-auto border rounded-lg relative overflow-hidden" :style="{ backgroundColor: formData.primaryColor || '#ffffff' }">
              <div class="absolute inset-0 flex items-center justify-center">
                <div class="text-center">
                  <div class="text-lg font-bold" :style="{ fontFamily: formData.fontFamily, color: formData.secondaryColor || '#000000' }">
                    گیفت کارت
                  </div>
                  <div class="text-sm mt-2" :style="{ fontFamily: formData.fontFamily, color: formData.secondaryColor || '#000000' }">
                    {{ formatAmount(formData.amount) }} تومان
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- تب گیرنده و ارسال -->
        <div v-if="activeTab === 'recipient'" class="space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- نام گیرنده -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نام گیرنده *</label>
              <input 
                v-model="formData.recipientName" 
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="نام و نام خانوادگی"
              />
            </div>

            <!-- ایمیل گیرنده -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ایمیل گیرنده *</label>
              <input 
                v-model="formData.recipientEmail" 
                type="email" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="example@email.com"
              />
            </div>

            <!-- شماره تلفن گیرنده -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">شماره تلفن گیرنده</label>
              <input 
                v-model="formData.recipientPhone" 
                type="tel" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="09123456789"
              />
            </div>

            <!-- روش ارسال -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">روش ارسال</label>
              <select v-model="formData.deliveryMethod" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
                <option value="email">ایمیل</option>
                <option value="sms">پیامک</option>
                <option value="both">ایمیل و پیامک</option>
                <option value="manual">ارسال دستی</option>
              </select>
            </div>
          </div>

          <!-- پیام شخصی -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">پیام شخصی</label>
            <textarea 
              v-model="formData.personalMessage" 
              rows="4" 
              maxlength="500"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="پیام شخصی خود را اینجا بنویسید..."
            ></textarea>
            <div class="text-xs text-gray-500 mt-1">
              {{ formData.personalMessage.length }}/500 کاراکتر
            </div>
          </div>

          <!-- تاریخ ارسال -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ ارسال</label>
            <input 
              v-model="formData.deliveryDate" 
              type="datetime-local" 
              :min="minDeliveryDate"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
        </div>

        <!-- تب تنظیمات پیشرفته -->
        <div v-if="activeTab === 'advanced'" class="space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- محدودیت استفاده -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">محدودیت استفاده</label>
              <select v-model="formData.usageLimit" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
                <option value="unlimited">نامحدود</option>
                <option value="once">یک بار</option>
                <option value="multiple">چند بار</option>
              </select>
            </div>

            <!-- تعداد استفاده مجاز -->
            <div v-if="formData.usageLimit === 'multiple'">
              <label class="block text-sm font-medium text-gray-700 mb-2">تعداد استفاده مجاز</label>
              <input 
                v-model.number="formData.maxUsageCount" 
                type="number" 
                min="2" 
                max="10"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>

            <!-- حداقل مبلغ سفارش -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ سفارش</label>
              <input 
                v-model.number="formData.minOrderAmount" 
                type="number" 
                min="0" 
                step="1000"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="0 = بدون محدودیت"
              />
            </div>

            <!-- حداکثر مبلغ سفارش -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر مبلغ سفارش</label>
              <input 
                v-model.number="formData.maxOrderAmount" 
                type="number" 
                min="0" 
                step="1000"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="0 = بدون محدودیت"
              />
            </div>
          </div>

          <!-- دسته‌بندی‌های مجاز -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی‌های مجاز</label>
            <div class="grid grid-cols-2 md:grid-cols-4 gap-2">
              <div v-for="category in availableCategories" :key="category.id" class="flex items-center">
                <input 
                  :id="'category-' + category.id" 
                  v-model="formData.allowedCategories"
                  type="checkbox" 
                  :value="category.id"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label :for="'category-' + category.id" class="mr-2 block text-sm text-gray-900">
                  {{ category.name }}
                </label>
              </div>
            </div>
          </div>

          <!-- تنظیمات امنیتی -->
          <div class="space-y-3">
            <h4 class="text-sm font-medium text-gray-700">تنظیمات امنیتی</h4>
            <div class="space-y-2">
              <div class="flex items-center">
                <input 
                  id="requireVerification" 
                  v-model="formData.requireVerification" 
                  type="checkbox" 
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label for="requireVerification" class="mr-2 block text-sm text-gray-900">نیاز به تأیید هویت</label>
              </div>
              <div class="flex items-center">
                <input 
                  id="allowPartialUsage" 
                  v-model="formData.allowPartialUsage" 
                  type="checkbox" 
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label for="allowPartialUsage" class="mr-2 block text-sm text-gray-900">اجازه استفاده جزئی</label>
              </div>
              <div class="flex items-center">
                <input 
                  id="autoRenew" 
                  v-model="formData.autoRenew" 
                  type="checkbox" 
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label for="autoRenew" class="mr-2 block text-sm text-gray-900">تمدید خودکار</label>
              </div>
            </div>
          </div>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex justify-end space-x-3 space-x-reverse pt-6 border-t border-gray-200">
          <button 
            type="button" 
            class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            @click="$emit('close')"
          >
            انصراف
          </button>
          <button 
            v-if="activeTab !== 'basic'" 
            type="button"
            class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            @click="previousTab"
          >
            قبلی
          </button>
          <button 
            v-if="activeTab !== 'advanced'" 
            type="button"
            class="px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            @click="nextTab"
          >
            بعدی
          </button>
          <button 
            v-if="activeTab === 'advanced'" 
            type="submit"
            :disabled="isSubmitting"
            class="px-6 py-2 text-sm font-medium text-white bg-green-600 border border-transparent rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 disabled:opacity-50"
          >
            {{ isSubmitting ? 'در حال ایجاد...' : 'ایجاد گیفت کارت' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'

// Props
const props = defineProps<{
  giftCard?: any
}>()

// Emits
const emit = defineEmits<{
  close: []
  created: [giftCard: any]
}>()

// Reactive data
const activeTab = ref('basic')
const isSubmitting = ref(false)

// فرم داده‌ها
const formData = ref({
  // اطلاعات پایه
  type: 'digital',
  category: 'general',
  amount: 500000,
  amountType: 'fixed',
  minAmount: 0,
  maxAmount: 0,
  expiryDate: '',
  quantity: 1,
  autoGenerateCode: true,
  customCode: '',

  // طراحی و ظاهر
  templateId: '',
  primaryColor: '#3B82F6',
  secondaryColor: '#1F2937',
  fontFamily: 'IRANSans',
  logo: null,

  // گیرنده و ارسال
  recipientName: '',
  recipientEmail: '',
  recipientPhone: '',
  deliveryMethod: 'email',
  personalMessage: '',
  deliveryDate: '',

  // تنظیمات پیشرفته
  usageLimit: 'unlimited',
  maxUsageCount: 1,
  minOrderAmount: 0,
  maxOrderAmount: 0,
  allowedCategories: [],
  requireVerification: false,
  allowPartialUsage: true,
  autoRenew: false
})

// داده‌های کمکی
const availableTemplates = ref([
  { id: 1, name: 'قالب تولد' },
  { id: 2, name: 'قالب عروسی' },
  { id: 3, name: 'قالب سال نو' },
  { id: 4, name: 'قالب شرکتی' },
  { id: 5, name: 'قالب عمومی' }
])

const availableCategories = ref([
  { id: 'electronics', name: 'الکترونیک' },
  { id: 'clothing', name: 'پوشاک' },
  { id: 'books', name: 'کتاب' },
  { id: 'home', name: 'خانه و آشپزخانه' },
  { id: 'sports', name: 'ورزشی' },
  { id: 'beauty', name: 'زیبایی' },
  { id: 'toys', name: 'اسباب بازی' },
  { id: 'food', name: 'مواد غذایی' }
])

// Computed properties
const minExpiryDate = computed(() => {
  const today = new Date()
  today.setDate(today.getDate() + 1)
  return today.toISOString().split('T')[0]
})

const minDeliveryDate = computed(() => {
  const now = new Date()
  now.setMinutes(now.getMinutes() + 30)
  return now.toISOString().slice(0, 16)
})

// Methods
const formatAmount = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount)
}

const handleLogoUpload = (event: Event) => {
  const file = (event.target as HTMLInputElement).files?.[0]
  if (file) {
    formData.value.logo = file
  }
}

const generateUniqueCode = () => {
  const prefix = 'GC'
  const timestamp = Date.now().toString(36)
  const random = Math.random().toString(36).substr(2, 5)
  return `${prefix}-${timestamp}-${random}`.toUpperCase()
}

const nextTab = () => {
  const tabs = ['basic', 'design', 'recipient', 'advanced']
  const currentIndex = tabs.indexOf(activeTab.value)
  if (currentIndex < tabs.length - 1) {
    activeTab.value = tabs[currentIndex + 1]
  }
}

const previousTab = () => {
  const tabs = ['basic', 'design', 'recipient', 'advanced']
  const currentIndex = tabs.indexOf(activeTab.value)
  if (currentIndex > 0) {
    activeTab.value = tabs[currentIndex - 1]
  }
}

const validateForm = () => {
  const errors = []

  // اعتبارسنجی اطلاعات پایه
  if (!formData.value.amount || formData.value.amount < 1000) {
    errors.push('مبلغ باید حداقل 1000 تومان باشد')
  }

  if (formData.value.amountType === 'range') {
    if (!formData.value.minAmount || !formData.value.maxAmount) {
      errors.push('حداقل و حداکثر مبلغ برای محدوده الزامی است')
    }
    if (formData.value.minAmount >= formData.value.maxAmount) {
      errors.push('حداقل مبلغ باید کمتر از حداکثر باشد')
    }
  }

  if (!formData.value.expiryDate) {
    errors.push('تاریخ انقضا الزامی است')
  }

  if (!formData.value.autoGenerateCode && !formData.value.customCode) {
    errors.push('کد سفارشی الزامی است')
  }

  // اعتبارسنجی گیرنده
  if (!formData.value.recipientName) {
    errors.push('نام گیرنده الزامی است')
  }

  if (!formData.value.recipientEmail) {
    errors.push('ایمیل گیرنده الزامی است')
  }

  return errors
}

const handleSubmit = async () => {
  const errors = validateForm()
  if (errors.length > 0) {
    alert('خطاهای زیر را برطرف کنید:\n' + errors.join('\n'))
    return
  }

  isSubmitting.value = true

  try {
    // تولید کد منحصر به فرد
    if (formData.value.autoGenerateCode) {
      formData.value.customCode = generateUniqueCode()
    }

    // ایجاد گیفت کارت
    const giftCard = {
      id: Date.now(),
      code: formData.value.customCode,
      amount: formData.value.amount,
      remainingAmount: formData.value.amount,
      status: 'active',
      type: formData.value.type,
      category: formData.value.category,
      templateId: formData.value.templateId,
      recipientName: formData.value.recipientName,
      recipientEmail: formData.value.recipientEmail,
      recipientPhone: formData.value.recipientPhone,
      personalMessage: formData.value.personalMessage,
      deliveryMethod: formData.value.deliveryMethod,
      deliveryDate: formData.value.deliveryDate,
      expiryDate: formData.value.expiryDate,
      usageLimit: formData.value.usageLimit,
      maxUsageCount: formData.value.maxUsageCount,
      minOrderAmount: formData.value.minOrderAmount,
      maxOrderAmount: formData.value.maxOrderAmount,
      allowedCategories: formData.value.allowedCategories,
      requireVerification: formData.value.requireVerification,
      allowPartialUsage: formData.value.allowPartialUsage,
      autoRenew: formData.value.autoRenew,
      createdAt: new Date(),
      createdBy: 'کاربر فعلی',
      isUsed: false,
      usageHistory: []
    }

    // در نسخه واقعی: ارسال به API
    console.log('Creating gift card:', giftCard)

    // شبیه‌سازی تأخیر
    await new Promise(resolve => setTimeout(resolve, 1000))

    emit('created', giftCard)
    emit('close')
  } catch (error) {
    console.error('Error creating gift card:', error)
    alert('خطا در ایجاد گیفت کارت')
  } finally {
    isSubmitting.value = false
  }
}

// Lifecycle
onMounted(() => {
  if (props.giftCard) {
    // ویرایش گیفت کارت موجود
    formData.value = { ...props.giftCard }
  }
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
