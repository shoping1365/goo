import { beforeEach, describe, expect, it, vi } from 'vitest'
import chunkErrorHandler from '../../app/middleware/chunk-error-handler.global'

vi.mock('nuxt/app', () => ({
  defineNuxtRouteMiddleware: (handler: unknown) => handler
}))

describe('Chunk Error Handler Middleware', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('should add error listener on client side', () => {
    vi.stubGlobal('import.meta', { client: true })
    const addEventListenerSpy = vi.spyOn(window, 'addEventListener')

    const to = {} as unknown
    const from = {} as unknown

    // @ts-expect-error - Mocking route
    chunkErrorHandler(to, from)

    expect(addEventListenerSpy).toHaveBeenCalledWith('error', expect.any(Function))
  })

  it('should reload page on chunk load error', () => {
    vi.stubGlobal('import.meta', { client: true })
    
    // Mock window.location.reload
    const reloadMock = vi.fn()
    Object.defineProperty(window, 'location', {
      value: { reload: reloadMock },
      writable: true
    })

    // Capture the event listener
    let errorListener: EventListener | undefined
    vi.spyOn(window, 'addEventListener').mockImplementation((event, handler) => {
      if (event === 'error') {
        errorListener = handler as EventListener
      }
    })

    const to = {} as unknown
    const from = {} as unknown

    // @ts-expect-error - Mocking route
    chunkErrorHandler(to, from)

    // Simulate error event
    const mockEvent = {
      error: { message: 'Loading chunk failed' },
      preventDefault: vi.fn()
    }

    vi.useFakeTimers()
    if (errorListener) {
      errorListener(mockEvent as unknown as Event)
    }
    
    expect(mockEvent.preventDefault).toHaveBeenCalled()
    
    vi.advanceTimersByTime(500)
    expect(reloadMock).toHaveBeenCalled()
    
    vi.useRealTimers()
  })
})
