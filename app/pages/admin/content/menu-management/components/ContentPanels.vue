<template>
  <div v-if="hasAccess" class="space-y-4">
      <!-- Pages Section -->
      <ContentPanel
        v-model:search-query="localSearchPages"
        v-model:selected-items="localSelectedPages"
        panel-name="pages"
        title="برگه‌ها"
        :is-open="openPanel === 'pages'"
        :items="pages"
        :filtered-items="filteredPages"
        item-label="برگه"
        select-all-text="انتخاب همه"
        deselect-all-text="حذف انتخاب همه"
        add-button-text="افزودن برگه‌های انتخاب شده"
        @toggle-panel="$emit('toggle-panel', 'pages')"
        @add-selected="emitAddSelectedPages"
      />

      <!-- Posts Section -->
      <ContentPanel
        v-model:search-query="localSearchPosts"
        v-model:selected-items="localSelectedPosts"
        panel-name="posts"
        title="نوشته‌ها"
        :is-open="openPanel === 'posts'"
        :items="posts"
        :filtered-items="filteredPosts"
        item-label="نوشته"
        select-all-text="انتخاب همه"
        deselect-all-text="حذف انتخاب همه"
        add-button-text="افزودن نوشته‌های انتخاب شده"
        @toggle-panel="$emit('toggle-panel', 'posts')"
        @add-selected="emitAddSelectedPosts"
      />

      <!-- Categories Section -->
      <ContentPanel
        v-model:search-query="localSearchCategories"
        v-model:selected-items="localSelectedCategories"
        panel-name="categories"
        title="دسته‌بندی‌ها"
        :is-open="openPanel === 'categories'"
        :items="categories"
        :filtered-items="filteredCategories"
        item-label="دسته‌بندی"
        select-all-text="انتخاب همه"
        deselect-all-text="حذف انتخاب همه"
        add-button-text="افزودن دسته‌بندی‌های انتخاب شده"
        @toggle-panel="$emit('toggle-panel', 'categories')"
        @add-selected="emitAddSelectedCategories"
      />

      <!-- Product Categories Section -->
      <ContentPanel
        v-model:search-query="localSearchProductCategories"
        v-model:selected-items="localSelectedProductCategories"
        panel-name="productCategories"
        title="دسته‌های محصولات"
        :is-open="openPanel === 'productCategories'"
        :items="productCategories"
        :filtered-items="filteredProductCategories"
        item-label="دسته محصول"
        select-all-text="انتخاب همه"
        deselect-all-text="حذف انتخاب همه"
        add-button-text="افزودن دسته‌های انتخاب شده"
        @toggle-panel="$emit('toggle-panel', 'productCategories')"
        @add-selected="emitAddSelectedProductCategories"
      />
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';
import ContentPanel from './ContentPanel.vue'

interface PageItem {
  id?: string | number
  title?: string
  slug?: string
  [key: string]: unknown
}

interface PostItem {
  id?: string | number
  title?: string
  slug?: string
  [key: string]: unknown
}

interface CategoryItem {
  id?: string | number
  name?: string
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
  openPanel?: string
  pages?: PageItem[]
  posts?: PostItem[]
  categories?: CategoryItem[]
  productCategories?: CategoryItem[]
}>()

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

const normalizeString = (value: unknown): string => (value || '').toString().toLowerCase()

const filteredPages = computed<PageItem[]>(() => {
  const query = normalizeString(localSearchPages.value)
  const list = Array.isArray(props.pages) ? props.pages : []
  if (!query) return list
  return list.filter((page: PageItem) =>
    normalizeString(page?.title).includes(query) ||
    normalizeString(page?.slug).includes(query)
  )
})

const filteredPosts = computed<PostItem[]>(() => {
  const query = normalizeString(localSearchPosts.value)
  const list = Array.isArray(props.posts) ? props.posts : []
  if (!query) return list
  return list.filter((post: PostItem) =>
    normalizeString(post?.title).includes(query) ||
    normalizeString(post?.slug).includes(query)
  )
})

const filteredCategories = computed<CategoryItem[]>(() => {
  const query = normalizeString(localSearchCategories.value)
  const list = Array.isArray(props.categories) ? props.categories : []
  if (!query) return list
  return list.filter((category: CategoryItem) =>
    normalizeString(category?.name).includes(query) ||
    normalizeString(category?.slug).includes(query)
  )
})

const filteredProductCategories = computed<CategoryItem[]>(() => {
  const query = normalizeString(localSearchProductCategories.value)
  const list = Array.isArray(props.productCategories) ? props.productCategories : []
  if (!query) return list
  return list.filter((category: CategoryItem) =>
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
