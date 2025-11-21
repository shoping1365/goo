<template>
  <div class="w-full">
    <!-- Title Section -->
    <div class="bg-white w-full shadow-sm border-2 border-blue-200 rounded-lg mt-4">
      <div class="flex justify-between items-center px-8 py-6">
        <h1 class="text-2xl font-bold text-gray-800">ایجاد {{ fetchedWidget?.title || 'ابزارک' }}</h1>
        <div class="flex items-center gap-6">
          <button
            :disabled="isSaving"
            class="bg-green-500 text-white px-6 py-2 rounded-md font-bold hover:bg-green-600 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
            @click="saveWidget"
          >
            <svg v-if="isSaving" class="w-5 h-5 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
            </svg>
            <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
            {{ isSaving ? 'در حال ذخیره...' : 'ذخیره' }}
          </button>
          <TemplateButton
            bg-gradient="bg-gradient-to-r from-purple-400 to-purple-600"
            text-color="text-white"
            border-color="border border-purple-500"
            hover-class="hover:from-purple-500 hover:to-purple-700"
            size="medium"
            class="flex items-center gap-2"
            @click="$router.push('/admin/content/banners')"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
            </svg>
            بازگشت به ابزارک ها
          </TemplateButton>
        </div>
      </div>
    </div>

    <!-- Form Section -->
    <div class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
      <form class="grid grid-cols-1 md:grid-cols-5 gap-4 items-end">
        <!-- نوع ابزارک -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">نوع ابزارک</label>
          <select
            v-model="formData.type"
            class="w-full border border-gray-300 rounded-md px-2 py-1.5 text-sm focus:outline-none focus:ring-2 focus:ring-purple-400 bg-gray-100 cursor-not-allowed"
            disabled
            required
          >
            <option value="">انتخاب کنید</option>
            <option v-for="(label, type) in WIDGET_TYPE_LABELS" :key="type" :value="type">
              {{ label }}
            </option>
          </select>
        </div>

        <!-- عنوان ابزارک -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700" style="font-family: 'Yekan', sans-serif;">عنوان ابزارک *</label>
          <input
            v-model="formData.title"
            type="text"
            class="w-full border border-gray-300 rounded-md px-2 py-1.5 text-sm focus:outline-none focus:ring-2 focus:ring-purple-400"
            style="font-family: 'Yekan', sans-serif;"
            :class="{ 'border-red-500 focus:ring-red-500': titleError }"
            placeholder="عنوان ابزارک را وارد کنید"
            required
            @input="validateTitle"
            @blur="validateTitle"
          />
        </div>

        <!-- وضعیت -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">وضعیت</label>
          <select
            v-model="formData.status"
            class="w-full border border-gray-300 rounded-md px-2 py-1.5 text-sm focus:outline-none focus:ring-2 focus:ring-purple-400"
          >
            <option value="active">فعال</option>
            <option value="inactive">غیرفعال</option>
            <option value="draft">پیش‌نویس</option>
          </select>
        </div>

        <!-- صفحه -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">صفحه</label>
          <select
            v-model="formData.page"
            class="w-full border border-gray-300 rounded-md px-2 py-1.5 text-sm focus:outline-none focus:ring-2 focus:ring-purple-400"
          >
            <option value="home">صفحه اصلی</option>
            <option value="category">دسته‌بندی</option>
            <option value="product">محصول</option>
            <option value="about">درباره ما</option>
            <option value="contact">تماس با ما</option>
          </select>
        </div>

        <!-- نمایش در موبایل -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">نمایش در موبایل</label>
          <select
            v-model="formData.show_on_mobile"
            class="w-full border border-gray-300 rounded-md px-2 py-1.5 text-sm focus:outline-none focus:ring-2 focus:ring-purple-400"
          >
            <option :value="true">فعال</option>
            <option :value="false">غیرفعال</option>
          </select>
        </div>
      </form>

      <!-- Error Message for Title Field -->
      <div v-if="titleError" class="mt-2">
        <p class="text-sm text-red-600" style="font-family: 'Yekan', sans-serif;">
          {{ titleError }}
        </p>
      </div>
    </div>

    <!-- تنظیمات بنر -->
    <DeviceTabs
      ref="deviceTabsRef"
      :banner-config="bannerConfig"
      :current-preview-banner="currentPreviewBanner"
      :open-add-banner-modal="openAddBannerModal"
      :edit-banner="editBanner"
      :remove-banner="removeBanner"
      :get-banner-classes="() => ''"
      :get-banner-width="() => ''"
    >
      <!-- محتوای تب دسکتاپ -->
      <template #desktop-content>
        <div class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
          <div class="flex justify-between items-center mb-6">
            <h3 class="text-lg font-bold text-gray-700">تنظیمات بنر</h3>
          </div>
          
          <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-6">
            <!-- پس‌زمینه فعال -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">پس‌زمینه فعال</label>
              <select
                v-model="bannerConfig.bg_enabled"
                class="w-full border border-gray-300 rounded-md px-2 py-1.5 focus:outline-none focus:ring-2 focus:ring-purple-400 text-sm"
              >
                <option :value="true">فعال</option>
                <option :value="false">غیرفعال</option>
              </select>
            </div>

            <!-- عریض پس‌زمینه -->
            <div v-if="bannerConfig.bg_enabled">
              <label class="block mb-2 text-sm font-medium text-gray-700">عریض پس‌زمینه</label>
              <select
                v-model="bannerConfig.wide_bg"
                class="w-full border border-gray-300 rounded-md px-2 py-1.5 focus:outline-none focus:ring-2 focus:ring-purple-400 text-sm"
              >
                <option :value="true">بله</option>
                <option :value="false">خیر</option>
              </select>
            </div>

            <!-- رنگ پس‌زمینه -->
            <div v-if="bannerConfig.bg_enabled">
              <label class="block mb-2 text-sm font-medium text-gray-700">رنگ پس‌زمینه</label>
              <input
                v-model="bannerConfig.bg_color"
                type="color"
                class="w-full h-8 border border-gray-300 rounded-md"
              />
            </div>

            <!-- ارتفاع بنر -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">ارتفاع بنر (پیکسل)</label>
              <select
                v-model="bannerConfig.height"
                class="w-full border border-gray-300 rounded-md px-2 py-1.5 focus:outline-none focus:ring-2 focus:ring-purple-400 text-sm"
              >
                <option :value="300">300 پیکسل</option>
                <option :value="400">400 پیکسل</option>
                <option :value="500">500 پیکسل</option>
                <option :value="600">600 پیکسل</option>
                <option :value="700">700 پیکسل</option>
              </select>
            </div>

            <!-- عرض بنر -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">عرض بنر</label>
              <select
                v-model="bannerConfig.banner_width"
                class="w-full border border-gray-300 rounded-md px-2 py-1.5 focus:outline-none focus:ring-2 focus:ring-purple-400 text-sm"
              >
                <option :value="800">تمام عرض</option>
                <option :value="600">وسط</option>
              </select>
            </div>

            <!-- پدینگ بالا -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">پدینگ بالا (px)</label>
              <input
                type="number"
                :value="bannerConfig.padding_top === undefined ? '' : bannerConfig.padding_top"
                min="0"
                max="100"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
                placeholder="0"
                @input="(e: Event) => {
                  const target = e.target as HTMLInputElement
                  bannerConfig.padding_top = target.value === '' ? undefined : Number(target.value)
                }"
              />
            </div>

            <!-- پدینگ پایین -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">پدینگ پایین (px)</label>
              <input
                type="number"
                :value="bannerConfig.padding_bottom === undefined ? '' : bannerConfig.padding_bottom"
                min="0"
                max="100"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
                placeholder="0"
                @input="(e: Event) => {
                  const target = e.target as HTMLInputElement
                  bannerConfig.padding_bottom = target.value === '' ? undefined : Number(target.value)
                }"
              />
            </div>

            <!-- مارجین راست -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">مارجین راست (px)</label>
              <input
                type="number"
                :value="bannerConfig.margin_right === undefined ? '' : bannerConfig.margin_right"
                min="0"
                max="100"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
                placeholder="0"
                @input="(e: Event) => {
                  const target = e.target as HTMLInputElement
                  bannerConfig.margin_right = target.value === '' ? undefined : Number(target.value)
                }"
              />
            </div>

            <!-- مارجین چپ -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">مارجین چپ (px)</label>
              <input
                type="number"
                :value="bannerConfig.margin_left === undefined ? '' : bannerConfig.margin_left"
                min="0"
                max="100"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
                placeholder="0"
                @input="(e: Event) => {
                  const target = e.target as HTMLInputElement
                  bannerConfig.margin_left = target.value === '' ? undefined : Number(target.value)
                }"
              />
            </div>
          </div>
        </div>
      </template>

      <!-- محتوای تب موبایل -->
      <template #mobile-content>
        <div class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
          <div class="flex justify-between items-center mb-6">
            <h3 class="text-lg font-bold text-gray-700">تنظیمات موبایل</h3>
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-6">
            <!-- نمایش عمودی در موبایل -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">نمایش عمودی</label>
              <select
                v-model="bannerConfig.mobile_vertical_display"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              >
                <option :value="true">فعال</option>
                <option :value="false">غیرفعال</option>
              </select>
            </div>

            <!-- پدینگ بالا موبایل -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">پدینگ بالا موبایل (px)</label>
              <input
                type="number"
                :value="bannerConfig.mobile_padding_top === undefined ? '' : bannerConfig.mobile_padding_top"
                min="0"
                max="100"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
                placeholder="0"
                @input="(e: Event) => {
                  const target = e.target as HTMLInputElement
                  bannerConfig.mobile_padding_top = target.value === '' ? undefined : Number(target.value)
                }"
              />
            </div>

            <!-- پدینگ پایین موبایل -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">پدینگ پایین موبایل (px)</label>
              <input
                type="number"
                :value="bannerConfig.mobile_padding_bottom === undefined ? '' : bannerConfig.mobile_padding_bottom"
                min="0"
                max="100"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
                placeholder="0"
                @input="(e: Event) => {
                  const target = e.target as HTMLInputElement
                  bannerConfig.mobile_padding_bottom = target.value === '' ? undefined : Number(target.value)
                }"
              />
            </div>
          </div>
        </div>
      </template>
    </DeviceTabs>

    <!-- پیش‌نمایش زنده -->
    <div v-if="deviceTabsRef?.activeTab === 'desktop'" class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
      <h3 class="text-lg font-bold text-gray-700 mb-6">پیش‌نمایش زنده</h3>
      
      <!-- پیش‌نمایش بنرهای دوتایی -->
      <div class="bg-gray-50 rounded-lg p-6 border-2 border-dashed border-gray-300">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- بنر اول -->
          <div class="relative overflow-hidden rounded-lg shadow-lg">
            <div
              class="relative"
              :style="{
                height: `${bannerConfig.height}px`,
                backgroundColor: bannerConfig.bg_enabled ? bannerConfig.bg_color : 'transparent'
              }"
            >
              <img
                v-if="bannerConfig.banners && bannerConfig.banners.length > 0"
                :src="bannerConfig.banners[0].image"
                :alt="bannerConfig.banners[0].title"
                class="w-full h-full object-cover"
              />
              <div v-else class="flex items-center justify-center h-full text-gray-400">
                <div class="text-center">
                  <svg class="w-12 h-12 mx-auto mb-2 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                  </svg>
                  <p class="text-sm">بنر اول</p>
                </div>
              </div>
              <div
                v-if="bannerConfig.show_title || bannerConfig.show_description"
                class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
              >
                <h4
                  v-if="bannerConfig.show_title && bannerConfig.banners && bannerConfig.banners[0] && bannerConfig.banners[0].title"
                  class="text-white text-lg font-bold mb-2"
                >
                  {{ bannerConfig.banners[0].title }}
                </h4>
                <p
                  v-if="bannerConfig.show_description && bannerConfig.banners && bannerConfig.banners[0] && bannerConfig.banners[0].description"
                  class="text-white/90 text-sm"
                >
                  {{ bannerConfig.banners[0].description }}
                </p>
              </div>
            </div>
          </div>

          <!-- بنر دوم -->
          <div class="relative overflow-hidden rounded-lg shadow-lg">
            <div
              class="relative"
              :style="{
                height: `${bannerConfig.height}px`,
                backgroundColor: bannerConfig.bg_enabled ? bannerConfig.bg_color : 'transparent'
              }"
            >
              <img
                v-if="bannerConfig.banners && bannerConfig.banners.length > 1"
                :src="bannerConfig.banners[1].image"
                :alt="bannerConfig.banners[1].title"
                class="w-full h-full object-cover"
              />
              <div v-else class="flex items-center justify-center h-full text-gray-400">
                <div class="text-center">
                  <svg class="w-12 h-12 mx-auto mb-2 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                  </svg>
                  <p class="text-sm">بنر دوم</p>
                </div>
              </div>
              <div
                v-if="bannerConfig.show_title || bannerConfig.show_description"
                class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
              >
                <h4
                  v-if="bannerConfig.show_title && bannerConfig.banners && bannerConfig.banners[1] && bannerConfig.banners[1].title"
                  class="text-white text-lg font-bold mb-2"
                >
                  {{ bannerConfig.banners[1].title }}
                </h4>
                <p
                  v-if="bannerConfig.show_description && bannerConfig.banners && bannerConfig.banners[1] && bannerConfig.banners[1].description"
                  class="text-white/90 text-sm"
                >
                  {{ bannerConfig.banners[1].description }}
                </p>
              </div>
            </div>
          </div>
        </div>
        
        <!-- اطلاعات تنظیمات -->
        <div class="mt-4 text-sm text-gray-600 text-center">
          <p>ارتفاع: {{ bannerConfig.height }}px | 
              عرض: {{ bannerConfig.banner_width === 800 ? 'تمام عرض' : 'وسط صفحه' }} | 
              پس‌زمینه: {{ bannerConfig.bg_enabled ? 'فعال' : 'غیرفعال' }}</p>
        </div>
      </div>
    </div>

    <!-- نمایش و مدیریت بنرها -->
    <div v-if="deviceTabsRef?.activeTab === 'desktop'" class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
      <div class="flex justify-between items-center mb-6">
        <h3 class="text-lg font-bold text-gray-700">بنرها</h3>
        <button
          class="bg-cyan-500 text-white px-4 py-2 rounded-md font-bold hover:bg-cyan-600 transition-colors"
          @click="openAddBannerModal"
        >
          افزودن بنر جدید
        </button>
      </div>
      <div v-if="bannerConfig.banners.length === 0" class="text-gray-400 text-center py-8">
        چیزی برای نمایش وجود ندارد!
      </div>
      <div v-else class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div
          v-for="(banner, idx) in bannerConfig.banners"
          :key="idx"
          class="flex items-center gap-6 p-3 border-2 border-purple-200 rounded-lg bg-purple-50"
        >
          <img
            :src="banner.image"
            alt="بنر"
            class="w-32 h-20 object-cover rounded border-2 border-blue-300"
          />
          <div class="flex flex-col flex-1">
            <div class="font-bold text-sm text-gray-700 mb-1">{{ banner.title }}</div>
            <div v-if="banner.description" class="text-xs text-gray-600 mb-1">{{ banner.description }}</div>
            <div v-if="banner.link" class="text-xs text-blue-600 break-all">{{ banner.link }}</div>
          </div>
          <div class="flex gap-2">
            <button
              class="text-blue-500 hover:text-blue-700 p-1"
              @click="editBanner(idx)"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536M9 11l6 6M3 17v4h4l10.293-10.293a1 1 0 00-1.414-1.414L3 17z"></path>
              </svg>
            </button>
            <button
              class="text-red-500 hover:text-red-700 p-1"
              @click="removeBanner(idx)"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Banner Modal Component -->
    <BannerModal
      :is-visible="showBannerModal"
      :is-editing="editingBannerIndex !== null"
      :banner-data="editingBanner"
      :show-title="showTitleInBanner"
      @update:is-visible="showBannerModal = $event"
      @update:show-title="showTitleInBanner = $event"
      @update:banner-data="editingBanner = $event"
      @save="handleBannerSave"
      @open-media-library="openMediaLibrary"
      @remove-image="removeImage"
    />

    <!-- Media Library Modal -->
    <MediaLibraryModal
      v-model="showMediaLibrary"
      file-type="image"
      default-category="banners"
      @confirm="onSelectFromLibrary"
    />

    <!-- Page content will be added here -->

    <!-- فاصله انتهای صفحه -->
    <div class="pb-16"></div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import BannerModal from '~/components/common/BannerModal.vue'
