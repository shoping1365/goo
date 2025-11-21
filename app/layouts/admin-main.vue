ای ذ<template>
  <div dir="rtl" :class="['flex flex-col min-h-screen', fontClass]">
    
    <header class="h-12 bg-[#222d32] shadow-md flex items-center px-6" dir="rtl">
      <div class="flex-1 flex justify-between items-center w-full">
        <!-- Buttons + welcome as a grouped flex item (RTL order preserved) -->
        <div class="flex items-center space-x-2 space-x-reverse">
          <!-- Mobile burger button -->
          <button class="md:hidden p-2 rounded-md hover:bg-gray-700 transition-colors duration-0" @click="toggleSidebar">
            <svg v-if="!isSidebarOpen" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 text-white">
              <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
            </svg>
            <svg v-else xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 text-white">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>

          <!-- Desktop collapse button -->
          <button class="hidden md:inline-flex p-2 rounded-md hover:bg-gray-700 transition-colors duration-0" :title="isSidebarCollapsed ? 'باز کردن منو' : 'بستن منو'" @click="toggleSidebarCollapse">
            <svg v-if="!isSidebarCollapsed" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 text-white">
              <path stroke-linecap="round" stroke-linejoin="round" d="M3 6h18M3 12h18M3 18h18" />
            </svg>
            <svg v-else xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 text-white">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>

          <!-- Welcome text -->
          <div class="text-xl font-semibold text-white whitespace-nowrap">
            خوش آمدید!
          </div>
        </div>

        <div class="absolute left-1/2 transform -translate-x-1/2 text-center">
          <div class="text-white">{{ currentDate }}</div>
          <div class="text-white">{{ currentTime }}</div>
        </div>

        <div class="flex items-center space-x-4 space-x-reverse">
          <!-- Admin Chat Button -->
          <button class="relative p-2 rounded-md hover:bg-gray-700 transition-colors duration-0" title="چت ادمین" @click="openAdminChat">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 text-white">
              <path stroke-linecap="round" stroke-linejoin="round" d="M8.625 12a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H8.25m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H12m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0h-.375M21 12c0 4.556-4.03 8.25-9 8.25a9.764 9.764 0 01-2.555-.337A5.972 5.972 0 015.41 20.97a5.969 5.969 0 01-.474-.065 4.48 4.48 0 00-.978-2.025c.09-.457-.133-.901-.467-1.226C3.93 16.178 3 14.189 3 12c0-4.556 4.03-8.25 9-8.25s9 3.694 9 8.25z" />
            </svg>
            <!-- نشان پیغام جدید -->
            <span class="absolute -top-1 -right-1 bg-red-500 text-white text-xs rounded-full h-5 w-5 flex items-center justify-center">1</span>
          </button>
          
          <button class="p-2 rounded-md hover:bg-gray-700 transition-colors duration-0" title="تنظیمات پوسته" @click="openThemeSettings">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 text-white">
              <path stroke-linecap="round" stroke-linejoin="round" d="M11.42 15.17L17.25 21 21 17.25l-5.83-5.83m-9.96.96 1.485 1.485m-1.485-1.485a2.25 2.25 0 010-3.182V8.04l3.182-3.182a2.25 2.25 0 013.182 0M4.5 12.75h.008v.008H4.5v-.008zm.75 0a.75.75 0 11-1.5 0 .75.75 0 011.5 0z" />
            </svg>
          </button>
          <button class="p-2 rounded-md hover:bg-gray-700 transition-colors duration-0" title="اعلانات">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 text-white">
              <path stroke-linecap="round" stroke-linejoin="round" d="M14.857 17.082a23.848 23.848 0 005.454-1.31A8.967 8.967 0 0118 9.75V9.75a3.375 3.375 0 00-6.107-1.78l-.001.001-.002-.001-.001-.001A9.01 9.01 0 003 9.75v1.082a23.858 23.858 0 01-5.454 1.31A8.967 8.967 0 016 12.75a8.967 8.967 0 015.982 2.275M12 21a2 2 0 100-4 2 2 0 000 4z" />
            </svg>
          </button>
          


          <!-- Profile dropdown -->
          <div class="relative">
            <button class="p-2 rounded-md hover:bg-gray-700 transition-colors duration-0" title="پروفایل" @click="toggleProfile">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 text-white">
              <path stroke-linecap="round" stroke-linejoin="round" d="M17.982 18.725A7.488 7.488 0 0012 15.75a7.488 7.488 0 00-5.982 2.975m11.963 0a9 9 0 10-11.963 0m11.963 0A8.966 8.966 0 0112 21a8.966 8.966 0 01-5.982-2.275M15 9.75a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
          </button>
            <div v-if="showProfile" class="absolute left-0 mt-2 w-44 bg-white rounded-lg shadow-lg z-50">
              <div class="px-4 py-2 text-sm text-gray-700 border-b" dir="rtl">
                {{ displayName }}
              </div>
              <button class="w-full text-right px-4 py-2 text-sm text-red-600" dir="rtl" @click="logout">خروج</button>
            </div>
          </div>
        </div>
      </div>
    </header>

    <div class="flex flex-1">
      <aside
:class="[
        'admin-sidebar bg-[#222d32] border-l border-gray-600 transform transition-transform duration-0 ease-in-out z-20',
        sidebarClasses
      ]">
        <div class="p-6 border-b border-gray-600">
          <h2 class="text-lg font-bold text-white">پنل مدیریت</h2>
        </div>

        <nav class="p-6 space-y-2 text-sm text-right">
          <NuxtLink
v-if="canShowMenu('/admin')" to="/admin" exact class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                    active-class="bg-gray-700 font-semibold text-white">داشبورد</NuxtLink>

          <!-- مثال برای مدیریت رسانه -->
          <CollapsibleSection
            v-if="canShowSection(['media_library.read','media_upload.read','media_image_compression.read','media_video_compression.read'])"
            title="مدیریت رسانه"
            :persist-key="'media-management'"
            :initial-open="isActivePrefix('/admin/media-management')"
          >
            <div class="collapsible-sub-links-wrapper">
              <NuxtLink
