export interface Product {
  id: number
  sku?: string
  slug?: string
  name: string
  description?: string
  full_description?: string
  meta_description?: string
  seo_title?: string
  price: number
  image?: string
  category?: string | {
    name: string
    slug: string
  }
  rating?: number
  stock?: number
  stock_quantity?: number
  in_stock?: boolean
  brand?: string
  discount?: number
  isNew?: boolean
  isFeatured?: boolean
  disable_buy_button?: boolean
  call_for_price?: boolean
  isSpecial?: boolean
  isFavorite?: boolean
  show_stock_to_customer?: boolean
  track_inventory?: boolean
  vip_only?: boolean
  hide_price_until_login?: boolean
}