import TemplateButton from '~/components/common/TemplateButton.vue'
import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue'
import { useToast } from '~/composables/useToast'
import { useWidget } from '~/composables/useWidget'
import type { BannerConfig, BannerItem, Widget, WidgetPage, WidgetStatus, WidgetType } from '~/types/widget'
import { WIDGET_TYPE_LABELS } from '~/types/widget'
import DeviceTabs from './components/DeviceTabs.vue'

interface MediaFile {
  id?: number | string
  url: string
  name?: string
  [key: string]: unknown
}

// تعریف definePageMeta، useHead برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void

definePageMeta({ layout: 'admin-main', middleware: 'admin' })
useHead({ title: 'ایجاد بنر دوتایی - پنل ادمین' })

// Route params
const route = useRoute()
const router = useRouter()
const widgetId = parseInt(route.params.id as string)

// Composables
const { fetchWidget, createWidget, loading: _loading, error: _error, clearError, widget: fetchedWidget } = useWidget()
const { showSuccess, showError } = useToast()

// Props
interface Props {
  widget?: Widget
}

const _props = defineProps<Props>()

// Emits
const _emit = defineEmits<{
  updated: [widget: Widget]
}>()

// State
const isSaving = ref(false)
const showBannerModal = ref(false)
const showMediaLibrary = ref(false)
const editingBannerIndex = ref<number | null>(null)
const currentPreviewBanner = ref(0)
const deviceTabsRef = ref()
const activeDeviceTab = ref<'desktop' | 'mobile'>('desktop')

