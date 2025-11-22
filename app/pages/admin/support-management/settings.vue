<template>
  <div class="p-6" dir="rtl">
    <div class="mb-6 bg-white p-6 rounded-lg shadow-md border border-gray-200">
      <h1 class="text-2xl font-bold text-gray-900 mb-2">تنظیمات چت آنلاین و پشتیبانی</h1>
      <p class="text-gray-600">مدیریت دپارتمان‌های پشتیبانی</p>
    </div>

         <!-- منوی تب‌ها -->
     <div class="mb-8">
       <nav class="flex flex-wrap gap-2 border-b border-gray-200">
         <button
:class="['px-4 py-2 text-sm font-medium rounded-t-md focus:outline-none', activeTab === 'departments' ? 'bg-white border-x border-t border-b-0 border-gray-200 text-blue-600' : 'bg-gray-100 text-gray-600 hover:bg-gray-200']"
           @click="activeTab = 'departments'">
           دپارتمان بندی
         </button>
         <button
:class="['px-4 py-2 text-sm font-medium rounded-t-md focus:outline-none', activeTab === 'workingHours' ? 'bg-white border-x border-t border-b-0 border-gray-200 text-blue-600' : 'bg-gray-100 text-gray-600 hover:bg-gray-200']"
           @click="activeTab = 'workingHours'">
           زمان‌های کاری
         </button>
         <button
:class="['px-4 py-2 text-sm font-medium rounded-t-md focus:outline-none', activeTab === 'waitingMode' ? 'bg-white border-x border-t border-b-0 border-gray-200 text-blue-600' : 'bg-gray-100 text-gray-600 hover:bg-gray-200']"
           @click="activeTab = 'waitingMode'">
           حالت انتظار
         </button>
         <button
:class="['px-4 py-2 text-sm font-medium rounded-t-md focus:outline-none', activeTab === 'smartMessages' ? 'bg-white border-x border-t border-b-0 border-gray-200 text-blue-600' : 'bg-gray-100 text-gray-600 hover:bg-gray-200']"
           @click="activeTab = 'smartMessages'">
           هوشمند سازی
         </button>
         <button
:class="['px-4 py-2 text-sm font-medium rounded-t-md focus:outline-none', activeTab === 'formBuilder' ? 'bg-white border-x border-t border-b-0 border-gray-200 text-blue-600' : 'bg-gray-100 text-gray-600 hover:bg-gray-200']"
           @click="activeTab = 'formBuilder'">
           ایجاد فرم
         </button>
        <button
:class="['px-4 py-2 text-sm font-medium rounded-t-md focus:outline-none', activeTab === 'rateLimit' ? 'bg-white border-x border-t border-b-0 border-gray-200 text-blue-600' : 'bg-gray-100 text-gray-600 hover:bg-gray-200']"
          @click="activeTab = 'rateLimit'">
          نرخ پیام
        </button>
        <button
:class="['px-4 py-2 text-sm font-medium rounded-t-md focus:outline-none', activeTab === 'security' ? 'bg-white border-x border-t border-b-0 border-gray-200 text-blue-600' : 'bg-gray-100 text-gray-600 hover:bg-gray-200']"
          @click="activeTab = 'security'">
          امنیت چت
        </button>

       </nav>
     </div>

    <!-- محتوای تب فعال -->
    <div>
      <!-- دپارتمان بندی -->
      <div v-if="activeTab === 'departments'" class="space-y-6">
        <!-- هدر بنفش -->
        <div class="bg-purple-600 text-white p-6 rounded-lg">
          <h2 class="text-2xl font-bold mb-2">ایجاد تیم</h2>
          <p class="text-purple-100">
            با استفاده از قابلیت "ایجاد تیم"، می‌توانید تیم‌هایی با تخصص‌های متفاوت بسازید و برای هر کدام اپراتور مخصوص اختصاص دهید. به عنوان مثال، یک تیم برای مشکلات فنی یا یک تیم مالی تشکیل دهید.
          </p>
        </div>

        <!-- محتوای اصلی -->
        <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
          <div class="flex flex-col lg:flex-row gap-6">
            <!-- دکمه ایجاد تیم جدید -->
            <div class="lg:w-1/3">
              <button class="w-full bg-purple-600 text-white py-4 px-6 rounded-lg hover:bg-purple-700 transition-colors font-medium" @click="showCreateTeamModal = true">
                ایجاد تیم جدید
              </button>
            </div>

                         <!-- لیست تیم‌ها -->
             <div class="lg:w-2/3">
               <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                 <div v-for="team in teams" :key="team.id" class="bg-gray-50 rounded-lg p-6 border border-gray-200">
                   <div class="flex items-center justify-between">
                     <div class="flex items-center space-x-3 space-x-reverse">
                       <div class="w-12 h-12 rounded-full flex items-center justify-center" :style="{ backgroundColor: team.color }">
                         <span class="text-white text-lg">{{ team.icon }}</span>
                       </div>
                       <div>
                         <h3 class="font-medium text-gray-900">{{ team.name }}</h3>
                         <p class="text-sm text-gray-600">{{ team.operatorCount }} اپراتور</p>
                       </div>
                     </div>
                     <div class="flex space-x-2 space-x-reverse">
                       <button class="w-6 h-6 bg-blue-500 rounded flex items-center justify-center" @click="showEditTeamModal = true; currentTeam = team">
                         <span class="text-white text-xs">✏️</span>
                       </button>
                       <button class="w-6 h-6 bg-red-500 rounded flex items-center justify-center" @click="deleteTeam(team.id)">
                         <span class="text-white text-xs">🗑️</span>
                       </button>
                     </div>
                   </div>
                 </div>
               </div>
             </div>
          </div>
        </div>
             </div>
     </div>

     <!-- زمان‌های کاری -->
     <div v-if="activeTab === 'workingHours'" class="space-y-6">
       <!-- هدر بنفش -->
       <div class="bg-purple-600 text-white p-6 rounded-lg">
         <h2 class="text-2xl font-bold mb-2">زمان‌های کاری</h2>
         <p class="text-purple-100">
           مشخص کردن ساعات کاری برای آگاهی کاربران. در زمان‌های غیرکاری نیز کاربران از دست نمی‌روند. توجه داشته باشید که آنلاین بودن در تلگرام به معنای آنلاین بودن در سیستم ما نیست.
         </p>
       </div>

       <!-- محتوای اصلی -->
       <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
         <!-- بخش زمان‌های کاری -->
         <div class="mb-8">
           <h3 class="text-xl font-bold text-gray-900 mb-4">زمان های کاری</h3>
           
           <!-- تصویر ساعت -->
           <div class="flex justify-center mb-6">
             <div class="relative">
               <!-- ساعت اصلی -->
               <div class="w-32 h-32 bg-purple-600 rounded-full flex items-center justify-center relative">
                 <div class="w-24 h-24 bg-white rounded-full flex items-center justify-center">
                   <div class="w-1 h-8 bg-purple-600 rounded-full transform rotate-45 origin-bottom"></div>
                   <div class="w-1 h-6 bg-purple-600 rounded-full transform -rotate-45 origin-bottom absolute"></div>
                 </div>
               </div>
               
               <!-- ماه و ستاره -->
               <div class="absolute -top-6 -right-4 text-2xl">🌙</div>
               <div class="absolute -top-2 -right-8 text-lg">⭐</div>
               
               <!-- خورشید -->
               <div class="absolute -top-6 -left-4 text-2xl">☀️</div>
               
               <!-- شکل‌های هندسی -->
               <div class="absolute top-2 -right-12 w-3 h-3 bg-yellow-400 transform rotate-45"></div>
               <div class="absolute -bottom-2 -right-8 w-2 h-2 bg-purple-400 rounded"></div>
               <div class="absolute -bottom-4 -left-6 w-2 h-2 bg-yellow-400 transform rotate-45"></div>
             </div>
           </div>
           
           <!-- متن توضیحات -->
           <div class="text-center mb-6">
             <p class="text-gray-600 mb-4">
               در صورتی که زمان کاری تعریف نکرده باشید، کل هفته زمان‌های کاری شما محسوب می‌شود و تنظیمات زمان‌های غیر کاری اعمال نخواهند شد.
             </p>
             <button class="bg-purple-600 text-white px-6 py-3 rounded-lg hover:bg-purple-700 transition-colors font-medium" @click="showWorkingHoursModal = true">
               افزودن زمان کاری
             </button>
           </div>
         </div>

         <!-- بخش زمان‌های غیرکاری -->
         <div>
           <h3 class="text-xl font-bold text-gray-900 mb-4">زمان های غیرکاری</h3>
           
           <!-- تنظیمات -->
           <div class="space-y-4">
             <div class="flex items-center justify-between p-6 bg-gray-50 rounded-lg">
               <div class="flex items-center space-x-3 space-x-reverse">
                 <div class="w-6 h-6 bg-gray-300 rounded-full flex items-center justify-center">
                   <span class="text-white text-xs">?</span>
                 </div>
                 <span class="text-sm text-gray-700">فعال/غیرفعال کردن پیام آفلاین: دریافت اطلاعات کاربران برای تماس مجدد</span>
               </div>
               <div class="flex items-center space-x-2 space-x-reverse">
                 <button class="text-blue-600 hover:text-blue-700">
                   <span class="text-sm">✏️</span>
                 </button>
                 <label class="relative inline-flex items-center cursor-pointer">
                   <input v-model="offlineSettings.offlineMessage" type="checkbox" class="sr-only peer">
                   <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                 </label>
               </div>
             </div>

             <div class="flex items-center justify-between p-6 bg-gray-50 rounded-lg">
               <div class="flex items-center space-x-3 space-x-reverse">
                 <div class="w-6 h-6 bg-gray-300 rounded-full flex items-center justify-center">
                   <span class="text-white text-xs">?</span>
                 </div>
                 <span class="text-sm text-gray-700">تغییر به حالت آنلاین</span>
               </div>
               <label class="relative inline-flex items-center cursor-pointer">
                 <input v-model="offlineSettings.changeToOnline" type="checkbox" class="sr-only peer">
                 <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
               </label>
             </div>

             <div class="flex items-center justify-between p-6 bg-gray-50 rounded-lg">
               <div class="flex items-center space-x-3 space-x-reverse">
                 <div class="w-6 h-6 bg-gray-300 rounded-full flex items-center justify-center">
                   <span class="text-white text-xs">?</span>
                 </div>
                 <span class="text-sm text-gray-700">فعال/غیرفعال کردن ارسال پیام‌های هوشمند</span>
               </div>
               <label class="relative inline-flex items-center cursor-pointer">
                 <input v-model="offlineSettings.smartMessages" type="checkbox" class="sr-only peer">
                 <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
               </label>
             </div>

             <div class="flex items-center justify-between p-6 bg-gray-50 rounded-lg">
               <div class="flex items-center space-x-3 space-x-reverse">
                 <div class="w-6 h-6 bg-gray-300 rounded-full flex items-center justify-center">
                   <span class="text-white text-xs">?</span>
                 </div>
                 <span class="text-sm text-gray-700">ارسال پیام‌ها به تلگرام</span>
               </div>
               <label class="relative inline-flex items-center cursor-pointer">
                 <input v-model="offlineSettings.sendToTelegram" type="checkbox" class="sr-only peer">
                 <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
               </label>
             </div>

             <div class="flex items-center justify-between p-6 bg-gray-50 rounded-lg">
               <div class="flex items-center space-x-3 space-x-reverse">
                 <div class="w-6 h-6 bg-gray-300 rounded-full flex items-center justify-center">
                   <span class="text-white text-xs">?</span>
                 </div>
                 <span class="text-sm text-gray-700">تغییر آیکون ابزارک به حالت خواب</span>
               </div>
               <label class="relative inline-flex items-center cursor-pointer">
                 <input v-model="offlineSettings.sleepMode" type="checkbox" class="sr-only peer">
                 <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
               </label>
             </div>

             <div class="flex items-center justify-between p-6 bg-gray-50 rounded-lg">
               <div class="flex items-center space-x-3 space-x-reverse">
                 <div class="w-6 h-6 bg-gray-300 rounded-full flex items-center justify-center">
                   <span class="text-white text-xs">?</span>
                 </div>
                 <span class="text-sm text-gray-700">مخفی کردن ابزارک</span>
               </div>
               <label class="relative inline-flex items-center cursor-pointer">
                 <input v-model="offlineSettings.hideWidget" type="checkbox" class="sr-only peer">
                 <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
               </label>
             </div>

             <div class="flex items-center justify-between p-6 bg-gray-50 rounded-lg">
               <div class="flex items-center space-x-3 space-x-reverse">
                 <div class="w-6 h-6 bg-gray-300 rounded-full flex items-center justify-center">
                   <span class="text-white text-xs">?</span>
                 </div>
                 <span class="text-sm text-gray-700">غیرفعال کردن امکان چت</span>
               </div>
               <label class="relative inline-flex items-center cursor-pointer">
                 <input v-model="offlineSettings.disableChat" type="checkbox" class="sr-only peer">
                 <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
               </label>
             </div>
           </div>
         </div>
       </div>
     </div>

     <!-- حالت انتظار -->
     <div v-if="activeTab === 'waitingMode'" class="space-y-6">
       <!-- هدر بنفش -->
       <div class="bg-purple-600 text-white p-6 rounded-lg">
         <h2 class="text-2xl font-bold mb-2">حالت انتظار</h2>
         <p class="text-purple-100">
           در این حالت، کاربران در صف انتظار قرار می‌گیرند و پیام خودکار برای آن‌ها ارسال می‌شود. این حالت برای زمانی مناسب است که اپراتورها مشغول هستند یا خارج از ساعات کاری می‌باشند.
         </p>
       </div>

       <!-- محتوای اصلی -->
       <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
         <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
           <!-- ستون راست - مدیریت عملیات -->
           <div class="space-y-6">
             <h3 class="text-xl font-bold text-gray-900 mb-4">مدیریت عملیات</h3>
             
             <!-- دکمه ایجاد عملیات جدید -->
             <div class="bg-purple-600 text-white p-6 rounded-lg">
               <h4 class="text-lg font-semibold mb-2">عملیات تاخیر در پاسخگویی</h4>
               <p class="text-purple-100 mb-4">
                 حتی زمانی که تمامی اپراتورهای شما مشغول هستند همچنان به مشتریان خود توجه کنید
               </p>
               <button class="bg-white text-purple-600 px-6 py-3 rounded-lg hover:bg-purple-50 transition-colors font-medium" @click="openCreateOperationModal">
                 ایجاد عملیات جدید
               </button>
             </div>

             <!-- لیست عملیات -->
             <div class="bg-white border border-gray-200 rounded-lg p-6">
               <h4 class="text-lg font-semibold text-gray-900 mb-4">لیست عملیات ها</h4>
               
               <!-- لیست عملیات‌ها -->
               <div v-if="operations.length > 0" class="space-y-4">
                 <div v-for="operation in operations" :key="operation.id" class="bg-purple-600 text-white p-6 rounded-lg relative">
                   <div class="flex justify-between items-start">
                     <div class="flex-1">
                       <p class="text-center">{{ operation.content }}</p>
                       <div class="flex justify-center space-x-2 space-x-reverse mt-2">
                         <div class="w-3 h-3 bg-white rounded-full opacity-60"></div>
                         <div class="w-3 h-3 bg-white rounded-full opacity-60"></div>
                         <div class="w-3 h-3 bg-white rounded-full opacity-60"></div>
                       </div>
                     </div>
                     <div class="flex space-x-2 space-x-reverse">
                       <button class="text-white hover:text-purple-200 text-sm" @click="openEditOperationModal(operation)">✏️</button>
                       <button class="text-white hover:text-red-200 text-sm" @click="deleteOperation(operation.id)">✕</button>
                     </div>
                   </div>
                 </div>
               </div>

               <!-- پیام خالی -->
               <div v-if="operations.length === 0" class="text-center py-8">
                 <p class="text-gray-500 font-bold">هیچ عملیاتی تعریف نشده است</p>
               </div>
             </div>
           </div>
         </div>
       </div>
     </div>

     <!-- هوشمند سازی -->
     <div v-if="activeTab === 'smartMessages'" class="space-y-6">
       <!-- هدر بنفش -->
       <div class="bg-gradient-to-r from-purple-600 to-purple-700 text-white p-6 rounded-lg">
         <div class="flex justify-between items-start">
           <div class="flex-1">
             <h2 class="text-2xl font-bold mb-2">۱ پیام هوشمند</h2>
             <p class="text-purple-100">
               آیا به دنبال راهی سریع و هوشمند برای آغاز کردن مکالمات هستید؟ هوشمندسازی ابزارک شما این امکان را به شما میدهد!
             </p>
           </div>
           <div class="flex flex-col space-y-3 space-y-reverse">
             <button class="bg-white text-purple-600 px-6 py-3 rounded-lg hover:bg-purple-50 transition-colors font-medium" @click="showCreateSmartMessageModal = true">
               پیام هوشمند جدید
             </button>
             <a href="#" class="text-purple-100 hover:text-white text-sm flex items-center space-x-2 space-x-reverse">
               <span>راهنمای خوش آمدگویی</span>
               <span>❓</span>
             </a>
           </div>
         </div>
       </div>

       <!-- محتوای اصلی -->
       <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
         <!-- جدول پیام‌های هوشمند -->
         <div class="overflow-x-auto">
           <table class="w-full">
             <thead>
               <tr class="border-b border-gray-200">
                 <th class="text-right py-3 px-4 font-medium text-gray-900">عنوان</th>
                 <th class="text-right py-3 px-4 font-medium text-gray-900">وضعیت</th>
                 <th class="text-right py-3 px-4 font-medium text-gray-900">ارسال کننده</th>
                 <th class="text-right py-3 px-4 font-medium text-gray-900">آخرین به روزرسانی</th>
                 <th class="text-right py-3 px-4 font-medium text-gray-900">به روز رسانی توسط</th>
                 <th class="text-right py-3 px-4 font-medium text-gray-900">ویرایش و حذف</th>
               </tr>
             </thead>
             <tbody>
               <tr v-for="message in smartMessages" :key="message.id" class="border-b border-gray-100 hover:bg-gray-50">
                 <td class="py-4 px-4 text-gray-900">{{ message.title }}</td>
                 <td class="py-4 px-4">
                   <div class="flex items-center space-x-2 space-x-reverse">
                     <div class="relative">
                       <select v-model="message.status" class="appearance-none bg-transparent border-none text-sm focus:outline-none">
                         <option value="active">فعال</option>
                         <option value="inactive">غیرفعال</option>
                       </select>
                       <div class="absolute inset-y-0 left-0 flex items-center pointer-events-none">
                         <span class="text-gray-400">▼</span>
                       </div>
                     </div>
                     <div v-if="message.status === 'active'" class="w-2 h-2 bg-green-500 rounded-full"></div>
                   </div>
                 </td>
                 <td class="py-4 px-4">
                   <div class="flex items-center space-x-2 space-x-reverse">
                     <div class="relative">
                       <select v-model="message.sender" class="appearance-none bg-transparent border-none text-sm focus:outline-none">
                         <option value="random">انتخاب تصادفی</option>
                         <option value="specific">انتخاب خاص</option>
                       </select>
                       <div class="absolute inset-y-0 left-0 flex items-center pointer-events-none">
                         <span class="text-gray-400">▼</span>
                       </div>
                     </div>
                     <span v-if="message.sender === 'random'" class="text-gray-500">🔀</span>
                   </div>
                 </td>
                 <td class="py-4 px-4 text-gray-600 text-sm">{{ message.lastUpdate }}</td>
                 <td class="py-4 px-4">
                   <div class="flex items-center space-x-2 space-x-reverse">
                     <span class="text-gray-900 text-sm">{{ message.updatedBy.name }}</span>
                     <span class="text-gray-500 text-xs">{{ message.updatedBy.email }}</span>
                     <div class="w-6 h-6 bg-orange-400 rounded-full flex items-center justify-center">
                       <span class="text-white text-xs">👤</span>
                     </div>
                   </div>
                 </td>
                 <td class="py-4 px-4">
                   <div class="flex items-center space-x-2 space-x-reverse">
                     <button class="text-gray-500 hover:text-blue-600" @click="editSmartMessage(message)">
                       <span class="text-lg">✏️</span>
                     </button>
                     <button class="text-gray-500 hover:text-red-600" @click="deleteSmartMessage(message.id)">
                       <span class="text-lg">🗑️</span>
                     </button>
                   </div>
                 </td>
               </tr>
             </tbody>
           </table>
         </div>

         <!-- پیام خالی -->
         <div v-if="smartMessages.length === 0" class="text-center py-12">
           <div class="text-gray-400 text-6xl mb-4">💬</div>
           <p class="text-gray-500 font-medium">هیچ پیام هوشمندی تعریف نشده است</p>
           <p class="text-gray-400 text-sm mt-2">برای شروع، یک پیام هوشمند جدید ایجاد کنید</p>
         </div>
       </div>
     </div>

     <!-- مودال ایجاد/ویرایش پیام هوشمند -->
     <div v-if="showCreateSmartMessageModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
       <div class="bg-white rounded-lg w-full max-w-6xl mx-4 max-h-[90vh] overflow-hidden flex flex-col">
         <!-- هدر مودال -->
         <div class="bg-gradient-to-r from-purple-600 to-purple-700 text-white p-6">
           <div class="flex justify-between items-center">
             <button class="text-white hover:text-purple-200 text-sm" @click="closeSmartMessageModal">
               بازگشت
             </button>
             <div class="flex items-center space-x-4 space-x-reverse">
               <h2 class="text-xl font-bold">{{ currentSmartMessage.title || 'شرط ۲' }}</h2>
               <input type="checkbox" class="rounded border-white text-purple-600 focus:ring-purple-500">
             </div>
           </div>
           
           <!-- تب‌های ناوبری -->
           <div class="mt-4">
             <nav class="flex space-x-1 space-x-reverse">
               <button
