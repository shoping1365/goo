<template>
  <div>
    <!-- Header and Action Buttons -->
    <div class="flex justify-between items-center mb-6 bg-white rounded-lg shadow px-4 py-4">
      <div class="flex items-center">
        <h1 class="text-2xl font-bold text-gray-800 mr-4">ویرایش برند {{ name || '' }}</h1>
      </div>
      <div class="flex flex-col items-end space-y-2">
        <div class="flex gap-2 flex-row-reverse">
          <button class="px-4 py-2 rounded-lg text-white bg-gradient-to-r from-red-400 to-red-600 hover:from-red-500 hover:to-red-700 focus:outline-none focus:ring-2 focus:ring-red-400 shadow-md transition-all duration-200 font-semibold">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 ml-2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m-1.022.165L5.93 19.673a2.25 2.25 0 002.244 2.077h7.412a2.25 2.25 0 002.244-2.077L19.504 5.79m-14.456 0a48.108 48.108 0 00-3.478-.397m12 .562c.34-.059.68-.114 1.022-.165" />
            </svg>
            حذف
          </button>
          <button class="px-4 py-2 rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-400 shadow-md transition-all duration-200 font-semibold" @click="saveBrand">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 ml-2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 3.75H6.912a2.25 2.25 0 00-2.15 1.588L2.25 16.5l1.75 1.75 5.125-5.125M15 1.5l1.5 1.5m-1.5-1.5l-3.182 3.182A9.962 9.962 0 0121.75 8.25c0 2.278-.674 4.42-1.902 6.25" />
            </svg>
            ذخیره و ادامه ویرایش
          </button>
          <button class="px-4 py-2 rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-400 shadow-md transition-all duration-200 font-semibold">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 ml-2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 3.75H6.912a2.25 2.25 0 00-2.15 1.588L2.25 16.5l1.75 1.75 5.125-5.125M15 1.5l1.5 1.5m-1.5-1.5l-3.182 3.182A9.962 9.962 0 0121.75 8.25c0 2.278-.674 4.42-1.902 6.25" />
            </svg>
            ذخیره
          </button>
          <button 
            class="px-4 py-2 rounded-lg text-white bg-gradient-to-r from-purple-400 to-purple-600 hover:from-purple-500 hover:to-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-400 shadow-md transition-all duration-200 font-semibold"
            @click="previewBrand">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 ml-2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z" />
            </svg>
            پیش نمایش
          </button>
        </div>
      </div>
    </div>

    <!-- Container for Tabs and Back link -->
    <div class="bg-white rounded-lg shadow px-4 py-4 flex items-center justify-between border-b border-gray-200 mt-4 mb-6">
      <!-- Tabs -->
      <nav class="-mb-px flex" aria-label="Tabs">
        <a href="#" :class="[activeTab === 'info' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm ml-6']" @click.prevent="selectTab('info')">اطلاعات برند</a>
        <a href="#" :class="[activeTab === 'display' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm ml-6']" @click.prevent="selectTab('display')">تنظیمات نمایش برند در دسته بندی</a>
        <a href="#" :class="[activeTab === 'seo' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm ml-6']" @click.prevent="selectTab('seo')">سئو</a>
        <a href="#" :class="[activeTab === 'products' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm ml-6']" @click.prevent="selectTab('products')">محصولات</a>
        <a href="#" :class="[activeTab === 'images' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm ml-16']" @click.prevent="selectTab('images')">تصاویر مرتبط</a>
        <a href="#" :class="[activeTab === 'video' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']" @click.prevent="selectTab('video')">ویدیو</a>
      </nav>
      <!-- Back link -->
      <NuxtLink to="/admin/product-management/brands" class="flex items-center text-blue-600 hover:text-blue-800 text-sm">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4 ml-1">
          <path stroke-linecap="round" stroke-linejoin="round" d="M13.5 9.75L9.75 13.5m0 0l3.75 3.75M9.75 13.5H21" />
        </svg>
        برگشت به لیست برندها
      </NuxtLink>
    </div>

    <!-- Tab Content -->
    <div class="mt-6">
      <div v-if="activeTab === 'info'">
        <!-- Name and Description -->
        <div class="bg-white rounded-lg shadow px-4 py-4 mb-6">
          <!-- Name Field -->
          <div class="mb-6">
            <label for="brand-name" class="block text-sm font-medium text-gray-700 mb-1 text-right">نام</label>
            <input id="brand-name" v-model="name" type="text" class="w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50 bg-gray-50 p-2 text-right" dir="rtl">
          </div>
          <!-- Description Field -->
          <div>
            <label for="brand-description" class="block text-sm font-medium text-gray-700 mb-1 text-right">توضیحات</label>
            <ClientOnly>
              <RichTextEditor v-model="brandDescription" :lang="'fa'" :direction="'rtl'" :height="300" />
            </ClientOnly>
          </div>
        </div>
        <!-- Official Name & Image Upload -->
        <div class="bg-white rounded-lg shadow px-4 py-4">
          <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
            <div>
              <label for="official-name" class="block text-sm font-medium text-gray-700 mb-1 text-right">نام رسمی</label>
              <input id="official-name" v-model="officialName" type="text" class="w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50 focus:outline-blue-500 bg-gray-50 p-2 text-right" dir="rtl" @blur="onEnglishBlur">
            </div>
            <div>
              <label for="brand-image" class="block text-sm font-medium text-gray-700 mb-1 text-right">تصویر</label>
              <div class="flex items-center space-x-4 space-x-reverse">
                <img :src="imageUrl || '/statics/images/default-image_100.png'" alt="No Image" class="w-16 h-16 object-contain opacity-60 border rounded" />
                <div class="flex items-center">
                  <div class="flex gap-2">
                    <label for="brand-image-upload" class="bg-green-500 text-white rounded px-3 py-1 text-xs hover:bg-green-600 transition cursor-pointer">ارسال تصویر</label>
                    <button class="bg-blue-500 text-white rounded px-3 py-1 text-xs hover:bg-blue-600 transition">انتخاب از رسانه</button>
                  </div>
                  <span class="text-gray-600 mr-2" v-text="selectedFileName"></span>
                </div>
                <input id="brand-image-upload" type="file" class="hidden" @change="handleFileChange" />
              </div>
            </div>
          </div>
        </div>
      </div>

      <div v-else-if="activeTab === 'display'">
        <!-- Display Settings -->
        <div class="bg-gray-50 px-4 py-4 rounded-md shadow-md">
          <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
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

              <!-- Display Order -->
              <div class="flex items-center justify-start">
                <label for="display-order" class="block text-sm font-medium text-gray-700 ml-2">ترتیب نمایش</label>
                <input id="display-order" v-model="displayOrder" type="number" placeholder="0" class="w-24 rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50 bg-gray-50 p-2 text-sm focus:outline-none">
              </div>
            </div>

            <!-- Left Column -->
            <div class="space-y-6">
              <!-- Additional settings can be added here in future -->
            </div>
          </div>
        </div>
      </div>

      <div v-else-if="activeTab === 'seo'">
        <div class="bg-gray-50 px-4 py-4 rounded-md shadow-md">
          <ProductSeoTab :slug="seoSlug" @update:slug="updateSlug" />
        </div>
      </div>

      <div v-else-if="activeTab === 'products'">
        <!-- Products -->
        <div class="bg-gray-50 px-4 py-4 rounded-md shadow-md">
          <h3 class="text-lg font-medium text-gray-900 mb-4 text-right">محصولات</h3>
          <!-- Products content will go here -->
        </div>
      </div>

      <div v-else-if="activeTab === 'images'">
        <!-- Related Images -->
        <div class="bg-gray-50 px-4 py-4 rounded-md shadow-md">
          <h3 class="text-lg font-medium text-gray-900 mb-4 text-right">تصاویر مرتبط</h3>
          <!-- Related images content will go here -->
        </div>
      </div>

      <div v-else-if="activeTab === 'video'">
        <!-- Video -->
        <div class="bg-gray-50 px-4 py-4 rounded-md shadow-md">
          <h3 class="text-lg font-medium text-gray-900 mb-4 text-right">ویدیو</h3>
          <!-- Video content will go here -->
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useFetch: <T>(url: string, options?: unknown) => Promise<{ data: { value: T }; pending: { value: boolean }; error: { value: unknown }; refresh: () => Promise<void> }>
</script>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import RichTextEditor from '~/components/common/RichTextEditor.vue';
import ProductSeoTab from '~/components/product/ProductSeoTab.vue';

