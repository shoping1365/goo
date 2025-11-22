<template>
  <div class="space-y-8" dir="rtl">
    
    <!-- 1. مدیریت ویژگی‌های جدید -->
    <div class="bg-gradient-to-br from-blue-50 to-cyan-50 rounded-2xl border border-blue-200 shadow-lg overflow-hidden">
      <div class="bg-gradient-to-r from-blue-600 to-cyan-600 px-6 py-4 cursor-pointer" @click="toggleSection('technicalSpecs')">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <div class="p-2 bg-white/20 rounded-xl">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4"></path>
              </svg>
            </div>
            <h3 class="text-xl font-bold text-white">مدیریت ویژگی‌های جدید</h3>
          </div>
          <div class="p-2 bg-white/20 rounded-lg">
            <span class="text-white font-bold text-lg">{{ sections.technicalSpecs ? '−' : '+' }}</span>
          </div>
        </div>
      </div>

      <div v-show="sections.technicalSpecs" class="p-8">
        <div class="grid grid-cols-1 gap-6">
          <!-- Row 1: Basic Info -->
          <div class="grid grid-cols-12 gap-6 items-end">
            <div class="col-span-4">
              <label class="block text-sm font-semibold text-gray-700 mb-2">نام ویژگی *</label>
              <div class="relative">
                <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                  <svg class="w-5 h-5 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                  </svg>
                </div>
                <input 
                  v-model="newAttr.name" 
                  type="text" 
                  class="w-full pr-10 pl-4 py-3 bg-white border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-gray-900 placeholder-gray-500 transition-all duration-200" 
                  placeholder="مثلاً رنگ، سایز، نوع" 
                />
              </div>
            </div>

            <div class="col-span-4">
              <label class="block text-sm font-semibold text-gray-700 mb-2">متن نمایشی</label>
              <div class="relative">
                <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                  <svg class="w-5 h-5 text-emerald-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 4V2a1 1 0 011-1h8a1 1 0 011 1v2m3 0V2a1 1 0 00-1-1H8a1 1 0 00-1 1v2H4a1 1 0 00-1 1v16a1 1 0 001 1h16a1 1 0 001-1V5a1 1 0 00-1-1h-3zM9 12h6m-6 4h4"></path>
                  </svg>
                </div>
                <input 
                  v-model="newAttr.display_name" 
                  type="text" 
                  class="w-full pr-10 pl-4 py-3 bg-white border border-gray-300 rounded-xl focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 text-gray-900 placeholder-gray-500 transition-all duration-200" 
                  placeholder="نام نمایشی برای کاربران" 
                />
              </div>
            </div>

            <div class="col-span-2">
              <label class="block text-sm font-semibold text-gray-700 mb-2">واحد اندازه‌گیری</label>
              <div class="relative">
                <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                  <svg class="w-5 h-5 text-orange-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
                  </svg>
                </div>
                <input 
                  v-model="newAttr.unit" 
                  type="text" 
                  class="w-full pr-10 pl-4 py-3 bg-white border border-gray-300 rounded-xl focus:ring-2 focus:ring-orange-500 focus:border-orange-500 text-gray-900 placeholder-gray-500 transition-all duration-200" 
                  placeholder="مثلاً cm, kg" 
                />
              </div>
            </div>

            <div class="col-span-2">
              <label class="block text-sm font-semibold text-gray-700 mb-2">ترتیب نمایش</label>
              <div class="relative">
                <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                  <svg class="w-5 h-5 text-purple-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4"></path>
                  </svg>
                </div>
                <input 
                  v-model.number="newAttr.sort_order" 
                  type="number" 
                  class="w-full pr-10 pl-4 py-3 bg-white border border-gray-300 rounded-xl focus:ring-2 focus:ring-purple-500 focus:border-purple-500 text-gray-900 placeholder-gray-500 transition-all duration-200" 
                  placeholder="0" 
                />
              </div>
            </div>
          </div>

          <!-- Row 2: Type & Control -->
          <div class="grid grid-cols-12 gap-6 items-end">
            <div class="col-span-3">
              <label class="block text-sm font-semibold text-gray-700 mb-2">نوع ویژگی *</label>
              <div class="relative">
                <select 
                  v-model="newAttr.type" 
                  class="w-full px-4 py-3 bg-white border border-gray-300 rounded-xl focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-900 transition-all duration-200 appearance-none"
                >
                  <option value="select">انتخاب از لیست</option>
                  <option value="custom_text">متن سفارشی</option>
                </select>
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                  </svg>
                </div>
              </div>
            </div>

            <div class="col-span-5">
              <label class="block text-sm font-semibold text-gray-700 mb-2">نوع کنترل نمایش *</label>
              <div class="relative">
                <select 
                  v-model="newAttr.control_type" 
                  class="w-full px-4 py-3 bg-white border border-gray-300 rounded-xl focus:ring-2 focus:ring-teal-500 focus:border-teal-500 text-gray-900 transition-all duration-200 appearance-none"
                >
                  <option value="dropdown_single">لیست باز شونده تک انتخابی</option>
                  <option value="dropdown_multi">لیست باز شونده چند انتخابی</option>
                  <option value="textbox_multiline">متن چند خطی</option>
                </select>
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                  </svg>
                </div>
              </div>
            </div>

            <div class="col-span-4">
              <button 
                :disabled="creatingAttr"
                class="w-full inline-flex items-center justify-center px-6 py-3 border border-transparent text-sm font-semibold rounded-xl text-white bg-gradient-to-r from-blue-600 to-cyan-600 hover:from-blue-700 hover:to-cyan-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-lg transition-all duration-200 hover:shadow-xl hover:scale-105 disabled:opacity-60 disabled:cursor-not-allowed disabled:transform-none"
                @click="createAttribute"
              >
                <svg v-if="!creatingAttr" class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
                </svg>
                <svg v-else class="animate-spin w-5 h-5 ml-2" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="m4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                {{ creatingAttr ? 'در حال ذخیره...' : 'افزودن ویژگی' }}
              </button>
            </div>
          </div>

          <!-- Row 3: Options -->
          <div class="grid grid-cols-12 gap-6">
            <div class="col-span-12">
              <h4 class="text-sm font-semibold text-gray-700 mb-3">تنظیمات اضافی:</h4>
              <div class="grid grid-cols-4 gap-6">
                <label class="flex items-center gap-2 p-3 bg-blue-50 rounded-xl border border-blue-200 cursor-pointer hover:bg-blue-100 transition-colors">
                  <input v-model="newAttr.has_filter" type="checkbox" class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2" />
                  <span class="text-sm font-medium text-blue-700">ابزار فیلتر</span>
                </label>
                <label class="flex items-center gap-2 p-3 bg-emerald-50 rounded-xl border border-emerald-200 cursor-pointer hover:bg-emerald-100 transition-colors">
                  <input v-model="newAttr.is_key" type="checkbox" class="w-4 h-4 text-emerald-600 bg-gray-100 border-gray-300 rounded focus:ring-emerald-500 focus:ring-2" />
                  <span class="text-sm font-medium text-emerald-700">مشخصه کلیدی</span>
                </label>
                <label class="flex items-center gap-2 p-3 bg-purple-50 rounded-xl border border-purple-200 cursor-pointer hover:bg-purple-100 transition-colors">
                  <input v-model="newAttr.show_on_product" type="checkbox" class="w-4 h-4 text-purple-600 bg-gray-100 border-gray-300 rounded focus:ring-purple-500 focus:ring-2" />
                  <span class="text-sm font-medium text-purple-700">نمایش روی صفحه محصول</span>
                </label>
                <label class="flex items-center gap-2 p-3 bg-red-50 rounded-xl border border-red-200 cursor-pointer hover:bg-red-100 transition-colors">
                  <input v-model="newAttr.is_required" type="checkbox" class="w-4 h-4 text-red-600 bg-gray-100 border-gray-300 rounded focus:ring-red-500 focus:ring-2" />
                  <span class="text-sm font-medium text-red-700">ضروری</span>
                </label>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 2. فراخوانی و مدیریت گروه‌های ویژگی -->
    <div class="bg-gradient-to-br from-emerald-50 to-teal-50 rounded-2xl border border-emerald-200 shadow-lg overflow-hidden">
      <div class="bg-gradient-to-r from-emerald-600 to-teal-600 px-6 py-4">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 rounded-xl">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="p-8">
        <div class="text-center mb-6">
          <button 
            :disabled="loadingGroups"
            class="inline-flex items-center px-6 py-3 border border-transparent text-sm font-semibold rounded-xl text-blue-700 bg-white border-blue-200 hover:bg-blue-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105 disabled:opacity-60 disabled:cursor-not-allowed disabled:transform-none"
            @click="fetchAttributeGroups"
          >
            <svg v-if="!loadingGroups" class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4-4m0 0L8 8m4-4v12"></path>
            </svg>
            <svg v-else class="animate-spin w-5 h-5 ml-2" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="m4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            {{ loadingGroups ? 'در حال دریافت...' : 'فراخوانی گروه‌های ویژگی' }}
          </button>
        </div>



        <!-- Attributes Table -->
        <div v-if="selectedGroupId && groupAttributes.length > 0" class="mt-6 max-w-7xl mx-auto">
          <div class="bg-white rounded-xl shadow-lg overflow-hidden border border-blue-200">
            <div class="bg-gradient-to-r from-blue-600 to-indigo-600 px-4 py-3">
              <h4 class="text-lg font-bold text-white flex items-center gap-2">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                </svg>
                مدیریت ویژگی‌های «{{ selectedGroupName }}»
              </h4>
            </div>
            
            <div class="overflow-x-auto">
              <table class="w-full">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-1 py-2 text-right text-xs font-medium text-gray-500 uppercase tracking-wider w-20">نوع</th>
                    <th class="px-1 py-2 text-right text-xs font-medium text-gray-500 uppercase tracking-wider w-32">ویژگی</th>
                    <th class="px-4 py-2 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مقدار</th>
                    <th class="px-4 py-2 text-right text-xs font-medium text-gray-500 uppercase tracking-wider w-28">عملیات</th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-for="attr in groupAttributes" :key="attr.id" class="hover:bg-gray-50">
                    <td class="px-1 py-2 whitespace-nowrap">
                      <span class="inline-flex items-center justify-center px-3 py-1.5 text-xs font-semibold rounded-full bg-blue-100 text-blue-800 w-24">
                        {{ attr.type === 'custom_text' ? 'متن سفارشی' : 'انتخاب از لیست' }}
                      </span>
                    </td>
                    <td class="px-1 py-2 whitespace-nowrap">
                      <div class="flex items-center">
                        <div class="flex-shrink-0 h-8 w-8">
                          <div class="h-8 w-8 rounded-full bg-emerald-100 flex items-center justify-center">
                            <svg class="w-4 h-4 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                            </svg>
                          </div>
                        </div>
                        <div class="mr-1">
                          <div class="text-sm font-medium text-gray-900">
                            {{ attr.name }}
                            <span v-if="attr.isRequired" class="text-red-500 mr-1">*</span>
                          </div>
                        </div>
                      </div>
                    </td>
                    <td class="px-4 py-2">
                      <!-- Custom Text Input -->
                      <div v-if="attr.type === 'custom_text'" class="w-full">
                        <textarea 
                          v-model="attr.valueText"
                          rows="4"
                          class="w-full px-3 py-2 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 text-gray-900 resize-none"
                          :placeholder="`مقدار ${attr.name}`"
                          @keydown="handleKeydown"
                          @keypress="handleKeypress"
                        ></textarea>
                      </div>

                      <!-- Multi Select -->
                      <div v-else-if="attr.type === 'multi_select'" class="w-full">
                        <div class="flex flex-wrap gap-2">
                          <label 
                            v-for="val in attr.values" 
                            :key="val.id"
                            class="inline-flex items-center"
                          >
                            <input 
                              v-model="attr.selectedOptionIds" 
                              type="checkbox"
                              :value="val.id"
                              class="rounded border-gray-300 text-emerald-600 shadow-sm focus:border-emerald-300 focus:ring focus:ring-emerald-200 focus:ring-opacity-50"
                            />
                            <span class="mr-2 text-sm">{{ val.value }}</span>
                          </label>
                        </div>
                      </div>

                      <!-- Single Select -->
                      <div v-else class="w-full">
                        <div class="relative">
                          <select 
                            v-model="attr.selectedOptionId"
                            class="w-full px-3 py-2 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 text-gray-900 appearance-none"
                          >
                            <option value="">انتخاب کنید...</option>
                            <option v-for="val in attr.values" :key="val.id" :value="val.id">
                              {{ val.value }}
                            </option>
                          </select>
                          <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                            <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                            </svg>
                          </div>
                        </div>
                      </div>
                    </td>
                    <td class="px-4 py-2 whitespace-nowrap text-sm font-medium">
                      <div class="flex flex-col gap-1">
                        <button 
                          v-if="attr.type !== 'custom_text'"
                          class="px-3 py-2 text-sm font-medium text-emerald-700 bg-emerald-100 border border-emerald-300 rounded-lg hover:bg-emerald-200 focus:outline-none focus:ring-2 focus:ring-emerald-500 transition-colors"
                          title="افزودن گزینه جدید"
                          @click="addNewOption(attr)"
                        >
                          افزودن ایتم
                        </button>
                        <button 
                          class="px-3 py-2 text-sm font-medium text-red-600 bg-red-100 border border-red-300 rounded-lg hover:bg-red-200 hover:text-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 transition-colors"
                          title="حذف ویژگی"
                          @click="removeAttribute(attr)"
                        >
                          حذف
                        </button>
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>

        <!-- Empty State -->
        <div v-else-if="selectedGroupId && groupAttributes.length === 0" class="text-center py-12">
          <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          <h3 class="mt-2 text-sm font-medium text-gray-900">هیچ ویژگی‌ای یافت نشد</h3>
          <p class="mt-1 text-sm text-gray-500">برای این گروه ویژگی موردی تعریف نشده است.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { useNuxtApp } from 'nuxt/app'