:class="['px-4 py-2 text-sm font-medium rounded-t-md focus:outline-none', activeSmartMessageTab === 'content' ? 'bg-white text-purple-600' : 'text-purple-200 hover:text-white']" 
                 @click="activeSmartMessageTab = 'content'">
                 محتوا
               </button>
               <button
:class="['px-4 py-2 text-sm font-medium rounded-t-md focus:outline-none', activeSmartMessageTab === 'condition' ? 'bg-white text-purple-600' : 'text-purple-200 hover:text-white']" 
                 @click="activeSmartMessageTab = 'condition'">
                 شرط
               </button>
               <button
:class="['px-4 py-2 text-sm font-medium rounded-t-md focus:outline-none', activeSmartMessageTab === 'publishTime' ? 'bg-white text-purple-600' : 'text-purple-200 hover:text-white']" 
                 @click="activeSmartMessageTab = 'publishTime'">
                 زمان انتشار
               </button>
             </nav>
           </div>
         </div>

         <!-- محتوای اصلی -->
         <div class="flex-1 flex overflow-hidden">
           <!-- بخش چپ - پیش‌نمایش چت -->
           <div class="w-1/2 bg-gray-50 p-6 flex flex-col">
             <!-- پیش‌نمایش چت -->
             <div class="flex-1 flex items-center justify-center">
               <div class="relative">
                 <!-- دایره سفید بزرگ -->
                 <div class="w-80 h-80 bg-white rounded-full shadow-lg flex items-center justify-center relative">
                   <!-- حباب پیام چت -->
                   <div class="bg-purple-100 p-6 rounded-lg max-w-48 mb-4">
                     <p class="text-sm text-gray-800">{{ smartMessageForm.content || 'به سایت ما خوش اومدین، در خدمتتون هستیم 👋' }}</p>
                   </div>
                   
                   <!-- فیلد ورودی -->
                   <div class="absolute bottom-8 left-1/2 transform -translate-x-1/2 w-64">
                     <div class="bg-white border border-gray-300 rounded-lg p-3 flex items-center space-x-2 space-x-reverse">
                       <span class="text-gray-400">💬</span>
                       <span class="text-gray-400">📎</span>
                       <span class="text-gray-400">@</span>
                       <input v-model="smartMessageForm.content" type="text" placeholder="اینجا تایپ کنید." class="flex-1 text-sm focus:outline-none">
                     </div>
                   </div>
                 </div>
                 
                 <!-- آیکون چت بنفش -->
                 <div class="absolute -right-8 -top-8 w-16 h-16 bg-purple-600 rounded-full flex items-center justify-center">
                   <span class="text-white text-2xl">💬</span>
                 </div>
               </div>
             </div>
             
             <!-- خلاصه تنظیمات -->
             <div class="mt-6 space-y-2">
               <div class="flex justify-between text-sm">
                 <span class="text-gray-600">محتوا:</span>
                 <span class="text-gray-900">{{ getDisplayTypeText(smartMessageForm.displayType) }}</span>
               </div>
               <div class="flex justify-between text-sm">
                 <span class="text-gray-600">شرط:</span>
                 <span class="text-gray-900">{{ smartMessageForm.conditions.length > 0 ? 'تعریف شده' : 'تعریف نشده' }}</span>
               </div>
               <div class="flex justify-between text-sm">
                 <span class="text-gray-600">زمان انتشار:</span>
                 <span class="text-gray-900">{{ getPublishTimeText() }}</span>
               </div>
             </div>
           </div>

           <!-- بخش راست - فرم تنظیمات -->
           <div class="w-1/2 p-6 overflow-y-auto">
             <!-- تب محتوا -->
             <div v-if="activeSmartMessageTab === 'content'" class="space-y-6">
               <!-- نوع عملیات -->
               <div>
                 <label class="block text-sm font-medium text-gray-700 mb-2">این عملیات انجام شود:</label>
                 <select v-model="smartMessageForm.operationType" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500">
                   <option value="message_only">پیام ارسال شود اما ابزارک باز نشود (نمایش کوتاه پیام)</option>
                   <option value="message_and_widget">پیام ارسال شود و ابزارک باز شود</option>
                   <option value="widget_only">فقط ابزارک باز شود</option>
                 </select>
               </div>

               <!-- متن پیام -->
               <div>
                 <label class="block text-sm font-medium text-gray-700 mb-2">متن پیام</label>
                 <textarea 
                   v-model="smartMessageForm.content" 
                   rows="6" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
                   placeholder="متن پیام هوشمند خود را وارد کنید..."
                 ></textarea>
                 
                 <!-- نوار ابزار فرمت -->
                 <div class="flex items-center space-x-2 space-x-reverse mt-2 p-2 bg-gray-50 rounded">
                   <button class="p-1 hover:bg-gray-200 rounded font-bold">B</button>
                   <button class="p-1 hover:bg-gray-200 rounded italic">I</button>
                   <button class="p-1 hover:bg-gray-200 rounded underline">U</button>
                   <button class="p-1 hover:bg-gray-200 rounded">☰</button>
                 </div>
                 
                 <!-- شمارنده کاراکتر -->
                 <div class="text-sm text-gray-500 mt-2">
                   کاراکتر باقی مانده: {{ 10000 - (smartMessageForm.content?.length || 0) }}
                 </div>
               </div>

               <!-- ارسال به عنوان -->
               <div>
                 <label class="block text-sm font-medium text-gray-700 mb-2">ارسال به عنوان</label>
                 <div class="space-y-3">
                   <div class="flex items-center space-x-3 space-x-reverse">
                     <input v-model="smartMessageForm.displayType" type="radio" value="full" class="text-purple-600 focus:ring-purple-500">
                     <div>
                       <div class="font-medium text-gray-900">پیام کامل</div>
                       <div class="text-sm text-gray-600">پیام به صورت کامل نمایش داده می‌شود.</div>
                     </div>
                   </div>
                   <div class="flex items-center space-x-3 space-x-reverse">
                     <input v-model="smartMessageForm.displayType" type="radio" value="short" class="text-purple-600 focus:ring-purple-500">
                     <div>
                       <div class="font-medium text-gray-900">پیام کوتاه</div>
                       <div class="text-sm text-gray-600">خلاصه‌ای از متن پیام نمایش داده می‌شود.</div>
                     </div>
                   </div>
                   <div class="flex items-center space-x-3 space-x-reverse">
                     <input v-model="smartMessageForm.displayType" type="radio" value="notification" class="text-purple-600 focus:ring-purple-500">
                     <div>
                       <div class="font-medium text-gray-900">اعلان</div>
                       <div class="text-sm text-gray-600">نمایش دایره کوچک در بالای ابزارک</div>
                     </div>
                   </div>
                 </div>
               </div>

               <!-- ارسال پیام از سمت -->
               <div>
                 <label class="block text-sm font-medium text-gray-700 mb-2">ارسال پیام از سمت</label>
                 <div class="flex items-center space-x-2 space-x-reverse">
                   <select v-model="smartMessageForm.sender" class="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500">
                     <option value="random">انتخاب تصادفی</option>
                     <option value="specific">انتخاب خاص</option>
                   </select>
                   <button class="p-2 text-gray-500 hover:text-gray-700">🔄</button>
                 </div>
               </div>

               <!-- تنظیمات اضافی -->
               <div class="space-y-4">
                 <div class="flex items-center justify-between">
                   <span class="text-sm font-medium text-gray-700">برای هر بازدید کننده یک بار ارسال</span>
                   <label class="relative inline-flex items-center cursor-pointer">
                     <input v-model="smartMessageForm.sendOncePerVisitor" type="checkbox" class="sr-only peer">
                     <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-purple-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-purple-600"></div>
                   </label>
                 </div>
                 
                 <div class="flex items-center justify-between">
                   <span class="text-sm font-medium text-gray-700">ارسال به همراه صدا</span>
                   <label class="relative inline-flex items-center cursor-pointer">
                     <input v-model="smartMessageForm.withSound" type="checkbox" class="sr-only peer">
                     <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-purple-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-purple-600"></div>
                   </label>
                 </div>
                 
                 <div class="flex items-center justify-between">
                   <span class="text-sm font-medium text-gray-700">در موبایل نیز ارسال شود</span>
                   <label class="relative inline-flex items-center cursor-pointer">
                     <input v-model="smartMessageForm.sendOnMobile" type="checkbox" class="sr-only peer">
                     <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-purple-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-purple-600"></div>
                   </label>
                 </div>
               </div>
             </div>

             <!-- تب شرط -->
             <div v-if="activeSmartMessageTab === 'condition'" class="space-y-6">
               <div>
                 <label class="block text-sm font-medium text-gray-700 mb-2">شرط</label>
                 <p class="text-sm text-gray-600 mb-4">در چه شرایطی ارسال شود</p>
                 
                 <select v-model="smartMessageForm.conditionType" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 mb-4">
                   <option value="at_least_one">حداقل یکی از شرط های زیر رخ دهد</option>
                   <option value="all">تمام شرط های زیر رخ دهد</option>
                 </select>

                 <!-- لیست شرط‌ها -->
                 <div class="space-y-3">
                   <div v-for="(condition, index) in smartMessageForm.conditions" :key="index" class="flex items-center space-x-2 space-x-reverse p-3 bg-gray-50 rounded-lg">
                     <button class="text-red-500 hover:text-red-700" @click="removeCondition(index)">✕</button>
                     <select v-model="condition.type" class="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500">
                       <option value="time_on_page">زمان در صفحه</option>
                       <option value="scroll_percentage">درصد اسکرول</option>
                       <option value="page_visit">بازدید صفحه</option>
                     </select>
                     <input v-model="condition.value" type="number" class="w-20 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500">
                     <span class="text-sm text-gray-600">{{ getConditionUnit(condition.type) }}</span>
                   </div>
                 </div>

                 <!-- دکمه افزودن شرط -->
                 <button class="flex items-center space-x-2 space-x-reverse text-purple-600 hover:text-purple-700 text-sm" @click="addCondition">
                   <span>+</span>
                   <span>افزودن شرط</span>
                 </button>
               </div>
             </div>

             <!-- تب زمان انتشار -->
             <div v-if="activeSmartMessageTab === 'publishTime'" class="space-y-6">
               <!-- شروع به ارسال -->
               <div>
                 <label class="block text-sm font-medium text-gray-700 mb-2">شروع به ارسال</label>
                 <div class="grid grid-cols-2 gap-6">
                   <input v-model="smartMessageForm.startDate" type="date" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500">
                   <input v-model="smartMessageForm.startTime" type="time" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500">
                 </div>
               </div>

               <!-- پایان ارسال -->
               <div>
                 <label class="block text-sm font-medium text-gray-700 mb-2">پایان ارسال</label>
                 <div class="space-y-3">
                   <div class="flex items-center space-x-3 space-x-reverse">
                     <input v-model="smartMessageForm.endType" type="radio" value="never" class="text-purple-600 focus:ring-purple-500">
                     <span class="text-sm text-gray-700">هیچ وقت</span>
                   </div>
                   <div class="flex items-center space-x-3 space-x-reverse">
                     <input v-model="smartMessageForm.endType" type="radio" value="specific" class="text-purple-600 focus:ring-purple-500">
                     <span class="text-sm text-gray-700">تنظیم زمان خاص</span>
                   </div>
                 </div>
                 
                 <div v-if="smartMessageForm.endType === 'specific'" class="mt-4 grid grid-cols-2 gap-6">
                   <input v-model="smartMessageForm.endDate" type="date" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500">
                   <input v-model="smartMessageForm.endTime" type="time" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500">
                 </div>
               </div>

               <!-- تنظیم روز و ساعت -->
               <div>
                 <label class="block text-sm font-medium text-gray-700 mb-2">تنظیم روز و ساعتهایی که این چت میتواند ارسال شود</label>
                 <div class="space-y-3">
                   <div class="flex items-center space-x-3 space-x-reverse">
                     <input v-model="smartMessageForm.scheduleType" type="radio" value="always" class="text-purple-600 focus:ring-purple-500">
                     <span class="text-sm text-gray-700">هر روز هر ساعتی</span>
                   </div>
                   <div class="flex items-center space-x-3 space-x-reverse">
                     <input v-model="smartMessageForm.scheduleType" type="radio" value="working_hours" class="text-purple-600 focus:ring-purple-500">
                     <span class="text-sm text-gray-700">در طول زمانهای کاری</span>
                   </div>
                   <div class="flex items-center space-x-3 space-x-reverse">
                     <input v-model="smartMessageForm.scheduleType" type="radio" value="non_working_hours" class="text-purple-600 focus:ring-purple-500">
                     <span class="text-sm text-gray-700">در طول زمانهای غیرکاری</span>
                   </div>
                 </div>
               </div>
             </div>
           </div>
         </div>

         <!-- فوتر -->
         <div class="border-t border-gray-200 p-6">
           <div class="flex justify-end space-x-4 space-x-reverse">
             <button class="px-6 py-2 text-sm font-medium text-purple-600 bg-purple-50 border border-purple-200 rounded-md hover:bg-purple-100" @click="closeSmartMessageModal">
               انصراف
             </button>
             <button class="px-6 py-2 text-sm font-medium text-white bg-purple-600 border border-transparent rounded-md hover:bg-purple-700" @click="saveSmartMessage">
               انتشار
             </button>
           </div>
         </div>
       </div>
     </div>

     <!-- ایجاد فرم -->
     <div v-if="activeTab === 'formBuilder'" class="space-y-6">
       <!-- هدر بنفش -->
       <div class="bg-gradient-to-r from-purple-600 to-purple-700 text-white p-6 rounded-lg">
         <div class="flex justify-between items-start">
           <div class="flex-1">
             <h2 class="text-2xl font-bold mb-2">صفحه ایجاد فرم ساعات کاری و غیر کاری</h2>
             <p class="text-purple-100">
               طراحی فرم اطلاعات آغاز مکالمه بر اساس نیازهای خودتان یا جمع آوری لیدهای خود با استفاده از رایچت
             </p>
           </div>
           <div class="flex flex-col space-y-3 space-y-reverse">
             <button class="bg-white text-purple-600 px-6 py-3 rounded-lg hover:bg-purple-50 transition-colors font-medium" @click="openCreateFormModal">
               ایجاد فرم جدید
             </button>
             <a href="#" class="text-purple-100 hover:text-white text-sm flex items-center space-x-2 space-x-reverse">
               <span>راهنمای ایجاد فرم</span>
               <span>❓</span>
             </a>
           </div>
         </div>
       </div>

       <!-- محتوای اصلی -->
       <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
         <!-- جدول فرم‌ها -->
         <div class="overflow-x-auto">
           <table class="w-full">
             <thead>
               <tr class="border-b border-gray-200">
                 <th class="text-right py-3 px-4 font-medium text-gray-900">عنوان فرم</th>
                 <th class="text-right py-3 px-4 font-medium text-gray-900">نوع فرم</th>
                 <th class="text-right py-3 px-4 font-medium text-gray-900">وضعیت</th>
                 <th class="text-right py-3 px-4 font-medium text-gray-900">تعداد فیلدها</th>
                 <th class="text-right py-3 px-4 font-medium text-gray-900">آخرین به روزرسانی</th>
                 <th class="text-right py-3 px-4 font-medium text-gray-900">عملیات</th>
               </tr>
             </thead>
             <tbody>
               <tr v-for="form in forms" :key="form.id" class="border-b border-gray-100 hover:bg-gray-50">
                 <td class="py-4 px-4 text-gray-900">{{ form.title }}</td>
                 <td class="py-4 px-4 text-gray-600">{{ form.type }}</td>
                 <td class="py-4 px-4">
                   <div class="flex items-center space-x-2 space-x-reverse">
                     <div class="relative">
                       <select v-model="form.status" class="appearance-none bg-transparent border-none text-sm focus:outline-none">
                         <option value="active">فعال</option>
                         <option value="inactive">غیرفعال</option>
                       </select>
                       <div class="absolute inset-y-0 left-0 flex items-center pointer-events-none">
                         <span class="text-gray-400">▼</span>
                       </div>
                     </div>
                     <div v-if="form.status === 'active'" class="w-2 h-2 bg-green-500 rounded-full"></div>
                   </div>
                 </td>
                 <td class="py-4 px-4 text-gray-600">{{ form.fieldCount }} فیلد</td>
                 <td class="py-4 px-4 text-gray-600 text-sm">{{ form.lastUpdate }}</td>
                 <td class="py-4 px-4">
                   <div class="flex items-center space-x-2 space-x-reverse">
                     <button class="text-gray-500 hover:text-blue-600" @click="editForm(form)">
                       <span class="text-lg">✏️</span>
                     </button>
                     <button class="text-gray-500 hover:text-red-600" @click="deleteForm(form.id)">
                       <span class="text-lg">🗑️</span>
                     </button>
                   </div>
                 </td>
               </tr>
             </tbody>
           </table>
         </div>

         <!-- پیام خالی -->
         <div v-if="forms.length === 0" class="text-center py-12">
           <div class="text-gray-400 text-6xl mb-4">📋</div>
           <p class="text-gray-500 font-medium">هیچ فرمی تعریف نشده است</p>
           <p class="text-gray-400 text-sm mt-2">برای شروع، یک فرم جدید ایجاد کنید</p>
         </div>
       </div>
     </div>

     <!-- مودال افزودن تیم جدید -->
     <div v-if="showCreateTeamModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
       <div class="bg-white rounded-lg p-6 w-full max-w-2xl mx-4 max-h-[90vh] overflow-y-auto">
         <div class="flex justify-between items-center mb-6">
           <h2 class="text-xl font-bold text-gray-900">افزودن تیم جدید</h2>
           <button class="text-gray-400 hover:text-gray-600" @click="showCreateTeamModal = false">
             <span class="text-2xl">×</span>
           </button>
         </div>

         <div class="space-y-6">
           <!-- اطلاعات تیم -->
           <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
             <div>
               <label class="block text-sm font-medium text-gray-700 mb-2">نام تیم</label>
               <input v-model="newTeam.name" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500" placeholder="نام تیم">
             </div>
             <div>
               <label class="block text-sm font-medium text-gray-700 mb-2">رنگ تیم</label>
               <input v-model="newTeam.color" type="color" class="w-full h-10 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500">
             </div>
           </div>

           <!-- لوگوی تیم -->
           <div class="flex items-center space-x-4 space-x-reverse">
             <div class="w-16 h-16 bg-purple-600 rounded-full flex items-center justify-center">
               <span class="text-white text-2xl">💬</span>
             </div>
             <div class="space-y-2">
               <button class="text-purple-600 hover:text-purple-700 text-sm">بارگذاری لوگو</button>
               <button class="text-red-600 hover:text-red-700 text-sm">حذف تصویر</button>
             </div>
           </div>

           <!-- افزودن اپراتور -->
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">افزودن اپراتور</label>
             <div class="relative">
               <input v-model="operatorSearch" type="text" class="w-full px-3 py-2 pr-10 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500" placeholder="اپراتورهای خود را به این تیم اضافه کنید">
               <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                 <span class="text-gray-400">🔍</span>
               </div>
             </div>
           </div>

           <!-- لیست اپراتورها -->
           <div v-if="selectedOperators.length > 0">
             <h3 class="text-sm font-medium text-gray-700 mb-2">اپراتورها</h3>
             <div class="space-y-2">
               <div v-for="operator in selectedOperators" :key="operator.id" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                 <div class="flex items-center space-x-3 space-x-reverse">
                   <div class="w-8 h-8 bg-orange-400 rounded-full flex items-center justify-center">
                     <span class="text-white text-sm">👤</span>
                   </div>
                   <div>
                     <p class="text-sm font-medium text-gray-900">{{ operator.name }}</p>
                     <p class="text-xs text-gray-600">{{ operator.email }}</p>
                   </div>
                 </div>
                 <button class="text-red-600 hover:text-red-700 text-sm" @click="removeOperator(operator.id)">حذف</button>
               </div>
             </div>
           </div>
         </div>

         <!-- دکمه‌های عملیات -->
         <div class="flex justify-end space-x-4 space-x-reverse mt-6">
           <button class="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-100 border border-gray-300 rounded-md hover:bg-gray-200" @click="showCreateTeamModal = false">
             انصراف
           </button>
           <button class="px-4 py-2 text-sm font-medium text-white bg-purple-600 border border-transparent rounded-md hover:bg-purple-700" @click="createTeam">
             تایید
           </button>
         </div>
       </div>
     </div>

     <!-- مودال ویرایش تیم -->
     <div v-if="showEditTeamModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
       <div class="bg-white rounded-lg p-6 w-full max-w-2xl mx-4 max-h-[90vh] overflow-y-auto">
         <div class="flex justify-between items-center mb-6">
           <h2 class="text-xl font-bold text-gray-900">ویرایش تیم</h2>
           <button class="text-gray-400 hover:text-gray-600" @click="showEditTeamModal = false">
             <span class="text-2xl">×</span>
           </button>
         </div>

         <div class="space-y-6">
           <!-- اطلاعات تیم -->
           <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
             <div>
               <label class="block text-sm font-medium text-gray-700 mb-2">نام تیم</label>
               <input v-model="currentTeam.name" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500" placeholder="نام تیم">
             </div>
             <div>
               <label class="block text-sm font-medium text-gray-700 mb-2">رنگ تیم</label>
               <input v-model="currentTeam.color" type="color" class="w-full h-10 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500">
             </div>
           </div>

           <!-- لوگوی تیم -->
           <div class="flex items-center space-x-4 space-x-reverse">
             <div class="w-16 h-16 rounded-full flex items-center justify-center" :style="{ backgroundColor: currentTeam.color }">
               <span class="text-white text-2xl">{{ currentTeam.icon }}</span>
             </div>
             <div class="space-y-2">
               <button class="text-purple-600 hover:text-purple-700 text-sm">بارگذاری لوگو</button>
               <button class="text-red-600 hover:text-red-700 text-sm">حذف تصویر</button>
             </div>
           </div>

           <!-- افزودن اپراتور -->
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">افزودن اپراتور</label>
             <div class="relative">
               <input v-model="operatorSearch" type="text" class="w-full px-3 py-2 pr-10 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500" placeholder="اپراتورهای خود را به این تیم اضافه کنید">
               <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                 <span class="text-gray-400">🔍</span>
               </div>
             </div>
           </div>

           <!-- لیست اپراتورها -->
           <div v-if="currentTeam.operators && currentTeam.operators.length > 0">
             <h3 class="text-sm font-medium text-gray-700 mb-2">اپراتورها</h3>
             <div class="space-y-2">
               <div v-for="operator in currentTeam.operators" :key="operator.id" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                 <div class="flex items-center space-x-3 space-x-reverse">
                   <div class="w-8 h-8 bg-orange-400 rounded-full flex items-center justify-center">
                     <span class="text-white text-sm">👤</span>
                   </div>
                   <div>
                     <p class="text-sm font-medium text-gray-900">{{ operator.name }}</p>
                     <p class="text-xs text-gray-600">{{ operator.email }}</p>
                   </div>
                 </div>
                 <button class="text-red-600 hover:text-red-700 text-sm" @click="removeOperatorFromTeam(operator.id)">حذف</button>
               </div>
             </div>
           </div>
         </div>

         <!-- دکمه‌های عملیات -->
         <div class="flex justify-end space-x-4 space-x-reverse mt-6">
           <button class="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-100 border border-gray-300 rounded-md hover:bg-gray-200" @click="showEditTeamModal = false">
             انصراف
           </button>
           <button class="px-4 py-2 text-sm font-medium text-white bg-purple-600 border border-transparent rounded-md hover:bg-purple-700" @click="updateTeam">
             تایید
           </button>
         </div>
       </div>
     </div>

     <!-- مودال ایجاد/ویرایش عملیات جدید -->
     <div v-if="showCreateOperationModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
       <div class="bg-white rounded-lg p-6 w-full max-w-6xl mx-4 max-h-[90vh] overflow-y-auto">
         <div class="flex justify-between items-center mb-6">
           <h2 class="text-xl font-bold text-gray-900">{{ isEditingOperation ? 'ویرایش عملیات' : 'ایجاد عملیات جدید' }}</h2>
           <button class="text-gray-400 hover:text-gray-600" @click="closeOperationModal">
             <span class="text-2xl">×</span>
           </button>
         </div>

         <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
           <!-- ستون اصلی - فرم عملیات -->
           <div class="lg:col-span-2 space-y-6">
             <!-- ارسال در مکالمات -->
             <div>
               <label class="block text-sm font-medium text-gray-700 mb-2">ارسال در مکالمات</label>
               <select v-model="operationForm.conversationType" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500">
                 <option value="awaiting">در انتظار پاسخ</option>
                 <option value="active">مکالمات فعال</option>
                 <option value="all">تمام مکالمات</option>
               </select>
             </div>

             <!-- نوع اعمال عملیات -->
             <div>
               <label class="block text-sm font-medium text-gray-700 mb-2">نوع اعمال عملیات</label>
               <select v-model="operationForm.operationType" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500">
                 <option value="all_departments">اعمال بر روی تمام دپارتمان ها</option>
                 <option value="specific_departments">دپارتمان های خاص</option>
                 <option value="specific_teams">تیم های خاص</option>
               </select>
             </div>

             <!-- ارسال به همراه صدا -->
             <div class="flex items-center justify-between">
               <span class="text-sm font-medium text-gray-700">ارسال به همراه صدا</span>
               <label class="relative inline-flex items-center cursor-pointer">
                 <input v-model="operationForm.withSound" type="checkbox" class="sr-only peer">
                 <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-purple-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-purple-600"></div>
               </label>
             </div>

             <!-- زمان انتظار -->
             <div>
               <label class="block text-sm font-medium text-gray-700 mb-2">زمان انتظار</label>
               <div class="flex items-center space-x-2 space-x-reverse">
                 <input 
                   v-model="operationForm.waitTime" 
                   type="number" 
                   min="0"
                   class="w-20 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
                 >
                 <span class="text-sm text-gray-700">ثانیه</span>
               </div>
             </div>

             <!-- متن پیام قبل از پاسخ -->
             <div>
               <label class="block text-sm font-medium text-gray-700 mb-2">متن پیام قبل از پاسخ</label>
               
               <!-- پیش‌نمایش پیام -->
               <div class="mb-2 p-3 bg-gray-50 rounded-lg border">
                 <p class="text-sm text-gray-700">{{ operationForm.messageText || 'به سایت ما خوش اومدین، در خدمتتون هستیم 👋' }}</p>
               </div>
               
               <textarea 
                 v-model="operationForm.messageText" 
                 rows="6" 
                 class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
                 placeholder="پیام خودکار برای کاربران در حالت انتظار"
               ></textarea>
               
               <!-- نوار ابزار فرمت -->
               <div class="flex items-center space-x-2 space-x-reverse mt-2 p-2 bg-gray-50 rounded">
                 <button class="p-1 hover:bg-gray-200 rounded">🖼️</button>
                 <button class="p-1 hover:bg-gray-200 rounded font-bold">B</button>
                 <button class="p-1 hover:bg-gray-200 rounded italic">I</button>
                 <button class="p-1 hover:bg-gray-200 rounded underline">U</button>
                 <button class="p-1 hover:bg-gray-200 rounded">☰</button>
               </div>
               
               <!-- شمارنده کاراکتر -->
               <div class="text-sm text-gray-500 mt-2">
                 کاراکتر باقی مانده: {{ 500 - (operationForm.messageText?.length || 0) }}
               </div>
             </div>

             <!-- ایجاد فرم -->
             <div>
               <label class="block text-sm font-medium text-gray-700 mb-2">ایجاد فرم</label>
               <button class="flex items-center space-x-2 space-x-reverse px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors">
                 <span class="text-lg">+</span>
                 <span>اضافه کردن فیلد جدید</span>
               </button>
             </div>
           </div>

           <!-- ستون کناری - پیش‌نمایش -->
           <div class="space-y-6">
             <!-- دکمه لیست عملیات -->
             <button class="w-full bg-purple-600 text-white py-3 px-4 rounded-lg hover:bg-purple-700 transition-colors font-medium">
               لیست عملیات ها
             </button>

             <!-- کارت پیش‌نمایش -->
             <div class="bg-purple-600 text-white p-6 rounded-lg relative">
               <button class="absolute top-2 left-2 text-white hover:text-purple-200">✕</button>
               <p class="text-center">پاسخگوی سوالات شما هستیم.</p>
               <div class="flex justify-center space-x-2 space-x-reverse mt-2">
                 <div class="w-3 h-3 bg-white rounded-full opacity-60"></div>
                 <div class="w-3 h-3 bg-white rounded-full opacity-60"></div>
                 <div class="w-3 h-3 bg-white rounded-full opacity-60"></div>
               </div>
             </div>

             <!-- پیش‌نمایش پیام -->
             <div class="bg-gray-50 p-6 rounded-lg">
               <p class="text-gray-700 text-sm">{{ operationForm.messageText || 'به سایت ما خوش اومدین، در خدمتتون هستیم 👋' }}</p>
             </div>
           </div>
         </div>

         <!-- دکمه‌های عملیات -->
         <div class="flex justify-end space-x-4 space-x-reverse mt-8">
           <button class="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-100 border border-gray-300 rounded-md hover:bg-gray-200" @click="closeOperationModal">
             انصراف
           </button>
           <button class="px-4 py-2 text-sm font-medium text-white bg-purple-600 border border-transparent rounded-md hover:bg-purple-700" @click="saveOperation">
             انتشار
           </button>
         </div>
       </div>
     </div>

     <!-- دکمه‌های عملیات -->
     <div class="mt-8 flex justify-end space-x-4 space-x-reverse">
       <button class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500" @click="resetForm">
         بازنشانی
       </button>
       <button class="px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500" @click="saveAllSettings">
         ذخیره تنظیمات
       </button>
     </div>
   </div>

   <!-- مودال ایجاد/ویرایش فرم -->
   <div v-if="showCreateFormModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
     <div class="bg-white rounded-lg w-full max-w-6xl mx-4 max-h-[90vh] overflow-hidden flex flex-col">
       <!-- هدر مودال -->
       <div class="bg-gradient-to-r from-purple-600 to-purple-700 text-white p-6">
         <div class="flex justify-between items-center">
           <button class="text-white hover:text-purple-200 text-sm" @click="closeFormModal">
             بازگشت
           </button>
           <div class="flex items-center space-x-4 space-x-reverse">
             <h2 class="text-xl font-bold">{{ isEditingForm ? 'ویرایش فرم' : 'ایجاد فرم جدید' }}</h2>
             <input v-model="formBuilderForm.isDisabled" type="checkbox" class="rounded border-white text-purple-600 focus:ring-purple-500">
           </div>
         </div>
       </div>

       <!-- محتوای اصلی -->
       <div class="flex-1 flex overflow-hidden">
         <!-- بخش چپ - پیش‌نمایش فرم -->
         <div class="w-1/2 bg-gray-50 p-6 flex flex-col">
           <!-- پیش‌نمایش فرم -->
           <div class="flex-1 flex items-center justify-center">
             <div class="relative">
               <!-- دایره سفید بزرگ -->
               <div class="w-80 h-80 bg-white rounded-full shadow-lg flex items-center justify-center relative">
                 <!-- فرم پیش‌نمایش -->
                 <div class="bg-white p-6 rounded-lg max-w-48 mb-4 border border-gray-200">
                   <h3 class="text-sm font-medium text-gray-900 mb-3">{{ formBuilderForm.title || 'عنوان فرم' }}</h3>
                   <p class="text-xs text-gray-600 mb-4">{{ formBuilderForm.content || 'توضیحات فرم' }}</p>
                   
                   <!-- فیلدهای فرم -->
                   <div class="space-y-3">
                     <div v-for="field in formBuilderForm.fields" :key="field.id" class="space-y-1">
                       <label class="text-xs text-gray-700">{{ field.label }}{{ field.required ? ' *' : '' }}</label>
                       <input v-if="field.type === 'text'" type="text" :placeholder="field.label" class="w-full text-xs p-2 border border-gray-300 rounded">
                       <input v-else-if="field.type === 'email'" type="email" :placeholder="field.label" class="w-full text-xs p-2 border border-gray-300 rounded">
                       <input v-else-if="field.type === 'phone'" type="tel" :placeholder="field.label" class="w-full text-xs p-2 border border-gray-300 rounded">
                       <textarea v-else-if="field.type === 'textarea'" :placeholder="field.label" class="w-full text-xs p-2 border border-gray-300 rounded" rows="3"></textarea>
                     </div>
                   </div>
                   
                   <button class="w-full mt-4 bg-purple-600 text-white text-xs py-2 rounded hover:bg-purple-700">
                     ارسال
                   </button>
                 </div>
                 
                 <!-- فیلد ورودی -->
                 <div class="absolute bottom-8 left-1/2 transform -translate-x-1/2 w-64">
                   <div class="bg-white border border-gray-300 rounded-lg p-3 flex items-center space-x-2 space-x-reverse">
                     <span class="text-gray-400">💬</span>
                     <span class="text-gray-400">📎</span>
                     <span class="text-gray-400">@</span>
                     <input type="text" placeholder="اینجا تایپ کنید." class="flex-1 text-sm focus:outline-none">
                   </div>
                 </div>
               </div>
               
               <!-- آیکون فرم بنفش -->
               <div class="absolute -right-8 -top-8 w-16 h-16 bg-purple-600 rounded-full flex items-center justify-center">
                 <span class="text-white text-2xl">📋</span>
               </div>
             </div>
           </div>
           
           <!-- خلاصه تنظیمات -->
           <div class="mt-6 space-y-2">
             <div class="flex justify-between text-sm">
               <span class="text-gray-600">نوع فرم:</span>
               <span class="text-gray-900">{{ formBuilderForm.type || 'انتخاب نشده' }}</span>
             </div>
             <div class="flex justify-between text-sm">
               <span class="text-gray-600">تعداد فیلدها:</span>
               <span class="text-gray-900">{{ formBuilderForm.fields.length }} فیلد</span>
             </div>
             <div class="flex justify-between text-sm">
               <span class="text-gray-600">وضعیت:</span>
               <span class="text-gray-900">{{ formBuilderForm.isDisabled ? 'غیرفعال' : 'فعال' }}</span>
             </div>
           </div>
         </div>

         <!-- بخش راست - فرم تنظیمات -->
         <div class="w-1/2 p-6 overflow-y-auto">
           <div class="space-y-6">
             <!-- اطلاعات اصلی -->
             <div>
               <h3 class="text-lg font-medium text-gray-900 mb-4">اطلاعات اصلی</h3>
               
               <!-- عنوان فرم -->
               <div class="mb-4">
                 <label class="block text-sm font-medium text-gray-700 mb-2">عنوان فرم *</label>
                 <input v-model="formBuilderForm.title" type="text" placeholder="عنوان فرم را وارد کنید" class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
               </div>
               
               <!-- نوع فرم -->
               <div class="mb-4">
                 <label class="block text-sm font-medium text-gray-700 mb-2">نوع فرم *</label>
                 <select v-model="formBuilderForm.type" class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
                   <option value="">انتخاب کنید</option>
                   <option value="فرم ساعات کاری">فرم ساعات کاری</option>
                   <option value="فرم اطلاعات تماس">فرم اطلاعات تماس</option>
                   <option value="فرم سفارشی">فرم سفارشی</option>
                 </select>
               </div>
               
               <!-- محتوای فرم -->
               <div class="mb-4">
                 <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات فرم</label>
                 <textarea v-model="formBuilderForm.content" placeholder="توضیحات فرم را وارد کنید" rows="3" class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"></textarea>
               </div>
             </div>

             <!-- تنظیمات فرم -->
             <div>
               <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات فرم</h3>
               
               <div class="space-y-3">
                 <label class="flex items-center">
                   <input v-model="formBuilderForm.isRequired" type="checkbox" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                   <span class="mr-2 text-sm text-gray-700">فرم اجباری است</span>
                 </label>
                 
                 <label class="flex items-center">
                   <input v-model="formBuilderForm.sendBeforeFirstMessage" type="checkbox" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                   <span class="mr-2 text-sm text-gray-700">ارسال قبل از پیام اول</span>
                 </label>
                 
                 <label class="flex items-center">
                   <input v-model="formBuilderForm.isDisabled" type="checkbox" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                   <span class="mr-2 text-sm text-gray-700">غیرفعال کردن فرم</span>
                 </label>
               </div>
             </div>

             <!-- مدیریت فیلدها -->
             <div>
               <div class="flex justify-between items-center mb-4">
                 <h3 class="text-lg font-medium text-gray-900">فیلدهای فرم</h3>
                 <button class="bg-purple-600 text-white px-4 py-2 rounded-lg hover:bg-purple-700 text-sm" @click="addFormField">
                   افزودن فیلد
                 </button>
               </div>
               
               <div class="space-y-4">
                 <div v-for="(field, index) in formBuilderForm.fields" :key="index" class="border border-gray-200 rounded-lg p-6">
                   <div class="flex justify-between items-start mb-3">
                     <h4 class="text-sm font-medium text-gray-900">فیلد {{ index + 1 }}</h4>
                     <button class="text-red-500 hover:text-red-700 text-sm" @click="removeFormField(index)">
                       حذف
                     </button>
                   </div>
                   
                   <div class="grid grid-cols-2 gap-3">
                     <div>
                       <label class="block text-xs text-gray-700 mb-1">نوع فیلد</label>
                       <select v-model="field.type" class="w-full p-2 border border-gray-300 rounded text-sm">
                         <option value="text">متن</option>
                         <option value="email">ایمیل</option>
                         <option value="phone">تلفن</option>
                         <option value="textarea">متن چند خطی</option>
                       </select>
                     </div>
                     <div>
                       <label class="block text-xs text-gray-700 mb-1">برچسب</label>
                       <input v-model="field.label" type="text" placeholder="برچسب فیلد" class="w-full p-2 border border-gray-300 rounded text-sm">
                     </div>
                   </div>
                   
                   <div class="mt-3">
                     <label class="flex items-center">
                       <input v-model="field.required" type="checkbox" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                       <span class="mr-2 text-xs text-gray-700">فیلد اجباری</span>
                     </label>
                   </div>
                 </div>
               </div>
               
               <div v-if="formBuilderForm.fields.length === 0" class="text-center py-8 text-gray-500">
                 <div class="text-4xl mb-2">📋</div>
                 <p class="text-sm">هیچ فیلدی تعریف نشده است</p>
                 <p class="text-xs">برای شروع، یک فیلد جدید اضافه کنید</p>
               </div>
             </div>
           </div>
         </div>
       </div>

       <!-- فوتر مودال -->
       <div class="bg-gray-50 px-6 py-4 flex justify-between items-center">
         <button class="px-4 py-2 text-gray-600 hover:text-gray-800" @click="closeFormModal">
           انصراف
         </button>
         <div class="flex space-x-3 space-x-reverse">
           <button class="px-4 py-2 text-gray-600 hover:text-gray-800" @click="resetFormBuilderForm">
             بازنشانی
           </button>
           <button class="bg-purple-600 text-white px-6 py-2 rounded-lg hover:bg-purple-700" @click="saveForm">
             {{ isEditingForm ? 'ویرایش' : 'ایجاد' }}
           </button>
         </div>
       </div>
     </div>




   </div>
   <!-- نرخ پیام -->
  <div v-if="activeTab === 'rateLimit'" class="space-y-6">
    <div class="bg-blue-600 text-white p-6 rounded-lg">
      <h2 class="text-2xl font-bold mb-2">تنظیم نرخ پیام</h2>
      <p class="text-blue-100">تعداد پیام مجاز در دقیقه برای هر نقش کاربری را تعیین کنید.</p>
    </div>
    <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6 space-y-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">مشتری (customer)</label>
          <input v-model.number="rateLimit.customer" type="number" min="1" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">اپراتور (operator)</label>
          <input v-model.number="rateLimit.operator" type="number" min="1" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">ادمین (admin)</label>
          <input v-model.number="rateLimit.admin" type="number" min="1" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">بازه زمانی (ثانیه)</label>
          <input v-model.number="rateLimit.window" type="number" min="10" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm" />
        </div>
      </div>
      <p class="text-sm text-gray-500">در صورت عبور کاربر از حد مجاز، پیام او رد شده و خطای <code>rate limit exceeded</code> دریافت می‌کند. پس از گذشت یک دقیقه امکان ارسال مجدد وجود دارد.</p>
      <div class="flex justify-end">
        <button class="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition-colors font-medium" @click="saveRateLimit">ثبت تغییرات</button>
      </div>
    </div>
  </div>

  <!-- امنیت چت -->
  <div v-if="activeTab === 'security'" class="space-y-6">
    <div class="bg-red-600 text-white p-6 rounded-lg">
      <h2 class="text-2xl font-bold mb-2">🛡️ تنظیمات امنیت چت</h2>
      <p class="text-red-100">مدیریت تنظیمات امنیتی برای محافظت از سیستم چت در برابر حملات مخرب</p>
    </div>
    
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- تنظیمات عمومی امنیت -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">🔧 تنظیمات عمومی</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">حداکثر طول پیام (کاراکتر)</label>
            <input
