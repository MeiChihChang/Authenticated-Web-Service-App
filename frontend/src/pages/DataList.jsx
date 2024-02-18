import { Suspense, useState,  useEffect, useRef, useContext} from 'react';
import {
  useRouteLoaderData,
  json,
  defer,
  Await,
} from 'react-router-dom';
import { Container, Row, Col } from "reactstrap"

import Organization from '../components/Organization';
import DataList from '../components/List';
import VerifyModel from '../components/VerifyModel';
import { getAuthToken} from '../util/auth';
import {VerifiedContext} from '../store/global-context';

/**
 * @description This component renders a DataListPage to the user.
 *
 * @param None.
 * @returns {DataListPage} A React element that renders a DataListPage including a combox list, a verifier model dialog and a table datalist.
 */
function DataListPage() {
  const { organizations } = useRouteLoaderData('datalist');

  const [ selectedOrganization, setSelectedOrganization] = useState("None")
  const [loadeddatalist, setLoadeddatalist] = useState([]);

  const {verified, toggle_verified} = useContext(VerifiedContext);
  const [test, setTest] = useState("");
  const [errorText, setErrorText] = useState("");
  const answear = useRef(0);

  const onDropdownItem_Click = (sender) => {   
      const dropDownValue = sender.currentTarget.getAttribute("dropDownValue");
      console.log("dropdownvalue %s", dropDownValue)
      setSelectedOrganization(dropDownValue);
  }

  const onModelSubmit_Click = (event) => {   
    event.preventDefault();
 
    const reply = event.target.answear.value; 
    console.log("isVerified " + verified);  
    if (parseInt(reply) === answear.current) {
      if (!verified) {
        toggle_verified(); 
      }     
    } else { setErrorText("The answear is incorrect. Please try again"); }
    console.log("isVerified " + verified); 
  }

  useEffect(() => {
    function generateTest() {
      const max = 100
      const min = 0
      const number1 = Math.floor(Math.random() * (max - min + 1) + min);
      const number2 = Math.floor(Math.random() * (max - min + 1) + min);
      setTest(`${number1} + ${number2} = '`);
      answear.current = number1 + number2;
    };
    async function fetchdataList() {
      
      try {
        const datalist = await loadDataList(selectedOrganization);
        setLoadeddatalist(datalist);
      } catch (error) {
        //setError({ message: error.message || 'Failed to fetch datalist.' });
      }

    };

    generateTest();
    fetchdataList();
  }, [selectedOrganization, setLoadeddatalist]);

  return (
    <>
    <Container className="mt-5 container-xs"> 
      <Row className="justify-content-start"> 
      <Col>
      <Suspense fallback={<p style={{ textAlign: 'center' }}>Loading...</p>}>
        <Await resolve={organizations}>
          {(loadedorganizations) => <Organization organization_list={loadedorganizations} selecthandler={onDropdownItem_Click}/>}
        </Await>
      </Suspense>
      <VerifyModel isopen={!verified} test={test} handler={onModelSubmit_Click} error={errorText}/>
      {loadeddatalist !== null && <DataList datalist={loadeddatalist}/>}
      </Col>
    </Row>
    </Container>
    </>
  );
}

export default DataListPage;

async function loadOrganizations() {
  const token = getAuthToken();
  //console.log("access_token:" + token)
  const response = await fetch(`${process.env.REACT_APP_BACKEND}/swissdata/organizations`, {
    method: "GET",
    headers: {
      'Authorization': 'Bearer ' + token
    }
  });

  if (!response.ok) {
    throw json(
      { message: 'Could not fetch organization list.' },
      {
        status: 500,
      }
    );
  } else {
    const resData = await response.json();
    console.log(resData.length)
    return resData;
  }
}

async function loadDataList(name) {
  if (name !== "" && name !== null) {
    const token = getAuthToken();
    const response = await fetch(`${process.env.REACT_APP_BACKEND}/swissdata/datalist/` + name , {
      method: "GET",
      headers: {
        'Authorization': 'Bearer ' + token
      }
   });

    if (!response.ok) {
      throw json(
        { message: 'Could not fetch data list.' },
        {
          status: 500,
        }
      );
    } else {
      const resData = await response.json();
      //console.log(resData);
      return resData;
    }
  }
  return null;
}

export async function loader({ request, params }) {
  return defer({
    organizations: await loadOrganizations()
  });
}