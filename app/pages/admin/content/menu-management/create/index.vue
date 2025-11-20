<template>
  <div class="min-h-screen">
    <!-- Header -->
    <div class="header-bg">
      <div class="page-header-flex">
        <h1 class="page-title">ایجاد منو جدید</h1>
        <button class="btn btn-primary new-item-btn" @click="addNewMenu">بازگشت به لیست منوها</button>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-7xl px-4 sm:px-6 lg:px-8 py-6">
        <div class="flex flex-col lg:flex-row gap-6">
        <!-- Add Menu Items Panel (Left) -->
        <AddMenuItemsPanel
          :custom-link="customLink"
          :pages="pages"
          :posts="posts"
          :categories="categories"
          :product-categories="productCategories"
          @refresh="refreshContent"
          @add-custom-link="addCustomLink"
          @add-selected-pages="addSelectedPages"
          @add-selected-posts="addSelectedPosts"
          @add-selected-categories="addSelectedCategories"
          @add-selected-product-categories="addSelectedProductCategories"
        />

        <!-- Menu Structure Panel (Right) -->
        <MenuPreview
          :menu="currentMenu"
          :is-saving="isSaving"
          @update-slug="updateSlug"
          @update-item="updateMenuItem"
          @remove-item="handleRemoveMenuItem"
          @toggle-expanded="toggleItemExpanded"
          @drop-item="handleDropItem"
          @save="saveCurrentMenu"
        />
      </div>
    </div>

    <!-- Icon Selector Modal -->
    <div v-if="showIconSelector" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 max-w-2xl w-full mx-4 max-h-[90vh] overflow-y-auto">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">انتخاب آیکون</h3>
          <button class="text-gray-400 hover:text-gray-600" @click="showIconSelector = false">
            <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"/>
            </svg>
          </button>
        </div>
        <IconSelector v-model="selectedIconForItem" @close="closeIconSelector" />
      </div>
    </div>

    <!-- Media Library Modal -->
    <div v-if="showImageSelector" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 max-w-4xl w-full mx-4 max-h-[90vh] overflow-y-auto">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">انتخاب تصویر</h3>
          <button class="text-gray-400 hover:text-gray-600" @click="showImageSelector = false">
            <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"/>
            </svg>
          </button>
        </div>
        <div class="mb-4">
          <ImageSelector :images="filteredImages" :selected-images="selectedImageIds" @select="handleImageSelect" />
        </div>
        <div class="flex space-x-2 space-x-reverse">
          <button :disabled="selectedImageIds.length === 0" class="flex-1 bg-blue-600 text-white px-4 py-2 rounded-lg" @click="confirmImageSelection">
            انتخاب تصویر
          </button>
          <button class="px-4 py-2 border border-gray-300 rounded-lg" @click="showImageSelector = false">
            انصراف
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, onUnmounted } from "vue"
import { useMenuManagement } from "~/composables/useMenuManagement"
import IconSelector from "~/components/admin/ui/IconSelector.vue"
import ImageSelector from "~/components/media/ImageSelector.vue"
import AddMenuItemsPanel from "../components/AddMenuItemsPanel.vue"
import MenuPreview from "../components/MenuPreview.vue"
import { useRoute, navigateTo } from "#app"
import { useToast } from "~/composables/useToast"

definePageMeta({
  layout: "admin-main"
})

const { currentMenu, isSaving, pages, posts, categories, productCategories, saveMenu, createNewMenu, addMenuItem, handleDropItem, initializeContent, refreshContent, fetchMenu, generateSlug, registerDeletedMenuBranch } = useMenuManagement()

const showIconSelector = ref(false)
const selectedIconForItem = ref("")
const currentEditingItemIndex = ref(-1)
const showImageSelector = ref(false)
const selectedImageIds = ref([])
const images = ref([])

const customLink = reactive({ title: "", url: "" })
const { showSuccess, showError } = useToast()

const loadImages = async () => {
  try {
    const { data } = await $fetch("/api/admin/media")
    images.value = data || []
  } catch (error) {
    console.error("خطا در بارگذاری تصاویر:", error)
  }
}

const filteredImages = computed(() => images.value)

const handleImageSelect = (imageId) => {
  const index = selectedImageIds.value.indexOf(imageId)
  if (index > -1) selectedImageIds.value.splice(index, 1)
  else selectedImageIds.value = [imageId]
}

const confirmImageSelection = () => {
  if (selectedImageIds.value.length > 0 && currentEditingItemIndex.value >= 0) {
    const img = images.value.find(i => i.id === selectedImageIds.value[0])
    if (img) currentMenu.value.items[currentEditingItemIndex.value].icon = JSON.stringify({ type: "image", url: img.thumbnail })
  }
  selectedImageIds.value = []
  showImageSelector.value = false
}

