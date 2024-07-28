<script lang="ts">
	import ToolList from '$lib/components/ToolList.svelte';
	import Introduction from '$lib/components/Introduction.svelte';
</script>

<div class="flex h-full justify-center">
	<div class="grid h-full max-w-7xl grid-cols-1 lg:grid-cols-2">
		<div class="mx-8 pt-8 lg:py-8">
			<Introduction />
		</div>
		<div class="mx-8 flex h-[1000px] flex-col overflow-hidden pb-8 lg:h-full lg:py-8">
			<ToolList />
		</div>
	</div>
</div>
