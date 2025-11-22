export interface HeaderLayer {
  id?: number
  name: string
  width: number
  height: number
  rowCount: number
  color: string
  opacity: number
  showSeparator: boolean
  separatorType: string
  separatorColor: string
  separatorOpacity: number
  separatorWidth: number
  items: string[]
  createdAt?: string
  updatedAt?: string
  [key: string]: unknown
}

export interface Header {
  id?: number
  name: string
  description: string
  pageSelection: string
  specificPages: string
  excludedPages: string
  layers: HeaderLayer[]
  is_active?: boolean
  createdAt?: string
  updatedAt?: string
  [key: string]: unknown
} 