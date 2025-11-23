<template>
  <div v-if="hasAccess" class="meta-box">
    <button
      class="meta-box__header"
      @click="$emit('toggle-panel')"
    >
      <span class="meta-box__title">{{ title }}</span>
      <span v-if="localSelectedItems.length > 0" class="meta-box__summary">{{ localSelectedItems.length }} انتخاب شده</span>
      <svg
        class="meta-box__chevron"
        :class="{ 'meta-box__chevron--open': isOpen }"
        viewBox="0 0 20 20"
      >
        <path d="M5.23 7.21a.75.75 0 011.06.02L10 10.94l3.71-3.71a.75.75 0 111.06 1.06l-4.24 4.24a.75.75 0 01-1.06 0L5.21 8.27a.75.75 0 01.02-1.06z" fill="currentColor" />
      </svg>
    </button>

    <div v-if="isOpen" class="meta-box__body">
      <div class="meta-box__tabs">
        <button
          v-for="tab in tabs"
          :key="tab.key"
          class="meta-box__tab"
          :class="{ 'meta-box__tab--active': activeTab === tab.key }"
          type="button"
          @click="setActiveTab(tab.key)"
        >
          {{ tab.label }}
        </button>
      </div>

      <div v-if="activeTab === 'search'" class="meta-box__search">
        <label class="meta-box__search-label" :for="`search-input-${panelName}`">جستجو</label>
        <div class="meta-box__search-field">
          <input
            :id="`search-input-${panelName}`"
            v-model="localSearchQuery"
            type="text"
            class="meta-box__search-input"
          />
        </div>
      </div>

      <div class="meta-box__list" :class="{ 'meta-box__list--empty': !displayedItems.length }">
        <div class="meta-box__list-scroll">
          <div
            v-for="item in displayedItems"
            :key="item.id"
            class="meta-box__item"
          >
            <label class="meta-box__item-label">
              <input
                :id="`${panelName}-${item.id}`"
                :value="item.id"
                type="checkbox"
                class="meta-box__checkbox"
                :checked="isItemSelected(item.id)"
                @change="toggleItemSelection(item.id)"
              />
              <span
                class="meta-box__item-title"
                :style="{ paddingInlineStart: `${Math.max(0, Number(item?.depth) || 0) * 16}px` }"
              >
                {{ item.displayName || item.title || item.name }}
              </span>
            </label>
          </div>
          <p v-if="!displayedItems.length" class="meta-box__empty">موردی یافت نشد.</p>
        </div>
      </div>

      <div class="meta-box__footer">
        <label class="meta-box__select-all">
          <input
            type="checkbox"
            class="meta-box__checkbox"
            :checked="allDisplayedSelected"
            @change="toggleSelectAll"
          />
          <span>انتخاب همه</span>
        </label>
        <button
          type="button"
          class="meta-box__submit"
          :disabled="localSelectedItems.length === 0"
          @click="handleAddSelected"
        >
          {{ addButtonText }}
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

interface ContentItem {
  id?: string | number
  name?: string
  title?: string
  displayName?: string
  depth?: number
  slug?: string
  [key: string]: unknown
}

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

const props = defineProps<{
  panelName: string
  title: string
  isOpen: boolean
  items?: ContentItem[]
  searchQuery?: string
  selectedItems?: (string | number)[]
  filteredItems?: ContentItem[]
  itemLabel?: string
  selectAllText?: string
  deselectAllText?: string
  addButtonText?: string
}>()

const emit = defineEmits([
  'toggle-panel',
  'add-selected',
  'update:search-query',
  'update:selected-items',
])

const tabs = [
  { key: 'mostUsed', label: 'پراستفاده' },
  { key: 'all', label: 'دیدن همه' },
  { key: 'search', label: 'جستجو' },
]

const activeTab = ref('mostUsed')
const localSelectedItems = ref([...(props.selectedItems || [])])
const localSearchQuery = ref(props.searchQuery || '')

watch(
  () => props.selectedItems,
  (value = []) => {
    localSelectedItems.value = [...value]
  },
  { deep: true }
)

watch(
  () => props.searchQuery,
  (value = '') => {
    if (value !== localSearchQuery.value) {
      localSearchQuery.value = value
    }
  }
)

watch(localSearchQuery, (value) => {
  emit('update:search-query', value)
})

const baseItems = computed<ContentItem[]>(() => (Array.isArray(props.items) ? props.items : []))

const mostUsedItems = computed<ContentItem[]>(() => {
  const items = baseItems.value.slice(0, 10)
  const selectedExtras = baseItems.value.filter(
    (item: ContentItem) =>
      localSelectedItems.value.includes(item?.id) &&
      !items.some((baseItem: ContentItem) => baseItem?.id === item?.id)
  )
  return [...items, ...selectedExtras]
})

const displayedItems = computed<ContentItem[]>(() => {
  if (activeTab.value === 'search') {
    if (!localSearchQuery.value) {
      return []
    }
    return Array.isArray(props.filteredItems) ? props.filteredItems : []
  }
  if (activeTab.value === 'all') {
    return baseItems.value
  }
  return mostUsedItems.value
})

