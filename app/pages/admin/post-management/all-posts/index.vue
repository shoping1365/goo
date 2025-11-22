<template>
  <div class="min-h-screen">

      
    
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 sticky top-0 z-30">
      <div class="px-6 py-4">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-2xl font-bold text-gray-900">همه نوشته‌ها</h1>
            <p class="text-sm text-gray-500 mt-1">مدیریت و ویرایش نوشته‌های موجود</p>
          </div>
          
          <div class="flex items-center gap-3">
            <!-- دکمه AI -->
            <TemplateButton
              v-if="hasPermission('post_add.create')"
              bg-gradient="bg-gradient-to-r from-purple-400 to-indigo-600"
              text-color="text-white"
              hover-class="hover:from-purple-500 hover:to-indigo-700 hover:shadow-lg hover:scale-105"
              focus-class="focus:ring-2 focus:ring-offset-2 focus:ring-purple-500"
              size="medium"
              @click="openAIChatModal"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
              </svg>
              AI نویسنده
            </TemplateButton>
            
            <!-- افزودن نوشته جدید -->
            <NuxtLink 
              v-if="hasPermission('post_add.create')"
              to="/admin/post-management/add-post"
            >
              <TemplateButton
                bg-gradient="bg-gradient-to-r from-green-400 to-emerald-600"
                text-color="text-white"
                hover-class="hover:from-green-500 hover:to-emerald-700 hover:shadow-lg hover:scale-105"
                focus-class="focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
                size="medium"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
                </svg>
                افزودن نوشته جدید
              </TemplateButton>
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار -->
    <div class="mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
        <TemplateCard
          title="کل نوشته‌ها"
          :value="totalPosts"
          variant="blue"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
          </template>
        </TemplateCard>
        <TemplateCard
          title="منتشر شده"
          :value="publishedPosts"
          variant="green"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
          </template>
        </TemplateCard>
        <TemplateCard
          title="پیش‌نویس"
          :value="draftPosts"
          variant="orange"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
            </svg>
          </template>
        </TemplateCard>
        <TemplateCard
          title="زمان‌بندی شده"
          :value="scheduledPosts"
          variant="purple"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
            </svg>
          </template>
        </TemplateCard>
      </div>
    </div>

    <!-- فیلترها و عملیات گروهی در یک ردیف -->
    <div class="bg-white border-b border-gray-200 px-6 py-4 mb-6 rounded-lg shadow-sm">
      <div class="flex flex-col sm:flex-row sm:items-center gap-6">
        <!-- جستجو در نوشته‌ها -->
        <div class="flex-1 order-1 sm:order-none">
          <div class="relative">
            <input
              v-model="searchQuery"
              type="text"
              placeholder="جستجو در نوشته‌ها..."
              class="w-full sm:w-80 pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
          </div>
        </div>
        <!-- فیلترها و عملیات گروهی -->
        <div class="flex gap-6 order-2 sm:order-none items-center">
          <!-- فیلتر وضعیت -->
          <select v-model="statusFilter" class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 text-sm">
            <option value="">همه وضعیت‌ها</option>
            <option value="draft">پیش‌نویس</option>
            <option value="published">منتشر شده</option>
            <option value="scheduled">زمان‌بندی شده</option>
          </select>
          <!-- فیلتر دسته‌بندی -->
          <select v-model="categoryFilter" class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 text-sm">
            <option value="">همه دسته‌بندی‌ها</option>
            <option v-for="category in categories" :key="category.id" :value="category.id">
              {{ category.name }}
            </option>
          </select>
          <!-- عملیات گروهی -->
          <div v-if="selectedPosts.length > 0" class="flex items-center gap-2">
            <div class="relative">
              <TemplateButton 
                type="button" 
                bg-gradient="bg-white"
                text-color="text-gray-700"
                border-color="border border-gray-300"
                focus-class="focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                size="small"
                @click="showBulkDropdown = !showBulkDropdown"
              >
                <span>{{ bulkActionLabel }}</span>
                <svg class="w-4 h-4 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                </svg>
              </TemplateButton>
              <div v-if="showBulkDropdown" class="absolute right-0 left-0 z-40 mt-1 bg-white border border-gray-200 rounded-lg shadow-lg py-1 animate-fade-in-up">
                <button
