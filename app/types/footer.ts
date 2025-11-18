export interface FooterLayer {
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
}

export interface Footer {
     id?: number
     name: string
     description: string
     pageSelection: string
     specificPages: string
     excludedPages: string
     layers: FooterLayer[]
     is_active?: boolean
     createdAt?: string
     updatedAt?: string
}
