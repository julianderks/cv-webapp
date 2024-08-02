<script lang="ts">
	import ToolList from '$lib/components/ToolList.svelte';
	import Introduction from '$lib/components/Introduction.svelte';
</script>

<!-- <div class="w-screen">
	{#each circles as circle}
		<div
			class="relative rounded-full bg-white bg-opacity-50"
			style="width: {circle.r * 2}px; height: {circle.r * 2}px; left: {circle.x *
				100}%; top: {circle.y * 100}%; transform: translate(-50%, -50%);"
		></div>
	{/each}
</div> -->

<div class="relative flex h-full justify-center">
	<div class="grid h-full max-w-7xl grid-cols-1 lg:grid-cols-2">
		<div class="mx-8 pt-8 lg:py-8">
			<Introduction />
		</div>
		<div class="mx-8 flex h-[1000px] flex-col overflow-hidden pb-8 lg:h-full lg:py-8">
			<ToolList />
		</div>
	</div>
</div>
