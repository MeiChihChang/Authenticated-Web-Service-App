import {Dropdown, DropdownToggle, DropdownMenu, DropdownItem, Row, Col} from 'reactstrap';
import {useState} from 'react';

function Organization({organization_list, selecthandler}) {
  const [dropdownOpen, setDropdownOpen] = useState(false);
  const toggle = () => setDropdownOpen((prevState) => !prevState);

  return (
    <>
    <Row className="justify-content-center"> 
      <Col>
      <h4>Organizations:</h4> 
      </Col>
      <Col>
   
      <Dropdown isOpen={dropdownOpen} toggle={toggle}>
        <DropdownToggle className="bg-primary" caret>Choose</DropdownToggle>
        <DropdownMenu>
        {organization_list.map((organization) => (<DropdownItem key={organization.id} onClick={selecthandler} dropDownValue={organization.name}>{organization.name}</DropdownItem >))}
        </DropdownMenu>
      </Dropdown>
      </Col>
    </Row>
    </>
  );
}



export default Organization;