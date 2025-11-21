<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-end justify-center z-50">
    <div class="bg-white rounded-t-lg shadow-xl w-full max-w-md max-h-[80vh] overflow-hidden">
      <!-- هدر کشو -->
      <div class="flex items-center justify-between p-6 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">تنظیمات آیتم</h3>
        <button
          class="text-gray-400 hover:text-gray-600 transition-colors"
          @click="$emit('close')"
        >
          <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <!-- محتوای کشو -->
      <div class="p-6 overflow-y-auto max-h-[60vh]">
        <div class="space-y-4">
          <!-- عرض آیتم -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">عرض آیتم (%)</label>
            <input
              v-model="local.width"
              type="number"
              min="1"
              max="100"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500 mt-1">درصد عرض آیتم در لایه</p>
          </div>

          <!-- چینش آیتم -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">چینش آیتم</label>
            <select
              v-model="local.align"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="left">چپ</option>
              <option value="center">وسط</option>
              <option value="right">راست</option>
            </select>
            <p class="text-xs text-gray-500 mt-1">موقعیت آیتم در لایه</p>
          </div>

          <!-- پدینگ راست -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">پدینگ راست (px)</label>
            <input
              v-model="local.paddingRight"
              type="number"
              min="0"
              max="100"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500 mt-1">فاصله داخلی از سمت راست</p>
          </div>

          <!-- پدینگ چپ -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">پدینگ چپ (px)</label>
            <input
              v-model="local.paddingLeft"
              type="number"
              min="0"
              max="100"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500 mt-1">فاصله داخلی از سمت چپ</p>
          </div>

          <!-- تنظیمات خاص آیتم -->
          <div v-if="item.id === 'logo'" class="logo-specific-settings">
            <h4 class="font-medium text-gray-900 mb-3">تنظیمات لوگو</h4>
            
            <!-- لینک لوگو -->
            <div class="mb-3">
              <label class="block text-sm font-medium text-gray-700 mb-2">لینک لوگو</label>
              <input
                v-model="local.props.logoLink"
                type="url"
                placeholder="https://example.com"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <!-- متن جایگزین -->
            <div class="mb-3">
              <label class="block text-sm font-medium text-gray-700 mb-2">متن جایگزین</label>
              <input
                v-model="local.props.logoAlt"
                type="text"
                placeholder="نام سایت"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>

          <div v-else-if="item.id === 'links'" class="links-specific-settings">
            <h4 class="font-medium text-gray-900 mb-3">تنظیمات لینک‌ها</h4>
            
            <!-- لینک‌های مفید -->
            <div class="space-y-2">
              <div
                v-for="(link, index) in local.props.links"
                :key="index"
                class="flex items-center space-x-2 space-x-reverse"
              >
                <input
                  v-model="link.title"
                  type="text"
                  placeholder="عنوان لینک"
                  class="flex-1 px-2 py-1 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-blue-500"
                />
                <input
                  v-model="link.url"
                  type="url"
                  placeholder="لینک"
                  class="flex-1 px-2 py-1 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-blue-500"
                />
                <button
                  class="text-red-500 hover:text-red-700 text-sm"
                  @click="removeLink(index)"
                >
                  <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
              
              <button
                class="w-full px-3 py-2 text-sm border border-dashed border-gray-300 rounded-md text-gray-500 hover:border-gray-400 hover:text-gray-600 transition-colors"
                @click="addLink"
              >
                + افزودن لینک جدید
              </button>
            </div>
          </div>

          <div v-else-if="item.id === 'contact'" class="contact-specific-settings">
            <h4 class="font-medium text-gray-900 mb-3">تنظیمات اطلاعات تماس</h4>
            
            <div class="p-4 bg-blue-50 border border-blue-200 rounded-md text-sm text-blue-800">
              <p class="font-medium mb-1">📌 نکته مهم</p>
              <p>اطلاعات تماس (شماره تلفن، آدرس، ایمیل) از بخش <strong>تنظیمات سایت</strong> خوانده می‌شود.</p>
              <p class="mt-2">برای ویرایش این اطلاعات به منوی <strong>تنظیمات → اطلاعات شرکت</strong> مراجعه کنید.</p>
            </div>
          </div>

          <div v-else-if="item.id === 'about'" class="about-specific-settings">
            <h4 class="font-medium text-gray-900 mb-3">تنظیمات درباره ما</h4>

            <div class="p-4 bg-blue-50 border border-blue-200 rounded-md text-sm text-blue-800">
              <p class="font-medium mb-1">نکته</p>
              <p>متن این ویجت از فیلد <strong>توضیحات کوتاه / درباره ما</strong> در صفحه <strong>تنظیمات → تنظیمات خانه</strong> خوانده می‌شود.</p>
              <p class="mt-2">برای تغییر متن به آن بخش از تنظیمات مراجعه کنید.</p>
            </div>
          </div>

          <div v-else-if="item.id === 'working-hours'" class="working-hours-specific-settings">
            <h4 class="font-medium text-gray-900 mb-3">تنظیمات ساعات کاری</h4>
            
            <div class="p-4 bg-blue-50 border border-blue-200 rounded-md text-sm text-blue-800">
              <p class="font-medium mb-1">📌 نکته مهم</p>
              <p>ساعات کاری از بخش <strong>تنظیمات سایت</strong> خوانده می‌شود.</p>
              <p class="mt-2">برای ویرایش ساعات کاری به منوی <strong>تنظیمات → ساعات کاری</strong> مراجعه کنید.</p>
            </div>
          </div>

          <div v-else-if="item.id === 'newsletter'" class="newsletter-specific-settings">
            <h4 class="font-medium text-gray-900 mb-3">تنظیمات خبرنامه</h4>
