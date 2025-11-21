import { computed, ref } from 'vue'

export interface MenuItem {
  id?: number
  clientId?: number
  menuId?: number
  title: string
  path: string
  icon?: string
  iconType?: 'icon' | 'image'
  enabled: boolean
  itemType: 'page' | 'post' | 'custom' | 'category' | 'product_category'
  order?: number
  parentId?: number | null
  parentClientId?: number | null
  depth?: number
  showOptions?: boolean
  children?: MenuItem[]
  expanded?: boolean
  openInNewTab?: boolean
  badge?: string
  badgeColor?: string
  isMegaMenu?: boolean
  megaWidth?: string
  megaColumns?: number
  description?: string
  imageURL?: string
  featured?: boolean
  showOnMobile?: boolean
  showOnDesktop?: boolean
  showOnTablet?: boolean
  isNew?: boolean
}

export interface Menu {
  id?: number
  name: string
  slug?: string
  items: MenuItem[]
  enabled?: boolean
  order?: number
}

export interface MenuLocation {
  id: number
  name: string
  slug: string
  description?: string
}

// Helper interface for raw data from API or other sources
interface RawMenuItem {
  id?: number
  item_id?: number
  clientId?: number
  client_id?: number
  menuId?: number
  menu_id?: number
  menuID?: number
  menu?: number
  title?: string
  name?: string
  label?: string
  path?: string
  url?: string
  link?: string
  icon?: string
  icon_class?: string
  iconName?: string
  iconType?: string
  icon_type?: string
  itemType?: string
  item_type?: string
  type?: string
  enabled?: boolean
  is_active?: boolean
  status?: string
  state?: string
  order?: number
  sort_order?: number
  parentId?: number
  parent_id?: number
  parentClientId?: number
  parent_client_id?: number
  depth?: number
  level?: number
  openInNewTab?: boolean
  open_in_new_tab?: boolean
  target?: string
  badge?: string
  badge_text?: string
  badgeColor?: string
  badge_color?: string
  isMegaMenu?: boolean
  is_mega_menu?: boolean
  megaWidth?: string
  mega_width?: string
  megaColumns?: number
  mega_columns?: number
  description?: string
  imageURL?: string
  image_url?: string
  featured?: boolean
  showOnMobile?: boolean
  show_on_mobile?: boolean
  showOnDesktop?: boolean
  show_on_desktop?: boolean
  showOnTablet?: boolean
  show_on_tablet?: boolean
  children?: RawMenuItem[]
  childrens?: RawMenuItem[]
  childCategories?: RawMenuItem[]
  child_categories?: RawMenuItem[]
  nodes?: RawMenuItem[]
  items?: RawMenuItem[]
  menuItems?: RawMenuItem[]
  isNew?: boolean
  slug?: string
  identifier?: string
  category_id?: number
  categoryId?: number
  parentID?: number
  parent?: number
  parentCategoryId?: number
  parent_category_id?: number
  targetId?: number
  target_id?: number
  targetType?: string
  target_type?: string
  displayName?: string
  [key: string]: unknown
}

