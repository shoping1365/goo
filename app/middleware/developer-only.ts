// @ts-ignore
export default defineNuxtRouteMiddleware((to: any, from: any) => {
  // اجازه دسترسی بده - بک‌اند خودش permission و authentication چک می‌کند
  // اگر کاربر دسترسی نداشته باشد، API های بک‌اند 401 یا 403 برمی‌گردانند
})