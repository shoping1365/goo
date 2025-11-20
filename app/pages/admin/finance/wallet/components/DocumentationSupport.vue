<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">مستندات و پشتیبانی</h2>
          <p class="text-gray-600 mt-1">راهنما، سوالات متداول و پشتیبانی فنی کیف پول</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'guide'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'guide'"
          >
            راهنما
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'faq'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'faq'"
          >
            سوالات متداول
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'support'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'support'"
          >
            پشتیبانی
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'api'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'api'"
          >
            مستندات API
          </button>
        </div>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <!-- راهنما -->
      <div v-if="activeTab === 'guide'" class="p-6">
        <div class="space-y-6">
          <!-- راهنمای شروع -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">راهنمای شروع کار</h3>
            <div class="space-y-4">
              <div class="flex items-start space-x-4 space-x-reverse">
                <div class="flex-shrink-0 w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                  <span class="text-blue-600 font-semibold text-sm">1</span>
                </div>
                <div>
                  <h4 class="font-medium text-gray-900">تنظیم کیف پول</h4>
                  <p class="text-gray-600 text-sm mt-1">
                    ابتدا تنظیمات اولیه کیف پول را انجام دهید. حداقل موجودی، کارمزدها و محدودیت‌ها را تعیین کنید.
                  </p>
                </div>
              </div>
              <div class="flex items-start space-x-4 space-x-reverse">
                <div class="flex-shrink-0 w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                  <span class="text-blue-600 font-semibold text-sm">2</span>
                </div>
                <div>
                  <h4 class="font-medium text-gray-900">اتصال به درگاه‌های پرداخت</h4>
                  <p class="text-gray-600 text-sm mt-1">
                    درگاه‌های پرداخت مورد نیاز را متصل کنید و تنظیمات امنیتی را انجام دهید.
                  </p>
                </div>
              </div>
              <div class="flex items-start space-x-4 space-x-reverse">
                <div class="flex-shrink-0 w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                  <span class="text-blue-600 font-semibold text-sm">3</span>
                </div>
                <div>
                  <h4 class="font-medium text-gray-900">مدیریت کاربران</h4>
                  <p class="text-gray-600 text-sm mt-1">
                    دسترسی کاربران به کیف پول را مدیریت کنید و نقش‌های مختلف را تعریف کنید.
                  </p>
                </div>
              </div>
            </div>
          </div>

          <!-- راهنمای امنیت -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">راهنمای امنیت</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="bg-red-50 border border-red-200 rounded-lg p-6">
                <h4 class="font-medium text-red-800 mb-2">نکات مهم امنیتی</h4>
                <ul class="text-sm text-red-700 space-y-1">
                  <li>• رمزهای عبور قوی استفاده کنید</li>
                  <li>• احراز هویت دو مرحله‌ای فعال کنید</li>
                  <li>• دسترسی‌ها را محدود کنید</li>
                  <li>• لاگ‌ها را بررسی کنید</li>
                </ul>
              </div>
              <div class="bg-green-50 border border-green-200 rounded-lg p-6">
                <h4 class="font-medium text-green-800 mb-2">بهترین شیوه‌ها</h4>
                <ul class="text-sm text-green-700 space-y-1">
                  <li>• پشتیبان‌گیری منظم</li>
                  <li>• به‌روزرسانی سیستم</li>
                  <li>• نظارت بر تراکنش‌ها</li>
                  <li>• آموزش کارکنان</li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- سوالات متداول -->
      <div v-if="activeTab === 'faq'" class="p-6">
        <div class="space-y-4">
          <div
            v-for="(faq, index) in faqs"
            :key="index"
            class="border border-gray-200 rounded-lg"
          >
            <button
              class="w-full px-6 py-4 text-right flex items-center justify-between hover:bg-gray-50 transition-colors"
              @click="toggleFaq(index)"
            >
              <span class="font-medium text-gray-900">{{ faq.question }}</span>
              <svg
                :class="[
                  'w-5 h-5 text-gray-500 transition-transform',
                  faq.isOpen ? 'rotate-180' : ''
                ]"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
              </svg>
            </button>
            <div
              v-show="faq.isOpen"
              class="px-6 pb-4 text-gray-600 text-sm leading-relaxed"
            >
              {{ faq.answer }}
            </div>
          </div>
        </div>
      </div>

      <!-- پشتیبانی -->
      <div v-if="activeTab === 'support'" class="p-6">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- فرم تماس -->
          <div class="space-y-4">
            <h3 class="text-lg font-semibold text-gray-900">تماس با پشتیبانی</h3>
            <form class="space-y-4" @submit.prevent="submitSupportForm">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">موضوع</label>
                <select
                  v-model="supportForm.subject"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                >
                  <option value="">انتخاب کنید</option>
                  <option value="technical">مشکل فنی</option>
                  <option value="billing">مشکل مالی</option>
                  <option value="security">مشکل امنیتی</option>
                  <option value="general">سوال عمومی</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">عنوان</label>
                <input
                  v-model="supportForm.title"
                  type="text"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="عنوان مشکل یا سوال"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
                <textarea
                  v-model="supportForm.description"
                  rows="4"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="جزئیات مشکل یا سوال خود را بنویسید..."
                ></textarea>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">اولویت</label>
                <select
                  v-model="supportForm.priority"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                >
                  <option value="low">کم</option>
                  <option value="medium">متوسط</option>
                  <option value="high">زیاد</option>
                  <option value="urgent">فوری</option>
                </select>
              </div>
              <button
                type="submit"
                class="w-full bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 transition-colors"
              >
                ارسال درخواست پشتیبانی
              </button>
            </form>
          </div>

          <!-- اطلاعات تماس -->
          <div class="space-y-4">
            <h3 class="text-lg font-semibold text-gray-900">اطلاعات تماس</h3>
            <div class="space-y-4">
              <div class="flex items-center space-x-3 space-x-reverse">
                <div class="w-10 h-10 bg-blue-100 rounded-full flex items-center justify-center">
                  <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z" />
                  </svg>
                </div>
                <div>
                  <p class="font-medium text-gray-900">تلفن پشتیبانی</p>
                  <p class="text-gray-600">021-12345678</p>
                </div>
              </div>
              <div class="flex items-center space-x-3 space-x-reverse">
                <div class="w-10 h-10 bg-green-100 rounded-full flex items-center justify-center">
                  <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                  </svg>
                </div>
                <div>
                  <p class="font-medium text-gray-900">ایمیل پشتیبانی</p>
                  <p class="text-gray-600">support@example.com</p>
                </div>
              </div>
              <div class="flex items-center space-x-3 space-x-reverse">
                <div class="w-10 h-10 bg-purple-100 rounded-full flex items-center justify-center">
                  <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
                  </svg>
                </div>
                <div>
                  <p class="font-medium text-gray-900">چت آنلاین</p>
                  <p class="text-gray-600">24/7 در دسترس</p>
                </div>
              </div>
            </div>

            <!-- ساعات کاری -->
            <div class="bg-gray-50 rounded-lg p-6">
              <h4 class="font-medium text-gray-900 mb-2">ساعات کاری پشتیبانی</h4>
              <div class="space-y-1 text-sm text-gray-600">
                <p>شنبه تا چهارشنبه: 8:00 - 18:00</p>
                <p>پنجشنبه: 8:00 - 14:00</p>
                <p>جمعه: تعطیل</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- مستندات API -->
      <div v-if="activeTab === 'api'" class="p-6">
        <div class="space-y-6">
          <!-- معرفی API -->
          <div class="bg-blue-50 border border-blue-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-blue-900 mb-2">مستندات API کیف پول</h3>
            <p class="text-blue-800 text-sm">
              این API برای مدیریت کیف پول، تراکنش‌ها و گزارش‌گیری طراحی شده است. تمام درخواست‌ها باید شامل توکن احراز هویت باشند.
            </p>
          </div>

          <!-- نقاط پایانی API -->
          <div class="space-y-4">
            <h3 class="text-lg font-semibold text-gray-900">نقاط پایانی اصلی</h3>
            
            <!-- دریافت موجودی -->
            <div class="border border-gray-200 rounded-lg p-6">
              <div class="flex items-center justify-between mb-2">
                <h4 class="font-medium text-gray-900">دریافت موجودی کیف پول</h4>
                <span class="px-2 py-1 bg-green-100 text-green-800 text-xs rounded">GET</span>
              </div>
              <code class="text-sm bg-gray-100 px-2 py-1 rounded">/api/wallet/balance/{wallet_id}</code>
              <p class="text-gray-600 text-sm mt-2">دریافت موجودی فعلی کیف پول</p>
            </div>

            <!-- ایجاد تراکنش -->
            <div class="border border-gray-200 rounded-lg p-6">
              <div class="flex items-center justify-between mb-2">
                <h4 class="font-medium text-gray-900">ایجاد تراکنش جدید</h4>
                <span class="px-2 py-1 bg-blue-100 text-blue-800 text-xs rounded">POST</span>
              </div>
              <code class="text-sm bg-gray-100 px-2 py-1 rounded">/api/wallet/transactions</code>
              <p class="text-gray-600 text-sm mt-2">ایجاد تراکنش جدید (شارژ، برداشت، انتقال)</p>
            </div>

            <!-- دریافت تراکنش‌ها -->
            <div class="border border-gray-200 rounded-lg p-6">
              <div class="flex items-center justify-between mb-2">
                <h4 class="font-medium text-gray-900">دریافت لیست تراکنش‌ها</h4>
                <span class="px-2 py-1 bg-green-100 text-green-800 text-xs rounded">GET</span>
              </div>
              <code class="text-sm bg-gray-100 px-2 py-1 rounded">/api/wallet/transactions</code>
              <p class="text-gray-600 text-sm mt-2">دریافت لیست تراکنش‌ها با فیلتر و صفحه‌بندی</p>
            </div>

            <!-- گزارش‌گیری -->
            <div class="border border-gray-200 rounded-lg p-6">
              <div class="flex items-center justify-between mb-2">
                <h4 class="font-medium text-gray-900">گزارش‌گیری مالی</h4>
                <span class="px-2 py-1 bg-green-100 text-green-800 text-xs rounded">GET</span>
              </div>
              <code class="text-sm bg-gray-100 px-2 py-1 rounded">/api/wallet/reports</code>
              <p class="text-gray-600 text-sm mt-2">دریافت گزارش‌های مالی و آماری</p>
            </div>
          </div>

          <!-- نمونه کد -->
          <div class="space-y-4">
            <h3 class="text-lg font-semibold text-gray-900">نمونه کد</h3>
            <div class="bg-gray-900 rounded-lg p-6">
              <pre class="text-green-400 text-sm overflow-x-auto"><code>// نمونه درخواست دریافت موجودی
