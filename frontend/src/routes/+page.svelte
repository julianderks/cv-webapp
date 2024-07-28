<script lang="ts">
	import ToolList from '$lib/components/ToolList.svelte';
	import Introduction from '$lib/components/Introduction.svelte';
</script>

<div class="flex flex-col lg:h-full lg:w-full lg:flex-row">
	<div class="flex flex-col items-center lg:w-1/2">
		<div class="mx-6 my-6">
			<Introduction></Introduction>
		</div>
	</div>
	<div class="flex flex-col lg:h-full lg:w-1/2">
		<div class="mx-6 my-6 rounded-xl lg:h-full">
			<ToolList></ToolList>
		</div>
	</div>
</div>