// Error states
const titleError = ref<string>('')

// Validation functions
const validateTitle = () => {
  if (!formData.value.title?.trim()) {
    titleError.value = 'لطفاً عنوان ابزارک را وارد کنید.'
  } else if (formData.value.title.trim().length < 2) {
    titleError.value = 'عنوان ابزارک باید حداقل ۲ کاراکتر باشد.'
  } else if (formData.value.title.trim().length > 100) {
    titleError.value = 'عنوان ابزارک نمی‌تواند بیشتر از ۱۰۰ کاراکتر باشد.'
  } else {
    titleError.value = ''
  }
}

// Editing banner data
const editingBanner = ref<BannerItem>({
  title: '',
  description: '',
  image: '',
  link: '',
  order: 1,
  status: 'active'
})

// Title display control
const showTitleInBanner = ref(true)

// Form data
const formData = ref({
  title: '',
  description: '',
  type: 'double-banner',
  status: 'active',
  page: 'home',
  show_on_mobile: true
})

// Initialize form data when widget is available
const initializeFormData = () => {
  if (fetchedWidget.value) {
    formData.value = {
      title: fetchedWidget.value.title || '',
      description: fetchedWidget.value.description || '',
      type: fetchedWidget.value.type || 'double-banner',
      status: fetchedWidget.value.status || 'active',
      page: fetchedWidget.value.page || 'home',
      show_on_mobile: fetchedWidget.value.show_on_mobile !== undefined ? fetchedWidget.value.show_on_mobile : true
    }
  }
}

