import './App.css';
import {useState} from 'react';

import HomePage from './pages/HomePage';
import RootLayout from './pages/RootLayout';
import AuthenticationPage, {action as authAction} from './pages/Authentication';
import {action as logoutAction} from './pages/Logout';
import DataListPage, {loader as datalistLoader} from './pages/DataList';
import {tokenLoader} from './util/auth';

import {VerifiedContext} from './store/global-context';

import {
  createBrowserRouter,
  RouterProvider,
} from 'react-router-dom';

// Router Setup with path
const router = createBrowserRouter([
  {
    path: '/',
    element: <RootLayout />,
    id: 'root',
    loader: tokenLoader,
    children: [
      { index: true, element: <HomePage /> },
      { path: 'login', element: <AuthenticationPage />, action: authAction},
      { path: 'logout', action: logoutAction},
      { path: 'datalist', id: 'datalist', element: <DataListPage />, loader: datalistLoader},
    ]},
  
]);

function App() {
  const [isVerified, setIsVerified] = useState(false);
  const toggleVerified = () => {
    setIsVerified((prev) => (prev === false ? true : false));
  };
  const ctxValue = {
    verified: isVerified,
    toggle_verified: toggleVerified
  };
  return (<VerifiedContext.Provider value={ctxValue}><RouterProvider router={router} /></VerifiedContext.Provider>) ;
}

export default App;
