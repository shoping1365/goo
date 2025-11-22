<template>
     <div>
       <!-- Header and Action Buttons -->
       <div class="flex justify-between items-center mb-6 bg-white rounded-lg shadow p-6">
         <div class="flex items-center">
           <h1 class="text-2xl font-bold text-gray-800 mr-4">ویرایش دسته بندی <span v-if="currentCategory">«{{ currentCategory.name }}»</span></h1>
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
              <button class="flex items-center px-4 py-2 bg-purple-500/20 text-purple-600 rounded-md hover:bg-purple-500/30 transition-colors font-semibold" @click="previewCategory">
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
             <a href="#" :class="[activeTab === 'info' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']" @click.prevent="selectTab('info')">اطلاعات دسته</a>
             <a href="#" :class="[activeTab === 'display' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']" @click.prevent="selectTab('display')">تنظیمات نمایش</a>
             <a href="#" :class="[activeTab === 'seo' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']" @click.prevent="selectTab('seo')">سئو</a>
             <a href="#" :class="[activeTab === 'products' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']" @click.prevent="selectTab('products')">محصولات</a>
             <a href="#" :class="[activeTab === 'images' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']" @click.prevent="selectTab('images')">بنر دسته بندی</a>
             <a href="#" :class="[activeTab === 'message' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']" @click.prevent="selectTab('message')">پیغام دسته بندی</a>
             <a href="#" :class="[activeTab === 'faq' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']" @click.prevent="selectTab('faq')">سوالات متداول</a>
             <a href="#" :class="[activeTab === 'video' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']" @click.prevent="selectTab('video')">ویدیو</a>
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
             
   
             <!-- Name Field -->
             <div class="bg-white rounded-lg shadow p-6">
               <div class="mb-8 flex flex-col md:flex-row gap-6">
                 <div class="flex-1">
                   <label for="category-name" class="block text-sm font-medium text-gray-700 mb-1">نام</label>
                   <input id="category-name" v-model="categoryName" type="text" class="w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50 bg-gray-50 p-2">
                 </div>
                 <div class="flex-1">
                   <label for="category-name-en" class="block text-sm font-medium text-gray-700 mb-1">نام رسمی (English)</label>
                   <input id="category-name-en" v-model="categoryNameEn" type="text" dir="ltr" class="w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50 bg-gray-50 p-2" @blur="onEnglishBlur">
                 </div>
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
             <div class="bg-white rounded-lg shadow p-6 mt-8 grid grid-cols-1 md:grid-cols-2 gap-6">
               <!-- Parent Category -->
               <div>
                  <label for="parent-category" class="block text-sm font-medium text-gray-700 mb-1">دسته پدر</label>
                  <select id="parent-category" v-model="parentId" class="w-full rounded-md border-gray-300 shadow-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-gray-50 p-2">
                    <option value="">هیچکدام</option>
                    <option v-for="c in categories.filter(c=>c.id!== Number(route.params.id))" :key="c.id" :value="c.id">{{ c.name }}</option>
                  </select>
               </div>
   
               <!-- Image Upload -->
               <div>
                 <label for="category-image" class="block text-sm font-medium text-gray-700 mb-1">تصویر</label>
                 <!-- Flex container with reversed order -->
                  <div class="flex items-center space-x-6 space-x-reverse">
   
                   <!-- Preview image -->
                   <img :src="imagePreview" alt="Preview" class="w-12 h-12 object-contain rounded border ml-3" />
   
                   <!-- Buttons container -->
                   <div class="flex items-center">
                     <div class="flex gap-2">
                       <!-- فقط انتخاب از رسانه -->
                       <button type="button" class="bg-blue-500/20 text-blue-600 rounded-md px-4 py-2 text-sm hover:bg-blue-500/30 transition-colors font-semibold" @click="openMedia">انتخاب از رسانه</button>
                     </div>
                     <!-- File name display -->
                     <span class="text-gray-600 mr-2" v-text="selectedFileName"></span>
                   </div>
   
                    <!-- ورودی آپلود بومی حذف شد -->
                 </div>
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
                 <input id="product-count-per-page" v-model="productCountPerPage" type="text" placeholder="20" class="w-24 rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50 bg-gray-50 p-2 text-sm focus:outline-none">
               </div>
             </div>
   
             <!-- Left Column -->
             <div class="space-y-6">
             </div>
           </div>
         </div>
   
         <!-- SEO Settings Tab Content -->
         <div v-if="activeTab === 'seo'" class="bg-gray-50 p-6 rounded-md shadow-md">
           <ProductSeoTab :enable-og-type="false" :slug="seoSlug" @update:slug="updateSlug" />
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
                   <!-- Dynamic rows -->
                   <template v-if="productsInCategory.length">
                     <tr v-for="p in productsInCategory" :key="p.id">
                                               <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                          <img :src="productImage(p)" class="object-cover rounded w-12 h-12 md:w-14 md:h-14" @error="onImgError($event)" />
                        </td>
                       <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700">{{ p.name }}</td>
                       <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700 text-center">
                         <span :class="isProductFeatured(p) ? 'text-green-700' : 'text-red-600'">{{ isProductFeatured(p) ? '✔' : '✖' }}</span>
                       </td>
                       <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700 text-center">{{ p.sort_order || '-' }}</td>
                       <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                        <NuxtLink :to="`/product/sku-${p.sku || p.id}`" target="_blank" class="text-blue-600 hover:text-blue-900">مشاهده</NuxtLink>
                       </td>
                       <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                         <NuxtLink :to="`/admin/product-management/products/edit?id=${p.id}`" class="text-indigo-600 hover:text-indigo-900">ویرایش</NuxtLink>
                       </td>
                       <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                         <button class="text-red-600 hover:text-red-900" @click="deleteProduct(p)">حذف</button>
                       </td>
                     </tr>
                   </template>
                   <tr v-else>
                     <td colspan="7" class="px-6 py-4 text-center text-gray-500">محصولی یافت نشد.</td>
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
   
     <!-- Media Library Modal -->
     <MediaLibraryModal v-model="showMediaModal" default-category="product-categories" @confirm="onMediaSelected" />
 </template>
   