```            
            <!-- متن دکمه -->
            <div class="mb-3">
              <label class="block text-sm font-medium text-gray-700 mb-2">متن دکمه</label>
              <input
                v-model="local.props.buttonText"
                type="text"
                placeholder="عضویت"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <!-- متن راهنما -->
            <div class="mb-3">
              <label class="block text-sm font-medium text-gray-700 mb-2">متن راهنما</label>
              <input
                v-model="local.props.placeholder"
                type="text"
                placeholder="ایمیل خود را وارد کنید..."
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>

          <div v-else-if="item.id === 'trust'" class="trust-specific-settings">
            <h4 class="font-medium text-gray-900 mb-4">نشان‌های اعتماد (اینماد، ساماندهی و...)</h4>

            <div class="space-y-4">
              <div
                v-for="(badge, index) in local.props.trustBadges"
                :key="badge.id || index"
                class="p-4 border border-gray-200 rounded-lg space-y-3"
              >
                <div class="flex items-center justify-between">
                  <input
                    v-model="badge.title"
                    type="text"
                    placeholder="عنوان نماد (مثلاً: اینماد)"
                    class="flex-1 px-3 py-2 text-sm border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                  />

                  <button
                    v-if="local.props.trustBadges.length > 1"
                    type="button"
                    class="mr-2 p-2 text-red-500 hover:text-red-600"
                    title="حذف نماد"
                    @click="removeTrustBadge(index)"
                  >
                    <svg class="h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <line x1="18" y1="6" x2="6" y2="18" />
                      <line x1="6" y1="6" x2="18" y2="18" />
                    </svg>
                  </button>
                </div>

                <div>
                  <label class="block text-xs font-medium text-gray-700 mb-1">کد HTML نماد</label>
                  <textarea
                    v-model="badge.htmlCode"
                    rows="4"
                    placeholder='<div><a href="..."><img src="..." alt="اینماد"></a></div>'
                    class="w-full px-3 py-2 text-sm border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 font-mono text-xs"
                  ></textarea>
                  <p class="text-xs text-gray-500 mt-1">کد HTML دریافتی از سایت نماد را اینجا قرار دهید</p>
                </div>

                <div>
                  <label class="block text-xs font-medium text-gray-700 mb-1">توضیحات (اختیاری)</label>
                  <input
                    v-model="badge.description"
                    type="text"
                    placeholder="توضیحات کوتاه"
                    class="w-full px-3 py-2 text-sm border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                  />
                </div>
              </div>

              <button
                type="button"
                class="w-full px-3 py-2 text-sm border border-dashed border-gray-300 rounded-md text-gray-600 hover:border-gray-400 hover:text-gray-700"
                @click="addTrustBadge"
              >
                + افزودن نماد جدید
              </button>
            </div>
          </div>

          <div v-else-if="item.id === 'menu'" class="menu-specific-settings">
            <h4 class="font-medium text-gray-900 mb-4">انتخاب منو</h4>

            <div class="space-y-3">
              <div v-if="menusLoading" class="text-center text-gray-500 text-sm py-4">
                در حال بارگذاری منوها...
              </div>
              <div v-else-if="menus && menus.length === 0" class="text-center text-gray-500 text-sm py-4">
                هیچ منویی یافت نشد. ابتدا از پنل ادمین یک منو ایجاد کنید.
              </div>
              <div v-else class="space-y-2">
                <div
                  v-for="menu in menus"
                  :key="menu.id"
                  class="flex items-center p-3 border border-gray-200 rounded-lg hover:bg-gray-50"
                >
                  <input
                    :id="`menu-${menu.id}`"
                    type="checkbox"
                    :value="menu.id"
                    :checked="isMenuSelected(menu.id)"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                    @change="toggleMenuSelection(menu.id)"
                  />
                  <label :for="`menu-${menu.id}`" class="mr-3 flex-1 text-sm cursor-pointer">
                    <div class="font-medium text-gray-900">{{ menu.name }}</div>
                    <div v-if="menu.items && menu.items.length > 0" class="text-xs text-gray-500">
                      {{ menu.items.length }} آیتم
                    </div>
                  </label>
                </div>
              </div>
            </div>
          </div>

          <div v-else-if="isSocialItem" class="social-specific-settings">
            <h4 class="font-medium text-gray-900 mb-3">تنظیمات شبکه‌های اجتماعی</h4>

            <div class="p-4 bg-blue-50 border border-blue-200 rounded-md text-sm text-blue-800 space-y-2">
              <p class="font-medium">📌 نکته مهم</p>
              <p>
                لینک‌ها و وضعیت شبکه‌های اجتماعی از بخش
                <strong>تنظیمات → شبکه‌های اجتماعی</strong>
                خوانده می‌شود.
              </p>
              <p class="leading-6">
                برای افزودن، حذف یا ویرایش هر شبکه ابتدا به آن بخش از تنظیمات بروید و سپس اینجا فقط چینش ظاهری ویجت را تعیین کنید.
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- فوتر کشو -->
      <div class="flex justify-end space-x-3 space-x-reverse p-6 border-t border-gray-200">
        <button
          class="px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="$emit('close')"
        >
          انصراف
        </button>
        <button
          class="px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="saveSettings"
        >
          ذخیره
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'

interface SocialItem {
  id: string
  platform: string
  label: string
  url: string
  enabled: boolean
  openInNewTab: boolean
  title?: string
  href?: string
  visible?: boolean
  value?: string
}

interface LinkItem {
  title: string
  url: string
}

interface TrustBadge {
  id: string
  title: string
  description: string
  htmlCode: string
  icon: string
}

interface ItemProps {
  logoLink?: string
  logoAlt?: string
  links?: LinkItem[]
  title?: string
  description?: string
  note?: string
  buttonText?: string
  placeholder?: string
  instagram?: string
  telegram?: string
  twitter?: string
  linkedin?: string
  youtube?: string
  facebook?: string
  whatsapp?: string
  socials?: SocialItem[]
  trustBadges?: TrustBadge[]
  [key: string]: unknown
}

interface ItemConfig {
  menuIds?: number[]
  [key: string]: unknown
}

interface Item {
  id?: string
  type?: string
  component?: string
  widget?: string
  widgetKey?: string
  width?: number
  align?: string
  paddingRight?: number
  paddingLeft?: number
  bgColor?: string
  config?: ItemConfig
  props?: ItemProps
  [key: string]: unknown
}

interface Props {
  item: Item
}

const props = defineProps<Props>()

// Emit events
const emit = defineEmits<{
  close: []
  save: [item: Item]
}>()

interface Menu {
  id: number | string
  name?: string
  [key: string]: unknown
}

// منوها
const menus = ref<Menu[]>([])
const menusLoading = ref(false)

// تنظیمات شرکت برای مقادیر پیش‌فرض
const companySettings = ref<Record<string, string>>({
  phone: '',
  address: '',
  email: ''
})

// بارگذاری تنظیمات شرکت
interface Setting {
  key: string
  value: string
  [key: string]: unknown
}

interface SettingsResponse {
  success?: boolean
  data?: Setting[]
  [key: string]: unknown
}

async function fetchCompanySettings() {
  try {
    const response = await $fetch<SettingsResponse>('/api/admin/settings?category=company')
    if (response && response.success && Array.isArray(response.data)) {
      response.data.forEach((setting: Setting) => {
        if (setting.key === 'company_phone' || setting.key === 'footer_phone_number' || setting.key === 'header_phone_number') {
          companySettings.value.phone = setting.value
        } else if (setting.key === 'company_address') {
          companySettings.value.address = setting.value
        } else if (setting.key === 'company_email' || setting.key === 'support_email') {
          companySettings.value.email = setting.value
        }
      })
    }
  } catch (_error) {
    // console.error('خطا در بارگذاری تنظیمات شرکت:', error)
  }
}

async function fetchMenus() {
  try {
    menusLoading.value = true
    const data = await $fetch('/api/admin/menus')
    if (data) {
      menus.value = Array.isArray(data) ? data : []
    }
  } catch (_err) {
    // console.error('Error loading menus:', err)
    menus.value = []
  } finally {
    menusLoading.value = false
  }
}

function isMenuSelected(menuId: number) {
  if (!local.value.config?.menuIds) return false
  return local.value.config.menuIds.includes(menuId)
}

function toggleMenuSelection(menuId: number) {
  if (!local.value.config) {
    local.value.config = { menuIds: [] }
  }
  if (!local.value.config.menuIds) {
    local.value.config.menuIds = []
  }
  
  const index = local.value.config.menuIds.indexOf(menuId)
  if (index > -1) {
    local.value.config.menuIds.splice(index, 1)
  } else {
    local.value.config.menuIds.push(menuId)
  }
}

const SOCIAL_PRESETS = [
  { platform: 'instagram', label: 'اینستاگرام', icon: '📷' },
  { platform: 'telegram', label: 'تلگرام', icon: '📱' },
  { platform: 'twitter', label: 'توییتر', icon: '🐦' },
  { platform: 'linkedin', label: 'لینکدین', icon: '💼' },
  { platform: 'youtube', label: 'یوتیوب', icon: '📺' },
  { platform: 'facebook', label: 'فیسبوک', icon: '📘' },
  { platform: 'whatsapp', label: 'واتساپ', icon: '💬' }
]

const createDefaultSocials = () => SOCIAL_PRESETS.map(preset => ({
  id: preset.platform,
  platform: preset.platform,
  label: preset.label,
  url: '',
  enabled: false,
  openInNewTab: true
}))

function findPresetLabel(platform?: string) {
  if (!platform) return ''
  const preset = SOCIAL_PRESETS.find(item => item.platform === platform)
  return preset ? preset.label : ''
}

const normalizeSocials = (source: ItemProps = {}) => {
  const list: SocialItem[] = Array.isArray(source.socials) ? source.socials : []

  if (list.length) {
    return list.map((item, index) => ({
      id: item.id ?? `${item.platform || 'social'}-${index}`,
      platform: item.platform || 'custom',
      label: item.label || item.title || findPresetLabel(item.platform) || 'شبکه اجتماعی',
      url: item.url || item.href || '',
      enabled: item.enabled !== false && item.visible !== false,
      openInNewTab: item.openInNewTab !== false
    }))
  }

  // fallback به فیلدهای قدیمی
  const legacy = [
    { platform: 'instagram', value: source.instagram },
    { platform: 'telegram', value: source.telegram },
    { platform: 'twitter', value: source.twitter },
    { platform: 'linkedin', value: source.linkedin },
    { platform: 'youtube', value: source.youtube },
    { platform: 'facebook', value: source.facebook },
    { platform: 'whatsapp', value: source.whatsapp }
  ]

  const fromLegacy = legacy
    .filter(entry => typeof entry.value === 'string' && entry.value.trim().length)
    .map(entry => ({
      id: entry.platform,
      platform: entry.platform,
      label: findPresetLabel(entry.platform) || entry.platform,
      url: entry.value,
      enabled: true,
      openInNewTab: true
    }))

  return fromLegacy.length ? fromLegacy : createDefaultSocials()
}

// کپی محلی از آیتم برای ویرایش
const local = ref<Item>({
  width: 25,
  align: 'center',
  paddingRight: 0,
  paddingLeft: 0,
  bgColor: 'transparent',
  config: {
    menuIds: []
  },
  props: {
    logoLink: '',
    logoAlt: '',
    links: [
      { title: 'درباره ما', url: '/about' },
      { title: 'تماس با ما', url: '/contact' }
    ],
    title: '',
    description: '',
    note: '',
    buttonText: 'عضویت',
    placeholder: 'ایمیل خود را وارد کنید...',
    instagram: '',
    telegram: '',
    twitter: '',
    linkedin: '',
    youtube: '',
    facebook: '',
    whatsapp: '',
    socials: createDefaultSocials(),
    trustBadges: []
  }
})

const resolveItemType = (item: Item): string => {
  if (!item) return ''
  const raw = (item.type || item.component || item.id || item.widget || item.widgetKey || '')
    .toString()
    .toLowerCase()

  return raw.replace(/[^a-z]/g, '')
}

const isSocialItem = computed(() => {
  const type = resolveItemType(props.item)
  if (!type) return false
  return type === 'social' || type.endsWith('social') || type.includes('footerwidgetsocial') || type.includes('socialwidget')
})

const applyItemToLocal = (item: Item) => {
  if (!item) {
    local.value = {
      ...local.value,
      props: {
        ...local.value.props,
        socials: createDefaultSocials()
      }
    }
    return
  }

  local.value = {
    width: item.width ?? 25,
    align: item.align || 'center',
    paddingRight: item.paddingRight || 0,
    paddingLeft: item.paddingLeft || 0,
    bgColor: item.bgColor || 'transparent',
    config: item.config || { menuIds: [] },
    props: {
      logoLink: item.props?.logoLink || '',
      logoAlt: item.props?.logoAlt || '',
      links: item.props?.links || [
        { title: 'درباره ما', url: '/about' },
        { title: 'تماس با ما', url: '/contact' }
      ],
      title: item.props?.title || '',
      description: item.props?.description || '',
      note: item.props?.note || '',
      buttonText: item.props?.buttonText || 'عضویت',
      placeholder: item.props?.placeholder || 'ایمیل خود را وارد کنید...',
      instagram: item.props?.instagram || '',
      telegram: item.props?.telegram || '',
      twitter: item.props?.twitter || '',
      linkedin: item.props?.linkedin || '',
      youtube: item.props?.youtube || '',
      facebook: item.props?.facebook || '',
      whatsapp: item.props?.whatsapp || '',
      socials: normalizeSocials(item.props),
      trustBadges: item.props?.trustBadges || []
    }
  }
}

// افزودن لینک جدید
const addLink = () => {
  if (!local.value.props) local.value.props = {}
  if (!local.value.props.links) local.value.props.links = []
  local.value.props.links.push({ title: '', url: '' })
}

// حذف لینک
const removeLink = (index: number) => {
  if (local.value.props?.links) {
    local.value.props.links.splice(index, 1)
  }
}

const _addSocial = () => {
  if (!local.value.props) local.value.props = {}
  if (!local.value.props.socials) local.value.props.socials = []
  local.value.props.socials.push({
    id: `custom-${Date.now()}`,
    platform: 'custom',
    label: 'شبکه جدید',
    url: '',
    enabled: true,
    openInNewTab: true
  })
}

const _removeSocial = (_index: number) => {
  if (local.value.props?.socials && local.value.props.socials.length > 1) {
    local.value.props.socials.splice(index, 1)
  }
}

// مدیریت نشان‌های اعتماد
const addTrustBadge = () => {
  if (!local.value.props) local.value.props = {}
  if (!local.value.props.trustBadges) {
    local.value.props.trustBadges = []
  }
  local.value.props.trustBadges.push({
    id: `badge-${Date.now()}`,
    title: 'نماد جدید',
    description: '',
    htmlCode: '',
    icon: '🏆'
  })
}

const removeTrustBadge = (index: number) => {
  if (local.value.props?.trustBadges && local.value.props.trustBadges.length > 1) {
    local.value.props.trustBadges.splice(index, 1)
  }
}

// مقداردهی اولیه trustBadges اگر خالی است
const ensureTrustBadges = () => {
  if (props.item?.id === 'trust' && (!local.value.props?.trustBadges || local.value.props.trustBadges.length === 0)) {
    if (!local.value.props) local.value.props = {}
    local.value.props.trustBadges = [{
      id: `badge-${Date.now()}`,
      title: 'نماد اعتماد',
      description: 'مثال: اینماد یا ساماندهی',
      htmlCode: '',
      icon: '🏆'
    }]
  }
}

// فراخوانی ensureTrustBadges بعد از mount
onMounted(async () => {
  // بارگذاری تنظیمات شرکت برای مقادیر پیش‌فرض
  await fetchCompanySettings()
  
  applyItemToLocal(props.item)
  ensureTrustBadges()
  // بارگذاری منوها برای انتخاب
  if (props.item?.id === 'menu') {
    fetchMenus()
  }
})

watch(() => props.item, async (newItem) => {
  // بارگذاری تنظیمات شرکت اگر هنوز بارگذاری نشده
  if (!companySettings.value.phone && !companySettings.value.address && !companySettings.value.email) {
    await fetchCompanySettings()
  }
  
  applyItemToLocal(newItem)
  ensureTrustBadges()
  // بارگذاری منوها اگر تغییر به menu
  if (newItem?.id === 'menu') {
    fetchMenus()
  }
})

const _resetSocials = () => {
  local.value.props.socials = createDefaultSocials()
}

const sanitizeSocialsForSave = (list: SocialItem[]) => {
  return (Array.isArray(list) ? list : [])
    .map((item, index) => {
      const platform = (item.platform || '').toLowerCase()
      const url = typeof item.url === 'string' ? item.url.trim() : ''
      const label = item.label || findPresetLabel(platform) || 'شبکه اجتماعی'
      return {
        id: item.id ?? `${platform || 'social'}-${index}`,
        platform: platform || 'custom',
        label,
        title: label,
        url,
        href: url,
        enabled: item.enabled !== false,
        visible: item.enabled !== false,
        openInNewTab: item.openInNewTab !== false
      }
    })
    .filter(item => item.url || item.label)
}

const prepareItemForSave = (current: Item) => {
  if (isSocialItem.value) {
    return {
      ...current,
      config: current.config || {},
      props: {
        ...current.props,
        socials: Array.isArray(current.props?.socials) ? current.props?.socials : []
      }
    }
  }

  const socials = sanitizeSocialsForSave(current.props?.socials || [])

  const legacy: Record<string, string> = {}
  socials.forEach(item => {
    if (item.enabled && item.url && !legacy[item.platform]) {
      legacy[item.platform] = item.url
    }
  })

  return {
    ...current,
    config: current.config || {},
    props: {
      ...current.props,
      socials,
      instagram: legacy.instagram || current.props?.instagram || '',
      telegram: legacy.telegram || current.props?.telegram || '',
      twitter: legacy.twitter || current.props?.twitter || '',
      linkedin: legacy.linkedin || current.props?.linkedin || '',
      youtube: legacy.youtube || current.props?.youtube || '',
      facebook: legacy.facebook || current.props?.facebook || '',
      whatsapp: legacy.whatsapp || current.props?.whatsapp || ''
    }
  }
}

// ذخیره تنظیمات
const saveSettings = () => {
  const prepared = prepareItemForSave({
    ...local.value,
    props: { ...local.value.props }
  })
  emit('save', prepared)
}
</script>
