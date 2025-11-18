<script setup>
import { useConfirmDialog } from '~/composables/useConfirmDialog'
import RichTextEditor from '~/components/common/RichTextEditor.vue'

const tinyApiKey = 'qwa4j6x5mh2e3241igpyi345b4uhe2d5qeq6f8hy9qfkw2ro'

// Section collapse/expand states for FAQ tab
const sections = reactive({
  faq: true
})

function toggleSection(sectionName) {
  sections[sectionName] = !sections[sectionName]
}

// FAQ management data
const faqs = ref([
  {
    question: "چگونه این محصول را استفاده کنم؟",
    answer: "برای استفاده از این محصول ابتدا دستورالعمل را مطالعه کنید...",
    category: "usage",
    status: "published",
    priority: 1
  }
])

const newFaq = ref({
  question: '',
  answer: '',
  category: '',
  status: 'published',
  priority: 1
})

const editingIndex = ref(-1)

const faqStats = computed(() => ({
  total: faqs.value.length,
  published: faqs.value.filter(f => f.status === 'published').length,
  pending: faqs.value.filter(f => f.status === 'pending').length,
  categories: [...new Set(faqs.value.map(f => f.category))].length
}))

function addFaq() {
  if (newFaq.value.question && newFaq.value.answer) {
    if (editingIndex.value >= 0) {
      faqs.value[editingIndex.value] = { ...newFaq.value }
      editingIndex.value = -1
    } else {
      faqs.value.push({ ...newFaq.value })
    }
    clearFaqForm()
  }
}

function clearFaqForm() {
  newFaq.value = {
    question: '',
    answer: '',
    category: '',
    status: 'published',
    priority: 1
  }
  editingIndex.value = -1
}

function editFaq(index) {
  newFaq.value = { ...faqs.value[index] }
  editingIndex.value = index
}

function deleteFaq(index) {
  const { confirm } = useConfirmDialog()
  confirm({
    title: 'تأیید حذف',
    message: 'آیا از حذف این سوال اطمینان دارید؟',
    confirmText: 'حذف',
    cancelText: 'انصراف',
    type: 'danger'
  }).then(ok => {
    if (!ok) return
    faqs.value.splice(index, 1)
    if (editingIndex.value === index) {
      clearFaqForm()
    }
  })
}

function getCategoryColor(category) {
  const colors = {
    general: 'bg-gradient-to-r from-blue-500 to-blue-600 text-white',
    technical: 'bg-gradient-to-r from-emerald-500 to-emerald-600 text-white',
    shipping: 'bg-gradient-to-r from-amber-500 to-amber-600 text-white',
    warranty: 'bg-gradient-to-r from-purple-500 to-purple-600 text-white',
    payment: 'bg-gradient-to-r from-rose-500 to-rose-600 text-white',
    usage: 'bg-gradient-to-r from-indigo-500 to-indigo-600 text-white'
  }
  return colors[category] || 'bg-gradient-to-r from-gray-500 to-gray-600 text-white'
}

function getCategoryLabel(category) {
  const labels = {
    general: 'عمومی',
    technical: 'فنی',
    shipping: 'حمل و نقل',
    warranty: 'گارانتی',
    payment: 'پرداخت',
    usage: 'نحوه استفاده'
  }
  return labels[category] || 'نامشخص'
}

function getStatusColor(status) {
  const colors = {
    published: 'bg-gradient-to-r from-green-500 to-green-600 text-white',
    draft: 'bg-gradient-to-r from-yellow-500 to-yellow-600 text-white',
    pending: 'bg-gradient-to-r from-orange-500 to-orange-600 text-white'
  }
  return colors[status] || 'bg-gradient-to-r from-gray-500 to-gray-600 text-white'
}

function getStatusLabel(status) {
  const labels = {
    published: 'منتشر شده',
    draft: 'پیش‌نویس',
    pending: 'در انتظار'
  }
  return labels[status] || 'نامشخص'
}
</script>

