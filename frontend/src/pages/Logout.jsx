import { redirect } from 'react-router-dom';
import { eraseAuthToken } from '../util/auth';

/**
 * @description This route action is called when Logout is submitted.
 *
 * @param None.
 * @returns {redirect} redirect to root path.
 */
export function action() {
  //localStorage.removeItem('access_token');
  eraseAuthToken('access_token')
  //localStorage.removeItem('refresh_token');
  eraseAuthToken('refresh_token')
  //localStorage.removeItem('expiration');
  eraseAuthToken('expiration')
  return redirect('/');
}