import { Box } from '@mui/material';
import PasswordForm from './components/PasswordForm';

function App() {
  return (
    <Box
      display="flex"
      justifyContent="center"
      alignItems="center"
      minHeight="100vh"
      bgcolor="#f0f0f0"
    >
      <PasswordForm />
    </Box>
  );
}

export default App;