v-model.number="securitySettings.maxMessageLength" type="number" min="100" max="5000" 
                   class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-red-500 focus:ring-red-500 sm:text-sm" />
            <p class="text-xs text-gray-500 mt-1">پیش‌فرض: 1000 کاراکتر</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">مدت مسدود کردن (دقیقه)</label>
            <input
v-model.number="securitySettings.blockDuration" type="number" min="1" max="1440" 
                   class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-red-500 focus:ring-red-500 sm:text-sm" />
            <p class="text-xs text-gray-500 mt-1">مدت مسدود کردن برای تهدیدات critical (پیش‌فرض: 10 دقیقه)</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">مدت cooldown (ثانیه)</label>
            <input
v-model.number="securitySettings.cooldownDuration" type="number" min="10" max="300" 
                   class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-red-500 focus:ring-red-500 sm:text-sm" />
            <p class="text-xs text-gray-500 mt-1">مدت cooldown برای تهدیدات high (پیش‌فرض: 30 ثانیه)</p>
          </div>
          
          <div class="flex items-center">
            <input
id="strictMode" v-model="securitySettings.strictMode" type="checkbox"
                   class="h-4 w-4 text-red-600 focus:ring-red-500 border-gray-300 rounded" />
            <label for="strictMode" class="mr-2 block text-sm text-gray-900">حالت امنیت سخت</label>
          </div>
          <p class="text-xs text-gray-500">در حالت سخت، فیلترهای بیشتری فعال می‌شود</p>
        </div>
      </div>

      <!-- فرمت‌های فایل مجاز -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">📁 فرمت‌های فایل مجاز</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">انواع فایل مجاز</label>
            <div class="space-y-2">
              <div
