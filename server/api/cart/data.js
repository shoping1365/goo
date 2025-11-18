// Session-based cart data store (server-side)
// این پیاده‌سازی از session ID برای ذخیره سبد خرید هر کاربر استفاده می‌کند
// برای جلوگیری از از دست رفتن داده‌ها در ریست سرور، از فایل سیستم استفاده می‌کنیم

import fs from 'fs'
import path from 'path'

const CART_DATA_FILE = path.join(process.cwd(), 'data', 'carts.json')

// اطمینان از وجود پوشه data
if (!fs.existsSync(path.dirname(CART_DATA_FILE))) {
  fs.mkdirSync(path.dirname(CART_DATA_FILE), { recursive: true })
}

// بارگذاری داده‌ها از فایل
let carts = new Map()
let globalNextId = 1

function loadCartsFromFile() {
  try {
    if (fs.existsSync(CART_DATA_FILE)) {
      const data = fs.readFileSync(CART_DATA_FILE, 'utf8')
      const parsed = JSON.parse(data)
      carts = new Map(Object.entries(parsed.carts || {}))
      globalNextId = parsed.globalNextId || 1
    }
  } catch (error) {
    console.error('خطا در بارگذاری سبدهای خرید از فایل:', error)
    carts = new Map()
    globalNextId = 1
  }
}

// ذخیره داده‌ها در فایل
function saveCartsToFile() {
  try {
    const data = {
      carts: Object.fromEntries(carts),
      globalNextId
    }
    fs.writeFileSync(CART_DATA_FILE, JSON.stringify(data, null, 2))
  } catch (error) {
    console.error('خطا در ذخیره سبدهای خرید در فایل:', error)
  }
}

// بارگذاری اولیه
loadCartsFromFile()

// ذخیره خودکار هر 30 ثانیه
setInterval(saveCartsToFile, 30000)

export function getCartItems(sessionId) {
  if (!sessionId) return []

  const cart = carts.get(sessionId)
  if (!cart) {
    return []
  }

  return cart.items
}

export function addCartItem(sessionId, product) {
  if (!sessionId) return

  let cart = carts.get(sessionId)
  if (!cart) {
    cart = {
      items: [],
      nextId: 1
    }
    carts.set(sessionId, cart)
  }

  const existing = cart.items.find(ci => ci.product_id === product.product_id)
  if (existing) {
    existing.quantity += product.quantity || 1
  } else {
    cart.items.push({
      id: globalNextId++,
      ...product,
      quantity: product.quantity || 1
    })
  }
  
  // ذخیره فوری پس از تغییر
  saveCartsToFile()
}

export function updateCartItem(sessionId, cartItemId, quantity) {
  if (!sessionId) return

  const cart = carts.get(sessionId)
  if (!cart) return

  const item = cart.items.find(ci => ci.id === cartItemId)
  if (item) {
    item.quantity = quantity
    // ذخیره فوری پس از تغییر
    saveCartsToFile()
  }
}

export function removeCartItem(sessionId, cartItemId) {
  if (!sessionId) return

  const cart = carts.get(sessionId)
  if (!cart) return

  cart.items = cart.items.filter(ci => ci.id !== cartItemId)
  // ذخیره فوری پس از تغییر
  saveCartsToFile()
}

export function clearCart(sessionId) {
  if (!sessionId) return

  const cart = carts.get(sessionId)
  if (cart) {
    cart.items = []
    cart.nextId = 1
    // ذخیره فوری پس از تغییر
    saveCartsToFile()
  }
}

// انتقال آیتم به خرید بعدی
export function moveItemToNext(sessionId, cartItemId) {
  if (!sessionId) return false

  const cart = carts.get(sessionId)
  if (!cart) return false

  const item = cart.items.find(ci => ci.id === cartItemId)
  if (item) {
    item.is_next = true
    // ذخیره فوری پس از تغییر
    saveCartsToFile()
    return true
  }
  return false
}

// بازگرداندن آیتم از خرید بعدی به سبد
export function moveItemToCart(sessionId, cartItemId) {
  if (!sessionId) return false

  const cart = carts.get(sessionId)
  if (!cart) return false

  const item = cart.items.find(ci => ci.id === cartItemId)
  if (item) {
    item.is_next = false
    // ذخیره فوری پس از تغییر
    saveCartsToFile()
    return true
  }
  return false
}

// دریافت تمام سبدهای خرید (برای debug)
export function getAllCarts() {
  return Object.fromEntries(carts)
}

// حذف سبد خرید یک session (برای cleanup)
export function removeCart(sessionId) {
  carts.delete(sessionId)
  // ذخیره فوری پس از تغییر
  saveCartsToFile()
} 