<template>
  <!-- عنوان باکس با عرض بزرگ‌تر و مرکزچین -->
  <div class="max-w-6xl mx-auto mt-8">
    <div class="bg-white shadow p-6 mb-6 rounded-lg">
      <h1 class="text-2xl font-bold text-blue-900 text-right w-full" style="margin-right:0;padding-right:0;">افزودن اسکیما جدید</h1>
    </div>
  </div>
  <!-- فرم اصلی جداگانه با عرض بزرگ‌تر -->
  <div class="max-w-6xl mx-auto">
    <div class="bg-white rounded-lg shadow p-8">
      <form class="space-y-8" @submit.prevent="handleSubmit">
        <!-- نوع اسکیما -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع اسکیما <span class="text-red-500">*</span></label>
          <select v-model="form.type" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500" :class="{'border-red-400': errors.type}">
            <option value="">انتخاب کنید</option>
            <option v-for="type in schemaTypes" :key="type" :value="type">
              {{ getSchemaTypeLabel(type) }} ({{ type }})
            </option>
          </select>
          <p v-if="errors.type" class="text-xs text-red-500 mt-1">نوع اسکیما الزامی است.</p>
        </div>
        <!-- نام و توضیحات -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام تمپلیت <span class="text-red-500">*</span></label>
            <input 
              v-model="form.name" 
              type="text" 
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500" 
              :class="{'border-red-400': errors.name}"
              placeholder="نام تمپلیت اسکیما"
            >
            <p v-if="errors.name" class="text-xs text-red-500 mt-1">نام تمپلیت الزامی است.</p>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
            <textarea v-model="form.description" rows="2" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="توضیحات تمپلیت"></textarea>
          </div>
        </div>
        <!-- دکمه‌ها -->
        <div class="flex justify-end gap-3 pt-8 border-t">
          <NuxtLink to="../SchemaManagement" class="px-4 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50">بازگشت</NuxtLink>
          <button 
            type="submit" 
            :disabled="submitting || loading"
            class="px-6 py-2 bg-blue-600 text-white rounded-md text-sm font-bold hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
          >
            <div v-if="submitting" class="animate-spin rounded-full h-4 w-4 border-b-2 border-white"></div>
            {{ submitting ? 'در حال ثبت...' : 'ثبت تمپلیت' }}
          </button>
        </div>
      </form>
    </div>
  </div>
  <!-- کانتینر جدا برای فیلدهای نوع اسکیما -->
  <div v-if="form.type === 'Article'" class="max-w-6xl mx-auto mt-8 mb-8">
    <!-- اطلاعات پایه مقاله -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-8 mb-6">
      <div class="text-lg font-bold text-gray-700 mb-6 border-b pb-2">اطلاعات پایه مقاله</div>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <!-- عنوان مقاله -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">عنوان مقاله <span class="text-red-500">*</span></label>
          <input v-model="form.title" type="text" class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right" placeholder="مثال: آموزش سئو برای مبتدیان">
        </div>
        <!-- اسلاگ -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">اسلاگ <span class="text-red-500">*</span></label>
          <input v-model="form.slug" type="text" class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right" placeholder="مثال: seo-for-beginners">
          <p class="text-xs text-gray-400 mt-1">آدرس یکتا برای مقاله (فقط حروف انگلیسی و خط تیره)</p>
        </div>
        <!-- خلاصه مقاله -->
        <div class="md:col-span-2">
          <label class="block text-sm font-medium text-gray-700 mb-2">خلاصه مقاله</label>
          <textarea v-model="form.excerpt" rows="2" class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right" placeholder="خلاصه‌ای کوتاه از مقاله..."></textarea>
        </div>
        <!-- محتوای مقاله -->
        <div class="md:col-span-2">
          <label class="block text-sm font-medium text-gray-700 mb-2">محتوای مقاله <span class="text-red-500">*</span></label>
          <textarea v-model="form.content" rows="6" class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right" placeholder="متن کامل مقاله..."></textarea>
        </div>
        <!-- تصویر شاخص -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تصویر شاخص</label>
          <input v-model="form.featured_image" type="text" class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right" placeholder="آدرس یا آپلود تصویر شاخص">
        </div>
        <!-- وضعیت انتشار -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت انتشار <span class="text-red-500">*</span></label>
          <select v-model="form.status" class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right">
            <option value="">انتخاب کنید</option>
            <option value="draft">پیش‌نویس</option>
            <option value="published">منتشر شده</option>
            <option value="scheduled">زمان‌بندی شده</option>
          </select>
        </div>
        <!-- تاریخ انتشار -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ انتشار</label>
          <input v-model="form.published_at" type="date" class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right">
        </div>
        <!-- دسته‌بندی -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
          <input v-model="form.category_id" type="text" class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right" placeholder="شناسه یا نام دسته‌بندی">
        </div>
        <!-- نویسنده -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نویسنده <span class="text-red-500">*</span></label>
          <input v-model="form.author" type="text" class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right" placeholder="نام نویسنده">
        </div>
      </div>
    </div>
    <!-- اطلاعات متا -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-8 mb-6">
      <div class="text-base font-bold text-gray-700 mb-6 border-b pb-2">اطلاعات متا</div>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">عنوان متا</label>
          <input v-model="form.meta_title" type="text" class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right" placeholder="عنوان متا برای سئو">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کلمات کلیدی متا</label>
          <input v-model="form.meta_keywords" type="text" class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right" placeholder="مثال: سئو, آموزش, مقاله">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات متا</label>
          <input v-model="form.meta_description" type="text" class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right" placeholder="توضیح کوتاه برای موتور جستجو">
        </div>
      </div>
    </div>
    <!-- اطلاعات شبکه‌های اجتماعی (Open Graph) -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-8 mb-6">
      <div class="text-base font-bold text-gray-700 mb-6 border-b pb-2">اطلاعات شبکه‌های اجتماعی (Open Graph)</div>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">OG Title</label>
          <input v-model="form.og_title" type="text" class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right" placeholder="عنوان برای شبکه‌های اجتماعی">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">OG Description</label>
          <input v-model="form.og_description" type="text" class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right" placeholder="توضیح برای شبکه‌های اجتماعی">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">OG Image</label>
          <input v-model="form.og_image" type="text" class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right" placeholder="آدرس تصویر OG">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">OG Type</label>
          <input v-model="form.og_type" type="text" class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right" placeholder="مثال: article">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">OG Site Name</label>
          <input v-model="form.og_site_name" type="text" class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right" placeholder="نام سایت برای OG">
        </div>
      </div>
    </div>
    <!-- نوع اسکیما (ثابت) -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-8">
      <div class="text-base font-bold text-gray-700 mb-6 border-b pb-2">نوع اسکیما</div>
      <input type="text" class="w-full px-4 py-3 border border-gray-200 rounded-lg bg-gray-100 text-gray-500 text-right" value="Article" readonly>
    </div>
  </div>
  <div v-if="form.type === 'Product'" class="bg-gray-50 rounded p-6">
    <h4 class="text-md font-semibold text-gray-900 mb-3">اطلاعات محصول</h4>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">قیمت</label>
        <input v-model="form.price" type="number" step="0.01" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="قیمت محصول">
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">واحد پول</label>
        <select v-model="form.currency" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          <option value="IRR">ریال ایران</option>
          <option value="USD">دلار آمریکا</option>
          <option value="EUR">یورو</option>
        </select>
      </div>
    </div>
  </div>
  <div v-if="form.type === 'Organization'" class="bg-gray-50 rounded p-6">
    <h4 class="text-md font-semibold text-gray-900 mb-3">اطلاعات سازمان</h4>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">آدرس</label>
        <input v-model="form.address" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="آدرس سازمان">
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">شماره تماس</label>
        <input v-model="form.telephone" type="tel" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="شماره تماس">
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
declare const navigateTo: (to: string) => Promise<void>
// eslint-disable-next-line @typescript-eslint/no-explicit-any
declare const useSchema: () => any
</script>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'