v-for="(fileType, index) in securitySettings.allowedFileTypes" :key="index" 
                   class="flex items-center justify-between p-2 bg-gray-50 rounded">
                <span class="text-sm font-mono">{{ fileType }}</span>
                <button class="text-red-600 hover:text-red-800" @click="removeFileType(index)">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                  </svg>
                </button>
              </div>
            </div>
            <div class="flex mt-2">
              <input
v-model="newFileType" type="text" placeholder="مثال: image/png" 
                     class="flex-1 rounded-r-md border-gray-300 shadow-sm focus:border-red-500 focus:ring-red-500 sm:text-sm" 
                     @keydown.enter="addFileType" />
              <button
class="px-3 py-2 bg-red-600 text-white rounded-l-md hover:bg-red-700 text-sm" 
                      @click="addFileType">
                افزودن
              </button>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">حداکثر اندازه فایل (MB)</label>
            <input
v-model.number="securitySettings.maxFileSize" type="number" min="1" max="100" 
                   class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-red-500 focus:ring-red-500 sm:text-sm" />
            <p class="text-xs text-gray-500 mt-1">پیش‌فرض: 5 مگابایت</p>
          </div>
        </div>
      </div>
    </div>

    <!-- کلمات و الگوهای ممنوع -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">🚫 کلمات و الگوهای ممنوع</h3>
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- کلمات ممنوع -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کلمات ممنوع</label>
          <div class="space-y-2 max-h-40 overflow-y-auto">
            <div