const allDisplayedSelected = computed(() => {
  const ids = displayedItems.value.map((item: ContentItem) => item?.id)
  if (!ids.length) return false
  return ids.every((id) => localSelectedItems.value.includes(id))
})

const isItemSelected = (itemId: string | number): boolean => {
  return localSelectedItems.value.includes(itemId)
}

const toggleItemSelection = (itemId: string | number): void => {
  const index = localSelectedItems.value.indexOf(itemId)
  if (index > -1) {
    localSelectedItems.value.splice(index, 1)
  } else {
    localSelectedItems.value.push(itemId)
  }
  emit('update:selected-items', [...localSelectedItems.value])
}

const toggleSelectAll = (): void => {
  const ids = displayedItems.value.map((item: ContentItem) => item?.id).filter(Boolean)
  if (!ids.length) return

  if (allDisplayedSelected.value) {
    localSelectedItems.value = localSelectedItems.value.filter((id) => !ids.includes(id))
  } else {
    const merged = new Set([...localSelectedItems.value, ...ids])
    localSelectedItems.value = Array.from(merged)
  }

  emit('update:selected-items', [...localSelectedItems.value])
}

const setActiveTab = (tabKey) => {
  activeTab.value = tabKey
  if (tabKey !== 'search' && localSearchQuery.value) {
    localSearchQuery.value = ''
  }
}

const handleAddSelected = () => {
  emit('add-selected')
  // Clear selections after emitting
  localSelectedItems.value = []
  emit('update:selected-items', [])
}
</script>

<style scoped>
.meta-box {
  border: 1px solid #d7d7d7;
  border-radius: 8px;
  background: #fff;
}

.meta-box__header {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0.9rem 1rem;
  background: #f8f9fb;
  border: none;
  border-radius: 8px 8px 0 0;
  font-weight: 600;
  color: #1f2937;
  text-align: right;
}

.meta-box__header:hover {
  background: #eef1f6;
}

.meta-box__title {
  flex: 1;
}

.meta-box__summary {
  font-size: 0.85rem;
  color: #6b7280;
  margin-left: 0.75rem;
}

.meta-box__chevron {
  width: 18px;
  height: 18px;
  transition: transform 0.2s ease;
  color: #6b7280;
}

.meta-box__chevron--open {
  transform: rotate(180deg);
}

.meta-box__body {
  border-top: 1px solid #e5e7eb;
  padding: 1rem;
}

.meta-box__tabs {
  display: flex;
  gap: 0.75rem;
  margin-bottom: 0.85rem;
}

.meta-box__tab {
  border: none;
  background: transparent;
  color: #2563eb;
  font-size: 0.85rem;
  padding: 0;
}

.meta-box__tab--active {
  color: #111827;
  font-weight: 600;
}

.meta-box__search {
  margin-bottom: 1rem;
}

.meta-box__search-label {
  display: block;
  font-size: 0.8rem;
  color: #374151;
  margin-bottom: 0.35rem;
}

.meta-box__search-field {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.meta-box__search-input {
  flex: 1;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  padding: 0.45rem 0.6rem;
  font-size: 0.85rem;
}

.meta-box__list {
  border: 1px solid #e5e7eb;
  border-radius: 6px;
  background: #fff;
}

.meta-box__list-scroll {
  max-height: 235px;
  overflow-y: auto;
}

.meta-box__item {
  border-bottom: 1px solid #f1f1f1;
  cursor: pointer;
  transition: background-color 0.15s ease;
}

.meta-box__item:hover {
  background-color: #f9fafb;
}

.meta-box__item:last-child {
  border-bottom: none;
}

.meta-box__item-label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.9rem;
  color: #374151;
  padding: 0.55rem 0.75rem;
  cursor: pointer;
  user-select: none;
}

.meta-box__checkbox {
  width: 16px;
  height: 16px;
  cursor: pointer;
  pointer-events: auto;
  flex-shrink: 0;
}

.meta-box__item-title {
  flex: 1;
  cursor: pointer;
  user-select: none;
  display: block;
}

.meta-box__empty {
  padding: 1rem;
  text-align: center;
  color: #9ca3af;
  font-size: 0.85rem;
}

.meta-box__footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 1rem;
  gap: 0.75rem;
}

.meta-box__select-all {
  display: flex;
  align-items: center;
  gap: 0.4rem;
  color: #374151;
  font-size: 0.85rem;
}

.meta-box__submit {
  border: 1px solid #dc2626;
  background: #fee2e2;
  color: #b91c1c;
  border-radius: 6px;
  padding: 0.5rem 1rem;
  font-size: 0.85rem;
  font-weight: 600;
  transition: background 0.2s ease;
}

.meta-box__submit:hover:not(:disabled) {
  background: #fecaca;
}

.meta-box__submit:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}
</style>
