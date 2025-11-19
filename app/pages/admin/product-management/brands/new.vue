<template>
  <div>
    <!-- Header and Action Buttons -->
    <div class="flex justify-between items-center mb-6 bg-white rounded-lg shadow p-6">
      <div class="flex items-center">
        <h1 class="text-2xl font-bold text-gray-800 mr-4"> ایجاد برند جدید</h1>
      </div>
      <div class="flex flex-col items-end space-y-2">
        <div class="flex gap-2 flex-row-reverse">
          <button class="px-4 py-2 rounded-lg text-white bg-gradient-to-r from-purple-400 to-purple-600 hover:from-purple-500 hover:to-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-400 shadow-md transition-all duration-200 font-semibold">پیش نمایش</button>
          <button @click="saveBrand" class="px-4 py-2 rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-400 shadow-md transition-all duration-200 font-semibold">ذخیره</button>
          <button class="px-4 py-2 rounded-lg text-white bg-gradient-to-r from-red-400 to-red-600 hover:from-red-500 hover:to-red-700 focus:outline-none focus:ring-2 focus:ring-red-400 shadow-md transition-all duration-200 font-semibold">حذف</button>
        </div>
      </div>
    </div>

    <!-- Container for Tabs and Back link -->
    <div class="bg-white rounded-lg shadow p-6 flex items-center justify-between border-b border-gray-200 mt-4 mb-6">
      <!-- Tabs -->
      <nav class="-mb-px flex" aria-label="Tabs">
        <a href="#" @click.prevent="selectTab('info')" :class="[activeTab === 'info' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm ml-6']">اطلاعات برند</a>
        <a href="#" @click.prevent="selectTab('display')" :class="[activeTab === 'display' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm ml-6']">تنظیمات نمایش برند در دسته بندی</a>
        <a href="#" @click.prevent="selectTab('seo')" :class="[activeTab === 'seo' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm ml-6']">سئو</a>
        <a href="#" @click.prevent="selectTab('products')" :class="[activeTab === 'products' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm ml-6']">محصولات</a>
        <a href="#" @click.prevent="selectTab('images')" :class="[activeTab === 'images' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm ml-16']">تصاویر مرتبط</a>
        <a href="#" @click.prevent="selectTab('video')" :class="[activeTab === 'video' ? 'border-blue-600 text-blue-600 font-semibold' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300', 'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm']">ویدیو</a>
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
        <!-- Name and Description in a shared container -->
        <div class="bg-white rounded-lg shadow p-6 mb-6">
          <!-- Name Field -->
          <div class="mb-6">
            <label for="brand-name" class="block text-sm font-medium text-gray-700 mb-1 text-right">نام</label>
            <input v-model="name" type="text" id="brand-name" class="w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50 bg-gray-50 p-2 text-right" dir="rtl">
          </div>
          <!-- Description Field -->
          <div>
            <label for="brand-description" class="block text-sm font-medium text-gray-700 mb-1 text-right">توضیحات</label>
            <ClientOnly>
              <RichTextEditor v-model="brandDescription" :lang="'fa'" :direction="'rtl'" :height="300" />
            </ClientOnly>
          </div>
        </div>
        <!-- Official Name and Image Upload in a separate container -->
        <div class="bg-white rounded-lg shadow p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label for="official-name" class="block text-sm font-medium text-gray-700 mb-1 text-right">نام رسمی</label>
              <input @blur="onEnglishBlur" v-model="officialName" type="text" id="official-name" class="w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50 focus:outline-blue-500 bg-gray-50 p-2 text-right" dir="rtl">
            </div>
            <div>
              <label for="brand-image" class="block text-sm font-medium text-gray-700 mb-1 text-right">تصویر</label>
              <div class="flex items-center space-x-4 space-x-reverse">
                <div class="flex items-center">
                  <!-- Buttons container -->
                  <div class="flex gap-2">
                    <label for="brand-image-upload" class="bg-green-500 text-white rounded px-3 py-1 text-xs hover:bg-green-600 transition cursor-pointer">
                      ارسال تصویر
                    </label>
                    <button class="bg-blue-500 text-white rounded px-3 py-1 text-xs hover:bg-blue-600 transition">انتخاب از رسانه</button>
                  </div>
                  <span class="text-gray-600 mr-2" v-text="selectedFileName"></span>
                </div>
                <input type="file" id="brand-image-upload" @change="handleFileChange" class="hidden"/>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div v-else-if="activeTab === 'display'">
        <!-- Display Settings -->
        <div class="bg-gray-50 p-6 rounded-md shadow-md">
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

              <!-- Display Order -->
              <div class="flex items-center justify-start">
                <label for="display-order" class="block text-sm font-medium text-gray-700 ml-2">ترتیب نمایش</label>
                <input type="number" id="display-order" v-model="displayOrder" placeholder="0" class="w-24 rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50 bg-gray-50 p-2 text-sm focus:outline-none">
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
        <div class="bg-gray-50 p-6 rounded-md shadow-md">
          <ProductSeoTab :slug="seoSlug" @update:slug="updateSlug" />
        </div>
      </div>

      <div v-else-if="activeTab === 'products'">
        <!-- Products -->
        <div class="bg-gray-50 p-6 rounded-md shadow-md">
          <h3 class="text-lg font-medium text-gray-900 mb-4 text-right">محصولات</h3>
          <!-- Products content will go here -->
        </div>
      </div>

      <div v-else-if="activeTab === 'images'">
        <!-- Related Images -->
        <div class="bg-gray-50 p-6 rounded-md shadow-md">
          <h3 class="text-lg font-medium text-gray-900 mb-4 text-right">تصاویر مرتبط</h3>
          <!-- Related images content will go here -->
        </div>
      </div>

      <div v-else-if="activeTab === 'video'">
        <!-- Video -->
        <div class="bg-gray-50 p-6 rounded-md shadow-md">
          <h3 class="text-lg font-medium text-gray-900 mb-1 text-right">ویدیو</h3>
          <!-- Video content will go here -->
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const navigateTo: (to: string) => Promise<void>
</script>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
// useFetch is auto-imported in Nuxt 3
import RichTextEditor from '~/components/common/RichTextEditor.vue';
import ProductSeoTab from '~/components/product/ProductSeoTab.vue';

