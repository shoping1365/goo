<template>
  <div v-if="showItemsModal" class="modal-overlay" @click="closeModal">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <h3>انتخاب آیتم‌ها برای لایه</h3>
        <button class="close-button" @click="closeModal">
          <span>&times;</span>
        </button>
      </div>
      
          <div class="modal-body">
            <div class="items-grid">
              <div
                v-for="item in availableItems"
                :key="item.id"
                class="item-card"
                :class="{ selected: isItemSelected(item.id) }"
                role="button"
                tabindex="0"
                @click="handleItemClick(item)"
                @keydown.enter.prevent="handleItemClick(item)"
                @keydown.space.prevent="handleItemClick(item)"
              >
                <div class="item-card__icon">
                  <Icon
                    v-if="item.icon"
                    :name="item.icon"
                    class="item-card__icon-visual"
                  />
                  <span v-else class="item-card__icon-fallback">
                    {{ item.name?.charAt(0) || '?' }}
                  </span>
                </div>
                <div class="item-card__info">
                  <span class="item-card__name">{{ item.name }}</span>
                  <span class="item-card__path">{{ item.path }}</span>
                </div>
                <span v-if="isItemSelected(item.id)" class="item-card__check">✓</span>
              </div>
            </div>
        
        <div class="selected-items-summary">
          <h4>آیتم‌های انتخاب شده:</h4>
          <div v-if="newLayer.items.length === 0" class="no-items">
            هیچ آیتمی انتخاب نشده است
          </div>
          <div v-else class="selected-items-list">
            <div
              v-for="(itemObj, index) in newLayer.items"
              :key="index"
              class="selected-item"
            >
              <span class="item-name">{{ getItemDisplayName(typeof itemObj === 'string' ? itemObj : itemObj.id) }}</span>
              <span class="item-width">{{ (itemObj.width || (100 / newLayer.items.length)).toFixed(1) }}%</span>
            </div>
          </div>
        </div>
      </div>
      
      <!-- Menu chooser panel (shown when user clicks the menu card) -->
      <div v-if="showMenuChooser" class="menu-chooser">
        <h4>انتخاب منو برای نمایش</h4>
        <div v-if="menus.length === 0" class="no-menus">هیچ منویی موجود نیست</div>
        <div v-else class="menus-list">
          <label v-for="m in menus" :key="m.id" class="menu-item">
            <input v-model="selectedMenuIds" type="checkbox" :value="m.id" />
            <span class="menu-title">{{ m.title || m.name || ('menu_' + m.id) }}</span>
          </label>
        </div>

        <div class="menu-chooser-actions">
          <button class="btn btn-primary" @click="addSelectedMenus">افزودن منو(ها)</button>
          <button class="btn btn-secondary" @click="cancelMenuChooser">انصراف</button>
        </div>
      </div>

      <!-- Trust badges editor panel -->
      <div v-if="showTrustEditor" class="trust-editor">
        <h4>پیکربندی نشان‌های اعتماد</h4>

        <div class="trust-field">
          <label>عنوان بخش</label>
          <input v-model="trustTitle" type="text" class="trust-input" placeholder="نشان‌های اعتماد" />
        </div>

        <div class="trust-field">
          <label>توضیحات (اختیاری)</label>
          <textarea v-model="trustDescription" rows="2" class="trust-textarea" placeholder="کوتاهی درباره گواهی‌ها"></textarea>
        </div>

        <div class="trust-checkbox">
          <label>
            <input v-model="trustAutoPlay" type="checkbox" />
            <span>حرکت خودکار اسلایدر</span>
          </label>
        </div>

        <div v-if="trustAutoPlay" class="trust-field trust-field--inline">
          <label>فاصله زمانی بین اسلایدها (میلی‌ثانیه)</label>
          <input
            v-model.number="trustAutoPlayInterval"
            type="number"
            min="2000"
            step="500"
            class="trust-input"
            placeholder="مثلاً 5000"
          />
          <small class="trust-help">حداقل 2000 میلی‌ثانیه توصیه می‌شود.</small>
        </div>

        <div class="trust-options">
          <label class="trust-checkbox">
            <input v-model="trustShowArrows" type="checkbox" />
            <span>نمایش دکمه‌های قبلی/بعدی</span>
          </label>
          <label class="trust-checkbox">
            <input v-model="trustShowIndicators" type="checkbox" />
            <span>نمایش نشانگرهای پایین</span>
          </label>
        </div>

        <div class="trust-badges">
          <div
            v-for="(badge, index) in trustBadges"
            :key="badge.id"
            class="trust-badge-card"
          >
            <div class="trust-badge-header">
              <span>نشان {{ index + 1 }}</span>
              <button type="button" class="trust-remove" @click="removeTrustBadge(badge.id)">حذف</button>
            </div>
            <div class="trust-field">
              <label>عنوان</label>
              <input v-model="badge.title" type="text" class="trust-input" placeholder="مثلاً اینماد" />
            </div>
            <div class="trust-field">
              <label>لینک (اختیاری)</label>
              <input v-model="badge.link" type="text" dir="ltr" class="trust-input" placeholder="https://example.com" />
            </div>
            <div class="trust-field">
              <label>آدرس تصویر (اختیاری)</label>
              <input v-model="badge.imageUrl" type="text" dir="ltr" class="trust-input" placeholder="https://example.com/badge.png" />
            </div>
            <div class="trust-field">
              <label>کد/HTML جایگزین</label>
              <textarea v-model="badge.html" rows="3" dir="ltr" class="trust-textarea" placeholder="&lt;script&gt;...&lt;/script&gt;"></textarea>
            </div>
          </div>

          <button type="button" class="btn btn-outline trust-add" @click="addTrustBadge">+ افزودن نشان جدید</button>
        </div>

        <div class="trust-editor-actions">
          <button class="btn btn-primary" @click="confirmTrustBadges">ذخیره نشان‌ها</button>
          <button class="btn btn-secondary" @click="cancelTrustEditor">انصراف</button>
        </div>
      </div>
      <div class="modal-footer">
        <button class="btn btn-secondary" @click="closeModal">
          بستن
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { inject, ref, onMounted } from 'vue'

