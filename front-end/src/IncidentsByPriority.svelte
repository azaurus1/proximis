<script>
    import PriorityChart from "./components/PriorityChart.svelte";
    export let records;

    let priorityCount = {}

    function countPriorities(item){
        if (item.priority in priorityCount){
            priorityCount[item.priority] += 1;
        }else{
            priorityCount[item.priority] = 1;
        }
    }

    function buildChartData(records_data){

        records_data.items.forEach(countPriorities);
        const data = {
            labels: ['Critical', 'High', 'Medium', 'Low'],
            datasets: [
            {
                data: [priorityCount['critical'], priorityCount['high'], priorityCount['medium'], priorityCount['low']],
                backgroundColor: [
                '#F7464A',
                '#FF781f',
                '#FDB45C',
                '#949FB1',
                ],
                hoverBackgroundColor: [
                '#FF5A5E',
                '#FFA500',
                '#FFC870',
                '#A8B3C5',
                ],
            },
            ],
        };
        return data;
    }
</script>

<h1 class="pb-4">Incidents by priority</h1>
<div>
    {#await records}
        Loading.. 
    {:then records_data}
        <PriorityChart data={buildChartData(records_data)}/>
    {/await}
</div>