import { reactive, ref, watch, watchEffect } from 'vue'
import { useAttributeGroup } from '~/composables/useAttributeGroup'
import { useConfirmDialog } from '~/composables/useConfirmDialog'
import { useNotifier } from '~/composables/useNotifier'
import { useProductCreateStore } from '~/stores/productCreate'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

const sections = reactive({
  technicalSpecs: true,
  groupsContainer: false
})

// const showFeaturesList = ref(false)

interface ApiClient {
  (url: string, options?: { method?: string; body?: unknown; [key: string]: unknown }): Promise<unknown>
  get?: <T>(url: string) => Promise<T>
  post?: <T>(url: string, body?: unknown) => Promise<T>
  put?: <T>(url: string, body?: unknown) => Promise<T>
  delete?: <T>(url: string) => Promise<T>
}

interface NuxtApp {
  $api: ApiClient
  [key: string]: unknown
}

const nuxtApp = useNuxtApp() as unknown as NuxtApp
const api = nuxtApp.$api
const pStore = useProductCreateStore()
const notifier = useNotifier()
const { confirm } = useConfirmDialog()

const attributeGroups = ref([])
const loadingGroups = ref(false)
const showGroupContainer = ref(false)
const selectedGroupId = ref('')
const selectedGroupName = ref('')

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const { attributes: groupAttributes, loadingGroup: loadingAttrGroup, loadGroup } = useAttributeGroup()

