// Composable to manage attribute groups and their attribute values
import { ref } from 'vue'

export interface AttributeValue {
  id: number
  value: string
}
export interface Attribute {
  id: number
  name: string
  values?: AttributeValue[]
  type?: string
  typeDisplay?: string
  controlType?: string
  controlTypeDisplay?: string
  isRequired?: boolean
  sort_order?: number
  has_filter?: boolean
  is_key?: boolean
  show_on_product?: boolean
  selectedOptionId?: number | string
  selectedOptionIds?: Array<number>
  valueText?: string
}

// Helper functions to normalise Persian labels that come from the API into
// consistent English slugs that the UI logic relies on.
function normaliseType(raw: string | undefined): string | undefined {
  if (!raw) return raw
  switch (raw.trim()) {
    case 'انتخاب':
      return 'select'
    case 'متن سفارشی':
      return 'custom_text'
    default:
      return raw
  }
}

function normaliseControlType(raw: string | undefined): string | undefined {
  if (!raw) return raw
  const val = raw.trim()
  switch (val) {
    case 'لیست باز شونده تک انتخابی':
      return 'dropdown_single'
    case 'لیست باز شونده چند انتخابی':
      return 'dropdown_multi'
    case 'تکست باکس متنی تک خطی':
      return 'textbox_single'
    case 'تکست باکس متنی چند خطی':
      return 'textbox_multiline'
    default:
      return val
  }
}

function faType(slug: string | undefined) {
  if (!slug) return ''
  return slug === 'select' ? 'انتخاب' : (slug === 'custom_text' ? 'متن سفارشی' : slug)
}

function faControl(slug: string | undefined) {
  if (!slug) return ''
  switch (slug) {
    case 'dropdown_single': return 'لیست باز شونده تک انتخابی'
    case 'dropdown_multi': return 'لیست باز شونده چند انتخابی'
    case 'textbox_multiline': return 'متن چند خطی'
    default: return slug
  }
}

export function useAttributeGroup() {
  const attributes = ref<Attribute[]>([])
  const loadingGroup = ref(false)

  /**
   * Load attributes of a given group ID.
   */
  async function loadGroup(groupId: string | number) {
    if (!groupId) return
    loadingGroup.value = true
    try {
      interface AttributeGroupResponse {
        attributes?: unknown[]
        Attributes?: unknown[]
        attribute_group_attributes?: unknown[]
        data?: unknown[]
        [key: string]: unknown
      }
      interface AttributeItem {
        attribute?: { id?: number | string; name?: string }
        Attribute?: { id?: number | string; name?: string }
        attribute_id?: number | string
        id?: number | string
        name?: string
        [key: string]: unknown
      }
      const res = await $fetch<AttributeGroupResponse>(`/api/attribute-groups/${groupId}`)
      const list = res?.attributes || res?.Attributes || res?.attribute_group_attributes || res?.data || []
      if (Array.isArray(list)) {
        attributes.value = (list as AttributeItem[]).map((a) => {
          const attrId = a.attribute?.id || a.Attribute?.id || a.attribute_id || a.id
          const idValue = typeof attrId === 'string' ? Number(attrId) : (attrId as number) || 0

          const typeValue = String(a.type || a.Type || a.attribute_type || a.AttributeType || '')
          const controlTypeValue = String(a.control_type || a.ControlType || a.controlType || '')

          return {
            id: idValue,
            name: String(a.attribute?.name || a.Attribute?.name || a.name || ''),
            // Normalise Persian labels to English slugs so that downstream components
            // can make simple equality checks.
            type: normaliseType(typeValue),
            controlType: normaliseControlType(controlTypeValue),
            typeDisplay: faType(normaliseType(typeValue)),
            controlTypeDisplay: faControl(normaliseControlType(controlTypeValue)),
            isRequired: !!(a.is_required ?? a.IsRequired),
            sort_order: Number(a.sort_order || a.SortOrder || 0),
            has_filter: !!(a.has_filter ?? a.HasFilter),
            is_key: !!(a.is_key ?? a.IsKey),
            show_on_product: !!(a.show_on_product ?? a.ShowOnProduct),
            values: [],
            selectedOptionId: a.selected_option_id || a.SelectedOptionId,
            selectedOptionIds: a.selected_option_ids || a.SelectedOptionIds || [],
            valueText: String(a.value_text || a.ValueText || '')
          } as Attribute
        })

        // Fetch possible values for each attribute in parallel
        await Promise.all(
          attributes.value.map(async (attr) => {
            const vals = await loadAttributeValues(attr.id)
            // Normalise values structure to ensure {id, value} shape
            if (Array.isArray(vals)) {
              attr.values = vals.map((v: Record<string, unknown>) => {
                const valueId = v.id ?? v.ID ?? v.value_id ?? v.ValueID ?? v.value_id
                const idNum = typeof valueId === 'string' ? Number(valueId) : (valueId as number) || 0
                return {
                  id: idNum,
                  value: String(v.value ?? v.Value ?? v.name ?? v.Name ?? v.label ?? '')
                }
              })
            } else {
              attr.values = []
            }
          })
        )
      } else {
        attributes.value = []
      }
    } catch (e) {
      console.error('Failed to load attribute group', e)
      attributes.value = []
    } finally {
      loadingGroup.value = false
    }
  }

  /**
   * Load possible values for a specific attribute.
   */
  async function loadAttributeValues(attrId: string | number) {
    if (!attrId) return []
    try {
      interface AttributeValuesResponse {
        data?: Array<Record<string, unknown>>
        [key: string]: unknown
      }
      const res = await $fetch<unknown>(`/api/attribute-values/by-attribute/${attrId}`)
      if (Array.isArray(res)) {
        return res as Array<Record<string, unknown>>
      }
      const response = res as AttributeValuesResponse
      if (response && typeof response === 'object' && 'data' in response && Array.isArray(response.data)) {
        return response.data
      }
      return []
    } catch (e) {
      console.error('Failed load values', e)
      return []
    }
  }

  return { attributes, loadingGroup, loadGroup, loadAttributeValues }
} 