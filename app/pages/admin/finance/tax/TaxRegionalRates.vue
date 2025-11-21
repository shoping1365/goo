<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <!-- هدر -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6 mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">نرخ‌های منطقه‌ای مالیات</h3>
        <p class="text-sm text-gray-600">مدیریت نرخ‌های متفاوت مالیات بر اساس مناطق جغرافیایی</p>
      </div>
      
      <!-- دکمه افزودن منطقه جدید -->
      <button 
        class="inline-flex items-center gap-2 px-4 py-2 bg-purple-600 hover:bg-purple-700 text-white rounded-lg transition-colors duration-200"
        @click="showAddRegionModal = true"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
        </svg>
        افزودن منطقه جدید
      </button>
    </div>

    <!-- نقشه مناطق -->
    <div class="mb-6">
      <div class="bg-gray-50 rounded-lg p-6">
        <h4 class="text-md font-semibold text-gray-900 mb-4">نقشه مناطق مالیاتی</h4>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          <div 
            v-for="region in regions" 
            :key="region.id"
            class="bg-white rounded-lg p-6 border-2 transition-colors duration-200 cursor-pointer"
            :class="selectedRegion?.id === region.id ? 'border-purple-500 bg-purple-50' : 'border-gray-200 hover:border-purple-300'"
            @click="selectRegion(region)"
          >
            <div class="flex items-center justify-between mb-2">
              <h5 class="font-medium text-gray-900">{{ region.name }}</h5>
              <div class="w-3 h-3 rounded-full" :class="getRegionColor(region.taxRate)"></div>
            </div>
            <div class="text-sm text-gray-600">
              <div>نرخ مالیات: {{ region.taxRate }}%</div>
              <div>تعداد شهرها: {{ region.citiesCount }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- جزئیات منطقه انتخاب شده -->
    <div v-if="selectedRegion" class="mb-6">
      <div class="bg-purple-50 rounded-lg p-6">
        <div class="flex items-center justify-between mb-4">
          <h4 class="text-lg font-semibold text-purple-900">{{ selectedRegion.name }}</h4>
          <button 
            class="inline-flex items-center gap-2 px-3 py-1.5 bg-purple-600 hover:bg-purple-700 text-white rounded-lg text-sm transition-colors duration-200"
            @click="editRegion(selectedRegion)"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
            </svg>
            ویرایش منطقه
          </button>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <!-- اطلاعات کلی -->
          <div>
            <h5 class="font-medium text-purple-800 mb-3">اطلاعات کلی</h5>
            <div class="space-y-2 text-sm">
              <div class="flex justify-between">
                <span class="text-gray-600">نرخ مالیات:</span>
                <span class="font-medium text-purple-900">{{ selectedRegion.taxRate }}%</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">تعداد شهرها:</span>
                <span class="font-medium text-purple-900">{{ selectedRegion.citiesCount }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">جمعیت:</span>
                <span class="font-medium text-purple-900">{{ formatNumber(selectedRegion.population) }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">مساحت:</span>
                <span class="font-medium text-purple-900">{{ formatArea(selectedRegion.area) }}</span>
              </div>
            </div>
          </div>

          <!-- آمار مالیاتی -->
          <div>
            <h5 class="font-medium text-purple-800 mb-3">آمار مالیاتی</h5>
            <div class="space-y-2 text-sm">
              <div class="flex justify-between">
                <span class="text-gray-600">مالیات ماهانه:</span>
                <span class="font-medium text-purple-900">{{ formatCurrency(selectedRegion.monthlyTax) }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">مالیات سالانه:</span>
                <span class="font-medium text-purple-900">{{ formatCurrency(selectedRegion.yearlyTax) }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">رشد نسبت به سال قبل:</span>
                <span class="font-medium" :class="selectedRegion.growth >= 0 ? 'text-green-600' : 'text-red-600'">
                  {{ selectedRegion.growth >= 0 ? '+' : '' }}{{ selectedRegion.growth }}%
                </span>
              </div>
            </div>
          </div>

          <!-- شهرهای منطقه -->
          <div>
            <h5 class="font-medium text-purple-800 mb-3">شهرهای منطقه</h5>
            <div class="space-y-1 text-sm">
              <div 
                v-for="city in selectedRegion.cities.slice(0, 5)" 
                :key="city.id"
                class="flex justify-between items-center py-1"
              >
                <span class="text-gray-600">{{ city.name }}</span>
                <span class="font-medium text-purple-900">{{ city.taxRate }}%</span>
              </div>
              <div v-if="selectedRegion.cities.length > 5" class="text-xs text-gray-500 text-center pt-2">
                و {{ selectedRegion.cities.length - 5 }} شهر دیگر...
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول مناطق -->
    <div class="overflow-x-auto">
      <table class="w-full text-sm">
        <thead>
          <tr class="border-b border-gray-200 bg-gray-50">
            <th class="text-right py-3 px-4 font-medium text-gray-600">منطقه</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">نرخ مالیات</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">تعداد شهرها</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">جمعیت</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">مالیات ماهانه</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">رشد</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="region in regions" :key="region.id" class="border-b border-gray-100 hover:bg-gray-50">
            <td class="py-3 px-4">
              <div class="flex items-center gap-2">
                <div class="w-8 h-8 rounded-lg flex items-center justify-center" :class="getRegionBgColor(region.taxRate)">
                  <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
                  </svg>
                </div>
                <span class="font-medium text-gray-900">{{ region.name }}</span>
              </div>
            </td>
            <td class="py-3 px-4">
              <span class="font-bold text-blue-600">{{ region.taxRate }}%</span>
            </td>
            <td class="py-3 px-4 text-gray-900">{{ region.citiesCount }}</td>
            <td class="py-3 px-4 text-gray-600">{{ formatNumber(region.population) }}</td>
            <td class="py-3 px-4 text-gray-900">{{ formatCurrency(region.monthlyTax) }}</td>
            <td class="py-3 px-4">
              <span :class="region.growth >= 0 ? 'text-green-600' : 'text-red-600'">
                {{ region.growth >= 0 ? '+' : '' }}{{ region.growth }}%
              </span>
            </td>
            <td class="py-3 px-4">
              <span 
                :class="[
                  'px-2 py-1 rounded-full text-xs font-medium',
                  region.status === 'active' ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'
                ]"
              >
                {{ region.status === 'active' ? 'فعال' : 'غیرفعال' }}
              </span>
            </td>
            <td class="py-3 px-4">
              <div class="flex items-center gap-2">
                <button 
                  class="p-1 text-blue-600 hover:text-blue-800 transition-colors"
                  @click="editRegion(region)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                  </svg>
                </button>
                <button 
                  :class="[
                    'p-1 transition-colors',
                    region.status === 'active' ? 'text-red-600 hover:text-red-800' : 'text-green-600 hover:text-green-800'
                  ]"
                  @click="toggleRegionStatus(region)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="region.status === 'active' ? 'M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728' : 'M5 13l4 4L19 7'"></path>
                  </svg>
                </button>
                <button 
                  v-if="canDeleteTaxRegion"
                  class="p-1 text-red-600 hover:text-red-800 transition-colors"
                  @click="deleteRegion(region)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                  </svg>
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- مودال افزودن/ویرایش منطقه -->
    <div v-if="showAddRegionModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl p-6 w-full max-w-lg mx-4">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">
            {{ editingRegion ? 'ویرایش منطقه' : 'افزودن منطقه جدید' }}
          </h3>
          <button class="text-gray-400 hover:text-gray-600" @click="closeModal">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <form class="space-y-4" @submit.prevent="saveRegion">
          <!-- نام منطقه -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام منطقه</label>
            <input 
              v-model="regionForm.name"
              type="text"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
            />
          </div>

          <!-- نرخ مالیات -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نرخ مالیات (%)</label>
            <input 
              v-model.number="regionForm.taxRate"
              type="number"
              step="0.1"
              min="0"
              max="100"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
            />
          </div>

          <!-- کد پستی -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">کد پستی (محدوده)</label>
            <div class="grid grid-cols-2 gap-2">
              <input 
                v-model="regionForm.postalCodeFrom"
                type="text"
                placeholder="از"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
              />
              <input 
                v-model="regionForm.postalCodeTo"
                type="text"
                placeholder="تا"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
              />
            </div>
          </div>

          <!-- شهرها -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">شهرهای منطقه</label>
            <textarea 
              v-model="regionForm.cities"
              rows="3"
              placeholder="نام شهرها را با کاما جدا کنید"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
            ></textarea>
          </div>

          <!-- توضیحات -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
            <textarea 
              v-model="regionForm.description"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
            ></textarea>
          </div>

          <!-- دکمه‌ها -->
          <div class="flex gap-3 pt-4">
            <button 
              type="submit"
              class="flex-1 px-4 py-2 bg-purple-600 hover:bg-purple-700 text-white rounded-lg transition-colors duration-200"
            >
              {{ editingRegion ? 'ویرایش' : 'افزودن' }}
            </button>
            <button 
              type="button"
              class="flex-1 px-4 py-2 bg-gray-200 hover:bg-gray-300 text-gray-700 rounded-lg transition-colors duration-200"
              @click="closeModal"
            >
              انصراف
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { useAuth } from '~/composables/useAuth';

// احراز هویت غیرفعال شده است - محدودیت دسترسی حذف شده
definePageMeta({
  // احراز هویت غیرفعال شده است
})

// دسترسی حذف منطقه مالیاتی فقط برای توسعه‌دهنده
const { user: _user } = useAuth()
const canDeleteTaxRegion = computed(() => {
  // const roles = _user.value?.roles || []
  // return roles.includes('developer')
  return true // احراز هویت غیرفعال شده است
})

// مناطق
const regions = ref([
  {
    id: 1,
    name: 'تهران',
    taxRate: 9,
    citiesCount: 16,
    population: 13500000,
    area: 730,
    monthlyTax: 45000000000,
    yearlyTax: 540000000000,
    growth: 12.5,
    status: 'active',
    cities: [
      { id: 1, name: 'تهران', taxRate: 9 },
      { id: 2, name: 'شهریار', taxRate: 9 },
      { id: 3, name: 'ورامین', taxRate: 9 },
      { id: 4, name: 'فیروزکوه', taxRate: 9 },
      { id: 5, name: 'دماوند', taxRate: 9 }
    ]
  },
  {
    id: 2,
    name: 'اصفهان',
    taxRate: 8,
    citiesCount: 24,
    population: 5200000,
    area: 107029,
    monthlyTax: 28000000000,
    yearlyTax: 336000000000,
    growth: 8.3,
    status: 'active',
    cities: [
      { id: 6, name: 'اصفهان', taxRate: 8 },
      { id: 7, name: 'کاشان', taxRate: 8 },
      { id: 8, name: 'نجف‌آباد', taxRate: 8 }
    ]
  },
  {
    id: 3,
    name: 'خراسان رضوی',
    taxRate: 7,
    citiesCount: 28,
    population: 6400000,
    area: 118851,
    monthlyTax: 32000000000,
    yearlyTax: 384000000000,
    growth: 15.7,
    status: 'active',
    cities: [
      { id: 9, name: 'مشهد', taxRate: 7 },
      { id: 10, name: 'نیشابور', taxRate: 7 },
      { id: 11, name: 'سبزوار', taxRate: 7 }
    ]
  },
  {
    id: 4,
    name: 'فارس',
    taxRate: 8.5,
    citiesCount: 29,
    population: 4900000,
    area: 122608,
    monthlyTax: 25000000000,
    yearlyTax: 300000000000,
    growth: 6.2,
    status: 'inactive',
    cities: [
      { id: 12, name: 'شیراز', taxRate: 8.5 },
      { id: 13, name: 'مرودشت', taxRate: 8.5 },
      { id: 14, name: 'جهرم', taxRate: 8.5 }
    ]
  }
])

// منطقه انتخاب شده
const selectedRegion = ref(null)

// مودال
const showAddRegionModal = ref(false)
const editingRegion = ref(null)

// فرم منطقه
const regionForm = ref({
  name: '',
  taxRate: 0,
  postalCodeFrom: '',
  postalCodeTo: '',
  cities: '',
  description: ''
})

interface Region {
  id?: number | string
  name?: string
  [key: string]: unknown
}

// انتخاب منطقه
const selectRegion = (region: Region) => {
  selectedRegion.value = region
}

// رنگ منطقه بر اساس نرخ مالیات
const getRegionColor = (taxRate: number) => {
  if (taxRate >= 9) return 'bg-red-500'
  if (taxRate >= 7) return 'bg-orange-500'
  if (taxRate >= 5) return 'bg-yellow-500'
  return 'bg-green-500'
}

// رنگ پس‌زمینه منطقه
const getRegionBgColor = (taxRate: number) => {
  if (taxRate >= 9) return 'bg-red-500'
  if (taxRate >= 7) return 'bg-orange-500'
  if (taxRate >= 5) return 'bg-yellow-500'
  return 'bg-green-500'
}

// فرمت عدد
const formatNumber = (num: number): string => {
  return new Intl.NumberFormat('fa-IR').format(num)
}

// فرمت مساحت
const formatArea = (area: number): string => {
  return `${formatNumber(area)} کیلومتر مربع`
}

// فرمت مبلغ
const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

// ویرایش منطقه
const editRegion = (region: Region) => {
  editingRegion.value = region
  regionForm.value = {
    name: region.name,
    taxRate: region.taxRate,
    postalCodeFrom: '',
    postalCodeTo: '',
    cities: Array.isArray(region.cities) ? region.cities.map((c: { name?: string }) => c.name || '').join(', ') : '',
    description: ''
  }
  showAddRegionModal.value = true
}

// تغییر وضعیت منطقه
const toggleRegionStatus = async (region: Region) => {
  try {
    // TODO: ارسال درخواست به API
    region.status = region.status === 'active' ? 'inactive' : 'active'
  } catch (error) {
    console.error('خطا در تغییر وضعیت منطقه:', error)
  }
}

// حذف منطقه
const deleteRegion = async (region: Region) => {
  if (confirm('آیا از حذف این منطقه اطمینان دارید؟')) {
    try {
      // TODO: ارسال درخواست به API
      const index = regions.value.findIndex(r => r.id === region.id)
      if (index > -1) {
        regions.value.splice(index, 1)
      }
      if (selectedRegion.value?.id === region.id) {
        selectedRegion.value = null
      }
    } catch (error) {
      console.error('خطا در حذف منطقه:', error)
    }
  }
}

// ذخیره منطقه
const saveRegion = async () => {
  try {
    if (editingRegion.value) {
      // ویرایش منطقه موجود
      const index = regions.value.findIndex(r => r.id === editingRegion.value.id)
      if (index > -1) {
        regions.value[index] = { ...editingRegion.value, ...regionForm.value }
      }
    } else {
      // افزودن منطقه جدید
      const newRegion = {
        id: Date.now(),
        ...regionForm.value,
        citiesCount: regionForm.value.cities.split(',').length,
        population: 0,
        area: 0,
        monthlyTax: 0,
        yearlyTax: 0,
        growth: 0,
        status: 'active',
        cities: regionForm.value.cities.split(',').map((name: string, index: number) => ({
          id: index + 1,
          name: name.trim(),
          taxRate: regionForm.value.taxRate
        }))
      }
      regions.value.push(newRegion)
    }
    
    closeModal()
  } catch (error) {
    console.error('خطا در ذخیره منطقه:', error)
  }
}

// بستن مودال
const closeModal = () => {
  showAddRegionModal.value = false
  editingRegion.value = null
  regionForm.value = {
    name: '',
    taxRate: 0,
    postalCodeFrom: '',
    postalCodeTo: '',
    cities: '',
    description: ''
  }
}

// بارگذاری اولیه
onMounted(() => {
  // TODO: بارگذاری مناطق از API
})
</script>

<!--
  کامپوننت نرخ‌های منطقه‌ای مالیات
  شامل:
  1. نقشه مناطق مالیاتی
  2. جزئیات هر منطقه
  3. مدیریت نرخ‌های متفاوت
  4. آمار منطقه‌ای
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 