// Inject functions and data from parent (loose typing for brevity)
const showItemsModal: any = inject('showItemsModal')
const closeItemsModal: any = inject('closeItemsModal')
const toggleItem: any = inject('toggleItem')
const isItemSelected: any = inject('isItemSelected')
const newLayer: any = inject('newLayer')
const availableItems: any = inject('availableItems')

// Local state for menu chooser
const showMenuChooser = ref(false)
const menus = ref<any[]>([])
const selectedMenuIds = ref<string[]>([])
const pendingMenuItem = ref<any>(null)

type TrustBadgeForm = {
  id: string
  title: string
  link: string
  imageUrl: string
  html: string
}

const showTrustEditor = ref(false)
const pendingTrustItem = ref<any>(null)
const trustTitle = ref('نشان‌های اعتماد')
const trustDescription = ref('')
const trustAutoPlay = ref(true)
const trustAutoPlayInterval = ref(5000)
const trustShowArrows = ref(true)
const trustShowIndicators = ref(true)
const trustBadges = ref<TrustBadgeForm[]>([])

const resolveLayerState = () => (newLayer && typeof newLayer.value !== 'undefined' ? newLayer.value : newLayer)
const generateBadgeId = () => `trust-${Date.now().toString(36)}-${Math.random().toString(36).slice(2, 8)}`
const createEmptyBadge = (): TrustBadgeForm => ({
  id: generateBadgeId(),
  title: '',
  link: '',
  imageUrl: '',
  html: ''
})

const resetTrustEditor = () => {
  trustTitle.value = 'نشان‌های اعتماد'
  trustDescription.value = ''
  trustAutoPlay.value = true
  trustAutoPlayInterval.value = 5000
  trustShowArrows.value = true
  trustShowIndicators.value = true
  trustBadges.value = [createEmptyBadge(), createEmptyBadge()]
}

const cancelTrustEditor = () => {
  showTrustEditor.value = false
  pendingTrustItem.value = null
  trustTitle.value = 'نشان‌های اعتماد'
  trustDescription.value = ''
  trustAutoPlay.value = true
  trustAutoPlayInterval.value = 5000
  trustShowArrows.value = true
  trustShowIndicators.value = true
  trustBadges.value = []
}

