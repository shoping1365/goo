<template>
  <div v-if="hasAccess" class="p-6 bg-white rounded-lg shadow-sm">
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900 mb-2">ویجت‌های نمایشی</h1>
      <p class="text-gray-600">مدیریت محتوای نمایشی در صفحات مختلف سایت</p>
    </div>

    <!-- تب‌های صفحات -->
    <div class="mb-6">
      <div class="border-b border-gray-200">
        <nav class="-mb-px flex space-x-8 space-x-reverse">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            :class="[
              'py-2 px-1 border-b-2 font-medium text-sm',
              activeTab === tab.id
                ? 'border-blue-500 text-blue-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
            @click="activeTab = tab.id"
          >
            {{ tab.name }}
          </button>
        </nav>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="space-y-6">
      <!-- تب صفحه محصول -->
      <div v-if="activeTab === 'product'" class="space-y-4">
        <div class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-semibold mb-3">صفحه محصول</h3>
          
          <!-- ویجت محصولات مرتبط -->
          <div class="mb-4">
            <label class="flex items-center">
              <input
                v-model="widgets.product.relatedProducts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm font-medium text-gray-700">نمایش محصولات مرتبط</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">نمایش محصولات مشابه در انتهای صفحه محصول</p>
          </div>

          <!-- ویجت محصولات پیشنهادی -->
          <div class="mb-4">
            <label class="flex items-center">
              <input
                v-model="widgets.product.recommendedProducts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm font-medium text-gray-700">محصولات پیشنهادی</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">نمایش محصولات پیشنهادی بر اساس رفتار کاربر</p>
          </div>

          <!-- ویجت نظرات -->
          <div class="mb-4">
            <label class="flex items-center">
              <input
                v-model="widgets.product.reviews"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm font-medium text-gray-700">بخش نظرات</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">نمایش نظرات و امتیازات کاربران</p>
          </div>
        </div>
      </div>

      <!-- تب سبد خرید -->
      <div v-if="activeTab === 'cart'" class="space-y-4">
        <div class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-semibold mb-3">سبد خرید</h3>
          
          <!-- ویجت محصولات پیشنهادی -->
          <div class="mb-4">
            <label class="flex items-center">
              <input
                v-model="widgets.cart.recommendedProducts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm font-medium text-gray-700">محصولات پیشنهادی</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">نمایش محصولات پیشنهادی در سبد خرید</p>
          </div>

          <!-- ویجت پیام هزینه ارسال -->
          <div class="mb-4">
            <label class="flex items-center">
              <input
                v-model="widgets.cart.shippingMessage"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm font-medium text-gray-700">پیام هزینه ارسال</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">نمایش اطلاعات هزینه ارسال</p>
          </div>

          <!-- ویجت کد تخفیف -->
          <div class="mb-4">
            <label class="flex items-center">
              <input
                v-model="widgets.cart.discountCode"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm font-medium text-gray-700">کد تخفیف</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">نمایش فیلد وارد کردن کد تخفیف</p>
          </div>
        </div>
      </div>

      <!-- تب صورتحساب -->
      <div v-if="activeTab === 'checkout'" class="space-y-4">
        <div class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-semibold mb-3">صورتحساب</h3>
          
          <!-- ویجت محصولات پیشنهادی -->
          <div class="mb-4">
            <label class="flex items-center">
              <input
                v-model="widgets.checkout.recommendedProducts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm font-medium text-gray-700">محصولات پیشنهادی</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">نمایش محصولات پیشنهادی در صفحه صورتحساب</p>
          </div>

          <!-- ویجت امنیت -->
          <div class="mb-4">
            <label class="flex items-center">
              <input
                v-model="widgets.checkout.securityBadge"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm font-medium text-gray-700">نشان امنیت</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">نمایش نشان‌های امنیت و اعتماد</p>
          </div>

          <!-- ویجت روش‌های پرداخت -->
          <div class="mb-4">
            <label class="flex items-center">
              <input
                v-model="widgets.checkout.paymentMethods"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm font-medium text-gray-700">روش‌های پرداخت</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">نمایش آیکون‌های روش‌های پرداخت</p>
          </div>
        </div>
      </div>

      <!-- تب صفحه اصلی -->
      <div v-if="activeTab === 'home'" class="space-y-4">
        <div class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-semibold mb-3">صفحه اصلی</h3>
          
          <!-- ویجت محصولات پرفروش -->
          <div class="mb-4">
            <label class="flex items-center">
              <input
                v-model="widgets.home.bestSellers"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm font-medium text-gray-700">محصولات پرفروش</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">نمایش محصولات پرفروش در صفحه اصلی</p>
          </div>

          <!-- ویجت محصولات جدید -->
          <div class="mb-4">
            <label class="flex items-center">
              <input
                v-model="widgets.home.newProducts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm font-medium text-gray-700">محصولات جدید</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">نمایش محصولات جدید در صفحه اصلی</p>
          </div>

          <!-- ویجت تخفیفات ویژه -->
          <div class="mb-4">
            <label class="flex items-center">
              <input
                v-model="widgets.home.specialOffers"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm font-medium text-gray-700">تخفیفات ویژه</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">نمایش تخفیفات ویژه در صفحه اصلی</p>
          </div>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="mt-8 flex justify-end space-x-3 space-x-reverse">
      <button
        class="px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        @click="resetSettings"
      >
        بازنشانی
      </button>
      <button
        class="px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        @click="saveSettings"
      >
        ذخیره تنظیمات
      </button>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, reactive, ref, watch } from 'vue'