definePageMeta({
  layout: 'admin-main'
});

const activeTab = ref('info');
const router = useRouter();

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
const showBrandLogo = ref(false);
const displayOrder = ref(0);

// SEO slug handling like categories
const seoSlug = ref('')
const slugTouched = ref(false)
let programmaticUpdate = false

function slugify(str){
  return str.toLowerCase().trim().replace(/[^a-z0-9\s-]/g,'').replace(/\s+/g,'-').replace(/-+/g,'-')
}

function onEnglishBlur(){
  if(!slugTouched.value){
    programmaticUpdate=true
    seoSlug.value = slugify(officialName.value||'')
    programmaticUpdate=false
  }
}

function updateSlug(val){
  seoSlug.value = val
  if(!programmaticUpdate) slugTouched.value=true
}

// SEO meta extras (for future)
const seoPageName = ref('');
const seoMetaTitle = ref('');
const seoMetaKeywords = ref('');
const seoMetaDescription = ref('');

// Rich description
const brandDescription = ref('')

// Form reactive variables
const name = ref('');
const officialName = ref('');
const imageUrl = ref('');
const videoUrl = ref('');

async function saveBrand() {
  if (!name.value.trim()) {
    alert('نام برند الزامی است');
    return;
  }

  try {
    await $fetch('/api/brands', {
      method: 'POST',
      body: {
        name: name.value,
        description: brandDescription.value,
        name_en: officialName.value,
        slug: seoSlug.value,
        published: isPublished.value,
        show_on_home: showOnHomepage.value,
        show_in_menu: showInMainMenu.value,
        meta_title: seoMetaTitle.value,
        meta_keywords: seoMetaKeywords.value,
        meta_description: seoMetaDescription.value,
        image_url: imageUrl.value,
        video_url: videoUrl.value
      },
    });

    alert('برند با موفقیت ایجاد شد');
    router.push('/admin/product-management/brands');
  } catch (err) {
    console.error(err);
    alert('خطا در ایجاد برند');
  }
}
</script>

<style scoped>
/* Add any specific styles here */
</style> 
