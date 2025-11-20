// @ts-ignore
export default defineNuxtRouteMiddleware((_to, _from) => {
  // اجازه دسترسی بده - بک‌اند خودش permission و authentication چک می‌کند
  // اگر کاربر دسترسی نداشته باشد، API های بک‌اند 401 یا 403 برمی‌گردانند
})