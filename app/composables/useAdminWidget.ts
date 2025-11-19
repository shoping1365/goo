import { readonly } from 'vue'
import { useWidgetCore } from './_useWidgetCore'

/**
 * Composable برای مدیریت ابزارک‌ها در پنل ادمین
 * شامل تمام عملیات CRUD و مدیریت state با احراز هویت
 */
export const useAdminWidget = () => {
  const core = useWidgetCore('/api/admin/widgets')
  
  return {
    ...core,
    // Override state with readonly where appropriate to match original API
    widgets: readonly(core.widgets),
    loading: readonly(core.loading),
    error: readonly(core.error)
  }
}