<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { computed, defineAsyncComponent, nextTick, onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue';
import ProductImagesTab from '~/components/product/ProductImagesTab.vue';
import ProductSeoTab from '~/components/product/ProductSeoTab.vue';
import ProductVideoTab from '~/components/product/ProductVideoTab.vue';
import { useNotifier } from '~/composables/useNotifier';
import ProductFAQTab from '~/pages/admin/product-management/products/new/faq.vue';
 
   
 definePageMeta({
   layout: 'admin-main',
   middleware: 'admin'
 });
   
 const activeTab = ref('info');
   
 const selectTab = (tabName: string) => {
   activeTab.value = tabName;
 };
   
 const categoryName = ref('');
 const categoryDescription = ref('');
 const categoryNameEn = ref('');
 const duplicateName = computed(()=>{
   const n = categoryName.value.trim().replace(/\s+/g,' ').toLowerCase()
   if(!n) return false
   // eslint-disable-next-line @typescript-eslint/no-explicit-any
   return categories.value.some((c:any)=>{
      if(route.params.id && Number(route.params.id) === Number(c.id)) return false
      return (c.name||'').trim().replace(/\s+/g,' ').toLowerCase() === n
   })
 })
   
 // const tinyApiKey = 'qwa4j6x5mh2e3241igpyi345b4uhe2d5qeq6f8hy9qfkw2ro';
   
 // Dynamically import the RichTextEditor component
 const RichTextEditor = defineAsyncComponent(() =>
   import('~/components/common/RichTextEditor.vue')
 );
   
 // File upload customization
 const selectedFileName = ref('');
 const imageUrl = ref('');
 const imagePreview = computed(()=> imageUrl.value || '/statics/images/default-image_100.png')
   
 // Media modal state
 const showMediaModal = ref(false)
   
 function openMedia(){ showMediaModal.value = true }
   
 // Display Settings reactive variables
 const isPublished = ref(false);
 const showOnHomepage = ref(false);
 const showInMainMenu = ref(false);
 const displayAsProductRow = ref(false);
 const productCountPerPage = ref('20');
   
 // SEO Settings reactive variables
 const seoSlug = ref('')
 const seoMetaTitle = ref('');
 const seoMetaKeywords = ref('');
 const seoMetaDescription = ref('');
   
 // Category Message reactive variable
 const categoryMessage = ref('');
 
 const tinyApiKey = ref('');
   
 const parentId = ref<string | number>('');
   
 const route = useRoute();
   
 const categories = ref([]);
 const currentCategory = ref(null);
   
 // ---------------------------------------------------------------------------
 // Products for this category
 // ---------------------------------------------------------------------------
 interface Product {
   id?: number | string
   category_id?: number | string
   category?: {
     id?: number | string
     [key: string]: unknown
   }
   [key: string]: unknown
 }
 const allProducts = ref<Product[]>([])
   
   /** Fetch all products once (simple implementation, could be paginated later) */
  async function loadProducts() {
    try {
      // Use fetch to avoid TypeScript issues
      const response = await fetch('/api/products')
      if (response.ok) {
        const data = await response.json()
        if (Array.isArray(data)) {
          allProducts.value = data
        } else if (Array.isArray(data?.data)) {
          allProducts.value = data.data
        } else {
          allProducts.value = []
        }
      } else {
        allProducts.value = []
      }
    } catch (e) {
      console.error('خطا در دریافت محصولات:', e)
      allProducts.value = []
    }
  }
   
 // computed products that belong to current category
 const productsInCategory = computed(() => {
   if (!currentCategory.value) return []
   const id = Number(currentCategory.value.id)
   return allProducts.value.filter((p: Product) => {
     if (p.category_id && Number(p.category_id) === id) return true
     if (p.category && Number(p.category.id) === id) return true
     return false
   })
 })
   
 // eslint-disable-next-line @typescript-eslint/no-explicit-any
 function productImage(p:any) {
   // Helper to convert original url to _thumbnail variant
   const toThumbnail = (url:string)=>{
     if(!url) return ''
     const idx = url.lastIndexOf('.')
     if(idx === -1) return url + '_thumbnail'
     return url.slice(0,idx) + '_thumbnail' + url.slice(idx)
   }

   // 1) Check images array
   const imgs = p.Images || p.images || []
   if (imgs.length) {
     const first = imgs[0]
     if(first.thumbnail || first.thumb) return first.thumbnail || first.thumb
     if(first.image_url) return toThumbnail(first.image_url)
     if(first.url) return toThumbnail(first.url)
   }

   // 2) Fallback to single image field
   if(p.image) return toThumbnail(p.image)

   return '/statics/images/default-image_100.png'
 }
   
 // eslint-disable-next-line @typescript-eslint/no-explicit-any
 function onImgError(e:any) {
   e.target.src = '/statics/images/default-image_100.png'
 }
   
 // eslint-disable-next-line @typescript-eslint/no-explicit-any
 function isProductFeatured(p:any):boolean {
   if (typeof p.is_featured === 'boolean') return p.is_featured
   if (typeof p.featured === 'boolean') return p.featured
   const status = (p.is_featured || p.featured || '').toString().toLowerCase()
   return ['1','true','yes','featured'].includes(status)
 }
   
 async function deleteProduct(p: Product) {
   if (!confirm(`آیا از حذف «${p.name || 'محصول'}» مطمئن هستید؟`)) return
   try {
     await $fetch(`/api/products/${p.id}`, { method: 'DELETE' })
     // Remove from local list
     // eslint-disable-next-line @typescript-eslint/no-explicit-any
     allProducts.value = allProducts.value.filter((x:any) => x.id !== p.id)
   } catch (e) {
     alert('خطا در حذف محصول')
   }
 }
   
  const { success, error, warning } = useNotifier()
  function previewCategory(){
    const slug = (seoSlug.value || '').trim()
    if(!slug){
      warning('برای پیش‌نمایش، ابتدا URL را تعیین کنید', 'پیش‌نمایش دسته‌بندی')
      return
    }
    const url = `/product-category/${slug}?preview=1`
    window.open(url, '_blank', 'noopener,noreferrer')
  }
  const saveCategory = async () => {
   const id = route.params.id;

   // جمع‌آوری داده‌های فرم
    if(duplicateName.value){
      warning('نام دسته‌بندی تکراری است', 'اعتبارسنجی')
      return
    }
    // Rule: URL یا نام رسمی انگلیسی الزامی است
    if(!seoSlug.value.trim() && !categoryNameEn.value.trim()){
      warning('URL یا نام رسمی (English) الزامی است. لطفاً یکی را وارد کنید تا URL ساخته شود', 'اعتبارسنجی')
      return
    }
   const payload = {
     name: categoryName.value.trim(),
     name_en: categoryNameEn.value.trim(),
     slug: seoSlug.value.trim(),
     description: categoryDescription.value,
     parent_id: parentId.value ? Number(parentId.value) : 0,
     image_url: imageUrl.value || selectedFileName.value,
     published: isPublished.value,
     show_on_home: showOnHomepage.value,
     show_in_menu: showInMainMenu.value,
     meta_title: seoMetaTitle.value,
     meta_keywords: seoMetaKeywords.value,
     meta_description: seoMetaDescription.value,
     notice_message: categoryMessage.value,
   };
    try {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      const res = await ($fetch as any)(`/api/product-categories/${id}`, {
        method: 'PUT',
        body: payload
      });
      success('دسته‌بندی با موفقیت ذخیره شد!')
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } catch (e: any) {
      const status = e?.statusCode || e?.response?.status || e?.status
      const msg = e?.data?.message || e?.data?.user_message || e?.statusMessage || e?.message || ''
      const meta = `status=${status||''} method=PUT url=/api/product-categories/${id}`
      error('خطا در ذخیره دسته‌بندی' + (msg?`: ${msg}`:''), 'ذخیره دسته‌بندی', meta)
    }
 }
   
 const slugTouched = ref(false)
 let programmaticUpdate = false
   
 function slugify(str:string){
   return str.toLowerCase().trim().replace(/[^a-z0-9\s-]/g,'').replace(/\s+/g,'-').replace(/-+/g,'-')
 }
   
 function onEnglishBlur(){
   if(!slugTouched.value){
     programmaticUpdate = true
     seoSlug.value = slugify(categoryNameEn.value || '')
     programmaticUpdate = false
   }
 }
   
 function updateSlug(val:string){
   seoSlug.value = val
   if(!programmaticUpdate){
     slugTouched.value = true
   }
 }
   
 onMounted(async () => {
   // Use fetch to avoid TypeScript issues
   try {
     const response = await fetch('/api/product-categories?all=1')
     if (response.ok) {
       const data = await response.json()
       if (Array.isArray(data)) {
         categories.value = data;
       } else if (Array.isArray(data?.data)) {
         categories.value = data.data;
       } else {
         categories.value = [];
       }
     } else {
       categories.value = [];
     }
   } catch (error) {
     console.error('خطا در دریافت دسته‌بندی‌ها:', error)
     categories.value = [];
   }

   // Fetch details of the category being edited
   const id = route.params.id;
   if (id) {
     try {
       const response = await fetch(`/api/product-categories/${id}`)
       if (response.ok) {
         const catRes = await response.json()
         if (catRes) {
           currentCategory.value = catRes;

           // Populate simple input fields (use v-model conversion later if needed)
           nextTick(() => {
             // eslint-disable-next-line @typescript-eslint/no-explicit-any
             categoryName.value = (catRes as any).name || '';
             // eslint-disable-next-line @typescript-eslint/no-explicit-any
             categoryNameEn.value = (catRes as any).name_en || '';
             // eslint-disable-next-line @typescript-eslint/no-explicit-any
             seoSlug.value = (catRes as any).slug || '';
           });

           // eslint-disable-next-line @typescript-eslint/no-explicit-any
           categoryDescription.value = (catRes as any).description || '';
           // eslint-disable-next-line @typescript-eslint/no-explicit-any
           if((catRes as any).image_url){
             // eslint-disable-next-line @typescript-eslint/no-explicit-any
             imageUrl.value = (catRes as any).image_url
             selectedFileName.value = imageUrl.value.split('/').pop() || ''
           }

           // eslint-disable-next-line @typescript-eslint/no-explicit-any
           isPublished.value = (catRes as any).published;
           // eslint-disable-next-line @typescript-eslint/no-explicit-any
           showOnHomepage.value = (catRes as any).show_on_home;
           // eslint-disable-next-line @typescript-eslint/no-explicit-any
           showInMainMenu.value = (catRes as any).show_in_menu;
           // Other fields like SEO
           // eslint-disable-next-line @typescript-eslint/no-explicit-any
           seoMetaTitle.value = (catRes as any).meta_title || '';
           // eslint-disable-next-line @typescript-eslint/no-explicit-any
           seoMetaKeywords.value = (catRes as any).meta_keywords || '';
           // eslint-disable-next-line @typescript-eslint/no-explicit-any
           seoMetaDescription.value = (catRes as any).meta_description || '';

           // eslint-disable-next-line @typescript-eslint/no-explicit-any
           categoryMessage.value = (catRes as any).notice_message || '';

           // Set parent select
           // eslint-disable-next-line @typescript-eslint/no-explicit-any
           parentId.value = (catRes as any).parent_id ?? '';

           // After we know currentCategory, load products (only once)
           await loadProducts()
         }
       }
     } catch (error) {
       console.error('خطا در دریافت جزئیات دسته‌بندی:', error)
     }
   }
 });
   
 // eslint-disable-next-line @typescript-eslint/no-explicit-any
 function onMediaSelected(files:any[]){
   if(files && files.length){
     const f = files[0]
     let url = f.thumbnail || f.url || ''
     if(url && !url.startsWith('/uploads/media')) url = '/uploads/media'+url
     imageUrl.value = url
   }
 }

 defineExpose({ onEnglishBlur, updateSlug })
 </script>
   
 <style scoped>
 /* Add any specific styles here */
 </style>