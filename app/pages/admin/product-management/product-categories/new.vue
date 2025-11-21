<template>
        <div>
          <!-- Header and Action Buttons -->
          <div class="flex justify-between items-center mb-6 bg-white rounded-lg shadow p-6">
            <div class="flex items-center">
              <h1 class="text-2xl font-bold text-gray-800 mr-4">افزودن دسته بندی</h1>
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
                      <p v-if="duplicateName" class="text-red-600 text-xs mt-1">نام دسته‌بندی تکراری است</p>
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
                      <div class="relative">
                        <input 
                          v-model="parentSearchTerm"
                          type="text"
                          placeholder="جستجو در دسته‌بندی‌ها..."
                          class="w-full rounded-md border-gray-300 shadow-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-gray-50 p-2"
                          @focus="showParentDropdown = true; loadParentCategories()"
                          @blur="hideParentDropdown"
                        />
                        <div v-if="showParentDropdown && filteredParentCategories.length > 0" class="absolute z-50 w-full mt-1 bg-white border border-gray-300 rounded-md shadow-lg max-h-60 overflow-y-auto">
                          <div
                            v-for="category in filteredParentCategories"
                            :key="category.id"
                            class="px-3 py-2 hover:bg-gray-100 cursor-pointer text-sm border-b border-gray-100 last:border-b-0"
                            @click="selectParentCategory(category)"
                          >
                            <div class="flex items-center justify-between">
                              <div>
                                <div class="font-medium text-gray-900">{{ category.name }}</div>
                                <div v-if="category.parent_name && category.parent_name !== '-'" class="text-xs text-gray-500">
                                  زیرمجموعه: {{ category.parent_name }}
                                </div>
                              </div>
                              <div class="flex items-center space-x-2 space-x-reverse">
                                <span
                                  class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium"
                                  :class="category.parent_id ? 'bg-green-100 text-green-800' : 'bg-purple-100 text-purple-800'"
                                >
                                  {{ category.parent_id ? 'فرعی' : 'اصلی' }}
                                </span>
                                <span v-if="category.product_count > 0" class="text-xs text-gray-500">
                                  {{ category.product_count }} محصول
                                </span>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                  </div>
      
                  <!-- Image Upload -->
                  <div>
                    <label for="category-image" class="block text-sm font-medium text-gray-700 mb-1">تصویر</label>
                    <!-- Flex container with reversed order -->
                      <div class="flex items-center space-x-6 space-x-reverse">
      
                      <!-- Preview image -->
                      <img :src="imagePreview" alt="Preview" class="w-12 h-12 object-contain rounded border ml-3" />
      
                          <div class="flex items-center">
                          <div class="flex gap-2">
                            <button type="button" class="bg-blue-500/20 text-blue-600 rounded-md px-4 py-2 text-sm hover:bg-blue-500/30 transition-colors font-semibold" @click="openMedia">انتخاب از رسانه</button>
                          </div>
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
                  <!-- Is Final Category -->
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
                      <!-- Placeholder row -->
                      <tr>
                        <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900"></td>
                        <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500"></td>
                        <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500"></td>
                        <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500"></td>
                        <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                          <button class="text-blue-600 hover:text-blue-900">مشاهده</button>
                        </td>
                        <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                          <button class="text-indigo-600 hover:text-indigo-900">ویرایش</button>
                        </td>
                        <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                          <button class="text-red-600 hover:text-red-900">حذف</button>
                        </td>
                      </tr>
                        <tr>
                        <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900"></td>
                        <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500"></td>
                        <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500"></td>
                        <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500"></td>
                        <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                          <button class="text-blue-600 hover:text-blue-900">مشاهده</button>
                        </td>
                        <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                          <button class="text-indigo-600 hover:text-indigo-900">ویرایش</button>
                        </td>
                        <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                          <button class="text-red-600 hover:text-red-900">حذف</button>
                        </td>
                      </tr>
                      <!-- More rows will be added here -->
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

          <!-- Media Library Modal -->
          <MediaLibraryModal v-model="showMediaModal" default-category="product-categories" @confirm="onMediaSelected" />
        </div>
      </template>
      
    <script lang="ts">
    declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
    declare const navigateTo: (to: string) => Promise<void>
    </script>

    <script setup lang="ts">
    import { computed, defineAsyncComponent, onMounted, ref } from 'vue';
      // import { useFetch } from '#app';
      import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue';