v-for="(word, index) in securitySettings.bannedWords" :key="index" 
                 class="flex items-center justify-between p-2 bg-gray-50 rounded">
              <span class="text-sm">{{ word }}</span>
              <button class="text-red-600 hover:text-red-800" @click="removeBannedWord(index)">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
              </button>
            </div>
          </div>
          <div class="flex mt-2">
            <input
v-model="newBannedWord" type="text" placeholder="کلمه ممنوع جدید" 
                   class="flex-1 rounded-r-md border-gray-300 shadow-sm focus:border-red-500 focus:ring-red-500 sm:text-sm" 
                   @keydown.enter="addBannedWord" />
            <button
class="px-3 py-2 bg-red-600 text-white rounded-l-md hover:bg-red-700 text-sm" 
                    @click="addBannedWord">
              افزودن
            </button>
          </div>
        </div>

        <!-- URL های مشکوک -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">دامنه‌های مسدود</label>
          <div class="space-y-2 max-h-40 overflow-y-auto">
            <div
v-for="(domain, index) in securitySettings.blockedDomains" :key="index" 
                 class="flex items-center justify-between p-2 bg-gray-50 rounded">
              <span class="text-sm font-mono">{{ domain }}</span>
              <button class="text-red-600 hover:text-red-800" @click="removeBlockedDomain(index)">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
              </button>
            </div>
          </div>
          <div class="flex mt-2">
            <input
