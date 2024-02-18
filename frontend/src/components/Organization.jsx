import {Dropdown, DropdownToggle, DropdownMenu, DropdownItem, Row, Col} from 'reactstrap';
import {useState} from 'react';

/**
 * @description This component renders a combox list.
 *
 * @param  {[Organization]} name organization_list list of organization name.
 * @param  {function} name selecthandler a handle to set the selected organization name.
 * @returns {Organization} A React element that renders a combox list with organization names.
 */
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