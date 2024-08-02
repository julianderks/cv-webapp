<script lang="ts">
	import * as Avatar from '$lib/components/ui/avatar';
	import * as Card from '$lib/components/ui/card';
	import avatar from '$lib/assets/julian.jpeg';
</script>

<Card.Root class="border-0 shadow-none">
	<Card.Header class="mb-2">
		<Card.Title class="text-4xl ">Julian Derks</Card.Title>
		<Card.Description class="text-base ">Data Engineer | DevOps Engineer</Card.Description>
	</Card.Header>
	<Card.Content>
		<div class="mb-8 flex justify-center">
			<Avatar.Root class="size-72 border-4 border-white">
				<Avatar.Image src={avatar} alt="@shadcn" class="grayscale" />
				<Avatar.Fallback>JD</Avatar.Fallback>
			</Avatar.Root>
		</div>
		<div class="text-start">
			<h1 class="text-ring mb-0 text-2xl underline">Introduction</h1>
			<p class="break-all text-base">
				Former Data Scientist turned Data Engineer, currently exploring the world of DevOps with a
				focus on Kubernetes (GKE), Helm & Automated testing. My journey has equipped me with
				expertise in building data infrastructures like Data warehouses, Data lakehouses, and Delta
				lakes, utilizing tools such as PySpark, SQL, Databricks, and Terraform. I have experience
				across both Azure and GCP platforms, embracing challenges and opportunities to innovate
				within data and infrastructure management.
				<br />
				<br />
				Working as a consultant has exposed me to a variety of customers in different sectors, enriching
				my experience with diverse business needs and technological environments. This role has allowed
				me to not only assist organizations in harnessing their data to drive progress but also to continually
				expand my knowledge and skills with new tools and practices.
				<br />
				<br />
				I am passionate about learning new tech‐related things and enjoy working on hobby projects in
				my spare time. In fact this very website is a result of a hobby project where i implemented a
				three‐tier app with Golang as the backend and Sveltekit/TypeScript/Tailwind CSS as the frontend!
			</p>
		</div>
	</Card.Content>
	<Card.Footer></Card.Footer>
</Card.Root>