const updateSlug = () => {
  if (currentMenu.value.name && (!currentMenu.value.slug || currentMenu.value.slug.startsWith("new-menu-"))) {
    currentMenu.value.slug = generateSlug(currentMenu.value.name)
  }
}

const toggleItemExpanded = (index) => {
  currentMenu.value.items[index].expanded = !currentMenu.value.items[index].expanded
}

const closeIconSelector = () => {
  if (currentEditingItemIndex.value >= 0) {
    currentMenu.value.items[currentEditingItemIndex.value].icon = selectedIconForItem.value
  }
  showIconSelector.value = false
}

const addSelectedPages = (selectedPages) => {
  console.log('🟢 addSelectedPages called with:', selectedPages)
  selectedPages?.forEach(page => addMenuItem({ title: page.title, path: `/page/${page.slug}`, itemType: "page" }))
}

const addSelectedPosts = (selectedPosts) => {
  console.log('🟢 addSelectedPosts called with:', selectedPosts)
  selectedPosts?.forEach(post => addMenuItem({ title: post.title, path: `/blog/${post.slug}`, itemType: "post" }))
}

const dedupeByKey = (list, makeKey) => {
  if (!Array.isArray(list)) return []
  const seen = new Set()
  const result = []
  list.forEach((item) => {
    const key = makeKey(item)
    if (key === undefined || key === null) return
    if (seen.has(key)) return
    seen.add(key)
    result.push(item)
  })
  return result
}

const addSelectedCategories = (selectedCategories) => {
  console.log('🟢 addSelectedCategories called with:', selectedCategories)
  const unique = dedupeByKey(selectedCategories, (cat) => cat?.id ?? cat?.slug)
  unique.forEach((cat) => {
    addMenuItem({
      title: cat?.name,
      path: `/category/${cat?.slug}`,
      itemType: "category"
    })
  })
}

const addSelectedProductCategories = (selectedProductCategories) => {
  console.log('🟢 addSelectedProductCategories called with:', selectedProductCategories)
  const unique = dedupeByKey(selectedProductCategories, (cat) => cat?.id ?? cat?.slug)
  unique.forEach((cat) => {
    addMenuItem({
      title: cat?.name,
      path: `/product-category/${cat?.slug}`,
      itemType: "product_category"
    })
  })
}

const addCustomLink = (linkData) => {
  console.log('🟢 addCustomLink called with:', linkData)
  if (linkData?.title && linkData?.url) {
    addMenuItem({ title: linkData.title, path: linkData.url, icon: "fas fa-link", enabled: true, itemType: "custom" })
  }
}

const handleRemoveMenuItem = (removedBranch) => {
  if (removedBranch) registerDeletedMenuBranch(removedBranch)
}

const updateMenuItem = (index, item) => {
  currentMenu.value.items[index] = item
}

const saveCurrentMenu = async () => {
  if (!currentMenu.value.name) return
  try {
    const response = await saveMenu(currentMenu.value)
    const successMessage = response?.message ?? "منو با موفقیت ذخیره شد"
    showSuccess(successMessage, 4000)
  } catch (error) {
    console.error("Error saving menu:", error)
    const errorMessage = error?.data?.message ?? error?.message ?? "خطا در ذخیره منو"
    showError(errorMessage, 5000)
  }
}

const addNewMenu = () => navigateTo("/admin/content/menu-management")

let autoRefreshInterval = null

onUnmounted(() => {
  if (autoRefreshInterval !== null) {
    clearInterval(autoRefreshInterval)
    autoRefreshInterval = null
  }
})

onMounted(async () => {
  const route = useRoute()
  const menuId = route.query.id
  if (menuId) {
    try {
      await fetchMenu(parseInt(menuId))
    } catch (error) {
      await navigateTo("/admin/content/menu-management")
      return
    }
  } else {
    createNewMenu()
  }
  await initializeContent()
  autoRefreshInterval = setInterval(async () => {
    await refreshContent()
  }, 30000)
})
</script>

<style scoped>
.header-bg {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 30px;
  border-radius: 12px;
  margin-bottom: 30px;
}

.page-header-flex {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.page-title {
  font-size: 2rem;
  font-weight: bold;
  margin: 0;
}

.new-item-btn {
  background: rgba(255, 255, 255, 0.2);
  border: 2px solid rgba(255, 255, 255, 0.3);
  color: white;
  padding: 12px 24px;
  border-radius: 8px;
  font-weight: 600;
  transition: all 0.3s ease;
}

.new-item-btn:hover {
  background: rgba(255, 255, 255, 0.3);
  border-color: rgba(255, 255, 255, 0.5);
}
</style>
