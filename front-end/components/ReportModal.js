import * as React from 'react';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import Modal from '@mui/material/Modal';
import TextField from '@mui/material/TextField';
import Grid from '@mui/material/Grid';
import { Select, MenuItem } from '@mui/material';
import { FormControl } from '@mui/material';
import PocketBase from 'pocketbase';


const pb = new PocketBase('http://127.0.0.1:8090')


const style = {
  position: 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 800,
  bgcolor: 'background.paper',
  border: '2px solid #000',
  boxShadow: 24,
  p: 4,
};

export default function ReportModal({ open, onClose }) {
  const handleClose = () => {
    onClose();
  };
  const [priorities, setPriorities] = React.useState([]);
  const [systems, setSystems] = React.useState([]);

  React.useEffect(() => {
    const fetchPriorities = async () => {
      const fetchedPriorities = await pb.collection('priorities').getFullList();
      setPriorities(fetchedPriorities);
      console.log(fetchedPriorities); // Log the fetched data
    };
    const fetchSystems = async () => {
      const fetchedSystems = await pb.collection('systems_and_services').getFullList();
      setSystems(fetchedSystems);
      console.log(fetchedSystems); // Log the fetched data
    };
    fetchSystems();
    fetchPriorities();
  }, []);
  

  return (
    <Modal
      open={open}
      onClose={handleClose}
      aria-labelledby="modal-modal-title"
      aria-describedby="modal-modal-description"
    >
      <Box sx={style}>
        <Typography id="modal-modal-title" variant="h5" component="h1">
          Report a new incident
        </Typography>
        <Box
          component="form"
          sx={{
            '& .MuiTextField-root': { m: 1, width: '100%' },
            mt:4
          }}
          noValidate
          autoComplete="off"
        >
          <Grid container spacing={2}>
            <Grid item xs={12}>
              <TextField 
              required 
              id="outlined-required" 
              label="Title" 
              fullWidth
              helperText="Please provide an accurate title for your incident"
              />
            </Grid>
            <Grid item xs={6}>
              <TextField
                id="outlined-select-currency"
                select
                label="Priority"
                helperText="Please select your priority"
              >
                {priorities.map((option) => (
                  <MenuItem key={option.name} value={option.name}>
                    {option.name}
                  </MenuItem>
                ))}
              </TextField>
            </Grid>
            <Grid item xs={6}>
            <TextField
                id="outlined-select-currency"
                select
                label="Affected System/Service"
                helperText="Please select the system/service that is affected"
              >
                {systems.map((option) => (
                  <MenuItem key={option.name} value={option.name}>
                    {option.name}
                  </MenuItem>
                ))}
              </TextField>
            </Grid>
            <Grid item xs={12}>
              <TextField 
              required 
              id="outlined-required" 
              label="Description"  
              fullWidth
              multiline
              rows={4}
              />
            </Grid>
            <Grid item xs={12}>
              <TextField id="outlined-required" label="Attachments" defaultValue="Hello World" fullWidth />
            </Grid>
          </Grid>
          <Box
          sx={{
            display: 'flex',
            justifyContent: 'flex-end',
            mt: 2,
          }}
        >
          <Button variant="contained" color="primary">
            Submit
          </Button>
        </Box>
      </Box>
      </Box>
    </Modal>
  );
}
