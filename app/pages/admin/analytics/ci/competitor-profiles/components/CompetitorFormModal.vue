<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-6">
    <div class="bg-white rounded-xl shadow-xl w-full max-w-4xl max-h-[90vh] overflow-hidden">
      <!-- Header -->
      <div class="flex items-center justify-between p-6 border-b border-gray-200">
        <div class="flex items-center space-x-3 space-x-reverse">
          <div class="w-10 h-10 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
            </svg>
          </div>
          <h2 class="text-xl font-bold text-gray-900">
            {{ isEditing ? 'ویرایش رقیب' : 'افزودن رقیب جدید' }}
          </h2>
        </div>
        <button 
          class="p-2 text-gray-400 hover:text-gray-600 hover:bg-gray-100 rounded-lg transition-colors"
          @click="$emit('close')"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <!-- Form -->
      <form class="overflow-y-auto max-h-[calc(90vh-120px)]" @submit.prevent="handleSubmit">
        <div class="p-6 space-y-6">
          <!-- Basic Information -->
          <div class="bg-gray-50 rounded-xl p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">اطلاعات پایه</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نام شرکت *</label>
                <input
                  v-model="form.name"
                  type="text"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="نام شرکت رقیب"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">وب‌سایت *</label>
                <input
                  v-model="form.website"
                  type="url"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="https://example.com"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">سال تأسیس</label>
                <input
                  v-model="form.foundedYear"
                  type="number"
                  min="1900"
                  :max="new Date().getFullYear()"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="2020"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">اندازه شرکت</label>
                <select
                  v-model="form.companySize"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                >
                  <option value="">انتخاب کنید</option>
                  <option value="1-10 نفر">1-10 نفر</option>
                  <option value="11-50 نفر">11-50 نفر</option>
                  <option value="51-100 نفر">51-100 نفر</option>
                  <option value="101-500 نفر">101-500 نفر</option>
                  <option value="501-1000 نفر">501-1000 نفر</option>
                  <option value="1000+ نفر">1000+ نفر</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">قیمت‌گذاری</label>
                <select
                  v-model="form.pricing"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                >
                  <option value="">انتخاب کنید</option>
                  <option value="کم">کم</option>
                  <option value="متوسط">متوسط</option>
                  <option value="متوسط تا بالا">متوسط تا بالا</option>
                  <option value="بالا">بالا</option>
                  <option value="بسیار بالا">بسیار بالا</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">سطح تهدید</label>
                <select
                  v-model="form.threatLevel"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                >
                  <option value="">انتخاب کنید</option>
                  <option value="low">کم</option>
                  <option value="medium">متوسط</option>
                  <option value="high">زیاد</option>
                </select>
              </div>
            </div>
          </div>

          <!-- Target Markets -->
          <div class="bg-gray-50 rounded-xl p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">بازارهای هدف</h3>
            <div class="space-y-4">
              <div class="flex items-center space-x-2 space-x-reverse">
                <input
                  v-model="newMarket"
                  type="text"
                  class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="نام بازار (مثل: ایران، ترکیه)"
                  @keyup.enter="addMarket"
                />
                <button
                  type="button"
                  class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                  @click="addMarket"
                >
                  افزودن
                </button>
              </div>
              <div class="flex flex-wrap gap-2">
                <span 
                  v-for="(market, index) in form.targetMarkets" 
                  :key="index"
                  class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-blue-100 text-blue-800"
                >
                  {{ market }}
                  <button
                    type="button"
                    class="mr-1 text-blue-600 hover:text-blue-800"
                    @click="removeMarket(index)"
                  >
                    ×
                  </button>
                </span>
              </div>
            </div>
          </div>

          <!-- Key Products -->
          <div class="bg-gray-50 rounded-xl p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">محصولات کلیدی</h3>
            <div class="space-y-4">
              <div class="flex items-center space-x-2 space-x-reverse">
                <input
                  v-model="newProduct"
                  type="text"
                  class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="نام محصول یا سرویس"
                  @keyup.enter="addProduct"
                />
                <button
                  type="button"
                  class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
                  @click="addProduct"
                >
                  افزودن
                </button>
              </div>
              <div class="flex flex-wrap gap-2">
                <span 
                  v-for="(product, index) in form.keyProducts" 
                  :key="index"
                  class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-green-100 text-green-800"
                >
                  {{ product }}
                  <button
                    type="button"
                    class="mr-1 text-green-600 hover:text-green-800"
                    @click="removeProduct(index)"
                  >
                    ×
                  </button>
                </span>
              </div>
            </div>
          </div>

          <!-- Strengths and Weaknesses -->
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div class="bg-gray-50 rounded-xl p-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">نقاط قوت</h3>
              <div class="space-y-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <input
                    v-model="newStrength"
                    type="text"
                    class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500"
                    placeholder="نقطه قوت"
                    @keyup.enter="addStrength"
                  />
                  <button
                    type="button"
                    class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
                    @click="addStrength"
                  >
                    +
                  </button>
                </div>
                <div class="space-y-2">
                  <div 
                    v-for="(strength, index) in form.strengths" 
                    :key="index"
                    class="flex items-center justify-between p-2 bg-green-100 rounded-lg"
                  >
                    <span class="text-sm text-green-800">{{ strength }}</span>
                    <button
                      type="button"
                      class="text-green-600 hover:text-green-800"
                      @click="removeStrength(index)"
                    >
                      ×
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <div class="bg-gray-50 rounded-xl p-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">نقاط ضعف</h3>
              <div class="space-y-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <input
                    v-model="newWeakness"
                    type="text"
                    class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500"
                    placeholder="نقطه ضعف"
                    @keyup.enter="addWeakness"
                  />
                  <button
                    type="button"
                    class="px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
                    @click="addWeakness"
                  >
                    +
                  </button>
                </div>
                <div class="space-y-2">
                  <div 
                    v-for="(weakness, index) in form.weaknesses" 
                    :key="index"
                    class="flex items-center justify-between p-2 bg-red-100 rounded-lg"
                  >
                    <span class="text-sm text-red-800">{{ weakness }}</span>
                    <button
                      type="button"
                      class="text-red-600 hover:text-red-800"
                      @click="removeWeakness(index)"
                    >
                      ×
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Recent Changes -->
          <div class="bg-gray-50 rounded-xl p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تغییرات اخیر</h3>
            <div class="space-y-4">
              <div class="flex items-center space-x-2 space-x-reverse">
                <input
                  v-model="newChange"
                  type="text"
                  class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="تغییر اخیر (مثل: لانچ محصول جدید)"
                  @keyup.enter="addChange"
                />
                <button
                  type="button"
                  class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                  @click="addChange"
                >
                  افزودن
                </button>
              </div>
              <div class="space-y-2">
                <div 
                  v-for="(change, index) in form.recentChanges" 
                  :key="index"
                  class="flex items-center justify-between p-2 bg-blue-100 rounded-lg"
                >
                  <span class="text-sm text-blue-800">{{ change }}</span>
                  <button
                    type="button"
                    class="text-blue-600 hover:text-blue-800"
                    @click="removeChange(index)"
                  >
                    ×
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Footer -->
        <div class="flex items-center justify-end space-x-3 space-x-reverse p-6 border-t border-gray-200 bg-gray-50">
          <button
            type="button"
            class="px-4 py-2 border border-gray-300 text-sm font-medium rounded-lg text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            @click="$emit('close')"
          >
            انصراف
          </button>
          <button
            type="submit"
            :disabled="isSubmitting"
            class="px-6 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <svg v-if="isSubmitting" class="animate-spin -ml-1 mr-3 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            {{ isSubmitting ? 'در حال ذخیره...' : (isEditing ? 'ویرایش' : 'افزودن') }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { type CompetitorProfile } from '../../composables/useCompetitiveIntelligence'

// Props
interface Props {
  competitor?: CompetitorProfile | null
}

const props = withDefaults(defineProps<Props>(), {
  competitor: null
})

// Emits
const emit = defineEmits<{
  close: []
  save: [competitorData: Omit<CompetitorProfile, 'id'>]
}>()

// State
const isSubmitting = ref(false)
const newMarket = ref('')
const newProduct = ref('')
const newStrength = ref('')
const newWeakness = ref('')
const newChange = ref('')

// Form data
const form = ref({
  name: '',
  website: '',
  foundedYear: new Date().getFullYear(),
  companySize: '',
  targetMarkets: [] as string[],
  keyProducts: [] as string[],
  pricing: '',
  strengths: [] as string[],
  weaknesses: [] as string[],
  recentChanges: [] as string[],
  threatLevel: ''
})

// Computed
const isEditing = computed(() => !!props.competitor)

// Methods
const addMarket = () => {
  if (newMarket.value.trim() && !form.value.targetMarkets.includes(newMarket.value.trim())) {
    form.value.targetMarkets.push(newMarket.value.trim())
    newMarket.value = ''
  }
}

const removeMarket = (index: number) => {
  form.value.targetMarkets.splice(index, 1)
}

const addProduct = () => {
  if (newProduct.value.trim() && !form.value.keyProducts.includes(newProduct.value.trim())) {
    form.value.keyProducts.push(newProduct.value.trim())
    newProduct.value = ''
  }
}

const removeProduct = (index: number) => {
  form.value.keyProducts.splice(index, 1)
}

const addStrength = () => {
  if (newStrength.value.trim() && !form.value.strengths.includes(newStrength.value.trim())) {
    form.value.strengths.push(newStrength.value.trim())
    newStrength.value = ''
  }
}

const removeStrength = (index: number) => {
  form.value.strengths.splice(index, 1)
}

const addWeakness = () => {
  if (newWeakness.value.trim() && !form.value.weaknesses.includes(newWeakness.value.trim())) {
    form.value.weaknesses.push(newWeakness.value.trim())
    newWeakness.value = ''
  }
}

const removeWeakness = (index: number) => {
  form.value.weaknesses.splice(index, 1)
}

const addChange = () => {
  if (newChange.value.trim() && !form.value.recentChanges.includes(newChange.value.trim())) {
    form.value.recentChanges.push(newChange.value.trim())
    newChange.value = ''
  }
}

const removeChange = (index: number) => {
  form.value.recentChanges.splice(index, 1)
}

const resetForm = () => {
  form.value = {
    name: '',
    website: '',
    foundedYear: new Date().getFullYear(),
    companySize: '',
    targetMarkets: [],
    keyProducts: [],
    pricing: '',
    strengths: [],
    weaknesses: [],
    recentChanges: [],
    threatLevel: ''
  }
}

const loadCompetitorData = (competitor: CompetitorProfile) => {
  form.value = {
    name: competitor.name,
    website: competitor.website,
    foundedYear: competitor.foundedYear,
    companySize: competitor.companySize,
    targetMarkets: [...competitor.targetMarkets],
    keyProducts: [...competitor.keyProducts],
    pricing: competitor.pricing,
    strengths: [...competitor.strengths],
    weaknesses: [...competitor.weaknesses],
    recentChanges: [...competitor.recentChanges],
    threatLevel: competitor.threatLevel
  }
}

const handleSubmit = async () => {
  if (isSubmitting.value) return

  // Validation
  if (!form.value.name.trim() || !form.value.website.trim() || !form.value.threatLevel) {
    alert('لطفاً فیلدهای اجباری را پر کنید')
    return
  }

  isSubmitting.value = true

  try {
    // Create competitor data object
    const competitorData: Omit<CompetitorProfile, 'id'> = {
      name: form.value.name.trim(),
      website: form.value.website.trim(),
      foundedYear: form.value.foundedYear,
      companySize: form.value.companySize,
      targetMarkets: form.value.targetMarkets,
      keyProducts: form.value.keyProducts,
      pricing: form.value.pricing,
      strengths: form.value.strengths,
      weaknesses: form.value.weaknesses,
      recentChanges: form.value.recentChanges,
      threatLevel: form.value.threatLevel as 'low' | 'medium' | 'high'
    }

    // Emit save event
    emit('save', competitorData)
  } catch (error) {
    console.error('Error saving competitor:', error)
    alert('خطا در ذخیره اطلاعات')
  } finally {
    isSubmitting.value = false
  }
}

// Lifecycle
onMounted(() => {
  if (props.competitor) {
    loadCompetitorData(props.competitor)
  } else {
    resetForm()
  }
})

// Watch for competitor changes
watch(() => props.competitor, (newCompetitor) => {
  if (newCompetitor) {
    loadCompetitorData(newCompetitor)
  } else {
    resetForm()
  }
})
</script>

<style scoped>
/* Custom styles */
</style>
