import type { Diff } from '~/i18n/types.d.ts'

import type { DefaultLocaleKeys } from './types.d.ts'

const locale = {
	// Index
	'Simple CORS proxy server to bypass browser restrictions locally.':
		'Serveur proxy CORS simple pour contourner les restrictions du navigateur localement.',
	'This tool allows you to format the proxy URL pointing to the target service you want to access without CORS restrictions.':
		'Cet outil vous permet de formater l\'URL du proxy pointant vers le service cible que vous souhaitez accéder sans restrictions CORS.',
	// AppForm
	'Start the CORS Proxy server locally.':
		'Démarrez le serveur proxy CORS localement.',
	'Run on port:': 'Exécuter sur le port :',
	'CORS Proxy server URL (leave empty for default):':
		'URL du serveur CORS Proxy (laissez vide pour la valeur par défaut) :',
	'Target service URL you want to access via your proxy:':
		'URL du service cible que vous souhaitez accéder via votre proxy :',
	'Target service path of your request (optional):':
		'Chemin du service cible de votre demande (facultatif) :',
	'Formatted proxy URL pointing to the target service:':
		'URL du proxy formatée pointant vers le service cible :',
	'Copy': 'Copier',
	'Copied!': 'Copié !',
	// Footer
	'Open source project': 'Projet open source',
	'See the source code on {0}': 'Voir le code source sur {0}',
	'Built with {0}, served by {1}': 'Construit avec {0}, servi par {1}',
	'Made with love by {0}': 'Créé avec amour par {0}',
	'Data privacy': 'Confidentialité des données',
	'No data is collected or processed over the network or on any server.':
		'Aucune donnée n\'est collectée ou traitée sur le réseau ou sur un serveur.',
	'All data is processed locally in your browser, and stays on your own device.':
		'Toutes les données sont traitées localement dans votre navigateur et restent sur votre propre appareil.',
	'This website uses no cookies and does no tracking.':
		'Ce site web n\'utilise pas de cookies et ne fait pas de suivi.',
} as const

export default locale satisfies
	// Static type check for missing keys
	Readonly<Record<Diff<DefaultLocaleKeys, keyof typeof locale>, string>> &
	// Static type check for extra keys
	Readonly<Record<Diff<keyof typeof locale, DefaultLocaleKeys>, never>>
