<script lang="ts">
	import Inner from '$lib/components/Inner.svelte';

	let m = { x: 0, y: 0 };

	function handleMove(event: MouseEvent) {
		m.x = event.clientX;
		m.y = event.clientY;
	}

	let helloText: string;
	function handeMessage(event) {
		helloText = event.detail.text;
	}
</script>

<div class="grid grid-cols-4 gap-8 p-8">
	<div class="h-96 border-2" on:pointermove={handleMove}>
		<div class="flex flex-col"></div>
		<h1 class="mt-8 basis-1/5 text-center underline">DOM EVENT</h1>
		<p class="m-8 basis-4/5">
			Events are fired to notify code of "interesting changes" that may affect code execution. These
			can arise from user interactions such as using a mouse or resizing a window, changes in the
			state of the underlying environment (e.g. low battery or media events from the operating
			system), and other causes. <br />
			<br />
			The pointer is at {m.x} x {m.y}
		</p>
	</div>
	<div
		class="h-96 border-2"
		on:pointermove={(e) => {
			m = { x: e.clientX, y: e.clientY };
		}}
	>
		<div class="flex flex-col"></div>
		<h1 class="mt-8 basis-1/5 text-center underline">INLINE HANDLER</h1>
		<p class="m-8 basis-4/5">
			You can also declare event handlers inline:
			<br />
			<br />
			The pointer is at {m.x} x {m.y}
		</p>
	</div>
	<div class="h-96 border-2">
		<div class="flex flex-col"></div>
		<h1 class="mt-8 basis-1/5 text-center underline">COMPONENT EVENT</h1>
		<p class="m-8 basis-4/5">
			<Inner on:message={handeMessage}></Inner>
		</p>
		<h1>{helloText ? helloText : ''}</h1>
	</div>
	<div class="h-96 border-2">
		<div class="flex flex-col"></div>
		<h1 class="mt-8 basis-1/5 text-center underline">EVENT FORWARDING</h1>
		<p class="m-8 basis-4/5">
			<Inner on:message={handeMessage}></Inner>
		</p>
		<h1>{helloText ? helloText : ''}</h1>
	</div>
	<div class="h-96 border-2">1</div>
	<div class="h-96 border-2">1</div>
</div>