const response = await fetch('/api/wallet/balance/123', {
  method: 'GET',
  headers: {
    'Authorization': 'Bearer YOUR_TOKEN',
    'Content-Type': 'application/json'
  }
});

const balance = await response.json();
console.log('موجودی:', balance.amount);</code></pre>
            </div>
          </div>

          <!-- کدهای خطا -->
          <div class="space-y-4">
            <h3 class="text-lg font-semibold text-gray-900">کدهای خطا</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="border border-red-200 rounded-lg p-6">
                <h4 class="font-medium text-red-800 mb-2">خطاهای احراز هویت</h4>
                <ul class="text-sm text-red-700 space-y-1">
                  <li>401 - توکن نامعتبر</li>
                  <li>403 - دسترسی غیرمجاز</li>
                  <li>429 - محدودیت نرخ درخواست</li>
                </ul>
              </div>
              <div class="border border-orange-200 rounded-lg p-6">
                <h4 class="font-medium text-orange-800 mb-2">خطاهای تراکنش</h4>
                <ul class="text-sm text-orange-700 space-y-1">
                  <li>400 - داده‌های نامعتبر</li>
                  <li>422 - موجودی ناکافی</li>
                  <li>500 - خطای سرور</li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// تعریف reactive data
const activeTab = ref('guide')

