<script lang="ts">
  import "./app.css";
  import ToolList from "$lib/components/ToolList.svelte";
  import TabList from "$lib/TabList.svelte";
  import { ScrollArea } from "$lib/components/ui/scroll-area/index.ts";
  import "iconify-icon";
  import * as Avatar from "$lib/components/ui/avatar";
  import * as Card from "$lib/components/ui/card";

  let selectedAreas: string[] = [];

  function handleValueChange(event: CustomEvent<string[]>) {
    console.log("Selected values in App:", event.detail);
    selectedAreas = event.detail;
  }
</script>

<main class="flex h-screen">
  <div class="flex flex-col items-center w-1/2 mx-6 my-6">
    <Card.Root class="border-0">
      <Card.Header>
        <Card.Title class="text-4xl">Julian Derks</Card.Title>
        <Card.Description class="text-base"
          >Data Engineer | DevOps Engineer</Card.Description
        >
      </Card.Header>
      <Card.Content>
        <div class="flex justify-center">
          <Avatar.Root class="size-72 border-white border-2">
            <Avatar.Image src="julian.jpeg" alt="@shadcn" class="grayscale" />
            <Avatar.Fallback>JD</Avatar.Fallback>
          </Avatar.Root>
        </div>
        <div class="my-6 text-start">
          <h1 class="text-3xl text-ring mb-0">Introduction</h1>
          <p class="text-base">
            Former Data Scientist turned Data Engineer, currently exploring the
            world of DevOps with a focus on Kubernetes (GKE), Helm & automated
            testing. My journey has equipped me with expertise in building data
            infrastructures like Data warehouses, Data lakehouses, and Delta
            lakes, utilizing tools such as PySpark, SQL, Databricks, and
            Terraform. I have experience across both Azure and GCP platforms,
            embracing challenges and opportunities to innovate within data and
            infrastructure management.
            <br />
            <br />
            Working as a consultant has exposed me to a variety of customers in different
            sectors, enriching my experience with diverse business needs and technological
            environments. This role has allowed me to not only assist organizations
            in harnessing their data to drive progress but also to continually expand
            my knowledge and skills with new tools and practices.
          </p>
        </div>
      </Card.Content>
      <Card.Footer></Card.Footer>
    </Card.Root>
  </div>
  <div class="flex flex-col w-1/2 mx-6 my-6">
    <div class="w-full">
      <TabList on:valueChange={handleValueChange}></TabList>
    </div>
    <div class="flex-grow border rounded-xl overflow-hidden">
      <ScrollArea class="h-full" scrollbarYClasses="w-0" orientation="vertical">
        <ToolList {selectedAreas}></ToolList>
      </ScrollArea>
    </div>
  </div>
</main>
