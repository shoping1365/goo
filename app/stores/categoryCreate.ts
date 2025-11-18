import { defineStore } from 'pinia'

export const useCategoryCreateStore = defineStore('categoryCreate', {
  state: () => ({
    images: [],
    sections: { images: true }
  }),
  actions: {
    addImage(img) { this.images.push(img) },
    removeImage(id) { this.images = this.images.filter(i => i.id !== id) },
    toggleSection(section) { this.sections[section] = !this.sections[section] }
  }
}) 