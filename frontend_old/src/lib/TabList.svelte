<script lang="ts">
  import * as HoverCard from "$lib/components/ui/hover-card";
  import * as ToggleGroup from "$lib/components/ui/toggle-group";
  import { createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();

  // Define a function to handle value changes
  function handleValueChange(values: string[] | undefined) {
    if (values) {
      console.log("Selected values:", values);
      dispatch("valueChange", values); // Dispatch the event with the selected values
    } else {
      console.log("No values selected");
    }
  }
</script>

<HoverCard.Root openDelay={200}>
  <HoverCard.Trigger asChild let:builder>
    <span
      class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
      use:builder.action
      {...builder}
    >
      <!-- Mode -->
    </span>
  </HoverCard.Trigger>
  <HoverCard.Content class="w-[320px] text-sm" side="left">
    Choose the interface that best suits your task. You can provide: a simple
    prompt to complete, starting and ending text to insert a completion within,
    or some text with instructions to edit it.
  </HoverCard.Content>
</HoverCard.Root>
<ToggleGroup.Root
  type="multiple"
  class="bg-opacity-10 border-transparent mb-1 ml-2 flex justify-start"
  onValueChange={handleValueChange}
>
  <ToggleGroup.Item value="DevOps">DevOps</ToggleGroup.Item>
  <ToggleGroup.Item value="Backend">Backend</ToggleGroup.Item>
  <ToggleGroup.Item value="Frontend">Frontend</ToggleGroup.Item>
  <ToggleGroup.Item value="Data">Data</ToggleGroup.Item>
</ToggleGroup.Root>
