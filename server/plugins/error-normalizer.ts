export default defineNitroPlugin((nitroApp) => {
  // Only log and best-effort normalize by mutating, do NOT return custom objects
  nitroApp.hooks.hook('error', (error: any) => {
    if (!(error instanceof Error)) {
      try {
        const normalized: any = createError({
          statusCode: error?.statusCode || error?.status || 500,
          message: error?.message || error?.statusMessage || 'Internal Server Error',
          data: error
        })
        // Attempt to copy useful fields onto original error (if it's an object)
        if (error && typeof error === 'object') {
          Object.assign(error, normalized)
        }
      } catch {
        // swallow
      }
    }
  })
})