v-model="newBlockedDomain" type="text" placeholder="مثال: evil.com" 
                   class="flex-1 rounded-r-md border-gray-300 shadow-sm focus:border-red-500 focus:ring-red-500 sm:text-sm" 
                   @keydown.enter="addBlockedDomain" />
            <button
class="px-3 py-2 bg-red-600 text-white rounded-l-md hover:bg-red-700 text-sm" 
                    @click="addBlockedDomain">
              افزودن
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار امنیتی -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">📊 آمار امنیتی (24 ساعت گذشته)</h3>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="bg-blue-50 p-6 rounded-lg">
          <div class="text-2xl font-bold text-blue-600">{{ securityStats.totalMessages }}</div>
          <div class="text-sm text-blue-600">کل پیام‌ها</div>
        </div>
        <div class="bg-red-50 p-6 rounded-lg">
          <div class="text-2xl font-bold text-red-600">{{ securityStats.blockedMessages }}</div>
          <div class="text-sm text-red-600">پیام‌های مسدود</div>
        </div>
        <div class="bg-orange-50 p-6 rounded-lg">
          <div class="text-2xl font-bold text-orange-600">{{ securityStats.suspiciousUsers }}</div>
          <div class="text-sm text-orange-600">کاربران مشکوک</div>
        </div>
        <div class="bg-green-50 p-6 rounded-lg">
          <div class="text-2xl font-bold text-green-600">{{ securityStats.threatLevel }}</div>
          <div class="text-sm text-green-600">سطح تهدید</div>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
      <div class="flex flex-wrap gap-6 justify-between items-center">
        <div class="flex flex-wrap gap-2">
          <button
class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors font-medium" 
                  @click="testSecurity">
            🧪 تست امنیت
          </button>
          <button
class="px-4 py-2 bg-gray-600 text-white rounded-lg hover:bg-gray-700 transition-colors font-medium" 
                  @click="resetToDefaults">
            🔄 بازگشت به پیش‌فرض
          </button>
          <button
class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors font-medium" 
                  @click="exportSecurityLogs">
            📥 دانلود لاگ‌ها
          </button>
        </div>
        <button
class="px-6 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 transition-colors font-medium" 
                @click="saveSecuritySettings">
          💾 ذخیره تنظیمات
        </button>
      </div>
    </div>
  </div>

</template>

<script lang="ts">
// تعریف interface ها
interface ApiResponse<T> {
  data?: T
  success?: boolean
  message?: string
}

interface RateLimitData {
  customer_limit?: number
  operator_limit?: number
  admin_limit?: number
  window_seconds?: number
}

interface SecurityData {
  maxMessageLength?: number
  blockDuration?: number
  cooldownDuration?: number
  strictMode?: boolean
  allowedFileTypes?: string[]
  maxFileSize?: number
  bannedWords?: string[]
  blockedDomains?: string[]
}

interface SecurityStats {
  totalMessages?: number
  blockedMessages?: number
  suspiciousUsers?: number
  threatLevel?: string
}

interface Team {
  id: number
  name: string
  color: string
  icon: string
  operatorCount: number
  operators: Array<{ id: number; name: string; email: string }>
}

interface Operation {
  id: number
  name: string
  description: string
  type: string
  content: string
  sound: boolean
  waitTime: number
  departments: string[]
  teams: string[]
  conversationType?: string
  operationType?: string
  createdAt: string
}

interface SmartMessage {
  id: number
  title: string
  content: string
  status: string
  sender: string
  triggerType: string
  operationType?: string
  displayType?: string
  sendOncePerVisitor?: boolean
  withSound?: boolean
  sendOnMobile?: boolean
  conditionType?: string
  conditions?: unknown[]
  startDate?: string
  startTime?: string
  endType?: string
  endDate?: string
  endTime?: string
  scheduleType?: string
  lastUpdate: string
  updatedBy: { name: string; email: string }
}

interface Form {
  id: number
  title: string
  type: string
  status: string
  fieldCount: number
  lastUpdate: string
  content: string
  isRequired: boolean
  sendBeforeFirstMessage: boolean
  isDisabled: boolean
  fields: { type: string; label: string; required: boolean }[]
}

declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; body?: unknown; responseType?: string }) => Promise<T>
declare const $toast: { success: (msg: string) => void; error: (msg: string) => void; warning: (msg: string) => void }

export default {
  name: 'SupportSettingsPage'
}
</script>

<script setup lang="ts">
import { ref } from 'vue'

definePageMeta({ layout: 'admin-main', middleware: 'admin' })

const activeTab = ref('departments')
const showCreateTeamModal = ref(false)
const showEditTeamModal = ref(false)
const showWorkingHoursModal = ref(false)
const showCreateOperationModal = ref(false)
const showCreateSmartMessageModal = ref(false)
const showCreateFormModal = ref(false)
const isEditingOperation = ref(false)
const isEditingSmartMessage = ref(false)
const isEditingForm = ref(false)
const activeSmartMessageTab = ref('content')
const operatorSearch = ref('')
const selectedOperators = ref([])
const currentTeam = ref<Partial<Team>>({})
const currentOperation = ref<Partial<Operation>>({})
const currentSmartMessage = ref<Partial<SmartMessage>>({})
const currentForm = ref<Partial<Form>>({})

// فرم عملیات جدید
const operationForm = ref({
  conversationType: 'awaiting',
  operationType: 'all_departments',
  withSound: true,
  waitTime: 0,
  messageText: 'به سایت ما خوش اومدین، در خدمتتون هستیم 👋',
  departments: [],
  teams: []
})

// فرم پیام هوشمند جدید
const smartMessageForm = ref({
  title: '',
  content: 'به سایت ما خوش اومدین، در خدمتتون هستیم 👋',
  status: 'active',
  sender: 'random',
  operationType: 'message_only',
  displayType: 'full',
  sendOncePerVisitor: true,
  withSound: true,
  sendOnMobile: true,
  conditionType: 'at_least_one',
  conditions: [],
  startDate: '',
  startTime: '',
  endType: 'never',
  endDate: '',
  endTime: '',
  scheduleType: 'always'
})

// فرم ایجاد فرم جدید
const formBuilderForm = ref({
  title: '',
  type: '',
  content: '',
  isRequired: false,
  sendBeforeFirstMessage: false,
  isDisabled: false,
  fields: []
})

// تنظیمات زمان‌های غیرکاری

// نرخ پیام
const rateLimit = ref({
  customer: 20,
  operator: 60,
  admin: 120,
  window: 60,
})

// تنظیمات امنیت
const securitySettings = ref({
  maxMessageLength: 1000,
  blockDuration: 10, // دقیقه
  cooldownDuration: 30, // ثانیه
  strictMode: false,
  allowedFileTypes: [
    'image/jpeg',
    'image/png', 
    'image/gif',
    'image/webp',
    'text/plain',
    'application/pdf'
  ],
  maxFileSize: 5, // MB
  bannedWords: [
    'eval',
    'function',
    'alert',
    'prompt',
    'confirm',
    'document.cookie',
    'localStorage',
    'sessionStorage',
    'window.location',
    'location.href'
  ],
  blockedDomains: [
    'bit.ly',
    'tinyurl.com',
    'example-malicious.com'
  ]
})