// Watch for widget changes
watch(fetchedWidget, (newWidget) => {
  if (newWidget) {
    initializeFormData()
  }
}, { immediate: true })

// Watch for device tab changes
watch(() => deviceTabsRef.value?.activeTab, (newTab) => {
  if (newTab) {
    activeDeviceTab.value = newTab
  }
})

// Computed properties for reactive form data (not currently used but may be needed for future features)
const _widgetTitle = computed(() => fetchedWidget.value?.title || '')
const _widgetType = computed(() => fetchedWidget.value?.type || 'double-banner')
const _widgetStatus = computed(() => fetchedWidget.value?.status || 'active')
const _widgetPage = computed(() => fetchedWidget.value?.page || 'home')

// Banner config
const bannerConfig = ref<BannerConfig>({
  wide_bg: false,
  bg_color: '#ffffff',
  height: 400,
  banners: [],
  mobile_banners: [],
  bg_enabled: false,
  banner_width: 800,
  show_title: true,
  show_description: true,
  padding_top: 0,
  padding_bottom: 0,
  margin_left: 0,
  margin_right: 0,
  // Mobile specific settings
  mobile_vertical_display: false,
  mobile_height: 150,
  mobile_slider_width: 300,
  // Mobile padding and margin settings
  mobile_padding_top: undefined,
  mobile_padding_bottom: undefined
})

