import { readFile } from 'fs/promises'
import { createError, defineEventHandler, getQuery } from 'h3'
import { join } from 'path'

interface SearchLog {
    id: string
    query: string
    timestamp: string
    userId?: string
    resultsCount: number
    clickedResults: number
    searchDuration: number
    filters?: Record<string, unknown>
    userAgent?: string
    ip?: string
}

interface SearchStats {
    totalSearches: number
    uniqueUsers: number
    avgResultsCount: number
    avgSearchDuration: number
    topQueries: Array<{ query: string; count: number }>
    searchesByDay: Array<{ date: string; count: number }>
    noResultQueries: Array<{ query: string; count: number }>
    popularFilters: Record<string, number>
}

export default defineEventHandler(async (event) => {
    // Check admin authorization
    const user = event.context.user
    if (!user || user.role !== 'admin') {
        throw createError({
            statusCode: 403,
            statusMessage: 'Unauthorized - Admin access required'
        })
    }

    const query = getQuery(event)
    const fromDate = query.from ? new Date(query.from as string) : new Date(Date.now() - 7 * 24 * 60 * 60 * 1000)
    const toDate = query.to ? new Date(query.to as string) : new Date()

    try {
        // Read search logs from file or database
        // For now, we'll return mock data structure
        const logs: SearchLog[] = []

        // In production, you would read from a database or log file
        // Example: Reading from NDJSON file
        try {
            const dataPath = join(process.cwd(), 'data', 'search', 'search_logs.ndjson')
            const fileContent = await readFile(dataPath, 'utf-8')
            const lines = fileContent.trim().split('\n')

            for (const line of lines) {
                try {
                    const log = JSON.parse(line)
                    const logDate = new Date(log.timestamp)

                    // Filter by date range
                    if (logDate >= fromDate && logDate <= toDate) {
                        logs.push(log)
                    }
                } catch (parseError) {
                    console.error('Error parsing log line:', parseError)
                }
            }
        } catch {
            console.warn('Search logs file not found, returning empty data')
        }

        // Calculate statistics
        const stats: SearchStats = {
            totalSearches: logs.length,
            uniqueUsers: new Set(logs.filter(l => l.userId).map(l => l.userId)).size,
            avgResultsCount: logs.length > 0
                ? logs.reduce((sum, log) => sum + log.resultsCount, 0) / logs.length
                : 0,
            avgSearchDuration: logs.length > 0
                ? logs.reduce((sum, log) => sum + log.searchDuration, 0) / logs.length
                : 0,
            topQueries: [],
            searchesByDay: [],
            noResultQueries: [],
            popularFilters: {}
        }

        // Calculate top queries
        const queryCount: Record<string, number> = {}
        logs.forEach(log => {
            queryCount[log.query] = (queryCount[log.query] || 0) + 1
        })

        stats.topQueries = Object.entries(queryCount)
            .map(([query, count]) => ({ query, count }))
            .sort((a, b) => b.count - a.count)
            .slice(0, 20)

        // Calculate no-result queries
        const noResultCount: Record<string, number> = {}
        logs.filter(log => log.resultsCount === 0).forEach(log => {
            noResultCount[log.query] = (noResultCount[log.query] || 0) + 1
        })

        stats.noResultQueries = Object.entries(noResultCount)
            .map(([query, count]) => ({ query, count }))
            .sort((a, b) => b.count - a.count)
            .slice(0, 20)

        // Calculate searches by day
        const searchesByDay: Record<string, number> = {}
        logs.forEach(log => {
            const date = new Date(log.timestamp).toISOString().split('T')[0]
            searchesByDay[date] = (searchesByDay[date] || 0) + 1
        })

        stats.searchesByDay = Object.entries(searchesByDay)
            .map(([date, count]) => ({ date, count }))
            .sort((a, b) => a.date.localeCompare(b.date))

        // Calculate popular filters
        const filterCount: Record<string, number> = {}
        logs.forEach(log => {
            if (log.filters) {
                Object.keys(log.filters).forEach(filterKey => {
                    filterCount[filterKey] = (filterCount[filterKey] || 0) + 1
                })
            }
        })

        stats.popularFilters = filterCount

        return {
            logs: logs.slice(0, 1000), // Limit to 1000 logs for performance
            stats
        }
    } catch (error) {
        console.error('Error fetching search analytics:', error)
        throw createError({
            statusCode: 500,
            statusMessage: 'Failed to fetch search analytics'
        })
    }
})
