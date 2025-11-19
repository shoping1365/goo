<template>
     <div>
       <!-- Header and Action Buttons -->
       <div class="flex justify-between items-center mb-6 bg-white rounded-lg shadow p-6">
         <div class="flex items-center">
           <h1 class="text-2xl font-bold text-gray-800 mr-4">دسته بندی+برند </h1>
         </div>
         <div class="flex flex-col items-end space-y-2">
           <div class="flex gap-3 flex-row-reverse">
             <button class="flex items-center px-4 py-2 bg-red-500/20 text-red-600 rounded-md hover:bg-red-500/30 transition-colors font-semibold">
               حذف
             </button>
             <button
               class="flex items-center px-4 py-2 bg-blue-500/20 text-blue-600 rounded-md hover:bg-blue-500/30 transition-colors font-semibold"
               @click="saveCategory"
             >
               ذخیره
             </button>
             <button class="flex items-center px-4 py-2 bg-purple-500/20 text-purple-600 rounded-md hover:bg-purple-500/30 transition-colors font-semibold">
               پیش نمایش
             </button>
           </div>
         </div>
       </div>

       <!-- Container for Back link and Tabs -->
       <div class="border-b border-gray-200 mt-8">
         <div class="bg-white rounded-t-lg shadow p-6 flex items-center justify-between">
           <!-- Tabs - Moved to the right -->
           <nav class="-mb-px flex gap-6 space-x-reverse" aria-label="Tabs">
             <a href="#" @click.prevent="selectTab('info')" :class="[activeTab === 'info' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']">اطلاعات دسته</a>
             <a href="#" @click.prevent="selectTab('display')" :class="[activeTab === 'display' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']">تنظیمات نمایش</a>
             <a href="#" @click.prevent="selectTab('seo')" :class="[activeTab === 'seo' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']">سئو</a>
             <a href="#" @click.prevent="selectTab('products')" :class="[activeTab === 'products' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']">محصولات</a>
             <a href="#" @click.prevent="selectTab('images')" :class="[activeTab === 'images' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']">بنر دسته بندی</a>
             <a href="#" @click.prevent="selectTab('message')" :class="[activeTab === 'message' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']">پیغام دسته بندی</a>
             <a href="#" @click.prevent="selectTab('faq')" :class="[activeTab === 'faq' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']">سوالات متداول</a>
             <a href="#" @click.prevent="selectTab('video')" :class="[activeTab === 'video' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']">ویدیو</a>
             <!-- Add other tabs as needed -->
           </nav>
           <!-- Back link - Moved to the left -->
           <NuxtLink to="/admin/product-management/product-categories" class="flex items-center text-blue-600 hover:text-blue-800 text-sm">
             <img src="/svg/arrow-left.svg" class="w-4 h-4 ml-1" alt="blue-arrow " />
             برگشت به لیست دسته‌بندی‌ها
           </NuxtLink>
         </div>
       </div>

       <!-- Tab Content will go here -->
       <div class="mt-6">
         <div v-if="activeTab === 'info'">
           <!-- Content for the active tab -->
           <div class="space-y-6">
             <!-- Language Selection - REMOVED -->


             <!-- Category & Brand Row -->
             <div class="mb-8 grid grid-cols-1 md:grid-cols-2 gap-6 order-first">
               <!-- Background wrapper -->
               <!-- Category -->
               <div>
                 <label class="block text-sm font-medium text-gray-700 mb-1">دسته‌بندی <span class="text-red-600">*</span></label>
                 <select v-model="page.category_id" class="w-full rounded-md border-gray-300 shadow-sm bg-gray-50 p-2">
                   <option value="">-- انتخاب --</option>
                   <option v-for="c in categories" :key="c.id" :value="c.id">{{ c.name }}</option>
                 </select>
                 <p v-if="errors.category" class="text-red-600 text-sm mt-1">{{ errors.category }}</p>
               </div>

               <!-- Brand -->
               <div>
                 <label class="block text-sm font-medium text-gray-700 mb-1">برند <span class="text-red-600">*</span></label>
                 <select v-model="page.brand_id" class="w-full rounded-md border-gray-300 shadow-sm bg-gray-50 p-2">
                   <option value="">-- انتخاب --</option>
                   <option v-for="b in brands" :key="b.id" :value="b.id">{{ b.name }}</option>
                 </select>
                 <p v-if="errors.brand" class="text-red-600 text-sm mt-1">{{ errors.brand }}</p>
               </div>
             </div>

             <!-- Name Field -->
             <div class="mb-8">
               <label for="category-name" class="block text-sm font-medium text-gray-700 mb-1">نام</label>
               <input type="text" id="category-name" v-model="page.name" class="w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50 bg-gray-50 p-2">
             </div>

             <!-- Description Field (Rich Text Editor Placeholder) -->
             <div>
               <label for="category-description" class="block text-sm font-medium text-gray-700 mb-1">توضیحات</label>
               <ClientOnly>
                 <RichTextEditor
                   v-model="categoryDescription"
                   :api-key="tinyApiKey"
                   :lang="'fa'"
                   :direction="'rtl'"
                   :height="400"
                 />
               </ClientOnly>
             </div>
           </div>

           <!-- Parent Category & Image Upload -->
           <div class="bg-white rounded-lg shadow p-6 mt-8">
             <!-- Image Upload -->
             <div>
               <label for="category-image" class="block text-sm font-medium text-gray-700 mb-1">تصویر</label>
               <!-- Flex container with reversed order -->
               <div class="flex items-center space-x-4 space-x-reverse">

                 <!-- Default image placeholder - moved to the right -->
                 <img src="/statics/images/default-image_100.png" alt="No Image" class="w-12 h-12 object-contain opacity-60" />

                 <!-- Custom file input appearance - moved to the left -->
                 <div class="flex items-center">
                   <!-- Buttons container -->
                   <div class="flex gap-2">
                     <label for="category-image-upload" class="bg-green-500/20 text-green-600 rounded-md px-4 py-2 text-sm hover:bg-green-500/30 transition-colors font-semibold cursor-pointer">
                       انتخاب تصویر
                     </label>
                     <button class="bg-blue-500/20 text-blue-600 rounded-md px-4 py-2 text-sm hover:bg-blue-500/30 transition-colors font-semibold">انتخاب از رسانه</button>
                   </div>
                   <!-- File name display -->
                   <span class="text-gray-600 mr-2" v-text="selectedFileName"></span>
                 </div>

                 <!-- Hidden native input -->
                 <input type="file" id="category-image-upload" @change="handleFileChange" class="hidden"/>
               </div>
             </div>
           </div>
         </div>

         <!-- Display Settings Tab Content -->
         <div v-if="activeTab === 'display'" class="bg-gray-50 p-6 rounded-md shadow-md">
           <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
             <!-- Right Column -->
             <div class="space-y-6">
               <!-- Published Checkbox -->
               <div class="flex items-center justify-start">
                 <input id="is-published" v-model="isPublished" type="checkbox" class="h-4 w-4 text-blue-600 border-gray-300 rounded ml-2">
                 <label for="is-published" class="block text-sm text-gray-900">منتشر شده</label>
               </div>

               <!-- Is Final Category -->
               <div class="flex items-center justify-start">
                 <input id="is-final-category" v-model="isFinalCategory" type="checkbox" class="h-4 w-4 text-blue-600 border-gray-300 rounded ml-2">
                 <label for="is-final-category" class="block text-sm text-gray-900">دسته آخر می باشد</label>
               </div>

               <!-- Show on Homepage -->
               <div class="flex items-center justify-start">
                 <input id="show-on-homepage" v-model="showOnHomepage" type="checkbox" class="h-4 w-4 text-blue-600 border-gray-300 rounded ml-2">
                 <label for="show-on-homepage" class="block text-sm text-gray-900">نمایش در صفحه اصلی</label>
               </div>

               <!-- Show in Main Menu -->
               <div class="flex items-center justify-start">
                 <input id="show-in-main-menu" v-model="showInMainMenu" type="checkbox" class="h-4 w-4 text-blue-600 border-gray-300 rounded ml-2">
                 <label for="show-in-main-menu" class="block text-sm text-gray-900">در منوی اصلی سایت نمایش داده بشود</label>
               </div>

               <!-- Display as Product Row -->
               <div class="flex items-center justify-start">
                 <input id="display-as-product-row" v-model="displayAsProductRow" type="checkbox" class="h-4 w-4 text-blue-600 border-gray-300 rounded ml-2">
                 <label for="display-as-product-row" class="block text-sm text-gray-900">نمایش بعنوان یک ردیف از محصولات در صفحه اصلی</label>
               </div>

               <!-- Product Count Per Page -->
               <div class="flex items-center justify-start">
                 <label for="product-count-per-page" class="block text-sm font-medium text-gray-700 ml-2">تعداد محصول در هر صفحه دسته بندی</label>
                 <input type="text" id="product-count-per-page" v-model="productCountPerPage" placeholder="20" class="w-24 rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50 bg-gray-50 p-2 text-sm focus:outline-none">
               </div>
             </div>

             <!-- Left Column -->
             <div class="space-y-6">
             </div>
           </div>
         </div>

         <!-- SEO Settings Tab Content -->
         <div v-if="activeTab === 'seo'" class="bg-gray-50 p-6 rounded-md shadow-md">
           <ProductSeoTab />
         </div>

         <!-- Products Tab Content -->
         <div v-if="activeTab === 'products'" class="bg-gray-50 p-6 rounded-md shadow-md">
           <div class="space-y-6">
             <!-- Products Table -->
             <div class="overflow-x-auto">
               <table class="min-w-full divide-y divide-gray-200">
                 <thead class="bg-white">
                   <tr>
                     <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تصویر</th>
                     <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">محصول</th>
                     <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آیا محصول ویژه است؟</th>
                     <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ترتیب نمایش</th>
                     <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مشاهده</th>
                     <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ویرایش</th>
                     <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">حذف</th>
                   </tr>
                 </thead>
                 <tbody class="bg-white divide-y divide-gray-200">
                   <!-- Loading State -->
                   <tr v-if="productsLoading">
                     <td colspan="7" class="px-6 py-8 whitespace-nowrap text-center text-sm text-gray-500">
                       در حال بارگذاری محصولات...
                     </td>
                   </tr>

                   <!-- Empty State -->
                   <tr v-else-if="products.length === 0">
                     <td colspan="7" class="px-6 py-8 whitespace-nowrap text-center text-sm text-gray-500">
                       محصولی برای این دسته بندی پیدا نشد.
                     </td>
                   </tr>

                   <!-- Product Rows -->
                   <tr v-else v-for="product in products" :key="product.id">
                     <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                       <img :src="getProductThumbnail(product)" alt="تصویر محصول" class="w-12 h-12 object-cover rounded" />
                     </td>
                     <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700">
                       {{ product.name }}
                     </td>
                     <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                       <span v-if="product.is_special">بله</span>
                       <span v-else>خیر</span>
                     </td>
                     <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ product.order || 0 }}</td>
                     <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                       <NuxtLink :to="`/product/sku-${product.sku || product.id}${product.slug ? '/' + encodeURIComponent(product.slug) : ''}`" class="text-blue-600 hover:text-blue-900">مشاهده</NuxtLink>
                     </td>
                      <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                        <NuxtLink :to="`/admin/product-management/products/edit?id=${product.id}`" class="text-indigo-600 hover:text-indigo-900">ویرایش</NuxtLink>
                      </td>
                     <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                       <button @click="removeProduct(product.id)" class="text-red-600 hover:text-red-900">حذف</button>
                     </td>
                   </tr>
                 </tbody>
               </table>
             </div>

             <!-- Pagination and Footer -->
             <div class="flex items-center justify-between mt-4">
               <div class="flex-1 flex justify-between sm:hidden">
                 <a href="#" class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-100">قبل</a>
                 <a href="#" class="relative ml-3 inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-100">بعد</a>
               </div>
               <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
                 <div>
                   <p class="text-sm text-gray-700">
                     نمایش
                     <span class="font-medium">1</span>
                     تا
                     <span class="font-medium">2</span>
                     از
                     <span class="font-medium">2</span>
                     آیتم
                   </p>
                 </div>
                 <div>
                   <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
                     <a href="#" class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-100">
                       <span class="sr-only">قبلی</span>
                       <!-- Heroicon name: solid/chevron-right -->
                       <img src="/svg/chevron-right.svg" class="h-5 w-5" alt="Chevron Right" />
                     </a>
                      <a href="#" aria-current="page" class="z-10 bg-blue-50 border-blue-500 text-blue-600 relative inline-flex items-center px-4 py-2 border text-sm font-medium">1</a>
                      <a href="#" class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-100">
                       <span class="sr-only">بعدی</span>
                       <!-- Heroicon name: solid/chevron-left -->
                       <img src="/svg/chevron-left.svg" class="h-5 w-5" alt="Chevron Left" />
                     </a>
                   </nav>
                 </div>
               </div>
             </div>

           </div>
         </div>

         <!-- Category Banner Tab Content -->
         <div v-if="activeTab === 'images'" class="bg-gray-50 p-6 rounded-md shadow-md">
           <ProductImagesTab />
         </div>

         <!-- Category Message Tab Content -->
         <div v-if="activeTab === 'message'" class="bg-gray-50 p-6 rounded-md shadow-md">
           <div class="space-y-6">
             <!-- Message Textarea -->
             <div>
               <label for="category-message" class="block text-sm font-medium text-gray-700 mb-1">پیغام</label>
               <textarea id="category-message" v-model="categoryMessage" rows="4" class="w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50 bg-gray-50 p-2 focus:outline-none"></textarea>
             </div>
           </div>
         </div>

         <!-- FAQ Tab Content -->
         <div v-if="activeTab === 'faq'" class="bg-gray-50 p-6 rounded-md shadow-md">
           <ProductFAQTab />
         </div>

         <!-- Video Tab Content -->
         <div v-if="activeTab === 'video'" class="bg-gray-50 p-6 rounded-md shadow-md">
           <ProductVideoTab />
         </div>

       </div>
     </div>

 </template>

 <script setup lang="ts">
 import { defineAsyncComponent, onMounted, reactive, ref, watch } from 'vue';