const newAttr = reactive({
  name: '',
  display_name: '',
  unit: '',
  sort_order: 0,
  type: 'select' as 'select' | 'custom_text',
  control_type: 'dropdown_single' as 'dropdown_single' | 'dropdown_multi' | 'textbox_multiline',
  has_filter: false,
  is_key: false,
  show_on_product: true,
  is_required: false
})

const creatingAttr = ref(false)

function syncSpecsToStore(){
  const list = groupAttributes.value.map(attr=>{
    if(attr.type==='multi_select'){
      if(Array.isArray(attr.selectedOptionIds) && attr.selectedOptionIds.length>0){
        return {attribute_id:attr.id, option_ids:attr.selectedOptionIds}
      }
      return null
    }
    if(attr.type!=='custom_text'){
      if(typeof attr.selectedOptionId==='number' && !Number.isNaN(attr.selectedOptionId)){
        return {attribute_id:attr.id, option_id:attr.selectedOptionId}
      }
      return null
    }
    const txt = attr.valueText ? String(attr.valueText).trim() : ''
    return txt ? {attribute_id:attr.id, value_text:txt} : null
  }).filter(Boolean)
  if(list.length>0) pStore.productSpecs = list
}

watchEffect(syncSpecsToStore)

// Apply persisted specs into UI whenever both lists are ready
watchEffect(() => {
  if (!pStore.productSpecs || pStore.productSpecs.length === 0) return
  for (const attr of groupAttributes.value) {
    const found = pStore.productSpecs.find(s => s.attribute_id === attr.id)
    if (found) {
      if(attr.type==='multi_select'){
        attr.selectedOptionIds = found.option_ids || []
      }else if (attr.type !== 'custom_text') {
        attr.selectedOptionId = (typeof found.option_id === 'number' && !Number.isNaN(found.option_id) && found.option_id !== 0)
          ? found.option_id
          : ''
      } else {
        attr.valueText = found.value_text || ''
      }
    }
  }
})

