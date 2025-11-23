<template>
  <div v-if="hasAccess" class="font-iranyekan">
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 px-4 py-4">
      <div class="flex items-center justify-between mb-4">
        <h3 class="text-lg font-bold text-gray-900 font-iranyekan">پیش‌نمایش فوتر</h3>
        <button
          class="text-gray-400 hover:text-gray-600 transition-colors font-iranyekan"
          @click="$emit('close')"
        >
          <span class="text-xl">✕</span>
        </button>
      </div>
      
      <div class="border-t border-gray-200 pt-4">
        <div v-if="footer" class="space-y-4">
          <!-- عنوان -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2 font-iranyekan">عنوان:</label>
            <div class="text-lg font-semibold text-gray-900 font-iranyekan">{{ footer.title }}</div>
          </div>
          
          <!-- وضعیت -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2 font-iranyekan">وضعیت:</label>
            <span
class="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium font-iranyekan"
                  :class="footer.is_active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'">
              <span
class="w-2 h-2 rounded-full mr-2"
                    :class="footer.is_active ? 'bg-green-500' : 'bg-red-500'"></span>
              {{ footer.is_active ? 'فعال' : 'غیرفعال' }}
            </span>
          </div>
          
          <!-- نوع نمایش -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2 font-iranyekan">نوع نمایش:</label>
            <div class="text-sm text-gray-900 font-iranyekan">{{ getPageSelectionLabel(footer.page_selection) }}</div>
          </div>
          
          <!-- صفحات خاص -->
          <div v-if="footer.page_selection === 'specific' && footer.specific_pages">
            <label class="block text-sm font-medium text-gray-700 mb-2 font-iranyekan">صفحات خاص:</label>
            <div class="text-sm text-gray-900 font-iranyekan bg-gray-50 p-3 rounded-lg">
              {{ footer.specific_pages }}
            </div>
          </div>
          
          <!-- صفحات مستثنی -->
          <div v-if="footer.page_selection === 'exclude' && footer.excluded_pages">
            <label class="block text-sm font-medium text-gray-700 mb-2 font-iranyekan">صفحات مستثنی:</label>
            <div class="text-sm text-gray-900 font-iranyekan bg-gray-50 p-3 rounded-lg">
              {{ footer.excluded_pages }}
            </div>
          </div>
          
          <!-- محتوا -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2 font-iranyekan">محتوای فوتر:</label>
            <div class="bg-gray-50 px-4 py-4 rounded-lg border border-gray-200">
              <SanitizedHtml :content="footer?.content || ''" class="text-sm text-gray-900 leading-relaxed font-iranyekan" />
            </div>
          </div>
          
          <!-- تاریخ ایجاد -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2 font-iranyekan">تاریخ ایجاد:</label>
            <div class="text-sm text-gray-900 font-iranyekan">{{ formatDate(footer.created_at) }}</div>
          </div>
        </div>
        
        <div v-else class="text-center text-gray-500 py-8 font-iranyekan">
          <div class="text-4xl mb-2">📄</div>
          <p>هیچ فوتری برای نمایش انتخاب نشده است</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';
import SanitizedHtml from '~/components/common/SanitizedHtml.vue'

// احراز هویت
const { user, isAuthenticated } = useAuth();

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false;
  }

  const userRole = user.value?.role?.toLowerCase() || '';
  const adminRoles = ['admin', 'developer'];
  return adminRoles.includes(userRole);
});

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async (): Promise<void> => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false });
  }
};

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth();
});

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth();
  }
});

// تعریف props
const _props = defineProps({
  footer: {
    type: Object,
    default: null
  }
})

// تعریف events
const _emit = defineEmits(['close'])

// Footer data
const footer = _props.footer

// برچسب نوع نمایش صفحه
const getPageSelectionLabel = (pageSelection: string) => {
  const labels: Record<string, string> = {
    'all': 'همه صفحات',
    'specific': 'صفحات خاص',
    'exclude': 'به جز صفحات خاص'
  }
  return labels[pageSelection] || pageSelection
}

// فرمت تاریخ
const formatDate = (dateString: string) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}
</script>

<style scoped>
.font-iranyekan {
  font-family: 'Yekan', 'Tahoma', sans-serif;
}

/* بهبود انیمیشن‌ها */
.transition-all {
  transition-property: all;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}

/* بهبود hover effects */
.hover\:shadow-lg:hover {
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

.hover\:shadow-md:hover {
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
}
</style>

