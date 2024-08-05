<script lang="ts">
	import * as ToggleGroup from '$lib/components/ui/toggle-group';
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	// Define a function to handle value changes
	function handleValueChange(values: string[] | undefined) {
		if (values) {
			console.log('Selected values:', values);
			dispatch('valueChange', values); // Dispatch the event with the selected values
		} else {
			console.log('No values selected');
		}
	}
</script>

<ToggleGroup.Root
	type="multiple"
	class="mb-1 ml-2 flex justify-start border-transparent bg-opacity-10"
	size="sm"
	onValueChange={handleValueChange}
>
	<ToggleGroup.Item class="text-lg" value="DevOps">DevOps</ToggleGroup.Item>
	<ToggleGroup.Item class="text-lg" value="Backend">Backend</ToggleGroup.Item>
	<ToggleGroup.Item class="text-lg" value="Frontend">Frontend</ToggleGroup.Item>
	<ToggleGroup.Item class="text-lg" value="Data">Data</ToggleGroup.Item>
</ToggleGroup.Root>
