import { computed } from 'vue'
import { useAuthState } from './useAuthState'

export const useAdmin = () => {
  const { user, isAuthenticated } = useAuthState()

  // بررسی اینکه آیا کاربر admin است یا نه
  const isAdmin = computed(() => {
    if (!isAuthenticated.value || !user.value) return false

    const role = (user.value.role || '').toLowerCase()
    const username = (user.value.name || '').toLowerCase()
    const roleId = user.value.role_id || 0

    // بررسی نقش‌های ادمین (هماهنگ با backend)
    const adminRoles = ['admin', 'developer', 'super_admin', 'manager', 'operator']
    const isAdminRole = adminRoles.includes(role)

    // بررسی username خاص (dev)
    const isDevUser = username === 'dev'

    // بررسی role_id (معمولاً role_id = 1 برای مشتریان، 2+ برای ادمین‌ها، 100 برای developer)
    const isAdminRoleId = roleId > 1 || roleId === 100

    return isAdminRole || isDevUser || isAdminRoleId
  })

  return {
    isAdmin
  }
}