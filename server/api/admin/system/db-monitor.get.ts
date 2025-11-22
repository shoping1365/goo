export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  try {
    const data = await $fetch(`${base}/api/admin/monitoring/overview`, {
      method: 'GET',
      credentials: 'include',
      headers: { cookie: getHeader(event, 'cookie') || '' }
    })
    // انتظار داریم پاسخ شامل activeQueries و slowQueries باشد
    return {
      activeQueries: (data as { activeQueries?: unknown[]; slowQueries?: unknown[] })?.activeQueries || [],
      slowQueries: (data as { activeQueries?: unknown[]; slowQueries?: unknown[] })?.slowQueries || []
    }
  } catch (_err) {
    return { activeQueries: [], slowQueries: [] }
  }
})