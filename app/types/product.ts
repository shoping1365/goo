export interface Product {
  id: number
  name: string
  description: string
  price: number
  image: string
  category: string
  rating: number
  stock: number
  brand: string
  discount?: number
  isNew?: boolean
  isFeatured?: boolean
  disable_buy_button?: boolean
  call_for_price?: boolean
  isSpecial?: boolean
  isFavorite?: boolean
} 