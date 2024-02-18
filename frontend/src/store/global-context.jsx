import { createContext } from 'react';

/**
 * @description This function creat a global context for verifier .
 *
 * @param None.
 * @returns None.
 */
export const VerifiedContext = createContext({
  verified: false,
  toggle_verified: () => {},
});