<template>
  <!-- لوگو با لینک به صفحه اصلی یا مقدار props.href -->
  <div :class="alignmentClass">
    <NuxtLink :to="href" class="logo-link">
      <img :src="logoSrc" alt="Logo" class="logo-image" />
    </NuxtLink>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'

interface ShopSettingsResponse {
  status: string
  data: {
    shopNameFa?: string
    shopNameEn?: string
    logo?: string
    logoRetina?: string
    favicon?: string
    defaultLanguage?: string
    timezone?: string
    defaultCurrency?: string
    maintenanceMode?: boolean
    maintenanceMessage?: string
    phones?: string[]
    email?: string
    adminPhones?: string[]
    address?: string
    latitude?: string
    longitude?: string
    workingHours?: string[]
    shortDescription?: string
  }
}

const props = defineProps<{ 
  src?: string; 
  href?: string;
  align?: 'left' | 'center' | 'right';
}>()

const href = props.href ?? '/'

const logoSrc = ref(props.src ?? '/default-product.svg')

// محاسبه کلاس چینش بر اساس prop align
const alignmentClass = computed(() => {
  const base = 'flex items-center h-full w-full'
  switch (props.align) {
    case 'left':
      return `${base} justify-start`
    case 'center':
      return `${base} justify-center`
    case 'right':
      return `${base} justify-end`
    default:
      return `${base} justify-center`
  }
})

// بارگذاری لوگو از تنظیمات عمومی فروشگاه
const loadLogo = async () => {
  try {
    const response = await $fetch<ShopSettingsResponse>('/api/shop-settings')
    
    if (response.status === 'success' && response.data?.logo) {
      logoSrc.value = response.data.logo
      // لوگو بارگذاری شد
    } else {
      // لوگو در تنظیمات یافت نشد، از لوگوی پیش‌فرض استفاده می‌شود
    }
  } catch {
    // خطا در بارگذاری لوگو
    // در صورت خطا، لوگوی پیش‌فرض نمایش داده می‌شود
  }
}

onMounted(() => {
  loadLogo()
})
</script>

<style scoped>
.logo-link {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
  padding: 0 8px;
  box-sizing: border-box;
}

.logo-image {
  max-height: 100%;
  width: auto;
  object-fit: contain;
}
</style>
