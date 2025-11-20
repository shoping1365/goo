<template>
  <div class="min-h-screen" dir="rtl">
    <!-- Header -->
    <div class="bg-white border-b border-gray-200 px-6 py-4">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-orange-700">آمار موجودی انبار</h1>
          <p class="text-gray-500 mt-1">تحلیل جامع موجودی، گردش کالا و عملکرد انبار</p>
        </div>
        <div class="flex gap-2">
          <button class="px-4 py-2 bg-orange-500 text-white rounded-lg shadow hover:bg-orange-600 transition">
            گزارش تفصیلی موجودی
          </button>
          <button class="px-4 py-2 bg-green-500 text-white rounded-lg shadow hover:bg-green-600 transition">
            خروجی اکسل
          </button>
          <button class="px-4 py-2 bg-blue-500 text-white rounded-lg shadow hover:bg-blue-600 transition">
            درخواست تامین
          </button>
        </div>
      </div>
    </div>

    <!-- Filters -->
    <div class="px-6 py-4 bg-white border-b">
      <div class="flex flex-wrap gap-6">
        <select class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-orange-500">
          <option>همه دسته‌بندی‌ها</option>
          <option>الکترونیک</option>
          <option>پوشاک</option>
          <option>کتاب</option>
          <option>خانه و آشپزخانه</option>
        </select>
        <select class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-orange-500">
          <option>همه انبارها</option>
          <option>انبار مرکزی تهران</option>
          <option>انبار شعبه اصفهان</option>
          <option>انبار شعبه مشهد</option>
        </select>
        <select class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-orange-500">
          <option>وضعیت موجودی</option>
          <option>موجود</option>
          <option>کمبود</option>
          <option>بحرانی</option>
          <option>ناموجود</option>
        </select>
        <input type="date" class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-orange-500" value="2024-01-01">
        <input type="date" class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-orange-500" value="2024-12-31">
        <button class="px-4 py-2 bg-orange-500 text-white rounded-lg text-sm hover:bg-orange-600 transition">اعمال فیلتر</button>
        <button class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg text-sm hover:bg-gray-200 transition">پاک کردن</button>
      </div>
    </div>
    
    <!-- Summary Cards -->
    <div class="px-6 py-6">
      <div class="grid grid-cols-2 md:grid-cols-4 xl:grid-cols-8 gap-6">
        <div class="bg-gradient-to-br from-orange-50 to-orange-100 rounded-lg shadow-sm p-6 border border-orange-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-orange-600">کل محصولات</p>
              <p class="text-xl font-bold text-orange-700">{{ inventoryStats.totalProducts.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-orange-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-orange-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 2L3 7v11a2 2 0 002 2h10a2 2 0 002-2V7l-7-5z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-green-50 to-green-100 rounded-lg shadow-sm p-6 border border-green-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-green-600">موجود</p>
              <p class="text-xl font-bold text-green-700">{{ inventoryStats.inStock.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-green-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-green-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-red-50 to-red-100 rounded-lg shadow-sm p-6 border border-red-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-red-600">کمبود</p>
              <p class="text-xl font-bold text-red-700">{{ inventoryStats.lowStock.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-red-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-red-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-gray-50 to-gray-100 rounded-lg shadow-sm p-6 border border-gray-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-gray-600">ناموجود</p>
              <p class="text-xl font-bold text-gray-700">{{ inventoryStats.outOfStock.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-gray-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-gray-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-purple-50 to-purple-100 rounded-lg shadow-sm p-6 border border-purple-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-purple-600">ارزش موجودی</p>
              <p class="text-lg font-bold text-purple-700">{{ inventoryStats.totalValue.toLocaleString('fa-IR') }}M</p>
            </div>
            <div class="bg-purple-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-purple-700" fill="currentColor" viewBox="0 0 20 20">
                <path d="M8.433 7.418c.155-.103.346-.196.567-.267v1.698a2.305 2.305 0 01-.567-.267C8.07 8.34 8 8.114 8 8c0-.114.07-.34.433-.582z"/>
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-13a1 1 0 10-2 0v.092a4.535 4.535 0 00-1.676.662C6.602 6.234 6 7.009 6 8c0 .99.602 1.765 1.324 2.246.48.32 1.054.545 1.676.662v1.941c-.391-.127-.68-.317-.843-.504a1 1 0 10-1.51 1.31c.562.649 1.413 1.076 2.353 1.253V15a1 1 0 102 0z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-lg shadow-sm p-6 border border-blue-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-blue-600">گردش روزانه</p>
              <p class="text-xl font-bold text-blue-700">{{ inventoryStats.dailyTurnover.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-blue-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-blue-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M4 2a1 1 0 011 1v2.101a7.002 7.002 0 0111.601 2.566 1 1 0 11-1.885.666A5.002 5.002 0 005.999 7H9a1 1 0 010 2H4a1 1 0 01-1-1V3a1 1 0 011-1zm.008 9.057a1 1 0 011.276.61A5.002 5.002 0 0014.001 13H11a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0v-2.101a7.002 7.002 0 01-11.601-2.566 1 1 0 01.61-1.276z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-yellow-50 to-yellow-100 rounded-lg shadow-sm p-6 border border-yellow-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-yellow-600">ورودی امروز</p>
              <p class="text-xl font-bold text-yellow-700">{{ inventoryStats.todayReceived.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-yellow-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-yellow-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM6.293 6.707a1 1 0 010-1.414l3-3a1 1 0 011.414 0l3 3a1 1 0 01-1.414 1.414L11 5.414V13a1 1 0 11-2 0V5.414L7.707 6.707a1 1 0 01-1.414 0z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-indigo-50 to-indigo-100 rounded-lg shadow-sm p-6 border border-indigo-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-indigo-600">خروجی امروز</p>
              <p class="text-xl font-bold text-indigo-700">{{ inventoryStats.todayShipped.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-indigo-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-indigo-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm3.293-7.707a1 1 0 011.414 0L9 10.586V3a1 1 0 112 0v7.586l1.293-1.293a1 1 0 111.414 1.414l-3 3a1 1 0 01-1.414 0l-3-3a1 1 0 010-1.414z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Content Sections -->
    <div class="px-6 space-y-6">
      <!-- Low Stock Alerts -->
      <div class="bg-white rounded-lg shadow-sm p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-red-700">هشدار کمبود موجودی</h3>
          <span class="bg-red-100 text-red-800 px-2 py-1 rounded-full text-xs">{{ lowStockItems.length }} مورد</span>
        </div>
        <div class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead class="bg-gray-50">
              <tr class="text-right">
                <th class="px-4 py-2">کد محصول</th>
                <th class="px-4 py-2">نام محصول</th>
                <th class="px-4 py-2">دسته‌بندی</th>
                <th class="px-4 py-2">موجودی فعلی</th>
                <th class="px-4 py-2">حداقل مورد نیاز</th>
                <th class="px-4 py-2">وضعیت</th>
                <th class="px-4 py-2">عملیات</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in lowStockItems" :key="item.id" class="border-t">
                <td class="px-4 py-3 font-mono text-gray-600">{{ item.sku }}</td>
                <td class="px-4 py-3 font-medium">{{ item.name }}</td>
                <td class="px-4 py-3">{{ item.category }}</td>
                <td class="px-4 py-3">{{ item.currentStock }}</td>
                <td class="px-4 py-3">{{ item.minRequired }}</td>
                <td class="px-4 py-3">
                  <span :class="getStockStatusClass(item.status)" class="px-2 py-1 rounded-full text-xs">
                    {{ item.status }}
                  </span>
                </td>
                <td class="px-4 py-3">
                  <div class="flex gap-2">
                    <button class="text-blue-600 hover:text-blue-800 text-sm">سفارش</button>
                    <button class="text-green-600 hover:text-green-800 text-sm">تاریخچه</button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Category Analysis & Top Products -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- Category Analysis -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <h3 class="text-lg font-semibold mb-4">تحلیل دسته‌بندی</h3>
          <div class="space-y-4">
            <div v-for="category in categoryAnalysis" :key="category.name" class="p-6 bg-gradient-to-r from-orange-50 to-orange-100 rounded-lg border border-orange-200">
              <div class="flex justify-between items-center mb-2">
                <span class="font-medium text-orange-800">{{ category.name }}</span>
                <span class="text-orange-700 text-sm">{{ category.percentage }}%</span>
              </div>
              <div class="grid grid-cols-3 gap-6 text-sm">
                <div>
                  <p class="text-gray-600">محصولات</p>
                  <p class="font-bold">{{ category.products }}</p>
                </div>
                <div>
                  <p class="text-gray-600">ارزش</p>
                  <p class="font-bold">{{ category.value }}M</p>
                </div>
                <div>
                  <p class="text-gray-600">گردش</p>
                  <p class="font-bold">{{ category.turnover }}%</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Top Selling Products -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <h3 class="text-lg font-semibold mb-4">پرفروش‌ترین محصولات</h3>
          <div class="space-y-3">
            <div v-for="(product, index) in topSellingProducts" :key="product.id" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <div class="flex items-center gap-3">
                <span class="flex items-center justify-center w-6 h-6 bg-orange-500 text-white rounded-full text-xs">{{ index + 1 }}</span>
                <div>
                  <p class="font-medium text-sm">{{ product.name }}</p>
                  <p class="text-xs text-gray-600">{{ product.sku }}</p>
                </div>
              </div>
              <div class="text-left">
                <p class="text-sm font-bold">{{ product.sold }} فروش</p>
                <p class="text-xs text-gray-600">{{ product.revenue }}M درآمد</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Inventory Movement History -->
      <div class="bg-white rounded-lg shadow-sm p-6">
        <h3 class="text-lg font-semibold mb-4">تاریخچه تحرکات موجودی</h3>
        <div class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead class="bg-gray-50">
              <tr class="text-right">
                <th class="px-4 py-2">تاریخ</th>
                <th class="px-4 py-2">نوع عملیات</th>
                <th class="px-4 py-2">محصول</th>
                <th class="px-4 py-2">تعداد</th>
                <th class="px-4 py-2">انبار</th>
                <th class="px-4 py-2">مسئول</th>
                <th class="px-4 py-2">توضیحات</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="movement in inventoryMovements" :key="movement.id" class="border-t">
                <td class="px-4 py-3">{{ movement.date }}</td>
                <td class="px-4 py-3">
                  <span :class="getMovementTypeClass(movement.type)" class="px-2 py-1 rounded-full text-xs">
                    {{ movement.type }}
                  </span>
                </td>
                <td class="px-4 py-3">{{ movement.product }}</td>
                <td class="px-4 py-3" :class="movement.quantity > 0 ? 'text-green-600' : 'text-red-600'">
                  {{ movement.quantity > 0 ? '+' : '' }}{{ movement.quantity }}
                </td>
                <td class="px-4 py-3">{{ movement.warehouse }}</td>
                <td class="px-4 py-3">{{ movement.responsible }}</td>
                <td class="px-4 py-3 text-gray-600">{{ movement.notes }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Warehouse Performance -->
      <div class="bg-white rounded-lg shadow-sm p-6">
        <h3 class="text-lg font-semibold mb-4">عملکرد انبارها</h3>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div v-for="warehouse in warehousePerformance" :key="warehouse.id" class="p-6 border border-gray-200 rounded-lg">
            <div class="flex items-center justify-between mb-3">
              <h4 class="font-medium">{{ warehouse.name }}</h4>
              <span :class="getPerformanceClass(warehouse.efficiency)" class="px-2 py-1 rounded-full text-xs">
                {{ warehouse.efficiency }}%
              </span>
            </div>
            <div class="space-y-2 text-sm">
              <div class="flex justify-between">
                <span class="text-gray-600">ظرفیت:</span>
                <span>{{ warehouse.capacity }}%</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">ورودی امروز:</span>
                <span class="text-green-600">+{{ warehouse.todayIn }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">خروجی امروز:</span>
                <span class="text-red-600">-{{ warehouse.todayOut }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">آخرین بازرسی:</span>
                <span>{{ warehouse.lastInspection }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Inventory Value Breakdown -->
      <div class="bg-white rounded-lg shadow-sm p-6">
        <h3 class="text-lg font-semibold mb-4">تفکیک ارزش موجودی</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          <div class="text-center p-6 bg-gradient-to-br from-blue-50 to-blue-100 rounded-lg border border-blue-200">
            <div class="text-2xl font-bold text-blue-700 mb-1">{{ valueBreakdown.rawMaterials }}M</div>
            <div class="text-sm text-blue-600">مواد اولیه</div>
          </div>
          <div class="text-center p-6 bg-gradient-to-br from-green-50 to-green-100 rounded-lg border border-green-200">
            <div class="text-2xl font-bold text-green-700 mb-1">{{ valueBreakdown.finishedGoods }}M</div>
            <div class="text-sm text-green-600">کالای آماده</div>
          </div>
          <div class="text-center p-6 bg-gradient-to-br from-yellow-50 to-yellow-100 rounded-lg border border-yellow-200">
            <div class="text-2xl font-bold text-yellow-700 mb-1">{{ valueBreakdown.workInProgress }}M</div>
            <div class="text-sm text-yellow-600">در حال تولید</div>
          </div>
          <div class="text-center p-6 bg-gradient-to-br from-purple-50 to-purple-100 rounded-lg border border-purple-200">
            <div class="text-2xl font-bold text-purple-700 mb-1">{{ valueBreakdown.deadStock }}M</div>
            <div class="text-sm text-purple-600">موجودی راکد</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
// declare const useAuth: () => { user: unknown; hasPermission: (perm: string) => boolean }
declare const useHead: (head: { title?: string; meta?: Record<string, unknown>[] }) => void
</script>

<script setup lang="ts">

// Set admin layout
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

// Sample data
const inventoryStats = {
  totalProducts: 1247,
  inStock: 892,
  lowStock: 245,
  outOfStock: 110,
  totalValue: 450,
  dailyTurnover: 156,
  todayReceived: 89,
  todayShipped: 134
}

const lowStockItems = [
  { id: 1, sku: 'IP13-128', name: 'گوشی آیفون ۱۳ ۱۲۸گیگ', category: 'الکترونیک', currentStock: 8, minRequired: 20, status: 'بحرانی' },
  { id: 2, sku: 'SW-APL', name: 'ساعت هوشمند اپل', category: 'الکترونیک', currentStock: 15, minRequired: 30, status: 'کم' },
  { id: 3, sku: 'LP-LEN', name: 'لپ‌تاپ لنوو ThinkPad', category: 'الکترونیک', currentStock: 5, minRequired: 15, status: 'بحرانی' },
  { id: 4, sku: 'HD-SS', name: 'هدفون سامسونگ', category: 'لوازم جانبی', currentStock: 12, minRequired: 25, status: 'کم' },
  { id: 5, sku: 'TB-10', name: 'تبلت ۱۰ اینچ', category: 'الکترونیک', currentStock: 3, minRequired: 12, status: 'بحرانی' }
]

const categoryAnalysis = [
  { name: 'الکترونیک', products: 485, value: 180, turnover: 85, percentage: 39 },
  { name: 'پوشاک', products: 324, value: 95, turnover: 72, percentage: 26 },
  { name: 'کتاب', products: 198, value: 45, turnover: 58, percentage: 16 },
  { name: 'خانه و آشپزخانه', products: 156, value: 78, turnover: 63, percentage: 12 },
  { name: 'ورزش', products: 84, value: 52, turnover: 69, percentage: 7 }
]

const topSellingProducts = [
  { id: 1, name: 'گوشی سامسونگ Galaxy S23', sku: 'GS23-256', sold: 145, revenue: 25.6 },
  { id: 2, name: 'لپ‌تاپ ایسوس ZenBook', sku: 'ZB-14', sold: 89, revenue: 18.9 },
  { id: 3, name: 'ساعت هوشمند شیائومی', sku: 'MI-WA', sold: 156, revenue: 12.3 },
  { id: 4, name: 'هدفون AirPods Pro', sku: 'APP-2', sold: 78, revenue: 9.8 },
  { id: 5, name: 'تبلت iPad Air', sku: 'IPA-64', sold: 45, revenue: 15.4 }
]

const inventoryMovements = [
  { id: 1, date: '۱۴۰۲/۱۲/۱۵', type: 'ورود', product: 'گوشی آیفون ۱۳', quantity: 50, warehouse: 'انبار مرکزی', responsible: 'احمد احمدی', notes: 'خرید از تامین‌کننده' },
  { id: 2, date: '۱۴۰۲/۱۲/۱۵', type: 'خروج', product: 'ساعت هوشمند', quantity: -25, warehouse: 'انبار مرکزی', responsible: 'مریم مرادی', notes: 'ارسال به مشتری' },
  { id: 3, date: '۱۴۰۲/۱۲/۱۴', type: 'انتقال', product: 'لپ‌تاپ لنوو', quantity: 10, warehouse: 'انبار اصفهان', responsible: 'علی علوی', notes: 'انتقال بین انبارها' },
  { id: 4, date: '۱۴۰۲/۱۲/۱۴', type: 'تعدیل', product: 'هدفون', quantity: -5, warehouse: 'انبار مشهد', responsible: 'زهرا زارعی', notes: 'کسری موجودی' },
  { id: 5, date: '۱۴۰۲/۱۲/۱۳', type: 'بازگشت', product: 'تبلت', quantity: 8, warehouse: 'انبار مرکزی', responsible: 'حسین حسینی', notes: 'بازگشت از مشتری' }
]

const warehousePerformance = [
  { id: 1, name: 'انبار مرکزی تهران', efficiency: 92, capacity: 78, todayIn: 156, todayOut: 134, lastInspection: '۱۴۰۲/۱۲/۱۰' },
  { id: 2, name: 'انبار شعبه اصفهان', efficiency: 85, capacity: 65, todayIn: 89, todayOut: 67, lastInspection: '۱۴۰۲/۱۲/۰۸' },
  { id: 3, name: 'انبار شعبه مشهد', efficiency: 78, capacity: 72, todayIn: 45, todayOut: 52, lastInspection: '۱۴۰۲/۱۲/۰۵' }
]

const valueBreakdown = {
  rawMaterials: 125,
  finishedGoods: 280,
  workInProgress: 35,
  deadStock: 10
}

const getStockStatusClass = (status: string) => {
  switch (status) {
    case 'بحرانی':
      return 'bg-red-100 text-red-800'
    case 'کم':
      return 'bg-yellow-100 text-yellow-800'
    case 'موجود':
      return 'bg-green-100 text-green-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getMovementTypeClass = (type: string) => {
  switch (type) {
    case 'ورود':
      return 'bg-green-100 text-green-800'
    case 'خروج':
      return 'bg-red-100 text-red-800'
    case 'انتقال':
      return 'bg-blue-100 text-blue-800'
    case 'تعدیل':
      return 'bg-yellow-100 text-yellow-800'
    case 'بازگشت':
      return 'bg-purple-100 text-purple-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getPerformanceClass = (efficiency: number) => {
  if (efficiency >= 90) return 'bg-green-100 text-green-800'
  if (efficiency >= 80) return 'bg-yellow-100 text-yellow-800'
  return 'bg-red-100 text-red-800'
}

// Meta and SEO
useHead({
  title: 'آمار موجودی انبار - پنل مدیریت',
  meta: [
    { name: 'description', content: 'تحلیل جامع موجودی، گردش کالا و عملکرد انبار' }
  ]
})
</script>

