import { computed, readonly, ref } from 'vue'

// تعریف interface برای Schema
interface Schema {
  id: string
  name: string
  type: string
  description: string
  isActive: boolean
  template: Record<string, unknown>
  meta: Record<string, unknown>
  openGraph: Record<string, unknown>
}

// تعریف interface برای index data
interface IndexData {
  templates: Array<{
    id: string
    name: string
    type: string
    file: string
    isActive: boolean
  }>
}

// تعریف interface برای template data
interface TemplateData {
  name: string
  type: string
  description: string
  isActive: boolean
  template: Record<string, unknown>
  meta: Record<string, unknown>
  openGraph: Record<string, unknown>
}

export const useSchema = () => {
  const schemas = ref<Schema[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  // دریافت تمام تمپلیت‌ها
  const fetchAllTemplates = async () => {
    loading.value = true
    error.value = null

    try {
      // دریافت لیست تمپلیت‌ها از فایل index.json
      const indexData = await $fetch<IndexData>('/templates/schema/index.json')
      const templates: Schema[] = []

      // دریافت جزئیات هر تمپلیت
      for (const templateInfo of indexData.templates) {
        if (templateInfo.isActive) {
          try {
            const templateData = await $fetch<TemplateData>(`/templates/schema/${templateInfo.file}`)
            templates.push({
              id: templateInfo.id,
              name: templateData.name,
              type: templateData.type,
              description: templateData.description,
              isActive: templateData.isActive,
              template: templateData.template,
              meta: templateData.meta,
              openGraph: templateData.openGraph
            })
          } catch (err) {
            console.error(`خطا در دریافت تمپلیت ${templateInfo.file}:`, err)
          }
        }
      }

      schemas.value = templates
    } catch (err) {
      error.value = 'خطا در دریافت تمپلیت‌ها'
      console.error('خطا در دریافت تمپلیت‌ها:', err)
    } finally {
      loading.value = false
    }
  }

  // دریافت تمپلیت‌ها بر اساس نوع
  const fetchTemplatesByType = async (type: string) => {
    await fetchAllTemplates()
    schemas.value = schemas.value.filter(s => s.type === type)
  }

  // دریافت تمپلیت بر اساس ID
  const fetchTemplateById = async (id: string): Promise<Schema | null> => {
    try {
      const templateData = await $fetch<TemplateData>(`/templates/schema/${id}.json`)
      return {
        id,
        name: templateData.name,
        type: templateData.type,
        description: templateData.description,
        isActive: templateData.isActive,
        template: templateData.template,
        meta: templateData.meta,
        openGraph: templateData.openGraph
      }
    } catch (err) {
      console.error('خطا در دریافت تمپلیت:', err)
      return null
    }
  }

  // تولید اسکیما از تمپلیت
  const generateSchemaFromTemplate = async (templateId: string, data: Record<string, unknown>): Promise<Record<string, unknown> | null> => {
    try {
      const template = await fetchTemplateById(templateId)
      if (!template) return null

      // جایگزینی متغیرها در تمپلیت
      const generatedSchema = JSON.parse(JSON.stringify(template.template))
      replaceTemplateVariables(generatedSchema, data)

      return generatedSchema
    } catch (err) {
      console.error('خطا در تولید اسکیما:', err)
      return null
    }
  }

  // جایگزینی متغیرها در تمپلیت
  const replaceTemplateVariables = (obj: Record<string, unknown>, data: Record<string, unknown>) => {
    for (const key in obj) {
      if (typeof obj[key] === 'string') {
        obj[key] = (obj[key] as string).replace(/\{\{(\w+)\}\}/g, (match: string, variable: string) => {
          return (data[variable] as string) || match
        })
      } else if (typeof obj[key] === 'object' && obj[key] !== null) {
        replaceTemplateVariables(obj[key] as Record<string, unknown>, data)
      }
    }
  }

  // دریافت انواع اسکیما
  const fetchSchemaTypes = async (): Promise<string[]> => {
    try {
      const indexData = await $fetch<IndexData>('/templates/schema/index.json')
      const types = indexData.templates.map((t) => t.type)
      return [...new Set(types)] as string[]
    } catch (err) {
      console.error('خطا در دریافت انواع اسکیما:', err)
      return []
    }
  }

  // Computed properties
  const activeTemplates = computed(() =>
    schemas.value.filter(s => s.isActive)
  )

  const templatesByType = computed(() => (type: string) =>
    schemas.value.filter(s => s.type === type && s.isActive)
  )

  const articleTemplates = computed(() =>
    templatesByType.value('Article')
  )

  const productTemplates = computed(() =>
    templatesByType.value('Product')
  )

  const organizationTemplates = computed(() =>
    templatesByType.value('Organization')
  )

  return {
    // State
    schemas: readonly(schemas),
    loading: readonly(loading),
    error: readonly(error),

    // Methods
    fetchAllTemplates,
    fetchTemplatesByType,
    fetchTemplateById,
    generateSchemaFromTemplate,
    fetchSchemaTypes,

    // Computed
    activeTemplates,
    templatesByType,
    articleTemplates,
    productTemplates,
    organizationTemplates,
  }
} 