v-if="canShowMenu('/admin/media-management/library')" to="/admin/media-management/library" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">کتابخانه رسانه</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/media-management/upload')" to="/admin/media-management/upload" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">افزودن رسانه جدید</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/media-management/image-compression')" to="/admin/media-management/image-compression" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تنظیمات تصاویر</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/media-management/video-compression')" to="/admin/media-management/video-compression" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تنظیمات ویدیو</NuxtLink>

            </div>
          </CollapsibleSection>

          <CollapsibleSection
            v-if="canShowSection(['post_add.read','post_all.read','post_categories.read'])"
            title="مدیریت نوشته‌ها"
            :persist-key="'posts'"
            :initial-open="isActivePrefix('/admin/post-management')"
          >
            <div class="collapsible-sub-links-wrapper">
              <NuxtLink
v-if="canShowMenu('/admin/post-management/add-post')" to="/admin/post-management/add-post" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">افزودن نوشته</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/post-management/all-posts')" to="/admin/post-management/all-posts" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">همه نوشته‌ها</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/post-management/categories')" to="/admin/post-management/categories" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">دسته‌بندی‌ها</NuxtLink>
            </div>
          </CollapsibleSection>

          <CollapsibleSection
            v-if="canShowSection(['page_create.read','page_all.read','page_categories.read'])"
            title="مدیریت صفحات"
            :persist-key="'page-management'"
            :initial-open="isActivePrefix('/admin/page-management')"
          >
            <div class="collapsible-sub-links-wrapper">
              <NuxtLink
v-if="canShowMenu('/admin/page-management/create-page')" to="/admin/page-management/create-page" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">ایجاد صفحه جدید</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/page-management/all-pages')" to="/admin/page-management/all-pages" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">همه صفحات</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/page-management/page-categories')" to="/admin/page-management/page-categories" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">دسته بندی صفحات</NuxtLink>
            </div>
          </CollapsibleSection>

          <CollapsibleSection
            v-if="canShowSection(['product_brands.read','product_notify_requests.read','product_attributes.read','product_categories.read','product_attribute_groups.read','product_qa.read','product_products.read','product_wishlists.read','product_reviews.read'])"
            title="مدیریت محصولات"
            :persist-key="'product-management'"
            :initial-open="isActivePrefix('/admin/product-management')"
          >
            <div class="collapsible-sub-links-wrapper">
              <NuxtLink
v-if="canShowMenu('/admin/product-management/brands')" to="/admin/product-management/brands" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">برندها</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/product-management/notify-requests')" to="/admin/product-management/notify-requests" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">خبرم کن</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/product-management/attributes')" to="/admin/product-management/attributes" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">ویژگی‌ها</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/product-management/product-categories')" to="/admin/product-management/product-categories" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">دسته‌بندی‌ها</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/product-management/attribute-groups')" to="/admin/product-management/attribute-groups" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">گروه ویژگی‌ها</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/product-management/qa')" to="/admin/product-management/qa" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">پرسش و پاسخ</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/product-management/products')" to="/admin/product-management/products" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">همه محصولات</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/product-management/wishlists')" to="/admin/product-management/wishlists" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">لیست علاقمندی</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/product-management/reviews')" to="/admin/product-management/reviews" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">نظرات محصولات</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/product-management/products/new')" to="/admin/product-management/products/new" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">افزودن محصول جدید</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/product-management/products/bulk-edit')" to="/admin/product-management/products/bulk-edit" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">ویرایش کلی محصولات</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/product-management/warehouse-management')" to="/admin/product-management/warehouse-management" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مدیریت انبارها</NuxtLink>
            </div>
          </CollapsibleSection>

          <CollapsibleSection
            v-if="canShowSection(['user_users.read','user_blacklist.read','user_rating_system.read'])"
            title="مدیریت کاربران"
            :persist-key="'users'"
            :initial-open="isActivePrefix('/admin/users')"
          >
            <div class="collapsible-sub-links-wrapper">
              <NuxtLink
v-if="canShowMenu('/admin/users')" to="/admin/users" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">کاربران</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/verifications')" to="/admin/verifications" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">احراز هویت کاربران</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/users/blacklist')" to="/admin/users/blacklist" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">لیست سیاه</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/users/user-rating-system')" to="/admin/users/user-rating-system" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">سیستم رتبه‌بندی کاربران</NuxtLink>
            </div>
          </CollapsibleSection>

          <CollapsibleSection
            v-if="canShowSection(['role_roles.read','permission_permissions.read'])"
            title="مدیریت دسترسی‌ها"
            :persist-key="'access-management'"
            :initial-open="isActivePrefix('/admin/access-management')"
          >
            <div class="collapsible-sub-links-wrapper">
              <NuxtLink
v-if="canShowMenu('/admin/access-management')" to="/admin/access-management" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">نقش‌های کاربری</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/access-management/permissions')" to="/admin/access-management/permissions" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مجوزها</NuxtLink>
            </div>
          </CollapsibleSection>

          <CollapsibleSection
            v-if="canShowSection(['order_orders.read','order_processing.read','order_cancelled.read','order_shipped.read','order_surveys.read','order_surveys2.read','order_fraud_detection.read'])"
            title="مدیریت سفارشات"
            :persist-key="'transactions'"
            :initial-open="isActivePrefix('/admin/transactions')"
          >
            <div class="collapsible-sub-links-wrapper">
              <NuxtLink
v-if="canShowMenu('/admin/transactions/orders')" to="/admin/transactions/orders" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">سفارشات</NuxtLink>
              <NuxtLink
to="/admin/transactions/orders/in-progress" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">در صف پردازش</NuxtLink>
              <NuxtLink
to="/admin/transactions/orders/processing" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">در حال انجام</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/transactions/orders/shipped')" to="/admin/transactions/orders/shipped" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">ارسال شده</NuxtLink>
              <NuxtLink
to="/admin/transactions/orders/returned" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مرجوع شده</NuxtLink>
              <NuxtLink
to="/admin/transactions/orders/refunded" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مسترد شده</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/transactions/orders/cancelled')" to="/admin/transactions/orders/cancelled" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">کنسل شده</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/transactions/fraud-detection')" to="/admin/transactions/fraud-detection" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تشخیص تقلب</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/transactions/order-surveys')" to="/admin/transactions/order-surveys" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">نظر سنجی بعد از خرید</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/transactions/order-surveys2')" to="/admin/transactions/order-surveys2" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">نظر سنجی بعد از خرید 2</NuxtLink>
            </div>
          </CollapsibleSection>

          <CollapsibleSection
            v-if="canShowSection(['finance_payments.read','finance_payment_gateways.read','finance_giftcard.read','finance_wallet.read','finance_tax.read','finance_accounting.read','finance_installment_payments.read'])"
            title="مالی و پرداخت"
            :persist-key="'finance'"
            :initial-open="isActivePrefix('/admin/finance')"
          >
            <div class="collapsible-sub-links-wrapper">
              <NuxtLink