v-for="item in bulkActions" :key="item.value" type="button" class="w-full text-right px-4 py-2 text-sm text-gray-700 hover:bg-blue-50 transition-colors"
                  @click="selectBulkAction(item)"
                >
                  {{ item.label }}
                </button>
              </div>
            </div>
            <TemplateButton 
              :disabled="!bulkAction" 
              bg-gradient="bg-gradient-to-r from-blue-500 to-blue-600" 
              text-color="text-white"
              hover-class="hover:from-blue-600 hover:to-blue-700 hover:shadow-lg hover:scale-105"
              focus-class="focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              size="small"
              :class="{'opacity-50': !bulkAction}"
              @click="handleBulkAction"
            >
              اجرا
            </TemplateButton>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول نوشته‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                <input 
                  type="checkbox" 
                  :checked="selectedPosts.length === filteredPosts.length && filteredPosts.length > 0"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                  @change="toggleSelectAll"
                >
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تصویر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عنوان نوشته</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نویسنده</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دسته‌بندی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="post in paginatedPosts" :key="post.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <input 
                  v-model="selectedPosts" 
                  type="checkbox"
                  :value="post.id"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                >
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="">
                  <img v-if="post.featured_image" :src="post.featured_image" alt="تصویر شاخص" class="h-10 w-10 rounded-lg object-cover">
                  <div v-else class="h-10 w-10 rounded-lg bg-gray-200 flex items-center justify-center">
                    <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                    </svg>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <NuxtLink
                  v-if="post.slug && getCategorySlug(post.category_id)"
                  :to="`/blog/${getCategorySlug(post.category_id)}/${post.slug}`"
                  class="text-sm font-medium text-blue-700 hover:underline cursor-pointer"
                  target="_blank"
                >
                  {{ post.title }}
                </NuxtLink>
                <NuxtLink
                  v-else-if="post.slug"
                  :to="`/blog/${post.slug}`"
                  class="text-sm font-medium text-blue-700 hover:underline cursor-pointer"
                  target="_blank"
                >
                  {{ post.title }}
                </NuxtLink>
                <NuxtLink
                  v-else
                  :to="`/admin/post-management/edit-post/${post.id}`"
                  class="text-sm font-medium text-orange-600 hover:underline cursor-pointer"
                  title="نوشته slug ندارد - برای ویرایش کلیک کنید"
                >
                  {{ post.title }}
                </NuxtLink>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ getAuthorName(post.author) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ getCategoryName(post.category_id) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusBadgeClass(post.status)" class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full">
                  {{ getStatusText(post.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(post.created_at) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex items-center gap-2">
                  <TemplateButton 
                    v-if="hasPermission('post_add.update')"
                    text-color="text-blue-600"
                    hover-class="hover:text-blue-900"
                    size="small"
                    :border-color="'border-0'"
                    :bg-gradient="'bg-transparent'"
                    title="ویرایش"
                    @click="editPost(post)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                    </svg>
                  </TemplateButton>
                  <TemplateButton 
                    v-if="hasPermission('post_add.delete')"
                    text-color="text-red-600"
                    hover-class="hover:text-red-900"
                    size="small"
                    :border-color="'border-0'"
                    :bg-gradient="'bg-transparent'"
                    title="حذف"
                    @click="deletePost(post)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                  </TemplateButton>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      
      <!-- Pagination پایین جدول -->
      <div class="flex flex-col items-end gap-6 mt-6">
        <Pagination
          :current-page="currentPage"
          :total-pages="totalPages"
          :total="filteredPosts.length"
          @page-changed="handlePageChange"
        />
      </div>
    </div>

    <!-- AI Chat Modal -->
    <div v-if="showAIChatModal" class="fixed inset-0 z-50 flex items-center justify-center">
      <!-- پس‌زمینه شفاف -->
      <div class="absolute inset-0 bg-black bg-opacity-30"></div>
      
      <!-- Modal Content -->
      <div class="relative bg-white rounded-lg shadow-xl w-full max-w-4xl max-h-[90vh] flex flex-col">
        <!-- Header -->
        <div class="flex items-center justify-between p-6 border-b border-gray-200">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-gradient-to-r from-purple-400 to-indigo-600 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
              </svg>
            </div>
            <div>
              <h2 class="text-xl font-bold text-gray-900">AI نویسنده</h2>
              <p class="text-sm text-gray-500">با AI صحبت کنید و مقاله تولید کنید</p>
            </div>
          </div>
          
          <div class="flex items-center gap-2">
            <!-- دکمه شروع چت جدید -->
            <button 
              class="p-2 hover:bg-gray-100 rounded-lg transition-colors"
              title="شروع چت جدید"
              @click="clearChat"
            >
              <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
              </svg>
            </button>
            
            <!-- دکمه راهنما -->
            <button 
              class="p-2 hover:bg-gray-100 rounded-lg transition-colors"
              title="راهنمای AI Writer"
              @click="showGuideModal = true"
            >
              <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </button>
            
            <!-- دکمه سابقه چت -->
            <button 
              class="p-2 hover:bg-gray-100 rounded-lg transition-colors"
              title="سابقه چت"
              @click="showChatHistoryModal"
            >
              <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </button>
            
            <!-- دکمه بستن -->
            <button 
              class="p-2 hover:bg-gray-100 rounded-lg transition-colors"
              @click="showAIChatModal = false"
            >
              <svg class="w-6 h-6 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>
          <!-- دکمه تست اتصال OpenAI -->
          <div class="flex items-center gap-2">
            <button class="px-3 py-1 rounded bg-blue-100 text-blue-700 text-sm hover:bg-blue-200 transition" @click="testOpenAIConnection">
              تست اتصال OpenAI
            </button>
            <span v-if="testResult" :class="testResult.success ? 'text-green-600' : 'text-red-600'" class="text-sm">
              {{ testResult.message }}
            </span>
          </div>
        </div>

                  <!-- Content -->
          <div class="flex-1 flex flex-col min-h-0">
            <!-- تنظیمات AI -->
            <div class="p-6 border-b border-gray-200 bg-gray-50 flex-shrink-0">
            <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
              <!-- انتخاب مدل -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">مدل AI</label>
                <select 
                  v-model="aiSettings.model"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 text-sm"
                >
                  <option v-for="model in availableModels" :key="model.id" :value="model.id">
                    {{ model.name }}
                  </option>
                </select>
              </div>
              
              <!-- تعداد کلمات -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">تعداد کلمات</label>
                <select 
                  v-model="aiSettings.wordCount"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 text-sm"
                >
                  <option value="300">300 کلمه</option>
                  <option value="500">500 کلمه</option>
                  <option value="800">800 کلمه</option>
                  <option value="1000">1000 کلمه</option>
                  <option value="1500">1500 کلمه</option>
                  <option value="2000">2000 کلمه</option>
                </select>
              </div>
              
              <!-- دما (Temperature) -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">دما (خلاقیت)</label>
                <select 
                  v-model="aiSettings.temperature"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 text-sm"
                >
                  <option value="0.1">0.1 - دقیق و قابل پیش‌بینی</option>
                  <option value="0.3">0.3 - متعادل</option>
                  <option value="0.5">0.5 - متعادل</option>
                  <option value="0.7">0.7 - خلاقانه</option>
                  <option value="0.9">0.9 - بسیار خلاقانه</option>
                  <option value="1.0">1.0 - کاملاً خلاقانه</option>
                </select>
              </div>
              
              <!-- سبک نوشتار -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">سبک نوشتار</label>
                <select 
                  v-model="aiSettings.writingStyle"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 text-sm"
                >
                  <option value="professional">حرفه‌ای</option>
                  <option value="casual">صمیمی</option>
                  <option value="academic">علمی</option>
                  <option value="creative">خلاقانه</option>
                  <option value="technical">تکنیکی</option>
                </select>
              </div>
            </div>
          </div>

          <!-- Chat Area -->
          <div class="flex-1 flex flex-col min-h-0">
            <!-- Messages -->
            <div class="flex-1 overflow-y-auto p-6 space-y-4 chat-messages-container relative max-h-[calc(100vh-400px)] min-h-[300px]">
              <!-- دکمه اسکرول به پایین -->
              <button 
                v-if="showScrollButton"
                class="absolute bottom-4 right-4 z-10 w-10 h-10 bg-purple-600 text-white rounded-full shadow-lg hover:bg-purple-700 transition-all duration-200 flex items-center justify-center"
                title="اسکرول به پایین"
                @click="scrollToBottom"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3"></path>
                </svg>
              </button>
              <!-- پیام خوشامدگویی -->
              <div v-if="messages.length === 0" class="text-center py-8">
                <div class="w-16 h-16 bg-gradient-to-r from-purple-400 to-indigo-600 rounded-full flex items-center justify-center mx-auto mb-4">
                  <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
                    </svg>
                  </div>
                  <h3 class="text-lg font-semibold text-gray-900 mb-2">به AI نویسنده خوش آمدید!</h3>
                  <p class="text-gray-600 mb-4">می‌توانید از من بخواهید:</p>
                                        <div class="grid grid-cols-1 md:grid-cols-2 gap-3 max-w-2xl mx-auto">
                          <div class="p-3 bg-purple-50 rounded-lg border border-purple-200">
                            <p class="text-sm text-purple-800">"در مورد عشق 5 تا عنوان قشنگ برای مقاله بده"</p>
                          </div>
                          <div class="p-3 bg-purple-50 rounded-lg border border-purple-200">
                            <p class="text-sm text-purple-800">"برای این عنوان بنویس"</p>
                          </div>
                          <div class="p-3 bg-purple-50 rounded-lg border border-purple-200">
                            <p class="text-sm text-purple-800">"این قسمت را ویرایش کن"</p>
                          </div>
                          <div class="p-3 bg-purple-50 rounded-lg border border-purple-200">
                            <p class="text-sm text-purple-800">"همین خوبه"</p>
                          </div>
                          <div class="p-3 bg-purple-50 rounded-lg border border-purple-200">
                            <p class="text-sm text-purple-800">"مقاله‌ای درباره هوش مصنوعی بنویس"</p>
                          </div>
                          <div class="p-3 bg-purple-50 rounded-lg border border-purple-200">
                            <p class="text-sm text-purple-800">"محتوا برای وبلاگ تولید کن"</p>
                          </div>
                        </div>
                </div>

                                    <!-- پیام‌های چت -->
                      <div v-for="(message, index) in messages" :key="index" class="flex" :class="message.role === 'user' ? 'justify-end' : 'justify-start'">
                        <div 
                          class="max-w-[80%] p-6 rounded-lg relative"
                          :class="message.role === 'user' 
                            ? 'bg-purple-600 text-white' 
                            : 'bg-gray-100 text-gray-900'"
                        >
                          <!-- تیک سبز برای پیام‌های کاربر که درخواست مقاله دارند -->
                          <div v-if="message.role === 'user' && message.isArticleRequest" class="absolute -top-2 -right-2 w-6 h-6 bg-green-500 rounded-full flex items-center justify-center shadow-lg">
                            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                            </svg>
                          </div>
                          
                          <div class="flex items-start gap-2">
                            <div v-if="message.role === 'assistant'" class="w-6 h-6 bg-gradient-to-r from-purple-400 to-indigo-600 rounded-full flex items-center justify-center flex-shrink-0">
                              <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
                              </svg>
                            </div>
                            <div class="flex-1">
                              <p class="text-sm whitespace-pre-wrap">{{ message.content }}</p>
                              
                                                           <!-- پیشنهادات عنوان -->
                             <div v-if="message.role === 'assistant' && 'suggestions' in message && Array.isArray(message.suggestions) && message.suggestions.length > 0" class="mt-3 space-y-2">
                               <p class="text-xs text-gray-600 mb-2">برای انتخاب عنوان، شماره آن را بنویسید:</p>
                               <div class="grid grid-cols-1 gap-2">
                                 <div 
                                   v-for="(suggestion, idx) in message.suggestions" 
                                   :key="idx"
                                   class="p-2 bg-blue-50 border border-blue-200 rounded text-xs"
                                 >
                                   <span class="font-medium text-blue-700">{{ idx + 1 }}.</span> {{ suggestion }}
                                 </div>
                               </div>
                             </div>
                             
                             <!-- دکمه‌های تأیید/ویرایش برای مقاله -->
                             <div v-if="message.role === 'assistant' && 'requiresConfirmation' in message && message.requiresConfirmation" class="mt-3 flex gap-2">
                                <button 
                                  class="inline-flex items-center gap-2 px-3 py-1 bg-green-600 text-white text-xs rounded-lg hover:bg-green-700 transition-colors"
                                  @click="confirmArticle(message)"
                                >
                                  <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                                  </svg>
                                  تأیید و ذخیره
                                </button>
                                <button 
                                  class="inline-flex items-center gap-2 px-3 py-1 bg-blue-600 text-white text-xs rounded-lg hover:bg-blue-700 transition-colors"
                                  @click="editArticle(message)"
                                >
                                  <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                                  </svg>
                                  ویرایش
                                </button>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>

              <!-- Loading indicator -->
              <div v-if="isGenerating" class="flex justify-start">
                <div class="bg-gray-100 text-gray-900 max-w-[80%] p-6 rounded-lg">
                  <div class="flex items-center gap-2">
                    <div class="w-6 h-6 bg-gradient-to-r from-purple-400 to-indigo-600 rounded-full flex items-center justify-center">
                      <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
                      </svg>
                    </div>
                    <div class="flex items-center gap-1">
                      <div class="w-2 h-2 bg-purple-600 rounded-full animate-bounce"></div>
                                  <div class="w-2 h-2 bg-purple-600 rounded-full animate-bounce" style="animation-delay: 0s"></div>
            <div class="w-2 h-2 bg-purple-600 rounded-full animate-bounce" style="animation-delay: 0s"></div>
                      <span class="text-sm text-gray-600 mr-2">در حال تولید...</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Input Area -->
            <div class="p-6 border-t border-gray-200 flex-shrink-0">
              <div class="flex gap-3">
                <div class="flex-1">
                  <textarea
                    v-model="userInput"
                    placeholder="پیام خود را بنویسید..."
                    class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 resize-none"
                    rows="3"
                    dir="rtl"
                    @keydown.enter.prevent="sendMessage"
                    @input="scrollToBottom"
                  ></textarea>
                </div>
                <div class="flex flex-col gap-2">
                  <button 
                    :disabled="!userInput.trim() || isGenerating"
                    class="px-6 py-3 bg-gradient-to-r from-purple-600 to-indigo-600 text-white rounded-lg hover:from-purple-700 hover:to-indigo-700 focus:ring-2 focus:ring-offset-2 focus:ring-purple-500 disabled:opacity-50 disabled:cursor-not-allowed transition-all"
                    @click="sendMessage"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"></path>
                    </svg>
                  </button>
                  
                  <button 
                    class="px-6 py-3 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 transition-colors"
                    title="پاک کردن چت"
                    @click="clearChat"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Chat History Modal -->
    <div v-if="showChatHistory" class="fixed inset-0 z-50 flex items-center justify-center">
      <!-- پس‌زمینه شفاف -->
      <div class="absolute inset-0 bg-black bg-opacity-30"></div>
      
      <!-- Modal Content -->
      <div class="relative bg-white rounded-lg shadow-xl w-full max-w-2xl max-h-[80vh] flex flex-col">
        <!-- Header -->
        <div class="flex items-center justify-between p-6 border-b border-gray-200">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-gradient-to-r from-blue-400 to-indigo-600 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div>
              <h2 class="text-xl font-bold text-gray-900">سابقه چت‌ها</h2>
              <p class="text-sm text-gray-500">جلسات چت قبلی شما</p>
            </div>
          </div>
          
          <div class="flex items-center gap-2">
            <button 
              class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors text-sm"
              @click="startNewChat"
            >
              چت جدید
            </button>
            <button 
              class="p-2 hover:bg-gray-100 rounded-lg transition-colors"
              @click="showChatHistory = false"
            >
              <svg class="w-6 h-6 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>
        </div>

        <!-- Content -->
        <div class="flex-1 overflow-y-auto p-6">
          <div v-if="chatHistory.length === 0" class="text-center py-8">
            <div class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4">
              <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">سابقه‌ای یافت نشد</h3>
            <p class="text-gray-600">هنوز هیچ جلسه چتی ندارید</p>
          </div>

          <div v-else class="space-y-3">
            <div 
              v-for="session in chatHistory" 
              :key="session.id"
              class="p-6 border border-gray-200 rounded-lg hover:border-blue-300 hover:bg-blue-50 transition-all cursor-pointer"
              @click="loadSession(session)"
            >
              <div class="flex items-start justify-between">
                <div class="flex-1">
                  <h4 class="font-medium text-gray-900 mb-1">{{ session.title }}</h4>
                  <div class="flex items-center gap-6 text-sm text-gray-500">
                    <span>{{ session.model }}</span>
                    <span>{{ (session.message_count as number) || 0 }} پیام</span>
                    <span>{{ formatDate((session.updated_at as string) || '') }}</span>
                  </div>
                </div>
                <button 
                  class="p-1 hover:bg-red-100 rounded text-red-600 hover:text-red-700 transition-colors"
                  title="حذف"
                  @click.stop="deleteSession(String(session.id))"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Guide Modal -->
    <div v-if="showGuideModal" class="fixed inset-0 z-50 flex items-center justify-center">
      <!-- پس‌زمینه شفاف -->
      <div class="absolute inset-0 bg-black bg-opacity-30" @click="showGuideModal = false"></div>
      
      <!-- Modal Content -->
      <div class="relative bg-white rounded-lg shadow-xl w-full max-w-4xl max-h-[90vh] flex flex-col" @click.stop>
        <!-- هدر مودال -->
        <div class="flex items-center justify-between p-6 border-b border-gray-200">
          <h3 class="text-xl font-bold text-gray-900">🎯 راهنمای استفاده از AI Writer</h3>
          <button 
            class="p-2 hover:bg-gray-100 rounded-lg transition-colors" 
            @click="showGuideModal = false"
          >
            <svg class="w-6 h-6 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <!-- محتوای مودال -->
        <div class="flex-1 overflow-y-auto p-6">
            <div class="space-y-6">
                                  <!-- بخش 1: نحوه کار -->
                      <div class="border-b border-gray-100 pb-6">
                        <h4 class="text-lg font-semibold text-gray-700 mb-3 flex items-center">
                          📝 نحوه کار سیستم تعاملی
                        </h4>
                        <div class="text-sm text-gray-600">
                          <p class="mb-3">سیستم AI Writer به صورت تعاملی و مرحله‌ای کار می‌کند:</p>
                          <ol class="list-decimal list-inside space-y-2">
                            <li><strong>درخواست عنوان:</strong> "در مورد عشق 5 تا عنوان قشنگ برای مقاله بده"</li>
                            <li><strong>انتخاب عنوان:</strong> شماره عنوان مورد نظر را بنویسید (مثل "1" یا "2")</li>
                            <li><strong>تولید مقاله:</strong> AI مقاله کامل تولید و نمایش می‌دهد</li>
                            <li><strong>بررسی و ویرایش:</strong> "این قسمت را ویرایش کن" یا "این جا را تغییر بده"</li>
                            <li><strong>تأیید نهایی:</strong> "همین خوبه" یا "تأیید" برای ذخیره</li>
                          </ol>
                        </div>
                      </div>

                                  <!-- بخش 2: انواع پیام‌ها -->
                      <div class="border-b border-gray-100 pb-6">
                        <h4 class="text-lg font-semibold text-gray-700 mb-3 flex items-center">
                          💬 انواع پیام‌های قابل تشخیص
                        </h4>
                        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                          <!-- درخواست عنوان -->
                          <div class="bg-gray-50 p-6 rounded-lg">
                            <h5 class="font-medium text-gray-700 mb-2">📝 درخواست عنوان</h5>
                            <ul class="space-y-1 text-sm text-gray-600">
                              <li class="bg-white px-2 py-1 rounded text-xs">"در مورد عشق 5 تا عنوان بده"</li>
                              <li class="bg-white px-2 py-1 rounded text-xs">"پیشنهاد عنوان برای مقاله"</li>
                              <li class="bg-white px-2 py-1 rounded text-xs">"چند تا تیتر قشنگ بده"</li>
                            </ul>
                          </div>
          
                          <!-- انتخاب عنوان -->
                          <div class="bg-gray-50 p-6 rounded-lg">
                            <h5 class="font-medium text-gray-700 mb-2">🎯 انتخاب عنوان</h5>
                            <ul class="space-y-1 text-sm text-gray-600">
                              <li class="bg-white px-2 py-1 rounded text-xs">"1" (شماره عنوان)</li>
                              <li class="bg-white px-2 py-1 rounded text-xs">"این" یا "همین"</li>
                              <li class="bg-white px-2 py-1 rounded text-xs">"برای این عنوان بنویس"</li>
                            </ul>
                          </div>
          
                          <!-- درخواست مقاله -->
                          <div class="bg-gray-50 p-6 rounded-lg">
                            <h5 class="font-medium text-gray-700 mb-2">📄 درخواست مقاله</h5>
                            <ul class="space-y-1 text-sm text-gray-600">
                              <li class="bg-white px-2 py-1 rounded text-xs">"برای این عنوان بنویس"</li>
                              <li class="bg-white px-2 py-1 rounded text-xs">"مقاله‌ای درباره هوش مصنوعی بنویس"</li>
                              <li class="bg-white px-2 py-1 rounded text-xs">"محتوا تولید کن"</li>
                            </ul>
                          </div>
          
                          <!-- ویرایش مقاله -->
                          <div class="bg-gray-50 p-6 rounded-lg">
                            <h5 class="font-medium text-gray-700 mb-2">✏️ ویرایش مقاله</h5>
                            <ul class="space-y-1 text-sm text-gray-600">
                              <li class="bg-white px-2 py-1 rounded text-xs">"این قسمت را ویرایش کن"</li>
                              <li class="bg-white px-2 py-1 rounded text-xs">"این جا را تغییر بده"</li>
                              <li class="bg-white px-2 py-1 rounded text-xs">"اصلاح کن"</li>
                            </ul>
                          </div>
          
                          <!-- تأیید نهایی -->
                          <div class="bg-gray-50 p-6 rounded-lg">
                            <h5 class="font-medium text-gray-700 mb-2">✅ تأیید نهایی</h5>
                            <ul class="space-y-1 text-sm text-gray-600">
                              <li class="bg-white px-2 py-1 rounded text-xs">"همین خوبه"</li>
                              <li class="bg-white px-2 py-1 rounded text-xs">"تأیید"</li>
                              <li class="bg-white px-2 py-1 rounded text-xs">"عالی شده"</li>
                            </ul>
                          </div>
                        </div>
                      </div>

              <!-- بخش 3: فیلدهای تولید شده -->
              <div class="border-b border-gray-100 pb-6">
                <h4 class="text-lg font-semibold text-gray-700 mb-3 flex items-center">
                  📋 فیلدهای تولید شده
                </h4>
                <div class="text-sm text-gray-600">
                  <p class="mb-3">وقتی مقاله تولید می‌شود، این فیلدها به صورت خودکار پر می‌شوند:</p>
                <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
                  <div class="bg-gray-50 p-3 rounded-lg">
                    <span class="font-medium text-gray-700 text-sm">عنوان مقاله</span>
                    <span class="block text-xs text-gray-500 mt-1">عنوان جذاب و SEO-friendly</span>
                  </div>
                  <div class="bg-gray-50 p-3 rounded-lg">
                    <span class="font-medium text-gray-700 text-sm">محتوای کامل</span>
                    <span class="block text-xs text-gray-500 mt-1">متن کامل مقاله</span>
                  </div>
                  <div class="bg-gray-50 p-3 rounded-lg">
                    <span class="font-medium text-gray-700 text-sm">خلاصه</span>
                    <span class="block text-xs text-gray-500 mt-1">خلاصه کوتاه مقاله</span>
                  </div>
                  <div class="bg-gray-50 p-3 rounded-lg">
                    <span class="font-medium text-gray-700 text-sm">توضیحات متا</span>
                    <span class="block text-xs text-gray-500 mt-1">برای SEO</span>
                  </div>

                  <div class="bg-gray-50 p-3 rounded-lg">
                    <span class="font-medium text-gray-700 text-sm">URL</span>
                    <span class="block text-xs text-gray-500 mt-1">آدرس مناسب</span>
                  </div>
                  <div class="bg-gray-50 p-3 rounded-lg">
                    <span class="font-medium text-gray-700 text-sm">تگ‌ها</span>
                    <span class="block text-xs text-gray-500 mt-1">برچسب‌های مرتبط</span>
                  </div>
                  <div class="bg-gray-50 p-3 rounded-lg">
                    <span class="font-medium text-gray-700 text-sm">وضعیت</span>
                    <span class="block text-xs text-gray-500 mt-1">پیش‌نویس (draft)</span>
                  </div>
                </div>
              </div>
            </div>

            <!-- بخش 4: نکات مهم -->
            <div>
              <h4 class="text-lg font-semibold text-gray-700 mb-3 flex items-center">
                ⚠️ نکات مهم
              </h4>
              <div class="text-sm text-gray-600">
                <ul class="space-y-2">
                  <li class="flex items-start">✅ مقالات به صورت خودکار در جدول <code class="bg-gray-100 px-1 py-0.5 rounded text-xs font-mono">posts</code> ذخیره می‌شوند</li>
                  <li class="flex items-start">✅ وضعیت پیش‌فرض: <code class="bg-gray-100 px-1 py-0.5 rounded text-xs font-mono">draft</code> (پیش‌نویس)</li>
                  <li class="flex items-start">✅ تمام فیلدهای SEO به صورت خودکار پر می‌شوند</li>
                  <li class="flex items-start">✅ ID مقاله در پاسخ برمی‌گردد</li>
                  <li class="flex items-start">✅ نیازی به ایجاد جدول جدید نیست</li>
                </ul>
              </div>
            </div>
          </div>
        </div>

        <!-- فوتر مودال -->
        <div class="flex justify-end p-6 border-t border-gray-200">
          <button 
            class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" 
            @click="showGuideModal = false"
          >
            بستن راهنما
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useHead: (head: { title?: string }) => void
declare const navigateTo: (to: string) => Promise<void>
</script>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import Pagination from '~/components/admin/common/Pagination.vue'
import { useAuth } from '@/composables/useAuth'

// Import useAuth for permission checking
const { user: _user, hasPermission } = useAuth()

// const { showError } = useErrorHandler()

// تعریف interface برای نوشته
interface Post {
  id: number
  title: string
  excerpt: string
  content: string
  status: 'draft' | 'published' | 'scheduled'
  category_id: number
  featured_image: string
  author: string | { name: string }
  created_at: string
  updated_at: string
  slug: string
}

// تعریف interface برای response types
interface ChatSessionResponse {
  sessions?: ChatSession[]
}

interface ChatSession {
  id: string | number
  title?: string
  model?: string
  messages?: ChatMessage[]
  message_count?: number
  updated_at?: string
  session?: {
    messages?: ChatMessage[]
    model?: string
  }
  [key: string]: unknown
}

interface ChatMessage {
  role: string
  content: string
  generatedContent?: unknown
  isArticleRequest?: boolean
  suggestions?: string[]
  requiresConfirmation?: boolean
}

interface CreateSessionResponse {
  id: string | number
}

interface AIChatResponse {
  message?: string
  article?: unknown
  type?: string
  post_id?: string | number
}

interface APISettingsResponse {
  openai?: {
    api_key?: string
    enabled?: boolean
    available_models?: Array<{ id: string; name?: string; [key: string]: unknown }>
    [key: string]: unknown
  }
  data?: {
    openai?: {
      api_key?: string
      enabled?: boolean
      available_models?: Array<{ id: string; name?: string; [key: string]: unknown }>
      [key: string]: unknown
    }
  }
}

interface TestOpenAIResponse {
  status: string
  message?: string
}

interface ErrorResponse {
  status?: number
  statusCode?: number
  data?: {
    message?: string
  }
  message?: string
}

// تعریف interface برای دسته‌بندی
interface Category {
  id: number
  name: string
  slug: string
}

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

useHead({
  title: 'همه نوشته‌ها - پنل مدیریت'
})

// Reactive data
const posts = ref<Post[]>([])
const categories = ref<Category[]>([])
const searchQuery = ref('')
const statusFilter = ref('')
const categoryFilter = ref('')
const sortBy = ref('created_at')
const perPage = ref(25)
const currentPage = ref(1)
const selectedPosts = ref<number[]>([])
const bulkActions = [
  { value: 'publish', label: 'انتشار' },
  { value: 'draft', label: 'پیش‌نویس' },
  { value: 'delete', label: 'حذف' },
  { value: 'index', label: 'تنظیم به index' },
  { value: 'nofollow', label: 'تنظیم به nofollow' },
]
const bulkAction = ref('')
const bulkActionLabel = computed(() => {
  const found = bulkActions.find(a => a.value === bulkAction.value)
  return found ? found.label : 'انتخاب عملیات گروهی'
})
const showBulkDropdown = ref(false)

// AI Chat Variables
const showAIChatModal = ref(false)
const showChatHistory = ref(false)
const showGuideModal = ref(false)
const userInput = ref('')
const isGenerating = ref(false)
const messages = ref<ChatMessage[]>([])
const testResult = ref<{ success: boolean; message: string } | null>(null)
const showScrollButton = ref(false)
const chatHistory = ref<ChatSession[]>([])

// AI Settings
const aiSettings = ref({
  model: 'gpt-4.1-nano-2025-04-14',
  wordCount: '1000',
  temperature: '0.7',
  writingStyle: 'professional'
})

// Available AI Models
const availableModels = ref([
  { id: 'o3-deep-research-2025-06-26', name: 'O3 Deep Research' },
  { id: 'o4-mini-deep-research-2025-06-26', name: 'O4 Mini Deep Research' },
  { id: 'gpt-4.1-2025-04-14', name: 'GPT-4.1' },
  { id: 'gpt-4.1-nano-2025-04-14', name: 'GPT-4.1 Nano' }
])

// دریافت مدل‌های موجود از تنظیمات
const fetchAvailableModels = async () => {
  try {
    const apiSettings = await $fetch('/api/admin/api-settings', {
      credentials: 'include'
    })
    
    const settings = apiSettings as APISettingsResponse
    const openaiSettings = (settings?.openai || settings?.data?.openai) as { available_models?: Array<{ id: string; name?: string; [key: string]: unknown }>; [key: string]: unknown } | undefined
    if (openaiSettings?.available_models) {
      interface AIModel {
        id: string
        name: string
        [key: string]: unknown
      }
      availableModels.value = (openaiSettings.available_models as AIModel[]).map((model: AIModel) => ({
        id: model.id,
        name: model.name
      }))
    }
  } catch (error) {
    // خطا در دریافت مدل‌های AI
  }
}
function selectBulkAction(item) {
  bulkAction.value = item.value
  showBulkDropdown.value = false
}

// Methods
const fetchPosts = async () => {
  try {
    const res = await $fetch('/api/posts?all=1')
    interface PostsResponse {
      data?: Post[]
    }
    if (Array.isArray(res)) {
      posts.value = res
    } else if (res && Array.isArray((res as PostsResponse).data)) {
      posts.value = (res as PostsResponse).data
    } else {
      posts.value = []
    }
  } catch (error) {
    posts.value = []
  }
}

const fetchCategories = async () => {
  try {
    // دریافت همه دسته‌بندی‌ها از API
    interface CategoriesResponse {
      data?: Category[]
    }
    const response = await $fetch<CategoriesResponse>('/api/post-categories?all=1')
    if (response && response.data) {
      categories.value = response.data
    }
  } catch (error) {
    categories.value = []
  }
}

// Computed properties
const filteredPosts = computed(() => {
  let filtered = [...posts.value]
  
  // فیلتر بر اساس جستجو
  if (searchQuery.value) {
    filtered = filtered.filter(post => 
      post.title.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      post.excerpt.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }
  
  // فیلتر بر اساس وضعیت
  if (statusFilter.value) {
    filtered = filtered.filter(post => post.status === statusFilter.value)
  }
  
  // فیلتر بر اساس دسته‌بندی
  if (categoryFilter.value) {
    filtered = filtered.filter(post => post.category_id === parseInt(categoryFilter.value))
  }
  
  // مرتب‌سازی
  filtered.sort((a, b) => {
    switch (sortBy.value) {
      case 'title':
        return a.title.localeCompare(b.title)
      case 'status':
        return a.status.localeCompare(b.status)
      case 'updated_at':
        return new Date(b.updated_at).getTime() - new Date(a.updated_at).getTime()
      case 'created_at':
      default:
        return new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
    }
  })
  
  return filtered
})

const paginatedPosts = computed(() => {
  const start = (currentPage.value - 1) * perPage.value
  const end = start + perPage.value
  return filteredPosts.value.slice(start, end)
})

const totalPages = computed(() => Math.ceil(filteredPosts.value.length / perPage.value))

const totalPosts = computed(() => posts.value.length)
const publishedPosts = computed(() => posts.value.filter(p => p.status === 'published').length)
const draftPosts = computed(() => posts.value.filter(p => p.status === 'draft').length)
const scheduledPosts = computed(() => posts.value.filter(p => p.status === 'scheduled').length)

// Methods
const getCategoryName = (categoryId: number) => {
  const category = categories.value.find(c => c.id === categoryId)
  return category ? category.name : 'بدون دسته‌بندی'
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'published': return 'منتشر شده'
    case 'draft': return 'پیش‌نویس'
    case 'scheduled': return 'زمان‌بندی شده'
    default: return status
  }
}

const getStatusBadgeClass = (status: string) => {
  switch (status) {
    case 'published': return 'bg-green-100 text-green-800'
    case 'draft': return 'bg-yellow-100 text-yellow-800'
    case 'scheduled': return 'bg-purple-100 text-purple-800'
    default: return 'bg-gray-100 text-gray-800'
  }
}

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

const toggleSelectAll = () => {
  if (selectedPosts.value.length === filteredPosts.value.length) {
    selectedPosts.value = []
  } else {
    selectedPosts.value = filteredPosts.value.map(p => p.id)
  }
}

const editPost = (post: Post) => {
  // هدایت به صفحه ویرایش نوشته
  navigateTo(`/admin/post-management/edit-post/${post.id}`)
}

const _viewPost = (post: Post) => {
  // اینجا باید به صفحه مشاهده هدایت شود
}

const deletePost = async (post: Post) => {
  if (confirm(`آیا از حذف نوشته "${post.title}" اطمینان دارید؟`)) {
    try {
      // اینجا باید API call برای حذف نوشته انجام شود
      posts.value = posts.value.filter(p => p.id !== post.id)
    } catch (error) {
      alert('خطا در حذف نوشته')
    }
  }
}

// تابع کمکی برای نمایش نام نویسنده
interface Author {
  name?: string
  [key: string]: unknown
}

const getAuthorName = (author: string | Author | null | undefined): string => {
  if (!author) return '-'
  if (typeof author === 'string') return author
  if (typeof author === 'object' && author.name) return author.name
  return '-'
}

// تابع کمکی برای گرفتن slug دسته‌بندی
const getCategorySlug = (categoryId: number) => {
  const category = categories.value.find(c => c.id === categoryId)
  return category ? category.slug : ''
}

// Pagination handlers
const handlePageChange = (page: number) => {
  currentPage.value = page
}

// Watch perPage change to reset page
watch(perPage, () => {
  currentPage.value = 1
})

// Watchers
watch([searchQuery, statusFilter, categoryFilter, sortBy], () => {
  currentPage.value = 1
})

// اسکرول خودکار وقتی پیام‌ها تغییر می‌کنند
watch(messages, (newMessages, oldMessages) => {
  // فقط وقتی پیام جدید اضافه شده اسکرول کن
  if (newMessages.length > (oldMessages?.length || 0)) {
    setTimeout(() => {
      scrollToBottom()
    }, 50)
  }
}, { deep: true })

// اسکرول وقتی مودال باز می‌شود
watch(showAIChatModal, (newValue) => {
  if (newValue) {
    // کمی تاخیر برای اطمینان از لود شدن کامل مودال
    setTimeout(() => {
      scrollToBottom()
      // اضافه کردن event listener برای اسکرول
      const chatContainer = document.querySelector('.chat-messages-container')
      if (chatContainer) {
        chatContainer.addEventListener('scroll', handleScroll)
      }
    }, 100)
  } else {
    // حذف event listener وقتی مودال بسته می‌شود
    const chatContainer = document.querySelector('.chat-messages-container')
    if (chatContainer) {
      chatContainer.removeEventListener('scroll', handleScroll)
    }
  }
})

function handleBulkAction() {
  // TODO: عملیات گروهی بر اساس bulkAction.value روی selectedPosts
  // اینجا باید درخواست API یا تغییر وضعیت انجام شود
  alert(`عملیات گروهی: ${bulkActionLabel.value} روی ${selectedPosts.value.length} آیتم`)
}

          // تابع تشخیص نوع پیام (به‌روزرسانی شده)
          const detectMessageType = (message: string): string => {
            const lowerMessage = message.toLowerCase()
            
            // کلمات کلیدی برای درخواست عنوان
            const titleKeywords = [
              'عنوان', 'تیتر', 'سرتیتر', '5 تا', 'چند تا', 'پیشنهاد عنوان'
            ]
            
            for (const keyword of titleKeywords) {
              if (lowerMessage.includes(keyword)) {
                return 'title_request'
              }
            }
            
            // کلمات کلیدی برای درخواست مقاله
            const articleKeywords = [
              'بنویس', 'نویس', 'تولید کن', 'بساز', 'ایجاد کن', 'ساخت',
              'مقاله', 'محتوا', 'متن', 'نوشته', 'پست', 'بلاگ',
              'برام', 'برای من', 'برام بنویس', 'برام بساز'
            ]
            
            for (const keyword of articleKeywords) {
              if (lowerMessage.includes(keyword)) {
                return 'article_request'
              }
            }
            
            // کلمات کلیدی برای درخواست پیشنهاد
            const suggestionKeywords = [
              'چی', 'پیشنهاد', 'نوع', 'کدام', 'چه', 'چطور'
            ]
            
            for (const keyword of suggestionKeywords) {
              if (lowerMessage.includes(keyword)) {
                return 'suggestion_request'
              }
            }
            
            // کلمات کلیدی برای تأیید
            const confirmationKeywords = [
              'خب', 'باشه', 'تأیید', 'بله', 'درست', 'اوکی', 'ok'
            ]
            
            for (const keyword of confirmationKeywords) {
              if (lowerMessage.includes(keyword)) {
                return 'confirmation'
              }
            }
            
            return 'general'
          }
          
          // تأیید مقاله
          const confirmArticle = (_message: ChatMessage) => {
            userInput.value = 'همین خوبه'
            sendMessage()
          }
          
          // ویرایش مقاله
          const editArticle = (_message: ChatMessage) => {
            userInput.value = 'این قسمت را ویرایش کن'
            sendMessage()
          }

  // تولید session ID منحصر به فرد
  const generateSessionId = () => {
    return `session_${Date.now()}_${Math.random().toString(36).substr(2, 9)}`
  }

  // session ID برای این چت
  const currentSessionId = ref(generateSessionId())

  // AI Chat Functions
  const sendMessage = async () => {
    if (!userInput.value.trim() || isGenerating.value) return
    
    const userMessage = userInput.value.trim()
    const messageType = detectMessageType(userMessage)
    
    messages.value.push({
      role: 'user',
      content: userMessage,
      isArticleRequest: messageType === 'article_request'
    })
    
    userInput.value = ''
    isGenerating.value = true
    
    // اسکرول به پایین بعد از اضافه کردن پیام کاربر
    scrollToBottom()
    
    try {
      // ارسال درخواست به AI با session ID
      const response = await $fetch('/api/admin/ai/chat', {
        method: 'POST',
        body: {
          message: userMessage,
          session_id: currentSessionId.value
        }
      })
      
      const aiResponse = response as AIChatResponse
      
      if (aiResponse && aiResponse.message) {
        messages.value.push({
          role: 'assistant',
          content: aiResponse.message,
          generatedContent: aiResponse.article || null,
          isArticleRequest: false
        })
        
        // اگر مقاله تولید شده، نمایش پیام موفقیت
        if (aiResponse.type === 'article_created' && aiResponse.post_id) {
          alert(`مقاله با موفقیت ایجاد شد! شناسه: ${aiResponse.post_id}`)
        }
      } else {
        messages.value.push({
          role: 'assistant',
          content: 'متأسفانه خطایی در تولید محتوا رخ داد. لطفاً دوباره تلاش کنید.'
        })
      }
      
      // اسکرول به پایین بعد از دریافت پاسخ AI
      scrollToBottom()
    } catch (error: unknown) {
      // استفاده از سیستم خطای جدید
      // const { showError } = useErrorHandler()
      // showError(error)
    
    // نمایش پیام خطا در چت
    messages.value.push({
      role: 'assistant',
      content: 'متأسفانه خطایی در تولید محتوا رخ داد. لطفاً دوباره تلاش کنید.'
    })
    
    // اسکرول به پایین بعد از نمایش خطا
    scrollToBottom()
  } finally {
    isGenerating.value = false
    // اسکرول بعد از اتمام loading
    setTimeout(() => {
      scrollToBottom()
    }, 100)
  }
}

const clearChat = () => {
  messages.value = []
  userInput.value = ''
  // ایجاد session جدید
  currentSessionId.value = generateSessionId()
}

// تابع اسکرول خودکار به آخرین پیام
const scrollToBottom = () => {
  setTimeout(() => {
    const chatContainer = document.querySelector('.chat-messages-container')
    if (chatContainer) {
      // اسکرول نرم به پایین
      chatContainer.scrollTo({
        top: chatContainer.scrollHeight,
        behavior: 'smooth'
      })
      // مخفی کردن دکمه اسکرول
      showScrollButton.value = false
    }
  })
}

// تابع کنترل نمایش دکمه اسکرول
const handleScroll = () => {
  const chatContainer = document.querySelector('.chat-messages-container')
  if (chatContainer) {
    const { scrollTop, scrollHeight, clientHeight } = chatContainer
    // اگر کاربر بالا اسکرول کرده و در پایین نیست، دکمه را نمایش بده
    showScrollButton.value = scrollTop + clientHeight < scrollHeight - 100
  }
}

// تابع نمایش سابقه چت
const showChatHistoryModal = () => {
  loadChatHistory()
  showChatHistory.value = true
}

// تابع بارگذاری سابقه چت از API
const loadChatHistory = async () => {
  try {
    const response = await $fetch<ChatSessionResponse>('/api/admin/chat/sessions')
    chatHistory.value = response.sessions || []
  } catch (_error) {
    chatHistory.value = []
  }
}

// تابع بارگذاری سابقه مکالمات برای AI
const loadAIConversationHistory = async () => {
  try {
    if (currentSessionId.value) {
      const response = await $fetch<ChatSession>(`/api/admin/chat/sessions/${currentSessionId.value}`)
      const sessionResponse = response
      if (sessionResponse.session && sessionResponse.session.messages) {
        // تبدیل پیام‌ها به فرمت مورد نیاز AI
        const aiMessages = sessionResponse.session.messages.map((msg: ChatMessage) => ({
          role: msg.role,
          content: msg.content
        }))
        return aiMessages
      }
    }
    return []
  } catch (error) {
    return []
  }
}

// تابع ذخیره سابقه چت در API
const _saveChatHistory = async () => {
  try {
    if (messages.value.length === 0) return
    
    const title = messages.value[0].content.substring(0, 50) + '...'
    
    if (currentSessionId.value) {
      // بروزرسانی session موجود
      await $fetch(`/api/admin/chat/sessions/${currentSessionId.value}`, {
        method: 'PUT',
        body: { title }
      })
    } else {
      // ایجاد session جدید
      const response = await $fetch('/api/admin/chat/sessions', {
        method: 'POST',
        body: {
          title,
          model: aiSettings.value.model
        }
      })
      const createResponse = response as CreateSessionResponse
      currentSessionId.value = String(createResponse.id)
    }
    
    // ذخیره پیام‌ها
    for (const message of messages.value) {
      await $fetch('/api/admin/chat/messages', {
        method: 'POST',
        body: {
          session_id: currentSessionId.value,
          role: message.role,
          content: message.content
        }
      })
    }
  } catch (error) {
    // خطا در ذخیره سابقه چت
  }
}

// تابع بارگذاری session خاص
const loadSession = async (session: ChatSession) => {
  try {
    const response = await $fetch<ChatSession>(`/api/admin/chat/sessions/${session.id}`)
    messages.value = response.session.messages || []
    aiSettings.value.model = response.session.model
    currentSessionId.value = String(session.id)
    showChatHistory.value = false
    

    
    scrollToBottom()
  } catch (error) {
    // خطا در بارگذاری session
  }
}

// تابع حذف session
const deleteSession = async (sessionId: string) => {
  try {
    await $fetch(`/api/admin/chat/sessions/${sessionId}`, {
      method: 'DELETE'
    })
    
    // بروزرسانی لیست
    await loadChatHistory()
    
    // اگر session فعلی حذف شده، پاک کردن چت
    if (currentSessionId.value === sessionId) {
      clearChat()
      currentSessionId.value = null
    }
  } catch (error) {
    // خطا در حذف session
  }
}

// تابع شروع چت جدید
const startNewChat = () => {
  clearChat()
  currentSessionId.value = null
  showChatHistory.value = false
}

// تابع بررسی تنظیمات OpenAI قبل از باز کردن مودال
const checkOpenAISettings = async () => {
  try {
    // تست اتصال به API
    const apiSettings = await $fetch('/api/admin/api-settings', {
      credentials: 'include'
    })
    // پشتیبانی از هر دو ساختار
    const settingsResponse = apiSettings as APISettingsResponse
    const openaiSettings = settingsResponse.openai || settingsResponse.data?.openai
    if (!openaiSettings) {
      alert('لطفاً ابتدا تنظیمات OpenAI را در بخش تنظیمات فعال کنید و API Key را وارد کنید.')
      return false
    }
    if (!openaiSettings.api_key || !openaiSettings.enabled) {
      alert('لطفاً ابتدا تنظیمات OpenAI را در بخش تنظیمات فعال کنید و API Key را وارد کنید.')
      return false
    }
    return true
  } catch (error: unknown) {
    const errorDetails = error as ErrorResponse
          if (errorDetails.status === 401 || errorDetails.statusCode === 401) {
        alert('خطا در احراز هویت. لطفاً ابتدا وارد شوید و دوباره تلاش کنید.')
        // navigateTo('/auth/login')
      } else if (errorDetails.status === 403 || errorDetails.statusCode === 403) {
        alert('دسترسی محدود شده است. لطفاً مجوزهای خود را بررسی کنید.')
      } else if (errorDetails.status === 404 || errorDetails.statusCode === 404) {
        alert('آدرس API یافت نشد. لطفاً با مدیر سیستم تماس بگیرید.')
      } else if (errorDetails.status >= 500 || errorDetails.statusCode >= 500) {
        alert('خطای سرور. لطفاً بعداً تلاش کنید.')
      } else {
        alert('خطا در بررسی تنظیمات OpenAI. لطفاً ابتدا تنظیمات API را در بخش تنظیمات ذخیره کنید.')
      }
    return false
  }
}

// تابع باز کردن مودال AI نویسنده با بررسی تنظیمات
const openAIChatModal = async () => {
  const hasSettings = await checkOpenAISettings()
  if (hasSettings) {
    showAIChatModal.value = true
    
    // بررسی سابقه مکالمات
    if (currentSessionId.value) {
      const conversationHistory = await loadAIConversationHistory()
    }
    // دریافت مدل‌های موجود
    await fetchAvailableModels()
  }
}

const _createArticleFromAI = (generatedContent: unknown) => {
  if (!generatedContent) {
    alert('محتوای تولید شده یافت نشد!')
    return
  }
  
  try {
    // ذخیره محتوای تولید شده در localStorage
    localStorage.setItem('aiGeneratedContent', JSON.stringify(generatedContent))
    
    // بستن modal
    showAIChatModal.value = false
    
    // هدایت به صفحه افزودن نوشته جدید
    navigateTo('/admin/post-management/add-post')
  } catch (error) {
    alert('خطا در ایجاد مقاله')
  }
}

const testOpenAIConnection = async () => {
  testResult.value = null
  try {
    const res = await $fetch('/api/admin/api-settings/test-openai', { method: 'POST', credentials: 'include' })
    const testResponse = res as TestOpenAIResponse
    if (testResponse.status === 'success') {
      testResult.value = { success: true, message: 'اتصال موفق به OpenAI ✅' }
    } else {
      testResult.value = { success: false, message: testResponse.message || 'اتصال ناموفق به OpenAI' }
    }
  } catch (e: unknown) {
    const error = e as ErrorResponse
    testResult.value = { success: false, message: 'خطا در تست اتصال: ' + (error.data?.message || error.message || 'خطای نامشخص') }
  }
}



// Lifecycle
onMounted(() => {
  fetchPosts()
  fetchCategories()
  fetchAvailableModels()
})
</script>

<style scoped>
/* استایل برای scrollbar چت */
.chat-messages-container {
  scrollbar-width: thin;
  scrollbar-color: #cbd5e0 #f7fafc;
}

.chat-messages-container::-webkit-scrollbar {
  width: 8px;
}

.chat-messages-container::-webkit-scrollbar-track {
  background: #f7fafc;
  border-radius: 4px;
}

.chat-messages-container::-webkit-scrollbar-thumb {
  background-color: #cbd5e0;
  border-radius: 4px;
  border: 1px solid #f7fafc;
}

.chat-messages-container::-webkit-scrollbar-thumb:hover {
  background-color: #a0aec0;
}

/* بهبود عملکرد اسکرول */
.chat-messages-container {
  scroll-behavior: smooth;
  -webkit-overflow-scrolling: touch;
}

/* استایل برای مودال چت */
.fixed.inset-0 {
  backdrop-filter: blur(4px);
}

/* بهبود نمایش پیام‌ها */
.whitespace-pre-wrap {
  word-wrap: break-word;
  overflow-wrap: break-word;
}
</style> 

