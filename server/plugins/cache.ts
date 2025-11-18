// Performance cache plugin
export default defineNitroPlugin((nitroApp) => {
  // Configure cache headers for better performance
  nitroApp.hooks.hook('render:response', (response, { event }) => {
    // Ensure response.headers exists
    if (!response.headers) {
      return
    }

    // Check if headers is a Headers object or a plain object
    const headers = response.headers as unknown as Headers | Record<string, string> | undefined
    if (!headers) {
      return
    }

    // Handle Headers object
    if (headers instanceof Headers) {
      // Add cache headers for static assets
      if (event.node.req.url?.includes('/_nuxt/')) {
        headers.set('Cache-Control', 'public,max-age=31536000,immutable')
      }

      // Add cache headers for API responses
      if (event.node.req.url?.startsWith('/api/')) {
        headers.set('Cache-Control', 'public,max-age=60,stale-while-revalidate=300')
      }
    } else if (headers !== null) {
      // Handle plain object
      // Add cache headers for static assets
      if (event.node.req.url?.includes('/_nuxt/')) {
        headers['Cache-Control'] = 'public,max-age=31536000,immutable'
      }

      // Add cache headers for API responses
      if (event.node.req.url?.startsWith('/api/')) {
        headers['Cache-Control'] = 'public,max-age=60,stale-while-revalidate=300'
      }
    }
  })
})
