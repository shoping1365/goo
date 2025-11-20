<template>
  <Teleport to="body">
    <div v-if="open" class="fixed inset-0 z-50 flex items-center justify-center">
      <!-- پنجره اصلی -->
      <div class="bg-white w-full max-w-2xl max-h-[90vh] overflow-y-auto overflow-x-hidden rounded-2xl shadow-lg p-6" dir="rtl">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-xl font-bold">آدرس‌های شما</h2>
          <button class="text-gray-600 hover:text-gray-800 text-xl font-bold" @click="emit('close')">×</button>
        </div>

        <!-- حالت فهرست -->
        <template v-if="mode==='list'">
          <div v-if="addressesLoading" class="flex justify-center py-8">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-[#e60023]"></div>
          </div>
          <div v-else class="space-y-3">
            <div v-for="addr in addresses" :key="addr.id" :class="cardClass(addr)" class="p-6 rounded-xl border relative">
              <!-- علامت پیش‌فرض -->
              <div v-if="addr.is_default" class="absolute -top-2 -left-2 bg-teal-500 text-white text-xs rounded-full px-2 py-0.5">پیش‌فرض</div>

              <div class="flex flex-col items-end w-full" dir="rtl">
                <p class="font-medium text-gray-800 text-right w-full">{{ addr.recipient_name }}</p>
                <p class="text-sm text-gray-600 mt-1 text-right w-full">کد پستی: <span>{{ addr.postal_code }}</span></p>
                <p class="text-sm text-gray-800 mt-1 text-right w-full">آدرس: {{ addr.full_address }}</p>
                <p class="text-sm text-gray-800 mt-1 text-right w-full">شماره موبایل: {{ addr.recipient_mobile }}</p>
                <p v-if="addr.phone" class="text-sm text-gray-800 mt-1 text-right w-full">تلفن ثابت: {{ addr.phone }}</p>
              </div>
              <div class="flex items-center gap-3 mt-3 w-full">
                <button v-if="!addr.is_default" class="text-xs text-gray-500 hover:text-teal-600" @click="makeDefault(addr.id)">پیش‌فرض کن</button>
              </div>
              <div class="flex items-center gap-3 mt-3 w-full">
                <button class="text-sm text-[#e60023] font-bold hover:underline" @click="selectAddress(addr)">انتخاب</button>
              </div>
              <div class="flex items-center gap-3 mt-2 w-full">
                <button class="text-xs text-gray-500 hover:text-gray-700 underline" @click="editAddress(addr)">ویرایش</button>
                <button class="text-xs text-red-500 hover:text-red-700 underline" @click="confirmDelete(addr.id)">حذف</button>
              </div>
            </div>
          </div>

          <button class="mt-6 w-full flex items-center justify-center gap-2 border-2 border-dashed border-gray-300 rounded-xl py-3 hover:bg-gray-50 text-sm font-bold text-gray-600" @click="mode='add'; resetForm();">
            + افزودن آدرس جدید
          </button>
        </template>

        <!-- حالت افزودن / ویرایش -->
        <template v-else>
          <form class="space-y-4" @submit.prevent="saveAddress">
            <h3 class="text-lg font-bold mb-4 pb-2 border-b border-gray-200">{{ mode==='add' ? 'افزودن آدرس جدید' : 'ویرایش آدرس' }}</h3>

            <div>
              <label class="form-label">نام گیرنده *</label>
              <input v-model="form.recipient_name" required class="w-full input-base" placeholder="نام و نام خانوادگی گیرنده" />
            </div>
            <div>
              <label class="form-label">موبایل گیرنده *</label>
              <input v-model="form.recipient_mobile" required class="w-full input-base" placeholder="09xxxxxxxxx" />
            </div>
            <div>
              <label class="form-label">تلفن ثابت</label>
              <input v-model="form.phone" class="w-full input-base" placeholder="021xxxxxxxx" />
            </div>
            <div>
              <label class="form-label">آدرس کامل *</label>
              <textarea v-model="form.full_address" required rows="3" class="w-full input-base" placeholder="آدرس کامل" />
            </div>
            <div class="grid grid-cols-2 gap-3">
              <div>
                <label class="form-label">استان *</label>
                <select v-model="selectedProvinceId" required class="w-full input-base">
                  <option value="" disabled>انتخاب استان</option>
                  <option v-for="p in provinces" :key="p.id" :value="p.id">{{ p.name }}</option>
                </select>
              </div>
              <div>
                <label class="form-label">شهر *</label>
                <select v-model="form.city_id" required class="w-full input-base" :disabled="!selectedProvinceId || citiesLoading">
                  <option value="" disabled>{{ citiesLoading ? 'در حال بارگذاری...' : 'انتخاب شهر' }}</option>
                  <option v-for="c in cities" :key="c.id" :value="c.id">{{ c.name }}</option>
                </select>
              </div>
            </div>
            <div>
              <label class="form-label">کد پستی</label>
              <input v-model="form.postal_code" class="w-full input-base" placeholder="کد پستی ۱۰ رقمی" />
            </div>
            <div class="flex items-center gap-2">
              <input id="is_default" v-model="form.is_default" type="checkbox" />
              <label for="is_default" class="text-sm">تنظیم به عنوان آدرس پیش‌فرض</label>
            </div>

            <div class="flex gap-3 mt-6 pt-4 border-t border-gray-200">
              <button type="submit" :disabled="saving" class="flex-1 py-3 rounded-xl bg-[#e60023] text-white text-sm font-bold hover:bg-[#c9001b] transition-colors shadow-sm disabled:opacity-50 disabled:cursor-not-allowed">
                <span v-if="saving" class="flex items-center justify-center gap-2">
                  <div class="animate-spin rounded-full h-4 w-4 border-b-2 border-white"></div>
                  در حال ذخیره...
                </span>
                <span v-else>{{ mode==='add' ? 'ذخیره آدرس' : 'ذخیره تغییرات' }}</span>
              </button>
              <button type="button" class="flex-1 py-3 rounded-xl border-2 border-gray-300 text-sm font-medium hover:bg-gray-50 transition-colors" @click="cancelForm">انصراف</button>
            </div>
          </form>
        </template>
      </div>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref, watch } from 'vue';