const prepareTrustEditor = (item: any, existing?: any) => {
  pendingTrustItem.value = item
  resetTrustEditor()
  if (existing && typeof existing === 'object') {
    trustTitle.value = existing.props?.title || 'نشان‌های اعتماد'
    trustDescription.value = existing.props?.description || ''
    trustAutoPlay.value = existing.props?.autoPlay !== undefined ? Boolean(existing.props.autoPlay) : true
    trustAutoPlayInterval.value = existing.props?.autoPlayInterval ? Number(existing.props.autoPlayInterval) : 5000
    if (!Number.isFinite(trustAutoPlayInterval.value) || trustAutoPlayInterval.value <= 0) {
      trustAutoPlayInterval.value = 5000
    }
    trustShowArrows.value = existing.props?.showArrows !== undefined ? Boolean(existing.props.showArrows) : true
    trustShowIndicators.value = existing.props?.showIndicators !== undefined ? Boolean(existing.props.showIndicators) : true
    if (Array.isArray(existing.props?.badges) && existing.props.badges.length) {
      trustBadges.value = existing.props.badges.map((badge: any, index: number) => ({
        id: badge.id ? String(badge.id) : generateBadgeId() + index,
        title: badge.title || '',
        link: badge.link || '',
        imageUrl: badge.imageUrl || '',
        html: badge.html || ''
      }))
    }
  }
  showTrustEditor.value = true
}

const addTrustBadge = () => {
  trustBadges.value = [...trustBadges.value, createEmptyBadge()]
}

const removeTrustBadge = (badgeId: string) => {
  if (trustBadges.value.length <= 1) {
    trustBadges.value = [createEmptyBadge()]
    return
  }
  trustBadges.value = trustBadges.value.filter(badge => badge.id !== badgeId)
  if (trustBadges.value.length === 0) {
    trustBadges.value = [createEmptyBadge()]
  }
}

const confirmTrustBadges = () => {
  if (!pendingTrustItem.value) return
  const layerState = resolveLayerState()
  if (!layerState) return

  if (!Array.isArray(layerState.items)) {
    layerState.items = []
  }

  const cleanedBadges = trustBadges.value
    .map(badge => ({
      id: badge.id,
      title: badge.title.trim(),
      link: badge.link.trim(),
      imageUrl: badge.imageUrl.trim(),
      html: badge.html.trim()
    }))
    .filter(badge => badge.title || badge.link || badge.imageUrl || badge.html)

  const itemPayload = {
    id: pendingTrustItem.value.id,
    name: pendingTrustItem.value.name,
    path: pendingTrustItem.value.path,
    icon: pendingTrustItem.value.icon,
    align: 'center',
    bgColor: 'transparent',
    width: 0,
    props: {
      title: trustTitle.value.trim() || 'نشان‌های اعتماد',
      description: trustDescription.value.trim(),
      autoPlay: Boolean(trustAutoPlay.value),
      autoPlayInterval: Math.max(2000, Number(trustAutoPlayInterval.value) || 5000),
      showArrows: Boolean(trustShowArrows.value),
      showIndicators: Boolean(trustShowIndicators.value),
      badges: cleanedBadges
    }
  }

  const existingIndex = layerState.items.findIndex((it: any) => typeof it === 'object' && it.id === itemPayload.id)
  if (existingIndex !== -1) {
    layerState.items.splice(existingIndex, 1, itemPayload)
  } else {
    layerState.items.push(itemPayload)
  }

  const total = layerState.items.length
  if (total > 0) {
    layerState.items.forEach((it: any) => {
      it.width = Math.round((100 / total) * 100) / 100
    })
  }

  cancelTrustEditor()
}

async function fetchMenus() {
  try {
    // call admin menus endpoint
  const res: any = await $fetch('/api/admin/menus').catch(() => null)
    if (!res) return
    // handle different shapes
    menus.value = res.data || res.menus || res || []
  } catch (e) {
    console.error('خطا در بارگذاری منوها:', e)
    menus.value = []
  }
}

onMounted(() => {
  // prefetch menus so chooser is snappy
  fetchMenus()
})

function getItemDisplayName(itemId: string): string {
  try {
    const items = (availableItems && (availableItems.value || availableItems)) || []
    const item = items.find((ai: any) => ai.id === itemId)
    return item ? item.name : itemId
  } catch (e) {
    return itemId
  }
}