<template>
  <!-- سوالات متداول -->
  <div class="bg-white border border-gray-200 rounded-xl shadow-lg overflow-hidden text-right mb-8">
    <!-- Header -->
    <div class="bg-gradient-to-r from-indigo-600 to-purple-600 px-4 py-4">
      <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('faq')">
        <div class="flex items-center gap-3">
          <div class="bg-white/20 p-2 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div>
            <h3 class="text-xl font-bold text-white">مدیریت سوالات متداول محصول</h3>
            <p class="text-indigo-100 text-sm mt-1">ایجاد و مدیریت سوالات پرتکرار محصول</p>
          </div>
        </div>
        <div class="bg-white/20 p-2 rounded-lg">
          <svg class="w-5 h-5 text-white transition-transform duration-200" :class="{ 'rotate-180': sections.faq }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
          </svg>
        </div>
      </div>
    </div>

    <div v-show="sections.faq" class="px-4 py-4">
      <!-- آمار سوالات متداول -->
      <div class="grid grid-cols-2 md:grid-cols-4 gapx-4 py-4 mb-8">
        <div class="bg-gradient-to-br from-purple-50 to-purple-100 border border-purple-200 rounded-xl px-4 py-4 text-center">
          <div class="text-2xl font-bold text-purple-600 mb-1">{{ faqStats.total }}</div>
          <div class="text-sm text-purple-700 font-medium">کل سوالات</div>
        </div>
        <div class="bg-gradient-to-br from-green-50 to-green-100 border border-green-200 rounded-xl px-4 py-4 text-center">
          <div class="text-2xl font-bold text-green-600 mb-1">{{ faqStats.published }}</div>
          <div class="text-sm text-green-700 font-medium">منتشر شده</div>
        </div>
        <div class="bg-gradient-to-br from-amber-50 to-amber-100 border border-amber-200 rounded-xl px-4 py-4 text-center">
          <div class="text-2xl font-bold text-amber-600 mb-1">{{ faqStats.pending }}</div>
          <div class="text-sm text-amber-700 font-medium">در انتظار</div>
        </div>
        <div class="bg-gradient-to-br from-blue-50 to-blue-100 border border-blue-200 rounded-xl px-4 py-4 text-center">
          <div class="text-2xl font-bold text-blue-600 mb-1">{{ faqStats.categories }}</div>
          <div class="text-sm text-blue-700 font-medium">دسته‌بندی</div>
        </div>
      </div>

      <!-- فرم افزودن سوال جدید -->
      <div class="bg-gradient-to-br from-gray-50 to-gray-100 border border-gray-200 rounded-xl px-4 py-4 mb-8">
        <div class="flex items-center gap-3 mb-6">
          <div class="bg-gradient-to-r from-indigo-500 to-purple-500 p-2 rounded-lg">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
          </div>
          <h4 class="text-lg font-bold text-gray-800">
            {{ editingIndex >= 0 ? 'ویرایش سوال متداول' : 'افزودن سوال متداول جدید' }}
          </h4>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gapx-4 py-4">
          <div class="space-y-2">
            <label class="block text-sm font-semibold text-gray-700">سوال</label>
            <input
                v-model="newFaq.question"
                type="text"
                class="w-full px-4 py-3 text-right border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors"
                dir="rtl"
                placeholder="سوال خود را وارد کنید..."
            />
          </div>

          <div class="space-y-2">
            <label class="block text-sm font-semibold text-gray-700">دسته‌بندی</label>
            <select v-model="newFaq.category" class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors">
              <option value="">انتخاب دسته‌بندی</option>
              <option value="general">عمومی</option>
              <option value="technical">فنی</option>
              <option value="shipping">حمل و نقل</option>
              <option value="warranty">گارانتی</option>
              <option value="payment">پرداخت</option>
              <option value="usage">نحوه استفاده</option>
            </select>
          </div>

          <div class="lg:col-span-2 space-y-2">
            <label class="block text-sm font-semibold text-gray-700">پاسخ</label>
            <div class="border border-gray-300 rounded-lg overflow-hidden">
              <ClientOnly>
                <RichTextEditor
                    v-model="newFaq.answer"
                    :api-key="tinyApiKey"
                    lang="fa"
                    direction="rtl"
                    :height="250"
                />
              </ClientOnly>
            </div>
          </div>

          <div class="space-y-2">
            <label class="block text-sm font-semibold text-gray-700">اولویت نمایش</label>
            <input
                v-model="newFaq.priority"
                type="number"
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors"
                min="1"
                max="100"
                placeholder="1-100"
            />
          </div>

          <div class="space-y-2">
            <label class="block text-sm font-semibold text-gray-700">وضعیت</label>
            <select v-model="newFaq.status" class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors">
              <option value="published">منتشر شده</option>
              <option value="draft">پیش‌نویس</option>
              <option value="pending">در انتظار بررسی</option>
            </select>
          </div>

          <div class="lg:col-span-2 flex gap-3 pt-4">
            <button
                @click="addFaq"
                class="bg-gradient-to-r from-indigo-600 to-purple-600 text-white rounded-lg px-6 py-3 font-medium hover:from-indigo-700 hover:to-purple-700 transition-all duration-200 flex items-center gap-2 shadow-lg hover:shadow-xl"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
              </svg>
              {{ editingIndex >= 0 ? 'به‌روزرسانی سوال' : 'افزودن سوال' }}
            </button>
            <button
                @click="clearFaqForm"
                class="bg-gray-500 text-white rounded-lg px-6 py-3 font-medium hover:bg-gray-600 transition-colors"
            >
              {{ editingIndex >= 0 ? 'لغو ویرایش' : 'پاک کردن فرم' }}
            </button>
          </div>
        </div>
      </div>

      <!-- لیست سوالات متداول -->
      <div class="space-y-4">
        <h5 class="text-lg font-bold text-gray-800 flex items-center gap-2">
          <svg class="w-5 h-5 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          لیست سوالات متداول
        </h5>
        
        <div v-if="faqs.length === 0" class="text-center py-12">
          <svg class="w-16 h-16 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <p class="text-gray-500 text-lg">هنوز سوالی اضافه نشده است</p>
          <p class="text-gray-400 text-sm">اولین سوال متداول خود را اضافه کنید</p>
        </div>

        <div v-else class="grid gapx-4 py-4">
          <div v-for="(faq, index) in faqs" :key="index" 
               class="bg-white border border-gray-200 rounded-lg px-4 py-4 hover:shadow-lg transition-shadow duration-200"
               :class="{ 'ring-2 ring-indigo-500': editingIndex === index }">
            <div class="flex justify-between items-start gapx-4 py-4">
              <div class="flex-1">
                <div class="flex items-center gap-3 mb-3">
                  <span class="px-3 py-1 rounded-full text-xs font-medium" :class="getCategoryColor(faq.category)">
                    {{ getCategoryLabel(faq.category) }}
                  </span>
                  <span class="px-3 py-1 rounded-full text-xs font-medium" :class="getStatusColor(faq.status)">
                    {{ getStatusLabel(faq.status) }}
                  </span>
                  <span class="text-xs text-gray-500 bg-gray-100 px-2 py-1 rounded">
                    اولویت: {{ faq.priority }}
                  </span>
                </div>
                <h6 class="text-lg font-semibold text-gray-900 mb-3">{{ faq.question }}</h6>
                <div class="text-gray-600 prose prose-sm max-w-none" v-html="faq.answer"></div>
              </div>
              <div class="flex gap-2 shrink-0">
                <button
                    @click="editFaq(index)"
                    class="bg-blue-100 text-blue-600 hover:bg-blue-200 p-2 rounded-lg transition-colors"
                    title="ویرایش"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                </button>
                <button
                    @click="deleteFaq(index)"
                    class="bg-red-100 text-red-600 hover:bg-red-200 p-2 rounded-lg transition-colors"
                    title="حذف"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- تنظیمات نمایش -->
      <div class="mt-8 bg-gradient-to-br from-amber-50 to-orange-50 border border-amber-200 rounded-xl px-4 py-4">
        <div class="flex items-center gap-3 mb-4">
          <div class="bg-gradient-to-r from-amber-500 to-orange-500 p-2 rounded-lg">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
          </div>
          <h5 class="text-lg font-bold text-gray-800">تنظیمات نمایش سوالات متداول</h5>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
          <label class="flex items-center gap-3 p-3 bg-white rounded-lg border border-amber-200 cursor-pointer hover:bg-amber-50 transition-colors">
            <input type="checkbox" class="w-4 h-4 text-indigo-600 border-gray-300 rounded focus:ring-indigo-500" />
            <span class="text-sm font-medium text-gray-700">نمایش تب سوالات متداول</span>
          </label>
          <label class="flex items-center gap-3 p-3 bg-white rounded-lg border border-amber-200 cursor-pointer hover:bg-amber-50 transition-colors">
            <input type="checkbox" class="w-4 h-4 text-indigo-600 border-gray-300 rounded focus:ring-indigo-500" />
            <span class="text-sm font-medium text-gray-700">نمایش جستجو در سوالات</span>
          </label>
          <label class="flex items-center gap-3 p-3 bg-white rounded-lg border border-amber-200 cursor-pointer hover:bg-amber-50 transition-colors">
            <input type="checkbox" class="w-4 h-4 text-indigo-600 border-gray-300 rounded focus:ring-indigo-500" />
            <span class="text-sm font-medium text-gray-700">نمایش فیلتر دسته‌بندی</span>
          </label>
        </div>
      </div>
    </div>
  </div>
</template>
