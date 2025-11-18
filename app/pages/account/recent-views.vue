<template>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- لایوت اصلی با سایدبار -->
      <div class="flex flex-col lg:flex-row gap-8 justify-center">
        <!-- محتوای اصلی: فقط لیست بازدیدهای اخیر -->
        <div class="flex-1 max-w-4xl" dir="rtl">
          <!-- عنوان -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
            <h1 class="text-sm font-bold text-gray-900 text-right">بازدیدهای اخیر</h1>
          </div>
          
          <div>
            <div v-if="loading" class="loading text-center py-8 text-gray-500">در حال بارگذاری...</div>
            <div v-else>
              <div v-if="views.length === 0" class="empty text-center py-8 text-gray-500">هیچ بازدیدی ثبت نشده است.</div>
              <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                <div v-for="view in views" :key="view.id" class="bg-white rounded-lg border border-gray-200 p-6 flex flex-col items-center hover:shadow-lg transition-shadow">
                  <a :href="`/product/${view.product.slug}`" class="w-full flex flex-col items-center">
                    <img :src="view.product.image_url || view.product.image" alt="product" class="w-32 h-32 object-cover rounded-lg mb-2">
                    <div class="text-center">
                      <div class="text-sm font-medium text-gray-900 mb-1">{{ view.product.name }}</div>
                      <div class="text-lg font-bold text-gray-900 mb-1">{{ view.product.price.toLocaleString() }} تومان</div>
                      <div class="text-xs text-gray-500 mb-1">
                        <div>اولین بازدید: {{ formatDate(view.viewed_at) }}</div>
                        <div v-if="view.view_count > 1" class="text-blue-600 font-medium">
                          {{ view.view_count }} بار بازدید کرده‌اید
                        </div>
                      </div>
                    </div>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- سایدبار -->
        <div class="lg:w-80 flex-shrink-0 lg:ml-8">
          <AccountSidebar />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import AccountSidebar from '@/components/account/AccountSidebar.vue'

// میدلور حذف شده
definePageMeta({ layout: 'default' })
useHead({ title: 'بازدیدهای اخیر - حساب کاربری' })

const views = ref([])
const loading = ref(true)

function formatDate(dateStr) {
  const d = new Date(dateStr)
  return d.toLocaleDateString('fa-IR') + ' ' + d.toLocaleTimeString('fa-IR', { hour: '2-digit', minute: '2-digit' })
}

onMounted(async () => {
  try {
    const resp = await fetch('/api/recent-views', { credentials: 'include' })
    const json = await resp.json()
    views.value = json && json.data ? json.data : []
  } catch (e) {
    console.error('Failed to load recent views', e)
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.recent-views {
  max-width: 700px;
  margin: 0 auto;
  padding: 2rem 1rem;
}
.views-list {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
}
.view-item {
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px #eee;
  padding: 1rem;
  width: 210px;
  display: flex;
  flex-direction: column;
  align-items: center;
}
.product-image {
  width: 120px;
  height: 120px;
  object-fit: cover;
  border-radius: 6px;
  margin-bottom: 0.5rem;
}
.product-info {
  text-align: center;
}
.product-name {
  font-weight: bold;
  margin-bottom: 0.3rem;
}
.product-price {
  color: #1976d2;
  margin-bottom: 0.3rem;
}
.viewed-at {
  font-size: 0.85em;
  color: #888;
}
.loading, .empty {
  text-align: center;
  margin: 2rem 0;
  color: #888;
}
</style>
