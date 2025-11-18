import { ref } from 'vue'

export const useDeleteConfirmModal = () => {
  const deleteModalRef = ref<{
    openDeleteConfirm: (id: number | string) => void
    openBulkDeleteConfirm: () => void
    closeDeleteConfirm: () => void
    closeBulkDeleteConfirm: () => void
  } | null>(null)

  // باز کردن مودال تأیید حذف تک آیتم
  const openDeleteConfirm = (id: number | string) => {
    if (deleteModalRef.value) {
      deleteModalRef.value.openDeleteConfirm(id)
    }
  }

  // باز کردن مودال تأیید حذف گروهی
  const openBulkDeleteConfirm = () => {
    if (deleteModalRef.value) {
      deleteModalRef.value.openBulkDeleteConfirm()
    }
  }

  // بستن مودال تأیید حذف تک آیتم
  const closeDeleteConfirm = () => {
    if (deleteModalRef.value) {
      deleteModalRef.value.closeDeleteConfirm()
    }
  }

  // بستن مودال تأیید حذف گروهی
  const closeBulkDeleteConfirm = () => {
    if (deleteModalRef.value) {
      deleteModalRef.value.closeBulkDeleteConfirm()
    }
  }

  return {
    deleteModalRef,
    openDeleteConfirm,
    openBulkDeleteConfirm,
    closeDeleteConfirm,
    closeBulkDeleteConfirm
  }
}
