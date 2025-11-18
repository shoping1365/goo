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
      activeQueries: (data as any)?.activeQueries || [],
      slowQueries: (data as any)?.slowQueries || []
    }
  } catch (err) {
    return { activeQueries: [], slowQueries: [] }
  }
})