definePageMeta({
  layout: 'admin-main'
});

const activeTab = ref('info');

const selectTab = (tabName: string) => {
  activeTab.value = tabName;
};

// File upload
const selectedFileName = ref('');

const handleFileChange = (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files && target.files.length > 0) {
    selectedFileName.value = target.files[0].name;
  } else {
    selectedFileName.value = '';
  }
};

// Display settings reactive variables
const isPublished = ref(true);
const showOnHomepage = ref(false);
const showInMainMenu = ref(false);
const displayOrder = ref(0);

// brand description
const brandDescription = ref('');

// SEO slug handling
const seoSlug = ref('')
const slugTouched = ref(false)
let programmaticUpdate=false

function slugify(str){return str.toLowerCase().trim().replace(/[^a-z0-9\s-]/g,'').replace(/\s+/g,'-').replace(/-+/g,'-')}
function onEnglishBlur(){if(!slugTouched.value){programmaticUpdate=true;seoSlug.value=slugify(officialName.value||'');programmaticUpdate=false}}
function updateSlug(v){seoSlug.value=v;if(!programmaticUpdate)slugTouched.value=true}

const seoMetaTitle = ref('');
const seoMetaKeywords = ref('');
const seoMetaDescription = ref('');

const route = useRoute();

// Preview function
function previewBrand() {
  // Get the brand slug from the SEO tab or generate it from the name
  const slug = seoSlug.value || slugify(name.value);
  if (!slug) {
    alert('لطفاً نام برند را وارد کنید');
    return;
  }
  // Open brand page in new tab
  window.open(`/brand/${slug}`, '_blank');
}

