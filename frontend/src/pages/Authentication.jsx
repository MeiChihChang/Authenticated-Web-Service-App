import AuthForm from '../components/AuthForm';
import { json, redirect} from 'react-router-dom';
import { setAuthToken } from '../util/auth';

function AuthenticationPage() {
  return <AuthForm />;
}

export default AuthenticationPage;

export async function action({ request }) {
    const data = await request.formData();
    const authData = {
        username: data.get('username'),
        password: data.get('password'),
    }
  
    let url = `${process.env.REACT_APP_BACKEND}/login`;
    
    const response = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(authData)
    });
  
    if (response.status === 422 || response.status === 401)   {
        return response;
    }
  
    if (!response.ok) {
        throw json({message: 'Could not authenticate user'}, {status: 500});
    }

    const resData = await response.json();
    const access_token = resData.access_token;
    const refresh_token = resData.refresh_token;
    let i = parseInt(resData.expires_in);
    i = i / 60;
    setAuthToken('access_token', access_token, i)
    setAuthToken('refresh_token', refresh_token, i)
    const expiration = new Date();  
    expiration.setMinutes(expiration.getMinutes() + i);
    setAuthToken('expiration', expiration.toISOString(), i)
  
    return redirect('/');
  }