// Methods for managing banners
const openAddBannerModal = () => {
  const banners = activeDeviceTab.value === 'mobile'
    ? bannerConfig.value.mobile_banners
    : bannerConfig.value.banners
  
  editingBannerIndex.value = null
  editingBanner.value = {
    title: '',
    description: '',
    image: '',
    mobile_image: '',
    link: '',
    order: banners.length + 1,
    status: 'active'
  }
  showTitleInBanner.value = true
  showBannerModal.value = true
}

const editBanner = (index: number) => {
  const banners = activeDeviceTab.value === 'mobile'
    ? bannerConfig.value.mobile_banners
    : bannerConfig.value.banners
  
  editingBannerIndex.value = index
  const banner = banners[index]
  editingBanner.value = {
    title: banner.title || '',
    description: banner.description || '',
    image: banner.image || '',
    mobile_image: banner.mobile_image || '',
    link: banner.link || '',
    order: banner.order || index + 1,
    status: banner.status || 'active'
  }
  showTitleInBanner.value = banner.showTitle !== undefined ? banner.showTitle : true
  showBannerModal.value = true
}

const removeBanner = (index: number) => {
  if (activeDeviceTab.value === 'mobile') {
    bannerConfig.value.mobile_banners.splice(index, 1)
    bannerConfig.value.mobile_banners.forEach((banner, idx) => {
      banner.order = idx + 1
    })
  } else {
    bannerConfig.value.banners.splice(index, 1)
    bannerConfig.value.banners.forEach((banner, idx) => {
      banner.order = idx + 1
    })
  }
}