watch(
  [() => pStore.productForm.category_id, () => pStore.productForm.sub_category_id],
  () => {
    if (pStore.editingProductId) {
      fetchAttributeGroups()
    } else {
      attributeGroups.value = []
      showGroupContainer.value = false
    }
  }
)

async function fetchAttributeGroups() {
  const categoryId = pStore.productForm.sub_category_id || pStore.productForm.category_id

  if (!categoryId) {
    notifier.warning('ابتدا دسته‌بندی اصلی را در تب اطلاعات محصول انتخاب کنید و سپس ذخیره نمایید.')
    return
  }

  if (!pStore.editingProductId) {
    notifier.warning('ابتدا روی «ذخیره و ادامه ویرایش» کلیک کنید تا محصول ذخیره شود، سپس این عملیات را انجام دهید.')
    return
  }

  loadingGroups.value = true
  try {
    const res = await api('/api/attribute-groups', {
      params: { category_id: categoryId, per_page: 200 }
    })

    const response = res as { data?: unknown[] } | unknown[]
    if (response && typeof response === 'object' && 'data' in response && Array.isArray(response.data)) {
      attributeGroups.value = response.data
    } else if (Array.isArray(response)) {
      attributeGroups.value = response
    } else {
      attributeGroups.value = []
    }

    if (attributeGroups.value.length > 0) {
      selectGroup(attributeGroups.value[0])
      sections.groupsContainer = true
    }

    showGroupContainer.value = true
  } catch (e) {
    console.error('Error fetching attribute groups', e)
    notifier.error('خطا در دریافت گروه ویژگی‌ها')
  } finally {
    loadingGroups.value = false
  }
}

