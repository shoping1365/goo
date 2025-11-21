import { describe, expect, it, vi } from 'vitest'
import useNewProductForm from '../../app/composables/useNewProductForm'

describe('useNewProductForm Composable', () => {
  // Since state is shared (singleton), we need to be careful about test isolation.
  // In a real app, we might want to reset the state, but the composable doesn't expose a reset method.
  // We will test the behavior as is.

  it('should initialize with default values', () => {
    const { productForm, sections, tinyApiKey } = useNewProductForm()
    
    expect(productForm.status).toBe('active')
    expect(sections.mainInfo).toBe(true)
    expect(tinyApiKey).toBeDefined()
  })

  it('should toggle sections correctly', () => {
    const { sections, toggleSection } = useNewProductForm()
    
    // Initial state
    const initialState = sections.mainInfo
    
    // Toggle
    toggleSection('mainInfo')
    expect(sections.mainInfo).toBe(!initialState)
    
    // Toggle back
    toggleSection('mainInfo')
    expect(sections.mainInfo).toBe(initialState)
  })

  it('should update product form data', () => {
    const { productForm } = useNewProductForm()
    
    productForm.name = 'New Product'
    productForm.description = 'Desc'
    
    expect(productForm.name).toBe('New Product')
    expect(productForm.description).toBe('Desc')
  })

  it('should have generateAIContent method', () => {
    const { generateAIContent } = useNewProductForm()
    
    // Mock alert
    const alertMock = vi.spyOn(window, 'alert').mockImplementation(() => {})
    
    generateAIContent('short')
    expect(alertMock).toHaveBeenCalled()
    
    alertMock.mockRestore()
  })
})