v-if="canShowMenu('/admin/finance/payments')" to="/admin/finance/payments" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">پرداخت‌ها</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/finance/installment-payments')" to="/admin/finance/installment-payments" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">خرید اقساطی</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/finance/payment-gateways')" to="/admin/finance/payment-gateways" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">درگاه‌های پرداخت</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/finance/giftcard')" to="/admin/finance/giftcard" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">گیفت‌کارت</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/finance/wallet')" to="/admin/finance/wallet" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">کیف پول</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/finance/tax')" to="/admin/finance/tax" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مالیات و تطابق</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/finance/accounting')" to="/admin/finance/accounting" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">اتصال حسابداری/ERP</NuxtLink>
            </div>
          </CollapsibleSection>

          <CollapsibleSection
            v-if="canShowSection(['shipping_methods.read','shipping_fleet.read','shipping_driver_tracking.read','shipping_delivery_slots.read'])"
            title="حمل‌ونقل و لجستیک"
            :persist-key="'shipping'"
            :initial-open="isActivePrefix('/admin/shipping-methods')"
          >
            <div class="collapsible-sub-links-wrapper">
              <NuxtLink
v-if="canShowMenu('/admin/shipping-methods')" to="/admin/shipping-methods" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">روش‌های ارسال</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/shipping-methods/fleet')" to="/admin/shipping-methods/fleet" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مدیریت ناوگان</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/shipping-methods/driver-tracking')" to="/admin/shipping-methods/driver-tracking" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">رهگیری رانندگان</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/shipping-methods/delivery-slots')" to="/admin/shipping-methods/delivery-slots" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">زمان‌بندی تحویل</NuxtLink>
            </div>
          </CollapsibleSection>

          <CollapsibleSection
            v-if="canShowSection(['statistics_visits.read','statistics_sales.read','statistics_orders.read','statistics_customers.read','statistics_inventory.read','statistics_abandoned_carts.read','statistics_user_behavior.read'])"
            title="آمار"
            :persist-key="'statistics'"
            :initial-open="isActivePrefix('/admin/statistics')"
          >
            <div class="collapsible-sub-links-wrapper">
              <NuxtLink
v-if="canShowMenu('/admin/statistics/visits')" to="/admin/statistics/visits" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">آمار بازدید</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/statistics/sales')" to="/admin/statistics/sales" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">آمار فروش</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/statistics/orders')" to="/admin/statistics/orders" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">آمار سفارش‌ها</NuxtLink>
               <NuxtLink
v-if="canShowMenu('/admin/statistics/customers')" to="/admin/statistics/customers" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">آمار مشتریان</NuxtLink>
               <NuxtLink
v-if="canShowMenu('/admin/statistics/inventory')" to="/admin/statistics/inventory" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">آمار موجودی</NuxtLink>
               <NuxtLink
v-if="canShowMenu('/admin/statistics/abandoned-carts')" to="/admin/statistics/abandoned-carts" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">سبدهای خرید رها شده</NuxtLink>
               <NuxtLink
v-if="canShowMenu('/admin/statistics/user-behavior')" to="/admin/statistics/user-behavior" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">آمار رفتار کاربران</NuxtLink>
            </div>
          </CollapsibleSection>

          <CollapsibleSection
            v-if="canShowSection(['marketing.read','marketing_discounts.read','marketing_integrations.read','marketing_referral.read','marketing_loyalty.read','marketing_ab_testing.read','marketing_recommender.read','marketing_fraud_detection.read'])"
            title="بازاریابی و فروش"
            :persist-key="'marketing'"
            :initial-open="isActivePrefix('/admin/marketing')"
          >
            <div class="collapsible-sub-links-wrapper">
              <NuxtLink
v-if="canShowMenu('/admin/marketing')" to="/admin/marketing" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">بازاریابی</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/marketing/discounts')" to="/admin/marketing/discounts" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تخفیف‌ها و کوپن‌ها</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/marketing/integrations')" to="/admin/marketing/integrations" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">ادغام‌ها</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/marketing/referral')" to="/admin/marketing/referral" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">سیستم معرفی</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/marketing/loyalty')" to="/admin/marketing/loyalty" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">برنامه وفاداری</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/marketing/ab-testing')" to="/admin/marketing/ab-testing" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تست A/B</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/marketing/recommender')" to="/admin/marketing/recommender" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">پیشنهادات هوشمند</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/marketing/fraud-detection')" to="/admin/marketing/fraud-detection" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تشخیص تقلب</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/marketing/lucky-draw')" to="/admin/marketing/lucky-draw" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">گرونه شانس</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/marketing/display-widgets')" to="/admin/marketing/display-widgets" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">ویجت‌های نمایشی</NuxtLink>
            </div>
          </CollapsibleSection>

          <!-- تحلیل هوش مصنوعی و BI -->
          <CollapsibleSection
            v-if="canShowSection(['analytics_bi.read','analytics_reports.read'])"
            title="تحلیل هوش مصنوعی و BI"
            :persist-key="'analytics-bi'"
            :initial-open="isActivePrefix('/admin/analytics')"
          >
            <div class="collapsible-sub-links-wrapper">
              <NuxtLink