// متغیرهای ورودی جدید
const newFileType = ref('')
const newBannedWord = ref('')
const newBlockedDomain = ref('')

// آمار امنیتی
const securityStats = ref({
  totalMessages: 1247,
  blockedMessages: 23,
  suspiciousUsers: 5,
  threatLevel: 'پایین'
})

async function _fetchRateLimit() {
  try {
    const response: ApiResponse<RateLimitData> = await $fetch('/api/admin/chat/settings/rate-limit')
    if (response.success && response.data) {
      rateLimit.value = {
        customer: response.data.customer_limit || 20,
        operator: response.data.operator_limit || 60,
        admin: response.data.admin_limit || 120,
        window: response.data.window_seconds || 60,
      }
    }
  } catch (err) {
    console.error('Error fetching rate limit:', err)
  }
}

async function saveRateLimit() {
  try {
    const response = await $fetch('/api/admin/chat/settings/rate-limit', {
      method: 'PUT',
      body: {
        customer_limit: rateLimit.value.customer,
        operator_limit: rateLimit.value.operator,
        admin_limit: rateLimit.value.admin,
        window_seconds: rateLimit.value.window,
      },
    }) as { success?: boolean }
    if (response.success) {
      $toast.success('تنظیمات با موفقیت ذخیره شد')
    }
  } catch (err) {
    // console.error('Error saving rate limit:', err)
    $toast.error('خطا در ذخیره تنظیمات')
  }
}

// توابع مدیریت امنیت
function addFileType() {
  if (newFileType.value.trim() && !securitySettings.value.allowedFileTypes.includes(newFileType.value.trim())) {
    securitySettings.value.allowedFileTypes.push(newFileType.value.trim())
    newFileType.value = ''
  }
}

function removeFileType(index: number) {
  securitySettings.value.allowedFileTypes.splice(index, 1)
}

function addBannedWord() {
  if (newBannedWord.value.trim() && !securitySettings.value.bannedWords.includes(newBannedWord.value.trim().toLowerCase())) {
    securitySettings.value.bannedWords.push(newBannedWord.value.trim().toLowerCase())
    newBannedWord.value = ''
  }
}

function removeBannedWord(index: number) {
  securitySettings.value.bannedWords.splice(index, 1)
}

function addBlockedDomain() {
  if (newBlockedDomain.value.trim() && !securitySettings.value.blockedDomains.includes(newBlockedDomain.value.trim().toLowerCase())) {
    securitySettings.value.blockedDomains.push(newBlockedDomain.value.trim().toLowerCase())
    newBlockedDomain.value = ''
  }
}

function removeBlockedDomain(index: number) {
  securitySettings.value.blockedDomains.splice(index, 1)
}

async function _fetchSecuritySettings() {
  try {
    const response: ApiResponse<SecurityData> = await $fetch('/api/admin/chat/settings/security')
    if (response.success && response.data) {
      securitySettings.value = { ...securitySettings.value, ...response.data }
    }
  } catch (err) {
    // console.error('Error fetching security settings:', err)
  }
}

async function saveSecuritySettings() {
  try {
    const response = await $fetch('/api/admin/chat/settings/security', {
      method: 'PUT',
      body: securitySettings.value
    }) as { success?: boolean }
    if (response.success) {
      $toast.success('تنظیمات امنیتی با موفقیت ذخیره شد')
    }
  } catch (err) {
    // console.error('Error saving security settings:', err)
    $toast.error('خطا در ذخیره تنظیمات امنیتی')
  }
}

async function _fetchSecurityStats() {
  try {
    const response: ApiResponse<SecurityStats> = await $fetch('/api/admin/chat/security/stats')
    if (response.success && response.data) {
      securityStats.value = {
        ...securityStats.value,
        ...response.data
      }
    }
  } catch (err) {
    // console.error('Error fetching security stats:', err)
  }
}

function testSecurity() {
  import('~/utils/securityTest').then(({ runSecurityTests }) => {
    const results = runSecurityTests()
    const passed = results.filter(r => r.passed).length
    const total = results.length
    const passRate = ((passed / total) * 100).toFixed(1)
    
    if (passed === total) {
      $toast.success(`✅ تست امنیت موفق: ${passed}/${total} تست (${passRate}%)`)
    } else {
      $toast.warning(`⚠️ برخی تست‌ها ناموفق: ${passed}/${total} تست (${passRate}%)`)
    }
  }).catch(err => {
    // console.error('Error running security tests:', err)
    $toast.error('خطا در اجرای تست‌های امنیت')
  })
}

function resetToDefaults() {
  if (confirm('آیا مطمئن هستید که می‌خواهید تنظیمات به حالت پیش‌فرض برگردد؟')) {
    securitySettings.value = {
      maxMessageLength: 1000,
      blockDuration: 10,
      cooldownDuration: 30,
      strictMode: false,
      allowedFileTypes: [
        'image/jpeg',
        'image/png', 
        'image/gif',
        'image/webp',
        'text/plain',
        'application/pdf'
      ],
      maxFileSize: 5,
      bannedWords: [
        'eval', 'function', 'alert', 'prompt', 'confirm',
        'document.cookie', 'localStorage', 'sessionStorage',
        'window.location', 'location.href'
      ],
      blockedDomains: [
        'bit.ly', 'tinyurl.com', 'example-malicious.com'
      ]
    }
    $toast.success('تنظیمات به حالت پیش‌فرض برگشت')
  }
}

async function exportSecurityLogs() {
  try {
    const response = await $fetch('/api/admin/chat/security/logs', {
      method: 'GET',
      responseType: 'blob'
    }) as Blob
    
    // ایجاد لینک دانلود
    const url = window.URL.createObjectURL(new Blob([response]))
    const link = document.createElement('a')
    link.href = url
    link.setAttribute('download', `security-logs-${new Date().toISOString().split('T')[0]}.json`)
    document.body.appendChild(link)
    link.click()
    link.remove()
    window.URL.revokeObjectURL(url)
    
    $toast.success('لاگ‌های امنیتی دانلود شد')
  } catch (err) {
    // console.error('Error exporting security logs:', err)
    $toast.error('خطا در دانلود لاگ‌های امنیتی')
  }
}
const offlineSettings = ref({
  offlineMessage: true,
  changeToOnline: false,
  smartMessages: false,
  sendToTelegram: false,
  sleepMode: false,
  hideWidget: false,
  disableChat: false
})



// فرم تیم جدید
const newTeam = ref({
  name: '',
  color: '#8B5CF6',
  icon: '💬',
  operators: []
})

// لیست تیم‌ها
const teams = ref([
  {
    id: 1,
    name: 'تیم پشتیبانی',
    color: '#F59E0B',
    icon: '🦊',
    operatorCount: 1,
    operators: [
      {
        id: 1,
        name: 'آ . آذری',
        email: 'azitaazari۵۸۵۰@gmail.com'
      }
    ]
  },
  {
    id: 2,
    name: 'تیم پیشفرض',
    color: '#8B5CF6',
    icon: '💬',
    operatorCount: 0,
    operators: []
  }
])

// لیست عملیات‌های جدید
const operations = ref([
  {
    id: 1,
    name: 'پاسخگویی',
    description: 'پاسخگویی به پیام‌های کاربران در حالت انتظار',
    type: 'auto_message',
    content: 'پاسخگوی سوالات شما هستیم.',
    sound: true,
    waitTime: 30,
    departments: [], // مثال: ['پشتیبانی', 'فنی']
    teams: [], // مثال: ['تیم پشتیبانی', 'تیم فنی']
    createdAt: '2023-10-27T10:00:00Z'
  },
  {
    id: 2,
    name: 'تاخیر در پاسخگویی',
    description: 'تاخیر در پاسخگویی به پیام‌های کاربران در حالت انتظار',
    type: 'delay',
    content: 'به زودی پاسخ می‌دهیم.',
    sound: false,
    waitTime: 60,
    departments: ['پشتیبانی'],
    teams: ['تیم پشتیبانی'],
    createdAt: '2023-10-27T11:00:00Z'
  }
])

// لیست پیام‌های هوشمند
const smartMessages = ref([
  {
    id: 1,
    title: 'شرط ۱',
    content: 'پیام خوش‌آمدگویی برای کاربران جدید',
    status: 'active',
    sender: 'random',
    triggerType: 'welcome',
    lastUpdate: '۱۴۰۳/۵/۵',
    updatedBy: {
      name: '۱. آذری',
      email: 'azitaazari۵۸۵۰@gmail.com'
    },
    conditions: []
  }
])

// لیست فرم‌ها
const forms = ref([
  {
    id: 1,
    title: 'فرم اطلاعات تماس',
    type: 'فرم ساعات کاری',
    status: 'active',
    fieldCount: 3,
    lastUpdate: '۱۴۰۳/۵/۵',
    content: 'فرم جمع‌آوری اطلاعات تماس کاربران',
    isRequired: true,
    sendBeforeFirstMessage: false,
    isDisabled: false,
    fields: [
      { type: 'text', label: 'نام', required: true },
      { type: 'email', label: 'ایمیل', required: true },
      { type: 'phone', label: 'تلفن', required: false }
    ]
  }
])


// ایجاد تیم جدید
function createTeam() {
  if (!newTeam.value.name) {
    alert('لطفاً نام تیم را وارد کنید')
    return
  }

  const team = {
    id: Date.now(),
    ...newTeam.value,
    operatorCount: selectedOperators.value.length,
    operators: [...selectedOperators.value]
  }

  teams.value.push(team)
  
  // پاک کردن فرم
  resetForm()
  showCreateTeamModal.value = false
  
  alert('تیم با موفقیت ایجاد شد')
}

// به‌روزرسانی تیم
function updateTeam() {
  if (!currentTeam.value.name) {
    alert('لطفاً نام تیم را وارد کنید')
    return
  }

  const index = teams.value.findIndex(team => team.id === currentTeam.value.id)
  if (index !== -1 && currentTeam.value.id) {
    teams.value[index] = { ...currentTeam.value } as Team
  }
  
  showEditTeamModal.value = false
  alert('تیم با موفقیت به‌روزرسانی شد')
}

// حذف تیم
function deleteTeam(id) {
  if (confirm('آیا مطمئن هستید که می‌خواهید این تیم را حذف کنید؟')) {
    teams.value = teams.value.filter(team => team.id !== id)
    alert('تیم با موفقیت حذف شد')
  }
}

// حذف اپراتور از لیست انتخاب شده
function removeOperator(id) {
  selectedOperators.value = selectedOperators.value.filter(op => op.id !== id)
}

// حذف اپراتور از تیم
function removeOperatorFromTeam(id) {
  if (currentTeam.value.operators) {
    currentTeam.value.operators = currentTeam.value.operators.filter(op => op.id !== id)
    currentTeam.value.operatorCount = currentTeam.value.operators.length
  }
}



// حذف عملیات
function deleteOperation(id) {
  if (confirm('آیا مطمئن هستید که می‌خواهید این عملیات را حذف کنید؟')) {
    operations.value = operations.value.filter(op => op.id !== id)
    alert('عملیات با موفقیت حذف شد')
  }
}

// باز کردن مودال ایجاد عملیات جدید
function openCreateOperationModal() {
  isEditingOperation.value = false
  resetOperationForm()
  showCreateOperationModal.value = true
}

// باز کردن مودال ویرایش عملیات
function openEditOperationModal(operation) {
  isEditingOperation.value = true
  currentOperation.value = { ...operation }
  operationForm.value = {
    conversationType: operation.conversationType || 'awaiting',
    operationType: operation.operationType || 'all_departments',
    withSound: operation.sound || true,
    waitTime: operation.waitTime || 0,
    messageText: operation.content || 'به سایت ما خوش اومدین، در خدمتتون هستیم 👋',
    departments: operation.departments || [],
    teams: operation.teams || []
  }
  showCreateOperationModal.value = true
}

// بستن مودال عملیات
function closeOperationModal() {
  showCreateOperationModal.value = false
  isEditingOperation.value = false
  currentOperation.value = {}
  resetOperationForm()
}

// بازنشانی فرم عملیات
function resetOperationForm() {
  operationForm.value = {
    conversationType: 'awaiting',
    operationType: 'all_departments',
    withSound: true,
    waitTime: 0,
    messageText: 'به سایت ما خوش اومدین، در خدمتتون هستیم 👋',
    departments: [],
    teams: []
  }
}

