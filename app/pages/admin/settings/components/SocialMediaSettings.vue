<template>
  <div v-if="hasAccess" class="settings-section">
    <div class="settings-header">
      <h2>تنظیمات شبکه‌های اجتماعی</h2>
      <p class="settings-description">
        آدرس صفحات شبکه‌های اجتماعی فروشگاه را مدیریت کنید
      </p>
    </div>

    <div class="social-media-form">
      <!-- Instagram -->
      <div class="social-media-item">
        <div class="social-header">
          <div class="social-title">
            <span class="i-mdi-instagram text-2xl text-pink-500"></span>
            <span class="social-name">اینستاگرام</span>
          </div>
          <label class="toggle-switch">
            <input v-model="localSettings.instagram_enabled" type="checkbox" />
            <span class="toggle-slider"></span>
          </label>
        </div>
        <input
          v-model="localSettings.instagram"
          type="url"
          class="form-input"
          :disabled="!localSettings.instagram_enabled"
          placeholder="https://instagram.com/username"
          dir="ltr"
        />
      </div>

      <!-- Telegram -->
      <div class="social-media-item">
        <div class="social-header">
          <div class="social-title">
            <span class="i-mdi-telegram text-2xl text-blue-400"></span>
            <span class="social-name">تلگرام</span>
          </div>
          <label class="toggle-switch">
            <input v-model="localSettings.telegram_enabled" type="checkbox" />
            <span class="toggle-slider"></span>
          </label>
        </div>
        <input
          v-model="localSettings.telegram"
          type="url"
          class="form-input"
          :disabled="!localSettings.telegram_enabled"
          placeholder="https://t.me/username"
          dir="ltr"
        />
      </div>

      <!-- WhatsApp -->
      <div class="social-media-item">
        <div class="social-header">
          <div class="social-title">
            <span class="i-mdi-whatsapp text-2xl text-green-500"></span>
            <span class="social-name">واتساپ</span>
          </div>
          <label class="toggle-switch">
            <input v-model="localSettings.whatsapp_enabled" type="checkbox" />
            <span class="toggle-slider"></span>
          </label>
        </div>
        <input
          v-model="localSettings.whatsapp"
          type="url"
          class="form-input"
          :disabled="!localSettings.whatsapp_enabled"
          placeholder="https://wa.me/989123456789"
          dir="ltr"
        />
      </div>

      <!-- Twitter (X) -->
      <div class="social-media-item">
        <div class="social-header">
          <div class="social-title">
            <span class="i-mdi-twitter text-2xl text-blue-400"></span>
            <span class="social-name">توییتر (X)</span>
          </div>
          <label class="toggle-switch">
            <input v-model="localSettings.twitter_enabled" type="checkbox" />
            <span class="toggle-slider"></span>
          </label>
        </div>
        <input
          v-model="localSettings.twitter"
          type="url"
          class="form-input"
          :disabled="!localSettings.twitter_enabled"
          placeholder="https://twitter.com/username"
          dir="ltr"
        />
      </div>

      <!-- LinkedIn -->
      <div class="social-media-item">
        <div class="social-header">
          <div class="social-title">
            <span class="i-mdi-linkedin text-2xl text-blue-600"></span>
            <span class="social-name">لینکدین</span>
          </div>
          <label class="toggle-switch">
            <input v-model="localSettings.linkedin_enabled" type="checkbox" />
            <span class="toggle-slider"></span>
          </label>
        </div>
        <input
          v-model="localSettings.linkedin"
          type="url"
          class="form-input"
          :disabled="!localSettings.linkedin_enabled"
          placeholder="https://linkedin.com/in/username"
          dir="ltr"
        />
      </div>

      <!-- Facebook -->
      <div class="social-media-item">
        <div class="social-header">
          <div class="social-title">
            <span class="i-mdi-facebook text-2xl text-blue-500"></span>
            <span class="social-name">فیسبوک</span>
          </div>
          <label class="toggle-switch">
            <input v-model="localSettings.facebook_enabled" type="checkbox" />
            <span class="toggle-slider"></span>
          </label>
        </div>
        <input
          v-model="localSettings.facebook"
          type="url"
          class="form-input"
          :disabled="!localSettings.facebook_enabled"
          placeholder="https://facebook.com/username"
          dir="ltr"
        />
      </div>

      <!-- YouTube -->
      <div class="social-media-item">
        <div class="social-header">
          <div class="social-title">
            <span class="i-mdi-youtube text-2xl text-red-500"></span>
            <span class="social-name">یوتیوب</span>
          </div>
          <label class="toggle-switch">
            <input v-model="localSettings.youtube_enabled" type="checkbox" />
            <span class="toggle-slider"></span>
          </label>
        </div>
        <input
          v-model="localSettings.youtube"
          type="url"
          class="form-input"
          :disabled="!localSettings.youtube_enabled"
          placeholder="https://youtube.com/@username"
          dir="ltr"
        />
      </div>

      <!-- Aparat -->
      <div class="social-media-item">
        <div class="social-header">
          <div class="social-title">
            <span class="i-heroicons-play-circle text-2xl text-orange-500"></span>
            <span class="social-name">آپارات</span>
          </div>
          <label class="toggle-switch">
            <input v-model="localSettings.aparat_enabled" type="checkbox" />
            <span class="toggle-slider"></span>
          </label>
        </div>
        <input
          v-model="localSettings.aparat"
          type="url"
          class="form-input"
          :disabled="!localSettings.aparat_enabled"
          placeholder="https://aparat.com/username"
          dir="ltr"
        />
      </div>

      <!-- Rubika -->
      <div class="social-media-item">
        <div class="social-header">
          <div class="social-title">
            <span class="i-heroicons-chat-bubble-left-right text-2xl text-purple-500"></span>
            <span class="social-name">روبیکا</span>
          </div>
          <label class="toggle-switch">
            <input v-model="localSettings.rubika_enabled" type="checkbox" />
            <span class="toggle-slider"></span>
          </label>
        </div>
        <input
          v-model="localSettings.rubika"
          type="url"
          class="form-input"
          :disabled="!localSettings.rubika_enabled"
          placeholder="https://rubika.ir/username"
          dir="ltr"
        />
      </div>

      <!-- Eitaa -->
      <div class="social-media-item">
        <div class="social-header">
          <div class="social-title">
            <span class="i-heroicons-chat-bubble-left text-2xl text-teal-500"></span>
            <span class="social-name">ایتا</span>
          </div>
          <label class="toggle-switch">
            <input v-model="localSettings.eitaa_enabled" type="checkbox" />
            <span class="toggle-slider"></span>
          </label>
        </div>
        <input
          v-model="localSettings.eitaa"
          type="url"
          class="form-input"
          :disabled="!localSettings.eitaa_enabled"
          placeholder="https://eitaa.com/username"
          dir="ltr"
        />
      </div>

      <!-- Bale -->
      <div class="social-media-item">
        <div class="social-header">
          <div class="social-title">
            <span class="i-heroicons-device-phone-mobile text-2xl text-cyan-500"></span>
            <span class="social-name">بله</span>
          </div>
          <label class="toggle-switch">
            <input v-model="localSettings.bale_enabled" type="checkbox" />
            <span class="toggle-slider"></span>
          </label>
        </div>
        <input
          v-model="localSettings.bale"
          type="url"
          class="form-input"
          :disabled="!localSettings.bale_enabled"
          placeholder="https://ble.ir/username"
          dir="ltr"
        />
      </div>

      <!-- TikTok -->
      <div class="social-media-item">
        <div class="social-header">
          <div class="social-title">
            <span class="i-mdi-music text-2xl text-gray-800"></span>
            <span class="social-name">تیک‌تاک</span>
          </div>
          <label class="toggle-switch">
            <input v-model="localSettings.tiktok_enabled" type="checkbox" />
            <span class="toggle-slider"></span>
          </label>
        </div>
        <input
          v-model="localSettings.tiktok"
          type="url"
          class="form-input"
          :disabled="!localSettings.tiktok_enabled"
          placeholder="https://tiktok.com/@username"
          dir="ltr"
        />
      </div>

      <!-- Pinterest -->
      <div class="social-media-item">
        <div class="social-header">
          <div class="social-title">
            <span class="i-mdi-pinterest text-2xl text-red-600"></span>
            <span class="social-name">پینترست</span>
          </div>
          <label class="toggle-switch">
            <input v-model="localSettings.pinterest_enabled" type="checkbox" />
            <span class="toggle-slider"></span>
          </label>
        </div>
        <input
          v-model="localSettings.pinterest"
          type="url"
          class="form-input"
          :disabled="!localSettings.pinterest_enabled"
          placeholder="https://pinterest.com/username"
          dir="ltr"
        />
      </div>

      <!-- Discord -->
      <div class="social-media-item">
        <div class="social-header">
          <div class="social-title">
            <span class="i-mdi-discord text-2xl text-indigo-500"></span>
            <span class="social-name">دیسکورد</span>
          </div>
          <label class="toggle-switch">
            <input v-model="localSettings.discord_enabled" type="checkbox" />
            <span class="toggle-slider"></span>
          </label>
        </div>
        <input
          v-model="localSettings.discord"
          type="url"
          class="form-input"
          :disabled="!localSettings.discord_enabled"
          placeholder="https://discord.gg/invite"
          dir="ltr"
        />
      </div>

      <!-- GitHub -->
      <div class="social-media-item">
        <div class="social-header">
          <div class="social-title">
            <span class="i-mdi-github text-2xl text-gray-800"></span>
            <span class="social-name">گیت‌هاب</span>
          </div>
          <label class="toggle-switch">
            <input v-model="localSettings.github_enabled" type="checkbox" />
            <span class="toggle-slider"></span>
          </label>
        </div>
        <input
          v-model="localSettings.github"
          type="url"
          class="form-input"
          :disabled="!localSettings.github_enabled"
          placeholder="https://github.com/username"
          dir="ltr"
        />
      </div>

      <!-- Custom Links Section -->
      <div class="custom-links-section mt-8">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-semibold">لینک‌های سفارشی</h3>
          <button class="btn-add" type="button" @click="addCustomLink">
            <span class="i-heroicons-plus"></span>
            افزودن لینک
          </button>
        </div>

        <div v-for="(link, index) in localSettings.customLinks" :key="index" class="custom-link-item">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div class="form-group">
              <label class="form-label">عنوان</label>
              <input
                v-model="link.title"
                type="text"
                class="form-input"
                placeholder="نام شبکه اجتماعی"
              />
            </div>
            <div class="form-group">
              <label class="form-label">آدرس (URL)</label>
              <div class="flex gap-2">
                <input
                  v-model="link.url"
                  type="url"
                  class="form-input flex-1"
                  placeholder="https://example.com/username"
                  dir="ltr"
                />
                <button 
                  class="btn-remove" 
                  type="button"
                  @click="removeCustomLink(index)"
                >
                  <span class="i-heroicons-trash"></span>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Action Buttons -->
    <div class="settings-actions">
      <button 
        :disabled="saving" 
        class="btn-primary" 
        type="button"
        @click="handleSave"
      >
        <span v-if="!saving">ذخیره تنظیمات</span>
        <span v-else class="flex items-center gap-2">
          <span class="i-heroicons-arrow-path animate-spin"></span>
          در حال ذخیره...
        </span>
      </button>
      <button 
        :disabled="saving" 
        class="btn-secondary" 
        type="button"
        @click="handleReset"
      >
        بازنشانی
      </button>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, reactive, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