v-if="canShowMenu('/admin/analytics/bi')" to="/admin/analytics/bi" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">داشبورد تحلیل هوش مصنوعی</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/analytics/bi/ProductAnalysisDashboard')" to="/admin/analytics/bi/ProductAnalysisDashboard" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تحلیل محصولات</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/analytics/bi/SalesAnalysisDashboard')" to="/admin/analytics/bi/SalesAnalysisDashboard" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تحلیل فروش</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/analytics/bi/CustomerAnalysisDashboard')" to="/admin/analytics/bi/CustomerAnalysisDashboard" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تحلیل مشتریان</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/analytics/bi/OverviewDashboard')" to="/admin/analytics/bi/OverviewDashboard" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">داشبورد کلی</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/analytics/bi/BISettingsModal')" to="/admin/analytics/bi/BISettingsModal" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تنظیمات BI</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/analytics/Reports')" to="/admin/analytics/Reports" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">گزارش‌ها</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/analytics/Reports/search-analytics')" to="/admin/analytics/Reports/search-analytics" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">گزارشات جستجو</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/analytics/ci')" to="/admin/analytics/ci" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">هوش رقابتی CI</NuxtLink>
            </div>
          </CollapsibleSection>

          <!-- ابزارهای سئو -->
          <CollapsibleSection
            v-if="canShowSection(['seo_schema.read','seo_keywords.read','seo_content.read','seo_page_analysis.read','seo_content_optimization.read','seo_site_structure.read','seo_link_management.read','seo_image_optimization.read','seo_local_seo.read','seo_performance.read','seo_user_experience.read','seo_monitoring.read','seo_redirects.read'])"
            title="ابزارهای سئو"
            :persist-key="'seo'"
            :initial-open="isActivePrefix('/admin/seo')"
          >
            <div class="collapsible-sub-links-wrapper">
              <NuxtLink
v-if="canShowMenu('/admin/seo/SchemaManagement')" to="/admin/seo/SchemaManagement" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مدیریت اسکیما</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/seo/keyword-tracking')" to="/admin/seo/keyword-tracking" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تحلیل و ردیابی کلمات کلیدی</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/seo/content-generation')" to="/admin/seo/content-generation" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تولید محتوا</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/seo/page-analysis')" to="/admin/seo/page-analysis" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">آنالیز صفحات</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/seo/content-optimization')" to="/admin/seo/content-optimization" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">بهینه‌سازی محتوا و متا</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/seo/site-structure')" to="/admin/seo/site-structure" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">ساختار و ناوبری سایت</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/seo/link-management')" to="/admin/seo/link-management" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مدیریت لینک‌ها</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/seo/image-optimization')" to="/admin/seo/image-optimization" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">بهینه‌سازی تصاویر</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/seo/local-seo')" to="/admin/seo/local-seo" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">سئو محلی</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/seo/performance')" to="/admin/seo/performance" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">عملکرد و سرعت</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/seo/user-experience')" to="/admin/seo/user-experience" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تجربه کاربری و ریسپانسیو</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/seo/monitoring')" to="/admin/seo/monitoring" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">نظارت و گزارش‌گیری</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/seo/redirects')" to="/admin/seo/redirects" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تغییر مسیر</NuxtLink>
            </div>
          </CollapsibleSection>

          <!-- مدیریت محتوا -->
          <CollapsibleSection
            v-if="canShowSection(['content_menus.read','content_header.read','content_footer.read','content_banners.read','mobile_app_header.read','mobile_app_footer.read','mobile_app_navigation.read'])"
            title="مدیریت محتوا"
            :persist-key="'content'"
            :initial-open="isActivePrefix('/admin/content') || isActivePrefix('/admin/menu-management')"
          >
            <div class="collapsible-sub-links-wrapper">
              <NuxtLink
v-if="canShowMenu('/admin/content/menu-management')" to="/admin/content/menu-management" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مدیریت منوها</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/content/header-management')" to="/admin/content/header-management" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مدیریت هدر</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/content/footer-management')" to="/admin/content/footer-management" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مدیریت فوتر</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/content/banners')" to="/admin/content/banners" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">بنرها و اسلایدر</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/content/mobile-app-header-management')" to="/admin/content/mobile-app-header-management" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">هدر موبایل و اپلیکیشن</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/content/mobile-app-footer-management')" to="/admin/content/mobile-app-footer-management" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">فوتر موبایل و اپلیکیشن</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/content/mobile-app-navigation-management')" to="/admin/content/mobile-app-navigation-management" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">ناوبری موبایل و اپلیکیشن</NuxtLink>
            </div>
          </CollapsibleSection>

          <!-- پیامک و اعلان‌ها -->
          <CollapsibleSection
            v-if="canShowSection(['notification_settings.read','notification_monitoring.read','notification_alerts.read','notification_patterns.read'])"
            title="پیامک و اعلان‌ها"
            :persist-key="'notifications'"
            :initial-open="isActivePrefix('/admin/notifications')"
          >
            <div class="collapsible-sub-links-wrapper">
              <NuxtLink
v-if="canShowMenu('/admin/notifications/system-settings')" to="/admin/notifications/system-settings" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تنظیمات سامانه</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/notifications/monitoring')" to="/admin/notifications/monitoring" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مانیتورینگ</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/notifications/alerts')" to="/admin/notifications/alerts" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">اعلان‌ها</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/notifications/patterns')" to="/admin/notifications/patterns" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">پترن‌ها</NuxtLink>
            </div>
          </CollapsibleSection>

          <!-- چت آنلاین و پشتیبانی -->
          <CollapsibleSection
            v-if="canShowSection(['support_live_chat.read','support_tickets.read','support_chat_window.read','support_widgets.read','support_operators.read','support_knowledge_base.read','support_statistics.read','support_settings.read','support_performance_monitoring.read'])"
            title="چت آنلاین و پشتیبانی"
            :persist-key="'support-management'"
            :initial-open="isActivePrefix('/admin/support-management')"
          >
            <div class="collapsible-sub-links-wrapper">
              <NuxtLink
v-if="canShowMenu('/admin/support-management/chat-window')" to="/admin/support-management/chat-window" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">پنجره چت</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/support-management/live-chat')" to="/admin/support-management/live-chat" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">چت زنده</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/support-management/tickets')" to="/admin/support-management/tickets" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تیکت‌ها</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/support-management/widgets')" to="/admin/support-management/widgets" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">ابزارک</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/support-management/operators')" to="/admin/support-management/operators" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">اپراتورها</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/support-management/knowledge-base')" to="/admin/support-management/knowledge-base" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">پایگاه دانش</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/support-management/statistics')" to="/admin/support-management/statistics" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">آمار</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/support-management/performance-monitoring')" to="/admin/support-management/performance-monitoring" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مانیتورینگ عملکرد</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/support-management/settings')" to="/admin/support-management/settings" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تنظیمات</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/support-management/ai-chatbot')" to="/admin/support-management/ai-chatbot" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">چت بات هوش مصنوعی</NuxtLink>

            </div>
          </CollapsibleSection>

          <CollapsibleSection
            v-if="canShowSection(['security_traffic.read','security_blocked.read','security_login.read','security_scan.read','security_firewall.read','security_two_factor.read','security_recaptcha.read','security_bots.read','security_cloudflare.read'])"
            title="امنیت سامانه"
            :persist-key="'system-security'"
            :initial-open="isActivePrefix('/admin/system-security')"
          >
            <div class="collapsible-sub-links-wrapper">

              <NuxtLink
