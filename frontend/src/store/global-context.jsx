import { createContext } from 'react';

export const VerifiedContext = createContext({
  verified: false,
  toggle_verified: () => {},
});