import { useAddresses } from '~/composables/useAddresses';

interface Address {
  id: number
  recipient_name: string
  recipient_mobile: string
  phone?: string
  full_address: string
  postal_code?: string
  province_id?: number
  city_id?: number
  province?: string
  city?: string
  is_default: boolean
}

interface Province {
  id: number
  name: string
}

interface City {
  id: number
  name: string
  province_id: number
}

// Props
interface Props {
  open: boolean
}
const props = defineProps<Props>()
const emit = defineEmits(['close','select','addresses-changed'])

// Composable
const { addresses, loading: addressesLoading, fetchAddresses, addAddress, updateAddress, deleteAddress, setDefaultAddress } = useAddresses()

const mode = ref<'list' | 'add' | 'edit'>('list')
const saving = ref(false)
const citiesLoading = ref(false)

// فرم آدرس
const form = reactive<Partial<Address>>({
  id: undefined,
  recipient_name: '',
  recipient_mobile: '',
  phone: '',
  full_address: '',
  postal_code: '',
  province_id: undefined,
  city_id: undefined,
  is_default: false
})

function resetForm(){
  Object.assign(form,{ id:undefined, recipient_name:'', recipient_mobile:'', phone:'', full_address:'', postal_code:'', province_id:undefined, city_id:undefined, is_default:false })
  selectedProvinceId.value = null
  cities.value = []
}

// استان و شهر - با کش
const provinces = ref<Province[]>([])
const cities = ref<City[]>([])
const selectedProvinceId = ref<number|null>(null)
const provincesLoaded = ref(false)
const citiesCache = new Map<number, City[]>()

// فراخوانی استان‌ها فقط یک بار
async function fetchProvinces(){
  if (provincesLoaded.value) return
  try {
    provinces.value = await $fetch<Province[]>('/api/geo/provinces')
    provincesLoaded.value = true
  } catch {
    // خطا در دریافت استان‌ها
  }
}

// فراخوانی شهرها با کش
async function fetchCities(pid: number){
  if (citiesCache.has(pid)) {
    cities.value = citiesCache.get(pid)!
    return
  }
  
  citiesLoading.value = true
  try {
    const citiesData = await $fetch<City[]>(`/api/geo/provinces/${pid}/cities`)
    cities.value = Array.isArray(citiesData) ? citiesData : []
    citiesCache.set(pid, Array.isArray(citiesData) ? citiesData : [])
  } catch {
    // خطا در دریافت شهرها
    cities.value = []
  } finally {
    citiesLoading.value = false
  }
}

// بهینه‌سازی watch با debounce
let citiesFetchTimeout: NodeJS.Timeout | null = null
watch(selectedProvinceId, (val) => {
  if (citiesFetchTimeout) {
    clearTimeout(citiesFetchTimeout)
  }
  
  if (val) {
    citiesFetchTimeout = setTimeout(() => {
      fetchCities(val)
    }, 300) // تاخیر 300ms برای جلوگیری از فراخوانی‌های مکرر
  } else {
    cities.value = []
  }
})