v-if="canShowMenu('/admin/system-security/login')" to="/admin/system-security/login" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">ورود</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/system-security/firewall')" to="/admin/system-security/firewall" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">فایروال</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/system-security/bots')" to="/admin/system-security/bots" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">ربات‌ها</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/system-security/scan')" to="/admin/system-security/scan" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">اسکن</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/system-security/settings')" to="/admin/system-security/settings" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تنظیمات امنیت</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/system-security/recaptcha')" to="/admin/system-security/recaptcha" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">ریکپچا</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/system-security/blocked')" to="/admin/system-security/blocked" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مسدودی</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/system-security/traffic')" to="/admin/system-security/traffic" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">ترافیک آنلاین</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/system-security/two-factor')" to="/admin/system-security/two-factor" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تایید دو مرحله‌ای</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/system-security/cloudflare')" to="/admin/system-security/cloudflare" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">کلادفلر</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/settings/security')" to="/admin/settings/security" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">محدودیت نرخ درخواست</NuxtLink>
            </div>
          </CollapsibleSection>

          <NuxtLink
v-if="canShowMenu('/admin/settings')" to="/admin/settings" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white flex items-center"
                    active-class="bg-gray-700 font-semibold text-white">
            <span class="i-heroicons-cog-6-tooth ml-2"></span>
            تنظیمات
          </NuxtLink>

          <!-- بخش Developer - فقط برای developer -->
          <CollapsibleSection v-if="isDeveloper" title="دولوپر" :persist-key="'developer'" :initial-open="isActivePrefix('/admin/developer')">
            <div class="collapsible-sub-links-wrapper">
              <NuxtLink
v-if="canShowMenu('/admin/developer/database-status')" to="/admin/developer/database-status" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">وضعیت دیتابیس</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/system-test')" to="/admin/developer/system-test" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تست سیستم</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/media-debug')" to="/admin/developer/media-debug" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">دیباگ رسانه</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/test-api')" to="/admin/developer/test-api" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تست API</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/api')" to="/admin/developer/api" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">API</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/test-connection')" to="/admin/developer/test-connection" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تست اتصال سیستم</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/wordpress-migration')" to="/admin/developer/wordpress-migration" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">انتقال محصولات و سفارشات</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/digikala-import')" to="/admin/developer/digikala-import" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">انتقال اطلاعات از دیجی کالا</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/system-performance')" to="/admin/developer/system-performance" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">عملکرد سیستم</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/performance-optimizer')" to="/admin/developer/performance-optimizer" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">بهینه‌ساز عملکرد</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/devops')" to="/admin/developer/devops" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">DevOps</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/log-viewer')" to="/admin/developer/log-viewer" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مشاهده لاگ‌ها</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/code-editor')" to="/admin/developer/code-editor" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">ویرایشگر کد</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/database-explorer')" to="/admin/developer/database-explorer" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">کاوشگر دیتابیس</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/git-manager')" to="/admin/developer/git-manager" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مدیریت گیت</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/package-manager')" to="/admin/developer/package-manager" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مدیریت پکیج‌ها</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/docker-manager')" to="/admin/developer/docker-manager" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مدیریت Docker</NuxtLink>
              <!-- حذف لینک‌های غیرضروری دولوپر طبق درخواست -->
              <NuxtLink
v-if="canShowMenu('/admin/developer/backup-manager')" to="/admin/developer/backup-manager" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مدیریت پشتیبان‌گیری</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/monitoring-dashboard')" to="/admin/developer/monitoring-dashboard" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">داشبورد مانیتورینگ</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/security-scanner')" to="/admin/developer/security-scanner" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">اسکنر امنیتی</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/performance-analyzer')" to="/admin/developer/performance-analyzer" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">تحلیل‌گر عملکرد</NuxtLink>
              <NuxtLink
v-if="canShowMenu('/admin/developer/deployment-manager')" to="/admin/developer/deployment-manager" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                        active-class="bg-gray-700 font-semibold text-white">مدیر استقرار</NuxtLink>
            </div>
          </CollapsibleSection>
        </nav>
      </aside>

      <main class="flex-1 p-6 sm:p-6 transition-all duration-0 ease-in-out bg-sky-50" :class="isSidebarCollapsed ? 'md:pr-4' : ''">
      <slot />
      </main>
    </div>

    <!-- Footer Component for Admin Pages - REMOVED -->
    <!-- Footer should not be shown in admin panel -->

    <div v-if="isSidebarOpen && windowWidth < 768" class="fixed inset-0 bg-black opacity-50 z-30 md:hidden" @click="isSidebarOpen = false"></div>

    <div v-if="showThemeSettings" class="fixed inset-0 flex items-center justify-center z-50">
      <!-- Backdrop شفاف -->
      <div class="absolute inset-0 bg-black bg-opacity-20 backdrop-blur-sm" @click="showThemeSettings = false"></div>
      <!-- مودال -->
      <div class="relative bg-white rounded-2xl shadow-2xl p-8 w-full max-w-md mx-4 transform transition-all duration-0 scale-100">
        <!-- دکمه بستن -->
        <button class="absolute top-6 left-4 text-gray-400 hover:text-gray-600 transition-colors duration-0" @click="showThemeSettings = false">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
        <h2 class="text-lg font-bold mb-4 text-center">تنظیمات پوسته</h2>
        <div class="space-y-4 text-right">
          <div>
            <label for="fontSelect" class="block text-sm font-medium text-gray-700 mb-1">نوع فونت:</label>
            <select id="fontSelect" v-model="selectedFont" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50" @change="adminTheme.font = selectedFont">
              <option value="iransans">ایران‌سن‌س</option>
              <option value="iransansfanum">ایران‌سن‌س فا‌نام</option>
              <option value="sahel">ساحل</option>
              <option value="shabnam">شبنم</option>
              <option value="samim">صمیم</option>
              <option value="parastoo">پرستو</option>
              <option value="tanha">تنها</option>
              <option value="vazir">وزیر</option>
              <option value="yekan">یکان</option>
            </select>
          </div>
          <div>
            <label for="colorScheme" class="block text-sm font-medium text-gray-700 mb-1">رنگ‌بندی:</label>
            <select id="colorScheme" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
              <option value="light">روشن</option>
              <option value="dark">تیره</option>
            </select>
          </div>
        </div>
        <div class="mt-6 flex justify-end space-x-2 space-x-reverse">
          <button class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors duration-0 font-medium" @click="showThemeSettings = false">انصراف</button>
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors duration-0 font-medium shadow-md" @click="applyThemeSettings">ذخیره</button>
        </div>
      </div>

    </div>



    <!-- Admin Button Bar (Bottom) -->
    <div class="fixed bottom-0 left-0 right-0 bg-[#222d32] text-white p-2 flex justify-center items-center gap-6 text-sm z-10 shadow-md">
      <span>پنل مدیریت</span>
      <span>|</span>
      <NuxtLink to="/" class="text-blue-300 hover:text-blue-100">بازگشت به سایت</NuxtLink>
      <span>|</span>
      <span>کاربر فعال: {{ displayName }}</span>
    </div>
  </div>
  
  <!-- Global Confirm Dialog for all admin pages -->
  <ConfirmDialog ref="confirmDialogRef" />
  
  <!-- Toast Container for notifications -->
  <ToastContainer />