export const useMenuManagement = () => {
  const menus = ref<Menu[]>([])
  const currentMenu = ref<Menu>({
    id: undefined,
    name: '',
    items: []
  })
  const deletedMenuItemIds = ref<number[]>([])
  let tempClientIdSequence = 0
  const generateTempClientId = (): number => {
    tempClientIdSequence = (tempClientIdSequence + 1) % 1000
    return Date.now() * 1000 + tempClientIdSequence
  }


  const isLoading = ref(false)
  const isSaving = ref(false)

  // Content data - wrapped in computed to ensure arrays
  const _pages = ref<RawMenuItem[]>([])
  const _posts = ref<RawMenuItem[]>([])
  const _categories = ref<RawMenuItem[]>([])
  const _productCategories = ref<RawMenuItem[]>([])

  const pages = computed(() => Array.isArray(_pages.value) ? _pages.value : [])
  const posts = computed(() => Array.isArray(_posts.value) ? _posts.value : [])
  const categories = computed(() => Array.isArray(_categories.value) ? _categories.value : [])
  const productCategories = computed(() => Array.isArray(_productCategories.value) ? _productCategories.value : [])

  const menuLocations = ref<MenuLocation[]>([
    { id: 1, name: 'منوی اصلی', slug: 'main-menu', description: 'منوی اصلی سایت' },
    { id: 2, name: 'منوی فوتر', slug: 'footer-menu', description: 'منوی پایین صفحه' },
    { id: 3, name: 'منوی کاربری', slug: 'user-menu', description: 'منوی کاربران' },
    { id: 4, name: 'مگامنو', slug: 'mega-menu', description: 'منوی بزرگ با ستون‌های مختلف' }
  ])

  const resetDeletedMenuItems = () => {
    deletedMenuItemIds.value = []
  }

  const markMenuItemDeleted = (id?: number | null) => {
    if (typeof id !== 'number' || id <= 0) return
    if (!deletedMenuItemIds.value.includes(id)) {
      deletedMenuItemIds.value.push(id)
    }
  }

  const registerDeletedMenuBranch = (item?: Partial<MenuItem> | null) => {
    if (!item) return
    const numericId = typeof item?.id === 'number' ? item.id : null
    const isClientOnly = (item as RawMenuItem)?.isNew === true || numericId === null || numericId <= 0
    if (!isClientOnly && typeof numericId === 'number') {
      markMenuItemDeleted(numericId)
    }
    if (isClientOnly) {
      return
    }
    const children = (item as RawMenuItem)?.children
    if (Array.isArray(children)) {
      (children as Partial<MenuItem>[]).forEach(child => registerDeletedMenuBranch(child))
    }
  }

  // API Methods
  const mapIconType = (value: unknown): MenuItem['iconType'] => {
    if (typeof value !== 'string') {
      return 'icon'
    }

    return value.toLowerCase() === 'image' ? 'image' : 'icon'
  }

  const mapItemType = (value: unknown): MenuItem['itemType'] => {
    if (typeof value !== 'string') {
      return 'custom'
    }

    const normalized = value.toLowerCase().replace(/-/g, '_')

    switch (normalized) {
      case 'page':
      case 'post':
      case 'custom':
      case 'category':
      case 'product_category':
        return normalized as MenuItem['itemType']
      default:
        return 'custom'
    }
  }

  const normalizeIncomingItem = (item: Partial<MenuItem>, order: number): MenuItem => {
    const raw = item as RawMenuItem
    return {
      id: undefined,
      clientId: generateTempClientId(),
      menuId: currentMenu.value.id,
      order,
      title: (item?.title ?? raw?.name ?? 'بدون عنوان').toString(),
      path: (item?.path ?? raw?.url ?? '').toString(),
      icon: item?.icon,
      iconType: mapIconType(item?.iconType ?? raw?.iconType ?? raw?.icon_type),
      enabled: item?.enabled !== false,
      itemType: mapItemType(item?.itemType ?? raw?.itemType ?? raw?.item_type ?? raw?.type),
      parentId: null,
      parentClientId: null,
      depth: 0,
      showOptions: false,
      children: [],
      expanded: false,
      openInNewTab: item?.openInNewTab ?? false,
      badge: item?.badge ?? '',
      badgeColor: item?.badgeColor ?? 'red',
      isMegaMenu: item?.isMegaMenu ?? false,
      megaWidth: item?.megaWidth ?? 'full',
      megaColumns: item?.megaColumns ?? 2,
      description: item?.description ?? '',
      imageURL: item?.imageURL ?? '',
      featured: item?.featured ?? false,
      showOnMobile: item?.showOnMobile ?? true,
      showOnDesktop: item?.showOnDesktop ?? true,
      showOnTablet: item?.showOnTablet ?? true,
      isNew: true
    }
  }

  const normalizeNumericId = (value: unknown): number | null => {
    if (value === null || value === undefined) return null
    if (typeof value === 'number') {
      return Number.isFinite(value) ? value : null
    }
    const numeric = Number(value)
    return Number.isFinite(numeric) ? numeric : null
  }

  const resolveCategoryParentId = (item: RawMenuItem): number | null => {
    const candidates = [
      item?.parentId,
      item?.parent_id,
      item?.parentID,
      item?.parent,
      item?.parentCategoryId,
      item?.parent_category_id
    ]
    for (const candidate of candidates) {
      const resolved = normalizeNumericId(candidate)
      if (resolved && resolved > 0) {
        return resolved
      }
    }
    return null
  }

  const getCategoryLabel = (item: RawMenuItem): string => {
    const name = (item?.name ?? item?.title ?? '').toString().trim()
    if (name) {
      return name
    }
    const slug = (item?.slug ?? '').toString().trim()
    if (slug) {
      return slug
    }
    return 'بدون عنوان'
  }

  const computeItemSignature = (item: Partial<MenuItem>): string => {
    const typeKey = (item?.itemType ?? '').toString().toLowerCase()
    const pathKey = (item?.path ?? '').toString().trim().toLowerCase()
    const titleKey = (item?.title ?? '').toString().trim().toLowerCase()
    const parentKey = item?.parentId ?? item?.parentClientId ?? 'root'
    const depthKey = item?.depth ?? 0

    // Signature شامل parent و depth هم میشه تا آیتم‌های یکسان در سطوح مختلف unique باشن
    return `${typeKey}::${pathKey || titleKey}::parent-${parentKey}::depth-${depthKey}`
  }

  const menuContainsSignature = (items: MenuItem[] | undefined, signature: string): boolean => {
    if (!Array.isArray(items) || !signature) return false
    for (const item of items) {
      if (!item) continue
      if (computeItemSignature(item) === signature) {
        return true
      }
      if (menuContainsSignature(item.children, signature)) {
        return true
      }
    }
    return false
  }

  const buildIndentedLabel = (base: string, depth: number): string => {
    if (depth <= 0) {
      return base
    }
    const spacer = '  '.repeat(Math.max(0, depth - 1))
    return `${spacer}-> ${base}`
  }

  const buildHierarchicalOptions = (items: RawMenuItem[]): RawMenuItem[] => {
    if (!Array.isArray(items) || items.length === 0) {
      return []
    }

    const getChildrenArray = (item: RawMenuItem): RawMenuItem[] => {
      const candidates = [
        item?.children,
        item?.childrens,
        item?.childCategories,
        item?.child_categories,
        item?.nodes
      ]
      for (const candidate of candidates) {
        if (Array.isArray(candidate) && candidate.length > 0) {
          return candidate
        }
      }
      return []
    }

    const hasExplicitChildren = items.some((item) => getChildrenArray(item).length > 0)

    const cloneWithMeta = (item: RawMenuItem, depth: number) => {
      const normalizedId = normalizeNumericId(item?.id ?? item?.category_id ?? item?.categoryId)
      const label = getCategoryLabel(item)
      return {
        ...item,
        id: normalizedId ?? item?.id,
        depth,
        displayName: buildIndentedLabel(label, depth)
      }
    }

    if (hasExplicitChildren) {
      const result: RawMenuItem[] = []
      const traverse = (list: RawMenuItem[], depth: number) => {
        if (!Array.isArray(list)) return
        list.forEach((raw) => {
          const cloned = cloneWithMeta(raw, depth)
          result.push(cloned)
          const childList = getChildrenArray(raw)
          if (childList.length > 0) {
            traverse(childList, depth + 1)
          }
        })
      }

      traverse(items, 0)
      return result
    }

    const nodes = items.map((item, index) => {
      const normalizedId = normalizeNumericId(item?.id ?? item?.category_id ?? item?.categoryId ?? index)
      return {
        id: normalizedId,
        parentId: resolveCategoryParentId(item),
        original: item,
        order: index
      }
    })

    const byParent = new Map<number | null, typeof nodes>()
    nodes.forEach((entry) => {
      const key = entry.parentId ?? null
      const bucket = byParent.get(key)
      if (bucket) {
        bucket.push(entry)
      } else {
        byParent.set(key, [entry])
      }
    })

    const sortByLabel = (list: typeof nodes): typeof nodes => {
      return [...list].sort((a, b) => {
        const labelA = getCategoryLabel(a.original)
        const labelB = getCategoryLabel(b.original)
        return labelA.localeCompare(labelB, 'fa', { sensitivity: 'base' })
      })
    }

    const result: RawMenuItem[] = []
    const visited = new Set<RawMenuItem>()

    const visit = (entry: typeof nodes[number], depth: number) => {
      if (!entry || visited.has(entry.original)) {
        return
      }
      visited.add(entry.original)

      const clone = cloneWithMeta(entry.original, depth)
      result.push(clone)

      if (entry.id !== null && entry.id !== undefined) {
        const children = byParent.get(entry.id)
        if (children && children.length > 0) {
          sortByLabel(children).forEach((child) => visit(child, depth + 1))
        }
      }
    }

    const roots = byParent.get(null) ? sortByLabel(byParent.get(null)!) : []
    roots.forEach((root) => visit(root, 0))

    nodes.forEach((entry) => {
      if (!visited.has(entry.original)) {
        visit(entry, 0)
      }
    })

    return result
  }

  const normalizeMenuItem = (
    raw: RawMenuItem,
    index: number = 0,
    parentIdOverride: number | null = null,
    menuIdOverride?: number
  ): MenuItem => {
    const title = raw?.title ?? raw?.name ?? raw?.label ?? ''
    const path = raw?.path ?? raw?.url ?? raw?.link ?? ''
    const icon = raw?.icon ?? raw?.icon_class ?? raw?.iconName ?? undefined
    const iconType = mapIconType(raw?.iconType ?? raw?.icon_type)
    const rawType = raw?.itemType ?? raw?.item_type ?? raw?.type
    const itemType = mapItemType(rawType)

    // Determine enabled status - default to true if not explicitly set to false
    let enabled = true
    if (raw?.enabled === false || raw?.is_active === false) {
      enabled = false
    } else if (raw?.status === 'inactive' || raw?.status === 'disabled' || raw?.state === 'inactive') {
      enabled = false
    }

    const persistedId = raw?.id ?? raw?.item_id
    const clientId = raw?.clientId ?? raw?.client_id ?? persistedId ?? index
    const menuId = menuIdOverride ?? raw?.menuId ?? raw?.menu_id ?? raw?.menuID ?? raw?.menu ?? undefined

    return {
      id: persistedId,
      clientId,
      menuId,
      title: title || 'بدون عنوان',
      path,
      icon,
      iconType,
      enabled: enabled,
      itemType,
      order: raw?.order ?? raw?.sort_order ?? index + 1,
      parentId: parentIdOverride ?? raw?.parentId ?? raw?.parent_id ?? null,
      parentClientId: raw?.parentClientId ?? raw?.parent_client_id ?? null,
      depth: typeof raw?.depth === 'number' ? raw.depth : (typeof raw?.level === 'number' ? raw.level : 0),
      showOptions: false,
      expanded: false,
      openInNewTab: raw?.openInNewTab ?? raw?.open_in_new_tab ?? raw?.target === '_blank',
      badge: raw?.badge ?? raw?.badge_text ?? '',
      badgeColor: raw?.badgeColor ?? raw?.badge_color ?? 'gray',
      isMegaMenu: raw?.isMegaMenu ?? raw?.is_mega_menu ?? false,
      megaWidth: raw?.megaWidth ?? raw?.mega_width ?? 'full',
      megaColumns: raw?.megaColumns ?? raw?.mega_columns ?? 1,
      description: raw?.description ?? '',
      imageURL: raw?.imageURL ?? raw?.image_url ?? '',
      featured: raw?.featured ?? false,
      showOnMobile: raw?.showOnMobile ?? raw?.show_on_mobile ?? true,
      showOnDesktop: raw?.showOnDesktop ?? raw?.show_on_desktop ?? true,
      showOnTablet: raw?.showOnTablet ?? raw?.show_on_tablet ?? true,
      isNew: false
    }
  }

  const flattenMenuItems = (items: RawMenuItem[], parentId: number | null = null, menuId?: number): MenuItem[] => {
    if (!Array.isArray(items)) {
      return []
    }

    const flattened: MenuItem[] = []

    items.forEach((item, index) => {
      const normalized = normalizeMenuItem(item, index, parentId, menuId)
      flattened.push(normalized)

      const rawChildren = item?.children || item?.childrens || []
      if (Array.isArray(rawChildren) && rawChildren.length > 0) {
        const nextParentId = item?.id ?? normalized.id ?? parentId
        flattened.push(...flattenMenuItems(rawChildren, nextParentId ?? null, normalized.menuId ?? menuId))
      }
    })

    return flattened
  }

  const buildMenuTree = (flatItems: MenuItem[]): MenuItem[] => {
    if (!Array.isArray(flatItems) || flatItems.length === 0) {
      return []
    }

    const sorted = [...flatItems].sort((a, b) => (a.order || 0) - (b.order || 0))
    const itemMap = new Map<number, MenuItem>()
    const rootItems: MenuItem[] = []

    sorted.forEach(item => {
      const itemCopy = { ...item, children: [] }
      if (item.id) {
        itemMap.set(item.id, itemCopy)
      }
    })

    sorted.forEach(item => {
      const itemCopy = itemMap.get(item.id!) || { ...item, children: [] }

      if (!item.parentId) {
        rootItems.push(itemCopy)
      } else {
        const parent = itemMap.get(item.parentId)
        if (parent) {
          if (!parent.children) parent.children = []
          parent.children.push(itemCopy)
        } else {
          rootItems.push(itemCopy)
        }
      }
    })

    return rootItems
  }

  const normalizeMenu = (raw: unknown): Menu => {
    const data = raw as RawMenuItem
    if (!data) {
      return { id: undefined, name: '', slug: '', items: [] }
    }

    const menuId = data?.id ?? data?.menu_id
    const rawItems = data?.items ?? data?.menuItems ?? []

    // First flatten all items
    const flatItems = flattenMenuItems(rawItems, null, menuId)

    // Then build tree structure
    const treeItems = buildMenuTree(flatItems)

    return {
      id: data.id ?? data.menu_id ?? undefined,
      name: data.name ?? data.title ?? '',
      slug: data.slug ?? data.identifier ?? '',
      enabled: data.enabled !== false && data.is_active !== false && data.status !== 'inactive',
      order: data.order ?? data.sort_order ?? undefined,
      items: treeItems
    }
  }

  const fetchMenus = async () => {
    isLoading.value = true
    try {
      const response = await $fetch<{ data: unknown[] }>('/api/admin/menus')
      const data = response?.data || response || []
      menus.value = Array.isArray(data) ? data.map(normalizeMenu) : []
    } catch (error) {
      // console.error('Error fetching menus:', error)
      throw error
    } finally {
      isLoading.value = false
    }
  }

  const fetchMenu = async (id: number) => {
    isLoading.value = true
    try {
      const response = await $fetch<{ data: unknown }>(`/api/admin/menus/${id}`)
      const data = response?.data || response || {}
      currentMenu.value = normalizeMenu(data)
      resetDeletedMenuItems()
    } catch (error) {
      // console.error('Error fetching menu:', error)
      throw error
    } finally {
      isLoading.value = false
    }
  }

  // Helper function to flatten menu items tree into array with parentId
  const flattenMenuItemsForSave = (items: MenuItem[], parentLink: number | null = null, currentOrder = 0): Record<string, unknown>[] => {
    const result: Record<string, unknown>[] = []
    let order = currentOrder

    items.forEach((item) => {
      order++
      const source = item as RawMenuItem
      const persistedId = typeof source?.id === 'number' && source.id > 0 ? source.id : undefined
      const clientId = typeof source?.clientId === 'number' ? source.clientId : undefined
      const shouldCreate = source?.isNew || !persistedId
      const parentDbId = typeof source?.parentId === 'number' ? source.parentId : null
      const parentClientId = typeof source?.parentClientId === 'number' ? source.parentClientId : null
      const depth = typeof source?.depth === 'number' ? source.depth : 0
      const effectiveParentLink = typeof parentLink === 'number' ? parentLink : null
      const resolvedParentClientId = parentClientId ?? effectiveParentLink

      // Add current item
      const flatItem = {
        id: shouldCreate ? undefined : persistedId,
        title: source?.title ?? source?.name ?? '',
        name: source?.title ?? source?.name ?? '',
        path: source?.path ?? source?.url ?? '',
        url: source?.path ?? source?.url ?? '',
        icon: source?.icon ?? null,
        icon_type: source?.iconType ?? source?.icon_type ?? 'icon',
        itemType: source?.itemType ?? source?.item_type ?? 'custom',
        item_type: source?.itemType ?? source?.item_type ?? 'custom',
        enabled: source?.enabled !== false,
        is_active: source?.enabled !== false,
        status: source?.enabled !== false ? 'active' : 'inactive',
        order: order,
        sort_order: order,
        parentId: parentDbId,
        parent_id: parentDbId,
        parentClientId: resolvedParentClientId,
        parent_client_id: resolvedParentClientId,
        clientId: clientId ?? null,
        client_id: clientId ?? null,
        depth,
        openInNewTab: source?.openInNewTab ?? source?.open_in_new_tab ?? false,
        open_in_new_tab: source?.openInNewTab ?? source?.open_in_new_tab ?? false,
        badge: source?.badge ?? source?.badge_text ?? '',
        badge_text: source?.badge ?? source?.badge_text ?? '',
        badgeColor: source?.badgeColor ?? source?.badge_color ?? 'gray',
        badge_color: source?.badgeColor ?? source?.badge_color ?? 'gray',
        isMegaMenu: source?.isMegaMenu ?? source?.is_mega_menu ?? false,
        is_mega_menu: source?.isMegaMenu ?? source?.is_mega_menu ?? false,
        megaWidth: source?.megaWidth ?? source?.mega_width ?? 'full',
        mega_width: source?.megaWidth ?? source?.mega_width ?? 'full',
        megaColumns: source?.megaColumns ?? source?.mega_columns ?? 1,
        mega_columns: source?.megaColumns ?? source?.mega_columns ?? 1,
        description: source?.description ?? '',
        imageURL: source?.imageURL ?? source?.image_url ?? '',
        image_url: source?.imageURL ?? source?.image_url ?? '',
        featured: source?.featured ?? false,
        showOnMobile: source?.showOnMobile ?? source?.show_on_mobile ?? true,
        show_on_mobile: source?.showOnMobile ?? source?.show_on_mobile ?? true,
        showOnDesktop: source?.showOnDesktop ?? source?.show_on_desktop ?? true,
        show_on_desktop: source?.showOnDesktop ?? source?.show_on_desktop ?? true,
        showOnTablet: source?.showOnTablet ?? source?.show_on_tablet ?? true,
        show_on_tablet: source?.showOnTablet ?? source?.show_on_tablet ?? true
      }

      result.push(flatItem)

      // Recursively process children
      if (item.children && item.children.length > 0) {
        const childItems = flattenMenuItemsForSave(
          item.children,
          persistedId ?? clientId ?? parentLink ?? null,
          order
        )
        result.push(...childItems)
        order += childItems.length
      }
    })

    return result
  }

  const buildMenuItemPayload = (item: RawMenuItem, menuId: number, parentId: number | null) => {
    const title = item?.title ?? item?.name ?? ''
    const path = item?.path ?? item?.url ?? ''
    const badge = item?.badge ?? item?.badge_text ?? ''
    const badgeColor = item?.badgeColor ?? item?.badge_color ?? 'red'
    const iconType = item?.iconType ?? item?.icon_type ?? 'icon'
    const itemType = item?.itemType ?? item?.item_type ?? 'custom'
    const enabled = item?.enabled !== false
    const orderValue = typeof item?.order === 'number' ? item.order : 0
    const openInNewTab = item?.openInNewTab ?? item?.open_in_new_tab ?? false
    const showOnMobile = item?.showOnMobile ?? item?.show_on_mobile ?? true
    const showOnDesktop = item?.showOnDesktop ?? item?.show_on_desktop ?? true
    const showOnTablet = item?.showOnTablet ?? item?.show_on_tablet ?? true

    return {
      menu_id: menuId,
      menuId,
      parent_id: parentId,
      parentId,
      title,
      name: title,
      path,
      url: path,
      icon: item?.icon ?? null,
      icon_type: iconType,
      item_type: itemType,
      itemType,
      badge,
      badge_text: badge,
      badgeColor,
      badge_color: badgeColor,
      enabled,
      is_active: enabled,
      status: enabled ? 'active' : 'inactive',
      order: orderValue,
      sort_order: orderValue,
      openInNewTab,
      open_in_new_tab: openInNewTab,
      showOnMobile,
      show_on_mobile: showOnMobile,
      showOnDesktop,
      show_on_desktop: showOnDesktop,
      showOnTablet,
      show_on_tablet: showOnTablet,
      imageURL: item?.imageURL ?? item?.image_url ?? '',
      image_url: item?.imageURL ?? item?.image_url ?? '',
      description: item?.description ?? '',
      featured: item?.featured ?? false,
      depth: item?.depth ?? 0,
      target_id: item?.targetId ?? item?.target_id ?? null,
      target_type: item?.targetType ?? item?.target_type ?? null
    }
  }

  const syncMenuItems = async (menuId: number | undefined, items: RawMenuItem[]) => {
    if (!menuId || !Array.isArray(items)) return

    const idMap = new Map<number | string, number>()
    const processedIds = new Set<number>()

    const registerPersisted = (key?: number | null, value?: number | null) => {
      if (typeof key === 'number' && typeof value === 'number') {
        idMap.set(key, value)
      }
    }

    const walkItems = async (itemList: RawMenuItem[], parentId: number | null = null) => {
      for (const item of itemList) {
        const hasNumericId = typeof item?.id === 'number'
        const persistedId = hasNumericId ? item.id : null
        const clientId = typeof item?.clientId === 'number' ? item.clientId : null
        const isPersisted = !item?.isNew && typeof persistedId === 'number' && persistedId > 0

        let resolvedParentId = parentId
        if (resolvedParentId === null && item?.parentId != null) {
          resolvedParentId = typeof item.parentId === 'number' ? item.parentId : null
        }
        if ((resolvedParentId === null || resolvedParentId === 0) && item?.parentClientId != null) {
          const mapped = idMap.get(item.parentClientId)
          if (typeof mapped === 'number') {
            resolvedParentId = mapped
          }
        }

        const payload = buildMenuItemPayload(item, menuId, resolvedParentId)

        let currentId = isPersisted ? persistedId : null
        if (isPersisted && currentId) {
          try {
            await $fetch(`/api/admin/menu-items/${currentId}`, {
              method: 'PUT',
              body: payload
            })
          } catch (error) {
            // console.error('Failed to update menu item', {
            //   id: currentId,
            //   clientId,
            //   parentClientId: item?.parentClientId,
            //   payload,
            //   error
            // })
            throw error
          }
        } else {
          try {
            const response = await $fetch<{ data: unknown }>('/api/admin/menu-items', {
              method: 'POST',
              body: payload
            })
            const created = (response?.data || response) as RawMenuItem
            currentId = typeof created?.id === 'number' ? created.id : null
          } catch (error) {
            // console.error('Failed to create menu item', {
            //   clientId,
            //   parentClientId: item?.parentClientId,
            //   payload,
            //   error
            // })
            throw error
          }
        }

        if (typeof currentId === 'number') {
          processedIds.add(currentId)
          registerPersisted(currentId, currentId)
          if (typeof clientId === 'number') registerPersisted(clientId, currentId)
          item.id = currentId
          if (item && typeof item === 'object') {
            ; (item as RawMenuItem).isNew = false
          }
        }

        if (Array.isArray(item?.children) && currentId) {
          await walkItems(item.children, currentId)
        }
      }
    }

    await walkItems(items)

    const uniqueDeletedIds = [...new Set(deletedMenuItemIds.value)].filter(id => typeof id === 'number' && Number.isFinite(id) && id > 0)
    const deleteTargets: number[] = []

    if (uniqueDeletedIds.length > 0) {
      for (const storedId of uniqueDeletedIds) {
        if (typeof storedId !== 'number') continue

        let targetId = storedId
        if (!processedIds.has(targetId)) {
          const mapped = idMap.get(storedId)
          if (typeof mapped === 'number') targetId = mapped
        }

        if (typeof targetId === 'number' && targetId > 0 && !processedIds.has(targetId)) {
          deleteTargets.push(targetId)
        }
      }
    }

    if (deleteTargets.length > 0) {
      const chunkSize = 6
      for (let idx = 0; idx < deleteTargets.length; idx += chunkSize) {
        const chunk = deleteTargets.slice(idx, idx + chunkSize)
        const results = await Promise.allSettled(
          chunk.map(targetId =>
            $fetch(`/api/admin/menu-items/${targetId}`, {
              method: 'DELETE'
            }).catch(error => {
              // اگر آیتم قبلاً حذف شده (404) یا وجود نداره، نادیده بگیر
              if (error?.status === 404 || error?.statusCode === 404) {
                // console.warn(`⚠️ Item ${targetId} not found (already deleted), skipping...`)
                return { deleted: true, id: targetId }
              }
              throw error
            })
          )
        )

        const rejected = results.find((result): result is PromiseRejectedResult => result.status === 'rejected')
        if (rejected) {
          // console.error('❌ Failed to delete menu items chunk:', {
          //   chunk,
          //   reason: rejected.reason
          // })
          throw rejected.reason
        }
      }
    }

    resetDeletedMenuItems()
  }

  const saveMenu = async (menu: Menu) => {
    isSaving.value = true
    try {
      // اطمینان از وجود slug
      if (!menu.slug || menu.slug === '') {
        menu.slug = generateSlug(menu.name)
      }

      // Flatten the tree structure
      const flattenedItems = flattenMenuItemsForSave(menu.items || [])

      // Add menuId to all items
      const itemsWithMenuId = flattenedItems.map(item => ({
        ...item,
        menuId: menu.id ?? undefined,
        menu_id: menu.id ?? undefined
      }))

      const payload = {
        name: menu.name?.trim?.() || '',
        slug: menu.slug?.trim?.() || generateSlug(menu.name || ''),
        enabled: menu.enabled !== false, // اضافه شد - حفظ وضعیت فعال/غیرفعال منو
        items: itemsWithMenuId
      }

      let response
      if (menu.id) {
        // Update existing menu
        response = await $fetch<{ data: unknown }>(`/api/admin/menus/${menu.id}`, {
          method: 'PUT',
          body: payload
        })
      } else {
        // Create new menu
        response = await $fetch<{ data: unknown }>('/api/admin/menus', {
          method: 'POST',
          body: payload
        })
      }

      const responseData = (response?.data || response) as RawMenuItem
      let resolvedMenuId = typeof responseData?.id === 'number' ? responseData.id : menu.id

      // Ensure menu id is available for syncing items
      if (!resolvedMenuId && typeof currentMenu.value.id === 'number') {
        resolvedMenuId = currentMenu.value.id
      }

      await syncMenuItems(resolvedMenuId, flattenedItems as RawMenuItem[])

      if (resolvedMenuId) {
        await fetchMenu(resolvedMenuId)
      } else {
        currentMenu.value = normalizeMenu(responseData)
        resetDeletedMenuItems()
      }

      await fetchMenus()

      return response
    } catch (error) {
      // console.error('Error saving menu:', error)
      throw error
    } finally {
      isSaving.value = false
    }
  }

  const deleteMenu = async (id: number) => {
    try {
      await $fetch(`/api/admin/menus/${id}`, {
        method: 'DELETE'
      })

      // Remove from local list
      menus.value = menus.value.filter(menu => menu.id !== id)

      // Clear current menu if it was the deleted one
      if (currentMenu.value.id === id) {
        currentMenu.value = { id: undefined, name: '', items: [] }
      }
    } catch (error) {
      // console.error('Error deleting menu:', error)
      throw error
    }
  }

  const reorderMenuItems = async (menuId: number, itemOrders: { id: number; order: number }[]) => {
    try {
      await $fetch(`/api/admin/menus/${menuId}/reorder`, {
        method: 'POST',
        body: itemOrders
      })
    } catch (error) {
      // console.error('Error reordering menu items:', error)
      throw error
    }
  }

  const assignMenuToLocation = async (menuId: number, locationId: number, order: number = 1) => {
    try {
      await $fetch('/api/admin/menu-assign', {
        method: 'POST',
        body: {
          menu_id: menuId,
          location_id: locationId,
          order
        }
      })
    } catch (error) {
      // console.error('Error assigning menu to location:', error)
      throw error
    }
  }

  const unassignMenuFromLocation = async (menuId: number, locationId: number) => {
    try {
      await $fetch('/api/admin/menu-assign', {
        method: 'DELETE',
        body: {
          menu_id: menuId,
          location_id: locationId
        }
      })
    } catch (error) {
      // console.error('Error unassigning menu from location:', error)
      throw error
    }
  }

  // Content API Methods
  const fetchPages = async () => {
    try {
      const response = await $fetch<{ data: unknown[] }>('/api/admin/menu-content/pages')
      const data = response?.data || response
      _pages.value = Array.isArray(data) ? (data as RawMenuItem[]) : []
    } catch {
      _pages.value = []
    }
  }

  const fetchPosts = async () => {
    try {
      const response = await $fetch<{ data: unknown[] }>('/api/admin/menu-content/posts')
      const data = response?.data || response
      _posts.value = Array.isArray(data) ? (data as RawMenuItem[]) : []
    } catch {
      _posts.value = []
    }
  }

  const fetchCategories = async () => {
    try {
      const response = await $fetch<{ data: unknown[] }>('/api/admin/menu-content/categories')
      const data = response?.data || response
      _categories.value = buildHierarchicalOptions(Array.isArray(data) ? (data as RawMenuItem[]) : [])
    } catch {
      _categories.value = []
    }
  }

  const fetchProductCategories = async () => {
    try {
      const response = await $fetch<{ data: unknown[] }>('/api/admin/menu-content/product-categories')
      const data = response?.data || response
      _productCategories.value = buildHierarchicalOptions(Array.isArray(data) ? (data as RawMenuItem[]) : [])
    } catch {
      _productCategories.value = []
    }
  }

  // Utility Methods
  const createNewMenu = () => {
    currentMenu.value = {
      id: undefined,
      name: 'منوی جدید',
      slug: 'new-menu-' + Date.now(), // تولید slug خودکار
      items: []
    }
    resetDeletedMenuItems()
  }

  // تابع تولید slug از نام
  const generateSlug = (name: string): string => {
    return name
      .toLowerCase()
      .replace(/[^\u0600-\u06FFa-zA-Z0-9\s-]/g, '') // حذف کاراکترهای غیرمجاز
      .replace(/\s+/g, '-') // تبدیل فاصله به خط تیره
      .replace(/-+/g, '-') // حذف خط تیره‌های تکراری
      .replace(/^-|-$/g, '') // حذف خط تیره از ابتدا و انتها
      + '-' + Date.now() // اضافه کردن timestamp برای یکتا بودن
  }

  const addMenuItem = (item: Partial<MenuItem>) => {
    const nextOrder = currentMenu.value.items.length + 1
    const prepared = normalizeIncomingItem(item, nextOrder)
    const signature = computeItemSignature(prepared)

    if (menuContainsSignature(currentMenu.value.items, signature)) {
      return
    }

    currentMenu.value.items.push(prepared)
  }

  const removeMenuItem = (index: number) => {
    const [removed] = currentMenu.value.items.splice(index, 1)
    if (removed) registerDeletedMenuBranch(removed)
    // Reorder remaining items
    currentMenu.value.items.forEach((item, idx) => {
      item.order = idx + 1
    })
  }

  const reorderItems = (fromIndex: number, toIndex: number) => {
    const items = [...currentMenu.value.items]
    const [movedItem] = items.splice(fromIndex, 1)
    items.splice(toIndex, 0, movedItem)

    // Update order
    items.forEach((item, index) => {
      item.order = index + 1
    })

    currentMenu.value.items = items
  }

  // Helper to get item by path
  const getItemByPath = (items: MenuItem[], path: number[]): MenuItem | null => {
    if (!path || path.length === 0) return null

    let current = items[path[0]]
    for (let i = 1; i < path.length; i++) {
      if (!current?.children || !current.children[path[i]]) return null
      current = current.children[path[i]]
    }
    return current
  }

  // Helper to remove item by path
  const removeItemByPath = (items: MenuItem[], path: number[]): MenuItem | null => {
    if (!path || path.length === 0) return null

    if (path.length === 1) {
      const [removed] = items.splice(path[0], 1)
      return removed
    }

    let current = items[path[0]]
    for (let i = 1; i < path.length - 1; i++) {
      if (!current?.children) return null
      current = current.children[path[i]]
    }

    if (!current?.children) return null
    const [removed] = current.children.splice(path[path.length - 1], 1)
    return removed
  }

  // Simple drag & drop: just change depth based on horizontal movement
  const handleDropItem = (dropData: { draggedPath: number[], targetPath: number[], newDepth: number }) => {
    const { draggedPath, targetPath, newDepth } = dropData

    // پیدا کردن آیتمی که drag شده
    const draggedItem = getItemByPath(currentMenu.value.items, draggedPath)
    if (!draggedItem) {
      return
    }

    // حذف از جای قبلی
    removeItemByPath(currentMenu.value.items, draggedPath)

    // تنظیم depth جدید با parentId
    // depth 0 = parentId: null
    // depth > 0 = parentId: آیتم بالایی
    if (newDepth === 0) {
      draggedItem.parentId = null
    } else {
      // پیدا کردن parent بر اساس target
      const targetItem = getItemByPath(currentMenu.value.items, targetPath)
      if (targetItem) {
        draggedItem.parentId = targetItem.id || targetItem.clientId
      }
    }

    // اضافه کردن به جای جدید
    currentMenu.value.items.push(draggedItem)

    // Rebuild tree
    const flatItems = flattenMenuItems(currentMenu.value.items, null, currentMenu.value.id)
    currentMenu.value.items = buildMenuTree(flatItems)
  }

  // Initialize content data
  const initializeContent = async () => {
    await Promise.all([
      fetchPages(),
      fetchPosts(),
      fetchCategories(),
      fetchProductCategories()
    ])
  }

  // Refresh content data (for when new content is added)
  const refreshContent = async () => {
    await initializeContent()
  }

  return {
    // State
    menus,
    currentMenu,
    deletedMenuItemIds,
    isLoading,
    isSaving,
    pages,
    posts,
    categories,
    productCategories,
    menuLocations,

    // API Methods
    fetchMenus,
    fetchMenu,
    saveMenu,
    deleteMenu,
    reorderMenuItems,
    assignMenuToLocation,
    unassignMenuFromLocation,
    markMenuItemDeleted,
    registerDeletedMenuBranch,

    // Content API Methods
    fetchPages,
    fetchPosts,
    fetchCategories,
    fetchProductCategories,
    initializeContent,
    refreshContent,

    // Utility Methods
    createNewMenu,
    generateSlug,
    addMenuItem,
    removeMenuItem,
    reorderItems,
    handleDropItem
  }
}
