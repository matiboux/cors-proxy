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

const placeholderProxyUrl: string = 'http://localhost'
const placeholderServiceUrl: string = 'http://api.example.com'
const placeholderServicePath: string = '/api/v1'

let proxyUrl: string = placeholderProxyUrl
let serviceUrl: string = ''
let servicePath: string = ''

let outputUrl: string = ''
let outputValueElement = null
let convertError: string | null = null

let onInputConvertTimeout: NodeJS.Timeout | undefined = undefined

function convert()
{
	if (onInputConvertTimeout)
	{
		clearTimeout(onInputConvertTimeout)
	}

	if (!proxyUrl)
	{
		outputUrl = ''
		convertError = null
		return
	}

	let newProxyUrl = proxyUrl.trim()

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
		outputUrl = `${newProxyUrl}${newServiceUrl}${newServicePath}`
		convertError = null
	}
	catch (error: any)
	{
		outputUrl = ''
		convertError = error.message
	}
}

onMount(() =>
	{
		convert()
	})

let allowDefaultInputValue: boolean = true

function onInput()
{
	if (onInputConvertTimeout)
	{
		clearTimeout(onInputConvertTimeout)
	}

	allowDefaultInputValue = false

	onInputConvertTimeout = setTimeout(() =>
		{
			convert()
		}, 400)
}

function onChange()
{
	allowDefaultInputValue = false

	convert()
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

		<label class="block space-y-2 flex flex-col">
			<span class="text-gray-700">
				{_('Start the CORS Proxy server locally.')}
			</span>
			<div class="flex gap-2 items-center">
				<span class="text-gray-700">
					{_('Run on port:')}
				</span>
				<input
					class="form-textarea bg-gray-100 block w-16 h-8 p-2 rounded-md flex-1 resize-none outline-gray-500"
					value="80"
				/>
			</div>
			<div class="h-8 sm:h-12">
				<input
					class="form-textarea bg-gray-100 block w-full h-full p-2 rounded-md flex-1 resize-none outline-gray-500"
					value="docker run -p 80:8080 ghcr.io/matiboux/cors-proxy"
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

		<label class="block space-y-2 flex flex-col">
			<span class="text-gray-700">
				{_('CORS Proxy server URL:')}
			</span>
			<div class="h-8 sm:h-12">
				<input
					class="form-textarea bg-gray-100 block w-full h-full p-2 rounded-md flex-1 resize-none outline-gray-500"
					placeholder={placeholderProxyUrl}
					bind:value={proxyUrl}
					on:input|preventDefault={onInput}
					on:change|preventDefault={onChange}
				/>
			</div>
		</label>

		<label class="block space-y-2 flex flex-col">
			<span class="text-gray-700">
				{_('Target service URL you want to access via your proxy:')}
			</span>
			<div class="h-8 sm:h-12">
				<input
					class="form-textarea bg-gray-100 block w-full h-full p-2 rounded-md flex-1 resize-none outline-gray-500"
					placeholder={placeholderServiceUrl}
					bind:value={serviceUrl}
					on:input|preventDefault={onInput}
					on:change|preventDefault={onChange}
				/>
			</div>
		</label>

		<label class="block space-y-2 flex flex-col">
			<span class="text-gray-700">
				{_('Target service path of your request:')}
			</span>
			<div class="h-8 sm:h-12">
				<input
					class="form-textarea bg-gray-100 block w-full h-full p-2 rounded-md flex-1 resize-none outline-gray-500"
					placeholder={placeholderServicePath}
					bind:value={servicePath}
					on:input|preventDefault={onInput}
					on:change|preventDefault={onChange}
				/>
			</div>
		</label>
	</div>

	<!-- Convertion separator -->
	<div class="pl-2 text-xl text-gray-600">
		<span class="icon-[mdi--arrow-down] align-icon-inline"></span>
	</div>

	<!-- Output group -->
	<div class="bg-gray-100 p-4 space-y-4 col-span-2 sm:col-auto border border-gray-300 rounded-md">
		<label class="block space-y-2 flex flex-col">
			<span class="text-gray-700">
				{_('Formatted proxy URL pointing to the target service:')}
			</span>
			<div class="h-8 sm:h-12">
				<input
					class="form-textarea bg-gray-200 block w-full h-full p-2 rounded-md flex-1 resize-none"
					placeholder="Enter the proxy and service URL"
					bind:this={outputValueElement}
					bind:value={outputUrl}
					disabled
				/>
			</div>
		</label>
	</div>

</div>
