const defaultLocale =
[
	// Index
	'Simple CORS proxy server to bypass browser restrictions locally.',
	'This tool allows you to format the proxy URL pointing to the target service you want to access without CORS restrictions.',
	// AppForm
	'Start the CORS Proxy server locally.',
	'Run on port:',
	'CORS Proxy server URL (leave empty for default):',
	'Target service URL you want to access via your proxy:',
	'Target service path of your request (optional):',
	'Formatted proxy URL pointing to the target service:',
	'Copy',
	'Copied!',
	// Footer
	'Open source project',
	'See the source code on {0}',
	'Built with {0}, served by {1}',
	'Made with love by {0}',
	'Data privacy',
	'No data is collected or processed over the network or on any server.',
	'All data is processed locally in your browser, and stays on your own device.',
	'This website uses no cookies and does no tracking.',
] as const

type Keys = typeof defaultLocale[number]
type Type = { [key in Keys]: key }

// Default locale uses the key as the value
const locale = defaultLocale
	.reduce<Type>((acc, key) =>
		{
			(acc as any)[key] = key
			return acc
		},
		{} as Type,
	)

export default locale as Readonly<typeof locale>
