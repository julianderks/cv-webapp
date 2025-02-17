<script lang="ts">
	import { onMount } from 'svelte';
	import ToggleGroup from './ToggleGroup.svelte';
	import { ScrollArea } from '$lib/components/ui/scroll-area';
	import * as Card from '$lib/components/ui/card';
	import * as Dialog from '$lib/components/ui/dialog';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';

	interface Tool {
		name: string;
		iconify_icon_name?: string;
		custom_icon_link?: string;
		experience_level: number;
		description: {
			introduction: string;
			used_where: string;
		};
		area: string;
		docs_link: string;
	}

	let tools: Tool[] = [];
	let filteredTools: Tool[] = [];
	let uniqueAreas: Set<string> = new Set();
	let selectedTool: Tool | null = null;
	let selectedAreas: string[] = [];

	onMount(async () => {
		const response = await fetch(PUBLIC_BACKEND_URL + '/api');
		tools = (await response.json()) as Tool[];
		tools.sort(
			(a, b) =>
				a.area.localeCompare(b.area) ||
				a.name.localeCompare(b.name) ||
				b.experience_level - a.experience_level
		);
		uniqueAreas = new Set(tools.map((tool) => tool.area));
		filterTools();
	});

	function filterTools() {
		filteredTools = selectedAreas.length
			? tools.filter((tool) => selectedAreas.includes(tool.area))
			: tools;
		console.log(`Filtered Tools:`, filteredTools);
	}

	$: selectedAreas, filterTools(); // Ensure this triggers whenever selectedAreas changes

	function handleToolClick(tool: Tool) {
		selectedTool = tool;
	}

	function handleValueChange(event: CustomEvent<string[]>) {
		selectedAreas = event.detail;
	}
</script>

<div class="flex h-full flex-col">
	<ToggleGroup on:valueChange={handleValueChange}></ToggleGroup>
	<div class="flex-grow overflow-auto rounded-3xl border-2 border-white">
		<ScrollArea class="" scrollbarYClasses="w-0" orientation="vertical">
			<div class="grid grid-cols-1 gap-4 p-4 sm:grid-cols-2 lg:grid-cols-3">
				{#each filteredTools as tool}
					<Dialog.Root>
						<Dialog.Trigger>
							<button on:click={() => handleToolClick(tool)} aria-label="Tool {tool.name}">
								<Card.Root class="border-0 text-left shadow-none duration-200 hover:scale-110">
									<Card.Header>
										<Card.Title class="overflow-hidden text-ellipsis whitespace-nowrap"
											>{tool.name}</Card.Title
										>
										<Card.Description class="overflow-hidden text-ellipsis whitespace-nowrap"
											>{tool.area}</Card.Description
										>
									</Card.Header>
									<Card.Content>
										<div class="flex h-16 w-full flex-col items-center">
											<div class="flex flex-grow items-center justify-center">
												{#if tool.iconify_icon_name}
													<iconify-icon icon={tool.iconify_icon_name} height="3.5em"></iconify-icon>
												{:else if tool.custom_icon_link}
													<img
														src={'http://localhost:8000/' + tool.custom_icon_link}
														alt="{tool.name} icon"
													/>
												{/if}
											</div>
										</div>
									</Card.Content>
									<Card.Footer>
										<div class="flex w-full justify-center gap-1">
											{#each Array(5) as _, i}
												<div
													class="h-2 w-4 rounded-full border-2 border-white {i <
													tool.experience_level
														? 'bg-white'
														: ''}"
												></div>
											{/each}
										</div>
									</Card.Footer>
								</Card.Root>
							</button>
						</Dialog.Trigger>
						<Dialog.Content>
							<Dialog.Header>
								<Dialog.Title class="mb-4"
									>{selectedTool ? selectedTool.name : 'Tool Details'}</Dialog.Title
								>
								<Dialog.Description>
									<h1 class="font-bold">Introduction</h1>
									<p class="mb-4">
										{selectedTool ? selectedTool.description.introduction : ''}
									</p>
									<h1 class="font-bold">Personal Experience</h1>
									<p class="mb-4">
										{selectedTool ? selectedTool.description.used_where : ''}
									</p>
									<p class="mb-4">
										Link to the
										<a
											href={selectedTool ? selectedTool.docs_link : '#'}
											class="text-blue-500 underline"
											target="_blank"
										>
											documentation
										</a>
									</p>
								</Dialog.Description>
							</Dialog.Header>
						</Dialog.Content>
					</Dialog.Root>
				{/each}
			</div>
		</ScrollArea>
	</div>
</div>
