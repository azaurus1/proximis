<script>
  import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from 'flowbite-svelte';

  /**
	 * @type {any}
	 */
   export let incidents;
   
</script>

<h1 class="pb-4">Queue Summary</h1>

<Table hoverable={true}>
  <TableHead>
    <TableHeadCell>Title</TableHeadCell>
    <TableHeadCell>Status</TableHeadCell>
    <TableHeadCell>Category</TableHeadCell>
    <TableHeadCell>Priority</TableHeadCell>
    <TableHeadCell>Assignee</TableHeadCell>
  </TableHead>
  <TableBody>
    {#await incidents}
      Loading...
    {:then data}
      {#each data.items as item}
        <TableBodyRow>
          <TableBodyCell>{item.title}</TableBodyCell>
          <TableBodyCell>{item.status}</TableBodyCell>
          <TableBodyCell>{item.expand.categorisation.title}</TableBodyCell>
          <TableBodyCell>{item.priority}</TableBodyCell>
          <TableBodyCell>{item.expand.assignee.username}</TableBodyCell>
        </TableBodyRow>
      {/each}
    {/await}
  </TableBody>
</Table>
