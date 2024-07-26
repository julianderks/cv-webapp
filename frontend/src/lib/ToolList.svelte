<script lang="ts">
  import { onMount } from "svelte";
  import * as Card from "$lib/components/ui/card";
  import * as Dialog from "$lib/components/ui/dialog";

  interface Tool {
    name: string;
    iconify_icon_name?: string;
    custom_icon_link?: string;
    experience_level: number;
    iconify_icon_height_factor: number;
    description: {
      introduction: string;
      used_where: string;
    };
    area: string;
    docs_link: string;
  }

  export let selectedAreas: string[] = [];
  let tools: Tool[] = [];
  let filteredTools: Tool[] = [];
  let uniqueAreas: Set<string> = new Set();
  let selectedTool: Tool | null = null;

  onMount(async () => {
    const response = await fetch("http://localhost:8000/api");
    tools = await response.json();
    tools.sort((a, b) => {
      if (a.area < b.area) return -1;
      if (a.area > b.area) return 1;
      if (a.name < b.name) return -1;
      if (a.name > b.name) return 1;
      return b.experience_level - a.experience_level;
    });
    uniqueAreas = new Set(tools.map((tool) => tool.area));
    filterTools();
  });

  function filterTools() {
    if (selectedAreas.length === 0) {
      filteredTools = tools;
    } else {
      filteredTools = tools.filter((tool) => selectedAreas.includes(tool.area));
    }
    console.log(`Filtered Tools:`, filteredTools);
  }

  $: selectedAreas, filterTools();

  function handleToolClick(tool: Tool) {
    selectedTool = tool;
  }
</script>

<div class="flex flex-wrap justify-center my-6">
  {#each filteredTools as tool}
    <Dialog.Root>
      <Dialog.Trigger>
        <button
          class="hover:scale-110 duration-200 m-1"
          on:click={() => handleToolClick(tool)}
          aria-label="Tool {tool.name}"
        >
          <Card.Root class="border-0 text-left">
            <Card.Header>
              <Card.Title class="">{tool.name}</Card.Title>
              <Card.Description>{tool.area}</Card.Description>
            </Card.Header>
            <Card.Content>
              <div class="flex flex-col items-center w-22 h-16">
                <div class="flex-grow flex items-center justify-center">
                  {#if tool.iconify_icon_name}
                    <iconify-icon icon={tool.iconify_icon_name} height="3.5em"
                    ></iconify-icon>
                  {:else if tool.custom_icon_link}
                    <img
                      src={"http://localhost:8000/" + tool.custom_icon_link}
                      alt="{tool.name} icon"
                    />
                  {/if}
                </div>
              </div>
            </Card.Content>
            <Card.Footer>
              <div class="flex justify-center gap-1 w-full">
                {#each Array(5) as _, i}
                  <div
                    class="w-4 h-2 rounded-full border-2 border-white {i <
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
          <Dialog.Title class="mb-4">
            {selectedTool ? selectedTool.name : "Tool Details"}
          </Dialog.Title>
          <Dialog.Description>
            <h1 class="font-bold">Introduction</h1>
            <p class=" mb-4">
              {selectedTool ? selectedTool.description.introduction : ""}
            </p>
            <h1 class="font-bold">Personal Experience</h1>
            <p class="mb-4">
              {selectedTool ? selectedTool.description.used_where : ""}
            </p>
            <p class="mb-4">
              Link to the
              <a
                href={tool.docs_link}
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
