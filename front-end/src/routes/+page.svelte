<script>
	import { Card } from "flowbite-svelte";
	import IncidentsByCategory from "../IncidentsByCategory.svelte";
	import IncidentsByPriority from "../IncidentsByPriority.svelte";
	import IncidentsBySeverity from "../IncidentsBySeverity.svelte";
	import IncidentsByUser from "../IncidentsByUser.svelte";
	import Metrics from "../Metrics.svelte";
	import QueueSummary from "../QueueSummary.svelte";
	import ResolutionTime from "../ResolutionTime.svelte";
	import Sla from "../SLA.svelte";
	import Trends from "../Trends.svelte";

	const page = writable(1)
	const perPage = writable(10)

	import { pbStore } from 'svelte-pocketbase';
	import { writable } from "svelte/store";
	const records = $pbStore.collection('incidents').getList($page,$perPage,{"expand":"categorisation,assignee"})

  </script>
  
  <div class="grid grid-cols-6 grid-rows-10 gap-4 pt-6">
	<div class="col-span-4 row-span-3">
	  <Card class="w-full h-full max-w-none max-h-none" size="lg">
		  <QueueSummary incidents={records}/>
	  </Card>
	</div>
	<div class="col-span-2 row-span-3">
	  <Card class="w-full h-full max-w-none" size="lg">
		  <IncidentsByPriority records={records}/>
	  </Card>
	</div>
	<div class="col-span-2 row-span-2">
		<Card class="w-full h-full max-w-none" size="lg">
			<IncidentsByCategory />
		</Card>
	</div>
	<div class="col-span-2 row-span-2">
		<Card class="w-full h-full max-w-none" size="lg">
			<IncidentsBySeverity />
		</Card>
	</div>
	<div class="col-span-2 row-span-1">
		<Card class="w-full h-full max-w-none" size="lg">
			<ResolutionTime />
		</Card>
	</div>
	<div class="col-span-2 row-span-1">
		<Card class="w-full h-full max-w-none" size="lg">
			<Sla />
		</Card>
	</div>
	<div class="col-span-2 row-span-4">
		<Card class="w-full h-full max-w-none" size="lg">
			<Trends />
		</Card>
	</div>
	<div class="col-span-2 row-span-4">
		<Card class="w-full h-full max-w-none" size="lg">
			<IncidentsByUser />
		</Card>
	</div>
	<div class="col-span-2 row-span-4" >
		<Card class="w-full h-full max-w-none" size="lg">
			<Metrics />
		</Card>
	</div>
  </div>
  