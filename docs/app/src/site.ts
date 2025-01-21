import type { I18nKeys } from '~/i18n/type.d.ts'

export interface Site
{
	title?: string
	description?: string | I18nKeys
	author?: string
	keywords?: string[]
	themeColor?: string
	favicon?: string
	lang?: string
}

export const site: Site = {
	title: 'CORS Proxy',
	description: 'Simple CORS proxy server to bypass browser restrictions locally.',
	author: 'Matiboux',
	themeColor: '#ffffff',
}