// احراز هویت
const { user, isAuthenticated } = useAuth();

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false;
  }

  const userRole = user.value?.role?.toLowerCase() || '';
  const adminRoles = ['admin', 'developer'];
  return adminRoles.includes(userRole);
});

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async (): Promise<void> => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false });
  }
};

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth();
});

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth();
  }
});

const props = defineProps({
  settings: {
    type: Object,
    required: true
  },
  saving: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['save', 'reset'])

// Local reactive copy of settings
const localSettings = reactive({
  instagram: props.settings.instagram || '',
  instagram_enabled: props.settings.instagram_enabled ?? true,
  telegram: props.settings.telegram || '',
  telegram_enabled: props.settings.telegram_enabled ?? true,
  whatsapp: props.settings.whatsapp || '',
  whatsapp_enabled: props.settings.whatsapp_enabled ?? true,
  twitter: props.settings.twitter || '',
  twitter_enabled: props.settings.twitter_enabled ?? true,
  linkedin: props.settings.linkedin || '',
  linkedin_enabled: props.settings.linkedin_enabled ?? true,
  facebook: props.settings.facebook || '',
  facebook_enabled: props.settings.facebook_enabled ?? true,
  youtube: props.settings.youtube || '',
  youtube_enabled: props.settings.youtube_enabled ?? true,
  aparat: props.settings.aparat || '',
  aparat_enabled: props.settings.aparat_enabled ?? true,
  rubika: props.settings.rubika || '',
  rubika_enabled: props.settings.rubika_enabled ?? true,
  eitaa: props.settings.eitaa || '',
  eitaa_enabled: props.settings.eitaa_enabled ?? true,
  bale: props.settings.bale || '',
  bale_enabled: props.settings.bale_enabled ?? true,
  tiktok: props.settings.tiktok || '',
  tiktok_enabled: props.settings.tiktok_enabled ?? true,
  pinterest: props.settings.pinterest || '',
  pinterest_enabled: props.settings.pinterest_enabled ?? true,
  discord: props.settings.discord || '',
  discord_enabled: props.settings.discord_enabled ?? true,
  github: props.settings.github || '',
  github_enabled: props.settings.github_enabled ?? true,
  customLinks: props.settings.customLinks || []
})

// Watch for external changes
watch(() => props.settings, (newSettings) => {
  Object.assign(localSettings, {
    instagram: newSettings.instagram || '',
    instagram_enabled: newSettings.instagram_enabled ?? true,
    telegram: newSettings.telegram || '',
    telegram_enabled: newSettings.telegram_enabled ?? true,
    whatsapp: newSettings.whatsapp || '',
    whatsapp_enabled: newSettings.whatsapp_enabled ?? true,
    twitter: newSettings.twitter || '',
    twitter_enabled: newSettings.twitter_enabled ?? true,
    linkedin: newSettings.linkedin || '',
    linkedin_enabled: newSettings.linkedin_enabled ?? true,
    facebook: newSettings.facebook || '',
    facebook_enabled: newSettings.facebook_enabled ?? true,
    youtube: newSettings.youtube || '',
    youtube_enabled: newSettings.youtube_enabled ?? true,
    aparat: newSettings.aparat || '',
    aparat_enabled: newSettings.aparat_enabled ?? true,
    rubika: newSettings.rubika || '',
    rubika_enabled: newSettings.rubika_enabled ?? true,
    eitaa: newSettings.eitaa || '',
    eitaa_enabled: newSettings.eitaa_enabled ?? true,
    bale: newSettings.bale || '',
    bale_enabled: newSettings.bale_enabled ?? true,
    tiktok: newSettings.tiktok || '',
    tiktok_enabled: newSettings.tiktok_enabled ?? true,
    pinterest: newSettings.pinterest || '',
    pinterest_enabled: newSettings.pinterest_enabled ?? true,
    discord: newSettings.discord || '',
    discord_enabled: newSettings.discord_enabled ?? true,
    github: newSettings.github || '',
    github_enabled: newSettings.github_enabled ?? true,
    customLinks: newSettings.customLinks || []
  })
}, { deep: true })

const addCustomLink = () => {
  localSettings.customLinks.push({
    title: '',
    url: ''
  })
}

const removeCustomLink = (index) => {
  localSettings.customLinks.splice(index, 1)
}

const handleSave = () => {
  emit('save', { ...localSettings })
}

const handleReset = () => {
  emit('reset')
}
</script>

<style scoped>
.settings-section {
  background: white;
  border-radius: 8px;
  padding: 24px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.settings-header {
  margin-bottom: 24px;
  border-bottom: 2px solid #e5e7eb;
  padding-bottom: 16px;
}

.settings-header h2 {
  font-size: 1.5rem;
  font-weight: bold;
  color: #1f2937;
  margin-bottom: 8px;
}

.settings-description {
  color: #6b7280;
  font-size: 0.875rem;
}

.social-media-form {
  display: grid;
  gap: 20px;
}

.social-media-item {
  background: #f9fafb;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 16px;
  transition: all 0.2s;
}

.social-media-item:hover {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
}

.social-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.social-title {
  display: flex;
  align-items: center;
  gap: 12px;
}

.social-name {
  font-weight: 600;
  color: #374151;
  font-size: 1rem;
}

.toggle-switch {
  position: relative;
  display: inline-block;
  width: 48px;
  height: 24px;
  cursor: pointer;
}

.toggle-switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.toggle-slider {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #cbd5e1;
  border-radius: 24px;
  transition: 0.3s;
}

.toggle-slider:before {
  position: absolute;
  content: "";
  height: 18px;
  width: 18px;
  left: 3px;
  bottom: 3px;
  background-color: white;
  border-radius: 50%;
  transition: 0.3s;
}

.toggle-switch input:checked + .toggle-slider {
  background-color: #10b981;
}

.toggle-switch input:checked + .toggle-slider:before {
  transform: translateX(24px);
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-label {
  font-weight: 600;
  color: #374151;
  font-size: 0.875rem;
  display: flex;
  align-items: center;
  gap: 8px;
}

.form-input {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  font-size: 0.875rem;
  transition: all 0.2s;
  background: white;
}

.form-input:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.form-input:disabled {
  background-color: #f3f4f6;
  color: #9ca3af;
  cursor: not-allowed;
  opacity: 0.6;
}

.form-input::placeholder {
  color: #9ca3af;
}

.custom-links-section {
  border-top: 2px solid #e5e7eb;
  padding-top: 24px;
}

.custom-link-item {
  background: #f9fafb;
  border: 1px solid #e5e7eb;
  border-radius: 6px;
  padding: 16px;
  margin-bottom: 12px;
}

.btn-add {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 8px 16px;
  background: #10b981;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 0.875rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-add:hover {
  background: #059669;
}

.btn-remove {
  padding: 10px;
  background: #ef4444;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
}

.btn-remove:hover {
  background: #dc2626;
}

.settings-actions {
  display: flex;
  gap: 12px;
  justify-content: flex-start;
  margin-top: 24px;
  padding-top: 24px;
  border-top: 2px solid #e5e7eb;
}

.btn-primary {
  padding: 10px 24px;
  background: #3b82f6;
  color: white;
  border: none;
  border-radius: 6px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-primary:hover:not(:disabled) {
  background: #2563eb;
}

.btn-primary:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn-secondary {
  padding: 10px 24px;
  background: #6b7280;
  color: white;
  border: none;
  border-radius: 6px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-secondary:hover:not(:disabled) {
  background: #4b5563;
}

.btn-secondary:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

@media (max-width: 768px) {
  .settings-section {
    padding: 16px;
  }
  
  .settings-actions {
    flex-direction: column;
  }
  
  .btn-primary,
  .btn-secondary {
    width: 100%;
  }
}
</style>
