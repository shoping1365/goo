<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-gradient-to-r from-emerald-50 to-green-50 rounded-lg p-6">
      <h2 class="text-2xl font-bold text-gray-900 mb-2">👥 مدیریت کاربران کیف پول</h2>
      <p class="text-gray-600">مدیریت کاربران و حساب‌های کیف پول</p>
    </div>

    <!-- آمار کاربران -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- کل کاربران -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">کل کاربران</h3>
          <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">ثبت شده</span>
        </div>
        <div class="text-3xl font-bold text-blue-600 mb-2">{{ userStats.totalUsers }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-green-500">+{{ userStats.dailyGrowth }}%</span>
          <span class="mx-2">از دیروز</span>
        </div>
      </div>

      <!-- کاربران فعال -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">کاربران فعال</h3>
          <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
        </div>
        <div class="text-3xl font-bold text-green-600 mb-2">{{ userStats.activeUsers }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-green-500">{{ userStats.activePercentage }}%</span>
          <span class="mx-2">از کل کاربران</span>
        </div>
      </div>

      <!-- کاربران جدید -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">کاربران جدید</h3>
          <span class="text-xs bg-yellow-100 text-yellow-800 rounded-full px-3 py-1">امروز</span>
        </div>
        <div class="text-3xl font-bold text-yellow-600 mb-2">{{ userStats.newUsers }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-yellow-500">+{{ userStats.newUserGrowth }}%</span>
          <span class="mx-2">از هفته قبل</span>
        </div>
      </div>

      <!-- کاربران مسدود -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">کاربران مسدود</h3>
          <span class="text-xs bg-red-100 text-red-800 rounded-full px-3 py-1">مسدود</span>
        </div>
        <div class="text-3xl font-bold text-red-600 mb-2">{{ userStats.blockedUsers }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-red-500">{{ userStats.blockedPercentage }}%</span>
          <span class="mx-2">از کل کاربران</span>
        </div>
      </div>
    </div>

    <!-- فیلترها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
        <!-- وضعیت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="">همه</option>
            <option value="active">فعال</option>
            <option value="inactive">غیرفعال</option>
            <option value="blocked">مسدود</option>
            <option value="pending">در انتظار تأیید</option>
          </select>
        </div>

        <!-- سطح کاربر -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">سطح کاربر</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="">همه</option>
            <option value="basic">پایه</option>
            <option value="premium">پریمیوم</option>
            <option value="vip">VIP</option>
          </select>
        </div>

        <!-- بازه موجودی -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه موجودی</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="">همه</option>
            <option value="low">کمتر از 1 میلیون</option>
            <option value="medium">1 تا 10 میلیون</option>
            <option value="high">بیش از 10 میلیون</option>
          </select>
        </div>

        <!-- جستجو -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
          <input
type="text" placeholder="نام، ایمیل یا شماره موبایل" 
                 class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
      </div>

      <div class="flex justify-between items-center">
        <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
          اعمال فیلتر
        </button>
        <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
          پاک کردن فیلترها
        </button>
      </div>
    </div>

    <!-- جدول کاربران -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-xl font-semibold text-gray-900">کاربران کیف پول</h3>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
            فعال کردن همه
          </button>
          <button class="px-4 py-2 bg-red-600 text-white text-sm rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500">
            مسدود کردن همه
          </button>
          <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            خروجی اکسل
          </button>
        </div>
      </div>
      
      <div class="overflow-x-auto">
        <table class="w-full text-sm text-right">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-gray-700 font-medium">
                <input type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
              </th>
              <th class="px-4 py-3 text-gray-700 font-medium">کاربر</th>
              <th class="px-4 py-3 text-gray-700 font-medium">موجودی</th>
              <th class="px-4 py-3 text-gray-700 font-medium">سطح</th>
              <th class="px-4 py-3 text-gray-700 font-medium">تاریخ عضویت</th>
              <th class="px-4 py-3 text-gray-700 font-medium">آخرین فعالیت</th>
              <th class="px-4 py-3 text-gray-700 font-medium">وضعیت</th>
              <th class="px-4 py-3 text-gray-700 font-medium">عملیات</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="user in users" :key="user.id" class="hover:bg-gray-50">
              <td class="px-4 py-3">
                <input type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
              </td>
              <td class="px-4 py-3">
                <div class="flex items-center">
                  <div class="w-10 h-10 bg-gray-200 rounded-full flex items-center justify-center text-sm font-medium text-gray-600">
                    {{ user.avatar }}
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ user.name }}</div>
                    <div class="text-xs text-gray-500">{{ user.email }}</div>
                    <div class="text-xs text-gray-500">{{ user.phone }}</div>
                  </div>
                </div>
              </td>
              <td class="px-4 py-3">
                <div class="text-sm font-medium text-gray-900">{{ formatCurrency(user.balance) }}</div>
                <div class="text-xs text-gray-500">{{ user.transactionCount }} تراکنش</div>
              </td>
              <td class="px-4 py-3">
                <span :class="getUserLevelClass(user.level)" class="px-2 py-1 text-xs rounded-full">
                  {{ user.level }}
                </span>
              </td>
              <td class="px-4 py-3 text-gray-600">{{ user.joinDate }}</td>
              <td class="px-4 py-3 text-gray-600">{{ user.lastActivity }}</td>
              <td class="px-4 py-3">
                <span :class="getStatusClass(user.status)" class="px-2 py-1 text-xs rounded-full">
                  {{ user.status }}
                </span>
              </td>
              <td class="px-4 py-3">
                <div class="flex space-x-2 space-x-reverse">
                  <button class="text-blue-600 hover:text-blue-800 text-sm">جزئیات</button>
                  <button
v-if="user.status === 'فعال'" 
                          class="text-red-600 hover:text-red-800 text-sm">مسدود</button>
                  <button
v-if="user.status === 'مسدود'" 
                          class="text-green-600 hover:text-green-800 text-sm">باز کردن</button>
                  <button class="text-purple-600 hover:text-purple-800 text-sm">تنظیمات</button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- صفحه‌بندی -->
      <div class="flex items-center justify-between mt-6">
        <div class="text-sm text-gray-700">
          نمایش {{ paginationInfo.start }} تا {{ paginationInfo.end }} از {{ paginationInfo.total }} کاربر
        </div>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">قبلی</button>
          <button class="px-3 py-1 text-sm bg-blue-600 text-white rounded-lg">1</button>
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">2</button>
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">3</button>
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">بعدی</button>
        </div>
      </div>
    </div>

    <!-- آمار تحلیلی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- نمودار کاربران بر اساس سطح -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">کاربران بر اساس سطح</h3>
        <div class="space-y-4">
          <div v-for="level in userLevelStats" :key="level.name" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-4 h-4 rounded-full mr-3" :style="{ backgroundColor: level.color }"></div>
              <span class="text-sm text-gray-700">{{ level.name }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-sm font-medium text-gray-900">{{ level.count }}</span>
              <span class="text-sm text-gray-500">({{ level.percentage }}%)</span>
            </div>
          </div>
        </div>
      </div>

      <!-- نمودار کاربران بر اساس وضعیت -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">کاربران بر اساس وضعیت</h3>
        <div class="space-y-4">
          <div v-for="status in userStatusStats" :key="status.name" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-4 h-4 rounded-full mr-3" :style="{ backgroundColor: status.color }"></div>
              <span class="text-sm text-gray-700">{{ status.name }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-sm font-medium text-gray-900">{{ status.count }}</span>
              <span class="text-sm text-gray-500">({{ status.percentage }}%)</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- عملیات دسته‌ای -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">عملیات دسته‌ای</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <button class="px-4 py-3 bg-green-100 text-green-700 rounded-lg hover:bg-green-200 focus:outline-none focus:ring-2 focus:ring-green-500">
          <div class="text-center">
            <div class="text-lg mb-1">✅</div>
            <div class="font-medium">فعال کردن</div>
            <div class="text-sm text-green-600">فعال کردن کاربران انتخاب شده</div>
          </div>
        </button>
        
        <button class="px-4 py-3 bg-red-100 text-red-700 rounded-lg hover:bg-red-200 focus:outline-none focus:ring-2 focus:ring-red-500">
          <div class="text-center">
            <div class="text-lg mb-1">🚫</div>
            <div class="font-medium">مسدود کردن</div>
            <div class="text-sm text-red-600">مسدود کردن کاربران انتخاب شده</div>
          </div>
        </button>
        
        <button class="px-4 py-3 bg-blue-100 text-blue-700 rounded-lg hover:bg-blue-200 focus:outline-none focus:ring-2 focus:ring-blue-500">
          <div class="text-center">
            <div class="text-lg mb-1">📧</div>
            <div class="font-medium">ارسال ایمیل</div>
            <div class="text-sm text-blue-600">ارسال ایمیل به کاربران انتخاب شده</div>
          </div>
        </button>
        
        <button class="px-4 py-3 bg-purple-100 text-purple-700 rounded-lg hover:bg-purple-200 focus:outline-none focus:ring-2 focus:ring-purple-500">
          <div class="text-center">
            <div class="text-lg mb-1">📊</div>
            <div class="font-medium">گزارش</div>
            <div class="text-sm text-purple-600">ایجاد گزارش از کاربران انتخاب شده</div>
          </div>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// آمار کاربران
const userStats = {
  totalUsers: 15420,
  dailyGrowth: 3.2,
  activeUsers: 12500,
  activePercentage: 81.1,
  newUsers: 45,
  newUserGrowth: 12.5,
  blockedUsers: 280,
  blockedPercentage: 1.8
}

// کاربران
const users = [
  {
    id: 1,
    name: 'علی احمدی',
    email: 'ali@example.com',
    phone: '09123456789',
    avatar: 'عا',
    balance: 5000000,
    transactionCount: 25,
    level: 'پریمیوم',
    joinDate: '1402/08/15',
    lastActivity: '2 ساعت پیش',
    status: 'فعال'
  },
  {
    id: 2,
    name: 'فاطمه محمدی',
    email: 'fateme@example.com',
    phone: '09187654321',
    avatar: 'فم',
    balance: 15000000,
    transactionCount: 42,
    level: 'VIP',
    joinDate: '1402/07/20',
    lastActivity: '1 روز پیش',
    status: 'فعال'
  },
  {
    id: 3,
    name: 'محمد رضایی',
    email: 'mohammad@example.com',
    phone: '09111111111',
    avatar: 'مر',
    balance: 800000,
    transactionCount: 8,
    level: 'پایه',
    joinDate: '1402/10/01',
    lastActivity: '3 روز پیش',
    status: 'غیرفعال'
  },
  {
    id: 4,
    name: 'زهرا کریمی',
    email: 'zahra@example.com',
    phone: '09222222222',
    avatar: 'زک',
    balance: 25000000,
    transactionCount: 67,
    level: 'VIP',
    joinDate: '1402/06/10',
    lastActivity: '5 دقیقه پیش',
    status: 'فعال'
  },
  {
    id: 5,
    name: 'حسین نوری',
    email: 'hossein@example.com',
    phone: '09333333333',
    avatar: 'حن',
    balance: 0,
    transactionCount: 0,
    level: 'پایه',
    joinDate: '1402/10/15',
    lastActivity: 'هرگز',
    status: 'مسدود'
  }
]

// اطلاعات صفحه‌بندی
const paginationInfo = {
  start: 1,
  end: 10,
  total: 15420
}

// آمار سطح کاربران
const userLevelStats = [
  { name: 'پایه', count: 9250, percentage: 60, color: '#6B7280' },
  { name: 'پریمیوم', count: 4620, percentage: 30, color: '#3B82F6' },
  { name: 'VIP', count: 1550, percentage: 10, color: '#F59E0B' }
]

// آمار وضعیت کاربران
const userStatusStats = [
  { name: 'فعال', count: 12500, percentage: 81.1, color: '#10B981' },
  { name: 'غیرفعال', count: 2640, percentage: 17.1, color: '#6B7280' },
  { name: 'مسدود', count: 280, percentage: 1.8, color: '#EF4444' }
]

// تابع فرمت کردن ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// تابع کلاس سطح کاربر
const getUserLevelClass = (level: string) => {
  switch (level) {
    case 'پایه':
      return 'bg-gray-100 text-gray-800'
    case 'پریمیوم':
      return 'bg-blue-100 text-blue-800'
    case 'VIP':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// تابع کلاس وضعیت
const getStatusClass = (status: string) => {
  switch (status) {
    case 'فعال':
      return 'bg-green-100 text-green-800'
    case 'غیرفعال':
      return 'bg-gray-100 text-gray-800'
    case 'مسدود':
      return 'bg-red-100 text-red-800'
    case 'در انتظار تأیید':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل بخش‌های زیر است:
  1. آمار کلی کاربران (کل، فعال، جدید، مسدود)
  2. فیلترهای پیشرفته برای جستجو
  3. جدول کاربران با امکان مدیریت
  4. نمودارهای تحلیلی بر اساس سطح و وضعیت
  5. عملیات دسته‌ای
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
--> 
