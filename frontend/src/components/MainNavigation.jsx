import {Button, NavbarBrand, Navbar, Nav,NavItem, NavLink, Container, Row, Col} from 'reactstrap';
import {Form, useRouteLoaderData} from 'react-router-dom';

function MainNavigation() {
  const token = useRouteLoaderData('root');

  return (
    <header >
    <Navbar fixed="top" color="light" light expand="xs" className="border-bottom border-gray bg-white" style={{ height: 80 }}>
      <Container>
        <Row g-0 className="position-relative w-100 align-items-center">
          <Col className="d-none d-lg-flex justify-content-start">
            <Nav className="mrx-auto" navbar>  
              <NavbarBrand href="/">Home</NavbarBrand>
              <NavItem className="d-flex align-items-center">
              {!token && (
                <NavLink className="font-weight-bold" href="/datalist" disabled>OpenData.Swiss</NavLink>
              )}
              {token && (
                <NavLink className="font-weight-bold" href="/datalist">OpenData.Swiss</NavLink>
              )}
              </NavItem> 
            </Nav>
          </Col>
          <Col className="d-none d-lg-flex justify-content-end">
          {!token && (
            <Form action="/login" inline>
              <Button color="info" outline size="sm">Login</Button>
            </Form>
          )}
          {token && (
            <Form action="/logout" method="post" inline>
              <Button color="info" outline size="sm">Logout</Button>
            </Form>
          )}
          </Col>
        </Row>
      </Container>
    </Navbar>
    </header>
  );
}

export default MainNavigation;