// ذخیره عملیات
function saveOperation() {
  if (!operationForm.value.messageText.trim()) {
    alert('لطفاً متن پیام را وارد کنید')
    return
  }

  const operationData = {
    conversationType: operationForm.value.conversationType,
    operationType: operationForm.value.operationType,
    sound: operationForm.value.withSound,
    waitTime: operationForm.value.waitTime,
    content: operationForm.value.messageText,
    departments: operationForm.value.departments,
    teams: operationForm.value.teams,
    createdAt: new Date().toISOString()
  }

  if (isEditingOperation.value) {
    // ویرایش عملیات موجود
    const index = operations.value.findIndex(op => op.id === currentOperation.value.id)
    if (index !== -1) {
      operations.value[index] = {
        ...operations.value[index],
        ...operationData
      }
    }
    alert('عملیات با موفقیت ویرایش شد')
  } else {
    // ایجاد عملیات جدید
    const newOperation = {
      id: Date.now(),
      name: 'عملیات جدید',
      description: 'عملیات خودکار برای پاسخگویی',
      type: 'auto_message',
      ...operationData
    }
    operations.value.push(newOperation)
    alert('عملیات با موفقیت ایجاد شد')
  }

  closeOperationModal()
}

// باز کردن مودال ایجاد پیام هوشمند جدید
function _openCreateSmartMessageModal() {
  isEditingSmartMessage.value = false
  currentSmartMessage.value = {}
  resetSmartMessageForm()
  showCreateSmartMessageModal.value = true
}

// ویرایش پیام هوشمند
function editSmartMessage(message) {
  isEditingSmartMessage.value = true
  currentSmartMessage.value = { ...message }
  smartMessageForm.value = {
    title: message.title,
    content: message.content,
    status: message.status,
    sender: message.sender,
    operationType: message.operationType || 'message_only',
    displayType: message.displayType || 'full',
    sendOncePerVisitor: message.sendOncePerVisitor !== false,
    withSound: message.withSound !== false,
    sendOnMobile: message.sendOnMobile !== false,
    conditionType: message.conditionType || 'at_least_one',
    conditions: message.conditions || [],
    startDate: message.startDate || '',
    startTime: message.startTime || '',
    endType: message.endType || 'never',
    endDate: message.endDate || '',
    endTime: message.endTime || '',
    scheduleType: message.scheduleType || 'always'
  }
  showCreateSmartMessageModal.value = true
}

// حذف پیام هوشمند
function deleteSmartMessage(id) {
  if (confirm('آیا از حذف این پیام هوشمند اطمینان دارید؟')) {
    const index = smartMessages.value.findIndex(msg => msg.id === id)
    if (index !== -1) {
      smartMessages.value.splice(index, 1)
      alert('پیام هوشمند با موفقیت حذف شد')
    }
  }
}

// بستن مودال پیام هوشمند
function closeSmartMessageModal() {
  showCreateSmartMessageModal.value = false
  isEditingSmartMessage.value = false
  currentSmartMessage.value = {}
  resetSmartMessageForm()
}

// بازنشانی فرم پیام هوشمند
function resetSmartMessageForm() {
  smartMessageForm.value = {
    title: '',
    content: 'به سایت ما خوش اومدین، در خدمتتون هستیم 👋',
    status: 'active',
    sender: 'random',
    operationType: 'message_only',
    displayType: 'full',
    sendOncePerVisitor: true,
    withSound: true,
    sendOnMobile: true,
    conditionType: 'at_least_one',
    conditions: [],
    startDate: '',
    startTime: '',
    endType: 'never',
    endDate: '',
    endTime: '',
    scheduleType: 'always'
  }
}

// دریافت متن نوع نمایش
function getDisplayTypeText(displayType) {
  switch (displayType) {
    case 'full': return 'پیام کامل'
    case 'short': return 'پیام کوتاه'
    case 'notification': return 'اعلان'
    default: return 'پیام کامل'
  }
}

// دریافت متن زمان انتشار
function getPublishTimeText() {
  const startDate = smartMessageForm.value.startDate
  const startTime = smartMessageForm.value.startTime
  const endType = smartMessageForm.value.endType
  const scheduleType = smartMessageForm.value.scheduleType
  
  let text = ''
  if (startDate && startTime) {
    text += `شروع ارسال: ${startDate} ساعت ${startTime} `
  }
  
  text += `پایان ارسال: ${endType === 'never' ? 'هیچوقت' : 'تنظیم شده'} `
  
  switch (scheduleType) {
    case 'always': text += 'هر روز هر ساعتی'; break
    case 'working_hours': text += 'در طول زمانهای کاری'; break
    case 'non_working_hours': text += 'در طول زمانهای غیرکاری'; break
  }
  
  return text
}

// دریافت واحد شرط
function getConditionUnit(conditionType) {
  switch (conditionType) {
    case 'time_on_page': return 'ثانیه'
    case 'scroll_percentage': return '%'
    case 'page_visit': return 'بار'
    default: return ''
  }
}

// افزودن شرط جدید
function addCondition() {
  smartMessageForm.value.conditions.push({
    type: 'time_on_page',
    value: 30
  })
}

// حذف شرط
function removeCondition(index) {
  smartMessageForm.value.conditions.splice(index, 1)
}

// باز کردن مودال ایجاد فرم جدید
function openCreateFormModal() {
  isEditingForm.value = false
  currentForm.value = {}
  resetFormBuilderForm()
  showCreateFormModal.value = true
}

// ویرایش فرم
function editForm(form) {
  isEditingForm.value = true
  currentForm.value = { ...form }
  formBuilderForm.value = {
    title: form.title,
    type: form.type,
    content: form.content,
    isRequired: form.isRequired,
    sendBeforeFirstMessage: form.sendBeforeFirstMessage,
    isDisabled: form.isDisabled,
    fields: form.fields || []
  }
  showCreateFormModal.value = true
}

// حذف فرم
function deleteForm(id) {
  if (confirm('آیا از حذف این فرم اطمینان دارید؟')) {
    const index = forms.value.findIndex(form => form.id === id)
    if (index !== -1) {
      forms.value.splice(index, 1)
      alert('فرم با موفقیت حذف شد')
    }
  }
}

// بستن مودال فرم
function closeFormModal() {
  showCreateFormModal.value = false
  isEditingForm.value = false
  currentForm.value = {}
  resetFormBuilderForm()
}

// بازنشانی فرم فرم‌ساز
function resetFormBuilderForm() {
  formBuilderForm.value = {
    title: '',
    type: '',
    content: '',
    isRequired: false,
    sendBeforeFirstMessage: false,
    isDisabled: false,
    fields: []
  }
}

// ذخیره فرم
function saveForm() {
  if (!formBuilderForm.value.title.trim()) {
    alert('لطفاً عنوان فرم را وارد کنید')
    return
  }
  
  if (!formBuilderForm.value.type) {
    alert('لطفاً نوع فرم را انتخاب کنید')
    return
  }

  const formData: Partial<Form> = {
    title: formBuilderForm.value.title,
    type: formBuilderForm.value.type,
    status: formBuilderForm.value.isDisabled ? 'inactive' : 'active',
    content: formBuilderForm.value.content,
    isRequired: formBuilderForm.value.isRequired,
    sendBeforeFirstMessage: formBuilderForm.value.sendBeforeFirstMessage,
    isDisabled: formBuilderForm.value.isDisabled,
    fields: formBuilderForm.value.fields,
    fieldCount: formBuilderForm.value.fields.length,
    lastUpdate: new Date().toLocaleDateString('fa-IR')
  }

  if (isEditingForm.value) {
    // ویرایش فرم موجود
    const index = forms.value.findIndex(form => form.id === currentForm.value.id)
    if (index !== -1) {
      forms.value[index] = {
        ...forms.value[index],
        ...formData
      }
    }
    alert('فرم با موفقیت ویرایش شد')
  } else {
    // ایجاد فرم جدید
    const newForm: Form = {
      id: Date.now(),
      title: formData.title || '',
      type: formData.type || '',
      status: formData.status || 'active',
      content: formData.content || '',
      isRequired: formData.isRequired || false,
      sendBeforeFirstMessage: formData.sendBeforeFirstMessage || false,
      isDisabled: formData.isDisabled || false,
      fields: (formData.fields || []) as { type: string; label: string; required: boolean }[],
      fieldCount: formData.fieldCount || 0,
      lastUpdate: formData.lastUpdate || new Date().toLocaleDateString('fa-IR')
    }
    forms.value.push(newForm)
    alert('فرم با موفقیت ایجاد شد')
  }

  closeFormModal()
}

// ذخیره پیام هوشمند
function saveSmartMessage() {
  if (!smartMessageForm.value.title.trim() || !smartMessageForm.value.content.trim()) {
    alert('لطفاً عنوان و محتوای پیام را وارد کنید')
    return
  }

  const messageData: Partial<SmartMessage> = {
    title: smartMessageForm.value.title,
    content: smartMessageForm.value.content,
    status: smartMessageForm.value.status,
    sender: smartMessageForm.value.sender,
    triggerType: smartMessageForm.value.operationType || 'welcome',
    operationType: smartMessageForm.value.operationType,
    displayType: smartMessageForm.value.displayType,
    sendOncePerVisitor: smartMessageForm.value.sendOncePerVisitor,
    withSound: smartMessageForm.value.withSound,
    sendOnMobile: smartMessageForm.value.sendOnMobile,
    conditionType: smartMessageForm.value.conditionType,
    conditions: smartMessageForm.value.conditions,
    startDate: smartMessageForm.value.startDate,
    startTime: smartMessageForm.value.startTime,
    endType: smartMessageForm.value.endType,
    endDate: smartMessageForm.value.endDate,
    endTime: smartMessageForm.value.endTime,
    scheduleType: smartMessageForm.value.scheduleType,
    lastUpdate: new Date().toLocaleDateString('fa-IR'),
    updatedBy: {
      name: '۱. آذری',
      email: 'azitaazari۵۸۵۰@gmail.com'
    }
  }

  if (isEditingSmartMessage.value) {
    // ویرایش پیام موجود
    const index = smartMessages.value.findIndex(msg => msg.id === currentSmartMessage.value.id)
    if (index !== -1) {
      smartMessages.value[index] = {
        ...smartMessages.value[index],
        ...messageData
      }
    }
    alert('پیام هوشمند با موفقیت ویرایش شد')
  } else {
    // ایجاد پیام جدید
    const newMessage: SmartMessage = {
      id: Date.now(),
      title: messageData.title || '',
      content: messageData.content || '',
      status: messageData.status || 'active',
      sender: messageData.sender || 'random',
      triggerType: messageData.triggerType || 'welcome',
      operationType: messageData.operationType,
      displayType: messageData.displayType,
      sendOncePerVisitor: messageData.sendOncePerVisitor,
      withSound: messageData.withSound,
      sendOnMobile: messageData.sendOnMobile,
      conditionType: messageData.conditionType,
      conditions: messageData.conditions || [],
      startDate: messageData.startDate,
      startTime: messageData.startTime,
      endType: messageData.endType,
      endDate: messageData.endDate,
      endTime: messageData.endTime,
      scheduleType: messageData.scheduleType,
      lastUpdate: messageData.lastUpdate || new Date().toLocaleDateString('fa-IR'),
      updatedBy: messageData.updatedBy || { name: '۱. آذری', email: 'azitaazari۵۸۵۰@gmail.com' }
    }
    smartMessages.value.push(newMessage as SmartMessage & { conditions: unknown[] })
    alert('پیام هوشمند با موفقیت ایجاد شد')
  }

  closeSmartMessageModal()
}


// بازنشانی فرم
function resetForm() {
  newTeam.value = {
    name: '',
    color: '#8B5CF6',
    icon: '💬',
    operators: []
  }
  selectedOperators.value = []
  operatorSearch.value = ''
}

// ذخیره تمام تنظیمات
function saveAllSettings() {
  // اینجا کد ذخیره تنظیمات در دیتابیس قرار می‌گیرد
  
  // نمایش پیام موفقیت
  alert('تنظیمات با موفقیت ذخیره شد')
}

// افزودن فیلد جدید به فرم
function addFormField() {
  formBuilderForm.value.fields.push({
    type: 'text',
    label: '',
    required: false
  })
}

// حذف فیلد از فرم
function removeFormField(index: number) {
  formBuilderForm.value.fields.splice(index, 1)
}
</script> 
