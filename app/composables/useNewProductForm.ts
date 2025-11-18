import { reactive, ref } from 'vue'

// shared singleton state for new product wizard
const productForm = reactive({
  name: '',
  description: '',
  fullDescription: '',
  status: 'active'
})

const sections = reactive({
  // Info tab sections
  mainInfo: true,
  technicalInfo: true,
  displaySettings: true,
  scheduling: true,
  management: true,
  // Pricing tab sections
  mainPrices: true,
  priceSettings: true,
  taxSettings: true,
  tierPricing: true,
  discountCodes: true,
  // Shipping tab sections
  dimensions: true,
  shippingSettings: true,
  shippingCosts: true,
  deliverySchedule: true,
  specialSettings: true,
  packaging: true,
  // Inventory tab sections
  inventoryMain: true,
  purchaseLimits: true,
  alerts: true,
  advancedSettings: true,
  multiWarehouse: true,
  // SEO tab sections
  seoBasic: true,
  seoSchema: true,
  seoSocial: true,
  seoIndexing: true,
  seoAdvanced: true,
  seoPreview: true,
  // Variants tab sections
  variantAttributes: true,
  variantCombinations: true,
  variantGeneration: true,
  variantBulkEdit: true,
  // Others
  strengthsWeaknesses: true,
  faq: true,
  videoForm: true,
  videoList: true,
  aparatVideos: true
})

const toggleSection = (key: keyof typeof sections) => {
  sections[key] = !sections[key]
}

const tinyApiKey = 'qwa4j6x5mh2e3241igpyi345b4uhe2d5qeq6f8hy9qfkw2ro'

function generateAIContent(type: 'short' | 'full') {
  alert(`قابلیت تولید محتوا با AI برای ${type === 'short' ? 'توضیحات کوتاه' : 'توضیحات کامل'} در آینده اضافه خواهد شد.`)
}

export default function useNewProductForm() {
  return {
    productForm,
    sections,
    toggleSection,
    tinyApiKey,
    generateAIContent
  }
} 