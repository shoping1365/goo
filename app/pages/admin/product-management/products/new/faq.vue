<script setup lang="ts">
import { computed, reactive, ref } from 'vue';
import RichTextEditor from '~/components/common/RichTextEditor.vue';
import { useConfirmDialog } from '~/composables/useConfirmDialog';

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
  <div class="bg-gray-50 border rounded shadow p-6 text-right mb-6">
    <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('faq')">
      <h3 class="text-sm font-semibold text-gray-700">مدیریت سوالات متداول محصول</h3>
      <span class="text-gray-500 text-lg">{{ sections.faq ? '−' : '+' }}</span>
    </div>

    <div v-show="sections.faq" class="mt-4">
      <!-- آمار سوالات متداول -->
      <div class="grid grid-cols-2 md:grid-cols-4 gap-6 mb-6 p-6 bg-gradient-to-r from-purple-50 to-blue-50 rounded-lg border border-blue-200">
        <div class="text-center">
          <div class="text-lg font-bold text-purple-600">{{ faqStats.total }}</div>
          <div class="text-xs text-gray-600">کل سوالات</div>
        </div>
        <div class="text-center">
          <div class="text-lg font-bold text-green-600">{{ faqStats.published }}</div>
          <div class="text-xs text-gray-600">منتشر شده</div>
        </div>
        <div class="text-center">
          <div class="text-lg font-bold text-yellow-600">{{ faqStats.pending }}</div>
          <div class="text-xs text-gray-600">در انتظار</div>
        </div>
        <div class="text-center">
          <div class="text-lg font-bold text-blue-600">{{ faqStats.categories }}</div>
          <div class="text-xs text-gray-600">دسته‌بندی</div>
        </div>
      </div>

      <!-- فرم افزودن سوال جدید -->
      <div class="border rounded-lg p-6 bg-gray-50 mb-6 border-blue-200">
        <h4 class="text-sm font-semibold text-gray-700 mb-4 flex items-center gap-2">
          <svg class="w-4 h-4 text-purple-600" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z" clip-rule="evenodd" />
          </svg>
          افزودن سوال متداول جدید
        </h4>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="flex flex-col gap-2">
            <label class="text-xs text-gray-700 font-semibold">سوال</label>
            <input
                v-model="newFaq.question"
                type="text"
                class="input input-bordered w-full text-right"
                dir="rtl"
                placeholder="سوال شما..."
            />
          </div>

          <div class="flex flex-col gap-2">
            <label class="text-xs text-gray-700 font-semibold">دسته‌بندی</label>
            <select v-model="newFaq.category" class="input input-bordered w-full">
              <option value="">انتخاب دسته‌بندی</option>
              <option value="general">عمومی</option>
              <option value="technical">فنی</option>
              <option value="shipping">حمل و نقل</option>
              <option value="warranty">گارانتی</option>
              <option value="payment">پرداخت</option>
              <option value="usage">نحوه استفاده</option>
            </select>
          </div>

          <div class="md:col-span-2 flex flex-col gap-2">
            <label class="text-xs text-gray-700 font-semibold">پاسخ</label>
            <ClientOnly>
              <RichTextEditor
                  v-model="newFaq.answer"
                  lang="fa"
                  direction="rtl"
                  :height="200"
              />
            </ClientOnly>
          </div>

          <div class="flex flex-col gap-2">
            <label class="text-xs text-gray-700 font-semibold">اولویت نمایش</label>
            <input
                v-model="newFaq.priority"
                type="number"
                class="input input-bordered w-full"
                min="1"
                max="100"
                placeholder="1-100"
            />
          </div>

          <div class="flex flex-col gap-2">
            <label class="text-xs text-gray-700 font-semibold">وضعیت</label>
            <select v-model="newFaq.status" class="input input-bordered w-full">
              <option value="published">منتشر شده</option>
              <option value="draft">پیش‌نویس</option>
              <option value="pending">در انتظار بررسی</option>
            </select>
          </div>

          <div class="md:col-span-2 flex gap-2 mt-2">
            <button
                class="bg-purple-600 text-white rounded px-4 py-2 text-xs hover:bg-purple-700 transition flex items-center gap-2"
                @click="addFaq"
            >
              <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z" clip-rule="evenodd" />
              </svg>
              افزودن سوال
            </button>
            <button
                class="bg-gray-500 text-white rounded px-4 py-2 text-xs hover:bg-gray-600 transition"
                @click="clearFaqForm"
            >
              پاک کردن فرم
            </button>
          </div>
        </div>
      </div>

      <!-- لیست سوالات متداول -->
      <div class="overflow-x-auto">
        <table class="min-w-full text-xs text-right rtl border-collapse border border-blue-200">
          <thead>
            <tr class="bg-gradient-to-r from-purple-100 to-blue-100 border-b border-blue-200">
              <th class="px-3 py-3 text-gray-700 font-semibold text-right border-r border-blue-200">سوال</th>
              <th class="px-3 py-3 text-gray-700 font-semibold text-right border-r border-blue-200">دسته‌بندی</th>
              <th class="px-3 py-3 text-gray-700 font-semibold text-right border-r border-blue-200">وضعیت</th>
              <th class="px-3 py-3 text-gray-700 font-semibold text-center">عملیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(faq, index) in faqs" :key="index" class="border-b border-blue-200 hover:bg-gray-50">
              <td class="px-3 py-3 text-right border-r border-blue-200 max-w-xs">
                <div class="font-medium text-gray-900 truncate">{{ faq.question }}</div>
              </td>
              <td class="px-3 py-3 text-right border-r border-blue-200">
                <span class="px-2 py-1 rounded-full text-xs" :class="getCategoryColor(faq.category)">
                  {{ getCategoryLabel(faq.category) }}
                </span>
              </td>
              <td class="px-3 py-3 text-right border-r border-blue-200">
                <span class="px-2 py-1 rounded-full text-xs" :class="getStatusColor(faq.status)">
                  {{ getStatusLabel(faq.status) }}
                </span>
              </td>
              <td class="px-3 py-3 text-center">
                <div class="flex justify-center gap-1">
                  <button
                      class="text-blue-600 hover:text-blue-800 p-1"
                      title="ویرایش"
                      @click="editFaq(index)"
                  >
                    ویرایش
                  </button>
                  <button
                      class="text-red-600 hover:text-red-800 p-1"
                      title="حذف"
                      @click="deleteFaq(index)"
                  >
                    حذف
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- تنظیمات نمایش -->
      <div class="mt-6 p-6 bg-yellow-50 rounded-lg border border-yellow-200">
        <h5 class="text-sm font-semibold text-gray-700 mb-3">تنظیمات نمایش سوالات متداول</h5>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="flex items-center gap-2">
            <input id="show-faq-tab" type="checkbox" class="form-checkbox" />
            <label for="show-faq-tab" class="text-xs text-gray-700">نمایش تب سوالات متداول</label>
          </div>
          <div class="flex items-center gap-2">
            <input id="show-faq-search" type="checkbox" class="form-checkbox" />
            <label for="show-faq-search" class="text-xs text-gray-700">نمایش جستجو در سوالات</label>
          </div>
          <div class="flex items-center gap-2">
            <input id="show-category-filter" type="checkbox" class="form-checkbox" />
            <label for="show-category-filter" class="text-xs text-gray-700">نمایش فیلتر دسته‌بندی</label>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