</template>

<script setup lang="ts">
import { computed, onMounted, onUnmounted, provide, reactive, ref } from 'vue';
import { useRoute } from 'vue-router';
import CollapsibleSection from '~/components/common/CollapsibleSection.vue';
import ConfirmDialog from '~/components/common/ConfirmDialog.vue';
import ToastContainer from '~/components/common/ToastContainer.vue';
import { useAuthState } from '~/composables/useAuthState';
import { useAdminThemeStore } from '~/stores/adminTheme';

// تعریف navigateTo برای Nuxt 3
declare const navigateTo: (to: string) => Promise<void>;

// استفاده از Global Auth State
const { setUser } = useAuthState()



const showThemeSettings = ref(false);
const isSidebarOpen = ref(false); // متغیر برای کنترل وضعیت سایدبار
const isSidebarCollapsed = ref(false) // وضعیت تا شدن سایدبار روی دسکتاپ
const windowWidth = ref(undefined); // متغیر reactive برای عرض پنجره

const currentDate = ref('');
const currentTime = ref('');
let timer: NodeJS.Timeout | null = null;

const showProfile = ref(false);
const route = useRoute();

// Provide confirm dialog for all admin pages
const confirmDialogRef = ref<{ show?: (options: Record<string, unknown>) => void; [key: string]: unknown } | null>(null);
provide('confirmDialogRef', confirmDialogRef);

// Session-based user state
const currentUser = ref(null);
const displayName = computed(() => {
  return currentUser.value?.name || currentUser.value?.username || 'کاربر';
});

// بررسی نقش developer
const isDeveloper = computed(() => {
  return currentUser.value?.role === 'developer';
});

// وضعیت بنر سراسری
const banner = reactive({ visible: false, title: '', message: '', isOffline: false })
const showOffline = () => {
  banner.visible = true
  banner.isOffline = true
  banner.title = 'عدم دسترسی به اینترنت'
  banner.message = 'اتصال اینترنت برقرار نیست. لطفاً اتصال خود را بررسی کنید.'
}
const hideOffline = () => {
  banner.isOffline = false
  banner.visible = false
}

// احراز هویت - دریافت اطلاعات کاربر واقعی
onMounted(async () => {
  try {
    const response = await $fetch('/api/auth/me', {
      credentials: 'include'
    }) as { authenticated?: boolean; user?: Record<string, unknown>; [key: string]: unknown }
    
    if (response?.authenticated && response?.user) {
      currentUser.value = response.user
      setUser(response.user) // sync با global state
    } else {
      // اگر احراز نشده، redirect به صفحه ورود
      await navigateTo('/auth/login')
    }
  } catch (_error) {
    // در صورت عدم دسترسی یا خطا، redirect به صفحه ورود
    await navigateTo('/auth/login')
  }

  window.addEventListener('offline', showOffline)
  window.addEventListener('online', hideOffline)
});

onUnmounted(() => {
  if (timer) {
    clearInterval(timer);
  }
  window.removeEventListener('resize', handleResizeSidebarMobile);
  window.removeEventListener('offline', showOffline)
  window.removeEventListener('online', hideOffline)
});

const updateDateTime = () => {
  const now = new Date();
  currentDate.value = now.toLocaleDateString('fa-IR', { 
    year: 'numeric', 
    month: 'long', 
    day: 'numeric' 
  });
  currentTime.value = now.toLocaleTimeString('fa-IR', { 
    hour: '2-digit', 
    minute: '2-digit',
    hour12: false
  });
};

onMounted(() => {
  updateDateTime();
  timer = setInterval(updateDateTime, 60000); // Update every minute instead of every second

  // تنظیم مقدار اولیه عرض پنجره
  windowWidth.value = window.innerWidth;

  // بستن سایدبار در موبایل به صورت پیش‌فرض
  if (windowWidth.value && windowWidth.value < 768) {
    isSidebarOpen.value = false;
  }
  // اگر کاربر سایز صفحه را تغییر داد و وارد موبایل شد، سایدبار بسته شود
  window.addEventListener('resize', handleResizeSidebarMobile, { passive: true });
  

});

function handleResizeSidebarMobile() {
  windowWidth.value = window.innerWidth;
  if (windowWidth.value && windowWidth.value < 768) {
    isSidebarOpen.value = false;
  }
}

const adminTheme = useAdminThemeStore()
const selectedFont = ref(adminTheme.font)

const fontClass = computed(() => {
  switch (adminTheme.font) {
    case 'sahel': return 'font-sahel'
    case 'shabnam': return 'font-shabnam'
    case 'samim': return 'font-samim'
    default: return 'font-iransans'
  }
})

// computed property برای کلاس‌های سایدبار
const sidebarClasses = computed(() => {
  // در موبایل: اگر isSidebarOpen false باشد، سایدبار مخفی شود
  // در دسکتاپ: اگر isSidebarCollapsed true باشد، سایدبار مخفی شود
  if (windowWidth.value && windowWidth.value < 768) {
    return !isSidebarOpen.value ? '-translate-x-full w-0' : 'w-64 translate-x-0'
  } else {
    return isSidebarCollapsed.value ? 'md:-translate-x-full md:w-0' : 'md:translate-x-0 md:w-64'
  }
})