import ProductImagesTab from '~/components/product/ProductImagesTab.vue';
import ProductSeoTab from '~/components/product/ProductSeoTab.vue';
import ProductVideoTab from '~/components/product/ProductVideoTab.vue';
import { useAuth } from '~/composables/useAuth';
import ProductFAQTab from '~/pages/admin/product-management/products/new/faq.vue';

// تعریف definePageMeta برای Nuxt 3
// @ts-ignore: Nuxt macro
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void

 definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
});

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

 const activeTab = ref('info');

 const selectTab = (tabName: string) => {
   activeTab.value = tabName;
 };

 const categoryDescription = ref('');

 const tinyApiKey = 'qwa4j6x5mh2e3241igpyi345b4uhe2d5qeq6f8hy9qfkw2ro';

 // Dynamically import the RichTextEditor component
 const RichTextEditor = defineAsyncComponent(() =>
  import('~/components/common/RichTextEditor.vue')
 );

 // File upload customization
 const selectedFileName = ref('');

 const handleFileChange = (event: Event) => {
   const target = event.target as HTMLInputElement;
   if (target.files && target.files.length > 0) {
     selectedFileName.value = target.files[0].name;
   } else {
     selectedFileName.value = '';
   }
 };

 // --------------------------------------------------
 // Main reactive form object (page)
 // --------------------------------------------------

 const page = reactive<any>({
   name: '',
   slug: '',
   category_id: '',
   brand_id: '',
   published: false,
   show_on_home: false,
   show_in_menu: false,
   order: 0,
   notice_message: ''
 })

 // Build name and slug when category or brand changes
 watch([() => page.category_id, () => page.brand_id], ([catId, brId]) => {
   const cat = categories.value.find(c => c.id === catId)
   const br = brands.value.find(b => b.id === brId)
   if (cat && br) {
     page.name = `${cat.name} ${br.name}`
     page.slug = `${cat.slug}/${br.slug}`
   }
 })

 // Display Settings reactive variables
 const isPublished = ref(false);
 const showOnHomepage = ref(false);
 const showInMainMenu = ref(false);
 const isFinalCategory = ref(false);
 const displayAsProductRow = ref(false);
 const productCountPerPage = ref('20');

 // SEO Settings reactive variables
 const seoPageName = ref('');
 const seoMetaTitle = ref('');
 const seoMetaKeywords = ref('');
 const seoMetaDescription = ref('');

 // Category Message reactive variable
 const categoryMessage = ref('');

 const errors = reactive<{ name?: string; slug?: string; category?: string; brand?: string }>({});

 const saveCategory = async () => {
   // basic validation
   errors.category = page.category_id ? '' : 'انتخاب دسته‌بندی الزامی است'
   errors.brand    = page.brand_id    ? '' : 'انتخاب برند الزامی است'
   if (errors.category || errors.brand) return

   const payload: any = {
     name: page.name,
     slug: page.slug,
     parent_id: page.category_id,
     category_id: page.category_id,
     brand_id: page.brand_id,
     description: categoryDescription.value,
     image_url: selectedFileName.value,
     published: isPublished.value,
     show_on_home: showOnHomepage.value,
     show_in_menu: showInMainMenu.value,
     is_final_category: isFinalCategory.value,
     meta_title: seoMetaTitle.value,
     meta_keywords: seoMetaKeywords.value,
     meta_description: seoMetaDescription.value,
     notice_message: categoryMessage.value,
   }

   try {
     const res = await ($fetch as any)('/api/product-categories', {
       method: 'POST',
       body: payload
     });
     alert('دسته‌بندی با موفقیت ذخیره شد!');
     // در صورت نیاز ریدایرکت:
     // navigateTo('/admin/product-categories')
   } catch (e: any) {
     alert('خطا در ذخیره دسته‌بندی: ' + (e?.statusMessage || e?.message || ''));
   }
 }

 const categories = ref<any[]>([]);
 const brands = ref<any[]>([]);
 
 // Declare useFetch for TypeScript
 declare const useFetch: <T>(url: string, options?: unknown) => Promise<{ data: { value: T } }>
 
  onMounted(async () => {
    const { data } = await useFetch<unknown>('/api/product-categories?all=1');
    if (Array.isArray((data as { value: unknown }).value)) {
      categories.value = (data as { value: unknown[] }).value;
    } else if (Array.isArray(((data as { value: { data?: unknown[] } }).value as { data?: unknown[] })?.data)) {
      categories.value = ((data as { value: { data: unknown[] } }).value).data;
    } else {
      categories.value = [];
    }
  });

  onMounted(async () => {
    const { data: brandRes } = await useFetch<unknown>('/api/brands?all=1');
    brands.value = Array.isArray((brandRes as { value: unknown }).value)
      ? (brandRes as { value: unknown[] }).value as unknown[]
      : ((((brandRes as { value: { data?: unknown[] } }).value as { data?: unknown[] })?.data) || []);
  });

 // --------------------------------------------------
 // Products list for selected category
 // --------------------------------------------------

 const products = ref<any[]>([])
 const productsLoading = ref(false)

 // Helper to build thumbnail URL (fallback to default placeholder)
 const getProductThumbnail = (product: any) => {
   if (product?.images?.length && product.images[0].image_url) {
     const url = product.images[0].image_url as string
     const dotIdx = url.lastIndexOf('.')
     if (dotIdx === -1) return url
     return url.slice(0, dotIdx) + '_thumbnail' + url.slice(dotIdx)
   }
   // Use absolute path for public assets to avoid build resolution issues
   return '/statics/images/default-image_100.png'
 }

 const fetchProducts = async () => {
   if (!page.category_id) {
     products.value = []
     return
   }
    try {
      productsLoading.value = true
      const res: any = await ($fetch as any)('/api/products')
     const list = Array.isArray(res) ? res : (res?.data || [])
     products.value = list.filter((p: any) => {
       // match by category id – the API may return either category_id or nested category object
       return p.category_id === page.category_id || (p.category && p.category.id === page.category_id)
     })
   } catch (e) {
     console.error('خطا در دریافت محصولات:', e)
     products.value = []
   } finally {
     productsLoading.value = false
   }
 }

 // Watch for category change to refresh products list
 watch(() => page.category_id, () => {
   fetchProducts()
 })

 // Simple delete handler (optional)
 const removeProduct = async (productId: number) => {
   if (!confirm('آیا از حذف این محصول مطمئن هستید؟')) return
   try {
      await ($fetch as any)(`/api/products/${productId}`, { method: 'DELETE' })
     await fetchProducts()
   } catch (e) {
     console.error('خطا در حذف محصول:', e)
     alert('خطا در حذف محصول')
   }
 }

 // Initial fetch if category already selected
 onMounted(() => {
   if (page.category_id) fetchProducts()
 })

 </script>

 <style scoped>
 /* Add any specific styles here */
 </style>