const toggleSection = (section) => {
  sections[section] = !sections[section]
}

// Force enable space key for text inputs
function handleKeydown(e: KeyboardEvent) {
  if (e.code === 'Space' || e.keyCode === 32) {
    e.stopPropagation()
    e.stopImmediatePropagation()
  }
}

function handleKeypress(e: KeyboardEvent) {
  if (e.code === 'Space' || e.keyCode === 32) {
    e.stopPropagation()
    e.stopImmediatePropagation()
  }
}

async function selectGroup(g){
  selectedGroupId.value = g.id
  selectedGroupName.value = g.name
  await loadGroup(g.id)
  // Prefill selections right after data loaded
  if (pStore.productSpecs && pStore.productSpecs.length>0){
    groupAttributes.value.forEach(attr=>{
      const found = pStore.productSpecs.find(s=>s.attribute_id===attr.id)
      if(found){
        if(attr.type!=='custom_text'){
          attr.selectedOptionId = (typeof found.option_id === 'number' && !Number.isNaN(found.option_id) && found.option_id !== 0)
            ? found.option_id
            : ''
        }else{
          attr.valueText = found.value_text || ''
        }
      }
    })
  }
}

async function addNewOption(attr){
  if(attr.type==='custom_text') return
  const val = prompt(`مقدار جدید برای "${attr.name}" را وارد کنید:`)
  if(!val || !val.trim()) return
  try{
    interface AttributeValueResponse {
      id?: number | string
      value?: string
      attribute_id?: number | string
      [key: string]: unknown
    }
    const res = await api(`/api/attribute-values/by-attribute/${attr.id}`, {method: 'POST', body: {value: val.trim()}}) as AttributeValueResponse
    if(res && res.id){
      const newVal = {id: Number(res.id), value: val.trim()}
      if(Array.isArray(attr.values)){
        attr.values.push(newVal)
      }else{
        attr.values = [newVal]
      }
      attr.selectedOptionId = newVal.id
      syncSpecsToStore()
    }
  }catch(e){
    notifier.error('خطا در افزودن مقدار')
    console.error(e)
  }
}

