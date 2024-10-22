import type { I18n, Diff } from '../type'

const locale =
{
	// Index
	'Simple CORS proxy server to bypass browser restrictions locally.':
		'Serveur proxy CORS simple pour contourner les restrictions du navigateur localement.',
	'This tool allows you to format the proxy URL pointing to the target service you want to access without CORS restrictions.':
		'Cet outil vous permet de formater l\'URL du proxy pointant vers le service cible que vous souhaitez accéder sans restrictions CORS.',
	// AppForm
	'CORS Proxy server URL:':
		'URL du serveur CORS Proxy :',
	'Target service URL you want to access via your proxy:':
		'URL du service cible que vous souhaitez accéder via votre proxy :',
	'Target service path of your request:':
		'Chemin du service cible de votre demande :',
	'Formatted proxy URL pointing to the target service:':
		'URL du proxy formatée pointant vers le service cible :',
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

// Static type check
export default locale satisfies
	// Check for missing keys:
	Readonly<Record<Diff<keyof I18n, keyof typeof locale>, string>> &
	// Check for extra keys:
	Readonly<Record<Diff<keyof typeof locale, keyof I18n>, never>>
