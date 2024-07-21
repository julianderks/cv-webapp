<script lang="ts">
  import { onMount } from "svelte";
  import * as Popover from "$lib/components/ui/popover";

  interface Tool {
    tool: string;
    icon_path: string;
    experience_level: number; // Assume this is a number from 0 to 3
  }

  let tools: Tool[] = [];

  onMount(async () => {
    const response = await fetch("http://localhost:8000/api");
    tools = await response.json();
    // Sort tools by experience_level in descending order
    tools.sort((a, b) => b.experience_level - a.experience_level);
    console.log(tools);
  });
</script>

<div class="flex flex-wrap justify-center gap-8">
  {#each tools as tool}
    <Popover.Root>
      <Popover.Trigger>
        <div class="flex flex-col items-center hover:scale-110 duration-200">
          <img class="w-32 h-32" src={tool.icon_path} alt={tool.tool} />
          <div class="flex justify-center mt-2 gap-1">
            {#each Array(5) as _, i}
              <div
                class="w-4 h-2 rounded-full my-2 border-2 border-white {i <
                tool.experience_level
                  ? 'bg-white'
                  : ''}"
              ></div>
            {/each}
          </div>
        </div>
      </Popover.Trigger>
      <Popover.Content class="border-2">
        Place content for the popover here.
      </Popover.Content>
    </Popover.Root>
  {/each}
</div>