// داده‌های سوالات متداول
const faqs = ref([
  {
    question: 'چگونه کیف پول جدید ایجاد کنم؟',
    answer: 'برای ایجاد کیف پول جدید، به بخش تنظیمات کیف پول بروید و روی دکمه "ایجاد کیف پول جدید" کلیک کنید. سپس اطلاعات مورد نیاز را وارد کنید.',
    isOpen: false
  },
  {
    question: 'آیا می‌توانم موجودی کیف پول را به صورت دستی تغییر دهم؟',
    answer: 'بله، در بخش مدیریت دستی کیف پول می‌توانید موجودی را به صورت دستی شارژ یا کسر کنید. این عملیات در لاگ‌ها ثبت می‌شود.',
    isOpen: false
  },
  {
    question: 'چگونه تراکنش مشکوک را گزارش دهم؟',
    answer: 'برای گزارش تراکنش مشکوک، به بخش امنیت بروید و روی "گزارش تراکنش مشکوک" کلیک کنید. سپس جزئیات را وارد کنید.',
    isOpen: false
  },
  {
    question: 'آیا امکان اتصال به سیستم‌های حسابداری وجود دارد؟',
    answer: 'بله، در بخش یکپارچه‌سازی می‌توانید کیف پول را به سیستم‌های حسابداری مختلف متصل کنید.',
    isOpen: false
  },
  {
    question: 'چگونه گزارش‌های مالی را دریافت کنم؟',
    answer: 'در بخش گزارش‌های مالی می‌توانید انواع مختلف گزارش‌ها را مشاهده و دانلود کنید.',
    isOpen: false
  }
])

// فرم پشتیبانی
const supportForm = ref({
  subject: '',
  title: '',
  description: '',
  priority: 'medium'
})

// توابع
const toggleFaq = (index: number) => {
  faqs.value[index].isOpen = !faqs.value[index].isOpen
}

const submitSupportForm = () => {
  // ارسال فرم پشتیبانی
  console.log('فرم پشتیبانی ارسال شد:', supportForm.value)
  
  // نمایش پیام موفقیت
  alert('درخواست پشتیبانی شما با موفقیت ارسال شد.')
  
  // پاک کردن فرم
  supportForm.value = {
    subject: '',
    title: '',
    description: '',
    priority: 'medium'
  }
}
</script> 