// Handle banner save from modal component
const handleBannerSave = (bannerData: BannerItem, showTitle: boolean) => {
  const banners = activeDeviceTab.value === 'mobile'
    ? bannerConfig.value.mobile_banners
    : bannerConfig.value.banners
  
  const banner: BannerItem = {
    ...bannerData,
    title: bannerData.title?.trim() || '',
    description: bannerData.description?.trim() || '',
    link: bannerData.link?.trim() || '',
    order: editingBannerIndex.value !== null ? banners[editingBannerIndex.value].order : banners.length + 1,
    status: 'active',
    showTitle: showTitle
  }
  
  if (activeDeviceTab.value === 'mobile' && bannerData.mobile_image) {
    banner.image = bannerData.mobile_image
  }
  
  if (editingBannerIndex.value !== null) {
    banners[editingBannerIndex.value] = banner
  } else {
    banners.push(banner)
  }
  
  if (activeDeviceTab.value === 'mobile') {
    bannerConfig.value.mobile_banners = banners
  } else {
    bannerConfig.value.banners = banners
  }
  
  showBannerModal.value = false
  editingBannerIndex.value = null
}

// Media library functions
const openMediaLibrary = () => {
  showMediaLibrary.value = true
}

const onSelectFromLibrary = (files: MediaFile[]) => {
  if (files && files.length > 0) {
    const file = files[0]
    if (activeDeviceTab.value === 'mobile') {
      editingBanner.value.mobile_image = file.url
    } else {
      editingBanner.value.image = file.url
    }
  }
  showMediaLibrary.value = false
}

// Image functions
const removeImage = () => {
  editingBanner.value.image = ''
}

