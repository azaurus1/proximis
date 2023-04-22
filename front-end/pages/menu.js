import * as React from 'react';
import Button from '@mui/material/Button';
import ButtonGroup from '@mui/material/ButtonGroup';
import Box from '@mui/material/Box';
import Link from '../src/Link';
import Layout from "./Layout";
import ReportModal from '../components/ReportModal';

export default function Menu() {
  const [modalOpen, setModalOpen] = React.useState(false);


  console.log(modalOpen)

  function handleModalOpen() {
    console.log("Modal open")
    setModalOpen(true);
  }

  function handleModalClose() {
    console.log("Modal closed")
    setModalOpen(false);
  }

  const buttons = [
    <ReportModal open={modalOpen} onClose={handleModalClose} key="basic-modal" />,
    <Button onClick={handleModalOpen} key="report" sx={{ my: 1 }}>
      Report an Incident
    </Button>,
    <Button component={Link} href="my_incidents" key="my_incidents" sx={{ my: 1 }}>
      My Incidents
    </Button>,
    <Button component={Link} href="profile" key="profile" sx={{ my: 1 }}>
      Profile
    </Button>,
    <Button component={Link} href="help" key="help" sx={{ my: 1 }}>
      Help
    </Button>,
    <Button key="logout" sx={{ my: 1 }}>
      Logout
    </Button>
  ];

  return (
    <Layout>
      <Box
        sx={{
          height: '33vh',
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center'
        }}
      >
        <Box
          sx={{
            width: '33%',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center'
          }}
        >
          <ButtonGroup
            sx={{
              width: '100%'
            }}
            orientation="vertical"
            aria-label="vertical contained button group"
            variant="contained"
          >
            {buttons}
          </ButtonGroup>
        </Box>
      </Box>
    </Layout>
  );
}
