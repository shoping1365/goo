export const faType = (v?: string): string => {
  if (!v) return ''
  switch (v.trim()) {
    case 'select':
      return 'انتخاب'
    case 'custom_text':
      return 'متن سفارشی'
    // Already Persian
    case 'انتخاب':
    case 'متن سفارشی':
      return v.trim()
    default:
      return v.trim()
  }
}

export const faControl = (v?: string): string => {
  if (!v) return ''
  const val = v.trim()
  switch (val) {
    case 'dropdown_single':
      return 'لیست باز شونده تک انتخابی'
    case 'dropdown_multi':
      return 'لیست باز شونده چند انتخابی'
    case 'textbox_single':
      return 'تکست باکس متنی تک خطی'
    case 'textbox_multiline':
      return 'تکست باکس متنی چند خطی'
    // Already Persian labels
    case 'لیست باز شونده تک انتخابی':
    case 'لیست باز شونده چند انتخابی':
    case 'تکست باکس متنی تک خطی':
    case 'تکست باکس متنی چند خطی':
      return val
    default:
      return val
  }
}

export const slugType = (v: string): string => {
  const val = v.trim()
  switch (val) {
    case 'انتخاب':
      return 'select'
    case 'متن سفارشی':
      return 'custom_text'
    default:
      return val
  }
}

export const slugControl = (v: string): string => {
  const val = v.trim()
  switch (val) {
    case 'لیست باز شونده تک انتخابی':
      return 'dropdown_single'
    case 'لیست باز شونده چند انتخابی':
      return 'dropdown_multi'
    case 'تکست باکس متنی تک خطی':
      return 'textbox_single'
    case 'تکست باکس متنی چند خطی':
      return 'textbox_multiline'
    default:
      return val
  }
} 