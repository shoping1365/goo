import { beforeEach, describe, expect, it, vi } from 'vitest'
import { ref } from 'vue'
import { useCart } from '../../app/composables/useCart'
import type { Product } from '../../app/types/widget'

// Mock Nuxt composables
const mockUseCookie = vi.fn()
const mockUseCSRF = vi.fn()
const mockFetch = vi.fn()

// Global mocks
vi.stubGlobal('useCookie', mockUseCookie)
vi.stubGlobal('$fetch', mockFetch)
vi.stubGlobal('useCSRF', () => ({
  getCSRFToken: mockUseCSRF
}))

describe('useCart Composable', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    
    // Default mocks
    mockUseCookie.mockReturnValue(ref('test-session-id'))
    mockUseCSRF.mockResolvedValue('test-csrf-token')
    
    // Reset shared state (if possible, though it's module level)
    // Since state is module-level, we might need to rely on fetchCart to reset it or just accept it persists
  })

  it('should initialize with default values', () => {
    const cart = useCart()
    expect(cart.cartItems.value).toEqual([])
    expect(cart.cartCount.value).toBe(0)
    expect(cart.cartTotal.value).toBe(0)
    expect(cart.loading.value).toBe(false)
  })

  describe('fetchCart', () => {
    it('should fetch and parse cart data correctly', async () => {
      const mockCartData = {
        items: [
          {
            id: 1,
            product_id: 101,
            quantity: 2,
            price: 1000,
            name: 'Test Product',
            image: '/test.jpg'
          }
        ],
        total_items: 2,
        total_price: 2000
      }

      mockFetch.mockResolvedValue({ data: mockCartData })

      const cart = useCart()
      await cart.fetchCart()

      expect(mockFetch).toHaveBeenCalledWith('/api/cart', expect.objectContaining({
        credentials: 'include'
      }))
      
      expect(cart.cartItems.value).toHaveLength(1)
      expect(cart.cartItems.value[0].id).toBe(1)
      expect(cart.cartItems.value[0].quantity).toBe(2)
      expect(cart.cartCount.value).toBe(2)
      expect(cart.cartTotal.value).toBe(2000)
    })

    it('should handle empty cart response', async () => {
      mockFetch.mockResolvedValue(null)

      const cart = useCart()
      await cart.fetchCart()

      expect(cart.cartItems.value).toEqual([])
      expect(cart.cartCount.value).toBe(0)
      expect(cart.cartTotal.value).toBe(0)
    })

    it('should handle fetch error gracefully', async () => {
      mockFetch.mockRejectedValue(new Error('Network error'))

      const cart = useCart()
      await cart.fetchCart()

      expect(cart.cartItems.value).toEqual([])
      expect(cart.cartCount.value).toBe(0)
      expect(cart.loading.value).toBe(false)
    })
  })

  describe('addToCart', () => {
    it('should add item successfully', async () => {
      mockFetch.mockResolvedValueOnce({ success: true, message: 'Added' }) // addToCart response
      mockFetch.mockResolvedValueOnce({ data: { items: [], total_items: 1, total_price: 100 } }) // fetchCart response

      const cart = useCart()
      const result = await cart.addToCart(101, 1, { name: 'Product 1', price: 100 } as Product)

      expect(mockUseCSRF).toHaveBeenCalled()
      expect(mockFetch).toHaveBeenCalledWith('/api/cart/add', expect.objectContaining({
        method: 'POST',
        body: expect.objectContaining({
          product_id: 101,
          quantity: 1
        })
      }))
      expect(result).toEqual({ success: true, message: 'Added' })
    })

    it('should validate input parameters', async () => {
      const cart = useCart()
      
      const resultInvalidId = await cart.addToCart(0, 1)
      expect(resultInvalidId.success).toBe(false)
      expect(resultInvalidId.message).toContain('شناسه محصول نامعتبر است')

      const resultInvalidQty = await cart.addToCart(101, 0)
      expect(resultInvalidQty.success).toBe(false)
      expect(resultInvalidQty.message).toContain('تعداد محصول نامعتبر است')
    })

    it('should handle CSRF token missing', async () => {
      mockUseCSRF.mockResolvedValue(null)
      
      const cart = useCart()
      const result = await cart.addToCart(101, 1)

      expect(result.success).toBe(false)
      expect(result.message).toContain('CSRF token')
    })
  })

  describe('updateCartItem', () => {
    it('should update item quantity successfully', async () => {
      mockFetch.mockResolvedValueOnce({ success: true, message: 'Updated' })
      mockFetch.mockResolvedValueOnce({ data: { items: [], total_items: 1, total_price: 100 } })

      const cart = useCart()
      const result = await cart.updateCartItem(1, 5)

      expect(mockFetch).toHaveBeenCalledWith('/api/cart/update', expect.objectContaining({
        method: 'PUT',
        body: expect.objectContaining({
          cart_item_id: 1,
          quantity: 5
        })
      }))
      expect(result.success).toBe(true)
    })
  })

  describe('removeFromCart', () => {
    it('should remove item successfully', async () => {
      mockFetch.mockResolvedValueOnce({ success: true, message: 'Removed' })
      mockFetch.mockResolvedValueOnce({ data: { items: [], total_items: 0, total_price: 0 } })

      const cart = useCart()
      const result = await cart.removeFromCart(1)

      expect(mockFetch).toHaveBeenCalledWith('/api/cart/remove', expect.objectContaining({
        method: 'DELETE',
        body: expect.objectContaining({
          cart_item_id: 1
        })
      }))
      expect(result.success).toBe(true)
    })
  })

  describe('Helper Methods', () => {
    it('should check if item is in cart', async () => {
      // Setup cart state
      const mockCartData = {
        items: [{ id: 1, product_id: 101, quantity: 1 }],
        total_items: 1,
        total_price: 100
      }
      mockFetch.mockResolvedValue({ data: mockCartData })
      
      const cart = useCart()
      await cart.fetchCart()

      expect(cart.isInCart(101)).toBe(true)
      expect(cart.isInCart(999)).toBe(false)
    })

    it('should get item quantity', async () => {
      // Setup cart state
      const mockCartData = {
        items: [{ id: 1, product_id: 101, quantity: 5 }],
        total_items: 5,
        total_price: 500
      }
      mockFetch.mockResolvedValue({ data: mockCartData })
      
      const cart = useCart()
      await cart.fetchCart()

      expect(cart.getItemQuantity(101)).toBe(5)
      expect(cart.getItemQuantity(999)).toBe(0)
    })
  })
})
