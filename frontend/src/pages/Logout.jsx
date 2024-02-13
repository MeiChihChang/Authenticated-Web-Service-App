import { redirect } from 'react-router-dom';
import { eraseAuthToken } from '../util/auth';


export function action() {
  //localStorage.removeItem('access_token');
  eraseAuthToken('access_token')
  //localStorage.removeItem('refresh_token');
  eraseAuthToken('refresh_token')
  //localStorage.removeItem('expiration');
  eraseAuthToken('expiration')
  return redirect('/');
}