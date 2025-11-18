import type { Component } from 'vue'

type AsyncComponent = () => Promise<Component>

type FooterWidgetRegistry = Record<string, AsyncComponent>

export const footerWidgetRegistry: FooterWidgetRegistry = {
  logo: () => import('./FooterWidgetLogo.vue'),
  image: () => import('./FooterWidgetImage.vue'),
  custom: () => import('./FooterWidgetCustom.vue'),
  menu: () => import('./FooterWidgetMenu.vue'),
  language: () => import('./FooterWidgetLanguage.vue'),
  social: () => import('./FooterWidgetSocial.vue'),
  contact: () => import('./FooterWidgetContact.vue'),
  about: () => import('./FooterWidgetAbout.vue'),
  'working-hours': () => import('./FooterWidgetWorkingHours.vue'),
  newsletter: () => import('./FooterWidgetNewsletter.vue'),
  copyright: () => import('./FooterWidgetCopyright.vue'),
  links: () => import('./FooterWidgetLinks.vue'),
  trust: () => import('./FooterWidgetTrust.vue'),
  'back-to-top': () => import('./FooterWidgetBackToTop.vue')
}