function handleItemClick(item: any) {
  if (!item) return
  
  // برای منو و trust، فقط آیتم را به لایه اضافه می‌کنیم
  // بعداً وقتی روی آیتم در لایه کلیک شد، تنظیمات باز می‌شود
  if (item.id === 'menu') {
    if (typeof toggleItem === 'function') {
      cancelTrustEditor()
      toggleItem(item.id)
    } else if (newLayer && newLayer.items) {
      const exists = newLayer.items.some((it: any) => 
        (typeof it === 'object' && it.id === 'menu') || it === 'menu'
      )
      if (!exists) {
        newLayer.items.push({ 
          id: 'menu', 
          name: 'منو', 
          path: '/menu', 
          icon: item.icon || 'heroicons:bars-3',
          align: 'center', 
          bgColor: 'transparent', 
          width: 0,
          config: { 
            menuIds: [] 
          }
        })
        // recalc widths
        const total = newLayer.items.length
        newLayer.items.forEach((it: any) => {
          if (typeof it === 'object') {
            it.width = Math.round((100 / total) * 100) / 100
          }
        })
      }
    }
    return
  }

  if (item.id === 'trust') {
    // به جای باز کردن editor، فقط آیتم را به لایه اضافه می‌کنیم
    // editor بعداً وقتی روی لایه کلیک شد باز خواهد شد
    if (typeof toggleItem === 'function') {
      cancelMenuChooser()
      toggleItem(item.id)
    } else if (newLayer && newLayer.items) {
      // اگر قبلاً اضافه نشده، آیتم trust با props پیش‌فرض اضافه می‌شود
      const exists = newLayer.items.some((it: any) => 
        (typeof it === 'object' && it.id === 'trust') || it === 'trust'
      )
      if (!exists) {
        newLayer.items.push({ 
          id: 'trust', 
          name: 'نشان‌های اعتماد', 
          path: '/trust', 
          icon: 'heroicons:shield-check',
          align: 'center', 
          bgColor: 'transparent', 
          width: 0,
          props: {
            title: 'نشان‌های اعتماد',
            trustBadges: []
          }
        })
        // recalc widths
        const total = newLayer.items.length
        newLayer.items.forEach((it: any) => {
          if (typeof it === 'object') {
            it.width = Math.round((100 / total) * 100) / 100
          }
        })
      }
    }
    return
  }

  // default toggle behaviour
  if (typeof toggleItem === 'function') {
    cancelTrustEditor()
    toggleItem(item.id)
  } else if (newLayer && newLayer.items) {
    // fallback: push minimal item
    newLayer.items.push({ id: item.id, name: item.name, path: item.path, icon: item.icon, align: 'right', bgColor: 'transparent', width: 0 })
  }
}

function cancelMenuChooser() {
  showMenuChooser.value = false
  pendingMenuItem.value = null
  selectedMenuIds.value = []
}

function addSelectedMenus() {
  if (!pendingMenuItem.value) return
  const cfg = {
    id: pendingMenuItem.value.id,
    name: pendingMenuItem.value.name,
    path: pendingMenuItem.value.path,
    icon: pendingMenuItem.value.icon,
    align: 'right',
    bgColor: 'transparent',
    width: 0,
    // store chosen menu ids under `config.menuIds`
    config: { menuIds: Array.isArray(selectedMenuIds.value) ? [...selectedMenuIds.value] : [] }
  }

  if (newLayer && newLayer.items) {
    newLayer.items.push(cfg)

    // recalc widths similarly to existing logic
    const total = newLayer.items.length
    newLayer.items.forEach((it: any) => {
      it.width = Math.round((100 / total) * 100) / 100
    })
  }

  // close chooser and keep modal open
  cancelMenuChooser()
}

const closeModal = () => {
  cancelMenuChooser()
  cancelTrustEditor()
  if (typeof closeItemsModal === 'function') {
    closeItemsModal()
  }
}
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  padding: 20px;
}

.modal-content {
  background: white;
  border-radius: 8px;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 800px;
  max-height: 90vh;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 24px;
  border-bottom: 1px solid #e5e7eb;
}

.modal-header h3 {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
}

.close-button {
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
  color: #6b7280;
  padding: 0;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 4px;
  transition: background-color 0.2s;
}

.close-button:hover {
  background: #f3f4f6;
}

.modal-body {
  flex: 1;
  padding: 24px;
  overflow-y: auto;
}

.trust-editor {
  padding: 24px;
  border-top: 1px solid #e5e7eb;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.trust-field {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.trust-field--inline {
  align-items: flex-start;
}

.trust-field label {
  font-size: 13px;
  font-weight: 500;
  color: #1f2937;
}

.trust-help {
  font-size: 11px;
  color: #6b7280;
}

.trust-input,
.trust-textarea {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  font-size: 13px;
  transition: border-color 0.2s ease, box-shadow 0.2s ease;
}

.trust-input:focus,
.trust-textarea:focus {
  outline: none;
  border-color: #2563eb;
  box-shadow: 0 0 0 2px rgba(37, 99, 235, 0.1);
}

.trust-textarea {
  font-family: 'Fira Code', 'Courier New', monospace;
  min-height: 90px;
}

