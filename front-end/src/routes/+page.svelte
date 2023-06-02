<script>
	import { Card } from "flowbite-svelte";
  </script>
  
  <div class="grid grid-cols-6 grid-rows-10 gap-4">
	<div class="col-span-4 row-span-4">
	  <Card class="w-full h-full max-w-none max-h-none">Queue Summary</Card>
	</div>
	<div class="col-span-2 row-span-4">
	  <Card class="w-full h-full max-w-none">Incidents by priority</Card>
	</div>
	<div class="col-span-2 row-span-4">
		<Card class="w-full h-full max-w-none">Incidents by category</Card>
	</div>
	<div class="col-span-2 row-span-4">
		<Card class="w-full h-full max-w-none">Incidents by severity</Card>
	</div>
	<div class="col-span-2 row-span-4">
		<Card class="w-full h-full max-w-none">Resolution Time</Card>
	</div>
	<div class="col-span-2 row-span-4">
		<Card class="w-full h-full max-w-none">SLA</Card>
	</div>
	<div class="col-span-2 row-span-4">
		<Card class="w-full h-full max-w-none">Trends</Card>
	</div>
	<div class="col-span-2 row-span-4">
		<Card class="w-full h-full max-w-none">Incidents by assigned user</Card>
	</div>
	<div class="col-span-2 row-span-4">
		<Card class="w-full h-full max-w-none">Metrics</Card>
	</div>
	<!-- 
		SLA Compliance: If you have Service Level Agreements (SLAs) in place, you'll want an indicator of compliance with these. How many incidents were resolved within the agreed time frame?
		Incident Trends: This can help identify recurring problems or spikes in incidents at particular times, which could indicate larger systemic issues.
		Incident Assignment: This shows who is assigned to what incident, their role, and possibly their contact info. It could also display the incident assignment history.
		Performance Metrics: How is the team doing? This could include metrics like incidents resolved per day, average response time, incidents escalated, etc. 
	-->
  </div>
  