const openThemeSettings = () => {
  showThemeSettings.value = true;
};

const openAdminChat = () => {
  // باز کردن چت در پنجره جداگانه مثل تلگرام
  const chatWindow = window.open(
    '/admin/chat-window', 
    'AdminChat',
    'width=400,height=600,resizable=yes,scrollbars=no,toolbar=no,menubar=no,location=no,status=no'
  );
  
  if (chatWindow) {
    // تنظیم موقعیت پنجره در گوشه پایین راست
    chatWindow.moveTo(
      screen.width - 420,
      screen.height - 650
    );
  }
};

// openUserSettings removed - not used

const applyThemeSettings = () => {
  // فونت از selectedFont گرفته شده و همین الان ذخیره شده است.
  alert(`تنظیمات پوسته ذخیره شد.`);

  showThemeSettings.value = false;
};

const toggleSidebar = () => {
  // در موبایل، isSidebarOpen را تغییر می‌دهیم
  // در دسکتاپ، isSidebarCollapsed را تغییر می‌دهیم
  if (windowWidth.value && windowWidth.value < 768) {
    isSidebarOpen.value = !isSidebarOpen.value;
  } else {
    isSidebarCollapsed.value = !isSidebarCollapsed.value;
  }
};



const toggleSidebarCollapse = () => {
  isSidebarCollapsed.value = !isSidebarCollapsed.value
}

function toggleProfile() {
  showProfile.value = !showProfile.value;
}


async function logout() {
  const auth = useAuthState()
  try {
    await auth.logout()

  } catch (e) {
    console.warn('Logout failed but proceeding with redirect', e)
  }

  showProfile.value = false
  // redirect after clearing client state
  window.location.href = '/auth/login'
}

const isActivePrefix = (prefix) => route.path.startsWith(prefix);

defineExpose({ displayName });

// const { hasPermission, permissions } = useAuth() // احراز هویت غیرفعال شده است