// اعمال برش خودکار موبایل هنگام ذخیره
const applyAutoMobileCrop = async () => {
  if (bannerConfig.value.mobile_image_mode !== 'auto') {
    return
  }
  
  if (!bannerConfig.value.banners.length) {
    return
  }
  
  const banner = bannerConfig.value.banners[0]
  if (!banner.image) {
    return
  }
  
  try {
    const cropWidth = bannerConfig.value.mobile_crop_width || 375
    const cropHeight = bannerConfig.value.mobile_crop_height || 150
    
    try {
      interface ImageCropResponse {
        success: boolean
        data?: {
          cropped_url: string
        }
      }
      
      const response = await $fetch<ImageCropResponse>('/api/media/image-crop', {
        method: 'POST',
        body: {
          src: banner.image,
          crop_width: cropWidth,
          crop_height: cropHeight,
          mode: 'cover',
          device: 'mobile',
          quality: 85
        }
      })
      
      if (response.success && response.data) {
        bannerConfig.value.mobile_cropped_image = response.data.cropped_url
      } else {
        throw new Error('API response failed')
      }
    } catch (_error) {
      return
    }
  } catch (_error) {
    // Silent fail
  }
}

// Save widget function
const saveWidget = async () => {
  validateTitle()
  
  if (titleError.value) {
    const titleInput = document.querySelector('input[placeholder="عنوان ابزارک را وارد کنید"]') as HTMLInputElement
    if (titleInput) {
      titleInput.focus()
      titleInput.scrollIntoView({ behavior: 'smooth', block: 'center' })
    }
    return
  }
  
  try {
    isSaving.value = true
    await applyAutoMobileCrop()
    
    const configToSave = { ...bannerConfig.value }
    configToSave.padding_top = bannerConfig.value.padding_top
    configToSave.padding_bottom = bannerConfig.value.padding_bottom
    configToSave.margin_left = bannerConfig.value.margin_left
    configToSave.margin_right = bannerConfig.value.margin_right
    configToSave.mobile_padding_top = bannerConfig.value.mobile_padding_top
    configToSave.mobile_padding_bottom = bannerConfig.value.mobile_padding_bottom
    
    const widgetData = {
      title: formData.value.title,
      description: formData.value.description,
      type: formData.value.type as WidgetType,
      status: formData.value.status as WidgetStatus,
      page: formData.value.page as WidgetPage,
      show_on_mobile: formData.value.show_on_mobile,
      config: configToSave
    }
    
    const createdWidget = await createWidget(widgetData)
    if (createdWidget && createdWidget.id) {
      showSuccess('ابزارک با موفقیت ایجاد شد!')
      const targetPath = `/admin/content/banners/DoubleBanner/edit/${createdWidget.id}`
      await router.push(targetPath)
    } else {
      showError('ابزارک ایجاد شد اما شناسه دریافت نشد')
    }
  } catch (_error) {
    showError('خطا در ذخیره ابزارک')
  } finally {
    isSaving.value = false
  }
}

// Initialize config from widget
onMounted(async () => {
  clearError()
  if (widgetId) {
    await fetchWidget(widgetId)
  }
  initializeFormData()
  
  if (fetchedWidget.value?.config) {
    const config = fetchedWidget.value.config as BannerConfig
    if (config.banners) bannerConfig.value.banners = config.banners
    if (config.height) bannerConfig.value.height = config.height
    if (config.banner_width) bannerConfig.value.banner_width = config.banner_width
    if (config.bg_enabled !== undefined) bannerConfig.value.bg_enabled = config.bg_enabled
    if (config.bg_color) bannerConfig.value.bg_color = config.bg_color
    if (config.wide_bg !== undefined) bannerConfig.value.wide_bg = config.wide_bg
    if (config.mobile_vertical_display !== undefined) bannerConfig.value.mobile_vertical_display = config.mobile_vertical_display
    if (config.mobile_height) bannerConfig.value.mobile_height = config.mobile_height
    if (config.mobile_padding_top !== undefined) bannerConfig.value.mobile_padding_top = config.mobile_padding_top
    if (config.mobile_padding_bottom !== undefined) bannerConfig.value.mobile_padding_bottom = config.mobile_padding_bottom
    if (config.padding_top !== undefined) bannerConfig.value.padding_top = config.padding_top
    if (config.padding_bottom !== undefined) bannerConfig.value.padding_bottom = config.padding_bottom
    if (config.margin_left !== undefined) bannerConfig.value.margin_left = config.margin_left
    if (config.margin_right !== undefined) bannerConfig.value.margin_right = config.margin_right
  }
})
</script>
