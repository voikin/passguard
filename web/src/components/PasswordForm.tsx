import React, { useState, useEffect } from 'react';
import { useQuery } from '@tanstack/react-query';
import { getPasswordScore } from '../api/password';
import {
  Box,
  Button,
  CircularProgress,
  FormControl,
  InputLabel,
  Input,
  Typography,
  Alert,
  Slider,
  Card,
  CardContent,
  CardActions,
} from '@mui/material';

const PasswordForm: React.FC = () => {
  const [password, setPassword] = useState('');
  const [score, setScore] = useState<number | null>(null);

  const { data, error, isLoading, refetch } = useQuery({
    queryKey: [password],
    queryFn: () => getPasswordScore(password),
    enabled: false,
  });

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    refetch();
  };

  useEffect(() => {
    if (data && data.score !== undefined) {
      setScore(data.score);
    }
  }, [data]);

  return (
    <Card sx={{ minWidth: 300, maxWidth: 600, mx: 'auto', mt: 5, borderRadius: 4, p: 4 }} raised>
      <CardContent>
        <Typography align='center' variant="h5" component="div" gutterBottom>
          Passguard
        </Typography>
        <form onSubmit={handleSubmit}>
          <FormControl fullWidth margin="normal">
            <InputLabel htmlFor="password">Password</InputLabel>
            <Input
              id="password"
              type="text"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </FormControl>
          <CardActions>
            <Button variant="contained" color="primary" type="submit" fullWidth>
              Check Password
            </Button>
          </CardActions>
        </form>
        {isLoading && (
          <Box mt={2} textAlign="center">
            <CircularProgress />
          </Box>
        )}
        {error && (
          <Box mt={2}>
            <Alert severity="error">{(error as Error).message}</Alert>
          </Box>
        )}
        {score !== null && (
          <Box mt={2} p={2} border={1} borderRadius={4}>
            <Typography>Score: {score}</Typography>
            <Slider
              value={score}
              min={0}
              max={50}
              valueLabelDisplay="auto"
              aria-labelledby="password-score-slider"
              disableSwap
            />
          </Box>
        )}
      </CardContent>
    </Card>
  );
};

export default PasswordForm;
