export interface MenuChild {
  id: string
  title: string
  path?: string
  icon?: string
  badge?: string
  enabled: boolean
  order?: number
}

export interface MenuItem {
  id: string
  title: string
  path?: string
  icon?: string
  badge?: string
  enabled: boolean
  order?: number
  children?: MenuChild[]
}

export interface MenuSettings {
  name: string
  location: 'header' | 'footer' | 'sidebar' | 'mobile'
  layout: 'horizontal' | 'vertical' | 'both'
  desktopLayout: 'horizontal' | 'vertical'
  mobileLayout: 'horizontal' | 'vertical'
  maxDepth: string
  theme: 'default' | 'modern' | 'minimal' | 'classic'
  enabled: boolean
  showIcons: boolean
  showBadges: boolean
}

export interface MenuConfiguration {
  id: string
  name: string
  items: MenuItem[]
  settings: MenuSettings
  createdAt?: string
  updatedAt?: string
} 