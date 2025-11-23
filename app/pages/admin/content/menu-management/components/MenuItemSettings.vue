<template>
  <div v-if="hasAccess" class="border-t border-gray-100 bg-gray-50">
    <div class="p-6">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- Basic Settings -->
        <div class="space-y-4">
          <h4 class="font-medium text-gray-900 flex items-center space-x-2 space-x-reverse">
            <svg class="w-4 h-4 text-gray-500" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947zM10 13a3 3 0 100-6 3 3 0 000 6z" clip-rule="evenodd"/>
            </svg>
            <span>تنظیمات اصلی</span>
          </h4>
          
          <!-- Title -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">عنوان</label>
            <input
              :value="item.title"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
              @input="(e: Event) => $emit('update-item', { ...item, title: (e.target as HTMLInputElement).value })"
            />
          </div>

          <!-- Path -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">مسیر</label>
            <input
              :value="item.path"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
              @input="(e: Event) => $emit('update-item', { ...item, path: (e.target as HTMLInputElement).value })"
            />
          </div>

          <!-- Icon/Image -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">آیکون یا تصویر</label>
            <div class="flex space-x-2 space-x-reverse">
              <input
                :value="item.icon"
                type="text"
                placeholder="کلاس آیکون یا JSON تصویر"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
                readonly
              />
              <select
                :value="item.iconType"
                class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm bg-white"
                @change="(e: Event) => $emit('update-item', { ...item, iconType: (e.target as HTMLSelectElement).value })"
              >
                <option value="icon">آیکون</option>
                <option value="image">تصویر</option>
              </select>
              <button
                class="px-3 py-2 bg-blue-100 hover:bg-blue-200 text-blue-700 rounded-md text-sm font-medium transition-colors"
                @click="$emit('open-icon-selector')"
              >
                {{ item.iconType === 'image' ? 'انتخاب تصویر' : 'انتخاب آیکون' }}
              </button>
            </div>
          </div>

          <!-- Enabled/Disabled - More Prominent -->
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg border-2 border-gray-200">
            <div class="flex items-center space-x-3 space-x-reverse">
              <input
                :id="`enabled-${item.id ?? item.clientId}`"
                :checked="item.enabled"
                type="checkbox"
                class="w-5 h-5 text-blue-600 border-gray-300 rounded focus:ring-blue-500 cursor-pointer"
                @change="(e: Event) => $emit('update-item', { ...item, enabled: (e.target as HTMLInputElement).checked })"
              />
              <label :for="`enabled-${item.id ?? item.clientId}`" class="text-sm font-semibold text-gray-900 cursor-pointer">
                وضعیت نمایش
              </label>
            </div>
            <span 
              :class="item.enabled ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'"
              class="px-3 py-1 text-xs rounded-full font-bold"
            >
              {{ item.enabled ? '✓ فعال' : '✗ غیرفعال' }}
            </span>
          </div>
        </div>

        <!-- Advanced Settings -->
        <div class="space-y-4">
          <h4 class="font-medium text-gray-900 flex items-center space-x-2 space-x-reverse">
            <svg class="w-4 h-4 text-gray-500" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M12.316 3.051a1 1 0 01.633 1.265l-4 12a1 1 0 11-1.898-.632l4-12a1 1 0 011.265-.633zM5.707 6.293a1 1 0 010 1.414L3.414 10l2.293 2.293a1 1 0 11-1.414 1.414l-3-3a1 1 0 010-1.414l3-3a1 1 0 011.414 0zm8.586 0a1 1 0 011.414 0l3 3a1 1 0 010 1.414l-3 3a1 1 0 11-1.414-1.414L16.586 10l-2.293-2.293a1 1 0 010-1.414z" clip-rule="evenodd"/>
            </svg>
            <span>تنظیمات پیشرفته</span>
          </h4>
          
          <!-- Open in New Tab -->
          <div class="flex items-center space-x-3 space-x-reverse">
            <input
              :id="`new-tab-${item.id ?? item.clientId}`"
              :checked="item.openInNewTab"
              type="checkbox"
              class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
              @change="(e: Event) => $emit('update-item', { ...item, openInNewTab: (e.target as HTMLInputElement).checked })"
            />
            <label :for="`new-tab-${item.id ?? item.clientId}`" class="text-sm font-medium text-gray-700">
              باز شدن در تب جدید
            </label>
          </div>

          <!-- Badge -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">نشان (Badge)</label>
            <input
              :value="item.badge"
              type="text"
              placeholder="متن نشان"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
              @input="(e: Event) => $emit('update-item', { ...item, badge: (e.target as HTMLInputElement).value })"
            />
          </div>

          <!-- Badge Color -->
          <div v-if="item.badge">
            <label class="block text-sm font-medium text-gray-700 mb-1">رنگ نشان</label>
            <select
              :value="item.badgeColor"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
              @change="(e: Event) => $emit('update-item', { ...item, badgeColor: (e.target as HTMLSelectElement).value })"
            >
              <option value="red">قرمز</option>
              <option value="green">سبز</option>
              <option value="blue">آبی</option>
              <option value="yellow">زرد</option>
              <option value="purple">بنفش</option>
              <option value="gray">خاکستری</option>
            </select>
          </div>

          <!-- Mega Menu Settings -->
          <div class="border-t border-gray-200 pt-4">
            <h5 class="font-medium text-gray-900 mb-3 flex items-center space-x-2 space-x-reverse">
              <svg class="w-4 h-4 text-purple-500" fill="currentColor" viewBox="0 0 20 20">
                <path d="M3 4a1 1 0 011-1h12a1 1 0 011 1v2a1 1 0 01-1 1H4a1 1 0 01-1-1V4zM3 10a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H4a1 1 0 01-1-1v-6zM14 9a1 1 0 00-1 1v6a1 1 0 001 1h2a1 1 0 001-1v-6a1 1 0 00-1-1h-2z"/>
              </svg>
              <span>تنظیمات مگا منو</span>
            </h5>
            
            <!-- Is Mega Menu -->
            <div class="flex items-center space-x-3 space-x-reverse mb-3">
              <input
                :id="`mega-menu-${item.id ?? item.clientId}`"
                :checked="item.isMegaMenu"
                type="checkbox"
                class="w-4 h-4 text-purple-600 border-gray-300 rounded focus:ring-purple-500"
                @change="(e: Event) => $emit('update-item', { ...item, isMegaMenu: (e.target as HTMLInputElement).checked })"
              />
              <label :for="`mega-menu-${item.id ?? item.clientId}`" class="text-sm font-medium text-gray-700">
                مگا منو باشد
              </label>
            </div>

            <!-- Mega Menu Width -->
            <div v-if="item.isMegaMenu" class="mb-3">
              <label class="block text-sm font-medium text-gray-700 mb-1">عرض مگا منو</label>
              <select
                :value="item.megaWidth"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
                @change="(e: Event) => $emit('update-item', { ...item, megaWidth: (e.target as HTMLSelectElement).value })"
              >
                <option value="full">کامل</option>
                <option value="container">کانتینر</option>
                <option value="custom">سفارشی</option>
              </select>
            </div>

            <!-- Mega Menu Columns -->
            <div v-if="item.isMegaMenu" class="mb-3">
              <label class="block text-sm font-medium text-gray-700 mb-1">تعداد ستون‌ها</label>
              <select
                :value="item.megaColumns"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
                @change="(e: Event) => $emit('update-item', { ...item, megaColumns: (e.target as HTMLSelectElement).value })"
              >
                <option value="2">2 ستون</option>
                <option value="3">3 ستون</option>
                <option value="4">4 ستون</option>
                <option value="5">5 ستون</option>
              </select>
            </div>
          </div>
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

defineProps({
  item: Object,
})

defineEmits(['update-item', 'open-icon-selector', 'open-image-selector'])
</script>