definePageMeta({ layout: 'admin-main' })

// استفاده از composable
const { createTemplate, fetchSchemaTypes, loading } = useSchema()

// فرم واکنشی برای اسکیما
const form = reactive({
  name: '',
  type: '',
  description: '',
  title: '',
  slug: '',
  excerpt: '',
  content: '',
  author: '',
  published_at: '',
  featured_image: '',
  status: '',
  category_id: '',
  price: undefined as number | undefined,
  currency: 'IRR',
  address: '',
  telephone: '',
  meta_title: '',
  meta_keywords: '',
  meta_description: '',
  og_title: '',
  og_description: '',
  og_image: '',
  og_type: '',
  og_site_name: '',
  extra_fields: {}
})

// اعتبارسنجی
const errors = reactive({ 
  name: false, 
  type: false 
})

const schemaTypes = ref<string[]>([])
const submitting = ref(false)

// دریافت انواع اسکیما
const loadSchemaTypes = async () => {
  schemaTypes.value = await fetchSchemaTypes()
}

// تابع تبدیل نوع اسکیما به برچسب فارسی
const getSchemaTypeLabel = (type: string): string => {
  const labels: Record<string, string> = {
    'Article': 'مقاله',
    'Product': 'محصول',
    'Organization': 'سازمان',
    'WebPage': 'صفحه وب',
    'BreadcrumbList': 'مسیر ناوبری',
    'FAQPage': 'سوالات متداول',
    'LocalBusiness': 'کسب و کار محلی',
    'Review': 'نظر'
  }
  return labels[type] || type
}

// ثبت فرم
const handleSubmit = async () => {
  // اعتبارسنجی
  errors.name = !form.name.trim()
  errors.type = !form.type
  
  if (errors.name || errors.type) {
    alert('لطفاً فیلدهای الزامی را پر کنید')
    return
  }

  submitting.value = true
  
  try {
    const result = await createTemplate({
      name: form.name,
      type: form.type,
      description: form.description,
      title: form.title,
      slug: form.slug,
      excerpt: form.excerpt,
      content: form.content,
      author: form.author,
      published_at: form.published_at,
      featured_image: form.featured_image,
      status: form.status,
      category_id: form.category_id,
      price: form.price,
      currency: form.currency,
      address: form.address,
      telephone: form.telephone,
      meta_title: form.meta_title,
      meta_keywords: form.meta_keywords,
      meta_description: form.meta_description,
      og_title: form.og_title,
      og_description: form.og_description,
      og_image: form.og_image,
      og_type: form.og_type,
      og_site_name: form.og_site_name,
      extra_fields: form.extra_fields
    })
    
    if (result) {
      alert('تمپلیت اسکیما با موفقیت ایجاد شد!')
      navigateTo('../SchemaManagement')
    }
  } catch (err) {
    console.error('خطا در ایجاد تمپلیت:', err)
  } finally {
    submitting.value = false
  }
}

// Lifecycle
onMounted(() => {
  loadSchemaTypes()
})
</script>

<!--
  این صفحه برای افزودن اسکیما جدید است و فرم آن تمیز و زیبا طراحی شده است.
  کامنت‌ها به فارسی برای توسعه‌دهندگان آینده درج شده است.
--> 