.trust-badges {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.trust-options {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-top: 12px;
}

.trust-checkbox {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: #1f2937;
}

.trust-checkbox label {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
}

.trust-checkbox span {
  cursor: pointer;
}

.trust-checkbox input[type="checkbox"] {
  width: 16px;
  height: 16px;
  margin: 0;
  accent-color: #2563eb;
}

.trust-badge-card {
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 16px;
  background: #f9fafb;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.trust-badge-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-weight: 600;
  color: #1f2937;
}

.trust-remove {
  background: none;
  border: none;
  color: #ef4444;
  cursor: pointer;
  font-size: 13px;
  padding: 0;
}

.trust-remove:hover {
  text-decoration: underline;
}

.trust-add {
  justify-content: center;
}

.trust-editor-actions {
  display: flex;
  gap: 12px;
  justify-content: flex-end;
}

.items-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(170px, 1fr));
  gap: 12px;
  margin-bottom: 24px;
}

.item-card {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 10px;
  padding: 14px;
  border: 1px solid #d1d5db;
  border-radius: 10px;
  background: #ffffff;
  cursor: pointer;
  transition: transform 0.18s ease, box-shadow 0.18s ease, border-color 0.18s ease;
  min-height: 130px;
  outline: none;
}

.item-card:hover,
.item-card:focus-visible {
  border-color: #3b82f6;
  box-shadow: 0 10px 20px rgba(59, 130, 246, 0.12);
  transform: translateY(-2px);
}

.item-card.selected {
  border-color: #2563eb;
  background: linear-gradient(180deg, rgba(37, 99, 235, 0.08) 0%, rgba(37, 99, 235, 0.02) 100%);
}

.item-card__icon {
  width: 48px;
  height: 48px;
  border-radius: 14px;
  background: #f3f4f6;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.18s ease, color 0.18s ease;
}

.item-card:hover .item-card__icon,
.item-card:focus-visible .item-card__icon {
  background: #e0e7ff;
  color: #1d4ed8;
}

.item-card__icon-visual {
  width: 28px;
  height: 28px;
  color: inherit;
}

.item-card__icon-visual :deep(svg) {
  width: 28px;
  height: 28px;
}

.item-card__icon-fallback {
  font-size: 20px;
  font-weight: 600;
  color: #1d4ed8;
}

.item-card__info {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.item-card__name {
  font-weight: 600;
  font-size: 14px;
  color: #1f2937;
}

.item-card__path {
  font-size: 11px;
  letter-spacing: 0.4px;
  color: #6b7280;
  text-transform: lowercase;
  font-family: 'Fira Mono', 'Courier New', monospace;
}

.item-card__check {
  position: absolute;
  top: 10px;
  left: 10px;
  width: 22px;
  height: 22px;
  border-radius: 6px;
  background: #2563eb;
  color: #ffffff;
  font-size: 12px;
  font-weight: 700;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 4px 10px rgba(37, 99, 235, 0.3);
}

.selected-items-summary {
  border-top: 1px solid #e5e7eb;
  padding-top: 24px;
}

.selected-items-summary h4 {
  margin: 0 0 16px 0;
  font-size: 16px;
  font-weight: 600;
  color: #1f2937;
}

.no-items {
  color: #6b7280;
  font-style: italic;
  text-align: center;
  padding: 16px;
  background: #f9fafb;
  border-radius: 6px;
}

.selected-items-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.selected-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px;
  background: #f3f4f6;
  border-radius: 6px;
  border: 1px solid #e5e7eb;
}

.item-width {
  font-family: monospace;
  font-size: 12px;
  color: #6b7280;
  background: white;
  padding: 2px 6px;
  border-radius: 4px;
  border: 1px solid #d1d5db;
}

.modal-footer {
  padding: 24px;
  border-top: 1px solid #e5e7eb;
  display: flex;
  justify-content: flex-end;
}

.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 6px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-outline {
  background: #ffffff;
  color: #2563eb;
  border: 1px dashed #93c5fd;
}

.btn-outline:hover {
  background: #eff6ff;
  border-color: #3b82f6;
}

.btn-primary {
  background: #2563eb;
  color: #ffffff;
}

.btn-primary:hover {
  background: #1d4ed8;
}

.btn-secondary {
  background: #6b7280;
  color: white;
}

.btn-secondary:hover {
  background: #4b5563;
}

@media (max-width: 768px) {
  .modal-overlay {
    padding: 10px;
  }
  
  .modal-content {
    max-height: 95vh;
  }
  
  .items-grid {
    grid-template-columns: 1fr;
  }
  
  .modal-header,
  .modal-body,
  .modal-footer {
    padding: 16px;
  }
}
</style>


