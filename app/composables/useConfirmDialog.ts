import { inject } from 'vue'

interface ConfirmDialogOptions {
  title?: string
  message: string
  confirmText?: string
  cancelText?: string
  type?: 'danger' | 'warning' | 'info'
}

export const useConfirmDialog = () => {
  // استفاده از inject در setup function
  const confirmDialogRef = inject<{ show: (options: ConfirmDialogOptions) => Promise<boolean> } | null>('confirmDialogRef', null)

  const confirm = async (options: ConfirmDialogOptions | string): Promise<boolean> => {
    if (!confirmDialogRef) {
      console.warn('ConfirmDialog component not found. Make sure to add it to your template.')
      return false
    }

    const config = typeof options === 'string'
      ? { message: options }
      : options

    try {
      return await confirmDialogRef.show(config)
    } catch (error) {
      console.error('Error showing confirm dialog:', error)
      return false
    }
  }

  const confirmDelete = async (itemName?: string): Promise<boolean> => {
    const message = itemName
      ? `آیا از حذف "${itemName}" اطمینان دارید؟`
      : 'آیا از حذف این مورد اطمینان دارید؟'

    return confirm({
      title: 'تأیید حذف',
      message,
      confirmText: 'حذف',
      cancelText: 'انصراف',
      type: 'danger'
    })
  }

  const confirmAction = async (actionName: string, itemName?: string): Promise<boolean> => {
    const message = itemName
      ? `آیا از ${actionName} "${itemName}" اطمینان دارید؟`
      : `آیا از ${actionName} این مورد اطمینان دارید؟`

    return confirm({
      title: 'تأیید عملیات',
      message,
      confirmText: 'تأیید',
      cancelText: 'انصراف',
      type: 'warning'
    })
  }

  return {
    confirmDialogRef,
    confirm,
    confirmDelete,
    confirmAction
  }
}