// load brand data
onMounted(async () => {
  const { data } = await useFetch(`/api/brands`);
  const list = Array.isArray(data.value) ? data.value : [];
  const b = list.find((x:any) => String(x.id) === String(route.params.id));
  if(b){
    name.value = b.name;
    officialName.value = b.official_name || '';
    onEnglishBlur();
    brandDescription.value = b.description || '';
    isPublished.value = b.published;
    showOnHomepage.value = b.show_on_home;
    showInMainMenu.value = b.show_in_menu;
    seoMetaTitle.value = b.meta_title || '';
    seoMetaKeywords.value = b.meta_keywords || '';
    seoMetaDescription.value = b.meta_description || '';
    imageUrl.value = b.image_url || '';
    displayOrder.value = b.order || 0;
  }
});

// form refs defined earlier
const name = ref('');
const officialName = ref('');
const imageUrl = ref('');

async function saveBrand(){
  const id = route.params.id;
  try{
    await $fetch(`/api/brands/${id}`,{
      method:'PUT' as any,
      body:{
        name: name.value,
        official_name: officialName.value,
        description: brandDescription.value,
        published: isPublished.value,
        show_on_home: showOnHomepage.value,
        show_in_menu: showInMainMenu.value,
        meta_title: seoMetaTitle.value,
        meta_keywords: seoMetaKeywords.value,
        meta_description: seoMetaDescription.value,
        image_url: imageUrl.value
      }
    });
    alert('برند ذخیره شد');
  }catch{alert('خطا در ذخیره');}
}
</script>

<style scoped>
/* Add any specific styles here */
</style> 