// --- اکشن‌ها ---
function cardClass(addr: Address){
  return addr.is_default ? 'border-teal-500 border-2' : 'border-gray-200'
}

function selectAddress(addr: Address){
  emit('select', addr)
  emit('close')
}

function editAddress(addr: Address){
  mode.value = 'edit'
  Object.assign(form, JSON.parse(JSON.stringify(addr)))
  selectedProvinceId.value = addr.province_id ?? null
  // اگر آدرس دارای نام استان و شهر است، آن‌ها را در فرم قرار بده
  if (addr.province && !provinces.value.find(p => p.name === addr.province)) {
    // اگر استان در لیست موجود نیست، اضافه کن
    if (addr.province_id) {
      provinces.value.push({ id: addr.province_id, name: addr.province })
    }
  }
  if (addr.city && !cities.value.find(c => c.name === addr.city)) {
    // اگر شهر در لیست موجود نیست، اضافه کن
    if (addr.city_id && addr.province_id) {
      cities.value.push({ id: addr.city_id, name: addr.city, province_id: addr.province_id })
    }
  }
  if(addr.province_id) fetchCities(addr.province_id)
}

async function saveAddress(){
  saving.value = true
  try {
    // اضافه کردن نام استان و شهر به فرم
    const selectedProvince = provinces.value.find(p => p.id === selectedProvinceId.value)
    const selectedCity = cities.value.find(c => c.id === form.city_id)
    
    const addressData = {
      ...form,
      province: selectedProvince?.name || '',
      city: selectedCity?.name || '',
    }
    
    if (mode.value === 'add'){
      const res = await addAddress(addressData)
      if (form.is_default) await setDefaultAddress(res.id)
    } else if (mode.value === 'edit' && form.id){
      await updateAddress(form.id, addressData)
      if (form.is_default) await setDefaultAddress(form.id)
    }
    await fetchAddresses()
    emit('addresses-changed')
    mode.value = 'list'
    resetForm()
  } catch {
    // خطا در ذخیره آدرس
    alert('خطا در ذخیره آدرس. لطفاً دوباره تلاش کنید.')
  } finally {
    saving.value = false
  }
}

function cancelForm(){
  mode.value = 'list'
  resetForm()
}

async function confirmDelete(id: number){
  if (confirm('آیا از حذف این آدرس مطمئن هستید؟')){
    try {
      await deleteAddress(id)
      emit('addresses-changed')
    } catch {
      // خطا در حذف آدرس
      alert('خطا در حذف آدرس. لطفاً دوباره تلاش کنید.')
    }
  }
}

async function makeDefault(id: number){
  try {
    await setDefaultAddress(id)
    emit('addresses-changed')
  } catch {
    // خطا در تنظیم آدرس پیش‌فرض
    alert('خطا در تنظیم آدرس پیش‌فرض. لطفاً دوباره تلاش کنید.')
  }
}

// بهینه‌سازی watch برای باز شدن مودال
let modalOpenTimeout: NodeJS.Timeout | null = null
watch(() => props.open, async (val) => {
  if (modalOpenTimeout) {
    clearTimeout(modalOpenTimeout)
  }
  
  if (val) {
    modalOpenTimeout = setTimeout(async () => {
      // فراخوانی همزمان برای بهبود عملکرد
      await Promise.all([
        fetchAddresses(),
        fetchProvinces()
      ])
      mode.value = 'list'
      resetForm()
    }, 100)
  }
})

// پاک‌سازی timeout ها
onMounted(() => {
  return () => {
    if (citiesFetchTimeout) clearTimeout(citiesFetchTimeout)
    if (modalOpenTimeout) clearTimeout(modalOpenTimeout)
  }
})
</script>

<style scoped lang="postcss">
@layer components {
  .input-base {
    /* استفاده از کلاس‌های معمولی به‌جای @apply برای جلوگیری از خطای Tailwind */
    padding-left: 1rem;
    padding-right: 1rem;
    padding-top: 0.75rem;
    padding-bottom: 0.75rem;
    border-width: 2px;
    border-color: rgb(209 213 219);
    border-radius: 0.75rem;
    background-color: white;
    outline: none;
  }

  .input-base:focus {
    border-color: #e60023;
    box-shadow: 0 0 0 2px #e60023;
    box-shadow-color: rgba(230, 0, 35, 0.2);
  }

  .text-primary-600 {
    color: #e60023;
  }

  .form-label {
    font-size: 0.875rem;
    font-weight: 500;
    color: rgb(55 65 81);
    margin-bottom: 0.5rem;
    display: block;
  }
}
</style> 