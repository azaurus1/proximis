import * as React from 'react';
import { DataGrid } from '@mui/x-data-grid';
import Layout from './Layout';
import PocketBase from 'pocketbase';

const pb = new PocketBase('http://127.0.0.1:8090')

const columns = [
  { field: 'title', headerName: 'Title', width: 130 },
  {
    field: 'assignee',
    headerName: 'Assignee',
    width: 130,
    valueGetter: (params) => params.row.expand?.assignee?.name || '',
  },
  {
    field: 'priority',
    headerName: 'Priority',
    width: 130,
    valueGetter: (params) => params.row.expand?.priority?.name || '',
  },
  {
    field: 'category',
    headerName: 'Category',
    width: 130,
    valueGetter: (params) => params.row.expand?.category?.name || '',
  },
  {
    field: 'status',
    headerName: 'Status',
    width: 130,
    valueGetter: (params) => params.row.expand?.status?.name || '',
  },
  { field: 'created', headerName: 'Created Date', width: 200 },
];

export default function MyIncidents() {
  const [incidents, setIncidents] = React.useState([]);

  React.useEffect(() => {
    const fetchIncidents = async () => {
      const fetchedIncidents = await pb.collection('incidents').getFullList({
        expand: 'category,priority,status,assignee,affected_system_or_service',
      });
      setIncidents(fetchedIncidents);
      console.log(fetchedIncidents); // Log the fetched data
    };

    fetchIncidents();
  }, []);
  
    return (
    <Layout> 
        <div style={{ height: 400, width: '100%' }}>
        <DataGrid
            rows={incidents}
            columns={columns}
            pageSize={5}
            rowsPerPageOptions={[5]}
            checkboxSelection
        />
        </div>
    </Layout>
  );
    }