import ProductImagesTab from '~/components/product/ProductImagesTab.vue';
import ProductSeoTab from '~/components/product/ProductSeoTab.vue';
import ProductVideoTab from '~/components/product/ProductVideoTab.vue';
import { useNotifier } from '~/composables/useNotifier';
import ProductFAQTab from '~/pages/admin/product-management/products/new/faq.vue';
// import { useConfirmDialog } from '~/composables/useConfirmDialog'
    
      
    definePageMeta({
      layout: 'admin-main'
    });
      
    interface Category {
      id?: number
      name?: string
      slug?: string
      parent_id?: number | null
      order?: number
      published?: boolean
      image_url?: string
      banner_url?: string
      product_count?: number
      parent_name?: string
      parent_slug?: string
      [key: string]: unknown
    }
    
    const activeTab = ref('info');
      
    const selectTab = (tabName: string) => {
      activeTab.value = tabName;
    };
      
    const categoryName = ref('');
    const categoryNameEn = ref('');
    // Categories list (loaded later) declared early for TS reference
    const categories = ref<Category[]>([]);
      
    const duplicateName = computed(() => {
      const n = categoryName.value.trim().replace(/\s+/g,' ').toLowerCase()
      if(!n) return false
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      return categories.value.some((c:any)=> (c.name||'').trim().replace(/\s+/g,' ').toLowerCase() === n)
    })
    const categoryDescription = ref('');
      
    // const tinyApiKey = 'qwa4j6x5mh2e3241igpyi345b4uhe2d5qeq6f8hy9qfkw2ro';
      
    // Dynamically import the RichTextEditor component
    const RichTextEditor = defineAsyncComponent(() =>
      import('~/components/common/RichTextEditor.vue')
    );
      
    // Image selection & preview
    const selectedFileName = ref('');
    const imageUrl = ref('');
    const imagePreview = computed(()=> imageUrl.value || '/statics/images/default-image_100.png')
      
    
      
    // async function handleFileChange(event: Event) {
    //   const target = event.target as HTMLInputElement;
    //   if (!target.files || !target.files.length) return;
    //   const file = target.files[0];
    //   selectedFileName.value = file.name;

    //   // Show a temporary preview while uploading
    //   imageUrl.value = URL.createObjectURL(file);

    //   const xhr = new XMLHttpRequest();
    //   xhr.open('POST', '/api/media/upload');
      

    //   xhr.onreadystatechange = () => {
    //     if (xhr.readyState === 4) {
    //       if (xhr.status === 200) {
    //         try {
    //           const json = JSON.parse(xhr.responseText);
    //           const payload = json.data || (Array.isArray(json.files) ? json.files[0] : null);
    //           if (payload) {
    //             let url = payload.url || payload.file_path || '';
    //             if (url && !url.startsWith('/uploads/media')) url = '/uploads/media/product-categories/' + url;
    //             imageUrl.value = url;
    //           }
    //         } catch (err) {
    //           console.error('Error parsing upload response', err);
    //         }
    //       } else {
    //         console.error('Upload failed', xhr.status);
    //       }
    //     }
    //   };

    //   const fd = new FormData();
    //   fd.append('file', file);
    //   fd.append('category', 'product-categories');
    //   xhr.send(fd);
    // }
      
    const showMediaModal = ref(false)
    function openMedia(){ showMediaModal.value = true }
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    function onMediaSelected(files:any[]){
      if(files && files.length){
        const f = files[0]
        imageUrl.value = f.thumbnail || f.url || ''
        selectedFileName.value = (f.url || '').split('/').pop() || ''
      }
      showMediaModal.value = false
    }
      
    // Display Settings reactive variables
      const isPublished = ref(false);
      const showOnHomepage = ref(false);
      const showInMainMenu = ref(false);
      const displayAsProductRow = ref(false);
    const productCountPerPage = ref('20');
      
    // SEO Settings reactive variables
    const seoSlug = ref('')
    const slugTouched = ref(false)
    let programmaticUpdate = false
      
    function slugify(str:string){
      return str
        .toLowerCase()
        .trim()
        .replace(/[^a-z0-9\s-]/g, '') // remove invalid chars
        .replace(/\s+/g, '-')          // spaces to hyphen
        .replace(/-+/g, '-')            // collapse multiple hyphens
    }
      
    // Generate slug when english name field loses focus
    function onEnglishBlur(){
      if(!slugTouched.value){
        programmaticUpdate = true
        seoSlug.value = slugify(categoryNameEn.value || '')
        programmaticUpdate = false
      }
    }
      
    // SEO Settings reactive variables
    // const seoPageName = ref('');
    const seoMetaTitle = ref('');
    const seoMetaKeywords = ref('');
    const seoMetaDescription = ref('');
      
    // Category Message reactive variable
    const categoryMessage = ref('');
    
    const tinyApiKey = ref('');
      
    const parentId = ref<string | number>('')
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const selectedParentCategory = ref<any>(null)
    const parentSearchTerm = ref('')
    const showParentDropdown = ref(false)
      
      const { error, warning } = useNotifier()
      function previewCategory(){
        const slug = (seoSlug.value || '').trim() || slugify(categoryNameEn.value || '')
        if(!slug){
          warning('برای پیش‌نمایش، ابتدا URL یا نام رسمی (English) را تعیین کنید', 'پیش‌نمایش دسته‌بندی')
          return
        }
        const url = `/product-category/${slug}?preview=1`
        window.open(url, '_blank', 'noopener,noreferrer')
      }
      const saveCategory = async () => {
        if(duplicateName.value){
          warning('نام دسته‌بندی تکراری است', 'اعتبارسنجی')
          return
        }
      // Validation: when publishing, required fields must be filled
        // Rule: URL یا نام رسمی انگلیسی الزامی است
        if(!seoSlug.value.trim() && !categoryNameEn.value.trim()){
          warning('URL یا نام رسمی (English) الزامی است. لطفاً یکی را وارد کنید تا URL ساخته شود', 'اعتبارسنجی')
          return
        }
      // جمع‌آوری داده‌های فرم
        interface CategoryPayload {
          name: string
          name_en: string
          slug?: string
          parent_id?: number | string | null
          description?: string
          [key: string]: unknown
        }
        const payload: CategoryPayload = {
        name: categoryName.value.trim(),
        name_en: categoryNameEn.value.trim(),
        slug: seoSlug.value.trim(),
        description: categoryDescription.value,
        image_url: imageUrl.value || selectedFileName.value,
        published: isPublished.value,
        show_on_home: showOnHomepage.value,
        show_in_menu: showInMainMenu.value,
        meta_title: seoMetaTitle.value,
        meta_keywords: seoMetaKeywords.value,
        meta_description: seoMetaDescription.value,
        notice_message: categoryMessage.value,
      }
      
      // parent_id فقط زمانی ارسال شود که انتخاب شده باشد
      if (parentId.value) {
        payload.parent_id = Number(parentId.value)
      }
      
        try {
          // اگر URL خالی است ولی نام انگلیسی داریم، پیش‌ساخت
          if(!seoSlug.value.trim() && categoryNameEn.value.trim()){
            seoSlug.value = slugify(categoryNameEn.value)
          }
          interface CreateCategoryResponse {
            id?: number | string
            success?: boolean
            [key: string]: unknown
          }
          const res = await $fetch<CreateCategoryResponse>('/api/product-categories', {
          method: 'POST',
          body: payload
        });
          const createdId = (res && (res.id || (res.data && res.data.id)))
          // پس از ایجاد، به حالت ویرایش برو
          if(createdId){
            navigateTo(`/admin/product-management/product-categories/${createdId}/edit`)
          }
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      } catch (e: any) {
          console.error('❌ خطا در ایجاد دسته‌بندی:', e)
          const status = e?.statusCode || e?.response?.status || e?.status
          const code = e?.data?.code || ''
          const msg = e?.data?.message || e?.data?.user_message || e?.message || ''
          const isDup = (code.toString().includes('DUPLICATE') || /slug/i.test(msg))
          const message = isDup
            ? 'URL انتخاب‌شده تکراری است. در صورت عدم تغییر، سیستم به‌صورت خودکار URL یکتا می‌سازد.'
            : ('خطا در ذخیره دسته‌بندی' + (msg ? `: ${msg}` : ''))
          const meta = `status=${status||''} method=POST url=/api/product-categories`
          error(message, 'ذخیره دسته‌بندی', meta)
      }
    }
      
    onMounted(async () => {
      // فقط وقتی لازم است دسته‌بندی‌ها لود کن (lazy loading)
      // یعنی زمانی که کاربر روی parent category dropdown کلیک کند
    });

    // تابع برای لود دسته‌بندی‌ها هنگام کلیک روی dropdown
    async function loadParentCategories() {
      if (categories.value.length > 0) return; // اگر قبلاً لود شده، دوباره نلود کن
      
      try {
        const response = await $fetch('/api/product-categories?all=1');
        let raw = []
        if (Array.isArray(response)) {
          raw = response;
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        } else if (Array.isArray((response as any)?.data)) {
          // eslint-disable-next-line @typescript-eslint/no-explicit-any
          raw = (response as any).data;
        } else {
          raw = [];
        }
        
        // Add parent_name for display
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        raw.forEach((cat: any) => {
          if (cat.parent_id) {
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            const parent = raw.find((c: any) => c.id === cat.parent_id)
            cat.parent_name = parent ? parent.name : '-'
          } else {
            cat.parent_name = '-'
          }
        })
        
        categories.value = raw;
      } catch (err) {
        console.error('خطا در بارگذاری دسته‌بندی‌ها:', err);
      }
    }
      
    // function getParentName(parentId) {
    //   if (!parentId) return '-';
    //   const parent = categories.value.find(c => c.id === parentId);
    //   return parent ? parent.name : '-';
    // }

    function updateSlug(val:string){
      if(seoSlug.value === val) return
      seoSlug.value = val
      if(!programmaticUpdate){
        slugTouched.value = true
      }
    }

    // Handle parent category selection change
     
    // function handleParentCategoryChange(category: any) {
    //   selectedParentCategory.value = category
    //   parentId.value = category ? category.id : ''
    // }

    // Computed filtered categories for parent selection
    const filteredParentCategories = computed(() => {
      if (!parentSearchTerm.value.trim()) {
        return categories.value.slice(0, 10) // Show first 10 when no search
      }
      
      const term = parentSearchTerm.value.toLowerCase()
      return categories.value.filter(category => 
        category.name.toLowerCase().includes(term) ||
        (category.parent_name && category.parent_name.toLowerCase().includes(term))
      )
    })

    // Select parent category
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    function selectParentCategory(category: any) {
      selectedParentCategory.value = category
      parentId.value = category.id
      parentSearchTerm.value = category.name
      showParentDropdown.value = false
    }

    // Hide parent dropdown with delay
    function hideParentDropdown() {
      setTimeout(() => {
        showParentDropdown.value = false
      }, 200)
    }

    defineExpose({ onEnglishBlur, updateSlug })
    
    // Register components
    // const components = {
    //   CategorySearchBox
    // }
    </script>
      
    <style scoped>
    /* Add any specific styles here */
    </style>