import { useAuth } from '~/composables/useAuth'

// احراز هویت
const { user, isAuthenticated } = useAuth()

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false
  }

  const userRole = user.value?.role?.toLowerCase() || ''
  const adminRoles = ['admin', 'developer']
  return adminRoles.includes(userRole)
})

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async () => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false });
  }
};

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth();
});

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth();
  }
});

// تنظیمات صفحه
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// تب فعال
const activeTab = ref('product')

// تب‌های موجود
const tabs = [
  { id: 'product', name: 'صفحه محصول' },
  { id: 'cart', name: 'سبد خرید' },
  { id: 'checkout', name: 'صورتحساب' },
  { id: 'home', name: 'صفحه اصلی' }
]

// تنظیمات ویجت‌ها
const widgets = reactive({
  product: {
    relatedProducts: true,
    recommendedProducts: true,
    reviews: true
  },
  cart: {
    recommendedProducts: true,
    shippingMessage: false,
    discountCode: true
  },
  checkout: {
    recommendedProducts: false,
    securityBadge: true,
    paymentMethods: true
  },
  home: {
    bestSellers: true,
    newProducts: true,
    specialOffers: true
  }
})

// ذخیره تنظیمات
const saveSettings = async () => {
  try {
    // TODO: ارسال تنظیمات به سرور
    
    // نمایش پیام موفقیت
    alert('تنظیمات با موفقیت ذخیره شد')
  } catch (_error) {
    alert('خطا در ذخیره تنظیمات')
  }
}

// بازنشانی تنظیمات
const resetSettings = () => {
  if (confirm('آیا مطمئن هستید که می‌خواهید تنظیمات را بازنشانی کنید؟')) {
    // بازنشانی به حالت پیش‌فرض
    Object.assign(widgets.product, {
      relatedProducts: true,
      recommendedProducts: true,
      reviews: true
    })
    Object.assign(widgets.cart, {
      recommendedProducts: true,
      shippingMessage: false,
      discountCode: true
    })
    Object.assign(widgets.checkout, {
      recommendedProducts: false,
      securityBadge: true,
      paymentMethods: true
    })
    Object.assign(widgets.home, {
      bestSellers: true,
      newProducts: true,
      specialOffers: true
    })
  }
}
</script>

<style scoped>
/* استایل‌های اضافی در صورت نیاز */
</style>
