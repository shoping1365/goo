<template>
  <div class="space-y-6">
    <!-- هدر صفحه -->
    <div class="bg-white rounded-lg shadow-sm p-6">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-gray-900">ایجاد صفحه جدید</h1>
          <p class="text-gray-600 mt-1">صفحه جدید خود را ایجاد کنید</p>
        </div>
        <div class="flex space-x-3 space-x-reverse">
          <TemplateButton
            bg-gradient="bg-gradient-to-r from-gray-400 to-gray-600"
            text-color="text-white"
            hover-class="hover:from-gray-500 hover:to-gray-700 hover:shadow-lg hover:scale-105"
            focus-class="focus:ring-2 focus:ring-offset-2 focus:ring-gray-500"
            size="medium"
            @click="saveAsDraft"
          >
            ذخیره پیش‌نویس
          </TemplateButton>
          <TemplateButton
            bg-gradient="bg-gradient-to-r from-blue-400 to-blue-600"
            text-color="text-white"
            hover-class="hover:from-blue-500 hover:to-blue-700 hover:shadow-lg hover:scale-105"
            focus-class="focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            size="medium"
            @click="publishPage"
          >
            انتشار صفحه
          </TemplateButton>
        </div>
      </div>
    </div>

    <!-- فرم ایجاد صفحه -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- ستون اصلی -->
      <div class="lg:col-span-2 space-y-6">
        <!-- عنوان صفحه -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <label for="pageTitle" class="block text-sm font-medium text-gray-700 mb-2">عنوان صفحه *</label>
          <input
            id="pageTitle"
            v-model="pageData.title"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="عنوان صفحه را وارد کنید"
          />
        </div>

        <!-- محتوای صفحه -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">محتوای صفحه *</label>
          <textarea
            v-model="pageData.content"
            rows="12"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="محتوای صفحه را وارد کنید"
          ></textarea>
        </div>

        <!-- خلاصه صفحه -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <label for="pageExcerpt" class="block text-sm font-medium text-gray-700 mb-2">خلاصه صفحه</label>
          <textarea
            id="pageExcerpt"
            v-model="pageData.excerpt"
            rows="4"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="خلاصه کوتاهی از صفحه"
          ></textarea>
        </div>
      </div>

      <!-- ستون کناری -->
      <div class="space-y-6">
        <!-- وضعیت انتشار -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">وضعیت انتشار</h3>
          <div class="space-y-3">
            <label class="flex items-center">
              <input
                v-model="pageData.status"
                type="radio"
                value="draft"
                class="ml-2 text-blue-600 focus:ring-blue-500"
              />
              <span class="text-sm text-gray-700">پیش‌نویس</span>
            </label>
            <label class="flex items-center">
              <input
                v-model="pageData.status"
                type="radio"
                value="published"
                class="ml-2 text-blue-600 focus:ring-blue-500"
              />
              <span class="text-sm text-gray-700">منتشر شده</span>
            </label>
            <label class="flex items-center">
              <input
                v-model="pageData.status"
                type="radio"
                value="private"
                class="ml-2 text-blue-600 focus:ring-blue-500"
              />
              <span class="text-sm text-gray-700">خصوصی</span>
            </label>
          </div>
        </div>

        <!-- دسته‌بندی -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <label for="pageCategory" class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
          <select
            id="pageCategory"
            v-model="pageData.categoryId"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">انتخاب دسته‌بندی</option>
            <option value="1">صفحات عمومی</option>
            <option value="2">صفحات اطلاعاتی</option>
            <option value="3">صفحات تماس</option>
          </select>
        </div>

        <!-- تصویر شاخص -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">تصویر شاخص</label>
          <div class="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center">
            <svg class="mx-auto h-12 w-12 text-gray-400" stroke="currentColor" fill="none" viewBox="0 0 48 48">
              <path d="M28 8H12a4 4 0 00-4 4v20m32-12v8m0 0v8a4 4 0 01-4 4H12a4 4 0 01-4-4v-4m32-4l-3.172-3.172a4 4 0 00-5.656 0L28 28M8 32l9.172-9.172a4 4 0 015.656 0L28 28m0 0l4 4m4-24h8m-4-4v8m-12 4h.02" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
            </svg>
            <div class="mt-4">
              <TemplateButton
                type="button"
                bg-gradient="bg-gradient-to-r from-blue-400 to-blue-600"
                text-color="text-white"
                hover-class="hover:from-blue-500 hover:to-blue-700 hover:shadow-lg hover:scale-105"
                focus-class="focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                size="medium"
              >
                انتخاب تصویر
              </TemplateButton>
            </div>
            <p class="text-xs text-gray-500 mt-2">PNG, JPG تا 10MB</p>
          </div>
        </div>

        <!-- تنظیمات SEO -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات SEO</h3>
          <div class="space-y-4">
            <div>
              <label for="metaTitle" class="block text-sm font-medium text-gray-700 mb-1">عنوان متا</label>
              <input
                id="metaTitle"
                v-model="pageData.metaTitle"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="عنوان برای موتورهای جستجو"
              />
            </div>
            <div>
              <label for="metaDescription" class="block text-sm font-medium text-gray-700 mb-1">توضیحات متا</label>
              <textarea
                id="metaDescription"
                v-model="pageData.metaDescription"
                rows="3"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="توضیحات برای موتورهای جستجو"
              ></textarea>
            </div>
            <div>
              <label for="metaKeywords" class="block text-sm font-medium text-gray-700 mb-1">کلمات کلیدی</label>
              <input
                id="metaKeywords"
                v-model="pageData.metaKeywords"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="کلمات کلیدی با کاما جدا کنید"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { ref } from 'vue'

definePageMeta({
  layout: 'admin-main'
})

// تعریف interface برای داده‌های صفحه
interface PageData {
  title: string
  content: string
  excerpt: string
  status: 'draft' | 'published' | 'private'
  categoryId: string
  metaTitle: string
  metaDescription: string
  metaKeywords: string
}

// داده‌های صفحه
const pageData = ref<PageData>({
  title: '',
  content: '',
  excerpt: '',
  status: 'draft',
  categoryId: '',
  metaTitle: '',
  metaDescription: '',
  metaKeywords: ''
})

// ذخیره به عنوان پیش‌نویس
const saveAsDraft = async () => {
  try {
    pageData.value.status = 'draft'
    // اینجا کد ذخیره در دیتابیس قرار می‌گیرد
    console.log('صفحه به عنوان پیش‌نویس ذخیره شد:', pageData.value)
    // نمایش پیام موفقیت
  } catch (error) {
    console.error('خطا در ذخیره پیش‌نویس:', error)
    // نمایش پیام خطا
  }
}

// انتشار صفحه
const publishPage = async () => {
  try {
    pageData.value.status = 'published'
    // اینجا کد ذخیره و انتشار در دیتابیس قرار می‌گیرد
    console.log('صفحه منتشر شد:', pageData.value)
    // نمایش پیام موفقیت
  } catch (error) {
    console.error('خطا در انتشار صفحه:', error)
    // نمایش پیام خطا
  }
}
</script> 
