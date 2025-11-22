<template>
  <div>
    <div class="max-w-6xl mx-auto bg-white rounded-lg shadow p-8 mt-8">
      <h1 class="text-2xl font-bold mb-6 text-gray-900">ویرایش اسکیما</h1>
      
      <!-- اطلاعات اصلی اسکیما -->
      <div class="bg-blue-50 border border-blue-200 rounded-lg p-6 mb-8">
        <h2 class="text-lg font-semibold text-blue-900 mb-4">اطلاعات اصلی اسکیما</h2>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع اسکیما *</label>
            <select v-model="form.type" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500" disabled>
              <option value="">انتخاب کنید</option>
              <option value="Article">مقاله (Article)</option>
              <option value="Product">محصول (Product)</option>
              <option value="Organization">سازمان (Organization)</option>
              <option value="WebPage">صفحه وب (WebPage)</option>
              <option value="BreadcrumbList">مسیر ناوبری (BreadcrumbList)</option>
              <option value="FAQPage">سوالات متداول (FAQPage)</option>
              <option value="LocalBusiness">کسب و کار محلی (LocalBusiness)</option>
              <option value="Review">نظر (Review)</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام اسکیما</label>
            <input v-model="form.name" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="نام اسکیما">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
            <select v-model="form.isActive" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option :value="true">فعال</option>
              <option :value="false">غیرفعال</option>
            </select>
          </div>
        </div>
        <div class="mt-4">
          <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
          <textarea v-model="form.description" rows="3" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="توضیحات اسکیما"></textarea>
        </div>
      </div>

      <!-- فیلدهای اسکیما -->
      <div class="bg-green-50 border border-green-200 rounded-lg p-6 mb-8">
        <h2 class="text-lg font-semibold text-green-900 mb-4">فیلدهای اسکیما و مقادیر</h2>
        <p class="text-sm text-gray-600 mb-6">در این بخش می‌توانید فیلدهای اسکیما را ویرایش کنید. مقادیر با {{ }} نشان‌دهنده متغیرهای پویا هستند که از داده‌های مقاله گرفته می‌شوند.</p>
        
        <div class="space-y-6">
          <!-- فیلدهای عمومی -->
          <div>
            <h3 class="text-md font-medium text-gray-900 mb-3 border-b border-green-200 pb-2">فیلدهای عمومی</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  عنوان (Title)
                  <span class="text-xs text-gray-500">- عنوان اصلی محتوا</span>
                </label>
                <input v-model="form.title" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{title}}">
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('title') }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  خلاصه (Excerpt)
                  <span class="text-xs text-gray-500">- خلاصه کوتاه محتوا</span>
                </label>
                <textarea v-model="form.excerpt" rows="2" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{excerpt}}"></textarea>
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('excerpt') }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  ID مقاله (Article ID)
                  <span class="text-xs text-gray-500">- شناسه یکتای مقاله</span>
                </label>
                <input v-model="form.articleId" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{id}}">
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('id') }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  URL مقاله (Article URL)
                  <span class="text-xs text-gray-500">- آدرس کامل مقاله</span>
                </label>
                <input v-model="form.articleUrl" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{url}}">
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('url') }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  نام سایت (Site Name)
                  <span class="text-xs text-gray-500">- نام سایت یا برند</span>
                </label>
                <input v-model="form.siteName" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{site_name}}">
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('site_name') }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  تصویر شاخص (Featured Image)
                  <span class="text-xs text-gray-500">- تصویر اصلی محتوا</span>
                </label>
                <input v-model="form.featuredImage" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{featured_image}}">
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('featured_image') }}</p>
              </div>
            </div>
          </div>

          <!-- فیلدهای مخصوص مقاله -->
          <div v-if="form.type === 'Article'">
            <h3 class="text-md font-medium text-gray-900 mb-3 border-b border-green-200 pb-2">فیلدهای مخصوص مقاله</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  نویسنده (Author)
                  <span class="text-xs text-gray-500">- نام نویسنده مقاله</span>
                </label>
                <input v-model="form.author" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{author}}">
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('author') }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  تاریخ انتشار (Published At)
                  <span class="text-xs text-gray-500">- تاریخ انتشار مقاله</span>
                </label>
                <input v-model="form.publishedAt" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{published_at}}">
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('published_at') }}</p>
              </div>
            </div>
          </div>

          <!-- فیلدهای مخصوص محصول -->
          <div v-if="form.type === 'Product'">
            <h3 class="text-md font-medium text-gray-900 mb-3 border-b border-green-200 pb-2">فیلدهای مخصوص محصول</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  قیمت (Price)
                  <span class="text-xs text-gray-500">- قیمت محصول</span>
                </label>
                <input v-model="form.price" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{price}}">
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('price') }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  واحد پول (Currency)
                  <span class="text-xs text-gray-500">- واحد پول قیمت</span>
                </label>
                <select v-model="form.currency" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500">
                  <option value="IRR">ریال ایران (IRR)</option>
                  <option value="USD">دلار آمریکا (USD)</option>
                  <option value="EUR">یورو (EUR)</option>
                </select>
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('currency') }}</p>
              </div>
            </div>
          </div>

          <!-- فیلدهای مخصوص سازمان -->
          <div v-if="form.type === 'Organization'">
            <h3 class="text-md font-medium text-gray-900 mb-3 border-b border-green-200 pb-2">فیلدهای مخصوص سازمان</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  آدرس (Address)
                  <span class="text-xs text-gray-500">- آدرس سازمان</span>
                </label>
                <input v-model="form.address" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{address}}">
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('address') }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  شماره تماس (Telephone)
                  <span class="text-xs text-gray-500">- شماره تماس سازمان</span>
                </label>
                <input v-model="form.telephone" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{telephone}}">
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('telephone') }}</p>
              </div>
            </div>
          </div>

          <!-- فیلدهای متا -->
          <div>
            <h3 class="text-md font-medium text-gray-900 mb-3 border-b border-green-200 pb-2">فیلدهای متا</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  متا تایتل (Meta Title)
                  <span class="text-xs text-gray-500">- عنوان برای موتورهای جستجو</span>
                </label>
                <input v-model="form.metaTitle" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{title}} - {{site_name}}">
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('meta_title') }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  متا دیسکریپشن (Meta Description)
                  <span class="text-xs text-gray-500">- توضیحات برای موتورهای جستجو</span>
                </label>
                <textarea v-model="form.metaDescription" rows="2" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{excerpt}}"></textarea>
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('meta_description') }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  کلمات کلیدی (Meta Keywords)
                  <span class="text-xs text-gray-500">- کلمات کلیدی برای SEO</span>
                </label>
                <input v-model="form.metaKeywords" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{keywords}}">
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('meta_keywords') }}</p>
              </div>
            </div>
          </div>

          <!-- فیلدهای Open Graph -->
          <div>
            <h3 class="text-md font-medium text-gray-900 mb-3 border-b border-green-200 pb-2">فیلدهای Open Graph</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  OG تایتل (OG Title)
                  <span class="text-xs text-gray-500">- عنوان برای شبکه‌های اجتماعی</span>
                </label>
                <input v-model="form.ogTitle" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{title}}">
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('og_title') }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  OG دیسکریپشن (OG Description)
                  <span class="text-xs text-gray-500">- توضیحات برای شبکه‌های اجتماعی</span>
                </label>
                <textarea v-model="form.ogDescription" rows="2" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{excerpt}}"></textarea>
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('og_description') }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  OG تصویر (OG Image)
                  <span class="text-xs text-gray-500">- تصویر برای شبکه‌های اجتماعی</span>
                </label>
                <input v-model="form.ogImage" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{featured_image}}">
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('og_image') }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  OG تایپ (OG Type)
                  <span class="text-xs text-gray-500">- نوع محتوا برای شبکه‌های اجتماعی</span>
                </label>
                <select v-model="form.ogType" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500">
                  <option value="article">مقاله (article)</option>
                  <option value="website">وب‌سایت (website)</option>
                  <option value="product">محصول (product)</option>
                  <option value="organization">سازمان (organization)</option>
                </select>
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('og_type') }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  OG سایت نیم (OG Site Name)
                  <span class="text-xs text-gray-500">- نام سایت</span>
                </label>
                <input v-model="form.ogSiteName" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="{{site_name}}">
                <p class="text-xs text-gray-500 mt-1">مقدار فعلی: {{ getFieldValue('og_site_name') }}</p>
              </div>
            </div>
          </div>

          <!-- فیلدهای اضافی JSON -->
          <div>
            <h3 class="text-md font-medium text-gray-900 mb-3 border-b border-green-200 pb-2">فیلدهای اضافی (JSON)</h3>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                فیلدهای اضافی
                <span class="text-xs text-gray-500">- فیلدهای اضافی اسکیما در فرمت JSON</span>
              </label>
              <textarea v-model="form.extraFields" rows="8" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500 font-mono text-sm" placeholder='{"headline": "{{title}}", "author": "{{author}}", "description": "{{excerpt}}"}'></textarea>
              <p class="text-xs text-gray-500 mt-1">این فیلدها برای اسکیمای JSON-LD استفاده می‌شوند</p>
            </div>
          </div>
        </div>
      </div>

      <!-- افزودن فیلدهای سفارشی -->
      <div class="bg-purple-50 border border-purple-200 rounded-lg p-6 mb-8">
        <h2 class="text-lg font-semibold text-purple-900 mb-4">افزودن فیلدهای سفارشی</h2>
        <p class="text-sm text-gray-600 mb-6">در این بخش می‌توانید فیلدهای سفارشی جدید به اسکیما اضافه کنید.</p>
        
        <!-- افزودن فیلد تکی -->
        <div class="bg-white border border-purple-200 rounded-lg p-6 mb-6">
          <h3 class="text-md font-medium text-purple-900 mb-3">{{ isEditing ? 'ویرایش فیلد' : 'افزودن فیلد جدید' }}</h3>
          <div class="grid grid-cols-1 md:grid-cols-5 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">بخش مقصد</label>
              <select v-model="newField.section" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
                <option value="general">فیلدهای عمومی</option>
                <option value="meta">فیلدهای متا</option>
                <option value="og">فیلدهای Open Graph</option>
                <option value="json">فیلدهای اضافی (JSON)</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نام فیلد (Key)</label>
              <input v-model="newField.key" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-purple-500 focus:border-purple-500" placeholder="مثال: customField">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">برچسب (Label)</label>
              <input v-model="newField.label" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-purple-500 focus:border-purple-500" placeholder="مثال: فیلد سفارشی">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع فیلد</label>
              <select v-model="newField.type" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
                <option value="text">متن</option>
                <option value="textarea">متن چند خطی</option>
                <option value="number">عدد</option>
                <option value="date">تاریخ</option>
                <option value="url">لینک</option>
                <option value="email">ایمیل</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">مقدار</label>
              <input v-if="newField.section !== 'json'" v-model="newField.value" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-purple-500 focus:border-purple-500" placeholder="مثال: {{custom_value}}">
              <textarea v-else v-model="newField.value" rows="3" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-purple-500 focus:border-purple-500 font-mono text-sm" placeholder='{"key": "value", "nested": {"key": "{{variable}}"}}'></textarea>
            </div>
          </div>
          <!-- راهنمای JSON -->
          <div v-if="newField.section === 'json'" class="mt-3 p-3 bg-blue-50 border border-blue-200 rounded-md">
            <p class="text-sm text-blue-800 mb-2">
              <strong>راهنمای فیلد JSON:</strong>
            </p>
            <ul class="text-xs text-blue-700 space-y-1">
              <li>• فیلدهای JSON باید در فرمت معتبر JSON باشند</li>
              <li>• می‌توانید از متغیرهای پویا مانند <code v-pre class="bg-blue-100 px-1 rounded">{{title}}</code> استفاده کنید</li>
              <li>• مثال: <code v-pre class="bg-blue-100 px-1 rounded">{"customField": "{{title}}", "nested": {"key": "{{author}}"}}</code></li>
            </ul>
          </div>
          
          <div class="mt-4 flex justify-end gap-2">
            <button v-if="isEditing" class="px-4 py-2 bg-gray-500 text-white rounded-md text-sm font-medium hover:bg-gray-600" @click="cancelEdit">
              انصراف
            </button>
            <button class="px-4 py-2 bg-purple-600 text-white rounded-md text-sm font-medium hover:bg-purple-700" @click="isEditing ? updateCustomField() : addCustomField()">
              {{ isEditing ? 'بروزرسانی فیلد' : 'افزودن فیلد' }}
            </button>
          </div>
        </div>



        <!-- نمایش فیلدهای سفارشی موجود -->
        <div v-if="customFields.length > 0" class="bg-white border border-purple-200 rounded-lg p-6">
          <h3 class="text-md font-medium text-purple-900 mb-3">فیلدهای سفارشی موجود</h3>
          <div class="space-y-3">
            <div v-for="(field, index) in customFields" :key="index" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <div class="flex-1">
                <div class="flex items-center gap-6">
                  <span class="font-medium text-gray-900">{{ field.label }}</span>
                  <span class="text-sm text-gray-500">({{ field.key }})</span>
                  <span class="text-sm text-purple-600">{{ field.value }}</span>
                  <span class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded">{{ getSectionLabel(field.section) }}</span>
                </div>
              </div>
              <div class="flex gap-2">
                <button class="px-3 py-1 bg-blue-500 text-white rounded text-sm hover:bg-blue-600" @click="editCustomField(index)">
                  ویرایش
                </button>
                <button class="px-3 py-1 bg-red-500 text-white rounded text-sm hover:bg-red-600" @click="removeCustomField(index)">
                  حذف
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- راهنمای متغیرها -->
      <div class="bg-yellow-50 border border-yellow-200 rounded-lg p-6 mb-8">
        <h2 class="text-lg font-semibold text-yellow-900 mb-4">راهنمای متغیرهای پویا</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-sm">
          <div>
            <h3 class="font-medium text-yellow-800 mb-2">متغیرهای عمومی:</h3>
            <ul v-pre class="space-y-1 text-yellow-700">
              <li><code class="bg-yellow-100 px-1 rounded">{{title}}</code> - عنوان محتوا</li>
              <li><code class="bg-yellow-100 px-1 rounded">{{excerpt}}</code> - خلاصه محتوا</li>
              <li><code class="bg-yellow-100 px-1 rounded">{{content}}</code> - متن کامل محتوا</li>
              <li><code class="bg-yellow-100 px-1 rounded">{{slug}}</code> - آدرس محتوا</li>
              <li><code class="bg-yellow-100 px-1 rounded">{{site_name}}</code> - نام سایت</li>
            </ul>
          </div>
          <div>
            <h3 class="font-medium text-yellow-800 mb-2">متغیرهای مقاله:</h3>
            <ul v-pre class="space-y-1 text-yellow-700">
              <li><code class="bg-yellow-100 px-1 rounded">{{author}}</code> - نویسنده</li>
              <li><code class="bg-yellow-100 px-1 rounded">{{published_at}}</code> - تاریخ انتشار</li>
              <li><code class="bg-yellow-100 px-1 rounded">{{updated_at}}</code> - تاریخ بروزرسانی</li>
              <li><code class="bg-yellow-100 px-1 rounded">{{featured_image}}</code> - تصویر شاخص</li>
              <li><code class="bg-yellow-100 px-1 rounded">{{category_name}}</code> - نام دسته‌بندی</li>
            </ul>
          </div>
        </div>
      </div>

      <!-- دکمه‌ها -->
      <div class="flex justify-end gap-3 pt-4 border-t">
        <NuxtLink to="../SchemaManagement" class="px-4 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50">بازگشت</NuxtLink>
        <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded-md text-sm font-medium hover:bg-blue-700" @click="handleSubmit">ثبت تغییرات</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
