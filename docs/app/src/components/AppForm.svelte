<script lang="ts">
import { onMount } from 'svelte'

// Props
let userClass: string | undefined = undefined
let style: string | undefined = undefined
let locale: string | undefined = undefined
export {
	userClass as class,
	style,
	locale,
}

import { i18nFactory } from '~/i18n'
const _ = i18nFactory(locale as any)

// Generate a random suffix for id attributes
const idSuffix = Math.random().toString(36).substring(2)

const placeholderProxyUrl: string = 'http://localhost'
const placeholderServiceUrl: string = 'http://api.example.com'
const placeholderServicePath: string = '/api/v1'

let proxyPort: string = '80'
let proxyUrl: string = ''
let serviceUrl: string = ''
let servicePath: string = ''

let convertError: string | null = null

let lastCopiedSelector: string | null = null

function parsePort(port: string): number
{
	const parsedPort = port ? Number(port) : 80
	if (!Number.isInteger(parsedPort))
	{
		return 80
	}
	return parsedPort
}

function defaultProxyUrl(port: string): string
{
	const parsedPort = parsePort(port)
	if (parsedPort === 80)
	{
		return placeholderProxyUrl
	}
	return `${placeholderProxyUrl}:${parsedPort}`
}

function outputUrl(
	port: string,
	proxyUrl: string,
	serviceUrl: string,
	servicePath: string,
): string
{
	let newProxyUrl = proxyUrl.trim()
	if (!newProxyUrl)
	{
		newProxyUrl = defaultProxyUrl(proxyPort)
	}

	let newServiceUrl = serviceUrl.trim()
	if (newServiceUrl)
	{
		newServiceUrl = '/' + encodeURIComponent(newServiceUrl)
	}

	let newServicePath = newServiceUrl ? servicePath : ''
	if (newServicePath && !newServicePath.startsWith('/'))
	{
		newServicePath = '/' + newServicePath
	}

	try
	{
		convertError = null
		return `${newProxyUrl}${newServiceUrl}${newServicePath}`
	}
	catch (error: any)
	{
		convertError = error.message
		return ''
	}
}

function copyInputValue(selector: string, event: MouseEvent)
{
	const target = document.getElementById(`input-${selector}-${idSuffix}`) as HTMLInputElement | null
	if (target)
	{
		lastCopiedSelector = selector
		navigator.clipboard.writeText(target.value)
	}
}
</script>

<div
	class={[
		userClass,
	].join(' ')}
	style={style}
>

	<!-- Setup group -->
	<div class="p-4 space-y-4 col-span-2 sm:col-auto border border-gray-200 rounded-md">

		<label class="flex flex-col gap-2">
			<span class="text-gray-700">
				{_('Start the CORS Proxy server locally.')}
				<button
					type="button" class="copy-button" class:copied={lastCopiedSelector === 'proxy'}
					on:click|preventDefault={copyInputValue.bind(null, 'proxy')}
				>
					<span class="icon icon-[mdi--content-copy] icon-align"></span>
					{lastCopiedSelector === 'proxy' ? _('Copied!') : _('Copy')}
				</button>
			</span>
			<div class="flex gap-2 items-center">
				<span class="text-gray-700">
					{_('Run on port:')}
				</span>
				<input
					class="form-textarea bg-gray-100 block w-16 h-8 p-2 rounded-md flex-1 resize-none outline-gray-500"
					bind:value={proxyPort}
				/>
			</div>
			<div class="h-8 sm:h-12">
				<input
					id={`input-proxy-${idSuffix}`}
					class="form-textarea bg-gray-100 block w-full h-full p-2 rounded-md flex-1 resize-none outline-gray-500"
					value={`docker run -p ${parsePort(proxyPort)}:8080 matiboux/cors-proxy`}
					disabled
				/>
			</div>
		</label>

	</div>

	<!-- Step separator -->
	<div class="pl-2 text-xl text-gray-600">
		<span class="icon-[mdi--plus] align-icon-inline"></span>
	</div>

	<!-- Input group -->
	<div class="p-4 space-y-4 col-span-2 sm:col-auto border border-gray-200 rounded-md">

		<label class="flex flex-col gap-2">
			<span class="text-gray-700">
				{_('CORS Proxy server URL (leave empty for default):')}
			</span>
			<div class="h-8 sm:h-12">
				<input
					class="form-textarea bg-gray-100 block w-full h-full p-2 placeholder:text-gray-600 rounded-md flex-1 resize-none outline-gray-500"
					placeholder={defaultProxyUrl(proxyPort)}
					bind:value={proxyUrl}
				/>
			</div>
		</label>

		<label class="flex flex-col gap-2">
			<span class="text-gray-700">
				{_('Target service URL you want to access via your proxy:')}
			</span>
			<div class="h-8 sm:h-12">
				<input
					class="form-textarea bg-gray-100 block w-full h-full p-2 placeholder:text-gray-400 rounded-md flex-1 resize-none outline-gray-500"
					placeholder={placeholderServiceUrl}
					bind:value={serviceUrl}
				/>
			</div>
		</label>

		<label class="flex flex-col gap-2">
			<span class="text-gray-700">
				{_('Target service path of your request (optional):')}
			</span>
			<div class="h-8 sm:h-12">
				<input
					class="form-textarea bg-gray-100 block w-full h-full p-2 placeholder:text-gray-400 rounded-md flex-1 resize-none outline-gray-500"
					placeholder={placeholderServicePath}
					bind:value={servicePath}
				/>
			</div>
		</label>
	</div>

	<!-- Convertion separator -->
	<div class="pl-2 text-xl text-gray-600">
		<span class="icon-[mdi--arrow-down] align-icon-inline"></span>
	</div>

	<!-- Output group -->
	<div class="output-group bg-gray-100 p-4 space-y-4 col-span-2 sm:col-auto border border-gray-300 rounded-md">
		<label class="flex flex-col gap-2">
			<span class="text-gray-700">
				{_('Formatted proxy URL pointing to the target service:')}
				<button
					type="button" class="copy-button" class:copied={lastCopiedSelector === 'output'}
					on:click|preventDefault={copyInputValue.bind(null, 'output')}
				>
					<span class="icon icon-[mdi--content-copy] icon-align"></span>
					{lastCopiedSelector === 'output' ? _('Copied!') : _('Copy')}
				</button>
			</span>
			<div class="h-8 sm:h-12">
				<input
					id={`input-output-${idSuffix}`}
					class="form-textarea bg-gray-200 block w-full h-full p-2 placeholder:text-gray-400 rounded-md flex-1 resize-none outline-gray-500"
					placeholder="Enter the proxy and service URL"
					value={outputUrl(proxyPort, proxyUrl, serviceUrl, servicePath)}
					disabled
				/>
			</div>
		</label>
	</div>

</div>

<style lang="scss">
.copy-button {
	@apply inline-flex items-center gap-1;
	@apply ml-1 px-2 py-0.5 rounded-full;
	@apply text-sm font-normal;
	@apply text-gray-600 active:text-gray-700;
	@apply bg-gray-100 active:bg-gray-200;

	&.copied {
		@apply bg-green-100 active:bg-green-200;
	}
}

.output-group .copy-button {
	@apply bg-gray-200 active:bg-gray-300;

	&.copied {
		@apply bg-green-200 active:bg-green-300;
	}
}
</style>