// نگاشت مسیر به permission - currently unused but kept for future use
const _menuPermissions = {
  '/admin': 'admin_panel_access',
  
  // مدیریت رسانه
  '/admin/media-management/library': 'media_library.read',
  '/admin/media-management/upload': 'media_upload.read',
  '/admin/media-management/image-compression': 'media_image_compression.read',
  '/admin/media-management/video-compression': 'media_video_compression.read',
  
  // مدیریت نوشته‌ها
  '/admin/post-management/add-post': 'post_add.read',
  '/admin/post-management/all-posts': 'post_all.read',
  '/admin/post-management/categories': 'post_categories.read',
  
  // مدیریت صفحات
  '/admin/page-management/create-page': 'page_create.read',
  '/admin/page-management/all-pages': 'page_all.read',
  '/admin/page-management/page-categories': 'page_categories.read',
  
  // مدیریت محصولات
  '/admin/product-management/brands': 'product_brands.read',
  '/admin/product-management/notify-requests': 'product_notify_requests.read',
  '/admin/product-management/attributes': 'product_attributes.read',
  '/admin/product-management/product-categories': 'product_categories.read',
  '/admin/product-management/attribute-groups': 'product_attribute_groups.read',
  '/admin/product-management/qa': 'product_qa.read',
  '/admin/product-management/products': 'product_products.read',
  '/admin/product-management/wishlists': 'product_wishlists.read',
  '/admin/product-management/reviews': 'product_reviews.read',
  '/admin/product-management/products/new': 'product_products.create',
  '/admin/product-management/products/bulk-edit': 'product_products.update',
  
  // مدیریت کاربران
  '/admin/users': 'user_users.read',
  '/admin/verifications': 'user_users.read',
  '/admin/users/blacklist': 'user_blacklist.read',
  '/admin/users/user-rating-system': 'user_rating_system.read',
  
  // مدیریت دسترسی‌ها
  '/admin/access-management': 'role_roles.read',
  '/admin/access-management/permissions': 'permission_permissions.read',
  
  // مدیریت سفارشات
  '/admin/transactions/orders': 'order_orders.read',
  '/admin/transactions/orders/processing': 'order_processing.read',
  '/admin/transactions/orders/cancelled': 'order_cancelled.read',
  '/admin/transactions/orders/shipped': 'order_shipped.read',
  '/admin/transactions/order-surveys': 'order_surveys.read',
        '/admin/transactions/order-surveys2': 'order_surveys2.read',
      '/admin/transactions/fraud-detection': 'order_fraud_detection.read',
  
  // مالی و پرداخت
  '/admin/finance/payments': 'finance_payments.read',
  '/admin/finance/installment-payments': 'finance_installment_payments.read',
  '/admin/finance/payment-gateways': 'finance_payment_gateways.read',
  '/admin/finance/giftcard': 'finance_giftcard.read',
  '/admin/finance/wallet': 'finance_wallet.read',
  '/admin/finance/tax': 'finance_tax.read',
  '/admin/finance/accounting': 'finance_accounting.read',
  
  // حمل‌ونقل و لجستیک
  '/admin/shipping-methods': 'shipping_methods.read',
  '/admin/shipping-methods/fleet': 'shipping_fleet.read',
  '/admin/shipping-methods/driver-tracking': 'shipping_driver_tracking.read',
  '/admin/shipping-methods/delivery-slots': 'shipping_delivery_slots.read',
  
  // بازاریابی و فروش
  '/admin/marketing': 'marketing.read',
  '/admin/marketing/discounts': 'marketing_discounts.read',
  '/admin/marketing/integrations': 'marketing_integrations.read',
  '/admin/marketing/referral': 'marketing_referral.read',
  '/admin/marketing/loyalty': 'marketing_loyalty.read',
  '/admin/marketing/ab-testing': 'marketing_ab_testing.read',
  '/admin/marketing/recommender': 'marketing_recommender.read',
  '/admin/marketing/lucky-draw': 'marketing.read',
  
  // ابزارهای سئو
  '/admin/seo/SchemaManagement': 'seo_schema.read',
  '/admin/seo/keyword-tracking': 'seo_keywords.read',
  '/admin/seo/content-generation': 'seo_content.read',
  '/admin/seo/page-analysis': 'seo_page_analysis.read',
  '/admin/seo/content-optimization': 'seo_content_optimization.read',
  '/admin/seo/site-structure': 'seo_site_structure.read',
  '/admin/seo/link-management': 'seo_link_management.read',
  '/admin/seo/image-optimization': 'seo_image_optimization.read',
  '/admin/seo/local-seo': 'seo_local_seo.read',
  '/admin/seo/performance': 'seo_performance.read',
  '/admin/seo/user-experience': 'seo_user_experience.read',
  '/admin/seo/monitoring': 'seo_monitoring.read',
  '/admin/seo/redirects': 'seo_redirects.read',
  
  // تحلیل و گزارش‌ها
  '/admin/analytics/reports': 'analytics_reports.read',
  '/admin/analytics/bi': 'analytics_bi.read',
  '/admin/analytics/bi/ProductAnalysisDashboard': 'analytics_bi.read',
  '/admin/analytics/bi/SalesAnalysisDashboard': 'analytics_bi.read',
  '/admin/analytics/bi/CustomerAnalysisDashboard': 'analytics_bi.read',
  '/admin/analytics/bi/OverviewDashboard': 'analytics_bi.read',
  '/admin/analytics/bi/BISettingsModal': 'analytics_bi.read',
  '/admin/analytics/Reports': 'analytics_reports.read',
  '/admin/analytics/ci': 'analytics_bi.read',
  
  // مدیریت پشتیبانی
  '/admin/support-management/live-chat': 'support_live_chat.read',
  '/admin/support-management/tickets': 'support_tickets.read',
  '/admin/support-management/chat-window': 'support_chat_window.read',
  '/admin/support-management/widgets': 'support_widgets.read',
  '/admin/support-management/operators': 'support_operators.read',
  '/admin/support-management/knowledge-base': 'support_knowledge_base.read',
  '/admin/support-management/statistics': 'support_statistics.read',
  '/admin/support-management/performance-monitoring': 'support_performance_monitoring.read',
  '/admin/support-management/settings': 'support_settings.read',
  '/admin/support-management/ai-chatbot': 'support_ai_chatbot.read',
  
  
  // مدیریت محتوا
  '/admin/content/menu-management': 'content_menus.read',
  '/admin/content/header-management': 'content_header.read',
  '/admin/content/footer-management': 'content_footer.read',
  '/admin/content/banners': 'content_banners.read',
  '/admin/content/mobile-app-header-management': 'mobile_app_header.read',
  '/admin/content/mobile-app-footer-management': 'mobile_app_footer.read',
  '/admin/content/mobile-app-navigation-management': 'mobile_app_navigation.read',
  
  // پیامک و اعلان‌ها
  '/admin/notifications/system-settings': 'notification_settings.read',
  '/admin/notifications/monitoring': 'notification_monitoring.read',
  '/admin/notifications/alerts': 'notification_alerts.read',
  '/admin/notifications/patterns': 'notification_patterns.read',
  
  // آمار
  '/admin/statistics/visits': 'statistics_visits.read',
  '/admin/statistics/sales': 'statistics_sales.read',
  '/admin/statistics/orders': 'statistics_orders.read',
  '/admin/statistics/customers': 'statistics_customers.read',
  '/admin/statistics/inventory': 'statistics_inventory.read',
  '/admin/statistics/abandoned-carts': 'statistics_abandoned_carts.read',
  '/admin/statistics/user-behavior': 'statistics_user_behavior.read',
  
  // امنیت سامانه
  '/admin/system-security/traffic': 'security_traffic.read',
  '/admin/system-security/blocked': 'security_blocked.read',
  '/admin/system-security/login': 'security_login.read',
  '/admin/system-security/two-factor': 'security_two_factor.read',
  '/admin/system-security/scan': 'security_scan.read',
  '/admin/system-security/firewall': 'security_firewall.read',
  '/admin/system-security/recaptcha': 'security_recaptcha.read',
  '/admin/system-security/bots': 'security_bots.read',
  '/admin/system-security/cloudflare': 'security_cloudflare.read',
  
  // تنظیمات
  '/admin/settings': 'settings.read',
  '/admin/settings/api-settings': 'settings.read',

  // صفحات توسعه
  '/admin/developer/database-status': 'dev',
  '/admin/developer/system-test': 'dev',
  '/admin/developer/media-debug': 'dev',
  '/admin/developer/test-api': 'dev',
  '/admin/developer/api': 'dev',
  '/admin/developer/test-connection': 'dev',
  '/admin/developer/wordpress-migration': 'dev',
  '/admin/developer/digikala-import': 'dev',
  '/admin/developer/devops': 'dev',
  '/admin/developer/log-viewer': 'dev',
  '/admin/developer/code-editor': 'dev',
  '/admin/developer/database-explorer': 'dev',
  '/admin/developer/git-manager': 'dev',
  '/admin/developer/package-manager': 'dev',
  '/admin/developer/docker-manager': 'dev',
  // حذف مسیرهای غیرضروری دولوپر
  '/admin/developer/backup-manager': 'dev',
  '/admin/developer/monitoring-dashboard': 'dev',
  '/admin/developer/security-scanner': 'dev',
  '/admin/developer/performance-analyzer': 'dev',
  '/admin/developer/deployment-manager': 'dev',
}

// بررسی permission برای نمایش menu item ها - حالت پیش‌فرض
const canShowMenu = (_path: string): boolean => {
  // حالت پیش‌فرض: همه menu items نمایش داده می‌شوند
  return true
}

// تابع کمکی برای بررسی اینکه آیا کاربر در هیچ‌کدام از زیرمنوهای یک سرگروه دسترسی مشاهده دارد یا نه - حالت پیش‌فرض
function canShowSection(_permissionList: string[]): boolean {
  // حالت پیش‌فرض: همه sections نمایش داده می‌شوند
  return true
}
</script>

<style scoped>
/* استایل‌های مربوط به تو رفتگی و خط عمودی برای لینک‌های زیرمجموعه */
.collapsible-sub-links-wrapper {
  padding-right: 1rem;
  border-right: 1px solid #4a5568; /* رنگ خط را برای تم تیره تنظیم کردم */
  margin-right: 0.5rem;
  padding-bottom: 0.25rem;
  padding-top: 0.25rem;
}

/* رفع FOUC - جلوگیری از نمایش محتوا قبل از بارگذاری CSS */
.admin-sidebar {
  visibility: visible;
  opacity: 1;
  transition: opacity 0.1s ease-in-out;
}

/* Loading state برای سایدبار */
.admin-sidebar.loading {
  opacity: 0.8;
}


</style>

