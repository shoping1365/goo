<template>
  <div class="space-y-4">
      <!-- Pages Section -->
      <ContentPanel
        panel-name="pages"
        title="برگه‌ها"
        :is-open="openPanel === 'pages'"
        @toggle-panel="$emit('toggle-panel', 'pages')"
        @add-selected="emitAddSelectedPages"
        :items="pages"
        v-model:search-query="localSearchPages"
        v-model:selected-items="localSelectedPages"
        :filtered-items="filteredPages"
        item-label="برگه"
        select-all-text="انتخاب همه"
        deselect-all-text="حذف انتخاب همه"
        add-button-text="افزودن برگه‌های انتخاب شده"
      />

      <!-- Posts Section -->
      <ContentPanel
        panel-name="posts"
        title="نوشته‌ها"
        :is-open="openPanel === 'posts'"
        @toggle-panel="$emit('toggle-panel', 'posts')"
        @add-selected="emitAddSelectedPosts"
        :items="posts"
        v-model:search-query="localSearchPosts"
        v-model:selected-items="localSelectedPosts"
        :filtered-items="filteredPosts"
        item-label="نوشته"
        select-all-text="انتخاب همه"
        deselect-all-text="حذف انتخاب همه"
        add-button-text="افزودن نوشته‌های انتخاب شده"
      />

      <!-- Categories Section -->
      <ContentPanel
        panel-name="categories"
        title="دسته‌بندی‌ها"
        :is-open="openPanel === 'categories'"
        @toggle-panel="$emit('toggle-panel', 'categories')"
        @add-selected="emitAddSelectedCategories"
        :items="categories"
        v-model:search-query="localSearchCategories"
        v-model:selected-items="localSelectedCategories"
        :filtered-items="filteredCategories"
        item-label="دسته‌بندی"
        select-all-text="انتخاب همه"
        deselect-all-text="حذف انتخاب همه"
        add-button-text="افزودن دسته‌بندی‌های انتخاب شده"
      />

      <!-- Product Categories Section -->
      <ContentPanel
        panel-name="productCategories"
        title="دسته‌های محصولات"
        :is-open="openPanel === 'productCategories'"
        @toggle-panel="$emit('toggle-panel', 'productCategories')"
        @add-selected="emitAddSelectedProductCategories"
        :items="productCategories"
        v-model:search-query="localSearchProductCategories"
        v-model:selected-items="localSelectedProductCategories"
        :filtered-items="filteredProductCategories"
        item-label="دسته محصول"
        select-all-text="انتخاب همه"
        deselect-all-text="حذف انتخاب همه"
        add-button-text="افزودن دسته‌های انتخاب شده"
      />
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import ContentPanel from './ContentPanel.vue'

const props = defineProps({
  openPanel: String,
  pages: Array,
  posts: Array,
  categories: Array,
  productCategories: Array,
})

const emit = defineEmits([
  'toggle-panel',
  'add-selected-pages',
  'add-selected-posts',
  'add-selected-categories',
  'add-selected-product-categories',
])

const localSearchPages = ref('')
const localSearchPosts = ref('')
const localSearchCategories = ref('')
const localSearchProductCategories = ref('')

const localSelectedPages = ref([])
const localSelectedPosts = ref([])
const localSelectedCategories = ref([])
const localSelectedProductCategories = ref([])

const normalizeString = (value) => (value || '').toString().toLowerCase()

const filteredPages = computed(() => {
  const query = normalizeString(localSearchPages.value)
  const list = Array.isArray(props.pages) ? props.pages : []
  if (!query) return list
  return list.filter(page =>
    normalizeString(page?.title).includes(query) ||
    normalizeString(page?.slug).includes(query)
  )
})

const filteredPosts = computed(() => {
  const query = normalizeString(localSearchPosts.value)
  const list = Array.isArray(props.posts) ? props.posts : []
  if (!query) return list
  return list.filter(post =>
    normalizeString(post?.title).includes(query) ||
    normalizeString(post?.slug).includes(query)
  )
})

const filteredCategories = computed(() => {
  const query = normalizeString(localSearchCategories.value)
  const list = Array.isArray(props.categories) ? props.categories : []
  if (!query) return list
  return list.filter(category =>
    normalizeString(category?.name).includes(query) ||
    normalizeString(category?.slug).includes(query)
  )
})

const filteredProductCategories = computed(() => {
  const query = normalizeString(localSearchProductCategories.value)
  const list = Array.isArray(props.productCategories) ? props.productCategories : []
  if (!query) return list
  return list.filter(category =>
    normalizeString(category?.name).includes(query) ||
    normalizeString(category?.slug).includes(query)
  )
})

const buildUniqueSelection = (source, selectedIds) => {
  const list = Array.isArray(source) ? source : []
  const selectedSet = new Set((Array.isArray(selectedIds) ? selectedIds : []).map((value) => {
    if (value === null || value === undefined) return value
    const numeric = Number(value)
    return Number.isFinite(numeric) ? numeric : value
  }))
  const seenKeys = new Set()
  const result = []

  list.forEach((item) => {
    const rawId = item?.id ?? item?.category_id ?? item?.ID ?? null
    const normalizedId = rawId === null || rawId === undefined
      ? rawId
      : (Number.isFinite(Number(rawId)) ? Number(rawId) : rawId)

    if (!selectedSet.has(normalizedId)) {
      return
    }
    const keySource = normalizedId ?? item?.slug ?? item?.path ?? item?.title ?? item?.name
    const key = keySource === undefined || keySource === null
      ? `__index_${result.length}`
      : `${typeof keySource}:${String(keySource)}`
    if (seenKeys.has(key)) {
      return
    }
    seenKeys.add(key)
    result.push(item)
  })

  return result
}

const emitAddSelectedPages = () => {
  const selectedItems = buildUniqueSelection(props.pages, localSelectedPages.value)
  emit('add-selected-pages', selectedItems)
  localSelectedPages.value = []
}

const emitAddSelectedPosts = () => {
  const selectedItems = buildUniqueSelection(props.posts, localSelectedPosts.value)
  emit('add-selected-posts', selectedItems)
  localSelectedPosts.value = []
}

const emitAddSelectedCategories = () => {
  const selectedItems = buildUniqueSelection(props.categories, localSelectedCategories.value)
  emit('add-selected-categories', selectedItems)
  localSelectedCategories.value = []
}

const emitAddSelectedProductCategories = () => {
  const selectedItems = buildUniqueSelection(props.productCategories, localSelectedProductCategories.value)
  emit('add-selected-product-categories', selectedItems)
  localSelectedProductCategories.value = []
}

</script>