async function removeAttribute(attr){
  const ok = await confirm({
    title: 'تأیید حذف',
    message: `با حذف این ویژگی، از گروه ویژگی «${selectedGroupName.value}» نیز حذف خواهد شد. آیا مطمئن هستید؟`,
    confirmText: 'حذف',
    cancelText: 'انصراف',
    type: 'danger'
  })
  if(!ok) return
  try{
    await api(`/api/attribute-groups/${selectedGroupId.value}/attributes/${attr.id}`, {method:'DELETE'})
    groupAttributes.value = groupAttributes.value.filter(a=>a.id!==attr.id)
    syncSpecsToStore()
  }catch(e){
    notifier.error('خطا در حذف ویژگی از گروه')
    console.error(e)
  }
}

async function createAttribute(){
  if(!newAttr.name.trim()){
    notifier.warning('نام ویژگی الزامی است')
    return
  }
  creatingAttr.value = true
  try{
    // 1) create attribute record
    const payloadAttr = {
      name: newAttr.name.trim(),
      display_name: newAttr.display_name.trim() || newAttr.name.trim(),
      unit: newAttr.unit.trim(),
      data_type: 'string',
      sort_order: newAttr.sort_order,
      is_required: newAttr.is_required,
      is_filterable: newAttr.has_filter,
      is_active: true
    }
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const attrRes:any = await api('/api/attributes', {method:'POST', body: payloadAttr})
    if(!attrRes || !attrRes.id){throw new Error('خطا در ایجاد ویژگی')}

    // 2) add to group via upsert (send full list incl new)
    const updatedList = [
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      ...groupAttributes.value.map((a:any)=>({
        attribute_id: a.id,
        sort_order: a.sort_order || 0,
        type: a.type,
        control_type: a.controlType,
        has_filter: a.has_filter || false,
        is_key: a.is_key || false,
        show_on_product: a.show_on_product ?? true,
        is_required: a.isRequired || false
      })),
      {
        attribute_id: attrRes.id,
        sort_order: newAttr.sort_order,
        type: newAttr.type,
        control_type: newAttr.control_type,
        has_filter: newAttr.has_filter,
        is_key: newAttr.is_key,
        show_on_product: newAttr.show_on_product,
        is_required: newAttr.is_required
      }
    ]
    await api(`/api/attribute-groups/${selectedGroupId.value}/attributes`, {method:'PUT', body:{attributes: updatedList}})

    // 3) update UI
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const attrObj:any = {
      id: attrRes.id,
      name: attrRes.name,
      type: newAttr.type,
      controlType: newAttr.control_type,
      isRequired: newAttr.is_required,
      values: [],
      selectedOptionId: '',
      valueText: ''
    }
    groupAttributes.value.push(attrObj)

    // reset form
    newAttr.name=''; newAttr.display_name=''; newAttr.unit=''; newAttr.sort_order=0;
    newAttr.type='select'; newAttr.control_type='dropdown_single'; newAttr.has_filter=false; newAttr.is_key=false; newAttr.show_on_product=true; newAttr.is_required=false

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  }catch(e:any){
    notifier.error(e?.error || e?.message || 'خطا در افزودن ویژگی')
    console.error(e)
  }finally{
    creatingAttr.value=false
  }
}

// ---------- DEBUG -----------
// eslint-disable-next-line @typescript-eslint/no-explicit-any
if (import.meta.client && (import.meta as any).hot) {
  watchEffect(() => {
    // @ts-ignore
    window._ga = { attrs: groupAttributes.value, specs: pStore.productSpecs }
  })
}
// ----------------------------
</script>