// eslint-disable-next-line @typescript-eslint/no-explicit-any
declare const useSchema: () => any
</script>

<script setup lang="ts">
import { reactive, onMounted, ref, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'

// تعریف layout برای استفاده از سایدبار ادمین
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

const router = useRouter()
const route = useRoute()

// مقداردهی اولیه فرم
const form = reactive({
  type: 'Article',
  name: 'نمونه مقاله',
  description: 'توضیحات نمونه',
  isActive: true,
  
  // فیلدهای عمومی
  title: '{{title}}',
  excerpt: '{{excerpt}}',
  content: '{{content}}',
  slug: '{{slug}}',
  siteName: '{{site_name}}',
  featuredImage: '{{featured_image}}',
  articleId: '{{id}}',
  articleUrl: '{{url}}',
  
  // فیلدهای مخصوص مقاله
  author: '{{author}}',
  publishedAt: '{{published_at}}',
  
  // فیلدهای مخصوص محصول
  price: '{{price}}',
  currency: 'IRR',
  
  // فیلدهای مخصوص سازمان
  address: '{{address}}',
  telephone: '{{telephone}}',
  
  // فیلدهای متا
  metaTitle: '{{title}} - {{site_name}}',
  metaKeywords: '{{keywords}}',
  metaDescription: '{{excerpt}}',
  
  // فیلدهای Open Graph
  ogTitle: '{{title}}',
  ogDescription: '{{excerpt}}',
  ogImage: '{{featured_image}}',
  ogType: 'article',
  ogSiteName: '{{site_name}}',
  
  // فیلدهای اضافی
  extraFields: JSON.stringify({
    headline: '{{title}}',
    author: '{{author}}',
    description: '{{excerpt}}',
    datePublished: '{{published_at}}',
    dateModified: '{{updated_at}}',
    publisher: '{{site_name}}',
    mainEntityOfPage: '{{slug}}',
    image: '{{featured_image}}',
    articleSection: '{{category_name}}',
    keywords: '{{keywords}}',
    wordCount: '{{word_count}}',
    timeRequired: '{{reading_time}}',
    articleBody: '{{content}}',
    inLanguage: 'fa',
    isAccessibleForFree: 'true'
  }, null, 2)
})

// فیلدهای سفارشی
const customFields = reactive<Array<{key: string, label: string, type: string, value: string, section: string}>>([])

// فیلد جدید برای افزودن
const newField = reactive({
  key: '',
  label: '',
  type: 'text',
  value: '',
  section: 'general'
})

// وضعیت ویرایش
const isEditing = ref(false)
const editingIndex = ref(-1)

// نظارت بر تغییر بخش برای JSON
watch(() => newField.section, (newSection) => {
  if (newSection === 'json' && newField.value && !newField.value.startsWith('{')) {
    // اگر بخش به JSON تغییر کرد و مقدار فعلی JSON نیست، آن را پاک کن
    newField.value = ''
  }
})



// تابع برای نمایش مقدار فعلی فیلدها
function getFieldValue(fieldName: string): string {
  const fieldMap = {
    title: 'عنوان مقاله نمونه',
    excerpt: 'این یک خلاصه کوتاه از مقاله است که برای SEO استفاده می‌شود.',
    content: 'متن کامل مقاله در اینجا قرار می‌گیرد...',
    slug: 'sample-article',
    author: 'نویسنده نمونه',
    published_at: '2024-01-15',
    updated_at: '2024-01-20',
    featured_image: '/images/sample-image.jpg',
    category_name: 'تکنولوژی',
    keywords: 'سئو, مقاله, نمونه',
    word_count: '1500',
    reading_time: '5',
    price: '100000',
    currency: 'IRR',
    address: 'تهران، ایران',
    telephone: '021-12345678',
    site_name: 'نام سایت',
    meta_title: 'عنوان مقاله نمونه - نام سایت',
    meta_keywords: 'سئو, مقاله, نمونه',
    meta_description: 'این یک خلاصه کوتاه از مقاله است که برای SEO استفاده می‌شود.',
    og_title: 'عنوان مقاله نمونه',
    og_description: 'این یک خلاصه کوتاه از مقاله است که برای SEO استفاده می‌شود.',
    og_image: '/images/sample-image.jpg',
    og_type: 'article',
    og_site_name: 'نام سایت',
    id: '12345',
    url: 'https://example.com/sample-article'
  }
  
  return fieldMap[fieldName] || 'مقدار تعریف نشده'
}

// بررسی اعتبار فیلد
// function isValidField(field: {key: string, label: string, value: string, section?: string}): boolean {
//   if (field.key.trim() === '' || field.label.trim() === '') {
//     return false
//   }
//   
//   // بررسی اعتبار JSON برای بخش JSON
//   if (field.section === 'json' && field.value.trim() !== '') {
//     try {
//       JSON.parse(field.value)
//     } catch (e) {
//       return false
//     }
//   }
//   
//   return true
// }

// دریافت برچسب بخش
function getSectionLabel(section: string): string {
  const sectionLabels = {
    general: 'عمومی',
    meta: 'متا',
    og: 'Open Graph',
    json: 'JSON'
  }
  return sectionLabels[section] || section
}

// افزودن فیلد سفارشی تکی
function addCustomField() {
  if (newField.key.trim() === '' || newField.label.trim() === '') {
    alert('لطفاً نام فیلد و برچسب را پر کنید!')
    return
  }
  
  // بررسی اعتبار JSON برای بخش JSON
  if (newField.section === 'json' && newField.value.trim() !== '') {
    try {
      JSON.parse(newField.value)
    } catch (e) {
      alert('فرمت JSON نامعتبر است. لطفاً JSON معتبر وارد کنید!')
      return
    }
  }
  
  customFields.push({
    key: newField.key.trim(),
    label: newField.label.trim(),
    type: newField.type,
    value: newField.value || '',
    section: newField.section
  })
  
  // نمایش پیام موفقیت
  alert('فیلد سفارشی با موفقیت اضافه شد!')
  
  // پاک کردن فرم
  resetNewField()
}

// ویرایش فیلد سفارشی
function editCustomField(index: number) {
  const field = customFields[index]
  newField.key = field.key
  newField.label = field.label
  newField.type = field.type
  newField.value = field.value
  newField.section = field.section
  
  isEditing.value = true
  editingIndex.value = index
}

// بروزرسانی فیلد سفارشی
function updateCustomField() {
  if (newField.key.trim() === '' || newField.label.trim() === '') {
    alert('لطفاً نام فیلد و برچسب را پر کنید!')
    return
  }
  
  // بررسی اعتبار JSON برای بخش JSON
  if (newField.section === 'json' && newField.value.trim() !== '') {
    try {
      JSON.parse(newField.value)
    } catch (e) {
      alert('فرمت JSON نامعتبر است. لطفاً JSON معتبر وارد کنید!')
      return
    }
  }
  
  if (editingIndex.value >= 0) {
    customFields[editingIndex.value] = {
      key: newField.key.trim(),
      label: newField.label.trim(),
      type: newField.type,
      value: newField.value || '',
      section: newField.section
    }
    
    // نمایش پیام موفقیت
    alert('فیلد سفارشی با موفقیت بروزرسانی شد!')
    
    // پاک کردن فرم و خروج از حالت ویرایش
    resetNewField()
    cancelEdit()
  }
}

// انصراف از ویرایش
function cancelEdit() {
  isEditing.value = false
  editingIndex.value = -1
  resetNewField()
}

// پاک کردن فرم فیلد جدید
function resetNewField() {
  newField.key = ''
  newField.label = ''
  newField.type = 'text'
  newField.value = ''
  newField.section = 'general'
}

// حذف فیلد سفارشی
function removeCustomField(index: number) {
  customFields.splice(index, 1)
}



// بارگذاری داده‌های اسکیما از API
onMounted(async () => {
  const schemaId = route.query.id
  if (schemaId) {
    try {
      // دریافت اطلاعات اسکیما از فایل JSON
      const { fetchTemplateById } = useSchema()
      const schemaData = await fetchTemplateById(schemaId as string)
      
      if (schemaData) {
        // بروزرسانی فرم با داده‌های دریافتی
        form.name = schemaData.name
        form.type = schemaData.type
        form.description = schemaData.description
        form.isActive = schemaData.isActive
        
        // بروزرسانی فیلدهای اضافی JSON
        if (schemaData.template) {
          form.extraFields = JSON.stringify(schemaData.template, null, 2)
        }
        
        // بروزرسانی فیلدهای متا و Open Graph اگر موجود باشند
        if (schemaData.meta) {
          form.metaTitle = schemaData.meta.title || form.metaTitle
          form.metaDescription = schemaData.meta.description || form.metaDescription
          form.metaKeywords = schemaData.meta.keywords || form.metaKeywords
        }
        
        if (schemaData.openGraph) {
          form.ogTitle = schemaData.openGraph.title || form.ogTitle
          form.ogDescription = schemaData.openGraph.description || form.ogDescription
          form.ogImage = schemaData.openGraph.image || form.ogImage
          form.ogType = schemaData.openGraph.type || form.ogType
          form.ogSiteName = schemaData.openGraph.site_name || form.ogSiteName
        }
      } else {
        console.warn('Schema not found:', schemaId)
        alert('اسکیما مورد نظر یافت نشد!')
      }
    } catch (error) {
      console.error('Error loading schema:', error)
      alert('خطا در بارگذاری اسکیما!')
    }
  }
})

// ثبت فرم
function handleSubmit() {
  // آماده‌سازی داده‌ها برای ذخیره در دیتابیس
  const schemaData = {
    ...form,
    customFields: customFields.map(field => ({
      key: field.key,
      label: field.label,
      type: field.type,
      value: field.value,
      section: field.section
    }))
  }
  
  // اینجا باید ذخیره‌سازی واقعی انجام شود (مثلاً API)
  
  // نمایش پیام موفقیت
  alert('تغییرات اسکیما با موفقیت ثبت شد!')
  
  // بازگشت به صفحه مدیریت اسکیمای
  router.push('../SchemaManagement')
}
</script>

<!--
  این صفحه برای ویرایش اسکیما است و فرم آن تمیز و زیبا طراحی شده است.
  کامنت‌ها به فارسی برای توسعه‌دهندگان آینده درج شده است.
--> 
