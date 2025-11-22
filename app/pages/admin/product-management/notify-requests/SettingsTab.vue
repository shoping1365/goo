<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import TemplateButton from '@/components/common/TemplateButton.vue'
import TemplateCard from '@/components/common/TemplateCard.vue'
import { useNotifyRequests } from '~/composables/useNotifyRequests'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

interface Props {
  activeNavItem: string
}

const props = defineProps<Props>()

const { 
  showSettings,
  activeSettingsTab,
  settings,
  successRate,
  todayRequests
} = useNotifyRequests()
</script>

<template>
  <!-- تب تنظیمات -->
  <div v-if="props.activeNavItem === 'settings'" class="w-full">
    <div class="bg-white rounded-xl shadow-md border border-gray-100 overflow-hidden">
      <div class="bg-gray-50 px-6 py-4 border-b border-gray-100">
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-lg font-semibold text-gray-900">تنظیمات سیستم اطلاع‌رسانی</h3>
            <p class="text-sm text-gray-600 mt-1">پیکربندی و مدیریت تنظیمات</p>
          </div>
          <TemplateButton
            bg-gradient="bg-gradient-to-r from-blue-400 to-blue-600"
            text-color="text-white"
            hover-class="hover:from-blue-500 hover:to-blue-700 hover:shadow-lg hover:scale-105"
            focus-class="focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            size="medium"
            @click="showSettings = true"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
            </svg>
            تنظیمات پیشرفته
          </TemplateButton>
        </div>
      </div>

      <div class="p-6">
        <!-- کارت‌های تنظیمات سریع -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
          <!-- عمومی -->
          <div class="bg-blue-50 rounded-lg p-6 border border-blue-200">
            <div class="flex items-center gap-3 mb-3">
              <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4"></path>
                </svg>
              </div>
              <div>
                <h4 class="font-medium text-blue-900">تنظیمات عمومی</h4>
                <p class="text-sm text-blue-700">پیکربندی پایه</p>
              </div>
            </div>
            <div class="space-y-2 text-sm text-blue-800">
              <div class="flex justify-between">
                <span>تاخیر ارسال:</span>
                <span>{{ settings.delayMinutes }} دقیقه</span>
              </div>
              <div class="flex justify-between">
                <span>محدودیت روزانه:</span>
                <span>{{ settings.dailyLimit }}</span>
              </div>
            </div>
          </div>

          <!-- ایمیل -->
          <div class="bg-green-50 rounded-lg p-6 border border-green-200">
            <div class="flex items-center gap-3 mb-3">
              <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
                </svg>
              </div>
              <div>
                <h4 class="font-medium text-green-900">ایمیل</h4>
                <p class="text-sm text-green-700">{{ settings.emailNotifications ? 'فعال' : 'غیرفعال' }}</p>
              </div>
            </div>
            <div class="space-y-2 text-sm text-green-800">
              <div class="flex justify-between">
                <span>ارائه‌دهنده:</span>
                <span>{{ settings.emailProvider.toUpperCase() }}</span>
              </div>
              <div class="flex justify-between">
                <span>فرستنده:</span>
                <span class="truncate">{{ settings.emailFromName }}</span>
              </div>
            </div>
          </div>

          <!-- پیامک -->
          <div class="bg-purple-50 rounded-lg p-6 border border-purple-200">
            <div class="flex items-center gap-3 mb-3">
              <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
                </svg>
              </div>
              <div>
                <h4 class="font-medium text-purple-900">پیامک</h4>
                <p class="text-sm text-purple-700">{{ settings.smsNotifications ? 'فعال' : 'غیرفعال' }}</p>
              </div>
            </div>
            <div class="space-y-2 text-sm text-purple-800">
              <div class="flex justify-between">
                <span>ارائه‌دهنده:</span>
                <span>{{ settings.smsProvider === 'kavenegar' ? 'کاوه‌نگار' : settings.smsProvider === 'melipayamak' ? 'ملی‌پیامک' : 'فراپیامک' }}</span>
              </div>
              <div class="flex justify-between">
                <span>وضعیت:</span>
                <span class="text-xs bg-purple-200 px-2 py-1 rounded">{{ settings.smsApiKey ? 'تنظیم شده' : 'نیاز به تنظیم' }}</span>
              </div>
            </div>
          </div>

          <!-- اطلاع‌رسانی خودکار -->
          <div class="bg-orange-50 rounded-lg p-6 border border-orange-200">
            <div class="flex items-center gap-3 mb-3">
              <div class="w-10 h-10 bg-orange-500 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
                </svg>
              </div>
              <div>
                <h4 class="font-medium text-orange-900">اطلاع‌رسانی خودکار</h4>
                <p class="text-sm text-orange-700">{{ settings.autoNotifyStock && settings.autoNotifyDiscount ? 'فعال' : 'جزئی' }}</p>
              </div>
            </div>
            <div class="space-y-2 text-sm text-orange-800">
              <div class="flex justify-between">
                <span>موجودی:</span>
                <span>{{ settings.autoNotifyStock ? '✓' : '✗' }}</span>
              </div>
              <div class="flex justify-between">
                <span>تخفیف:</span>
                <span>{{ settings.autoNotifyDiscount ? '✓' : '✗' }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- دسترسی سریع به تنظیمات -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          <TemplateCard
            title="تنظیمات عمومی"
            :value="''"
            variant="blue"
            @click="showSettings = true; activeSettingsTab = 'general'"
          >
            <template #icon>
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4"></path>
              </svg>
            </template>
          </TemplateCard>

          <TemplateCard
            title="تنظیمات ایمیل"
            :value="''"
            variant="green"
            @click="showSettings = true; activeSettingsTab = 'email'"
          >
            <template #icon>
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
              </svg>
            </template>
          </TemplateCard>

          <TemplateCard
            title="تنظیمات پیامک"
            :value="''"
            variant="purple"
            @click="showSettings = true; activeSettingsTab = 'sms'"
          >
            <template #icon>
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
              </svg>
            </template>
          </TemplateCard>

          <TemplateCard
            title="مشاهده لاگ‌ها"
            :value="''"
            variant="orange"
            @click="showSettings = true; activeSettingsTab = 'logs'"
          >
            <template #icon>
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01"></path>
              </svg>
            </template>
          </TemplateCard>
        </div>

        <!-- آمار سریع -->
        <div class="mt-8 grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-lg p-6 text-white">
            <div class="flex items-center justify-between">
              <div>
                <h3 class="text-lg font-semibold">وضعیت سیستم</h3>
                <p class="text-blue-100 text-sm mt-1">همه سرویس‌ها فعال</p>
              </div>
              <div class="w-12 h-12 bg-white/20 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
            </div>
            <div class="mt-4 flex items-center gap-2">
              <div class="w-2 h-2 bg-green-400 rounded-full"></div>
              <span class="text-sm">آخرین بررسی: اکنون</span>
            </div>
          </div>

          <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-lg p-6 text-white">
            <div class="flex items-center justify-between">
              <div>
                <h3 class="text-lg font-semibold">نرخ موفقیت</h3>
                <p class="text-green-100 text-sm mt-1">ارسال اطلاع‌رسانی‌ها</p>
              </div>
              <div class="w-12 h-12 bg-white/20 rounded-lg flex items-center justify-center">
                <span class="text-xl font-bold">{{ successRate }}%</span>
              </div>
            </div>
            <div class="mt-4">
              <div class="w-full bg-green-400/30 rounded-full h-2">
                <div class="bg-green-200 h-2 rounded-full" :style="{ width: successRate + '%' }"></div>
              </div>
            </div>
          </div>

          <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-lg p-6 text-white">
            <div class="flex items-center justify-between">
              <div>
                <h3 class="text-lg font-semibold">درخواست‌های امروز</h3>
                <p class="text-purple-100 text-sm mt-1">اطلاع‌رسانی جدید</p>
              </div>
              <div class="w-12 h-12 bg-white/20 rounded-lg flex items-center justify-center">
                <span class="text-xl font-bold">{{ todayRequests }}</span>
              </div>
            </div>
            <div class="mt-4 text-sm text-purple-100">
              <span>+12% نسبت به دیروز</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>
