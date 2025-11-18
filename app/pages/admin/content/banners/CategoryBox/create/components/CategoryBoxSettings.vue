<template>
  <div class="widget-settings-panel bg-white rounded-lg shadow mb-6">
    <div class="px-6 py-4 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <h2 class="text-lg font-medium text-gray-900">تنظیمات ویجت</h2>
        <div class="flex items-center space-x-8">
          <div class="flex items-center">
            <input
              v-model="config.show_title"
              type="checkbox"
              class="h-4 w-4 text-purple-600 focus:ring-purple-500 border-gray-300 rounded"
            />
            <label class="mr-2 text-sm text-gray-700">نمایش عنوان بخش</label>
          </div>
          <div class="flex items-center">
            <input
              v-model="config.show_product_count"
              type="checkbox"
              class="h-4 w-4 text-purple-600 focus:ring-purple-500 border-gray-300 rounded"
            />
            <label class="mr-2 text-sm text-gray-700">نمایش تعداد محصولات</label>
          </div>
        </div>
      </div>
    </div>
    <div class="p-6">
      <div class="grid grid-cols-1 lg:grid-cols-4 gap-6">
        <!-- اطلاعات پایه -->
        <div class="space-y-3">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">عنوان ویجت</label>
            <input
              v-model="formData.title"
              type="text"
              placeholder="عنوان ویجت را وارد کنید"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
              required
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">صفحه نمایش</label>
            <select
              v-model="formData.page"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            >
              <option value="home">صفحه اصلی</option>
              <option value="category">صفحه دسته‌بندی</option>
              <option value="product">صفحه محصول</option>
              <option value="other">سایر صفحات</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاصله داخلی (Padding)</label>
            <input
              v-model="config.padding"
              type="text"
              placeholder="مثال: 24px"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاصله داخلی بالا (Padding Top)</label>
            <input
              v-model="config.padding_top"
              type="text"
              placeholder="مثال: 12px"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            />
          </div>
          <div v-if="config.show_background">
            <label class="block text-sm font-medium text-gray-700 mb-2">رنگ پس‌زمینه</label>
            <input
              v-model="config.background_color"
              type="color"
              class="w-full h-10 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            />
          </div>

          <div v-if="config.show_border">
            <label class="block text-sm font-medium text-gray-700 mb-2">رنگ خط</label>
            <input
              v-model="config.border_color"
              type="color"
              class="w-full h-10 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            />
          </div>
        </div>

        <!-- تنظیمات عنوان و محتوا -->
        <div class="space-y-3">
          <div v-if="config.show_title">
            <label class="block text-sm font-medium text-gray-700 mb-2">عنوان بخش</label>
            <input
              v-model="config.title"
              type="text"
              placeholder="عنوان بخش را وارد کنید"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">منبع دسته‌بندی‌ها</label>
            <select
              v-model="config.category_source"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            >
              <option value="manual">دلخواه (انتخاب دستی)</option>
              <option value="best_selling">پرفروش‌ترین دسته‌بندی‌ها</option>
              <option value="user_search">بر اساس سرچ کاربر</option>
              <option value="user_interests">بر اساس علایق کاربر</option>
            </select>
          </div>
          
          <div v-if="config.show_title">
            <label class="block text-sm font-medium text-gray-700 mb-2">رنگ عنوان</label>
            <input
              v-model="config.title_color"
              type="color"
              class="w-full h-10 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            />
          </div>
          
          <div v-if="config.show_title">
            <label class="block text-sm font-medium text-gray-700 mb-2">محل قرار گیری عنوان</label>
            <select
              v-model="config.title_alignment"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            >
              <option value="left">چپ</option>
              <option value="center">وسط</option>
              <option value="right">راست</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاصله خارجی (Margin)</label>
            <input
              v-model="config.margin"
              type="text"
              placeholder="مثال: 0px"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">گردی گوشه‌های باکس</label>
            <select
              v-model="config.border_radius"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            >
              <option value="none">بدون گردی</option>
              <option value="small">گردی کم</option>
              <option value="medium" selected>گردی متوسط</option>
              <option value="large">گردی زیاد</option>
            </select>
          </div>

          
          <div v-if="!config.show_title">
            <label class="block text-sm font-medium text-gray-700 mb-2">فاصله داخلی پایین (Padding Bottom)</label>
            <input
              v-model="config.padding_bottom"
              type="text"
              placeholder="مثال: 12px"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            />
          </div>
          
          <div v-if="!config.show_title">
            <label class="block text-sm font-medium text-gray-700 mb-2">تراز متن</label>
            <select
              v-model="config.text_alignment"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            >
              <option value="left">چپ</option>
              <option value="center">وسط</option>
              <option value="right">راست</option>
            </select>
          </div>
        </div>

        <!-- تنظیمات چیدمان -->
        <div class="space-y-3">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">چیدمان</label>
            <select
              v-model="config.layout"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            >
              <option value="grid">شبکه‌ای</option>
              <option value="list">لیستی</option>
            </select>
          </div>

          <div v-if="config.layout === 'grid'">
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد ستون‌ها</label>
            <select
              v-model="config.columns"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            >
              <option :value="1">1 ستون</option>
              <option :value="2">2 ستون</option>
              <option :value="3">3 ستون</option>
              <option :value="4">4 ستون</option>
              <option :value="5">5 ستون</option>
              <option :value="6">6 ستون</option>
            </select>
          </div>

          <div v-if="config.layout === 'grid'">
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد ستون در موبایل</label>
            <select
              v-model="config.mobile_columns"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            >
              <option :value="1">1 ستون</option>
              <option :value="2">2 ستون</option>
              <option :value="3">3 ستون</option>
              <option :value="4">4 ستون</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تراز متن</label>
            <select
              v-model="config.text_alignment"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            >
              <option value="left">چپ</option>
              <option value="center">وسط</option>
              <option value="right">راست</option>
            </select>
          </div>

          <div v-if="config.show_border">
            <label class="block text-sm font-medium text-gray-700 mb-2">ضخامت خط</label>
            <select
              v-model="config.border_width"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 text-sm"
            >
              <option value="thin">نازک</option>
              <option value="medium" selected>متوسط</option>
              <option value="thick">ضخیم</option>
            </select>
          </div>

          <div v-if="config.show_background">
            <label class="block text-sm font-medium text-gray-700 mb-2">پس زمینه تمام عرض</label>
            <select
              v-model="config.full_width_background"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 text-sm"
            >
              <option :value="false">خاموش</option>
              <option :value="true">روشن</option>
            </select>
          </div>

        </div>

        <!-- تنظیمات ظاهری و رنگ -->
        <div class="space-y-3">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">چیدمان عمودی</label>
            <select
              v-model="config.vertical_layout"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 text-sm"
            >
              <option :value="false">خاموش</option>
              <option :value="true">روشن</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">استایل کارت</label>
            <select
              v-model="config.card_style"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            >
              <option value="default">پیش‌فرض</option>
              <option value="rounded">گرد</option>
              <option value="shadow">سایه‌دار</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">اندازه تصویر</label>
            <select
              v-model="config.image_size"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            >
              <option value="small">کوچک</option>
              <option value="medium">متوسط</option>
              <option value="large">بزرگ</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">پس زمینه فعال</label>
            <select
              v-model="config.show_background"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 text-sm"
            >
              <option :value="false">خاموش</option>
              <option :value="true">روشن</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">خط دور دسته‌بندی‌ها</label>
            <select
              v-model="config.show_border"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 text-sm"
            >
              <option :value="false">خاموش</option>
              <option :value="true">روشن</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">اندازه باکس</label>
            <select
              v-model="config.box_width"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            >
              <option value="full">تمام عرض</option>
              <option value="center">وسط صفحه</option>
            </select>
          </div>

        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { CategoryConfig } from '~/types/widget'

interface Props {
  formData: {
    title: string
    page: string
  }
  config: CategoryConfig
}

defineProps<Props>()
</script>
