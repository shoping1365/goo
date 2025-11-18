// Simple in-memory cart data store (server-side)
// توجه: این پیاده‌سازی صرفاً برای محیط توسعه/ماک است و باید در آینده با دیتابیس جایگزین شود.

let cartItems = []
let nextId = 1

export function getCartItems() {
  return cartItems
}

export function addCartItem(product) {
  const existing = cartItems.find(ci => ci.product_id === product.product_id)
  if (existing) {
    existing.quantity += product.quantity || 1
  } else {
    cartItems.push({
      id: nextId++,
      ...product,
      quantity: product.quantity || 1
    })
  }
}

export function updateCartItem(cartItemId, quantity) {
  const item = cartItems.find(ci => ci.id === cartItemId)
  if (item) {
    item.quantity = quantity
  }
}

export function removeCartItem(cartItemId) {
  cartItems = cartItems.filter(ci => ci.id !== cartItemId)
}

export function clearCart() {
  cartItems = []
  nextId = 1
} 