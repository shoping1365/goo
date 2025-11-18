// Performance optimization: Component lazy loading and optimization
// این plugin موقتاً غیرفعال شده است چون باعث خطاهای Vue Router می‌شد
export default defineNitroPlugin((nitroApp) => {
  // Plugin disabled - was causing Vue Router warnings
  // The prefetch links were pointing to non-existent component files
})